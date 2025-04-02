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

// D8xPoolManagerMetaData contains all meta data concerning the D8xPoolManager contract.
var D8xPoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getShareTokenPriceD18\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_poolId\",\"type\":\"uint8\"}],\"name\":\"getLiquidityPool\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isRunning\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"iPerpetualCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"fCeilPnLShare\",\"type\":\"int32\"},{\"internalType\":\"uint8\",\"name\":\"marginTokenDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"iTargetPoolSizeUpdateTime\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"marginTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"prevAnchor\",\"type\":\"uint64\"},{\"internalType\":\"int128\",\"name\":\"fRedemptionRate\",\"type\":\"int128\"},{\"internalType\":\"address\",\"name\":\"shareTokenAddress\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"fPnLparticipantsCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetAMMFundSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fDefaultFundCashCC\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fTargetDFSize\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"fBrokerCollateralLotSize\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"prevTokenAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"nextTokenAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalSupplyShareToken\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"fBrokerFundCashCC\",\"type\":\"int128\"}],\"internalType\":\"structLiquidityPoolData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// D8xPoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use D8xPoolManagerMetaData.ABI instead.
var D8xPoolManagerABI = D8xPoolManagerMetaData.ABI

// D8xPoolManager is an auto generated Go binding around an Ethereum contract.
type D8xPoolManager struct {
	D8xPoolManagerCaller     // Read-only binding to the contract
	D8xPoolManagerTransactor // Write-only binding to the contract
	D8xPoolManagerFilterer   // Log filterer for contract events
}

// D8xPoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type D8xPoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D8xPoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type D8xPoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D8xPoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type D8xPoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D8xPoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type D8xPoolManagerSession struct {
	Contract     *D8xPoolManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// D8xPoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type D8xPoolManagerCallerSession struct {
	Contract *D8xPoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// D8xPoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type D8xPoolManagerTransactorSession struct {
	Contract     *D8xPoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// D8xPoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type D8xPoolManagerRaw struct {
	Contract *D8xPoolManager // Generic contract binding to access the raw methods on
}

// D8xPoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type D8xPoolManagerCallerRaw struct {
	Contract *D8xPoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// D8xPoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type D8xPoolManagerTransactorRaw struct {
	Contract *D8xPoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewD8xPoolManager creates a new instance of D8xPoolManager, bound to a specific deployed contract.
func NewD8xPoolManager(address common.Address, backend bind.ContractBackend) (*D8xPoolManager, error) {
	contract, err := bindD8xPoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &D8xPoolManager{D8xPoolManagerCaller: D8xPoolManagerCaller{contract: contract}, D8xPoolManagerTransactor: D8xPoolManagerTransactor{contract: contract}, D8xPoolManagerFilterer: D8xPoolManagerFilterer{contract: contract}}, nil
}

// NewD8xPoolManagerCaller creates a new read-only instance of D8xPoolManager, bound to a specific deployed contract.
func NewD8xPoolManagerCaller(address common.Address, caller bind.ContractCaller) (*D8xPoolManagerCaller, error) {
	contract, err := bindD8xPoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &D8xPoolManagerCaller{contract: contract}, nil
}

// NewD8xPoolManagerTransactor creates a new write-only instance of D8xPoolManager, bound to a specific deployed contract.
func NewD8xPoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*D8xPoolManagerTransactor, error) {
	contract, err := bindD8xPoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &D8xPoolManagerTransactor{contract: contract}, nil
}

// NewD8xPoolManagerFilterer creates a new log filterer instance of D8xPoolManager, bound to a specific deployed contract.
func NewD8xPoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*D8xPoolManagerFilterer, error) {
	contract, err := bindD8xPoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &D8xPoolManagerFilterer{contract: contract}, nil
}

