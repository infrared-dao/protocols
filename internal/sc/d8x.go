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

// LiquidityPoolData is an auto generated low-level Go binding around an user-defined struct.
type LiquidityPoolData struct {
	IsRunning                 bool
	IPerpetualCount           uint8
	Id                        uint8
	FCeilPnLShare             int32
	MarginTokenDecimals       uint8
	ITargetPoolSizeUpdateTime uint16
	MarginTokenAddress        common.Address
	PrevAnchor                uint64
	FRedemptionRate           *big.Int
	ShareTokenAddress         common.Address
	FPnLparticipantsCashCC    *big.Int
	FTargetAMMFundSize        *big.Int
	FDefaultFundCashCC        *big.Int
	FTargetDFSize             *big.Int
	FBrokerCollateralLotSize  *big.Int
	PrevTokenAmount           *big.Int
	NextTokenAmount           *big.Int
	TotalSupplyShareToken     *big.Int
	FBrokerFundCashCC         *big.Int
}

// D8xMetaData contains all meta data concerning the D8x contract.
var D8xMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getShareTokenPriceD18\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getLiquidityPool\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRunning\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"iPerpetualCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"fCeilPnLShare\",\"type\":\"int32\"},{\"internalType\":\"uint8\",\"name\":\"marginTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"iTargetPoolSizeUpdateTime\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"marginTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"prevAnchor\",\"type\":\"uint64\"},{\"internalType\":\"int128\",\"name\":\"fRedemptionRate\",\"type\":\"int128\"},{\"internalType\":\"address\",\"name\":\"shareTokenAddress\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fPnLparticipantsCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetAMMFundSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fDefaultFundCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetDFSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fBrokerCollateralLotSize\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"prevTokenAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"nextTokenAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalSupplyShareToken\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"fBrokerFundCashCC\",\"type\":\"int128\"}],\"internalType\":\"structLiquidityPoolData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// D8xABI is the input ABI used to generate the binding from.
// Deprecated: Use D8xMetaData.ABI instead.
var D8xABI = D8xMetaData.ABI

// D8x is an auto generated Go binding around an Ethereum contract.
type D8x struct {
	D8xCaller     // Read-only binding to the contract
	D8xTransactor // Write-only binding to the contract
	D8xFilterer   // Log filterer for contract events
}

// D8xCaller is an auto generated read-only Go binding around an Ethereum contract.
type D8xCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D8xTransactor is an auto generated write-only Go binding around an Ethereum contract.
type D8xTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D8xFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type D8xFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D8xSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type D8xSession struct {
	Contract     *D8x              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// D8xCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type D8xCallerSession struct {
	Contract *D8xCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// D8xTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type D8xTransactorSession struct {
	Contract     *D8xTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// D8xRaw is an auto generated low-level Go binding around an Ethereum contract.
type D8xRaw struct {
	Contract *D8x // Generic contract binding to access the raw methods on
}

// D8xCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type D8xCallerRaw struct {
	Contract *D8xCaller // Generic read-only contract binding to access the raw methods on
}

// D8xTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type D8xTransactorRaw struct {
	Contract *D8xTransactor // Generic write-only contract binding to access the raw methods on
}

