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

// EtherfiAccountantMetaData contains all meta data concerning the EtherfiAccountant contract.
var EtherfiAccountantMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"startingExchangeRate\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"allowedExchangeRateChangeUpper\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"allowedExchangeRateChangeLower\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"minimumUpdateDelayInSeconds\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"platformFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"performanceFee\",\"type\":\"uint16\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountantWithFixedRate__HighWaterMarkCannotChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithFixedRate__OnlyCallableByYieldDistributor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithFixedRate__StartingExchangeRateCannotBeGreaterThanFixed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithFixedRate__UnsafeUint96Cast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithFixedRate__ZeroYieldOwed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithRateProviders__ExchangeRateAboveHighwaterMark\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithRateProviders__LowerBoundTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithRateProviders__OnlyCallableByBoringVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithRateProviders__Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithRateProviders__PerformanceFeeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithRateProviders__PlatformFeeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithRateProviders__UpdateDelayTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithRateProviders__UpperBoundTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountantWithRateProviders__ZeroFeesOwed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractAuthority\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"oldDelay\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"newDelay\",\"type\":\"uint24\"}],\"name\":\"DelayInSecondsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"oldRate\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"newRate\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"currentTime\",\"type\":\"uint64\"}],\"name\":\"ExchangeRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"HighwaterMarkReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldBound\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newBound\",\"type\":\"uint16\"}],\"name\":\"LowerBoundUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPayout\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPayout\",\"type\":\"address\"}],\"name\":\"PayoutAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldFee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"}],\"name\":\"PerformanceFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldFee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"}],\"name\":\"PlatformFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPegged\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rateProvider\",\"type\":\"address\"}],\"name\":\"RateProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldBound\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newBound\",\"type\":\"uint16\"}],\"name\":\"UpperBoundUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"yieldAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"YieldClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"yieldDistributor\",\"type\":\"address\"}],\"name\":\"YieldDistributorUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accountantState\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"highwaterMark\",\"type\":\"uint96\"},{\"internalType\":\"uint128\",\"name\":\"feesOwedInBase\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalSharesLastUpdate\",\"type\":\"uint128\"},{\"internalType\":\"uint96\",\"name\":\"exchangeRate\",\"type\":\"uint96\"},{\"internalType\":\"uint16\",\"name\":\"allowedExchangeRateChangeUpper\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"allowedExchangeRateChangeLower\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"minimumUpdateDelayInSeconds\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"platformFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"performanceFee\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractAuthority\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"base\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"feeAsset\",\"type\":\"address\"}],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"yieldAsset\",\"type\":\"address\"}],\"name\":\"claimYield\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fixedRateAccountantState\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"yieldEarnedInBase\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"yieldDistributor\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"getRateInQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rateInQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"quote\",\"type\":\"address\"}],\"name\":\"getRateInQuoteSafe\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rateInQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateSafe\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"newExchangeRate\",\"type\":\"uint96\"}],\"name\":\"previewUpdateExchangeRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"updateWillPause\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"newFeesOwedInBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeesOwedInBase\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rateProviderData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isPeggedToBase\",\"type\":\"bool\"},{\"internalType\":\"contractIRateProvider\",\"name\":\"rateProvider\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetHighwaterMark\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAuthority\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isPeggedToBase\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"rateProvider\",\"type\":\"address\"}],\"name\":\"setRateProviderData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"yieldDistributor\",\"type\":\"address\"}],\"name\":\"setYieldDistributor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"minimumUpdateDelayInSeconds\",\"type\":\"uint24\"}],\"name\":\"updateDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"newExchangeRate\",\"type\":\"uint96\"}],\"name\":\"updateExchangeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"allowedExchangeRateChangeLower\",\"type\":\"uint16\"}],\"name\":\"updateLower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payoutAddress\",\"type\":\"address\"}],\"name\":\"updatePayoutAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"performanceFee\",\"type\":\"uint16\"}],\"name\":\"updatePerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"platformFee\",\"type\":\"uint16\"}],\"name\":\"updatePlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"allowedExchangeRateChangeUpper\",\"type\":\"uint16\"}],\"name\":\"updateUpper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractBoringVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EtherfiAccountantABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherfiAccountantMetaData.ABI instead.
var EtherfiAccountantABI = EtherfiAccountantMetaData.ABI

// EtherfiAccountant is an auto generated Go binding around an Ethereum contract.
type EtherfiAccountant struct {
	EtherfiAccountantCaller     // Read-only binding to the contract
	EtherfiAccountantTransactor // Write-only binding to the contract
	EtherfiAccountantFilterer   // Log filterer for contract events
}

// EtherfiAccountantCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherfiAccountantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAccountantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherfiAccountantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAccountantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherfiAccountantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiAccountantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherfiAccountantSession struct {
	Contract     *EtherfiAccountant // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EtherfiAccountantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherfiAccountantCallerSession struct {
	Contract *EtherfiAccountantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EtherfiAccountantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherfiAccountantTransactorSession struct {
	Contract     *EtherfiAccountantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EtherfiAccountantRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherfiAccountantRaw struct {
	Contract *EtherfiAccountant // Generic contract binding to access the raw methods on
}

// EtherfiAccountantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherfiAccountantCallerRaw struct {
	Contract *EtherfiAccountantCaller // Generic read-only contract binding to access the raw methods on
}

// EtherfiAccountantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherfiAccountantTransactorRaw struct {
	Contract *EtherfiAccountantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherfiAccountant creates a new instance of EtherfiAccountant, bound to a specific deployed contract.
func NewEtherfiAccountant(address common.Address, backend bind.ContractBackend) (*EtherfiAccountant, error) {
	contract, err := bindEtherfiAccountant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountant{EtherfiAccountantCaller: EtherfiAccountantCaller{contract: contract}, EtherfiAccountantTransactor: EtherfiAccountantTransactor{contract: contract}, EtherfiAccountantFilterer: EtherfiAccountantFilterer{contract: contract}}, nil
}

// NewEtherfiAccountantCaller creates a new read-only instance of EtherfiAccountant, bound to a specific deployed contract.
func NewEtherfiAccountantCaller(address common.Address, caller bind.ContractCaller) (*EtherfiAccountantCaller, error) {
	contract, err := bindEtherfiAccountant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantCaller{contract: contract}, nil
}

// NewEtherfiAccountantTransactor creates a new write-only instance of EtherfiAccountant, bound to a specific deployed contract.
func NewEtherfiAccountantTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherfiAccountantTransactor, error) {
	contract, err := bindEtherfiAccountant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantTransactor{contract: contract}, nil
}

// NewEtherfiAccountantFilterer creates a new log filterer instance of EtherfiAccountant, bound to a specific deployed contract.
func NewEtherfiAccountantFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherfiAccountantFilterer, error) {
	contract, err := bindEtherfiAccountant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantFilterer{contract: contract}, nil
}

// bindEtherfiAccountant binds a generic wrapper to an already deployed contract.
func bindEtherfiAccountant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherfiAccountantMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiAccountant *EtherfiAccountantRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiAccountant.Contract.EtherfiAccountantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiAccountant *EtherfiAccountantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.EtherfiAccountantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiAccountant *EtherfiAccountantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.EtherfiAccountantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiAccountant *EtherfiAccountantCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiAccountant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiAccountant *EtherfiAccountantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiAccountant *EtherfiAccountantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.contract.Transact(opts, method, params...)
}

