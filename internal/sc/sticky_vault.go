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

// SwapData is an auto generated low-level Go binding around an user-defined struct.
type SwapData struct {
	Router       common.Address
	AmountIn     *big.Int
	MinAmountOut *big.Int
	ZeroForOne   bool
	RouteData    []byte
}

// StickyVaultMetaData contains all meta data concerning the StickyVault contract.
var StickyVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AllowanceOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllowanceUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPermit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Permit2AllowanceIsFixedAtInfinity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermitExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TotalSupplyOverflow\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBurned\",\"type\":\"uint128\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesEarned0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesEarned1\",\"type\":\"uint256\"}],\"name\":\"FeesEarned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityMinted\",\"type\":\"uint128\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"PauserSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"compounder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"lowerTick_\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"upperTick_\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityBefore\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidityAfter\",\"type\":\"uint128\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RestrictedMintSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"RouterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"managerFeeBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"managerTreasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"compounderSlippageBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"compounderSlippageInterval\",\"type\":\"uint32\"}],\"name\":\"UpdateManagerParams\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityBurned\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compounderSlippageBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compounderSlippageInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newUpperTick\",\"type\":\"int24\"},{\"internalType\":\"uint160\",\"name\":\"swapThresholdPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint256\",\"name\":\"swapAmountBPS\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"}],\"name\":\"executiveRebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newLowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"newUpperTick\",\"type\":\"int24\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"routeData\",\"type\":\"bytes\"}],\"internalType\":\"structSwapData\",\"name\":\"swapData\",\"type\":\"tuple\"}],\"name\":\"executiveRebalanceWithRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"interval\",\"type\":\"uint32\"}],\"name\":\"getAvgPrice\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"avgSqrtPriceX96\",\"type\":\"uint160\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Max\",\"type\":\"uint256\"}],\"name\":\"getMintAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositionID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"positionID\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnderlyingBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Current\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint160\",\"name\":\"sqrtRatioX96\",\"type\":\"uint160\"}],\"name\":\"getUnderlyingBalancesAtPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Current\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_managerFeeBPS\",\"type\":\"uint16\"},{\"internalType\":\"int24\",\"name\":\"_lowerTick\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_upperTick\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"_manager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_managerTreasury\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isManaged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowerTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerFeeBPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"liquidityMinted\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIUniswapV3Pool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restrictedMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setRestrictedMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stickyVaultFactory\",\"outputs\":[{\"internalType\":\"contractIStickyVaultFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapRouter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"syncToFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Owed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Owed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int16\",\"name\":\"newManagerFeeBPS\",\"type\":\"int16\"},{\"internalType\":\"address\",\"name\":\"newManagerTreasury\",\"type\":\"address\"},{\"internalType\":\"int16\",\"name\":\"newSlippageBPS\",\"type\":\"int16\"},{\"internalType\":\"int32\",\"name\":\"newSlippageInterval\",\"type\":\"int32\"}],\"name\":\"updateManagerParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawManagerBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"slippageBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint160\",\"name\":\"avgSqrtPriceX96\",\"type\":\"uint160\"},{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"}],\"name\":\"worstAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// StickyVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use StickyVaultMetaData.ABI instead.
var StickyVaultABI = StickyVaultMetaData.ABI

// StickyVault is an auto generated Go binding around an Ethereum contract.
type StickyVault struct {
	StickyVaultCaller     // Read-only binding to the contract
	StickyVaultTransactor // Write-only binding to the contract
	StickyVaultFilterer   // Log filterer for contract events
}

// StickyVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type StickyVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StickyVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StickyVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StickyVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StickyVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StickyVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StickyVaultSession struct {
	Contract     *StickyVault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StickyVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StickyVaultCallerSession struct {
	Contract *StickyVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StickyVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StickyVaultTransactorSession struct {
	Contract     *StickyVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StickyVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type StickyVaultRaw struct {
	Contract *StickyVault // Generic contract binding to access the raw methods on
}

// StickyVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StickyVaultCallerRaw struct {
	Contract *StickyVaultCaller // Generic read-only contract binding to access the raw methods on
}

// StickyVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StickyVaultTransactorRaw struct {
	Contract *StickyVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStickyVault creates a new instance of StickyVault, bound to a specific deployed contract.
func NewStickyVault(address common.Address, backend bind.ContractBackend) (*StickyVault, error) {
	contract, err := bindStickyVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StickyVault{StickyVaultCaller: StickyVaultCaller{contract: contract}, StickyVaultTransactor: StickyVaultTransactor{contract: contract}, StickyVaultFilterer: StickyVaultFilterer{contract: contract}}, nil
}

// NewStickyVaultCaller creates a new read-only instance of StickyVault, bound to a specific deployed contract.
func NewStickyVaultCaller(address common.Address, caller bind.ContractCaller) (*StickyVaultCaller, error) {
	contract, err := bindStickyVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StickyVaultCaller{contract: contract}, nil
}

// NewStickyVaultTransactor creates a new write-only instance of StickyVault, bound to a specific deployed contract.
func NewStickyVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*StickyVaultTransactor, error) {
	contract, err := bindStickyVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StickyVaultTransactor{contract: contract}, nil
}

// NewStickyVaultFilterer creates a new log filterer instance of StickyVault, bound to a specific deployed contract.
func NewStickyVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*StickyVaultFilterer, error) {
	contract, err := bindStickyVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StickyVaultFilterer{contract: contract}, nil
}

