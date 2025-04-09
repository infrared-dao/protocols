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

// WasabeeVaultMetaData contains all meta data concerning the WasabeeVault contract.
var WasabeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowToken0\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_allowToken1\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"__owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_twapPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_vaultIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"Affiliate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ammFeeRecipient\",\"type\":\"address\"}],\"name\":\"AmmFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowToken0\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowToken1\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"twapPeriod\",\"type\":\"uint256\"}],\"name\":\"DeployICHIVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit0Max\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit1Max\",\"type\":\"uint256\"}],\"name\":\"DepositMax\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hysteresis\",\"type\":\"uint256\"}],\"name\":\"Hysteresis\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newAuxTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"SetAuxTwapPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"SetTwapPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"affiliate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"algebraMintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"algebraSwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowToken0\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowToken1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ammFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auxTwapPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fees0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fees1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit0Max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit1Max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"fee_\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasePosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimitPosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hysteresis\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ichiVaultFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_baseLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_baseUpper\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_limitLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_limitUpper\",\"type\":\"int24\"},{\"internalType\":\"int256\",\"name\":\"swapQuantity\",\"type\":\"int256\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_affiliate\",\"type\":\"address\"}],\"name\":\"setAffiliate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ammFeeRecipient\",\"type\":\"address\"}],\"name\":\"setAmmFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newAuxTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"setAuxTwapPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deposit0Max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deposit1Max\",\"type\":\"uint256\"}],\"name\":\"setDepositMax\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hysteresis\",\"type\":\"uint256\"}],\"name\":\"setHysteresis\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newTwapPeriod\",\"type\":\"uint32\"}],\"name\":\"setTwapPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twapPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WasabeeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use WasabeeVaultMetaData.ABI instead.
var WasabeeVaultABI = WasabeeVaultMetaData.ABI

// WasabeeVault is an auto generated Go binding around an Ethereum contract.
type WasabeeVault struct {
	WasabeeVaultCaller     // Read-only binding to the contract
	WasabeeVaultTransactor // Write-only binding to the contract
	WasabeeVaultFilterer   // Log filterer for contract events
}

// WasabeeVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type WasabeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WasabeeVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WasabeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WasabeeVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WasabeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WasabeeVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WasabeeVaultSession struct {
	Contract     *WasabeeVault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WasabeeVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WasabeeVaultCallerSession struct {
	Contract *WasabeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WasabeeVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WasabeeVaultTransactorSession struct {
	Contract     *WasabeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WasabeeVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type WasabeeVaultRaw struct {
	Contract *WasabeeVault // Generic contract binding to access the raw methods on
}

// WasabeeVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WasabeeVaultCallerRaw struct {
	Contract *WasabeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// WasabeeVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WasabeeVaultTransactorRaw struct {
	Contract *WasabeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWasabeeVault creates a new instance of WasabeeVault, bound to a specific deployed contract.
func NewWasabeeVault(address common.Address, backend bind.ContractBackend) (*WasabeeVault, error) {
	contract, err := bindWasabeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WasabeeVault{WasabeeVaultCaller: WasabeeVaultCaller{contract: contract}, WasabeeVaultTransactor: WasabeeVaultTransactor{contract: contract}, WasabeeVaultFilterer: WasabeeVaultFilterer{contract: contract}}, nil
}

// NewWasabeeVaultCaller creates a new read-only instance of WasabeeVault, bound to a specific deployed contract.
func NewWasabeeVaultCaller(address common.Address, caller bind.ContractCaller) (*WasabeeVaultCaller, error) {
	contract, err := bindWasabeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultCaller{contract: contract}, nil
}

// NewWasabeeVaultTransactor creates a new write-only instance of WasabeeVault, bound to a specific deployed contract.
func NewWasabeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*WasabeeVaultTransactor, error) {
	contract, err := bindWasabeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultTransactor{contract: contract}, nil
}

// NewWasabeeVaultFilterer creates a new log filterer instance of WasabeeVault, bound to a specific deployed contract.
func NewWasabeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*WasabeeVaultFilterer, error) {
	contract, err := bindWasabeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultFilterer{contract: contract}, nil
}

// bindWasabeeVault binds a generic wrapper to an already deployed contract.
func bindWasabeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WasabeeVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WasabeeVault *WasabeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WasabeeVault.Contract.WasabeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WasabeeVault *WasabeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WasabeeVault.Contract.WasabeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WasabeeVault *WasabeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WasabeeVault.Contract.WasabeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WasabeeVault *WasabeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WasabeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WasabeeVault *WasabeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WasabeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WasabeeVault *WasabeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WasabeeVault.Contract.contract.Transact(opts, method, params...)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_WasabeeVault *WasabeeVaultCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_WasabeeVault *WasabeeVaultSession) PRECISION() (*big.Int, error) {
	return _WasabeeVault.Contract.PRECISION(&_WasabeeVault.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_WasabeeVault *WasabeeVaultCallerSession) PRECISION() (*big.Int, error) {
	return _WasabeeVault.Contract.PRECISION(&_WasabeeVault.CallOpts)
}

// Affiliate is a free data retrieval call binding the contract method 0x45e05f43.
//
// Solidity: function affiliate() view returns(address)
func (_WasabeeVault *WasabeeVaultCaller) Affiliate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "affiliate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Affiliate is a free data retrieval call binding the contract method 0x45e05f43.
//
// Solidity: function affiliate() view returns(address)
func (_WasabeeVault *WasabeeVaultSession) Affiliate() (common.Address, error) {
	return _WasabeeVault.Contract.Affiliate(&_WasabeeVault.CallOpts)
}

// Affiliate is a free data retrieval call binding the contract method 0x45e05f43.
//
// Solidity: function affiliate() view returns(address)
func (_WasabeeVault *WasabeeVaultCallerSession) Affiliate() (common.Address, error) {
	return _WasabeeVault.Contract.Affiliate(&_WasabeeVault.CallOpts)
}

// AllowToken0 is a free data retrieval call binding the contract method 0x7f7a1eec.
//
// Solidity: function allowToken0() view returns(bool)
func (_WasabeeVault *WasabeeVaultCaller) AllowToken0(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "allowToken0")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowToken0 is a free data retrieval call binding the contract method 0x7f7a1eec.
//
// Solidity: function allowToken0() view returns(bool)
func (_WasabeeVault *WasabeeVaultSession) AllowToken0() (bool, error) {
	return _WasabeeVault.Contract.AllowToken0(&_WasabeeVault.CallOpts)
}

// AllowToken0 is a free data retrieval call binding the contract method 0x7f7a1eec.
//
// Solidity: function allowToken0() view returns(bool)
func (_WasabeeVault *WasabeeVaultCallerSession) AllowToken0() (bool, error) {
	return _WasabeeVault.Contract.AllowToken0(&_WasabeeVault.CallOpts)
}

// AllowToken1 is a free data retrieval call binding the contract method 0x37e41b40.
//
// Solidity: function allowToken1() view returns(bool)
func (_WasabeeVault *WasabeeVaultCaller) AllowToken1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "allowToken1")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowToken1 is a free data retrieval call binding the contract method 0x37e41b40.
//
// Solidity: function allowToken1() view returns(bool)
func (_WasabeeVault *WasabeeVaultSession) AllowToken1() (bool, error) {
	return _WasabeeVault.Contract.AllowToken1(&_WasabeeVault.CallOpts)
}