// AccountantState is a free data retrieval call binding the contract method 0x433255de.
//
// Solidity: function accountantState() view returns(address payoutAddress, uint96 highwaterMark, uint128 feesOwedInBase, uint128 totalSharesLastUpdate, uint96 exchangeRate, uint16 allowedExchangeRateChangeUpper, uint16 allowedExchangeRateChangeLower, uint64 lastUpdateTimestamp, bool isPaused, uint24 minimumUpdateDelayInSeconds, uint16 platformFee, uint16 performanceFee)
func (_EtherfiAccountant *EtherfiAccountantCaller) AccountantState(opts *bind.CallOpts) (struct {
	PayoutAddress                  common.Address
	HighwaterMark                  *big.Int
	FeesOwedInBase                 *big.Int
	TotalSharesLastUpdate          *big.Int
	ExchangeRate                   *big.Int
	AllowedExchangeRateChangeUpper uint16
	AllowedExchangeRateChangeLower uint16
	LastUpdateTimestamp            uint64
	IsPaused                       bool
	MinimumUpdateDelayInSeconds    *big.Int
	PlatformFee                    uint16
	PerformanceFee                 uint16
}, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "accountantState")

	outstruct := new(struct {
		PayoutAddress                  common.Address
		HighwaterMark                  *big.Int
		FeesOwedInBase                 *big.Int
		TotalSharesLastUpdate          *big.Int
		ExchangeRate                   *big.Int
		AllowedExchangeRateChangeUpper uint16
		AllowedExchangeRateChangeLower uint16
		LastUpdateTimestamp            uint64
		IsPaused                       bool
		MinimumUpdateDelayInSeconds    *big.Int
		PlatformFee                    uint16
		PerformanceFee                 uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PayoutAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.HighwaterMark = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FeesOwedInBase = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalSharesLastUpdate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExchangeRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AllowedExchangeRateChangeUpper = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.AllowedExchangeRateChangeLower = *abi.ConvertType(out[6], new(uint16)).(*uint16)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.IsPaused = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.MinimumUpdateDelayInSeconds = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.PlatformFee = *abi.ConvertType(out[10], new(uint16)).(*uint16)
	outstruct.PerformanceFee = *abi.ConvertType(out[11], new(uint16)).(*uint16)

	return *outstruct, err

}

// AccountantState is a free data retrieval call binding the contract method 0x433255de.
//
// Solidity: function accountantState() view returns(address payoutAddress, uint96 highwaterMark, uint128 feesOwedInBase, uint128 totalSharesLastUpdate, uint96 exchangeRate, uint16 allowedExchangeRateChangeUpper, uint16 allowedExchangeRateChangeLower, uint64 lastUpdateTimestamp, bool isPaused, uint24 minimumUpdateDelayInSeconds, uint16 platformFee, uint16 performanceFee)
func (_EtherfiAccountant *EtherfiAccountantSession) AccountantState() (struct {
	PayoutAddress                  common.Address
	HighwaterMark                  *big.Int
	FeesOwedInBase                 *big.Int
	TotalSharesLastUpdate          *big.Int
	ExchangeRate                   *big.Int
	AllowedExchangeRateChangeUpper uint16
	AllowedExchangeRateChangeLower uint16
	LastUpdateTimestamp            uint64
	IsPaused                       bool
	MinimumUpdateDelayInSeconds    *big.Int
	PlatformFee                    uint16
	PerformanceFee                 uint16
}, error) {
	return _EtherfiAccountant.Contract.AccountantState(&_EtherfiAccountant.CallOpts)
}

// AccountantState is a free data retrieval call binding the contract method 0x433255de.
//
// Solidity: function accountantState() view returns(address payoutAddress, uint96 highwaterMark, uint128 feesOwedInBase, uint128 totalSharesLastUpdate, uint96 exchangeRate, uint16 allowedExchangeRateChangeUpper, uint16 allowedExchangeRateChangeLower, uint64 lastUpdateTimestamp, bool isPaused, uint24 minimumUpdateDelayInSeconds, uint16 platformFee, uint16 performanceFee)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) AccountantState() (struct {
	PayoutAddress                  common.Address
	HighwaterMark                  *big.Int
	FeesOwedInBase                 *big.Int
	TotalSharesLastUpdate          *big.Int
	ExchangeRate                   *big.Int
	AllowedExchangeRateChangeUpper uint16
	AllowedExchangeRateChangeLower uint16
	LastUpdateTimestamp            uint64
	IsPaused                       bool
	MinimumUpdateDelayInSeconds    *big.Int
	PlatformFee                    uint16
	PerformanceFee                 uint16
}, error) {
	return _EtherfiAccountant.Contract.AccountantState(&_EtherfiAccountant.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantSession) Authority() (common.Address, error) {
	return _EtherfiAccountant.Contract.Authority(&_EtherfiAccountant.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) Authority() (common.Address, error) {
	return _EtherfiAccountant.Contract.Authority(&_EtherfiAccountant.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantCaller) Base(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "base")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantSession) Base() (common.Address, error) {
	return _EtherfiAccountant.Contract.Base(&_EtherfiAccountant.CallOpts)
}

// Base is a free data retrieval call binding the contract method 0x5001f3b5.
//
// Solidity: function base() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) Base() (common.Address, error) {
	return _EtherfiAccountant.Contract.Base(&_EtherfiAccountant.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EtherfiAccountant *EtherfiAccountantCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EtherfiAccountant *EtherfiAccountantSession) Decimals() (uint8, error) {
	return _EtherfiAccountant.Contract.Decimals(&_EtherfiAccountant.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) Decimals() (uint8, error) {
	return _EtherfiAccountant.Contract.Decimals(&_EtherfiAccountant.CallOpts)
}

// FixedRateAccountantState is a free data retrieval call binding the contract method 0x0a4f02d7.
//
// Solidity: function fixedRateAccountantState() view returns(uint96 yieldEarnedInBase, address yieldDistributor)
func (_EtherfiAccountant *EtherfiAccountantCaller) FixedRateAccountantState(opts *bind.CallOpts) (struct {
	YieldEarnedInBase *big.Int
	YieldDistributor  common.Address
}, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "fixedRateAccountantState")

	outstruct := new(struct {
		YieldEarnedInBase *big.Int
		YieldDistributor  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.YieldEarnedInBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.YieldDistributor = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// FixedRateAccountantState is a free data retrieval call binding the contract method 0x0a4f02d7.
//
// Solidity: function fixedRateAccountantState() view returns(uint96 yieldEarnedInBase, address yieldDistributor)
func (_EtherfiAccountant *EtherfiAccountantSession) FixedRateAccountantState() (struct {
	YieldEarnedInBase *big.Int
	YieldDistributor  common.Address
}, error) {
	return _EtherfiAccountant.Contract.FixedRateAccountantState(&_EtherfiAccountant.CallOpts)
}

// FixedRateAccountantState is a free data retrieval call binding the contract method 0x0a4f02d7.
//
// Solidity: function fixedRateAccountantState() view returns(uint96 yieldEarnedInBase, address yieldDistributor)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) FixedRateAccountantState() (struct {
	YieldEarnedInBase *big.Int
	YieldDistributor  common.Address
}, error) {
	return _EtherfiAccountant.Contract.FixedRateAccountantState(&_EtherfiAccountant.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256 rate)
func (_EtherfiAccountant *EtherfiAccountantCaller) GetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256 rate)
func (_EtherfiAccountant *EtherfiAccountantSession) GetRate() (*big.Int, error) {
	return _EtherfiAccountant.Contract.GetRate(&_EtherfiAccountant.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256 rate)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) GetRate() (*big.Int, error) {
	return _EtherfiAccountant.Contract.GetRate(&_EtherfiAccountant.CallOpts)
}

// GetRateInQuote is a free data retrieval call binding the contract method 0x1dcbb110.
//
// Solidity: function getRateInQuote(address quote) view returns(uint256 rateInQuote)
func (_EtherfiAccountant *EtherfiAccountantCaller) GetRateInQuote(opts *bind.CallOpts, quote common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "getRateInQuote", quote)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateInQuote is a free data retrieval call binding the contract method 0x1dcbb110.
//
// Solidity: function getRateInQuote(address quote) view returns(uint256 rateInQuote)
func (_EtherfiAccountant *EtherfiAccountantSession) GetRateInQuote(quote common.Address) (*big.Int, error) {
	return _EtherfiAccountant.Contract.GetRateInQuote(&_EtherfiAccountant.CallOpts, quote)
}

// GetRateInQuote is a free data retrieval call binding the contract method 0x1dcbb110.
//
// Solidity: function getRateInQuote(address quote) view returns(uint256 rateInQuote)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) GetRateInQuote(quote common.Address) (*big.Int, error) {
	return _EtherfiAccountant.Contract.GetRateInQuote(&_EtherfiAccountant.CallOpts, quote)
}

// GetRateInQuoteSafe is a free data retrieval call binding the contract method 0x820973da.
//
// Solidity: function getRateInQuoteSafe(address quote) view returns(uint256 rateInQuote)
func (_EtherfiAccountant *EtherfiAccountantCaller) GetRateInQuoteSafe(opts *bind.CallOpts, quote common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "getRateInQuoteSafe", quote)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateInQuoteSafe is a free data retrieval call binding the contract method 0x820973da.
//
// Solidity: function getRateInQuoteSafe(address quote) view returns(uint256 rateInQuote)
func (_EtherfiAccountant *EtherfiAccountantSession) GetRateInQuoteSafe(quote common.Address) (*big.Int, error) {
	return _EtherfiAccountant.Contract.GetRateInQuoteSafe(&_EtherfiAccountant.CallOpts, quote)
}

// GetRateInQuoteSafe is a free data retrieval call binding the contract method 0x820973da.
//
// Solidity: function getRateInQuoteSafe(address quote) view returns(uint256 rateInQuote)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) GetRateInQuoteSafe(quote common.Address) (*big.Int, error) {
	return _EtherfiAccountant.Contract.GetRateInQuoteSafe(&_EtherfiAccountant.CallOpts, quote)
}

// GetRateSafe is a free data retrieval call binding the contract method 0x282a8700.
//
// Solidity: function getRateSafe() view returns(uint256 rate)
func (_EtherfiAccountant *EtherfiAccountantCaller) GetRateSafe(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "getRateSafe")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateSafe is a free data retrieval call binding the contract method 0x282a8700.
//
// Solidity: function getRateSafe() view returns(uint256 rate)
func (_EtherfiAccountant *EtherfiAccountantSession) GetRateSafe() (*big.Int, error) {
	return _EtherfiAccountant.Contract.GetRateSafe(&_EtherfiAccountant.CallOpts)
}

// GetRateSafe is a free data retrieval call binding the contract method 0x282a8700.
//
// Solidity: function getRateSafe() view returns(uint256 rate)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) GetRateSafe() (*big.Int, error) {
	return _EtherfiAccountant.Contract.GetRateSafe(&_EtherfiAccountant.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantSession) Owner() (common.Address, error) {
	return _EtherfiAccountant.Contract.Owner(&_EtherfiAccountant.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) Owner() (common.Address, error) {
	return _EtherfiAccountant.Contract.Owner(&_EtherfiAccountant.CallOpts)
}

// PreviewUpdateExchangeRate is a free data retrieval call binding the contract method 0x6183fb95.
//
// Solidity: function previewUpdateExchangeRate(uint96 newExchangeRate) view returns(bool updateWillPause, uint256 newFeesOwedInBase, uint256 totalFeesOwedInBase)
func (_EtherfiAccountant *EtherfiAccountantCaller) PreviewUpdateExchangeRate(opts *bind.CallOpts, newExchangeRate *big.Int) (struct {
	UpdateWillPause     bool
	NewFeesOwedInBase   *big.Int
	TotalFeesOwedInBase *big.Int
}, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "previewUpdateExchangeRate", newExchangeRate)

	outstruct := new(struct {
		UpdateWillPause     bool
		NewFeesOwedInBase   *big.Int
		TotalFeesOwedInBase *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UpdateWillPause = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.NewFeesOwedInBase = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalFeesOwedInBase = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PreviewUpdateExchangeRate is a free data retrieval call binding the contract method 0x6183fb95.
//
// Solidity: function previewUpdateExchangeRate(uint96 newExchangeRate) view returns(bool updateWillPause, uint256 newFeesOwedInBase, uint256 totalFeesOwedInBase)
func (_EtherfiAccountant *EtherfiAccountantSession) PreviewUpdateExchangeRate(newExchangeRate *big.Int) (struct {
	UpdateWillPause     bool
	NewFeesOwedInBase   *big.Int
	TotalFeesOwedInBase *big.Int
}, error) {
	return _EtherfiAccountant.Contract.PreviewUpdateExchangeRate(&_EtherfiAccountant.CallOpts, newExchangeRate)
}

// PreviewUpdateExchangeRate is a free data retrieval call binding the contract method 0x6183fb95.
//
// Solidity: function previewUpdateExchangeRate(uint96 newExchangeRate) view returns(bool updateWillPause, uint256 newFeesOwedInBase, uint256 totalFeesOwedInBase)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) PreviewUpdateExchangeRate(newExchangeRate *big.Int) (struct {
	UpdateWillPause     bool
	NewFeesOwedInBase   *big.Int
	TotalFeesOwedInBase *big.Int
}, error) {
	return _EtherfiAccountant.Contract.PreviewUpdateExchangeRate(&_EtherfiAccountant.CallOpts, newExchangeRate)
}

// RateProviderData is a free data retrieval call binding the contract method 0x12e2d8f3.
//
// Solidity: function rateProviderData(address ) view returns(bool isPeggedToBase, address rateProvider)
func (_EtherfiAccountant *EtherfiAccountantCaller) RateProviderData(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsPeggedToBase bool
	RateProvider   common.Address
}, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "rateProviderData", arg0)

	outstruct := new(struct {
		IsPeggedToBase bool
		RateProvider   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsPeggedToBase = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RateProvider = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RateProviderData is a free data retrieval call binding the contract method 0x12e2d8f3.
//
// Solidity: function rateProviderData(address ) view returns(bool isPeggedToBase, address rateProvider)
func (_EtherfiAccountant *EtherfiAccountantSession) RateProviderData(arg0 common.Address) (struct {
	IsPeggedToBase bool
	RateProvider   common.Address
}, error) {
	return _EtherfiAccountant.Contract.RateProviderData(&_EtherfiAccountant.CallOpts, arg0)
}

// RateProviderData is a free data retrieval call binding the contract method 0x12e2d8f3.
//
// Solidity: function rateProviderData(address ) view returns(bool isPeggedToBase, address rateProvider)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) RateProviderData(arg0 common.Address) (struct {
	IsPeggedToBase bool
	RateProvider   common.Address
}, error) {
	return _EtherfiAccountant.Contract.RateProviderData(&_EtherfiAccountant.CallOpts, arg0)
}

// ResetHighwaterMark is a free data retrieval call binding the contract method 0xe059ac07.
//
// Solidity: function resetHighwaterMark() view returns()
func (_EtherfiAccountant *EtherfiAccountantCaller) ResetHighwaterMark(opts *bind.CallOpts) error {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "resetHighwaterMark")

	if err != nil {
		return err
	}

	return err

}

