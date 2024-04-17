// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IBRC20GatewayIncomingInvoiceResponse is an auto generated low-level Go binding around an user-defined struct.
type IBRC20GatewayIncomingInvoiceResponse struct {
	InvoiceId     *big.Int
	Utxo          string
	TokenId       string
	Amount        *big.Int
	Recipient     common.Address
	Status        uint8
	Validators    []common.Address
	Confirmations []bool
	UtxoRefund    string
}

// IBRC20GatewayOutgoingInvoiceBasicInfo is an auto generated low-level Go binding around an user-defined struct.
type IBRC20GatewayOutgoingInvoiceBasicInfo struct {
	InvoiceId     *big.Int
	TokenContract common.Address
	Amount        *big.Int
	Recipient     string
}

// IBRC20GatewayOutgoingInvoiceResponse is an auto generated low-level Go binding around an user-defined struct.
type IBRC20GatewayOutgoingInvoiceResponse struct {
	InvoiceId     *big.Int
	From          common.Address
	TokenContract common.Address
	Amount        *big.Int
	Recipient     string
	Status        uint8
	TxId          *big.Int
}

// IBRC20GatewayOutgoingTxInfo is an auto generated low-level Go binding around an user-defined struct.
type IBRC20GatewayOutgoingTxInfo struct {
	TxId            *big.Int
	ContractAddress []common.Address
	Amounts         []*big.Int
	InvoiceIds      []*big.Int
	TxContent       string
	Validators      []common.Address
	Signatures      []string
	Status          uint8
	TxHash          string
}

// IBRC20GatewayValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IBRC20GatewayValidatorInfo struct {
	Validator           common.Address
	NextIncomingInvoice *big.Int
	NextOutgoingTx      *big.Int
}

// Brc20GatewayMetaData contains all meta data concerning the Brc20Gateway contract.
var Brc20GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ProposerUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"CreateNewToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIBRC20Gateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"IncomingInvoiceCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"utxo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"IncomingInvoiceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"utxo_refund\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"IncomingInvoiceRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIBRC20Gateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"IncomingInvoiceRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"IncomingInvoiceVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"}],\"name\":\"OutgoingInvoiceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIBRC20Gateway.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"OutgoingTxProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txContent\",\"type\":\"string\"}],\"name\":\"OutgoingTxSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"}],\"name\":\"OutgoingTxVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nextIncomingInvoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextOutgoingTx\",\"type\":\"uint256\"}],\"internalType\":\"structIBRC20Gateway.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"brc20Tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"completeIncomingInvoice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"concensusThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_utxo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"createIncomingInvoice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_tokenId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"createNewToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_utxo\",\"type\":\"string\"}],\"name\":\"incomingInvoice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"utxo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"enumIBRC20Gateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"confirmations\",\"type\":\"bool[]\"},{\"internalType\":\"string\",\"name\":\"utxo_refund\",\"type\":\"string\"}],\"internalType\":\"structIBRC20Gateway.IncomingInvoiceResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"incomingInvoices\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"utxo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"enumIBRC20Gateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"utxo_refund\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incomingInvoicesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_invoiceId\",\"type\":\"uint256\"}],\"name\":\"outgoingInvoice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"enumIBRC20Gateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"}],\"internalType\":\"structIBRC20Gateway.OutgoingInvoiceResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outgoingInvoices\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"enumIBRC20Gateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"txContent\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outgoingInvoicesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txId\",\"type\":\"uint256\"}],\"name\":\"outgoingTx\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"contractAddress\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"invoiceIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"txContent\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"enumIBRC20Gateway.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"internalType\":\"structIBRC20Gateway.OutgoingTxInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outgoingTxCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_txHash\",\"type\":\"string\"}],\"name\":\"processOutgoingTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_utxo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_utxo_refund\",\"type\":\"string\"}],\"name\":\"refundIncomingInvoice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"}],\"internalType\":\"structIBRC20Gateway.OutgoingInvoiceBasicInfo[]\",\"name\":\"_invoices\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"_txContent\",\"type\":\"string\"}],\"name\":\"submitTxContent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenIds\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_concensusThreshold\",\"type\":\"uint256\"}],\"name\":\"updateConcensusThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"}],\"name\":\"updateProposer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"utxos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"validator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nextIncomingInvoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextOutgoingTx\",\"type\":\"uint256\"}],\"internalType\":\"structIBRC20Gateway.ValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_utxo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_tokenId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isVerified\",\"type\":\"bool\"}],\"name\":\"verifyIncomingInvoice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isVerified\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_signature\",\"type\":\"string\"}],\"name\":\"verifyOutgoingTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Brc20GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use Brc20GatewayMetaData.ABI instead.
var Brc20GatewayABI = Brc20GatewayMetaData.ABI

// Brc20Gateway is an auto generated Go binding around an Ethereum contract.
type Brc20Gateway struct {
	Brc20GatewayCaller     // Read-only binding to the contract
	Brc20GatewayTransactor // Write-only binding to the contract
	Brc20GatewayFilterer   // Log filterer for contract events
}

// Brc20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type Brc20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Brc20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Brc20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Brc20GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Brc20GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Brc20GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Brc20GatewaySession struct {
	Contract     *Brc20Gateway     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Brc20GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Brc20GatewayCallerSession struct {
	Contract *Brc20GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Brc20GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Brc20GatewayTransactorSession struct {
	Contract     *Brc20GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Brc20GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type Brc20GatewayRaw struct {
	Contract *Brc20Gateway // Generic contract binding to access the raw methods on
}

// Brc20GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Brc20GatewayCallerRaw struct {
	Contract *Brc20GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// Brc20GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Brc20GatewayTransactorRaw struct {
	Contract *Brc20GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrc20Gateway creates a new instance of Brc20Gateway, bound to a specific deployed contract.
func NewBrc20Gateway(address common.Address, backend bind.ContractBackend) (*Brc20Gateway, error) {
	contract, err := bindBrc20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Brc20Gateway{Brc20GatewayCaller: Brc20GatewayCaller{contract: contract}, Brc20GatewayTransactor: Brc20GatewayTransactor{contract: contract}, Brc20GatewayFilterer: Brc20GatewayFilterer{contract: contract}}, nil
}

// NewBrc20GatewayCaller creates a new read-only instance of Brc20Gateway, bound to a specific deployed contract.
func NewBrc20GatewayCaller(address common.Address, caller bind.ContractCaller) (*Brc20GatewayCaller, error) {
	contract, err := bindBrc20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayCaller{contract: contract}, nil
}

// NewBrc20GatewayTransactor creates a new write-only instance of Brc20Gateway, bound to a specific deployed contract.
func NewBrc20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*Brc20GatewayTransactor, error) {
	contract, err := bindBrc20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayTransactor{contract: contract}, nil
}

// NewBrc20GatewayFilterer creates a new log filterer instance of Brc20Gateway, bound to a specific deployed contract.
func NewBrc20GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*Brc20GatewayFilterer, error) {
	contract, err := bindBrc20Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayFilterer{contract: contract}, nil
}

// bindBrc20Gateway binds a generic wrapper to an already deployed contract.
func bindBrc20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Brc20GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brc20Gateway *Brc20GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brc20Gateway.Contract.Brc20GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brc20Gateway *Brc20GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Brc20GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brc20Gateway *Brc20GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Brc20GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Brc20Gateway *Brc20GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Brc20Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Brc20Gateway *Brc20GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Brc20Gateway *Brc20GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.contract.Transact(opts, method, params...)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns((address,uint256,uint256)[])
func (_Brc20Gateway *Brc20GatewayCaller) AllValidators(opts *bind.CallOpts) ([]IBRC20GatewayValidatorInfo, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "allValidators")

	if err != nil {
		return *new([]IBRC20GatewayValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBRC20GatewayValidatorInfo)).(*[]IBRC20GatewayValidatorInfo)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns((address,uint256,uint256)[])
