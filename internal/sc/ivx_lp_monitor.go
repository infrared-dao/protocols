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

// IVXLPMonitorMetaData contains all meta data concerning the IVXLPMonitor contract.
var IVXLPMonitorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_beaconAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBeacon\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSharePrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTVL\",\"inputs\":[],\"outputs\":[{\"name\":\"tvl\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"AddressRequired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MathOverflowedMulDiv\",\"inputs\":[]}]",
}

// IVXLPMonitorABI is the input ABI used to generate the binding from.
// Deprecated: Use IVXLPMonitorMetaData.ABI instead.
var IVXLPMonitorABI = IVXLPMonitorMetaData.ABI

// IVXLPMonitor is an auto generated Go binding around an Ethereum contract.
type IVXLPMonitor struct {
	IVXLPMonitorCaller     // Read-only binding to the contract
	IVXLPMonitorTransactor // Write-only binding to the contract
	IVXLPMonitorFilterer   // Log filterer for contract events
}

// IVXLPMonitorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVXLPMonitorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVXLPMonitorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVXLPMonitorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVXLPMonitorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVXLPMonitorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVXLPMonitorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVXLPMonitorSession struct {
	Contract     *IVXLPMonitor     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVXLPMonitorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVXLPMonitorCallerSession struct {
	Contract *IVXLPMonitorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IVXLPMonitorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVXLPMonitorTransactorSession struct {
	Contract     *IVXLPMonitorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IVXLPMonitorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVXLPMonitorRaw struct {
	Contract *IVXLPMonitor // Generic contract binding to access the raw methods on
}

// IVXLPMonitorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVXLPMonitorCallerRaw struct {
	Contract *IVXLPMonitorCaller // Generic read-only contract binding to access the raw methods on
}

// IVXLPMonitorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVXLPMonitorTransactorRaw struct {
	Contract *IVXLPMonitorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVXLPMonitor creates a new instance of IVXLPMonitor, bound to a specific deployed contract.
func NewIVXLPMonitor(address common.Address, backend bind.ContractBackend) (*IVXLPMonitor, error) {
	contract, err := bindIVXLPMonitor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVXLPMonitor{IVXLPMonitorCaller: IVXLPMonitorCaller{contract: contract}, IVXLPMonitorTransactor: IVXLPMonitorTransactor{contract: contract}, IVXLPMonitorFilterer: IVXLPMonitorFilterer{contract: contract}}, nil
}

// NewIVXLPMonitorCaller creates a new read-only instance of IVXLPMonitor, bound to a specific deployed contract.
func NewIVXLPMonitorCaller(address common.Address, caller bind.ContractCaller) (*IVXLPMonitorCaller, error) {
	contract, err := bindIVXLPMonitor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVXLPMonitorCaller{contract: contract}, nil
}

// NewIVXLPMonitorTransactor creates a new write-only instance of IVXLPMonitor, bound to a specific deployed contract.
func NewIVXLPMonitorTransactor(address common.Address, transactor bind.ContractTransactor) (*IVXLPMonitorTransactor, error) {
	contract, err := bindIVXLPMonitor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVXLPMonitorTransactor{contract: contract}, nil
}

// NewIVXLPMonitorFilterer creates a new log filterer instance of IVXLPMonitor, bound to a specific deployed contract.
func NewIVXLPMonitorFilterer(address common.Address, filterer bind.ContractFilterer) (*IVXLPMonitorFilterer, error) {
	contract, err := bindIVXLPMonitor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVXLPMonitorFilterer{contract: contract}, nil
}

// bindIVXLPMonitor binds a generic wrapper to an already deployed contract.
func bindIVXLPMonitor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVXLPMonitorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVXLPMonitor *IVXLPMonitorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVXLPMonitor.Contract.IVXLPMonitorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVXLPMonitor *IVXLPMonitorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVXLPMonitor.Contract.IVXLPMonitorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVXLPMonitor *IVXLPMonitorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVXLPMonitor.Contract.IVXLPMonitorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVXLPMonitor *IVXLPMonitorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVXLPMonitor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVXLPMonitor *IVXLPMonitorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVXLPMonitor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVXLPMonitor *IVXLPMonitorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVXLPMonitor.Contract.contract.Transact(opts, method, params...)
}

// GetBeacon is a free data retrieval call binding the contract method 0x2d6b3a6b.
//
// Solidity: function getBeacon() view returns(address)
func (_IVXLPMonitor *IVXLPMonitorCaller) GetBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IVXLPMonitor.contract.Call(opts, &out, "getBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBeacon is a free data retrieval call binding the contract method 0x2d6b3a6b.
//
// Solidity: function getBeacon() view returns(address)
func (_IVXLPMonitor *IVXLPMonitorSession) GetBeacon() (common.Address, error) {
	return _IVXLPMonitor.Contract.GetBeacon(&_IVXLPMonitor.CallOpts)
}

// GetBeacon is a free data retrieval call binding the contract method 0x2d6b3a6b.
//
// Solidity: function getBeacon() view returns(address)
func (_IVXLPMonitor *IVXLPMonitorCallerSession) GetBeacon() (common.Address, error) {
	return _IVXLPMonitor.Contract.GetBeacon(&_IVXLPMonitor.CallOpts)
}

// GetSharePrice is a free data retrieval call binding the contract method 0x5b1dac60.
//
// Solidity: function getSharePrice() view returns(uint256)
func (_IVXLPMonitor *IVXLPMonitorCaller) GetSharePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVXLPMonitor.contract.Call(opts, &out, "getSharePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharePrice is a free data retrieval call binding the contract method 0x5b1dac60.
//
// Solidity: function getSharePrice() view returns(uint256)
func (_IVXLPMonitor *IVXLPMonitorSession) GetSharePrice() (*big.Int, error) {
	return _IVXLPMonitor.Contract.GetSharePrice(&_IVXLPMonitor.CallOpts)
}

// GetSharePrice is a free data retrieval call binding the contract method 0x5b1dac60.
//
// Solidity: function getSharePrice() view returns(uint256)
func (_IVXLPMonitor *IVXLPMonitorCallerSession) GetSharePrice() (*big.Int, error) {
	return _IVXLPMonitor.Contract.GetSharePrice(&_IVXLPMonitor.CallOpts)
}

// GetTVL is a free data retrieval call binding the contract method 0x97b3fcaa.
//
// Solidity: function getTVL() view returns(uint256 tvl)
func (_IVXLPMonitor *IVXLPMonitorCaller) GetTVL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVXLPMonitor.contract.Call(opts, &out, "getTVL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTVL is a free data retrieval call binding the contract method 0x97b3fcaa.
//
// Solidity: function getTVL() view returns(uint256 tvl)
func (_IVXLPMonitor *IVXLPMonitorSession) GetTVL() (*big.Int, error) {
	return _IVXLPMonitor.Contract.GetTVL(&_IVXLPMonitor.CallOpts)
}

// GetTVL is a free data retrieval call binding the contract method 0x97b3fcaa.
//
// Solidity: function getTVL() view returns(uint256 tvl)
func (_IVXLPMonitor *IVXLPMonitorCallerSession) GetTVL() (*big.Int, error) {
	return _IVXLPMonitor.Contract.GetTVL(&_IVXLPMonitor.CallOpts)
}