// ResetHighwaterMark is a free data retrieval call binding the contract method 0xe059ac07.
//
// Solidity: function resetHighwaterMark() view returns()
func (_EtherfiAccountant *EtherfiAccountantSession) ResetHighwaterMark() error {
	return _EtherfiAccountant.Contract.ResetHighwaterMark(&_EtherfiAccountant.CallOpts)
}

// ResetHighwaterMark is a free data retrieval call binding the contract method 0xe059ac07.
//
// Solidity: function resetHighwaterMark() view returns()
func (_EtherfiAccountant *EtherfiAccountantCallerSession) ResetHighwaterMark() error {
	return _EtherfiAccountant.Contract.ResetHighwaterMark(&_EtherfiAccountant.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiAccountant.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantSession) Vault() (common.Address, error) {
	return _EtherfiAccountant.Contract.Vault(&_EtherfiAccountant.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_EtherfiAccountant *EtherfiAccountantCallerSession) Vault() (common.Address, error) {
	return _EtherfiAccountant.Contract.Vault(&_EtherfiAccountant.CallOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x15a0ea6a.
//
// Solidity: function claimFees(address feeAsset) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) ClaimFees(opts *bind.TransactOpts, feeAsset common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "claimFees", feeAsset)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x15a0ea6a.
//
// Solidity: function claimFees(address feeAsset) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) ClaimFees(feeAsset common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.ClaimFees(&_EtherfiAccountant.TransactOpts, feeAsset)
}

// ClaimFees is a paid mutator transaction binding the contract method 0x15a0ea6a.
//
// Solidity: function claimFees(address feeAsset) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) ClaimFees(feeAsset common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.ClaimFees(&_EtherfiAccountant.TransactOpts, feeAsset)
}

// ClaimYield is a paid mutator transaction binding the contract method 0x999927df.
//
// Solidity: function claimYield(address yieldAsset) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) ClaimYield(opts *bind.TransactOpts, yieldAsset common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "claimYield", yieldAsset)
}

// ClaimYield is a paid mutator transaction binding the contract method 0x999927df.
//
// Solidity: function claimYield(address yieldAsset) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) ClaimYield(yieldAsset common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.ClaimYield(&_EtherfiAccountant.TransactOpts, yieldAsset)
}