func (_Brc20Gateway *Brc20GatewaySession) AllValidators() ([]IBRC20GatewayValidatorInfo, error) {
	return _Brc20Gateway.Contract.AllValidators(&_Brc20Gateway.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns((address,uint256,uint256)[])
func (_Brc20Gateway *Brc20GatewayCallerSession) AllValidators() ([]IBRC20GatewayValidatorInfo, error) {
	return _Brc20Gateway.Contract.AllValidators(&_Brc20Gateway.CallOpts)
}

// Brc20Tokens is a free data retrieval call binding the contract method 0xcc91148a.
//
// Solidity: function brc20Tokens(string ) view returns(address)
func (_Brc20Gateway *Brc20GatewayCaller) Brc20Tokens(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "brc20Tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Brc20Tokens is a free data retrieval call binding the contract method 0xcc91148a.
//
// Solidity: function brc20Tokens(string ) view returns(address)
func (_Brc20Gateway *Brc20GatewaySession) Brc20Tokens(arg0 string) (common.Address, error) {
	return _Brc20Gateway.Contract.Brc20Tokens(&_Brc20Gateway.CallOpts, arg0)
}

// Brc20Tokens is a free data retrieval call binding the contract method 0xcc91148a.
//
// Solidity: function brc20Tokens(string ) view returns(address)
func (_Brc20Gateway *Brc20GatewayCallerSession) Brc20Tokens(arg0 string) (common.Address, error) {
	return _Brc20Gateway.Contract.Brc20Tokens(&_Brc20Gateway.CallOpts, arg0)
}

// ConcensusThreshold is a free data retrieval call binding the contract method 0x99bcf2dd.
//
// Solidity: function concensusThreshold() view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCaller) ConcensusThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "concensusThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConcensusThreshold is a free data retrieval call binding the contract method 0x99bcf2dd.
//
// Solidity: function concensusThreshold() view returns(uint256)
func (_Brc20Gateway *Brc20GatewaySession) ConcensusThreshold() (*big.Int, error) {
	return _Brc20Gateway.Contract.ConcensusThreshold(&_Brc20Gateway.CallOpts)
}

// ConcensusThreshold is a free data retrieval call binding the contract method 0x99bcf2dd.
//
// Solidity: function concensusThreshold() view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCallerSession) ConcensusThreshold() (*big.Int, error) {
	return _Brc20Gateway.Contract.ConcensusThreshold(&_Brc20Gateway.CallOpts)
}

// IncomingInvoice is a free data retrieval call binding the contract method 0xe7c777eb.
//
// Solidity: function incomingInvoice(string _utxo) view returns((uint256,string,string,uint256,address,uint8,address[],bool[],string))
func (_Brc20Gateway *Brc20GatewayCaller) IncomingInvoice(opts *bind.CallOpts, _utxo string) (IBRC20GatewayIncomingInvoiceResponse, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "incomingInvoice", _utxo)

	if err != nil {
		return *new(IBRC20GatewayIncomingInvoiceResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(IBRC20GatewayIncomingInvoiceResponse)).(*IBRC20GatewayIncomingInvoiceResponse)

	return out0, err

}

// IncomingInvoice is a free data retrieval call binding the contract method 0xe7c777eb.
//
// Solidity: function incomingInvoice(string _utxo) view returns((uint256,string,string,uint256,address,uint8,address[],bool[],string))
func (_Brc20Gateway *Brc20GatewaySession) IncomingInvoice(_utxo string) (IBRC20GatewayIncomingInvoiceResponse, error) {
	return _Brc20Gateway.Contract.IncomingInvoice(&_Brc20Gateway.CallOpts, _utxo)
}

// IncomingInvoice is a free data retrieval call binding the contract method 0xe7c777eb.
//
// Solidity: function incomingInvoice(string _utxo) view returns((uint256,string,string,uint256,address,uint8,address[],bool[],string))
func (_Brc20Gateway *Brc20GatewayCallerSession) IncomingInvoice(_utxo string) (IBRC20GatewayIncomingInvoiceResponse, error) {
	return _Brc20Gateway.Contract.IncomingInvoice(&_Brc20Gateway.CallOpts, _utxo)
}

// IncomingInvoices is a free data retrieval call binding the contract method 0xf5d730c8.
//
// Solidity: function incomingInvoices(uint256 ) view returns(string utxo, string tokenId, uint256 amount, address recipient, uint8 status, string utxo_refund)
func (_Brc20Gateway *Brc20GatewayCaller) IncomingInvoices(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Utxo       string
	TokenId    string
	Amount     *big.Int
	Recipient  common.Address
	Status     uint8
	UtxoRefund string
}, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "incomingInvoices", arg0)

	outstruct := new(struct {
		Utxo       string
		TokenId    string
		Amount     *big.Int
		Recipient  common.Address
		Status     uint8
		UtxoRefund string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Utxo = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.TokenId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Recipient = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.UtxoRefund = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// IncomingInvoices is a free data retrieval call binding the contract method 0xf5d730c8.
//
// Solidity: function incomingInvoices(uint256 ) view returns(string utxo, string tokenId, uint256 amount, address recipient, uint8 status, string utxo_refund)
func (_Brc20Gateway *Brc20GatewaySession) IncomingInvoices(arg0 *big.Int) (struct {
	Utxo       string
	TokenId    string
	Amount     *big.Int
	Recipient  common.Address
	Status     uint8
	UtxoRefund string
}, error) {
	return _Brc20Gateway.Contract.IncomingInvoices(&_Brc20Gateway.CallOpts, arg0)
}

// IncomingInvoices is a free data retrieval call binding the contract method 0xf5d730c8.
//
// Solidity: function incomingInvoices(uint256 ) view returns(string utxo, string tokenId, uint256 amount, address recipient, uint8 status, string utxo_refund)
func (_Brc20Gateway *Brc20GatewayCallerSession) IncomingInvoices(arg0 *big.Int) (struct {
	Utxo       string
	TokenId    string
	Amount     *big.Int
	Recipient  common.Address
	Status     uint8
	UtxoRefund string
}, error) {
	return _Brc20Gateway.Contract.IncomingInvoices(&_Brc20Gateway.CallOpts, arg0)
}

// IncomingInvoicesCount is a free data retrieval call binding the contract method 0xd6988f61.
//
// Solidity: function incomingInvoicesCount() view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCaller) IncomingInvoicesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "incomingInvoicesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncomingInvoicesCount is a free data retrieval call binding the contract method 0xd6988f61.
//
// Solidity: function incomingInvoicesCount() view returns(uint256)
func (_Brc20Gateway *Brc20GatewaySession) IncomingInvoicesCount() (*big.Int, error) {
	return _Brc20Gateway.Contract.IncomingInvoicesCount(&_Brc20Gateway.CallOpts)
}

// IncomingInvoicesCount is a free data retrieval call binding the contract method 0xd6988f61.
//
// Solidity: function incomingInvoicesCount() view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCallerSession) IncomingInvoicesCount() (*big.Int, error) {
	return _Brc20Gateway.Contract.IncomingInvoicesCount(&_Brc20Gateway.CallOpts)
}

// OutgoingInvoice is a free data retrieval call binding the contract method 0x455578b3.
//
// Solidity: function outgoingInvoice(uint256 _invoiceId) view returns((uint256,address,address,uint256,string,uint8,uint256))
func (_Brc20Gateway *Brc20GatewayCaller) OutgoingInvoice(opts *bind.CallOpts, _invoiceId *big.Int) (IBRC20GatewayOutgoingInvoiceResponse, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "outgoingInvoice", _invoiceId)

	if err != nil {
		return *new(IBRC20GatewayOutgoingInvoiceResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(IBRC20GatewayOutgoingInvoiceResponse)).(*IBRC20GatewayOutgoingInvoiceResponse)

	return out0, err

}

// OutgoingInvoice is a free data retrieval call binding the contract method 0x455578b3.
//
// Solidity: function outgoingInvoice(uint256 _invoiceId) view returns((uint256,address,address,uint256,string,uint8,uint256))
func (_Brc20Gateway *Brc20GatewaySession) OutgoingInvoice(_invoiceId *big.Int) (IBRC20GatewayOutgoingInvoiceResponse, error) {
	return _Brc20Gateway.Contract.OutgoingInvoice(&_Brc20Gateway.CallOpts, _invoiceId)
}

// OutgoingInvoice is a free data retrieval call binding the contract method 0x455578b3.
//
// Solidity: function outgoingInvoice(uint256 _invoiceId) view returns((uint256,address,address,uint256,string,uint8,uint256))
func (_Brc20Gateway *Brc20GatewayCallerSession) OutgoingInvoice(_invoiceId *big.Int) (IBRC20GatewayOutgoingInvoiceResponse, error) {
	return _Brc20Gateway.Contract.OutgoingInvoice(&_Brc20Gateway.CallOpts, _invoiceId)
}

// OutgoingInvoices is a free data retrieval call binding the contract method 0xb576bab1.
//
// Solidity: function outgoingInvoices(uint256 ) view returns(address from, address tokenContract, uint256 amount, string recipient, uint8 status, string txContent, uint256 txId)
func (_Brc20Gateway *Brc20GatewayCaller) OutgoingInvoices(opts *bind.CallOpts, arg0 *big.Int) (struct {
	From          common.Address
	TokenContract common.Address
	Amount        *big.Int
	Recipient     string
	Status        uint8
	TxContent     string
	TxId          *big.Int
}, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "outgoingInvoices", arg0)

	outstruct := new(struct {
		From          common.Address
		TokenContract common.Address
		Amount        *big.Int
		Recipient     string
		Status        uint8
		TxContent     string
		TxId          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.From = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Recipient = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.TxContent = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.TxId = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OutgoingInvoices is a free data retrieval call binding the contract method 0xb576bab1.
//
// Solidity: function outgoingInvoices(uint256 ) view returns(address from, address tokenContract, uint256 amount, string recipient, uint8 status, string txContent, uint256 txId)
func (_Brc20Gateway *Brc20GatewaySession) OutgoingInvoices(arg0 *big.Int) (struct {
	From          common.Address
	TokenContract common.Address
	Amount        *big.Int
	Recipient     string
	Status        uint8
	TxContent     string
	TxId          *big.Int
}, error) {
	return _Brc20Gateway.Contract.OutgoingInvoices(&_Brc20Gateway.CallOpts, arg0)
}

// OutgoingInvoices is a free data retrieval call binding the contract method 0xb576bab1.
//
// Solidity: function outgoingInvoices(uint256 ) view returns(address from, address tokenContract, uint256 amount, string recipient, uint8 status, string txContent, uint256 txId)
func (_Brc20Gateway *Brc20GatewayCallerSession) OutgoingInvoices(arg0 *big.Int) (struct {
	From          common.Address
	TokenContract common.Address
	Amount        *big.Int
	Recipient     string
	Status        uint8
	TxContent     string
	TxId          *big.Int
}, error) {
	return _Brc20Gateway.Contract.OutgoingInvoices(&_Brc20Gateway.CallOpts, arg0)
}

// OutgoingInvoicesCount is a free data retrieval call binding the contract method 0xc8c52a6a.
//
// Solidity: function outgoingInvoicesCount() view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCaller) OutgoingInvoicesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "outgoingInvoicesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutgoingInvoicesCount is a free data retrieval call binding the contract method 0xc8c52a6a.
//
// Solidity: function outgoingInvoicesCount() view returns(uint256)
func (_Brc20Gateway *Brc20GatewaySession) OutgoingInvoicesCount() (*big.Int, error) {
	return _Brc20Gateway.Contract.OutgoingInvoicesCount(&_Brc20Gateway.CallOpts)
}

// OutgoingInvoicesCount is a free data retrieval call binding the contract method 0xc8c52a6a.
//
// Solidity: function outgoingInvoicesCount() view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCallerSession) OutgoingInvoicesCount() (*big.Int, error) {
	return _Brc20Gateway.Contract.OutgoingInvoicesCount(&_Brc20Gateway.CallOpts)
}

// OutgoingTx is a free data retrieval call binding the contract method 0x2823434e.
//
// Solidity: function outgoingTx(uint256 _txId) view returns((uint256,address[],uint256[],uint256[],string,address[],string[],uint8,string))
func (_Brc20Gateway *Brc20GatewayCaller) OutgoingTx(opts *bind.CallOpts, _txId *big.Int) (IBRC20GatewayOutgoingTxInfo, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "outgoingTx", _txId)

	if err != nil {
		return *new(IBRC20GatewayOutgoingTxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBRC20GatewayOutgoingTxInfo)).(*IBRC20GatewayOutgoingTxInfo)

	return out0, err

}

// OutgoingTx is a free data retrieval call binding the contract method 0x2823434e.
//
// Solidity: function outgoingTx(uint256 _txId) view returns((uint256,address[],uint256[],uint256[],string,address[],string[],uint8,string))
func (_Brc20Gateway *Brc20GatewaySession) OutgoingTx(_txId *big.Int) (IBRC20GatewayOutgoingTxInfo, error) {
	return _Brc20Gateway.Contract.OutgoingTx(&_Brc20Gateway.CallOpts, _txId)
}

// OutgoingTx is a free data retrieval call binding the contract method 0x2823434e.
//
// Solidity: function outgoingTx(uint256 _txId) view returns((uint256,address[],uint256[],uint256[],string,address[],string[],uint8,string))
func (_Brc20Gateway *Brc20GatewayCallerSession) OutgoingTx(_txId *big.Int) (IBRC20GatewayOutgoingTxInfo, error) {
	return _Brc20Gateway.Contract.OutgoingTx(&_Brc20Gateway.CallOpts, _txId)
}

// OutgoingTxCount is a free data retrieval call binding the contract method 0x56770381.
//
// Solidity: function outgoingTxCount() view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCaller) OutgoingTxCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "outgoingTxCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutgoingTxCount is a free data retrieval call binding the contract method 0x56770381.
//
// Solidity: function outgoingTxCount() view returns(uint256)
func (_Brc20Gateway *Brc20GatewaySession) OutgoingTxCount() (*big.Int, error) {
	return _Brc20Gateway.Contract.OutgoingTxCount(&_Brc20Gateway.CallOpts)
}

// OutgoingTxCount is a free data retrieval call binding the contract method 0x56770381.
//
// Solidity: function outgoingTxCount() view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCallerSession) OutgoingTxCount() (*big.Int, error) {
	return _Brc20Gateway.Contract.OutgoingTxCount(&_Brc20Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brc20Gateway *Brc20GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brc20Gateway *Brc20GatewaySession) Owner() (common.Address, error) {
	return _Brc20Gateway.Contract.Owner(&_Brc20Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Brc20Gateway *Brc20GatewayCallerSession) Owner() (common.Address, error) {
	return _Brc20Gateway.Contract.Owner(&_Brc20Gateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Brc20Gateway *Brc20GatewayCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Brc20Gateway *Brc20GatewaySession) Paused() (bool, error) {
	return _Brc20Gateway.Contract.Paused(&_Brc20Gateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Brc20Gateway *Brc20GatewayCallerSession) Paused() (bool, error) {
	return _Brc20Gateway.Contract.Paused(&_Brc20Gateway.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Brc20Gateway *Brc20GatewayCaller) Proposer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "proposer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Brc20Gateway *Brc20GatewaySession) Proposer() (common.Address, error) {
	return _Brc20Gateway.Contract.Proposer(&_Brc20Gateway.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Brc20Gateway *Brc20GatewayCallerSession) Proposer() (common.Address, error) {
	return _Brc20Gateway.Contract.Proposer(&_Brc20Gateway.CallOpts)
}

// TokenIds is a free data retrieval call binding the contract method 0xfc97a303.
//
// Solidity: function tokenIds(address ) view returns(string)
func (_Brc20Gateway *Brc20GatewayCaller) TokenIds(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "tokenIds", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenIds is a free data retrieval call binding the contract method 0xfc97a303.
//
// Solidity: function tokenIds(address ) view returns(string)
func (_Brc20Gateway *Brc20GatewaySession) TokenIds(arg0 common.Address) (string, error) {
	return _Brc20Gateway.Contract.TokenIds(&_Brc20Gateway.CallOpts, arg0)
}

// TokenIds is a free data retrieval call binding the contract method 0xfc97a303.
//
// Solidity: function tokenIds(address ) view returns(string)
func (_Brc20Gateway *Brc20GatewayCallerSession) TokenIds(arg0 common.Address) (string, error) {
	return _Brc20Gateway.Contract.TokenIds(&_Brc20Gateway.CallOpts, arg0)
}

// Utxos is a free data retrieval call binding the contract method 0xf1fbbee0.
//
// Solidity: function utxos(string ) view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCaller) Utxos(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "utxos", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Utxos is a free data retrieval call binding the contract method 0xf1fbbee0.
//
// Solidity: function utxos(string ) view returns(uint256)
func (_Brc20Gateway *Brc20GatewaySession) Utxos(arg0 string) (*big.Int, error) {
	return _Brc20Gateway.Contract.Utxos(&_Brc20Gateway.CallOpts, arg0)
}

// Utxos is a free data retrieval call binding the contract method 0xf1fbbee0.
//
// Solidity: function utxos(string ) view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCallerSession) Utxos(arg0 string) (*big.Int, error) {
	return _Brc20Gateway.Contract.Utxos(&_Brc20Gateway.CallOpts, arg0)
}

// Validator is a free data retrieval call binding the contract method 0x223b3b7a.
//
// Solidity: function validator(address _validator) view returns((address,uint256,uint256))
func (_Brc20Gateway *Brc20GatewayCaller) Validator(opts *bind.CallOpts, _validator common.Address) (IBRC20GatewayValidatorInfo, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "validator", _validator)

	if err != nil {
		return *new(IBRC20GatewayValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBRC20GatewayValidatorInfo)).(*IBRC20GatewayValidatorInfo)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x223b3b7a.
//
// Solidity: function validator(address _validator) view returns((address,uint256,uint256))
func (_Brc20Gateway *Brc20GatewaySession) Validator(_validator common.Address) (IBRC20GatewayValidatorInfo, error) {
	return _Brc20Gateway.Contract.Validator(&_Brc20Gateway.CallOpts, _validator)
}

// Validator is a free data retrieval call binding the contract method 0x223b3b7a.
//
// Solidity: function validator(address _validator) view returns((address,uint256,uint256))
func (_Brc20Gateway *Brc20GatewayCallerSession) Validator(_validator common.Address) (IBRC20GatewayValidatorInfo, error) {
	return _Brc20Gateway.Contract.Validator(&_Brc20Gateway.CallOpts, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) AddValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "addValidator", _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_Brc20Gateway *Brc20GatewaySession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.AddValidator(&_Brc20Gateway.TransactOpts, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.AddValidator(&_Brc20Gateway.TransactOpts, _validator)
}

// CompleteIncomingInvoice is a paid mutator transaction binding the contract method 0xa0522b9d.
//
// Solidity: function completeIncomingInvoice(uint256 _invoiceId, address _recipient) returns(uint256)
func (_Brc20Gateway *Brc20GatewayTransactor) CompleteIncomingInvoice(opts *bind.TransactOpts, _invoiceId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "completeIncomingInvoice", _invoiceId, _recipient)
}

// CompleteIncomingInvoice is a paid mutator transaction binding the contract method 0xa0522b9d.
//
// Solidity: function completeIncomingInvoice(uint256 _invoiceId, address _recipient) returns(uint256)
func (_Brc20Gateway *Brc20GatewaySession) CompleteIncomingInvoice(_invoiceId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.CompleteIncomingInvoice(&_Brc20Gateway.TransactOpts, _invoiceId, _recipient)
}

// CompleteIncomingInvoice is a paid mutator transaction binding the contract method 0xa0522b9d.
//
// Solidity: function completeIncomingInvoice(uint256 _invoiceId, address _recipient) returns(uint256)
func (_Brc20Gateway *Brc20GatewayTransactorSession) CompleteIncomingInvoice(_invoiceId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.CompleteIncomingInvoice(&_Brc20Gateway.TransactOpts, _invoiceId, _recipient)
}

// CreateIncomingInvoice is a paid mutator transaction binding the contract method 0x5ca2f5b2.
//
// Solidity: function createIncomingInvoice(string _utxo, string _tokenId, uint256 _amount, address _recipient) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) CreateIncomingInvoice(opts *bind.TransactOpts, _utxo string, _tokenId string, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "createIncomingInvoice", _utxo, _tokenId, _amount, _recipient)
}

// CreateIncomingInvoice is a paid mutator transaction binding the contract method 0x5ca2f5b2.
//
// Solidity: function createIncomingInvoice(string _utxo, string _tokenId, uint256 _amount, address _recipient) returns()
func (_Brc20Gateway *Brc20GatewaySession) CreateIncomingInvoice(_utxo string, _tokenId string, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.CreateIncomingInvoice(&_Brc20Gateway.TransactOpts, _utxo, _tokenId, _amount, _recipient)
}

// CreateIncomingInvoice is a paid mutator transaction binding the contract method 0x5ca2f5b2.
//
// Solidity: function createIncomingInvoice(string _utxo, string _tokenId, uint256 _amount, address _recipient) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) CreateIncomingInvoice(_utxo string, _tokenId string, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.CreateIncomingInvoice(&_Brc20Gateway.TransactOpts, _utxo, _tokenId, _amount, _recipient)
}

// CreateNewToken is a paid mutator transaction binding the contract method 0x53021fae.
//
// Solidity: function createNewToken(string _tokenId, string _name, string _symbol, uint8 _decimals) returns(address)
func (_Brc20Gateway *Brc20GatewayTransactor) CreateNewToken(opts *bind.TransactOpts, _tokenId string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "createNewToken", _tokenId, _name, _symbol, _decimals)
}

// CreateNewToken is a paid mutator transaction binding the contract method 0x53021fae.
//
// Solidity: function createNewToken(string _tokenId, string _name, string _symbol, uint8 _decimals) returns(address)
func (_Brc20Gateway *Brc20GatewaySession) CreateNewToken(_tokenId string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.CreateNewToken(&_Brc20Gateway.TransactOpts, _tokenId, _name, _symbol, _decimals)
}

// CreateNewToken is a paid mutator transaction binding the contract method 0x53021fae.
//
// Solidity: function createNewToken(string _tokenId, string _name, string _symbol, uint8 _decimals) returns(address)
func (_Brc20Gateway *Brc20GatewayTransactorSession) CreateNewToken(_tokenId string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.CreateNewToken(&_Brc20Gateway.TransactOpts, _tokenId, _name, _symbol, _decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Brc20Gateway *Brc20GatewayTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Brc20Gateway *Brc20GatewaySession) Initialize() (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Initialize(&_Brc20Gateway.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) Initialize() (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Initialize(&_Brc20Gateway.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Brc20Gateway *Brc20GatewayTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Brc20Gateway *Brc20GatewaySession) Pause() (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Pause(&_Brc20Gateway.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) Pause() (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Pause(&_Brc20Gateway.TransactOpts)
}

// ProcessOutgoingTx is a paid mutator transaction binding the contract method 0x0f6dc9e2.
//
// Solidity: function processOutgoingTx(uint256 _txId, string _txHash) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) ProcessOutgoingTx(opts *bind.TransactOpts, _txId *big.Int, _txHash string) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "processOutgoingTx", _txId, _txHash)
}

// ProcessOutgoingTx is a paid mutator transaction binding the contract method 0x0f6dc9e2.
//
// Solidity: function processOutgoingTx(uint256 _txId, string _txHash) returns()
func (_Brc20Gateway *Brc20GatewaySession) ProcessOutgoingTx(_txId *big.Int, _txHash string) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.ProcessOutgoingTx(&_Brc20Gateway.TransactOpts, _txId, _txHash)
}

// ProcessOutgoingTx is a paid mutator transaction binding the contract method 0x0f6dc9e2.
//
// Solidity: function processOutgoingTx(uint256 _txId, string _txHash) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) ProcessOutgoingTx(_txId *big.Int, _txHash string) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.ProcessOutgoingTx(&_Brc20Gateway.TransactOpts, _txId, _txHash)
}

// RefundIncomingInvoice is a paid mutator transaction binding the contract method 0xfb1b9730.
//
// Solidity: function refundIncomingInvoice(uint256 _invoiceId, string _utxo, string _tokenId, uint256 _amount, address _recipient, string _utxo_refund) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) RefundIncomingInvoice(opts *bind.TransactOpts, _invoiceId *big.Int, _utxo string, _tokenId string, _amount *big.Int, _recipient common.Address, _utxo_refund string) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "refundIncomingInvoice", _invoiceId, _utxo, _tokenId, _amount, _recipient, _utxo_refund)
}

// RefundIncomingInvoice is a paid mutator transaction binding the contract method 0xfb1b9730.
//
// Solidity: function refundIncomingInvoice(uint256 _invoiceId, string _utxo, string _tokenId, uint256 _amount, address _recipient, string _utxo_refund) returns()
func (_Brc20Gateway *Brc20GatewaySession) RefundIncomingInvoice(_invoiceId *big.Int, _utxo string, _tokenId string, _amount *big.Int, _recipient common.Address, _utxo_refund string) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.RefundIncomingInvoice(&_Brc20Gateway.TransactOpts, _invoiceId, _utxo, _tokenId, _amount, _recipient, _utxo_refund)
}

// RefundIncomingInvoice is a paid mutator transaction binding the contract method 0xfb1b9730.
//
// Solidity: function refundIncomingInvoice(uint256 _invoiceId, string _utxo, string _tokenId, uint256 _amount, address _recipient, string _utxo_refund) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) RefundIncomingInvoice(_invoiceId *big.Int, _utxo string, _tokenId string, _amount *big.Int, _recipient common.Address, _utxo_refund string) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.RefundIncomingInvoice(&_Brc20Gateway.TransactOpts, _invoiceId, _utxo, _tokenId, _amount, _recipient, _utxo_refund)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) RemoveValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "removeValidator", _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_Brc20Gateway *Brc20GatewaySession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.RemoveValidator(&_Brc20Gateway.TransactOpts, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.RemoveValidator(&_Brc20Gateway.TransactOpts, _validator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brc20Gateway *Brc20GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brc20Gateway *Brc20GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _Brc20Gateway.Contract.RenounceOwnership(&_Brc20Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Brc20Gateway.Contract.RenounceOwnership(&_Brc20Gateway.TransactOpts)
}

// SubmitTxContent is a paid mutator transaction binding the contract method 0x24c70323.
//
// Solidity: function submitTxContent((uint256,address,uint256,string)[] _invoices, string _txContent) returns(uint256)
func (_Brc20Gateway *Brc20GatewayTransactor) SubmitTxContent(opts *bind.TransactOpts, _invoices []IBRC20GatewayOutgoingInvoiceBasicInfo, _txContent string) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "submitTxContent", _invoices, _txContent)
}

