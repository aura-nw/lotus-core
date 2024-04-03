package evm

import (
	"context"
	"log/slog"
	"math/big"
	"strconv"
	"time"

	"github.com/aura-nw/lotus-core/clients/evm/contracts"
	"github.com/aura-nw/lotus-core/config"
	"github.com/aura-nw/lotus-core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	ChainID() int64
	GetAddress() common.Address

	CreateIncomingInvoice(deposit *types.BtcDeposit) error

	FilterOutgoingInvoice() ([]types.BtcWithdraw, error)
	SubmitTXContent(btcWithdraws []types.BtcWithdraw, txHex string) error
	//GetOutgoingInvoiceCount() (*big.Int, error)
	//GetOutgoingInvoice(id uint64) (contracts.IGatewayOutgoingInvoiceResponse, error)
	GetOutgoingTxCount() (*big.Int, error)
	GetOutgoingTx(id uint64) (contracts.IGatewayOutgoingTxInfo, error)
}

type clientImpl struct {
	logger *slog.Logger
	info   config.EvmInfo

	client          *ethclient.Client
	auth            *bind.TransactOpts
	gatewayContract *contracts.Gateway
}

// CreateIncomingInvoice implements Client.
func (c *clientImpl) CreateIncomingInvoice(deposit *types.BtcDeposit) error {
	utxo := deposit.TxId
	amount, err := strconv.ParseInt(deposit.Amount, 10, 64)
	if err != nil {
		return err
	}
	amountInt := big.NewInt(amount)
	tx, err := c.gatewayContract.CreateIncomingInvoice(c.auth, utxo, amountInt, common.HexToAddress(deposit.Receiver))
	if err != nil {
		c.logger.Error("call CreateIncomingInvoice error", "err", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.info.CallTimeout)*time.Second)
	defer cancel()
	receipt, err := bind.WaitMined(ctx, c.client, tx)
	if err != nil {
		c.logger.Error("call WaitMined error", "err", err)
		return err
	}
	c.logger.Info("call WaitMined ok", "tx_hash", receipt.TxHash.Hex())
	return nil
}

// GetAddress implements Client.
func (c *clientImpl) GetAddress() common.Address {
	return c.auth.From
}

// ChainID implements Client.
func (c *clientImpl) ChainID() int64 {
	return c.info.ChainID
}

// FilterOutgoingInvoice implements Client.
func (c *clientImpl) FilterOutgoingInvoice() ([]types.BtcWithdraw, error) {
	// TODO: implement
	return nil, nil
}

// SubmitTXContent implements Client.
func (c *clientImpl) SubmitTXContent(btcWithdraws []types.BtcWithdraw, txHex string) error {
	invoices := make([]contracts.IGatewayOutgoingInvoiceBasicInfo, 0, len(btcWithdraws))
	for _, btcWithdraw := range btcWithdraws {
		amountBig := new(big.Int).SetInt64(int64(btcWithdraw.Amount))
		invoice := contracts.IGatewayOutgoingInvoiceBasicInfo{
			Recipient: btcWithdraw.Address,
			Amount:    amountBig,
			InvoiceId: big.NewInt(btcWithdraw.InvoiceId),
		}
		invoices = append(invoices, invoice)
	}
	_, err := c.gatewayContract.SubmitTxContent(&bind.TransactOpts{}, invoices, txHex)
	if err != nil {
		return err
	}
	return nil
}

//// GetOutgoingInvoiceCount implements Client.
//func (c *clientImpl) GetOutgoingInvoiceCount() (*big.Int, error) {
//	return c.gatewayContract.OutgoingInvoicesCount(&bind.CallOpts{})
//}

// GetOutgoingTxCount implements Client.
func (c *clientImpl) GetOutgoingTxCount() (*big.Int, error) {
	return c.gatewayContract.OutgoingTxCount(&bind.CallOpts{})

}

// GetOutgoingTx implements Client.
func (c *clientImpl) GetOutgoingTx(id uint64) (contracts.IGatewayOutgoingTxInfo, error) {
	return c.gatewayContract.OutgoingTx(&bind.CallOpts{}, big.NewInt(int64(id)))
}

// GetOutgoingInvoice implements Client.
//func (c *clientImpl) GetOutgoingInvoice(id uint64) (contracts.IGatewayOutgoingInvoiceResponse, error) {
//	return c.gatewayContract.OutgoingInvoice(&bind.CallOpts{}, big.NewInt(int64(id)))
//}

func NewClient(logger *slog.Logger, info config.EvmInfo) (Client, error) {
	client, err := ethclient.Dial(info.Url)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(info.PrivateKey)
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(info.ChainID))
	if err != nil {
		return nil, err
	}

	gatewayContract, err := contracts.NewGateway(common.HexToAddress(info.Contracts.GatewayAddr), client)
	if err != nil {
		return nil, err
	}

	return &clientImpl{
		logger:          logger,
		info:            info,
		client:          client,
		auth:            auth,
		gatewayContract: gatewayContract,
	}, nil
}

var _ Client = &clientImpl{}
