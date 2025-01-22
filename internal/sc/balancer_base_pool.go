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

// IPoolSwapStructsSwapRequest is an auto generated low-level Go binding around an user-defined struct.
type IPoolSwapStructsSwapRequest struct {
	Kind            uint8
	TokenIn         common.Address
	TokenOut        common.Address
	Amount          *big.Int
	PoolId          [32]byte
	LastChangeBlock *big.Int
	From            common.Address
	To              common.Address
	UserData        []byte
}

// WeightedPoolNewPoolParams is an auto generated low-level Go binding around an user-defined struct.
type WeightedPoolNewPoolParams struct {
	Name              string
	Symbol            string
	Tokens            []common.Address
	NormalizedWeights []*big.Int
	RateProviders     []common.Address
	AssetManagers     []common.Address
	SwapFeePercentage *big.Int
}

// BalancerBasePoolMetaData contains all meta data concerning the BalancerBasePool contract.
var BalancerBasePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"normalizedWeights\",\"type\":\"uint256[]\"},{\"internalType\":\"contractIRateProvider[]\",\"name\":\"rateProviders\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"assetManagers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"internalType\":\"structWeightedPool.NewPoolParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"contractIVault\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"contractIProtocolFeePercentagesProvider\",\"name\":\"protocolFeeProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeePercentage\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeePercentageCacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"RecoveryModeStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableRecoveryMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getATHRateProduct\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"getActionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActualSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInvariant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastPostJoinExitInvariant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNormalizedWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferPeriodEndTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"}],\"name\":\"getProtocolFeePercentageCache\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeesCollector\",\"outputs\":[{\"internalType\":\"contractIProtocolFeesCollector\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolSwapFeeDelegation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRateProviders\",\"outputs\":[{\"internalType\":\"contractIRateProvider[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getScalingFactors\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inRecoveryMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onExitPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"onJoinPool\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIVault.SwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIPoolSwapStructs.SwapRequest\",\"name\":\"request\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceTokenOut\",\"type\":\"uint256\"}],\"name\":\"onSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"lastChangeBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"queryJoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"poolConfig\",\"type\":\"bytes\"}],\"name\":\"setAssetManagerPoolConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setSwapFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateProtocolFeePercentageCache\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BalancerBasePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use BalancerBasePoolMetaData.ABI instead.
var BalancerBasePoolABI = BalancerBasePoolMetaData.ABI

// BalancerBasePool is an auto generated Go binding around an Ethereum contract.
type BalancerBasePool struct {
	BalancerBasePoolCaller     // Read-only binding to the contract
	BalancerBasePoolTransactor // Write-only binding to the contract
	BalancerBasePoolFilterer   // Log filterer for contract events
}

// BalancerBasePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerBasePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerBasePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerBasePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerBasePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerBasePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerBasePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerBasePoolSession struct {
	Contract     *BalancerBasePool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancerBasePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerBasePoolCallerSession struct {
	Contract *BalancerBasePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BalancerBasePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerBasePoolTransactorSession struct {
	Contract     *BalancerBasePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BalancerBasePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerBasePoolRaw struct {
	Contract *BalancerBasePool // Generic contract binding to access the raw methods on
}

// BalancerBasePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerBasePoolCallerRaw struct {
	Contract *BalancerBasePoolCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerBasePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerBasePoolTransactorRaw struct {
	Contract *BalancerBasePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerBasePool creates a new instance of BalancerBasePool, bound to a specific deployed contract.
func NewBalancerBasePool(address common.Address, backend bind.ContractBackend) (*BalancerBasePool, error) {
	contract, err := bindBalancerBasePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerBasePool{BalancerBasePoolCaller: BalancerBasePoolCaller{contract: contract}, BalancerBasePoolTransactor: BalancerBasePoolTransactor{contract: contract}, BalancerBasePoolFilterer: BalancerBasePoolFilterer{contract: contract}}, nil
}

// NewBalancerBasePoolCaller creates a new read-only instance of BalancerBasePool, bound to a specific deployed contract.
func NewBalancerBasePoolCaller(address common.Address, caller bind.ContractCaller) (*BalancerBasePoolCaller, error) {
	contract, err := bindBalancerBasePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerBasePoolCaller{contract: contract}, nil
}

// NewBalancerBasePoolTransactor creates a new write-only instance of BalancerBasePool, bound to a specific deployed contract.
func NewBalancerBasePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerBasePoolTransactor, error) {
	contract, err := bindBalancerBasePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerBasePoolTransactor{contract: contract}, nil
}

// NewBalancerBasePoolFilterer creates a new log filterer instance of BalancerBasePool, bound to a specific deployed contract.
func NewBalancerBasePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerBasePoolFilterer, error) {
	contract, err := bindBalancerBasePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerBasePoolFilterer{contract: contract}, nil
}

// bindBalancerBasePool binds a generic wrapper to an already deployed contract.
func bindBalancerBasePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BalancerBasePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerBasePool *BalancerBasePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerBasePool.Contract.BalancerBasePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerBasePool *BalancerBasePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.BalancerBasePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerBasePool *BalancerBasePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.BalancerBasePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerBasePool *BalancerBasePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BalancerBasePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerBasePool *BalancerBasePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerBasePool *BalancerBasePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.contract.Transact(opts, method, params...)
}

