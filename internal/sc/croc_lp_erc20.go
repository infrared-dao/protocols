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

// CrocLPERC20MetaData contains all meta data concerning the CrocLPERC20 contract.
var CrocLPERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CrocLPERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use CrocLPERC20MetaData.ABI instead.
var CrocLPERC20ABI = CrocLPERC20MetaData.ABI

// CrocLPERC20 is an auto generated Go binding around an Ethereum contract.
type CrocLPERC20 struct {
	CrocLPERC20Caller     // Read-only binding to the contract
	CrocLPERC20Transactor // Write-only binding to the contract
	CrocLPERC20Filterer   // Log filterer for contract events
}

// CrocLPERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type CrocLPERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrocLPERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CrocLPERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrocLPERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrocLPERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrocLPERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrocLPERC20Session struct {
	Contract     *CrocLPERC20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrocLPERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrocLPERC20CallerSession struct {
	Contract *CrocLPERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CrocLPERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrocLPERC20TransactorSession struct {
	Contract     *CrocLPERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CrocLPERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type CrocLPERC20Raw struct {
	Contract *CrocLPERC20 // Generic contract binding to access the raw methods on
}

// CrocLPERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrocLPERC20CallerRaw struct {
	Contract *CrocLPERC20Caller // Generic read-only contract binding to access the raw methods on
}

// CrocLPERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrocLPERC20TransactorRaw struct {
	Contract *CrocLPERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCrocLPERC20 creates a new instance of CrocLPERC20, bound to a specific deployed contract.
func NewCrocLPERC20(address common.Address, backend bind.ContractBackend) (*CrocLPERC20, error) {
	contract, err := bindCrocLPERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrocLPERC20{CrocLPERC20Caller: CrocLPERC20Caller{contract: contract}, CrocLPERC20Transactor: CrocLPERC20Transactor{contract: contract}, CrocLPERC20Filterer: CrocLPERC20Filterer{contract: contract}}, nil
}

// NewCrocLPERC20Caller creates a new read-only instance of CrocLPERC20, bound to a specific deployed contract.
func NewCrocLPERC20Caller(address common.Address, caller bind.ContractCaller) (*CrocLPERC20Caller, error) {
	contract, err := bindCrocLPERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrocLPERC20Caller{contract: contract}, nil
}

// NewCrocLPERC20Transactor creates a new write-only instance of CrocLPERC20, bound to a specific deployed contract.
func NewCrocLPERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*CrocLPERC20Transactor, error) {
	contract, err := bindCrocLPERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrocLPERC20Transactor{contract: contract}, nil
}

// NewCrocLPERC20Filterer creates a new log filterer instance of CrocLPERC20, bound to a specific deployed contract.
func NewCrocLPERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*CrocLPERC20Filterer, error) {
	contract, err := bindCrocLPERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrocLPERC20Filterer{contract: contract}, nil
}

// bindCrocLPERC20 binds a generic wrapper to an already deployed contract.
func bindCrocLPERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrocLPERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrocLPERC20 *CrocLPERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrocLPERC20.Contract.CrocLPERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrocLPERC20 *CrocLPERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrocLPERC20.Contract.CrocLPERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrocLPERC20 *CrocLPERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrocLPERC20.Contract.CrocLPERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrocLPERC20 *CrocLPERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrocLPERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrocLPERC20 *CrocLPERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrocLPERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrocLPERC20 *CrocLPERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrocLPERC20.Contract.contract.Transact(opts, method, params...)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_CrocLPERC20 *CrocLPERC20Caller) BaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrocLPERC20.contract.Call(opts, &out, "baseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_CrocLPERC20 *CrocLPERC20Session) BaseToken() (common.Address, error) {
	return _CrocLPERC20.Contract.BaseToken(&_CrocLPERC20.CallOpts)
}

// BaseToken is a free data retrieval call binding the contract method 0xc55dae63.
//
// Solidity: function baseToken() view returns(address)
func (_CrocLPERC20 *CrocLPERC20CallerSession) BaseToken() (common.Address, error) {
	return _CrocLPERC20.Contract.BaseToken(&_CrocLPERC20.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CrocLPERC20 *CrocLPERC20Caller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrocLPERC20.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CrocLPERC20 *CrocLPERC20Session) Factory() (common.Address, error) {
	return _CrocLPERC20.Contract.Factory(&_CrocLPERC20.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_CrocLPERC20 *CrocLPERC20CallerSession) Factory() (common.Address, error) {
	return _CrocLPERC20.Contract.Factory(&_CrocLPERC20.CallOpts)
}

// PoolHash is a free data retrieval call binding the contract method 0x9b503daf.
//
// Solidity: function poolHash() view returns(bytes32)
func (_CrocLPERC20 *CrocLPERC20Caller) PoolHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrocLPERC20.contract.Call(opts, &out, "poolHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PoolHash is a free data retrieval call binding the contract method 0x9b503daf.
//
// Solidity: function poolHash() view returns(bytes32)
func (_CrocLPERC20 *CrocLPERC20Session) PoolHash() ([32]byte, error) {
	return _CrocLPERC20.Contract.PoolHash(&_CrocLPERC20.CallOpts)
}

// PoolHash is a free data retrieval call binding the contract method 0x9b503daf.
//
// Solidity: function poolHash() view returns(bytes32)
func (_CrocLPERC20 *CrocLPERC20CallerSession) PoolHash() ([32]byte, error) {
	return _CrocLPERC20.Contract.PoolHash(&_CrocLPERC20.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint256)
func (_CrocLPERC20 *CrocLPERC20Caller) PoolType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrocLPERC20.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint256)
func (_CrocLPERC20 *CrocLPERC20Session) PoolType() (*big.Int, error) {
	return _CrocLPERC20.Contract.PoolType(&_CrocLPERC20.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint256)
func (_CrocLPERC20 *CrocLPERC20CallerSession) PoolType() (*big.Int, error) {
	return _CrocLPERC20.Contract.PoolType(&_CrocLPERC20.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_CrocLPERC20 *CrocLPERC20Caller) QuoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrocLPERC20.contract.Call(opts, &out, "quoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_CrocLPERC20 *CrocLPERC20Session) QuoteToken() (common.Address, error) {
	return _CrocLPERC20.Contract.QuoteToken(&_CrocLPERC20.CallOpts)
}

// QuoteToken is a free data retrieval call binding the contract method 0x217a4b70.
//
// Solidity: function quoteToken() view returns(address)
func (_CrocLPERC20 *CrocLPERC20CallerSession) QuoteToken() (common.Address, error) {
	return _CrocLPERC20.Contract.QuoteToken(&_CrocLPERC20.CallOpts)
}
