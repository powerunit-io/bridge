package main

import (
	"fmt"
	"os"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/connections/mysql"
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

	wmanager := manager.NewWorkerManager(logger)
	hmanager := helpers.NewManager(logger)

	service := NewBridgeService(logger, cnfservice, wmanager, hmanager)

	// Ensure that we can boost our thing as maximum possible.
	service.SetGoMaxProcs("PU_BRIDGE_GO_MAX_PROCS")

	// @TODO - One day we should get this connection stuff as helper or something ...
	cnfmysql, err := config.NewConfigManager(fmt.Sprintf("%s-MYSQL", ServiceName), Config_MySqlPrimary)

	if err != nil {
		logger.Error("Failed to build mysql config (error: %s)", err)
		os.Exit(2)
	}

	mysqlconn, err := mysql.NewConnection(logger, cnfmysql)

	if err != nil {
		logger.Error("Failed to create new connection (error: %s)", err)
		os.Exit(2)
	}

	mpw, err := NewMqttWorker(
		PrimaryMessageBrokerName, service, logger,
		Config_MqttPrimaryWorker, mysqlconn,
	)

	if err != nil {
		logger.Error("Could not make new primary mqtt worker due to (error: %s)", err)
		os.Exit(2)
	}

	//msw, err := NewMqttWorker("Secondary Message Bridge", logger, Config_MqttSecondaryWorker)

	if err != nil {
		logger.Error("Could not make new secondary mqtt worker due to (error: %s)", err)
		os.Exit(2)
	}

	if err := wmanager.AttachWorker(mpw.WorkerName(), mpw); err != nil {
		logger.Error("Could not attach (worker: %s) due to (err: %s)", mpw.WorkerName(), err)
		os.Exit(2)
	}

	/**
	if err := wmanager.AttachWorker(msw.String(), msw); err != nil {
		logger.Error("Could not attach (worker: %s) due to (err: %s)", msw.String(), err)
		os.Exit(2)
	}
	**/

	if err := service.Start(); err != nil {
		logger.Error("Could not start Bridge due to (err: %s)", err)
		os.Exit(2)
	}

}
