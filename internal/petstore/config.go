package petstore

import (
	"github.com/anz-bank/sysl-go/config"
)

// ConfigContainer struct
type ConfigContainer struct {
	config.DefaultConfig
	CustomConfig
}

// RootConfig struct
type RootConfig struct {
	CustomConfig CustomConfig `yaml:"custom"`
}

// CustomConfig struct
type CustomConfig struct {
	AdminServer config.CommonHTTPServerConfig `yaml:"adminServer"`
}
