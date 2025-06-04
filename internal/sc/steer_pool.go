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

// KodiakMultiPositionLiquidityManagerLiquidityPositions is an auto generated low-level Go binding around an user-defined struct.
type KodiakMultiPositionLiquidityManagerLiquidityPositions struct {
	LowerTick      []*big.Int
	UpperTick      []*big.Int
	RelativeWeight []uint16
}

// SteerPoolMetaData contains all meta data concerning the SteerPool contract.
var SteerPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Earned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Earned\",\"type\":\"uint256\"}],\"name\":\"FeesEarned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STEER_FRACTION_OF_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STRATEGIST_FRACTION_OF_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"accruedFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"accruedFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedSteerFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedSteerFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedStrategistFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedStrategistFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"feeIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"collectFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Used\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Used\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"name\":\"emergencyBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositions\",\"outputs\":[{\"internalType\":\"int24[]\",\"name\":\"\",\"type\":\"int24[]\"},{\"internalType\":\"int24[]\",\"name\":\"\",\"type\":\"int24[]\"},{\"internalType\":\"uint16[]\",\"name\":\"\",\"type\":\"uint16[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_steer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTickChange\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int24[]\",\"name\":\"lowerTick\",\"type\":\"int24[]\"},{\"internalType\":\"int24[]\",\"name\":\"upperTick\",\"type\":\"int24[]\"},{\"internalType\":\"uint16[]\",\"name\":\"relativeWeight\",\"type\":\"uint16[]\"}],\"internalType\":\"structKodiakMultiPositionLiquidityManager.LiquidityPositions\",\"name\":\"newPositions\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"timeSensitiveData\",\"type\":\"bytes\"}],\"name\":\"tend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twapInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Wanted\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Wanted\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SteerPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use SteerPoolMetaData.ABI instead.
var SteerPoolABI = SteerPoolMetaData.ABI

// SteerPool is an auto generated Go binding around an Ethereum contract.
type SteerPool struct {
	SteerPoolCaller     // Read-only binding to the contract
	SteerPoolTransactor // Write-only binding to the contract
	SteerPoolFilterer   // Log filterer for contract events
}

// SteerPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SteerPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SteerPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SteerPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SteerPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SteerPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SteerPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SteerPoolSession struct {
	Contract     *SteerPool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SteerPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SteerPoolCallerSession struct {
	Contract *SteerPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SteerPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SteerPoolTransactorSession struct {
	Contract     *SteerPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SteerPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SteerPoolRaw struct {
	Contract *SteerPool // Generic contract binding to access the raw methods on
}

// SteerPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SteerPoolCallerRaw struct {
	Contract *SteerPoolCaller // Generic read-only contract binding to access the raw methods on
}

// SteerPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SteerPoolTransactorRaw struct {
	Contract *SteerPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSteerPool creates a new instance of SteerPool, bound to a specific deployed contract.
func NewSteerPool(address common.Address, backend bind.ContractBackend) (*SteerPool, error) {
	contract, err := bindSteerPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SteerPool{SteerPoolCaller: SteerPoolCaller{contract: contract}, SteerPoolTransactor: SteerPoolTransactor{contract: contract}, SteerPoolFilterer: SteerPoolFilterer{contract: contract}}, nil
}

// NewSteerPoolCaller creates a new read-only instance of SteerPool, bound to a specific deployed contract.
func NewSteerPoolCaller(address common.Address, caller bind.ContractCaller) (*SteerPoolCaller, error) {
	contract, err := bindSteerPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SteerPoolCaller{contract: contract}, nil
}

// NewSteerPoolTransactor creates a new write-only instance of SteerPool, bound to a specific deployed contract.
func NewSteerPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SteerPoolTransactor, error) {
	contract, err := bindSteerPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SteerPoolTransactor{contract: contract}, nil
}

// NewSteerPoolFilterer creates a new log filterer instance of SteerPool, bound to a specific deployed contract.
func NewSteerPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SteerPoolFilterer, error) {
	contract, err := bindSteerPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SteerPoolFilterer{contract: contract}, nil
}

// bindSteerPool binds a generic wrapper to an already deployed contract.
func bindSteerPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SteerPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SteerPool *SteerPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SteerPool.Contract.SteerPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SteerPool *SteerPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerPool.Contract.SteerPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SteerPool *SteerPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SteerPool.Contract.SteerPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SteerPool *SteerPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SteerPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SteerPool *SteerPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SteerPool *SteerPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SteerPool.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SteerPool *SteerPoolCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SteerPool *SteerPoolSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SteerPool.Contract.DEFAULTADMINROLE(&_SteerPool.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SteerPool *SteerPoolCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SteerPool.Contract.DEFAULTADMINROLE(&_SteerPool.CallOpts)
}

// STEERFRACTIONOFFEE is a free data retrieval call binding the contract method 0x1c1f1936.
//
// Solidity: function STEER_FRACTION_OF_FEE() view returns(uint256)
func (_SteerPool *SteerPoolCaller) STEERFRACTIONOFFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "STEER_FRACTION_OF_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STEERFRACTIONOFFEE is a free data retrieval call binding the contract method 0x1c1f1936.
//
// Solidity: function STEER_FRACTION_OF_FEE() view returns(uint256)
func (_SteerPool *SteerPoolSession) STEERFRACTIONOFFEE() (*big.Int, error) {
	return _SteerPool.Contract.STEERFRACTIONOFFEE(&_SteerPool.CallOpts)
}

