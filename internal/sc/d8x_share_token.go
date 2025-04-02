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

// D8xShareTokenMetaData contains all meta data concerning the D8xShareToken contract.
var D8xShareTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountD18\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceD18\",\"type\":\"uint256\"}],\"name\":\"P2PTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolId\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"setTransferRestricted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// D8xShareTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use D8xShareTokenMetaData.ABI instead.
var D8xShareTokenABI = D8xShareTokenMetaData.ABI

// D8xShareToken is an auto generated Go binding around an Ethereum contract.
type D8xShareToken struct {
	D8xShareTokenCaller     // Read-only binding to the contract
	D8xShareTokenTransactor // Write-only binding to the contract
	D8xShareTokenFilterer   // Log filterer for contract events
}

// D8xShareTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type D8xShareTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D8xShareTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type D8xShareTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D8xShareTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type D8xShareTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D8xShareTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type D8xShareTokenSession struct {
	Contract     *D8xShareToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// D8xShareTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type D8xShareTokenCallerSession struct {
	Contract *D8xShareTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// D8xShareTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type D8xShareTokenTransactorSession struct {
	Contract     *D8xShareTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// D8xShareTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type D8xShareTokenRaw struct {
	Contract *D8xShareToken // Generic contract binding to access the raw methods on
}

// D8xShareTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type D8xShareTokenCallerRaw struct {
	Contract *D8xShareTokenCaller // Generic read-only contract binding to access the raw methods on
}

// D8xShareTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type D8xShareTokenTransactorRaw struct {
	Contract *D8xShareTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewD8xShareToken creates a new instance of D8xShareToken, bound to a specific deployed contract.
func NewD8xShareToken(address common.Address, backend bind.ContractBackend) (*D8xShareToken, error) {
	contract, err := bindD8xShareToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &D8xShareToken{D8xShareTokenCaller: D8xShareTokenCaller{contract: contract}, D8xShareTokenTransactor: D8xShareTokenTransactor{contract: contract}, D8xShareTokenFilterer: D8xShareTokenFilterer{contract: contract}}, nil
}

// NewD8xShareTokenCaller creates a new read-only instance of D8xShareToken, bound to a specific deployed contract.
func NewD8xShareTokenCaller(address common.Address, caller bind.ContractCaller) (*D8xShareTokenCaller, error) {
	contract, err := bindD8xShareToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &D8xShareTokenCaller{contract: contract}, nil
}

// NewD8xShareTokenTransactor creates a new write-only instance of D8xShareToken, bound to a specific deployed contract.
func NewD8xShareTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*D8xShareTokenTransactor, error) {
	contract, err := bindD8xShareToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &D8xShareTokenTransactor{contract: contract}, nil
}

// NewD8xShareTokenFilterer creates a new log filterer instance of D8xShareToken, bound to a specific deployed contract.
func NewD8xShareTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*D8xShareTokenFilterer, error) {
	contract, err := bindD8xShareToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &D8xShareTokenFilterer{contract: contract}, nil
}

