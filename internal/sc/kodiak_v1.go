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

// KodiakV1MetaData contains all meta data concerning the KodiakV1 contract.
var KodiakV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_gelato\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_kodiakTreasury\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBurned\",\"type\":\"uint128\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesEarned0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesEarned1\",\"type\":\"uint256\"}],\"name\":\"FeesEarned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityMinted\",\"type\":\"uint128\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"lowerTick_\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"upperTick_\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBefore\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAfter\",\"type\":\"uint128\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"managerFeeBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"managerTreasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gelatoRebalanceBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gelatoSlippageBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gelatoSlippageInterval\",\"type\":\"uint32\"}],\"name\":\"UpdateManagerParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GELATO\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESTRICTED_MINT_ENABLED\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityBurned\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newUpperTick\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"swapThresholdPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"swapAmountBPS\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"}],\"name\":\"executiveRebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoRebalanceBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoSlippageBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoSlippageInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Max\",\"type\":\"uint256\"}],\"name\":\"getMintAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositionID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"positionID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnderlyingBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Current\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtRatioX96\",\"type\":\"uint160\"}],\"name\":\"getUnderlyingBalancesAtPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Current\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_managerFeeBPS\",\"type\":\"uint16\"},{\"internalType\":\"int24\",\"name\":\"_lowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_upperTick\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"_manager_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kodiakBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kodiakBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kodiakFeeBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kodiakTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowerTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerFeeBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityMinted\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"swapThresholdPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"swapAmountBPS\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restrictedMintToggle\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleRestrictMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"newManagerFeeBPS\",\"type\":\"int16\"},{\"internalType\":\"address\",\"name\":\"newManagerTreasury\",\"type\":\"address\"},{\"internalType\":\"int16\",\"name\":\"newRebalanceBPS\",\"type\":\"int16\"},{\"internalType\":\"int16\",\"name\":\"newSlippageBPS\",\"type\":\"int16\"},{\"internalType\":\"int32\",\"name\":\"newSlippageInterval\",\"type\":\"int32\"}],\"name\":\"updateManagerParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawKodiakBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawManagerBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KodiakV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use KodiakV1MetaData.ABI instead.
var KodiakV1ABI = KodiakV1MetaData.ABI

// KodiakV1 is an auto generated Go binding around an Ethereum contract.
type KodiakV1 struct {
	KodiakV1Caller     // Read-only binding to the contract
	KodiakV1Transactor // Write-only binding to the contract
	KodiakV1Filterer   // Log filterer for contract events
}

// KodiakV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type KodiakV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KodiakV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type KodiakV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KodiakV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KodiakV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KodiakV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KodiakV1Session struct {
	Contract     *KodiakV1         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KodiakV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KodiakV1CallerSession struct {
	Contract *KodiakV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// KodiakV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KodiakV1TransactorSession struct {
	Contract     *KodiakV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KodiakV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type KodiakV1Raw struct {
	Contract *KodiakV1 // Generic contract binding to access the raw methods on
}

// KodiakV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KodiakV1CallerRaw struct {
	Contract *KodiakV1Caller // Generic read-only contract binding to access the raw methods on
}

// KodiakV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KodiakV1TransactorRaw struct {
	Contract *KodiakV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKodiakV1 creates a new instance of KodiakV1, bound to a specific deployed contract.
func NewKodiakV1(address common.Address, backend bind.ContractBackend) (*KodiakV1, error) {
	contract, err := bindKodiakV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KodiakV1{KodiakV1Caller: KodiakV1Caller{contract: contract}, KodiakV1Transactor: KodiakV1Transactor{contract: contract}, KodiakV1Filterer: KodiakV1Filterer{contract: contract}}, nil
}

// NewKodiakV1Caller creates a new read-only instance of KodiakV1, bound to a specific deployed contract.
func NewKodiakV1Caller(address common.Address, caller bind.ContractCaller) (*KodiakV1Caller, error) {
	contract, err := bindKodiakV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KodiakV1Caller{contract: contract}, nil
}

// NewKodiakV1Transactor creates a new write-only instance of KodiakV1, bound to a specific deployed contract.
func NewKodiakV1Transactor(address common.Address, transactor bind.ContractTransactor) (*KodiakV1Transactor, error) {
	contract, err := bindKodiakV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KodiakV1Transactor{contract: contract}, nil
}

// NewKodiakV1Filterer creates a new log filterer instance of KodiakV1, bound to a specific deployed contract.
func NewKodiakV1Filterer(address common.Address, filterer bind.ContractFilterer) (*KodiakV1Filterer, error) {
	contract, err := bindKodiakV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KodiakV1Filterer{contract: contract}, nil
}

// bindKodiakV1 binds a generic wrapper to an already deployed contract.
func bindKodiakV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KodiakV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KodiakV1 *KodiakV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KodiakV1.Contract.KodiakV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KodiakV1 *KodiakV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakV1.Contract.KodiakV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KodiakV1 *KodiakV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KodiakV1.Contract.KodiakV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KodiakV1 *KodiakV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KodiakV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KodiakV1 *KodiakV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KodiakV1 *KodiakV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KodiakV1.Contract.contract.Transact(opts, method, params...)
}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_KodiakV1 *KodiakV1Caller) GELATO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "GELATO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_KodiakV1 *KodiakV1Session) GELATO() (common.Address, error) {
	return _KodiakV1.Contract.GELATO(&_KodiakV1.CallOpts)
}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_KodiakV1 *KodiakV1CallerSession) GELATO() (common.Address, error) {
	return _KodiakV1.Contract.GELATO(&_KodiakV1.CallOpts)
}

// RESTRICTEDMINTENABLED is a free data retrieval call binding the contract method 0xd839fb3f.
//
// Solidity: function RESTRICTED_MINT_ENABLED() view returns(uint16)
func (_KodiakV1 *KodiakV1Caller) RESTRICTEDMINTENABLED(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "RESTRICTED_MINT_ENABLED")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RESTRICTEDMINTENABLED is a free data retrieval call binding the contract method 0xd839fb3f.
//
// Solidity: function RESTRICTED_MINT_ENABLED() view returns(uint16)
func (_KodiakV1 *KodiakV1Session) RESTRICTEDMINTENABLED() (uint16, error) {
	return _KodiakV1.Contract.RESTRICTEDMINTENABLED(&_KodiakV1.CallOpts)
}