// DELEGATEPROTOCOLSWAPFEESSENTINEL is a free data retrieval call binding the contract method 0xddf4627b.
//
// Solidity: function DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) DELEGATEPROTOCOLSWAPFEESSENTINEL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DELEGATEPROTOCOLSWAPFEESSENTINEL is a free data retrieval call binding the contract method 0xddf4627b.
//
// Solidity: function DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) DELEGATEPROTOCOLSWAPFEESSENTINEL() (*big.Int, error) {
	return _BalancerBasePool.Contract.DELEGATEPROTOCOLSWAPFEESSENTINEL(&_BalancerBasePool.CallOpts)
}

// DELEGATEPROTOCOLSWAPFEESSENTINEL is a free data retrieval call binding the contract method 0xddf4627b.
//
// Solidity: function DELEGATE_PROTOCOL_SWAP_FEES_SENTINEL() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) DELEGATEPROTOCOLSWAPFEESSENTINEL() (*big.Int, error) {
	return _BalancerBasePool.Contract.DELEGATEPROTOCOLSWAPFEESSENTINEL(&_BalancerBasePool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BalancerBasePool.Contract.DOMAINSEPARATOR(&_BalancerBasePool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BalancerBasePool.Contract.DOMAINSEPARATOR(&_BalancerBasePool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BalancerBasePool.Contract.Allowance(&_BalancerBasePool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BalancerBasePool.Contract.Allowance(&_BalancerBasePool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BalancerBasePool.Contract.BalanceOf(&_BalancerBasePool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BalancerBasePool.Contract.BalanceOf(&_BalancerBasePool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerBasePool *BalancerBasePoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerBasePool *BalancerBasePoolSession) Decimals() (uint8, error) {
	return _BalancerBasePool.Contract.Decimals(&_BalancerBasePool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerBasePool *BalancerBasePoolCallerSession) Decimals() (uint8, error) {
	return _BalancerBasePool.Contract.Decimals(&_BalancerBasePool.CallOpts)
}

// GetATHRateProduct is a free data retrieval call binding the contract method 0x23ef89ed.
//
// Solidity: function getATHRateProduct() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) GetATHRateProduct(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getATHRateProduct")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetATHRateProduct is a free data retrieval call binding the contract method 0x23ef89ed.
//
// Solidity: function getATHRateProduct() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) GetATHRateProduct() (*big.Int, error) {
	return _BalancerBasePool.Contract.GetATHRateProduct(&_BalancerBasePool.CallOpts)
}

// GetATHRateProduct is a free data retrieval call binding the contract method 0x23ef89ed.
//
// Solidity: function getATHRateProduct() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetATHRateProduct() (*big.Int, error) {
	return _BalancerBasePool.Contract.GetATHRateProduct(&_BalancerBasePool.CallOpts)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolCaller) GetActionId(opts *bind.CallOpts, selector [4]byte) ([32]byte, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getActionId", selector)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BalancerBasePool.Contract.GetActionId(&_BalancerBasePool.CallOpts, selector)
}

// GetActionId is a free data retrieval call binding the contract method 0x851c1bb3.
//
// Solidity: function getActionId(bytes4 selector) view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetActionId(selector [4]byte) ([32]byte, error) {
	return _BalancerBasePool.Contract.GetActionId(&_BalancerBasePool.CallOpts, selector)
}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) GetActualSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getActualSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) GetActualSupply() (*big.Int, error) {
	return _BalancerBasePool.Contract.GetActualSupply(&_BalancerBasePool.CallOpts)
}

// GetActualSupply is a free data retrieval call binding the contract method 0x876f303b.
//
// Solidity: function getActualSupply() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetActualSupply() (*big.Int, error) {
	return _BalancerBasePool.Contract.GetActualSupply(&_BalancerBasePool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerBasePool *BalancerBasePoolCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerBasePool *BalancerBasePoolSession) GetAuthorizer() (common.Address, error) {
	return _BalancerBasePool.Contract.GetAuthorizer(&_BalancerBasePool.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetAuthorizer() (common.Address, error) {
	return _BalancerBasePool.Contract.GetAuthorizer(&_BalancerBasePool.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolSession) GetDomainSeparator() ([32]byte, error) {
	return _BalancerBasePool.Contract.GetDomainSeparator(&_BalancerBasePool.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _BalancerBasePool.Contract.GetDomainSeparator(&_BalancerBasePool.CallOpts)
}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) GetInvariant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getInvariant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) GetInvariant() (*big.Int, error) {
	return _BalancerBasePool.Contract.GetInvariant(&_BalancerBasePool.CallOpts)
}

// GetInvariant is a free data retrieval call binding the contract method 0xc0ff1a15.
//
// Solidity: function getInvariant() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetInvariant() (*big.Int, error) {
	return _BalancerBasePool.Contract.GetInvariant(&_BalancerBasePool.CallOpts)
}

// GetLastPostJoinExitInvariant is a free data retrieval call binding the contract method 0xb1096278.
//
// Solidity: function getLastPostJoinExitInvariant() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) GetLastPostJoinExitInvariant(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getLastPostJoinExitInvariant")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastPostJoinExitInvariant is a free data retrieval call binding the contract method 0xb1096278.
//
// Solidity: function getLastPostJoinExitInvariant() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) GetLastPostJoinExitInvariant() (*big.Int, error) {
	return _BalancerBasePool.Contract.GetLastPostJoinExitInvariant(&_BalancerBasePool.CallOpts)
}

// GetLastPostJoinExitInvariant is a free data retrieval call binding the contract method 0xb1096278.
//
// Solidity: function getLastPostJoinExitInvariant() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetLastPostJoinExitInvariant() (*big.Int, error) {
	return _BalancerBasePool.Contract.GetLastPostJoinExitInvariant(&_BalancerBasePool.CallOpts)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) GetNextNonce(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getNextNonce", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _BalancerBasePool.Contract.GetNextNonce(&_BalancerBasePool.CallOpts, account)
}

// GetNextNonce is a free data retrieval call binding the contract method 0x90193b7c.
//
// Solidity: function getNextNonce(address account) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetNextNonce(account common.Address) (*big.Int, error) {
	return _BalancerBasePool.Contract.GetNextNonce(&_BalancerBasePool.CallOpts, account)
}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BalancerBasePool *BalancerBasePoolCaller) GetNormalizedWeights(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getNormalizedWeights")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BalancerBasePool *BalancerBasePoolSession) GetNormalizedWeights() ([]*big.Int, error) {
	return _BalancerBasePool.Contract.GetNormalizedWeights(&_BalancerBasePool.CallOpts)
}

