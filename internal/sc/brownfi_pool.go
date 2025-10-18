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

// BrownFiPoolMetaData contains all meta data concerning the BrownFiPool contract.
var BrownFiPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Q64\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"k\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lambda\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_fee\",\"type\":\"uint32\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_k\",\"type\":\"uint256\"}],\"name\":\"setK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_lambda\",\"type\":\"uint64\"}],\"name\":\"setLambda\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_protocolFee\",\"type\":\"uint32\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0Decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1Decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BrownFiPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use BrownFiPoolMetaData.ABI instead.
var BrownFiPoolABI = BrownFiPoolMetaData.ABI

// BrownFiPool is an auto generated Go binding around an Ethereum contract.
type BrownFiPool struct {
	BrownFiPoolCaller     // Read-only binding to the contract
	BrownFiPoolTransactor // Write-only binding to the contract
	BrownFiPoolFilterer   // Log filterer for contract events
}

// BrownFiPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BrownFiPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrownFiPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BrownFiPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrownFiPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BrownFiPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BrownFiPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BrownFiPoolSession struct {
	Contract     *BrownFiPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BrownFiPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BrownFiPoolCallerSession struct {
	Contract *BrownFiPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BrownFiPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BrownFiPoolTransactorSession struct {
	Contract     *BrownFiPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BrownFiPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BrownFiPoolRaw struct {
	Contract *BrownFiPool // Generic contract binding to access the raw methods on
}

// BrownFiPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BrownFiPoolCallerRaw struct {
	Contract *BrownFiPoolCaller // Generic read-only contract binding to access the raw methods on
}

// BrownFiPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BrownFiPoolTransactorRaw struct {
	Contract *BrownFiPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBrownFiPool creates a new instance of BrownFiPool, bound to a specific deployed contract.
func NewBrownFiPool(address common.Address, backend bind.ContractBackend) (*BrownFiPool, error) {
	contract, err := bindBrownFiPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BrownFiPool{BrownFiPoolCaller: BrownFiPoolCaller{contract: contract}, BrownFiPoolTransactor: BrownFiPoolTransactor{contract: contract}, BrownFiPoolFilterer: BrownFiPoolFilterer{contract: contract}}, nil
}

// NewBrownFiPoolCaller creates a new read-only instance of BrownFiPool, bound to a specific deployed contract.
func NewBrownFiPoolCaller(address common.Address, caller bind.ContractCaller) (*BrownFiPoolCaller, error) {
	contract, err := bindBrownFiPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BrownFiPoolCaller{contract: contract}, nil
}

// NewBrownFiPoolTransactor creates a new write-only instance of BrownFiPool, bound to a specific deployed contract.
func NewBrownFiPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BrownFiPoolTransactor, error) {
	contract, err := bindBrownFiPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BrownFiPoolTransactor{contract: contract}, nil
}

// NewBrownFiPoolFilterer creates a new log filterer instance of BrownFiPool, bound to a specific deployed contract.
func NewBrownFiPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BrownFiPoolFilterer, error) {
	contract, err := bindBrownFiPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BrownFiPoolFilterer{contract: contract}, nil
}

