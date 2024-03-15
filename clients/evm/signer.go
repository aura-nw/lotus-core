package evm

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/aura-nw/btc-bridge/clients/evm/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ContractSigner struct {
	logger *slog.Logger
	auth   *bind.TransactOpts
	sender common.Address
	client *Client

	// Contracts
	WBitcoin *contracts.AuraWrappedBitcoin
	Gateway  *contracts.Gateway
}

type ContractInfo struct {
	WBitcoinAddr common.Address
	GatewayAddr  common.Address
}

func NewContractSigner(
	logger *slog.Logger,
	c *Client,
	privateKey *ecdsa.PrivateKey,
	info ContractInfo,
) (*ContractSigner, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(c.ChainID()))
	if err != nil {
		return nil, err
	}

	wbtcContract, err := contracts.NewAuraWrappedBitcoin(info.WBitcoinAddr, c.client)
	if err != nil {
		return nil, err
	}
	gatewayContract, err := contracts.NewGateway(info.GatewayAddr, c.client)
	if err != nil {
		return nil, err
	}

	return &ContractSigner{
		logger:   logger,
		sender:   crypto.PubkeyToAddress(*publicKeyECDSA),
		auth:     auth,
		client:   c,
		WBitcoin: wbtcContract,
		Gateway:  gatewayContract,
	}, nil
}

func (cs *ContractSigner) CreateIncomingInvoice(utxo string, amount *big.Int, recipient common.Address) error {
	tx, err := cs.Gateway.CreateIncomingInvoice(cs.auth, utxo, amount, recipient)
	if err != nil {
		cs.logger.Error("CreateIncomingInvoice failed", "err", err)
		return err
	}
	_ = tx
	return nil
}

func SignMsg(privateKey *ecdsa.PrivateKey, msg []byte) ([]byte, error) {
	// Hash the data using Keccak-256
	hash := crypto.Keccak256Hash(msg)
	sig, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

func VerifySig(publicKey *ecdsa.PublicKey, msg []byte, sig []byte) (bool, error) {
	// Hash the data using Keccak-256
	hash := crypto.Keccak256Hash(msg)

	// Verify the signature
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), sig)
	if err != nil {
		return false, err
	}

	// Compare the recovered public key with the expected public key
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	if !bytes.Equal(sigPublicKey, publicKeyBytes) {
		return false, nil
	}
	return true, nil
}
