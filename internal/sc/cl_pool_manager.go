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

// CLPoolManagerMetaData contains all meta data concerning the CLPoolManager contract.
var CLPoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getSlot0\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"uint24\",\"name\":\"protocolFee\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"lpFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getLiquidity\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"poolIdToPoolKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hooks\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolManager\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"parameters\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CLPoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CLPoolManagerMetaData.ABI instead.
var CLPoolManagerABI = CLPoolManagerMetaData.ABI

// CLPoolManager is an auto generated Go binding around an Ethereum contract.
type CLPoolManager struct {
	CLPoolManagerCaller     // Read-only binding to the contract
	CLPoolManagerTransactor // Write-only binding to the contract
	CLPoolManagerFilterer   // Log filterer for contract events
}

// CLPoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CLPoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CLPoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CLPoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CLPoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CLPoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CLPoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CLPoolManagerSession struct {
	Contract     *CLPoolManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CLPoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CLPoolManagerCallerSession struct {
	Contract *CLPoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CLPoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CLPoolManagerTransactorSession struct {
	Contract     *CLPoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CLPoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CLPoolManagerRaw struct {
	Contract *CLPoolManager // Generic contract binding to access the raw methods on
}

// CLPoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CLPoolManagerCallerRaw struct {
	Contract *CLPoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// CLPoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CLPoolManagerTransactorRaw struct {
	Contract *CLPoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCLPoolManager creates a new instance of CLPoolManager, bound to a specific deployed contract.
func NewCLPoolManager(address common.Address, backend bind.ContractBackend) (*CLPoolManager, error) {
	contract, err := bindCLPoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CLPoolManager{CLPoolManagerCaller: CLPoolManagerCaller{contract: contract}, CLPoolManagerTransactor: CLPoolManagerTransactor{contract: contract}, CLPoolManagerFilterer: CLPoolManagerFilterer{contract: contract}}, nil
}

// NewCLPoolManagerCaller creates a new read-only instance of CLPoolManager, bound to a specific deployed contract.
func NewCLPoolManagerCaller(address common.Address, caller bind.ContractCaller) (*CLPoolManagerCaller, error) {
	contract, err := bindCLPoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CLPoolManagerCaller{contract: contract}, nil
}

// NewCLPoolManagerTransactor creates a new write-only instance of CLPoolManager, bound to a specific deployed contract.
func NewCLPoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CLPoolManagerTransactor, error) {
	contract, err := bindCLPoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CLPoolManagerTransactor{contract: contract}, nil
}

// NewCLPoolManagerFilterer creates a new log filterer instance of CLPoolManager, bound to a specific deployed contract.
func NewCLPoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CLPoolManagerFilterer, error) {
	contract, err := bindCLPoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CLPoolManagerFilterer{contract: contract}, nil
}

// bindCLPoolManager binds a generic wrapper to an already deployed contract.
func bindCLPoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CLPoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CLPoolManager *CLPoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CLPoolManager.Contract.CLPoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CLPoolManager *CLPoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CLPoolManager.Contract.CLPoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CLPoolManager *CLPoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CLPoolManager.Contract.CLPoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CLPoolManager *CLPoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CLPoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CLPoolManager *CLPoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CLPoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CLPoolManager *CLPoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CLPoolManager.Contract.contract.Transact(opts, method, params...)
}

