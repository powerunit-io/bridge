package main

import (
	"fmt"

	"github.com/powerunit-io/platform/config"
	"github.com/powerunit-io/platform/logging"
	"github.com/powerunit-io/platform/service"

	mqttworkers "github.com/powerunit-io/platform/workers/mqtt"
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
