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

// AquaBeraMetaData contains all meta data concerning the AquaBera contract.
var AquaBeraMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowToken0\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_allowToken1\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"__owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_twapPeriod\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"Affiliate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ammFeeRecipient\",\"type\":\"address\"}],\"name\":\"AmmFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowToken0\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowToken1\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"twapPeriod\",\"type\":\"uint256\"}],\"name\":\"DeployICHIVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit0Max\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit1Max\",\"type\":\"uint256\"}],\"name\":\"DepositMax\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hysteresis\",\"type\":\"uint256\"}],\"name\":\"Hysteresis\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newAuxTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"SetAuxTwapPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"SetTwapPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"affiliate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowToken0\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowToken1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ammFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auxTwapPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fees1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit0Max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit1Max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasePosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimitPosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hysteresis\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ichiVaultFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_baseLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_baseUpper\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_limitLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_limitUpper\",\"type\":\"int24\"},{\"internalType\":\"int256\",\"name\":\"swapQuantity\",\"type\":\"int256\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_affiliate\",\"type\":\"address\"}],\"name\":\"setAffiliate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ammFeeRecipient\",\"type\":\"address\"}],\"name\":\"setAmmFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newAuxTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"setAuxTwapPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deposit0Max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deposit1Max\",\"type\":\"uint256\"}],\"name\":\"setDepositMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hysteresis\",\"type\":\"uint256\"}],\"name\":\"setHysteresis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"setTwapPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twapPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AquaBeraABI is the input ABI used to generate the binding from.
// Deprecated: Use AquaBeraMetaData.ABI instead.
var AquaBeraABI = AquaBeraMetaData.ABI

// AquaBera is an auto generated Go binding around an Ethereum contract.
type AquaBera struct {
	AquaBeraCaller     // Read-only binding to the contract
	AquaBeraTransactor // Write-only binding to the contract
	AquaBeraFilterer   // Log filterer for contract events
}

