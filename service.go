// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"fmt"

	"github.com/powerunit-io/bridge/models"
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/connections"
	"github.com/powerunit-io/platform/connections/adapters/mqtt"
	"github.com/powerunit-io/platform/connections/adapters/mysql"
	"github.com/powerunit-io/platform/devices"
	"github.com/powerunit-io/platform/events"
	"github.com/powerunit-io/platform/service"
	"github.com/powerunit-io/platform/workers"
)

// BridgeService - Proxy above platform base service
type BridgeService struct {
	service.BaseService

	Building models.Building
	Rooms    []models.Room

	Events chan events.Event
}

// Setup - Will setup all workers/devices/connections that service requires before
// it can be started
func (bs *BridgeService) Setup() error {
	bs.Info("Setting up (service: %s) ...", bs.Name())

	if err := bs.AttachSQLAdapter(DbAccessName, ConfigMySqlPrimary); err != nil {
		return fmt.Errorf("Failed to create and attach database adapter (err: %s)", err)
	}

	if err := bs.AttachMqttAdapter(PrimaryMqttConnection, ConfigPrimaryMqttConnection); err != nil {
		return fmt.Errorf("Failed to create and attach mqtt adapter (err: %s)", err)
	}

	// We need to start connections in order to be able aggregate rooms n stuff...
	if err := bs.StartConnections(); err != nil {
		return fmt.Errorf("Failed to start connections due to (err: %s)", err)
	}

	//
	if err := bs.ScanForBuilding(); err != nil {
		return fmt.Errorf("Failed to scan building data due to (err: %s)", err)
	}

	//
	if err := bs.ScanForRooms(); err != nil {
		return fmt.Errorf("Failed to scan rooms data due to (err: %s)", err)
	}

	if err := bs.AttachRoomWorker(PrimaryDeviceWorker, ConfigPrimaryDeviceWorker); err != nil {
		return fmt.Errorf("Failed to create and attach device worker (err: %s)", err)
	}

	return nil
}

// AttachRoomWorker -
func (bs *BridgeService) AttachRoomWorker(n string, c map[string]interface{}) (err error) {

	if db, err = bs.GetDb(); err != nil {
		return err
	}

	for _, room := range bs.Rooms {

		var floor models.Floor

		db.Model(&room).Related(&floor)

		bs.Info("Setting up (service: %s) room (worker: %s)...",
			bs.Name(), fmt.Sprintf("%s --- %s", room.Name, floor.Name),
		)

		var worker workers.Worker

		if worker, err = NewRoomWorker(room, floor, c, bs); err != nil {
			return err
		}

		bs.Workers.Attach(worker.Name(), worker)
	}

	return
}

// AttachSQLAdapter - Helper func designed to create and attach (but not start)
// mysql adapter. Additionally this is a DRY attempt to make sure code is as organized
// as possible.
func (bs *BridgeService) AttachSQLAdapter(n string, c map[string]interface{}) (err error) {
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
			Runtime:     make(chan string),
			Done:        make(chan bool),
		},
		Events: make(chan events.Event),
	})

	if err := service.Setup(); err != nil {
		return nil, err
	}

	return service, nil
}