// SubmitTxContent is a paid mutator transaction binding the contract method 0x24c70323.
//
// Solidity: function submitTxContent((uint256,address,uint256,string)[] _invoices, string _txContent) returns(uint256)
func (_Brc20Gateway *Brc20GatewaySession) SubmitTxContent(_invoices []IBRC20GatewayOutgoingInvoiceBasicInfo, _txContent string) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.SubmitTxContent(&_Brc20Gateway.TransactOpts, _invoices, _txContent)
}

// SubmitTxContent is a paid mutator transaction binding the contract method 0x24c70323.
//
// Solidity: function submitTxContent((uint256,address,uint256,string)[] _invoices, string _txContent) returns(uint256)
func (_Brc20Gateway *Brc20GatewayTransactorSession) SubmitTxContent(_invoices []IBRC20GatewayOutgoingInvoiceBasicInfo, _txContent string) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.SubmitTxContent(&_Brc20Gateway.TransactOpts, _invoices, _txContent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brc20Gateway *Brc20GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.TransferOwnership(&_Brc20Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.TransferOwnership(&_Brc20Gateway.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Brc20Gateway *Brc20GatewayTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Brc20Gateway *Brc20GatewaySession) Unpause() (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Unpause(&_Brc20Gateway.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) Unpause() (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Unpause(&_Brc20Gateway.TransactOpts)
}

// UpdateConcensusThreshold is a paid mutator transaction binding the contract method 0x2c3228d5.
//
// Solidity: function updateConcensusThreshold(uint256 _concensusThreshold) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) UpdateConcensusThreshold(opts *bind.TransactOpts, _concensusThreshold *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "updateConcensusThreshold", _concensusThreshold)
}

// UpdateConcensusThreshold is a paid mutator transaction binding the contract method 0x2c3228d5.
//
// Solidity: function updateConcensusThreshold(uint256 _concensusThreshold) returns()
func (_Brc20Gateway *Brc20GatewaySession) UpdateConcensusThreshold(_concensusThreshold *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.UpdateConcensusThreshold(&_Brc20Gateway.TransactOpts, _concensusThreshold)
}

// UpdateConcensusThreshold is a paid mutator transaction binding the contract method 0x2c3228d5.
//
// Solidity: function updateConcensusThreshold(uint256 _concensusThreshold) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) UpdateConcensusThreshold(_concensusThreshold *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.UpdateConcensusThreshold(&_Brc20Gateway.TransactOpts, _concensusThreshold)
}

// UpdateProposer is a paid mutator transaction binding the contract method 0xddf9f09f.
//
// Solidity: function updateProposer(address _proposer) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) UpdateProposer(opts *bind.TransactOpts, _proposer common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "updateProposer", _proposer)
}

