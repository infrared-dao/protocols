// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sc

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

// PendleWrapperMetaData contains all meta data concerning the PendleWrapper contract.
var PendleWrapperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LP\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRewardRedemptionDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"redeemRewardsAndTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isRewardRedemptionDisabled\",\"type\":\"bool\"}],\"name\":\"setRewardRedemptionDisabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netWrapIn\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"netLpIn\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PendleWrapperABI is the input ABI used to generate the binding from.
// Deprecated: Use PendleWrapperMetaData.ABI instead.
var PendleWrapperABI = PendleWrapperMetaData.ABI

// PendleWrapper is an auto generated Go binding around an Ethereum contract.
type PendleWrapper struct {
	PendleWrapperCaller     // Read-only binding to the contract
	PendleWrapperTransactor // Write-only binding to the contract
	PendleWrapperFilterer   // Log filterer for contract events
}

// PendleWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type PendleWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PendleWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleWrapperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PendleWrapperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PendleWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PendleWrapperSession struct {
	Contract     *PendleWrapper    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PendleWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PendleWrapperCallerSession struct {
	Contract *PendleWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PendleWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PendleWrapperTransactorSession struct {
	Contract     *PendleWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PendleWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type PendleWrapperRaw struct {
	Contract *PendleWrapper // Generic contract binding to access the raw methods on
}

// PendleWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PendleWrapperCallerRaw struct {
	Contract *PendleWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// PendleWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PendleWrapperTransactorRaw struct {
	Contract *PendleWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPendleWrapper creates a new instance of PendleWrapper, bound to a specific deployed contract.
func NewPendleWrapper(address common.Address, backend bind.ContractBackend) (*PendleWrapper, error) {
	contract, err := bindPendleWrapper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PendleWrapper{PendleWrapperCaller: PendleWrapperCaller{contract: contract}, PendleWrapperTransactor: PendleWrapperTransactor{contract: contract}, PendleWrapperFilterer: PendleWrapperFilterer{contract: contract}}, nil
}

// NewPendleWrapperCaller creates a new read-only instance of PendleWrapper, bound to a specific deployed contract.
func NewPendleWrapperCaller(address common.Address, caller bind.ContractCaller) (*PendleWrapperCaller, error) {
	contract, err := bindPendleWrapper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PendleWrapperCaller{contract: contract}, nil
}

// NewPendleWrapperTransactor creates a new write-only instance of PendleWrapper, bound to a specific deployed contract.
func NewPendleWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*PendleWrapperTransactor, error) {
	contract, err := bindPendleWrapper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PendleWrapperTransactor{contract: contract}, nil
}

// NewPendleWrapperFilterer creates a new log filterer instance of PendleWrapper, bound to a specific deployed contract.
func NewPendleWrapperFilterer(address common.Address, filterer bind.ContractFilterer) (*PendleWrapperFilterer, error) {
	contract, err := bindPendleWrapper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PendleWrapperFilterer{contract: contract}, nil
}

// bindPendleWrapper binds a generic wrapper to an already deployed contract.
func bindPendleWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PendleWrapperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleWrapper *PendleWrapperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleWrapper.Contract.PendleWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleWrapper *PendleWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleWrapper.Contract.PendleWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleWrapper *PendleWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleWrapper.Contract.PendleWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PendleWrapper *PendleWrapperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PendleWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PendleWrapper *PendleWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PendleWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PendleWrapper *PendleWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PendleWrapper.Contract.contract.Transact(opts, method, params...)
}

