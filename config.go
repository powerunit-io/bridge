package main

import "os"

var (

	// ServiceConfig - Configuration related to base service ConfigManager
	ServiceConfig = map[string]interface{}{
		"service_name":        ServiceName,
		"service_description": ServiceDescription,
		"service_version":     ServiceVersion,
	}

	// Config_MqttPrimaryWorker -
	Config_MqttPrimaryWorker = map[string]interface{}{
		"connection": map[string]interface{}{
			"network":  os.Getenv("PU_BRIDGE_MQTT_PRIMARY_NETWORK"),
			"username": os.Getenv("PU_BRIDGE_MQTT_PRIMARY_USERNAME"),
			"password": os.Getenv("PU_BRIDGE_MQTT_PRIMARY_PASSWORD"),
			"address":  os.Getenv("PU_BRIDGE_MQTT_PRIMARY_ADDRESS"),
			"topic":    os.Getenv("PU_BRIDGE_MQTT_PRIMARY_TOPIC"),
			"clientId": os.Getenv("PU_BRIDGE_MQTT_PRIMARY_CLIENT_ID"),
		},
	}

	// Config_MqttSecondaryWorker -
	Config_MqttSecondaryWorker = map[string]interface{}{
		"connection": map[string]interface{}{
			"network":  os.Getenv("PU_BRIDGE_MQTT_SECONDARY_NETWORK"),
			"address":  os.Getenv("PU_BRIDGE_MQTT_SECONDARY_ADDRESS"),
			"username": os.Getenv("PU_BRIDGE_MQTT_SECONDARY_USERNAME"),
			"password": os.Getenv("PU_BRIDGE_MQTT_SECONDARY_PASSWORD"),
			"topic":    os.Getenv("PU_BRIDGE_MQTT_SECONDARY_TOPIC"),
			"clientId": os.Getenv("PU_BRIDGE_MQTT_SECONDARY_CLIENT_ID"),
		},
	}
)
