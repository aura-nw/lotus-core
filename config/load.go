package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// LoadConfig loads config from toml file to BridgeConfig
func LoadConfig(path string) (Config, error) {
	var config Config

	// Decode the TOML file into the config struct
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return config, fmt.Errorf("failed to decode TOML configuration: %w", err)
	}

	return config, nil
}
