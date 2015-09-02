package main

import (
	"github.com/powerunit-io/platform/config"
	helpers "github.com/powerunit-io/platform/helpers/manager"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/service"
	"github.com/powerunit-io/platform/workers/manager"
)

// BridgeService -
type BridgeService struct {
	*service.BaseService
}

// NewBridgeService - Will initiate new powerunit-io platform service
func NewBridgeService(logger *logging.Logger, config *config.Config, wm manager.Manager, hm helpers.Manager) service.Service {
	serv := BridgeService{
		&service.BaseService{
			Logger:         logger,
			Config:         config,
			Manager:        wm,
			HelpersManager: hm,
		},
	}

	return service.NewService(serv)
}
