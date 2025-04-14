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

// KodiakIslandMetaData contains all meta data concerning the KodiakIsland contract.
var KodiakIslandMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_gelato\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_kodiakTreasury\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBurned\",\"type\":\"uint128\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesEarned0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesEarned1\",\"type\":\"uint256\"}],\"name\":\"FeesEarned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityMinted\",\"type\":\"uint128\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"lowerTick_\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"upperTick_\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBefore\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAfter\",\"type\":\"uint128\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"managerFeeBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"managerTreasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gelatoRebalanceBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"gelatoSlippageBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gelatoSlippageInterval\",\"type\":\"uint32\"}],\"name\":\"UpdateManagerParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GELATO\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESTRICTED_MINT_ENABLED\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityBurned\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newUpperTick\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"swapThresholdPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"swapAmountBPS\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"}],\"name\":\"executiveRebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoRebalanceBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoSlippageBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gelatoSlippageInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Max\",\"type\":\"uint256\"}],\"name\":\"getMintAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositionID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"positionID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnderlyingBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Current\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtRatioX96\",\"type\":\"uint160\"}],\"name\":\"getUnderlyingBalancesAtPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Current\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_managerFeeBPS\",\"type\":\"uint16\"},{\"internalType\":\"int24\",\"name\":\"_lowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_upperTick\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"_manager_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kodiakBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kodiakBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kodiakFeeBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kodiakTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowerTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerFeeBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityMinted\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"swapThresholdPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"swapAmountBPS\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restrictedMintToggle\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleRestrictMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"newManagerFeeBPS\",\"type\":\"int16\"},{\"internalType\":\"address\",\"name\":\"newManagerTreasury\",\"type\":\"address\"},{\"internalType\":\"int16\",\"name\":\"newRebalanceBPS\",\"type\":\"int16\"},{\"internalType\":\"int16\",\"name\":\"newSlippageBPS\",\"type\":\"int16\"},{\"internalType\":\"int32\",\"name\":\"newSlippageInterval\",\"type\":\"int32\"}],\"name\":\"updateManagerParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawKodiakBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawManagerBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KodiakIslandABI is the input ABI used to generate the binding from.
// Deprecated: Use KodiakIslandMetaData.ABI instead.
var KodiakIslandABI = KodiakIslandMetaData.ABI

// KodiakIsland is an auto generated Go binding around an Ethereum contract.
type KodiakIsland struct {
	KodiakIslandCaller     // Read-only binding to the contract
	KodiakIslandTransactor // Write-only binding to the contract
	KodiakIslandFilterer   // Log filterer for contract events
}