// UpdateProposer is a paid mutator transaction binding the contract method 0xddf9f09f.
//
// Solidity: function updateProposer(address _proposer) returns()
func (_Brc20Gateway *Brc20GatewaySession) UpdateProposer(_proposer common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.UpdateProposer(&_Brc20Gateway.TransactOpts, _proposer)
}

// UpdateProposer is a paid mutator transaction binding the contract method 0xddf9f09f.
//
// Solidity: function updateProposer(address _proposer) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) UpdateProposer(_proposer common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.UpdateProposer(&_Brc20Gateway.TransactOpts, _proposer)
}

// VerifyIncomingInvoice is a paid mutator transaction binding the contract method 0x3cb77a4b.
//
// Solidity: function verifyIncomingInvoice(uint256 _invoiceId, string _utxo, string _tokenId, uint256 _amount, address _recipient, bool _isVerified) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) VerifyIncomingInvoice(opts *bind.TransactOpts, _invoiceId *big.Int, _utxo string, _tokenId string, _amount *big.Int, _recipient common.Address, _isVerified bool) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "verifyIncomingInvoice", _invoiceId, _utxo, _tokenId, _amount, _recipient, _isVerified)
}

// VerifyIncomingInvoice is a paid mutator transaction binding the contract method 0x3cb77a4b.
//
// Solidity: function verifyIncomingInvoice(uint256 _invoiceId, string _utxo, string _tokenId, uint256 _amount, address _recipient, bool _isVerified) returns()
func (_Brc20Gateway *Brc20GatewaySession) VerifyIncomingInvoice(_invoiceId *big.Int, _utxo string, _tokenId string, _amount *big.Int, _recipient common.Address, _isVerified bool) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.VerifyIncomingInvoice(&_Brc20Gateway.TransactOpts, _invoiceId, _utxo, _tokenId, _amount, _recipient, _isVerified)
}

