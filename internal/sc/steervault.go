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

// SteerVaultMetaData contains all meta data concerning the SteerVault contract.
var SteerVaultMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Earned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Earned\",\"type\":\"uint256\"}],\"name\":\"FeesEarned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STEER_FRACTION_OF_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STRATEGIST_FRACTION_OF_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"accruedFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"accruedFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedSteerFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedSteerFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedStrategistFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedStrategistFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"feeIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"collectFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Used\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Used\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"name\":\"emergencyBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositions\",\"outputs\":[{\"internalType\":\"int24[]\",\"name\":\"\",\"type\":\"int24[]\"},{\"internalType\":\"int24[]\",\"name\":\"\",\"type\":\"int24[]\"},{\"internalType\":\"uint16[]\",\"name\":\"\",\"type\":\"uint16[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_steer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTickChange\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"int24[]\",\"name\":\"lowerTick\",\"type\":\"int24[]\"},{\"internalType\":\"int24[]\",\"name\":\"upperTick\",\"type\":\"int24[]\"},{\"internalType\":\"uint16[]\",\"name\":\"relativeWeight\",\"type\":\"uint16[]\"}],\"internalType\":\"structKodiakMultiPositionLiquidityManager.LiquidityPositions\",\"name\":\"newPositions\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"timeSensitiveData\",\"type\":\"bytes\"}],\"name\":\"tend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFees0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFees1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twapInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Wanted\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Wanted\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SteerVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use SteerVaultMetaData.ABI instead.
var SteerVaultABI = SteerVaultMetaData.ABI

// SteerVault is an auto generated Go binding around an Ethereum contract.
type SteerVault struct {
	SteerVaultCaller     // Read-only binding to the contract
	SteerVaultTransactor // Write-only binding to the contract
	SteerVaultFilterer   // Log filterer for contract events
}

// SteerVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type SteerVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SteerVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SteerVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SteerVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SteerVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SteerVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SteerVaultSession struct {
	Contract     *SteerVault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SteerVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SteerVaultCallerSession struct {
	Contract *SteerVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SteerVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SteerVaultTransactorSession struct {
	Contract     *SteerVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SteerVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type SteerVaultRaw struct {
	Contract *SteerVault // Generic contract binding to access the raw methods on
}

// SteerVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SteerVaultCallerRaw struct {
	Contract *SteerVaultCaller // Generic read-only contract binding to access the raw methods on
}

// SteerVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SteerVaultTransactorRaw struct {
	Contract *SteerVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSteerVault creates a new instance of SteerVault, bound to a specific deployed contract.
func NewSteerVault(address common.Address, backend bind.ContractBackend) (*SteerVault, error) {
	contract, err := bindSteerVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SteerVault{SteerVaultCaller: SteerVaultCaller{contract: contract}, SteerVaultTransactor: SteerVaultTransactor{contract: contract}, SteerVaultFilterer: SteerVaultFilterer{contract: contract}}, nil
}

// NewSteerVaultCaller creates a new read-only instance of SteerVault, bound to a specific deployed contract.
func NewSteerVaultCaller(address common.Address, caller bind.ContractCaller) (*SteerVaultCaller, error) {
	contract, err := bindSteerVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SteerVaultCaller{contract: contract}, nil
}

// NewSteerVaultTransactor creates a new write-only instance of SteerVault, bound to a specific deployed contract.
func NewSteerVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*SteerVaultTransactor, error) {
	contract, err := bindSteerVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SteerVaultTransactor{contract: contract}, nil
}

// NewSteerVaultFilterer creates a new log filterer instance of SteerVault, bound to a specific deployed contract.
func NewSteerVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*SteerVaultFilterer, error) {
	contract, err := bindSteerVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SteerVaultFilterer{contract: contract}, nil
}

// bindSteerVault binds a generic wrapper to an already deployed contract.
func bindSteerVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SteerVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SteerVault *SteerVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SteerVault.Contract.SteerVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SteerVault *SteerVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerVault.Contract.SteerVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SteerVault *SteerVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SteerVault.Contract.SteerVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SteerVault *SteerVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SteerVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SteerVault *SteerVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SteerVault *SteerVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SteerVault.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SteerVault *SteerVaultCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SteerVault *SteerVaultSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SteerVault.Contract.DEFAULTADMINROLE(&_SteerVault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SteerVault *SteerVaultCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SteerVault.Contract.DEFAULTADMINROLE(&_SteerVault.CallOpts)
}

// STEERFRACTIONOFFEE is a free data retrieval call binding the contract method 0x1c1f1936.
//
// Solidity: function STEER_FRACTION_OF_FEE() view returns(uint256)
func (_SteerVault *SteerVaultCaller) STEERFRACTIONOFFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "STEER_FRACTION_OF_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STEERFRACTIONOFFEE is a free data retrieval call binding the contract method 0x1c1f1936.
//
// Solidity: function STEER_FRACTION_OF_FEE() view returns(uint256)
func (_SteerVault *SteerVaultSession) STEERFRACTIONOFFEE() (*big.Int, error) {
	return _SteerVault.Contract.STEERFRACTIONOFFEE(&_SteerVault.CallOpts)
}

// STEERFRACTIONOFFEE is a free data retrieval call binding the contract method 0x1c1f1936.
//
// Solidity: function STEER_FRACTION_OF_FEE() view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) STEERFRACTIONOFFEE() (*big.Int, error) {
	return _SteerVault.Contract.STEERFRACTIONOFFEE(&_SteerVault.CallOpts)
}