// KodiakIslandCaller is an auto generated read-only Go binding around an Ethereum contract.
type KodiakIslandCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KodiakIslandTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KodiakIslandTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KodiakIslandFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KodiakIslandFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KodiakIslandSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KodiakIslandSession struct {
	Contract     *KodiakIsland     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KodiakIslandCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KodiakIslandCallerSession struct {
	Contract *KodiakIslandCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KodiakIslandTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KodiakIslandTransactorSession struct {
	Contract     *KodiakIslandTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KodiakIslandRaw is an auto generated low-level Go binding around an Ethereum contract.
type KodiakIslandRaw struct {
	Contract *KodiakIsland // Generic contract binding to access the raw methods on
}

// KodiakIslandCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KodiakIslandCallerRaw struct {
	Contract *KodiakIslandCaller // Generic read-only contract binding to access the raw methods on
}

// KodiakIslandTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KodiakIslandTransactorRaw struct {
	Contract *KodiakIslandTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKodiakIsland creates a new instance of KodiakIsland, bound to a specific deployed contract.
func NewKodiakIsland(address common.Address, backend bind.ContractBackend) (*KodiakIsland, error) {
	contract, err := bindKodiakIsland(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KodiakIsland{KodiakIslandCaller: KodiakIslandCaller{contract: contract}, KodiakIslandTransactor: KodiakIslandTransactor{contract: contract}, KodiakIslandFilterer: KodiakIslandFilterer{contract: contract}}, nil
}

// NewKodiakIslandCaller creates a new read-only instance of KodiakIsland, bound to a specific deployed contract.
func NewKodiakIslandCaller(address common.Address, caller bind.ContractCaller) (*KodiakIslandCaller, error) {
	contract, err := bindKodiakIsland(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KodiakIslandCaller{contract: contract}, nil
}

// NewKodiakIslandTransactor creates a new write-only instance of KodiakIsland, bound to a specific deployed contract.
func NewKodiakIslandTransactor(address common.Address, transactor bind.ContractTransactor) (*KodiakIslandTransactor, error) {
	contract, err := bindKodiakIsland(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KodiakIslandTransactor{contract: contract}, nil
}

// NewKodiakIslandFilterer creates a new log filterer instance of KodiakIsland, bound to a specific deployed contract.
func NewKodiakIslandFilterer(address common.Address, filterer bind.ContractFilterer) (*KodiakIslandFilterer, error) {
	contract, err := bindKodiakIsland(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KodiakIslandFilterer{contract: contract}, nil
}

// bindKodiakIsland binds a generic wrapper to an already deployed contract.
func bindKodiakIsland(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KodiakIslandMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KodiakIsland *KodiakIslandRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KodiakIsland.Contract.KodiakIslandCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KodiakIsland *KodiakIslandRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakIsland.Contract.KodiakIslandTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KodiakIsland *KodiakIslandRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KodiakIsland.Contract.KodiakIslandTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KodiakIsland *KodiakIslandCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KodiakIsland.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KodiakIsland *KodiakIslandTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakIsland.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KodiakIsland *KodiakIslandTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KodiakIsland.Contract.contract.Transact(opts, method, params...)
}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_KodiakIsland *KodiakIslandCaller) GELATO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "GELATO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_KodiakIsland *KodiakIslandSession) GELATO() (common.Address, error) {
	return _KodiakIsland.Contract.GELATO(&_KodiakIsland.CallOpts)
}

// GELATO is a free data retrieval call binding the contract method 0xeff557a7.
//
// Solidity: function GELATO() view returns(address)
func (_KodiakIsland *KodiakIslandCallerSession) GELATO() (common.Address, error) {
	return _KodiakIsland.Contract.GELATO(&_KodiakIsland.CallOpts)
}

// RESTRICTEDMINTENABLED is a free data retrieval call binding the contract method 0xd839fb3f.
//
// Solidity: function RESTRICTED_MINT_ENABLED() view returns(uint16)
func (_KodiakIsland *KodiakIslandCaller) RESTRICTEDMINTENABLED(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "RESTRICTED_MINT_ENABLED")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RESTRICTEDMINTENABLED is a free data retrieval call binding the contract method 0xd839fb3f.
//
// Solidity: function RESTRICTED_MINT_ENABLED() view returns(uint16)
func (_KodiakIsland *KodiakIslandSession) RESTRICTEDMINTENABLED() (uint16, error) {
	return _KodiakIsland.Contract.RESTRICTEDMINTENABLED(&_KodiakIsland.CallOpts)
}

// RESTRICTEDMINTENABLED is a free data retrieval call binding the contract method 0xd839fb3f.
//
// Solidity: function RESTRICTED_MINT_ENABLED() view returns(uint16)
func (_KodiakIsland *KodiakIslandCallerSession) RESTRICTEDMINTENABLED() (uint16, error) {
	return _KodiakIsland.Contract.RESTRICTEDMINTENABLED(&_KodiakIsland.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KodiakIsland *KodiakIslandCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KodiakIsland *KodiakIslandSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KodiakIsland.Contract.Allowance(&_KodiakIsland.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KodiakIsland *KodiakIslandCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KodiakIsland.Contract.Allowance(&_KodiakIsland.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KodiakIsland *KodiakIslandCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KodiakIsland *KodiakIslandSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KodiakIsland.Contract.BalanceOf(&_KodiakIsland.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KodiakIsland *KodiakIslandCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KodiakIsland.Contract.BalanceOf(&_KodiakIsland.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KodiakIsland *KodiakIslandCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KodiakIsland *KodiakIslandSession) Decimals() (uint8, error) {
	return _KodiakIsland.Contract.Decimals(&_KodiakIsland.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KodiakIsland *KodiakIslandCallerSession) Decimals() (uint8, error) {
	return _KodiakIsland.Contract.Decimals(&_KodiakIsland.CallOpts)
}

// GelatoRebalanceBPS is a free data retrieval call binding the contract method 0xa50b1fe7.
//
// Solidity: function gelatoRebalanceBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandCaller) GelatoRebalanceBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "gelatoRebalanceBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GelatoRebalanceBPS is a free data retrieval call binding the contract method 0xa50b1fe7.
//
// Solidity: function gelatoRebalanceBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandSession) GelatoRebalanceBPS() (uint16, error) {
	return _KodiakIsland.Contract.GelatoRebalanceBPS(&_KodiakIsland.CallOpts)
}

// GelatoRebalanceBPS is a free data retrieval call binding the contract method 0xa50b1fe7.
//
// Solidity: function gelatoRebalanceBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandCallerSession) GelatoRebalanceBPS() (uint16, error) {
	return _KodiakIsland.Contract.GelatoRebalanceBPS(&_KodiakIsland.CallOpts)
}

// GelatoSlippageBPS is a free data retrieval call binding the contract method 0xd6e7ff39.
//
// Solidity: function gelatoSlippageBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandCaller) GelatoSlippageBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "gelatoSlippageBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GelatoSlippageBPS is a free data retrieval call binding the contract method 0xd6e7ff39.
//
// Solidity: function gelatoSlippageBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandSession) GelatoSlippageBPS() (uint16, error) {
	return _KodiakIsland.Contract.GelatoSlippageBPS(&_KodiakIsland.CallOpts)
}

// GelatoSlippageBPS is a free data retrieval call binding the contract method 0xd6e7ff39.
//
// Solidity: function gelatoSlippageBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandCallerSession) GelatoSlippageBPS() (uint16, error) {
	return _KodiakIsland.Contract.GelatoSlippageBPS(&_KodiakIsland.CallOpts)
}

// GelatoSlippageInterval is a free data retrieval call binding the contract method 0x672152bd.
//
// Solidity: function gelatoSlippageInterval() view returns(uint32)
func (_KodiakIsland *KodiakIslandCaller) GelatoSlippageInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "gelatoSlippageInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GelatoSlippageInterval is a free data retrieval call binding the contract method 0x672152bd.
//
// Solidity: function gelatoSlippageInterval() view returns(uint32)
func (_KodiakIsland *KodiakIslandSession) GelatoSlippageInterval() (uint32, error) {
	return _KodiakIsland.Contract.GelatoSlippageInterval(&_KodiakIsland.CallOpts)
}

// GelatoSlippageInterval is a free data retrieval call binding the contract method 0x672152bd.
//
// Solidity: function gelatoSlippageInterval() view returns(uint32)
func (_KodiakIsland *KodiakIslandCallerSession) GelatoSlippageInterval() (uint32, error) {
	return _KodiakIsland.Contract.GelatoSlippageInterval(&_KodiakIsland.CallOpts)
}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_KodiakIsland *KodiakIslandCaller) GetMintAmounts(opts *bind.CallOpts, amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "getMintAmounts", amount0Max, amount1Max)

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
func (_KodiakIsland *KodiakIslandSession) GetMintAmounts(amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	return _KodiakIsland.Contract.GetMintAmounts(&_KodiakIsland.CallOpts, amount0Max, amount1Max)
}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_KodiakIsland *KodiakIslandCallerSession) GetMintAmounts(amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	return _KodiakIsland.Contract.GetMintAmounts(&_KodiakIsland.CallOpts, amount0Max, amount1Max)
}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_KodiakIsland *KodiakIslandCaller) GetPositionID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "getPositionID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_KodiakIsland *KodiakIslandSession) GetPositionID() ([32]byte, error) {
	return _KodiakIsland.Contract.GetPositionID(&_KodiakIsland.CallOpts)
}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_KodiakIsland *KodiakIslandCallerSession) GetPositionID() ([32]byte, error) {
	return _KodiakIsland.Contract.GetPositionID(&_KodiakIsland.CallOpts)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_KodiakIsland *KodiakIslandCaller) GetUnderlyingBalances(opts *bind.CallOpts) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "getUnderlyingBalances")

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
func (_KodiakIsland *KodiakIslandSession) GetUnderlyingBalances() (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _KodiakIsland.Contract.GetUnderlyingBalances(&_KodiakIsland.CallOpts)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_KodiakIsland *KodiakIslandCallerSession) GetUnderlyingBalances() (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _KodiakIsland.Contract.GetUnderlyingBalances(&_KodiakIsland.CallOpts)
}

