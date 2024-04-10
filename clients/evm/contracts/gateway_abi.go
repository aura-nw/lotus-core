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

// IGatewayIncomingInvoiceResponse is an auto generated low-level Go binding around an user-defined struct.
type IGatewayIncomingInvoiceResponse struct {
	InvoiceId     *big.Int
	Utxo          string
	Amount        *big.Int
	Recipient     common.Address
	Status        uint8
	Validators    []common.Address
	Confirmations []bool
	UtxoRefund    string
}

// IGatewayOutgoingInvoiceBasicInfo is an auto generated low-level Go binding around an user-defined struct.
type IGatewayOutgoingInvoiceBasicInfo struct {
	InvoiceId *big.Int
	Amount    *big.Int
	Recipient string
}

// IGatewayOutgoingInvoiceResponse is an auto generated low-level Go binding around an user-defined struct.
type IGatewayOutgoingInvoiceResponse struct {
	InvoiceId *big.Int
	From      common.Address
	Amount    *big.Int
	Recipient string
	Status    uint8
	TxId      *big.Int
}

// IGatewayOutgoingTxInfo is an auto generated low-level Go binding around an user-defined struct.
type IGatewayOutgoingTxInfo struct {
	TxId        *big.Int
	TotalAmount *big.Int
	InvoiceIds  []*big.Int
	TxContent   string
	Validators  []common.Address
	Signatures  []string
	Status      uint8
	TxHash      string
}

// IGatewayValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IGatewayValidatorInfo struct {
	Validator           common.Address
	NextIncomingInvoice *big.Int
	NextOutgoingTx      *big.Int
}

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ProposerUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ValidatorUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIGateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"IncomingInvoiceCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"utxo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"IncomingInvoiceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"utxo_refund\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"IncomingInvoiceRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIGateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"IncomingInvoiceRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"IncomingInvoiceVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"}],\"name\":\"OutgoingInvoiceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumIGateway.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"OutgoingTxProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txContent\",\"type\":\"string\"}],\"name\":\"OutgoingTxSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"}],\"name\":\"OutgoingTxVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nextIncomingInvoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextOutgoingTx\",\"type\":\"uint256\"}],\"internalType\":\"structIGateway.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"btcAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"completeIncomingInvoice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"concensusThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_utxo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"createIncomingInvoice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_utxo\",\"type\":\"string\"}],\"name\":\"incomingInvoice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"utxo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"enumIGateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"confirmations\",\"type\":\"bool[]\"},{\"internalType\":\"string\",\"name\":\"utxo_refund\",\"type\":\"string\"}],\"internalType\":\"structIGateway.IncomingInvoiceResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"incomingInvoices\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"utxo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"enumIGateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"utxo_refund\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incomingInvoicesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_invoiceId\",\"type\":\"uint256\"}],\"name\":\"outgoingInvoice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"enumIGateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"}],\"internalType\":\"structIGateway.OutgoingInvoiceResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"outgoingInvoices\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"},{\"internalType\":\"enumIGateway.InvoiceStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"txContent\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outgoingInvoicesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txId\",\"type\":\"uint256\"}],\"name\":\"outgoingTx\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"invoiceIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"txContent\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"enumIGateway.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"internalType\":\"structIGateway.OutgoingTxInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outgoingTxCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_txHash\",\"type\":\"string\"}],\"name\":\"processOutgoingTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_utxo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_utxo_refund\",\"type\":\"string\"}],\"name\":\"refundIncomingInvoice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"}],\"internalType\":\"structIGateway.OutgoingInvoiceBasicInfo[]\",\"name\":\"_invoices\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"_txContent\",\"type\":\"string\"}],\"name\":\"submitTxContent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_btcAddress\",\"type\":\"address\"}],\"name\":\"updateBtcAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_concensusThreshold\",\"type\":\"uint256\"}],\"name\":\"updateConcensusThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"}],\"name\":\"updateProposer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"utxos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"validator\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nextIncomingInvoice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextOutgoingTx\",\"type\":\"uint256\"}],\"internalType\":\"structIGateway.ValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_utxo\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isVerified\",\"type\":\"bool\"}],\"name\":\"verifyIncomingInvoice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isVerified\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_signature\",\"type\":\"string\"}],\"name\":\"verifyOutgoingTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMetaData.ABI instead.
var GatewayABI = GatewayMetaData.ABI

// Gateway is an auto generated Go binding around an Ethereum contract.
type Gateway struct {
	GatewayCaller     // Read-only binding to the contract
	GatewayTransactor // Write-only binding to the contract
	GatewayFilterer   // Log filterer for contract events
}

// GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewaySession struct {
	Contract     *Gateway          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayCallerSession struct {
	Contract *GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayTransactorSession struct {
	Contract     *GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRaw struct {
	Contract *Gateway // Generic contract binding to access the raw methods on
}

// GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayCallerRaw struct {
	Contract *GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayTransactorRaw struct {
	Contract *GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGateway creates a new instance of Gateway, bound to a specific deployed contract.
func NewGateway(address common.Address, backend bind.ContractBackend) (*Gateway, error) {
	contract, err := bindGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// NewGatewayCaller creates a new read-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayCaller(address common.Address, caller bind.ContractCaller) (*GatewayCaller, error) {
	contract, err := bindGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayCaller{contract: contract}, nil
}

// NewGatewayTransactor creates a new write-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayTransactor, error) {
	contract, err := bindGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayTransactor{contract: contract}, nil
}

// NewGatewayFilterer creates a new log filterer instance of Gateway, bound to a specific deployed contract.
func NewGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFilterer, error) {
	contract, err := bindGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFilterer{contract: contract}, nil
}

// bindGateway binds a generic wrapper to an already deployed contract.
func bindGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transact(opts, method, params...)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns((address,uint256,uint256)[])
func (_Gateway *GatewayCaller) AllValidators(opts *bind.CallOpts) ([]IGatewayValidatorInfo, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "allValidators")

	if err != nil {
		return *new([]IGatewayValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGatewayValidatorInfo)).(*[]IGatewayValidatorInfo)

	return out0, err

}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns((address,uint256,uint256)[])
func (_Gateway *GatewaySession) AllValidators() ([]IGatewayValidatorInfo, error) {
	return _Gateway.Contract.AllValidators(&_Gateway.CallOpts)
}

// AllValidators is a free data retrieval call binding the contract method 0xf30589c3.
//
// Solidity: function allValidators() view returns((address,uint256,uint256)[])
func (_Gateway *GatewayCallerSession) AllValidators() ([]IGatewayValidatorInfo, error) {
	return _Gateway.Contract.AllValidators(&_Gateway.CallOpts)
}

// BtcAddress is a free data retrieval call binding the contract method 0xf7d76ae5.
//
// Solidity: function btcAddress() view returns(address)
func (_Gateway *GatewayCaller) BtcAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "btcAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BtcAddress is a free data retrieval call binding the contract method 0xf7d76ae5.
//
// Solidity: function btcAddress() view returns(address)
func (_Gateway *GatewaySession) BtcAddress() (common.Address, error) {
	return _Gateway.Contract.BtcAddress(&_Gateway.CallOpts)
}

// BtcAddress is a free data retrieval call binding the contract method 0xf7d76ae5.
//
// Solidity: function btcAddress() view returns(address)
func (_Gateway *GatewayCallerSession) BtcAddress() (common.Address, error) {
	return _Gateway.Contract.BtcAddress(&_Gateway.CallOpts)
}

// ConcensusThreshold is a free data retrieval call binding the contract method 0x99bcf2dd.
//
// Solidity: function concensusThreshold() view returns(uint256)
func (_Gateway *GatewayCaller) ConcensusThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "concensusThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConcensusThreshold is a free data retrieval call binding the contract method 0x99bcf2dd.
//
// Solidity: function concensusThreshold() view returns(uint256)
func (_Gateway *GatewaySession) ConcensusThreshold() (*big.Int, error) {
	return _Gateway.Contract.ConcensusThreshold(&_Gateway.CallOpts)
}

