// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"fmt"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/connections"
	"github.com/powerunit-io/platform/connections/adapters/mqtt"
	"github.com/powerunit-io/platform/connections/adapters/mysql"
	"github.com/powerunit-io/platform/devices"
	"github.com/powerunit-io/platform/events"
	"github.com/powerunit-io/platform/models"
	"github.com/powerunit-io/platform/service"
	"github.com/powerunit-io/platform/workers"
)

// BridgeService - Proxy above platform base service
type BridgeService struct {
	service.BaseService

	Rooms  models.Model
	Events chan events.Event
}

// Setup - Will setup all workers/devices/connections that service requires before
// it can be started
func (bs *BridgeService) Setup() error {
	bs.Info("Setting up (service: %s) ...", bs.Name())

	if err := bs.AttachMySQLAdapter(DbAccessName, ConfigMySqlPrimary); err != nil {
		return fmt.Errorf("Failed to create and attach database adapter (err: %s)", err)
	}

	if err := bs.AttachMqttAdapter(PrimaryMqttConnection, ConfigPrimaryMqttConnection); err != nil {
		return fmt.Errorf("Failed to create and attach mqtt adapter (err: %s)", err)
	}

	if err := bs.AttachRooms(); err != nil {
		return fmt.Errorf("Failed to attach rooms due to (err: %s)", err)
	}

	if err := bs.AttachDeviceWorker(PrimaryDeviceWorker, ConfigPrimaryDeviceWorker); err != nil {
		return fmt.Errorf("Failed to create and attach device worker (err: %s)", err)
	}

	return nil
}

// AttachRooms -
func (bs *BridgeService) AttachRooms() (err error) {
	return
}

// AttachDeviceWorker -
func (bs *BridgeService) AttachDeviceWorker(n string, c map[string]interface{}) (err error) {
	bs.Info("Setting up (service: %s) device (worker: %s)...", bs.Name(), n)

	var worker workers.Worker

	if worker, err = NewDeviceWorker(n, c, bs); err != nil {
		return err
	}

	bs.Workers.Attach(worker.Name(), worker)

	return
}

// AttachMySQLAdapter - Helper func designed to create and attach (but not start)
// mysql adapter. Additionally this is a DRY attempt to make sure code is as organized
// as possible.
func (bs *BridgeService) AttachMySQLAdapter(n string, c map[string]interface{}) (err error) {
	bs.Info("Setting up (service: %s) mysql database (adapter: %s)...", bs.Name(), n)

	var adapter mysql.Adapter

	if adapter, err = mysql.NewAdapter(n, c, bs.Logger); err != nil {
		return
	}

	bs.Connections.Attach(adapter.Name(), adapter)
	return
}

// AttachMqttAdapter -
func (bs *BridgeService) AttachMqttAdapter(n string, c map[string]interface{}) (err error) {
	bs.Info("Setting up (service: %s) mqtt connection (adapter: %s)...", bs.Name(), n)

	var adapter mqtt.Adapter

	if adapter, err = mqtt.NewAdapter(n, c, bs.Logger); err != nil {
		return
	}

	bs.Connections.Attach(adapter.Name(), adapter)
	return
}

// NewBridgeService - Will initiate new powerunit-io platform service. In case
// of any errors we will be returning error
func NewBridgeService() (service.Service, error) {
	var err error
	var cnf *config.Config

	if cnf, err = config.NewConfigManager(ServiceName, ServiceConfig); err != nil {
		return nil, fmt.Errorf("Received error while building configuration (err: %s)", err)
	}

	service := service.Service(&BridgeService{
		BaseService: service.BaseService{
			Logger:      logger,
			Config:      cnf,
			Workers:     workers.NewManager(logger),
			Connections: connections.NewManager(logger),
			Devices:     devices.NewManager(logger),
		},
		Rooms:  NewRooms(),
		Events: make(chan events.Event),
	})

	if err := service.Setup(); err != nil {
		return nil, err
	}

	return service, nil
}