// STEERFRACTIONOFFEE is a free data retrieval call binding the contract method 0x1c1f1936.
//
// Solidity: function STEER_FRACTION_OF_FEE() view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) STEERFRACTIONOFFEE() (*big.Int, error) {
	return _SteerPool.Contract.STEERFRACTIONOFFEE(&_SteerPool.CallOpts)
}

// STRATEGISTFRACTIONOFFEE is a free data retrieval call binding the contract method 0x30faab2a.
//
// Solidity: function STRATEGIST_FRACTION_OF_FEE() view returns(uint256)
func (_SteerPool *SteerPoolCaller) STRATEGISTFRACTIONOFFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "STRATEGIST_FRACTION_OF_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STRATEGISTFRACTIONOFFEE is a free data retrieval call binding the contract method 0x30faab2a.
//
// Solidity: function STRATEGIST_FRACTION_OF_FEE() view returns(uint256)
func (_SteerPool *SteerPoolSession) STRATEGISTFRACTIONOFFEE() (*big.Int, error) {
	return _SteerPool.Contract.STRATEGISTFRACTIONOFFEE(&_SteerPool.CallOpts)
}

// STRATEGISTFRACTIONOFFEE is a free data retrieval call binding the contract method 0x30faab2a.
//
// Solidity: function STRATEGIST_FRACTION_OF_FEE() view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) STRATEGISTFRACTIONOFFEE() (*big.Int, error) {
	return _SteerPool.Contract.STRATEGISTFRACTIONOFFEE(&_SteerPool.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_SteerPool *SteerPoolCaller) TOTALFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "TOTAL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_SteerPool *SteerPoolSession) TOTALFEE() (*big.Int, error) {
	return _SteerPool.Contract.TOTALFEE(&_SteerPool.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) TOTALFEE() (*big.Int, error) {
	return _SteerPool.Contract.TOTALFEE(&_SteerPool.CallOpts)
}

// AccruedFees0 is a free data retrieval call binding the contract method 0x3bec034a.
//
// Solidity: function accruedFees0(string ) view returns(uint256)
func (_SteerPool *SteerPoolCaller) AccruedFees0(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "accruedFees0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedFees0 is a free data retrieval call binding the contract method 0x3bec034a.
//
// Solidity: function accruedFees0(string ) view returns(uint256)
func (_SteerPool *SteerPoolSession) AccruedFees0(arg0 string) (*big.Int, error) {
	return _SteerPool.Contract.AccruedFees0(&_SteerPool.CallOpts, arg0)
}

// AccruedFees0 is a free data retrieval call binding the contract method 0x3bec034a.
//
// Solidity: function accruedFees0(string ) view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) AccruedFees0(arg0 string) (*big.Int, error) {
	return _SteerPool.Contract.AccruedFees0(&_SteerPool.CallOpts, arg0)
}

// AccruedFees1 is a free data retrieval call binding the contract method 0xb7ab45b6.
//
// Solidity: function accruedFees1(string ) view returns(uint256)
func (_SteerPool *SteerPoolCaller) AccruedFees1(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "accruedFees1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedFees1 is a free data retrieval call binding the contract method 0xb7ab45b6.
//
// Solidity: function accruedFees1(string ) view returns(uint256)
func (_SteerPool *SteerPoolSession) AccruedFees1(arg0 string) (*big.Int, error) {
	return _SteerPool.Contract.AccruedFees1(&_SteerPool.CallOpts, arg0)
}

// AccruedFees1 is a free data retrieval call binding the contract method 0xb7ab45b6.
//
// Solidity: function accruedFees1(string ) view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) AccruedFees1(arg0 string) (*big.Int, error) {
	return _SteerPool.Contract.AccruedFees1(&_SteerPool.CallOpts, arg0)
}

// AccruedSteerFees0 is a free data retrieval call binding the contract method 0xef94c279.
//
// Solidity: function accruedSteerFees0() view returns(uint256)
func (_SteerPool *SteerPoolCaller) AccruedSteerFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "accruedSteerFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedSteerFees0 is a free data retrieval call binding the contract method 0xef94c279.
//
// Solidity: function accruedSteerFees0() view returns(uint256)
func (_SteerPool *SteerPoolSession) AccruedSteerFees0() (*big.Int, error) {
	return _SteerPool.Contract.AccruedSteerFees0(&_SteerPool.CallOpts)
}

// AccruedSteerFees0 is a free data retrieval call binding the contract method 0xef94c279.
//
// Solidity: function accruedSteerFees0() view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) AccruedSteerFees0() (*big.Int, error) {
	return _SteerPool.Contract.AccruedSteerFees0(&_SteerPool.CallOpts)
}

// AccruedSteerFees1 is a free data retrieval call binding the contract method 0x02bcd6ca.
//
// Solidity: function accruedSteerFees1() view returns(uint256)
func (_SteerPool *SteerPoolCaller) AccruedSteerFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "accruedSteerFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedSteerFees1 is a free data retrieval call binding the contract method 0x02bcd6ca.
//
// Solidity: function accruedSteerFees1() view returns(uint256)
func (_SteerPool *SteerPoolSession) AccruedSteerFees1() (*big.Int, error) {
	return _SteerPool.Contract.AccruedSteerFees1(&_SteerPool.CallOpts)
}

// AccruedSteerFees1 is a free data retrieval call binding the contract method 0x02bcd6ca.
//
// Solidity: function accruedSteerFees1() view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) AccruedSteerFees1() (*big.Int, error) {
	return _SteerPool.Contract.AccruedSteerFees1(&_SteerPool.CallOpts)
}

