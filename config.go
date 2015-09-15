// Copyright 2015 The PowerUnit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package main ...
package main

import "os"

var (

	// ServiceConfig - Configuration related to base service ConfigManager
	ServiceConfig = map[string]interface{}{
		"service_name":        ServiceName,
		"service_description": ServiceDescription,
		"service_version":     ServiceVersion,
	}

	// LoggingConfig -
	LoggingConfig = map[string]interface{}{
		"output":                     os.Stderr,
		"level":                      os.Getenv("PU_BRIDGE_LOG_LEVEL"),
		"formatter_force_colors":     true,
		"formatter_timestamp_format": "Mon Jan _2 15:04:05 2015",
	}

	// ConfigMySqlPrimary -
	ConfigMySqlPrimary = map[string]interface{}{
		"uri": os.Getenv("PU_BRIDGE_MYSQL_PRIMARY_URI"),
	}

	// ConfigPrimaryMqttConnection -
	ConfigPrimaryMqttConnection = map[string]interface{}{
		"connection": map[string]interface{}{
			"network":  os.Getenv("PU_BRIDGE_MQTT_PRIMARY_NETWORK"),
			"username": os.Getenv("PU_BRIDGE_MQTT_PRIMARY_USERNAME"),
			"password": os.Getenv("PU_BRIDGE_MQTT_PRIMARY_PASSWORD"),
			"address":  os.Getenv("PU_BRIDGE_MQTT_PRIMARY_ADDRESS"),
			"topic":    os.Getenv("PU_BRIDGE_MQTT_PRIMARY_TOPIC"),
			"clientId": os.Getenv("PU_BRIDGE_MQTT_PRIMARY_CLIENT_ID"),
		},
	}

	// ConfigPrimaryDeviceWorker -
	ConfigPrimaryDeviceWorker = map[string]interface{}{}
)
