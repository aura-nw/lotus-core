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

	if err := config.Validate(); err != nil {
		return config, err
	}

	return config, nil
}

func (c Config) Validate() error {
	if c.BridgeInterval <= 0 {
		return fmt.Errorf("invalid bridge interval: got %d", c.BridgeInterval)
	}

	return nil
}
