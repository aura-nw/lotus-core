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

// AuraWrappedBitcoinMetaData contains all meta data concerning the AuraWrappedBitcoin contract.
var AuraWrappedBitcoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gateway_\",\"type\":\"address\"}],\"name\":\"setGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"recipient\",\"type\":\"string\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AuraWrappedBitcoinABI is the input ABI used to generate the binding from.
// Deprecated: Use AuraWrappedBitcoinMetaData.ABI instead.
var AuraWrappedBitcoinABI = AuraWrappedBitcoinMetaData.ABI

// AuraWrappedBitcoin is an auto generated Go binding around an Ethereum contract.
type AuraWrappedBitcoin struct {
	AuraWrappedBitcoinCaller     // Read-only binding to the contract
	AuraWrappedBitcoinTransactor // Write-only binding to the contract
	AuraWrappedBitcoinFilterer   // Log filterer for contract events
}

// AuraWrappedBitcoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuraWrappedBitcoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraWrappedBitcoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuraWrappedBitcoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraWrappedBitcoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuraWrappedBitcoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraWrappedBitcoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuraWrappedBitcoinSession struct {
	Contract     *AuraWrappedBitcoin // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AuraWrappedBitcoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuraWrappedBitcoinCallerSession struct {
	Contract *AuraWrappedBitcoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AuraWrappedBitcoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuraWrappedBitcoinTransactorSession struct {
	Contract     *AuraWrappedBitcoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AuraWrappedBitcoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuraWrappedBitcoinRaw struct {
	Contract *AuraWrappedBitcoin // Generic contract binding to access the raw methods on
}

// AuraWrappedBitcoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuraWrappedBitcoinCallerRaw struct {
	Contract *AuraWrappedBitcoinCaller // Generic read-only contract binding to access the raw methods on
}

// AuraWrappedBitcoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuraWrappedBitcoinTransactorRaw struct {
	Contract *AuraWrappedBitcoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuraWrappedBitcoin creates a new instance of AuraWrappedBitcoin, bound to a specific deployed contract.
func NewAuraWrappedBitcoin(address common.Address, backend bind.ContractBackend) (*AuraWrappedBitcoin, error) {
	contract, err := bindAuraWrappedBitcoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBitcoin{AuraWrappedBitcoinCaller: AuraWrappedBitcoinCaller{contract: contract}, AuraWrappedBitcoinTransactor: AuraWrappedBitcoinTransactor{contract: contract}, AuraWrappedBitcoinFilterer: AuraWrappedBitcoinFilterer{contract: contract}}, nil
}

// NewAuraWrappedBitcoinCaller creates a new read-only instance of AuraWrappedBitcoin, bound to a specific deployed contract.
func NewAuraWrappedBitcoinCaller(address common.Address, caller bind.ContractCaller) (*AuraWrappedBitcoinCaller, error) {
	contract, err := bindAuraWrappedBitcoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBitcoinCaller{contract: contract}, nil
}

// NewAuraWrappedBitcoinTransactor creates a new write-only instance of AuraWrappedBitcoin, bound to a specific deployed contract.
func NewAuraWrappedBitcoinTransactor(address common.Address, transactor bind.ContractTransactor) (*AuraWrappedBitcoinTransactor, error) {
	contract, err := bindAuraWrappedBitcoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBitcoinTransactor{contract: contract}, nil
}

// NewAuraWrappedBitcoinFilterer creates a new log filterer instance of AuraWrappedBitcoin, bound to a specific deployed contract.
func NewAuraWrappedBitcoinFilterer(address common.Address, filterer bind.ContractFilterer) (*AuraWrappedBitcoinFilterer, error) {
	contract, err := bindAuraWrappedBitcoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBitcoinFilterer{contract: contract}, nil
}

