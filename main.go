package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/powerunit-io/platform/logging"
)

func main() {

	// Setting up loger ...
	logger := logging.New()

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "Mon Jan _2 15:04:05 2006",
		FullTimestamp:   true,
		DisableSorting:  true,
	})

	logger.SetOutput(os.Stderr)
	logger.SetLevel("PU_BRIDGE_LOG_LEVEL")

	logger.Info("About to start configuring PowerUnit Bridge...")

}
