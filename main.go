package main

import (
	"os"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/workers/manager"
)

func main() {
	logger := logging.New()
	logger.SetOutput(os.Stderr)
	logger.SetLevel("PU_BRIDGE_LOG_LEVEL")

	logger.Info("About to start configuring Bridge...")

	config, err := config.NewConfigManager(ServiceName, ServiceConfig)

	if err != nil {
		logger.Error("Received error while building configuration (error: %s)", err)
		os.Exit(2)
	}

	wmanager := manager.NewWorkerManager(logger)

	service := NewService(logger, config, wmanager)

	// Ensure that we can boost our thing as maximum possible.
	service.SetGoMaxProcs("PU_BRIDGE_GO_MAX_PROCS")

	mpw, err := NewMqttWorker("Primary Message Bridge", logger, Config_MqttPrimaryWorker)

	if err != nil {
		logger.Error("Could not make new primary mqtt worker due to (error: %s)", err)
		os.Exit(2)
	}

	//msw, err := NewMqttWorker("Secondary Message Bridge", logger, Config_MqttSecondaryWorker)

	if err != nil {
		logger.Error("Could not make new secondary mqtt worker due to (error: %s)", err)
		os.Exit(2)
	}

	if err := wmanager.AttachWorker(mpw.String(), mpw); err != nil {
		logger.Error("Could not attach (worker: %s) due to (err: %s)", mpw.String(), err)
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