// ConcensusThreshold is a free data retrieval call binding the contract method 0x99bcf2dd.
//
// Solidity: function concensusThreshold() view returns(uint256)
func (_Gateway *GatewayCallerSession) ConcensusThreshold() (*big.Int, error) {
	return _Gateway.Contract.ConcensusThreshold(&_Gateway.CallOpts)
}

// IncomingInvoice is a free data retrieval call binding the contract method 0xe7c777eb.
//
// Solidity: function incomingInvoice(string _utxo) view returns((uint256,string,uint256,address,uint8,address[],bool[],string))
func (_Gateway *GatewayCaller) IncomingInvoice(opts *bind.CallOpts, _utxo string) (IGatewayIncomingInvoiceResponse, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "incomingInvoice", _utxo)

	if err != nil {
		return *new(IGatewayIncomingInvoiceResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(IGatewayIncomingInvoiceResponse)).(*IGatewayIncomingInvoiceResponse)

	return out0, err

}

// IncomingInvoice is a free data retrieval call binding the contract method 0xe7c777eb.
//
// Solidity: function incomingInvoice(string _utxo) view returns((uint256,string,uint256,address,uint8,address[],bool[],string))
func (_Gateway *GatewaySession) IncomingInvoice(_utxo string) (IGatewayIncomingInvoiceResponse, error) {
	return _Gateway.Contract.IncomingInvoice(&_Gateway.CallOpts, _utxo)
}

// IncomingInvoice is a free data retrieval call binding the contract method 0xe7c777eb.
//
// Solidity: function incomingInvoice(string _utxo) view returns((uint256,string,uint256,address,uint8,address[],bool[],string))
func (_Gateway *GatewayCallerSession) IncomingInvoice(_utxo string) (IGatewayIncomingInvoiceResponse, error) {
	return _Gateway.Contract.IncomingInvoice(&_Gateway.CallOpts, _utxo)
}