// bindBrownFiPool binds a generic wrapper to an already deployed contract.
func bindBrownFiPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BrownFiPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrownFiPool *BrownFiPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrownFiPool.Contract.BrownFiPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrownFiPool *BrownFiPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrownFiPool.Contract.BrownFiPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrownFiPool *BrownFiPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrownFiPool.Contract.BrownFiPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BrownFiPool *BrownFiPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BrownFiPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BrownFiPool *BrownFiPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrownFiPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BrownFiPool *BrownFiPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BrownFiPool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BrownFiPool *BrownFiPoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BrownFiPool *BrownFiPoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BrownFiPool.Contract.DOMAINSEPARATOR(&_BrownFiPool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BrownFiPool *BrownFiPoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BrownFiPool.Contract.DOMAINSEPARATOR(&_BrownFiPool.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_BrownFiPool *BrownFiPoolCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_BrownFiPool *BrownFiPoolSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _BrownFiPool.Contract.MINIMUMLIQUIDITY(&_BrownFiPool.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_BrownFiPool *BrownFiPoolCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _BrownFiPool.Contract.MINIMUMLIQUIDITY(&_BrownFiPool.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_BrownFiPool *BrownFiPoolCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_BrownFiPool *BrownFiPoolSession) PERMITTYPEHASH() ([32]byte, error) {
	return _BrownFiPool.Contract.PERMITTYPEHASH(&_BrownFiPool.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_BrownFiPool *BrownFiPoolCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _BrownFiPool.Contract.PERMITTYPEHASH(&_BrownFiPool.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint32)
func (_BrownFiPool *BrownFiPoolCaller) PRECISION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint32)
func (_BrownFiPool *BrownFiPoolSession) PRECISION() (uint32, error) {
	return _BrownFiPool.Contract.PRECISION(&_BrownFiPool.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint32)
func (_BrownFiPool *BrownFiPoolCallerSession) PRECISION() (uint32, error) {
	return _BrownFiPool.Contract.PRECISION(&_BrownFiPool.CallOpts)
}

// Q64 is a free data retrieval call binding the contract method 0x33fca187.
//
// Solidity: function Q64() view returns(uint256)
func (_BrownFiPool *BrownFiPoolCaller) Q64(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "Q64")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Q64 is a free data retrieval call binding the contract method 0x33fca187.
//
// Solidity: function Q64() view returns(uint256)
func (_BrownFiPool *BrownFiPoolSession) Q64() (*big.Int, error) {
	return _BrownFiPool.Contract.Q64(&_BrownFiPool.CallOpts)
}

// Q64 is a free data retrieval call binding the contract method 0x33fca187.
//
// Solidity: function Q64() view returns(uint256)
func (_BrownFiPool *BrownFiPoolCallerSession) Q64() (*big.Int, error) {
	return _BrownFiPool.Contract.Q64(&_BrownFiPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_BrownFiPool *BrownFiPoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_BrownFiPool *BrownFiPoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BrownFiPool.Contract.Allowance(&_BrownFiPool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_BrownFiPool *BrownFiPoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BrownFiPool.Contract.Allowance(&_BrownFiPool.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BrownFiPool *BrownFiPoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BrownFiPool *BrownFiPoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BrownFiPool.Contract.BalanceOf(&_BrownFiPool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BrownFiPool *BrownFiPoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BrownFiPool.Contract.BalanceOf(&_BrownFiPool.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BrownFiPool *BrownFiPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BrownFiPool *BrownFiPoolSession) Decimals() (uint8, error) {
	return _BrownFiPool.Contract.Decimals(&_BrownFiPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BrownFiPool *BrownFiPoolCallerSession) Decimals() (uint8, error) {
	return _BrownFiPool.Contract.Decimals(&_BrownFiPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_BrownFiPool *BrownFiPoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_BrownFiPool *BrownFiPoolSession) Factory() (common.Address, error) {
	return _BrownFiPool.Contract.Factory(&_BrownFiPool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_BrownFiPool *BrownFiPoolCallerSession) Factory() (common.Address, error) {
	return _BrownFiPool.Contract.Factory(&_BrownFiPool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint32)
func (_BrownFiPool *BrownFiPoolCaller) Fee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint32)
func (_BrownFiPool *BrownFiPoolSession) Fee() (uint32, error) {
	return _BrownFiPool.Contract.Fee(&_BrownFiPool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint32)
func (_BrownFiPool *BrownFiPoolCallerSession) Fee() (uint32, error) {
	return _BrownFiPool.Contract.Fee(&_BrownFiPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_BrownFiPool *BrownFiPoolCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_BrownFiPool *BrownFiPoolSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _BrownFiPool.Contract.GetReserves(&_BrownFiPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_BrownFiPool *BrownFiPoolCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _BrownFiPool.Contract.GetReserves(&_BrownFiPool.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xb4f40c61.
//
// Solidity: function k() view returns(uint256)
func (_BrownFiPool *BrownFiPoolCaller) K(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "k")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// K is a free data retrieval call binding the contract method 0xb4f40c61.
//
// Solidity: function k() view returns(uint256)
func (_BrownFiPool *BrownFiPoolSession) K() (*big.Int, error) {
	return _BrownFiPool.Contract.K(&_BrownFiPool.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xb4f40c61.
//
// Solidity: function k() view returns(uint256)
func (_BrownFiPool *BrownFiPoolCallerSession) K() (*big.Int, error) {
	return _BrownFiPool.Contract.K(&_BrownFiPool.CallOpts)
}

// Lambda is a free data retrieval call binding the contract method 0xdad0be61.
//
// Solidity: function lambda() view returns(uint64)
func (_BrownFiPool *BrownFiPoolCaller) Lambda(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "lambda")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Lambda is a free data retrieval call binding the contract method 0xdad0be61.
//
// Solidity: function lambda() view returns(uint64)
func (_BrownFiPool *BrownFiPoolSession) Lambda() (uint64, error) {
	return _BrownFiPool.Contract.Lambda(&_BrownFiPool.CallOpts)
}

// Lambda is a free data retrieval call binding the contract method 0xdad0be61.
//
// Solidity: function lambda() view returns(uint64)
func (_BrownFiPool *BrownFiPoolCallerSession) Lambda() (uint64, error) {
	return _BrownFiPool.Contract.Lambda(&_BrownFiPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BrownFiPool *BrownFiPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BrownFiPool *BrownFiPoolSession) Name() (string, error) {
	return _BrownFiPool.Contract.Name(&_BrownFiPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BrownFiPool *BrownFiPoolCallerSession) Name() (string, error) {
	return _BrownFiPool.Contract.Name(&_BrownFiPool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_BrownFiPool *BrownFiPoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_BrownFiPool *BrownFiPoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _BrownFiPool.Contract.Nonces(&_BrownFiPool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_BrownFiPool *BrownFiPoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _BrownFiPool.Contract.Nonces(&_BrownFiPool.CallOpts, arg0)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint64)
func (_BrownFiPool *BrownFiPoolCaller) ProtocolFee(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint64)
func (_BrownFiPool *BrownFiPoolSession) ProtocolFee() (uint64, error) {
	return _BrownFiPool.Contract.ProtocolFee(&_BrownFiPool.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint64)
func (_BrownFiPool *BrownFiPoolCallerSession) ProtocolFee() (uint64, error) {
	return _BrownFiPool.Contract.ProtocolFee(&_BrownFiPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BrownFiPool *BrownFiPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BrownFiPool *BrownFiPoolSession) Symbol() (string, error) {
	return _BrownFiPool.Contract.Symbol(&_BrownFiPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BrownFiPool *BrownFiPoolCallerSession) Symbol() (string, error) {
	return _BrownFiPool.Contract.Symbol(&_BrownFiPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_BrownFiPool *BrownFiPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_BrownFiPool *BrownFiPoolSession) Token0() (common.Address, error) {
	return _BrownFiPool.Contract.Token0(&_BrownFiPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_BrownFiPool *BrownFiPoolCallerSession) Token0() (common.Address, error) {
	return _BrownFiPool.Contract.Token0(&_BrownFiPool.CallOpts)
}

// Token0Decimals is a free data retrieval call binding the contract method 0xb31ac6e2.
//
// Solidity: function token0Decimals() view returns(uint8)
func (_BrownFiPool *BrownFiPoolCaller) Token0Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "token0Decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Token0Decimals is a free data retrieval call binding the contract method 0xb31ac6e2.
//
// Solidity: function token0Decimals() view returns(uint8)
func (_BrownFiPool *BrownFiPoolSession) Token0Decimals() (uint8, error) {
	return _BrownFiPool.Contract.Token0Decimals(&_BrownFiPool.CallOpts)
}

// Token0Decimals is a free data retrieval call binding the contract method 0xb31ac6e2.
//
// Solidity: function token0Decimals() view returns(uint8)
func (_BrownFiPool *BrownFiPoolCallerSession) Token0Decimals() (uint8, error) {
	return _BrownFiPool.Contract.Token0Decimals(&_BrownFiPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_BrownFiPool *BrownFiPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_BrownFiPool *BrownFiPoolSession) Token1() (common.Address, error) {
	return _BrownFiPool.Contract.Token1(&_BrownFiPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_BrownFiPool *BrownFiPoolCallerSession) Token1() (common.Address, error) {
	return _BrownFiPool.Contract.Token1(&_BrownFiPool.CallOpts)
}

// Token1Decimals is a free data retrieval call binding the contract method 0x0b77884d.
//
// Solidity: function token1Decimals() view returns(uint8)
func (_BrownFiPool *BrownFiPoolCaller) Token1Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "token1Decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Token1Decimals is a free data retrieval call binding the contract method 0x0b77884d.
//
// Solidity: function token1Decimals() view returns(uint8)
func (_BrownFiPool *BrownFiPoolSession) Token1Decimals() (uint8, error) {
	return _BrownFiPool.Contract.Token1Decimals(&_BrownFiPool.CallOpts)
}

// Token1Decimals is a free data retrieval call binding the contract method 0x0b77884d.
//
// Solidity: function token1Decimals() view returns(uint8)
func (_BrownFiPool *BrownFiPoolCallerSession) Token1Decimals() (uint8, error) {
	return _BrownFiPool.Contract.Token1Decimals(&_BrownFiPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BrownFiPool *BrownFiPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BrownFiPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BrownFiPool *BrownFiPoolSession) TotalSupply() (*big.Int, error) {
	return _BrownFiPool.Contract.TotalSupply(&_BrownFiPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BrownFiPool *BrownFiPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _BrownFiPool.Contract.TotalSupply(&_BrownFiPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BrownFiPool *BrownFiPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BrownFiPool *BrownFiPoolSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Approve(&_BrownFiPool.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BrownFiPool *BrownFiPoolTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Approve(&_BrownFiPool.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_BrownFiPool *BrownFiPoolTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_BrownFiPool *BrownFiPoolSession) Burn(to common.Address) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Burn(&_BrownFiPool.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_BrownFiPool *BrownFiPoolTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Burn(&_BrownFiPool.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_BrownFiPool *BrownFiPoolTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_BrownFiPool *BrownFiPoolSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Initialize(&_BrownFiPool.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_BrownFiPool *BrownFiPoolTransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Initialize(&_BrownFiPool.TransactOpts, _token0, _token1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_BrownFiPool *BrownFiPoolTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_BrownFiPool *BrownFiPoolSession) Mint(to common.Address) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Mint(&_BrownFiPool.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_BrownFiPool *BrownFiPoolTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Mint(&_BrownFiPool.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BrownFiPool *BrownFiPoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BrownFiPool *BrownFiPoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Permit(&_BrownFiPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BrownFiPool *BrownFiPoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Permit(&_BrownFiPool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// SetFee is a paid mutator transaction binding the contract method 0x1ab971ab.
//
// Solidity: function setFee(uint32 _fee) returns()
func (_BrownFiPool *BrownFiPoolTransactor) SetFee(opts *bind.TransactOpts, _fee uint32) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "setFee", _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x1ab971ab.
//
// Solidity: function setFee(uint32 _fee) returns()
func (_BrownFiPool *BrownFiPoolSession) SetFee(_fee uint32) (*types.Transaction, error) {
	return _BrownFiPool.Contract.SetFee(&_BrownFiPool.TransactOpts, _fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x1ab971ab.
//
// Solidity: function setFee(uint32 _fee) returns()
func (_BrownFiPool *BrownFiPoolTransactorSession) SetFee(_fee uint32) (*types.Transaction, error) {
	return _BrownFiPool.Contract.SetFee(&_BrownFiPool.TransactOpts, _fee)
}

// SetK is a paid mutator transaction binding the contract method 0x67de8be9.
//
// Solidity: function setK(uint256 _k) returns()
func (_BrownFiPool *BrownFiPoolTransactor) SetK(opts *bind.TransactOpts, _k *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "setK", _k)
}

// SetK is a paid mutator transaction binding the contract method 0x67de8be9.
//
// Solidity: function setK(uint256 _k) returns()
func (_BrownFiPool *BrownFiPoolSession) SetK(_k *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.Contract.SetK(&_BrownFiPool.TransactOpts, _k)
}

// SetK is a paid mutator transaction binding the contract method 0x67de8be9.
//
// Solidity: function setK(uint256 _k) returns()
func (_BrownFiPool *BrownFiPoolTransactorSession) SetK(_k *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.Contract.SetK(&_BrownFiPool.TransactOpts, _k)
}

// SetLambda is a paid mutator transaction binding the contract method 0xcde5bf53.
//
// Solidity: function setLambda(uint64 _lambda) returns()
func (_BrownFiPool *BrownFiPoolTransactor) SetLambda(opts *bind.TransactOpts, _lambda uint64) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "setLambda", _lambda)
}

// SetLambda is a paid mutator transaction binding the contract method 0xcde5bf53.
//
// Solidity: function setLambda(uint64 _lambda) returns()
func (_BrownFiPool *BrownFiPoolSession) SetLambda(_lambda uint64) (*types.Transaction, error) {
	return _BrownFiPool.Contract.SetLambda(&_BrownFiPool.TransactOpts, _lambda)
}

// SetLambda is a paid mutator transaction binding the contract method 0xcde5bf53.
//
// Solidity: function setLambda(uint64 _lambda) returns()
func (_BrownFiPool *BrownFiPoolTransactorSession) SetLambda(_lambda uint64) (*types.Transaction, error) {
	return _BrownFiPool.Contract.SetLambda(&_BrownFiPool.TransactOpts, _lambda)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0xaa9f7628.
//
// Solidity: function setProtocolFee(uint32 _protocolFee) returns()
func (_BrownFiPool *BrownFiPoolTransactor) SetProtocolFee(opts *bind.TransactOpts, _protocolFee uint32) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "setProtocolFee", _protocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0xaa9f7628.
//
// Solidity: function setProtocolFee(uint32 _protocolFee) returns()
func (_BrownFiPool *BrownFiPoolSession) SetProtocolFee(_protocolFee uint32) (*types.Transaction, error) {
	return _BrownFiPool.Contract.SetProtocolFee(&_BrownFiPool.TransactOpts, _protocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0xaa9f7628.
//
// Solidity: function setProtocolFee(uint32 _protocolFee) returns()
func (_BrownFiPool *BrownFiPoolTransactorSession) SetProtocolFee(_protocolFee uint32) (*types.Transaction, error) {
	return _BrownFiPool.Contract.SetProtocolFee(&_BrownFiPool.TransactOpts, _protocolFee)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_BrownFiPool *BrownFiPoolTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_BrownFiPool *BrownFiPoolSession) Skim(to common.Address) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Skim(&_BrownFiPool.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_BrownFiPool *BrownFiPoolTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Skim(&_BrownFiPool.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_BrownFiPool *BrownFiPoolTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_BrownFiPool *BrownFiPoolSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Swap(&_BrownFiPool.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_BrownFiPool *BrownFiPoolTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Swap(&_BrownFiPool.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_BrownFiPool *BrownFiPoolTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_BrownFiPool *BrownFiPoolSession) Sync() (*types.Transaction, error) {
	return _BrownFiPool.Contract.Sync(&_BrownFiPool.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_BrownFiPool *BrownFiPoolTransactorSession) Sync() (*types.Transaction, error) {
	return _BrownFiPool.Contract.Sync(&_BrownFiPool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BrownFiPool *BrownFiPoolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BrownFiPool *BrownFiPoolSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Transfer(&_BrownFiPool.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BrownFiPool *BrownFiPoolTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.Contract.Transfer(&_BrownFiPool.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BrownFiPool *BrownFiPoolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BrownFiPool *BrownFiPoolSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.Contract.TransferFrom(&_BrownFiPool.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BrownFiPool *BrownFiPoolTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BrownFiPool.Contract.TransferFrom(&_BrownFiPool.TransactOpts, from, to, value)
}

// BrownFiPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BrownFiPool contract.
type BrownFiPoolApprovalIterator struct {
	Event *BrownFiPoolApproval // Event containing the contract specifics and raw log

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
func (it *BrownFiPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrownFiPoolApproval)
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
		it.Event = new(BrownFiPoolApproval)
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
func (it *BrownFiPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrownFiPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrownFiPoolApproval represents a Approval event raised by the BrownFiPool contract.
type BrownFiPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BrownFiPool *BrownFiPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BrownFiPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BrownFiPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BrownFiPoolApprovalIterator{contract: _BrownFiPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BrownFiPool *BrownFiPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BrownFiPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BrownFiPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrownFiPoolApproval)
				if err := _BrownFiPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BrownFiPool *BrownFiPoolFilterer) ParseApproval(log types.Log) (*BrownFiPoolApproval, error) {
	event := new(BrownFiPoolApproval)
	if err := _BrownFiPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrownFiPoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BrownFiPool contract.
type BrownFiPoolBurnIterator struct {
	Event *BrownFiPoolBurn // Event containing the contract specifics and raw log

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
func (it *BrownFiPoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrownFiPoolBurn)
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
		it.Event = new(BrownFiPoolBurn)
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
func (it *BrownFiPoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrownFiPoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrownFiPoolBurn represents a Burn event raised by the BrownFiPool contract.
type BrownFiPoolBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_BrownFiPool *BrownFiPoolFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*BrownFiPoolBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BrownFiPool.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BrownFiPoolBurnIterator{contract: _BrownFiPool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_BrownFiPool *BrownFiPoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BrownFiPoolBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BrownFiPool.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrownFiPoolBurn)
				if err := _BrownFiPool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_BrownFiPool *BrownFiPoolFilterer) ParseBurn(log types.Log) (*BrownFiPoolBurn, error) {
	event := new(BrownFiPoolBurn)
	if err := _BrownFiPool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrownFiPoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the BrownFiPool contract.
type BrownFiPoolMintIterator struct {
	Event *BrownFiPoolMint // Event containing the contract specifics and raw log

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
func (it *BrownFiPoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrownFiPoolMint)
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
		it.Event = new(BrownFiPoolMint)
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
func (it *BrownFiPoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrownFiPoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrownFiPoolMint represents a Mint event raised by the BrownFiPool contract.
type BrownFiPoolMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Price0  *big.Int
	Price1  *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x265ee4cff6cdf714e68c02e61a7864cf66bc04e372a41b6cc425acbb737cd395.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 price0, uint256 price1, address indexed to)
func (_BrownFiPool *BrownFiPoolFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*BrownFiPoolMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BrownFiPool.contract.FilterLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BrownFiPoolMintIterator{contract: _BrownFiPool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x265ee4cff6cdf714e68c02e61a7864cf66bc04e372a41b6cc425acbb737cd395.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 price0, uint256 price1, address indexed to)
func (_BrownFiPool *BrownFiPoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *BrownFiPoolMint, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BrownFiPool.contract.WatchLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrownFiPoolMint)
				if err := _BrownFiPool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x265ee4cff6cdf714e68c02e61a7864cf66bc04e372a41b6cc425acbb737cd395.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 price0, uint256 price1, address indexed to)
func (_BrownFiPool *BrownFiPoolFilterer) ParseMint(log types.Log) (*BrownFiPoolMint, error) {
	event := new(BrownFiPoolMint)
	if err := _BrownFiPool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrownFiPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the BrownFiPool contract.
type BrownFiPoolSwapIterator struct {
	Event *BrownFiPoolSwap // Event containing the contract specifics and raw log

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
func (it *BrownFiPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrownFiPoolSwap)
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
		it.Event = new(BrownFiPoolSwap)
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
func (it *BrownFiPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrownFiPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrownFiPoolSwap represents a Swap event raised by the BrownFiPool contract.
type BrownFiPoolSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	Price0     *big.Int
	Price1     *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x72362f8601a26e309bef1da8d6795550a593bd38f121bedcaf24be6e4b7c6bee.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, uint256 price0, uint256 price1, address indexed to)
func (_BrownFiPool *BrownFiPoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*BrownFiPoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BrownFiPool.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BrownFiPoolSwapIterator{contract: _BrownFiPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x72362f8601a26e309bef1da8d6795550a593bd38f121bedcaf24be6e4b7c6bee.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, uint256 price0, uint256 price1, address indexed to)
func (_BrownFiPool *BrownFiPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *BrownFiPoolSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BrownFiPool.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrownFiPoolSwap)
				if err := _BrownFiPool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x72362f8601a26e309bef1da8d6795550a593bd38f121bedcaf24be6e4b7c6bee.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, uint256 price0, uint256 price1, address indexed to)
func (_BrownFiPool *BrownFiPoolFilterer) ParseSwap(log types.Log) (*BrownFiPoolSwap, error) {
	event := new(BrownFiPoolSwap)
	if err := _BrownFiPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrownFiPoolSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the BrownFiPool contract.
type BrownFiPoolSyncIterator struct {
	Event *BrownFiPoolSync // Event containing the contract specifics and raw log

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
func (it *BrownFiPoolSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrownFiPoolSync)
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
		it.Event = new(BrownFiPoolSync)
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
func (it *BrownFiPoolSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrownFiPoolSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrownFiPoolSync represents a Sync event raised by the BrownFiPool contract.
type BrownFiPoolSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_BrownFiPool *BrownFiPoolFilterer) FilterSync(opts *bind.FilterOpts) (*BrownFiPoolSyncIterator, error) {

	logs, sub, err := _BrownFiPool.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &BrownFiPoolSyncIterator{contract: _BrownFiPool.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_BrownFiPool *BrownFiPoolFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *BrownFiPoolSync) (event.Subscription, error) {

	logs, sub, err := _BrownFiPool.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrownFiPoolSync)
				if err := _BrownFiPool.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_BrownFiPool *BrownFiPoolFilterer) ParseSync(log types.Log) (*BrownFiPoolSync, error) {
	event := new(BrownFiPoolSync)
	if err := _BrownFiPool.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BrownFiPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BrownFiPool contract.
type BrownFiPoolTransferIterator struct {
	Event *BrownFiPoolTransfer // Event containing the contract specifics and raw log

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
func (it *BrownFiPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BrownFiPoolTransfer)
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
		it.Event = new(BrownFiPoolTransfer)
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
func (it *BrownFiPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BrownFiPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BrownFiPoolTransfer represents a Transfer event raised by the BrownFiPool contract.
type BrownFiPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BrownFiPool *BrownFiPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BrownFiPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BrownFiPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BrownFiPoolTransferIterator{contract: _BrownFiPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BrownFiPool *BrownFiPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BrownFiPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BrownFiPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BrownFiPoolTransfer)
				if err := _BrownFiPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BrownFiPool *BrownFiPoolFilterer) ParseTransfer(log types.Log) (*BrownFiPoolTransfer, error) {
	event := new(BrownFiPoolTransfer)
	if err := _BrownFiPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
