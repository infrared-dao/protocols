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

// IchiMetaData contains all meta data concerning the Ichi contract.
var IchiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowToken0\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_allowToken1\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"__owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_twapPeriod\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"Affiliate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ammFeeRecipient\",\"type\":\"address\"}],\"name\":\"AmmFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowToken0\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowToken1\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"twapPeriod\",\"type\":\"uint256\"}],\"name\":\"DeployICHIVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit0Max\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit1Max\",\"type\":\"uint256\"}],\"name\":\"DepositMax\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hysteresis\",\"type\":\"uint256\"}],\"name\":\"Hysteresis\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newAuxTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"SetAuxTwapPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"SetTwapPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"affiliate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowToken0\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowToken1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ammFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auxTwapPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fees1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit0Max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit1Max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasePosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimitPosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hysteresis\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ichiVaultFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_baseLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_baseUpper\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_limitLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_limitUpper\",\"type\":\"int24\"},{\"internalType\":\"int256\",\"name\":\"swapQuantity\",\"type\":\"int256\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_affiliate\",\"type\":\"address\"}],\"name\":\"setAffiliate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ammFeeRecipient\",\"type\":\"address\"}],\"name\":\"setAmmFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newAuxTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"setAuxTwapPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deposit0Max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deposit1Max\",\"type\":\"uint256\"}],\"name\":\"setDepositMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hysteresis\",\"type\":\"uint256\"}],\"name\":\"setHysteresis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"setTwapPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twapPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IchiABI is the input ABI used to generate the binding from.
// Deprecated: Use IchiMetaData.ABI instead.
var IchiABI = IchiMetaData.ABI

// Ichi is an auto generated Go binding around an Ethereum contract.
type Ichi struct {
	IchiCaller     // Read-only binding to the contract
	IchiTransactor // Write-only binding to the contract
	IchiFilterer   // Log filterer for contract events
}