// IncomingInvoices is a free data retrieval call binding the contract method 0xf5d730c8.
//
// Solidity: function incomingInvoices(uint256 ) view returns(string utxo, uint256 amount, address recipient, uint8 status, string utxo_refund)
func (_Gateway *GatewayCaller) IncomingInvoices(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Utxo       string
	Amount     *big.Int
	Recipient  common.Address
	Status     uint8
	UtxoRefund string
}, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "incomingInvoices", arg0)

	outstruct := new(struct {
		Utxo       string
		Amount     *big.Int
		Recipient  common.Address
		Status     uint8
		UtxoRefund string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Utxo = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Recipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.UtxoRefund = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// IncomingInvoices is a free data retrieval call binding the contract method 0xf5d730c8.
//
// Solidity: function incomingInvoices(uint256 ) view returns(string utxo, uint256 amount, address recipient, uint8 status, string utxo_refund)
func (_Gateway *GatewaySession) IncomingInvoices(arg0 *big.Int) (struct {
	Utxo       string
	Amount     *big.Int
	Recipient  common.Address
	Status     uint8
	UtxoRefund string
}, error) {
	return _Gateway.Contract.IncomingInvoices(&_Gateway.CallOpts, arg0)
}

// IncomingInvoices is a free data retrieval call binding the contract method 0xf5d730c8.
//
// Solidity: function incomingInvoices(uint256 ) view returns(string utxo, uint256 amount, address recipient, uint8 status, string utxo_refund)
func (_Gateway *GatewayCallerSession) IncomingInvoices(arg0 *big.Int) (struct {
	Utxo       string
	Amount     *big.Int
	Recipient  common.Address
	Status     uint8
	UtxoRefund string
}, error) {
	return _Gateway.Contract.IncomingInvoices(&_Gateway.CallOpts, arg0)
}

// IncomingInvoicesCount is a free data retrieval call binding the contract method 0xd6988f61.
//
// Solidity: function incomingInvoicesCount() view returns(uint256)
func (_Gateway *GatewayCaller) IncomingInvoicesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "incomingInvoicesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncomingInvoicesCount is a free data retrieval call binding the contract method 0xd6988f61.
//
// Solidity: function incomingInvoicesCount() view returns(uint256)
func (_Gateway *GatewaySession) IncomingInvoicesCount() (*big.Int, error) {
	return _Gateway.Contract.IncomingInvoicesCount(&_Gateway.CallOpts)
}

// IncomingInvoicesCount is a free data retrieval call binding the contract method 0xd6988f61.
//
// Solidity: function incomingInvoicesCount() view returns(uint256)
func (_Gateway *GatewayCallerSession) IncomingInvoicesCount() (*big.Int, error) {
	return _Gateway.Contract.IncomingInvoicesCount(&_Gateway.CallOpts)
}

// OutgoingInvoice is a free data retrieval call binding the contract method 0x455578b3.
//
// Solidity: function outgoingInvoice(uint256 _invoiceId) view returns((uint256,address,uint256,string,uint8,uint256))
func (_Gateway *GatewayCaller) OutgoingInvoice(opts *bind.CallOpts, _invoiceId *big.Int) (IGatewayOutgoingInvoiceResponse, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "outgoingInvoice", _invoiceId)

	if err != nil {
		return *new(IGatewayOutgoingInvoiceResponse), err
	}

	out0 := *abi.ConvertType(out[0], new(IGatewayOutgoingInvoiceResponse)).(*IGatewayOutgoingInvoiceResponse)

	return out0, err

}

// OutgoingInvoice is a free data retrieval call binding the contract method 0x455578b3.
//
// Solidity: function outgoingInvoice(uint256 _invoiceId) view returns((uint256,address,uint256,string,uint8,uint256))
func (_Gateway *GatewaySession) OutgoingInvoice(_invoiceId *big.Int) (IGatewayOutgoingInvoiceResponse, error) {
	return _Gateway.Contract.OutgoingInvoice(&_Gateway.CallOpts, _invoiceId)
}

// OutgoingInvoice is a free data retrieval call binding the contract method 0x455578b3.
//
// Solidity: function outgoingInvoice(uint256 _invoiceId) view returns((uint256,address,uint256,string,uint8,uint256))
func (_Gateway *GatewayCallerSession) OutgoingInvoice(_invoiceId *big.Int) (IGatewayOutgoingInvoiceResponse, error) {
	return _Gateway.Contract.OutgoingInvoice(&_Gateway.CallOpts, _invoiceId)
}

// OutgoingInvoices is a free data retrieval call binding the contract method 0xb576bab1.
//
// Solidity: function outgoingInvoices(uint256 ) view returns(address from, uint256 amount, string recipient, uint8 status, string txContent, uint256 txId)
func (_Gateway *GatewayCaller) OutgoingInvoices(opts *bind.CallOpts, arg0 *big.Int) (struct {
	From      common.Address
	Amount    *big.Int
	Recipient string
	Status    uint8
	TxContent string
	TxId      *big.Int
}, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "outgoingInvoices", arg0)

	outstruct := new(struct {
		From      common.Address
		Amount    *big.Int
		Recipient string
		Status    uint8
		TxContent string
		TxId      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.From = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Recipient = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.TxContent = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.TxId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OutgoingInvoices is a free data retrieval call binding the contract method 0xb576bab1.
//
// Solidity: function outgoingInvoices(uint256 ) view returns(address from, uint256 amount, string recipient, uint8 status, string txContent, uint256 txId)
func (_Gateway *GatewaySession) OutgoingInvoices(arg0 *big.Int) (struct {
	From      common.Address
	Amount    *big.Int
	Recipient string
	Status    uint8
	TxContent string
	TxId      *big.Int
}, error) {
	return _Gateway.Contract.OutgoingInvoices(&_Gateway.CallOpts, arg0)
}

// OutgoingInvoices is a free data retrieval call binding the contract method 0xb576bab1.
//
// Solidity: function outgoingInvoices(uint256 ) view returns(address from, uint256 amount, string recipient, uint8 status, string txContent, uint256 txId)
func (_Gateway *GatewayCallerSession) OutgoingInvoices(arg0 *big.Int) (struct {
	From      common.Address
	Amount    *big.Int
	Recipient string
	Status    uint8
	TxContent string
	TxId      *big.Int
}, error) {
	return _Gateway.Contract.OutgoingInvoices(&_Gateway.CallOpts, arg0)
}

// OutgoingInvoicesCount is a free data retrieval call binding the contract method 0xc8c52a6a.
//
// Solidity: function outgoingInvoicesCount() view returns(uint256)
func (_Gateway *GatewayCaller) OutgoingInvoicesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "outgoingInvoicesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutgoingInvoicesCount is a free data retrieval call binding the contract method 0xc8c52a6a.
//
// Solidity: function outgoingInvoicesCount() view returns(uint256)
func (_Gateway *GatewaySession) OutgoingInvoicesCount() (*big.Int, error) {
	return _Gateway.Contract.OutgoingInvoicesCount(&_Gateway.CallOpts)
}

// OutgoingInvoicesCount is a free data retrieval call binding the contract method 0xc8c52a6a.
//
// Solidity: function outgoingInvoicesCount() view returns(uint256)
func (_Gateway *GatewayCallerSession) OutgoingInvoicesCount() (*big.Int, error) {
	return _Gateway.Contract.OutgoingInvoicesCount(&_Gateway.CallOpts)
}

// OutgoingTx is a free data retrieval call binding the contract method 0x2823434e.
//
// Solidity: function outgoingTx(uint256 _txId) view returns((uint256,uint256,uint256[],string,address[],string[],uint8,string))
func (_Gateway *GatewayCaller) OutgoingTx(opts *bind.CallOpts, _txId *big.Int) (IGatewayOutgoingTxInfo, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "outgoingTx", _txId)

	if err != nil {
		return *new(IGatewayOutgoingTxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IGatewayOutgoingTxInfo)).(*IGatewayOutgoingTxInfo)

	return out0, err

}

// OutgoingTx is a free data retrieval call binding the contract method 0x2823434e.
//
// Solidity: function outgoingTx(uint256 _txId) view returns((uint256,uint256,uint256[],string,address[],string[],uint8,string))
func (_Gateway *GatewaySession) OutgoingTx(_txId *big.Int) (IGatewayOutgoingTxInfo, error) {
	return _Gateway.Contract.OutgoingTx(&_Gateway.CallOpts, _txId)
}

// OutgoingTx is a free data retrieval call binding the contract method 0x2823434e.
//
// Solidity: function outgoingTx(uint256 _txId) view returns((uint256,uint256,uint256[],string,address[],string[],uint8,string))
func (_Gateway *GatewayCallerSession) OutgoingTx(_txId *big.Int) (IGatewayOutgoingTxInfo, error) {
	return _Gateway.Contract.OutgoingTx(&_Gateway.CallOpts, _txId)
}

// OutgoingTxCount is a free data retrieval call binding the contract method 0x56770381.
//
// Solidity: function outgoingTxCount() view returns(uint256)
func (_Gateway *GatewayCaller) OutgoingTxCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "outgoingTxCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutgoingTxCount is a free data retrieval call binding the contract method 0x56770381.
//
// Solidity: function outgoingTxCount() view returns(uint256)
func (_Gateway *GatewaySession) OutgoingTxCount() (*big.Int, error) {
	return _Gateway.Contract.OutgoingTxCount(&_Gateway.CallOpts)
}

// OutgoingTxCount is a free data retrieval call binding the contract method 0x56770381.
//
// Solidity: function outgoingTxCount() view returns(uint256)
func (_Gateway *GatewayCallerSession) OutgoingTxCount() (*big.Int, error) {
	return _Gateway.Contract.OutgoingTxCount(&_Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewaySession) Owner() (common.Address, error) {
	return _Gateway.Contract.Owner(&_Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewayCallerSession) Owner() (common.Address, error) {
	return _Gateway.Contract.Owner(&_Gateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gateway *GatewayCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gateway *GatewaySession) Paused() (bool, error) {
	return _Gateway.Contract.Paused(&_Gateway.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Gateway *GatewayCallerSession) Paused() (bool, error) {
	return _Gateway.Contract.Paused(&_Gateway.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Gateway *GatewayCaller) Proposer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "proposer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Gateway *GatewaySession) Proposer() (common.Address, error) {
	return _Gateway.Contract.Proposer(&_Gateway.CallOpts)
}

// Proposer is a free data retrieval call binding the contract method 0xa8e4fb90.
//
// Solidity: function proposer() view returns(address)
func (_Gateway *GatewayCallerSession) Proposer() (common.Address, error) {
	return _Gateway.Contract.Proposer(&_Gateway.CallOpts)
}

// Utxos is a free data retrieval call binding the contract method 0xf1fbbee0.
//
// Solidity: function utxos(string ) view returns(uint256)
func (_Gateway *GatewayCaller) Utxos(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "utxos", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Utxos is a free data retrieval call binding the contract method 0xf1fbbee0.
//
// Solidity: function utxos(string ) view returns(uint256)
func (_Gateway *GatewaySession) Utxos(arg0 string) (*big.Int, error) {
	return _Gateway.Contract.Utxos(&_Gateway.CallOpts, arg0)
}

// Utxos is a free data retrieval call binding the contract method 0xf1fbbee0.
//
// Solidity: function utxos(string ) view returns(uint256)
func (_Gateway *GatewayCallerSession) Utxos(arg0 string) (*big.Int, error) {
	return _Gateway.Contract.Utxos(&_Gateway.CallOpts, arg0)
}

// Validator is a free data retrieval call binding the contract method 0x223b3b7a.
//
// Solidity: function validator(address _validator) view returns((address,uint256,uint256))
func (_Gateway *GatewayCaller) Validator(opts *bind.CallOpts, _validator common.Address) (IGatewayValidatorInfo, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "validator", _validator)

	if err != nil {
		return *new(IGatewayValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IGatewayValidatorInfo)).(*IGatewayValidatorInfo)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x223b3b7a.
//
// Solidity: function validator(address _validator) view returns((address,uint256,uint256))
func (_Gateway *GatewaySession) Validator(_validator common.Address) (IGatewayValidatorInfo, error) {
	return _Gateway.Contract.Validator(&_Gateway.CallOpts, _validator)
}

// Validator is a free data retrieval call binding the contract method 0x223b3b7a.
//
// Solidity: function validator(address _validator) view returns((address,uint256,uint256))
func (_Gateway *GatewayCallerSession) Validator(_validator common.Address) (IGatewayValidatorInfo, error) {
	return _Gateway.Contract.Validator(&_Gateway.CallOpts, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_Gateway *GatewayTransactor) AddValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "addValidator", _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_Gateway *GatewaySession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.AddValidator(&_Gateway.TransactOpts, _validator)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validator) returns()
func (_Gateway *GatewayTransactorSession) AddValidator(_validator common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.AddValidator(&_Gateway.TransactOpts, _validator)
}

// CompleteIncomingInvoice is a paid mutator transaction binding the contract method 0xa0522b9d.
//
// Solidity: function completeIncomingInvoice(uint256 _invoiceId, address _recipient) returns(uint256)
func (_Gateway *GatewayTransactor) CompleteIncomingInvoice(opts *bind.TransactOpts, _invoiceId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "completeIncomingInvoice", _invoiceId, _recipient)
}

// CompleteIncomingInvoice is a paid mutator transaction binding the contract method 0xa0522b9d.
//
// Solidity: function completeIncomingInvoice(uint256 _invoiceId, address _recipient) returns(uint256)
func (_Gateway *GatewaySession) CompleteIncomingInvoice(_invoiceId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.CompleteIncomingInvoice(&_Gateway.TransactOpts, _invoiceId, _recipient)
}

// CompleteIncomingInvoice is a paid mutator transaction binding the contract method 0xa0522b9d.
//
// Solidity: function completeIncomingInvoice(uint256 _invoiceId, address _recipient) returns(uint256)
func (_Gateway *GatewayTransactorSession) CompleteIncomingInvoice(_invoiceId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.CompleteIncomingInvoice(&_Gateway.TransactOpts, _invoiceId, _recipient)
}

// CreateIncomingInvoice is a paid mutator transaction binding the contract method 0x6d5d6daf.
//
// Solidity: function createIncomingInvoice(string _utxo, uint256 _amount, address _recipient) returns()
func (_Gateway *GatewayTransactor) CreateIncomingInvoice(opts *bind.TransactOpts, _utxo string, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "createIncomingInvoice", _utxo, _amount, _recipient)
}

// CreateIncomingInvoice is a paid mutator transaction binding the contract method 0x6d5d6daf.
//
// Solidity: function createIncomingInvoice(string _utxo, uint256 _amount, address _recipient) returns()
func (_Gateway *GatewaySession) CreateIncomingInvoice(_utxo string, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.CreateIncomingInvoice(&_Gateway.TransactOpts, _utxo, _amount, _recipient)
}

// CreateIncomingInvoice is a paid mutator transaction binding the contract method 0x6d5d6daf.
//
// Solidity: function createIncomingInvoice(string _utxo, uint256 _amount, address _recipient) returns()
func (_Gateway *GatewayTransactorSession) CreateIncomingInvoice(_utxo string, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.CreateIncomingInvoice(&_Gateway.TransactOpts, _utxo, _amount, _recipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Gateway *GatewayTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Gateway *GatewaySession) Initialize() (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Gateway *GatewayTransactorSession) Initialize() (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gateway *GatewayTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gateway *GatewaySession) Pause() (*types.Transaction, error) {
	return _Gateway.Contract.Pause(&_Gateway.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Gateway *GatewayTransactorSession) Pause() (*types.Transaction, error) {
	return _Gateway.Contract.Pause(&_Gateway.TransactOpts)
}

// ProcessOutgoingTx is a paid mutator transaction binding the contract method 0x0f6dc9e2.
//
// Solidity: function processOutgoingTx(uint256 _txId, string _txHash) returns()
func (_Gateway *GatewayTransactor) ProcessOutgoingTx(opts *bind.TransactOpts, _txId *big.Int, _txHash string) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "processOutgoingTx", _txId, _txHash)
}

// ProcessOutgoingTx is a paid mutator transaction binding the contract method 0x0f6dc9e2.
//
// Solidity: function processOutgoingTx(uint256 _txId, string _txHash) returns()
func (_Gateway *GatewaySession) ProcessOutgoingTx(_txId *big.Int, _txHash string) (*types.Transaction, error) {
	return _Gateway.Contract.ProcessOutgoingTx(&_Gateway.TransactOpts, _txId, _txHash)
}

// ProcessOutgoingTx is a paid mutator transaction binding the contract method 0x0f6dc9e2.
//
// Solidity: function processOutgoingTx(uint256 _txId, string _txHash) returns()
func (_Gateway *GatewayTransactorSession) ProcessOutgoingTx(_txId *big.Int, _txHash string) (*types.Transaction, error) {
	return _Gateway.Contract.ProcessOutgoingTx(&_Gateway.TransactOpts, _txId, _txHash)
}

// RefundIncomingInvoice is a paid mutator transaction binding the contract method 0xdc8cc82f.
//
// Solidity: function refundIncomingInvoice(uint256 _invoiceId, string _utxo, uint256 _amount, address _recipient, string _utxo_refund) returns()
func (_Gateway *GatewayTransactor) RefundIncomingInvoice(opts *bind.TransactOpts, _invoiceId *big.Int, _utxo string, _amount *big.Int, _recipient common.Address, _utxo_refund string) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "refundIncomingInvoice", _invoiceId, _utxo, _amount, _recipient, _utxo_refund)
}

// RefundIncomingInvoice is a paid mutator transaction binding the contract method 0xdc8cc82f.
//
// Solidity: function refundIncomingInvoice(uint256 _invoiceId, string _utxo, uint256 _amount, address _recipient, string _utxo_refund) returns()
func (_Gateway *GatewaySession) RefundIncomingInvoice(_invoiceId *big.Int, _utxo string, _amount *big.Int, _recipient common.Address, _utxo_refund string) (*types.Transaction, error) {
	return _Gateway.Contract.RefundIncomingInvoice(&_Gateway.TransactOpts, _invoiceId, _utxo, _amount, _recipient, _utxo_refund)
}

// RefundIncomingInvoice is a paid mutator transaction binding the contract method 0xdc8cc82f.
//
// Solidity: function refundIncomingInvoice(uint256 _invoiceId, string _utxo, uint256 _amount, address _recipient, string _utxo_refund) returns()
func (_Gateway *GatewayTransactorSession) RefundIncomingInvoice(_invoiceId *big.Int, _utxo string, _amount *big.Int, _recipient common.Address, _utxo_refund string) (*types.Transaction, error) {
	return _Gateway.Contract.RefundIncomingInvoice(&_Gateway.TransactOpts, _invoiceId, _utxo, _amount, _recipient, _utxo_refund)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_Gateway *GatewayTransactor) RemoveValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "removeValidator", _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_Gateway *GatewaySession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.RemoveValidator(&_Gateway.TransactOpts, _validator)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validator) returns()
func (_Gateway *GatewayTransactorSession) RemoveValidator(_validator common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.RemoveValidator(&_Gateway.TransactOpts, _validator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _Gateway.Contract.RenounceOwnership(&_Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gateway.Contract.RenounceOwnership(&_Gateway.TransactOpts)
}

// SubmitTxContent is a paid mutator transaction binding the contract method 0xadfac6a1.
//
// Solidity: function submitTxContent((uint256,uint256,string)[] _invoices, string _txContent) returns(uint256)
func (_Gateway *GatewayTransactor) SubmitTxContent(opts *bind.TransactOpts, _invoices []IGatewayOutgoingInvoiceBasicInfo, _txContent string) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "submitTxContent", _invoices, _txContent)
}

// SubmitTxContent is a paid mutator transaction binding the contract method 0xadfac6a1.
//
// Solidity: function submitTxContent((uint256,uint256,string)[] _invoices, string _txContent) returns(uint256)
func (_Gateway *GatewaySession) SubmitTxContent(_invoices []IGatewayOutgoingInvoiceBasicInfo, _txContent string) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitTxContent(&_Gateway.TransactOpts, _invoices, _txContent)
}

// SubmitTxContent is a paid mutator transaction binding the contract method 0xadfac6a1.
//
// Solidity: function submitTxContent((uint256,uint256,string)[] _invoices, string _txContent) returns(uint256)
func (_Gateway *GatewayTransactorSession) SubmitTxContent(_invoices []IGatewayOutgoingInvoiceBasicInfo, _txContent string) (*types.Transaction, error) {
	return _Gateway.Contract.SubmitTxContent(&_Gateway.TransactOpts, _invoices, _txContent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.TransferOwnership(&_Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.TransferOwnership(&_Gateway.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gateway *GatewayTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gateway *GatewaySession) Unpause() (*types.Transaction, error) {
	return _Gateway.Contract.Unpause(&_Gateway.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Gateway *GatewayTransactorSession) Unpause() (*types.Transaction, error) {
	return _Gateway.Contract.Unpause(&_Gateway.TransactOpts)
}

// UpdateBtcAddress is a paid mutator transaction binding the contract method 0xb005d9a7.
//
// Solidity: function updateBtcAddress(address _btcAddress) returns()
func (_Gateway *GatewayTransactor) UpdateBtcAddress(opts *bind.TransactOpts, _btcAddress common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "updateBtcAddress", _btcAddress)
}

// UpdateBtcAddress is a paid mutator transaction binding the contract method 0xb005d9a7.
//
// Solidity: function updateBtcAddress(address _btcAddress) returns()
func (_Gateway *GatewaySession) UpdateBtcAddress(_btcAddress common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateBtcAddress(&_Gateway.TransactOpts, _btcAddress)
}

// UpdateBtcAddress is a paid mutator transaction binding the contract method 0xb005d9a7.
//
// Solidity: function updateBtcAddress(address _btcAddress) returns()
func (_Gateway *GatewayTransactorSession) UpdateBtcAddress(_btcAddress common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateBtcAddress(&_Gateway.TransactOpts, _btcAddress)
}

// UpdateConcensusThreshold is a paid mutator transaction binding the contract method 0x2c3228d5.
//
// Solidity: function updateConcensusThreshold(uint256 _concensusThreshold) returns()
func (_Gateway *GatewayTransactor) UpdateConcensusThreshold(opts *bind.TransactOpts, _concensusThreshold *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "updateConcensusThreshold", _concensusThreshold)
}

// UpdateConcensusThreshold is a paid mutator transaction binding the contract method 0x2c3228d5.
//
// Solidity: function updateConcensusThreshold(uint256 _concensusThreshold) returns()
func (_Gateway *GatewaySession) UpdateConcensusThreshold(_concensusThreshold *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateConcensusThreshold(&_Gateway.TransactOpts, _concensusThreshold)
}

// UpdateConcensusThreshold is a paid mutator transaction binding the contract method 0x2c3228d5.
//
// Solidity: function updateConcensusThreshold(uint256 _concensusThreshold) returns()
func (_Gateway *GatewayTransactorSession) UpdateConcensusThreshold(_concensusThreshold *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateConcensusThreshold(&_Gateway.TransactOpts, _concensusThreshold)
}

// UpdateProposer is a paid mutator transaction binding the contract method 0xddf9f09f.
//
// Solidity: function updateProposer(address _proposer) returns()
func (_Gateway *GatewayTransactor) UpdateProposer(opts *bind.TransactOpts, _proposer common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "updateProposer", _proposer)
}

// UpdateProposer is a paid mutator transaction binding the contract method 0xddf9f09f.
//
// Solidity: function updateProposer(address _proposer) returns()
func (_Gateway *GatewaySession) UpdateProposer(_proposer common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateProposer(&_Gateway.TransactOpts, _proposer)
}

// UpdateProposer is a paid mutator transaction binding the contract method 0xddf9f09f.
//
// Solidity: function updateProposer(address _proposer) returns()
func (_Gateway *GatewayTransactorSession) UpdateProposer(_proposer common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateProposer(&_Gateway.TransactOpts, _proposer)
}

// VerifyIncomingInvoice is a paid mutator transaction binding the contract method 0x2b7a2c7a.
//
// Solidity: function verifyIncomingInvoice(uint256 _invoiceId, string _utxo, uint256 _amount, address _recipient, bool _isVerified) returns()
func (_Gateway *GatewayTransactor) VerifyIncomingInvoice(opts *bind.TransactOpts, _invoiceId *big.Int, _utxo string, _amount *big.Int, _recipient common.Address, _isVerified bool) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "verifyIncomingInvoice", _invoiceId, _utxo, _amount, _recipient, _isVerified)
}