// RESTRICTEDMINTENABLED is a free data retrieval call binding the contract method 0xd839fb3f.
//
// Solidity: function RESTRICTED_MINT_ENABLED() view returns(uint16)
func (_KodiakV1 *KodiakV1CallerSession) RESTRICTEDMINTENABLED() (uint16, error) {
	return _KodiakV1.Contract.RESTRICTEDMINTENABLED(&_KodiakV1.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KodiakV1 *KodiakV1Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KodiakV1 *KodiakV1Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KodiakV1.Contract.Allowance(&_KodiakV1.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KodiakV1 *KodiakV1CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KodiakV1.Contract.Allowance(&_KodiakV1.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KodiakV1 *KodiakV1Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KodiakV1 *KodiakV1Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _KodiakV1.Contract.BalanceOf(&_KodiakV1.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KodiakV1 *KodiakV1CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KodiakV1.Contract.BalanceOf(&_KodiakV1.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KodiakV1 *KodiakV1Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KodiakV1 *KodiakV1Session) Decimals() (uint8, error) {
	return _KodiakV1.Contract.Decimals(&_KodiakV1.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KodiakV1 *KodiakV1CallerSession) Decimals() (uint8, error) {
	return _KodiakV1.Contract.Decimals(&_KodiakV1.CallOpts)
}

// GelatoRebalanceBPS is a free data retrieval call binding the contract method 0xa50b1fe7.
//
// Solidity: function gelatoRebalanceBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1Caller) GelatoRebalanceBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "gelatoRebalanceBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GelatoRebalanceBPS is a free data retrieval call binding the contract method 0xa50b1fe7.
//
// Solidity: function gelatoRebalanceBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1Session) GelatoRebalanceBPS() (uint16, error) {
	return _KodiakV1.Contract.GelatoRebalanceBPS(&_KodiakV1.CallOpts)
}

// GelatoRebalanceBPS is a free data retrieval call binding the contract method 0xa50b1fe7.
//
// Solidity: function gelatoRebalanceBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1CallerSession) GelatoRebalanceBPS() (uint16, error) {
	return _KodiakV1.Contract.GelatoRebalanceBPS(&_KodiakV1.CallOpts)
}

// GelatoSlippageBPS is a free data retrieval call binding the contract method 0xd6e7ff39.
//
// Solidity: function gelatoSlippageBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1Caller) GelatoSlippageBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "gelatoSlippageBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GelatoSlippageBPS is a free data retrieval call binding the contract method 0xd6e7ff39.
//
// Solidity: function gelatoSlippageBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1Session) GelatoSlippageBPS() (uint16, error) {
	return _KodiakV1.Contract.GelatoSlippageBPS(&_KodiakV1.CallOpts)
}

// GelatoSlippageBPS is a free data retrieval call binding the contract method 0xd6e7ff39.
//
// Solidity: function gelatoSlippageBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1CallerSession) GelatoSlippageBPS() (uint16, error) {
	return _KodiakV1.Contract.GelatoSlippageBPS(&_KodiakV1.CallOpts)
}

// GelatoSlippageInterval is a free data retrieval call binding the contract method 0x672152bd.
//
// Solidity: function gelatoSlippageInterval() view returns(uint32)
func (_KodiakV1 *KodiakV1Caller) GelatoSlippageInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "gelatoSlippageInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GelatoSlippageInterval is a free data retrieval call binding the contract method 0x672152bd.
//
// Solidity: function gelatoSlippageInterval() view returns(uint32)
func (_KodiakV1 *KodiakV1Session) GelatoSlippageInterval() (uint32, error) {
	return _KodiakV1.Contract.GelatoSlippageInterval(&_KodiakV1.CallOpts)
}

// GelatoSlippageInterval is a free data retrieval call binding the contract method 0x672152bd.
//
// Solidity: function gelatoSlippageInterval() view returns(uint32)
func (_KodiakV1 *KodiakV1CallerSession) GelatoSlippageInterval() (uint32, error) {
	return _KodiakV1.Contract.GelatoSlippageInterval(&_KodiakV1.CallOpts)
}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_KodiakV1 *KodiakV1Caller) GetMintAmounts(opts *bind.CallOpts, amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "getMintAmounts", amount0Max, amount1Max)

	outstruct := new(struct {
		Amount0    *big.Int
		Amount1    *big.Int
		MintAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MintAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_KodiakV1 *KodiakV1Session) GetMintAmounts(amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	return _KodiakV1.Contract.GetMintAmounts(&_KodiakV1.CallOpts, amount0Max, amount1Max)
}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_KodiakV1 *KodiakV1CallerSession) GetMintAmounts(amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	return _KodiakV1.Contract.GetMintAmounts(&_KodiakV1.CallOpts, amount0Max, amount1Max)
}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_KodiakV1 *KodiakV1Caller) GetPositionID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "getPositionID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_KodiakV1 *KodiakV1Session) GetPositionID() ([32]byte, error) {
	return _KodiakV1.Contract.GetPositionID(&_KodiakV1.CallOpts)
}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_KodiakV1 *KodiakV1CallerSession) GetPositionID() ([32]byte, error) {
	return _KodiakV1.Contract.GetPositionID(&_KodiakV1.CallOpts)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_KodiakV1 *KodiakV1Caller) GetUnderlyingBalances(opts *bind.CallOpts) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "getUnderlyingBalances")

	outstruct := new(struct {
		Amount0Current *big.Int
		Amount1Current *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0Current = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1Current = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_KodiakV1 *KodiakV1Session) GetUnderlyingBalances() (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _KodiakV1.Contract.GetUnderlyingBalances(&_KodiakV1.CallOpts)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_KodiakV1 *KodiakV1CallerSession) GetUnderlyingBalances() (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _KodiakV1.Contract.GetUnderlyingBalances(&_KodiakV1.CallOpts)
}

// GetUnderlyingBalancesAtPrice is a free data retrieval call binding the contract method 0xb670ed7d.
//
// Solidity: function getUnderlyingBalancesAtPrice(uint160 sqrtRatioX96) view returns(uint256 amount0Current, uint256 amount1Current)
func (_KodiakV1 *KodiakV1Caller) GetUnderlyingBalancesAtPrice(opts *bind.CallOpts, sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "getUnderlyingBalancesAtPrice", sqrtRatioX96)

	outstruct := new(struct {
		Amount0Current *big.Int
		Amount1Current *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0Current = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1Current = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUnderlyingBalancesAtPrice is a free data retrieval call binding the contract method 0xb670ed7d.
//
// Solidity: function getUnderlyingBalancesAtPrice(uint160 sqrtRatioX96) view returns(uint256 amount0Current, uint256 amount1Current)
func (_KodiakV1 *KodiakV1Session) GetUnderlyingBalancesAtPrice(sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _KodiakV1.Contract.GetUnderlyingBalancesAtPrice(&_KodiakV1.CallOpts, sqrtRatioX96)
}

// GetUnderlyingBalancesAtPrice is a free data retrieval call binding the contract method 0xb670ed7d.
//
// Solidity: function getUnderlyingBalancesAtPrice(uint160 sqrtRatioX96) view returns(uint256 amount0Current, uint256 amount1Current)
func (_KodiakV1 *KodiakV1CallerSession) GetUnderlyingBalancesAtPrice(sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _KodiakV1.Contract.GetUnderlyingBalancesAtPrice(&_KodiakV1.CallOpts, sqrtRatioX96)
}

// KodiakBalance0 is a free data retrieval call binding the contract method 0xd8d24145.
//
// Solidity: function kodiakBalance0() view returns(uint256)
func (_KodiakV1 *KodiakV1Caller) KodiakBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "kodiakBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KodiakBalance0 is a free data retrieval call binding the contract method 0xd8d24145.
//
// Solidity: function kodiakBalance0() view returns(uint256)
func (_KodiakV1 *KodiakV1Session) KodiakBalance0() (*big.Int, error) {
	return _KodiakV1.Contract.KodiakBalance0(&_KodiakV1.CallOpts)
}

// KodiakBalance0 is a free data retrieval call binding the contract method 0xd8d24145.
//
// Solidity: function kodiakBalance0() view returns(uint256)
func (_KodiakV1 *KodiakV1CallerSession) KodiakBalance0() (*big.Int, error) {
	return _KodiakV1.Contract.KodiakBalance0(&_KodiakV1.CallOpts)
}

// KodiakBalance1 is a free data retrieval call binding the contract method 0xf5338bcf.
//
// Solidity: function kodiakBalance1() view returns(uint256)
func (_KodiakV1 *KodiakV1Caller) KodiakBalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "kodiakBalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KodiakBalance1 is a free data retrieval call binding the contract method 0xf5338bcf.
//
// Solidity: function kodiakBalance1() view returns(uint256)
func (_KodiakV1 *KodiakV1Session) KodiakBalance1() (*big.Int, error) {
	return _KodiakV1.Contract.KodiakBalance1(&_KodiakV1.CallOpts)
}

// KodiakBalance1 is a free data retrieval call binding the contract method 0xf5338bcf.
//
// Solidity: function kodiakBalance1() view returns(uint256)
func (_KodiakV1 *KodiakV1CallerSession) KodiakBalance1() (*big.Int, error) {
	return _KodiakV1.Contract.KodiakBalance1(&_KodiakV1.CallOpts)
}

// KodiakFeeBPS is a free data retrieval call binding the contract method 0xce311070.
//
// Solidity: function kodiakFeeBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1Caller) KodiakFeeBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "kodiakFeeBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// KodiakFeeBPS is a free data retrieval call binding the contract method 0xce311070.
//
// Solidity: function kodiakFeeBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1Session) KodiakFeeBPS() (uint16, error) {
	return _KodiakV1.Contract.KodiakFeeBPS(&_KodiakV1.CallOpts)
}

// KodiakFeeBPS is a free data retrieval call binding the contract method 0xce311070.
//
// Solidity: function kodiakFeeBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1CallerSession) KodiakFeeBPS() (uint16, error) {
	return _KodiakV1.Contract.KodiakFeeBPS(&_KodiakV1.CallOpts)
}

// KodiakTreasury is a free data retrieval call binding the contract method 0x94a113ec.
//
// Solidity: function kodiakTreasury() view returns(address)
func (_KodiakV1 *KodiakV1Caller) KodiakTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "kodiakTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KodiakTreasury is a free data retrieval call binding the contract method 0x94a113ec.
//
// Solidity: function kodiakTreasury() view returns(address)
func (_KodiakV1 *KodiakV1Session) KodiakTreasury() (common.Address, error) {
	return _KodiakV1.Contract.KodiakTreasury(&_KodiakV1.CallOpts)
}

// KodiakTreasury is a free data retrieval call binding the contract method 0x94a113ec.
//
// Solidity: function kodiakTreasury() view returns(address)
func (_KodiakV1 *KodiakV1CallerSession) KodiakTreasury() (common.Address, error) {
	return _KodiakV1.Contract.KodiakTreasury(&_KodiakV1.CallOpts)
}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_KodiakV1 *KodiakV1Caller) LowerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "lowerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_KodiakV1 *KodiakV1Session) LowerTick() (*big.Int, error) {
	return _KodiakV1.Contract.LowerTick(&_KodiakV1.CallOpts)
}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_KodiakV1 *KodiakV1CallerSession) LowerTick() (*big.Int, error) {
	return _KodiakV1.Contract.LowerTick(&_KodiakV1.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_KodiakV1 *KodiakV1Caller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_KodiakV1 *KodiakV1Session) Manager() (common.Address, error) {
	return _KodiakV1.Contract.Manager(&_KodiakV1.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_KodiakV1 *KodiakV1CallerSession) Manager() (common.Address, error) {
	return _KodiakV1.Contract.Manager(&_KodiakV1.CallOpts)
}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_KodiakV1 *KodiakV1Caller) ManagerBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "managerBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_KodiakV1 *KodiakV1Session) ManagerBalance0() (*big.Int, error) {
	return _KodiakV1.Contract.ManagerBalance0(&_KodiakV1.CallOpts)
}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_KodiakV1 *KodiakV1CallerSession) ManagerBalance0() (*big.Int, error) {
	return _KodiakV1.Contract.ManagerBalance0(&_KodiakV1.CallOpts)
}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_KodiakV1 *KodiakV1Caller) ManagerBalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "managerBalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_KodiakV1 *KodiakV1Session) ManagerBalance1() (*big.Int, error) {
	return _KodiakV1.Contract.ManagerBalance1(&_KodiakV1.CallOpts)
}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_KodiakV1 *KodiakV1CallerSession) ManagerBalance1() (*big.Int, error) {
	return _KodiakV1.Contract.ManagerBalance1(&_KodiakV1.CallOpts)
}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1Caller) ManagerFeeBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "managerFeeBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1Session) ManagerFeeBPS() (uint16, error) {
	return _KodiakV1.Contract.ManagerFeeBPS(&_KodiakV1.CallOpts)
}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_KodiakV1 *KodiakV1CallerSession) ManagerFeeBPS() (uint16, error) {
	return _KodiakV1.Contract.ManagerFeeBPS(&_KodiakV1.CallOpts)
}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_KodiakV1 *KodiakV1Caller) ManagerTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "managerTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_KodiakV1 *KodiakV1Session) ManagerTreasury() (common.Address, error) {
	return _KodiakV1.Contract.ManagerTreasury(&_KodiakV1.CallOpts)
}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_KodiakV1 *KodiakV1CallerSession) ManagerTreasury() (common.Address, error) {
	return _KodiakV1.Contract.ManagerTreasury(&_KodiakV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KodiakV1 *KodiakV1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KodiakV1 *KodiakV1Session) Name() (string, error) {
	return _KodiakV1.Contract.Name(&_KodiakV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KodiakV1 *KodiakV1CallerSession) Name() (string, error) {
	return _KodiakV1.Contract.Name(&_KodiakV1.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_KodiakV1 *KodiakV1Caller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_KodiakV1 *KodiakV1Session) Pool() (common.Address, error) {
	return _KodiakV1.Contract.Pool(&_KodiakV1.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_KodiakV1 *KodiakV1CallerSession) Pool() (common.Address, error) {
	return _KodiakV1.Contract.Pool(&_KodiakV1.CallOpts)
}

// RestrictedMintToggle is a free data retrieval call binding the contract method 0xdd40e322.
//
// Solidity: function restrictedMintToggle() view returns(uint16)
func (_KodiakV1 *KodiakV1Caller) RestrictedMintToggle(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "restrictedMintToggle")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RestrictedMintToggle is a free data retrieval call binding the contract method 0xdd40e322.
//
// Solidity: function restrictedMintToggle() view returns(uint16)
func (_KodiakV1 *KodiakV1Session) RestrictedMintToggle() (uint16, error) {
	return _KodiakV1.Contract.RestrictedMintToggle(&_KodiakV1.CallOpts)
}

// RestrictedMintToggle is a free data retrieval call binding the contract method 0xdd40e322.
//
// Solidity: function restrictedMintToggle() view returns(uint16)
func (_KodiakV1 *KodiakV1CallerSession) RestrictedMintToggle() (uint16, error) {
	return _KodiakV1.Contract.RestrictedMintToggle(&_KodiakV1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KodiakV1 *KodiakV1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KodiakV1 *KodiakV1Session) Symbol() (string, error) {
	return _KodiakV1.Contract.Symbol(&_KodiakV1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KodiakV1 *KodiakV1CallerSession) Symbol() (string, error) {
	return _KodiakV1.Contract.Symbol(&_KodiakV1.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KodiakV1 *KodiakV1Caller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KodiakV1 *KodiakV1Session) Token0() (common.Address, error) {
	return _KodiakV1.Contract.Token0(&_KodiakV1.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KodiakV1 *KodiakV1CallerSession) Token0() (common.Address, error) {
	return _KodiakV1.Contract.Token0(&_KodiakV1.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KodiakV1 *KodiakV1Caller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KodiakV1 *KodiakV1Session) Token1() (common.Address, error) {
	return _KodiakV1.Contract.Token1(&_KodiakV1.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KodiakV1 *KodiakV1CallerSession) Token1() (common.Address, error) {
	return _KodiakV1.Contract.Token1(&_KodiakV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KodiakV1 *KodiakV1Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KodiakV1 *KodiakV1Session) TotalSupply() (*big.Int, error) {
	return _KodiakV1.Contract.TotalSupply(&_KodiakV1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KodiakV1 *KodiakV1CallerSession) TotalSupply() (*big.Int, error) {
	return _KodiakV1.Contract.TotalSupply(&_KodiakV1.CallOpts)
}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_KodiakV1 *KodiakV1Caller) UpperTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "upperTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_KodiakV1 *KodiakV1Session) UpperTick() (*big.Int, error) {
	return _KodiakV1.Contract.UpperTick(&_KodiakV1.CallOpts)
}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_KodiakV1 *KodiakV1CallerSession) UpperTick() (*big.Int, error) {
	return _KodiakV1.Contract.UpperTick(&_KodiakV1.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KodiakV1 *KodiakV1Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KodiakV1.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KodiakV1 *KodiakV1Session) Version() (string, error) {
	return _KodiakV1.Contract.Version(&_KodiakV1.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KodiakV1 *KodiakV1CallerSession) Version() (string, error) {
	return _KodiakV1.Contract.Version(&_KodiakV1.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KodiakV1 *KodiakV1Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KodiakV1 *KodiakV1Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakV1.Contract.Approve(&_KodiakV1.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KodiakV1 *KodiakV1TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakV1.Contract.Approve(&_KodiakV1.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_KodiakV1 *KodiakV1Transactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "burn", burnAmount, receiver)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_KodiakV1 *KodiakV1Session) Burn(burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakV1.Contract.Burn(&_KodiakV1.TransactOpts, burnAmount, receiver)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_KodiakV1 *KodiakV1TransactorSession) Burn(burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakV1.Contract.Burn(&_KodiakV1.TransactOpts, burnAmount, receiver)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KodiakV1 *KodiakV1Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KodiakV1 *KodiakV1Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KodiakV1.Contract.DecreaseAllowance(&_KodiakV1.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KodiakV1 *KodiakV1TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KodiakV1.Contract.DecreaseAllowance(&_KodiakV1.TransactOpts, spender, subtractedValue)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_KodiakV1 *KodiakV1Transactor) ExecutiveRebalance(opts *bind.TransactOpts, newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "executiveRebalance", newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_KodiakV1 *KodiakV1Session) ExecutiveRebalance(newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _KodiakV1.Contract.ExecutiveRebalance(&_KodiakV1.TransactOpts, newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_KodiakV1 *KodiakV1TransactorSession) ExecutiveRebalance(newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _KodiakV1.Contract.ExecutiveRebalance(&_KodiakV1.TransactOpts, newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KodiakV1 *KodiakV1Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KodiakV1 *KodiakV1Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KodiakV1.Contract.IncreaseAllowance(&_KodiakV1.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KodiakV1 *KodiakV1TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KodiakV1.Contract.IncreaseAllowance(&_KodiakV1.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xe25e15e3.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_) returns()
func (_KodiakV1 *KodiakV1Transactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "initialize", _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe25e15e3.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_) returns()
func (_KodiakV1 *KodiakV1Session) Initialize(_name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address) (*types.Transaction, error) {
	return _KodiakV1.Contract.Initialize(&_KodiakV1.TransactOpts, _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe25e15e3.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_) returns()
func (_KodiakV1 *KodiakV1TransactorSession) Initialize(_name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address) (*types.Transaction, error) {
	return _KodiakV1.Contract.Initialize(&_KodiakV1.TransactOpts, _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_KodiakV1 *KodiakV1Transactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "mint", mintAmount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_KodiakV1 *KodiakV1Session) Mint(mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakV1.Contract.Mint(&_KodiakV1.TransactOpts, mintAmount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_KodiakV1 *KodiakV1TransactorSession) Mint(mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakV1.Contract.Mint(&_KodiakV1.TransactOpts, mintAmount, receiver)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb135c99f.
//
// Solidity: function rebalance(uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne, uint256 feeAmount, address paymentToken) returns()
func (_KodiakV1 *KodiakV1Transactor) Rebalance(opts *bind.TransactOpts, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool, feeAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "rebalance", swapThresholdPrice, swapAmountBPS, zeroForOne, feeAmount, paymentToken)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb135c99f.
//
// Solidity: function rebalance(uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne, uint256 feeAmount, address paymentToken) returns()
func (_KodiakV1 *KodiakV1Session) Rebalance(swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool, feeAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _KodiakV1.Contract.Rebalance(&_KodiakV1.TransactOpts, swapThresholdPrice, swapAmountBPS, zeroForOne, feeAmount, paymentToken)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb135c99f.
//
// Solidity: function rebalance(uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne, uint256 feeAmount, address paymentToken) returns()
func (_KodiakV1 *KodiakV1TransactorSession) Rebalance(swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool, feeAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _KodiakV1.Contract.Rebalance(&_KodiakV1.TransactOpts, swapThresholdPrice, swapAmountBPS, zeroForOne, feeAmount, paymentToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KodiakV1 *KodiakV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KodiakV1 *KodiakV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _KodiakV1.Contract.RenounceOwnership(&_KodiakV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KodiakV1 *KodiakV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KodiakV1.Contract.RenounceOwnership(&_KodiakV1.TransactOpts)
}

// ToggleRestrictMint is a paid mutator transaction binding the contract method 0x5f59dce5.
//
// Solidity: function toggleRestrictMint() returns()
func (_KodiakV1 *KodiakV1Transactor) ToggleRestrictMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "toggleRestrictMint")
}

// ToggleRestrictMint is a paid mutator transaction binding the contract method 0x5f59dce5.
//
// Solidity: function toggleRestrictMint() returns()
func (_KodiakV1 *KodiakV1Session) ToggleRestrictMint() (*types.Transaction, error) {
	return _KodiakV1.Contract.ToggleRestrictMint(&_KodiakV1.TransactOpts)
}

// ToggleRestrictMint is a paid mutator transaction binding the contract method 0x5f59dce5.
//
// Solidity: function toggleRestrictMint() returns()
func (_KodiakV1 *KodiakV1TransactorSession) ToggleRestrictMint() (*types.Transaction, error) {
	return _KodiakV1.Contract.ToggleRestrictMint(&_KodiakV1.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KodiakV1 *KodiakV1Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KodiakV1 *KodiakV1Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakV1.Contract.Transfer(&_KodiakV1.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KodiakV1 *KodiakV1TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakV1.Contract.Transfer(&_KodiakV1.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KodiakV1 *KodiakV1Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KodiakV1 *KodiakV1Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakV1.Contract.TransferFrom(&_KodiakV1.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KodiakV1 *KodiakV1TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakV1.Contract.TransferFrom(&_KodiakV1.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KodiakV1 *KodiakV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KodiakV1 *KodiakV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KodiakV1.Contract.TransferOwnership(&_KodiakV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KodiakV1 *KodiakV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KodiakV1.Contract.TransferOwnership(&_KodiakV1.TransactOpts, newOwner)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_KodiakV1 *KodiakV1Transactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "uniswapV3MintCallback", amount0Owed, amount1Owed, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_KodiakV1 *KodiakV1Session) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakV1.Contract.UniswapV3MintCallback(&_KodiakV1.TransactOpts, amount0Owed, amount1Owed, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_KodiakV1 *KodiakV1TransactorSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakV1.Contract.UniswapV3MintCallback(&_KodiakV1.TransactOpts, amount0Owed, amount1Owed, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_KodiakV1 *KodiakV1Transactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_KodiakV1 *KodiakV1Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakV1.Contract.UniswapV3SwapCallback(&_KodiakV1.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_KodiakV1 *KodiakV1TransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakV1.Contract.UniswapV3SwapCallback(&_KodiakV1.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UpdateManagerParams is a paid mutator transaction binding the contract method 0xc540e483.
//
// Solidity: function updateManagerParams(int16 newManagerFeeBPS, address newManagerTreasury, int16 newRebalanceBPS, int16 newSlippageBPS, int32 newSlippageInterval) returns()
func (_KodiakV1 *KodiakV1Transactor) UpdateManagerParams(opts *bind.TransactOpts, newManagerFeeBPS int16, newManagerTreasury common.Address, newRebalanceBPS int16, newSlippageBPS int16, newSlippageInterval int32) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "updateManagerParams", newManagerFeeBPS, newManagerTreasury, newRebalanceBPS, newSlippageBPS, newSlippageInterval)
}

// UpdateManagerParams is a paid mutator transaction binding the contract method 0xc540e483.
//
// Solidity: function updateManagerParams(int16 newManagerFeeBPS, address newManagerTreasury, int16 newRebalanceBPS, int16 newSlippageBPS, int32 newSlippageInterval) returns()
func (_KodiakV1 *KodiakV1Session) UpdateManagerParams(newManagerFeeBPS int16, newManagerTreasury common.Address, newRebalanceBPS int16, newSlippageBPS int16, newSlippageInterval int32) (*types.Transaction, error) {
	return _KodiakV1.Contract.UpdateManagerParams(&_KodiakV1.TransactOpts, newManagerFeeBPS, newManagerTreasury, newRebalanceBPS, newSlippageBPS, newSlippageInterval)
}

// UpdateManagerParams is a paid mutator transaction binding the contract method 0xc540e483.
//
// Solidity: function updateManagerParams(int16 newManagerFeeBPS, address newManagerTreasury, int16 newRebalanceBPS, int16 newSlippageBPS, int32 newSlippageInterval) returns()
func (_KodiakV1 *KodiakV1TransactorSession) UpdateManagerParams(newManagerFeeBPS int16, newManagerTreasury common.Address, newRebalanceBPS int16, newSlippageBPS int16, newSlippageInterval int32) (*types.Transaction, error) {
	return _KodiakV1.Contract.UpdateManagerParams(&_KodiakV1.TransactOpts, newManagerFeeBPS, newManagerTreasury, newRebalanceBPS, newSlippageBPS, newSlippageInterval)
}

// WithdrawKodiakBalance is a paid mutator transaction binding the contract method 0xd83bfb90.
//
// Solidity: function withdrawKodiakBalance() returns()
func (_KodiakV1 *KodiakV1Transactor) WithdrawKodiakBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "withdrawKodiakBalance")
}

// WithdrawKodiakBalance is a paid mutator transaction binding the contract method 0xd83bfb90.
//
// Solidity: function withdrawKodiakBalance() returns()
func (_KodiakV1 *KodiakV1Session) WithdrawKodiakBalance() (*types.Transaction, error) {
	return _KodiakV1.Contract.WithdrawKodiakBalance(&_KodiakV1.TransactOpts)
}

// WithdrawKodiakBalance is a paid mutator transaction binding the contract method 0xd83bfb90.
//
// Solidity: function withdrawKodiakBalance() returns()
func (_KodiakV1 *KodiakV1TransactorSession) WithdrawKodiakBalance() (*types.Transaction, error) {
	return _KodiakV1.Contract.WithdrawKodiakBalance(&_KodiakV1.TransactOpts)
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x7ecd6717.
//
// Solidity: function withdrawManagerBalance() returns()
func (_KodiakV1 *KodiakV1Transactor) WithdrawManagerBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakV1.contract.Transact(opts, "withdrawManagerBalance")
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x7ecd6717.
//
// Solidity: function withdrawManagerBalance() returns()
func (_KodiakV1 *KodiakV1Session) WithdrawManagerBalance() (*types.Transaction, error) {
	return _KodiakV1.Contract.WithdrawManagerBalance(&_KodiakV1.TransactOpts)
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x7ecd6717.
//
// Solidity: function withdrawManagerBalance() returns()
func (_KodiakV1 *KodiakV1TransactorSession) WithdrawManagerBalance() (*types.Transaction, error) {
	return _KodiakV1.Contract.WithdrawManagerBalance(&_KodiakV1.TransactOpts)
}

// KodiakV1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KodiakV1 contract.
type KodiakV1ApprovalIterator struct {
	Event *KodiakV1Approval // Event containing the contract specifics and raw log

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
func (it *KodiakV1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakV1Approval)
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
		it.Event = new(KodiakV1Approval)
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
func (it *KodiakV1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakV1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakV1Approval represents a Approval event raised by the KodiakV1 contract.
type KodiakV1Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KodiakV1 *KodiakV1Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KodiakV1ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KodiakV1.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KodiakV1ApprovalIterator{contract: _KodiakV1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KodiakV1 *KodiakV1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KodiakV1Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KodiakV1.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakV1Approval)
				if err := _KodiakV1.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KodiakV1 *KodiakV1Filterer) ParseApproval(log types.Log) (*KodiakV1Approval, error) {
	event := new(KodiakV1Approval)
	if err := _KodiakV1.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakV1BurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the KodiakV1 contract.
type KodiakV1BurnedIterator struct {
	Event *KodiakV1Burned // Event containing the contract specifics and raw log

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
func (it *KodiakV1BurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakV1Burned)
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
		it.Event = new(KodiakV1Burned)
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
func (it *KodiakV1BurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakV1BurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakV1Burned represents a Burned event raised by the KodiakV1 contract.
type KodiakV1Burned struct {
	Receiver        common.Address
	BurnAmount      *big.Int
	Amount0Out      *big.Int
	Amount1Out      *big.Int
	LiquidityBurned *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x7239dff1718b550db7f36cbf69c665cfeb56d0e96b4fb76a5cba712961b65509.
//
// Solidity: event Burned(address receiver, uint256 burnAmount, uint256 amount0Out, uint256 amount1Out, uint128 liquidityBurned)
func (_KodiakV1 *KodiakV1Filterer) FilterBurned(opts *bind.FilterOpts) (*KodiakV1BurnedIterator, error) {

	logs, sub, err := _KodiakV1.contract.FilterLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return &KodiakV1BurnedIterator{contract: _KodiakV1.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x7239dff1718b550db7f36cbf69c665cfeb56d0e96b4fb76a5cba712961b65509.
//
// Solidity: event Burned(address receiver, uint256 burnAmount, uint256 amount0Out, uint256 amount1Out, uint128 liquidityBurned)
func (_KodiakV1 *KodiakV1Filterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *KodiakV1Burned) (event.Subscription, error) {

	logs, sub, err := _KodiakV1.contract.WatchLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakV1Burned)
				if err := _KodiakV1.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x7239dff1718b550db7f36cbf69c665cfeb56d0e96b4fb76a5cba712961b65509.
//
// Solidity: event Burned(address receiver, uint256 burnAmount, uint256 amount0Out, uint256 amount1Out, uint128 liquidityBurned)
func (_KodiakV1 *KodiakV1Filterer) ParseBurned(log types.Log) (*KodiakV1Burned, error) {
	event := new(KodiakV1Burned)
	if err := _KodiakV1.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakV1FeesEarnedIterator is returned from FilterFeesEarned and is used to iterate over the raw logs and unpacked data for FeesEarned events raised by the KodiakV1 contract.
type KodiakV1FeesEarnedIterator struct {
	Event *KodiakV1FeesEarned // Event containing the contract specifics and raw log

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
func (it *KodiakV1FeesEarnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakV1FeesEarned)
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
		it.Event = new(KodiakV1FeesEarned)
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
func (it *KodiakV1FeesEarnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakV1FeesEarnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakV1FeesEarned represents a FeesEarned event raised by the KodiakV1 contract.
type KodiakV1FeesEarned struct {
	FeesEarned0 *big.Int
	FeesEarned1 *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeesEarned is a free log retrieval operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 feesEarned0, uint256 feesEarned1)
func (_KodiakV1 *KodiakV1Filterer) FilterFeesEarned(opts *bind.FilterOpts) (*KodiakV1FeesEarnedIterator, error) {

	logs, sub, err := _KodiakV1.contract.FilterLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return &KodiakV1FeesEarnedIterator{contract: _KodiakV1.contract, event: "FeesEarned", logs: logs, sub: sub}, nil
}

// WatchFeesEarned is a free log subscription operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 feesEarned0, uint256 feesEarned1)
func (_KodiakV1 *KodiakV1Filterer) WatchFeesEarned(opts *bind.WatchOpts, sink chan<- *KodiakV1FeesEarned) (event.Subscription, error) {

	logs, sub, err := _KodiakV1.contract.WatchLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakV1FeesEarned)
				if err := _KodiakV1.contract.UnpackLog(event, "FeesEarned", log); err != nil {
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

// ParseFeesEarned is a log parse operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 feesEarned0, uint256 feesEarned1)
func (_KodiakV1 *KodiakV1Filterer) ParseFeesEarned(log types.Log) (*KodiakV1FeesEarned, error) {
	event := new(KodiakV1FeesEarned)
	if err := _KodiakV1.contract.UnpackLog(event, "FeesEarned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakV1MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the KodiakV1 contract.
type KodiakV1MintedIterator struct {
	Event *KodiakV1Minted // Event containing the contract specifics and raw log

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
func (it *KodiakV1MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakV1Minted)
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
		it.Event = new(KodiakV1Minted)
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
func (it *KodiakV1MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakV1MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakV1Minted represents a Minted event raised by the KodiakV1 contract.
type KodiakV1Minted struct {
	Receiver        common.Address
	MintAmount      *big.Int
	Amount0In       *big.Int
	Amount1In       *big.Int
	LiquidityMinted *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x55801cfe493000b734571da1694b21e7f66b11e8ce9fdaa0524ecb59105e73e7.
//
// Solidity: event Minted(address receiver, uint256 mintAmount, uint256 amount0In, uint256 amount1In, uint128 liquidityMinted)
func (_KodiakV1 *KodiakV1Filterer) FilterMinted(opts *bind.FilterOpts) (*KodiakV1MintedIterator, error) {

	logs, sub, err := _KodiakV1.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &KodiakV1MintedIterator{contract: _KodiakV1.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x55801cfe493000b734571da1694b21e7f66b11e8ce9fdaa0524ecb59105e73e7.
//
// Solidity: event Minted(address receiver, uint256 mintAmount, uint256 amount0In, uint256 amount1In, uint128 liquidityMinted)
func (_KodiakV1 *KodiakV1Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *KodiakV1Minted) (event.Subscription, error) {

	logs, sub, err := _KodiakV1.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakV1Minted)
				if err := _KodiakV1.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x55801cfe493000b734571da1694b21e7f66b11e8ce9fdaa0524ecb59105e73e7.
//
// Solidity: event Minted(address receiver, uint256 mintAmount, uint256 amount0In, uint256 amount1In, uint128 liquidityMinted)
func (_KodiakV1 *KodiakV1Filterer) ParseMinted(log types.Log) (*KodiakV1Minted, error) {
	event := new(KodiakV1Minted)
	if err := _KodiakV1.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KodiakV1 contract.
type KodiakV1OwnershipTransferredIterator struct {
	Event *KodiakV1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KodiakV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakV1OwnershipTransferred)
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
		it.Event = new(KodiakV1OwnershipTransferred)
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
func (it *KodiakV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakV1OwnershipTransferred represents a OwnershipTransferred event raised by the KodiakV1 contract.
type KodiakV1OwnershipTransferred struct {
	PreviousManager common.Address
	NewManager      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousManager, address indexed newManager)
func (_KodiakV1 *KodiakV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousManager []common.Address, newManager []common.Address) (*KodiakV1OwnershipTransferredIterator, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _KodiakV1.contract.FilterLogs(opts, "OwnershipTransferred", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &KodiakV1OwnershipTransferredIterator{contract: _KodiakV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousManager, address indexed newManager)
func (_KodiakV1 *KodiakV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KodiakV1OwnershipTransferred, previousManager []common.Address, newManager []common.Address) (event.Subscription, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _KodiakV1.contract.WatchLogs(opts, "OwnershipTransferred", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakV1OwnershipTransferred)
				if err := _KodiakV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousManager, address indexed newManager)
func (_KodiakV1 *KodiakV1Filterer) ParseOwnershipTransferred(log types.Log) (*KodiakV1OwnershipTransferred, error) {
	event := new(KodiakV1OwnershipTransferred)
	if err := _KodiakV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakV1RebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the KodiakV1 contract.
type KodiakV1RebalanceIterator struct {
	Event *KodiakV1Rebalance // Event containing the contract specifics and raw log

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
func (it *KodiakV1RebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakV1Rebalance)
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
		it.Event = new(KodiakV1Rebalance)
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
func (it *KodiakV1RebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakV1RebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakV1Rebalance represents a Rebalance event raised by the KodiakV1 contract.
type KodiakV1Rebalance struct {
	LowerTick       *big.Int
	UpperTick       *big.Int
	LiquidityBefore *big.Int
	LiquidityAfter  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRebalance is a free log retrieval operation binding the contract event 0xc749f9ae947d4734cf1569606a8a347391ae94a063478aa853aeff48ac5f99e8.
//
// Solidity: event Rebalance(int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_KodiakV1 *KodiakV1Filterer) FilterRebalance(opts *bind.FilterOpts) (*KodiakV1RebalanceIterator, error) {

	logs, sub, err := _KodiakV1.contract.FilterLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return &KodiakV1RebalanceIterator{contract: _KodiakV1.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0xc749f9ae947d4734cf1569606a8a347391ae94a063478aa853aeff48ac5f99e8.
//
// Solidity: event Rebalance(int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_KodiakV1 *KodiakV1Filterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *KodiakV1Rebalance) (event.Subscription, error) {

	logs, sub, err := _KodiakV1.contract.WatchLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakV1Rebalance)
				if err := _KodiakV1.contract.UnpackLog(event, "Rebalance", log); err != nil {
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

// ParseRebalance is a log parse operation binding the contract event 0xc749f9ae947d4734cf1569606a8a347391ae94a063478aa853aeff48ac5f99e8.
//
// Solidity: event Rebalance(int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_KodiakV1 *KodiakV1Filterer) ParseRebalance(log types.Log) (*KodiakV1Rebalance, error) {
	event := new(KodiakV1Rebalance)
	if err := _KodiakV1.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakV1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KodiakV1 contract.
type KodiakV1TransferIterator struct {
	Event *KodiakV1Transfer // Event containing the contract specifics and raw log

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
func (it *KodiakV1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakV1Transfer)
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
		it.Event = new(KodiakV1Transfer)
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
func (it *KodiakV1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakV1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakV1Transfer represents a Transfer event raised by the KodiakV1 contract.
type KodiakV1Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KodiakV1 *KodiakV1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KodiakV1TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KodiakV1.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KodiakV1TransferIterator{contract: _KodiakV1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KodiakV1 *KodiakV1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KodiakV1Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KodiakV1.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakV1Transfer)
				if err := _KodiakV1.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KodiakV1 *KodiakV1Filterer) ParseTransfer(log types.Log) (*KodiakV1Transfer, error) {
	event := new(KodiakV1Transfer)
	if err := _KodiakV1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakV1UpdateManagerParamsIterator is returned from FilterUpdateManagerParams and is used to iterate over the raw logs and unpacked data for UpdateManagerParams events raised by the KodiakV1 contract.
type KodiakV1UpdateManagerParamsIterator struct {
	Event *KodiakV1UpdateManagerParams // Event containing the contract specifics and raw log

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
func (it *KodiakV1UpdateManagerParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakV1UpdateManagerParams)
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
		it.Event = new(KodiakV1UpdateManagerParams)
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
func (it *KodiakV1UpdateManagerParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakV1UpdateManagerParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakV1UpdateManagerParams represents a UpdateManagerParams event raised by the KodiakV1 contract.
type KodiakV1UpdateManagerParams struct {
	ManagerFeeBPS          uint16
	ManagerTreasury        common.Address
	GelatoRebalanceBPS     uint16
	GelatoSlippageBPS      uint16
	GelatoSlippageInterval uint32
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagerParams is a free log retrieval operation binding the contract event 0xf459b381c988c676562c20a01f42f488c5560ed2e3957b0f27f26573418b6193.
//
// Solidity: event UpdateManagerParams(uint16 managerFeeBPS, address managerTreasury, uint16 gelatoRebalanceBPS, uint16 gelatoSlippageBPS, uint32 gelatoSlippageInterval)
func (_KodiakV1 *KodiakV1Filterer) FilterUpdateManagerParams(opts *bind.FilterOpts) (*KodiakV1UpdateManagerParamsIterator, error) {

	logs, sub, err := _KodiakV1.contract.FilterLogs(opts, "UpdateManagerParams")
	if err != nil {
		return nil, err
	}
	return &KodiakV1UpdateManagerParamsIterator{contract: _KodiakV1.contract, event: "UpdateManagerParams", logs: logs, sub: sub}, nil
}

// WatchUpdateManagerParams is a free log subscription operation binding the contract event 0xf459b381c988c676562c20a01f42f488c5560ed2e3957b0f27f26573418b6193.
//
// Solidity: event UpdateManagerParams(uint16 managerFeeBPS, address managerTreasury, uint16 gelatoRebalanceBPS, uint16 gelatoSlippageBPS, uint32 gelatoSlippageInterval)
func (_KodiakV1 *KodiakV1Filterer) WatchUpdateManagerParams(opts *bind.WatchOpts, sink chan<- *KodiakV1UpdateManagerParams) (event.Subscription, error) {

	logs, sub, err := _KodiakV1.contract.WatchLogs(opts, "UpdateManagerParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakV1UpdateManagerParams)
				if err := _KodiakV1.contract.UnpackLog(event, "UpdateManagerParams", log); err != nil {
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

// ParseUpdateManagerParams is a log parse operation binding the contract event 0xf459b381c988c676562c20a01f42f488c5560ed2e3957b0f27f26573418b6193.
//
// Solidity: event UpdateManagerParams(uint16 managerFeeBPS, address managerTreasury, uint16 gelatoRebalanceBPS, uint16 gelatoSlippageBPS, uint32 gelatoSlippageInterval)
func (_KodiakV1 *KodiakV1Filterer) ParseUpdateManagerParams(log types.Log) (*KodiakV1UpdateManagerParams, error) {
	event := new(KodiakV1UpdateManagerParams)
	if err := _KodiakV1.contract.UnpackLog(event, "UpdateManagerParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
