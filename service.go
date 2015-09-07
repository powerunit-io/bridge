package main

import (
	"github.com/powerunit-io/platform/config"
	devices "github.com/powerunit-io/platform/devices/manager"
	helpers "github.com/powerunit-io/platform/helpers/manager"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/service"
	workers "github.com/powerunit-io/platform/workers/manager"
)

// BridgeService -
type BridgeService struct {
	*service.BaseService
}

// NewBridgeService - Will initiate new powerunit-io platform service
func NewBridgeService(logger *logging.Logger, config *config.Config, wm workers.Manager, hm helpers.Manager, dm devices.Manager) service.Service {
	return service.NewService(BridgeService{
		&service.BaseService{
			Logger:         logger,
			Config:         config,
			Manager:        wm,
			HelpersManager: hm,
			DevicesManager: dm,
			Binds:          make(map[string]service.BindManager),
		},
	})
}