// bindD8xPoolManager binds a generic wrapper to an already deployed contract.
func bindD8xPoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := D8xPoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D8xPoolManager *D8xPoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D8xPoolManager.Contract.D8xPoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_D8xPoolManager *D8xPoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D8xPoolManager.Contract.D8xPoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_D8xPoolManager *D8xPoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D8xPoolManager.Contract.D8xPoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D8xPoolManager *D8xPoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D8xPoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_D8xPoolManager *D8xPoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D8xPoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_D8xPoolManager *D8xPoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D8xPoolManager.Contract.contract.Transact(opts, method, params...)
}

// GetLiquidityPool is a free data retrieval call binding the contract method 0x30af7c72.
//
// Solidity: function getLiquidityPool(uint8 _poolId) view returns((bool,uint8,uint8,int32,uint8,uint16,address,uint64,int128,address,int128,int128,int128,int128,int128,uint128,uint128,uint128,int128))
func (_D8xPoolManager *D8xPoolManagerCaller) GetLiquidityPool(opts *bind.CallOpts, _poolId uint8) (LiquidityPoolData, error) {
	var out []interface{}
	err := _D8xPoolManager.contract.Call(opts, &out, "getLiquidityPool", _poolId)

	if err != nil {
		return *new(LiquidityPoolData), err
	}

	out0 := *abi.ConvertType(out[0], new(LiquidityPoolData)).(*LiquidityPoolData)

	return out0, err

}

// GetLiquidityPool is a free data retrieval call binding the contract method 0x30af7c72.
//
// Solidity: function getLiquidityPool(uint8 _poolId) view returns((bool,uint8,uint8,int32,uint8,uint16,address,uint64,int128,address,int128,int128,int128,int128,int128,uint128,uint128,uint128,int128))
func (_D8xPoolManager *D8xPoolManagerSession) GetLiquidityPool(_poolId uint8) (LiquidityPoolData, error) {
	return _D8xPoolManager.Contract.GetLiquidityPool(&_D8xPoolManager.CallOpts, _poolId)
}

// GetLiquidityPool is a free data retrieval call binding the contract method 0x30af7c72.
//
// Solidity: function getLiquidityPool(uint8 _poolId) view returns((bool,uint8,uint8,int32,uint8,uint16,address,uint64,int128,address,int128,int128,int128,int128,int128,uint128,uint128,uint128,int128))
func (_D8xPoolManager *D8xPoolManagerCallerSession) GetLiquidityPool(_poolId uint8) (LiquidityPoolData, error) {
	return _D8xPoolManager.Contract.GetLiquidityPool(&_D8xPoolManager.CallOpts, _poolId)
}

// GetShareTokenPriceD18 is a free data retrieval call binding the contract method 0x2faee618.
//
// Solidity: function getShareTokenPriceD18(uint8 _poolId) view returns(uint256 price)
func (_D8xPoolManager *D8xPoolManagerCaller) GetShareTokenPriceD18(opts *bind.CallOpts, _poolId uint8) (*big.Int, error) {
	var out []interface{}
	err := _D8xPoolManager.contract.Call(opts, &out, "getShareTokenPriceD18", _poolId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetShareTokenPriceD18 is a free data retrieval call binding the contract method 0x2faee618.
//
// Solidity: function getShareTokenPriceD18(uint8 _poolId) view returns(uint256 price)
func (_D8xPoolManager *D8xPoolManagerSession) GetShareTokenPriceD18(_poolId uint8) (*big.Int, error) {
	return _D8xPoolManager.Contract.GetShareTokenPriceD18(&_D8xPoolManager.CallOpts, _poolId)
}

// GetShareTokenPriceD18 is a free data retrieval call binding the contract method 0x2faee618.
//
// Solidity: function getShareTokenPriceD18(uint8 _poolId) view returns(uint256 price)
func (_D8xPoolManager *D8xPoolManagerCallerSession) GetShareTokenPriceD18(_poolId uint8) (*big.Int, error) {
	return _D8xPoolManager.Contract.GetShareTokenPriceD18(&_D8xPoolManager.CallOpts, _poolId)
}
