package main

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/service"
	"github.com/powerunit-io/platform/workers/manager"
)

// BridgeService -
type BridgeService struct {
	*service.BaseService
}

// NewService -
func NewService(logger *logging.Logger, config *config.Config, wm manager.Manager) service.Service {
	serv := BridgeService{
		&service.BaseService{
			Logger:  logger,
			Config:  config,
			Manager: wm,
		},
	}

	return service.NewService(serv)
}
