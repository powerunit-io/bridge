package main

import (
	"os"

	"github.com/powerunit-io/platform/config"
	devices "github.com/powerunit-io/platform/devices/manager"
	helpers "github.com/powerunit-io/platform/helpers/manager"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/workers/manager"
)

func main() {
	logger := logging.New()
	logger.SetOutput(os.Stderr)
	logger.SetLevel("PU_BRIDGE_LOG_LEVEL")

	logger.Info("About to start configuring Bridge...")

	cnfservice, err := config.NewConfigManager(ServiceName, ServiceConfig)

	if err != nil {
		logger.Error("Received error while building configuration (error: %s)", err)
		os.Exit(2)
	}

	db, err := NewDb(DbAccessName, Config_MySqlPrimary, logger)

	if err != nil {
		logger.Error("Failed to create database connection (err: %s)", err)
		os.Exit(2)
	}

	wmanager := manager.NewWorkerManager(logger)
	hmanager := helpers.NewManager(logger)
	dmanager := devices.NewDeviceManager(logger, db)

	service := NewBridgeService(logger, cnfservice, wmanager, hmanager, dmanager)

	// Ensure that we can boost our thing as maximum possible.
	service.SetGoMaxProcs("PU_BRIDGE_GO_MAX_PROCS")

	if err := service.Bind(db); err != nil {
		logger.Error("Could not create new bind for (connection: %s) (error: %s)", db.Name(), err)
		os.Exit(2)
	}

	mpw, err := NewMqttWorker(
		PrimaryMessageBrokerName, service, logger, Config_MqttPrimaryWorker,
	)

	if err != nil {
		logger.Error("Could not make new primary mqtt worker due to (error: %s)", err)
		os.Exit(2)
	}

	if err := wmanager.AttachWorker(mpw.WorkerName(), mpw); err != nil {
		logger.Error("Could not attach (worker: %s) due to (err: %s)", mpw.WorkerName(), err)
		os.Exit(2)
	}

	if err := service.Start(); err != nil {
		logger.Error("Could not start Bridge due to (err: %s)", err)
		os.Exit(2)
	}

}