// GetNormalizedWeights is a free data retrieval call binding the contract method 0xf89f27ed.
//
// Solidity: function getNormalizedWeights() view returns(uint256[])
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetNormalizedWeights() ([]*big.Int, error) {
	return _BalancerBasePool.Contract.GetNormalizedWeights(&_BalancerBasePool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerBasePool *BalancerBasePoolCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerBasePool *BalancerBasePoolSession) GetOwner() (common.Address, error) {
	return _BalancerBasePool.Contract.GetOwner(&_BalancerBasePool.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetOwner() (common.Address, error) {
	return _BalancerBasePool.Contract.GetOwner(&_BalancerBasePool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerBasePool *BalancerBasePoolCaller) GetPausedState(opts *bind.CallOpts) (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getPausedState")

	outstruct := new(struct {
		Paused              bool
		PauseWindowEndTime  *big.Int
		BufferPeriodEndTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PauseWindowEndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BufferPeriodEndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerBasePool *BalancerBasePoolSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BalancerBasePool.Contract.GetPausedState(&_BalancerBasePool.CallOpts)
}

// GetPausedState is a free data retrieval call binding the contract method 0x1c0de051.
//
// Solidity: function getPausedState() view returns(bool paused, uint256 pauseWindowEndTime, uint256 bufferPeriodEndTime)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetPausedState() (struct {
	Paused              bool
	PauseWindowEndTime  *big.Int
	BufferPeriodEndTime *big.Int
}, error) {
	return _BalancerBasePool.Contract.GetPausedState(&_BalancerBasePool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolCaller) GetPoolId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getPoolId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolSession) GetPoolId() ([32]byte, error) {
	return _BalancerBasePool.Contract.GetPoolId(&_BalancerBasePool.CallOpts)
}

// GetPoolId is a free data retrieval call binding the contract method 0x38fff2d0.
//
// Solidity: function getPoolId() view returns(bytes32)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetPoolId() ([32]byte, error) {
	return _BalancerBasePool.Contract.GetPoolId(&_BalancerBasePool.CallOpts)
}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) GetProtocolFeePercentageCache(opts *bind.CallOpts, feeType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getProtocolFeePercentageCache", feeType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) GetProtocolFeePercentageCache(feeType *big.Int) (*big.Int, error) {
	return _BalancerBasePool.Contract.GetProtocolFeePercentageCache(&_BalancerBasePool.CallOpts, feeType)
}

// GetProtocolFeePercentageCache is a free data retrieval call binding the contract method 0x70464016.
//
// Solidity: function getProtocolFeePercentageCache(uint256 feeType) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetProtocolFeePercentageCache(feeType *big.Int) (*big.Int, error) {
	return _BalancerBasePool.Contract.GetProtocolFeePercentageCache(&_BalancerBasePool.CallOpts, feeType)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BalancerBasePool *BalancerBasePoolCaller) GetProtocolFeesCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getProtocolFeesCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BalancerBasePool *BalancerBasePoolSession) GetProtocolFeesCollector() (common.Address, error) {
	return _BalancerBasePool.Contract.GetProtocolFeesCollector(&_BalancerBasePool.CallOpts)
}

// GetProtocolFeesCollector is a free data retrieval call binding the contract method 0xd2946c2b.
//
// Solidity: function getProtocolFeesCollector() view returns(address)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetProtocolFeesCollector() (common.Address, error) {
	return _BalancerBasePool.Contract.GetProtocolFeesCollector(&_BalancerBasePool.CallOpts)
}

