package main

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/events"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/service"
	"github.com/powerunit-io/platform/workers/worker"
)

// DevicesWorker -
type DevicesWorker struct {
	*worker.BaseWorker
}

// Handle -
func (dw *DevicesWorker) Handle(e *events.Event) error {
	return nil
}

// String - We use it to name worker when needed.
func (dw *DevicesWorker) String() string {
	return "Devices Worker"
}

// NewDevicesWorker -
func NewDevicesWorker(serv service.Service, log *logging.Logger, cnf map[string]interface{}) worker.Worker {
	var w *DevicesWorker

	w = &DevicesWorker{
		BaseWorker: &worker.BaseWorker{
			Config: config.NewConfigManager(w.String(), cnf),
			Logger: log,
		},
	}

	return worker.NewWorker(w)
}