// AquaBeraCaller is an auto generated read-only Go binding around an Ethereum contract.
type AquaBeraCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AquaBeraTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AquaBeraTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AquaBeraFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AquaBeraFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AquaBeraSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AquaBeraSession struct {
	Contract     *AquaBera         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AquaBeraCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AquaBeraCallerSession struct {
	Contract *AquaBeraCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AquaBeraTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AquaBeraTransactorSession struct {
	Contract     *AquaBeraTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AquaBeraRaw is an auto generated low-level Go binding around an Ethereum contract.
type AquaBeraRaw struct {
	Contract *AquaBera // Generic contract binding to access the raw methods on
}

// AquaBeraCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AquaBeraCallerRaw struct {
	Contract *AquaBeraCaller // Generic read-only contract binding to access the raw methods on
}

// AquaBeraTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AquaBeraTransactorRaw struct {
	Contract *AquaBeraTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAquaBera creates a new instance of AquaBera, bound to a specific deployed contract.
func NewAquaBera(address common.Address, backend bind.ContractBackend) (*AquaBera, error) {
	contract, err := bindAquaBera(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AquaBera{AquaBeraCaller: AquaBeraCaller{contract: contract}, AquaBeraTransactor: AquaBeraTransactor{contract: contract}, AquaBeraFilterer: AquaBeraFilterer{contract: contract}}, nil
}

// NewAquaBeraCaller creates a new read-only instance of AquaBera, bound to a specific deployed contract.
func NewAquaBeraCaller(address common.Address, caller bind.ContractCaller) (*AquaBeraCaller, error) {
	contract, err := bindAquaBera(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AquaBeraCaller{contract: contract}, nil
}

// NewAquaBeraTransactor creates a new write-only instance of AquaBera, bound to a specific deployed contract.
func NewAquaBeraTransactor(address common.Address, transactor bind.ContractTransactor) (*AquaBeraTransactor, error) {
	contract, err := bindAquaBera(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AquaBeraTransactor{contract: contract}, nil
}

// NewAquaBeraFilterer creates a new log filterer instance of AquaBera, bound to a specific deployed contract.
func NewAquaBeraFilterer(address common.Address, filterer bind.ContractFilterer) (*AquaBeraFilterer, error) {
	contract, err := bindAquaBera(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AquaBeraFilterer{contract: contract}, nil
}

// bindAquaBera binds a generic wrapper to an already deployed contract.
func bindAquaBera(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AquaBeraMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AquaBera *AquaBeraRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AquaBera.Contract.AquaBeraCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AquaBera *AquaBeraRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AquaBera.Contract.AquaBeraTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AquaBera *AquaBeraRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AquaBera.Contract.AquaBeraTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AquaBera *AquaBeraCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AquaBera.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AquaBera *AquaBeraTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AquaBera.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AquaBera *AquaBeraTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AquaBera.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_AquaBera *AquaBeraCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_AquaBera *AquaBeraSession) PRECISION() (*big.Int, error) {
	return _AquaBera.Contract.PRECISION(&_AquaBera.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_AquaBera *AquaBeraCallerSession) PRECISION() (*big.Int, error) {
	return _AquaBera.Contract.PRECISION(&_AquaBera.CallOpts)
}

// Affiliate is a free data retrieval call binding the contract method 0x45e05f43.
//
// Solidity: function affiliate() view returns(address)
func (_AquaBera *AquaBeraCaller) Affiliate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "affiliate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Affiliate is a free data retrieval call binding the contract method 0x45e05f43.
//
// Solidity: function affiliate() view returns(address)
func (_AquaBera *AquaBeraSession) Affiliate() (common.Address, error) {
	return _AquaBera.Contract.Affiliate(&_AquaBera.CallOpts)
}

// Affiliate is a free data retrieval call binding the contract method 0x45e05f43.
//
// Solidity: function affiliate() view returns(address)
func (_AquaBera *AquaBeraCallerSession) Affiliate() (common.Address, error) {
	return _AquaBera.Contract.Affiliate(&_AquaBera.CallOpts)
}

// AllowToken0 is a free data retrieval call binding the contract method 0x7f7a1eec.
//
// Solidity: function allowToken0() view returns(bool)
func (_AquaBera *AquaBeraCaller) AllowToken0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "allowToken0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowToken0 is a free data retrieval call binding the contract method 0x7f7a1eec.
//
// Solidity: function allowToken0() view returns(bool)
func (_AquaBera *AquaBeraSession) AllowToken0() (bool, error) {
	return _AquaBera.Contract.AllowToken0(&_AquaBera.CallOpts)
}

// AllowToken0 is a free data retrieval call binding the contract method 0x7f7a1eec.
//
// Solidity: function allowToken0() view returns(bool)
func (_AquaBera *AquaBeraCallerSession) AllowToken0() (bool, error) {
	return _AquaBera.Contract.AllowToken0(&_AquaBera.CallOpts)
}

// AllowToken1 is a free data retrieval call binding the contract method 0x37e41b40.
//
// Solidity: function allowToken1() view returns(bool)
func (_AquaBera *AquaBeraCaller) AllowToken1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "allowToken1")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowToken1 is a free data retrieval call binding the contract method 0x37e41b40.
//
// Solidity: function allowToken1() view returns(bool)
func (_AquaBera *AquaBeraSession) AllowToken1() (bool, error) {
	return _AquaBera.Contract.AllowToken1(&_AquaBera.CallOpts)
}

// AllowToken1 is a free data retrieval call binding the contract method 0x37e41b40.
//
// Solidity: function allowToken1() view returns(bool)
func (_AquaBera *AquaBeraCallerSession) AllowToken1() (bool, error) {
	return _AquaBera.Contract.AllowToken1(&_AquaBera.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AquaBera *AquaBeraCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AquaBera *AquaBeraSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AquaBera.Contract.Allowance(&_AquaBera.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AquaBera *AquaBeraCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AquaBera.Contract.Allowance(&_AquaBera.CallOpts, owner, spender)
}

// AmmFeeRecipient is a free data retrieval call binding the contract method 0x897f078c.
//
// Solidity: function ammFeeRecipient() view returns(address)
func (_AquaBera *AquaBeraCaller) AmmFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "ammFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmmFeeRecipient is a free data retrieval call binding the contract method 0x897f078c.
//
// Solidity: function ammFeeRecipient() view returns(address)
func (_AquaBera *AquaBeraSession) AmmFeeRecipient() (common.Address, error) {
	return _AquaBera.Contract.AmmFeeRecipient(&_AquaBera.CallOpts)
}

// AmmFeeRecipient is a free data retrieval call binding the contract method 0x897f078c.
//
// Solidity: function ammFeeRecipient() view returns(address)
func (_AquaBera *AquaBeraCallerSession) AmmFeeRecipient() (common.Address, error) {
	return _AquaBera.Contract.AmmFeeRecipient(&_AquaBera.CallOpts)
}

// AuxTwapPeriod is a free data retrieval call binding the contract method 0x91563d32.
//
// Solidity: function auxTwapPeriod() view returns(uint32)
func (_AquaBera *AquaBeraCaller) AuxTwapPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "auxTwapPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AuxTwapPeriod is a free data retrieval call binding the contract method 0x91563d32.
//
// Solidity: function auxTwapPeriod() view returns(uint32)
func (_AquaBera *AquaBeraSession) AuxTwapPeriod() (uint32, error) {
	return _AquaBera.Contract.AuxTwapPeriod(&_AquaBera.CallOpts)
}

// AuxTwapPeriod is a free data retrieval call binding the contract method 0x91563d32.
//
// Solidity: function auxTwapPeriod() view returns(uint32)
func (_AquaBera *AquaBeraCallerSession) AuxTwapPeriod() (uint32, error) {
	return _AquaBera.Contract.AuxTwapPeriod(&_AquaBera.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AquaBera *AquaBeraCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AquaBera *AquaBeraSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AquaBera.Contract.BalanceOf(&_AquaBera.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AquaBera *AquaBeraCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AquaBera.Contract.BalanceOf(&_AquaBera.CallOpts, account)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_AquaBera *AquaBeraCaller) BaseLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "baseLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_AquaBera *AquaBeraSession) BaseLower() (*big.Int, error) {
	return _AquaBera.Contract.BaseLower(&_AquaBera.CallOpts)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_AquaBera *AquaBeraCallerSession) BaseLower() (*big.Int, error) {
	return _AquaBera.Contract.BaseLower(&_AquaBera.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_AquaBera *AquaBeraCaller) BaseUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "baseUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_AquaBera *AquaBeraSession) BaseUpper() (*big.Int, error) {
	return _AquaBera.Contract.BaseUpper(&_AquaBera.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_AquaBera *AquaBeraCallerSession) BaseUpper() (*big.Int, error) {
	return _AquaBera.Contract.BaseUpper(&_AquaBera.CallOpts)
}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_AquaBera *AquaBeraCaller) CurrentTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "currentTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_AquaBera *AquaBeraSession) CurrentTick() (*big.Int, error) {
	return _AquaBera.Contract.CurrentTick(&_AquaBera.CallOpts)
}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_AquaBera *AquaBeraCallerSession) CurrentTick() (*big.Int, error) {
	return _AquaBera.Contract.CurrentTick(&_AquaBera.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AquaBera *AquaBeraCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AquaBera *AquaBeraSession) Decimals() (uint8, error) {
	return _AquaBera.Contract.Decimals(&_AquaBera.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AquaBera *AquaBeraCallerSession) Decimals() (uint8, error) {
	return _AquaBera.Contract.Decimals(&_AquaBera.CallOpts)
}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_AquaBera *AquaBeraCaller) Deposit0Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "deposit0Max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_AquaBera *AquaBeraSession) Deposit0Max() (*big.Int, error) {
	return _AquaBera.Contract.Deposit0Max(&_AquaBera.CallOpts)
}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_AquaBera *AquaBeraCallerSession) Deposit0Max() (*big.Int, error) {
	return _AquaBera.Contract.Deposit0Max(&_AquaBera.CallOpts)
}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_AquaBera *AquaBeraCaller) Deposit1Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "deposit1Max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_AquaBera *AquaBeraSession) Deposit1Max() (*big.Int, error) {
	return _AquaBera.Contract.Deposit1Max(&_AquaBera.CallOpts)
}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_AquaBera *AquaBeraCallerSession) Deposit1Max() (*big.Int, error) {
	return _AquaBera.Contract.Deposit1Max(&_AquaBera.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_AquaBera *AquaBeraCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_AquaBera *AquaBeraSession) Fee() (*big.Int, error) {
	return _AquaBera.Contract.Fee(&_AquaBera.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_AquaBera *AquaBeraCallerSession) Fee() (*big.Int, error) {
	return _AquaBera.Contract.Fee(&_AquaBera.CallOpts)
}

// GetBasePosition is a free data retrieval call binding the contract method 0xd2eabcfc.
//
// Solidity: function getBasePosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_AquaBera *AquaBeraCaller) GetBasePosition(opts *bind.CallOpts) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "getBasePosition")

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
func (_AquaBera *AquaBeraSession) GetBasePosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _AquaBera.Contract.GetBasePosition(&_AquaBera.CallOpts)
}

// GetBasePosition is a free data retrieval call binding the contract method 0xd2eabcfc.
//
// Solidity: function getBasePosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_AquaBera *AquaBeraCallerSession) GetBasePosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _AquaBera.Contract.GetBasePosition(&_AquaBera.CallOpts)
}

// GetLimitPosition is a free data retrieval call binding the contract method 0xa049de6b.
//
// Solidity: function getLimitPosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_AquaBera *AquaBeraCaller) GetLimitPosition(opts *bind.CallOpts) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "getLimitPosition")

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
func (_AquaBera *AquaBeraSession) GetLimitPosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _AquaBera.Contract.GetLimitPosition(&_AquaBera.CallOpts)
}

// GetLimitPosition is a free data retrieval call binding the contract method 0xa049de6b.
//
// Solidity: function getLimitPosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_AquaBera *AquaBeraCallerSession) GetLimitPosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _AquaBera.Contract.GetLimitPosition(&_AquaBera.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_AquaBera *AquaBeraCaller) GetTotalAmounts(opts *bind.CallOpts) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "getTotalAmounts")

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
func (_AquaBera *AquaBeraSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _AquaBera.Contract.GetTotalAmounts(&_AquaBera.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_AquaBera *AquaBeraCallerSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _AquaBera.Contract.GetTotalAmounts(&_AquaBera.CallOpts)
}

// Hysteresis is a free data retrieval call binding the contract method 0x7aea5309.
//
// Solidity: function hysteresis() view returns(uint256)
func (_AquaBera *AquaBeraCaller) Hysteresis(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "hysteresis")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hysteresis is a free data retrieval call binding the contract method 0x7aea5309.
//
// Solidity: function hysteresis() view returns(uint256)
func (_AquaBera *AquaBeraSession) Hysteresis() (*big.Int, error) {
	return _AquaBera.Contract.Hysteresis(&_AquaBera.CallOpts)
}

// Hysteresis is a free data retrieval call binding the contract method 0x7aea5309.
//
// Solidity: function hysteresis() view returns(uint256)
func (_AquaBera *AquaBeraCallerSession) Hysteresis() (*big.Int, error) {
	return _AquaBera.Contract.Hysteresis(&_AquaBera.CallOpts)
}

// IchiVaultFactory is a free data retrieval call binding the contract method 0xdd81fa63.
//
// Solidity: function ichiVaultFactory() view returns(address)
func (_AquaBera *AquaBeraCaller) IchiVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "ichiVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IchiVaultFactory is a free data retrieval call binding the contract method 0xdd81fa63.
//
// Solidity: function ichiVaultFactory() view returns(address)
func (_AquaBera *AquaBeraSession) IchiVaultFactory() (common.Address, error) {
	return _AquaBera.Contract.IchiVaultFactory(&_AquaBera.CallOpts)
}

// IchiVaultFactory is a free data retrieval call binding the contract method 0xdd81fa63.
//
// Solidity: function ichiVaultFactory() view returns(address)
func (_AquaBera *AquaBeraCallerSession) IchiVaultFactory() (common.Address, error) {
	return _AquaBera.Contract.IchiVaultFactory(&_AquaBera.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_AquaBera *AquaBeraCaller) LimitLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "limitLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_AquaBera *AquaBeraSession) LimitLower() (*big.Int, error) {
	return _AquaBera.Contract.LimitLower(&_AquaBera.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_AquaBera *AquaBeraCallerSession) LimitLower() (*big.Int, error) {
	return _AquaBera.Contract.LimitLower(&_AquaBera.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_AquaBera *AquaBeraCaller) LimitUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "limitUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_AquaBera *AquaBeraSession) LimitUpper() (*big.Int, error) {
	return _AquaBera.Contract.LimitUpper(&_AquaBera.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_AquaBera *AquaBeraCallerSession) LimitUpper() (*big.Int, error) {
	return _AquaBera.Contract.LimitUpper(&_AquaBera.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AquaBera *AquaBeraCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AquaBera *AquaBeraSession) Name() (string, error) {
	return _AquaBera.Contract.Name(&_AquaBera.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AquaBera *AquaBeraCallerSession) Name() (string, error) {
	return _AquaBera.Contract.Name(&_AquaBera.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AquaBera *AquaBeraCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AquaBera *AquaBeraSession) Owner() (common.Address, error) {
	return _AquaBera.Contract.Owner(&_AquaBera.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AquaBera *AquaBeraCallerSession) Owner() (common.Address, error) {
	return _AquaBera.Contract.Owner(&_AquaBera.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_AquaBera *AquaBeraCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_AquaBera *AquaBeraSession) Pool() (common.Address, error) {
	return _AquaBera.Contract.Pool(&_AquaBera.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_AquaBera *AquaBeraCallerSession) Pool() (common.Address, error) {
	return _AquaBera.Contract.Pool(&_AquaBera.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AquaBera *AquaBeraCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AquaBera *AquaBeraSession) Symbol() (string, error) {
	return _AquaBera.Contract.Symbol(&_AquaBera.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AquaBera *AquaBeraCallerSession) Symbol() (string, error) {
	return _AquaBera.Contract.Symbol(&_AquaBera.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AquaBera *AquaBeraCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AquaBera *AquaBeraSession) TickSpacing() (*big.Int, error) {
	return _AquaBera.Contract.TickSpacing(&_AquaBera.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AquaBera *AquaBeraCallerSession) TickSpacing() (*big.Int, error) {
	return _AquaBera.Contract.TickSpacing(&_AquaBera.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AquaBera *AquaBeraCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AquaBera *AquaBeraSession) Token0() (common.Address, error) {
	return _AquaBera.Contract.Token0(&_AquaBera.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AquaBera *AquaBeraCallerSession) Token0() (common.Address, error) {
	return _AquaBera.Contract.Token0(&_AquaBera.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AquaBera *AquaBeraCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AquaBera *AquaBeraSession) Token1() (common.Address, error) {
	return _AquaBera.Contract.Token1(&_AquaBera.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AquaBera *AquaBeraCallerSession) Token1() (common.Address, error) {
	return _AquaBera.Contract.Token1(&_AquaBera.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AquaBera *AquaBeraCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AquaBera *AquaBeraSession) TotalSupply() (*big.Int, error) {
	return _AquaBera.Contract.TotalSupply(&_AquaBera.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AquaBera *AquaBeraCallerSession) TotalSupply() (*big.Int, error) {
	return _AquaBera.Contract.TotalSupply(&_AquaBera.CallOpts)
}

// TwapPeriod is a free data retrieval call binding the contract method 0xf6207326.
//
// Solidity: function twapPeriod() view returns(uint32)
func (_AquaBera *AquaBeraCaller) TwapPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AquaBera.contract.Call(opts, &out, "twapPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TwapPeriod is a free data retrieval call binding the contract method 0xf6207326.
//
// Solidity: function twapPeriod() view returns(uint32)
func (_AquaBera *AquaBeraSession) TwapPeriod() (uint32, error) {
	return _AquaBera.Contract.TwapPeriod(&_AquaBera.CallOpts)
}

// TwapPeriod is a free data retrieval call binding the contract method 0xf6207326.
//
// Solidity: function twapPeriod() view returns(uint32)
func (_AquaBera *AquaBeraCallerSession) TwapPeriod() (uint32, error) {
	return _AquaBera.Contract.TwapPeriod(&_AquaBera.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AquaBera *AquaBeraTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AquaBera *AquaBeraSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.Approve(&_AquaBera.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_AquaBera *AquaBeraTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.Approve(&_AquaBera.TransactOpts, spender, amount)
}

// CollectFees is a paid mutator transaction binding the contract method 0xc8796572.
//
// Solidity: function collectFees() returns(uint256 fees0, uint256 fees1)
func (_AquaBera *AquaBeraTransactor) CollectFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "collectFees")
}

// CollectFees is a paid mutator transaction binding the contract method 0xc8796572.
//
// Solidity: function collectFees() returns(uint256 fees0, uint256 fees1)
func (_AquaBera *AquaBeraSession) CollectFees() (*types.Transaction, error) {
	return _AquaBera.Contract.CollectFees(&_AquaBera.TransactOpts)
}

// CollectFees is a paid mutator transaction binding the contract method 0xc8796572.
//
// Solidity: function collectFees() returns(uint256 fees0, uint256 fees1)
func (_AquaBera *AquaBeraTransactorSession) CollectFees() (*types.Transaction, error) {
	return _AquaBera.Contract.CollectFees(&_AquaBera.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AquaBera *AquaBeraTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AquaBera *AquaBeraSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.DecreaseAllowance(&_AquaBera.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_AquaBera *AquaBeraTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.DecreaseAllowance(&_AquaBera.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to) returns(uint256 shares)
func (_AquaBera *AquaBeraTransactor) Deposit(opts *bind.TransactOpts, deposit0 *big.Int, deposit1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "deposit", deposit0, deposit1, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to) returns(uint256 shares)
func (_AquaBera *AquaBeraSession) Deposit(deposit0 *big.Int, deposit1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _AquaBera.Contract.Deposit(&_AquaBera.TransactOpts, deposit0, deposit1, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to) returns(uint256 shares)
func (_AquaBera *AquaBeraTransactorSession) Deposit(deposit0 *big.Int, deposit1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _AquaBera.Contract.Deposit(&_AquaBera.TransactOpts, deposit0, deposit1, to)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AquaBera *AquaBeraTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AquaBera *AquaBeraSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.IncreaseAllowance(&_AquaBera.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_AquaBera *AquaBeraTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.IncreaseAllowance(&_AquaBera.TransactOpts, spender, addedValue)
}

// Rebalance is a paid mutator transaction binding the contract method 0xd87346aa.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, int256 swapQuantity) returns()
func (_AquaBera *AquaBeraTransactor) Rebalance(opts *bind.TransactOpts, _baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, swapQuantity *big.Int) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "rebalance", _baseLower, _baseUpper, _limitLower, _limitUpper, swapQuantity)
}

// Rebalance is a paid mutator transaction binding the contract method 0xd87346aa.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, int256 swapQuantity) returns()
func (_AquaBera *AquaBeraSession) Rebalance(_baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, swapQuantity *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.Rebalance(&_AquaBera.TransactOpts, _baseLower, _baseUpper, _limitLower, _limitUpper, swapQuantity)
}

// Rebalance is a paid mutator transaction binding the contract method 0xd87346aa.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, int256 swapQuantity) returns()
func (_AquaBera *AquaBeraTransactorSession) Rebalance(_baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, swapQuantity *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.Rebalance(&_AquaBera.TransactOpts, _baseLower, _baseUpper, _limitLower, _limitUpper, swapQuantity)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AquaBera *AquaBeraTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AquaBera *AquaBeraSession) RenounceOwnership() (*types.Transaction, error) {
	return _AquaBera.Contract.RenounceOwnership(&_AquaBera.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AquaBera *AquaBeraTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AquaBera.Contract.RenounceOwnership(&_AquaBera.TransactOpts)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x2bbb56d9.
//
// Solidity: function setAffiliate(address _affiliate) returns()
func (_AquaBera *AquaBeraTransactor) SetAffiliate(opts *bind.TransactOpts, _affiliate common.Address) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "setAffiliate", _affiliate)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x2bbb56d9.
//
// Solidity: function setAffiliate(address _affiliate) returns()
func (_AquaBera *AquaBeraSession) SetAffiliate(_affiliate common.Address) (*types.Transaction, error) {
	return _AquaBera.Contract.SetAffiliate(&_AquaBera.TransactOpts, _affiliate)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x2bbb56d9.
//
// Solidity: function setAffiliate(address _affiliate) returns()
func (_AquaBera *AquaBeraTransactorSession) SetAffiliate(_affiliate common.Address) (*types.Transaction, error) {
	return _AquaBera.Contract.SetAffiliate(&_AquaBera.TransactOpts, _affiliate)
}

// SetAmmFeeRecipient is a paid mutator transaction binding the contract method 0x81de128b.
//
// Solidity: function setAmmFeeRecipient(address _ammFeeRecipient) returns()
func (_AquaBera *AquaBeraTransactor) SetAmmFeeRecipient(opts *bind.TransactOpts, _ammFeeRecipient common.Address) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "setAmmFeeRecipient", _ammFeeRecipient)
}

// SetAmmFeeRecipient is a paid mutator transaction binding the contract method 0x81de128b.
//
// Solidity: function setAmmFeeRecipient(address _ammFeeRecipient) returns()
func (_AquaBera *AquaBeraSession) SetAmmFeeRecipient(_ammFeeRecipient common.Address) (*types.Transaction, error) {
	return _AquaBera.Contract.SetAmmFeeRecipient(&_AquaBera.TransactOpts, _ammFeeRecipient)
}

// SetAmmFeeRecipient is a paid mutator transaction binding the contract method 0x81de128b.
//
// Solidity: function setAmmFeeRecipient(address _ammFeeRecipient) returns()
func (_AquaBera *AquaBeraTransactorSession) SetAmmFeeRecipient(_ammFeeRecipient common.Address) (*types.Transaction, error) {
	return _AquaBera.Contract.SetAmmFeeRecipient(&_AquaBera.TransactOpts, _ammFeeRecipient)
}

// SetAuxTwapPeriod is a paid mutator transaction binding the contract method 0x400f0ceb.
//
// Solidity: function setAuxTwapPeriod(uint32 newAuxTwapPeriod) returns()
func (_AquaBera *AquaBeraTransactor) SetAuxTwapPeriod(opts *bind.TransactOpts, newAuxTwapPeriod uint32) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "setAuxTwapPeriod", newAuxTwapPeriod)
}

// SetAuxTwapPeriod is a paid mutator transaction binding the contract method 0x400f0ceb.
//
// Solidity: function setAuxTwapPeriod(uint32 newAuxTwapPeriod) returns()
func (_AquaBera *AquaBeraSession) SetAuxTwapPeriod(newAuxTwapPeriod uint32) (*types.Transaction, error) {
	return _AquaBera.Contract.SetAuxTwapPeriod(&_AquaBera.TransactOpts, newAuxTwapPeriod)
}

// SetAuxTwapPeriod is a paid mutator transaction binding the contract method 0x400f0ceb.
//
// Solidity: function setAuxTwapPeriod(uint32 newAuxTwapPeriod) returns()
func (_AquaBera *AquaBeraTransactorSession) SetAuxTwapPeriod(newAuxTwapPeriod uint32) (*types.Transaction, error) {
	return _AquaBera.Contract.SetAuxTwapPeriod(&_AquaBera.TransactOpts, newAuxTwapPeriod)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0x3e091ee9.
//
// Solidity: function setDepositMax(uint256 _deposit0Max, uint256 _deposit1Max) returns()
func (_AquaBera *AquaBeraTransactor) SetDepositMax(opts *bind.TransactOpts, _deposit0Max *big.Int, _deposit1Max *big.Int) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "setDepositMax", _deposit0Max, _deposit1Max)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0x3e091ee9.
//
// Solidity: function setDepositMax(uint256 _deposit0Max, uint256 _deposit1Max) returns()
func (_AquaBera *AquaBeraSession) SetDepositMax(_deposit0Max *big.Int, _deposit1Max *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.SetDepositMax(&_AquaBera.TransactOpts, _deposit0Max, _deposit1Max)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0x3e091ee9.
//
// Solidity: function setDepositMax(uint256 _deposit0Max, uint256 _deposit1Max) returns()
func (_AquaBera *AquaBeraTransactorSession) SetDepositMax(_deposit0Max *big.Int, _deposit1Max *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.SetDepositMax(&_AquaBera.TransactOpts, _deposit0Max, _deposit1Max)
}

// SetHysteresis is a paid mutator transaction binding the contract method 0x5ffc1ff7.
//
// Solidity: function setHysteresis(uint256 _hysteresis) returns()
func (_AquaBera *AquaBeraTransactor) SetHysteresis(opts *bind.TransactOpts, _hysteresis *big.Int) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "setHysteresis", _hysteresis)
}

// SetHysteresis is a paid mutator transaction binding the contract method 0x5ffc1ff7.
//
// Solidity: function setHysteresis(uint256 _hysteresis) returns()
func (_AquaBera *AquaBeraSession) SetHysteresis(_hysteresis *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.SetHysteresis(&_AquaBera.TransactOpts, _hysteresis)
}

// SetHysteresis is a paid mutator transaction binding the contract method 0x5ffc1ff7.
//
// Solidity: function setHysteresis(uint256 _hysteresis) returns()
func (_AquaBera *AquaBeraTransactorSession) SetHysteresis(_hysteresis *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.SetHysteresis(&_AquaBera.TransactOpts, _hysteresis)
}

// SetTwapPeriod is a paid mutator transaction binding the contract method 0xf9c95d46.
//
// Solidity: function setTwapPeriod(uint32 newTwapPeriod) returns()
func (_AquaBera *AquaBeraTransactor) SetTwapPeriod(opts *bind.TransactOpts, newTwapPeriod uint32) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "setTwapPeriod", newTwapPeriod)
}

// SetTwapPeriod is a paid mutator transaction binding the contract method 0xf9c95d46.
//
// Solidity: function setTwapPeriod(uint32 newTwapPeriod) returns()
func (_AquaBera *AquaBeraSession) SetTwapPeriod(newTwapPeriod uint32) (*types.Transaction, error) {
	return _AquaBera.Contract.SetTwapPeriod(&_AquaBera.TransactOpts, newTwapPeriod)
}

// SetTwapPeriod is a paid mutator transaction binding the contract method 0xf9c95d46.
//
// Solidity: function setTwapPeriod(uint32 newTwapPeriod) returns()
func (_AquaBera *AquaBeraTransactorSession) SetTwapPeriod(newTwapPeriod uint32) (*types.Transaction, error) {
	return _AquaBera.Contract.SetTwapPeriod(&_AquaBera.TransactOpts, newTwapPeriod)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AquaBera *AquaBeraTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AquaBera *AquaBeraSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.Transfer(&_AquaBera.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_AquaBera *AquaBeraTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.Transfer(&_AquaBera.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AquaBera *AquaBeraTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AquaBera *AquaBeraSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.TransferFrom(&_AquaBera.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_AquaBera *AquaBeraTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AquaBera.Contract.TransferFrom(&_AquaBera.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AquaBera *AquaBeraTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AquaBera *AquaBeraSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AquaBera.Contract.TransferOwnership(&_AquaBera.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AquaBera *AquaBeraTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AquaBera.Contract.TransferOwnership(&_AquaBera.TransactOpts, newOwner)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_AquaBera *AquaBeraTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "uniswapV3MintCallback", amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_AquaBera *AquaBeraSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AquaBera.Contract.UniswapV3MintCallback(&_AquaBera.TransactOpts, amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_AquaBera *AquaBeraTransactorSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AquaBera.Contract.UniswapV3MintCallback(&_AquaBera.TransactOpts, amount0, amount1, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_AquaBera *AquaBeraTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_AquaBera *AquaBeraSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _AquaBera.Contract.UniswapV3SwapCallback(&_AquaBera.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_AquaBera *AquaBeraTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _AquaBera.Contract.UniswapV3SwapCallback(&_AquaBera.TransactOpts, amount0Delta, amount1Delta, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address to) returns(uint256 amount0, uint256 amount1)
func (_AquaBera *AquaBeraTransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int, to common.Address) (*types.Transaction, error) {
	return _AquaBera.contract.Transact(opts, "withdraw", shares, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address to) returns(uint256 amount0, uint256 amount1)
func (_AquaBera *AquaBeraSession) Withdraw(shares *big.Int, to common.Address) (*types.Transaction, error) {
	return _AquaBera.Contract.Withdraw(&_AquaBera.TransactOpts, shares, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address to) returns(uint256 amount0, uint256 amount1)
func (_AquaBera *AquaBeraTransactorSession) Withdraw(shares *big.Int, to common.Address) (*types.Transaction, error) {
	return _AquaBera.Contract.Withdraw(&_AquaBera.TransactOpts, shares, to)
}

// AquaBeraAffiliateIterator is returned from FilterAffiliate and is used to iterate over the raw logs and unpacked data for Affiliate events raised by the AquaBera contract.
type AquaBeraAffiliateIterator struct {
	Event *AquaBeraAffiliate // Event containing the contract specifics and raw log

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
func (it *AquaBeraAffiliateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraAffiliate)
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
		it.Event = new(AquaBeraAffiliate)
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
func (it *AquaBeraAffiliateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraAffiliateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraAffiliate represents a Affiliate event raised by the AquaBera contract.
type AquaBeraAffiliate struct {
	Sender    common.Address
	Affiliate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAffiliate is a free log retrieval operation binding the contract event 0x3066ef5dd340e8b2ea28d62f5a8391eb7a82d3ee87532724a1ca4386d34f7523.
//
// Solidity: event Affiliate(address indexed sender, address affiliate)
func (_AquaBera *AquaBeraFilterer) FilterAffiliate(opts *bind.FilterOpts, sender []common.Address) (*AquaBeraAffiliateIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "Affiliate", senderRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraAffiliateIterator{contract: _AquaBera.contract, event: "Affiliate", logs: logs, sub: sub}, nil
}

// WatchAffiliate is a free log subscription operation binding the contract event 0x3066ef5dd340e8b2ea28d62f5a8391eb7a82d3ee87532724a1ca4386d34f7523.
//
// Solidity: event Affiliate(address indexed sender, address affiliate)
func (_AquaBera *AquaBeraFilterer) WatchAffiliate(opts *bind.WatchOpts, sink chan<- *AquaBeraAffiliate, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "Affiliate", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraAffiliate)
				if err := _AquaBera.contract.UnpackLog(event, "Affiliate", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseAffiliate(log types.Log) (*AquaBeraAffiliate, error) {
	event := new(AquaBeraAffiliate)
	if err := _AquaBera.contract.UnpackLog(event, "Affiliate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraAmmFeeRecipientIterator is returned from FilterAmmFeeRecipient and is used to iterate over the raw logs and unpacked data for AmmFeeRecipient events raised by the AquaBera contract.
type AquaBeraAmmFeeRecipientIterator struct {
	Event *AquaBeraAmmFeeRecipient // Event containing the contract specifics and raw log

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
func (it *AquaBeraAmmFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraAmmFeeRecipient)
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
		it.Event = new(AquaBeraAmmFeeRecipient)
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
func (it *AquaBeraAmmFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraAmmFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraAmmFeeRecipient represents a AmmFeeRecipient event raised by the AquaBera contract.
type AquaBeraAmmFeeRecipient struct {
	Sender          common.Address
	AmmFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAmmFeeRecipient is a free log retrieval operation binding the contract event 0xbb78b7c13893a913fa8c9ecb9fdaf97597aa412a39c778bf976790555f0942f7.
//
// Solidity: event AmmFeeRecipient(address indexed sender, address ammFeeRecipient)
func (_AquaBera *AquaBeraFilterer) FilterAmmFeeRecipient(opts *bind.FilterOpts, sender []common.Address) (*AquaBeraAmmFeeRecipientIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "AmmFeeRecipient", senderRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraAmmFeeRecipientIterator{contract: _AquaBera.contract, event: "AmmFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchAmmFeeRecipient is a free log subscription operation binding the contract event 0xbb78b7c13893a913fa8c9ecb9fdaf97597aa412a39c778bf976790555f0942f7.
//
// Solidity: event AmmFeeRecipient(address indexed sender, address ammFeeRecipient)
func (_AquaBera *AquaBeraFilterer) WatchAmmFeeRecipient(opts *bind.WatchOpts, sink chan<- *AquaBeraAmmFeeRecipient, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "AmmFeeRecipient", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraAmmFeeRecipient)
				if err := _AquaBera.contract.UnpackLog(event, "AmmFeeRecipient", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseAmmFeeRecipient(log types.Log) (*AquaBeraAmmFeeRecipient, error) {
	event := new(AquaBeraAmmFeeRecipient)
	if err := _AquaBera.contract.UnpackLog(event, "AmmFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AquaBera contract.
type AquaBeraApprovalIterator struct {
	Event *AquaBeraApproval // Event containing the contract specifics and raw log

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
func (it *AquaBeraApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraApproval)
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
		it.Event = new(AquaBeraApproval)
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
func (it *AquaBeraApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraApproval represents a Approval event raised by the AquaBera contract.
type AquaBeraApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AquaBera *AquaBeraFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AquaBeraApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraApprovalIterator{contract: _AquaBera.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AquaBera *AquaBeraFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AquaBeraApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraApproval)
				if err := _AquaBera.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseApproval(log types.Log) (*AquaBeraApproval, error) {
	event := new(AquaBeraApproval)
	if err := _AquaBera.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraCollectFeesIterator is returned from FilterCollectFees and is used to iterate over the raw logs and unpacked data for CollectFees events raised by the AquaBera contract.
type AquaBeraCollectFeesIterator struct {
	Event *AquaBeraCollectFees // Event containing the contract specifics and raw log

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
func (it *AquaBeraCollectFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraCollectFees)
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
		it.Event = new(AquaBeraCollectFees)
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
func (it *AquaBeraCollectFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraCollectFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraCollectFees represents a CollectFees event raised by the AquaBera contract.
type AquaBeraCollectFees struct {
	Sender     common.Address
	FeeAmount0 *big.Int
	FeeAmount1 *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0xec8208dd791fa8ffdc0d7427f3ba9c0ed06f1bce9a86254e6940c10cc1802fef.
//
// Solidity: event CollectFees(address indexed sender, uint256 feeAmount0, uint256 feeAmount1)
func (_AquaBera *AquaBeraFilterer) FilterCollectFees(opts *bind.FilterOpts, sender []common.Address) (*AquaBeraCollectFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "CollectFees", senderRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraCollectFeesIterator{contract: _AquaBera.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0xec8208dd791fa8ffdc0d7427f3ba9c0ed06f1bce9a86254e6940c10cc1802fef.
//
// Solidity: event CollectFees(address indexed sender, uint256 feeAmount0, uint256 feeAmount1)
func (_AquaBera *AquaBeraFilterer) WatchCollectFees(opts *bind.WatchOpts, sink chan<- *AquaBeraCollectFees, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "CollectFees", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraCollectFees)
				if err := _AquaBera.contract.UnpackLog(event, "CollectFees", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseCollectFees(log types.Log) (*AquaBeraCollectFees, error) {
	event := new(AquaBeraCollectFees)
	if err := _AquaBera.contract.UnpackLog(event, "CollectFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraDeployICHIVaultIterator is returned from FilterDeployICHIVault and is used to iterate over the raw logs and unpacked data for DeployICHIVault events raised by the AquaBera contract.
type AquaBeraDeployICHIVaultIterator struct {
	Event *AquaBeraDeployICHIVault // Event containing the contract specifics and raw log

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
func (it *AquaBeraDeployICHIVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraDeployICHIVault)
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
		it.Event = new(AquaBeraDeployICHIVault)
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
func (it *AquaBeraDeployICHIVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraDeployICHIVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraDeployICHIVault represents a DeployICHIVault event raised by the AquaBera contract.
type AquaBeraDeployICHIVault struct {
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
func (_AquaBera *AquaBeraFilterer) FilterDeployICHIVault(opts *bind.FilterOpts, sender []common.Address, pool []common.Address) (*AquaBeraDeployICHIVaultIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "DeployICHIVault", senderRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraDeployICHIVaultIterator{contract: _AquaBera.contract, event: "DeployICHIVault", logs: logs, sub: sub}, nil
}

// WatchDeployICHIVault is a free log subscription operation binding the contract event 0x3e708ccf7d0e6de8558e020ea36189511cb3435bbfec54e721a48ee4df0d4f8c.
//
// Solidity: event DeployICHIVault(address indexed sender, address indexed pool, bool allowToken0, bool allowToken1, address owner, uint256 twapPeriod)
func (_AquaBera *AquaBeraFilterer) WatchDeployICHIVault(opts *bind.WatchOpts, sink chan<- *AquaBeraDeployICHIVault, sender []common.Address, pool []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "DeployICHIVault", senderRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraDeployICHIVault)
				if err := _AquaBera.contract.UnpackLog(event, "DeployICHIVault", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseDeployICHIVault(log types.Log) (*AquaBeraDeployICHIVault, error) {
	event := new(AquaBeraDeployICHIVault)
	if err := _AquaBera.contract.UnpackLog(event, "DeployICHIVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the AquaBera contract.
type AquaBeraDepositIterator struct {
	Event *AquaBeraDeposit // Event containing the contract specifics and raw log

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
func (it *AquaBeraDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraDeposit)
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
		it.Event = new(AquaBeraDeposit)
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
func (it *AquaBeraDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraDeposit represents a Deposit event raised by the AquaBera contract.
type AquaBeraDeposit struct {
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
func (_AquaBera *AquaBeraFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*AquaBeraDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraDepositIterator{contract: _AquaBera.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_AquaBera *AquaBeraFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AquaBeraDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraDeposit)
				if err := _AquaBera.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseDeposit(log types.Log) (*AquaBeraDeposit, error) {
	event := new(AquaBeraDeposit)
	if err := _AquaBera.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraDepositMaxIterator is returned from FilterDepositMax and is used to iterate over the raw logs and unpacked data for DepositMax events raised by the AquaBera contract.
type AquaBeraDepositMaxIterator struct {
	Event *AquaBeraDepositMax // Event containing the contract specifics and raw log

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
func (it *AquaBeraDepositMaxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraDepositMax)
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
		it.Event = new(AquaBeraDepositMax)
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
func (it *AquaBeraDepositMaxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraDepositMaxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraDepositMax represents a DepositMax event raised by the AquaBera contract.
type AquaBeraDepositMax struct {
	Sender      common.Address
	Deposit0Max *big.Int
	Deposit1Max *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositMax is a free log retrieval operation binding the contract event 0xafd3b05a4086b378b6f291200a528d8aed8c5e0317af77436b001f1bec28821a.
//
// Solidity: event DepositMax(address indexed sender, uint256 deposit0Max, uint256 deposit1Max)
func (_AquaBera *AquaBeraFilterer) FilterDepositMax(opts *bind.FilterOpts, sender []common.Address) (*AquaBeraDepositMaxIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "DepositMax", senderRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraDepositMaxIterator{contract: _AquaBera.contract, event: "DepositMax", logs: logs, sub: sub}, nil
}

// WatchDepositMax is a free log subscription operation binding the contract event 0xafd3b05a4086b378b6f291200a528d8aed8c5e0317af77436b001f1bec28821a.
//
// Solidity: event DepositMax(address indexed sender, uint256 deposit0Max, uint256 deposit1Max)
func (_AquaBera *AquaBeraFilterer) WatchDepositMax(opts *bind.WatchOpts, sink chan<- *AquaBeraDepositMax, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "DepositMax", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraDepositMax)
				if err := _AquaBera.contract.UnpackLog(event, "DepositMax", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseDepositMax(log types.Log) (*AquaBeraDepositMax, error) {
	event := new(AquaBeraDepositMax)
	if err := _AquaBera.contract.UnpackLog(event, "DepositMax", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraHysteresisIterator is returned from FilterHysteresis and is used to iterate over the raw logs and unpacked data for Hysteresis events raised by the AquaBera contract.
type AquaBeraHysteresisIterator struct {
	Event *AquaBeraHysteresis // Event containing the contract specifics and raw log

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
func (it *AquaBeraHysteresisIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraHysteresis)
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
		it.Event = new(AquaBeraHysteresis)
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
func (it *AquaBeraHysteresisIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraHysteresisIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraHysteresis represents a Hysteresis event raised by the AquaBera contract.
type AquaBeraHysteresis struct {
	Sender     common.Address
	Hysteresis *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterHysteresis is a free log retrieval operation binding the contract event 0x529698f34660760dcb172def5c99d62e1b5b74b444df322e8f7da31f2bd0a86b.
//
// Solidity: event Hysteresis(address indexed sender, uint256 hysteresis)
func (_AquaBera *AquaBeraFilterer) FilterHysteresis(opts *bind.FilterOpts, sender []common.Address) (*AquaBeraHysteresisIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "Hysteresis", senderRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraHysteresisIterator{contract: _AquaBera.contract, event: "Hysteresis", logs: logs, sub: sub}, nil
}

// WatchHysteresis is a free log subscription operation binding the contract event 0x529698f34660760dcb172def5c99d62e1b5b74b444df322e8f7da31f2bd0a86b.
//
// Solidity: event Hysteresis(address indexed sender, uint256 hysteresis)
func (_AquaBera *AquaBeraFilterer) WatchHysteresis(opts *bind.WatchOpts, sink chan<- *AquaBeraHysteresis, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "Hysteresis", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraHysteresis)
				if err := _AquaBera.contract.UnpackLog(event, "Hysteresis", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseHysteresis(log types.Log) (*AquaBeraHysteresis, error) {
	event := new(AquaBeraHysteresis)
	if err := _AquaBera.contract.UnpackLog(event, "Hysteresis", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AquaBera contract.
type AquaBeraOwnershipTransferredIterator struct {
	Event *AquaBeraOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AquaBeraOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraOwnershipTransferred)
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
		it.Event = new(AquaBeraOwnershipTransferred)
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
func (it *AquaBeraOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraOwnershipTransferred represents a OwnershipTransferred event raised by the AquaBera contract.
type AquaBeraOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AquaBera *AquaBeraFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AquaBeraOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraOwnershipTransferredIterator{contract: _AquaBera.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AquaBera *AquaBeraFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AquaBeraOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraOwnershipTransferred)
				if err := _AquaBera.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseOwnershipTransferred(log types.Log) (*AquaBeraOwnershipTransferred, error) {
	event := new(AquaBeraOwnershipTransferred)
	if err := _AquaBera.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the AquaBera contract.
type AquaBeraRebalanceIterator struct {
	Event *AquaBeraRebalance // Event containing the contract specifics and raw log

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
func (it *AquaBeraRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraRebalance)
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
		it.Event = new(AquaBeraRebalance)
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
func (it *AquaBeraRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraRebalance represents a Rebalance event raised by the AquaBera contract.
type AquaBeraRebalance struct {
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
func (_AquaBera *AquaBeraFilterer) FilterRebalance(opts *bind.FilterOpts) (*AquaBeraRebalanceIterator, error) {

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return &AquaBeraRebalanceIterator{contract: _AquaBera.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0xbc4c20ad04f161d631d9ce94d27659391196415aa3c42f6a71c62e905ece782d.
//
// Solidity: event Rebalance(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 feeAmount0, uint256 feeAmount1, uint256 totalSupply)
func (_AquaBera *AquaBeraFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *AquaBeraRebalance) (event.Subscription, error) {

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraRebalance)
				if err := _AquaBera.contract.UnpackLog(event, "Rebalance", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseRebalance(log types.Log) (*AquaBeraRebalance, error) {
	event := new(AquaBeraRebalance)
	if err := _AquaBera.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraSetAuxTwapPeriodIterator is returned from FilterSetAuxTwapPeriod and is used to iterate over the raw logs and unpacked data for SetAuxTwapPeriod events raised by the AquaBera contract.
type AquaBeraSetAuxTwapPeriodIterator struct {
	Event *AquaBeraSetAuxTwapPeriod // Event containing the contract specifics and raw log

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
func (it *AquaBeraSetAuxTwapPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraSetAuxTwapPeriod)
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
		it.Event = new(AquaBeraSetAuxTwapPeriod)
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
func (it *AquaBeraSetAuxTwapPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraSetAuxTwapPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraSetAuxTwapPeriod represents a SetAuxTwapPeriod event raised by the AquaBera contract.
type AquaBeraSetAuxTwapPeriod struct {
	Sender           common.Address
	NewAuxTwapPeriod uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetAuxTwapPeriod is a free log retrieval operation binding the contract event 0x39da19f5960a3f182ced1ff1853b7be54f37150799b3003a40bf4e0d4c740c85.
//
// Solidity: event SetAuxTwapPeriod(address sender, uint32 newAuxTwapPeriod)
func (_AquaBera *AquaBeraFilterer) FilterSetAuxTwapPeriod(opts *bind.FilterOpts) (*AquaBeraSetAuxTwapPeriodIterator, error) {

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "SetAuxTwapPeriod")
	if err != nil {
		return nil, err
	}
	return &AquaBeraSetAuxTwapPeriodIterator{contract: _AquaBera.contract, event: "SetAuxTwapPeriod", logs: logs, sub: sub}, nil
}

// WatchSetAuxTwapPeriod is a free log subscription operation binding the contract event 0x39da19f5960a3f182ced1ff1853b7be54f37150799b3003a40bf4e0d4c740c85.
//
// Solidity: event SetAuxTwapPeriod(address sender, uint32 newAuxTwapPeriod)
func (_AquaBera *AquaBeraFilterer) WatchSetAuxTwapPeriod(opts *bind.WatchOpts, sink chan<- *AquaBeraSetAuxTwapPeriod) (event.Subscription, error) {

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "SetAuxTwapPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraSetAuxTwapPeriod)
				if err := _AquaBera.contract.UnpackLog(event, "SetAuxTwapPeriod", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseSetAuxTwapPeriod(log types.Log) (*AquaBeraSetAuxTwapPeriod, error) {
	event := new(AquaBeraSetAuxTwapPeriod)
	if err := _AquaBera.contract.UnpackLog(event, "SetAuxTwapPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraSetTwapPeriodIterator is returned from FilterSetTwapPeriod and is used to iterate over the raw logs and unpacked data for SetTwapPeriod events raised by the AquaBera contract.
type AquaBeraSetTwapPeriodIterator struct {
	Event *AquaBeraSetTwapPeriod // Event containing the contract specifics and raw log

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
func (it *AquaBeraSetTwapPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraSetTwapPeriod)
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
		it.Event = new(AquaBeraSetTwapPeriod)
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
func (it *AquaBeraSetTwapPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraSetTwapPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraSetTwapPeriod represents a SetTwapPeriod event raised by the AquaBera contract.
type AquaBeraSetTwapPeriod struct {
	Sender        common.Address
	NewTwapPeriod uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTwapPeriod is a free log retrieval operation binding the contract event 0xe4c60f4984caeb7f45b0cfe6d4233c115601ab11d141bc2cbf68b48346cdef38.
//
// Solidity: event SetTwapPeriod(address sender, uint32 newTwapPeriod)
func (_AquaBera *AquaBeraFilterer) FilterSetTwapPeriod(opts *bind.FilterOpts) (*AquaBeraSetTwapPeriodIterator, error) {

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "SetTwapPeriod")
	if err != nil {
		return nil, err
	}
	return &AquaBeraSetTwapPeriodIterator{contract: _AquaBera.contract, event: "SetTwapPeriod", logs: logs, sub: sub}, nil
}

// WatchSetTwapPeriod is a free log subscription operation binding the contract event 0xe4c60f4984caeb7f45b0cfe6d4233c115601ab11d141bc2cbf68b48346cdef38.
//
// Solidity: event SetTwapPeriod(address sender, uint32 newTwapPeriod)
func (_AquaBera *AquaBeraFilterer) WatchSetTwapPeriod(opts *bind.WatchOpts, sink chan<- *AquaBeraSetTwapPeriod) (event.Subscription, error) {

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "SetTwapPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraSetTwapPeriod)
				if err := _AquaBera.contract.UnpackLog(event, "SetTwapPeriod", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseSetTwapPeriod(log types.Log) (*AquaBeraSetTwapPeriod, error) {
	event := new(AquaBeraSetTwapPeriod)
	if err := _AquaBera.contract.UnpackLog(event, "SetTwapPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AquaBera contract.
type AquaBeraTransferIterator struct {
	Event *AquaBeraTransfer // Event containing the contract specifics and raw log

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
func (it *AquaBeraTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraTransfer)
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
		it.Event = new(AquaBeraTransfer)
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
func (it *AquaBeraTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraTransfer represents a Transfer event raised by the AquaBera contract.
type AquaBeraTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AquaBera *AquaBeraFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AquaBeraTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraTransferIterator{contract: _AquaBera.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AquaBera *AquaBeraFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AquaBeraTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraTransfer)
				if err := _AquaBera.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseTransfer(log types.Log) (*AquaBeraTransfer, error) {
	event := new(AquaBeraTransfer)
	if err := _AquaBera.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AquaBeraWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AquaBera contract.
type AquaBeraWithdrawIterator struct {
	Event *AquaBeraWithdraw // Event containing the contract specifics and raw log

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
func (it *AquaBeraWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AquaBeraWithdraw)
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
		it.Event = new(AquaBeraWithdraw)
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
func (it *AquaBeraWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AquaBeraWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AquaBeraWithdraw represents a Withdraw event raised by the AquaBera contract.
type AquaBeraWithdraw struct {
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
func (_AquaBera *AquaBeraFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*AquaBeraWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AquaBera.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AquaBeraWithdrawIterator{contract: _AquaBera.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_AquaBera *AquaBeraFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AquaBeraWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AquaBera.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AquaBeraWithdraw)
				if err := _AquaBera.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_AquaBera *AquaBeraFilterer) ParseWithdraw(log types.Log) (*AquaBeraWithdraw, error) {
	event := new(AquaBeraWithdraw)
	if err := _AquaBera.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