// STRATEGISTFRACTIONOFFEE is a free data retrieval call binding the contract method 0x30faab2a.
//
// Solidity: function STRATEGIST_FRACTION_OF_FEE() view returns(uint256)
func (_SteerVault *SteerVaultCaller) STRATEGISTFRACTIONOFFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "STRATEGIST_FRACTION_OF_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STRATEGISTFRACTIONOFFEE is a free data retrieval call binding the contract method 0x30faab2a.
//
// Solidity: function STRATEGIST_FRACTION_OF_FEE() view returns(uint256)
func (_SteerVault *SteerVaultSession) STRATEGISTFRACTIONOFFEE() (*big.Int, error) {
	return _SteerVault.Contract.STRATEGISTFRACTIONOFFEE(&_SteerVault.CallOpts)
}

// STRATEGISTFRACTIONOFFEE is a free data retrieval call binding the contract method 0x30faab2a.
//
// Solidity: function STRATEGIST_FRACTION_OF_FEE() view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) STRATEGISTFRACTIONOFFEE() (*big.Int, error) {
	return _SteerVault.Contract.STRATEGISTFRACTIONOFFEE(&_SteerVault.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_SteerVault *SteerVaultCaller) TOTALFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "TOTAL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_SteerVault *SteerVaultSession) TOTALFEE() (*big.Int, error) {
	return _SteerVault.Contract.TOTALFEE(&_SteerVault.CallOpts)
}

// TOTALFEE is a free data retrieval call binding the contract method 0x63db7eae.
//
// Solidity: function TOTAL_FEE() view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) TOTALFEE() (*big.Int, error) {
	return _SteerVault.Contract.TOTALFEE(&_SteerVault.CallOpts)
}

// AccruedFees0 is a free data retrieval call binding the contract method 0x3bec034a.
//
// Solidity: function accruedFees0(string ) view returns(uint256)
func (_SteerVault *SteerVaultCaller) AccruedFees0(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "accruedFees0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedFees0 is a free data retrieval call binding the contract method 0x3bec034a.
//
// Solidity: function accruedFees0(string ) view returns(uint256)
func (_SteerVault *SteerVaultSession) AccruedFees0(arg0 string) (*big.Int, error) {
	return _SteerVault.Contract.AccruedFees0(&_SteerVault.CallOpts, arg0)
}

// AccruedFees0 is a free data retrieval call binding the contract method 0x3bec034a.
//
// Solidity: function accruedFees0(string ) view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) AccruedFees0(arg0 string) (*big.Int, error) {
	return _SteerVault.Contract.AccruedFees0(&_SteerVault.CallOpts, arg0)
}

// AccruedFees1 is a free data retrieval call binding the contract method 0xb7ab45b6.
//
// Solidity: function accruedFees1(string ) view returns(uint256)
func (_SteerVault *SteerVaultCaller) AccruedFees1(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "accruedFees1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedFees1 is a free data retrieval call binding the contract method 0xb7ab45b6.
//
// Solidity: function accruedFees1(string ) view returns(uint256)
func (_SteerVault *SteerVaultSession) AccruedFees1(arg0 string) (*big.Int, error) {
	return _SteerVault.Contract.AccruedFees1(&_SteerVault.CallOpts, arg0)
}

// AccruedFees1 is a free data retrieval call binding the contract method 0xb7ab45b6.
//
// Solidity: function accruedFees1(string ) view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) AccruedFees1(arg0 string) (*big.Int, error) {
	return _SteerVault.Contract.AccruedFees1(&_SteerVault.CallOpts, arg0)
}

// AccruedSteerFees0 is a free data retrieval call binding the contract method 0xef94c279.
//
// Solidity: function accruedSteerFees0() view returns(uint256)
func (_SteerVault *SteerVaultCaller) AccruedSteerFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "accruedSteerFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedSteerFees0 is a free data retrieval call binding the contract method 0xef94c279.
//
// Solidity: function accruedSteerFees0() view returns(uint256)
func (_SteerVault *SteerVaultSession) AccruedSteerFees0() (*big.Int, error) {
	return _SteerVault.Contract.AccruedSteerFees0(&_SteerVault.CallOpts)
}

// AccruedSteerFees0 is a free data retrieval call binding the contract method 0xef94c279.
//
// Solidity: function accruedSteerFees0() view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) AccruedSteerFees0() (*big.Int, error) {
	return _SteerVault.Contract.AccruedSteerFees0(&_SteerVault.CallOpts)
}

// AccruedSteerFees1 is a free data retrieval call binding the contract method 0x02bcd6ca.
//
// Solidity: function accruedSteerFees1() view returns(uint256)
func (_SteerVault *SteerVaultCaller) AccruedSteerFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "accruedSteerFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedSteerFees1 is a free data retrieval call binding the contract method 0x02bcd6ca.
//
// Solidity: function accruedSteerFees1() view returns(uint256)
func (_SteerVault *SteerVaultSession) AccruedSteerFees1() (*big.Int, error) {
	return _SteerVault.Contract.AccruedSteerFees1(&_SteerVault.CallOpts)
}

// AccruedSteerFees1 is a free data retrieval call binding the contract method 0x02bcd6ca.
//
// Solidity: function accruedSteerFees1() view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) AccruedSteerFees1() (*big.Int, error) {
	return _SteerVault.Contract.AccruedSteerFees1(&_SteerVault.CallOpts)
}