// IchiCaller is an auto generated read-only Go binding around an Ethereum contract.
type IchiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IchiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IchiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IchiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IchiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IchiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IchiSession struct {
	Contract     *Ichi             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IchiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IchiCallerSession struct {
	Contract *IchiCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IchiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IchiTransactorSession struct {
	Contract     *IchiTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IchiRaw is an auto generated low-level Go binding around an Ethereum contract.
type IchiRaw struct {
	Contract *Ichi // Generic contract binding to access the raw methods on
}

// IchiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IchiCallerRaw struct {
	Contract *IchiCaller // Generic read-only contract binding to access the raw methods on
}

// IchiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IchiTransactorRaw struct {
	Contract *IchiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIchi creates a new instance of Ichi, bound to a specific deployed contract.
func NewIchi(address common.Address, backend bind.ContractBackend) (*Ichi, error) {
	contract, err := bindIchi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ichi{IchiCaller: IchiCaller{contract: contract}, IchiTransactor: IchiTransactor{contract: contract}, IchiFilterer: IchiFilterer{contract: contract}}, nil
}

// NewIchiCaller creates a new read-only instance of Ichi, bound to a specific deployed contract.
func NewIchiCaller(address common.Address, caller bind.ContractCaller) (*IchiCaller, error) {
	contract, err := bindIchi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IchiCaller{contract: contract}, nil
}

// NewIchiTransactor creates a new write-only instance of Ichi, bound to a specific deployed contract.
func NewIchiTransactor(address common.Address, transactor bind.ContractTransactor) (*IchiTransactor, error) {
	contract, err := bindIchi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IchiTransactor{contract: contract}, nil
}

// NewIchiFilterer creates a new log filterer instance of Ichi, bound to a specific deployed contract.
func NewIchiFilterer(address common.Address, filterer bind.ContractFilterer) (*IchiFilterer, error) {
	contract, err := bindIchi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IchiFilterer{contract: contract}, nil
}

// bindIchi binds a generic wrapper to an already deployed contract.
func bindIchi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IchiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ichi *IchiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ichi.Contract.IchiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ichi *IchiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ichi.Contract.IchiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ichi *IchiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ichi.Contract.IchiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ichi *IchiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ichi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ichi *IchiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ichi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ichi *IchiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ichi.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Ichi *IchiCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Ichi *IchiSession) PRECISION() (*big.Int, error) {
	return _Ichi.Contract.PRECISION(&_Ichi.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Ichi *IchiCallerSession) PRECISION() (*big.Int, error) {
	return _Ichi.Contract.PRECISION(&_Ichi.CallOpts)
}

// Affiliate is a free data retrieval call binding the contract method 0x45e05f43.
//
// Solidity: function affiliate() view returns(address)
func (_Ichi *IchiCaller) Affiliate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "affiliate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Affiliate is a free data retrieval call binding the contract method 0x45e05f43.
//
// Solidity: function affiliate() view returns(address)
func (_Ichi *IchiSession) Affiliate() (common.Address, error) {
	return _Ichi.Contract.Affiliate(&_Ichi.CallOpts)
}

// Affiliate is a free data retrieval call binding the contract method 0x45e05f43.
//
// Solidity: function affiliate() view returns(address)
func (_Ichi *IchiCallerSession) Affiliate() (common.Address, error) {
	return _Ichi.Contract.Affiliate(&_Ichi.CallOpts)
}

// AllowToken0 is a free data retrieval call binding the contract method 0x7f7a1eec.
//
// Solidity: function allowToken0() view returns(bool)
func (_Ichi *IchiCaller) AllowToken0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "allowToken0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowToken0 is a free data retrieval call binding the contract method 0x7f7a1eec.
//
// Solidity: function allowToken0() view returns(bool)
func (_Ichi *IchiSession) AllowToken0() (bool, error) {
	return _Ichi.Contract.AllowToken0(&_Ichi.CallOpts)
}

// AllowToken0 is a free data retrieval call binding the contract method 0x7f7a1eec.
//
// Solidity: function allowToken0() view returns(bool)
func (_Ichi *IchiCallerSession) AllowToken0() (bool, error) {
	return _Ichi.Contract.AllowToken0(&_Ichi.CallOpts)
}

// AllowToken1 is a free data retrieval call binding the contract method 0x37e41b40.
//
// Solidity: function allowToken1() view returns(bool)
func (_Ichi *IchiCaller) AllowToken1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "allowToken1")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowToken1 is a free data retrieval call binding the contract method 0x37e41b40.
//
// Solidity: function allowToken1() view returns(bool)
func (_Ichi *IchiSession) AllowToken1() (bool, error) {
	return _Ichi.Contract.AllowToken1(&_Ichi.CallOpts)
}

// AllowToken1 is a free data retrieval call binding the contract method 0x37e41b40.
//
// Solidity: function allowToken1() view returns(bool)
func (_Ichi *IchiCallerSession) AllowToken1() (bool, error) {
	return _Ichi.Contract.AllowToken1(&_Ichi.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ichi *IchiCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ichi *IchiSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ichi.Contract.Allowance(&_Ichi.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ichi *IchiCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ichi.Contract.Allowance(&_Ichi.CallOpts, owner, spender)
}

// AmmFeeRecipient is a free data retrieval call binding the contract method 0x897f078c.
//
// Solidity: function ammFeeRecipient() view returns(address)
func (_Ichi *IchiCaller) AmmFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "ammFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmmFeeRecipient is a free data retrieval call binding the contract method 0x897f078c.
//
// Solidity: function ammFeeRecipient() view returns(address)
func (_Ichi *IchiSession) AmmFeeRecipient() (common.Address, error) {
	return _Ichi.Contract.AmmFeeRecipient(&_Ichi.CallOpts)
}

// AmmFeeRecipient is a free data retrieval call binding the contract method 0x897f078c.
//
// Solidity: function ammFeeRecipient() view returns(address)
func (_Ichi *IchiCallerSession) AmmFeeRecipient() (common.Address, error) {
	return _Ichi.Contract.AmmFeeRecipient(&_Ichi.CallOpts)
}

// AuxTwapPeriod is a free data retrieval call binding the contract method 0x91563d32.
//
// Solidity: function auxTwapPeriod() view returns(uint32)
func (_Ichi *IchiCaller) AuxTwapPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "auxTwapPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AuxTwapPeriod is a free data retrieval call binding the contract method 0x91563d32.
//
// Solidity: function auxTwapPeriod() view returns(uint32)
func (_Ichi *IchiSession) AuxTwapPeriod() (uint32, error) {
	return _Ichi.Contract.AuxTwapPeriod(&_Ichi.CallOpts)
}

// AuxTwapPeriod is a free data retrieval call binding the contract method 0x91563d32.
//
// Solidity: function auxTwapPeriod() view returns(uint32)
func (_Ichi *IchiCallerSession) AuxTwapPeriod() (uint32, error) {
	return _Ichi.Contract.AuxTwapPeriod(&_Ichi.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ichi *IchiCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ichi *IchiSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ichi.Contract.BalanceOf(&_Ichi.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ichi *IchiCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ichi.Contract.BalanceOf(&_Ichi.CallOpts, account)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_Ichi *IchiCaller) BaseLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "baseLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_Ichi *IchiSession) BaseLower() (*big.Int, error) {
	return _Ichi.Contract.BaseLower(&_Ichi.CallOpts)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_Ichi *IchiCallerSession) BaseLower() (*big.Int, error) {
	return _Ichi.Contract.BaseLower(&_Ichi.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_Ichi *IchiCaller) BaseUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "baseUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_Ichi *IchiSession) BaseUpper() (*big.Int, error) {
	return _Ichi.Contract.BaseUpper(&_Ichi.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_Ichi *IchiCallerSession) BaseUpper() (*big.Int, error) {
	return _Ichi.Contract.BaseUpper(&_Ichi.CallOpts)
}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_Ichi *IchiCaller) CurrentTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "currentTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_Ichi *IchiSession) CurrentTick() (*big.Int, error) {
	return _Ichi.Contract.CurrentTick(&_Ichi.CallOpts)
}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_Ichi *IchiCallerSession) CurrentTick() (*big.Int, error) {
	return _Ichi.Contract.CurrentTick(&_Ichi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ichi *IchiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ichi *IchiSession) Decimals() (uint8, error) {
	return _Ichi.Contract.Decimals(&_Ichi.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ichi *IchiCallerSession) Decimals() (uint8, error) {
	return _Ichi.Contract.Decimals(&_Ichi.CallOpts)
}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_Ichi *IchiCaller) Deposit0Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "deposit0Max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_Ichi *IchiSession) Deposit0Max() (*big.Int, error) {
	return _Ichi.Contract.Deposit0Max(&_Ichi.CallOpts)
}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_Ichi *IchiCallerSession) Deposit0Max() (*big.Int, error) {
	return _Ichi.Contract.Deposit0Max(&_Ichi.CallOpts)
}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_Ichi *IchiCaller) Deposit1Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "deposit1Max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_Ichi *IchiSession) Deposit1Max() (*big.Int, error) {
	return _Ichi.Contract.Deposit1Max(&_Ichi.CallOpts)
}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_Ichi *IchiCallerSession) Deposit1Max() (*big.Int, error) {
	return _Ichi.Contract.Deposit1Max(&_Ichi.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Ichi *IchiCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Ichi *IchiSession) Fee() (*big.Int, error) {
	return _Ichi.Contract.Fee(&_Ichi.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Ichi *IchiCallerSession) Fee() (*big.Int, error) {
	return _Ichi.Contract.Fee(&_Ichi.CallOpts)
}

// GetBasePosition is a free data retrieval call binding the contract method 0xd2eabcfc.
//
// Solidity: function getBasePosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Ichi *IchiCaller) GetBasePosition(opts *bind.CallOpts) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "getBasePosition")

	outstruct := new(struct {
		Liquidity *big.Int
		Amount0   *big.Int
		Amount1   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount0 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBasePosition is a free data retrieval call binding the contract method 0xd2eabcfc.
//
// Solidity: function getBasePosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Ichi *IchiSession) GetBasePosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _Ichi.Contract.GetBasePosition(&_Ichi.CallOpts)
}

// GetBasePosition is a free data retrieval call binding the contract method 0xd2eabcfc.
//
// Solidity: function getBasePosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Ichi *IchiCallerSession) GetBasePosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _Ichi.Contract.GetBasePosition(&_Ichi.CallOpts)
}

// GetLimitPosition is a free data retrieval call binding the contract method 0xa049de6b.
//
// Solidity: function getLimitPosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Ichi *IchiCaller) GetLimitPosition(opts *bind.CallOpts) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "getLimitPosition")

	outstruct := new(struct {
		Liquidity *big.Int
		Amount0   *big.Int
		Amount1   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount0 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLimitPosition is a free data retrieval call binding the contract method 0xa049de6b.
//
// Solidity: function getLimitPosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Ichi *IchiSession) GetLimitPosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _Ichi.Contract.GetLimitPosition(&_Ichi.CallOpts)
}

// GetLimitPosition is a free data retrieval call binding the contract method 0xa049de6b.
//
// Solidity: function getLimitPosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Ichi *IchiCallerSession) GetLimitPosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _Ichi.Contract.GetLimitPosition(&_Ichi.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_Ichi *IchiCaller) GetTotalAmounts(opts *bind.CallOpts) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "getTotalAmounts")

	outstruct := new(struct {
		Total0 *big.Int
		Total1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Total0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Total1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_Ichi *IchiSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _Ichi.Contract.GetTotalAmounts(&_Ichi.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_Ichi *IchiCallerSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _Ichi.Contract.GetTotalAmounts(&_Ichi.CallOpts)
}

// Hysteresis is a free data retrieval call binding the contract method 0x7aea5309.
//
// Solidity: function hysteresis() view returns(uint256)
func (_Ichi *IchiCaller) Hysteresis(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "hysteresis")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hysteresis is a free data retrieval call binding the contract method 0x7aea5309.
//
// Solidity: function hysteresis() view returns(uint256)
func (_Ichi *IchiSession) Hysteresis() (*big.Int, error) {
	return _Ichi.Contract.Hysteresis(&_Ichi.CallOpts)
}

// Hysteresis is a free data retrieval call binding the contract method 0x7aea5309.
//
// Solidity: function hysteresis() view returns(uint256)
func (_Ichi *IchiCallerSession) Hysteresis() (*big.Int, error) {
	return _Ichi.Contract.Hysteresis(&_Ichi.CallOpts)
}

// IchiVaultFactory is a free data retrieval call binding the contract method 0xdd81fa63.
//
// Solidity: function ichiVaultFactory() view returns(address)
func (_Ichi *IchiCaller) IchiVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "ichiVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IchiVaultFactory is a free data retrieval call binding the contract method 0xdd81fa63.
//
// Solidity: function ichiVaultFactory() view returns(address)
func (_Ichi *IchiSession) IchiVaultFactory() (common.Address, error) {
	return _Ichi.Contract.IchiVaultFactory(&_Ichi.CallOpts)
}

// IchiVaultFactory is a free data retrieval call binding the contract method 0xdd81fa63.
//
// Solidity: function ichiVaultFactory() view returns(address)
func (_Ichi *IchiCallerSession) IchiVaultFactory() (common.Address, error) {
	return _Ichi.Contract.IchiVaultFactory(&_Ichi.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_Ichi *IchiCaller) LimitLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "limitLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_Ichi *IchiSession) LimitLower() (*big.Int, error) {
	return _Ichi.Contract.LimitLower(&_Ichi.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_Ichi *IchiCallerSession) LimitLower() (*big.Int, error) {
	return _Ichi.Contract.LimitLower(&_Ichi.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_Ichi *IchiCaller) LimitUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "limitUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_Ichi *IchiSession) LimitUpper() (*big.Int, error) {
	return _Ichi.Contract.LimitUpper(&_Ichi.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_Ichi *IchiCallerSession) LimitUpper() (*big.Int, error) {
	return _Ichi.Contract.LimitUpper(&_Ichi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ichi *IchiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ichi *IchiSession) Name() (string, error) {
	return _Ichi.Contract.Name(&_Ichi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ichi *IchiCallerSession) Name() (string, error) {
	return _Ichi.Contract.Name(&_Ichi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ichi *IchiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ichi *IchiSession) Owner() (common.Address, error) {
	return _Ichi.Contract.Owner(&_Ichi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ichi *IchiCallerSession) Owner() (common.Address, error) {
	return _Ichi.Contract.Owner(&_Ichi.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Ichi *IchiCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Ichi *IchiSession) Pool() (common.Address, error) {
	return _Ichi.Contract.Pool(&_Ichi.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Ichi *IchiCallerSession) Pool() (common.Address, error) {
	return _Ichi.Contract.Pool(&_Ichi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ichi *IchiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ichi *IchiSession) Symbol() (string, error) {
	return _Ichi.Contract.Symbol(&_Ichi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ichi *IchiCallerSession) Symbol() (string, error) {
	return _Ichi.Contract.Symbol(&_Ichi.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Ichi *IchiCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Ichi *IchiSession) TickSpacing() (*big.Int, error) {
	return _Ichi.Contract.TickSpacing(&_Ichi.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Ichi *IchiCallerSession) TickSpacing() (*big.Int, error) {
	return _Ichi.Contract.TickSpacing(&_Ichi.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Ichi *IchiCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Ichi *IchiSession) Token0() (common.Address, error) {
	return _Ichi.Contract.Token0(&_Ichi.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Ichi *IchiCallerSession) Token0() (common.Address, error) {
	return _Ichi.Contract.Token0(&_Ichi.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Ichi *IchiCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Ichi *IchiSession) Token1() (common.Address, error) {
	return _Ichi.Contract.Token1(&_Ichi.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Ichi *IchiCallerSession) Token1() (common.Address, error) {
	return _Ichi.Contract.Token1(&_Ichi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ichi *IchiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ichi *IchiSession) TotalSupply() (*big.Int, error) {
	return _Ichi.Contract.TotalSupply(&_Ichi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ichi *IchiCallerSession) TotalSupply() (*big.Int, error) {
	return _Ichi.Contract.TotalSupply(&_Ichi.CallOpts)
}

// TwapPeriod is a free data retrieval call binding the contract method 0xf6207326.
//
// Solidity: function twapPeriod() view returns(uint32)
func (_Ichi *IchiCaller) TwapPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Ichi.contract.Call(opts, &out, "twapPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TwapPeriod is a free data retrieval call binding the contract method 0xf6207326.
//
// Solidity: function twapPeriod() view returns(uint32)
func (_Ichi *IchiSession) TwapPeriod() (uint32, error) {
	return _Ichi.Contract.TwapPeriod(&_Ichi.CallOpts)
}

// TwapPeriod is a free data retrieval call binding the contract method 0xf6207326.
//
// Solidity: function twapPeriod() view returns(uint32)
func (_Ichi *IchiCallerSession) TwapPeriod() (uint32, error) {
	return _Ichi.Contract.TwapPeriod(&_Ichi.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ichi *IchiTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ichi *IchiSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.Approve(&_Ichi.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ichi *IchiTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.Approve(&_Ichi.TransactOpts, spender, amount)
}

// CollectFees is a paid mutator transaction binding the contract method 0xc8796572.
//
// Solidity: function collectFees() returns(uint256 fees0, uint256 fees1)
func (_Ichi *IchiTransactor) CollectFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "collectFees")
}

// CollectFees is a paid mutator transaction binding the contract method 0xc8796572.
//
// Solidity: function collectFees() returns(uint256 fees0, uint256 fees1)
func (_Ichi *IchiSession) CollectFees() (*types.Transaction, error) {
	return _Ichi.Contract.CollectFees(&_Ichi.TransactOpts)
}

// CollectFees is a paid mutator transaction binding the contract method 0xc8796572.
//
// Solidity: function collectFees() returns(uint256 fees0, uint256 fees1)
func (_Ichi *IchiTransactorSession) CollectFees() (*types.Transaction, error) {
	return _Ichi.Contract.CollectFees(&_Ichi.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Ichi *IchiTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Ichi *IchiSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.DecreaseAllowance(&_Ichi.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Ichi *IchiTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.DecreaseAllowance(&_Ichi.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to) returns(uint256 shares)
func (_Ichi *IchiTransactor) Deposit(opts *bind.TransactOpts, deposit0 *big.Int, deposit1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "deposit", deposit0, deposit1, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to) returns(uint256 shares)
func (_Ichi *IchiSession) Deposit(deposit0 *big.Int, deposit1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _Ichi.Contract.Deposit(&_Ichi.TransactOpts, deposit0, deposit1, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to) returns(uint256 shares)
func (_Ichi *IchiTransactorSession) Deposit(deposit0 *big.Int, deposit1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _Ichi.Contract.Deposit(&_Ichi.TransactOpts, deposit0, deposit1, to)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Ichi *IchiTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Ichi *IchiSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.IncreaseAllowance(&_Ichi.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Ichi *IchiTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.IncreaseAllowance(&_Ichi.TransactOpts, spender, addedValue)
}

// Rebalance is a paid mutator transaction binding the contract method 0xd87346aa.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, int256 swapQuantity) returns()
func (_Ichi *IchiTransactor) Rebalance(opts *bind.TransactOpts, _baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, swapQuantity *big.Int) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "rebalance", _baseLower, _baseUpper, _limitLower, _limitUpper, swapQuantity)
}

// Rebalance is a paid mutator transaction binding the contract method 0xd87346aa.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, int256 swapQuantity) returns()
func (_Ichi *IchiSession) Rebalance(_baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, swapQuantity *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.Rebalance(&_Ichi.TransactOpts, _baseLower, _baseUpper, _limitLower, _limitUpper, swapQuantity)
}

// Rebalance is a paid mutator transaction binding the contract method 0xd87346aa.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, int256 swapQuantity) returns()
func (_Ichi *IchiTransactorSession) Rebalance(_baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, swapQuantity *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.Rebalance(&_Ichi.TransactOpts, _baseLower, _baseUpper, _limitLower, _limitUpper, swapQuantity)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ichi *IchiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ichi *IchiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ichi.Contract.RenounceOwnership(&_Ichi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ichi *IchiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ichi.Contract.RenounceOwnership(&_Ichi.TransactOpts)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x2bbb56d9.
//
// Solidity: function setAffiliate(address _affiliate) returns()
func (_Ichi *IchiTransactor) SetAffiliate(opts *bind.TransactOpts, _affiliate common.Address) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "setAffiliate", _affiliate)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x2bbb56d9.
//
// Solidity: function setAffiliate(address _affiliate) returns()
func (_Ichi *IchiSession) SetAffiliate(_affiliate common.Address) (*types.Transaction, error) {
	return _Ichi.Contract.SetAffiliate(&_Ichi.TransactOpts, _affiliate)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x2bbb56d9.
//
// Solidity: function setAffiliate(address _affiliate) returns()
func (_Ichi *IchiTransactorSession) SetAffiliate(_affiliate common.Address) (*types.Transaction, error) {
	return _Ichi.Contract.SetAffiliate(&_Ichi.TransactOpts, _affiliate)
}

// SetAmmFeeRecipient is a paid mutator transaction binding the contract method 0x81de128b.
//
// Solidity: function setAmmFeeRecipient(address _ammFeeRecipient) returns()
func (_Ichi *IchiTransactor) SetAmmFeeRecipient(opts *bind.TransactOpts, _ammFeeRecipient common.Address) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "setAmmFeeRecipient", _ammFeeRecipient)
}

// SetAmmFeeRecipient is a paid mutator transaction binding the contract method 0x81de128b.
//
// Solidity: function setAmmFeeRecipient(address _ammFeeRecipient) returns()
func (_Ichi *IchiSession) SetAmmFeeRecipient(_ammFeeRecipient common.Address) (*types.Transaction, error) {
	return _Ichi.Contract.SetAmmFeeRecipient(&_Ichi.TransactOpts, _ammFeeRecipient)
}

// SetAmmFeeRecipient is a paid mutator transaction binding the contract method 0x81de128b.
//
// Solidity: function setAmmFeeRecipient(address _ammFeeRecipient) returns()
func (_Ichi *IchiTransactorSession) SetAmmFeeRecipient(_ammFeeRecipient common.Address) (*types.Transaction, error) {
	return _Ichi.Contract.SetAmmFeeRecipient(&_Ichi.TransactOpts, _ammFeeRecipient)
}

// SetAuxTwapPeriod is a paid mutator transaction binding the contract method 0x400f0ceb.
//
// Solidity: function setAuxTwapPeriod(uint32 newAuxTwapPeriod) returns()
func (_Ichi *IchiTransactor) SetAuxTwapPeriod(opts *bind.TransactOpts, newAuxTwapPeriod uint32) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "setAuxTwapPeriod", newAuxTwapPeriod)
}

// SetAuxTwapPeriod is a paid mutator transaction binding the contract method 0x400f0ceb.
//
// Solidity: function setAuxTwapPeriod(uint32 newAuxTwapPeriod) returns()
func (_Ichi *IchiSession) SetAuxTwapPeriod(newAuxTwapPeriod uint32) (*types.Transaction, error) {
	return _Ichi.Contract.SetAuxTwapPeriod(&_Ichi.TransactOpts, newAuxTwapPeriod)
}

// SetAuxTwapPeriod is a paid mutator transaction binding the contract method 0x400f0ceb.
//
// Solidity: function setAuxTwapPeriod(uint32 newAuxTwapPeriod) returns()
func (_Ichi *IchiTransactorSession) SetAuxTwapPeriod(newAuxTwapPeriod uint32) (*types.Transaction, error) {
	return _Ichi.Contract.SetAuxTwapPeriod(&_Ichi.TransactOpts, newAuxTwapPeriod)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0x3e091ee9.
//
// Solidity: function setDepositMax(uint256 _deposit0Max, uint256 _deposit1Max) returns()
func (_Ichi *IchiTransactor) SetDepositMax(opts *bind.TransactOpts, _deposit0Max *big.Int, _deposit1Max *big.Int) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "setDepositMax", _deposit0Max, _deposit1Max)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0x3e091ee9.
//
// Solidity: function setDepositMax(uint256 _deposit0Max, uint256 _deposit1Max) returns()
func (_Ichi *IchiSession) SetDepositMax(_deposit0Max *big.Int, _deposit1Max *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.SetDepositMax(&_Ichi.TransactOpts, _deposit0Max, _deposit1Max)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0x3e091ee9.
//
// Solidity: function setDepositMax(uint256 _deposit0Max, uint256 _deposit1Max) returns()
func (_Ichi *IchiTransactorSession) SetDepositMax(_deposit0Max *big.Int, _deposit1Max *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.SetDepositMax(&_Ichi.TransactOpts, _deposit0Max, _deposit1Max)
}

// SetHysteresis is a paid mutator transaction binding the contract method 0x5ffc1ff7.
//
// Solidity: function setHysteresis(uint256 _hysteresis) returns()
func (_Ichi *IchiTransactor) SetHysteresis(opts *bind.TransactOpts, _hysteresis *big.Int) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "setHysteresis", _hysteresis)
}

// SetHysteresis is a paid mutator transaction binding the contract method 0x5ffc1ff7.
//
// Solidity: function setHysteresis(uint256 _hysteresis) returns()
func (_Ichi *IchiSession) SetHysteresis(_hysteresis *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.SetHysteresis(&_Ichi.TransactOpts, _hysteresis)
}

// SetHysteresis is a paid mutator transaction binding the contract method 0x5ffc1ff7.
//
// Solidity: function setHysteresis(uint256 _hysteresis) returns()
func (_Ichi *IchiTransactorSession) SetHysteresis(_hysteresis *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.SetHysteresis(&_Ichi.TransactOpts, _hysteresis)
}

// SetTwapPeriod is a paid mutator transaction binding the contract method 0xf9c95d46.
//
// Solidity: function setTwapPeriod(uint32 newTwapPeriod) returns()
func (_Ichi *IchiTransactor) SetTwapPeriod(opts *bind.TransactOpts, newTwapPeriod uint32) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "setTwapPeriod", newTwapPeriod)
}

// SetTwapPeriod is a paid mutator transaction binding the contract method 0xf9c95d46.
//
// Solidity: function setTwapPeriod(uint32 newTwapPeriod) returns()
func (_Ichi *IchiSession) SetTwapPeriod(newTwapPeriod uint32) (*types.Transaction, error) {
	return _Ichi.Contract.SetTwapPeriod(&_Ichi.TransactOpts, newTwapPeriod)
}

// SetTwapPeriod is a paid mutator transaction binding the contract method 0xf9c95d46.
//
// Solidity: function setTwapPeriod(uint32 newTwapPeriod) returns()
func (_Ichi *IchiTransactorSession) SetTwapPeriod(newTwapPeriod uint32) (*types.Transaction, error) {
	return _Ichi.Contract.SetTwapPeriod(&_Ichi.TransactOpts, newTwapPeriod)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Ichi *IchiTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Ichi *IchiSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.Transfer(&_Ichi.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Ichi *IchiTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.Transfer(&_Ichi.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Ichi *IchiTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Ichi *IchiSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.TransferFrom(&_Ichi.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Ichi *IchiTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ichi.Contract.TransferFrom(&_Ichi.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ichi *IchiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ichi *IchiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ichi.Contract.TransferOwnership(&_Ichi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ichi *IchiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ichi.Contract.TransferOwnership(&_Ichi.TransactOpts, newOwner)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Ichi *IchiTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "uniswapV3MintCallback", amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Ichi *IchiSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Ichi.Contract.UniswapV3MintCallback(&_Ichi.TransactOpts, amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Ichi *IchiTransactorSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Ichi.Contract.UniswapV3MintCallback(&_Ichi.TransactOpts, amount0, amount1, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Ichi *IchiTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Ichi *IchiSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Ichi.Contract.UniswapV3SwapCallback(&_Ichi.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Ichi *IchiTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Ichi.Contract.UniswapV3SwapCallback(&_Ichi.TransactOpts, amount0Delta, amount1Delta, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address to) returns(uint256 amount0, uint256 amount1)
func (_Ichi *IchiTransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int, to common.Address) (*types.Transaction, error) {
	return _Ichi.contract.Transact(opts, "withdraw", shares, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address to) returns(uint256 amount0, uint256 amount1)
func (_Ichi *IchiSession) Withdraw(shares *big.Int, to common.Address) (*types.Transaction, error) {
	return _Ichi.Contract.Withdraw(&_Ichi.TransactOpts, shares, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address to) returns(uint256 amount0, uint256 amount1)
func (_Ichi *IchiTransactorSession) Withdraw(shares *big.Int, to common.Address) (*types.Transaction, error) {
	return _Ichi.Contract.Withdraw(&_Ichi.TransactOpts, shares, to)
}

// IchiAffiliateIterator is returned from FilterAffiliate and is used to iterate over the raw logs and unpacked data for Affiliate events raised by the Ichi contract.
type IchiAffiliateIterator struct {
	Event *IchiAffiliate // Event containing the contract specifics and raw log

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
func (it *IchiAffiliateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiAffiliate)
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
		it.Event = new(IchiAffiliate)
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
func (it *IchiAffiliateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiAffiliateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiAffiliate represents a Affiliate event raised by the Ichi contract.
type IchiAffiliate struct {
	Sender    common.Address
	Affiliate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAffiliate is a free log retrieval operation binding the contract event 0x3066ef5dd340e8b2ea28d62f5a8391eb7a82d3ee87532724a1ca4386d34f7523.
//
// Solidity: event Affiliate(address indexed sender, address affiliate)
func (_Ichi *IchiFilterer) FilterAffiliate(opts *bind.FilterOpts, sender []common.Address) (*IchiAffiliateIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "Affiliate", senderRule)
	if err != nil {
		return nil, err
	}
	return &IchiAffiliateIterator{contract: _Ichi.contract, event: "Affiliate", logs: logs, sub: sub}, nil
}

// WatchAffiliate is a free log subscription operation binding the contract event 0x3066ef5dd340e8b2ea28d62f5a8391eb7a82d3ee87532724a1ca4386d34f7523.
//
// Solidity: event Affiliate(address indexed sender, address affiliate)
func (_Ichi *IchiFilterer) WatchAffiliate(opts *bind.WatchOpts, sink chan<- *IchiAffiliate, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "Affiliate", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiAffiliate)
				if err := _Ichi.contract.UnpackLog(event, "Affiliate", log); err != nil {
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

// ParseAffiliate is a log parse operation binding the contract event 0x3066ef5dd340e8b2ea28d62f5a8391eb7a82d3ee87532724a1ca4386d34f7523.
//
// Solidity: event Affiliate(address indexed sender, address affiliate)
func (_Ichi *IchiFilterer) ParseAffiliate(log types.Log) (*IchiAffiliate, error) {
	event := new(IchiAffiliate)
	if err := _Ichi.contract.UnpackLog(event, "Affiliate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiAmmFeeRecipientIterator is returned from FilterAmmFeeRecipient and is used to iterate over the raw logs and unpacked data for AmmFeeRecipient events raised by the Ichi contract.
type IchiAmmFeeRecipientIterator struct {
	Event *IchiAmmFeeRecipient // Event containing the contract specifics and raw log

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
func (it *IchiAmmFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiAmmFeeRecipient)
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
		it.Event = new(IchiAmmFeeRecipient)
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
func (it *IchiAmmFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiAmmFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiAmmFeeRecipient represents a AmmFeeRecipient event raised by the Ichi contract.
type IchiAmmFeeRecipient struct {
	Sender          common.Address
	AmmFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAmmFeeRecipient is a free log retrieval operation binding the contract event 0xbb78b7c13893a913fa8c9ecb9fdaf97597aa412a39c778bf976790555f0942f7.
//
// Solidity: event AmmFeeRecipient(address indexed sender, address ammFeeRecipient)
func (_Ichi *IchiFilterer) FilterAmmFeeRecipient(opts *bind.FilterOpts, sender []common.Address) (*IchiAmmFeeRecipientIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "AmmFeeRecipient", senderRule)
	if err != nil {
		return nil, err
	}
	return &IchiAmmFeeRecipientIterator{contract: _Ichi.contract, event: "AmmFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchAmmFeeRecipient is a free log subscription operation binding the contract event 0xbb78b7c13893a913fa8c9ecb9fdaf97597aa412a39c778bf976790555f0942f7.
//
// Solidity: event AmmFeeRecipient(address indexed sender, address ammFeeRecipient)
func (_Ichi *IchiFilterer) WatchAmmFeeRecipient(opts *bind.WatchOpts, sink chan<- *IchiAmmFeeRecipient, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "AmmFeeRecipient", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiAmmFeeRecipient)
				if err := _Ichi.contract.UnpackLog(event, "AmmFeeRecipient", log); err != nil {
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

// ParseAmmFeeRecipient is a log parse operation binding the contract event 0xbb78b7c13893a913fa8c9ecb9fdaf97597aa412a39c778bf976790555f0942f7.
//
// Solidity: event AmmFeeRecipient(address indexed sender, address ammFeeRecipient)
func (_Ichi *IchiFilterer) ParseAmmFeeRecipient(log types.Log) (*IchiAmmFeeRecipient, error) {
	event := new(IchiAmmFeeRecipient)
	if err := _Ichi.contract.UnpackLog(event, "AmmFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ichi contract.
type IchiApprovalIterator struct {
	Event *IchiApproval // Event containing the contract specifics and raw log

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
func (it *IchiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiApproval)
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
		it.Event = new(IchiApproval)
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
func (it *IchiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiApproval represents a Approval event raised by the Ichi contract.
type IchiApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ichi *IchiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IchiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IchiApprovalIterator{contract: _Ichi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ichi *IchiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IchiApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiApproval)
				if err := _Ichi.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Ichi *IchiFilterer) ParseApproval(log types.Log) (*IchiApproval, error) {
	event := new(IchiApproval)
	if err := _Ichi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiCollectFeesIterator is returned from FilterCollectFees and is used to iterate over the raw logs and unpacked data for CollectFees events raised by the Ichi contract.
type IchiCollectFeesIterator struct {
	Event *IchiCollectFees // Event containing the contract specifics and raw log

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
func (it *IchiCollectFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiCollectFees)
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
		it.Event = new(IchiCollectFees)
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
func (it *IchiCollectFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiCollectFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiCollectFees represents a CollectFees event raised by the Ichi contract.
type IchiCollectFees struct {
	Sender     common.Address
	FeeAmount0 *big.Int
	FeeAmount1 *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0xec8208dd791fa8ffdc0d7427f3ba9c0ed06f1bce9a86254e6940c10cc1802fef.
//
// Solidity: event CollectFees(address indexed sender, uint256 feeAmount0, uint256 feeAmount1)
func (_Ichi *IchiFilterer) FilterCollectFees(opts *bind.FilterOpts, sender []common.Address) (*IchiCollectFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "CollectFees", senderRule)
	if err != nil {
		return nil, err
	}
	return &IchiCollectFeesIterator{contract: _Ichi.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0xec8208dd791fa8ffdc0d7427f3ba9c0ed06f1bce9a86254e6940c10cc1802fef.
//
// Solidity: event CollectFees(address indexed sender, uint256 feeAmount0, uint256 feeAmount1)
func (_Ichi *IchiFilterer) WatchCollectFees(opts *bind.WatchOpts, sink chan<- *IchiCollectFees, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "CollectFees", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiCollectFees)
				if err := _Ichi.contract.UnpackLog(event, "CollectFees", log); err != nil {
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

// ParseCollectFees is a log parse operation binding the contract event 0xec8208dd791fa8ffdc0d7427f3ba9c0ed06f1bce9a86254e6940c10cc1802fef.
//
// Solidity: event CollectFees(address indexed sender, uint256 feeAmount0, uint256 feeAmount1)
func (_Ichi *IchiFilterer) ParseCollectFees(log types.Log) (*IchiCollectFees, error) {
	event := new(IchiCollectFees)
	if err := _Ichi.contract.UnpackLog(event, "CollectFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiDeployICHIVaultIterator is returned from FilterDeployICHIVault and is used to iterate over the raw logs and unpacked data for DeployICHIVault events raised by the Ichi contract.
type IchiDeployICHIVaultIterator struct {
	Event *IchiDeployICHIVault // Event containing the contract specifics and raw log

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
func (it *IchiDeployICHIVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiDeployICHIVault)
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
		it.Event = new(IchiDeployICHIVault)
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
func (it *IchiDeployICHIVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiDeployICHIVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiDeployICHIVault represents a DeployICHIVault event raised by the Ichi contract.
type IchiDeployICHIVault struct {
	Sender      common.Address
	Pool        common.Address
	AllowToken0 bool
	AllowToken1 bool
	Owner       common.Address
	TwapPeriod  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeployICHIVault is a free log retrieval operation binding the contract event 0x3e708ccf7d0e6de8558e020ea36189511cb3435bbfec54e721a48ee4df0d4f8c.
//
// Solidity: event DeployICHIVault(address indexed sender, address indexed pool, bool allowToken0, bool allowToken1, address owner, uint256 twapPeriod)
func (_Ichi *IchiFilterer) FilterDeployICHIVault(opts *bind.FilterOpts, sender []common.Address, pool []common.Address) (*IchiDeployICHIVaultIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "DeployICHIVault", senderRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &IchiDeployICHIVaultIterator{contract: _Ichi.contract, event: "DeployICHIVault", logs: logs, sub: sub}, nil
}

// WatchDeployICHIVault is a free log subscription operation binding the contract event 0x3e708ccf7d0e6de8558e020ea36189511cb3435bbfec54e721a48ee4df0d4f8c.
//
// Solidity: event DeployICHIVault(address indexed sender, address indexed pool, bool allowToken0, bool allowToken1, address owner, uint256 twapPeriod)
func (_Ichi *IchiFilterer) WatchDeployICHIVault(opts *bind.WatchOpts, sink chan<- *IchiDeployICHIVault, sender []common.Address, pool []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "DeployICHIVault", senderRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiDeployICHIVault)
				if err := _Ichi.contract.UnpackLog(event, "DeployICHIVault", log); err != nil {
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

// ParseDeployICHIVault is a log parse operation binding the contract event 0x3e708ccf7d0e6de8558e020ea36189511cb3435bbfec54e721a48ee4df0d4f8c.
//
// Solidity: event DeployICHIVault(address indexed sender, address indexed pool, bool allowToken0, bool allowToken1, address owner, uint256 twapPeriod)
func (_Ichi *IchiFilterer) ParseDeployICHIVault(log types.Log) (*IchiDeployICHIVault, error) {
	event := new(IchiDeployICHIVault)
	if err := _Ichi.contract.UnpackLog(event, "DeployICHIVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Ichi contract.
type IchiDepositIterator struct {
	Event *IchiDeposit // Event containing the contract specifics and raw log

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
func (it *IchiDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiDeposit)
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
		it.Event = new(IchiDeposit)
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
func (it *IchiDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiDeposit represents a Deposit event raised by the Ichi contract.
type IchiDeposit struct {
	Sender  common.Address
	To      common.Address
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Ichi *IchiFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IchiDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IchiDepositIterator{contract: _Ichi.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Ichi *IchiFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IchiDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiDeposit)
				if err := _Ichi.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Ichi *IchiFilterer) ParseDeposit(log types.Log) (*IchiDeposit, error) {
	event := new(IchiDeposit)
	if err := _Ichi.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiDepositMaxIterator is returned from FilterDepositMax and is used to iterate over the raw logs and unpacked data for DepositMax events raised by the Ichi contract.
type IchiDepositMaxIterator struct {
	Event *IchiDepositMax // Event containing the contract specifics and raw log

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
func (it *IchiDepositMaxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiDepositMax)
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
		it.Event = new(IchiDepositMax)
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
func (it *IchiDepositMaxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiDepositMaxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiDepositMax represents a DepositMax event raised by the Ichi contract.
type IchiDepositMax struct {
	Sender      common.Address
	Deposit0Max *big.Int
	Deposit1Max *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositMax is a free log retrieval operation binding the contract event 0xafd3b05a4086b378b6f291200a528d8aed8c5e0317af77436b001f1bec28821a.
//
// Solidity: event DepositMax(address indexed sender, uint256 deposit0Max, uint256 deposit1Max)
func (_Ichi *IchiFilterer) FilterDepositMax(opts *bind.FilterOpts, sender []common.Address) (*IchiDepositMaxIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "DepositMax", senderRule)
	if err != nil {
		return nil, err
	}
	return &IchiDepositMaxIterator{contract: _Ichi.contract, event: "DepositMax", logs: logs, sub: sub}, nil
}

// WatchDepositMax is a free log subscription operation binding the contract event 0xafd3b05a4086b378b6f291200a528d8aed8c5e0317af77436b001f1bec28821a.
//
// Solidity: event DepositMax(address indexed sender, uint256 deposit0Max, uint256 deposit1Max)
func (_Ichi *IchiFilterer) WatchDepositMax(opts *bind.WatchOpts, sink chan<- *IchiDepositMax, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "DepositMax", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiDepositMax)
				if err := _Ichi.contract.UnpackLog(event, "DepositMax", log); err != nil {
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

// ParseDepositMax is a log parse operation binding the contract event 0xafd3b05a4086b378b6f291200a528d8aed8c5e0317af77436b001f1bec28821a.
//
// Solidity: event DepositMax(address indexed sender, uint256 deposit0Max, uint256 deposit1Max)
func (_Ichi *IchiFilterer) ParseDepositMax(log types.Log) (*IchiDepositMax, error) {
	event := new(IchiDepositMax)
	if err := _Ichi.contract.UnpackLog(event, "DepositMax", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiHysteresisIterator is returned from FilterHysteresis and is used to iterate over the raw logs and unpacked data for Hysteresis events raised by the Ichi contract.
type IchiHysteresisIterator struct {
	Event *IchiHysteresis // Event containing the contract specifics and raw log

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
func (it *IchiHysteresisIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiHysteresis)
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
		it.Event = new(IchiHysteresis)
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
func (it *IchiHysteresisIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiHysteresisIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiHysteresis represents a Hysteresis event raised by the Ichi contract.
type IchiHysteresis struct {
	Sender     common.Address
	Hysteresis *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterHysteresis is a free log retrieval operation binding the contract event 0x529698f34660760dcb172def5c99d62e1b5b74b444df322e8f7da31f2bd0a86b.
//
// Solidity: event Hysteresis(address indexed sender, uint256 hysteresis)
func (_Ichi *IchiFilterer) FilterHysteresis(opts *bind.FilterOpts, sender []common.Address) (*IchiHysteresisIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "Hysteresis", senderRule)
	if err != nil {
		return nil, err
	}
	return &IchiHysteresisIterator{contract: _Ichi.contract, event: "Hysteresis", logs: logs, sub: sub}, nil
}

// WatchHysteresis is a free log subscription operation binding the contract event 0x529698f34660760dcb172def5c99d62e1b5b74b444df322e8f7da31f2bd0a86b.
//
// Solidity: event Hysteresis(address indexed sender, uint256 hysteresis)
func (_Ichi *IchiFilterer) WatchHysteresis(opts *bind.WatchOpts, sink chan<- *IchiHysteresis, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "Hysteresis", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiHysteresis)
				if err := _Ichi.contract.UnpackLog(event, "Hysteresis", log); err != nil {
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

// ParseHysteresis is a log parse operation binding the contract event 0x529698f34660760dcb172def5c99d62e1b5b74b444df322e8f7da31f2bd0a86b.
//
// Solidity: event Hysteresis(address indexed sender, uint256 hysteresis)
func (_Ichi *IchiFilterer) ParseHysteresis(log types.Log) (*IchiHysteresis, error) {
	event := new(IchiHysteresis)
	if err := _Ichi.contract.UnpackLog(event, "Hysteresis", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ichi contract.
type IchiOwnershipTransferredIterator struct {
	Event *IchiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IchiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiOwnershipTransferred)
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
		it.Event = new(IchiOwnershipTransferred)
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
func (it *IchiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiOwnershipTransferred represents a OwnershipTransferred event raised by the Ichi contract.
type IchiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ichi *IchiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IchiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IchiOwnershipTransferredIterator{contract: _Ichi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ichi *IchiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IchiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiOwnershipTransferred)
				if err := _Ichi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ichi *IchiFilterer) ParseOwnershipTransferred(log types.Log) (*IchiOwnershipTransferred, error) {
	event := new(IchiOwnershipTransferred)
	if err := _Ichi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the Ichi contract.
type IchiRebalanceIterator struct {
	Event *IchiRebalance // Event containing the contract specifics and raw log

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
func (it *IchiRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiRebalance)
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
		it.Event = new(IchiRebalance)
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
func (it *IchiRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiRebalance represents a Rebalance event raised by the Ichi contract.
type IchiRebalance struct {
	Tick         *big.Int
	TotalAmount0 *big.Int
	TotalAmount1 *big.Int
	FeeAmount0   *big.Int
	FeeAmount1   *big.Int
	TotalSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRebalance is a free log retrieval operation binding the contract event 0xbc4c20ad04f161d631d9ce94d27659391196415aa3c42f6a71c62e905ece782d.
//
// Solidity: event Rebalance(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 feeAmount0, uint256 feeAmount1, uint256 totalSupply)
func (_Ichi *IchiFilterer) FilterRebalance(opts *bind.FilterOpts) (*IchiRebalanceIterator, error) {

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return &IchiRebalanceIterator{contract: _Ichi.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0xbc4c20ad04f161d631d9ce94d27659391196415aa3c42f6a71c62e905ece782d.
//
// Solidity: event Rebalance(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 feeAmount0, uint256 feeAmount1, uint256 totalSupply)
func (_Ichi *IchiFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *IchiRebalance) (event.Subscription, error) {

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiRebalance)
				if err := _Ichi.contract.UnpackLog(event, "Rebalance", log); err != nil {
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

// ParseRebalance is a log parse operation binding the contract event 0xbc4c20ad04f161d631d9ce94d27659391196415aa3c42f6a71c62e905ece782d.
//
// Solidity: event Rebalance(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 feeAmount0, uint256 feeAmount1, uint256 totalSupply)
func (_Ichi *IchiFilterer) ParseRebalance(log types.Log) (*IchiRebalance, error) {
	event := new(IchiRebalance)
	if err := _Ichi.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiSetAuxTwapPeriodIterator is returned from FilterSetAuxTwapPeriod and is used to iterate over the raw logs and unpacked data for SetAuxTwapPeriod events raised by the Ichi contract.
type IchiSetAuxTwapPeriodIterator struct {
	Event *IchiSetAuxTwapPeriod // Event containing the contract specifics and raw log

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
func (it *IchiSetAuxTwapPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiSetAuxTwapPeriod)
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
		it.Event = new(IchiSetAuxTwapPeriod)
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
func (it *IchiSetAuxTwapPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiSetAuxTwapPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiSetAuxTwapPeriod represents a SetAuxTwapPeriod event raised by the Ichi contract.
type IchiSetAuxTwapPeriod struct {
	Sender           common.Address
	NewAuxTwapPeriod uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetAuxTwapPeriod is a free log retrieval operation binding the contract event 0x39da19f5960a3f182ced1ff1853b7be54f37150799b3003a40bf4e0d4c740c85.
//
// Solidity: event SetAuxTwapPeriod(address sender, uint32 newAuxTwapPeriod)
func (_Ichi *IchiFilterer) FilterSetAuxTwapPeriod(opts *bind.FilterOpts) (*IchiSetAuxTwapPeriodIterator, error) {

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "SetAuxTwapPeriod")
	if err != nil {
		return nil, err
	}
	return &IchiSetAuxTwapPeriodIterator{contract: _Ichi.contract, event: "SetAuxTwapPeriod", logs: logs, sub: sub}, nil
}

// WatchSetAuxTwapPeriod is a free log subscription operation binding the contract event 0x39da19f5960a3f182ced1ff1853b7be54f37150799b3003a40bf4e0d4c740c85.
//
// Solidity: event SetAuxTwapPeriod(address sender, uint32 newAuxTwapPeriod)
func (_Ichi *IchiFilterer) WatchSetAuxTwapPeriod(opts *bind.WatchOpts, sink chan<- *IchiSetAuxTwapPeriod) (event.Subscription, error) {

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "SetAuxTwapPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiSetAuxTwapPeriod)
				if err := _Ichi.contract.UnpackLog(event, "SetAuxTwapPeriod", log); err != nil {
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

// ParseSetAuxTwapPeriod is a log parse operation binding the contract event 0x39da19f5960a3f182ced1ff1853b7be54f37150799b3003a40bf4e0d4c740c85.
//
// Solidity: event SetAuxTwapPeriod(address sender, uint32 newAuxTwapPeriod)
func (_Ichi *IchiFilterer) ParseSetAuxTwapPeriod(log types.Log) (*IchiSetAuxTwapPeriod, error) {
	event := new(IchiSetAuxTwapPeriod)
	if err := _Ichi.contract.UnpackLog(event, "SetAuxTwapPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiSetTwapPeriodIterator is returned from FilterSetTwapPeriod and is used to iterate over the raw logs and unpacked data for SetTwapPeriod events raised by the Ichi contract.
type IchiSetTwapPeriodIterator struct {
	Event *IchiSetTwapPeriod // Event containing the contract specifics and raw log

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
func (it *IchiSetTwapPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiSetTwapPeriod)
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
		it.Event = new(IchiSetTwapPeriod)
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
func (it *IchiSetTwapPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiSetTwapPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiSetTwapPeriod represents a SetTwapPeriod event raised by the Ichi contract.
type IchiSetTwapPeriod struct {
	Sender        common.Address
	NewTwapPeriod uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTwapPeriod is a free log retrieval operation binding the contract event 0xe4c60f4984caeb7f45b0cfe6d4233c115601ab11d141bc2cbf68b48346cdef38.
//
// Solidity: event SetTwapPeriod(address sender, uint32 newTwapPeriod)
func (_Ichi *IchiFilterer) FilterSetTwapPeriod(opts *bind.FilterOpts) (*IchiSetTwapPeriodIterator, error) {

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "SetTwapPeriod")
	if err != nil {
		return nil, err
	}
	return &IchiSetTwapPeriodIterator{contract: _Ichi.contract, event: "SetTwapPeriod", logs: logs, sub: sub}, nil
}

// WatchSetTwapPeriod is a free log subscription operation binding the contract event 0xe4c60f4984caeb7f45b0cfe6d4233c115601ab11d141bc2cbf68b48346cdef38.
//
// Solidity: event SetTwapPeriod(address sender, uint32 newTwapPeriod)
func (_Ichi *IchiFilterer) WatchSetTwapPeriod(opts *bind.WatchOpts, sink chan<- *IchiSetTwapPeriod) (event.Subscription, error) {

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "SetTwapPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiSetTwapPeriod)
				if err := _Ichi.contract.UnpackLog(event, "SetTwapPeriod", log); err != nil {
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

// ParseSetTwapPeriod is a log parse operation binding the contract event 0xe4c60f4984caeb7f45b0cfe6d4233c115601ab11d141bc2cbf68b48346cdef38.
//
// Solidity: event SetTwapPeriod(address sender, uint32 newTwapPeriod)
func (_Ichi *IchiFilterer) ParseSetTwapPeriod(log types.Log) (*IchiSetTwapPeriod, error) {
	event := new(IchiSetTwapPeriod)
	if err := _Ichi.contract.UnpackLog(event, "SetTwapPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ichi contract.
type IchiTransferIterator struct {
	Event *IchiTransfer // Event containing the contract specifics and raw log

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
func (it *IchiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiTransfer)
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
		it.Event = new(IchiTransfer)
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
func (it *IchiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiTransfer represents a Transfer event raised by the Ichi contract.
type IchiTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ichi *IchiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IchiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IchiTransferIterator{contract: _Ichi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ichi *IchiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IchiTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiTransfer)
				if err := _Ichi.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Ichi *IchiFilterer) ParseTransfer(log types.Log) (*IchiTransfer, error) {
	event := new(IchiTransfer)
	if err := _Ichi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IchiWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Ichi contract.
type IchiWithdrawIterator struct {
	Event *IchiWithdraw // Event containing the contract specifics and raw log

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
func (it *IchiWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IchiWithdraw)
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
		it.Event = new(IchiWithdraw)
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
func (it *IchiWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IchiWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IchiWithdraw represents a Withdraw event raised by the Ichi contract.
type IchiWithdraw struct {
	Sender  common.Address
	To      common.Address
	Shares  *big.Int
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Ichi *IchiFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*IchiWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ichi.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IchiWithdrawIterator{contract: _Ichi.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Ichi *IchiFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IchiWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ichi.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IchiWithdraw)
				if err := _Ichi.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Ichi *IchiFilterer) ParseWithdraw(log types.Log) (*IchiWithdraw, error) {
	event := new(IchiWithdraw)
	if err := _Ichi.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