// GetLiquidity is a free data retrieval call binding the contract method 0xfa6793d5.
//
// Solidity: function getLiquidity(bytes32 id) view returns(uint128 liquidity)
func (_CLPoolManager *CLPoolManagerCaller) GetLiquidity(opts *bind.CallOpts, id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CLPoolManager.contract.Call(opts, &out, "getLiquidity", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidity is a free data retrieval call binding the contract method 0xfa6793d5.
//
// Solidity: function getLiquidity(bytes32 id) view returns(uint128 liquidity)
func (_CLPoolManager *CLPoolManagerSession) GetLiquidity(id [32]byte) (*big.Int, error) {
	return _CLPoolManager.Contract.GetLiquidity(&_CLPoolManager.CallOpts, id)
}

// GetLiquidity is a free data retrieval call binding the contract method 0xfa6793d5.
//
// Solidity: function getLiquidity(bytes32 id) view returns(uint128 liquidity)
func (_CLPoolManager *CLPoolManagerCallerSession) GetLiquidity(id [32]byte) (*big.Int, error) {
	return _CLPoolManager.Contract.GetLiquidity(&_CLPoolManager.CallOpts, id)
}

// GetSlot0 is a free data retrieval call binding the contract method 0xc815641c.
//
// Solidity: function getSlot0(bytes32 id) view returns(uint160 sqrtPriceX96, int24 tick, uint24 protocolFee, uint24 lpFee)
func (_CLPoolManager *CLPoolManagerCaller) GetSlot0(opts *bind.CallOpts, id [32]byte) (struct {
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	ProtocolFee  *big.Int
	LpFee        *big.Int
}, error) {
	var out []interface{}
	err := _CLPoolManager.contract.Call(opts, &out, "getSlot0", id)

	outstruct := new(struct {
		SqrtPriceX96 *big.Int
		Tick         *big.Int
		ProtocolFee  *big.Int
		LpFee        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProtocolFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LpFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSlot0 is a free data retrieval call binding the contract method 0xc815641c.
//
// Solidity: function getSlot0(bytes32 id) view returns(uint160 sqrtPriceX96, int24 tick, uint24 protocolFee, uint24 lpFee)
func (_CLPoolManager *CLPoolManagerSession) GetSlot0(id [32]byte) (struct {
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	ProtocolFee  *big.Int
	LpFee        *big.Int
}, error) {
	return _CLPoolManager.Contract.GetSlot0(&_CLPoolManager.CallOpts, id)
}

// GetSlot0 is a free data retrieval call binding the contract method 0xc815641c.
//
// Solidity: function getSlot0(bytes32 id) view returns(uint160 sqrtPriceX96, int24 tick, uint24 protocolFee, uint24 lpFee)
func (_CLPoolManager *CLPoolManagerCallerSession) GetSlot0(id [32]byte) (struct {
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	ProtocolFee  *big.Int
	LpFee        *big.Int
}, error) {
	return _CLPoolManager.Contract.GetSlot0(&_CLPoolManager.CallOpts, id)
}

// PoolIdToPoolKey is a free data retrieval call binding the contract method 0x0e2d484a.
//
// Solidity: function poolIdToPoolKey(bytes32 id) view returns(address currency0, address currency1, address hooks, address poolManager, uint24 fee, bytes32 parameters)
func (_CLPoolManager *CLPoolManagerCaller) PoolIdToPoolKey(opts *bind.CallOpts, id [32]byte) (struct {
	Currency0   common.Address
	Currency1   common.Address
	Hooks       common.Address
	PoolManager common.Address
	Fee         *big.Int
	Parameters  [32]byte
}, error) {
	var out []interface{}
	err := _CLPoolManager.contract.Call(opts, &out, "poolIdToPoolKey", id)

	outstruct := new(struct {
		Currency0   common.Address
		Currency1   common.Address
		Hooks       common.Address
		PoolManager common.Address
		Fee         *big.Int
		Parameters  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Currency0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Currency1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Hooks = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.PoolManager = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Parameters = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// PoolIdToPoolKey is a free data retrieval call binding the contract method 0x0e2d484a.
//
// Solidity: function poolIdToPoolKey(bytes32 id) view returns(address currency0, address currency1, address hooks, address poolManager, uint24 fee, bytes32 parameters)
func (_CLPoolManager *CLPoolManagerSession) PoolIdToPoolKey(id [32]byte) (struct {
	Currency0   common.Address
	Currency1   common.Address
	Hooks       common.Address
	PoolManager common.Address
	Fee         *big.Int
	Parameters  [32]byte
}, error) {
	return _CLPoolManager.Contract.PoolIdToPoolKey(&_CLPoolManager.CallOpts, id)
}

// PoolIdToPoolKey is a free data retrieval call binding the contract method 0x0e2d484a.
//
// Solidity: function poolIdToPoolKey(bytes32 id) view returns(address currency0, address currency1, address hooks, address poolManager, uint24 fee, bytes32 parameters)
func (_CLPoolManager *CLPoolManagerCallerSession) PoolIdToPoolKey(id [32]byte) (struct {
	Currency0   common.Address
	Currency1   common.Address
	Hooks       common.Address
	PoolManager common.Address
	Fee         *big.Int
	Parameters  [32]byte
}, error) {
	return _CLPoolManager.Contract.PoolIdToPoolKey(&_CLPoolManager.CallOpts, id)
}