// VerifyIncomingInvoice is a paid mutator transaction binding the contract method 0x2b7a2c7a.
//
// Solidity: function verifyIncomingInvoice(uint256 _invoiceId, string _utxo, uint256 _amount, address _recipient, bool _isVerified) returns()
func (_Gateway *GatewaySession) VerifyIncomingInvoice(_invoiceId *big.Int, _utxo string, _amount *big.Int, _recipient common.Address, _isVerified bool) (*types.Transaction, error) {
	return _Gateway.Contract.VerifyIncomingInvoice(&_Gateway.TransactOpts, _invoiceId, _utxo, _amount, _recipient, _isVerified)
}

// VerifyIncomingInvoice is a paid mutator transaction binding the contract method 0x2b7a2c7a.
//
// Solidity: function verifyIncomingInvoice(uint256 _invoiceId, string _utxo, uint256 _amount, address _recipient, bool _isVerified) returns()
func (_Gateway *GatewayTransactorSession) VerifyIncomingInvoice(_invoiceId *big.Int, _utxo string, _amount *big.Int, _recipient common.Address, _isVerified bool) (*types.Transaction, error) {
	return _Gateway.Contract.VerifyIncomingInvoice(&_Gateway.TransactOpts, _invoiceId, _utxo, _amount, _recipient, _isVerified)
}