// bindStickyVault binds a generic wrapper to an already deployed contract.
func bindStickyVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StickyVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StickyVault *StickyVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StickyVault.Contract.StickyVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StickyVault *StickyVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StickyVault.Contract.StickyVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StickyVault *StickyVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StickyVault.Contract.StickyVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StickyVault *StickyVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StickyVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StickyVault *StickyVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StickyVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StickyVault *StickyVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StickyVault.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_StickyVault *StickyVaultCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_StickyVault *StickyVaultSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StickyVault.Contract.DOMAINSEPARATOR(&_StickyVault.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32 result)
func (_StickyVault *StickyVaultCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StickyVault.Contract.DOMAINSEPARATOR(&_StickyVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_StickyVault *StickyVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_StickyVault *StickyVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StickyVault.Contract.Allowance(&_StickyVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 result)
func (_StickyVault *StickyVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StickyVault.Contract.Allowance(&_StickyVault.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_StickyVault *StickyVaultCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_StickyVault *StickyVaultSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _StickyVault.Contract.BalanceOf(&_StickyVault.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 result)
func (_StickyVault *StickyVaultCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _StickyVault.Contract.BalanceOf(&_StickyVault.CallOpts, owner)
}

// CompounderSlippageBPS is a free data retrieval call binding the contract method 0xdcf0ff28.
//
// Solidity: function compounderSlippageBPS() view returns(uint16)
func (_StickyVault *StickyVaultCaller) CompounderSlippageBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "compounderSlippageBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// CompounderSlippageBPS is a free data retrieval call binding the contract method 0xdcf0ff28.
//
// Solidity: function compounderSlippageBPS() view returns(uint16)
func (_StickyVault *StickyVaultSession) CompounderSlippageBPS() (uint16, error) {
	return _StickyVault.Contract.CompounderSlippageBPS(&_StickyVault.CallOpts)
}

// CompounderSlippageBPS is a free data retrieval call binding the contract method 0xdcf0ff28.
//
// Solidity: function compounderSlippageBPS() view returns(uint16)
func (_StickyVault *StickyVaultCallerSession) CompounderSlippageBPS() (uint16, error) {
	return _StickyVault.Contract.CompounderSlippageBPS(&_StickyVault.CallOpts)
}

// CompounderSlippageInterval is a free data retrieval call binding the contract method 0x38503345.
//
// Solidity: function compounderSlippageInterval() view returns(uint32)
func (_StickyVault *StickyVaultCaller) CompounderSlippageInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "compounderSlippageInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CompounderSlippageInterval is a free data retrieval call binding the contract method 0x38503345.
//
// Solidity: function compounderSlippageInterval() view returns(uint32)
func (_StickyVault *StickyVaultSession) CompounderSlippageInterval() (uint32, error) {
	return _StickyVault.Contract.CompounderSlippageInterval(&_StickyVault.CallOpts)
}

// CompounderSlippageInterval is a free data retrieval call binding the contract method 0x38503345.
//
// Solidity: function compounderSlippageInterval() view returns(uint32)
func (_StickyVault *StickyVaultCallerSession) CompounderSlippageInterval() (uint32, error) {
	return _StickyVault.Contract.CompounderSlippageInterval(&_StickyVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StickyVault *StickyVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StickyVault *StickyVaultSession) Decimals() (uint8, error) {
	return _StickyVault.Contract.Decimals(&_StickyVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StickyVault *StickyVaultCallerSession) Decimals() (uint8, error) {
	return _StickyVault.Contract.Decimals(&_StickyVault.CallOpts)
}

// GetAvgPrice is a free data retrieval call binding the contract method 0x51a2dc92.
//
// Solidity: function getAvgPrice(uint32 interval) view returns(uint160 avgSqrtPriceX96)
func (_StickyVault *StickyVaultCaller) GetAvgPrice(opts *bind.CallOpts, interval uint32) (*big.Int, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "getAvgPrice", interval)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvgPrice is a free data retrieval call binding the contract method 0x51a2dc92.
//
// Solidity: function getAvgPrice(uint32 interval) view returns(uint160 avgSqrtPriceX96)
func (_StickyVault *StickyVaultSession) GetAvgPrice(interval uint32) (*big.Int, error) {
	return _StickyVault.Contract.GetAvgPrice(&_StickyVault.CallOpts, interval)
}

// GetAvgPrice is a free data retrieval call binding the contract method 0x51a2dc92.
//
// Solidity: function getAvgPrice(uint32 interval) view returns(uint160 avgSqrtPriceX96)
func (_StickyVault *StickyVaultCallerSession) GetAvgPrice(interval uint32) (*big.Int, error) {
	return _StickyVault.Contract.GetAvgPrice(&_StickyVault.CallOpts, interval)
}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_StickyVault *StickyVaultCaller) GetMintAmounts(opts *bind.CallOpts, amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "getMintAmounts", amount0Max, amount1Max)

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
func (_StickyVault *StickyVaultSession) GetMintAmounts(amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	return _StickyVault.Contract.GetMintAmounts(&_StickyVault.CallOpts, amount0Max, amount1Max)
}

// GetMintAmounts is a free data retrieval call binding the contract method 0x9894f21a.
//
// Solidity: function getMintAmounts(uint256 amount0Max, uint256 amount1Max) view returns(uint256 amount0, uint256 amount1, uint256 mintAmount)
func (_StickyVault *StickyVaultCallerSession) GetMintAmounts(amount0Max *big.Int, amount1Max *big.Int) (struct {
	Amount0    *big.Int
	Amount1    *big.Int
	MintAmount *big.Int
}, error) {
	return _StickyVault.Contract.GetMintAmounts(&_StickyVault.CallOpts, amount0Max, amount1Max)
}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_StickyVault *StickyVaultCaller) GetPositionID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "getPositionID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_StickyVault *StickyVaultSession) GetPositionID() ([32]byte, error) {
	return _StickyVault.Contract.GetPositionID(&_StickyVault.CallOpts)
}

// GetPositionID is a free data retrieval call binding the contract method 0xdf28408a.
//
// Solidity: function getPositionID() view returns(bytes32 positionID)
func (_StickyVault *StickyVaultCallerSession) GetPositionID() ([32]byte, error) {
	return _StickyVault.Contract.GetPositionID(&_StickyVault.CallOpts)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_StickyVault *StickyVaultCaller) GetUnderlyingBalances(opts *bind.CallOpts) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "getUnderlyingBalances")

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
func (_StickyVault *StickyVaultSession) GetUnderlyingBalances() (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _StickyVault.Contract.GetUnderlyingBalances(&_StickyVault.CallOpts)
}

// GetUnderlyingBalances is a free data retrieval call binding the contract method 0x1322d954.
//
// Solidity: function getUnderlyingBalances() view returns(uint256 amount0Current, uint256 amount1Current)
func (_StickyVault *StickyVaultCallerSession) GetUnderlyingBalances() (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _StickyVault.Contract.GetUnderlyingBalances(&_StickyVault.CallOpts)
}

// GetUnderlyingBalancesAtPrice is a free data retrieval call binding the contract method 0xb670ed7d.
//
// Solidity: function getUnderlyingBalancesAtPrice(uint160 sqrtRatioX96) view returns(uint256 amount0Current, uint256 amount1Current)
func (_StickyVault *StickyVaultCaller) GetUnderlyingBalancesAtPrice(opts *bind.CallOpts, sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "getUnderlyingBalancesAtPrice", sqrtRatioX96)

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
func (_StickyVault *StickyVaultSession) GetUnderlyingBalancesAtPrice(sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _StickyVault.Contract.GetUnderlyingBalancesAtPrice(&_StickyVault.CallOpts, sqrtRatioX96)
}

// GetUnderlyingBalancesAtPrice is a free data retrieval call binding the contract method 0xb670ed7d.
//
// Solidity: function getUnderlyingBalancesAtPrice(uint160 sqrtRatioX96) view returns(uint256 amount0Current, uint256 amount1Current)
func (_StickyVault *StickyVaultCallerSession) GetUnderlyingBalancesAtPrice(sqrtRatioX96 *big.Int) (struct {
	Amount0Current *big.Int
	Amount1Current *big.Int
}, error) {
	return _StickyVault.Contract.GetUnderlyingBalancesAtPrice(&_StickyVault.CallOpts, sqrtRatioX96)
}

// IsManaged is a free data retrieval call binding the contract method 0x5e6663cc.
//
// Solidity: function isManaged() view returns(bool)
func (_StickyVault *StickyVaultCaller) IsManaged(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "isManaged")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManaged is a free data retrieval call binding the contract method 0x5e6663cc.
//
// Solidity: function isManaged() view returns(bool)
func (_StickyVault *StickyVaultSession) IsManaged() (bool, error) {
	return _StickyVault.Contract.IsManaged(&_StickyVault.CallOpts)
}

// IsManaged is a free data retrieval call binding the contract method 0x5e6663cc.
//
// Solidity: function isManaged() view returns(bool)
func (_StickyVault *StickyVaultCallerSession) IsManaged() (bool, error) {
	return _StickyVault.Contract.IsManaged(&_StickyVault.CallOpts)
}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_StickyVault *StickyVaultCaller) LowerTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "lowerTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_StickyVault *StickyVaultSession) LowerTick() (*big.Int, error) {
	return _StickyVault.Contract.LowerTick(&_StickyVault.CallOpts)
}

// LowerTick is a free data retrieval call binding the contract method 0x9b1344ac.
//
// Solidity: function lowerTick() view returns(int24)
func (_StickyVault *StickyVaultCallerSession) LowerTick() (*big.Int, error) {
	return _StickyVault.Contract.LowerTick(&_StickyVault.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_StickyVault *StickyVaultCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_StickyVault *StickyVaultSession) Manager() (common.Address, error) {
	return _StickyVault.Contract.Manager(&_StickyVault.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_StickyVault *StickyVaultCallerSession) Manager() (common.Address, error) {
	return _StickyVault.Contract.Manager(&_StickyVault.CallOpts)
}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_StickyVault *StickyVaultCaller) ManagerBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "managerBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_StickyVault *StickyVaultSession) ManagerBalance0() (*big.Int, error) {
	return _StickyVault.Contract.ManagerBalance0(&_StickyVault.CallOpts)
}