// GetProtocolSwapFeeDelegation is a free data retrieval call binding the contract method 0x15b0015b.
//
// Solidity: function getProtocolSwapFeeDelegation() view returns(bool)
func (_BalancerBasePool *BalancerBasePoolCaller) GetProtocolSwapFeeDelegation(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getProtocolSwapFeeDelegation")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetProtocolSwapFeeDelegation is a free data retrieval call binding the contract method 0x15b0015b.
//
// Solidity: function getProtocolSwapFeeDelegation() view returns(bool)
func (_BalancerBasePool *BalancerBasePoolSession) GetProtocolSwapFeeDelegation() (bool, error) {
	return _BalancerBasePool.Contract.GetProtocolSwapFeeDelegation(&_BalancerBasePool.CallOpts)
}

// GetProtocolSwapFeeDelegation is a free data retrieval call binding the contract method 0x15b0015b.
//
// Solidity: function getProtocolSwapFeeDelegation() view returns(bool)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetProtocolSwapFeeDelegation() (bool, error) {
	return _BalancerBasePool.Contract.GetProtocolSwapFeeDelegation(&_BalancerBasePool.CallOpts)
}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BalancerBasePool *BalancerBasePoolCaller) GetRateProviders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getRateProviders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BalancerBasePool *BalancerBasePoolSession) GetRateProviders() ([]common.Address, error) {
	return _BalancerBasePool.Contract.GetRateProviders(&_BalancerBasePool.CallOpts)
}

