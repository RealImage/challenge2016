package config

import (
	"os"
)

type ServerConfig interface {
	GetAddress() string
	GetPort() string
}

type envServerConfig struct{}

func (c envServerConfig) GetAddress() string {
	if addr := os.Getenv("SERVER_ADDRESS"); addr != "" {
		return addr
	}
	return "localhost"
}

func (c envServerConfig) GetPort() string {
	if port := os.Getenv("SERVER_PORT"); port != "" {
		return port
	}
	return "8080"
}

func NewEnvServerConfig() ServerConfig {
	return envServerConfig{}
}