// AllowToken1 is a free data retrieval call binding the contract method 0x37e41b40.
//
// Solidity: function allowToken1() view returns(bool)
func (_WasabeeVault *WasabeeVaultCallerSession) AllowToken1() (bool, error) {
	return _WasabeeVault.Contract.AllowToken1(&_WasabeeVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WasabeeVault *WasabeeVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WasabeeVault *WasabeeVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WasabeeVault.Contract.Allowance(&_WasabeeVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WasabeeVault *WasabeeVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WasabeeVault.Contract.Allowance(&_WasabeeVault.CallOpts, owner, spender)
}

// AmmFeeRecipient is a free data retrieval call binding the contract method 0x897f078c.
//
// Solidity: function ammFeeRecipient() view returns(address)
func (_WasabeeVault *WasabeeVaultCaller) AmmFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "ammFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmmFeeRecipient is a free data retrieval call binding the contract method 0x897f078c.
//
// Solidity: function ammFeeRecipient() view returns(address)
func (_WasabeeVault *WasabeeVaultSession) AmmFeeRecipient() (common.Address, error) {
	return _WasabeeVault.Contract.AmmFeeRecipient(&_WasabeeVault.CallOpts)
}

// AmmFeeRecipient is a free data retrieval call binding the contract method 0x897f078c.
//
// Solidity: function ammFeeRecipient() view returns(address)
func (_WasabeeVault *WasabeeVaultCallerSession) AmmFeeRecipient() (common.Address, error) {
	return _WasabeeVault.Contract.AmmFeeRecipient(&_WasabeeVault.CallOpts)
}

// AuxTwapPeriod is a free data retrieval call binding the contract method 0x91563d32.
//
// Solidity: function auxTwapPeriod() view returns(uint32)
func (_WasabeeVault *WasabeeVaultCaller) AuxTwapPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "auxTwapPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AuxTwapPeriod is a free data retrieval call binding the contract method 0x91563d32.
//
// Solidity: function auxTwapPeriod() view returns(uint32)
func (_WasabeeVault *WasabeeVaultSession) AuxTwapPeriod() (uint32, error) {
	return _WasabeeVault.Contract.AuxTwapPeriod(&_WasabeeVault.CallOpts)
}

// AuxTwapPeriod is a free data retrieval call binding the contract method 0x91563d32.
//
// Solidity: function auxTwapPeriod() view returns(uint32)
func (_WasabeeVault *WasabeeVaultCallerSession) AuxTwapPeriod() (uint32, error) {
	return _WasabeeVault.Contract.AuxTwapPeriod(&_WasabeeVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WasabeeVault *WasabeeVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WasabeeVault *WasabeeVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WasabeeVault.Contract.BalanceOf(&_WasabeeVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WasabeeVault *WasabeeVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WasabeeVault.Contract.BalanceOf(&_WasabeeVault.CallOpts, account)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_WasabeeVault *WasabeeVaultCaller) BaseLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "baseLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_WasabeeVault *WasabeeVaultSession) BaseLower() (*big.Int, error) {
	return _WasabeeVault.Contract.BaseLower(&_WasabeeVault.CallOpts)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_WasabeeVault *WasabeeVaultCallerSession) BaseLower() (*big.Int, error) {
	return _WasabeeVault.Contract.BaseLower(&_WasabeeVault.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_WasabeeVault *WasabeeVaultCaller) BaseUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "baseUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_WasabeeVault *WasabeeVaultSession) BaseUpper() (*big.Int, error) {
	return _WasabeeVault.Contract.BaseUpper(&_WasabeeVault.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_WasabeeVault *WasabeeVaultCallerSession) BaseUpper() (*big.Int, error) {
	return _WasabeeVault.Contract.BaseUpper(&_WasabeeVault.CallOpts)
}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_WasabeeVault *WasabeeVaultCaller) CurrentTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "currentTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_WasabeeVault *WasabeeVaultSession) CurrentTick() (*big.Int, error) {
	return _WasabeeVault.Contract.CurrentTick(&_WasabeeVault.CallOpts)
}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_WasabeeVault *WasabeeVaultCallerSession) CurrentTick() (*big.Int, error) {
	return _WasabeeVault.Contract.CurrentTick(&_WasabeeVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WasabeeVault *WasabeeVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WasabeeVault *WasabeeVaultSession) Decimals() (uint8, error) {
	return _WasabeeVault.Contract.Decimals(&_WasabeeVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WasabeeVault *WasabeeVaultCallerSession) Decimals() (uint8, error) {
	return _WasabeeVault.Contract.Decimals(&_WasabeeVault.CallOpts)
}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_WasabeeVault *WasabeeVaultCaller) Deposit0Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "deposit0Max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_WasabeeVault *WasabeeVaultSession) Deposit0Max() (*big.Int, error) {
	return _WasabeeVault.Contract.Deposit0Max(&_WasabeeVault.CallOpts)
}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_WasabeeVault *WasabeeVaultCallerSession) Deposit0Max() (*big.Int, error) {
	return _WasabeeVault.Contract.Deposit0Max(&_WasabeeVault.CallOpts)
}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_WasabeeVault *WasabeeVaultCaller) Deposit1Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "deposit1Max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_WasabeeVault *WasabeeVaultSession) Deposit1Max() (*big.Int, error) {
	return _WasabeeVault.Contract.Deposit1Max(&_WasabeeVault.CallOpts)
}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_WasabeeVault *WasabeeVaultCallerSession) Deposit1Max() (*big.Int, error) {
	return _WasabeeVault.Contract.Deposit1Max(&_WasabeeVault.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24 fee_)
func (_WasabeeVault *WasabeeVaultCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24 fee_)
func (_WasabeeVault *WasabeeVaultSession) Fee() (*big.Int, error) {
	return _WasabeeVault.Contract.Fee(&_WasabeeVault.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24 fee_)
func (_WasabeeVault *WasabeeVaultCallerSession) Fee() (*big.Int, error) {
	return _WasabeeVault.Contract.Fee(&_WasabeeVault.CallOpts)
}

// GetBasePosition is a free data retrieval call binding the contract method 0xd2eabcfc.
//
// Solidity: function getBasePosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_WasabeeVault *WasabeeVaultCaller) GetBasePosition(opts *bind.CallOpts) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "getBasePosition")

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
func (_WasabeeVault *WasabeeVaultSession) GetBasePosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _WasabeeVault.Contract.GetBasePosition(&_WasabeeVault.CallOpts)
}

// GetBasePosition is a free data retrieval call binding the contract method 0xd2eabcfc.
//
// Solidity: function getBasePosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_WasabeeVault *WasabeeVaultCallerSession) GetBasePosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _WasabeeVault.Contract.GetBasePosition(&_WasabeeVault.CallOpts)
}

