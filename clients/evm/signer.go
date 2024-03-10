package evm

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	bridgeEvmPrivkeyEnv = "BRIDGE_EVM_PRIVKEY"
)

type ContractSigner struct {
	auth   *bind.TransactOpts
	sender common.Address
	client *Client
}

func NewContractSigner(client *Client) (*ContractSigner, error) {
	privKeyStr := os.Getenv(bridgeEvmPrivkeyEnv)
	if privKeyStr == "" {
		return nil, fmt.Errorf("private key is not set to BRIDGE_EVM_PRIVKEY")
	}

	privKey, err := crypto.HexToECDSA(privKeyStr)
	if err != nil {
		return nil, err
	}
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, big.NewInt(client.ChainID()))
	if err != nil {
		return nil, err
	}

	return &ContractSigner{
		sender: crypto.PubkeyToAddress(*publicKeyECDSA),
		auth:   auth,
		client: client,
	}, nil
}