// GetUnderlyingBalancesAtPrice is a free data retrieval call binding the contract method 0xb670ed7d.
//
// Solidity: function getUnderlyingBalancesAtPrice(uint160 sqrtRatioX96) view returns(uint256 amount0Current, uint256 amount1Current)
func (_KodiakIsland *KodiakIslandCaller) GetUnderlyingBalancesAtPrice(opts *bind.CallOpts, sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "getUnderlyingBalancesAtPrice", sqrtRatioX96)

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
func (_KodiakIsland *KodiakIslandSession) GetUnderlyingBalancesAtPrice(sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _KodiakIsland.Contract.GetUnderlyingBalancesAtPrice(&_KodiakIsland.CallOpts, sqrtRatioX96)
}

// GetUnderlyingBalancesAtPrice is a free data retrieval call binding the contract method 0xb670ed7d.
//
// Solidity: function getUnderlyingBalancesAtPrice(uint160 sqrtRatioX96) view returns(uint256 amount0Current, uint256 amount1Current)
func (_KodiakIsland *KodiakIslandCallerSession) GetUnderlyingBalancesAtPrice(sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _KodiakIsland.Contract.GetUnderlyingBalancesAtPrice(&_KodiakIsland.CallOpts, sqrtRatioX96)
}

// KodiakBalance0 is a free data retrieval call binding the contract method 0xd8d24145.
//
// Solidity: function kodiakBalance0() view returns(uint256)
func (_KodiakIsland *KodiakIslandCaller) KodiakBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "kodiakBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KodiakBalance0 is a free data retrieval call binding the contract method 0xd8d24145.
//
// Solidity: function kodiakBalance0() view returns(uint256)
func (_KodiakIsland *KodiakIslandSession) KodiakBalance0() (*big.Int, error) {
	return _KodiakIsland.Contract.KodiakBalance0(&_KodiakIsland.CallOpts)
}

// KodiakBalance0 is a free data retrieval call binding the contract method 0xd8d24145.
//
// Solidity: function kodiakBalance0() view returns(uint256)
func (_KodiakIsland *KodiakIslandCallerSession) KodiakBalance0() (*big.Int, error) {
	return _KodiakIsland.Contract.KodiakBalance0(&_KodiakIsland.CallOpts)
}

// KodiakBalance1 is a free data retrieval call binding the contract method 0xf5338bcf.
//
// Solidity: function kodiakBalance1() view returns(uint256)
func (_KodiakIsland *KodiakIslandCaller) KodiakBalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "kodiakBalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KodiakBalance1 is a free data retrieval call binding the contract method 0xf5338bcf.
//
// Solidity: function kodiakBalance1() view returns(uint256)
func (_KodiakIsland *KodiakIslandSession) KodiakBalance1() (*big.Int, error) {
	return _KodiakIsland.Contract.KodiakBalance1(&_KodiakIsland.CallOpts)
}

// KodiakBalance1 is a free data retrieval call binding the contract method 0xf5338bcf.
//
// Solidity: function kodiakBalance1() view returns(uint256)
func (_KodiakIsland *KodiakIslandCallerSession) KodiakBalance1() (*big.Int, error) {
	return _KodiakIsland.Contract.KodiakBalance1(&_KodiakIsland.CallOpts)
}

// KodiakFeeBPS is a free data retrieval call binding the contract method 0xce311070.
//
// Solidity: function kodiakFeeBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandCaller) KodiakFeeBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "kodiakFeeBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// KodiakFeeBPS is a free data retrieval call binding the contract method 0xce311070.
//
// Solidity: function kodiakFeeBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandSession) KodiakFeeBPS() (uint16, error) {
	return _KodiakIsland.Contract.KodiakFeeBPS(&_KodiakIsland.CallOpts)
}

// KodiakFeeBPS is a free data retrieval call binding the contract method 0xce311070.
//
// Solidity: function kodiakFeeBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandCallerSession) KodiakFeeBPS() (uint16, error) {
	return _KodiakIsland.Contract.KodiakFeeBPS(&_KodiakIsland.CallOpts)
}

// KodiakTreasury is a free data retrieval call binding the contract method 0x94a113ec.
//
// Solidity: function kodiakTreasury() view returns(address)
func (_KodiakIsland *KodiakIslandCaller) KodiakTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "kodiakTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KodiakTreasury is a free data retrieval call binding the contract method 0x94a113ec.
//
// Solidity: function kodiakTreasury() view returns(address)
func (_KodiakIsland *KodiakIslandSession) KodiakTreasury() (common.Address, error) {
	return _KodiakIsland.Contract.KodiakTreasury(&_KodiakIsland.CallOpts)
}