// NewD8x creates a new instance of D8x, bound to a specific deployed contract.
func NewD8x(address common.Address, backend bind.ContractBackend) (*D8x, error) {
	contract, err := bindD8x(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &D8x{D8xCaller: D8xCaller{contract: contract}, D8xTransactor: D8xTransactor{contract: contract}, D8xFilterer: D8xFilterer{contract: contract}}, nil
}

// NewD8xCaller creates a new read-only instance of D8x, bound to a specific deployed contract.
func NewD8xCaller(address common.Address, caller bind.ContractCaller) (*D8xCaller, error) {
	contract, err := bindD8x(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &D8xCaller{contract: contract}, nil
}

// NewD8xTransactor creates a new write-only instance of D8x, bound to a specific deployed contract.
func NewD8xTransactor(address common.Address, transactor bind.ContractTransactor) (*D8xTransactor, error) {
	contract, err := bindD8x(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &D8xTransactor{contract: contract}, nil
}

// NewD8xFilterer creates a new log filterer instance of D8x, bound to a specific deployed contract.
func NewD8xFilterer(address common.Address, filterer bind.ContractFilterer) (*D8xFilterer, error) {
	contract, err := bindD8x(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &D8xFilterer{contract: contract}, nil
}

// bindD8x binds a generic wrapper to an already deployed contract.
func bindD8x(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := D8xMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D8x *D8xRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D8x.Contract.D8xCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_D8x *D8xRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D8x.Contract.D8xTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_D8x *D8xRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D8x.Contract.D8xTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D8x *D8xCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D8x.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_D8x *D8xTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D8x.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_D8x *D8xTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D8x.Contract.contract.Transact(opts, method, params...)
}

// GetLiquidityPool is a free data retrieval call binding the contract method 0x30af7c72.
//
// Solidity: function getLiquidityPool(uint8 _poolId) view returns((bool,uint8,uint8,int32,uint8,uint16,address,uint64,int128,address,int128,int128,int128,int128,int128,uint128,uint128,uint128,int128))
func (_D8x *D8xCaller) GetLiquidityPool(opts *bind.CallOpts, _poolId uint8) (LiquidityPoolData, error) {
	var out []interface{}
	err := _D8x.contract.Call(opts, &out, "getLiquidityPool", _poolId)

	if err != nil {
		return *new(LiquidityPoolData), err
	}

	out0 := *abi.ConvertType(out[0], new(LiquidityPoolData)).(*LiquidityPoolData)

	return out0, err

}

// GetLiquidityPool is a free data retrieval call binding the contract method 0x30af7c72.
//
// Solidity: function getLiquidityPool(uint8 _poolId) view returns((bool,uint8,uint8,int32,uint8,uint16,address,uint64,int128,address,int128,int128,int128,int128,int128,uint128,uint128,uint128,int128))
func (_D8x *D8xSession) GetLiquidityPool(_poolId uint8) (LiquidityPoolData, error) {
	return _D8x.Contract.GetLiquidityPool(&_D8x.CallOpts, _poolId)
}

// GetLiquidityPool is a free data retrieval call binding the contract method 0x30af7c72.
//
// Solidity: function getLiquidityPool(uint8 _poolId) view returns((bool,uint8,uint8,int32,uint8,uint16,address,uint64,int128,address,int128,int128,int128,int128,int128,uint128,uint128,uint128,int128))
func (_D8x *D8xCallerSession) GetLiquidityPool(_poolId uint8) (LiquidityPoolData, error) {
	return _D8x.Contract.GetLiquidityPool(&_D8x.CallOpts, _poolId)
}

// GetShareTokenPriceD18 is a free data retrieval call binding the contract method 0x2faee618.
//
// Solidity: function getShareTokenPriceD18(uint8 _poolId) view returns(uint256 price)
func (_D8x *D8xCaller) GetShareTokenPriceD18(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _D8x.contract.Call(opts, &out, "getShareTokenPriceD18", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetShareTokenPriceD18 is a free data retrieval call binding the contract method 0x2faee618.
//
// Solidity: function getShareTokenPriceD18(uint8 _poolId) view returns(uint256 price)
func (_D8x *D8xSession) GetShareTokenPriceD18(_poolId uint8) (*big.Int, error) {
	return _D8x.Contract.GetShareTokenPriceD18(&_D8x.CallOpts, _poolId)
}

// GetShareTokenPriceD18 is a free data retrieval call binding the contract method 0x2faee618.
//
// Solidity: function getShareTokenPriceD18(uint8 _poolId) view returns(uint256 price)
func (_D8x *D8xCallerSession) GetShareTokenPriceD18(_poolId uint8) (*big.Int, error) {
	return _D8x.Contract.GetShareTokenPriceD18(&_D8x.CallOpts, _poolId)
}