// GetLimitPosition is a free data retrieval call binding the contract method 0xa049de6b.
//
// Solidity: function getLimitPosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_WasabeeVault *WasabeeVaultCaller) GetLimitPosition(opts *bind.CallOpts) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "getLimitPosition")

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
func (_WasabeeVault *WasabeeVaultSession) GetLimitPosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _WasabeeVault.Contract.GetLimitPosition(&_WasabeeVault.CallOpts)
}

// GetLimitPosition is a free data retrieval call binding the contract method 0xa049de6b.
//
// Solidity: function getLimitPosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_WasabeeVault *WasabeeVaultCallerSession) GetLimitPosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _WasabeeVault.Contract.GetLimitPosition(&_WasabeeVault.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_WasabeeVault *WasabeeVaultCaller) GetTotalAmounts(opts *bind.CallOpts) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "getTotalAmounts")

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
func (_WasabeeVault *WasabeeVaultSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _WasabeeVault.Contract.GetTotalAmounts(&_WasabeeVault.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_WasabeeVault *WasabeeVaultCallerSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _WasabeeVault.Contract.GetTotalAmounts(&_WasabeeVault.CallOpts)
}

// Hysteresis is a free data retrieval call binding the contract method 0x7aea5309.
//
// Solidity: function hysteresis() view returns(uint256)
func (_WasabeeVault *WasabeeVaultCaller) Hysteresis(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "hysteresis")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hysteresis is a free data retrieval call binding the contract method 0x7aea5309.
//
// Solidity: function hysteresis() view returns(uint256)
func (_WasabeeVault *WasabeeVaultSession) Hysteresis() (*big.Int, error) {
	return _WasabeeVault.Contract.Hysteresis(&_WasabeeVault.CallOpts)
}

// Hysteresis is a free data retrieval call binding the contract method 0x7aea5309.
//
// Solidity: function hysteresis() view returns(uint256)
func (_WasabeeVault *WasabeeVaultCallerSession) Hysteresis() (*big.Int, error) {
	return _WasabeeVault.Contract.Hysteresis(&_WasabeeVault.CallOpts)
}

// IchiVaultFactory is a free data retrieval call binding the contract method 0xdd81fa63.
//
// Solidity: function ichiVaultFactory() view returns(address)
func (_WasabeeVault *WasabeeVaultCaller) IchiVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "ichiVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IchiVaultFactory is a free data retrieval call binding the contract method 0xdd81fa63.
//
// Solidity: function ichiVaultFactory() view returns(address)
func (_WasabeeVault *WasabeeVaultSession) IchiVaultFactory() (common.Address, error) {
	return _WasabeeVault.Contract.IchiVaultFactory(&_WasabeeVault.CallOpts)
}

// IchiVaultFactory is a free data retrieval call binding the contract method 0xdd81fa63.
//
// Solidity: function ichiVaultFactory() view returns(address)
func (_WasabeeVault *WasabeeVaultCallerSession) IchiVaultFactory() (common.Address, error) {
	return _WasabeeVault.Contract.IchiVaultFactory(&_WasabeeVault.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_WasabeeVault *WasabeeVaultCaller) LimitLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "limitLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_WasabeeVault *WasabeeVaultSession) LimitLower() (*big.Int, error) {
	return _WasabeeVault.Contract.LimitLower(&_WasabeeVault.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_WasabeeVault *WasabeeVaultCallerSession) LimitLower() (*big.Int, error) {
	return _WasabeeVault.Contract.LimitLower(&_WasabeeVault.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_WasabeeVault *WasabeeVaultCaller) LimitUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "limitUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_WasabeeVault *WasabeeVaultSession) LimitUpper() (*big.Int, error) {
	return _WasabeeVault.Contract.LimitUpper(&_WasabeeVault.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_WasabeeVault *WasabeeVaultCallerSession) LimitUpper() (*big.Int, error) {
	return _WasabeeVault.Contract.LimitUpper(&_WasabeeVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WasabeeVault *WasabeeVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WasabeeVault *WasabeeVaultSession) Name() (string, error) {
	return _WasabeeVault.Contract.Name(&_WasabeeVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WasabeeVault *WasabeeVaultCallerSession) Name() (string, error) {
	return _WasabeeVault.Contract.Name(&_WasabeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WasabeeVault *WasabeeVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WasabeeVault *WasabeeVaultSession) Owner() (common.Address, error) {
	return _WasabeeVault.Contract.Owner(&_WasabeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WasabeeVault *WasabeeVaultCallerSession) Owner() (common.Address, error) {
	return _WasabeeVault.Contract.Owner(&_WasabeeVault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_WasabeeVault *WasabeeVaultCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_WasabeeVault *WasabeeVaultSession) Pool() (common.Address, error) {
	return _WasabeeVault.Contract.Pool(&_WasabeeVault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_WasabeeVault *WasabeeVaultCallerSession) Pool() (common.Address, error) {
	return _WasabeeVault.Contract.Pool(&_WasabeeVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WasabeeVault *WasabeeVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WasabeeVault *WasabeeVaultSession) Symbol() (string, error) {
	return _WasabeeVault.Contract.Symbol(&_WasabeeVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WasabeeVault *WasabeeVaultCallerSession) Symbol() (string, error) {
	return _WasabeeVault.Contract.Symbol(&_WasabeeVault.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_WasabeeVault *WasabeeVaultCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_WasabeeVault *WasabeeVaultSession) TickSpacing() (*big.Int, error) {
	return _WasabeeVault.Contract.TickSpacing(&_WasabeeVault.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_WasabeeVault *WasabeeVaultCallerSession) TickSpacing() (*big.Int, error) {
	return _WasabeeVault.Contract.TickSpacing(&_WasabeeVault.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_WasabeeVault *WasabeeVaultCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_WasabeeVault *WasabeeVaultSession) Token0() (common.Address, error) {
	return _WasabeeVault.Contract.Token0(&_WasabeeVault.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_WasabeeVault *WasabeeVaultCallerSession) Token0() (common.Address, error) {
	return _WasabeeVault.Contract.Token0(&_WasabeeVault.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_WasabeeVault *WasabeeVaultCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_WasabeeVault *WasabeeVaultSession) Token1() (common.Address, error) {
	return _WasabeeVault.Contract.Token1(&_WasabeeVault.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_WasabeeVault *WasabeeVaultCallerSession) Token1() (common.Address, error) {
	return _WasabeeVault.Contract.Token1(&_WasabeeVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WasabeeVault *WasabeeVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WasabeeVault *WasabeeVaultSession) TotalSupply() (*big.Int, error) {
	return _WasabeeVault.Contract.TotalSupply(&_WasabeeVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WasabeeVault *WasabeeVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _WasabeeVault.Contract.TotalSupply(&_WasabeeVault.CallOpts)
}

// TwapPeriod is a free data retrieval call binding the contract method 0xf6207326.
//
// Solidity: function twapPeriod() view returns(uint32)
func (_WasabeeVault *WasabeeVaultCaller) TwapPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _WasabeeVault.contract.Call(opts, &out, "twapPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TwapPeriod is a free data retrieval call binding the contract method 0xf6207326.
//
// Solidity: function twapPeriod() view returns(uint32)
func (_WasabeeVault *WasabeeVaultSession) TwapPeriod() (uint32, error) {
	return _WasabeeVault.Contract.TwapPeriod(&_WasabeeVault.CallOpts)
}

// TwapPeriod is a free data retrieval call binding the contract method 0xf6207326.
//
// Solidity: function twapPeriod() view returns(uint32)
func (_WasabeeVault *WasabeeVaultCallerSession) TwapPeriod() (uint32, error) {
	return _WasabeeVault.Contract.TwapPeriod(&_WasabeeVault.CallOpts)
}

// AlgebraMintCallback is a paid mutator transaction binding the contract method 0x3dd657c5.
//
// Solidity: function algebraMintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_WasabeeVault *WasabeeVaultTransactor) AlgebraMintCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "algebraMintCallback", amount0, amount1, data)
}

// AlgebraMintCallback is a paid mutator transaction binding the contract method 0x3dd657c5.
//
// Solidity: function algebraMintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_WasabeeVault *WasabeeVaultSession) AlgebraMintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _WasabeeVault.Contract.AlgebraMintCallback(&_WasabeeVault.TransactOpts, amount0, amount1, data)
}

// AlgebraMintCallback is a paid mutator transaction binding the contract method 0x3dd657c5.
//
// Solidity: function algebraMintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) AlgebraMintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _WasabeeVault.Contract.AlgebraMintCallback(&_WasabeeVault.TransactOpts, amount0, amount1, data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_WasabeeVault *WasabeeVaultTransactor) AlgebraSwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "algebraSwapCallback", amount0Delta, amount1Delta, data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_WasabeeVault *WasabeeVaultSession) AlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _WasabeeVault.Contract.AlgebraSwapCallback(&_WasabeeVault.TransactOpts, amount0Delta, amount1Delta, data)
}

// AlgebraSwapCallback is a paid mutator transaction binding the contract method 0x2c8958f6.
//
// Solidity: function algebraSwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) AlgebraSwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _WasabeeVault.Contract.AlgebraSwapCallback(&_WasabeeVault.TransactOpts, amount0Delta, amount1Delta, data)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WasabeeVault *WasabeeVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WasabeeVault *WasabeeVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.Approve(&_WasabeeVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WasabeeVault *WasabeeVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.Approve(&_WasabeeVault.TransactOpts, spender, amount)
}

// CollectFees is a paid mutator transaction binding the contract method 0xc8796572.
//
// Solidity: function collectFees() returns(uint256 fees0, uint256 fees1)
func (_WasabeeVault *WasabeeVaultTransactor) CollectFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "collectFees")
}

// CollectFees is a paid mutator transaction binding the contract method 0xc8796572.
//
// Solidity: function collectFees() returns(uint256 fees0, uint256 fees1)
func (_WasabeeVault *WasabeeVaultSession) CollectFees() (*types.Transaction, error) {
	return _WasabeeVault.Contract.CollectFees(&_WasabeeVault.TransactOpts)
}

// CollectFees is a paid mutator transaction binding the contract method 0xc8796572.
//
// Solidity: function collectFees() returns(uint256 fees0, uint256 fees1)
func (_WasabeeVault *WasabeeVaultTransactorSession) CollectFees() (*types.Transaction, error) {
	return _WasabeeVault.Contract.CollectFees(&_WasabeeVault.TransactOpts)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WasabeeVault *WasabeeVaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WasabeeVault *WasabeeVaultSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.DecreaseAllowance(&_WasabeeVault.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WasabeeVault *WasabeeVaultTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.DecreaseAllowance(&_WasabeeVault.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to) returns(uint256 shares)
func (_WasabeeVault *WasabeeVaultTransactor) Deposit(opts *bind.TransactOpts, deposit0 *big.Int, deposit1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "deposit", deposit0, deposit1, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to) returns(uint256 shares)
func (_WasabeeVault *WasabeeVaultSession) Deposit(deposit0 *big.Int, deposit1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _WasabeeVault.Contract.Deposit(&_WasabeeVault.TransactOpts, deposit0, deposit1, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x8dbdbe6d.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to) returns(uint256 shares)
func (_WasabeeVault *WasabeeVaultTransactorSession) Deposit(deposit0 *big.Int, deposit1 *big.Int, to common.Address) (*types.Transaction, error) {
	return _WasabeeVault.Contract.Deposit(&_WasabeeVault.TransactOpts, deposit0, deposit1, to)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WasabeeVault *WasabeeVaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WasabeeVault *WasabeeVaultSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.IncreaseAllowance(&_WasabeeVault.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WasabeeVault *WasabeeVaultTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.IncreaseAllowance(&_WasabeeVault.TransactOpts, spender, addedValue)
}

// Rebalance is a paid mutator transaction binding the contract method 0xd87346aa.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, int256 swapQuantity) returns()
func (_WasabeeVault *WasabeeVaultTransactor) Rebalance(opts *bind.TransactOpts, _baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, swapQuantity *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "rebalance", _baseLower, _baseUpper, _limitLower, _limitUpper, swapQuantity)
}

// Rebalance is a paid mutator transaction binding the contract method 0xd87346aa.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, int256 swapQuantity) returns()
func (_WasabeeVault *WasabeeVaultSession) Rebalance(_baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, swapQuantity *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.Rebalance(&_WasabeeVault.TransactOpts, _baseLower, _baseUpper, _limitLower, _limitUpper, swapQuantity)
}

// Rebalance is a paid mutator transaction binding the contract method 0xd87346aa.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, int256 swapQuantity) returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) Rebalance(_baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, swapQuantity *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.Rebalance(&_WasabeeVault.TransactOpts, _baseLower, _baseUpper, _limitLower, _limitUpper, swapQuantity)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WasabeeVault *WasabeeVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WasabeeVault *WasabeeVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _WasabeeVault.Contract.RenounceOwnership(&_WasabeeVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WasabeeVault.Contract.RenounceOwnership(&_WasabeeVault.TransactOpts)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x2bbb56d9.
//
// Solidity: function setAffiliate(address _affiliate) returns()
func (_WasabeeVault *WasabeeVaultTransactor) SetAffiliate(opts *bind.TransactOpts, _affiliate common.Address) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "setAffiliate", _affiliate)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x2bbb56d9.
//
// Solidity: function setAffiliate(address _affiliate) returns()
func (_WasabeeVault *WasabeeVaultSession) SetAffiliate(_affiliate common.Address) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetAffiliate(&_WasabeeVault.TransactOpts, _affiliate)
}

// SetAffiliate is a paid mutator transaction binding the contract method 0x2bbb56d9.
//
// Solidity: function setAffiliate(address _affiliate) returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) SetAffiliate(_affiliate common.Address) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetAffiliate(&_WasabeeVault.TransactOpts, _affiliate)
}

// SetAmmFeeRecipient is a paid mutator transaction binding the contract method 0x81de128b.
//
// Solidity: function setAmmFeeRecipient(address _ammFeeRecipient) returns()
func (_WasabeeVault *WasabeeVaultTransactor) SetAmmFeeRecipient(opts *bind.TransactOpts, _ammFeeRecipient common.Address) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "setAmmFeeRecipient", _ammFeeRecipient)
}

// SetAmmFeeRecipient is a paid mutator transaction binding the contract method 0x81de128b.
//
// Solidity: function setAmmFeeRecipient(address _ammFeeRecipient) returns()
func (_WasabeeVault *WasabeeVaultSession) SetAmmFeeRecipient(_ammFeeRecipient common.Address) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetAmmFeeRecipient(&_WasabeeVault.TransactOpts, _ammFeeRecipient)
}

// SetAmmFeeRecipient is a paid mutator transaction binding the contract method 0x81de128b.
//
// Solidity: function setAmmFeeRecipient(address _ammFeeRecipient) returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) SetAmmFeeRecipient(_ammFeeRecipient common.Address) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetAmmFeeRecipient(&_WasabeeVault.TransactOpts, _ammFeeRecipient)
}

// SetAuxTwapPeriod is a paid mutator transaction binding the contract method 0x400f0ceb.
//
// Solidity: function setAuxTwapPeriod(uint32 newAuxTwapPeriod) returns()
func (_WasabeeVault *WasabeeVaultTransactor) SetAuxTwapPeriod(opts *bind.TransactOpts, newAuxTwapPeriod uint32) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "setAuxTwapPeriod", newAuxTwapPeriod)
}

// SetAuxTwapPeriod is a paid mutator transaction binding the contract method 0x400f0ceb.
//
// Solidity: function setAuxTwapPeriod(uint32 newAuxTwapPeriod) returns()
func (_WasabeeVault *WasabeeVaultSession) SetAuxTwapPeriod(newAuxTwapPeriod uint32) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetAuxTwapPeriod(&_WasabeeVault.TransactOpts, newAuxTwapPeriod)
}

// SetAuxTwapPeriod is a paid mutator transaction binding the contract method 0x400f0ceb.
//
// Solidity: function setAuxTwapPeriod(uint32 newAuxTwapPeriod) returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) SetAuxTwapPeriod(newAuxTwapPeriod uint32) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetAuxTwapPeriod(&_WasabeeVault.TransactOpts, newAuxTwapPeriod)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0x3e091ee9.
//
// Solidity: function setDepositMax(uint256 _deposit0Max, uint256 _deposit1Max) returns()
func (_WasabeeVault *WasabeeVaultTransactor) SetDepositMax(opts *bind.TransactOpts, _deposit0Max *big.Int, _deposit1Max *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "setDepositMax", _deposit0Max, _deposit1Max)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0x3e091ee9.
//
// Solidity: function setDepositMax(uint256 _deposit0Max, uint256 _deposit1Max) returns()
func (_WasabeeVault *WasabeeVaultSession) SetDepositMax(_deposit0Max *big.Int, _deposit1Max *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetDepositMax(&_WasabeeVault.TransactOpts, _deposit0Max, _deposit1Max)
}

// SetDepositMax is a paid mutator transaction binding the contract method 0x3e091ee9.
//
// Solidity: function setDepositMax(uint256 _deposit0Max, uint256 _deposit1Max) returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) SetDepositMax(_deposit0Max *big.Int, _deposit1Max *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetDepositMax(&_WasabeeVault.TransactOpts, _deposit0Max, _deposit1Max)
}

// SetHysteresis is a paid mutator transaction binding the contract method 0x5ffc1ff7.
//
// Solidity: function setHysteresis(uint256 _hysteresis) returns()
func (_WasabeeVault *WasabeeVaultTransactor) SetHysteresis(opts *bind.TransactOpts, _hysteresis *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "setHysteresis", _hysteresis)
}

// SetHysteresis is a paid mutator transaction binding the contract method 0x5ffc1ff7.
//
// Solidity: function setHysteresis(uint256 _hysteresis) returns()
func (_WasabeeVault *WasabeeVaultSession) SetHysteresis(_hysteresis *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetHysteresis(&_WasabeeVault.TransactOpts, _hysteresis)
}

// SetHysteresis is a paid mutator transaction binding the contract method 0x5ffc1ff7.
//
// Solidity: function setHysteresis(uint256 _hysteresis) returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) SetHysteresis(_hysteresis *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetHysteresis(&_WasabeeVault.TransactOpts, _hysteresis)
}

// SetTwapPeriod is a paid mutator transaction binding the contract method 0xf9c95d46.
//
// Solidity: function setTwapPeriod(uint32 newTwapPeriod) returns()
func (_WasabeeVault *WasabeeVaultTransactor) SetTwapPeriod(opts *bind.TransactOpts, newTwapPeriod uint32) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "setTwapPeriod", newTwapPeriod)
}

// SetTwapPeriod is a paid mutator transaction binding the contract method 0xf9c95d46.
//
// Solidity: function setTwapPeriod(uint32 newTwapPeriod) returns()
func (_WasabeeVault *WasabeeVaultSession) SetTwapPeriod(newTwapPeriod uint32) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetTwapPeriod(&_WasabeeVault.TransactOpts, newTwapPeriod)
}

// SetTwapPeriod is a paid mutator transaction binding the contract method 0xf9c95d46.
//
// Solidity: function setTwapPeriod(uint32 newTwapPeriod) returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) SetTwapPeriod(newTwapPeriod uint32) (*types.Transaction, error) {
	return _WasabeeVault.Contract.SetTwapPeriod(&_WasabeeVault.TransactOpts, newTwapPeriod)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WasabeeVault *WasabeeVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WasabeeVault *WasabeeVaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.Transfer(&_WasabeeVault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_WasabeeVault *WasabeeVaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.Transfer(&_WasabeeVault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WasabeeVault *WasabeeVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WasabeeVault *WasabeeVaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.TransferFrom(&_WasabeeVault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_WasabeeVault *WasabeeVaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WasabeeVault.Contract.TransferFrom(&_WasabeeVault.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WasabeeVault *WasabeeVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WasabeeVault *WasabeeVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WasabeeVault.Contract.TransferOwnership(&_WasabeeVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WasabeeVault *WasabeeVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WasabeeVault.Contract.TransferOwnership(&_WasabeeVault.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address to) returns(uint256 amount0, uint256 amount1)
func (_WasabeeVault *WasabeeVaultTransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int, to common.Address) (*types.Transaction, error) {
	return _WasabeeVault.contract.Transact(opts, "withdraw", shares, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address to) returns(uint256 amount0, uint256 amount1)
func (_WasabeeVault *WasabeeVaultSession) Withdraw(shares *big.Int, to common.Address) (*types.Transaction, error) {
	return _WasabeeVault.Contract.Withdraw(&_WasabeeVault.TransactOpts, shares, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 shares, address to) returns(uint256 amount0, uint256 amount1)
func (_WasabeeVault *WasabeeVaultTransactorSession) Withdraw(shares *big.Int, to common.Address) (*types.Transaction, error) {
	return _WasabeeVault.Contract.Withdraw(&_WasabeeVault.TransactOpts, shares, to)
}

// WasabeeVaultAffiliateIterator is returned from FilterAffiliate and is used to iterate over the raw logs and unpacked data for Affiliate events raised by the WasabeeVault contract.
type WasabeeVaultAffiliateIterator struct {
	Event *WasabeeVaultAffiliate // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultAffiliateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultAffiliate)
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
		it.Event = new(WasabeeVaultAffiliate)
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
func (it *WasabeeVaultAffiliateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultAffiliateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultAffiliate represents a Affiliate event raised by the WasabeeVault contract.
type WasabeeVaultAffiliate struct {
	Sender    common.Address
	Affiliate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAffiliate is a free log retrieval operation binding the contract event 0x3066ef5dd340e8b2ea28d62f5a8391eb7a82d3ee87532724a1ca4386d34f7523.
//
// Solidity: event Affiliate(address indexed sender, address affiliate)
func (_WasabeeVault *WasabeeVaultFilterer) FilterAffiliate(opts *bind.FilterOpts, sender []common.Address) (*WasabeeVaultAffiliateIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "Affiliate", senderRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultAffiliateIterator{contract: _WasabeeVault.contract, event: "Affiliate", logs: logs, sub: sub}, nil
}

// WatchAffiliate is a free log subscription operation binding the contract event 0x3066ef5dd340e8b2ea28d62f5a8391eb7a82d3ee87532724a1ca4386d34f7523.
//
// Solidity: event Affiliate(address indexed sender, address affiliate)
func (_WasabeeVault *WasabeeVaultFilterer) WatchAffiliate(opts *bind.WatchOpts, sink chan<- *WasabeeVaultAffiliate, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "Affiliate", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultAffiliate)
				if err := _WasabeeVault.contract.UnpackLog(event, "Affiliate", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseAffiliate(log types.Log) (*WasabeeVaultAffiliate, error) {
	event := new(WasabeeVaultAffiliate)
	if err := _WasabeeVault.contract.UnpackLog(event, "Affiliate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultAmmFeeRecipientIterator is returned from FilterAmmFeeRecipient and is used to iterate over the raw logs and unpacked data for AmmFeeRecipient events raised by the WasabeeVault contract.
type WasabeeVaultAmmFeeRecipientIterator struct {
	Event *WasabeeVaultAmmFeeRecipient // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultAmmFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultAmmFeeRecipient)
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
		it.Event = new(WasabeeVaultAmmFeeRecipient)
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
func (it *WasabeeVaultAmmFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultAmmFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultAmmFeeRecipient represents a AmmFeeRecipient event raised by the WasabeeVault contract.
type WasabeeVaultAmmFeeRecipient struct {
	Sender          common.Address
	AmmFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAmmFeeRecipient is a free log retrieval operation binding the contract event 0xbb78b7c13893a913fa8c9ecb9fdaf97597aa412a39c778bf976790555f0942f7.
//
// Solidity: event AmmFeeRecipient(address indexed sender, address ammFeeRecipient)
func (_WasabeeVault *WasabeeVaultFilterer) FilterAmmFeeRecipient(opts *bind.FilterOpts, sender []common.Address) (*WasabeeVaultAmmFeeRecipientIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "AmmFeeRecipient", senderRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultAmmFeeRecipientIterator{contract: _WasabeeVault.contract, event: "AmmFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchAmmFeeRecipient is a free log subscription operation binding the contract event 0xbb78b7c13893a913fa8c9ecb9fdaf97597aa412a39c778bf976790555f0942f7.
//
// Solidity: event AmmFeeRecipient(address indexed sender, address ammFeeRecipient)
func (_WasabeeVault *WasabeeVaultFilterer) WatchAmmFeeRecipient(opts *bind.WatchOpts, sink chan<- *WasabeeVaultAmmFeeRecipient, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "AmmFeeRecipient", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultAmmFeeRecipient)
				if err := _WasabeeVault.contract.UnpackLog(event, "AmmFeeRecipient", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseAmmFeeRecipient(log types.Log) (*WasabeeVaultAmmFeeRecipient, error) {
	event := new(WasabeeVaultAmmFeeRecipient)
	if err := _WasabeeVault.contract.UnpackLog(event, "AmmFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WasabeeVault contract.
type WasabeeVaultApprovalIterator struct {
	Event *WasabeeVaultApproval // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultApproval)
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
		it.Event = new(WasabeeVaultApproval)
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
func (it *WasabeeVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultApproval represents a Approval event raised by the WasabeeVault contract.
type WasabeeVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WasabeeVault *WasabeeVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WasabeeVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultApprovalIterator{contract: _WasabeeVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WasabeeVault *WasabeeVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WasabeeVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultApproval)
				if err := _WasabeeVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseApproval(log types.Log) (*WasabeeVaultApproval, error) {
	event := new(WasabeeVaultApproval)
	if err := _WasabeeVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultCollectFeesIterator is returned from FilterCollectFees and is used to iterate over the raw logs and unpacked data for CollectFees events raised by the WasabeeVault contract.
type WasabeeVaultCollectFeesIterator struct {
	Event *WasabeeVaultCollectFees // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultCollectFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultCollectFees)
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
		it.Event = new(WasabeeVaultCollectFees)
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
func (it *WasabeeVaultCollectFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultCollectFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultCollectFees represents a CollectFees event raised by the WasabeeVault contract.
type WasabeeVaultCollectFees struct {
	Sender     common.Address
	FeeAmount0 *big.Int
	FeeAmount1 *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0xec8208dd791fa8ffdc0d7427f3ba9c0ed06f1bce9a86254e6940c10cc1802fef.
//
// Solidity: event CollectFees(address indexed sender, uint256 feeAmount0, uint256 feeAmount1)
func (_WasabeeVault *WasabeeVaultFilterer) FilterCollectFees(opts *bind.FilterOpts, sender []common.Address) (*WasabeeVaultCollectFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "CollectFees", senderRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultCollectFeesIterator{contract: _WasabeeVault.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0xec8208dd791fa8ffdc0d7427f3ba9c0ed06f1bce9a86254e6940c10cc1802fef.
//
// Solidity: event CollectFees(address indexed sender, uint256 feeAmount0, uint256 feeAmount1)
func (_WasabeeVault *WasabeeVaultFilterer) WatchCollectFees(opts *bind.WatchOpts, sink chan<- *WasabeeVaultCollectFees, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "CollectFees", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultCollectFees)
				if err := _WasabeeVault.contract.UnpackLog(event, "CollectFees", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseCollectFees(log types.Log) (*WasabeeVaultCollectFees, error) {
	event := new(WasabeeVaultCollectFees)
	if err := _WasabeeVault.contract.UnpackLog(event, "CollectFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultDeployICHIVaultIterator is returned from FilterDeployICHIVault and is used to iterate over the raw logs and unpacked data for DeployICHIVault events raised by the WasabeeVault contract.
type WasabeeVaultDeployICHIVaultIterator struct {
	Event *WasabeeVaultDeployICHIVault // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultDeployICHIVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultDeployICHIVault)
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
		it.Event = new(WasabeeVaultDeployICHIVault)
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
func (it *WasabeeVaultDeployICHIVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultDeployICHIVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultDeployICHIVault represents a DeployICHIVault event raised by the WasabeeVault contract.
type WasabeeVaultDeployICHIVault struct {
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
func (_WasabeeVault *WasabeeVaultFilterer) FilterDeployICHIVault(opts *bind.FilterOpts, sender []common.Address, pool []common.Address) (*WasabeeVaultDeployICHIVaultIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "DeployICHIVault", senderRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultDeployICHIVaultIterator{contract: _WasabeeVault.contract, event: "DeployICHIVault", logs: logs, sub: sub}, nil
}

// WatchDeployICHIVault is a free log subscription operation binding the contract event 0x3e708ccf7d0e6de8558e020ea36189511cb3435bbfec54e721a48ee4df0d4f8c.
//
// Solidity: event DeployICHIVault(address indexed sender, address indexed pool, bool allowToken0, bool allowToken1, address owner, uint256 twapPeriod)
func (_WasabeeVault *WasabeeVaultFilterer) WatchDeployICHIVault(opts *bind.WatchOpts, sink chan<- *WasabeeVaultDeployICHIVault, sender []common.Address, pool []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "DeployICHIVault", senderRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultDeployICHIVault)
				if err := _WasabeeVault.contract.UnpackLog(event, "DeployICHIVault", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseDeployICHIVault(log types.Log) (*WasabeeVaultDeployICHIVault, error) {
	event := new(WasabeeVaultDeployICHIVault)
	if err := _WasabeeVault.contract.UnpackLog(event, "DeployICHIVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WasabeeVault contract.
type WasabeeVaultDepositIterator struct {
	Event *WasabeeVaultDeposit // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultDeposit)
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
		it.Event = new(WasabeeVaultDeposit)
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
func (it *WasabeeVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultDeposit represents a Deposit event raised by the WasabeeVault contract.
type WasabeeVaultDeposit struct {
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
func (_WasabeeVault *WasabeeVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*WasabeeVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultDepositIterator{contract: _WasabeeVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_WasabeeVault *WasabeeVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WasabeeVaultDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultDeposit)
				if err := _WasabeeVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseDeposit(log types.Log) (*WasabeeVaultDeposit, error) {
	event := new(WasabeeVaultDeposit)
	if err := _WasabeeVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultDepositMaxIterator is returned from FilterDepositMax and is used to iterate over the raw logs and unpacked data for DepositMax events raised by the WasabeeVault contract.
type WasabeeVaultDepositMaxIterator struct {
	Event *WasabeeVaultDepositMax // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultDepositMaxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultDepositMax)
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
		it.Event = new(WasabeeVaultDepositMax)
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
func (it *WasabeeVaultDepositMaxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultDepositMaxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultDepositMax represents a DepositMax event raised by the WasabeeVault contract.
type WasabeeVaultDepositMax struct {
	Sender      common.Address
	Deposit0Max *big.Int
	Deposit1Max *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositMax is a free log retrieval operation binding the contract event 0xafd3b05a4086b378b6f291200a528d8aed8c5e0317af77436b001f1bec28821a.
//
// Solidity: event DepositMax(address indexed sender, uint256 deposit0Max, uint256 deposit1Max)
func (_WasabeeVault *WasabeeVaultFilterer) FilterDepositMax(opts *bind.FilterOpts, sender []common.Address) (*WasabeeVaultDepositMaxIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "DepositMax", senderRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultDepositMaxIterator{contract: _WasabeeVault.contract, event: "DepositMax", logs: logs, sub: sub}, nil
}

// WatchDepositMax is a free log subscription operation binding the contract event 0xafd3b05a4086b378b6f291200a528d8aed8c5e0317af77436b001f1bec28821a.
//
// Solidity: event DepositMax(address indexed sender, uint256 deposit0Max, uint256 deposit1Max)
func (_WasabeeVault *WasabeeVaultFilterer) WatchDepositMax(opts *bind.WatchOpts, sink chan<- *WasabeeVaultDepositMax, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "DepositMax", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultDepositMax)
				if err := _WasabeeVault.contract.UnpackLog(event, "DepositMax", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseDepositMax(log types.Log) (*WasabeeVaultDepositMax, error) {
	event := new(WasabeeVaultDepositMax)
	if err := _WasabeeVault.contract.UnpackLog(event, "DepositMax", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultHysteresisIterator is returned from FilterHysteresis and is used to iterate over the raw logs and unpacked data for Hysteresis events raised by the WasabeeVault contract.
type WasabeeVaultHysteresisIterator struct {
	Event *WasabeeVaultHysteresis // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultHysteresisIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultHysteresis)
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
		it.Event = new(WasabeeVaultHysteresis)
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
func (it *WasabeeVaultHysteresisIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultHysteresisIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultHysteresis represents a Hysteresis event raised by the WasabeeVault contract.
type WasabeeVaultHysteresis struct {
	Sender     common.Address
	Hysteresis *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterHysteresis is a free log retrieval operation binding the contract event 0x529698f34660760dcb172def5c99d62e1b5b74b444df322e8f7da31f2bd0a86b.
//
// Solidity: event Hysteresis(address indexed sender, uint256 hysteresis)
func (_WasabeeVault *WasabeeVaultFilterer) FilterHysteresis(opts *bind.FilterOpts, sender []common.Address) (*WasabeeVaultHysteresisIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "Hysteresis", senderRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultHysteresisIterator{contract: _WasabeeVault.contract, event: "Hysteresis", logs: logs, sub: sub}, nil
}

// WatchHysteresis is a free log subscription operation binding the contract event 0x529698f34660760dcb172def5c99d62e1b5b74b444df322e8f7da31f2bd0a86b.
//
// Solidity: event Hysteresis(address indexed sender, uint256 hysteresis)
func (_WasabeeVault *WasabeeVaultFilterer) WatchHysteresis(opts *bind.WatchOpts, sink chan<- *WasabeeVaultHysteresis, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "Hysteresis", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultHysteresis)
				if err := _WasabeeVault.contract.UnpackLog(event, "Hysteresis", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseHysteresis(log types.Log) (*WasabeeVaultHysteresis, error) {
	event := new(WasabeeVaultHysteresis)
	if err := _WasabeeVault.contract.UnpackLog(event, "Hysteresis", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WasabeeVault contract.
type WasabeeVaultOwnershipTransferredIterator struct {
	Event *WasabeeVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultOwnershipTransferred)
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
		it.Event = new(WasabeeVaultOwnershipTransferred)
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
func (it *WasabeeVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultOwnershipTransferred represents a OwnershipTransferred event raised by the WasabeeVault contract.
type WasabeeVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WasabeeVault *WasabeeVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WasabeeVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultOwnershipTransferredIterator{contract: _WasabeeVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WasabeeVault *WasabeeVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WasabeeVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultOwnershipTransferred)
				if err := _WasabeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseOwnershipTransferred(log types.Log) (*WasabeeVaultOwnershipTransferred, error) {
	event := new(WasabeeVaultOwnershipTransferred)
	if err := _WasabeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the WasabeeVault contract.
type WasabeeVaultRebalanceIterator struct {
	Event *WasabeeVaultRebalance // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultRebalance)
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
		it.Event = new(WasabeeVaultRebalance)
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
func (it *WasabeeVaultRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultRebalance represents a Rebalance event raised by the WasabeeVault contract.
type WasabeeVaultRebalance struct {
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
func (_WasabeeVault *WasabeeVaultFilterer) FilterRebalance(opts *bind.FilterOpts) (*WasabeeVaultRebalanceIterator, error) {

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultRebalanceIterator{contract: _WasabeeVault.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0xbc4c20ad04f161d631d9ce94d27659391196415aa3c42f6a71c62e905ece782d.
//
// Solidity: event Rebalance(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 feeAmount0, uint256 feeAmount1, uint256 totalSupply)
func (_WasabeeVault *WasabeeVaultFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *WasabeeVaultRebalance) (event.Subscription, error) {

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultRebalance)
				if err := _WasabeeVault.contract.UnpackLog(event, "Rebalance", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseRebalance(log types.Log) (*WasabeeVaultRebalance, error) {
	event := new(WasabeeVaultRebalance)
	if err := _WasabeeVault.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultSetAuxTwapPeriodIterator is returned from FilterSetAuxTwapPeriod and is used to iterate over the raw logs and unpacked data for SetAuxTwapPeriod events raised by the WasabeeVault contract.
type WasabeeVaultSetAuxTwapPeriodIterator struct {
	Event *WasabeeVaultSetAuxTwapPeriod // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultSetAuxTwapPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultSetAuxTwapPeriod)
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
		it.Event = new(WasabeeVaultSetAuxTwapPeriod)
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
func (it *WasabeeVaultSetAuxTwapPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultSetAuxTwapPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultSetAuxTwapPeriod represents a SetAuxTwapPeriod event raised by the WasabeeVault contract.
type WasabeeVaultSetAuxTwapPeriod struct {
	Sender           common.Address
	NewAuxTwapPeriod uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetAuxTwapPeriod is a free log retrieval operation binding the contract event 0x39da19f5960a3f182ced1ff1853b7be54f37150799b3003a40bf4e0d4c740c85.
//
// Solidity: event SetAuxTwapPeriod(address sender, uint32 newAuxTwapPeriod)
func (_WasabeeVault *WasabeeVaultFilterer) FilterSetAuxTwapPeriod(opts *bind.FilterOpts) (*WasabeeVaultSetAuxTwapPeriodIterator, error) {

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "SetAuxTwapPeriod")
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultSetAuxTwapPeriodIterator{contract: _WasabeeVault.contract, event: "SetAuxTwapPeriod", logs: logs, sub: sub}, nil
}

// WatchSetAuxTwapPeriod is a free log subscription operation binding the contract event 0x39da19f5960a3f182ced1ff1853b7be54f37150799b3003a40bf4e0d4c740c85.
//
// Solidity: event SetAuxTwapPeriod(address sender, uint32 newAuxTwapPeriod)
func (_WasabeeVault *WasabeeVaultFilterer) WatchSetAuxTwapPeriod(opts *bind.WatchOpts, sink chan<- *WasabeeVaultSetAuxTwapPeriod) (event.Subscription, error) {

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "SetAuxTwapPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultSetAuxTwapPeriod)
				if err := _WasabeeVault.contract.UnpackLog(event, "SetAuxTwapPeriod", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseSetAuxTwapPeriod(log types.Log) (*WasabeeVaultSetAuxTwapPeriod, error) {
	event := new(WasabeeVaultSetAuxTwapPeriod)
	if err := _WasabeeVault.contract.UnpackLog(event, "SetAuxTwapPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultSetTwapPeriodIterator is returned from FilterSetTwapPeriod and is used to iterate over the raw logs and unpacked data for SetTwapPeriod events raised by the WasabeeVault contract.
type WasabeeVaultSetTwapPeriodIterator struct {
	Event *WasabeeVaultSetTwapPeriod // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultSetTwapPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultSetTwapPeriod)
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
		it.Event = new(WasabeeVaultSetTwapPeriod)
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
func (it *WasabeeVaultSetTwapPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultSetTwapPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultSetTwapPeriod represents a SetTwapPeriod event raised by the WasabeeVault contract.
type WasabeeVaultSetTwapPeriod struct {
	Sender        common.Address
	NewTwapPeriod uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTwapPeriod is a free log retrieval operation binding the contract event 0xe4c60f4984caeb7f45b0cfe6d4233c115601ab11d141bc2cbf68b48346cdef38.
//
// Solidity: event SetTwapPeriod(address sender, uint32 newTwapPeriod)
func (_WasabeeVault *WasabeeVaultFilterer) FilterSetTwapPeriod(opts *bind.FilterOpts) (*WasabeeVaultSetTwapPeriodIterator, error) {

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "SetTwapPeriod")
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultSetTwapPeriodIterator{contract: _WasabeeVault.contract, event: "SetTwapPeriod", logs: logs, sub: sub}, nil
}

// WatchSetTwapPeriod is a free log subscription operation binding the contract event 0xe4c60f4984caeb7f45b0cfe6d4233c115601ab11d141bc2cbf68b48346cdef38.
//
// Solidity: event SetTwapPeriod(address sender, uint32 newTwapPeriod)
func (_WasabeeVault *WasabeeVaultFilterer) WatchSetTwapPeriod(opts *bind.WatchOpts, sink chan<- *WasabeeVaultSetTwapPeriod) (event.Subscription, error) {

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "SetTwapPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultSetTwapPeriod)
				if err := _WasabeeVault.contract.UnpackLog(event, "SetTwapPeriod", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseSetTwapPeriod(log types.Log) (*WasabeeVaultSetTwapPeriod, error) {
	event := new(WasabeeVaultSetTwapPeriod)
	if err := _WasabeeVault.contract.UnpackLog(event, "SetTwapPeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WasabeeVault contract.
type WasabeeVaultTransferIterator struct {
	Event *WasabeeVaultTransfer // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultTransfer)
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
		it.Event = new(WasabeeVaultTransfer)
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
func (it *WasabeeVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultTransfer represents a Transfer event raised by the WasabeeVault contract.
type WasabeeVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WasabeeVault *WasabeeVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WasabeeVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultTransferIterator{contract: _WasabeeVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WasabeeVault *WasabeeVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WasabeeVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultTransfer)
				if err := _WasabeeVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseTransfer(log types.Log) (*WasabeeVaultTransfer, error) {
	event := new(WasabeeVaultTransfer)
	if err := _WasabeeVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WasabeeVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the WasabeeVault contract.
type WasabeeVaultWithdrawIterator struct {
	Event *WasabeeVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *WasabeeVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WasabeeVaultWithdraw)
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
		it.Event = new(WasabeeVaultWithdraw)
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
func (it *WasabeeVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WasabeeVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WasabeeVaultWithdraw represents a Withdraw event raised by the WasabeeVault contract.
type WasabeeVaultWithdraw struct {
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
func (_WasabeeVault *WasabeeVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*WasabeeVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WasabeeVault.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WasabeeVaultWithdrawIterator{contract: _WasabeeVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_WasabeeVault *WasabeeVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *WasabeeVaultWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WasabeeVault.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WasabeeVaultWithdraw)
				if err := _WasabeeVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_WasabeeVault *WasabeeVaultFilterer) ParseWithdraw(log types.Log) (*WasabeeVaultWithdraw, error) {
	event := new(WasabeeVaultWithdraw)
	if err := _WasabeeVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
