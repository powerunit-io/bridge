package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/utils"
)

func main() {
	logger := setupLogger()

	logger.Info("About to start configuring Bridge...")

	logger.Info("Setting up configuration ...")

	config := config.NewConfigManager(ServiceName, map[string]interface{}{
		"ServiceName":        ServiceName,
		"ServiceDescription": ServiceDescription,
		"ServiceVersion":     ServiceVersion,
	})

	service := NewService(logger, config)

	service.SetGoMaxProcs("PU_BRIDGE_GO_MAX_PROCS")

	dw := NewDevicesWorker(&service, map[string]interface{}{
		"mqtt_consumers": []string{
			"", "",
		},
	})

	logger.Debug("%v", dw)
	//service.AddWorker()

	if err := service.Start(); err != nil {
		logger.Error("Could not start Bridge due to (err: %s)", err)
		utils.ShutdownSignal()
	}

}

func setupLogger() *logging.Logger {
	logger := logging.New()

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "Mon Jan _2 15:04:05 2006",
		FullTimestamp:   true,
		DisableSorting:  true,
	})

	logger.SetOutput(os.Stderr)
	logger.SetLevel("PU_BRIDGE_LOG_LEVEL")

	return logger
}