// ClaimYield is a paid mutator transaction binding the contract method 0x999927df.
//
// Solidity: function claimYield(address yieldAsset) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) ClaimYield(yieldAsset common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.ClaimYield(&_EtherfiAccountant.TransactOpts, yieldAsset)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EtherfiAccountant *EtherfiAccountantSession) Pause() (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.Pause(&_EtherfiAccountant.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) Pause() (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.Pause(&_EtherfiAccountant.TransactOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) SetAuthority(opts *bind.TransactOpts, newAuthority common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "setAuthority", newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.SetAuthority(&_EtherfiAccountant.TransactOpts, newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.SetAuthority(&_EtherfiAccountant.TransactOpts, newAuthority)
}

// SetRateProviderData is a paid mutator transaction binding the contract method 0x4d8be07e.
//
// Solidity: function setRateProviderData(address asset, bool isPeggedToBase, address rateProvider) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) SetRateProviderData(opts *bind.TransactOpts, asset common.Address, isPeggedToBase bool, rateProvider common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "setRateProviderData", asset, isPeggedToBase, rateProvider)
}

// SetRateProviderData is a paid mutator transaction binding the contract method 0x4d8be07e.
//
// Solidity: function setRateProviderData(address asset, bool isPeggedToBase, address rateProvider) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) SetRateProviderData(asset common.Address, isPeggedToBase bool, rateProvider common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.SetRateProviderData(&_EtherfiAccountant.TransactOpts, asset, isPeggedToBase, rateProvider)
}

// SetRateProviderData is a paid mutator transaction binding the contract method 0x4d8be07e.
//
// Solidity: function setRateProviderData(address asset, bool isPeggedToBase, address rateProvider) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) SetRateProviderData(asset common.Address, isPeggedToBase bool, rateProvider common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.SetRateProviderData(&_EtherfiAccountant.TransactOpts, asset, isPeggedToBase, rateProvider)
}

// SetYieldDistributor is a paid mutator transaction binding the contract method 0x3038a60d.
//
// Solidity: function setYieldDistributor(address yieldDistributor) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) SetYieldDistributor(opts *bind.TransactOpts, yieldDistributor common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "setYieldDistributor", yieldDistributor)
}

// SetYieldDistributor is a paid mutator transaction binding the contract method 0x3038a60d.
//
// Solidity: function setYieldDistributor(address yieldDistributor) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) SetYieldDistributor(yieldDistributor common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.SetYieldDistributor(&_EtherfiAccountant.TransactOpts, yieldDistributor)
}

// SetYieldDistributor is a paid mutator transaction binding the contract method 0x3038a60d.
//
// Solidity: function setYieldDistributor(address yieldDistributor) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) SetYieldDistributor(yieldDistributor common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.SetYieldDistributor(&_EtherfiAccountant.TransactOpts, yieldDistributor)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.TransferOwnership(&_EtherfiAccountant.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.TransferOwnership(&_EtherfiAccountant.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EtherfiAccountant *EtherfiAccountantSession) Unpause() (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.Unpause(&_EtherfiAccountant.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) Unpause() (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.Unpause(&_EtherfiAccountant.TransactOpts)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x6a054dc9.
//
// Solidity: function updateDelay(uint24 minimumUpdateDelayInSeconds) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) UpdateDelay(opts *bind.TransactOpts, minimumUpdateDelayInSeconds *big.Int) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "updateDelay", minimumUpdateDelayInSeconds)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x6a054dc9.
//
// Solidity: function updateDelay(uint24 minimumUpdateDelayInSeconds) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) UpdateDelay(minimumUpdateDelayInSeconds *big.Int) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdateDelay(&_EtherfiAccountant.TransactOpts, minimumUpdateDelayInSeconds)
}

// UpdateDelay is a paid mutator transaction binding the contract method 0x6a054dc9.
//
// Solidity: function updateDelay(uint24 minimumUpdateDelayInSeconds) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) UpdateDelay(minimumUpdateDelayInSeconds *big.Int) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdateDelay(&_EtherfiAccountant.TransactOpts, minimumUpdateDelayInSeconds)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0x3458113d.
//
// Solidity: function updateExchangeRate(uint96 newExchangeRate) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) UpdateExchangeRate(opts *bind.TransactOpts, newExchangeRate *big.Int) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "updateExchangeRate", newExchangeRate)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0x3458113d.
//
// Solidity: function updateExchangeRate(uint96 newExchangeRate) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) UpdateExchangeRate(newExchangeRate *big.Int) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdateExchangeRate(&_EtherfiAccountant.TransactOpts, newExchangeRate)
}

// UpdateExchangeRate is a paid mutator transaction binding the contract method 0x3458113d.
//
// Solidity: function updateExchangeRate(uint96 newExchangeRate) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) UpdateExchangeRate(newExchangeRate *big.Int) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdateExchangeRate(&_EtherfiAccountant.TransactOpts, newExchangeRate)
}

// UpdateLower is a paid mutator transaction binding the contract method 0x207ec0e7.
//
// Solidity: function updateLower(uint16 allowedExchangeRateChangeLower) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) UpdateLower(opts *bind.TransactOpts, allowedExchangeRateChangeLower uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "updateLower", allowedExchangeRateChangeLower)
}

// UpdateLower is a paid mutator transaction binding the contract method 0x207ec0e7.
//
// Solidity: function updateLower(uint16 allowedExchangeRateChangeLower) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) UpdateLower(allowedExchangeRateChangeLower uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdateLower(&_EtherfiAccountant.TransactOpts, allowedExchangeRateChangeLower)
}

// UpdateLower is a paid mutator transaction binding the contract method 0x207ec0e7.
//
// Solidity: function updateLower(uint16 allowedExchangeRateChangeLower) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) UpdateLower(allowedExchangeRateChangeLower uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdateLower(&_EtherfiAccountant.TransactOpts, allowedExchangeRateChangeLower)
}

// UpdatePayoutAddress is a paid mutator transaction binding the contract method 0x56200819.
//
// Solidity: function updatePayoutAddress(address payoutAddress) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) UpdatePayoutAddress(opts *bind.TransactOpts, payoutAddress common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "updatePayoutAddress", payoutAddress)
}

// UpdatePayoutAddress is a paid mutator transaction binding the contract method 0x56200819.
//
// Solidity: function updatePayoutAddress(address payoutAddress) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) UpdatePayoutAddress(payoutAddress common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdatePayoutAddress(&_EtherfiAccountant.TransactOpts, payoutAddress)
}

// UpdatePayoutAddress is a paid mutator transaction binding the contract method 0x56200819.
//
// Solidity: function updatePayoutAddress(address payoutAddress) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) UpdatePayoutAddress(payoutAddress common.Address) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdatePayoutAddress(&_EtherfiAccountant.TransactOpts, payoutAddress)
}

// UpdatePerformanceFee is a paid mutator transaction binding the contract method 0x709ac1c3.
//
// Solidity: function updatePerformanceFee(uint16 performanceFee) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) UpdatePerformanceFee(opts *bind.TransactOpts, performanceFee uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "updatePerformanceFee", performanceFee)
}

// UpdatePerformanceFee is a paid mutator transaction binding the contract method 0x709ac1c3.
//
// Solidity: function updatePerformanceFee(uint16 performanceFee) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) UpdatePerformanceFee(performanceFee uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdatePerformanceFee(&_EtherfiAccountant.TransactOpts, performanceFee)
}

// UpdatePerformanceFee is a paid mutator transaction binding the contract method 0x709ac1c3.
//
// Solidity: function updatePerformanceFee(uint16 performanceFee) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) UpdatePerformanceFee(performanceFee uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdatePerformanceFee(&_EtherfiAccountant.TransactOpts, performanceFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 platformFee) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) UpdatePlatformFee(opts *bind.TransactOpts, platformFee uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "updatePlatformFee", platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 platformFee) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) UpdatePlatformFee(platformFee uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdatePlatformFee(&_EtherfiAccountant.TransactOpts, platformFee)
}

