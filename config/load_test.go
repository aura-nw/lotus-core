package config_test

import (
	"testing"

	"github.com/aura-nw/lotus-core/config"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	_, err := config.LoadConfig("../bridge.toml")
	require.NoError(t, err)
}
