package main

import (
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/service"
)

// BridgeService -
type BridgeService struct {
	*service.BaseService
}

// NewService -
func NewService(logger *logging.Logger) service.Service {
	serv := BridgeService{
		&service.BaseService{
			logger,
		},
	}

	return service.NewService(serv)
}