// LP is a free data retrieval call binding the contract method 0xb6fccf8a.
//
// Solidity: function LP() view returns(address)
func (_PendleWrapper *PendleWrapperCaller) LP(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleWrapper.contract.Call(opts, &out, "LP")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LP is a free data retrieval call binding the contract method 0xb6fccf8a.
//
// Solidity: function LP() view returns(address)
func (_PendleWrapper *PendleWrapperSession) LP() (common.Address, error) {
	return _PendleWrapper.Contract.LP(&_PendleWrapper.CallOpts)
}

// LP is a free data retrieval call binding the contract method 0xb6fccf8a.
//
// Solidity: function LP() view returns(address)
func (_PendleWrapper *PendleWrapperCallerSession) LP() (common.Address, error) {
	return _PendleWrapper.Contract.LP(&_PendleWrapper.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PendleWrapper *PendleWrapperCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PendleWrapper.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PendleWrapper *PendleWrapperSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PendleWrapper.Contract.Allowance(&_PendleWrapper.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PendleWrapper *PendleWrapperCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PendleWrapper.Contract.Allowance(&_PendleWrapper.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PendleWrapper *PendleWrapperCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PendleWrapper.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PendleWrapper *PendleWrapperSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PendleWrapper.Contract.BalanceOf(&_PendleWrapper.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PendleWrapper *PendleWrapperCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PendleWrapper.Contract.BalanceOf(&_PendleWrapper.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PendleWrapper *PendleWrapperCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PendleWrapper.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PendleWrapper *PendleWrapperSession) Decimals() (uint8, error) {
	return _PendleWrapper.Contract.Decimals(&_PendleWrapper.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PendleWrapper *PendleWrapperCallerSession) Decimals() (uint8, error) {
	return _PendleWrapper.Contract.Decimals(&_PendleWrapper.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PendleWrapper *PendleWrapperCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PendleWrapper.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PendleWrapper *PendleWrapperSession) Factory() (common.Address, error) {
	return _PendleWrapper.Contract.Factory(&_PendleWrapper.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PendleWrapper *PendleWrapperCallerSession) Factory() (common.Address, error) {
	return _PendleWrapper.Contract.Factory(&_PendleWrapper.CallOpts)
}

// IsRewardRedemptionDisabled is a free data retrieval call binding the contract method 0x69c1ca5d.
//
// Solidity: function isRewardRedemptionDisabled() view returns(bool)
func (_PendleWrapper *PendleWrapperCaller) IsRewardRedemptionDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PendleWrapper.contract.Call(opts, &out, "isRewardRedemptionDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewardRedemptionDisabled is a free data retrieval call binding the contract method 0x69c1ca5d.
//
// Solidity: function isRewardRedemptionDisabled() view returns(bool)
func (_PendleWrapper *PendleWrapperSession) IsRewardRedemptionDisabled() (bool, error) {
	return _PendleWrapper.Contract.IsRewardRedemptionDisabled(&_PendleWrapper.CallOpts)
}

// IsRewardRedemptionDisabled is a free data retrieval call binding the contract method 0x69c1ca5d.
//
// Solidity: function isRewardRedemptionDisabled() view returns(bool)
func (_PendleWrapper *PendleWrapperCallerSession) IsRewardRedemptionDisabled() (bool, error) {
	return _PendleWrapper.Contract.IsRewardRedemptionDisabled(&_PendleWrapper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PendleWrapper *PendleWrapperCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PendleWrapper.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PendleWrapper *PendleWrapperSession) Name() (string, error) {
	return _PendleWrapper.Contract.Name(&_PendleWrapper.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PendleWrapper *PendleWrapperCallerSession) Name() (string, error) {
	return _PendleWrapper.Contract.Name(&_PendleWrapper.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PendleWrapper *PendleWrapperCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PendleWrapper.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PendleWrapper *PendleWrapperSession) Symbol() (string, error) {
	return _PendleWrapper.Contract.Symbol(&_PendleWrapper.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PendleWrapper *PendleWrapperCallerSession) Symbol() (string, error) {
	return _PendleWrapper.Contract.Symbol(&_PendleWrapper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PendleWrapper *PendleWrapperCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PendleWrapper.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PendleWrapper *PendleWrapperSession) TotalSupply() (*big.Int, error) {
	return _PendleWrapper.Contract.TotalSupply(&_PendleWrapper.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PendleWrapper *PendleWrapperCallerSession) TotalSupply() (*big.Int, error) {
	return _PendleWrapper.Contract.TotalSupply(&_PendleWrapper.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PendleWrapper *PendleWrapperTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PendleWrapper *PendleWrapperSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.Contract.Approve(&_PendleWrapper.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PendleWrapper *PendleWrapperTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.Contract.Approve(&_PendleWrapper.TransactOpts, spender, amount)
}

// RedeemRewardsAndTransfer is a paid mutator transaction binding the contract method 0x1dfa02a8.
//
// Solidity: function redeemRewardsAndTransfer(address receiver) returns()
func (_PendleWrapper *PendleWrapperTransactor) RedeemRewardsAndTransfer(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _PendleWrapper.contract.Transact(opts, "redeemRewardsAndTransfer", receiver)
}

// RedeemRewardsAndTransfer is a paid mutator transaction binding the contract method 0x1dfa02a8.
//
// Solidity: function redeemRewardsAndTransfer(address receiver) returns()
func (_PendleWrapper *PendleWrapperSession) RedeemRewardsAndTransfer(receiver common.Address) (*types.Transaction, error) {
	return _PendleWrapper.Contract.RedeemRewardsAndTransfer(&_PendleWrapper.TransactOpts, receiver)
}

// RedeemRewardsAndTransfer is a paid mutator transaction binding the contract method 0x1dfa02a8.
//
// Solidity: function redeemRewardsAndTransfer(address receiver) returns()
func (_PendleWrapper *PendleWrapperTransactorSession) RedeemRewardsAndTransfer(receiver common.Address) (*types.Transaction, error) {
	return _PendleWrapper.Contract.RedeemRewardsAndTransfer(&_PendleWrapper.TransactOpts, receiver)
}

// SetRewardRedemptionDisabled is a paid mutator transaction binding the contract method 0xbb0d5de1.
//
// Solidity: function setRewardRedemptionDisabled(bool _isRewardRedemptionDisabled) returns()
func (_PendleWrapper *PendleWrapperTransactor) SetRewardRedemptionDisabled(opts *bind.TransactOpts, _isRewardRedemptionDisabled bool) (*types.Transaction, error) {
	return _PendleWrapper.contract.Transact(opts, "setRewardRedemptionDisabled", _isRewardRedemptionDisabled)
}

// SetRewardRedemptionDisabled is a paid mutator transaction binding the contract method 0xbb0d5de1.
//
// Solidity: function setRewardRedemptionDisabled(bool _isRewardRedemptionDisabled) returns()
func (_PendleWrapper *PendleWrapperSession) SetRewardRedemptionDisabled(_isRewardRedemptionDisabled bool) (*types.Transaction, error) {
	return _PendleWrapper.Contract.SetRewardRedemptionDisabled(&_PendleWrapper.TransactOpts, _isRewardRedemptionDisabled)
}

// SetRewardRedemptionDisabled is a paid mutator transaction binding the contract method 0xbb0d5de1.
//
// Solidity: function setRewardRedemptionDisabled(bool _isRewardRedemptionDisabled) returns()
func (_PendleWrapper *PendleWrapperTransactorSession) SetRewardRedemptionDisabled(_isRewardRedemptionDisabled bool) (*types.Transaction, error) {
	return _PendleWrapper.Contract.SetRewardRedemptionDisabled(&_PendleWrapper.TransactOpts, _isRewardRedemptionDisabled)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PendleWrapper *PendleWrapperTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PendleWrapper *PendleWrapperSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.Contract.Transfer(&_PendleWrapper.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PendleWrapper *PendleWrapperTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.Contract.Transfer(&_PendleWrapper.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PendleWrapper *PendleWrapperTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PendleWrapper *PendleWrapperSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.Contract.TransferFrom(&_PendleWrapper.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_PendleWrapper *PendleWrapperTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.Contract.TransferFrom(&_PendleWrapper.TransactOpts, from, to, amount)
}

// Unwrap is a paid mutator transaction binding the contract method 0x39f47693.
//
// Solidity: function unwrap(address receiver, uint256 netWrapIn) returns()
func (_PendleWrapper *PendleWrapperTransactor) Unwrap(opts *bind.TransactOpts, receiver common.Address, netWrapIn *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.contract.Transact(opts, "unwrap", receiver, netWrapIn)
}

// Unwrap is a paid mutator transaction binding the contract method 0x39f47693.
//
// Solidity: function unwrap(address receiver, uint256 netWrapIn) returns()
func (_PendleWrapper *PendleWrapperSession) Unwrap(receiver common.Address, netWrapIn *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.Contract.Unwrap(&_PendleWrapper.TransactOpts, receiver, netWrapIn)
}

// Unwrap is a paid mutator transaction binding the contract method 0x39f47693.
//
// Solidity: function unwrap(address receiver, uint256 netWrapIn) returns()
func (_PendleWrapper *PendleWrapperTransactorSession) Unwrap(receiver common.Address, netWrapIn *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.Contract.Unwrap(&_PendleWrapper.TransactOpts, receiver, netWrapIn)
}

// Wrap is a paid mutator transaction binding the contract method 0xbf376c7a.
//
// Solidity: function wrap(address receiver, uint256 netLpIn) returns()
func (_PendleWrapper *PendleWrapperTransactor) Wrap(opts *bind.TransactOpts, receiver common.Address, netLpIn *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.contract.Transact(opts, "wrap", receiver, netLpIn)
}

// Wrap is a paid mutator transaction binding the contract method 0xbf376c7a.
//
// Solidity: function wrap(address receiver, uint256 netLpIn) returns()
func (_PendleWrapper *PendleWrapperSession) Wrap(receiver common.Address, netLpIn *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.Contract.Wrap(&_PendleWrapper.TransactOpts, receiver, netLpIn)
}

// Wrap is a paid mutator transaction binding the contract method 0xbf376c7a.
//
// Solidity: function wrap(address receiver, uint256 netLpIn) returns()
func (_PendleWrapper *PendleWrapperTransactorSession) Wrap(receiver common.Address, netLpIn *big.Int) (*types.Transaction, error) {
	return _PendleWrapper.Contract.Wrap(&_PendleWrapper.TransactOpts, receiver, netLpIn)
}

// PendleWrapperApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PendleWrapper contract.
type PendleWrapperApprovalIterator struct {
	Event *PendleWrapperApproval // Event containing the contract specifics and raw log

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
func (it *PendleWrapperApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleWrapperApproval)
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
		it.Event = new(PendleWrapperApproval)
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
func (it *PendleWrapperApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleWrapperApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleWrapperApproval represents a Approval event raised by the PendleWrapper contract.
type PendleWrapperApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PendleWrapper *PendleWrapperFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PendleWrapperApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PendleWrapper.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PendleWrapperApprovalIterator{contract: _PendleWrapper.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PendleWrapper *PendleWrapperFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PendleWrapperApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PendleWrapper.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleWrapperApproval)
				if err := _PendleWrapper.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PendleWrapper *PendleWrapperFilterer) ParseApproval(log types.Log) (*PendleWrapperApproval, error) {
	event := new(PendleWrapperApproval)
	if err := _PendleWrapper.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PendleWrapperTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PendleWrapper contract.
type PendleWrapperTransferIterator struct {
	Event *PendleWrapperTransfer // Event containing the contract specifics and raw log

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
func (it *PendleWrapperTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PendleWrapperTransfer)
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
		it.Event = new(PendleWrapperTransfer)
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
func (it *PendleWrapperTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PendleWrapperTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PendleWrapperTransfer represents a Transfer event raised by the PendleWrapper contract.
type PendleWrapperTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PendleWrapper *PendleWrapperFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PendleWrapperTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PendleWrapper.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PendleWrapperTransferIterator{contract: _PendleWrapper.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PendleWrapper *PendleWrapperFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PendleWrapperTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PendleWrapper.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PendleWrapperTransfer)
				if err := _PendleWrapper.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PendleWrapper *PendleWrapperFilterer) ParseTransfer(log types.Log) (*PendleWrapperTransfer, error) {
	event := new(PendleWrapperTransfer)
	if err := _PendleWrapper.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
