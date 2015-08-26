package main

import "os"

var (

	// ServiceConfig - Configuration related to base service ConfigManager
	ServiceConfig = map[string]interface{}{
		"ServiceName":        ServiceName,
		"ServiceDescription": ServiceDescription,
		"ServiceVersion":     ServiceVersion,
	}

	// DevicesWorkerConfig -
	DevicesWorkerConfig = map[string]interface{}{
		"connections": map[string]interface{}{
			"mqtt": map[string]interface{}{
				"mqtt_primary": map[string]interface{}{
					"network":  os.Getenv("PU_BRIDGE_MQTT_PRIMARY_NETWORK"),
					"address":  os.Getenv("PU_BRIDGE_MQTT_PRIMARY_ADDRESS"),
					"clientId": os.Getenv("PU_BRIDGE_MQTT_PRIMARY_CLIENT_ID"),
				},
				"mqtt_secondary": map[string]interface{}{
					"network":  os.Getenv("PU_BRIDGE_MQTT_SECONDARY_NETWORK"),
					"address":  os.Getenv("PU_BRIDGE_MQTT_SECONDARY_ADDRESS"),
					"clientId": os.Getenv("PU_BRIDGE_MQTT_SECONDARY_CLIENT_ID"),
				},
			},
			"mysql": map[string]interface{}{
				"mysql_primary": map[string]interface{}{
					"uri": os.Getenv("PU_BRIDGE_MYSQL_PRIMARY_URI"),
				},
			},
			"amqp": map[string]interface{}{
				"amqp_primary": map[string]interface{}{
					"uri":      os.Getenv("PU_BRIDGE_AMQP_PRIMARY_URI"),
					"exchange": os.Getenv("PU_BRIDGE_AMQP_PRIMARY_EXCHANGE"),
					"type":     os.Getenv("PU_BRIDGE_AMQP_PRIMARY_TYPE"),
					"key":      os.Getenv("PU_BRIDGE_AMQP_PRIMARY_KEY"),
					"queue":    os.Getenv("PU_BRIDGE_AMQP_PRIMARY_QUEUE"),
				},
			},
		},
	}
)
