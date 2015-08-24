package main

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/service"
	"github.com/powerunit-io/platform/workers"
)

//
type DevicesWorker struct {
	*workers.BaseWorker
}

// String -
func (dw *DevicesWorker) String() string {
	return "DevicesWorker"
}

// NewDevicesWorker -
func NewDevicesWorker(serv *service.Service, c map[string]interface{}) *workers.Worker {
	var worker *DevicesWorker

	worker = &DevicesWorker{
		BaseWorker: &workers.BaseWorker{
			Service: serv,
			Config:  config.NewConfigManager(worker.String(), c),
		},
	}

	return workers.NewWorker(worker)
}