// AccruedStrategistFees0 is a free data retrieval call binding the contract method 0x3f2218f1.
//
// Solidity: function accruedStrategistFees0() view returns(uint256)
func (_SteerPool *SteerPoolCaller) AccruedStrategistFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "accruedStrategistFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedStrategistFees0 is a free data retrieval call binding the contract method 0x3f2218f1.
//
// Solidity: function accruedStrategistFees0() view returns(uint256)
func (_SteerPool *SteerPoolSession) AccruedStrategistFees0() (*big.Int, error) {
	return _SteerPool.Contract.AccruedStrategistFees0(&_SteerPool.CallOpts)
}

// AccruedStrategistFees0 is a free data retrieval call binding the contract method 0x3f2218f1.
//
// Solidity: function accruedStrategistFees0() view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) AccruedStrategistFees0() (*big.Int, error) {
	return _SteerPool.Contract.AccruedStrategistFees0(&_SteerPool.CallOpts)
}

// AccruedStrategistFees1 is a free data retrieval call binding the contract method 0x366e897b.
//
// Solidity: function accruedStrategistFees1() view returns(uint256)
func (_SteerPool *SteerPoolCaller) AccruedStrategistFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "accruedStrategistFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedStrategistFees1 is a free data retrieval call binding the contract method 0x366e897b.
//
// Solidity: function accruedStrategistFees1() view returns(uint256)
func (_SteerPool *SteerPoolSession) AccruedStrategistFees1() (*big.Int, error) {
	return _SteerPool.Contract.AccruedStrategistFees1(&_SteerPool.CallOpts)
}