// UpdatePlatformFee is a paid mutator transaction binding the contract method 0xafb06952.
//
// Solidity: function updatePlatformFee(uint16 platformFee) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) UpdatePlatformFee(platformFee uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdatePlatformFee(&_EtherfiAccountant.TransactOpts, platformFee)
}

// UpdateUpper is a paid mutator transaction binding the contract method 0x634da58f.
//
// Solidity: function updateUpper(uint16 allowedExchangeRateChangeUpper) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactor) UpdateUpper(opts *bind.TransactOpts, allowedExchangeRateChangeUpper uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.contract.Transact(opts, "updateUpper", allowedExchangeRateChangeUpper)
}

// UpdateUpper is a paid mutator transaction binding the contract method 0x634da58f.
//
// Solidity: function updateUpper(uint16 allowedExchangeRateChangeUpper) returns()
func (_EtherfiAccountant *EtherfiAccountantSession) UpdateUpper(allowedExchangeRateChangeUpper uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdateUpper(&_EtherfiAccountant.TransactOpts, allowedExchangeRateChangeUpper)
}

// UpdateUpper is a paid mutator transaction binding the contract method 0x634da58f.
//
// Solidity: function updateUpper(uint16 allowedExchangeRateChangeUpper) returns()
func (_EtherfiAccountant *EtherfiAccountantTransactorSession) UpdateUpper(allowedExchangeRateChangeUpper uint16) (*types.Transaction, error) {
	return _EtherfiAccountant.Contract.UpdateUpper(&_EtherfiAccountant.TransactOpts, allowedExchangeRateChangeUpper)
}

// EtherfiAccountantAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the EtherfiAccountant contract.
type EtherfiAccountantAuthorityUpdatedIterator struct {
	Event *EtherfiAccountantAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantAuthorityUpdated)
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
		it.Event = new(EtherfiAccountantAuthorityUpdated)
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
func (it *EtherfiAccountantAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantAuthorityUpdated represents a AuthorityUpdated event raised by the EtherfiAccountant contract.
type EtherfiAccountantAuthorityUpdated struct {
	User         common.Address
	NewAuthority common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0xa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b76389980198.
//
// Solidity: event AuthorityUpdated(address indexed user, address indexed newAuthority)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts, user []common.Address, newAuthority []common.Address) (*EtherfiAccountantAuthorityUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAuthorityRule []interface{}
	for _, newAuthorityItem := range newAuthority {
		newAuthorityRule = append(newAuthorityRule, newAuthorityItem)
	}

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "AuthorityUpdated", userRule, newAuthorityRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantAuthorityUpdatedIterator{contract: _EtherfiAccountant.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0xa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b76389980198.
//
// Solidity: event AuthorityUpdated(address indexed user, address indexed newAuthority)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantAuthorityUpdated, user []common.Address, newAuthority []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAuthorityRule []interface{}
	for _, newAuthorityItem := range newAuthority {
		newAuthorityRule = append(newAuthorityRule, newAuthorityItem)
	}

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "AuthorityUpdated", userRule, newAuthorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantAuthorityUpdated)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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

// ParseAuthorityUpdated is a log parse operation binding the contract event 0xa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b76389980198.
//
// Solidity: event AuthorityUpdated(address indexed user, address indexed newAuthority)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseAuthorityUpdated(log types.Log) (*EtherfiAccountantAuthorityUpdated, error) {
	event := new(EtherfiAccountantAuthorityUpdated)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantDelayInSecondsUpdatedIterator is returned from FilterDelayInSecondsUpdated and is used to iterate over the raw logs and unpacked data for DelayInSecondsUpdated events raised by the EtherfiAccountant contract.
type EtherfiAccountantDelayInSecondsUpdatedIterator struct {
	Event *EtherfiAccountantDelayInSecondsUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantDelayInSecondsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantDelayInSecondsUpdated)
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
		it.Event = new(EtherfiAccountantDelayInSecondsUpdated)
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
func (it *EtherfiAccountantDelayInSecondsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantDelayInSecondsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantDelayInSecondsUpdated represents a DelayInSecondsUpdated event raised by the EtherfiAccountant contract.
type EtherfiAccountantDelayInSecondsUpdated struct {
	OldDelay *big.Int
	NewDelay *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayInSecondsUpdated is a free log retrieval operation binding the contract event 0x5f7db254db512f40348d8a7ca15d574c051dfe59c19b47e273d926f2f4318606.
//
// Solidity: event DelayInSecondsUpdated(uint24 oldDelay, uint24 newDelay)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterDelayInSecondsUpdated(opts *bind.FilterOpts) (*EtherfiAccountantDelayInSecondsUpdatedIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "DelayInSecondsUpdated")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantDelayInSecondsUpdatedIterator{contract: _EtherfiAccountant.contract, event: "DelayInSecondsUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayInSecondsUpdated is a free log subscription operation binding the contract event 0x5f7db254db512f40348d8a7ca15d574c051dfe59c19b47e273d926f2f4318606.
//
// Solidity: event DelayInSecondsUpdated(uint24 oldDelay, uint24 newDelay)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchDelayInSecondsUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantDelayInSecondsUpdated) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "DelayInSecondsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantDelayInSecondsUpdated)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "DelayInSecondsUpdated", log); err != nil {
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

// ParseDelayInSecondsUpdated is a log parse operation binding the contract event 0x5f7db254db512f40348d8a7ca15d574c051dfe59c19b47e273d926f2f4318606.
//
// Solidity: event DelayInSecondsUpdated(uint24 oldDelay, uint24 newDelay)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseDelayInSecondsUpdated(log types.Log) (*EtherfiAccountantDelayInSecondsUpdated, error) {
	event := new(EtherfiAccountantDelayInSecondsUpdated)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "DelayInSecondsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantExchangeRateUpdatedIterator is returned from FilterExchangeRateUpdated and is used to iterate over the raw logs and unpacked data for ExchangeRateUpdated events raised by the EtherfiAccountant contract.
type EtherfiAccountantExchangeRateUpdatedIterator struct {
	Event *EtherfiAccountantExchangeRateUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantExchangeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantExchangeRateUpdated)
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
		it.Event = new(EtherfiAccountantExchangeRateUpdated)
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
func (it *EtherfiAccountantExchangeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantExchangeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantExchangeRateUpdated represents a ExchangeRateUpdated event raised by the EtherfiAccountant contract.
type EtherfiAccountantExchangeRateUpdated struct {
	OldRate     *big.Int
	NewRate     *big.Int
	CurrentTime uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExchangeRateUpdated is a free log retrieval operation binding the contract event 0xa95bc6aba40bbc4d95fc35f118c4cd8b53fc5d5b89ed264002af03503a7a9439.
//
// Solidity: event ExchangeRateUpdated(uint96 oldRate, uint96 newRate, uint64 currentTime)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterExchangeRateUpdated(opts *bind.FilterOpts) (*EtherfiAccountantExchangeRateUpdatedIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "ExchangeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantExchangeRateUpdatedIterator{contract: _EtherfiAccountant.contract, event: "ExchangeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchExchangeRateUpdated is a free log subscription operation binding the contract event 0xa95bc6aba40bbc4d95fc35f118c4cd8b53fc5d5b89ed264002af03503a7a9439.
//
// Solidity: event ExchangeRateUpdated(uint96 oldRate, uint96 newRate, uint64 currentTime)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchExchangeRateUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantExchangeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "ExchangeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantExchangeRateUpdated)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "ExchangeRateUpdated", log); err != nil {
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

// ParseExchangeRateUpdated is a log parse operation binding the contract event 0xa95bc6aba40bbc4d95fc35f118c4cd8b53fc5d5b89ed264002af03503a7a9439.
//
// Solidity: event ExchangeRateUpdated(uint96 oldRate, uint96 newRate, uint64 currentTime)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseExchangeRateUpdated(log types.Log) (*EtherfiAccountantExchangeRateUpdated, error) {
	event := new(EtherfiAccountantExchangeRateUpdated)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "ExchangeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the EtherfiAccountant contract.
type EtherfiAccountantFeesClaimedIterator struct {
	Event *EtherfiAccountantFeesClaimed // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantFeesClaimed)
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
		it.Event = new(EtherfiAccountantFeesClaimed)
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
func (it *EtherfiAccountantFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantFeesClaimed represents a FeesClaimed event raised by the EtherfiAccountant contract.
type EtherfiAccountantFeesClaimed struct {
	FeeAsset common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed feeAsset, uint256 amount)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterFeesClaimed(opts *bind.FilterOpts, feeAsset []common.Address) (*EtherfiAccountantFeesClaimedIterator, error) {

	var feeAssetRule []interface{}
	for _, feeAssetItem := range feeAsset {
		feeAssetRule = append(feeAssetRule, feeAssetItem)
	}

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "FeesClaimed", feeAssetRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantFeesClaimedIterator{contract: _EtherfiAccountant.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed feeAsset, uint256 amount)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantFeesClaimed, feeAsset []common.Address) (event.Subscription, error) {

	var feeAssetRule []interface{}
	for _, feeAssetItem := range feeAsset {
		feeAssetRule = append(feeAssetRule, feeAssetItem)
	}

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "FeesClaimed", feeAssetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantFeesClaimed)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// ParseFeesClaimed is a log parse operation binding the contract event 0x9493e5bbe4e8e0ac67284469a2d677403d0378a85a59e341d3abc433d0d9a209.
//
// Solidity: event FeesClaimed(address indexed feeAsset, uint256 amount)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseFeesClaimed(log types.Log) (*EtherfiAccountantFeesClaimed, error) {
	event := new(EtherfiAccountantFeesClaimed)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantHighwaterMarkResetIterator is returned from FilterHighwaterMarkReset and is used to iterate over the raw logs and unpacked data for HighwaterMarkReset events raised by the EtherfiAccountant contract.
type EtherfiAccountantHighwaterMarkResetIterator struct {
	Event *EtherfiAccountantHighwaterMarkReset // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantHighwaterMarkResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantHighwaterMarkReset)
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
		it.Event = new(EtherfiAccountantHighwaterMarkReset)
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
func (it *EtherfiAccountantHighwaterMarkResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantHighwaterMarkResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantHighwaterMarkReset represents a HighwaterMarkReset event raised by the EtherfiAccountant contract.
type EtherfiAccountantHighwaterMarkReset struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterHighwaterMarkReset is a free log retrieval operation binding the contract event 0x98637d475d52bc596e25457cb3385a05269c42e57d4d9f7561dacbbe8583eb89.
//
// Solidity: event HighwaterMarkReset()
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterHighwaterMarkReset(opts *bind.FilterOpts) (*EtherfiAccountantHighwaterMarkResetIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "HighwaterMarkReset")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantHighwaterMarkResetIterator{contract: _EtherfiAccountant.contract, event: "HighwaterMarkReset", logs: logs, sub: sub}, nil
}

// WatchHighwaterMarkReset is a free log subscription operation binding the contract event 0x98637d475d52bc596e25457cb3385a05269c42e57d4d9f7561dacbbe8583eb89.
//
// Solidity: event HighwaterMarkReset()
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchHighwaterMarkReset(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantHighwaterMarkReset) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "HighwaterMarkReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantHighwaterMarkReset)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "HighwaterMarkReset", log); err != nil {
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

// ParseHighwaterMarkReset is a log parse operation binding the contract event 0x98637d475d52bc596e25457cb3385a05269c42e57d4d9f7561dacbbe8583eb89.
//
// Solidity: event HighwaterMarkReset()
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseHighwaterMarkReset(log types.Log) (*EtherfiAccountantHighwaterMarkReset, error) {
	event := new(EtherfiAccountantHighwaterMarkReset)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "HighwaterMarkReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantLowerBoundUpdatedIterator is returned from FilterLowerBoundUpdated and is used to iterate over the raw logs and unpacked data for LowerBoundUpdated events raised by the EtherfiAccountant contract.
type EtherfiAccountantLowerBoundUpdatedIterator struct {
	Event *EtherfiAccountantLowerBoundUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantLowerBoundUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantLowerBoundUpdated)
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
		it.Event = new(EtherfiAccountantLowerBoundUpdated)
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
func (it *EtherfiAccountantLowerBoundUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantLowerBoundUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantLowerBoundUpdated represents a LowerBoundUpdated event raised by the EtherfiAccountant contract.
type EtherfiAccountantLowerBoundUpdated struct {
	OldBound uint16
	NewBound uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLowerBoundUpdated is a free log retrieval operation binding the contract event 0x76fe3c3557dd03afa5caf76f66f4019444ef3999e784ba08f47a33428fcc64d5.
//
// Solidity: event LowerBoundUpdated(uint16 oldBound, uint16 newBound)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterLowerBoundUpdated(opts *bind.FilterOpts) (*EtherfiAccountantLowerBoundUpdatedIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "LowerBoundUpdated")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantLowerBoundUpdatedIterator{contract: _EtherfiAccountant.contract, event: "LowerBoundUpdated", logs: logs, sub: sub}, nil
}

// WatchLowerBoundUpdated is a free log subscription operation binding the contract event 0x76fe3c3557dd03afa5caf76f66f4019444ef3999e784ba08f47a33428fcc64d5.
//
// Solidity: event LowerBoundUpdated(uint16 oldBound, uint16 newBound)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchLowerBoundUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantLowerBoundUpdated) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "LowerBoundUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantLowerBoundUpdated)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "LowerBoundUpdated", log); err != nil {
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

// ParseLowerBoundUpdated is a log parse operation binding the contract event 0x76fe3c3557dd03afa5caf76f66f4019444ef3999e784ba08f47a33428fcc64d5.
//
// Solidity: event LowerBoundUpdated(uint16 oldBound, uint16 newBound)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseLowerBoundUpdated(log types.Log) (*EtherfiAccountantLowerBoundUpdated, error) {
	event := new(EtherfiAccountantLowerBoundUpdated)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "LowerBoundUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EtherfiAccountant contract.
type EtherfiAccountantOwnershipTransferredIterator struct {
	Event *EtherfiAccountantOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantOwnershipTransferred)
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
		it.Event = new(EtherfiAccountantOwnershipTransferred)
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
func (it *EtherfiAccountantOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantOwnershipTransferred represents a OwnershipTransferred event raised by the EtherfiAccountant contract.
type EtherfiAccountantOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*EtherfiAccountantOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantOwnershipTransferredIterator{contract: _EtherfiAccountant.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantOwnershipTransferred)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseOwnershipTransferred(log types.Log) (*EtherfiAccountantOwnershipTransferred, error) {
	event := new(EtherfiAccountantOwnershipTransferred)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EtherfiAccountant contract.
type EtherfiAccountantPausedIterator struct {
	Event *EtherfiAccountantPaused // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantPaused)
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
		it.Event = new(EtherfiAccountantPaused)
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
func (it *EtherfiAccountantPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantPaused represents a Paused event raised by the EtherfiAccountant contract.
type EtherfiAccountantPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterPaused(opts *bind.FilterOpts) (*EtherfiAccountantPausedIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantPausedIterator{contract: _EtherfiAccountant.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantPaused) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantPaused)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParsePaused(log types.Log) (*EtherfiAccountantPaused, error) {
	event := new(EtherfiAccountantPaused)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantPayoutAddressUpdatedIterator is returned from FilterPayoutAddressUpdated and is used to iterate over the raw logs and unpacked data for PayoutAddressUpdated events raised by the EtherfiAccountant contract.
type EtherfiAccountantPayoutAddressUpdatedIterator struct {
	Event *EtherfiAccountantPayoutAddressUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantPayoutAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantPayoutAddressUpdated)
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
		it.Event = new(EtherfiAccountantPayoutAddressUpdated)
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
func (it *EtherfiAccountantPayoutAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantPayoutAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantPayoutAddressUpdated represents a PayoutAddressUpdated event raised by the EtherfiAccountant contract.
type EtherfiAccountantPayoutAddressUpdated struct {
	OldPayout common.Address
	NewPayout common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayoutAddressUpdated is a free log retrieval operation binding the contract event 0xba2be5e898fed1646bc0814dee1cc9a2aee98f51fced7d5fc4699c47d9907753.
//
// Solidity: event PayoutAddressUpdated(address oldPayout, address newPayout)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterPayoutAddressUpdated(opts *bind.FilterOpts) (*EtherfiAccountantPayoutAddressUpdatedIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "PayoutAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantPayoutAddressUpdatedIterator{contract: _EtherfiAccountant.contract, event: "PayoutAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchPayoutAddressUpdated is a free log subscription operation binding the contract event 0xba2be5e898fed1646bc0814dee1cc9a2aee98f51fced7d5fc4699c47d9907753.
//
// Solidity: event PayoutAddressUpdated(address oldPayout, address newPayout)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchPayoutAddressUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantPayoutAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "PayoutAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantPayoutAddressUpdated)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "PayoutAddressUpdated", log); err != nil {
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

// ParsePayoutAddressUpdated is a log parse operation binding the contract event 0xba2be5e898fed1646bc0814dee1cc9a2aee98f51fced7d5fc4699c47d9907753.
//
// Solidity: event PayoutAddressUpdated(address oldPayout, address newPayout)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParsePayoutAddressUpdated(log types.Log) (*EtherfiAccountantPayoutAddressUpdated, error) {
	event := new(EtherfiAccountantPayoutAddressUpdated)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "PayoutAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantPerformanceFeeUpdatedIterator is returned from FilterPerformanceFeeUpdated and is used to iterate over the raw logs and unpacked data for PerformanceFeeUpdated events raised by the EtherfiAccountant contract.
type EtherfiAccountantPerformanceFeeUpdatedIterator struct {
	Event *EtherfiAccountantPerformanceFeeUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantPerformanceFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantPerformanceFeeUpdated)
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
		it.Event = new(EtherfiAccountantPerformanceFeeUpdated)
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
func (it *EtherfiAccountantPerformanceFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantPerformanceFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantPerformanceFeeUpdated represents a PerformanceFeeUpdated event raised by the EtherfiAccountant contract.
type EtherfiAccountantPerformanceFeeUpdated struct {
	OldFee uint16
	NewFee uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPerformanceFeeUpdated is a free log retrieval operation binding the contract event 0xba8506b6cb85330fea21cbca8490aafb6a69b166f06201ef755eb511b2709fc1.
//
// Solidity: event PerformanceFeeUpdated(uint16 oldFee, uint16 newFee)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterPerformanceFeeUpdated(opts *bind.FilterOpts) (*EtherfiAccountantPerformanceFeeUpdatedIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "PerformanceFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantPerformanceFeeUpdatedIterator{contract: _EtherfiAccountant.contract, event: "PerformanceFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPerformanceFeeUpdated is a free log subscription operation binding the contract event 0xba8506b6cb85330fea21cbca8490aafb6a69b166f06201ef755eb511b2709fc1.
//
// Solidity: event PerformanceFeeUpdated(uint16 oldFee, uint16 newFee)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchPerformanceFeeUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantPerformanceFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "PerformanceFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantPerformanceFeeUpdated)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "PerformanceFeeUpdated", log); err != nil {
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

// ParsePerformanceFeeUpdated is a log parse operation binding the contract event 0xba8506b6cb85330fea21cbca8490aafb6a69b166f06201ef755eb511b2709fc1.
//
// Solidity: event PerformanceFeeUpdated(uint16 oldFee, uint16 newFee)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParsePerformanceFeeUpdated(log types.Log) (*EtherfiAccountantPerformanceFeeUpdated, error) {
	event := new(EtherfiAccountantPerformanceFeeUpdated)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "PerformanceFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantPlatformFeeUpdatedIterator is returned from FilterPlatformFeeUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeUpdated events raised by the EtherfiAccountant contract.
type EtherfiAccountantPlatformFeeUpdatedIterator struct {
	Event *EtherfiAccountantPlatformFeeUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantPlatformFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantPlatformFeeUpdated)
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
		it.Event = new(EtherfiAccountantPlatformFeeUpdated)
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
func (it *EtherfiAccountantPlatformFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantPlatformFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantPlatformFeeUpdated represents a PlatformFeeUpdated event raised by the EtherfiAccountant contract.
type EtherfiAccountantPlatformFeeUpdated struct {
	OldFee uint16
	NewFee uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeUpdated is a free log retrieval operation binding the contract event 0x84e4fe32bf74c4011a7e1fde79c63acdffaf92a0112cde153e7b0abee665bc6b.
//
// Solidity: event PlatformFeeUpdated(uint16 oldFee, uint16 newFee)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterPlatformFeeUpdated(opts *bind.FilterOpts) (*EtherfiAccountantPlatformFeeUpdatedIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantPlatformFeeUpdatedIterator{contract: _EtherfiAccountant.contract, event: "PlatformFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeUpdated is a free log subscription operation binding the contract event 0x84e4fe32bf74c4011a7e1fde79c63acdffaf92a0112cde153e7b0abee665bc6b.
//
// Solidity: event PlatformFeeUpdated(uint16 oldFee, uint16 newFee)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchPlatformFeeUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantPlatformFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantPlatformFeeUpdated)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
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

// ParsePlatformFeeUpdated is a log parse operation binding the contract event 0x84e4fe32bf74c4011a7e1fde79c63acdffaf92a0112cde153e7b0abee665bc6b.
//
// Solidity: event PlatformFeeUpdated(uint16 oldFee, uint16 newFee)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParsePlatformFeeUpdated(log types.Log) (*EtherfiAccountantPlatformFeeUpdated, error) {
	event := new(EtherfiAccountantPlatformFeeUpdated)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantRateProviderUpdatedIterator is returned from FilterRateProviderUpdated and is used to iterate over the raw logs and unpacked data for RateProviderUpdated events raised by the EtherfiAccountant contract.
type EtherfiAccountantRateProviderUpdatedIterator struct {
	Event *EtherfiAccountantRateProviderUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantRateProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantRateProviderUpdated)
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
		it.Event = new(EtherfiAccountantRateProviderUpdated)
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
func (it *EtherfiAccountantRateProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantRateProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantRateProviderUpdated represents a RateProviderUpdated event raised by the EtherfiAccountant contract.
type EtherfiAccountantRateProviderUpdated struct {
	Asset        common.Address
	IsPegged     bool
	RateProvider common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRateProviderUpdated is a free log retrieval operation binding the contract event 0x59f9adfe8cf4c9d4b77fb03aa2ae5f373632c97cb8caf6b61f0643d3d170a8fe.
//
// Solidity: event RateProviderUpdated(address asset, bool isPegged, address rateProvider)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterRateProviderUpdated(opts *bind.FilterOpts) (*EtherfiAccountantRateProviderUpdatedIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "RateProviderUpdated")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantRateProviderUpdatedIterator{contract: _EtherfiAccountant.contract, event: "RateProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchRateProviderUpdated is a free log subscription operation binding the contract event 0x59f9adfe8cf4c9d4b77fb03aa2ae5f373632c97cb8caf6b61f0643d3d170a8fe.
//
// Solidity: event RateProviderUpdated(address asset, bool isPegged, address rateProvider)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchRateProviderUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantRateProviderUpdated) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "RateProviderUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantRateProviderUpdated)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "RateProviderUpdated", log); err != nil {
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

// ParseRateProviderUpdated is a log parse operation binding the contract event 0x59f9adfe8cf4c9d4b77fb03aa2ae5f373632c97cb8caf6b61f0643d3d170a8fe.
//
// Solidity: event RateProviderUpdated(address asset, bool isPegged, address rateProvider)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseRateProviderUpdated(log types.Log) (*EtherfiAccountantRateProviderUpdated, error) {
	event := new(EtherfiAccountantRateProviderUpdated)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "RateProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EtherfiAccountant contract.
type EtherfiAccountantUnpausedIterator struct {
	Event *EtherfiAccountantUnpaused // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantUnpaused)
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
		it.Event = new(EtherfiAccountantUnpaused)
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
func (it *EtherfiAccountantUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantUnpaused represents a Unpaused event raised by the EtherfiAccountant contract.
type EtherfiAccountantUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterUnpaused(opts *bind.FilterOpts) (*EtherfiAccountantUnpausedIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantUnpausedIterator{contract: _EtherfiAccountant.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantUnpaused) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantUnpaused)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseUnpaused(log types.Log) (*EtherfiAccountantUnpaused, error) {
	event := new(EtherfiAccountantUnpaused)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantUpperBoundUpdatedIterator is returned from FilterUpperBoundUpdated and is used to iterate over the raw logs and unpacked data for UpperBoundUpdated events raised by the EtherfiAccountant contract.
type EtherfiAccountantUpperBoundUpdatedIterator struct {
	Event *EtherfiAccountantUpperBoundUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantUpperBoundUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantUpperBoundUpdated)
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
		it.Event = new(EtherfiAccountantUpperBoundUpdated)
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
func (it *EtherfiAccountantUpperBoundUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantUpperBoundUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantUpperBoundUpdated represents a UpperBoundUpdated event raised by the EtherfiAccountant contract.
type EtherfiAccountantUpperBoundUpdated struct {
	OldBound uint16
	NewBound uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpperBoundUpdated is a free log retrieval operation binding the contract event 0x67d3a3f6bebb5b894324217d5224ff719d5d95dfc67f1bb2645dddbfcd43cadb.
//
// Solidity: event UpperBoundUpdated(uint16 oldBound, uint16 newBound)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterUpperBoundUpdated(opts *bind.FilterOpts) (*EtherfiAccountantUpperBoundUpdatedIterator, error) {

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "UpperBoundUpdated")
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantUpperBoundUpdatedIterator{contract: _EtherfiAccountant.contract, event: "UpperBoundUpdated", logs: logs, sub: sub}, nil
}

// WatchUpperBoundUpdated is a free log subscription operation binding the contract event 0x67d3a3f6bebb5b894324217d5224ff719d5d95dfc67f1bb2645dddbfcd43cadb.
//
// Solidity: event UpperBoundUpdated(uint16 oldBound, uint16 newBound)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchUpperBoundUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantUpperBoundUpdated) (event.Subscription, error) {

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "UpperBoundUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantUpperBoundUpdated)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "UpperBoundUpdated", log); err != nil {
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

// ParseUpperBoundUpdated is a log parse operation binding the contract event 0x67d3a3f6bebb5b894324217d5224ff719d5d95dfc67f1bb2645dddbfcd43cadb.
//
// Solidity: event UpperBoundUpdated(uint16 oldBound, uint16 newBound)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseUpperBoundUpdated(log types.Log) (*EtherfiAccountantUpperBoundUpdated, error) {
	event := new(EtherfiAccountantUpperBoundUpdated)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "UpperBoundUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantYieldClaimedIterator is returned from FilterYieldClaimed and is used to iterate over the raw logs and unpacked data for YieldClaimed events raised by the EtherfiAccountant contract.
type EtherfiAccountantYieldClaimedIterator struct {
	Event *EtherfiAccountantYieldClaimed // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantYieldClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantYieldClaimed)
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
		it.Event = new(EtherfiAccountantYieldClaimed)
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
func (it *EtherfiAccountantYieldClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantYieldClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantYieldClaimed represents a YieldClaimed event raised by the EtherfiAccountant contract.
type EtherfiAccountantYieldClaimed struct {
	YieldAsset common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterYieldClaimed is a free log retrieval operation binding the contract event 0xc04825ba3f383b602255d2a13065a68e325c65c9e0ed5d031ea2b06f641873af.
//
// Solidity: event YieldClaimed(address indexed yieldAsset, uint256 amount)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterYieldClaimed(opts *bind.FilterOpts, yieldAsset []common.Address) (*EtherfiAccountantYieldClaimedIterator, error) {

	var yieldAssetRule []interface{}
	for _, yieldAssetItem := range yieldAsset {
		yieldAssetRule = append(yieldAssetRule, yieldAssetItem)
	}

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "YieldClaimed", yieldAssetRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantYieldClaimedIterator{contract: _EtherfiAccountant.contract, event: "YieldClaimed", logs: logs, sub: sub}, nil
}

// WatchYieldClaimed is a free log subscription operation binding the contract event 0xc04825ba3f383b602255d2a13065a68e325c65c9e0ed5d031ea2b06f641873af.
//
// Solidity: event YieldClaimed(address indexed yieldAsset, uint256 amount)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchYieldClaimed(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantYieldClaimed, yieldAsset []common.Address) (event.Subscription, error) {

	var yieldAssetRule []interface{}
	for _, yieldAssetItem := range yieldAsset {
		yieldAssetRule = append(yieldAssetRule, yieldAssetItem)
	}

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "YieldClaimed", yieldAssetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantYieldClaimed)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "YieldClaimed", log); err != nil {
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

// ParseYieldClaimed is a log parse operation binding the contract event 0xc04825ba3f383b602255d2a13065a68e325c65c9e0ed5d031ea2b06f641873af.
//
// Solidity: event YieldClaimed(address indexed yieldAsset, uint256 amount)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseYieldClaimed(log types.Log) (*EtherfiAccountantYieldClaimed, error) {
	event := new(EtherfiAccountantYieldClaimed)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "YieldClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiAccountantYieldDistributorUpdatedIterator is returned from FilterYieldDistributorUpdated and is used to iterate over the raw logs and unpacked data for YieldDistributorUpdated events raised by the EtherfiAccountant contract.
type EtherfiAccountantYieldDistributorUpdatedIterator struct {
	Event *EtherfiAccountantYieldDistributorUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiAccountantYieldDistributorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiAccountantYieldDistributorUpdated)
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
		it.Event = new(EtherfiAccountantYieldDistributorUpdated)
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
func (it *EtherfiAccountantYieldDistributorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiAccountantYieldDistributorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiAccountantYieldDistributorUpdated represents a YieldDistributorUpdated event raised by the EtherfiAccountant contract.
type EtherfiAccountantYieldDistributorUpdated struct {
	YieldDistributor common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterYieldDistributorUpdated is a free log retrieval operation binding the contract event 0x7d9c3ef9e65227fa9a8638f9e876cf890ef686bad3ab18e6c3a3f7cb9de258a0.
//
// Solidity: event YieldDistributorUpdated(address indexed yieldDistributor)
func (_EtherfiAccountant *EtherfiAccountantFilterer) FilterYieldDistributorUpdated(opts *bind.FilterOpts, yieldDistributor []common.Address) (*EtherfiAccountantYieldDistributorUpdatedIterator, error) {

	var yieldDistributorRule []interface{}
	for _, yieldDistributorItem := range yieldDistributor {
		yieldDistributorRule = append(yieldDistributorRule, yieldDistributorItem)
	}

	logs, sub, err := _EtherfiAccountant.contract.FilterLogs(opts, "YieldDistributorUpdated", yieldDistributorRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiAccountantYieldDistributorUpdatedIterator{contract: _EtherfiAccountant.contract, event: "YieldDistributorUpdated", logs: logs, sub: sub}, nil
}

// WatchYieldDistributorUpdated is a free log subscription operation binding the contract event 0x7d9c3ef9e65227fa9a8638f9e876cf890ef686bad3ab18e6c3a3f7cb9de258a0.
//
// Solidity: event YieldDistributorUpdated(address indexed yieldDistributor)
func (_EtherfiAccountant *EtherfiAccountantFilterer) WatchYieldDistributorUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiAccountantYieldDistributorUpdated, yieldDistributor []common.Address) (event.Subscription, error) {

	var yieldDistributorRule []interface{}
	for _, yieldDistributorItem := range yieldDistributor {
		yieldDistributorRule = append(yieldDistributorRule, yieldDistributorItem)
	}

	logs, sub, err := _EtherfiAccountant.contract.WatchLogs(opts, "YieldDistributorUpdated", yieldDistributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiAccountantYieldDistributorUpdated)
				if err := _EtherfiAccountant.contract.UnpackLog(event, "YieldDistributorUpdated", log); err != nil {
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

// ParseYieldDistributorUpdated is a log parse operation binding the contract event 0x7d9c3ef9e65227fa9a8638f9e876cf890ef686bad3ab18e6c3a3f7cb9de258a0.
//
// Solidity: event YieldDistributorUpdated(address indexed yieldDistributor)
func (_EtherfiAccountant *EtherfiAccountantFilterer) ParseYieldDistributorUpdated(log types.Log) (*EtherfiAccountantYieldDistributorUpdated, error) {
	event := new(EtherfiAccountantYieldDistributorUpdated)
	if err := _EtherfiAccountant.contract.UnpackLog(event, "YieldDistributorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