// VerifyOutgoingTx is a paid mutator transaction binding the contract method 0xbd784e7a.
//
// Solidity: function verifyOutgoingTx(uint256 _txId, bool _isVerified, string _signature) returns()
func (_Gateway *GatewayTransactor) VerifyOutgoingTx(opts *bind.TransactOpts, _txId *big.Int, _isVerified bool, _signature string) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "verifyOutgoingTx", _txId, _isVerified, _signature)
}

// VerifyOutgoingTx is a paid mutator transaction binding the contract method 0xbd784e7a.
//
// Solidity: function verifyOutgoingTx(uint256 _txId, bool _isVerified, string _signature) returns()
func (_Gateway *GatewaySession) VerifyOutgoingTx(_txId *big.Int, _isVerified bool, _signature string) (*types.Transaction, error) {
	return _Gateway.Contract.VerifyOutgoingTx(&_Gateway.TransactOpts, _txId, _isVerified, _signature)
}

// VerifyOutgoingTx is a paid mutator transaction binding the contract method 0xbd784e7a.
//
// Solidity: function verifyOutgoingTx(uint256 _txId, bool _isVerified, string _signature) returns()
func (_Gateway *GatewayTransactorSession) VerifyOutgoingTx(_txId *big.Int, _isVerified bool, _signature string) (*types.Transaction, error) {
	return _Gateway.Contract.VerifyOutgoingTx(&_Gateway.TransactOpts, _txId, _isVerified, _signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Gateway *GatewayTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Gateway *GatewaySession) Withdraw() (*types.Transaction, error) {
	return _Gateway.Contract.Withdraw(&_Gateway.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Gateway *GatewayTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Gateway.Contract.Withdraw(&_Gateway.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x607686f2.
//
// Solidity: function withdrawToken(uint256 amount, string recipient) returns()
func (_Gateway *GatewayTransactor) WithdrawToken(opts *bind.TransactOpts, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "withdrawToken", amount, recipient)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x607686f2.
//
// Solidity: function withdrawToken(uint256 amount, string recipient) returns()
func (_Gateway *GatewaySession) WithdrawToken(amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Gateway.Contract.WithdrawToken(&_Gateway.TransactOpts, amount, recipient)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x607686f2.
//
// Solidity: function withdrawToken(uint256 amount, string recipient) returns()
func (_Gateway *GatewayTransactorSession) WithdrawToken(amount *big.Int, recipient string) (*types.Transaction, error) {
	return _Gateway.Contract.WithdrawToken(&_Gateway.TransactOpts, amount, recipient)
}

// GatewayIncomingInvoiceCompletedIterator is returned from FilterIncomingInvoiceCompleted and is used to iterate over the raw logs and unpacked data for IncomingInvoiceCompleted events raised by the Gateway contract.
type GatewayIncomingInvoiceCompletedIterator struct {
	Event *GatewayIncomingInvoiceCompleted // Event containing the contract specifics and raw log

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
func (it *GatewayIncomingInvoiceCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayIncomingInvoiceCompleted)
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
		it.Event = new(GatewayIncomingInvoiceCompleted)
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
func (it *GatewayIncomingInvoiceCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayIncomingInvoiceCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayIncomingInvoiceCompleted represents a IncomingInvoiceCompleted event raised by the Gateway contract.
type GatewayIncomingInvoiceCompleted struct {
	InvoiceId *big.Int
	Status    uint8
	Executor  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncomingInvoiceCompleted is a free log retrieval operation binding the contract event 0x0040b1b3943ee0be2d706b8aaa4143a310c038463192e0d26b446e7edff7f80f.
//
// Solidity: event IncomingInvoiceCompleted(uint256 indexed invoiceId, uint8 status, address executor)
func (_Gateway *GatewayFilterer) FilterIncomingInvoiceCompleted(opts *bind.FilterOpts, invoiceId []*big.Int) (*GatewayIncomingInvoiceCompletedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "IncomingInvoiceCompleted", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayIncomingInvoiceCompletedIterator{contract: _Gateway.contract, event: "IncomingInvoiceCompleted", logs: logs, sub: sub}, nil
}

// WatchIncomingInvoiceCompleted is a free log subscription operation binding the contract event 0x0040b1b3943ee0be2d706b8aaa4143a310c038463192e0d26b446e7edff7f80f.
//
// Solidity: event IncomingInvoiceCompleted(uint256 indexed invoiceId, uint8 status, address executor)
func (_Gateway *GatewayFilterer) WatchIncomingInvoiceCompleted(opts *bind.WatchOpts, sink chan<- *GatewayIncomingInvoiceCompleted, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "IncomingInvoiceCompleted", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayIncomingInvoiceCompleted)
				if err := _Gateway.contract.UnpackLog(event, "IncomingInvoiceCompleted", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseIncomingInvoiceCompleted(log types.Log) (*GatewayIncomingInvoiceCompleted, error) {
	event := new(GatewayIncomingInvoiceCompleted)
	if err := _Gateway.contract.UnpackLog(event, "IncomingInvoiceCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayIncomingInvoiceCreatedIterator is returned from FilterIncomingInvoiceCreated and is used to iterate over the raw logs and unpacked data for IncomingInvoiceCreated events raised by the Gateway contract.
type GatewayIncomingInvoiceCreatedIterator struct {
	Event *GatewayIncomingInvoiceCreated // Event containing the contract specifics and raw log

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
func (it *GatewayIncomingInvoiceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayIncomingInvoiceCreated)
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
		it.Event = new(GatewayIncomingInvoiceCreated)
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
func (it *GatewayIncomingInvoiceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayIncomingInvoiceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayIncomingInvoiceCreated represents a IncomingInvoiceCreated event raised by the Gateway contract.
type GatewayIncomingInvoiceCreated struct {
	InvoiceId *big.Int
	Utxo      string
	Amount    *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncomingInvoiceCreated is a free log retrieval operation binding the contract event 0xf9ef8cea38fef9dce3302475d87b4a1b3a67134a277ee66e644d173758fab41c.
//
// Solidity: event IncomingInvoiceCreated(uint256 indexed invoiceId, string utxo, uint256 amount, address recipient)
func (_Gateway *GatewayFilterer) FilterIncomingInvoiceCreated(opts *bind.FilterOpts, invoiceId []*big.Int) (*GatewayIncomingInvoiceCreatedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "IncomingInvoiceCreated", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayIncomingInvoiceCreatedIterator{contract: _Gateway.contract, event: "IncomingInvoiceCreated", logs: logs, sub: sub}, nil
}

// WatchIncomingInvoiceCreated is a free log subscription operation binding the contract event 0xf9ef8cea38fef9dce3302475d87b4a1b3a67134a277ee66e644d173758fab41c.
//
// Solidity: event IncomingInvoiceCreated(uint256 indexed invoiceId, string utxo, uint256 amount, address recipient)
func (_Gateway *GatewayFilterer) WatchIncomingInvoiceCreated(opts *bind.WatchOpts, sink chan<- *GatewayIncomingInvoiceCreated, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "IncomingInvoiceCreated", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayIncomingInvoiceCreated)
				if err := _Gateway.contract.UnpackLog(event, "IncomingInvoiceCreated", log); err != nil {
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

// ParseIncomingInvoiceCreated is a log parse operation binding the contract event 0xf9ef8cea38fef9dce3302475d87b4a1b3a67134a277ee66e644d173758fab41c.
//
// Solidity: event IncomingInvoiceCreated(uint256 indexed invoiceId, string utxo, uint256 amount, address recipient)
func (_Gateway *GatewayFilterer) ParseIncomingInvoiceCreated(log types.Log) (*GatewayIncomingInvoiceCreated, error) {
	event := new(GatewayIncomingInvoiceCreated)
	if err := _Gateway.contract.UnpackLog(event, "IncomingInvoiceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayIncomingInvoiceRefundedIterator is returned from FilterIncomingInvoiceRefunded and is used to iterate over the raw logs and unpacked data for IncomingInvoiceRefunded events raised by the Gateway contract.
type GatewayIncomingInvoiceRefundedIterator struct {
	Event *GatewayIncomingInvoiceRefunded // Event containing the contract specifics and raw log

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
func (it *GatewayIncomingInvoiceRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayIncomingInvoiceRefunded)
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
		it.Event = new(GatewayIncomingInvoiceRefunded)
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
func (it *GatewayIncomingInvoiceRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayIncomingInvoiceRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayIncomingInvoiceRefunded represents a IncomingInvoiceRefunded event raised by the Gateway contract.
type GatewayIncomingInvoiceRefunded struct {
	InvoiceId  *big.Int
	UtxoRefund string
	Executor   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterIncomingInvoiceRefunded is a free log retrieval operation binding the contract event 0x5d71f72ffb55a011fed5421c0fe881100f1524d6dcee520bd3a2b83695fdc2fa.
//
// Solidity: event IncomingInvoiceRefunded(uint256 indexed invoiceId, string utxo_refund, address executor)
func (_Gateway *GatewayFilterer) FilterIncomingInvoiceRefunded(opts *bind.FilterOpts, invoiceId []*big.Int) (*GatewayIncomingInvoiceRefundedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "IncomingInvoiceRefunded", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayIncomingInvoiceRefundedIterator{contract: _Gateway.contract, event: "IncomingInvoiceRefunded", logs: logs, sub: sub}, nil
}

// WatchIncomingInvoiceRefunded is a free log subscription operation binding the contract event 0x5d71f72ffb55a011fed5421c0fe881100f1524d6dcee520bd3a2b83695fdc2fa.
//
// Solidity: event IncomingInvoiceRefunded(uint256 indexed invoiceId, string utxo_refund, address executor)
func (_Gateway *GatewayFilterer) WatchIncomingInvoiceRefunded(opts *bind.WatchOpts, sink chan<- *GatewayIncomingInvoiceRefunded, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "IncomingInvoiceRefunded", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayIncomingInvoiceRefunded)
				if err := _Gateway.contract.UnpackLog(event, "IncomingInvoiceRefunded", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseIncomingInvoiceRefunded(log types.Log) (*GatewayIncomingInvoiceRefunded, error) {
	event := new(GatewayIncomingInvoiceRefunded)
	if err := _Gateway.contract.UnpackLog(event, "IncomingInvoiceRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayIncomingInvoiceRejectedIterator is returned from FilterIncomingInvoiceRejected and is used to iterate over the raw logs and unpacked data for IncomingInvoiceRejected events raised by the Gateway contract.
type GatewayIncomingInvoiceRejectedIterator struct {
	Event *GatewayIncomingInvoiceRejected // Event containing the contract specifics and raw log

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
func (it *GatewayIncomingInvoiceRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayIncomingInvoiceRejected)
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
		it.Event = new(GatewayIncomingInvoiceRejected)
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
func (it *GatewayIncomingInvoiceRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayIncomingInvoiceRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayIncomingInvoiceRejected represents a IncomingInvoiceRejected event raised by the Gateway contract.
type GatewayIncomingInvoiceRejected struct {
	InvoiceId *big.Int
	Status    uint8
	Executor  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncomingInvoiceRejected is a free log retrieval operation binding the contract event 0x7da96e6608f6e569e80aaef406ea92a86644ee8234e16f847c93880a0d8c16ff.
//
// Solidity: event IncomingInvoiceRejected(uint256 indexed invoiceId, uint8 status, address executor)
func (_Gateway *GatewayFilterer) FilterIncomingInvoiceRejected(opts *bind.FilterOpts, invoiceId []*big.Int) (*GatewayIncomingInvoiceRejectedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "IncomingInvoiceRejected", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayIncomingInvoiceRejectedIterator{contract: _Gateway.contract, event: "IncomingInvoiceRejected", logs: logs, sub: sub}, nil
}

// WatchIncomingInvoiceRejected is a free log subscription operation binding the contract event 0x7da96e6608f6e569e80aaef406ea92a86644ee8234e16f847c93880a0d8c16ff.
//
// Solidity: event IncomingInvoiceRejected(uint256 indexed invoiceId, uint8 status, address executor)
func (_Gateway *GatewayFilterer) WatchIncomingInvoiceRejected(opts *bind.WatchOpts, sink chan<- *GatewayIncomingInvoiceRejected, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "IncomingInvoiceRejected", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayIncomingInvoiceRejected)
				if err := _Gateway.contract.UnpackLog(event, "IncomingInvoiceRejected", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseIncomingInvoiceRejected(log types.Log) (*GatewayIncomingInvoiceRejected, error) {
	event := new(GatewayIncomingInvoiceRejected)
	if err := _Gateway.contract.UnpackLog(event, "IncomingInvoiceRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayIncomingInvoiceVerifiedIterator is returned from FilterIncomingInvoiceVerified and is used to iterate over the raw logs and unpacked data for IncomingInvoiceVerified events raised by the Gateway contract.
type GatewayIncomingInvoiceVerifiedIterator struct {
	Event *GatewayIncomingInvoiceVerified // Event containing the contract specifics and raw log

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
func (it *GatewayIncomingInvoiceVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayIncomingInvoiceVerified)
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
		it.Event = new(GatewayIncomingInvoiceVerified)
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
func (it *GatewayIncomingInvoiceVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayIncomingInvoiceVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayIncomingInvoiceVerified represents a IncomingInvoiceVerified event raised by the Gateway contract.
type GatewayIncomingInvoiceVerified struct {
	InvoiceId *big.Int
	Validator common.Address
	Verified  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncomingInvoiceVerified is a free log retrieval operation binding the contract event 0x2d88fccf7a5b293b2b7dee2f4671af10b2df68f488bc92af96fea710ce36a756.
//
// Solidity: event IncomingInvoiceVerified(uint256 indexed invoiceId, address validator, bool verified)
func (_Gateway *GatewayFilterer) FilterIncomingInvoiceVerified(opts *bind.FilterOpts, invoiceId []*big.Int) (*GatewayIncomingInvoiceVerifiedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "IncomingInvoiceVerified", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayIncomingInvoiceVerifiedIterator{contract: _Gateway.contract, event: "IncomingInvoiceVerified", logs: logs, sub: sub}, nil
}

// WatchIncomingInvoiceVerified is a free log subscription operation binding the contract event 0x2d88fccf7a5b293b2b7dee2f4671af10b2df68f488bc92af96fea710ce36a756.
//
// Solidity: event IncomingInvoiceVerified(uint256 indexed invoiceId, address validator, bool verified)
func (_Gateway *GatewayFilterer) WatchIncomingInvoiceVerified(opts *bind.WatchOpts, sink chan<- *GatewayIncomingInvoiceVerified, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "IncomingInvoiceVerified", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayIncomingInvoiceVerified)
				if err := _Gateway.contract.UnpackLog(event, "IncomingInvoiceVerified", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseIncomingInvoiceVerified(log types.Log) (*GatewayIncomingInvoiceVerified, error) {
	event := new(GatewayIncomingInvoiceVerified)
	if err := _Gateway.contract.UnpackLog(event, "IncomingInvoiceVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Gateway contract.
type GatewayInitializedIterator struct {
	Event *GatewayInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayInitialized)
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
		it.Event = new(GatewayInitialized)
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
func (it *GatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayInitialized represents a Initialized event raised by the Gateway contract.
type GatewayInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Gateway *GatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayInitializedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayInitializedIterator{contract: _Gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Gateway *GatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayInitialized)
				if err := _Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseInitialized(log types.Log) (*GatewayInitialized, error) {
	event := new(GatewayInitialized)
	if err := _Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayOutgoingInvoiceCreatedIterator is returned from FilterOutgoingInvoiceCreated and is used to iterate over the raw logs and unpacked data for OutgoingInvoiceCreated events raised by the Gateway contract.
type GatewayOutgoingInvoiceCreatedIterator struct {
	Event *GatewayOutgoingInvoiceCreated // Event containing the contract specifics and raw log

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
func (it *GatewayOutgoingInvoiceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayOutgoingInvoiceCreated)
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
		it.Event = new(GatewayOutgoingInvoiceCreated)
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
func (it *GatewayOutgoingInvoiceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayOutgoingInvoiceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayOutgoingInvoiceCreated represents a OutgoingInvoiceCreated event raised by the Gateway contract.
type GatewayOutgoingInvoiceCreated struct {
	InvoiceId *big.Int
	Amount    *big.Int
	From      common.Address
	Recipient string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutgoingInvoiceCreated is a free log retrieval operation binding the contract event 0xb134d5a9363aa8d267d6d14febe01876d86730056c14d084eadc8541a4d35a1e.
//
// Solidity: event OutgoingInvoiceCreated(uint256 indexed invoiceId, uint256 amount, address from, string recipient)
func (_Gateway *GatewayFilterer) FilterOutgoingInvoiceCreated(opts *bind.FilterOpts, invoiceId []*big.Int) (*GatewayOutgoingInvoiceCreatedIterator, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "OutgoingInvoiceCreated", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayOutgoingInvoiceCreatedIterator{contract: _Gateway.contract, event: "OutgoingInvoiceCreated", logs: logs, sub: sub}, nil
}

// WatchOutgoingInvoiceCreated is a free log subscription operation binding the contract event 0xb134d5a9363aa8d267d6d14febe01876d86730056c14d084eadc8541a4d35a1e.
//
// Solidity: event OutgoingInvoiceCreated(uint256 indexed invoiceId, uint256 amount, address from, string recipient)
func (_Gateway *GatewayFilterer) WatchOutgoingInvoiceCreated(opts *bind.WatchOpts, sink chan<- *GatewayOutgoingInvoiceCreated, invoiceId []*big.Int) (event.Subscription, error) {

	var invoiceIdRule []interface{}
	for _, invoiceIdItem := range invoiceId {
		invoiceIdRule = append(invoiceIdRule, invoiceIdItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "OutgoingInvoiceCreated", invoiceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayOutgoingInvoiceCreated)
				if err := _Gateway.contract.UnpackLog(event, "OutgoingInvoiceCreated", log); err != nil {
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

// ParseOutgoingInvoiceCreated is a log parse operation binding the contract event 0xb134d5a9363aa8d267d6d14febe01876d86730056c14d084eadc8541a4d35a1e.
//
// Solidity: event OutgoingInvoiceCreated(uint256 indexed invoiceId, uint256 amount, address from, string recipient)
func (_Gateway *GatewayFilterer) ParseOutgoingInvoiceCreated(log types.Log) (*GatewayOutgoingInvoiceCreated, error) {
	event := new(GatewayOutgoingInvoiceCreated)
	if err := _Gateway.contract.UnpackLog(event, "OutgoingInvoiceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayOutgoingTxProcessedIterator is returned from FilterOutgoingTxProcessed and is used to iterate over the raw logs and unpacked data for OutgoingTxProcessed events raised by the Gateway contract.
type GatewayOutgoingTxProcessedIterator struct {
	Event *GatewayOutgoingTxProcessed // Event containing the contract specifics and raw log

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
func (it *GatewayOutgoingTxProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayOutgoingTxProcessed)
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
		it.Event = new(GatewayOutgoingTxProcessed)
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
func (it *GatewayOutgoingTxProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayOutgoingTxProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayOutgoingTxProcessed represents a OutgoingTxProcessed event raised by the Gateway contract.
type GatewayOutgoingTxProcessed struct {
	TxId   *big.Int
	Status uint8
	TxHash string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOutgoingTxProcessed is a free log retrieval operation binding the contract event 0xbb51d23336167c8e20a9ca20fbe3f55ebd9339040e3d56a545f2527bc55d3f3b.
//
// Solidity: event OutgoingTxProcessed(uint256 indexed txId, uint8 status, string txHash)
func (_Gateway *GatewayFilterer) FilterOutgoingTxProcessed(opts *bind.FilterOpts, txId []*big.Int) (*GatewayOutgoingTxProcessedIterator, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "OutgoingTxProcessed", txIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayOutgoingTxProcessedIterator{contract: _Gateway.contract, event: "OutgoingTxProcessed", logs: logs, sub: sub}, nil
}

// WatchOutgoingTxProcessed is a free log subscription operation binding the contract event 0xbb51d23336167c8e20a9ca20fbe3f55ebd9339040e3d56a545f2527bc55d3f3b.
//
// Solidity: event OutgoingTxProcessed(uint256 indexed txId, uint8 status, string txHash)
func (_Gateway *GatewayFilterer) WatchOutgoingTxProcessed(opts *bind.WatchOpts, sink chan<- *GatewayOutgoingTxProcessed, txId []*big.Int) (event.Subscription, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "OutgoingTxProcessed", txIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayOutgoingTxProcessed)
				if err := _Gateway.contract.UnpackLog(event, "OutgoingTxProcessed", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseOutgoingTxProcessed(log types.Log) (*GatewayOutgoingTxProcessed, error) {
	event := new(GatewayOutgoingTxProcessed)
	if err := _Gateway.contract.UnpackLog(event, "OutgoingTxProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayOutgoingTxSubmittedIterator is returned from FilterOutgoingTxSubmitted and is used to iterate over the raw logs and unpacked data for OutgoingTxSubmitted events raised by the Gateway contract.
type GatewayOutgoingTxSubmittedIterator struct {
	Event *GatewayOutgoingTxSubmitted // Event containing the contract specifics and raw log

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
func (it *GatewayOutgoingTxSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayOutgoingTxSubmitted)
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
		it.Event = new(GatewayOutgoingTxSubmitted)
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
func (it *GatewayOutgoingTxSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayOutgoingTxSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayOutgoingTxSubmitted represents a OutgoingTxSubmitted event raised by the Gateway contract.
type GatewayOutgoingTxSubmitted struct {
	TxId      *big.Int
	TxContent string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutgoingTxSubmitted is a free log retrieval operation binding the contract event 0xda987441b8bab19820139d670528f4e4de7f914c8018f6ef2894e83f69e67f21.
//
// Solidity: event OutgoingTxSubmitted(uint256 indexed txId, string txContent)
func (_Gateway *GatewayFilterer) FilterOutgoingTxSubmitted(opts *bind.FilterOpts, txId []*big.Int) (*GatewayOutgoingTxSubmittedIterator, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "OutgoingTxSubmitted", txIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayOutgoingTxSubmittedIterator{contract: _Gateway.contract, event: "OutgoingTxSubmitted", logs: logs, sub: sub}, nil
}

// WatchOutgoingTxSubmitted is a free log subscription operation binding the contract event 0xda987441b8bab19820139d670528f4e4de7f914c8018f6ef2894e83f69e67f21.
//
// Solidity: event OutgoingTxSubmitted(uint256 indexed txId, string txContent)
func (_Gateway *GatewayFilterer) WatchOutgoingTxSubmitted(opts *bind.WatchOpts, sink chan<- *GatewayOutgoingTxSubmitted, txId []*big.Int) (event.Subscription, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "OutgoingTxSubmitted", txIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayOutgoingTxSubmitted)
				if err := _Gateway.contract.UnpackLog(event, "OutgoingTxSubmitted", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseOutgoingTxSubmitted(log types.Log) (*GatewayOutgoingTxSubmitted, error) {
	event := new(GatewayOutgoingTxSubmitted)
	if err := _Gateway.contract.UnpackLog(event, "OutgoingTxSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayOutgoingTxVerifiedIterator is returned from FilterOutgoingTxVerified and is used to iterate over the raw logs and unpacked data for OutgoingTxVerified events raised by the Gateway contract.
type GatewayOutgoingTxVerifiedIterator struct {
	Event *GatewayOutgoingTxVerified // Event containing the contract specifics and raw log

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
func (it *GatewayOutgoingTxVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayOutgoingTxVerified)
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
		it.Event = new(GatewayOutgoingTxVerified)
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
func (it *GatewayOutgoingTxVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayOutgoingTxVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayOutgoingTxVerified represents a OutgoingTxVerified event raised by the Gateway contract.
type GatewayOutgoingTxVerified struct {
	TxId      *big.Int
	Validator common.Address
	Verified  bool
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOutgoingTxVerified is a free log retrieval operation binding the contract event 0x9191d6ec9841818acc620ffc69e1390adf521156b9383a417dfe403447a7a0c7.
//
// Solidity: event OutgoingTxVerified(uint256 indexed txId, address validator, bool verified, string signature)
func (_Gateway *GatewayFilterer) FilterOutgoingTxVerified(opts *bind.FilterOpts, txId []*big.Int) (*GatewayOutgoingTxVerifiedIterator, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "OutgoingTxVerified", txIdRule)
	if err != nil {
		return nil, err
	}
	return &GatewayOutgoingTxVerifiedIterator{contract: _Gateway.contract, event: "OutgoingTxVerified", logs: logs, sub: sub}, nil
}

// WatchOutgoingTxVerified is a free log subscription operation binding the contract event 0x9191d6ec9841818acc620ffc69e1390adf521156b9383a417dfe403447a7a0c7.
//
// Solidity: event OutgoingTxVerified(uint256 indexed txId, address validator, bool verified, string signature)
func (_Gateway *GatewayFilterer) WatchOutgoingTxVerified(opts *bind.WatchOpts, sink chan<- *GatewayOutgoingTxVerified, txId []*big.Int) (event.Subscription, error) {

	var txIdRule []interface{}
	for _, txIdItem := range txId {
		txIdRule = append(txIdRule, txIdItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "OutgoingTxVerified", txIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayOutgoingTxVerified)
				if err := _Gateway.contract.UnpackLog(event, "OutgoingTxVerified", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseOutgoingTxVerified(log types.Log) (*GatewayOutgoingTxVerified, error) {
	event := new(GatewayOutgoingTxVerified)
	if err := _Gateway.contract.UnpackLog(event, "OutgoingTxVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gateway contract.
type GatewayOwnershipTransferredIterator struct {
	Event *GatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayOwnershipTransferred)
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
		it.Event = new(GatewayOwnershipTransferred)
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
func (it *GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the Gateway contract.
type GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gateway *GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayOwnershipTransferredIterator{contract: _Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gateway *GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayOwnershipTransferred)
				if err := _Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayOwnershipTransferred, error) {
	event := new(GatewayOwnershipTransferred)
	if err := _Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Gateway contract.
type GatewayPausedIterator struct {
	Event *GatewayPaused // Event containing the contract specifics and raw log

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
func (it *GatewayPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayPaused)
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
		it.Event = new(GatewayPaused)
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
func (it *GatewayPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayPaused represents a Paused event raised by the Gateway contract.
type GatewayPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Gateway *GatewayFilterer) FilterPaused(opts *bind.FilterOpts) (*GatewayPausedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GatewayPausedIterator{contract: _Gateway.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Gateway *GatewayFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GatewayPaused) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayPaused)
				if err := _Gateway.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParsePaused(log types.Log) (*GatewayPaused, error) {
	event := new(GatewayPaused)
	if err := _Gateway.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Gateway contract.
type GatewayUnpausedIterator struct {
	Event *GatewayUnpaused // Event containing the contract specifics and raw log

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
func (it *GatewayUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayUnpaused)
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
		it.Event = new(GatewayUnpaused)
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
func (it *GatewayUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayUnpaused represents a Unpaused event raised by the Gateway contract.
type GatewayUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Gateway *GatewayFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GatewayUnpausedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GatewayUnpausedIterator{contract: _Gateway.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Gateway *GatewayFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GatewayUnpaused) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayUnpaused)
				if err := _Gateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseUnpaused(log types.Log) (*GatewayUnpaused, error) {
	event := new(GatewayUnpaused)
	if err := _Gateway.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
