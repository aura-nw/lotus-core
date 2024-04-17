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

// Brc20GatewayMetaData contains all meta data concerning the Brc20Gateway contract.
var Brc20GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"__decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"__gateway\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gateway_\",\"type\":\"address\"}],\"name\":\"setGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Brc20Gateway *Brc20GatewaySession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Brc20Gateway.Contract.Allowance(&_Brc20Gateway.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Brc20Gateway.Contract.Allowance(&_Brc20Gateway.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Brc20Gateway *Brc20GatewaySession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Brc20Gateway.Contract.BalanceOf(&_Brc20Gateway.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Brc20Gateway.Contract.BalanceOf(&_Brc20Gateway.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Brc20Gateway *Brc20GatewayCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Brc20Gateway *Brc20GatewaySession) Decimals() (uint8, error) {
	return _Brc20Gateway.Contract.Decimals(&_Brc20Gateway.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Brc20Gateway *Brc20GatewayCallerSession) Decimals() (uint8, error) {
	return _Brc20Gateway.Contract.Decimals(&_Brc20Gateway.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Brc20Gateway *Brc20GatewayCaller) Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Brc20Gateway *Brc20GatewaySession) Gateway() (common.Address, error) {
	return _Brc20Gateway.Contract.Gateway(&_Brc20Gateway.CallOpts)
}

