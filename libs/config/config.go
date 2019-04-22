// Package implement logic for config.{yaml|json}.
package config

import (
	"github.com/spf13/viper"
	"go-microservices/libs/logger"
)

var isInit bool = false

// Init dirs and paths to config(yaml/json) file.
func initConfig() error {
	viper.SetEnvPrefix("")

	// Common variables.
	viper.BindEnv("dev")
	viper.BindEnv("docker")
	viper.SetDefault("dev", true)
	viper.SetDefault("docker", false)

	// NATS.
	viper.BindEnv("nats_host")
	viper.BindEnv("nats_port")
	viper.BindEnv("nats_user")
	viper.BindEnv("nats_pass")
	viper.BindEnv("nats_cluster")
	viper.SetDefault("nats_host", "localhost")
	viper.SetDefault("nats_port", 4222)
	viper.SetDefault("nats_user", "")
	viper.SetDefault("nats_pass", "")
	viper.SetDefault("nats_cluster", "test-cluster")

	// HTTP service.
	viper.BindEnv("client_service_port")
	viper.SetDefault("client_service_port", 9110)

	isInit = true
	logger.GetConfig().Info(10)
	return nil
}

// Get INT value from config.
func GetInt(key string) int {
	if !isInit {
		initConfig()
	}
	value := viper.GetInt(key)
	logger.GetConfig().Info(20, map[string]interface{}{
		"type":  "INT",
		"key":   key,
		"value": value,
	})
	return value
}

// Get INT64 value from config.
func GetInt64(key string) int64 {
	if !isInit {
		initConfig()
	}
	value := viper.GetInt64(key)
	logger.GetConfig().Info(30, map[string]interface{}{
		"type":  "INT64",
		"key":   key,
		"value": value,
	})
	return value
}

// Get FLOAT value from config.
func GetFloat64(key string) float64 {
	if !isInit {
		initConfig()
	}
	value := viper.GetFloat64(key)
	logger.GetConfig().Info(40, map[string]interface{}{
		"type":  "FLOAT64",
		"key":   key,
		"value": value,
	})
	return value
}

// Get BOOLEAN value from config.
func GetBool(key string) bool {
	if !isInit {
		initConfig()
	}
	value := viper.GetBool(key)
	logger.GetConfig().Info(50, map[string]interface{}{
		"type":  "BOOl",
		"key":   key,
		"value": value,
	})
	return value
}

// Get STRING value from config.
func GetString(key string) string {
	if !isInit {
		initConfig()
	}
	value := viper.GetString(key)
	logger.GetConfig().Info(60, map[string]interface{}{
		"type":  "STRING",
		"key":   key,
		"value": value,
	})
	return value
}