// KodiakTreasury is a free data retrieval call binding the contract method 0x94a113ec.
//
// Solidity: function kodiakTreasury() view returns(address)
func (_KodiakIsland *KodiakIslandCallerSession) KodiakTreasury() (common.Address, error) {
	return _KodiakIsland.Contract.KodiakTreasury(&_KodiakIsland.CallOpts)
}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_KodiakIsland *KodiakIslandCaller) LowerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "lowerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_KodiakIsland *KodiakIslandSession) LowerTick() (*big.Int, error) {
	return _KodiakIsland.Contract.LowerTick(&_KodiakIsland.CallOpts)
}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_KodiakIsland *KodiakIslandCallerSession) LowerTick() (*big.Int, error) {
	return _KodiakIsland.Contract.LowerTick(&_KodiakIsland.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_KodiakIsland *KodiakIslandCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_KodiakIsland *KodiakIslandSession) Manager() (common.Address, error) {
	return _KodiakIsland.Contract.Manager(&_KodiakIsland.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_KodiakIsland *KodiakIslandCallerSession) Manager() (common.Address, error) {
	return _KodiakIsland.Contract.Manager(&_KodiakIsland.CallOpts)
}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_KodiakIsland *KodiakIslandCaller) ManagerBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "managerBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_KodiakIsland *KodiakIslandSession) ManagerBalance0() (*big.Int, error) {
	return _KodiakIsland.Contract.ManagerBalance0(&_KodiakIsland.CallOpts)
}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_KodiakIsland *KodiakIslandCallerSession) ManagerBalance0() (*big.Int, error) {
	return _KodiakIsland.Contract.ManagerBalance0(&_KodiakIsland.CallOpts)
}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_KodiakIsland *KodiakIslandCaller) ManagerBalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "managerBalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_KodiakIsland *KodiakIslandSession) ManagerBalance1() (*big.Int, error) {
	return _KodiakIsland.Contract.ManagerBalance1(&_KodiakIsland.CallOpts)
}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_KodiakIsland *KodiakIslandCallerSession) ManagerBalance1() (*big.Int, error) {
	return _KodiakIsland.Contract.ManagerBalance1(&_KodiakIsland.CallOpts)
}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandCaller) ManagerFeeBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "managerFeeBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandSession) ManagerFeeBPS() (uint16, error) {
	return _KodiakIsland.Contract.ManagerFeeBPS(&_KodiakIsland.CallOpts)
}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_KodiakIsland *KodiakIslandCallerSession) ManagerFeeBPS() (uint16, error) {
	return _KodiakIsland.Contract.ManagerFeeBPS(&_KodiakIsland.CallOpts)
}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_KodiakIsland *KodiakIslandCaller) ManagerTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "managerTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_KodiakIsland *KodiakIslandSession) ManagerTreasury() (common.Address, error) {
	return _KodiakIsland.Contract.ManagerTreasury(&_KodiakIsland.CallOpts)
}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_KodiakIsland *KodiakIslandCallerSession) ManagerTreasury() (common.Address, error) {
	return _KodiakIsland.Contract.ManagerTreasury(&_KodiakIsland.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KodiakIsland *KodiakIslandCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KodiakIsland *KodiakIslandSession) Name() (string, error) {
	return _KodiakIsland.Contract.Name(&_KodiakIsland.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KodiakIsland *KodiakIslandCallerSession) Name() (string, error) {
	return _KodiakIsland.Contract.Name(&_KodiakIsland.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_KodiakIsland *KodiakIslandCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_KodiakIsland *KodiakIslandSession) Pool() (common.Address, error) {
	return _KodiakIsland.Contract.Pool(&_KodiakIsland.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_KodiakIsland *KodiakIslandCallerSession) Pool() (common.Address, error) {
	return _KodiakIsland.Contract.Pool(&_KodiakIsland.CallOpts)
}

// RestrictedMintToggle is a free data retrieval call binding the contract method 0xdd40e322.
//
// Solidity: function restrictedMintToggle() view returns(uint16)
func (_KodiakIsland *KodiakIslandCaller) RestrictedMintToggle(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "restrictedMintToggle")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RestrictedMintToggle is a free data retrieval call binding the contract method 0xdd40e322.
//
// Solidity: function restrictedMintToggle() view returns(uint16)
func (_KodiakIsland *KodiakIslandSession) RestrictedMintToggle() (uint16, error) {
	return _KodiakIsland.Contract.RestrictedMintToggle(&_KodiakIsland.CallOpts)
}

// RestrictedMintToggle is a free data retrieval call binding the contract method 0xdd40e322.
//
// Solidity: function restrictedMintToggle() view returns(uint16)
func (_KodiakIsland *KodiakIslandCallerSession) RestrictedMintToggle() (uint16, error) {
	return _KodiakIsland.Contract.RestrictedMintToggle(&_KodiakIsland.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KodiakIsland *KodiakIslandCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KodiakIsland *KodiakIslandSession) Symbol() (string, error) {
	return _KodiakIsland.Contract.Symbol(&_KodiakIsland.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KodiakIsland *KodiakIslandCallerSession) Symbol() (string, error) {
	return _KodiakIsland.Contract.Symbol(&_KodiakIsland.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KodiakIsland *KodiakIslandCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KodiakIsland *KodiakIslandSession) Token0() (common.Address, error) {
	return _KodiakIsland.Contract.Token0(&_KodiakIsland.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_KodiakIsland *KodiakIslandCallerSession) Token0() (common.Address, error) {
	return _KodiakIsland.Contract.Token0(&_KodiakIsland.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KodiakIsland *KodiakIslandCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KodiakIsland *KodiakIslandSession) Token1() (common.Address, error) {
	return _KodiakIsland.Contract.Token1(&_KodiakIsland.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_KodiakIsland *KodiakIslandCallerSession) Token1() (common.Address, error) {
	return _KodiakIsland.Contract.Token1(&_KodiakIsland.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KodiakIsland *KodiakIslandCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KodiakIsland *KodiakIslandSession) TotalSupply() (*big.Int, error) {
	return _KodiakIsland.Contract.TotalSupply(&_KodiakIsland.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KodiakIsland *KodiakIslandCallerSession) TotalSupply() (*big.Int, error) {
	return _KodiakIsland.Contract.TotalSupply(&_KodiakIsland.CallOpts)
}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_KodiakIsland *KodiakIslandCaller) UpperTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "upperTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_KodiakIsland *KodiakIslandSession) UpperTick() (*big.Int, error) {
	return _KodiakIsland.Contract.UpperTick(&_KodiakIsland.CallOpts)
}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_KodiakIsland *KodiakIslandCallerSession) UpperTick() (*big.Int, error) {
	return _KodiakIsland.Contract.UpperTick(&_KodiakIsland.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KodiakIsland *KodiakIslandCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KodiakIsland.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KodiakIsland *KodiakIslandSession) Version() (string, error) {
	return _KodiakIsland.Contract.Version(&_KodiakIsland.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_KodiakIsland *KodiakIslandCallerSession) Version() (string, error) {
	return _KodiakIsland.Contract.Version(&_KodiakIsland.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KodiakIsland *KodiakIslandTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KodiakIsland *KodiakIslandSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Approve(&_KodiakIsland.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KodiakIsland *KodiakIslandTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Approve(&_KodiakIsland.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_KodiakIsland *KodiakIslandTransactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "burn", burnAmount, receiver)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_KodiakIsland *KodiakIslandSession) Burn(burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Burn(&_KodiakIsland.TransactOpts, burnAmount, receiver)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_KodiakIsland *KodiakIslandTransactorSession) Burn(burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Burn(&_KodiakIsland.TransactOpts, burnAmount, receiver)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KodiakIsland *KodiakIslandTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KodiakIsland *KodiakIslandSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.Contract.DecreaseAllowance(&_KodiakIsland.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KodiakIsland *KodiakIslandTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.Contract.DecreaseAllowance(&_KodiakIsland.TransactOpts, spender, subtractedValue)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_KodiakIsland *KodiakIslandTransactor) ExecutiveRebalance(opts *bind.TransactOpts, newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "executiveRebalance", newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_KodiakIsland *KodiakIslandSession) ExecutiveRebalance(newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _KodiakIsland.Contract.ExecutiveRebalance(&_KodiakIsland.TransactOpts, newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_KodiakIsland *KodiakIslandTransactorSession) ExecutiveRebalance(newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _KodiakIsland.Contract.ExecutiveRebalance(&_KodiakIsland.TransactOpts, newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KodiakIsland *KodiakIslandTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KodiakIsland *KodiakIslandSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.Contract.IncreaseAllowance(&_KodiakIsland.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KodiakIsland *KodiakIslandTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.Contract.IncreaseAllowance(&_KodiakIsland.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xe25e15e3.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_) returns()
func (_KodiakIsland *KodiakIslandTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "initialize", _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe25e15e3.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_) returns()
func (_KodiakIsland *KodiakIslandSession) Initialize(_name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Initialize(&_KodiakIsland.TransactOpts, _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_)
}

// Initialize is a paid mutator transaction binding the contract method 0xe25e15e3.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_) returns()
func (_KodiakIsland *KodiakIslandTransactorSession) Initialize(_name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Initialize(&_KodiakIsland.TransactOpts, _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_KodiakIsland *KodiakIslandTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "mint", mintAmount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_KodiakIsland *KodiakIslandSession) Mint(mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Mint(&_KodiakIsland.TransactOpts, mintAmount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_KodiakIsland *KodiakIslandTransactorSession) Mint(mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Mint(&_KodiakIsland.TransactOpts, mintAmount, receiver)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb135c99f.
//
// Solidity: function rebalance(uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne, uint256 feeAmount, address paymentToken) returns()
func (_KodiakIsland *KodiakIslandTransactor) Rebalance(opts *bind.TransactOpts, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool, feeAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "rebalance", swapThresholdPrice, swapAmountBPS, zeroForOne, feeAmount, paymentToken)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb135c99f.
//
// Solidity: function rebalance(uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne, uint256 feeAmount, address paymentToken) returns()
func (_KodiakIsland *KodiakIslandSession) Rebalance(swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool, feeAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Rebalance(&_KodiakIsland.TransactOpts, swapThresholdPrice, swapAmountBPS, zeroForOne, feeAmount, paymentToken)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb135c99f.
//
// Solidity: function rebalance(uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne, uint256 feeAmount, address paymentToken) returns()
func (_KodiakIsland *KodiakIslandTransactorSession) Rebalance(swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool, feeAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Rebalance(&_KodiakIsland.TransactOpts, swapThresholdPrice, swapAmountBPS, zeroForOne, feeAmount, paymentToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KodiakIsland *KodiakIslandTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KodiakIsland *KodiakIslandSession) RenounceOwnership() (*types.Transaction, error) {
	return _KodiakIsland.Contract.RenounceOwnership(&_KodiakIsland.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KodiakIsland *KodiakIslandTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KodiakIsland.Contract.RenounceOwnership(&_KodiakIsland.TransactOpts)
}

// ToggleRestrictMint is a paid mutator transaction binding the contract method 0x5f59dce5.
//
// Solidity: function toggleRestrictMint() returns()
func (_KodiakIsland *KodiakIslandTransactor) ToggleRestrictMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "toggleRestrictMint")
}

// ToggleRestrictMint is a paid mutator transaction binding the contract method 0x5f59dce5.
//
// Solidity: function toggleRestrictMint() returns()
func (_KodiakIsland *KodiakIslandSession) ToggleRestrictMint() (*types.Transaction, error) {
	return _KodiakIsland.Contract.ToggleRestrictMint(&_KodiakIsland.TransactOpts)
}

// ToggleRestrictMint is a paid mutator transaction binding the contract method 0x5f59dce5.
//
// Solidity: function toggleRestrictMint() returns()
func (_KodiakIsland *KodiakIslandTransactorSession) ToggleRestrictMint() (*types.Transaction, error) {
	return _KodiakIsland.Contract.ToggleRestrictMint(&_KodiakIsland.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KodiakIsland *KodiakIslandTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KodiakIsland *KodiakIslandSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Transfer(&_KodiakIsland.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_KodiakIsland *KodiakIslandTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.Contract.Transfer(&_KodiakIsland.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KodiakIsland *KodiakIslandTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KodiakIsland *KodiakIslandSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.Contract.TransferFrom(&_KodiakIsland.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_KodiakIsland *KodiakIslandTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KodiakIsland.Contract.TransferFrom(&_KodiakIsland.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KodiakIsland *KodiakIslandTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KodiakIsland *KodiakIslandSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KodiakIsland.Contract.TransferOwnership(&_KodiakIsland.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KodiakIsland *KodiakIslandTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KodiakIsland.Contract.TransferOwnership(&_KodiakIsland.TransactOpts, newOwner)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_KodiakIsland *KodiakIslandTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "uniswapV3MintCallback", amount0Owed, amount1Owed, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_KodiakIsland *KodiakIslandSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakIsland.Contract.UniswapV3MintCallback(&_KodiakIsland.TransactOpts, amount0Owed, amount1Owed, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_KodiakIsland *KodiakIslandTransactorSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakIsland.Contract.UniswapV3MintCallback(&_KodiakIsland.TransactOpts, amount0Owed, amount1Owed, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_KodiakIsland *KodiakIslandTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_KodiakIsland *KodiakIslandSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakIsland.Contract.UniswapV3SwapCallback(&_KodiakIsland.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_KodiakIsland *KodiakIslandTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _KodiakIsland.Contract.UniswapV3SwapCallback(&_KodiakIsland.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UpdateManagerParams is a paid mutator transaction binding the contract method 0xc540e483.
//
// Solidity: function updateManagerParams(int16 newManagerFeeBPS, address newManagerTreasury, int16 newRebalanceBPS, int16 newSlippageBPS, int32 newSlippageInterval) returns()
func (_KodiakIsland *KodiakIslandTransactor) UpdateManagerParams(opts *bind.TransactOpts, newManagerFeeBPS int16, newManagerTreasury common.Address, newRebalanceBPS int16, newSlippageBPS int16, newSlippageInterval int32) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "updateManagerParams", newManagerFeeBPS, newManagerTreasury, newRebalanceBPS, newSlippageBPS, newSlippageInterval)
}

// UpdateManagerParams is a paid mutator transaction binding the contract method 0xc540e483.
//
// Solidity: function updateManagerParams(int16 newManagerFeeBPS, address newManagerTreasury, int16 newRebalanceBPS, int16 newSlippageBPS, int32 newSlippageInterval) returns()
func (_KodiakIsland *KodiakIslandSession) UpdateManagerParams(newManagerFeeBPS int16, newManagerTreasury common.Address, newRebalanceBPS int16, newSlippageBPS int16, newSlippageInterval int32) (*types.Transaction, error) {
	return _KodiakIsland.Contract.UpdateManagerParams(&_KodiakIsland.TransactOpts, newManagerFeeBPS, newManagerTreasury, newRebalanceBPS, newSlippageBPS, newSlippageInterval)
}

// UpdateManagerParams is a paid mutator transaction binding the contract method 0xc540e483.
//
// Solidity: function updateManagerParams(int16 newManagerFeeBPS, address newManagerTreasury, int16 newRebalanceBPS, int16 newSlippageBPS, int32 newSlippageInterval) returns()
func (_KodiakIsland *KodiakIslandTransactorSession) UpdateManagerParams(newManagerFeeBPS int16, newManagerTreasury common.Address, newRebalanceBPS int16, newSlippageBPS int16, newSlippageInterval int32) (*types.Transaction, error) {
	return _KodiakIsland.Contract.UpdateManagerParams(&_KodiakIsland.TransactOpts, newManagerFeeBPS, newManagerTreasury, newRebalanceBPS, newSlippageBPS, newSlippageInterval)
}

// WithdrawKodiakBalance is a paid mutator transaction binding the contract method 0xd83bfb90.
//
// Solidity: function withdrawKodiakBalance() returns()
func (_KodiakIsland *KodiakIslandTransactor) WithdrawKodiakBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "withdrawKodiakBalance")
}

// WithdrawKodiakBalance is a paid mutator transaction binding the contract method 0xd83bfb90.
//
// Solidity: function withdrawKodiakBalance() returns()
func (_KodiakIsland *KodiakIslandSession) WithdrawKodiakBalance() (*types.Transaction, error) {
	return _KodiakIsland.Contract.WithdrawKodiakBalance(&_KodiakIsland.TransactOpts)
}

// WithdrawKodiakBalance is a paid mutator transaction binding the contract method 0xd83bfb90.
//
// Solidity: function withdrawKodiakBalance() returns()
func (_KodiakIsland *KodiakIslandTransactorSession) WithdrawKodiakBalance() (*types.Transaction, error) {
	return _KodiakIsland.Contract.WithdrawKodiakBalance(&_KodiakIsland.TransactOpts)
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x7ecd6717.
//
// Solidity: function withdrawManagerBalance() returns()
func (_KodiakIsland *KodiakIslandTransactor) WithdrawManagerBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KodiakIsland.contract.Transact(opts, "withdrawManagerBalance")
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x7ecd6717.
//
// Solidity: function withdrawManagerBalance() returns()
func (_KodiakIsland *KodiakIslandSession) WithdrawManagerBalance() (*types.Transaction, error) {
	return _KodiakIsland.Contract.WithdrawManagerBalance(&_KodiakIsland.TransactOpts)
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x7ecd6717.
//
// Solidity: function withdrawManagerBalance() returns()
func (_KodiakIsland *KodiakIslandTransactorSession) WithdrawManagerBalance() (*types.Transaction, error) {
	return _KodiakIsland.Contract.WithdrawManagerBalance(&_KodiakIsland.TransactOpts)
}

// KodiakIslandApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KodiakIsland contract.
type KodiakIslandApprovalIterator struct {
	Event *KodiakIslandApproval // Event containing the contract specifics and raw log

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
func (it *KodiakIslandApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakIslandApproval)
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
		it.Event = new(KodiakIslandApproval)
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
func (it *KodiakIslandApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakIslandApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakIslandApproval represents a Approval event raised by the KodiakIsland contract.
type KodiakIslandApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KodiakIsland *KodiakIslandFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KodiakIslandApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KodiakIsland.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KodiakIslandApprovalIterator{contract: _KodiakIsland.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KodiakIsland *KodiakIslandFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KodiakIslandApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KodiakIsland.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakIslandApproval)
				if err := _KodiakIsland.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KodiakIsland *KodiakIslandFilterer) ParseApproval(log types.Log) (*KodiakIslandApproval, error) {
	event := new(KodiakIslandApproval)
	if err := _KodiakIsland.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakIslandBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the KodiakIsland contract.
type KodiakIslandBurnedIterator struct {
	Event *KodiakIslandBurned // Event containing the contract specifics and raw log

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
func (it *KodiakIslandBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakIslandBurned)
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
		it.Event = new(KodiakIslandBurned)
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
func (it *KodiakIslandBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakIslandBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakIslandBurned represents a Burned event raised by the KodiakIsland contract.
type KodiakIslandBurned struct {
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
func (_KodiakIsland *KodiakIslandFilterer) FilterBurned(opts *bind.FilterOpts) (*KodiakIslandBurnedIterator, error) {

	logs, sub, err := _KodiakIsland.contract.FilterLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return &KodiakIslandBurnedIterator{contract: _KodiakIsland.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x7239dff1718b550db7f36cbf69c665cfeb56d0e96b4fb76a5cba712961b65509.
//
// Solidity: event Burned(address receiver, uint256 burnAmount, uint256 amount0Out, uint256 amount1Out, uint128 liquidityBurned)
func (_KodiakIsland *KodiakIslandFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *KodiakIslandBurned) (event.Subscription, error) {

	logs, sub, err := _KodiakIsland.contract.WatchLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakIslandBurned)
				if err := _KodiakIsland.contract.UnpackLog(event, "Burned", log); err != nil {
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
func (_KodiakIsland *KodiakIslandFilterer) ParseBurned(log types.Log) (*KodiakIslandBurned, error) {
	event := new(KodiakIslandBurned)
	if err := _KodiakIsland.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakIslandFeesEarnedIterator is returned from FilterFeesEarned and is used to iterate over the raw logs and unpacked data for FeesEarned events raised by the KodiakIsland contract.
type KodiakIslandFeesEarnedIterator struct {
	Event *KodiakIslandFeesEarned // Event containing the contract specifics and raw log

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
func (it *KodiakIslandFeesEarnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakIslandFeesEarned)
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
		it.Event = new(KodiakIslandFeesEarned)
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
func (it *KodiakIslandFeesEarnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakIslandFeesEarnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakIslandFeesEarned represents a FeesEarned event raised by the KodiakIsland contract.
type KodiakIslandFeesEarned struct {
	FeesEarned0 *big.Int
	FeesEarned1 *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeesEarned is a free log retrieval operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 feesEarned0, uint256 feesEarned1)
func (_KodiakIsland *KodiakIslandFilterer) FilterFeesEarned(opts *bind.FilterOpts) (*KodiakIslandFeesEarnedIterator, error) {

	logs, sub, err := _KodiakIsland.contract.FilterLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return &KodiakIslandFeesEarnedIterator{contract: _KodiakIsland.contract, event: "FeesEarned", logs: logs, sub: sub}, nil
}

// WatchFeesEarned is a free log subscription operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 feesEarned0, uint256 feesEarned1)
func (_KodiakIsland *KodiakIslandFilterer) WatchFeesEarned(opts *bind.WatchOpts, sink chan<- *KodiakIslandFeesEarned) (event.Subscription, error) {

	logs, sub, err := _KodiakIsland.contract.WatchLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakIslandFeesEarned)
				if err := _KodiakIsland.contract.UnpackLog(event, "FeesEarned", log); err != nil {
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
func (_KodiakIsland *KodiakIslandFilterer) ParseFeesEarned(log types.Log) (*KodiakIslandFeesEarned, error) {
	event := new(KodiakIslandFeesEarned)
	if err := _KodiakIsland.contract.UnpackLog(event, "FeesEarned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakIslandMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the KodiakIsland contract.
type KodiakIslandMintedIterator struct {
	Event *KodiakIslandMinted // Event containing the contract specifics and raw log

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
func (it *KodiakIslandMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakIslandMinted)
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
		it.Event = new(KodiakIslandMinted)
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
func (it *KodiakIslandMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakIslandMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakIslandMinted represents a Minted event raised by the KodiakIsland contract.
type KodiakIslandMinted struct {
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
func (_KodiakIsland *KodiakIslandFilterer) FilterMinted(opts *bind.FilterOpts) (*KodiakIslandMintedIterator, error) {

	logs, sub, err := _KodiakIsland.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &KodiakIslandMintedIterator{contract: _KodiakIsland.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x55801cfe493000b734571da1694b21e7f66b11e8ce9fdaa0524ecb59105e73e7.
//
// Solidity: event Minted(address receiver, uint256 mintAmount, uint256 amount0In, uint256 amount1In, uint128 liquidityMinted)
func (_KodiakIsland *KodiakIslandFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *KodiakIslandMinted) (event.Subscription, error) {

	logs, sub, err := _KodiakIsland.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakIslandMinted)
				if err := _KodiakIsland.contract.UnpackLog(event, "Minted", log); err != nil {
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
func (_KodiakIsland *KodiakIslandFilterer) ParseMinted(log types.Log) (*KodiakIslandMinted, error) {
	event := new(KodiakIslandMinted)
	if err := _KodiakIsland.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakIslandOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KodiakIsland contract.
type KodiakIslandOwnershipTransferredIterator struct {
	Event *KodiakIslandOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KodiakIslandOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakIslandOwnershipTransferred)
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
		it.Event = new(KodiakIslandOwnershipTransferred)
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
func (it *KodiakIslandOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakIslandOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakIslandOwnershipTransferred represents a OwnershipTransferred event raised by the KodiakIsland contract.
type KodiakIslandOwnershipTransferred struct {
	PreviousManager common.Address
	NewManager      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousManager, address indexed newManager)
func (_KodiakIsland *KodiakIslandFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousManager []common.Address, newManager []common.Address) (*KodiakIslandOwnershipTransferredIterator, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _KodiakIsland.contract.FilterLogs(opts, "OwnershipTransferred", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &KodiakIslandOwnershipTransferredIterator{contract: _KodiakIsland.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousManager, address indexed newManager)
func (_KodiakIsland *KodiakIslandFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KodiakIslandOwnershipTransferred, previousManager []common.Address, newManager []common.Address) (event.Subscription, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _KodiakIsland.contract.WatchLogs(opts, "OwnershipTransferred", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakIslandOwnershipTransferred)
				if err := _KodiakIsland.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KodiakIsland *KodiakIslandFilterer) ParseOwnershipTransferred(log types.Log) (*KodiakIslandOwnershipTransferred, error) {
	event := new(KodiakIslandOwnershipTransferred)
	if err := _KodiakIsland.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakIslandRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the KodiakIsland contract.
type KodiakIslandRebalanceIterator struct {
	Event *KodiakIslandRebalance // Event containing the contract specifics and raw log

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
func (it *KodiakIslandRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakIslandRebalance)
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
		it.Event = new(KodiakIslandRebalance)
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
func (it *KodiakIslandRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakIslandRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakIslandRebalance represents a Rebalance event raised by the KodiakIsland contract.
type KodiakIslandRebalance struct {
	LowerTick       *big.Int
	UpperTick       *big.Int
	LiquidityBefore *big.Int
	LiquidityAfter  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRebalance is a free log retrieval operation binding the contract event 0xc749f9ae947d4734cf1569606a8a347391ae94a063478aa853aeff48ac5f99e8.
//
// Solidity: event Rebalance(int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_KodiakIsland *KodiakIslandFilterer) FilterRebalance(opts *bind.FilterOpts) (*KodiakIslandRebalanceIterator, error) {

	logs, sub, err := _KodiakIsland.contract.FilterLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return &KodiakIslandRebalanceIterator{contract: _KodiakIsland.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0xc749f9ae947d4734cf1569606a8a347391ae94a063478aa853aeff48ac5f99e8.
//
// Solidity: event Rebalance(int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_KodiakIsland *KodiakIslandFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *KodiakIslandRebalance) (event.Subscription, error) {

	logs, sub, err := _KodiakIsland.contract.WatchLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakIslandRebalance)
				if err := _KodiakIsland.contract.UnpackLog(event, "Rebalance", log); err != nil {
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
func (_KodiakIsland *KodiakIslandFilterer) ParseRebalance(log types.Log) (*KodiakIslandRebalance, error) {
	event := new(KodiakIslandRebalance)
	if err := _KodiakIsland.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakIslandTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KodiakIsland contract.
type KodiakIslandTransferIterator struct {
	Event *KodiakIslandTransfer // Event containing the contract specifics and raw log

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
func (it *KodiakIslandTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakIslandTransfer)
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
		it.Event = new(KodiakIslandTransfer)
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
func (it *KodiakIslandTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakIslandTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakIslandTransfer represents a Transfer event raised by the KodiakIsland contract.
type KodiakIslandTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KodiakIsland *KodiakIslandFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KodiakIslandTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KodiakIsland.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KodiakIslandTransferIterator{contract: _KodiakIsland.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KodiakIsland *KodiakIslandFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KodiakIslandTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KodiakIsland.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakIslandTransfer)
				if err := _KodiakIsland.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KodiakIsland *KodiakIslandFilterer) ParseTransfer(log types.Log) (*KodiakIslandTransfer, error) {
	event := new(KodiakIslandTransfer)
	if err := _KodiakIsland.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KodiakIslandUpdateManagerParamsIterator is returned from FilterUpdateManagerParams and is used to iterate over the raw logs and unpacked data for UpdateManagerParams events raised by the KodiakIsland contract.
type KodiakIslandUpdateManagerParamsIterator struct {
	Event *KodiakIslandUpdateManagerParams // Event containing the contract specifics and raw log

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
func (it *KodiakIslandUpdateManagerParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KodiakIslandUpdateManagerParams)
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
		it.Event = new(KodiakIslandUpdateManagerParams)
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
func (it *KodiakIslandUpdateManagerParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KodiakIslandUpdateManagerParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KodiakIslandUpdateManagerParams represents a UpdateManagerParams event raised by the KodiakIsland contract.
type KodiakIslandUpdateManagerParams struct {
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
func (_KodiakIsland *KodiakIslandFilterer) FilterUpdateManagerParams(opts *bind.FilterOpts) (*KodiakIslandUpdateManagerParamsIterator, error) {

	logs, sub, err := _KodiakIsland.contract.FilterLogs(opts, "UpdateManagerParams")
	if err != nil {
		return nil, err
	}
	return &KodiakIslandUpdateManagerParamsIterator{contract: _KodiakIsland.contract, event: "UpdateManagerParams", logs: logs, sub: sub}, nil
}

// WatchUpdateManagerParams is a free log subscription operation binding the contract event 0xf459b381c988c676562c20a01f42f488c5560ed2e3957b0f27f26573418b6193.
//
// Solidity: event UpdateManagerParams(uint16 managerFeeBPS, address managerTreasury, uint16 gelatoRebalanceBPS, uint16 gelatoSlippageBPS, uint32 gelatoSlippageInterval)
func (_KodiakIsland *KodiakIslandFilterer) WatchUpdateManagerParams(opts *bind.WatchOpts, sink chan<- *KodiakIslandUpdateManagerParams) (event.Subscription, error) {

	logs, sub, err := _KodiakIsland.contract.WatchLogs(opts, "UpdateManagerParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KodiakIslandUpdateManagerParams)
				if err := _KodiakIsland.contract.UnpackLog(event, "UpdateManagerParams", log); err != nil {
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
func (_KodiakIsland *KodiakIslandFilterer) ParseUpdateManagerParams(log types.Log) (*KodiakIslandUpdateManagerParams, error) {
	event := new(KodiakIslandUpdateManagerParams)
	if err := _KodiakIsland.contract.UnpackLog(event, "UpdateManagerParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