// Gateway is a free data retrieval call binding the contract method 0x116191b6.
//
// Solidity: function gateway() view returns(address)
func (_Brc20Gateway *Brc20GatewayCallerSession) Gateway() (common.Address, error) {
	return _Brc20Gateway.Contract.Gateway(&_Brc20Gateway.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Brc20Gateway *Brc20GatewayCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Brc20Gateway *Brc20GatewaySession) Name() (string, error) {
	return _Brc20Gateway.Contract.Name(&_Brc20Gateway.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Brc20Gateway *Brc20GatewayCallerSession) Name() (string, error) {
	return _Brc20Gateway.Contract.Name(&_Brc20Gateway.CallOpts)
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

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Brc20Gateway *Brc20GatewayCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Brc20Gateway *Brc20GatewaySession) Symbol() (string, error) {
	return _Brc20Gateway.Contract.Symbol(&_Brc20Gateway.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Brc20Gateway *Brc20GatewayCallerSession) Symbol() (string, error) {
	return _Brc20Gateway.Contract.Symbol(&_Brc20Gateway.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Brc20Gateway.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brc20Gateway *Brc20GatewaySession) TotalSupply() (*big.Int, error) {
	return _Brc20Gateway.Contract.TotalSupply(&_Brc20Gateway.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Brc20Gateway *Brc20GatewayCallerSession) TotalSupply() (*big.Int, error) {
	return _Brc20Gateway.Contract.TotalSupply(&_Brc20Gateway.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Brc20Gateway *Brc20GatewayTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Brc20Gateway *Brc20GatewaySession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Approve(&_Brc20Gateway.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Brc20Gateway *Brc20GatewayTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Approve(&_Brc20Gateway.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Brc20Gateway *Brc20GatewaySession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Burn(&_Brc20Gateway.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Burn(&_Brc20Gateway.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x6a38eaa9.
//
// Solidity: function initialize(address initialOwner, string name, string symbol, uint8 __decimals, address __gateway) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, name string, symbol string, __decimals uint8, __gateway common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "initialize", initialOwner, name, symbol, __decimals, __gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x6a38eaa9.
//
// Solidity: function initialize(address initialOwner, string name, string symbol, uint8 __decimals, address __gateway) returns()
func (_Brc20Gateway *Brc20GatewaySession) Initialize(initialOwner common.Address, name string, symbol string, __decimals uint8, __gateway common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Initialize(&_Brc20Gateway.TransactOpts, initialOwner, name, symbol, __decimals, __gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x6a38eaa9.
//
// Solidity: function initialize(address initialOwner, string name, string symbol, uint8 __decimals, address __gateway) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) Initialize(initialOwner common.Address, name string, symbol string, __decimals uint8, __gateway common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Initialize(&_Brc20Gateway.TransactOpts, initialOwner, name, symbol, __decimals, __gateway)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount, uint256 invoiceId) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "mint", to, amount, invoiceId)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount, uint256 invoiceId) returns()
func (_Brc20Gateway *Brc20GatewaySession) Mint(to common.Address, amount *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Mint(&_Brc20Gateway.TransactOpts, to, amount, invoiceId)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 amount, uint256 invoiceId) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) Mint(to common.Address, amount *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Mint(&_Brc20Gateway.TransactOpts, to, amount, invoiceId)
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

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address gateway_) returns()
func (_Brc20Gateway *Brc20GatewayTransactor) SetGateway(opts *bind.TransactOpts, gateway_ common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "setGateway", gateway_)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address gateway_) returns()
func (_Brc20Gateway *Brc20GatewaySession) SetGateway(gateway_ common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.SetGateway(&_Brc20Gateway.TransactOpts, gateway_)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address gateway_) returns()
func (_Brc20Gateway *Brc20GatewayTransactorSession) SetGateway(gateway_ common.Address) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.SetGateway(&_Brc20Gateway.TransactOpts, gateway_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Brc20Gateway *Brc20GatewayTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Brc20Gateway *Brc20GatewaySession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Transfer(&_Brc20Gateway.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Brc20Gateway *Brc20GatewayTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.Transfer(&_Brc20Gateway.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Brc20Gateway *Brc20GatewayTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Brc20Gateway *Brc20GatewaySession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.TransferFrom(&_Brc20Gateway.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Brc20Gateway *Brc20GatewayTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Brc20Gateway.Contract.TransferFrom(&_Brc20Gateway.TransactOpts, from, to, value)
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

// Brc20GatewayApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Brc20Gateway contract.
type Brc20GatewayApprovalIterator struct {
	Event *Brc20GatewayApproval // Event containing the contract specifics and raw log

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
func (it *Brc20GatewayApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayApproval)
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
		it.Event = new(Brc20GatewayApproval)
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
func (it *Brc20GatewayApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayApproval represents a Approval event raised by the Brc20Gateway contract.
type Brc20GatewayApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Brc20GatewayApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayApprovalIterator{contract: _Brc20Gateway.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Brc20GatewayApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayApproval)
				if err := _Brc20Gateway.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Brc20Gateway *Brc20GatewayFilterer) ParseApproval(log types.Log) (*Brc20GatewayApproval, error) {
	event := new(Brc20GatewayApproval)
	if err := _Brc20Gateway.contract.UnpackLog(event, "Approval", log); err != nil {
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

// Brc20GatewayTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Brc20Gateway contract.
type Brc20GatewayTransferIterator struct {
	Event *Brc20GatewayTransfer // Event containing the contract specifics and raw log

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
func (it *Brc20GatewayTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Brc20GatewayTransfer)
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
		it.Event = new(Brc20GatewayTransfer)
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
func (it *Brc20GatewayTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Brc20GatewayTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Brc20GatewayTransfer represents a Transfer event raised by the Brc20Gateway contract.
type Brc20GatewayTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Brc20Gateway *Brc20GatewayFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Brc20GatewayTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Brc20Gateway.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Brc20GatewayTransferIterator{contract: _Brc20Gateway.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Brc20Gateway *Brc20GatewayFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Brc20GatewayTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Brc20Gateway.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Brc20GatewayTransfer)
				if err := _Brc20Gateway.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Brc20Gateway *Brc20GatewayFilterer) ParseTransfer(log types.Log) (*Brc20GatewayTransfer, error) {
	event := new(Brc20GatewayTransfer)
	if err := _Brc20Gateway.contract.UnpackLog(event, "Transfer", log); err != nil {
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
