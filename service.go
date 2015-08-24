package main

import (
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/service"
)

// BridgeService -
type BridgeService struct {
	*service.BaseService
}

// NewService -
func NewService(logger *logging.Logger, config *config.ConfigManager) service.Service {
	serv := BridgeService{
		&service.BaseService{
			Logger: logger,
			Config: config,
		},
	}

	return service.NewService(serv)
}