// AccruedStrategistFees1 is a free data retrieval call binding the contract method 0x366e897b.
//
// Solidity: function accruedStrategistFees1() view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) AccruedStrategistFees1() (*big.Int, error) {
	return _SteerPool.Contract.AccruedStrategistFees1(&_SteerPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SteerPool *SteerPoolCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SteerPool *SteerPoolSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SteerPool.Contract.Allowance(&_SteerPool.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SteerPool.Contract.Allowance(&_SteerPool.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SteerPool *SteerPoolCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SteerPool *SteerPoolSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SteerPool.Contract.BalanceOf(&_SteerPool.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SteerPool.Contract.BalanceOf(&_SteerPool.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SteerPool *SteerPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SteerPool *SteerPoolSession) Decimals() (uint8, error) {
	return _SteerPool.Contract.Decimals(&_SteerPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SteerPool *SteerPoolCallerSession) Decimals() (uint8, error) {
	return _SteerPool.Contract.Decimals(&_SteerPool.CallOpts)
}

// FeeDetails is a free data retrieval call binding the contract method 0x1f2c5443.
//
// Solidity: function feeDetails() view returns(uint256, address[], string[], uint256[])
func (_SteerPool *SteerPoolCaller) FeeDetails(opts *bind.CallOpts) (*big.Int, []common.Address, []string, []*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "feeDetails")

	if err != nil {
		return *new(*big.Int), *new([]common.Address), *new([]string), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new([]string)).(*[]string)
	out3 := *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, out2, out3, err

}

// FeeDetails is a free data retrieval call binding the contract method 0x1f2c5443.
//
// Solidity: function feeDetails() view returns(uint256, address[], string[], uint256[])
func (_SteerPool *SteerPoolSession) FeeDetails() (*big.Int, []common.Address, []string, []*big.Int, error) {
	return _SteerPool.Contract.FeeDetails(&_SteerPool.CallOpts)
}

// FeeDetails is a free data retrieval call binding the contract method 0x1f2c5443.
//
// Solidity: function feeDetails() view returns(uint256, address[], string[], uint256[])
func (_SteerPool *SteerPoolCallerSession) FeeDetails() (*big.Int, []common.Address, []string, []*big.Int, error) {
	return _SteerPool.Contract.FeeDetails(&_SteerPool.CallOpts)
}

// GetPositions is a free data retrieval call binding the contract method 0x80275860.
//
// Solidity: function getPositions() view returns(int24[], int24[], uint16[])
func (_SteerPool *SteerPoolCaller) GetPositions(opts *bind.CallOpts) ([]*big.Int, []*big.Int, []uint16, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "getPositions")

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	out2 := *abi.ConvertType(out[2], new([]uint16)).(*[]uint16)

	return out0, out1, out2, err

}

// GetPositions is a free data retrieval call binding the contract method 0x80275860.
//
// Solidity: function getPositions() view returns(int24[], int24[], uint16[])
func (_SteerPool *SteerPoolSession) GetPositions() ([]*big.Int, []*big.Int, []uint16, error) {
	return _SteerPool.Contract.GetPositions(&_SteerPool.CallOpts)
}

// GetPositions is a free data retrieval call binding the contract method 0x80275860.
//
// Solidity: function getPositions() view returns(int24[], int24[], uint16[])
func (_SteerPool *SteerPoolCallerSession) GetPositions() ([]*big.Int, []*big.Int, []uint16, error) {
	return _SteerPool.Contract.GetPositions(&_SteerPool.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SteerPool *SteerPoolCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SteerPool *SteerPoolSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SteerPool.Contract.GetRoleAdmin(&_SteerPool.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SteerPool *SteerPoolCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SteerPool.Contract.GetRoleAdmin(&_SteerPool.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SteerPool *SteerPoolCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SteerPool *SteerPoolSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _SteerPool.Contract.GetRoleMember(&_SteerPool.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SteerPool *SteerPoolCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _SteerPool.Contract.GetRoleMember(&_SteerPool.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SteerPool *SteerPoolCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SteerPool *SteerPoolSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _SteerPool.Contract.GetRoleMemberCount(&_SteerPool.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _SteerPool.Contract.GetRoleMemberCount(&_SteerPool.CallOpts, role)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_SteerPool *SteerPoolCaller) GetTotalAmounts(opts *bind.CallOpts) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "getTotalAmounts")

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
func (_SteerPool *SteerPoolSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _SteerPool.Contract.GetTotalAmounts(&_SteerPool.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_SteerPool *SteerPoolCallerSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _SteerPool.Contract.GetTotalAmounts(&_SteerPool.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SteerPool *SteerPoolCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SteerPool *SteerPoolSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SteerPool.Contract.HasRole(&_SteerPool.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SteerPool *SteerPoolCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SteerPool.Contract.HasRole(&_SteerPool.CallOpts, role, account)
}

// MaxTickChange is a free data retrieval call binding the contract method 0x1053f871.
//
// Solidity: function maxTickChange() view returns(int24)
func (_SteerPool *SteerPoolCaller) MaxTickChange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "maxTickChange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTickChange is a free data retrieval call binding the contract method 0x1053f871.
//
// Solidity: function maxTickChange() view returns(int24)
func (_SteerPool *SteerPoolSession) MaxTickChange() (*big.Int, error) {
	return _SteerPool.Contract.MaxTickChange(&_SteerPool.CallOpts)
}

// MaxTickChange is a free data retrieval call binding the contract method 0x1053f871.
//
// Solidity: function maxTickChange() view returns(int24)
func (_SteerPool *SteerPoolCallerSession) MaxTickChange() (*big.Int, error) {
	return _SteerPool.Contract.MaxTickChange(&_SteerPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SteerPool *SteerPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SteerPool *SteerPoolSession) Name() (string, error) {
	return _SteerPool.Contract.Name(&_SteerPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SteerPool *SteerPoolCallerSession) Name() (string, error) {
	return _SteerPool.Contract.Name(&_SteerPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SteerPool *SteerPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SteerPool *SteerPoolSession) Paused() (bool, error) {
	return _SteerPool.Contract.Paused(&_SteerPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SteerPool *SteerPoolCallerSession) Paused() (bool, error) {
	return _SteerPool.Contract.Paused(&_SteerPool.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SteerPool *SteerPoolCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SteerPool *SteerPoolSession) Pool() (common.Address, error) {
	return _SteerPool.Contract.Pool(&_SteerPool.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SteerPool *SteerPoolCallerSession) Pool() (common.Address, error) {
	return _SteerPool.Contract.Pool(&_SteerPool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SteerPool *SteerPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SteerPool *SteerPoolSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SteerPool.Contract.SupportsInterface(&_SteerPool.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SteerPool *SteerPoolCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SteerPool.Contract.SupportsInterface(&_SteerPool.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SteerPool *SteerPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SteerPool *SteerPoolSession) Symbol() (string, error) {
	return _SteerPool.Contract.Symbol(&_SteerPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SteerPool *SteerPoolCallerSession) Symbol() (string, error) {
	return _SteerPool.Contract.Symbol(&_SteerPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SteerPool *SteerPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SteerPool *SteerPoolSession) Token0() (common.Address, error) {
	return _SteerPool.Contract.Token0(&_SteerPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SteerPool *SteerPoolCallerSession) Token0() (common.Address, error) {
	return _SteerPool.Contract.Token0(&_SteerPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SteerPool *SteerPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SteerPool *SteerPoolSession) Token1() (common.Address, error) {
	return _SteerPool.Contract.Token1(&_SteerPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SteerPool *SteerPoolCallerSession) Token1() (common.Address, error) {
	return _SteerPool.Contract.Token1(&_SteerPool.CallOpts)
}

// TotalFees0 is a free data retrieval call binding the contract method 0x80814ffb.
//
// Solidity: function totalFees0() view returns(uint256)
func (_SteerPool *SteerPoolCaller) TotalFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "totalFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFees0 is a free data retrieval call binding the contract method 0x80814ffb.
//
// Solidity: function totalFees0() view returns(uint256)
func (_SteerPool *SteerPoolSession) TotalFees0() (*big.Int, error) {
	return _SteerPool.Contract.TotalFees0(&_SteerPool.CallOpts)
}

// TotalFees0 is a free data retrieval call binding the contract method 0x80814ffb.
//
// Solidity: function totalFees0() view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) TotalFees0() (*big.Int, error) {
	return _SteerPool.Contract.TotalFees0(&_SteerPool.CallOpts)
}

// TotalFees1 is a free data retrieval call binding the contract method 0xef095f8d.
//
// Solidity: function totalFees1() view returns(uint256)
func (_SteerPool *SteerPoolCaller) TotalFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "totalFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFees1 is a free data retrieval call binding the contract method 0xef095f8d.
//
// Solidity: function totalFees1() view returns(uint256)
func (_SteerPool *SteerPoolSession) TotalFees1() (*big.Int, error) {
	return _SteerPool.Contract.TotalFees1(&_SteerPool.CallOpts)
}

// TotalFees1 is a free data retrieval call binding the contract method 0xef095f8d.
//
// Solidity: function totalFees1() view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) TotalFees1() (*big.Int, error) {
	return _SteerPool.Contract.TotalFees1(&_SteerPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SteerPool *SteerPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SteerPool *SteerPoolSession) TotalSupply() (*big.Int, error) {
	return _SteerPool.Contract.TotalSupply(&_SteerPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SteerPool *SteerPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _SteerPool.Contract.TotalSupply(&_SteerPool.CallOpts)
}

// TwapInterval is a free data retrieval call binding the contract method 0x3c1d5df0.
//
// Solidity: function twapInterval() view returns(uint32)
func (_SteerPool *SteerPoolCaller) TwapInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SteerPool.contract.Call(opts, &out, "twapInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TwapInterval is a free data retrieval call binding the contract method 0x3c1d5df0.
//
// Solidity: function twapInterval() view returns(uint32)
func (_SteerPool *SteerPoolSession) TwapInterval() (uint32, error) {
	return _SteerPool.Contract.TwapInterval(&_SteerPool.CallOpts)
}

// TwapInterval is a free data retrieval call binding the contract method 0x3c1d5df0.
//
// Solidity: function twapInterval() view returns(uint32)
func (_SteerPool *SteerPoolCallerSession) TwapInterval() (uint32, error) {
	return _SteerPool.Contract.TwapInterval(&_SteerPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SteerPool *SteerPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SteerPool *SteerPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.Approve(&_SteerPool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SteerPool *SteerPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.Approve(&_SteerPool.TransactOpts, spender, amount)
}

// CollectFees is a paid mutator transaction binding the contract method 0x8ec4daf6.
//
// Solidity: function collectFees(string feeIdentifier, uint256 amount0, uint256 amount1) returns()
func (_SteerPool *SteerPoolTransactor) CollectFees(opts *bind.TransactOpts, feeIdentifier string, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "collectFees", feeIdentifier, amount0, amount1)
}

// CollectFees is a paid mutator transaction binding the contract method 0x8ec4daf6.
//
// Solidity: function collectFees(string feeIdentifier, uint256 amount0, uint256 amount1) returns()
func (_SteerPool *SteerPoolSession) CollectFees(feeIdentifier string, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.CollectFees(&_SteerPool.TransactOpts, feeIdentifier, amount0, amount1)
}

// CollectFees is a paid mutator transaction binding the contract method 0x8ec4daf6.
//
// Solidity: function collectFees(string feeIdentifier, uint256 amount0, uint256 amount1) returns()
func (_SteerPool *SteerPoolTransactorSession) CollectFees(feeIdentifier string, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.CollectFees(&_SteerPool.TransactOpts, feeIdentifier, amount0, amount1)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SteerPool *SteerPoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SteerPool *SteerPoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.DecreaseAllowance(&_SteerPool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SteerPool *SteerPoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.DecreaseAllowance(&_SteerPool.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0Used, uint256 amount1Used)
func (_SteerPool *SteerPoolTransactor) Deposit(opts *bind.TransactOpts, amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "deposit", amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0Used, uint256 amount1Used)
func (_SteerPool *SteerPoolSession) Deposit(amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerPool.Contract.Deposit(&_SteerPool.TransactOpts, amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0Used, uint256 amount1Used)
func (_SteerPool *SteerPoolTransactorSession) Deposit(amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerPool.Contract.Deposit(&_SteerPool.TransactOpts, amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xabbffcb9.
//
// Solidity: function emergencyBurn(int24 tickLower, int24 tickUpper, uint128 liquidity) returns(uint256 amount0, uint256 amount1)
func (_SteerPool *SteerPoolTransactor) EmergencyBurn(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "emergencyBurn", tickLower, tickUpper, liquidity)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xabbffcb9.
//
// Solidity: function emergencyBurn(int24 tickLower, int24 tickUpper, uint128 liquidity) returns(uint256 amount0, uint256 amount1)
func (_SteerPool *SteerPoolSession) EmergencyBurn(tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.EmergencyBurn(&_SteerPool.TransactOpts, tickLower, tickUpper, liquidity)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xabbffcb9.
//
// Solidity: function emergencyBurn(int24 tickLower, int24 tickUpper, uint128 liquidity) returns(uint256 amount0, uint256 amount1)
func (_SteerPool *SteerPoolTransactorSession) EmergencyBurn(tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.EmergencyBurn(&_SteerPool.TransactOpts, tickLower, tickUpper, liquidity)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SteerPool *SteerPoolTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SteerPool *SteerPoolSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerPool.Contract.GrantRole(&_SteerPool.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SteerPool *SteerPoolTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerPool.Contract.GrantRole(&_SteerPool.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SteerPool *SteerPoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SteerPool *SteerPoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.IncreaseAllowance(&_SteerPool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SteerPool *SteerPoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.IncreaseAllowance(&_SteerPool.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x246581f7.
//
// Solidity: function initialize(address _vaultManager, address , address _steer, bytes _params) returns()
func (_SteerPool *SteerPoolTransactor) Initialize(opts *bind.TransactOpts, _vaultManager common.Address, arg1 common.Address, _steer common.Address, _params []byte) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "initialize", _vaultManager, arg1, _steer, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0x246581f7.
//
// Solidity: function initialize(address _vaultManager, address , address _steer, bytes _params) returns()
func (_SteerPool *SteerPoolSession) Initialize(_vaultManager common.Address, arg1 common.Address, _steer common.Address, _params []byte) (*types.Transaction, error) {
	return _SteerPool.Contract.Initialize(&_SteerPool.TransactOpts, _vaultManager, arg1, _steer, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0x246581f7.
//
// Solidity: function initialize(address _vaultManager, address , address _steer, bytes _params) returns()
func (_SteerPool *SteerPoolTransactorSession) Initialize(_vaultManager common.Address, arg1 common.Address, _steer common.Address, _params []byte) (*types.Transaction, error) {
	return _SteerPool.Contract.Initialize(&_SteerPool.TransactOpts, _vaultManager, arg1, _steer, _params)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_SteerPool *SteerPoolTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_SteerPool *SteerPoolSession) Migrate() (*types.Transaction, error) {
	return _SteerPool.Contract.Migrate(&_SteerPool.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_SteerPool *SteerPoolTransactorSession) Migrate() (*types.Transaction, error) {
	return _SteerPool.Contract.Migrate(&_SteerPool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SteerPool *SteerPoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SteerPool *SteerPoolSession) Pause() (*types.Transaction, error) {
	return _SteerPool.Contract.Pause(&_SteerPool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SteerPool *SteerPoolTransactorSession) Pause() (*types.Transaction, error) {
	return _SteerPool.Contract.Pause(&_SteerPool.TransactOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_SteerPool *SteerPoolTransactor) Poke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "poke")
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_SteerPool *SteerPoolSession) Poke() (*types.Transaction, error) {
	return _SteerPool.Contract.Poke(&_SteerPool.TransactOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_SteerPool *SteerPoolTransactorSession) Poke() (*types.Transaction, error) {
	return _SteerPool.Contract.Poke(&_SteerPool.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SteerPool *SteerPoolTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SteerPool *SteerPoolSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerPool.Contract.RenounceRole(&_SteerPool.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SteerPool *SteerPoolTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerPool.Contract.RenounceRole(&_SteerPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SteerPool *SteerPoolTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SteerPool *SteerPoolSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerPool.Contract.RevokeRole(&_SteerPool.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SteerPool *SteerPoolTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerPool.Contract.RevokeRole(&_SteerPool.TransactOpts, role, account)
}

// Tend is a paid mutator transaction binding the contract method 0x8f747bb8.
//
// Solidity: function tend(uint256 totalWeight, (int24[],int24[],uint16[]) newPositions, bytes timeSensitiveData) returns()
func (_SteerPool *SteerPoolTransactor) Tend(opts *bind.TransactOpts, totalWeight *big.Int, newPositions KodiakMultiPositionLiquidityManagerLiquidityPositions, timeSensitiveData []byte) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "tend", totalWeight, newPositions, timeSensitiveData)
}

// Tend is a paid mutator transaction binding the contract method 0x8f747bb8.
//
// Solidity: function tend(uint256 totalWeight, (int24[],int24[],uint16[]) newPositions, bytes timeSensitiveData) returns()
func (_SteerPool *SteerPoolSession) Tend(totalWeight *big.Int, newPositions KodiakMultiPositionLiquidityManagerLiquidityPositions, timeSensitiveData []byte) (*types.Transaction, error) {
	return _SteerPool.Contract.Tend(&_SteerPool.TransactOpts, totalWeight, newPositions, timeSensitiveData)
}

// Tend is a paid mutator transaction binding the contract method 0x8f747bb8.
//
// Solidity: function tend(uint256 totalWeight, (int24[],int24[],uint16[]) newPositions, bytes timeSensitiveData) returns()
func (_SteerPool *SteerPoolTransactorSession) Tend(totalWeight *big.Int, newPositions KodiakMultiPositionLiquidityManagerLiquidityPositions, timeSensitiveData []byte) (*types.Transaction, error) {
	return _SteerPool.Contract.Tend(&_SteerPool.TransactOpts, totalWeight, newPositions, timeSensitiveData)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SteerPool *SteerPoolTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SteerPool *SteerPoolSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.Transfer(&_SteerPool.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SteerPool *SteerPoolTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.Transfer(&_SteerPool.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SteerPool *SteerPoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SteerPool *SteerPoolSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.TransferFrom(&_SteerPool.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SteerPool *SteerPoolTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerPool.Contract.TransferFrom(&_SteerPool.TransactOpts, sender, recipient, amount)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes ) returns()
func (_SteerPool *SteerPoolTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "uniswapV3MintCallback", amount0, amount1, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes ) returns()
func (_SteerPool *SteerPoolSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerPool.Contract.UniswapV3MintCallback(&_SteerPool.TransactOpts, amount0, amount1, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes ) returns()
func (_SteerPool *SteerPoolTransactorSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerPool.Contract.UniswapV3MintCallback(&_SteerPool.TransactOpts, amount0, amount1, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Wanted, int256 amount1Wanted, bytes ) returns()
func (_SteerPool *SteerPoolTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Wanted *big.Int, amount1Wanted *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "uniswapV3SwapCallback", amount0Wanted, amount1Wanted, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Wanted, int256 amount1Wanted, bytes ) returns()
func (_SteerPool *SteerPoolSession) UniswapV3SwapCallback(amount0Wanted *big.Int, amount1Wanted *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerPool.Contract.UniswapV3SwapCallback(&_SteerPool.TransactOpts, amount0Wanted, amount1Wanted, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Wanted, int256 amount1Wanted, bytes ) returns()
func (_SteerPool *SteerPoolTransactorSession) UniswapV3SwapCallback(amount0Wanted *big.Int, amount1Wanted *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerPool.Contract.UniswapV3SwapCallback(&_SteerPool.TransactOpts, amount0Wanted, amount1Wanted, arg2)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SteerPool *SteerPoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SteerPool *SteerPoolSession) Unpause() (*types.Transaction, error) {
	return _SteerPool.Contract.Unpause(&_SteerPool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SteerPool *SteerPoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _SteerPool.Contract.Unpause(&_SteerPool.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_SteerPool *SteerPoolTransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerPool.contract.Transact(opts, "withdraw", shares, amount0Min, amount1Min, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_SteerPool *SteerPoolSession) Withdraw(shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerPool.Contract.Withdraw(&_SteerPool.TransactOpts, shares, amount0Min, amount1Min, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_SteerPool *SteerPoolTransactorSession) Withdraw(shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerPool.Contract.Withdraw(&_SteerPool.TransactOpts, shares, amount0Min, amount1Min, to)
}

// SteerPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SteerPool contract.
type SteerPoolApprovalIterator struct {
	Event *SteerPoolApproval // Event containing the contract specifics and raw log

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
func (it *SteerPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolApproval)
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
		it.Event = new(SteerPoolApproval)
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
func (it *SteerPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolApproval represents a Approval event raised by the SteerPool contract.
type SteerPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SteerPool *SteerPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SteerPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SteerPoolApprovalIterator{contract: _SteerPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SteerPool *SteerPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SteerPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolApproval)
				if err := _SteerPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SteerPool *SteerPoolFilterer) ParseApproval(log types.Log) (*SteerPoolApproval, error) {
	event := new(SteerPoolApproval)
	if err := _SteerPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SteerPool contract.
type SteerPoolDepositIterator struct {
	Event *SteerPoolDeposit // Event containing the contract specifics and raw log

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
func (it *SteerPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolDeposit)
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
		it.Event = new(SteerPoolDeposit)
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
func (it *SteerPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolDeposit represents a Deposit event raised by the SteerPool contract.
type SteerPoolDeposit struct {
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
func (_SteerPool *SteerPoolFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*SteerPoolDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SteerPoolDepositIterator{contract: _SteerPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_SteerPool *SteerPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SteerPoolDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolDeposit)
				if err := _SteerPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_SteerPool *SteerPoolFilterer) ParseDeposit(log types.Log) (*SteerPoolDeposit, error) {
	event := new(SteerPoolDeposit)
	if err := _SteerPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerPoolFeesEarnedIterator is returned from FilterFeesEarned and is used to iterate over the raw logs and unpacked data for FeesEarned events raised by the SteerPool contract.
type SteerPoolFeesEarnedIterator struct {
	Event *SteerPoolFeesEarned // Event containing the contract specifics and raw log

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
func (it *SteerPoolFeesEarnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolFeesEarned)
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
		it.Event = new(SteerPoolFeesEarned)
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
func (it *SteerPoolFeesEarnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolFeesEarnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolFeesEarned represents a FeesEarned event raised by the SteerPool contract.
type SteerPoolFeesEarned struct {
	Amount0Earned *big.Int
	Amount1Earned *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeesEarned is a free log retrieval operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 amount0Earned, uint256 amount1Earned)
func (_SteerPool *SteerPoolFilterer) FilterFeesEarned(opts *bind.FilterOpts) (*SteerPoolFeesEarnedIterator, error) {

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return &SteerPoolFeesEarnedIterator{contract: _SteerPool.contract, event: "FeesEarned", logs: logs, sub: sub}, nil
}

// WatchFeesEarned is a free log subscription operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 amount0Earned, uint256 amount1Earned)
func (_SteerPool *SteerPoolFilterer) WatchFeesEarned(opts *bind.WatchOpts, sink chan<- *SteerPoolFeesEarned) (event.Subscription, error) {

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolFeesEarned)
				if err := _SteerPool.contract.UnpackLog(event, "FeesEarned", log); err != nil {
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
// Solidity: event FeesEarned(uint256 amount0Earned, uint256 amount1Earned)
func (_SteerPool *SteerPoolFilterer) ParseFeesEarned(log types.Log) (*SteerPoolFeesEarned, error) {
	event := new(SteerPoolFeesEarned)
	if err := _SteerPool.contract.UnpackLog(event, "FeesEarned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerPoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SteerPool contract.
type SteerPoolPausedIterator struct {
	Event *SteerPoolPaused // Event containing the contract specifics and raw log

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
func (it *SteerPoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolPaused)
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
		it.Event = new(SteerPoolPaused)
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
func (it *SteerPoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolPaused represents a Paused event raised by the SteerPool contract.
type SteerPoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SteerPool *SteerPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*SteerPoolPausedIterator, error) {

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SteerPoolPausedIterator{contract: _SteerPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SteerPool *SteerPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SteerPoolPaused) (event.Subscription, error) {

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolPaused)
				if err := _SteerPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SteerPool *SteerPoolFilterer) ParsePaused(log types.Log) (*SteerPoolPaused, error) {
	event := new(SteerPoolPaused)
	if err := _SteerPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerPoolRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SteerPool contract.
type SteerPoolRoleAdminChangedIterator struct {
	Event *SteerPoolRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SteerPoolRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolRoleAdminChanged)
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
		it.Event = new(SteerPoolRoleAdminChanged)
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
func (it *SteerPoolRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolRoleAdminChanged represents a RoleAdminChanged event raised by the SteerPool contract.
type SteerPoolRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SteerPool *SteerPoolFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SteerPoolRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SteerPoolRoleAdminChangedIterator{contract: _SteerPool.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SteerPool *SteerPoolFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SteerPoolRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolRoleAdminChanged)
				if err := _SteerPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SteerPool *SteerPoolFilterer) ParseRoleAdminChanged(log types.Log) (*SteerPoolRoleAdminChanged, error) {
	event := new(SteerPoolRoleAdminChanged)
	if err := _SteerPool.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerPoolRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SteerPool contract.
type SteerPoolRoleGrantedIterator struct {
	Event *SteerPoolRoleGranted // Event containing the contract specifics and raw log

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
func (it *SteerPoolRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolRoleGranted)
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
		it.Event = new(SteerPoolRoleGranted)
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
func (it *SteerPoolRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolRoleGranted represents a RoleGranted event raised by the SteerPool contract.
type SteerPoolRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SteerPool *SteerPoolFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SteerPoolRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SteerPoolRoleGrantedIterator{contract: _SteerPool.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SteerPool *SteerPoolFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SteerPoolRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolRoleGranted)
				if err := _SteerPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SteerPool *SteerPoolFilterer) ParseRoleGranted(log types.Log) (*SteerPoolRoleGranted, error) {
	event := new(SteerPoolRoleGranted)
	if err := _SteerPool.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerPoolRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SteerPool contract.
type SteerPoolRoleRevokedIterator struct {
	Event *SteerPoolRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SteerPoolRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolRoleRevoked)
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
		it.Event = new(SteerPoolRoleRevoked)
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
func (it *SteerPoolRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolRoleRevoked represents a RoleRevoked event raised by the SteerPool contract.
type SteerPoolRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SteerPool *SteerPoolFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SteerPoolRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SteerPoolRoleRevokedIterator{contract: _SteerPool.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SteerPool *SteerPoolFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SteerPoolRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolRoleRevoked)
				if err := _SteerPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SteerPool *SteerPoolFilterer) ParseRoleRevoked(log types.Log) (*SteerPoolRoleRevoked, error) {
	event := new(SteerPoolRoleRevoked)
	if err := _SteerPool.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerPoolSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the SteerPool contract.
type SteerPoolSnapshotIterator struct {
	Event *SteerPoolSnapshot // Event containing the contract specifics and raw log

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
func (it *SteerPoolSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolSnapshot)
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
		it.Event = new(SteerPoolSnapshot)
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
func (it *SteerPoolSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolSnapshot represents a Snapshot event raised by the SteerPool contract.
type SteerPoolSnapshot struct {
	SqrtPriceX96 *big.Int
	TotalAmount0 *big.Int
	TotalAmount1 *big.Int
	TotalSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x2fd4737bac0995700ade358ae308c92f20ee63884dfbe658bf55d2f99f60f077.
//
// Solidity: event Snapshot(uint160 sqrtPriceX96, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_SteerPool *SteerPoolFilterer) FilterSnapshot(opts *bind.FilterOpts) (*SteerPoolSnapshotIterator, error) {

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &SteerPoolSnapshotIterator{contract: _SteerPool.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x2fd4737bac0995700ade358ae308c92f20ee63884dfbe658bf55d2f99f60f077.
//
// Solidity: event Snapshot(uint160 sqrtPriceX96, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_SteerPool *SteerPoolFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *SteerPoolSnapshot) (event.Subscription, error) {

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolSnapshot)
				if err := _SteerPool.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x2fd4737bac0995700ade358ae308c92f20ee63884dfbe658bf55d2f99f60f077.
//
// Solidity: event Snapshot(uint160 sqrtPriceX96, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_SteerPool *SteerPoolFilterer) ParseSnapshot(log types.Log) (*SteerPoolSnapshot, error) {
	event := new(SteerPoolSnapshot)
	if err := _SteerPool.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SteerPool contract.
type SteerPoolTransferIterator struct {
	Event *SteerPoolTransfer // Event containing the contract specifics and raw log

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
func (it *SteerPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolTransfer)
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
		it.Event = new(SteerPoolTransfer)
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
func (it *SteerPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolTransfer represents a Transfer event raised by the SteerPool contract.
type SteerPoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SteerPool *SteerPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SteerPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SteerPoolTransferIterator{contract: _SteerPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SteerPool *SteerPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SteerPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolTransfer)
				if err := _SteerPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SteerPool *SteerPoolFilterer) ParseTransfer(log types.Log) (*SteerPoolTransfer, error) {
	event := new(SteerPoolTransfer)
	if err := _SteerPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerPoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SteerPool contract.
type SteerPoolUnpausedIterator struct {
	Event *SteerPoolUnpaused // Event containing the contract specifics and raw log

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
func (it *SteerPoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolUnpaused)
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
		it.Event = new(SteerPoolUnpaused)
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
func (it *SteerPoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolUnpaused represents a Unpaused event raised by the SteerPool contract.
type SteerPoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SteerPool *SteerPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SteerPoolUnpausedIterator, error) {

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SteerPoolUnpausedIterator{contract: _SteerPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SteerPool *SteerPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SteerPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolUnpaused)
				if err := _SteerPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SteerPool *SteerPoolFilterer) ParseUnpaused(log types.Log) (*SteerPoolUnpaused, error) {
	event := new(SteerPoolUnpaused)
	if err := _SteerPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SteerPool contract.
type SteerPoolWithdrawIterator struct {
	Event *SteerPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *SteerPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerPoolWithdraw)
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
		it.Event = new(SteerPoolWithdraw)
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
func (it *SteerPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerPoolWithdraw represents a Withdraw event raised by the SteerPool contract.
type SteerPoolWithdraw struct {
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
func (_SteerPool *SteerPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*SteerPoolWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerPool.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SteerPoolWithdrawIterator{contract: _SteerPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_SteerPool *SteerPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SteerPoolWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerPool.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerPoolWithdraw)
				if err := _SteerPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_SteerPool *SteerPoolFilterer) ParseWithdraw(log types.Log) (*SteerPoolWithdraw, error) {
	event := new(SteerPoolWithdraw)
	if err := _SteerPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
