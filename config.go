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
			"address":  os.Getenv("PU_BRIDGE_MQTT_PRIMARY_ADDRESS"),
			"clientId": os.Getenv("PU_BRIDGE_MQTT_PRIMARY_CLIENT_ID"),
		},
	}

	// Config_MqttSecondaryWorker -
	Config_MqttSecondaryWorker = map[string]interface{}{
		"connection": map[string]interface{}{
			"network":  os.Getenv("PU_BRIDGE_MQTT_SECONDARY_NETWORK"),
			"address":  os.Getenv("PU_BRIDGE_MQTT_SECONDARY_ADDRESS"),
			"clientId": os.Getenv("PU_BRIDGE_MQTT_SECONDARY_CLIENT_ID"),
		},
	}

	/**
	DevicesWorkerConfig = map[string]interface{}{
		"available_connections": []string{"mqtt", "amqp", "mysql"},

		"connections": map[string]interface{}{
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
	**/
)
