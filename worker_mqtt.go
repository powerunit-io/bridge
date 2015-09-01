package main

import (
	"fmt"
	"time"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/connections/mqtt"
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
	mw.Warning("Starting up (worker: %s) ...", mw.String())

	errors := mw.MqttConnection.Start(done)

	// Just one error for now ...
	select {
	case err := <-errors:
		return fmt.Errorf(
			"Failed to start mqtt connection for (worker: %s) due to (err: %s)",
			mw.String(), err,
		)
	case <-time.After(2 * time.Second):
		break
	}

	go mw.Handle(done)

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

// Stop -
func (mw *MqttWorker) Stop() error {
	mw.Warning("Stopping (worker: %s) ...", mw.String())

	// Stopping MqttConnection worker ...
	mw.MqttConnection.Stop()

	return nil
}

// Handle -
func (mw *MqttWorker) Handle(done chan bool) {
	mw.Info("Listening for events now (worker: %s)", mw.String())

	for {
		select {
		case event := <-mw.MqttConnection.DrainEvents():
			mw.Debug("Got new (event: %s) for (worker: %s)", event, mw.String())
		case <-done:
			mw.Warning(
				"Received kill signal from service. Killing (worker: %s) event handler now ...",
				mw.String(),
			)
			return
		}
	}

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

	return worker.New(w), nil
}