// GetRateProviders is a free data retrieval call binding the contract method 0x238a2d59.
//
// Solidity: function getRateProviders() view returns(address[])
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetRateProviders() ([]common.Address, error) {
	return _BalancerBasePool.Contract.GetRateProviders(&_BalancerBasePool.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BalancerBasePool *BalancerBasePoolCaller) GetScalingFactors(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getScalingFactors")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BalancerBasePool *BalancerBasePoolSession) GetScalingFactors() ([]*big.Int, error) {
	return _BalancerBasePool.Contract.GetScalingFactors(&_BalancerBasePool.CallOpts)
}

// GetScalingFactors is a free data retrieval call binding the contract method 0x1dd746ea.
//
// Solidity: function getScalingFactors() view returns(uint256[])
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetScalingFactors() ([]*big.Int, error) {
	return _BalancerBasePool.Contract.GetScalingFactors(&_BalancerBasePool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) GetSwapFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getSwapFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BalancerBasePool.Contract.GetSwapFeePercentage(&_BalancerBasePool.CallOpts)
}

// GetSwapFeePercentage is a free data retrieval call binding the contract method 0x55c67628.
//
// Solidity: function getSwapFeePercentage() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetSwapFeePercentage() (*big.Int, error) {
	return _BalancerBasePool.Contract.GetSwapFeePercentage(&_BalancerBasePool.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerBasePool *BalancerBasePoolCaller) GetVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "getVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerBasePool *BalancerBasePoolSession) GetVault() (common.Address, error) {
	return _BalancerBasePool.Contract.GetVault(&_BalancerBasePool.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x8d928af8.
//
// Solidity: function getVault() view returns(address)
func (_BalancerBasePool *BalancerBasePoolCallerSession) GetVault() (common.Address, error) {
	return _BalancerBasePool.Contract.GetVault(&_BalancerBasePool.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BalancerBasePool *BalancerBasePoolCaller) InRecoveryMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "inRecoveryMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BalancerBasePool *BalancerBasePoolSession) InRecoveryMode() (bool, error) {
	return _BalancerBasePool.Contract.InRecoveryMode(&_BalancerBasePool.CallOpts)
}

// InRecoveryMode is a free data retrieval call binding the contract method 0xb35056b8.
//
// Solidity: function inRecoveryMode() view returns(bool)
func (_BalancerBasePool *BalancerBasePoolCallerSession) InRecoveryMode() (bool, error) {
	return _BalancerBasePool.Contract.InRecoveryMode(&_BalancerBasePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerBasePool *BalancerBasePoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerBasePool *BalancerBasePoolSession) Name() (string, error) {
	return _BalancerBasePool.Contract.Name(&_BalancerBasePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerBasePool *BalancerBasePoolCallerSession) Name() (string, error) {
	return _BalancerBasePool.Contract.Name(&_BalancerBasePool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BalancerBasePool.Contract.Nonces(&_BalancerBasePool.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BalancerBasePool.Contract.Nonces(&_BalancerBasePool.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerBasePool *BalancerBasePoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerBasePool *BalancerBasePoolSession) Symbol() (string, error) {
	return _BalancerBasePool.Contract.Symbol(&_BalancerBasePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerBasePool *BalancerBasePoolCallerSession) Symbol() (string, error) {
	return _BalancerBasePool.Contract.Symbol(&_BalancerBasePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) TotalSupply() (*big.Int, error) {
	return _BalancerBasePool.Contract.TotalSupply(&_BalancerBasePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerBasePool *BalancerBasePoolCallerSession) TotalSupply() (*big.Int, error) {
	return _BalancerBasePool.Contract.TotalSupply(&_BalancerBasePool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BalancerBasePool *BalancerBasePoolCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BalancerBasePool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BalancerBasePool *BalancerBasePoolSession) Version() (string, error) {
	return _BalancerBasePool.Contract.Version(&_BalancerBasePool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BalancerBasePool *BalancerBasePoolCallerSession) Version() (string, error) {
	return _BalancerBasePool.Contract.Version(&_BalancerBasePool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.Approve(&_BalancerBasePool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.Approve(&_BalancerBasePool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "decreaseAllowance", spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.DecreaseAllowance(&_BalancerBasePool.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolTransactorSession) DecreaseAllowance(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.DecreaseAllowance(&_BalancerBasePool.TransactOpts, spender, amount)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BalancerBasePool *BalancerBasePoolTransactor) DisableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "disableRecoveryMode")
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BalancerBasePool *BalancerBasePoolSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _BalancerBasePool.Contract.DisableRecoveryMode(&_BalancerBasePool.TransactOpts)
}

// DisableRecoveryMode is a paid mutator transaction binding the contract method 0xb7b814fc.
//
// Solidity: function disableRecoveryMode() returns()
func (_BalancerBasePool *BalancerBasePoolTransactorSession) DisableRecoveryMode() (*types.Transaction, error) {
	return _BalancerBasePool.Contract.DisableRecoveryMode(&_BalancerBasePool.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BalancerBasePool *BalancerBasePoolTransactor) EnableRecoveryMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "enableRecoveryMode")
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BalancerBasePool *BalancerBasePoolSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _BalancerBasePool.Contract.EnableRecoveryMode(&_BalancerBasePool.TransactOpts)
}

// EnableRecoveryMode is a paid mutator transaction binding the contract method 0x54a844ba.
//
// Solidity: function enableRecoveryMode() returns()
func (_BalancerBasePool *BalancerBasePoolTransactorSession) EnableRecoveryMode() (*types.Transaction, error) {
	return _BalancerBasePool.Contract.EnableRecoveryMode(&_BalancerBasePool.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerBasePool *BalancerBasePoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerBasePool *BalancerBasePoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.IncreaseAllowance(&_BalancerBasePool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BalancerBasePool *BalancerBasePoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.IncreaseAllowance(&_BalancerBasePool.TransactOpts, spender, addedValue)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerBasePool *BalancerBasePoolTransactor) OnExitPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "onExitPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerBasePool *BalancerBasePoolSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.OnExitPool(&_BalancerBasePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnExitPool is a paid mutator transaction binding the contract method 0x74f3b009.
//
// Solidity: function onExitPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerBasePool *BalancerBasePoolTransactorSession) OnExitPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.OnExitPool(&_BalancerBasePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerBasePool *BalancerBasePoolTransactor) OnJoinPool(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "onJoinPool", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerBasePool *BalancerBasePoolSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.OnJoinPool(&_BalancerBasePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnJoinPool is a paid mutator transaction binding the contract method 0xd5c096c4.
//
// Solidity: function onJoinPool(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256[], uint256[])
func (_BalancerBasePool *BalancerBasePoolTransactorSession) OnJoinPool(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.OnJoinPool(&_BalancerBasePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BalancerBasePool *BalancerBasePoolTransactor) OnSwap(opts *bind.TransactOpts, request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "onSwap", request, balanceTokenIn, balanceTokenOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BalancerBasePool *BalancerBasePoolSession) OnSwap(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.OnSwap(&_BalancerBasePool.TransactOpts, request, balanceTokenIn, balanceTokenOut)
}

// OnSwap is a paid mutator transaction binding the contract method 0x9d2c110c.
//
// Solidity: function onSwap((uint8,address,address,uint256,bytes32,uint256,address,address,bytes) request, uint256 balanceTokenIn, uint256 balanceTokenOut) returns(uint256)
func (_BalancerBasePool *BalancerBasePoolTransactorSession) OnSwap(request IPoolSwapStructsSwapRequest, balanceTokenIn *big.Int, balanceTokenOut *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.OnSwap(&_BalancerBasePool.TransactOpts, request, balanceTokenIn, balanceTokenOut)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BalancerBasePool *BalancerBasePoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BalancerBasePool *BalancerBasePoolSession) Pause() (*types.Transaction, error) {
	return _BalancerBasePool.Contract.Pause(&_BalancerBasePool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BalancerBasePool *BalancerBasePoolTransactorSession) Pause() (*types.Transaction, error) {
	return _BalancerBasePool.Contract.Pause(&_BalancerBasePool.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerBasePool *BalancerBasePoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerBasePool *BalancerBasePoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.Permit(&_BalancerBasePool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BalancerBasePool *BalancerBasePoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.Permit(&_BalancerBasePool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerBasePool *BalancerBasePoolTransactor) QueryExit(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "queryExit", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerBasePool *BalancerBasePoolSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.QueryExit(&_BalancerBasePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryExit is a paid mutator transaction binding the contract method 0x6028bfd4.
//
// Solidity: function queryExit(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptIn, uint256[] amountsOut)
func (_BalancerBasePool *BalancerBasePoolTransactorSession) QueryExit(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.QueryExit(&_BalancerBasePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerBasePool *BalancerBasePoolTransactor) QueryJoin(opts *bind.TransactOpts, poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "queryJoin", poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerBasePool *BalancerBasePoolSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.QueryJoin(&_BalancerBasePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// QueryJoin is a paid mutator transaction binding the contract method 0x87ec6817.
//
// Solidity: function queryJoin(bytes32 poolId, address sender, address recipient, uint256[] balances, uint256 lastChangeBlock, uint256 protocolSwapFeePercentage, bytes userData) returns(uint256 bptOut, uint256[] amountsIn)
func (_BalancerBasePool *BalancerBasePoolTransactorSession) QueryJoin(poolId [32]byte, sender common.Address, recipient common.Address, balances []*big.Int, lastChangeBlock *big.Int, protocolSwapFeePercentage *big.Int, userData []byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.QueryJoin(&_BalancerBasePool.TransactOpts, poolId, sender, recipient, balances, lastChangeBlock, protocolSwapFeePercentage, userData)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerBasePool *BalancerBasePoolTransactor) SetAssetManagerPoolConfig(opts *bind.TransactOpts, token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "setAssetManagerPoolConfig", token, poolConfig)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerBasePool *BalancerBasePoolSession) SetAssetManagerPoolConfig(token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.SetAssetManagerPoolConfig(&_BalancerBasePool.TransactOpts, token, poolConfig)
}

// SetAssetManagerPoolConfig is a paid mutator transaction binding the contract method 0x50dd6ed9.
//
// Solidity: function setAssetManagerPoolConfig(address token, bytes poolConfig) returns()
func (_BalancerBasePool *BalancerBasePoolTransactorSession) SetAssetManagerPoolConfig(token common.Address, poolConfig []byte) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.SetAssetManagerPoolConfig(&_BalancerBasePool.TransactOpts, token, poolConfig)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerBasePool *BalancerBasePoolTransactor) SetSwapFeePercentage(opts *bind.TransactOpts, swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "setSwapFeePercentage", swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerBasePool *BalancerBasePoolSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.SetSwapFeePercentage(&_BalancerBasePool.TransactOpts, swapFeePercentage)
}

// SetSwapFeePercentage is a paid mutator transaction binding the contract method 0x38e9922e.
//
// Solidity: function setSwapFeePercentage(uint256 swapFeePercentage) returns()
func (_BalancerBasePool *BalancerBasePoolTransactorSession) SetSwapFeePercentage(swapFeePercentage *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.SetSwapFeePercentage(&_BalancerBasePool.TransactOpts, swapFeePercentage)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.Transfer(&_BalancerBasePool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.Transfer(&_BalancerBasePool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.TransferFrom(&_BalancerBasePool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BalancerBasePool *BalancerBasePoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BalancerBasePool.Contract.TransferFrom(&_BalancerBasePool.TransactOpts, sender, recipient, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BalancerBasePool *BalancerBasePoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BalancerBasePool *BalancerBasePoolSession) Unpause() (*types.Transaction, error) {
	return _BalancerBasePool.Contract.Unpause(&_BalancerBasePool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BalancerBasePool *BalancerBasePoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _BalancerBasePool.Contract.Unpause(&_BalancerBasePool.TransactOpts)
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BalancerBasePool *BalancerBasePoolTransactor) UpdateProtocolFeePercentageCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerBasePool.contract.Transact(opts, "updateProtocolFeePercentageCache")
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BalancerBasePool *BalancerBasePoolSession) UpdateProtocolFeePercentageCache() (*types.Transaction, error) {
	return _BalancerBasePool.Contract.UpdateProtocolFeePercentageCache(&_BalancerBasePool.TransactOpts)
}

// UpdateProtocolFeePercentageCache is a paid mutator transaction binding the contract method 0x0da0669c.
//
// Solidity: function updateProtocolFeePercentageCache() returns()
func (_BalancerBasePool *BalancerBasePoolTransactorSession) UpdateProtocolFeePercentageCache() (*types.Transaction, error) {
	return _BalancerBasePool.Contract.UpdateProtocolFeePercentageCache(&_BalancerBasePool.TransactOpts)
}

// BalancerBasePoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BalancerBasePool contract.
type BalancerBasePoolApprovalIterator struct {
	Event *BalancerBasePoolApproval // Event containing the contract specifics and raw log

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
func (it *BalancerBasePoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerBasePoolApproval)
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
		it.Event = new(BalancerBasePoolApproval)
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
func (it *BalancerBasePoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerBasePoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerBasePoolApproval represents a Approval event raised by the BalancerBasePool contract.
type BalancerBasePoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BalancerBasePool *BalancerBasePoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BalancerBasePoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BalancerBasePool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BalancerBasePoolApprovalIterator{contract: _BalancerBasePool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BalancerBasePool *BalancerBasePoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BalancerBasePoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BalancerBasePool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerBasePoolApproval)
				if err := _BalancerBasePool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BalancerBasePool *BalancerBasePoolFilterer) ParseApproval(log types.Log) (*BalancerBasePoolApproval, error) {
	event := new(BalancerBasePoolApproval)
	if err := _BalancerBasePool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerBasePoolPausedStateChangedIterator is returned from FilterPausedStateChanged and is used to iterate over the raw logs and unpacked data for PausedStateChanged events raised by the BalancerBasePool contract.
type BalancerBasePoolPausedStateChangedIterator struct {
	Event *BalancerBasePoolPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *BalancerBasePoolPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerBasePoolPausedStateChanged)
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
		it.Event = new(BalancerBasePoolPausedStateChanged)
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
func (it *BalancerBasePoolPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerBasePoolPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerBasePoolPausedStateChanged represents a PausedStateChanged event raised by the BalancerBasePool contract.
type BalancerBasePoolPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPausedStateChanged is a free log retrieval operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerBasePool *BalancerBasePoolFilterer) FilterPausedStateChanged(opts *bind.FilterOpts) (*BalancerBasePoolPausedStateChangedIterator, error) {

	logs, sub, err := _BalancerBasePool.contract.FilterLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerBasePoolPausedStateChangedIterator{contract: _BalancerBasePool.contract, event: "PausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPausedStateChanged is a free log subscription operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerBasePool *BalancerBasePoolFilterer) WatchPausedStateChanged(opts *bind.WatchOpts, sink chan<- *BalancerBasePoolPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerBasePool.contract.WatchLogs(opts, "PausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerBasePoolPausedStateChanged)
				if err := _BalancerBasePool.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
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

// ParsePausedStateChanged is a log parse operation binding the contract event 0x9e3a5e37224532dea67b89face185703738a228a6e8a23dee546960180d3be64.
//
// Solidity: event PausedStateChanged(bool paused)
func (_BalancerBasePool *BalancerBasePoolFilterer) ParsePausedStateChanged(log types.Log) (*BalancerBasePoolPausedStateChanged, error) {
	event := new(BalancerBasePoolPausedStateChanged)
	if err := _BalancerBasePool.contract.UnpackLog(event, "PausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerBasePoolProtocolFeePercentageCacheUpdatedIterator is returned from FilterProtocolFeePercentageCacheUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeePercentageCacheUpdated events raised by the BalancerBasePool contract.
type BalancerBasePoolProtocolFeePercentageCacheUpdatedIterator struct {
	Event *BalancerBasePoolProtocolFeePercentageCacheUpdated // Event containing the contract specifics and raw log

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
func (it *BalancerBasePoolProtocolFeePercentageCacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerBasePoolProtocolFeePercentageCacheUpdated)
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
		it.Event = new(BalancerBasePoolProtocolFeePercentageCacheUpdated)
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
func (it *BalancerBasePoolProtocolFeePercentageCacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerBasePoolProtocolFeePercentageCacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerBasePoolProtocolFeePercentageCacheUpdated represents a ProtocolFeePercentageCacheUpdated event raised by the BalancerBasePool contract.
type BalancerBasePoolProtocolFeePercentageCacheUpdated struct {
	FeeType               *big.Int
	ProtocolFeePercentage *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeePercentageCacheUpdated is a free log retrieval operation binding the contract event 0x6bfb689528fa96ec1ad670ad6d6064be1ae96bfd5d2ee35c837fd0fe0c11959a.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(uint256 indexed feeType, uint256 protocolFeePercentage)
func (_BalancerBasePool *BalancerBasePoolFilterer) FilterProtocolFeePercentageCacheUpdated(opts *bind.FilterOpts, feeType []*big.Int) (*BalancerBasePoolProtocolFeePercentageCacheUpdatedIterator, error) {

	var feeTypeRule []interface{}
	for _, feeTypeItem := range feeType {
		feeTypeRule = append(feeTypeRule, feeTypeItem)
	}

	logs, sub, err := _BalancerBasePool.contract.FilterLogs(opts, "ProtocolFeePercentageCacheUpdated", feeTypeRule)
	if err != nil {
		return nil, err
	}
	return &BalancerBasePoolProtocolFeePercentageCacheUpdatedIterator{contract: _BalancerBasePool.contract, event: "ProtocolFeePercentageCacheUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeePercentageCacheUpdated is a free log subscription operation binding the contract event 0x6bfb689528fa96ec1ad670ad6d6064be1ae96bfd5d2ee35c837fd0fe0c11959a.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(uint256 indexed feeType, uint256 protocolFeePercentage)
func (_BalancerBasePool *BalancerBasePoolFilterer) WatchProtocolFeePercentageCacheUpdated(opts *bind.WatchOpts, sink chan<- *BalancerBasePoolProtocolFeePercentageCacheUpdated, feeType []*big.Int) (event.Subscription, error) {

	var feeTypeRule []interface{}
	for _, feeTypeItem := range feeType {
		feeTypeRule = append(feeTypeRule, feeTypeItem)
	}

	logs, sub, err := _BalancerBasePool.contract.WatchLogs(opts, "ProtocolFeePercentageCacheUpdated", feeTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerBasePoolProtocolFeePercentageCacheUpdated)
				if err := _BalancerBasePool.contract.UnpackLog(event, "ProtocolFeePercentageCacheUpdated", log); err != nil {
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

// ParseProtocolFeePercentageCacheUpdated is a log parse operation binding the contract event 0x6bfb689528fa96ec1ad670ad6d6064be1ae96bfd5d2ee35c837fd0fe0c11959a.
//
// Solidity: event ProtocolFeePercentageCacheUpdated(uint256 indexed feeType, uint256 protocolFeePercentage)
func (_BalancerBasePool *BalancerBasePoolFilterer) ParseProtocolFeePercentageCacheUpdated(log types.Log) (*BalancerBasePoolProtocolFeePercentageCacheUpdated, error) {
	event := new(BalancerBasePoolProtocolFeePercentageCacheUpdated)
	if err := _BalancerBasePool.contract.UnpackLog(event, "ProtocolFeePercentageCacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerBasePoolRecoveryModeStateChangedIterator is returned from FilterRecoveryModeStateChanged and is used to iterate over the raw logs and unpacked data for RecoveryModeStateChanged events raised by the BalancerBasePool contract.
type BalancerBasePoolRecoveryModeStateChangedIterator struct {
	Event *BalancerBasePoolRecoveryModeStateChanged // Event containing the contract specifics and raw log

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
func (it *BalancerBasePoolRecoveryModeStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerBasePoolRecoveryModeStateChanged)
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
		it.Event = new(BalancerBasePoolRecoveryModeStateChanged)
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
func (it *BalancerBasePoolRecoveryModeStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerBasePoolRecoveryModeStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerBasePoolRecoveryModeStateChanged represents a RecoveryModeStateChanged event raised by the BalancerBasePool contract.
type BalancerBasePoolRecoveryModeStateChanged struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRecoveryModeStateChanged is a free log retrieval operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BalancerBasePool *BalancerBasePoolFilterer) FilterRecoveryModeStateChanged(opts *bind.FilterOpts) (*BalancerBasePoolRecoveryModeStateChangedIterator, error) {

	logs, sub, err := _BalancerBasePool.contract.FilterLogs(opts, "RecoveryModeStateChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerBasePoolRecoveryModeStateChangedIterator{contract: _BalancerBasePool.contract, event: "RecoveryModeStateChanged", logs: logs, sub: sub}, nil
}

// WatchRecoveryModeStateChanged is a free log subscription operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BalancerBasePool *BalancerBasePoolFilterer) WatchRecoveryModeStateChanged(opts *bind.WatchOpts, sink chan<- *BalancerBasePoolRecoveryModeStateChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerBasePool.contract.WatchLogs(opts, "RecoveryModeStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerBasePoolRecoveryModeStateChanged)
				if err := _BalancerBasePool.contract.UnpackLog(event, "RecoveryModeStateChanged", log); err != nil {
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

// ParseRecoveryModeStateChanged is a log parse operation binding the contract event 0xeff3d4d215b42bf0960be9c6d5e05c22cba4df6627a3a523e2acee733b5854c8.
//
// Solidity: event RecoveryModeStateChanged(bool enabled)
func (_BalancerBasePool *BalancerBasePoolFilterer) ParseRecoveryModeStateChanged(log types.Log) (*BalancerBasePoolRecoveryModeStateChanged, error) {
	event := new(BalancerBasePoolRecoveryModeStateChanged)
	if err := _BalancerBasePool.contract.UnpackLog(event, "RecoveryModeStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerBasePoolSwapFeePercentageChangedIterator is returned from FilterSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for SwapFeePercentageChanged events raised by the BalancerBasePool contract.
type BalancerBasePoolSwapFeePercentageChangedIterator struct {
	Event *BalancerBasePoolSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *BalancerBasePoolSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerBasePoolSwapFeePercentageChanged)
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
		it.Event = new(BalancerBasePoolSwapFeePercentageChanged)
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
func (it *BalancerBasePoolSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerBasePoolSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerBasePoolSwapFeePercentageChanged represents a SwapFeePercentageChanged event raised by the BalancerBasePool contract.
type BalancerBasePoolSwapFeePercentageChanged struct {
	SwapFeePercentage *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerBasePool *BalancerBasePoolFilterer) FilterSwapFeePercentageChanged(opts *bind.FilterOpts) (*BalancerBasePoolSwapFeePercentageChangedIterator, error) {

	logs, sub, err := _BalancerBasePool.contract.FilterLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &BalancerBasePoolSwapFeePercentageChangedIterator{contract: _BalancerBasePool.contract, event: "SwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeePercentageChanged is a free log subscription operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerBasePool *BalancerBasePoolFilterer) WatchSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *BalancerBasePoolSwapFeePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _BalancerBasePool.contract.WatchLogs(opts, "SwapFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerBasePoolSwapFeePercentageChanged)
				if err := _BalancerBasePool.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
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

// ParseSwapFeePercentageChanged is a log parse operation binding the contract event 0xa9ba3ffe0b6c366b81232caab38605a0699ad5398d6cce76f91ee809e322dafc.
//
// Solidity: event SwapFeePercentageChanged(uint256 swapFeePercentage)
func (_BalancerBasePool *BalancerBasePoolFilterer) ParseSwapFeePercentageChanged(log types.Log) (*BalancerBasePoolSwapFeePercentageChanged, error) {
	event := new(BalancerBasePoolSwapFeePercentageChanged)
	if err := _BalancerBasePool.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BalancerBasePoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BalancerBasePool contract.
type BalancerBasePoolTransferIterator struct {
	Event *BalancerBasePoolTransfer // Event containing the contract specifics and raw log

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
func (it *BalancerBasePoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerBasePoolTransfer)
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
		it.Event = new(BalancerBasePoolTransfer)
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
func (it *BalancerBasePoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerBasePoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerBasePoolTransfer represents a Transfer event raised by the BalancerBasePool contract.
type BalancerBasePoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BalancerBasePool *BalancerBasePoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BalancerBasePoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BalancerBasePool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BalancerBasePoolTransferIterator{contract: _BalancerBasePool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BalancerBasePool *BalancerBasePoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BalancerBasePoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BalancerBasePool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerBasePoolTransfer)
				if err := _BalancerBasePool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BalancerBasePool *BalancerBasePoolFilterer) ParseTransfer(log types.Log) (*BalancerBasePoolTransfer, error) {
	event := new(BalancerBasePoolTransfer)
	if err := _BalancerBasePool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