// bindD8xShareToken binds a generic wrapper to an already deployed contract.
func bindD8xShareToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := D8xShareTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D8xShareToken *D8xShareTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D8xShareToken.Contract.D8xShareTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_D8xShareToken *D8xShareTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D8xShareToken.Contract.D8xShareTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_D8xShareToken *D8xShareTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D8xShareToken.Contract.D8xShareTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D8xShareToken *D8xShareTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D8xShareToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_D8xShareToken *D8xShareTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D8xShareToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_D8xShareToken *D8xShareTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D8xShareToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_D8xShareToken *D8xShareTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D8xShareToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_D8xShareToken *D8xShareTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _D8xShareToken.Contract.Allowance(&_D8xShareToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_D8xShareToken *D8xShareTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _D8xShareToken.Contract.Allowance(&_D8xShareToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_D8xShareToken *D8xShareTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D8xShareToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_D8xShareToken *D8xShareTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _D8xShareToken.Contract.BalanceOf(&_D8xShareToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_D8xShareToken *D8xShareTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _D8xShareToken.Contract.BalanceOf(&_D8xShareToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_D8xShareToken *D8xShareTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _D8xShareToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_D8xShareToken *D8xShareTokenSession) Decimals() (uint8, error) {
	return _D8xShareToken.Contract.Decimals(&_D8xShareToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_D8xShareToken *D8xShareTokenCallerSession) Decimals() (uint8, error) {
	return _D8xShareToken.Contract.Decimals(&_D8xShareToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_D8xShareToken *D8xShareTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _D8xShareToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_D8xShareToken *D8xShareTokenSession) Name() (string, error) {
	return _D8xShareToken.Contract.Name(&_D8xShareToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_D8xShareToken *D8xShareTokenCallerSession) Name() (string, error) {
	return _D8xShareToken.Contract.Name(&_D8xShareToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_D8xShareToken *D8xShareTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D8xShareToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_D8xShareToken *D8xShareTokenSession) Owner() (common.Address, error) {
	return _D8xShareToken.Contract.Owner(&_D8xShareToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_D8xShareToken *D8xShareTokenCallerSession) Owner() (common.Address, error) {
	return _D8xShareToken.Contract.Owner(&_D8xShareToken.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_D8xShareToken *D8xShareTokenCaller) PoolId(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _D8xShareToken.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_D8xShareToken *D8xShareTokenSession) PoolId() (uint8, error) {
	return _D8xShareToken.Contract.PoolId(&_D8xShareToken.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint8)
func (_D8xShareToken *D8xShareTokenCallerSession) PoolId() (uint8, error) {
	return _D8xShareToken.Contract.PoolId(&_D8xShareToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_D8xShareToken *D8xShareTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _D8xShareToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_D8xShareToken *D8xShareTokenSession) Symbol() (string, error) {
	return _D8xShareToken.Contract.Symbol(&_D8xShareToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_D8xShareToken *D8xShareTokenCallerSession) Symbol() (string, error) {
	return _D8xShareToken.Contract.Symbol(&_D8xShareToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_D8xShareToken *D8xShareTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D8xShareToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_D8xShareToken *D8xShareTokenSession) TotalSupply() (*big.Int, error) {
	return _D8xShareToken.Contract.TotalSupply(&_D8xShareToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_D8xShareToken *D8xShareTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _D8xShareToken.Contract.TotalSupply(&_D8xShareToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_D8xShareToken *D8xShareTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_D8xShareToken *D8xShareTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.Approve(&_D8xShareToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_D8xShareToken *D8xShareTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.Approve(&_D8xShareToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_D8xShareToken *D8xShareTokenTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_D8xShareToken *D8xShareTokenSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.Burn(&_D8xShareToken.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_D8xShareToken *D8xShareTokenTransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.Burn(&_D8xShareToken.TransactOpts, _account, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_D8xShareToken *D8xShareTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_D8xShareToken *D8xShareTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.DecreaseAllowance(&_D8xShareToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_D8xShareToken *D8xShareTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.DecreaseAllowance(&_D8xShareToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_D8xShareToken *D8xShareTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_D8xShareToken *D8xShareTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.IncreaseAllowance(&_D8xShareToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_D8xShareToken *D8xShareTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.IncreaseAllowance(&_D8xShareToken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_D8xShareToken *D8xShareTokenTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_D8xShareToken *D8xShareTokenSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.Mint(&_D8xShareToken.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_D8xShareToken *D8xShareTokenTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.Mint(&_D8xShareToken.TransactOpts, _account, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_D8xShareToken *D8xShareTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D8xShareToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_D8xShareToken *D8xShareTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _D8xShareToken.Contract.RenounceOwnership(&_D8xShareToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_D8xShareToken *D8xShareTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _D8xShareToken.Contract.RenounceOwnership(&_D8xShareToken.TransactOpts)
}

// SetTransferRestricted is a paid mutator transaction binding the contract method 0x80860c03.
//
// Solidity: function setTransferRestricted(address _account) returns()
func (_D8xShareToken *D8xShareTokenTransactor) SetTransferRestricted(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _D8xShareToken.contract.Transact(opts, "setTransferRestricted", _account)
}

// SetTransferRestricted is a paid mutator transaction binding the contract method 0x80860c03.
//
// Solidity: function setTransferRestricted(address _account) returns()
func (_D8xShareToken *D8xShareTokenSession) SetTransferRestricted(_account common.Address) (*types.Transaction, error) {
	return _D8xShareToken.Contract.SetTransferRestricted(&_D8xShareToken.TransactOpts, _account)
}

// SetTransferRestricted is a paid mutator transaction binding the contract method 0x80860c03.
//
// Solidity: function setTransferRestricted(address _account) returns()
func (_D8xShareToken *D8xShareTokenTransactorSession) SetTransferRestricted(_account common.Address) (*types.Transaction, error) {
	return _D8xShareToken.Contract.SetTransferRestricted(&_D8xShareToken.TransactOpts, _account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_D8xShareToken *D8xShareTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_D8xShareToken *D8xShareTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.Transfer(&_D8xShareToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_D8xShareToken *D8xShareTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.Transfer(&_D8xShareToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_D8xShareToken *D8xShareTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_D8xShareToken *D8xShareTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.TransferFrom(&_D8xShareToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_D8xShareToken *D8xShareTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D8xShareToken.Contract.TransferFrom(&_D8xShareToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_D8xShareToken *D8xShareTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _D8xShareToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_D8xShareToken *D8xShareTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _D8xShareToken.Contract.TransferOwnership(&_D8xShareToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_D8xShareToken *D8xShareTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _D8xShareToken.Contract.TransferOwnership(&_D8xShareToken.TransactOpts, newOwner)
}

// D8xShareTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the D8xShareToken contract.
type D8xShareTokenApprovalIterator struct {
	Event *D8xShareTokenApproval // Event containing the contract specifics and raw log

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
func (it *D8xShareTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D8xShareTokenApproval)
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
		it.Event = new(D8xShareTokenApproval)
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
func (it *D8xShareTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D8xShareTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D8xShareTokenApproval represents a Approval event raised by the D8xShareToken contract.
type D8xShareTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_D8xShareToken *D8xShareTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*D8xShareTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _D8xShareToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &D8xShareTokenApprovalIterator{contract: _D8xShareToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_D8xShareToken *D8xShareTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *D8xShareTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _D8xShareToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D8xShareTokenApproval)
				if err := _D8xShareToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_D8xShareToken *D8xShareTokenFilterer) ParseApproval(log types.Log) (*D8xShareTokenApproval, error) {
	event := new(D8xShareTokenApproval)
	if err := _D8xShareToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D8xShareTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the D8xShareToken contract.
type D8xShareTokenOwnershipTransferredIterator struct {
	Event *D8xShareTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *D8xShareTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D8xShareTokenOwnershipTransferred)
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
		it.Event = new(D8xShareTokenOwnershipTransferred)
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
func (it *D8xShareTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D8xShareTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D8xShareTokenOwnershipTransferred represents a OwnershipTransferred event raised by the D8xShareToken contract.
type D8xShareTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_D8xShareToken *D8xShareTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*D8xShareTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _D8xShareToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &D8xShareTokenOwnershipTransferredIterator{contract: _D8xShareToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_D8xShareToken *D8xShareTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *D8xShareTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _D8xShareToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D8xShareTokenOwnershipTransferred)
				if err := _D8xShareToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_D8xShareToken *D8xShareTokenFilterer) ParseOwnershipTransferred(log types.Log) (*D8xShareTokenOwnershipTransferred, error) {
	event := new(D8xShareTokenOwnershipTransferred)
	if err := _D8xShareToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D8xShareTokenP2PTransferIterator is returned from FilterP2PTransfer and is used to iterate over the raw logs and unpacked data for P2PTransfer events raised by the D8xShareToken contract.
type D8xShareTokenP2PTransferIterator struct {
	Event *D8xShareTokenP2PTransfer // Event containing the contract specifics and raw log

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
func (it *D8xShareTokenP2PTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D8xShareTokenP2PTransfer)
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
		it.Event = new(D8xShareTokenP2PTransfer)
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
func (it *D8xShareTokenP2PTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D8xShareTokenP2PTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D8xShareTokenP2PTransfer represents a P2PTransfer event raised by the D8xShareToken contract.
type D8xShareTokenP2PTransfer struct {
	From      common.Address
	To        common.Address
	AmountD18 *big.Int
	PriceD18  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterP2PTransfer is a free log retrieval operation binding the contract event 0x78e5f93474d987925844427924e2f97e7053fbb290ae237610bc27386d34cdba.
//
// Solidity: event P2PTransfer(address from, address to, uint256 amountD18, uint256 priceD18)
func (_D8xShareToken *D8xShareTokenFilterer) FilterP2PTransfer(opts *bind.FilterOpts) (*D8xShareTokenP2PTransferIterator, error) {

	logs, sub, err := _D8xShareToken.contract.FilterLogs(opts, "P2PTransfer")
	if err != nil {
		return nil, err
	}
	return &D8xShareTokenP2PTransferIterator{contract: _D8xShareToken.contract, event: "P2PTransfer", logs: logs, sub: sub}, nil
}

// WatchP2PTransfer is a free log subscription operation binding the contract event 0x78e5f93474d987925844427924e2f97e7053fbb290ae237610bc27386d34cdba.
//
// Solidity: event P2PTransfer(address from, address to, uint256 amountD18, uint256 priceD18)
func (_D8xShareToken *D8xShareTokenFilterer) WatchP2PTransfer(opts *bind.WatchOpts, sink chan<- *D8xShareTokenP2PTransfer) (event.Subscription, error) {

	logs, sub, err := _D8xShareToken.contract.WatchLogs(opts, "P2PTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D8xShareTokenP2PTransfer)
				if err := _D8xShareToken.contract.UnpackLog(event, "P2PTransfer", log); err != nil {
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

// ParseP2PTransfer is a log parse operation binding the contract event 0x78e5f93474d987925844427924e2f97e7053fbb290ae237610bc27386d34cdba.
//
// Solidity: event P2PTransfer(address from, address to, uint256 amountD18, uint256 priceD18)
func (_D8xShareToken *D8xShareTokenFilterer) ParseP2PTransfer(log types.Log) (*D8xShareTokenP2PTransfer, error) {
	event := new(D8xShareTokenP2PTransfer)
	if err := _D8xShareToken.contract.UnpackLog(event, "P2PTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D8xShareTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the D8xShareToken contract.
type D8xShareTokenTransferIterator struct {
	Event *D8xShareTokenTransfer // Event containing the contract specifics and raw log

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
func (it *D8xShareTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D8xShareTokenTransfer)
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
		it.Event = new(D8xShareTokenTransfer)
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
func (it *D8xShareTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D8xShareTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D8xShareTokenTransfer represents a Transfer event raised by the D8xShareToken contract.
type D8xShareTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_D8xShareToken *D8xShareTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*D8xShareTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _D8xShareToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &D8xShareTokenTransferIterator{contract: _D8xShareToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_D8xShareToken *D8xShareTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *D8xShareTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _D8xShareToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D8xShareTokenTransfer)
				if err := _D8xShareToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_D8xShareToken *D8xShareTokenFilterer) ParseTransfer(log types.Log) (*D8xShareTokenTransfer, error) {
	event := new(D8xShareTokenTransfer)
	if err := _D8xShareToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