// AccruedStrategistFees0 is a free data retrieval call binding the contract method 0x3f2218f1.
//
// Solidity: function accruedStrategistFees0() view returns(uint256)
func (_SteerVault *SteerVaultCaller) AccruedStrategistFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "accruedStrategistFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedStrategistFees0 is a free data retrieval call binding the contract method 0x3f2218f1.
//
// Solidity: function accruedStrategistFees0() view returns(uint256)
func (_SteerVault *SteerVaultSession) AccruedStrategistFees0() (*big.Int, error) {
	return _SteerVault.Contract.AccruedStrategistFees0(&_SteerVault.CallOpts)
}

// AccruedStrategistFees0 is a free data retrieval call binding the contract method 0x3f2218f1.
//
// Solidity: function accruedStrategistFees0() view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) AccruedStrategistFees0() (*big.Int, error) {
	return _SteerVault.Contract.AccruedStrategistFees0(&_SteerVault.CallOpts)
}

// AccruedStrategistFees1 is a free data retrieval call binding the contract method 0x366e897b.
//
// Solidity: function accruedStrategistFees1() view returns(uint256)
func (_SteerVault *SteerVaultCaller) AccruedStrategistFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "accruedStrategistFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedStrategistFees1 is a free data retrieval call binding the contract method 0x366e897b.
//
// Solidity: function accruedStrategistFees1() view returns(uint256)
func (_SteerVault *SteerVaultSession) AccruedStrategistFees1() (*big.Int, error) {
	return _SteerVault.Contract.AccruedStrategistFees1(&_SteerVault.CallOpts)
}

