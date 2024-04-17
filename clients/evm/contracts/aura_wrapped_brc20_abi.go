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

// AuraWrappedBrc20MetaData contains all meta data concerning the AuraWrappedBrc20 contract.
var AuraWrappedBrc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"__decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"__gateway\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gateway_\",\"type\":\"address\"}],\"name\":\"setGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AuraWrappedBrc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use AuraWrappedBrc20MetaData.ABI instead.
var AuraWrappedBrc20ABI = AuraWrappedBrc20MetaData.ABI

// AuraWrappedBrc20 is an auto generated Go binding around an Ethereum contract.
type AuraWrappedBrc20 struct {
	AuraWrappedBrc20Caller     // Read-only binding to the contract
	AuraWrappedBrc20Transactor // Write-only binding to the contract
	AuraWrappedBrc20Filterer   // Log filterer for contract events
}

// AuraWrappedBrc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type AuraWrappedBrc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraWrappedBrc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AuraWrappedBrc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraWrappedBrc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuraWrappedBrc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuraWrappedBrc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuraWrappedBrc20Session struct {
	Contract     *AuraWrappedBrc20 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuraWrappedBrc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuraWrappedBrc20CallerSession struct {
	Contract *AuraWrappedBrc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AuraWrappedBrc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuraWrappedBrc20TransactorSession struct {
	Contract     *AuraWrappedBrc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AuraWrappedBrc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type AuraWrappedBrc20Raw struct {
	Contract *AuraWrappedBrc20 // Generic contract binding to access the raw methods on
}

// AuraWrappedBrc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuraWrappedBrc20CallerRaw struct {
	Contract *AuraWrappedBrc20Caller // Generic read-only contract binding to access the raw methods on
}

// AuraWrappedBrc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuraWrappedBrc20TransactorRaw struct {
	Contract *AuraWrappedBrc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAuraWrappedBrc20 creates a new instance of AuraWrappedBrc20, bound to a specific deployed contract.
func NewAuraWrappedBrc20(address common.Address, backend bind.ContractBackend) (*AuraWrappedBrc20, error) {
	contract, err := bindAuraWrappedBrc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBrc20{AuraWrappedBrc20Caller: AuraWrappedBrc20Caller{contract: contract}, AuraWrappedBrc20Transactor: AuraWrappedBrc20Transactor{contract: contract}, AuraWrappedBrc20Filterer: AuraWrappedBrc20Filterer{contract: contract}}, nil
}

// NewAuraWrappedBrc20Caller creates a new read-only instance of AuraWrappedBrc20, bound to a specific deployed contract.
func NewAuraWrappedBrc20Caller(address common.Address, caller bind.ContractCaller) (*AuraWrappedBrc20Caller, error) {
	contract, err := bindAuraWrappedBrc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBrc20Caller{contract: contract}, nil
}

// NewAuraWrappedBrc20Transactor creates a new write-only instance of AuraWrappedBrc20, bound to a specific deployed contract.
func NewAuraWrappedBrc20Transactor(address common.Address, transactor bind.ContractTransactor) (*AuraWrappedBrc20Transactor, error) {
	contract, err := bindAuraWrappedBrc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBrc20Transactor{contract: contract}, nil
}

// NewAuraWrappedBrc20Filterer creates a new log filterer instance of AuraWrappedBrc20, bound to a specific deployed contract.
func NewAuraWrappedBrc20Filterer(address common.Address, filterer bind.ContractFilterer) (*AuraWrappedBrc20Filterer, error) {
	contract, err := bindAuraWrappedBrc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBrc20Filterer{contract: contract}, nil
}

// bindAuraWrappedBrc20 binds a generic wrapper to an already deployed contract.
func bindAuraWrappedBrc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuraWrappedBrc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuraWrappedBrc20 *AuraWrappedBrc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuraWrappedBrc20.Contract.AuraWrappedBrc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuraWrappedBrc20 *AuraWrappedBrc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.AuraWrappedBrc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuraWrappedBrc20 *AuraWrappedBrc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.AuraWrappedBrc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuraWrappedBrc20 *AuraWrappedBrc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuraWrappedBrc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraWrappedBrc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AuraWrappedBrc20.Contract.Allowance(&_AuraWrappedBrc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AuraWrappedBrc20 *AuraWrappedBrc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AuraWrappedBrc20.Contract.Allowance(&_AuraWrappedBrc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AuraWrappedBrc20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _AuraWrappedBrc20.Contract.BalanceOf(&_AuraWrappedBrc20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AuraWrappedBrc20 *AuraWrappedBrc20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AuraWrappedBrc20.Contract.BalanceOf(&_AuraWrappedBrc20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AuraWrappedBrc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Decimals() (uint8, error) {
	return _AuraWrappedBrc20.Contract.Decimals(&_AuraWrappedBrc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AuraWrappedBrc20 *AuraWrappedBrc20CallerSession) Decimals() (uint8, error) {
	return _AuraWrappedBrc20.Contract.Decimals(&_AuraWrappedBrc20.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Caller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraWrappedBrc20.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Gateway() (common.Address, error) {
	return _AuraWrappedBrc20.Contract.Gateway(&_AuraWrappedBrc20.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_AuraWrappedBrc20 *AuraWrappedBrc20CallerSession) Gateway() (common.Address, error) {
	return _AuraWrappedBrc20.Contract.Gateway(&_AuraWrappedBrc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AuraWrappedBrc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Name() (string, error) {
	return _AuraWrappedBrc20.Contract.Name(&_AuraWrappedBrc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AuraWrappedBrc20 *AuraWrappedBrc20CallerSession) Name() (string, error) {
	return _AuraWrappedBrc20.Contract.Name(&_AuraWrappedBrc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuraWrappedBrc20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Owner() (common.Address, error) {
	return _AuraWrappedBrc20.Contract.Owner(&_AuraWrappedBrc20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuraWrappedBrc20 *AuraWrappedBrc20CallerSession) Owner() (common.Address, error) {
	return _AuraWrappedBrc20.Contract.Owner(&_AuraWrappedBrc20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AuraWrappedBrc20.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Paused() (bool, error) {
	return _AuraWrappedBrc20.Contract.Paused(&_AuraWrappedBrc20.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20CallerSession) Paused() (bool, error) {
	return _AuraWrappedBrc20.Contract.Paused(&_AuraWrappedBrc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AuraWrappedBrc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Symbol() (string, error) {
	return _AuraWrappedBrc20.Contract.Symbol(&_AuraWrappedBrc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AuraWrappedBrc20 *AuraWrappedBrc20CallerSession) Symbol() (string, error) {
	return _AuraWrappedBrc20.Contract.Symbol(&_AuraWrappedBrc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuraWrappedBrc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) TotalSupply() (*big.Int, error) {
	return _AuraWrappedBrc20.Contract.TotalSupply(&_AuraWrappedBrc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AuraWrappedBrc20 *AuraWrappedBrc20CallerSession) TotalSupply() (*big.Int, error) {
	return _AuraWrappedBrc20.Contract.TotalSupply(&_AuraWrappedBrc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Approve(&_AuraWrappedBrc20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Approve(&_AuraWrappedBrc20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Burn(&_AuraWrappedBrc20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Burn(&_AuraWrappedBrc20.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x6a38eaa9.
//
// Solidity: function initialize(address initialOwner, string name, string symbol, uint8 __decimals, address __gateway) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, name string, symbol string, __decimals uint8, __gateway common.Address) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "initialize", initialOwner, name, symbol, __decimals, __gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x6a38eaa9.
//
// Solidity: function initialize(address initialOwner, string name, string symbol, uint8 __decimals, address __gateway) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Initialize(initialOwner common.Address, name string, symbol string, __decimals uint8, __gateway common.Address) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Initialize(&_AuraWrappedBrc20.TransactOpts, initialOwner, name, symbol, __decimals, __gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x6a38eaa9.
//
// Solidity: function initialize(address initialOwner, string name, string symbol, uint8 __decimals, address __gateway) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) Initialize(initialOwner common.Address, name string, symbol string, __decimals uint8, __gateway common.Address) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Initialize(&_AuraWrappedBrc20.TransactOpts, initialOwner, name, symbol, __decimals, __gateway)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount, uint256 invoiceId) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "mint", to, amount, invoiceId)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount, uint256 invoiceId) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Mint(to common.Address, amount *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Mint(&_AuraWrappedBrc20.TransactOpts, to, amount, invoiceId)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount, uint256 invoiceId) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) Mint(to common.Address, amount *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Mint(&_AuraWrappedBrc20.TransactOpts, to, amount, invoiceId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Pause() (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Pause(&_AuraWrappedBrc20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) Pause() (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Pause(&_AuraWrappedBrc20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) RenounceOwnership() (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.RenounceOwnership(&_AuraWrappedBrc20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.RenounceOwnership(&_AuraWrappedBrc20.TransactOpts)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address gateway_) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) SetGateway(opts *bind.TransactOpts, gateway_ common.Address) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "setGateway", gateway_)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address gateway_) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) SetGateway(gateway_ common.Address) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.SetGateway(&_AuraWrappedBrc20.TransactOpts, gateway_)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address gateway_) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) SetGateway(gateway_ common.Address) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.SetGateway(&_AuraWrappedBrc20.TransactOpts, gateway_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Transfer(&_AuraWrappedBrc20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Transfer(&_AuraWrappedBrc20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.TransferFrom(&_AuraWrappedBrc20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.TransferFrom(&_AuraWrappedBrc20.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.TransferOwnership(&_AuraWrappedBrc20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.TransferOwnership(&_AuraWrappedBrc20.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuraWrappedBrc20.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20Session) Unpause() (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Unpause(&_AuraWrappedBrc20.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuraWrappedBrc20 *AuraWrappedBrc20TransactorSession) Unpause() (*types.Transaction, error) {
	return _AuraWrappedBrc20.Contract.Unpause(&_AuraWrappedBrc20.TransactOpts)
}

// AuraWrappedBrc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20ApprovalIterator struct {
	Event *AuraWrappedBrc20Approval // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBrc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBrc20Approval)
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
		it.Event = new(AuraWrappedBrc20Approval)
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
func (it *AuraWrappedBrc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBrc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBrc20Approval represents a Approval event raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AuraWrappedBrc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AuraWrappedBrc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBrc20ApprovalIterator{contract: _AuraWrappedBrc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AuraWrappedBrc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AuraWrappedBrc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBrc20Approval)
				if err := _AuraWrappedBrc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) ParseApproval(log types.Log) (*AuraWrappedBrc20Approval, error) {
	event := new(AuraWrappedBrc20Approval)
	if err := _AuraWrappedBrc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraWrappedBrc20InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20InitializedIterator struct {
	Event *AuraWrappedBrc20Initialized // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBrc20InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBrc20Initialized)
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
		it.Event = new(AuraWrappedBrc20Initialized)
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
func (it *AuraWrappedBrc20InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBrc20InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBrc20Initialized represents a Initialized event raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) FilterInitialized(opts *bind.FilterOpts) (*AuraWrappedBrc20InitializedIterator, error) {

	logs, sub, err := _AuraWrappedBrc20.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBrc20InitializedIterator{contract: _AuraWrappedBrc20.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AuraWrappedBrc20Initialized) (event.Subscription, error) {

	logs, sub, err := _AuraWrappedBrc20.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBrc20Initialized)
				if err := _AuraWrappedBrc20.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) ParseInitialized(log types.Log) (*AuraWrappedBrc20Initialized, error) {
	event := new(AuraWrappedBrc20Initialized)
	if err := _AuraWrappedBrc20.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraWrappedBrc20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20OwnershipTransferredIterator struct {
	Event *AuraWrappedBrc20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBrc20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBrc20OwnershipTransferred)
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
		it.Event = new(AuraWrappedBrc20OwnershipTransferred)
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
func (it *AuraWrappedBrc20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBrc20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBrc20OwnershipTransferred represents a OwnershipTransferred event raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuraWrappedBrc20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuraWrappedBrc20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBrc20OwnershipTransferredIterator{contract: _AuraWrappedBrc20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuraWrappedBrc20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuraWrappedBrc20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBrc20OwnershipTransferred)
				if err := _AuraWrappedBrc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) ParseOwnershipTransferred(log types.Log) (*AuraWrappedBrc20OwnershipTransferred, error) {
	event := new(AuraWrappedBrc20OwnershipTransferred)
	if err := _AuraWrappedBrc20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraWrappedBrc20PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20PausedIterator struct {
	Event *AuraWrappedBrc20Paused // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBrc20PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBrc20Paused)
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
		it.Event = new(AuraWrappedBrc20Paused)
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
func (it *AuraWrappedBrc20PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBrc20PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBrc20Paused represents a Paused event raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) FilterPaused(opts *bind.FilterOpts) (*AuraWrappedBrc20PausedIterator, error) {

	logs, sub, err := _AuraWrappedBrc20.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBrc20PausedIterator{contract: _AuraWrappedBrc20.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AuraWrappedBrc20Paused) (event.Subscription, error) {

	logs, sub, err := _AuraWrappedBrc20.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBrc20Paused)
				if err := _AuraWrappedBrc20.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) ParsePaused(log types.Log) (*AuraWrappedBrc20Paused, error) {
	event := new(AuraWrappedBrc20Paused)
	if err := _AuraWrappedBrc20.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraWrappedBrc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20TransferIterator struct {
	Event *AuraWrappedBrc20Transfer // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBrc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBrc20Transfer)
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
		it.Event = new(AuraWrappedBrc20Transfer)
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
func (it *AuraWrappedBrc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBrc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBrc20Transfer represents a Transfer event raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AuraWrappedBrc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AuraWrappedBrc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBrc20TransferIterator{contract: _AuraWrappedBrc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AuraWrappedBrc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AuraWrappedBrc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBrc20Transfer)
				if err := _AuraWrappedBrc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) ParseTransfer(log types.Log) (*AuraWrappedBrc20Transfer, error) {
	event := new(AuraWrappedBrc20Transfer)
	if err := _AuraWrappedBrc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuraWrappedBrc20UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20UnpausedIterator struct {
	Event *AuraWrappedBrc20Unpaused // Event containing the contract specifics and raw log

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
func (it *AuraWrappedBrc20UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuraWrappedBrc20Unpaused)
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
		it.Event = new(AuraWrappedBrc20Unpaused)
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
func (it *AuraWrappedBrc20UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuraWrappedBrc20UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuraWrappedBrc20Unpaused represents a Unpaused event raised by the AuraWrappedBrc20 contract.
type AuraWrappedBrc20Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) FilterUnpaused(opts *bind.FilterOpts) (*AuraWrappedBrc20UnpausedIterator, error) {

	logs, sub, err := _AuraWrappedBrc20.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AuraWrappedBrc20UnpausedIterator{contract: _AuraWrappedBrc20.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AuraWrappedBrc20Unpaused) (event.Subscription, error) {

	logs, sub, err := _AuraWrappedBrc20.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuraWrappedBrc20Unpaused)
				if err := _AuraWrappedBrc20.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AuraWrappedBrc20 *AuraWrappedBrc20Filterer) ParseUnpaused(log types.Log) (*AuraWrappedBrc20Unpaused, error) {
	event := new(AuraWrappedBrc20Unpaused)
	if err := _AuraWrappedBrc20.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
