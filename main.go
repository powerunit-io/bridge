package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/utils"
	"github.com/powerunit-io/platform/workers/manager"
)

func main() {
	logger := logging.New()

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "Mon Jan _2 15:04:05 2006",
		FullTimestamp:   true,
		DisableSorting:  true,
	})

	logger.SetOutput(os.Stderr)
	logger.SetLevel("PU_BRIDGE_LOG_LEVEL")

	logger.Info("About to start configuring Bridge...")
	logger.Info("Setting up configuration ...")

	config := config.NewConfigManager(ServiceName, ServiceConfig)
	wmanager := manager.NewWorkerManager(logger)

	service := NewService(logger, config, wmanager)

	service.SetGoMaxProcs("PU_BRIDGE_GO_MAX_PROCS")

	dw := NewDevicesWorker(service, logger, DevicesWorkerConfig)

	if err := wmanager.AttachWorker(dw.String(), dw); err != nil {
		logger.Error("Could not attach (worker: %s) due to (err: %s)", dw.String(), err)
		utils.ShutdownSignal()
	}

	if err := service.Start(); err != nil {
		logger.Error("Could not start Bridge due to (err: %s)", err)
		utils.ShutdownSignal()
	}

}
