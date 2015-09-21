// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import (
	"fmt"

	"github.com/powerunit-io/bridge/models"
	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/events"
	"github.com/powerunit-io/platform/workers"
)

// RoomWorker -
type RoomWorker struct {
	workers.WorkerBase

	Building models.Building
	Floor    models.Floor
	Room     models.Room

	*BridgeService
}

// Handle -
func (w *RoomWorker) Handle(e <-chan events.Event) {

}

// ----------------------- DEFAULTS --------------------------------------------

// Start -
func (w *RoomWorker) Start(done chan bool) error {
	return nil // Nothing to additionally start at this moment.
}

// Validate -
func (w *RoomWorker) Validate() error {
	return nil // Nothing extra to validate in this moment
}

// Stop -
func (w *RoomWorker) Stop() error {
	return nil // As start has nothing, stop have nothing to do as well.
}

// NewRoomWorker -
func NewRoomWorker(room models.Room, floor models.Floor, c map[string]interface{}, s *BridgeService) (workers.Worker, error) {
	name := fmt.Sprintf("%s of %s", room.Name, floor.Name)

	conf, err := config.NewConfigManager(name, c)

	if err != nil {
		s.Error("Failed to configure configuration manager for (room_worker: %s) (error: %s)", name, err)
		return nil, err
	}

	conf.Set("name", name)

	return workers.Worker(&RoomWorker{
		BridgeService: s,
		Room:          room,
		Floor:         floor,
		Building:      s.Building,
		WorkerBase: workers.WorkerBase{
			Config: conf,
		},
	}), nil
}

/**
import (
	"fmt"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/service"

	"github.com/powerunit-io/platform/workers/mqtt"
	"github.com/powerunit-io/platform/workers/worker"
)

// MqttWorker -
type MqttWorker struct {
	Service service.Service
	*logging.Logger
	*mqttworkers.Worker
}

// Handle -
func (w *MqttWorker) Handle(done chan bool) {
	w.Info("Listening for events now (worker: %s)", w.WorkerName())

handlerloop:
	for {
		select {
		case event := <-w.Worker.Drain():
			w.Debug("Got new (event: %s) for (worker: %s)", event, w.WorkerName())
		case <-done:
			w.Warning(
				"Received kill signal from service. Killing (worker: %s) event handler now ...",
				w.String(),
			)
			break handlerloop
		}
	}

	return

}



// NewMqttWorker - Will abstract MqttWorker and return back new Worker instance
func NewMqttWorker(name string, serv service.Service, log *logging.Logger, cnf map[string]interface{}) (worker.Worker, error) {
	conf, err := config.NewConfigManager(name, cnf)

	if err != nil {
		return nil, fmt.Errorf("Could not make new mqtt worker due to (error: %s)", err)
	}

	conf.Set("worker_name", name)

	if mqttworker, err := mqttworkers.NewWorker(name, log, conf); err != nil {
		return nil, err
	} else {
		return worker.New(&MqttWorker{Service: serv, Worker: mqttworker, Logger: log}), nil
	}

}
**/
