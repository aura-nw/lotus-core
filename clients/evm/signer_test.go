package evm

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSigner(t *testing.T) {
	signer, err := NewContractSigner(slog.Default(), nil, nil, ContractInfo{})
	require.NoError(t, err)
	require.NotNil(t, signer)
}
