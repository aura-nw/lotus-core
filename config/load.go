package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// LoadBridgeConfig loads config from toml file to BridgeConfig
func LoadBridgeConfig(path string) (BridgeConfig, error) {
	var config BridgeConfig

	// Decode the TOML file into the config struct
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return config, fmt.Errorf("failed to decode TOML configuration: %w", err)
	}

	return config, nil
}

// LoadOperatorConfig loads config from toml file to OperatorConfig
func LoadOperatorConfig(path string) (OperatorConfig, error) {
	var config OperatorConfig

	// Decode the TOML file into the config struct
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return config, fmt.Errorf("failed to decode TOML configuration: %w", err)
	}

	return config, nil
}