// bindAuraWrappedBitcoin binds a generic wrapper to an already deployed contract.
func bindAuraWrappedBitcoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuraWrappedBitcoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuraWrappedBitcoin *AuraWrappedBitcoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuraWrappedBitcoin.Contract.AuraWrappedBitcoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuraWrappedBitcoin *AuraWrappedBitcoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.AuraWrappedBitcoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuraWrappedBitcoin *AuraWrappedBitcoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.AuraWrappedBitcoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuraWrappedBitcoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraWrappedBitcoin.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AuraWrappedBitcoin.Contract.Allowance(&_AuraWrappedBitcoin.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AuraWrappedBitcoin.Contract.Allowance(&_AuraWrappedBitcoin.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraWrappedBitcoin.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AuraWrappedBitcoin.Contract.BalanceOf(&_AuraWrappedBitcoin.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AuraWrappedBitcoin.Contract.BalanceOf(&_AuraWrappedBitcoin.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AuraWrappedBitcoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Decimals() (uint8, error) {
	return _AuraWrappedBitcoin.Contract.Decimals(&_AuraWrappedBitcoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerSession) Decimals() (uint8, error) {
	return _AuraWrappedBitcoin.Contract.Decimals(&_AuraWrappedBitcoin.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraWrappedBitcoin.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Gateway() (common.Address, error) {
	return _AuraWrappedBitcoin.Contract.Gateway(&_AuraWrappedBitcoin.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerSession) Gateway() (common.Address, error) {
	return _AuraWrappedBitcoin.Contract.Gateway(&_AuraWrappedBitcoin.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraWrappedBitcoin.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) MaxSupply() (*big.Int, error) {
	return _AuraWrappedBitcoin.Contract.MaxSupply(&_AuraWrappedBitcoin.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerSession) MaxSupply() (*big.Int, error) {
	return _AuraWrappedBitcoin.Contract.MaxSupply(&_AuraWrappedBitcoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AuraWrappedBitcoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Name() (string, error) {
	return _AuraWrappedBitcoin.Contract.Name(&_AuraWrappedBitcoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerSession) Name() (string, error) {
	return _AuraWrappedBitcoin.Contract.Name(&_AuraWrappedBitcoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraWrappedBitcoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Owner() (common.Address, error) {
	return _AuraWrappedBitcoin.Contract.Owner(&_AuraWrappedBitcoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerSession) Owner() (common.Address, error) {
	return _AuraWrappedBitcoin.Contract.Owner(&_AuraWrappedBitcoin.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AuraWrappedBitcoin.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Paused() (bool, error) {
	return _AuraWrappedBitcoin.Contract.Paused(&_AuraWrappedBitcoin.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerSession) Paused() (bool, error) {
	return _AuraWrappedBitcoin.Contract.Paused(&_AuraWrappedBitcoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AuraWrappedBitcoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Symbol() (string, error) {
	return _AuraWrappedBitcoin.Contract.Symbol(&_AuraWrappedBitcoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerSession) Symbol() (string, error) {
	return _AuraWrappedBitcoin.Contract.Symbol(&_AuraWrappedBitcoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraWrappedBitcoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) TotalSupply() (*big.Int, error) {
	return _AuraWrappedBitcoin.Contract.TotalSupply(&_AuraWrappedBitcoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinCallerSession) TotalSupply() (*big.Int, error) {
	return _AuraWrappedBitcoin.Contract.TotalSupply(&_AuraWrappedBitcoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Approve(&_AuraWrappedBitcoin.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Approve(&_AuraWrappedBitcoin.TransactOpts, spender, value)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Initialize(&_AuraWrappedBitcoin.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Initialize(&_AuraWrappedBitcoin.TransactOpts, initialOwner)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount, uint256 invoiceId) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "mint", to, amount, invoiceId)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount, uint256 invoiceId) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Mint(to common.Address, amount *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Mint(&_AuraWrappedBitcoin.TransactOpts, to, amount, invoiceId)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount, uint256 invoiceId) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) Mint(to common.Address, amount *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Mint(&_AuraWrappedBitcoin.TransactOpts, to, amount, invoiceId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Pause() (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Pause(&_AuraWrappedBitcoin.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) Pause() (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Pause(&_AuraWrappedBitcoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.RenounceOwnership(&_AuraWrappedBitcoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.RenounceOwnership(&_AuraWrappedBitcoin.TransactOpts)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address gateway_) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) SetGateway(opts *bind.TransactOpts, gateway_ common.Address) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "setGateway", gateway_)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address gateway_) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) SetGateway(gateway_ common.Address) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.SetGateway(&_AuraWrappedBitcoin.TransactOpts, gateway_)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address gateway_) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) SetGateway(gateway_ common.Address) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.SetGateway(&_AuraWrappedBitcoin.TransactOpts, gateway_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Transfer(&_AuraWrappedBitcoin.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Transfer(&_AuraWrappedBitcoin.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.TransferFrom(&_AuraWrappedBitcoin.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.TransferFrom(&_AuraWrappedBitcoin.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.TransferOwnership(&_AuraWrappedBitcoin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.TransferOwnership(&_AuraWrappedBitcoin.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Unpause() (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Unpause(&_AuraWrappedBitcoin.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) Unpause() (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Unpause(&_AuraWrappedBitcoin.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xdc048cf2.
//
// Solidity: function withdraw(uint256 amount, string recipient) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, recipient string) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.contract.Transact(opts, "withdraw", amount, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0xdc048cf2.
//
// Solidity: function withdraw(uint256 amount, string recipient) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinSession) Withdraw(amount *big.Int, recipient string) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Withdraw(&_AuraWrappedBitcoin.TransactOpts, amount, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0xdc048cf2.
//
// Solidity: function withdraw(uint256 amount, string recipient) returns()
func (_AuraWrappedBitcoin *AuraWrappedBitcoinTransactorSession) Withdraw(amount *big.Int, recipient string) (*types.Transaction, error) {
	return _AuraWrappedBitcoin.Contract.Withdraw(&_AuraWrappedBitcoin.TransactOpts, amount, recipient)
}

// AuraWrappedBitcoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinApprovalIterator struct {
	Event *AuraWrappedBitcoinApproval // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBitcoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBitcoinApproval)
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
		it.Event = new(AuraWrappedBitcoinApproval)
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
func (it *AuraWrappedBitcoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBitcoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBitcoinApproval represents a Approval event raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AuraWrappedBitcoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AuraWrappedBitcoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBitcoinApprovalIterator{contract: _AuraWrappedBitcoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AuraWrappedBitcoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AuraWrappedBitcoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBitcoinApproval)
				if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) ParseApproval(log types.Log) (*AuraWrappedBitcoinApproval, error) {
	event := new(AuraWrappedBitcoinApproval)
	if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraWrappedBitcoinInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinInitializedIterator struct {
	Event *AuraWrappedBitcoinInitialized // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBitcoinInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBitcoinInitialized)
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
		it.Event = new(AuraWrappedBitcoinInitialized)
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
func (it *AuraWrappedBitcoinInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBitcoinInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBitcoinInitialized represents a Initialized event raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) FilterInitialized(opts *bind.FilterOpts) (*AuraWrappedBitcoinInitializedIterator, error) {

	logs, sub, err := _AuraWrappedBitcoin.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBitcoinInitializedIterator{contract: _AuraWrappedBitcoin.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AuraWrappedBitcoinInitialized) (event.Subscription, error) {

	logs, sub, err := _AuraWrappedBitcoin.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBitcoinInitialized)
				if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) ParseInitialized(log types.Log) (*AuraWrappedBitcoinInitialized, error) {
	event := new(AuraWrappedBitcoinInitialized)
	if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraWrappedBitcoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinOwnershipTransferredIterator struct {
	Event *AuraWrappedBitcoinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBitcoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBitcoinOwnershipTransferred)
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
		it.Event = new(AuraWrappedBitcoinOwnershipTransferred)
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
func (it *AuraWrappedBitcoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBitcoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBitcoinOwnershipTransferred represents a OwnershipTransferred event raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuraWrappedBitcoinOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuraWrappedBitcoin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBitcoinOwnershipTransferredIterator{contract: _AuraWrappedBitcoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuraWrappedBitcoinOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuraWrappedBitcoin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBitcoinOwnershipTransferred)
				if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) ParseOwnershipTransferred(log types.Log) (*AuraWrappedBitcoinOwnershipTransferred, error) {
	event := new(AuraWrappedBitcoinOwnershipTransferred)
	if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraWrappedBitcoinPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinPausedIterator struct {
	Event *AuraWrappedBitcoinPaused // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBitcoinPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBitcoinPaused)
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
		it.Event = new(AuraWrappedBitcoinPaused)
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
func (it *AuraWrappedBitcoinPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBitcoinPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBitcoinPaused represents a Paused event raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) FilterPaused(opts *bind.FilterOpts) (*AuraWrappedBitcoinPausedIterator, error) {

	logs, sub, err := _AuraWrappedBitcoin.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBitcoinPausedIterator{contract: _AuraWrappedBitcoin.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AuraWrappedBitcoinPaused) (event.Subscription, error) {

	logs, sub, err := _AuraWrappedBitcoin.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBitcoinPaused)
				if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) ParsePaused(log types.Log) (*AuraWrappedBitcoinPaused, error) {
	event := new(AuraWrappedBitcoinPaused)
	if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraWrappedBitcoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinTransferIterator struct {
	Event *AuraWrappedBitcoinTransfer // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBitcoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBitcoinTransfer)
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
		it.Event = new(AuraWrappedBitcoinTransfer)
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
func (it *AuraWrappedBitcoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBitcoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBitcoinTransfer represents a Transfer event raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AuraWrappedBitcoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AuraWrappedBitcoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBitcoinTransferIterator{contract: _AuraWrappedBitcoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AuraWrappedBitcoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AuraWrappedBitcoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBitcoinTransfer)
				if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) ParseTransfer(log types.Log) (*AuraWrappedBitcoinTransfer, error) {
	event := new(AuraWrappedBitcoinTransfer)
	if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraWrappedBitcoinUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinUnpausedIterator struct {
	Event *AuraWrappedBitcoinUnpaused // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBitcoinUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBitcoinUnpaused)
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
		it.Event = new(AuraWrappedBitcoinUnpaused)
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
func (it *AuraWrappedBitcoinUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBitcoinUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBitcoinUnpaused represents a Unpaused event raised by the AuraWrappedBitcoin contract.
type AuraWrappedBitcoinUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AuraWrappedBitcoinUnpausedIterator, error) {

	logs, sub, err := _AuraWrappedBitcoin.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBitcoinUnpausedIterator{contract: _AuraWrappedBitcoin.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AuraWrappedBitcoinUnpaused) (event.Subscription, error) {

	logs, sub, err := _AuraWrappedBitcoin.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBitcoinUnpaused)
				if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AuraWrappedBitcoin *AuraWrappedBitcoinFilterer) ParseUnpaused(log types.Log) (*AuraWrappedBitcoinUnpaused, error) {
	event := new(AuraWrappedBitcoinUnpaused)
	if err := _AuraWrappedBitcoin.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