// ManagerBalance0 is a free data retrieval call binding the contract method 0x065756db.
//
// Solidity: function managerBalance0() view returns(uint256)
func (_StickyVault *StickyVaultCallerSession) ManagerBalance0() (*big.Int, error) {
	return _StickyVault.Contract.ManagerBalance0(&_StickyVault.CallOpts)
}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_StickyVault *StickyVaultCaller) ManagerBalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "managerBalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_StickyVault *StickyVaultSession) ManagerBalance1() (*big.Int, error) {
	return _StickyVault.Contract.ManagerBalance1(&_StickyVault.CallOpts)
}

// ManagerBalance1 is a free data retrieval call binding the contract method 0x42fb9d44.
//
// Solidity: function managerBalance1() view returns(uint256)
func (_StickyVault *StickyVaultCallerSession) ManagerBalance1() (*big.Int, error) {
	return _StickyVault.Contract.ManagerBalance1(&_StickyVault.CallOpts)
}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_StickyVault *StickyVaultCaller) ManagerFeeBPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "managerFeeBPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_StickyVault *StickyVaultSession) ManagerFeeBPS() (uint16, error) {
	return _StickyVault.Contract.ManagerFeeBPS(&_StickyVault.CallOpts)
}

// ManagerFeeBPS is a free data retrieval call binding the contract method 0xccdf7a02.
//
// Solidity: function managerFeeBPS() view returns(uint16)
func (_StickyVault *StickyVaultCallerSession) ManagerFeeBPS() (uint16, error) {
	return _StickyVault.Contract.ManagerFeeBPS(&_StickyVault.CallOpts)
}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_StickyVault *StickyVaultCaller) ManagerTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "managerTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_StickyVault *StickyVaultSession) ManagerTreasury() (common.Address, error) {
	return _StickyVault.Contract.ManagerTreasury(&_StickyVault.CallOpts)
}