// VerifyIncomingInvoice is a paid mutator transaction binding the contract method 0x3cb77a4b.
//
// Solidity: function verifyIncomingInvoice(uint256 _invoiceId, string _utxo, string _tokenId, uint256 _amount, address _recipient, bool _isVerified) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) VerifyIncomingInvoice(_invoiceId *big.Int, _utxo string, _tokenId string, _amount *big.Int, _recipient common.Address, _isVerified bool) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.VerifyIncomingInvoice(&_Brc20Gateway.TransactOpts, _invoiceId, _utxo, _tokenId, _amount, _recipient, _isVerified)
}

// VerifyOutgoingTx is a paid mutator transaction binding the contract method 0xbd784e7a.
//
// Solidity: function verifyOutgoingTx(uint256 _txId, bool _isVerified, string _signature) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) VerifyOutgoingTx(opts *bind.TransactOpts, _txId *big.Int, _isVerified bool, _signature string) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "verifyOutgoingTx", _txId, _isVerified, _signature)
}

// VerifyOutgoingTx is a paid mutator transaction binding the contract method 0xbd784e7a.
//
// Solidity: function verifyOutgoingTx(uint256 _txId, bool _isVerified, string _signature) returns()
func (_Brc20Gateway *Brc20GatewaySession) VerifyOutgoingTx(_txId *big.Int, _isVerified bool, _signature string) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.VerifyOutgoingTx(&_Brc20Gateway.TransactOpts, _txId, _isVerified, _signature)
}

// VerifyOutgoingTx is a paid mutator transaction binding the contract method 0xbd784e7a.
//
// Solidity: function verifyOutgoingTx(uint256 _txId, bool _isVerified, string _signature) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) VerifyOutgoingTx(_txId *big.Int, _isVerified bool, _signature string) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.VerifyOutgoingTx(&_Brc20Gateway.TransactOpts, _txId, _isVerified, _signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Brc20Gateway *Brc20GatewayTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Brc20Gateway *Brc20GatewaySession) Withdraw() (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Withdraw(&_Brc20Gateway.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Withdraw(&_Brc20Gateway.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x14eaa43b.
//
// Solidity: function withdrawToken(address tokenContract, uint256 amount, string recipient) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) WithdrawToken(opts *bind.TransactOpts, tokenContract common.Address, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "withdrawToken", tokenContract, amount, recipient)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x14eaa43b.
//
// Solidity: function withdrawToken(address tokenContract, uint256 amount, string recipient) returns()
func (_Brc20Gateway *Brc20GatewaySession) WithdrawToken(tokenContract common.Address, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.WithdrawToken(&_Brc20Gateway.TransactOpts, tokenContract, amount, recipient)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x14eaa43b.
//
// Solidity: function withdrawToken(address tokenContract, uint256 amount, string recipient) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) WithdrawToken(tokenContract common.Address, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.WithdrawToken(&_Brc20Gateway.TransactOpts, tokenContract, amount, recipient)
}