// AccruedStrategistFees1 is a free data retrieval call binding the contract method 0x366e897b.
//
// Solidity: function accruedStrategistFees1() view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) AccruedStrategistFees1() (*big.Int, error) {
	return _SteerVault.Contract.AccruedStrategistFees1(&_SteerVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SteerVault *SteerVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SteerVault *SteerVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SteerVault.Contract.Allowance(&_SteerVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SteerVault.Contract.Allowance(&_SteerVault.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SteerVault *SteerVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SteerVault *SteerVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SteerVault.Contract.BalanceOf(&_SteerVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SteerVault.Contract.BalanceOf(&_SteerVault.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SteerVault *SteerVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SteerVault *SteerVaultSession) Decimals() (uint8, error) {
	return _SteerVault.Contract.Decimals(&_SteerVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SteerVault *SteerVaultCallerSession) Decimals() (uint8, error) {
	return _SteerVault.Contract.Decimals(&_SteerVault.CallOpts)
}

// FeeDetails is a free data retrieval call binding the contract method 0x1f2c5443.
//
// Solidity: function feeDetails() view returns(uint256, address[], string[], uint256[])
func (_SteerVault *SteerVaultCaller) FeeDetails(opts *bind.CallOpts) (*big.Int, []common.Address, []string, []*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "feeDetails")

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
func (_SteerVault *SteerVaultSession) FeeDetails() (*big.Int, []common.Address, []string, []*big.Int, error) {
	return _SteerVault.Contract.FeeDetails(&_SteerVault.CallOpts)
}

// FeeDetails is a free data retrieval call binding the contract method 0x1f2c5443.
//
// Solidity: function feeDetails() view returns(uint256, address[], string[], uint256[])
func (_SteerVault *SteerVaultCallerSession) FeeDetails() (*big.Int, []common.Address, []string, []*big.Int, error) {
	return _SteerVault.Contract.FeeDetails(&_SteerVault.CallOpts)
}

// GetPositions is a free data retrieval call binding the contract method 0x80275860.
//
// Solidity: function getPositions() view returns(int24[], int24[], uint16[])
func (_SteerVault *SteerVaultCaller) GetPositions(opts *bind.CallOpts) ([]*big.Int, []*big.Int, []uint16, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "getPositions")

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
func (_SteerVault *SteerVaultSession) GetPositions() ([]*big.Int, []*big.Int, []uint16, error) {
	return _SteerVault.Contract.GetPositions(&_SteerVault.CallOpts)
}

// GetPositions is a free data retrieval call binding the contract method 0x80275860.
//
// Solidity: function getPositions() view returns(int24[], int24[], uint16[])
func (_SteerVault *SteerVaultCallerSession) GetPositions() ([]*big.Int, []*big.Int, []uint16, error) {
	return _SteerVault.Contract.GetPositions(&_SteerVault.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SteerVault *SteerVaultCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SteerVault *SteerVaultSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SteerVault.Contract.GetRoleAdmin(&_SteerVault.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SteerVault *SteerVaultCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SteerVault.Contract.GetRoleAdmin(&_SteerVault.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SteerVault *SteerVaultCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SteerVault *SteerVaultSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _SteerVault.Contract.GetRoleMember(&_SteerVault.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_SteerVault *SteerVaultCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _SteerVault.Contract.GetRoleMember(&_SteerVault.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SteerVault *SteerVaultCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SteerVault *SteerVaultSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _SteerVault.Contract.GetRoleMemberCount(&_SteerVault.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _SteerVault.Contract.GetRoleMemberCount(&_SteerVault.CallOpts, role)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_SteerVault *SteerVaultCaller) GetTotalAmounts(opts *bind.CallOpts) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "getTotalAmounts")

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
func (_SteerVault *SteerVaultSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _SteerVault.Contract.GetTotalAmounts(&_SteerVault.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_SteerVault *SteerVaultCallerSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _SteerVault.Contract.GetTotalAmounts(&_SteerVault.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SteerVault *SteerVaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SteerVault *SteerVaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SteerVault.Contract.HasRole(&_SteerVault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SteerVault *SteerVaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SteerVault.Contract.HasRole(&_SteerVault.CallOpts, role, account)
}

// MaxTickChange is a free data retrieval call binding the contract method 0x1053f871.
//
// Solidity: function maxTickChange() view returns(int24)
func (_SteerVault *SteerVaultCaller) MaxTickChange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "maxTickChange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTickChange is a free data retrieval call binding the contract method 0x1053f871.
//
// Solidity: function maxTickChange() view returns(int24)
func (_SteerVault *SteerVaultSession) MaxTickChange() (*big.Int, error) {
	return _SteerVault.Contract.MaxTickChange(&_SteerVault.CallOpts)
}

// MaxTickChange is a free data retrieval call binding the contract method 0x1053f871.
//
// Solidity: function maxTickChange() view returns(int24)
func (_SteerVault *SteerVaultCallerSession) MaxTickChange() (*big.Int, error) {
	return _SteerVault.Contract.MaxTickChange(&_SteerVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SteerVault *SteerVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SteerVault *SteerVaultSession) Name() (string, error) {
	return _SteerVault.Contract.Name(&_SteerVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SteerVault *SteerVaultCallerSession) Name() (string, error) {
	return _SteerVault.Contract.Name(&_SteerVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SteerVault *SteerVaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SteerVault *SteerVaultSession) Paused() (bool, error) {
	return _SteerVault.Contract.Paused(&_SteerVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SteerVault *SteerVaultCallerSession) Paused() (bool, error) {
	return _SteerVault.Contract.Paused(&_SteerVault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SteerVault *SteerVaultCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SteerVault *SteerVaultSession) Pool() (common.Address, error) {
	return _SteerVault.Contract.Pool(&_SteerVault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SteerVault *SteerVaultCallerSession) Pool() (common.Address, error) {
	return _SteerVault.Contract.Pool(&_SteerVault.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SteerVault *SteerVaultCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SteerVault *SteerVaultSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SteerVault.Contract.SupportsInterface(&_SteerVault.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SteerVault *SteerVaultCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SteerVault.Contract.SupportsInterface(&_SteerVault.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SteerVault *SteerVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SteerVault *SteerVaultSession) Symbol() (string, error) {
	return _SteerVault.Contract.Symbol(&_SteerVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SteerVault *SteerVaultCallerSession) Symbol() (string, error) {
	return _SteerVault.Contract.Symbol(&_SteerVault.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SteerVault *SteerVaultCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SteerVault *SteerVaultSession) Token0() (common.Address, error) {
	return _SteerVault.Contract.Token0(&_SteerVault.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SteerVault *SteerVaultCallerSession) Token0() (common.Address, error) {
	return _SteerVault.Contract.Token0(&_SteerVault.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SteerVault *SteerVaultCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SteerVault *SteerVaultSession) Token1() (common.Address, error) {
	return _SteerVault.Contract.Token1(&_SteerVault.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SteerVault *SteerVaultCallerSession) Token1() (common.Address, error) {
	return _SteerVault.Contract.Token1(&_SteerVault.CallOpts)
}

// TotalFees0 is a free data retrieval call binding the contract method 0x80814ffb.
//
// Solidity: function totalFees0() view returns(uint256)
func (_SteerVault *SteerVaultCaller) TotalFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "totalFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFees0 is a free data retrieval call binding the contract method 0x80814ffb.
//
// Solidity: function totalFees0() view returns(uint256)
func (_SteerVault *SteerVaultSession) TotalFees0() (*big.Int, error) {
	return _SteerVault.Contract.TotalFees0(&_SteerVault.CallOpts)
}

// TotalFees0 is a free data retrieval call binding the contract method 0x80814ffb.
//
// Solidity: function totalFees0() view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) TotalFees0() (*big.Int, error) {
	return _SteerVault.Contract.TotalFees0(&_SteerVault.CallOpts)
}

// TotalFees1 is a free data retrieval call binding the contract method 0xef095f8d.
//
// Solidity: function totalFees1() view returns(uint256)
func (_SteerVault *SteerVaultCaller) TotalFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "totalFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFees1 is a free data retrieval call binding the contract method 0xef095f8d.
//
// Solidity: function totalFees1() view returns(uint256)
func (_SteerVault *SteerVaultSession) TotalFees1() (*big.Int, error) {
	return _SteerVault.Contract.TotalFees1(&_SteerVault.CallOpts)
}

// TotalFees1 is a free data retrieval call binding the contract method 0xef095f8d.
//
// Solidity: function totalFees1() view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) TotalFees1() (*big.Int, error) {
	return _SteerVault.Contract.TotalFees1(&_SteerVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SteerVault *SteerVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SteerVault *SteerVaultSession) TotalSupply() (*big.Int, error) {
	return _SteerVault.Contract.TotalSupply(&_SteerVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SteerVault *SteerVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _SteerVault.Contract.TotalSupply(&_SteerVault.CallOpts)
}

// TwapInterval is a free data retrieval call binding the contract method 0x3c1d5df0.
//
// Solidity: function twapInterval() view returns(uint32)
func (_SteerVault *SteerVaultCaller) TwapInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SteerVault.contract.Call(opts, &out, "twapInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TwapInterval is a free data retrieval call binding the contract method 0x3c1d5df0.
//
// Solidity: function twapInterval() view returns(uint32)
func (_SteerVault *SteerVaultSession) TwapInterval() (uint32, error) {
	return _SteerVault.Contract.TwapInterval(&_SteerVault.CallOpts)
}

// TwapInterval is a free data retrieval call binding the contract method 0x3c1d5df0.
//
// Solidity: function twapInterval() view returns(uint32)
func (_SteerVault *SteerVaultCallerSession) TwapInterval() (uint32, error) {
	return _SteerVault.Contract.TwapInterval(&_SteerVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SteerVault *SteerVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SteerVault *SteerVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.Approve(&_SteerVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SteerVault *SteerVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.Approve(&_SteerVault.TransactOpts, spender, amount)
}

// CollectFees is a paid mutator transaction binding the contract method 0x8ec4daf6.
//
// Solidity: function collectFees(string feeIdentifier, uint256 amount0, uint256 amount1) returns()
func (_SteerVault *SteerVaultTransactor) CollectFees(opts *bind.TransactOpts, feeIdentifier string, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "collectFees", feeIdentifier, amount0, amount1)
}

// CollectFees is a paid mutator transaction binding the contract method 0x8ec4daf6.
//
// Solidity: function collectFees(string feeIdentifier, uint256 amount0, uint256 amount1) returns()
func (_SteerVault *SteerVaultSession) CollectFees(feeIdentifier string, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.CollectFees(&_SteerVault.TransactOpts, feeIdentifier, amount0, amount1)
}

// CollectFees is a paid mutator transaction binding the contract method 0x8ec4daf6.
//
// Solidity: function collectFees(string feeIdentifier, uint256 amount0, uint256 amount1) returns()
func (_SteerVault *SteerVaultTransactorSession) CollectFees(feeIdentifier string, amount0 *big.Int, amount1 *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.CollectFees(&_SteerVault.TransactOpts, feeIdentifier, amount0, amount1)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SteerVault *SteerVaultTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SteerVault *SteerVaultSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.DecreaseAllowance(&_SteerVault.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SteerVault *SteerVaultTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.DecreaseAllowance(&_SteerVault.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0Used, uint256 amount1Used)
func (_SteerVault *SteerVaultTransactor) Deposit(opts *bind.TransactOpts, amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "deposit", amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0Used, uint256 amount1Used)
func (_SteerVault *SteerVaultSession) Deposit(amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerVault.Contract.Deposit(&_SteerVault.TransactOpts, amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0Used, uint256 amount1Used)
func (_SteerVault *SteerVaultTransactorSession) Deposit(amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerVault.Contract.Deposit(&_SteerVault.TransactOpts, amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xabbffcb9.
//
// Solidity: function emergencyBurn(int24 tickLower, int24 tickUpper, uint128 liquidity) returns(uint256 amount0, uint256 amount1)
func (_SteerVault *SteerVaultTransactor) EmergencyBurn(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "emergencyBurn", tickLower, tickUpper, liquidity)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xabbffcb9.
//
// Solidity: function emergencyBurn(int24 tickLower, int24 tickUpper, uint128 liquidity) returns(uint256 amount0, uint256 amount1)
func (_SteerVault *SteerVaultSession) EmergencyBurn(tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.EmergencyBurn(&_SteerVault.TransactOpts, tickLower, tickUpper, liquidity)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xabbffcb9.
//
// Solidity: function emergencyBurn(int24 tickLower, int24 tickUpper, uint128 liquidity) returns(uint256 amount0, uint256 amount1)
func (_SteerVault *SteerVaultTransactorSession) EmergencyBurn(tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.EmergencyBurn(&_SteerVault.TransactOpts, tickLower, tickUpper, liquidity)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SteerVault *SteerVaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SteerVault *SteerVaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerVault.Contract.GrantRole(&_SteerVault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SteerVault *SteerVaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerVault.Contract.GrantRole(&_SteerVault.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SteerVault *SteerVaultTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SteerVault *SteerVaultSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.IncreaseAllowance(&_SteerVault.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SteerVault *SteerVaultTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.IncreaseAllowance(&_SteerVault.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x246581f7.
//
// Solidity: function initialize(address _vaultManager, address , address _steer, bytes _params) returns()
func (_SteerVault *SteerVaultTransactor) Initialize(opts *bind.TransactOpts, _vaultManager common.Address, arg1 common.Address, _steer common.Address, _params []byte) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "initialize", _vaultManager, arg1, _steer, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0x246581f7.
//
// Solidity: function initialize(address _vaultManager, address , address _steer, bytes _params) returns()
func (_SteerVault *SteerVaultSession) Initialize(_vaultManager common.Address, arg1 common.Address, _steer common.Address, _params []byte) (*types.Transaction, error) {
	return _SteerVault.Contract.Initialize(&_SteerVault.TransactOpts, _vaultManager, arg1, _steer, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0x246581f7.
//
// Solidity: function initialize(address _vaultManager, address , address _steer, bytes _params) returns()
func (_SteerVault *SteerVaultTransactorSession) Initialize(_vaultManager common.Address, arg1 common.Address, _steer common.Address, _params []byte) (*types.Transaction, error) {
	return _SteerVault.Contract.Initialize(&_SteerVault.TransactOpts, _vaultManager, arg1, _steer, _params)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_SteerVault *SteerVaultTransactor) Migrate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "migrate")
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_SteerVault *SteerVaultSession) Migrate() (*types.Transaction, error) {
	return _SteerVault.Contract.Migrate(&_SteerVault.TransactOpts)
}

// Migrate is a paid mutator transaction binding the contract method 0x8fd3ab80.
//
// Solidity: function migrate() returns()
func (_SteerVault *SteerVaultTransactorSession) Migrate() (*types.Transaction, error) {
	return _SteerVault.Contract.Migrate(&_SteerVault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SteerVault *SteerVaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SteerVault *SteerVaultSession) Pause() (*types.Transaction, error) {
	return _SteerVault.Contract.Pause(&_SteerVault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SteerVault *SteerVaultTransactorSession) Pause() (*types.Transaction, error) {
	return _SteerVault.Contract.Pause(&_SteerVault.TransactOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_SteerVault *SteerVaultTransactor) Poke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "poke")
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_SteerVault *SteerVaultSession) Poke() (*types.Transaction, error) {
	return _SteerVault.Contract.Poke(&_SteerVault.TransactOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_SteerVault *SteerVaultTransactorSession) Poke() (*types.Transaction, error) {
	return _SteerVault.Contract.Poke(&_SteerVault.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SteerVault *SteerVaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SteerVault *SteerVaultSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerVault.Contract.RenounceRole(&_SteerVault.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_SteerVault *SteerVaultTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerVault.Contract.RenounceRole(&_SteerVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SteerVault *SteerVaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SteerVault *SteerVaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerVault.Contract.RevokeRole(&_SteerVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SteerVault *SteerVaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SteerVault.Contract.RevokeRole(&_SteerVault.TransactOpts, role, account)
}

// Tend is a paid mutator transaction binding the contract method 0x8f747bb8.
//
// Solidity: function tend(uint256 totalWeight, (int24[],int24[],uint16[]) newPositions, bytes timeSensitiveData) returns()
func (_SteerVault *SteerVaultTransactor) Tend(opts *bind.TransactOpts, totalWeight *big.Int, newPositions KodiakMultiPositionLiquidityManagerLiquidityPositions, timeSensitiveData []byte) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "tend", totalWeight, newPositions, timeSensitiveData)
}

// Tend is a paid mutator transaction binding the contract method 0x8f747bb8.
//
// Solidity: function tend(uint256 totalWeight, (int24[],int24[],uint16[]) newPositions, bytes timeSensitiveData) returns()
func (_SteerVault *SteerVaultSession) Tend(totalWeight *big.Int, newPositions KodiakMultiPositionLiquidityManagerLiquidityPositions, timeSensitiveData []byte) (*types.Transaction, error) {
	return _SteerVault.Contract.Tend(&_SteerVault.TransactOpts, totalWeight, newPositions, timeSensitiveData)
}

// Tend is a paid mutator transaction binding the contract method 0x8f747bb8.
//
// Solidity: function tend(uint256 totalWeight, (int24[],int24[],uint16[]) newPositions, bytes timeSensitiveData) returns()
func (_SteerVault *SteerVaultTransactorSession) Tend(totalWeight *big.Int, newPositions KodiakMultiPositionLiquidityManagerLiquidityPositions, timeSensitiveData []byte) (*types.Transaction, error) {
	return _SteerVault.Contract.Tend(&_SteerVault.TransactOpts, totalWeight, newPositions, timeSensitiveData)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SteerVault *SteerVaultTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SteerVault *SteerVaultSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.Transfer(&_SteerVault.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SteerVault *SteerVaultTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.Transfer(&_SteerVault.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SteerVault *SteerVaultTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SteerVault *SteerVaultSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.TransferFrom(&_SteerVault.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SteerVault *SteerVaultTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SteerVault.Contract.TransferFrom(&_SteerVault.TransactOpts, sender, recipient, amount)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes ) returns()
func (_SteerVault *SteerVaultTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "uniswapV3MintCallback", amount0, amount1, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes ) returns()
func (_SteerVault *SteerVaultSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerVault.Contract.UniswapV3MintCallback(&_SteerVault.TransactOpts, amount0, amount1, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes ) returns()
func (_SteerVault *SteerVaultTransactorSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerVault.Contract.UniswapV3MintCallback(&_SteerVault.TransactOpts, amount0, amount1, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Wanted, int256 amount1Wanted, bytes ) returns()
func (_SteerVault *SteerVaultTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Wanted *big.Int, amount1Wanted *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "uniswapV3SwapCallback", amount0Wanted, amount1Wanted, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Wanted, int256 amount1Wanted, bytes ) returns()
func (_SteerVault *SteerVaultSession) UniswapV3SwapCallback(amount0Wanted *big.Int, amount1Wanted *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerVault.Contract.UniswapV3SwapCallback(&_SteerVault.TransactOpts, amount0Wanted, amount1Wanted, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Wanted, int256 amount1Wanted, bytes ) returns()
func (_SteerVault *SteerVaultTransactorSession) UniswapV3SwapCallback(amount0Wanted *big.Int, amount1Wanted *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _SteerVault.Contract.UniswapV3SwapCallback(&_SteerVault.TransactOpts, amount0Wanted, amount1Wanted, arg2)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SteerVault *SteerVaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SteerVault *SteerVaultSession) Unpause() (*types.Transaction, error) {
	return _SteerVault.Contract.Unpause(&_SteerVault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SteerVault *SteerVaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _SteerVault.Contract.Unpause(&_SteerVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_SteerVault *SteerVaultTransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerVault.contract.Transact(opts, "withdraw", shares, amount0Min, amount1Min, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_SteerVault *SteerVaultSession) Withdraw(shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerVault.Contract.Withdraw(&_SteerVault.TransactOpts, shares, amount0Min, amount1Min, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_SteerVault *SteerVaultTransactorSession) Withdraw(shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _SteerVault.Contract.Withdraw(&_SteerVault.TransactOpts, shares, amount0Min, amount1Min, to)
}

// SteerVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SteerVault contract.
type SteerVaultApprovalIterator struct {
	Event *SteerVaultApproval // Event containing the contract specifics and raw log

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
func (it *SteerVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultApproval)
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
		it.Event = new(SteerVaultApproval)
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
func (it *SteerVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultApproval represents a Approval event raised by the SteerVault contract.
type SteerVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SteerVault *SteerVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SteerVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SteerVaultApprovalIterator{contract: _SteerVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SteerVault *SteerVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SteerVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultApproval)
				if err := _SteerVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParseApproval(log types.Log) (*SteerVaultApproval, error) {
	event := new(SteerVaultApproval)
	if err := _SteerVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SteerVault contract.
type SteerVaultDepositIterator struct {
	Event *SteerVaultDeposit // Event containing the contract specifics and raw log

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
func (it *SteerVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultDeposit)
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
		it.Event = new(SteerVaultDeposit)
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
func (it *SteerVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultDeposit represents a Deposit event raised by the SteerVault contract.
type SteerVaultDeposit struct {
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
func (_SteerVault *SteerVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*SteerVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SteerVaultDepositIterator{contract: _SteerVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_SteerVault *SteerVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SteerVaultDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultDeposit)
				if err := _SteerVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParseDeposit(log types.Log) (*SteerVaultDeposit, error) {
	event := new(SteerVaultDeposit)
	if err := _SteerVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerVaultFeesEarnedIterator is returned from FilterFeesEarned and is used to iterate over the raw logs and unpacked data for FeesEarned events raised by the SteerVault contract.
type SteerVaultFeesEarnedIterator struct {
	Event *SteerVaultFeesEarned // Event containing the contract specifics and raw log

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
func (it *SteerVaultFeesEarnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultFeesEarned)
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
		it.Event = new(SteerVaultFeesEarned)
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
func (it *SteerVaultFeesEarnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultFeesEarnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultFeesEarned represents a FeesEarned event raised by the SteerVault contract.
type SteerVaultFeesEarned struct {
	Amount0Earned *big.Int
	Amount1Earned *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeesEarned is a free log retrieval operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 amount0Earned, uint256 amount1Earned)
func (_SteerVault *SteerVaultFilterer) FilterFeesEarned(opts *bind.FilterOpts) (*SteerVaultFeesEarnedIterator, error) {

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return &SteerVaultFeesEarnedIterator{contract: _SteerVault.contract, event: "FeesEarned", logs: logs, sub: sub}, nil
}

// WatchFeesEarned is a free log subscription operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 amount0Earned, uint256 amount1Earned)
func (_SteerVault *SteerVaultFilterer) WatchFeesEarned(opts *bind.WatchOpts, sink chan<- *SteerVaultFeesEarned) (event.Subscription, error) {

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultFeesEarned)
				if err := _SteerVault.contract.UnpackLog(event, "FeesEarned", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParseFeesEarned(log types.Log) (*SteerVaultFeesEarned, error) {
	event := new(SteerVaultFeesEarned)
	if err := _SteerVault.contract.UnpackLog(event, "FeesEarned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerVaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SteerVault contract.
type SteerVaultPausedIterator struct {
	Event *SteerVaultPaused // Event containing the contract specifics and raw log

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
func (it *SteerVaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultPaused)
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
		it.Event = new(SteerVaultPaused)
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
func (it *SteerVaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultPaused represents a Paused event raised by the SteerVault contract.
type SteerVaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SteerVault *SteerVaultFilterer) FilterPaused(opts *bind.FilterOpts) (*SteerVaultPausedIterator, error) {

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SteerVaultPausedIterator{contract: _SteerVault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SteerVault *SteerVaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SteerVaultPaused) (event.Subscription, error) {

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultPaused)
				if err := _SteerVault.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParsePaused(log types.Log) (*SteerVaultPaused, error) {
	event := new(SteerVaultPaused)
	if err := _SteerVault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerVaultRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SteerVault contract.
type SteerVaultRoleAdminChangedIterator struct {
	Event *SteerVaultRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SteerVaultRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultRoleAdminChanged)
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
		it.Event = new(SteerVaultRoleAdminChanged)
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
func (it *SteerVaultRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultRoleAdminChanged represents a RoleAdminChanged event raised by the SteerVault contract.
type SteerVaultRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SteerVault *SteerVaultFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SteerVaultRoleAdminChangedIterator, error) {

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

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SteerVaultRoleAdminChangedIterator{contract: _SteerVault.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SteerVault *SteerVaultFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SteerVaultRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultRoleAdminChanged)
				if err := _SteerVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParseRoleAdminChanged(log types.Log) (*SteerVaultRoleAdminChanged, error) {
	event := new(SteerVaultRoleAdminChanged)
	if err := _SteerVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerVaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SteerVault contract.
type SteerVaultRoleGrantedIterator struct {
	Event *SteerVaultRoleGranted // Event containing the contract specifics and raw log

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
func (it *SteerVaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultRoleGranted)
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
		it.Event = new(SteerVaultRoleGranted)
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
func (it *SteerVaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultRoleGranted represents a RoleGranted event raised by the SteerVault contract.
type SteerVaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SteerVault *SteerVaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SteerVaultRoleGrantedIterator, error) {

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

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SteerVaultRoleGrantedIterator{contract: _SteerVault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SteerVault *SteerVaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SteerVaultRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultRoleGranted)
				if err := _SteerVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParseRoleGranted(log types.Log) (*SteerVaultRoleGranted, error) {
	event := new(SteerVaultRoleGranted)
	if err := _SteerVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerVaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SteerVault contract.
type SteerVaultRoleRevokedIterator struct {
	Event *SteerVaultRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SteerVaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultRoleRevoked)
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
		it.Event = new(SteerVaultRoleRevoked)
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
func (it *SteerVaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultRoleRevoked represents a RoleRevoked event raised by the SteerVault contract.
type SteerVaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SteerVault *SteerVaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SteerVaultRoleRevokedIterator, error) {

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

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SteerVaultRoleRevokedIterator{contract: _SteerVault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SteerVault *SteerVaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SteerVaultRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultRoleRevoked)
				if err := _SteerVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParseRoleRevoked(log types.Log) (*SteerVaultRoleRevoked, error) {
	event := new(SteerVaultRoleRevoked)
	if err := _SteerVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerVaultSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the SteerVault contract.
type SteerVaultSnapshotIterator struct {
	Event *SteerVaultSnapshot // Event containing the contract specifics and raw log

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
func (it *SteerVaultSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultSnapshot)
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
		it.Event = new(SteerVaultSnapshot)
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
func (it *SteerVaultSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultSnapshot represents a Snapshot event raised by the SteerVault contract.
type SteerVaultSnapshot struct {
	SqrtPriceX96 *big.Int
	TotalAmount0 *big.Int
	TotalAmount1 *big.Int
	TotalSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x2fd4737bac0995700ade358ae308c92f20ee63884dfbe658bf55d2f99f60f077.
//
// Solidity: event Snapshot(uint160 sqrtPriceX96, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_SteerVault *SteerVaultFilterer) FilterSnapshot(opts *bind.FilterOpts) (*SteerVaultSnapshotIterator, error) {

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &SteerVaultSnapshotIterator{contract: _SteerVault.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x2fd4737bac0995700ade358ae308c92f20ee63884dfbe658bf55d2f99f60f077.
//
// Solidity: event Snapshot(uint160 sqrtPriceX96, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_SteerVault *SteerVaultFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *SteerVaultSnapshot) (event.Subscription, error) {

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultSnapshot)
				if err := _SteerVault.contract.UnpackLog(event, "Snapshot", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParseSnapshot(log types.Log) (*SteerVaultSnapshot, error) {
	event := new(SteerVaultSnapshot)
	if err := _SteerVault.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SteerVault contract.
type SteerVaultTransferIterator struct {
	Event *SteerVaultTransfer // Event containing the contract specifics and raw log

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
func (it *SteerVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultTransfer)
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
		it.Event = new(SteerVaultTransfer)
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
func (it *SteerVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultTransfer represents a Transfer event raised by the SteerVault contract.
type SteerVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SteerVault *SteerVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SteerVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SteerVaultTransferIterator{contract: _SteerVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SteerVault *SteerVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SteerVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultTransfer)
				if err := _SteerVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParseTransfer(log types.Log) (*SteerVaultTransfer, error) {
	event := new(SteerVaultTransfer)
	if err := _SteerVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerVaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SteerVault contract.
type SteerVaultUnpausedIterator struct {
	Event *SteerVaultUnpaused // Event containing the contract specifics and raw log

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
func (it *SteerVaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultUnpaused)
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
		it.Event = new(SteerVaultUnpaused)
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
func (it *SteerVaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultUnpaused represents a Unpaused event raised by the SteerVault contract.
type SteerVaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SteerVault *SteerVaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SteerVaultUnpausedIterator, error) {

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SteerVaultUnpausedIterator{contract: _SteerVault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SteerVault *SteerVaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SteerVaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultUnpaused)
				if err := _SteerVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParseUnpaused(log types.Log) (*SteerVaultUnpaused, error) {
	event := new(SteerVaultUnpaused)
	if err := _SteerVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SteerVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SteerVault contract.
type SteerVaultWithdrawIterator struct {
	Event *SteerVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *SteerVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SteerVaultWithdraw)
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
		it.Event = new(SteerVaultWithdraw)
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
func (it *SteerVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SteerVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SteerVaultWithdraw represents a Withdraw event raised by the SteerVault contract.
type SteerVaultWithdraw struct {
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
func (_SteerVault *SteerVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*SteerVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerVault.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SteerVaultWithdrawIterator{contract: _SteerVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_SteerVault *SteerVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SteerVaultWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SteerVault.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SteerVaultWithdraw)
				if err := _SteerVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_SteerVault *SteerVaultFilterer) ParseWithdraw(log types.Log) (*SteerVaultWithdraw, error) {
	event := new(SteerVaultWithdraw)
	if err := _SteerVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