// ManagerTreasury is a free data retrieval call binding the contract method 0xcc95353e.
//
// Solidity: function managerTreasury() view returns(address)
func (_StickyVault *StickyVaultCallerSession) ManagerTreasury() (common.Address, error) {
	return _StickyVault.Contract.ManagerTreasury(&_StickyVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StickyVault *StickyVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StickyVault *StickyVaultSession) Name() (string, error) {
	return _StickyVault.Contract.Name(&_StickyVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StickyVault *StickyVaultCallerSession) Name() (string, error) {
	return _StickyVault.Contract.Name(&_StickyVault.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_StickyVault *StickyVaultCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_StickyVault *StickyVaultSession) Nonces(owner common.Address) (*big.Int, error) {
	return _StickyVault.Contract.Nonces(&_StickyVault.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256 result)
func (_StickyVault *StickyVaultCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _StickyVault.Contract.Nonces(&_StickyVault.CallOpts, owner)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StickyVault *StickyVaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StickyVault *StickyVaultSession) Paused() (bool, error) {
	return _StickyVault.Contract.Paused(&_StickyVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StickyVault *StickyVaultCallerSession) Paused() (bool, error) {
	return _StickyVault.Contract.Paused(&_StickyVault.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x0a1289ad.
//
// Solidity: function pauser(address ) view returns(bool)
func (_StickyVault *StickyVaultCaller) Pauser(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "pauser", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x0a1289ad.
//
// Solidity: function pauser(address ) view returns(bool)
func (_StickyVault *StickyVaultSession) Pauser(arg0 common.Address) (bool, error) {
	return _StickyVault.Contract.Pauser(&_StickyVault.CallOpts, arg0)
}

// Pauser is a free data retrieval call binding the contract method 0x0a1289ad.
//
// Solidity: function pauser(address ) view returns(bool)
func (_StickyVault *StickyVaultCallerSession) Pauser(arg0 common.Address) (bool, error) {
	return _StickyVault.Contract.Pauser(&_StickyVault.CallOpts, arg0)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_StickyVault *StickyVaultCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_StickyVault *StickyVaultSession) Pool() (common.Address, error) {
	return _StickyVault.Contract.Pool(&_StickyVault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_StickyVault *StickyVaultCallerSession) Pool() (common.Address, error) {
	return _StickyVault.Contract.Pool(&_StickyVault.CallOpts)
}

// RestrictedMint is a free data retrieval call binding the contract method 0xc7a4ae17.
//
// Solidity: function restrictedMint() view returns(bool)
func (_StickyVault *StickyVaultCaller) RestrictedMint(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "restrictedMint")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RestrictedMint is a free data retrieval call binding the contract method 0xc7a4ae17.
//
// Solidity: function restrictedMint() view returns(bool)
func (_StickyVault *StickyVaultSession) RestrictedMint() (bool, error) {
	return _StickyVault.Contract.RestrictedMint(&_StickyVault.CallOpts)
}

// RestrictedMint is a free data retrieval call binding the contract method 0xc7a4ae17.
//
// Solidity: function restrictedMint() view returns(bool)
func (_StickyVault *StickyVaultCallerSession) RestrictedMint() (bool, error) {
	return _StickyVault.Contract.RestrictedMint(&_StickyVault.CallOpts)
}

// StickyVaultFactory is a free data retrieval call binding the contract method 0x34aa394b.
//
// Solidity: function stickyVaultFactory() view returns(address)
func (_StickyVault *StickyVaultCaller) StickyVaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "stickyVaultFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StickyVaultFactory is a free data retrieval call binding the contract method 0x34aa394b.
//
// Solidity: function stickyVaultFactory() view returns(address)
func (_StickyVault *StickyVaultSession) StickyVaultFactory() (common.Address, error) {
	return _StickyVault.Contract.StickyVaultFactory(&_StickyVault.CallOpts)
}

// StickyVaultFactory is a free data retrieval call binding the contract method 0x34aa394b.
//
// Solidity: function stickyVaultFactory() view returns(address)
func (_StickyVault *StickyVaultCallerSession) StickyVaultFactory() (common.Address, error) {
	return _StickyVault.Contract.StickyVaultFactory(&_StickyVault.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0x99e7061b.
//
// Solidity: function swapRouter(address ) view returns(bool)
func (_StickyVault *StickyVaultCaller) SwapRouter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "swapRouter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0x99e7061b.
//
// Solidity: function swapRouter(address ) view returns(bool)
func (_StickyVault *StickyVaultSession) SwapRouter(arg0 common.Address) (bool, error) {
	return _StickyVault.Contract.SwapRouter(&_StickyVault.CallOpts, arg0)
}

// SwapRouter is a free data retrieval call binding the contract method 0x99e7061b.
//
// Solidity: function swapRouter(address ) view returns(bool)
func (_StickyVault *StickyVaultCallerSession) SwapRouter(arg0 common.Address) (bool, error) {
	return _StickyVault.Contract.SwapRouter(&_StickyVault.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StickyVault *StickyVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StickyVault *StickyVaultSession) Symbol() (string, error) {
	return _StickyVault.Contract.Symbol(&_StickyVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StickyVault *StickyVaultCallerSession) Symbol() (string, error) {
	return _StickyVault.Contract.Symbol(&_StickyVault.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_StickyVault *StickyVaultCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_StickyVault *StickyVaultSession) Token0() (common.Address, error) {
	return _StickyVault.Contract.Token0(&_StickyVault.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_StickyVault *StickyVaultCallerSession) Token0() (common.Address, error) {
	return _StickyVault.Contract.Token0(&_StickyVault.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_StickyVault *StickyVaultCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_StickyVault *StickyVaultSession) Token1() (common.Address, error) {
	return _StickyVault.Contract.Token1(&_StickyVault.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_StickyVault *StickyVaultCallerSession) Token1() (common.Address, error) {
	return _StickyVault.Contract.Token1(&_StickyVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_StickyVault *StickyVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_StickyVault *StickyVaultSession) TotalSupply() (*big.Int, error) {
	return _StickyVault.Contract.TotalSupply(&_StickyVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 result)
func (_StickyVault *StickyVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _StickyVault.Contract.TotalSupply(&_StickyVault.CallOpts)
}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_StickyVault *StickyVaultCaller) UpperTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "upperTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_StickyVault *StickyVaultSession) UpperTick() (*big.Int, error) {
	return _StickyVault.Contract.UpperTick(&_StickyVault.CallOpts)
}

// UpperTick is a free data retrieval call binding the contract method 0x727dd228.
//
// Solidity: function upperTick() view returns(int24)
func (_StickyVault *StickyVaultCallerSession) UpperTick() (*big.Int, error) {
	return _StickyVault.Contract.UpperTick(&_StickyVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StickyVault *StickyVaultCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StickyVault *StickyVaultSession) Version() (string, error) {
	return _StickyVault.Contract.Version(&_StickyVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StickyVault *StickyVaultCallerSession) Version() (string, error) {
	return _StickyVault.Contract.Version(&_StickyVault.CallOpts)
}

// WorstAmountOut is a free data retrieval call binding the contract method 0xd8329c35.
//
// Solidity: function worstAmountOut(uint256 amountIn, uint16 slippageBPS, uint160 avgSqrtPriceX96, bool zeroForOne) pure returns(uint256)
func (_StickyVault *StickyVaultCaller) WorstAmountOut(opts *bind.CallOpts, amountIn *big.Int, slippageBPS uint16, avgSqrtPriceX96 *big.Int, zeroForOne bool) (*big.Int, error) {
	var out []interface{}
	err := _StickyVault.contract.Call(opts, &out, "worstAmountOut", amountIn, slippageBPS, avgSqrtPriceX96, zeroForOne)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorstAmountOut is a free data retrieval call binding the contract method 0xd8329c35.
//
// Solidity: function worstAmountOut(uint256 amountIn, uint16 slippageBPS, uint160 avgSqrtPriceX96, bool zeroForOne) pure returns(uint256)
func (_StickyVault *StickyVaultSession) WorstAmountOut(amountIn *big.Int, slippageBPS uint16, avgSqrtPriceX96 *big.Int, zeroForOne bool) (*big.Int, error) {
	return _StickyVault.Contract.WorstAmountOut(&_StickyVault.CallOpts, amountIn, slippageBPS, avgSqrtPriceX96, zeroForOne)
}

// WorstAmountOut is a free data retrieval call binding the contract method 0xd8329c35.
//
// Solidity: function worstAmountOut(uint256 amountIn, uint16 slippageBPS, uint160 avgSqrtPriceX96, bool zeroForOne) pure returns(uint256)
func (_StickyVault *StickyVaultCallerSession) WorstAmountOut(amountIn *big.Int, slippageBPS uint16, avgSqrtPriceX96 *big.Int, zeroForOne bool) (*big.Int, error) {
	return _StickyVault.Contract.WorstAmountOut(&_StickyVault.CallOpts, amountIn, slippageBPS, avgSqrtPriceX96, zeroForOne)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StickyVault *StickyVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StickyVault *StickyVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StickyVault.Contract.Approve(&_StickyVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StickyVault *StickyVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StickyVault.Contract.Approve(&_StickyVault.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_StickyVault *StickyVaultTransactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "burn", burnAmount, receiver)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_StickyVault *StickyVaultSession) Burn(burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StickyVault.Contract.Burn(&_StickyVault.TransactOpts, burnAmount, receiver)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 burnAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityBurned)
func (_StickyVault *StickyVaultTransactorSession) Burn(burnAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StickyVault.Contract.Burn(&_StickyVault.TransactOpts, burnAmount, receiver)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_StickyVault *StickyVaultTransactor) ExecutiveRebalance(opts *bind.TransactOpts, newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "executiveRebalance", newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_StickyVault *StickyVaultSession) ExecutiveRebalance(newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _StickyVault.Contract.ExecutiveRebalance(&_StickyVault.TransactOpts, newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// ExecutiveRebalance is a paid mutator transaction binding the contract method 0x24b8fd1b.
//
// Solidity: function executiveRebalance(int24 newLowerTick, int24 newUpperTick, uint160 swapThresholdPrice, uint256 swapAmountBPS, bool zeroForOne) returns()
func (_StickyVault *StickyVaultTransactorSession) ExecutiveRebalance(newLowerTick *big.Int, newUpperTick *big.Int, swapThresholdPrice *big.Int, swapAmountBPS *big.Int, zeroForOne bool) (*types.Transaction, error) {
	return _StickyVault.Contract.ExecutiveRebalance(&_StickyVault.TransactOpts, newLowerTick, newUpperTick, swapThresholdPrice, swapAmountBPS, zeroForOne)
}

// ExecutiveRebalanceWithRouter is a paid mutator transaction binding the contract method 0xd1c122ea.
//
// Solidity: function executiveRebalanceWithRouter(int24 newLowerTick, int24 newUpperTick, (address,uint256,uint256,bool,bytes) swapData) returns()
func (_StickyVault *StickyVaultTransactor) ExecutiveRebalanceWithRouter(opts *bind.TransactOpts, newLowerTick *big.Int, newUpperTick *big.Int, swapData SwapData) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "executiveRebalanceWithRouter", newLowerTick, newUpperTick, swapData)
}

// ExecutiveRebalanceWithRouter is a paid mutator transaction binding the contract method 0xd1c122ea.
//
// Solidity: function executiveRebalanceWithRouter(int24 newLowerTick, int24 newUpperTick, (address,uint256,uint256,bool,bytes) swapData) returns()
func (_StickyVault *StickyVaultSession) ExecutiveRebalanceWithRouter(newLowerTick *big.Int, newUpperTick *big.Int, swapData SwapData) (*types.Transaction, error) {
	return _StickyVault.Contract.ExecutiveRebalanceWithRouter(&_StickyVault.TransactOpts, newLowerTick, newUpperTick, swapData)
}

// ExecutiveRebalanceWithRouter is a paid mutator transaction binding the contract method 0xd1c122ea.
//
// Solidity: function executiveRebalanceWithRouter(int24 newLowerTick, int24 newUpperTick, (address,uint256,uint256,bool,bytes) swapData) returns()
func (_StickyVault *StickyVaultTransactorSession) ExecutiveRebalanceWithRouter(newLowerTick *big.Int, newUpperTick *big.Int, swapData SwapData) (*types.Transaction, error) {
	return _StickyVault.Contract.ExecutiveRebalanceWithRouter(&_StickyVault.TransactOpts, newLowerTick, newUpperTick, swapData)
}

// Initialize is a paid mutator transaction binding the contract method 0xdb812fd0.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_, address _managerTreasury) returns()
func (_StickyVault *StickyVaultTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address, _managerTreasury common.Address) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "initialize", _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_, _managerTreasury)
}

// Initialize is a paid mutator transaction binding the contract method 0xdb812fd0.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_, address _managerTreasury) returns()
func (_StickyVault *StickyVaultSession) Initialize(_name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address, _managerTreasury common.Address) (*types.Transaction, error) {
	return _StickyVault.Contract.Initialize(&_StickyVault.TransactOpts, _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_, _managerTreasury)
}

// Initialize is a paid mutator transaction binding the contract method 0xdb812fd0.
//
// Solidity: function initialize(string _name, string _symbol, address _pool, uint16 _managerFeeBPS, int24 _lowerTick, int24 _upperTick, address _manager_, address _managerTreasury) returns()
func (_StickyVault *StickyVaultTransactorSession) Initialize(_name string, _symbol string, _pool common.Address, _managerFeeBPS uint16, _lowerTick *big.Int, _upperTick *big.Int, _manager_ common.Address, _managerTreasury common.Address) (*types.Transaction, error) {
	return _StickyVault.Contract.Initialize(&_StickyVault.TransactOpts, _name, _symbol, _pool, _managerFeeBPS, _lowerTick, _upperTick, _manager_, _managerTreasury)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_StickyVault *StickyVaultTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "mint", mintAmount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_StickyVault *StickyVaultSession) Mint(mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StickyVault.Contract.Mint(&_StickyVault.TransactOpts, mintAmount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 mintAmount, address receiver) returns(uint256 amount0, uint256 amount1, uint128 liquidityMinted)
func (_StickyVault *StickyVaultTransactorSession) Mint(mintAmount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StickyVault.Contract.Mint(&_StickyVault.TransactOpts, mintAmount, receiver)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StickyVault *StickyVaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StickyVault *StickyVaultSession) Pause() (*types.Transaction, error) {
	return _StickyVault.Contract.Pause(&_StickyVault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StickyVault *StickyVaultTransactorSession) Pause() (*types.Transaction, error) {
	return _StickyVault.Contract.Pause(&_StickyVault.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_StickyVault *StickyVaultTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_StickyVault *StickyVaultSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StickyVault.Contract.Permit(&_StickyVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_StickyVault *StickyVaultTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _StickyVault.Contract.Permit(&_StickyVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_StickyVault *StickyVaultTransactor) Rebalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "rebalance")
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_StickyVault *StickyVaultSession) Rebalance() (*types.Transaction, error) {
	return _StickyVault.Contract.Rebalance(&_StickyVault.TransactOpts)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_StickyVault *StickyVaultTransactorSession) Rebalance() (*types.Transaction, error) {
	return _StickyVault.Contract.Rebalance(&_StickyVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StickyVault *StickyVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StickyVault *StickyVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _StickyVault.Contract.RenounceOwnership(&_StickyVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StickyVault *StickyVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StickyVault.Contract.RenounceOwnership(&_StickyVault.TransactOpts)
}

// SetPauser is a paid mutator transaction binding the contract method 0x7180c8ca.
//
// Solidity: function setPauser(address _pauser, bool _status) returns()
func (_StickyVault *StickyVaultTransactor) SetPauser(opts *bind.TransactOpts, _pauser common.Address, _status bool) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "setPauser", _pauser, _status)
}

// SetPauser is a paid mutator transaction binding the contract method 0x7180c8ca.
//
// Solidity: function setPauser(address _pauser, bool _status) returns()
func (_StickyVault *StickyVaultSession) SetPauser(_pauser common.Address, _status bool) (*types.Transaction, error) {
	return _StickyVault.Contract.SetPauser(&_StickyVault.TransactOpts, _pauser, _status)
}

// SetPauser is a paid mutator transaction binding the contract method 0x7180c8ca.
//
// Solidity: function setPauser(address _pauser, bool _status) returns()
func (_StickyVault *StickyVaultTransactorSession) SetPauser(_pauser common.Address, _status bool) (*types.Transaction, error) {
	return _StickyVault.Contract.SetPauser(&_StickyVault.TransactOpts, _pauser, _status)
}

// SetRestrictedMint is a paid mutator transaction binding the contract method 0x67582f4a.
//
// Solidity: function setRestrictedMint(bool _status) returns()
func (_StickyVault *StickyVaultTransactor) SetRestrictedMint(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "setRestrictedMint", _status)
}

// SetRestrictedMint is a paid mutator transaction binding the contract method 0x67582f4a.
//
// Solidity: function setRestrictedMint(bool _status) returns()
func (_StickyVault *StickyVaultSession) SetRestrictedMint(_status bool) (*types.Transaction, error) {
	return _StickyVault.Contract.SetRestrictedMint(&_StickyVault.TransactOpts, _status)
}

// SetRestrictedMint is a paid mutator transaction binding the contract method 0x67582f4a.
//
// Solidity: function setRestrictedMint(bool _status) returns()
func (_StickyVault *StickyVaultTransactorSession) SetRestrictedMint(_status bool) (*types.Transaction, error) {
	return _StickyVault.Contract.SetRestrictedMint(&_StickyVault.TransactOpts, _status)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc3c64674.
//
// Solidity: function setRouter(address _router, bool _status) returns()
func (_StickyVault *StickyVaultTransactor) SetRouter(opts *bind.TransactOpts, _router common.Address, _status bool) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "setRouter", _router, _status)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc3c64674.
//
// Solidity: function setRouter(address _router, bool _status) returns()
func (_StickyVault *StickyVaultSession) SetRouter(_router common.Address, _status bool) (*types.Transaction, error) {
	return _StickyVault.Contract.SetRouter(&_StickyVault.TransactOpts, _router, _status)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc3c64674.
//
// Solidity: function setRouter(address _router, bool _status) returns()
func (_StickyVault *StickyVaultTransactorSession) SetRouter(_router common.Address, _status bool) (*types.Transaction, error) {
	return _StickyVault.Contract.SetRouter(&_StickyVault.TransactOpts, _router, _status)
}

// SyncToFactory is a paid mutator transaction binding the contract method 0x688f7218.
//
// Solidity: function syncToFactory() returns()
func (_StickyVault *StickyVaultTransactor) SyncToFactory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "syncToFactory")
}

// SyncToFactory is a paid mutator transaction binding the contract method 0x688f7218.
//
// Solidity: function syncToFactory() returns()
func (_StickyVault *StickyVaultSession) SyncToFactory() (*types.Transaction, error) {
	return _StickyVault.Contract.SyncToFactory(&_StickyVault.TransactOpts)
}

// SyncToFactory is a paid mutator transaction binding the contract method 0x688f7218.
//
// Solidity: function syncToFactory() returns()
func (_StickyVault *StickyVaultTransactorSession) SyncToFactory() (*types.Transaction, error) {
	return _StickyVault.Contract.SyncToFactory(&_StickyVault.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StickyVault *StickyVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StickyVault *StickyVaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StickyVault.Contract.Transfer(&_StickyVault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StickyVault *StickyVaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StickyVault.Contract.Transfer(&_StickyVault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StickyVault *StickyVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StickyVault *StickyVaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StickyVault.Contract.TransferFrom(&_StickyVault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StickyVault *StickyVaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StickyVault.Contract.TransferFrom(&_StickyVault.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StickyVault *StickyVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StickyVault *StickyVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StickyVault.Contract.TransferOwnership(&_StickyVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StickyVault *StickyVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StickyVault.Contract.TransferOwnership(&_StickyVault.TransactOpts, newOwner)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_StickyVault *StickyVaultTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "uniswapV3MintCallback", amount0Owed, amount1Owed, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_StickyVault *StickyVaultSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _StickyVault.Contract.UniswapV3MintCallback(&_StickyVault.TransactOpts, amount0Owed, amount1Owed, arg2)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0Owed, uint256 amount1Owed, bytes ) returns()
func (_StickyVault *StickyVaultTransactorSession) UniswapV3MintCallback(amount0Owed *big.Int, amount1Owed *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _StickyVault.Contract.UniswapV3MintCallback(&_StickyVault.TransactOpts, amount0Owed, amount1Owed, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_StickyVault *StickyVaultTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_StickyVault *StickyVaultSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _StickyVault.Contract.UniswapV3SwapCallback(&_StickyVault.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_StickyVault *StickyVaultTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _StickyVault.Contract.UniswapV3SwapCallback(&_StickyVault.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StickyVault *StickyVaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StickyVault *StickyVaultSession) Unpause() (*types.Transaction, error) {
	return _StickyVault.Contract.Unpause(&_StickyVault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StickyVault *StickyVaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _StickyVault.Contract.Unpause(&_StickyVault.TransactOpts)
}

// UpdateManagerParams is a paid mutator transaction binding the contract method 0x18e91fb5.
//
// Solidity: function updateManagerParams(int16 newManagerFeeBPS, address newManagerTreasury, int16 newSlippageBPS, int32 newSlippageInterval) returns()
func (_StickyVault *StickyVaultTransactor) UpdateManagerParams(opts *bind.TransactOpts, newManagerFeeBPS int16, newManagerTreasury common.Address, newSlippageBPS int16, newSlippageInterval int32) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "updateManagerParams", newManagerFeeBPS, newManagerTreasury, newSlippageBPS, newSlippageInterval)
}

// UpdateManagerParams is a paid mutator transaction binding the contract method 0x18e91fb5.
//
// Solidity: function updateManagerParams(int16 newManagerFeeBPS, address newManagerTreasury, int16 newSlippageBPS, int32 newSlippageInterval) returns()
func (_StickyVault *StickyVaultSession) UpdateManagerParams(newManagerFeeBPS int16, newManagerTreasury common.Address, newSlippageBPS int16, newSlippageInterval int32) (*types.Transaction, error) {
	return _StickyVault.Contract.UpdateManagerParams(&_StickyVault.TransactOpts, newManagerFeeBPS, newManagerTreasury, newSlippageBPS, newSlippageInterval)
}

// UpdateManagerParams is a paid mutator transaction binding the contract method 0x18e91fb5.
//
// Solidity: function updateManagerParams(int16 newManagerFeeBPS, address newManagerTreasury, int16 newSlippageBPS, int32 newSlippageInterval) returns()
func (_StickyVault *StickyVaultTransactorSession) UpdateManagerParams(newManagerFeeBPS int16, newManagerTreasury common.Address, newSlippageBPS int16, newSlippageInterval int32) (*types.Transaction, error) {
	return _StickyVault.Contract.UpdateManagerParams(&_StickyVault.TransactOpts, newManagerFeeBPS, newManagerTreasury, newSlippageBPS, newSlippageInterval)
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x7ecd6717.
//
// Solidity: function withdrawManagerBalance() returns()
func (_StickyVault *StickyVaultTransactor) WithdrawManagerBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StickyVault.contract.Transact(opts, "withdrawManagerBalance")
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x7ecd6717.
//
// Solidity: function withdrawManagerBalance() returns()
func (_StickyVault *StickyVaultSession) WithdrawManagerBalance() (*types.Transaction, error) {
	return _StickyVault.Contract.WithdrawManagerBalance(&_StickyVault.TransactOpts)
}

// WithdrawManagerBalance is a paid mutator transaction binding the contract method 0x7ecd6717.
//
// Solidity: function withdrawManagerBalance() returns()
func (_StickyVault *StickyVaultTransactorSession) WithdrawManagerBalance() (*types.Transaction, error) {
	return _StickyVault.Contract.WithdrawManagerBalance(&_StickyVault.TransactOpts)
}

// StickyVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StickyVault contract.
type StickyVaultApprovalIterator struct {
	Event *StickyVaultApproval // Event containing the contract specifics and raw log

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
func (it *StickyVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultApproval)
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
		it.Event = new(StickyVaultApproval)
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
func (it *StickyVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultApproval represents a Approval event raised by the StickyVault contract.
type StickyVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_StickyVault *StickyVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StickyVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StickyVaultApprovalIterator{contract: _StickyVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_StickyVault *StickyVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StickyVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultApproval)
				if err := _StickyVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_StickyVault *StickyVaultFilterer) ParseApproval(log types.Log) (*StickyVaultApproval, error) {
	event := new(StickyVaultApproval)
	if err := _StickyVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the StickyVault contract.
type StickyVaultBurnedIterator struct {
	Event *StickyVaultBurned // Event containing the contract specifics and raw log

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
func (it *StickyVaultBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultBurned)
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
		it.Event = new(StickyVaultBurned)
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
func (it *StickyVaultBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultBurned represents a Burned event raised by the StickyVault contract.
type StickyVaultBurned struct {
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
func (_StickyVault *StickyVaultFilterer) FilterBurned(opts *bind.FilterOpts) (*StickyVaultBurnedIterator, error) {

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return &StickyVaultBurnedIterator{contract: _StickyVault.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x7239dff1718b550db7f36cbf69c665cfeb56d0e96b4fb76a5cba712961b65509.
//
// Solidity: event Burned(address receiver, uint256 burnAmount, uint256 amount0Out, uint256 amount1Out, uint128 liquidityBurned)
func (_StickyVault *StickyVaultFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *StickyVaultBurned) (event.Subscription, error) {

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "Burned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultBurned)
				if err := _StickyVault.contract.UnpackLog(event, "Burned", log); err != nil {
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
func (_StickyVault *StickyVaultFilterer) ParseBurned(log types.Log) (*StickyVaultBurned, error) {
	event := new(StickyVaultBurned)
	if err := _StickyVault.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultFeesEarnedIterator is returned from FilterFeesEarned and is used to iterate over the raw logs and unpacked data for FeesEarned events raised by the StickyVault contract.
type StickyVaultFeesEarnedIterator struct {
	Event *StickyVaultFeesEarned // Event containing the contract specifics and raw log

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
func (it *StickyVaultFeesEarnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultFeesEarned)
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
		it.Event = new(StickyVaultFeesEarned)
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
func (it *StickyVaultFeesEarnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultFeesEarnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultFeesEarned represents a FeesEarned event raised by the StickyVault contract.
type StickyVaultFeesEarned struct {
	FeesEarned0 *big.Int
	FeesEarned1 *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeesEarned is a free log retrieval operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 feesEarned0, uint256 feesEarned1)
func (_StickyVault *StickyVaultFilterer) FilterFeesEarned(opts *bind.FilterOpts) (*StickyVaultFeesEarnedIterator, error) {

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return &StickyVaultFeesEarnedIterator{contract: _StickyVault.contract, event: "FeesEarned", logs: logs, sub: sub}, nil
}

// WatchFeesEarned is a free log subscription operation binding the contract event 0xc28ad1de9c0c32e5394ba60323e44d8d9536312236a47231772e448a3e49de42.
//
// Solidity: event FeesEarned(uint256 feesEarned0, uint256 feesEarned1)
func (_StickyVault *StickyVaultFilterer) WatchFeesEarned(opts *bind.WatchOpts, sink chan<- *StickyVaultFeesEarned) (event.Subscription, error) {

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "FeesEarned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultFeesEarned)
				if err := _StickyVault.contract.UnpackLog(event, "FeesEarned", log); err != nil {
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
func (_StickyVault *StickyVaultFilterer) ParseFeesEarned(log types.Log) (*StickyVaultFeesEarned, error) {
	event := new(StickyVaultFeesEarned)
	if err := _StickyVault.contract.UnpackLog(event, "FeesEarned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the StickyVault contract.
type StickyVaultMintedIterator struct {
	Event *StickyVaultMinted // Event containing the contract specifics and raw log

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
func (it *StickyVaultMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultMinted)
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
		it.Event = new(StickyVaultMinted)
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
func (it *StickyVaultMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultMinted represents a Minted event raised by the StickyVault contract.
type StickyVaultMinted struct {
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
func (_StickyVault *StickyVaultFilterer) FilterMinted(opts *bind.FilterOpts) (*StickyVaultMintedIterator, error) {

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &StickyVaultMintedIterator{contract: _StickyVault.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x55801cfe493000b734571da1694b21e7f66b11e8ce9fdaa0524ecb59105e73e7.
//
// Solidity: event Minted(address receiver, uint256 mintAmount, uint256 amount0In, uint256 amount1In, uint128 liquidityMinted)
func (_StickyVault *StickyVaultFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *StickyVaultMinted) (event.Subscription, error) {

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultMinted)
				if err := _StickyVault.contract.UnpackLog(event, "Minted", log); err != nil {
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
func (_StickyVault *StickyVaultFilterer) ParseMinted(log types.Log) (*StickyVaultMinted, error) {
	event := new(StickyVaultMinted)
	if err := _StickyVault.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StickyVault contract.
type StickyVaultOwnershipTransferredIterator struct {
	Event *StickyVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StickyVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultOwnershipTransferred)
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
		it.Event = new(StickyVaultOwnershipTransferred)
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
func (it *StickyVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultOwnershipTransferred represents a OwnershipTransferred event raised by the StickyVault contract.
type StickyVaultOwnershipTransferred struct {
	PreviousManager common.Address
	NewManager      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousManager, address indexed newManager)
func (_StickyVault *StickyVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousManager []common.Address, newManager []common.Address) (*StickyVaultOwnershipTransferredIterator, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "OwnershipTransferred", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &StickyVaultOwnershipTransferredIterator{contract: _StickyVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousManager, address indexed newManager)
func (_StickyVault *StickyVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StickyVaultOwnershipTransferred, previousManager []common.Address, newManager []common.Address) (event.Subscription, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "OwnershipTransferred", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultOwnershipTransferred)
				if err := _StickyVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StickyVault *StickyVaultFilterer) ParseOwnershipTransferred(log types.Log) (*StickyVaultOwnershipTransferred, error) {
	event := new(StickyVaultOwnershipTransferred)
	if err := _StickyVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StickyVault contract.
type StickyVaultPausedIterator struct {
	Event *StickyVaultPaused // Event containing the contract specifics and raw log

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
func (it *StickyVaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultPaused)
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
		it.Event = new(StickyVaultPaused)
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
func (it *StickyVaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultPaused represents a Paused event raised by the StickyVault contract.
type StickyVaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StickyVault *StickyVaultFilterer) FilterPaused(opts *bind.FilterOpts) (*StickyVaultPausedIterator, error) {

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StickyVaultPausedIterator{contract: _StickyVault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StickyVault *StickyVaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StickyVaultPaused) (event.Subscription, error) {

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultPaused)
				if err := _StickyVault.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StickyVault *StickyVaultFilterer) ParsePaused(log types.Log) (*StickyVaultPaused, error) {
	event := new(StickyVaultPaused)
	if err := _StickyVault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultPauserSetIterator is returned from FilterPauserSet and is used to iterate over the raw logs and unpacked data for PauserSet events raised by the StickyVault contract.
type StickyVaultPauserSetIterator struct {
	Event *StickyVaultPauserSet // Event containing the contract specifics and raw log

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
func (it *StickyVaultPauserSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultPauserSet)
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
		it.Event = new(StickyVaultPauserSet)
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
func (it *StickyVaultPauserSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultPauserSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultPauserSet represents a PauserSet event raised by the StickyVault contract.
type StickyVaultPauserSet struct {
	Pauser common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPauserSet is a free log retrieval operation binding the contract event 0xa11b5803b8a35081b8f993e0dee5bc30301a3d83f644e5ab2ff39f972f0a807f.
//
// Solidity: event PauserSet(address indexed pauser, bool status)
func (_StickyVault *StickyVaultFilterer) FilterPauserSet(opts *bind.FilterOpts, pauser []common.Address) (*StickyVaultPauserSetIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "PauserSet", pauserRule)
	if err != nil {
		return nil, err
	}
	return &StickyVaultPauserSetIterator{contract: _StickyVault.contract, event: "PauserSet", logs: logs, sub: sub}, nil
}

// WatchPauserSet is a free log subscription operation binding the contract event 0xa11b5803b8a35081b8f993e0dee5bc30301a3d83f644e5ab2ff39f972f0a807f.
//
// Solidity: event PauserSet(address indexed pauser, bool status)
func (_StickyVault *StickyVaultFilterer) WatchPauserSet(opts *bind.WatchOpts, sink chan<- *StickyVaultPauserSet, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "PauserSet", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultPauserSet)
				if err := _StickyVault.contract.UnpackLog(event, "PauserSet", log); err != nil {
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

// ParsePauserSet is a log parse operation binding the contract event 0xa11b5803b8a35081b8f993e0dee5bc30301a3d83f644e5ab2ff39f972f0a807f.
//
// Solidity: event PauserSet(address indexed pauser, bool status)
func (_StickyVault *StickyVaultFilterer) ParsePauserSet(log types.Log) (*StickyVaultPauserSet, error) {
	event := new(StickyVaultPauserSet)
	if err := _StickyVault.contract.UnpackLog(event, "PauserSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the StickyVault contract.
type StickyVaultRebalanceIterator struct {
	Event *StickyVaultRebalance // Event containing the contract specifics and raw log

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
func (it *StickyVaultRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultRebalance)
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
		it.Event = new(StickyVaultRebalance)
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
func (it *StickyVaultRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultRebalance represents a Rebalance event raised by the StickyVault contract.
type StickyVaultRebalance struct {
	Compounder      common.Address
	LowerTick       *big.Int
	UpperTick       *big.Int
	LiquidityBefore *big.Int
	LiquidityAfter  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRebalance is a free log retrieval operation binding the contract event 0xff52cb3b4d8fbd8cdf8f14ffefacf8704a332a0167fd0fc3c804ead362430d85.
//
// Solidity: event Rebalance(address indexed compounder, int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_StickyVault *StickyVaultFilterer) FilterRebalance(opts *bind.FilterOpts, compounder []common.Address) (*StickyVaultRebalanceIterator, error) {

	var compounderRule []interface{}
	for _, compounderItem := range compounder {
		compounderRule = append(compounderRule, compounderItem)
	}

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "Rebalance", compounderRule)
	if err != nil {
		return nil, err
	}
	return &StickyVaultRebalanceIterator{contract: _StickyVault.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0xff52cb3b4d8fbd8cdf8f14ffefacf8704a332a0167fd0fc3c804ead362430d85.
//
// Solidity: event Rebalance(address indexed compounder, int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_StickyVault *StickyVaultFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *StickyVaultRebalance, compounder []common.Address) (event.Subscription, error) {

	var compounderRule []interface{}
	for _, compounderItem := range compounder {
		compounderRule = append(compounderRule, compounderItem)
	}

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "Rebalance", compounderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultRebalance)
				if err := _StickyVault.contract.UnpackLog(event, "Rebalance", log); err != nil {
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

// ParseRebalance is a log parse operation binding the contract event 0xff52cb3b4d8fbd8cdf8f14ffefacf8704a332a0167fd0fc3c804ead362430d85.
//
// Solidity: event Rebalance(address indexed compounder, int24 lowerTick_, int24 upperTick_, uint128 liquidityBefore, uint128 liquidityAfter)
func (_StickyVault *StickyVaultFilterer) ParseRebalance(log types.Log) (*StickyVaultRebalance, error) {
	event := new(StickyVaultRebalance)
	if err := _StickyVault.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultRestrictedMintSetIterator is returned from FilterRestrictedMintSet and is used to iterate over the raw logs and unpacked data for RestrictedMintSet events raised by the StickyVault contract.
type StickyVaultRestrictedMintSetIterator struct {
	Event *StickyVaultRestrictedMintSet // Event containing the contract specifics and raw log

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
func (it *StickyVaultRestrictedMintSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultRestrictedMintSet)
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
		it.Event = new(StickyVaultRestrictedMintSet)
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
func (it *StickyVaultRestrictedMintSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultRestrictedMintSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultRestrictedMintSet represents a RestrictedMintSet event raised by the StickyVault contract.
type StickyVaultRestrictedMintSet struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRestrictedMintSet is a free log retrieval operation binding the contract event 0x466ad1943abf95179c60a8c1ab56e7ba1467e287ca7d923b7d98965bc382c8f6.
//
// Solidity: event RestrictedMintSet(bool status)
func (_StickyVault *StickyVaultFilterer) FilterRestrictedMintSet(opts *bind.FilterOpts) (*StickyVaultRestrictedMintSetIterator, error) {

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "RestrictedMintSet")
	if err != nil {
		return nil, err
	}
	return &StickyVaultRestrictedMintSetIterator{contract: _StickyVault.contract, event: "RestrictedMintSet", logs: logs, sub: sub}, nil
}

// WatchRestrictedMintSet is a free log subscription operation binding the contract event 0x466ad1943abf95179c60a8c1ab56e7ba1467e287ca7d923b7d98965bc382c8f6.
//
// Solidity: event RestrictedMintSet(bool status)
func (_StickyVault *StickyVaultFilterer) WatchRestrictedMintSet(opts *bind.WatchOpts, sink chan<- *StickyVaultRestrictedMintSet) (event.Subscription, error) {

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "RestrictedMintSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultRestrictedMintSet)
				if err := _StickyVault.contract.UnpackLog(event, "RestrictedMintSet", log); err != nil {
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

// ParseRestrictedMintSet is a log parse operation binding the contract event 0x466ad1943abf95179c60a8c1ab56e7ba1467e287ca7d923b7d98965bc382c8f6.
//
// Solidity: event RestrictedMintSet(bool status)
func (_StickyVault *StickyVaultFilterer) ParseRestrictedMintSet(log types.Log) (*StickyVaultRestrictedMintSet, error) {
	event := new(StickyVaultRestrictedMintSet)
	if err := _StickyVault.contract.UnpackLog(event, "RestrictedMintSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultRouterSetIterator is returned from FilterRouterSet and is used to iterate over the raw logs and unpacked data for RouterSet events raised by the StickyVault contract.
type StickyVaultRouterSetIterator struct {
	Event *StickyVaultRouterSet // Event containing the contract specifics and raw log

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
func (it *StickyVaultRouterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultRouterSet)
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
		it.Event = new(StickyVaultRouterSet)
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
func (it *StickyVaultRouterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultRouterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultRouterSet represents a RouterSet event raised by the StickyVault contract.
type StickyVaultRouterSet struct {
	Router common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterSet is a free log retrieval operation binding the contract event 0x8443a265af33599abe6c9ed0c3cabd8e2a6c5b8ea35c1d1ed40c513242f95d45.
//
// Solidity: event RouterSet(address indexed router, bool status)
func (_StickyVault *StickyVaultFilterer) FilterRouterSet(opts *bind.FilterOpts, router []common.Address) (*StickyVaultRouterSetIterator, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "RouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return &StickyVaultRouterSetIterator{contract: _StickyVault.contract, event: "RouterSet", logs: logs, sub: sub}, nil
}

// WatchRouterSet is a free log subscription operation binding the contract event 0x8443a265af33599abe6c9ed0c3cabd8e2a6c5b8ea35c1d1ed40c513242f95d45.
//
// Solidity: event RouterSet(address indexed router, bool status)
func (_StickyVault *StickyVaultFilterer) WatchRouterSet(opts *bind.WatchOpts, sink chan<- *StickyVaultRouterSet, router []common.Address) (event.Subscription, error) {

	var routerRule []interface{}
	for _, routerItem := range router {
		routerRule = append(routerRule, routerItem)
	}

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "RouterSet", routerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultRouterSet)
				if err := _StickyVault.contract.UnpackLog(event, "RouterSet", log); err != nil {
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

// ParseRouterSet is a log parse operation binding the contract event 0x8443a265af33599abe6c9ed0c3cabd8e2a6c5b8ea35c1d1ed40c513242f95d45.
//
// Solidity: event RouterSet(address indexed router, bool status)
func (_StickyVault *StickyVaultFilterer) ParseRouterSet(log types.Log) (*StickyVaultRouterSet, error) {
	event := new(StickyVaultRouterSet)
	if err := _StickyVault.contract.UnpackLog(event, "RouterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StickyVault contract.
type StickyVaultTransferIterator struct {
	Event *StickyVaultTransfer // Event containing the contract specifics and raw log

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
func (it *StickyVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultTransfer)
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
		it.Event = new(StickyVaultTransfer)
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
func (it *StickyVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultTransfer represents a Transfer event raised by the StickyVault contract.
type StickyVaultTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_StickyVault *StickyVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StickyVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StickyVaultTransferIterator{contract: _StickyVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_StickyVault *StickyVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StickyVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultTransfer)
				if err := _StickyVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_StickyVault *StickyVaultFilterer) ParseTransfer(log types.Log) (*StickyVaultTransfer, error) {
	event := new(StickyVaultTransfer)
	if err := _StickyVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StickyVault contract.
type StickyVaultUnpausedIterator struct {
	Event *StickyVaultUnpaused // Event containing the contract specifics and raw log

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
func (it *StickyVaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultUnpaused)
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
		it.Event = new(StickyVaultUnpaused)
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
func (it *StickyVaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultUnpaused represents a Unpaused event raised by the StickyVault contract.
type StickyVaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StickyVault *StickyVaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StickyVaultUnpausedIterator, error) {

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StickyVaultUnpausedIterator{contract: _StickyVault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StickyVault *StickyVaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StickyVaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultUnpaused)
				if err := _StickyVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StickyVault *StickyVaultFilterer) ParseUnpaused(log types.Log) (*StickyVaultUnpaused, error) {
	event := new(StickyVaultUnpaused)
	if err := _StickyVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StickyVaultUpdateManagerParamsIterator is returned from FilterUpdateManagerParams and is used to iterate over the raw logs and unpacked data for UpdateManagerParams events raised by the StickyVault contract.
type StickyVaultUpdateManagerParamsIterator struct {
	Event *StickyVaultUpdateManagerParams // Event containing the contract specifics and raw log

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
func (it *StickyVaultUpdateManagerParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StickyVaultUpdateManagerParams)
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
		it.Event = new(StickyVaultUpdateManagerParams)
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
func (it *StickyVaultUpdateManagerParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StickyVaultUpdateManagerParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StickyVaultUpdateManagerParams represents a UpdateManagerParams event raised by the StickyVault contract.
type StickyVaultUpdateManagerParams struct {
	ManagerFeeBPS              uint16
	ManagerTreasury            common.Address
	CompounderSlippageBPS      uint16
	CompounderSlippageInterval uint32
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagerParams is a free log retrieval operation binding the contract event 0xd61f3e2d0a1abe8c7a06f409e2069f4ba0ad684d8dcb174b9101f0bccf5e283f.
//
// Solidity: event UpdateManagerParams(uint16 managerFeeBPS, address managerTreasury, uint16 compounderSlippageBPS, uint32 compounderSlippageInterval)
func (_StickyVault *StickyVaultFilterer) FilterUpdateManagerParams(opts *bind.FilterOpts) (*StickyVaultUpdateManagerParamsIterator, error) {

	logs, sub, err := _StickyVault.contract.FilterLogs(opts, "UpdateManagerParams")
	if err != nil {
		return nil, err
	}
	return &StickyVaultUpdateManagerParamsIterator{contract: _StickyVault.contract, event: "UpdateManagerParams", logs: logs, sub: sub}, nil
}

// WatchUpdateManagerParams is a free log subscription operation binding the contract event 0xd61f3e2d0a1abe8c7a06f409e2069f4ba0ad684d8dcb174b9101f0bccf5e283f.
//
// Solidity: event UpdateManagerParams(uint16 managerFeeBPS, address managerTreasury, uint16 compounderSlippageBPS, uint32 compounderSlippageInterval)
func (_StickyVault *StickyVaultFilterer) WatchUpdateManagerParams(opts *bind.WatchOpts, sink chan<- *StickyVaultUpdateManagerParams) (event.Subscription, error) {

	logs, sub, err := _StickyVault.contract.WatchLogs(opts, "UpdateManagerParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StickyVaultUpdateManagerParams)
				if err := _StickyVault.contract.UnpackLog(event, "UpdateManagerParams", log); err != nil {
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

// ParseUpdateManagerParams is a log parse operation binding the contract event 0xd61f3e2d0a1abe8c7a06f409e2069f4ba0ad684d8dcb174b9101f0bccf5e283f.
//
// Solidity: event UpdateManagerParams(uint16 managerFeeBPS, address managerTreasury, uint16 compounderSlippageBPS, uint32 compounderSlippageInterval)
func (_StickyVault *StickyVaultFilterer) ParseUpdateManagerParams(log types.Log) (*StickyVaultUpdateManagerParams, error) {
	event := new(StickyVaultUpdateManagerParams)
	if err := _StickyVault.contract.UnpackLog(event, "UpdateManagerParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
