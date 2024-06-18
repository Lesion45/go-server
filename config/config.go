package config

import (
	"os"
)

type ServerConfig struct {
	Addr string
}

// New returns a new Config struct
func New() *ServerConfig {
	return &ServerConfig{
		Addr: getEnv("SERVER-ADDRESS", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
