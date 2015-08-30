package main

import (
	"fmt"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/connections/mqtt"
	"github.com/powerunit-io/platform/events"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/workers/worker"
)

// MqttWorker -
type MqttWorker struct {
	*logging.Logger
	mqtt.MqttConnection
	worker.BaseWorker
}

// Start -
func (mw *MqttWorker) Start(done chan bool) error {
	mw.Warning("Starting up (device: %s) ...", mw.String())
	return nil
}

// Validate -
func (mw *MqttWorker) Validate() error {

	// Validate basic stuff ...
	if err := mw.BaseWorker.Validate(); err != nil {
		return err
	}

	// Validate MQTT connection requirements ...
	if err := mw.MqttConnection.Validate(); err != nil {
		return err
	}

	return nil
}

// Start -
func (mw *MqttWorker) Stop() error {
	mw.Warning("Starting up (device: %s) ...", mw.String())
	return nil
}

// Handle -
func (dw *MqttWorker) Handle(e *events.Event) error {
	return nil
}

// NewMqttWorker -
func NewMqttWorker(name string, log *logging.Logger, data map[string]interface{}) (worker.Worker, error) {
	conf, err := config.NewConfigManager(name, data)

	if err != nil {
		return nil, fmt.Errorf("Could not make new mqtt worker due to (error: %s)", err)
	}

	conf.Set("worker_name", name)

	w := &MqttWorker{
		Logger:         log,
		BaseWorker:     worker.BaseWorker{Logger: log, Config: conf},
		MqttConnection: mqtt.MqttConnection{Logger: log, Config: conf},
	}

	return worker.NewWorker(w), nil
}