// Brc20GatewayCreateNewTokenIterator is returned from FilterCreateNewToken and is used to iterate over the raw logs and unpacked data for CreateNewToken events raised by the Brc20Gateway contract.
type Brc20GatewayCreateNewTokenIterator struct {
	Event *Brc20GatewayCreateNewToken // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayCreateNewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayCreateNewToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayCreateNewToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayCreateNewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayCreateNewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayCreateNewToken represents a CreateNewToken event raised by the Brc20Gateway contract.
type Brc20GatewayCreateNewToken struct {
	Id           common.Hash
	Name         string
	Symbol       string
	Decimals     uint8
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreateNewToken is a free log retrieval operation binding the contract event 0x1979f4a8e3ab204d1cc74c15c8cdef4fa926cfc1dba6547c340eef202f24fc79.
//
// Solidity: event CreateNewToken(string indexed id, string name, string symbol, uint8 decimals, address tokenAddress)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterCreateNewToken(opts *bind.FilterOpts, id []string) (*Brc20GatewayCreateNewTokenIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "CreateNewToken", idRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayCreateNewTokenIterator{contract: _Brc20Gateway.contract, event: "CreateNewToken", logs: logs, sub: sub}, nil
}

// WatchCreateNewToken is a free log subscription operation binding the contract event 0x1979f4a8e3ab204d1cc74c15c8cdef4fa926cfc1dba6547c340eef202f24fc79.
//
// Solidity: event CreateNewToken(string indexed id, string name, string symbol, uint8 decimals, address tokenAddress)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchCreateNewToken(opts *bind.WatchOpts, sink chan<- *Brc20GatewayCreateNewToken, id []string) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "CreateNewToken", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayCreateNewToken)
				if err := _Brc20Gateway.contract.UnpackLog(event, "CreateNewToken", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateNewToken is a log parse operation binding the contract event 0x1979f4a8e3ab204d1cc74c15c8cdef4fa926cfc1dba6547c340eef202f24fc79.
//
// Solidity: event CreateNewToken(string indexed id, string name, string symbol, uint8 decimals, address tokenAddress)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseCreateNewToken(log types.Log) (*Brc20GatewayCreateNewToken, error) {
	event := new(Brc20GatewayCreateNewToken)
	if err := _Brc20Gateway.contract.UnpackLog(event, "CreateNewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayIncomingInvoiceCompletedIterator is returned from FilterIncomingInvoiceCompleted and is used to iterate over the raw logs and unpacked data for IncomingInvoiceCompleted events raised by the Brc20Gateway contract.
type Brc20GatewayIncomingInvoiceCompletedIterator struct {
	Event *Brc20GatewayIncomingInvoiceCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayIncomingInvoiceCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayIncomingInvoiceCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayIncomingInvoiceCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayIncomingInvoiceCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayIncomingInvoiceCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayIncomingInvoiceCompleted represents a IncomingInvoiceCompleted event raised by the Brc20Gateway contract.
type Brc20GatewayIncomingInvoiceCompleted struct {
	InvoiceId *big.Int
	Status    uint8
	Executor  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncomingInvoiceCompleted is a free log retrieval operation binding the contract event 0x0040b1b3943ee0be2d706b8aaa4143a310c038463192e0d26b446e7edff7f80f.
//
// Solidity: event IncomingInvoiceCompleted(uint256 indexed invoiceId, uint8 status, address executor)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterIncomingInvoiceCompleted(opts *bind.FilterOpts, invoiceId []*big.Int) (*Brc20GatewayIncomingInvoiceCompletedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "IncomingInvoiceCompleted", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayIncomingInvoiceCompletedIterator{contract: _Brc20Gateway.contract, event: "IncomingInvoiceCompleted", logs: logs, sub: sub}, nil
}

// WatchIncomingInvoiceCompleted is a free log subscription operation binding the contract event 0x0040b1b3943ee0be2d706b8aaa4143a310c038463192e0d26b446e7edff7f80f.
//
// Solidity: event IncomingInvoiceCompleted(uint256 indexed invoiceId, uint8 status, address executor)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchIncomingInvoiceCompleted(opts *bind.WatchOpts, sink chan<- *Brc20GatewayIncomingInvoiceCompleted, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "IncomingInvoiceCompleted", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayIncomingInvoiceCompleted)
				if err := _Brc20Gateway.contract.UnpackLog(event, "IncomingInvoiceCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncomingInvoiceCompleted is a log parse operation binding the contract event 0x0040b1b3943ee0be2d706b8aaa4143a310c038463192e0d26b446e7edff7f80f.
//
// Solidity: event IncomingInvoiceCompleted(uint256 indexed invoiceId, uint8 status, address executor)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseIncomingInvoiceCompleted(log types.Log) (*Brc20GatewayIncomingInvoiceCompleted, error) {
	event := new(Brc20GatewayIncomingInvoiceCompleted)
	if err := _Brc20Gateway.contract.UnpackLog(event, "IncomingInvoiceCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayIncomingInvoiceCreatedIterator is returned from FilterIncomingInvoiceCreated and is used to iterate over the raw logs and unpacked data for IncomingInvoiceCreated events raised by the Brc20Gateway contract.
type Brc20GatewayIncomingInvoiceCreatedIterator struct {
	Event *Brc20GatewayIncomingInvoiceCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayIncomingInvoiceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayIncomingInvoiceCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayIncomingInvoiceCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayIncomingInvoiceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayIncomingInvoiceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayIncomingInvoiceCreated represents a IncomingInvoiceCreated event raised by the Brc20Gateway contract.
type Brc20GatewayIncomingInvoiceCreated struct {
	InvoiceId *big.Int
	Utxo      string
	TokenId   string
	Amount    *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncomingInvoiceCreated is a free log retrieval operation binding the contract event 0xc753bd85896668260462902523612f8ea802e6dfeb9de1b9824bdeeb35f18b4c.
//
// Solidity: event IncomingInvoiceCreated(uint256 indexed invoiceId, string utxo, string tokenId, uint256 amount, address recipient)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterIncomingInvoiceCreated(opts *bind.FilterOpts, invoiceId []*big.Int) (*Brc20GatewayIncomingInvoiceCreatedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "IncomingInvoiceCreated", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayIncomingInvoiceCreatedIterator{contract: _Brc20Gateway.contract, event: "IncomingInvoiceCreated", logs: logs, sub: sub}, nil
}

// WatchIncomingInvoiceCreated is a free log subscription operation binding the contract event 0xc753bd85896668260462902523612f8ea802e6dfeb9de1b9824bdeeb35f18b4c.
//
// Solidity: event IncomingInvoiceCreated(uint256 indexed invoiceId, string utxo, string tokenId, uint256 amount, address recipient)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchIncomingInvoiceCreated(opts *bind.WatchOpts, sink chan<- *Brc20GatewayIncomingInvoiceCreated, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "IncomingInvoiceCreated", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayIncomingInvoiceCreated)
				if err := _Brc20Gateway.contract.UnpackLog(event, "IncomingInvoiceCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncomingInvoiceCreated is a log parse operation binding the contract event 0xc753bd85896668260462902523612f8ea802e6dfeb9de1b9824bdeeb35f18b4c.
//
// Solidity: event IncomingInvoiceCreated(uint256 indexed invoiceId, string utxo, string tokenId, uint256 amount, address recipient)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseIncomingInvoiceCreated(log types.Log) (*Brc20GatewayIncomingInvoiceCreated, error) {
	event := new(Brc20GatewayIncomingInvoiceCreated)
	if err := _Brc20Gateway.contract.UnpackLog(event, "IncomingInvoiceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayIncomingInvoiceRefundedIterator is returned from FilterIncomingInvoiceRefunded and is used to iterate over the raw logs and unpacked data for IncomingInvoiceRefunded events raised by the Brc20Gateway contract.
type Brc20GatewayIncomingInvoiceRefundedIterator struct {
	Event *Brc20GatewayIncomingInvoiceRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayIncomingInvoiceRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayIncomingInvoiceRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayIncomingInvoiceRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayIncomingInvoiceRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayIncomingInvoiceRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayIncomingInvoiceRefunded represents a IncomingInvoiceRefunded event raised by the Brc20Gateway contract.
type Brc20GatewayIncomingInvoiceRefunded struct {
	InvoiceId  *big.Int
	UtxoRefund string
	Executor   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterIncomingInvoiceRefunded is a free log retrieval operation binding the contract event 0x5d71f72ffb55a011fed5421c0fe881100f1524d6dcee520bd3a2b83695fdc2fa.
//
// Solidity: event IncomingInvoiceRefunded(uint256 indexed invoiceId, string utxo_refund, address executor)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterIncomingInvoiceRefunded(opts *bind.FilterOpts, invoiceId []*big.Int) (*Brc20GatewayIncomingInvoiceRefundedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "IncomingInvoiceRefunded", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayIncomingInvoiceRefundedIterator{contract: _Brc20Gateway.contract, event: "IncomingInvoiceRefunded", logs: logs, sub: sub}, nil
}

// WatchIncomingInvoiceRefunded is a free log subscription operation binding the contract event 0x5d71f72ffb55a011fed5421c0fe881100f1524d6dcee520bd3a2b83695fdc2fa.
//
// Solidity: event IncomingInvoiceRefunded(uint256 indexed invoiceId, string utxo_refund, address executor)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchIncomingInvoiceRefunded(opts *bind.WatchOpts, sink chan<- *Brc20GatewayIncomingInvoiceRefunded, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "IncomingInvoiceRefunded", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayIncomingInvoiceRefunded)
				if err := _Brc20Gateway.contract.UnpackLog(event, "IncomingInvoiceRefunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncomingInvoiceRefunded is a log parse operation binding the contract event 0x5d71f72ffb55a011fed5421c0fe881100f1524d6dcee520bd3a2b83695fdc2fa.
//
// Solidity: event IncomingInvoiceRefunded(uint256 indexed invoiceId, string utxo_refund, address executor)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseIncomingInvoiceRefunded(log types.Log) (*Brc20GatewayIncomingInvoiceRefunded, error) {
	event := new(Brc20GatewayIncomingInvoiceRefunded)
	if err := _Brc20Gateway.contract.UnpackLog(event, "IncomingInvoiceRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayIncomingInvoiceRejectedIterator is returned from FilterIncomingInvoiceRejected and is used to iterate over the raw logs and unpacked data for IncomingInvoiceRejected events raised by the Brc20Gateway contract.
type Brc20GatewayIncomingInvoiceRejectedIterator struct {
	Event *Brc20GatewayIncomingInvoiceRejected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayIncomingInvoiceRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayIncomingInvoiceRejected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayIncomingInvoiceRejected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayIncomingInvoiceRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayIncomingInvoiceRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayIncomingInvoiceRejected represents a IncomingInvoiceRejected event raised by the Brc20Gateway contract.
type Brc20GatewayIncomingInvoiceRejected struct {
	InvoiceId *big.Int
	Status    uint8
	Executor  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncomingInvoiceRejected is a free log retrieval operation binding the contract event 0x7da96e6608f6e569e80aaef406ea92a86644ee8234e16f847c93880a0d8c16ff.
//
// Solidity: event IncomingInvoiceRejected(uint256 indexed invoiceId, uint8 status, address executor)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterIncomingInvoiceRejected(opts *bind.FilterOpts, invoiceId []*big.Int) (*Brc20GatewayIncomingInvoiceRejectedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "IncomingInvoiceRejected", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayIncomingInvoiceRejectedIterator{contract: _Brc20Gateway.contract, event: "IncomingInvoiceRejected", logs: logs, sub: sub}, nil
}

// WatchIncomingInvoiceRejected is a free log subscription operation binding the contract event 0x7da96e6608f6e569e80aaef406ea92a86644ee8234e16f847c93880a0d8c16ff.
//
// Solidity: event IncomingInvoiceRejected(uint256 indexed invoiceId, uint8 status, address executor)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchIncomingInvoiceRejected(opts *bind.WatchOpts, sink chan<- *Brc20GatewayIncomingInvoiceRejected, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "IncomingInvoiceRejected", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayIncomingInvoiceRejected)
				if err := _Brc20Gateway.contract.UnpackLog(event, "IncomingInvoiceRejected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncomingInvoiceRejected is a log parse operation binding the contract event 0x7da96e6608f6e569e80aaef406ea92a86644ee8234e16f847c93880a0d8c16ff.
//
// Solidity: event IncomingInvoiceRejected(uint256 indexed invoiceId, uint8 status, address executor)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseIncomingInvoiceRejected(log types.Log) (*Brc20GatewayIncomingInvoiceRejected, error) {
	event := new(Brc20GatewayIncomingInvoiceRejected)
	if err := _Brc20Gateway.contract.UnpackLog(event, "IncomingInvoiceRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayIncomingInvoiceVerifiedIterator is returned from FilterIncomingInvoiceVerified and is used to iterate over the raw logs and unpacked data for IncomingInvoiceVerified events raised by the Brc20Gateway contract.
type Brc20GatewayIncomingInvoiceVerifiedIterator struct {
	Event *Brc20GatewayIncomingInvoiceVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayIncomingInvoiceVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayIncomingInvoiceVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayIncomingInvoiceVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayIncomingInvoiceVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayIncomingInvoiceVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayIncomingInvoiceVerified represents a IncomingInvoiceVerified event raised by the Brc20Gateway contract.
type Brc20GatewayIncomingInvoiceVerified struct {
	InvoiceId *big.Int
	Validator common.Address
	Verified  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncomingInvoiceVerified is a free log retrieval operation binding the contract event 0x2d88fccf7a5b293b2b7dee2f4671af10b2df68f488bc92af96fea710ce36a756.
//
// Solidity: event IncomingInvoiceVerified(uint256 indexed invoiceId, address validator, bool verified)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterIncomingInvoiceVerified(opts *bind.FilterOpts, invoiceId []*big.Int) (*Brc20GatewayIncomingInvoiceVerifiedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "IncomingInvoiceVerified", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayIncomingInvoiceVerifiedIterator{contract: _Brc20Gateway.contract, event: "IncomingInvoiceVerified", logs: logs, sub: sub}, nil
}

// WatchIncomingInvoiceVerified is a free log subscription operation binding the contract event 0x2d88fccf7a5b293b2b7dee2f4671af10b2df68f488bc92af96fea710ce36a756.
//
// Solidity: event IncomingInvoiceVerified(uint256 indexed invoiceId, address validator, bool verified)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchIncomingInvoiceVerified(opts *bind.WatchOpts, sink chan<- *Brc20GatewayIncomingInvoiceVerified, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "IncomingInvoiceVerified", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayIncomingInvoiceVerified)
				if err := _Brc20Gateway.contract.UnpackLog(event, "IncomingInvoiceVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncomingInvoiceVerified is a log parse operation binding the contract event 0x2d88fccf7a5b293b2b7dee2f4671af10b2df68f488bc92af96fea710ce36a756.
//
// Solidity: event IncomingInvoiceVerified(uint256 indexed invoiceId, address validator, bool verified)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseIncomingInvoiceVerified(log types.Log) (*Brc20GatewayIncomingInvoiceVerified, error) {
	event := new(Brc20GatewayIncomingInvoiceVerified)
	if err := _Brc20Gateway.contract.UnpackLog(event, "IncomingInvoiceVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Brc20Gateway contract.
type Brc20GatewayInitializedIterator struct {
	Event *Brc20GatewayInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayInitialized represents a Initialized event raised by the Brc20Gateway contract.
type Brc20GatewayInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*Brc20GatewayInitializedIterator, error) {

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayInitializedIterator{contract: _Brc20Gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Brc20GatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayInitialized)
				if err := _Brc20Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseInitialized(log types.Log) (*Brc20GatewayInitialized, error) {
	event := new(Brc20GatewayInitialized)
	if err := _Brc20Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayOutgoingInvoiceCreatedIterator is returned from FilterOutgoingInvoiceCreated and is used to iterate over the raw logs and unpacked data for OutgoingInvoiceCreated events raised by the Brc20Gateway contract.
type Brc20GatewayOutgoingInvoiceCreatedIterator struct {
	Event *Brc20GatewayOutgoingInvoiceCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayOutgoingInvoiceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayOutgoingInvoiceCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayOutgoingInvoiceCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayOutgoingInvoiceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayOutgoingInvoiceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayOutgoingInvoiceCreated represents a OutgoingInvoiceCreated event raised by the Brc20Gateway contract.
type Brc20GatewayOutgoingInvoiceCreated struct {
	InvoiceId     *big.Int
	TokenContract common.Address
	Amount        *big.Int
	From          common.Address
	Recipient     string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOutgoingInvoiceCreated is a free log retrieval operation binding the contract event 0x9fea7cb33a49dcdd816a69e66384544ccd8367e24b703f6bf63710de3a92ab41.
//
// Solidity: event OutgoingInvoiceCreated(uint256 indexed invoiceId, address tokenContract, uint256 amount, address from, string recipient)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterOutgoingInvoiceCreated(opts *bind.FilterOpts, invoiceId []*big.Int) (*Brc20GatewayOutgoingInvoiceCreatedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "OutgoingInvoiceCreated", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayOutgoingInvoiceCreatedIterator{contract: _Brc20Gateway.contract, event: "OutgoingInvoiceCreated", logs: logs, sub: sub}, nil
}

// WatchOutgoingInvoiceCreated is a free log subscription operation binding the contract event 0x9fea7cb33a49dcdd816a69e66384544ccd8367e24b703f6bf63710de3a92ab41.
//
// Solidity: event OutgoingInvoiceCreated(uint256 indexed invoiceId, address tokenContract, uint256 amount, address from, string recipient)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchOutgoingInvoiceCreated(opts *bind.WatchOpts, sink chan<- *Brc20GatewayOutgoingInvoiceCreated, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "OutgoingInvoiceCreated", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayOutgoingInvoiceCreated)
				if err := _Brc20Gateway.contract.UnpackLog(event, "OutgoingInvoiceCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOutgoingInvoiceCreated is a log parse operation binding the contract event 0x9fea7cb33a49dcdd816a69e66384544ccd8367e24b703f6bf63710de3a92ab41.
//
// Solidity: event OutgoingInvoiceCreated(uint256 indexed invoiceId, address tokenContract, uint256 amount, address from, string recipient)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseOutgoingInvoiceCreated(log types.Log) (*Brc20GatewayOutgoingInvoiceCreated, error) {
	event := new(Brc20GatewayOutgoingInvoiceCreated)
	if err := _Brc20Gateway.contract.UnpackLog(event, "OutgoingInvoiceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayOutgoingTxProcessedIterator is returned from FilterOutgoingTxProcessed and is used to iterate over the raw logs and unpacked data for OutgoingTxProcessed events raised by the Brc20Gateway contract.
type Brc20GatewayOutgoingTxProcessedIterator struct {
	Event *Brc20GatewayOutgoingTxProcessed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayOutgoingTxProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayOutgoingTxProcessed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayOutgoingTxProcessed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayOutgoingTxProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayOutgoingTxProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayOutgoingTxProcessed represents a OutgoingTxProcessed event raised by the Brc20Gateway contract.
type Brc20GatewayOutgoingTxProcessed struct {
	TxId   *big.Int
	Status uint8
	TxHash string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOutgoingTxProcessed is a free log retrieval operation binding the contract event 0xbb51d23336167c8e20a9ca20fbe3f55ebd9339040e3d56a545f2527bc55d3f3b.
//
// Solidity: event OutgoingTxProcessed(uint256 indexed txId, uint8 status, string txHash)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterOutgoingTxProcessed(opts *bind.FilterOpts, txId []*big.Int) (*Brc20GatewayOutgoingTxProcessedIterator, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "OutgoingTxProcessed", txIdRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayOutgoingTxProcessedIterator{contract: _Brc20Gateway.contract, event: "OutgoingTxProcessed", logs: logs, sub: sub}, nil
}

// WatchOutgoingTxProcessed is a free log subscription operation binding the contract event 0xbb51d23336167c8e20a9ca20fbe3f55ebd9339040e3d56a545f2527bc55d3f3b.
//
// Solidity: event OutgoingTxProcessed(uint256 indexed txId, uint8 status, string txHash)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchOutgoingTxProcessed(opts *bind.WatchOpts, sink chan<- *Brc20GatewayOutgoingTxProcessed, txId []*big.Int) (event.Subscription, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "OutgoingTxProcessed", txIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayOutgoingTxProcessed)
				if err := _Brc20Gateway.contract.UnpackLog(event, "OutgoingTxProcessed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOutgoingTxProcessed is a log parse operation binding the contract event 0xbb51d23336167c8e20a9ca20fbe3f55ebd9339040e3d56a545f2527bc55d3f3b.
//
// Solidity: event OutgoingTxProcessed(uint256 indexed txId, uint8 status, string txHash)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseOutgoingTxProcessed(log types.Log) (*Brc20GatewayOutgoingTxProcessed, error) {
	event := new(Brc20GatewayOutgoingTxProcessed)
	if err := _Brc20Gateway.contract.UnpackLog(event, "OutgoingTxProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayOutgoingTxSubmittedIterator is returned from FilterOutgoingTxSubmitted and is used to iterate over the raw logs and unpacked data for OutgoingTxSubmitted events raised by the Brc20Gateway contract.
type Brc20GatewayOutgoingTxSubmittedIterator struct {
	Event *Brc20GatewayOutgoingTxSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayOutgoingTxSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayOutgoingTxSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayOutgoingTxSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayOutgoingTxSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayOutgoingTxSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayOutgoingTxSubmitted represents a OutgoingTxSubmitted event raised by the Brc20Gateway contract.
type Brc20GatewayOutgoingTxSubmitted struct {
	TxId      *big.Int
	TxContent string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutgoingTxSubmitted is a free log retrieval operation binding the contract event 0xda987441b8bab19820139d670528f4e4de7f914c8018f6ef2894e83f69e67f21.
//
// Solidity: event OutgoingTxSubmitted(uint256 indexed txId, string txContent)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterOutgoingTxSubmitted(opts *bind.FilterOpts, txId []*big.Int) (*Brc20GatewayOutgoingTxSubmittedIterator, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "OutgoingTxSubmitted", txIdRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayOutgoingTxSubmittedIterator{contract: _Brc20Gateway.contract, event: "OutgoingTxSubmitted", logs: logs, sub: sub}, nil
}

// WatchOutgoingTxSubmitted is a free log subscription operation binding the contract event 0xda987441b8bab19820139d670528f4e4de7f914c8018f6ef2894e83f69e67f21.
//
// Solidity: event OutgoingTxSubmitted(uint256 indexed txId, string txContent)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchOutgoingTxSubmitted(opts *bind.WatchOpts, sink chan<- *Brc20GatewayOutgoingTxSubmitted, txId []*big.Int) (event.Subscription, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "OutgoingTxSubmitted", txIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayOutgoingTxSubmitted)
				if err := _Brc20Gateway.contract.UnpackLog(event, "OutgoingTxSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOutgoingTxSubmitted is a log parse operation binding the contract event 0xda987441b8bab19820139d670528f4e4de7f914c8018f6ef2894e83f69e67f21.
//
// Solidity: event OutgoingTxSubmitted(uint256 indexed txId, string txContent)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseOutgoingTxSubmitted(log types.Log) (*Brc20GatewayOutgoingTxSubmitted, error) {
	event := new(Brc20GatewayOutgoingTxSubmitted)
	if err := _Brc20Gateway.contract.UnpackLog(event, "OutgoingTxSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayOutgoingTxVerifiedIterator is returned from FilterOutgoingTxVerified and is used to iterate over the raw logs and unpacked data for OutgoingTxVerified events raised by the Brc20Gateway contract.
type Brc20GatewayOutgoingTxVerifiedIterator struct {
	Event *Brc20GatewayOutgoingTxVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayOutgoingTxVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayOutgoingTxVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayOutgoingTxVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayOutgoingTxVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayOutgoingTxVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayOutgoingTxVerified represents a OutgoingTxVerified event raised by the Brc20Gateway contract.
type Brc20GatewayOutgoingTxVerified struct {
	TxId      *big.Int
	Validator common.Address
	Verified  bool
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutgoingTxVerified is a free log retrieval operation binding the contract event 0x9191d6ec9841818acc620ffc69e1390adf521156b9383a417dfe403447a7a0c7.
//
// Solidity: event OutgoingTxVerified(uint256 indexed txId, address validator, bool verified, string signature)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterOutgoingTxVerified(opts *bind.FilterOpts, txId []*big.Int) (*Brc20GatewayOutgoingTxVerifiedIterator, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "OutgoingTxVerified", txIdRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayOutgoingTxVerifiedIterator{contract: _Brc20Gateway.contract, event: "OutgoingTxVerified", logs: logs, sub: sub}, nil
}

// WatchOutgoingTxVerified is a free log subscription operation binding the contract event 0x9191d6ec9841818acc620ffc69e1390adf521156b9383a417dfe403447a7a0c7.
//
// Solidity: event OutgoingTxVerified(uint256 indexed txId, address validator, bool verified, string signature)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchOutgoingTxVerified(opts *bind.WatchOpts, sink chan<- *Brc20GatewayOutgoingTxVerified, txId []*big.Int) (event.Subscription, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "OutgoingTxVerified", txIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayOutgoingTxVerified)
				if err := _Brc20Gateway.contract.UnpackLog(event, "OutgoingTxVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOutgoingTxVerified is a log parse operation binding the contract event 0x9191d6ec9841818acc620ffc69e1390adf521156b9383a417dfe403447a7a0c7.
//
// Solidity: event OutgoingTxVerified(uint256 indexed txId, address validator, bool verified, string signature)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseOutgoingTxVerified(log types.Log) (*Brc20GatewayOutgoingTxVerified, error) {
	event := new(Brc20GatewayOutgoingTxVerified)
	if err := _Brc20Gateway.contract.UnpackLog(event, "OutgoingTxVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Brc20Gateway contract.
type Brc20GatewayOwnershipTransferredIterator struct {
	Event *Brc20GatewayOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the Brc20Gateway contract.
type Brc20GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Brc20GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayOwnershipTransferredIterator{contract: _Brc20Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Brc20GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayOwnershipTransferred)
				if err := _Brc20Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*Brc20GatewayOwnershipTransferred, error) {
	event := new(Brc20GatewayOwnershipTransferred)
	if err := _Brc20Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Brc20Gateway contract.
type Brc20GatewayPausedIterator struct {
	Event *Brc20GatewayPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayPaused represents a Paused event raised by the Brc20Gateway contract.
type Brc20GatewayPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterPaused(opts *bind.FilterOpts) (*Brc20GatewayPausedIterator, error) {

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayPausedIterator{contract: _Brc20Gateway.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Brc20GatewayPaused) (event.Subscription, error) {

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayPaused)
				if err := _Brc20Gateway.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Brc20Gateway *Brc20GatewayFilterer) ParsePaused(log types.Log) (*Brc20GatewayPaused, error) {
	event := new(Brc20GatewayPaused)
	if err := _Brc20Gateway.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Brc20GatewayUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Brc20Gateway contract.
type Brc20GatewayUnpausedIterator struct {
	Event *Brc20GatewayUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Brc20GatewayUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Brc20GatewayUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Brc20GatewayUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayUnpaused represents a Unpaused event raised by the Brc20Gateway contract.
type Brc20GatewayUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterUnpaused(opts *bind.FilterOpts) (*Brc20GatewayUnpausedIterator, error) {

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayUnpausedIterator{contract: _Brc20Gateway.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Brc20GatewayUnpaused) (event.Subscription, error) {

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayUnpaused)
				if err := _Brc20Gateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Brc20Gateway *Brc20GatewayFilterer) ParseUnpaused(log types.Log) (*Brc20GatewayUnpaused, error) {
	event := new(Brc20GatewayUnpaused)
	if err := _Brc20Gateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
