package config_test

import (
	"testing"

	"github.com/aura-nw/btc-bridge/config"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	_, err := config.LoadBridgeConfig("../bridge.toml")
	require.NoError(t, err)

	_, err = config.LoadOperatorConfig("../operator.toml")
	require.NoError(t, err)
}
