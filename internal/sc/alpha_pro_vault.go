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

// VaultParams is an auto generated low-level Go binding around an user-defined struct.
type VaultParams struct {
	Pool             common.Address
	Manager          common.Address
	ManagerFee       *big.Int
	MaxTotalSupply   *big.Int
	WideRangeWeight  *big.Int
	WideThreshold    *big.Int
	BaseThreshold    *big.Int
	LimitThreshold   *big.Int
	Period           uint32
	MinTickMove      *big.Int
	MaxTwapDeviation *big.Int
	TwapDuration     uint32
	Name             string
	Symbol           string
}

// AlphaProVaultMetaData contains all meta data concerning the AlphaProVault contract.
var AlphaProVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"APV_Amount0Min\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_Amount1Min\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_ManagerFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_MaxTotalSupply\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_MaxTwapDeviation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_MinTickMove\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_NotDepositDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_NotGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_NotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_NotManagerOrRebalanceDelegate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_NotPendingManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_NotPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_PeriodNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_PriceOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_ProtocolFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_SweepToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_ThresholdNotMultipleOfTickSpacing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_ThresholdNotPositive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_ThresholdsCannotBeSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_TickNotMoved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_TwapDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_TwapPriceDeviation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_WideRangeWeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_ZeroCross\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_ZeroDepositAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"APV_ZeroShares\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"T\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToVault0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToVault1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToProtocol0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToProtocol1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToManager0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feesToManager1\",\"type\":\"uint256\"}],\"name\":\"CollectFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"CollectManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"CollectProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Snapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"threshold\",\"type\":\"int24\"}],\"name\":\"UpdateBaseThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"UpdateDepositDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"threshold\",\"type\":\"int24\"}],\"name\":\"UpdateLimitThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"UpdateManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"managerFee\",\"type\":\"uint24\"}],\"name\":\"UpdateManagerFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTotalSupply\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxTotalSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"maxTwapDeviation\",\"type\":\"int24\"}],\"name\":\"UpdateMaxTwapDeviation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"minTickMove\",\"type\":\"int24\"}],\"name\":\"UpdateMinTickMove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"UpdatePendingManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"UpdatePeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"protocolFee\",\"type\":\"uint24\"}],\"name\":\"UpdateProtocolFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"UpdateRebalanceDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"twapDuration\",\"type\":\"uint32\"}],\"name\":\"UpdateTwapDuration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"weight\",\"type\":\"uint24\"}],\"name\":\"UpdateWideRangeWeight\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"threshold\",\"type\":\"int24\"}],\"name\":\"UpdateWideThreshold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"HUNDRED_PERCENT\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedManagerFees0\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedManagerFees1\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedProtocolFees0\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedProtocolFees1\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseThreshold\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkCanRebalance\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkPriceNearTwap\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"collectManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"collectProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Desired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"}],\"name\":\"emergencyBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractAlphaProVaultFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"bool\",\"name\":\"roundUp\",\"type\":\"bool\"}],\"name\":\"getPositionAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPositions\",\"outputs\":[{\"internalType\":\"int24[2][3]\",\"name\":\"\",\"type\":\"int24[2][3]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"roundUp\",\"type\":\"bool\"}],\"name\":\"getTotalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTwap\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"managerFee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint24\",\"name\":\"wideRangeWeight\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"wideThreshold\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"baseThreshold\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"limitThreshold\",\"type\":\"int24\"},{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"},{\"internalType\":\"int24\",\"name\":\"minTickMove\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"maxTwapDeviation\",\"type\":\"int24\"},{\"internalType\":\"uint32\",\"name\":\"twapDuration\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"internalType\":\"structVaultParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitThreshold\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"managerFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTwapDeviation\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTickMove\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingManagerFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingProtocolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"period\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIKodiakPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalanceDelegate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_baseThreshold\",\"type\":\"int24\"}],\"name\":\"setBaseThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositDelegate\",\"type\":\"address\"}],\"name\":\"setDepositDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_limitThreshold\",\"type\":\"int24\"}],\"name\":\"setLimitThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_pendingManagerFee\",\"type\":\"uint24\"}],\"name\":\"setManagerFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_maxTwapDeviation\",\"type\":\"int24\"}],\"name\":\"setMaxTwapDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_minTickMove\",\"type\":\"int24\"}],\"name\":\"setMinTickMove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_period\",\"type\":\"uint32\"}],\"name\":\"setPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_pendingProtocolFee\",\"type\":\"uint24\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rebalanceDelegate\",\"type\":\"address\"}],\"name\":\"setRebalanceDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_twapDuration\",\"type\":\"uint32\"}],\"name\":\"setTwapDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_wideRangeWeight\",\"type\":\"uint24\"}],\"name\":\"setWideRangeWeight\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_wideThreshold\",\"type\":\"int24\"}],\"name\":\"setWideThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"twapDuration\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3MintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wideLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wideRangeWeight\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wideThreshold\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wideUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Min\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AlphaProVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use AlphaProVaultMetaData.ABI instead.
var AlphaProVaultABI = AlphaProVaultMetaData.ABI

// AlphaProVault is an auto generated Go binding around an Ethereum contract.
type AlphaProVault struct {
	AlphaProVaultCaller     // Read-only binding to the contract
	AlphaProVaultTransactor // Write-only binding to the contract
	AlphaProVaultFilterer   // Log filterer for contract events
}

// AlphaProVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlphaProVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphaProVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlphaProVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphaProVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlphaProVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlphaProVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlphaProVaultSession struct {
	Contract     *AlphaProVault    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlphaProVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlphaProVaultCallerSession struct {
	Contract *AlphaProVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AlphaProVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlphaProVaultTransactorSession struct {
	Contract     *AlphaProVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AlphaProVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlphaProVaultRaw struct {
	Contract *AlphaProVault // Generic contract binding to access the raw methods on
}

// AlphaProVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlphaProVaultCallerRaw struct {
	Contract *AlphaProVaultCaller // Generic read-only contract binding to access the raw methods on
}

// AlphaProVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlphaProVaultTransactorRaw struct {
	Contract *AlphaProVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlphaProVault creates a new instance of AlphaProVault, bound to a specific deployed contract.
func NewAlphaProVault(address common.Address, backend bind.ContractBackend) (*AlphaProVault, error) {
	contract, err := bindAlphaProVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AlphaProVault{AlphaProVaultCaller: AlphaProVaultCaller{contract: contract}, AlphaProVaultTransactor: AlphaProVaultTransactor{contract: contract}, AlphaProVaultFilterer: AlphaProVaultFilterer{contract: contract}}, nil
}

// NewAlphaProVaultCaller creates a new read-only instance of AlphaProVault, bound to a specific deployed contract.
func NewAlphaProVaultCaller(address common.Address, caller bind.ContractCaller) (*AlphaProVaultCaller, error) {
	contract, err := bindAlphaProVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultCaller{contract: contract}, nil
}

// NewAlphaProVaultTransactor creates a new write-only instance of AlphaProVault, bound to a specific deployed contract.
func NewAlphaProVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*AlphaProVaultTransactor, error) {
	contract, err := bindAlphaProVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultTransactor{contract: contract}, nil
}

// NewAlphaProVaultFilterer creates a new log filterer instance of AlphaProVault, bound to a specific deployed contract.
func NewAlphaProVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*AlphaProVaultFilterer, error) {
	contract, err := bindAlphaProVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultFilterer{contract: contract}, nil
}

// bindAlphaProVault binds a generic wrapper to an already deployed contract.
func bindAlphaProVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AlphaProVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlphaProVault *AlphaProVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlphaProVault.Contract.AlphaProVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlphaProVault *AlphaProVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlphaProVault.Contract.AlphaProVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlphaProVault *AlphaProVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlphaProVault.Contract.AlphaProVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlphaProVault *AlphaProVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlphaProVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlphaProVault *AlphaProVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlphaProVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlphaProVault *AlphaProVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlphaProVault.Contract.contract.Transact(opts, method, params...)
}

// HUNDREDPERCENT is a free data retrieval call binding the contract method 0x6ed93dd0.
//
// Solidity: function HUNDRED_PERCENT() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCaller) HUNDREDPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "HUNDRED_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HUNDREDPERCENT is a free data retrieval call binding the contract method 0x6ed93dd0.
//
// Solidity: function HUNDRED_PERCENT() view returns(uint24)
func (_AlphaProVault *AlphaProVaultSession) HUNDREDPERCENT() (*big.Int, error) {
	return _AlphaProVault.Contract.HUNDREDPERCENT(&_AlphaProVault.CallOpts)
}

// HUNDREDPERCENT is a free data retrieval call binding the contract method 0x6ed93dd0.
//
// Solidity: function HUNDRED_PERCENT() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCallerSession) HUNDREDPERCENT() (*big.Int, error) {
	return _AlphaProVault.Contract.HUNDREDPERCENT(&_AlphaProVault.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint24)
func (_AlphaProVault *AlphaProVaultSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _AlphaProVault.Contract.MINIMUMLIQUIDITY(&_AlphaProVault.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _AlphaProVault.Contract.MINIMUMLIQUIDITY(&_AlphaProVault.CallOpts)
}

// AccruedManagerFees0 is a free data retrieval call binding the contract method 0xecb5468e.
//
// Solidity: function accruedManagerFees0() view returns(uint128)
func (_AlphaProVault *AlphaProVaultCaller) AccruedManagerFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "accruedManagerFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedManagerFees0 is a free data retrieval call binding the contract method 0xecb5468e.
//
// Solidity: function accruedManagerFees0() view returns(uint128)
func (_AlphaProVault *AlphaProVaultSession) AccruedManagerFees0() (*big.Int, error) {
	return _AlphaProVault.Contract.AccruedManagerFees0(&_AlphaProVault.CallOpts)
}

// AccruedManagerFees0 is a free data retrieval call binding the contract method 0xecb5468e.
//
// Solidity: function accruedManagerFees0() view returns(uint128)
func (_AlphaProVault *AlphaProVaultCallerSession) AccruedManagerFees0() (*big.Int, error) {
	return _AlphaProVault.Contract.AccruedManagerFees0(&_AlphaProVault.CallOpts)
}

// AccruedManagerFees1 is a free data retrieval call binding the contract method 0x60e1592a.
//
// Solidity: function accruedManagerFees1() view returns(uint128)
func (_AlphaProVault *AlphaProVaultCaller) AccruedManagerFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "accruedManagerFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedManagerFees1 is a free data retrieval call binding the contract method 0x60e1592a.
//
// Solidity: function accruedManagerFees1() view returns(uint128)
func (_AlphaProVault *AlphaProVaultSession) AccruedManagerFees1() (*big.Int, error) {
	return _AlphaProVault.Contract.AccruedManagerFees1(&_AlphaProVault.CallOpts)
}

// AccruedManagerFees1 is a free data retrieval call binding the contract method 0x60e1592a.
//
// Solidity: function accruedManagerFees1() view returns(uint128)
func (_AlphaProVault *AlphaProVaultCallerSession) AccruedManagerFees1() (*big.Int, error) {
	return _AlphaProVault.Contract.AccruedManagerFees1(&_AlphaProVault.CallOpts)
}

// AccruedProtocolFees0 is a free data retrieval call binding the contract method 0xeae989a2.
//
// Solidity: function accruedProtocolFees0() view returns(uint128)
func (_AlphaProVault *AlphaProVaultCaller) AccruedProtocolFees0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "accruedProtocolFees0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedProtocolFees0 is a free data retrieval call binding the contract method 0xeae989a2.
//
// Solidity: function accruedProtocolFees0() view returns(uint128)
func (_AlphaProVault *AlphaProVaultSession) AccruedProtocolFees0() (*big.Int, error) {
	return _AlphaProVault.Contract.AccruedProtocolFees0(&_AlphaProVault.CallOpts)
}

// AccruedProtocolFees0 is a free data retrieval call binding the contract method 0xeae989a2.
//
// Solidity: function accruedProtocolFees0() view returns(uint128)
func (_AlphaProVault *AlphaProVaultCallerSession) AccruedProtocolFees0() (*big.Int, error) {
	return _AlphaProVault.Contract.AccruedProtocolFees0(&_AlphaProVault.CallOpts)
}

// AccruedProtocolFees1 is a free data retrieval call binding the contract method 0xa00fa77f.
//
// Solidity: function accruedProtocolFees1() view returns(uint128)
func (_AlphaProVault *AlphaProVaultCaller) AccruedProtocolFees1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "accruedProtocolFees1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedProtocolFees1 is a free data retrieval call binding the contract method 0xa00fa77f.
//
// Solidity: function accruedProtocolFees1() view returns(uint128)
func (_AlphaProVault *AlphaProVaultSession) AccruedProtocolFees1() (*big.Int, error) {
	return _AlphaProVault.Contract.AccruedProtocolFees1(&_AlphaProVault.CallOpts)
}

// AccruedProtocolFees1 is a free data retrieval call binding the contract method 0xa00fa77f.
//
// Solidity: function accruedProtocolFees1() view returns(uint128)
func (_AlphaProVault *AlphaProVaultCallerSession) AccruedProtocolFees1() (*big.Int, error) {
	return _AlphaProVault.Contract.AccruedProtocolFees1(&_AlphaProVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AlphaProVault *AlphaProVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AlphaProVault *AlphaProVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AlphaProVault.Contract.Allowance(&_AlphaProVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_AlphaProVault *AlphaProVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _AlphaProVault.Contract.Allowance(&_AlphaProVault.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AlphaProVault *AlphaProVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AlphaProVault *AlphaProVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AlphaProVault.Contract.BalanceOf(&_AlphaProVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_AlphaProVault *AlphaProVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _AlphaProVault.Contract.BalanceOf(&_AlphaProVault.CallOpts, account)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) BaseLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "baseLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) BaseLower() (*big.Int, error) {
	return _AlphaProVault.Contract.BaseLower(&_AlphaProVault.CallOpts)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) BaseLower() (*big.Int, error) {
	return _AlphaProVault.Contract.BaseLower(&_AlphaProVault.CallOpts)
}

// BaseThreshold is a free data retrieval call binding the contract method 0xa87bab9c.
//
// Solidity: function baseThreshold() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) BaseThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "baseThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseThreshold is a free data retrieval call binding the contract method 0xa87bab9c.
//
// Solidity: function baseThreshold() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) BaseThreshold() (*big.Int, error) {
	return _AlphaProVault.Contract.BaseThreshold(&_AlphaProVault.CallOpts)
}

// BaseThreshold is a free data retrieval call binding the contract method 0xa87bab9c.
//
// Solidity: function baseThreshold() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) BaseThreshold() (*big.Int, error) {
	return _AlphaProVault.Contract.BaseThreshold(&_AlphaProVault.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) BaseUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "baseUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) BaseUpper() (*big.Int, error) {
	return _AlphaProVault.Contract.BaseUpper(&_AlphaProVault.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) BaseUpper() (*big.Int, error) {
	return _AlphaProVault.Contract.BaseUpper(&_AlphaProVault.CallOpts)
}

// CheckCanRebalance is a free data retrieval call binding the contract method 0x71bd0ea7.
//
// Solidity: function checkCanRebalance() view returns()
func (_AlphaProVault *AlphaProVaultCaller) CheckCanRebalance(opts *bind.CallOpts) error {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "checkCanRebalance")

	if err != nil {
		return err
	}

	return err

}

// CheckCanRebalance is a free data retrieval call binding the contract method 0x71bd0ea7.
//
// Solidity: function checkCanRebalance() view returns()
func (_AlphaProVault *AlphaProVaultSession) CheckCanRebalance() error {
	return _AlphaProVault.Contract.CheckCanRebalance(&_AlphaProVault.CallOpts)
}

// CheckCanRebalance is a free data retrieval call binding the contract method 0x71bd0ea7.
//
// Solidity: function checkCanRebalance() view returns()
func (_AlphaProVault *AlphaProVaultCallerSession) CheckCanRebalance() error {
	return _AlphaProVault.Contract.CheckCanRebalance(&_AlphaProVault.CallOpts)
}

// CheckPriceNearTwap is a free data retrieval call binding the contract method 0x345d90e2.
//
// Solidity: function checkPriceNearTwap() view returns()
func (_AlphaProVault *AlphaProVaultCaller) CheckPriceNearTwap(opts *bind.CallOpts) error {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "checkPriceNearTwap")

	if err != nil {
		return err
	}

	return err

}

// CheckPriceNearTwap is a free data retrieval call binding the contract method 0x345d90e2.
//
// Solidity: function checkPriceNearTwap() view returns()
func (_AlphaProVault *AlphaProVaultSession) CheckPriceNearTwap() error {
	return _AlphaProVault.Contract.CheckPriceNearTwap(&_AlphaProVault.CallOpts)
}

// CheckPriceNearTwap is a free data retrieval call binding the contract method 0x345d90e2.
//
// Solidity: function checkPriceNearTwap() view returns()
func (_AlphaProVault *AlphaProVaultCallerSession) CheckPriceNearTwap() error {
	return _AlphaProVault.Contract.CheckPriceNearTwap(&_AlphaProVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AlphaProVault *AlphaProVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AlphaProVault *AlphaProVaultSession) Decimals() (uint8, error) {
	return _AlphaProVault.Contract.Decimals(&_AlphaProVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AlphaProVault *AlphaProVaultCallerSession) Decimals() (uint8, error) {
	return _AlphaProVault.Contract.Decimals(&_AlphaProVault.CallOpts)
}

// DepositDelegate is a free data retrieval call binding the contract method 0x4efed314.
//
// Solidity: function depositDelegate() view returns(address)
func (_AlphaProVault *AlphaProVaultCaller) DepositDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "depositDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositDelegate is a free data retrieval call binding the contract method 0x4efed314.
//
// Solidity: function depositDelegate() view returns(address)
func (_AlphaProVault *AlphaProVaultSession) DepositDelegate() (common.Address, error) {
	return _AlphaProVault.Contract.DepositDelegate(&_AlphaProVault.CallOpts)
}

// DepositDelegate is a free data retrieval call binding the contract method 0x4efed314.
//
// Solidity: function depositDelegate() view returns(address)
func (_AlphaProVault *AlphaProVaultCallerSession) DepositDelegate() (common.Address, error) {
	return _AlphaProVault.Contract.DepositDelegate(&_AlphaProVault.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AlphaProVault *AlphaProVaultCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AlphaProVault *AlphaProVaultSession) Factory() (common.Address, error) {
	return _AlphaProVault.Contract.Factory(&_AlphaProVault.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_AlphaProVault *AlphaProVaultCallerSession) Factory() (common.Address, error) {
	return _AlphaProVault.Contract.Factory(&_AlphaProVault.CallOpts)
}

// GetBalance0 is a free data retrieval call binding the contract method 0x629d9405.
//
// Solidity: function getBalance0() view returns(uint256)
func (_AlphaProVault *AlphaProVaultCaller) GetBalance0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "getBalance0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance0 is a free data retrieval call binding the contract method 0x629d9405.
//
// Solidity: function getBalance0() view returns(uint256)
func (_AlphaProVault *AlphaProVaultSession) GetBalance0() (*big.Int, error) {
	return _AlphaProVault.Contract.GetBalance0(&_AlphaProVault.CallOpts)
}

// GetBalance0 is a free data retrieval call binding the contract method 0x629d9405.
//
// Solidity: function getBalance0() view returns(uint256)
func (_AlphaProVault *AlphaProVaultCallerSession) GetBalance0() (*big.Int, error) {
	return _AlphaProVault.Contract.GetBalance0(&_AlphaProVault.CallOpts)
}

// GetBalance1 is a free data retrieval call binding the contract method 0x41aec538.
//
// Solidity: function getBalance1() view returns(uint256)
func (_AlphaProVault *AlphaProVaultCaller) GetBalance1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "getBalance1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance1 is a free data retrieval call binding the contract method 0x41aec538.
//
// Solidity: function getBalance1() view returns(uint256)
func (_AlphaProVault *AlphaProVaultSession) GetBalance1() (*big.Int, error) {
	return _AlphaProVault.Contract.GetBalance1(&_AlphaProVault.CallOpts)
}

// GetBalance1 is a free data retrieval call binding the contract method 0x41aec538.
//
// Solidity: function getBalance1() view returns(uint256)
func (_AlphaProVault *AlphaProVaultCallerSession) GetBalance1() (*big.Int, error) {
	return _AlphaProVault.Contract.GetBalance1(&_AlphaProVault.CallOpts)
}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xc11c64e5.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper, bool roundUp) view returns(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultCaller) GetPositionAmounts(opts *bind.CallOpts, tickLower *big.Int, tickUpper *big.Int, roundUp bool) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "getPositionAmounts", tickLower, tickUpper, roundUp)

	outstruct := new(struct {
		Amount0 *big.Int
		Amount1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xc11c64e5.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper, bool roundUp) view returns(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultSession) GetPositionAmounts(tickLower *big.Int, tickUpper *big.Int, roundUp bool) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _AlphaProVault.Contract.GetPositionAmounts(&_AlphaProVault.CallOpts, tickLower, tickUpper, roundUp)
}

// GetPositionAmounts is a free data retrieval call binding the contract method 0xc11c64e5.
//
// Solidity: function getPositionAmounts(int24 tickLower, int24 tickUpper, bool roundUp) view returns(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultCallerSession) GetPositionAmounts(tickLower *big.Int, tickUpper *big.Int, roundUp bool) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _AlphaProVault.Contract.GetPositionAmounts(&_AlphaProVault.CallOpts, tickLower, tickUpper, roundUp)
}

// GetPositions is a free data retrieval call binding the contract method 0x80275860.
//
// Solidity: function getPositions() view returns(int24[2][3])
func (_AlphaProVault *AlphaProVaultCaller) GetPositions(opts *bind.CallOpts) ([3][2]*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "getPositions")

	if err != nil {
		return *new([3][2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([3][2]*big.Int)).(*[3][2]*big.Int)

	return out0, err

}

// GetPositions is a free data retrieval call binding the contract method 0x80275860.
//
// Solidity: function getPositions() view returns(int24[2][3])
func (_AlphaProVault *AlphaProVaultSession) GetPositions() ([3][2]*big.Int, error) {
	return _AlphaProVault.Contract.GetPositions(&_AlphaProVault.CallOpts)
}

// GetPositions is a free data retrieval call binding the contract method 0x80275860.
//
// Solidity: function getPositions() view returns(int24[2][3])
func (_AlphaProVault *AlphaProVaultCallerSession) GetPositions() ([3][2]*big.Int, error) {
	return _AlphaProVault.Contract.GetPositions(&_AlphaProVault.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0x47c570fd.
//
// Solidity: function getTotalAmounts(bool roundUp) view returns(uint256 total0, uint256 total1)
func (_AlphaProVault *AlphaProVaultCaller) GetTotalAmounts(opts *bind.CallOpts, roundUp bool) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "getTotalAmounts", roundUp)

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

// GetTotalAmounts is a free data retrieval call binding the contract method 0x47c570fd.
//
// Solidity: function getTotalAmounts(bool roundUp) view returns(uint256 total0, uint256 total1)
func (_AlphaProVault *AlphaProVaultSession) GetTotalAmounts(roundUp bool) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _AlphaProVault.Contract.GetTotalAmounts(&_AlphaProVault.CallOpts, roundUp)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0x47c570fd.
//
// Solidity: function getTotalAmounts(bool roundUp) view returns(uint256 total0, uint256 total1)
func (_AlphaProVault *AlphaProVaultCallerSession) GetTotalAmounts(roundUp bool) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _AlphaProVault.Contract.GetTotalAmounts(&_AlphaProVault.CallOpts, roundUp)
}

// GetTotalAmounts0 is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_AlphaProVault *AlphaProVaultCaller) GetTotalAmounts0(opts *bind.CallOpts) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "getTotalAmounts0")

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

// GetTotalAmounts0 is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_AlphaProVault *AlphaProVaultSession) GetTotalAmounts0() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _AlphaProVault.Contract.GetTotalAmounts0(&_AlphaProVault.CallOpts)
}

// GetTotalAmounts0 is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_AlphaProVault *AlphaProVaultCallerSession) GetTotalAmounts0() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _AlphaProVault.Contract.GetTotalAmounts0(&_AlphaProVault.CallOpts)
}

// GetTwap is a free data retrieval call binding the contract method 0x5d752a9a.
//
// Solidity: function getTwap() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) GetTwap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "getTwap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTwap is a free data retrieval call binding the contract method 0x5d752a9a.
//
// Solidity: function getTwap() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) GetTwap() (*big.Int, error) {
	return _AlphaProVault.Contract.GetTwap(&_AlphaProVault.CallOpts)
}

// GetTwap is a free data retrieval call binding the contract method 0x5d752a9a.
//
// Solidity: function getTwap() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) GetTwap() (*big.Int, error) {
	return _AlphaProVault.Contract.GetTwap(&_AlphaProVault.CallOpts)
}

// LastTick is a free data retrieval call binding the contract method 0x3dfa5d87.
//
// Solidity: function lastTick() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) LastTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "lastTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTick is a free data retrieval call binding the contract method 0x3dfa5d87.
//
// Solidity: function lastTick() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) LastTick() (*big.Int, error) {
	return _AlphaProVault.Contract.LastTick(&_AlphaProVault.CallOpts)
}

// LastTick is a free data retrieval call binding the contract method 0x3dfa5d87.
//
// Solidity: function lastTick() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) LastTick() (*big.Int, error) {
	return _AlphaProVault.Contract.LastTick(&_AlphaProVault.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint40)
func (_AlphaProVault *AlphaProVaultCaller) LastTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "lastTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint40)
func (_AlphaProVault *AlphaProVaultSession) LastTimestamp() (*big.Int, error) {
	return _AlphaProVault.Contract.LastTimestamp(&_AlphaProVault.CallOpts)
}

// LastTimestamp is a free data retrieval call binding the contract method 0x19d8ac61.
//
// Solidity: function lastTimestamp() view returns(uint40)
func (_AlphaProVault *AlphaProVaultCallerSession) LastTimestamp() (*big.Int, error) {
	return _AlphaProVault.Contract.LastTimestamp(&_AlphaProVault.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) LimitLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "limitLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) LimitLower() (*big.Int, error) {
	return _AlphaProVault.Contract.LimitLower(&_AlphaProVault.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) LimitLower() (*big.Int, error) {
	return _AlphaProVault.Contract.LimitLower(&_AlphaProVault.CallOpts)
}

// LimitThreshold is a free data retrieval call binding the contract method 0x16c3e29d.
//
// Solidity: function limitThreshold() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) LimitThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "limitThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitThreshold is a free data retrieval call binding the contract method 0x16c3e29d.
//
// Solidity: function limitThreshold() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) LimitThreshold() (*big.Int, error) {
	return _AlphaProVault.Contract.LimitThreshold(&_AlphaProVault.CallOpts)
}

// LimitThreshold is a free data retrieval call binding the contract method 0x16c3e29d.
//
// Solidity: function limitThreshold() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) LimitThreshold() (*big.Int, error) {
	return _AlphaProVault.Contract.LimitThreshold(&_AlphaProVault.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) LimitUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "limitUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) LimitUpper() (*big.Int, error) {
	return _AlphaProVault.Contract.LimitUpper(&_AlphaProVault.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) LimitUpper() (*big.Int, error) {
	return _AlphaProVault.Contract.LimitUpper(&_AlphaProVault.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_AlphaProVault *AlphaProVaultCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_AlphaProVault *AlphaProVaultSession) Manager() (common.Address, error) {
	return _AlphaProVault.Contract.Manager(&_AlphaProVault.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_AlphaProVault *AlphaProVaultCallerSession) Manager() (common.Address, error) {
	return _AlphaProVault.Contract.Manager(&_AlphaProVault.CallOpts)
}

// ManagerFee is a free data retrieval call binding the contract method 0x9c7632fc.
//
// Solidity: function managerFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCaller) ManagerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "managerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ManagerFee is a free data retrieval call binding the contract method 0x9c7632fc.
//
// Solidity: function managerFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultSession) ManagerFee() (*big.Int, error) {
	return _AlphaProVault.Contract.ManagerFee(&_AlphaProVault.CallOpts)
}

// ManagerFee is a free data retrieval call binding the contract method 0x9c7632fc.
//
// Solidity: function managerFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCallerSession) ManagerFee() (*big.Int, error) {
	return _AlphaProVault.Contract.ManagerFee(&_AlphaProVault.CallOpts)
}

// MaxTick is a free data retrieval call binding the contract method 0x1db3b6dd.
//
// Solidity: function maxTick() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) MaxTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "maxTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTick is a free data retrieval call binding the contract method 0x1db3b6dd.
//
// Solidity: function maxTick() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) MaxTick() (*big.Int, error) {
	return _AlphaProVault.Contract.MaxTick(&_AlphaProVault.CallOpts)
}

// MaxTick is a free data retrieval call binding the contract method 0x1db3b6dd.
//
// Solidity: function maxTick() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) MaxTick() (*big.Int, error) {
	return _AlphaProVault.Contract.MaxTick(&_AlphaProVault.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_AlphaProVault *AlphaProVaultCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_AlphaProVault *AlphaProVaultSession) MaxTotalSupply() (*big.Int, error) {
	return _AlphaProVault.Contract.MaxTotalSupply(&_AlphaProVault.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_AlphaProVault *AlphaProVaultCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _AlphaProVault.Contract.MaxTotalSupply(&_AlphaProVault.CallOpts)
}

// MaxTwapDeviation is a free data retrieval call binding the contract method 0xe7c7cb91.
//
// Solidity: function maxTwapDeviation() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) MaxTwapDeviation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "maxTwapDeviation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTwapDeviation is a free data retrieval call binding the contract method 0xe7c7cb91.
//
// Solidity: function maxTwapDeviation() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) MaxTwapDeviation() (*big.Int, error) {
	return _AlphaProVault.Contract.MaxTwapDeviation(&_AlphaProVault.CallOpts)
}

// MaxTwapDeviation is a free data retrieval call binding the contract method 0xe7c7cb91.
//
// Solidity: function maxTwapDeviation() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) MaxTwapDeviation() (*big.Int, error) {
	return _AlphaProVault.Contract.MaxTwapDeviation(&_AlphaProVault.CallOpts)
}

// MinTickMove is a free data retrieval call binding the contract method 0x854360ef.
//
// Solidity: function minTickMove() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) MinTickMove(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "minTickMove")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTickMove is a free data retrieval call binding the contract method 0x854360ef.
//
// Solidity: function minTickMove() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) MinTickMove() (*big.Int, error) {
	return _AlphaProVault.Contract.MinTickMove(&_AlphaProVault.CallOpts)
}

// MinTickMove is a free data retrieval call binding the contract method 0x854360ef.
//
// Solidity: function minTickMove() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) MinTickMove() (*big.Int, error) {
	return _AlphaProVault.Contract.MinTickMove(&_AlphaProVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AlphaProVault *AlphaProVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AlphaProVault *AlphaProVaultSession) Name() (string, error) {
	return _AlphaProVault.Contract.Name(&_AlphaProVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AlphaProVault *AlphaProVaultCallerSession) Name() (string, error) {
	return _AlphaProVault.Contract.Name(&_AlphaProVault.CallOpts)
}

// PendingManager is a free data retrieval call binding the contract method 0xa00fff6f.
//
// Solidity: function pendingManager() view returns(address)
func (_AlphaProVault *AlphaProVaultCaller) PendingManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "pendingManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingManager is a free data retrieval call binding the contract method 0xa00fff6f.
//
// Solidity: function pendingManager() view returns(address)
func (_AlphaProVault *AlphaProVaultSession) PendingManager() (common.Address, error) {
	return _AlphaProVault.Contract.PendingManager(&_AlphaProVault.CallOpts)
}

// PendingManager is a free data retrieval call binding the contract method 0xa00fff6f.
//
// Solidity: function pendingManager() view returns(address)
func (_AlphaProVault *AlphaProVaultCallerSession) PendingManager() (common.Address, error) {
	return _AlphaProVault.Contract.PendingManager(&_AlphaProVault.CallOpts)
}

// PendingManagerFee is a free data retrieval call binding the contract method 0xe4d2b8aa.
//
// Solidity: function pendingManagerFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCaller) PendingManagerFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "pendingManagerFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingManagerFee is a free data retrieval call binding the contract method 0xe4d2b8aa.
//
// Solidity: function pendingManagerFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultSession) PendingManagerFee() (*big.Int, error) {
	return _AlphaProVault.Contract.PendingManagerFee(&_AlphaProVault.CallOpts)
}

// PendingManagerFee is a free data retrieval call binding the contract method 0xe4d2b8aa.
//
// Solidity: function pendingManagerFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCallerSession) PendingManagerFee() (*big.Int, error) {
	return _AlphaProVault.Contract.PendingManagerFee(&_AlphaProVault.CallOpts)
}

// PendingProtocolFee is a free data retrieval call binding the contract method 0xc3fa0e3b.
//
// Solidity: function pendingProtocolFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCaller) PendingProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "pendingProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingProtocolFee is a free data retrieval call binding the contract method 0xc3fa0e3b.
//
// Solidity: function pendingProtocolFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultSession) PendingProtocolFee() (*big.Int, error) {
	return _AlphaProVault.Contract.PendingProtocolFee(&_AlphaProVault.CallOpts)
}

// PendingProtocolFee is a free data retrieval call binding the contract method 0xc3fa0e3b.
//
// Solidity: function pendingProtocolFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCallerSession) PendingProtocolFee() (*big.Int, error) {
	return _AlphaProVault.Contract.PendingProtocolFee(&_AlphaProVault.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint32)
func (_AlphaProVault *AlphaProVaultCaller) Period(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "period")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint32)
func (_AlphaProVault *AlphaProVaultSession) Period() (uint32, error) {
	return _AlphaProVault.Contract.Period(&_AlphaProVault.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint32)
func (_AlphaProVault *AlphaProVaultCallerSession) Period() (uint32, error) {
	return _AlphaProVault.Contract.Period(&_AlphaProVault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_AlphaProVault *AlphaProVaultCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_AlphaProVault *AlphaProVaultSession) Pool() (common.Address, error) {
	return _AlphaProVault.Contract.Pool(&_AlphaProVault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_AlphaProVault *AlphaProVaultCallerSession) Pool() (common.Address, error) {
	return _AlphaProVault.Contract.Pool(&_AlphaProVault.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultSession) ProtocolFee() (*big.Int, error) {
	return _AlphaProVault.Contract.ProtocolFee(&_AlphaProVault.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCallerSession) ProtocolFee() (*big.Int, error) {
	return _AlphaProVault.Contract.ProtocolFee(&_AlphaProVault.CallOpts)
}

// RebalanceDelegate is a free data retrieval call binding the contract method 0x5da07868.
//
// Solidity: function rebalanceDelegate() view returns(address)
func (_AlphaProVault *AlphaProVaultCaller) RebalanceDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "rebalanceDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RebalanceDelegate is a free data retrieval call binding the contract method 0x5da07868.
//
// Solidity: function rebalanceDelegate() view returns(address)
func (_AlphaProVault *AlphaProVaultSession) RebalanceDelegate() (common.Address, error) {
	return _AlphaProVault.Contract.RebalanceDelegate(&_AlphaProVault.CallOpts)
}

// RebalanceDelegate is a free data retrieval call binding the contract method 0x5da07868.
//
// Solidity: function rebalanceDelegate() view returns(address)
func (_AlphaProVault *AlphaProVaultCallerSession) RebalanceDelegate() (common.Address, error) {
	return _AlphaProVault.Contract.RebalanceDelegate(&_AlphaProVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AlphaProVault *AlphaProVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AlphaProVault *AlphaProVaultSession) Symbol() (string, error) {
	return _AlphaProVault.Contract.Symbol(&_AlphaProVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AlphaProVault *AlphaProVaultCallerSession) Symbol() (string, error) {
	return _AlphaProVault.Contract.Symbol(&_AlphaProVault.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) TickSpacing() (*big.Int, error) {
	return _AlphaProVault.Contract.TickSpacing(&_AlphaProVault.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) TickSpacing() (*big.Int, error) {
	return _AlphaProVault.Contract.TickSpacing(&_AlphaProVault.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AlphaProVault *AlphaProVaultCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AlphaProVault *AlphaProVaultSession) Token0() (common.Address, error) {
	return _AlphaProVault.Contract.Token0(&_AlphaProVault.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_AlphaProVault *AlphaProVaultCallerSession) Token0() (common.Address, error) {
	return _AlphaProVault.Contract.Token0(&_AlphaProVault.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AlphaProVault *AlphaProVaultCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AlphaProVault *AlphaProVaultSession) Token1() (common.Address, error) {
	return _AlphaProVault.Contract.Token1(&_AlphaProVault.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_AlphaProVault *AlphaProVaultCallerSession) Token1() (common.Address, error) {
	return _AlphaProVault.Contract.Token1(&_AlphaProVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AlphaProVault *AlphaProVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AlphaProVault *AlphaProVaultSession) TotalSupply() (*big.Int, error) {
	return _AlphaProVault.Contract.TotalSupply(&_AlphaProVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AlphaProVault *AlphaProVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _AlphaProVault.Contract.TotalSupply(&_AlphaProVault.CallOpts)
}

// TwapDuration is a free data retrieval call binding the contract method 0x26d89545.
//
// Solidity: function twapDuration() view returns(uint32)
func (_AlphaProVault *AlphaProVaultCaller) TwapDuration(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "twapDuration")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TwapDuration is a free data retrieval call binding the contract method 0x26d89545.
//
// Solidity: function twapDuration() view returns(uint32)
func (_AlphaProVault *AlphaProVaultSession) TwapDuration() (uint32, error) {
	return _AlphaProVault.Contract.TwapDuration(&_AlphaProVault.CallOpts)
}

// TwapDuration is a free data retrieval call binding the contract method 0x26d89545.
//
// Solidity: function twapDuration() view returns(uint32)
func (_AlphaProVault *AlphaProVaultCallerSession) TwapDuration() (uint32, error) {
	return _AlphaProVault.Contract.TwapDuration(&_AlphaProVault.CallOpts)
}

// WideLower is a free data retrieval call binding the contract method 0x4021ee8e.
//
// Solidity: function wideLower() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) WideLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "wideLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WideLower is a free data retrieval call binding the contract method 0x4021ee8e.
//
// Solidity: function wideLower() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) WideLower() (*big.Int, error) {
	return _AlphaProVault.Contract.WideLower(&_AlphaProVault.CallOpts)
}

// WideLower is a free data retrieval call binding the contract method 0x4021ee8e.
//
// Solidity: function wideLower() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) WideLower() (*big.Int, error) {
	return _AlphaProVault.Contract.WideLower(&_AlphaProVault.CallOpts)
}

// WideRangeWeight is a free data retrieval call binding the contract method 0x466360fc.
//
// Solidity: function wideRangeWeight() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCaller) WideRangeWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "wideRangeWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WideRangeWeight is a free data retrieval call binding the contract method 0x466360fc.
//
// Solidity: function wideRangeWeight() view returns(uint24)
func (_AlphaProVault *AlphaProVaultSession) WideRangeWeight() (*big.Int, error) {
	return _AlphaProVault.Contract.WideRangeWeight(&_AlphaProVault.CallOpts)
}

// WideRangeWeight is a free data retrieval call binding the contract method 0x466360fc.
//
// Solidity: function wideRangeWeight() view returns(uint24)
func (_AlphaProVault *AlphaProVaultCallerSession) WideRangeWeight() (*big.Int, error) {
	return _AlphaProVault.Contract.WideRangeWeight(&_AlphaProVault.CallOpts)
}

// WideThreshold is a free data retrieval call binding the contract method 0x6a37b9d0.
//
// Solidity: function wideThreshold() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) WideThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "wideThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WideThreshold is a free data retrieval call binding the contract method 0x6a37b9d0.
//
// Solidity: function wideThreshold() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) WideThreshold() (*big.Int, error) {
	return _AlphaProVault.Contract.WideThreshold(&_AlphaProVault.CallOpts)
}

// WideThreshold is a free data retrieval call binding the contract method 0x6a37b9d0.
//
// Solidity: function wideThreshold() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) WideThreshold() (*big.Int, error) {
	return _AlphaProVault.Contract.WideThreshold(&_AlphaProVault.CallOpts)
}

// WideUpper is a free data retrieval call binding the contract method 0x05bd3672.
//
// Solidity: function wideUpper() view returns(int24)
func (_AlphaProVault *AlphaProVaultCaller) WideUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlphaProVault.contract.Call(opts, &out, "wideUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WideUpper is a free data retrieval call binding the contract method 0x05bd3672.
//
// Solidity: function wideUpper() view returns(int24)
func (_AlphaProVault *AlphaProVaultSession) WideUpper() (*big.Int, error) {
	return _AlphaProVault.Contract.WideUpper(&_AlphaProVault.CallOpts)
}

// WideUpper is a free data retrieval call binding the contract method 0x05bd3672.
//
// Solidity: function wideUpper() view returns(int24)
func (_AlphaProVault *AlphaProVaultCallerSession) WideUpper() (*big.Int, error) {
	return _AlphaProVault.Contract.WideUpper(&_AlphaProVault.CallOpts)
}

// AcceptManager is a paid mutator transaction binding the contract method 0x48ff15b3.
//
// Solidity: function acceptManager() returns()
func (_AlphaProVault *AlphaProVaultTransactor) AcceptManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "acceptManager")
}

// AcceptManager is a paid mutator transaction binding the contract method 0x48ff15b3.
//
// Solidity: function acceptManager() returns()
func (_AlphaProVault *AlphaProVaultSession) AcceptManager() (*types.Transaction, error) {
	return _AlphaProVault.Contract.AcceptManager(&_AlphaProVault.TransactOpts)
}

// AcceptManager is a paid mutator transaction binding the contract method 0x48ff15b3.
//
// Solidity: function acceptManager() returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) AcceptManager() (*types.Transaction, error) {
	return _AlphaProVault.Contract.AcceptManager(&_AlphaProVault.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AlphaProVault *AlphaProVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AlphaProVault *AlphaProVaultSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Approve(&_AlphaProVault.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_AlphaProVault *AlphaProVaultTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Approve(&_AlphaProVault.TransactOpts, spender, value)
}

// CollectManager is a paid mutator transaction binding the contract method 0xfa9430e1.
//
// Solidity: function collectManager(address to) returns()
func (_AlphaProVault *AlphaProVaultTransactor) CollectManager(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "collectManager", to)
}

// CollectManager is a paid mutator transaction binding the contract method 0xfa9430e1.
//
// Solidity: function collectManager(address to) returns()
func (_AlphaProVault *AlphaProVaultSession) CollectManager(to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.CollectManager(&_AlphaProVault.TransactOpts, to)
}

// CollectManager is a paid mutator transaction binding the contract method 0xfa9430e1.
//
// Solidity: function collectManager(address to) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) CollectManager(to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.CollectManager(&_AlphaProVault.TransactOpts, to)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x1d769828.
//
// Solidity: function collectProtocol(address to) returns()
func (_AlphaProVault *AlphaProVaultTransactor) CollectProtocol(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "collectProtocol", to)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x1d769828.
//
// Solidity: function collectProtocol(address to) returns()
func (_AlphaProVault *AlphaProVaultSession) CollectProtocol(to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.CollectProtocol(&_AlphaProVault.TransactOpts, to)
}

// CollectProtocol is a paid mutator transaction binding the contract method 0x1d769828.
//
// Solidity: function collectProtocol(address to) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) CollectProtocol(to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.CollectProtocol(&_AlphaProVault.TransactOpts, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultTransactor) Deposit(opts *bind.TransactOpts, amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "deposit", amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultSession) Deposit(amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Deposit(&_AlphaProVault.TransactOpts, amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// Deposit is a paid mutator transaction binding the contract method 0x365d0ed7.
//
// Solidity: function deposit(uint256 amount0Desired, uint256 amount1Desired, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 shares, uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultTransactorSession) Deposit(amount0Desired *big.Int, amount1Desired *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Deposit(&_AlphaProVault.TransactOpts, amount0Desired, amount1Desired, amount0Min, amount1Min, to)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xabbffcb9.
//
// Solidity: function emergencyBurn(int24 tickLower, int24 tickUpper, uint128 liquidity) returns()
func (_AlphaProVault *AlphaProVaultTransactor) EmergencyBurn(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "emergencyBurn", tickLower, tickUpper, liquidity)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xabbffcb9.
//
// Solidity: function emergencyBurn(int24 tickLower, int24 tickUpper, uint128 liquidity) returns()
func (_AlphaProVault *AlphaProVaultSession) EmergencyBurn(tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.EmergencyBurn(&_AlphaProVault.TransactOpts, tickLower, tickUpper, liquidity)
}

// EmergencyBurn is a paid mutator transaction binding the contract method 0xabbffcb9.
//
// Solidity: function emergencyBurn(int24 tickLower, int24 tickUpper, uint128 liquidity) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) EmergencyBurn(tickLower *big.Int, tickUpper *big.Int, liquidity *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.EmergencyBurn(&_AlphaProVault.TransactOpts, tickLower, tickUpper, liquidity)
}

// Initialize is a paid mutator transaction binding the contract method 0x5fa3ffb5.
//
// Solidity: function initialize((address,address,uint24,uint256,uint24,int24,int24,int24,uint32,int24,int24,uint32,string,string) _params, address _factory) returns()
func (_AlphaProVault *AlphaProVaultTransactor) Initialize(opts *bind.TransactOpts, _params VaultParams, _factory common.Address) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "initialize", _params, _factory)
}

// Initialize is a paid mutator transaction binding the contract method 0x5fa3ffb5.
//
// Solidity: function initialize((address,address,uint24,uint256,uint24,int24,int24,int24,uint32,int24,int24,uint32,string,string) _params, address _factory) returns()
func (_AlphaProVault *AlphaProVaultSession) Initialize(_params VaultParams, _factory common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Initialize(&_AlphaProVault.TransactOpts, _params, _factory)
}

// Initialize is a paid mutator transaction binding the contract method 0x5fa3ffb5.
//
// Solidity: function initialize((address,address,uint24,uint256,uint24,int24,int24,int24,uint32,int24,int24,uint32,string,string) _params, address _factory) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) Initialize(_params VaultParams, _factory common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Initialize(&_AlphaProVault.TransactOpts, _params, _factory)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_AlphaProVault *AlphaProVaultTransactor) Rebalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "rebalance")
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_AlphaProVault *AlphaProVaultSession) Rebalance() (*types.Transaction, error) {
	return _AlphaProVault.Contract.Rebalance(&_AlphaProVault.TransactOpts)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7d7c2a1c.
//
// Solidity: function rebalance() returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) Rebalance() (*types.Transaction, error) {
	return _AlphaProVault.Contract.Rebalance(&_AlphaProVault.TransactOpts)
}

// SetBaseThreshold is a paid mutator transaction binding the contract method 0xa6329355.
//
// Solidity: function setBaseThreshold(int24 _baseThreshold) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetBaseThreshold(opts *bind.TransactOpts, _baseThreshold *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setBaseThreshold", _baseThreshold)
}

// SetBaseThreshold is a paid mutator transaction binding the contract method 0xa6329355.
//
// Solidity: function setBaseThreshold(int24 _baseThreshold) returns()
func (_AlphaProVault *AlphaProVaultSession) SetBaseThreshold(_baseThreshold *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetBaseThreshold(&_AlphaProVault.TransactOpts, _baseThreshold)
}

// SetBaseThreshold is a paid mutator transaction binding the contract method 0xa6329355.
//
// Solidity: function setBaseThreshold(int24 _baseThreshold) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetBaseThreshold(_baseThreshold *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetBaseThreshold(&_AlphaProVault.TransactOpts, _baseThreshold)
}

// SetDepositDelegate is a paid mutator transaction binding the contract method 0x5d9b6fd4.
//
// Solidity: function setDepositDelegate(address _depositDelegate) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetDepositDelegate(opts *bind.TransactOpts, _depositDelegate common.Address) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setDepositDelegate", _depositDelegate)
}

// SetDepositDelegate is a paid mutator transaction binding the contract method 0x5d9b6fd4.
//
// Solidity: function setDepositDelegate(address _depositDelegate) returns()
func (_AlphaProVault *AlphaProVaultSession) SetDepositDelegate(_depositDelegate common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetDepositDelegate(&_AlphaProVault.TransactOpts, _depositDelegate)
}

// SetDepositDelegate is a paid mutator transaction binding the contract method 0x5d9b6fd4.
//
// Solidity: function setDepositDelegate(address _depositDelegate) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetDepositDelegate(_depositDelegate common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetDepositDelegate(&_AlphaProVault.TransactOpts, _depositDelegate)
}

// SetLimitThreshold is a paid mutator transaction binding the contract method 0xaf794480.
//
// Solidity: function setLimitThreshold(int24 _limitThreshold) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetLimitThreshold(opts *bind.TransactOpts, _limitThreshold *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setLimitThreshold", _limitThreshold)
}

// SetLimitThreshold is a paid mutator transaction binding the contract method 0xaf794480.
//
// Solidity: function setLimitThreshold(int24 _limitThreshold) returns()
func (_AlphaProVault *AlphaProVaultSession) SetLimitThreshold(_limitThreshold *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetLimitThreshold(&_AlphaProVault.TransactOpts, _limitThreshold)
}

// SetLimitThreshold is a paid mutator transaction binding the contract method 0xaf794480.
//
// Solidity: function setLimitThreshold(int24 _limitThreshold) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetLimitThreshold(_limitThreshold *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetLimitThreshold(&_AlphaProVault.TransactOpts, _limitThreshold)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setManager", _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_AlphaProVault *AlphaProVaultSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetManager(&_AlphaProVault.TransactOpts, _manager)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address _manager) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetManager(_manager common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetManager(&_AlphaProVault.TransactOpts, _manager)
}

// SetManagerFee is a paid mutator transaction binding the contract method 0x16947fec.
//
// Solidity: function setManagerFee(uint24 _pendingManagerFee) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetManagerFee(opts *bind.TransactOpts, _pendingManagerFee *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setManagerFee", _pendingManagerFee)
}

// SetManagerFee is a paid mutator transaction binding the contract method 0x16947fec.
//
// Solidity: function setManagerFee(uint24 _pendingManagerFee) returns()
func (_AlphaProVault *AlphaProVaultSession) SetManagerFee(_pendingManagerFee *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetManagerFee(&_AlphaProVault.TransactOpts, _pendingManagerFee)
}

// SetManagerFee is a paid mutator transaction binding the contract method 0x16947fec.
//
// Solidity: function setManagerFee(uint24 _pendingManagerFee) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetManagerFee(_pendingManagerFee *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetManagerFee(&_AlphaProVault.TransactOpts, _pendingManagerFee)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetMaxTotalSupply(opts *bind.TransactOpts, _maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setMaxTotalSupply", _maxTotalSupply)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_AlphaProVault *AlphaProVaultSession) SetMaxTotalSupply(_maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetMaxTotalSupply(&_AlphaProVault.TransactOpts, _maxTotalSupply)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetMaxTotalSupply(_maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetMaxTotalSupply(&_AlphaProVault.TransactOpts, _maxTotalSupply)
}

// SetMaxTwapDeviation is a paid mutator transaction binding the contract method 0x3cbff3fe.
//
// Solidity: function setMaxTwapDeviation(int24 _maxTwapDeviation) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetMaxTwapDeviation(opts *bind.TransactOpts, _maxTwapDeviation *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setMaxTwapDeviation", _maxTwapDeviation)
}

// SetMaxTwapDeviation is a paid mutator transaction binding the contract method 0x3cbff3fe.
//
// Solidity: function setMaxTwapDeviation(int24 _maxTwapDeviation) returns()
func (_AlphaProVault *AlphaProVaultSession) SetMaxTwapDeviation(_maxTwapDeviation *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetMaxTwapDeviation(&_AlphaProVault.TransactOpts, _maxTwapDeviation)
}

// SetMaxTwapDeviation is a paid mutator transaction binding the contract method 0x3cbff3fe.
//
// Solidity: function setMaxTwapDeviation(int24 _maxTwapDeviation) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetMaxTwapDeviation(_maxTwapDeviation *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetMaxTwapDeviation(&_AlphaProVault.TransactOpts, _maxTwapDeviation)
}

// SetMinTickMove is a paid mutator transaction binding the contract method 0xf4005cae.
//
// Solidity: function setMinTickMove(int24 _minTickMove) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetMinTickMove(opts *bind.TransactOpts, _minTickMove *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setMinTickMove", _minTickMove)
}

// SetMinTickMove is a paid mutator transaction binding the contract method 0xf4005cae.
//
// Solidity: function setMinTickMove(int24 _minTickMove) returns()
func (_AlphaProVault *AlphaProVaultSession) SetMinTickMove(_minTickMove *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetMinTickMove(&_AlphaProVault.TransactOpts, _minTickMove)
}

// SetMinTickMove is a paid mutator transaction binding the contract method 0xf4005cae.
//
// Solidity: function setMinTickMove(int24 _minTickMove) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetMinTickMove(_minTickMove *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetMinTickMove(&_AlphaProVault.TransactOpts, _minTickMove)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xb8ec2d38.
//
// Solidity: function setPeriod(uint32 _period) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetPeriod(opts *bind.TransactOpts, _period uint32) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setPeriod", _period)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xb8ec2d38.
//
// Solidity: function setPeriod(uint32 _period) returns()
func (_AlphaProVault *AlphaProVaultSession) SetPeriod(_period uint32) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetPeriod(&_AlphaProVault.TransactOpts, _period)
}

// SetPeriod is a paid mutator transaction binding the contract method 0xb8ec2d38.
//
// Solidity: function setPeriod(uint32 _period) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetPeriod(_period uint32) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetPeriod(&_AlphaProVault.TransactOpts, _period)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x7feedaa1.
//
// Solidity: function setProtocolFee(uint24 _pendingProtocolFee) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetProtocolFee(opts *bind.TransactOpts, _pendingProtocolFee *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setProtocolFee", _pendingProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x7feedaa1.
//
// Solidity: function setProtocolFee(uint24 _pendingProtocolFee) returns()
func (_AlphaProVault *AlphaProVaultSession) SetProtocolFee(_pendingProtocolFee *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetProtocolFee(&_AlphaProVault.TransactOpts, _pendingProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x7feedaa1.
//
// Solidity: function setProtocolFee(uint24 _pendingProtocolFee) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetProtocolFee(_pendingProtocolFee *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetProtocolFee(&_AlphaProVault.TransactOpts, _pendingProtocolFee)
}

// SetRebalanceDelegate is a paid mutator transaction binding the contract method 0xf8d01c9d.
//
// Solidity: function setRebalanceDelegate(address _rebalanceDelegate) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetRebalanceDelegate(opts *bind.TransactOpts, _rebalanceDelegate common.Address) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setRebalanceDelegate", _rebalanceDelegate)
}

// SetRebalanceDelegate is a paid mutator transaction binding the contract method 0xf8d01c9d.
//
// Solidity: function setRebalanceDelegate(address _rebalanceDelegate) returns()
func (_AlphaProVault *AlphaProVaultSession) SetRebalanceDelegate(_rebalanceDelegate common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetRebalanceDelegate(&_AlphaProVault.TransactOpts, _rebalanceDelegate)
}

// SetRebalanceDelegate is a paid mutator transaction binding the contract method 0xf8d01c9d.
//
// Solidity: function setRebalanceDelegate(address _rebalanceDelegate) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetRebalanceDelegate(_rebalanceDelegate common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetRebalanceDelegate(&_AlphaProVault.TransactOpts, _rebalanceDelegate)
}

// SetTwapDuration is a paid mutator transaction binding the contract method 0xc433c80a.
//
// Solidity: function setTwapDuration(uint32 _twapDuration) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetTwapDuration(opts *bind.TransactOpts, _twapDuration uint32) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setTwapDuration", _twapDuration)
}

// SetTwapDuration is a paid mutator transaction binding the contract method 0xc433c80a.
//
// Solidity: function setTwapDuration(uint32 _twapDuration) returns()
func (_AlphaProVault *AlphaProVaultSession) SetTwapDuration(_twapDuration uint32) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetTwapDuration(&_AlphaProVault.TransactOpts, _twapDuration)
}

// SetTwapDuration is a paid mutator transaction binding the contract method 0xc433c80a.
//
// Solidity: function setTwapDuration(uint32 _twapDuration) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetTwapDuration(_twapDuration uint32) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetTwapDuration(&_AlphaProVault.TransactOpts, _twapDuration)
}

// SetWideRangeWeight is a paid mutator transaction binding the contract method 0xc24e55b5.
//
// Solidity: function setWideRangeWeight(uint24 _wideRangeWeight) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetWideRangeWeight(opts *bind.TransactOpts, _wideRangeWeight *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setWideRangeWeight", _wideRangeWeight)
}

// SetWideRangeWeight is a paid mutator transaction binding the contract method 0xc24e55b5.
//
// Solidity: function setWideRangeWeight(uint24 _wideRangeWeight) returns()
func (_AlphaProVault *AlphaProVaultSession) SetWideRangeWeight(_wideRangeWeight *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetWideRangeWeight(&_AlphaProVault.TransactOpts, _wideRangeWeight)
}

// SetWideRangeWeight is a paid mutator transaction binding the contract method 0xc24e55b5.
//
// Solidity: function setWideRangeWeight(uint24 _wideRangeWeight) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetWideRangeWeight(_wideRangeWeight *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetWideRangeWeight(&_AlphaProVault.TransactOpts, _wideRangeWeight)
}

// SetWideThreshold is a paid mutator transaction binding the contract method 0xc4b4c30f.
//
// Solidity: function setWideThreshold(int24 _wideThreshold) returns()
func (_AlphaProVault *AlphaProVaultTransactor) SetWideThreshold(opts *bind.TransactOpts, _wideThreshold *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "setWideThreshold", _wideThreshold)
}

// SetWideThreshold is a paid mutator transaction binding the contract method 0xc4b4c30f.
//
// Solidity: function setWideThreshold(int24 _wideThreshold) returns()
func (_AlphaProVault *AlphaProVaultSession) SetWideThreshold(_wideThreshold *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetWideThreshold(&_AlphaProVault.TransactOpts, _wideThreshold)
}

// SetWideThreshold is a paid mutator transaction binding the contract method 0xc4b4c30f.
//
// Solidity: function setWideThreshold(int24 _wideThreshold) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) SetWideThreshold(_wideThreshold *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.SetWideThreshold(&_AlphaProVault.TransactOpts, _wideThreshold)
}

// Sweep is a paid mutator transaction binding the contract method 0xdc2c256f.
//
// Solidity: function sweep(address token, uint256 amount, address to) returns()
func (_AlphaProVault *AlphaProVaultTransactor) Sweep(opts *bind.TransactOpts, token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "sweep", token, amount, to)
}

// Sweep is a paid mutator transaction binding the contract method 0xdc2c256f.
//
// Solidity: function sweep(address token, uint256 amount, address to) returns()
func (_AlphaProVault *AlphaProVaultSession) Sweep(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Sweep(&_AlphaProVault.TransactOpts, token, amount, to)
}

// Sweep is a paid mutator transaction binding the contract method 0xdc2c256f.
//
// Solidity: function sweep(address token, uint256 amount, address to) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) Sweep(token common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Sweep(&_AlphaProVault.TransactOpts, token, amount, to)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AlphaProVault *AlphaProVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AlphaProVault *AlphaProVaultSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Transfer(&_AlphaProVault.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_AlphaProVault *AlphaProVaultTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Transfer(&_AlphaProVault.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AlphaProVault *AlphaProVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AlphaProVault *AlphaProVaultSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.TransferFrom(&_AlphaProVault.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_AlphaProVault *AlphaProVaultTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _AlphaProVault.Contract.TransferFrom(&_AlphaProVault.TransactOpts, from, to, value)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_AlphaProVault *AlphaProVaultTransactor) UniswapV3MintCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "uniswapV3MintCallback", amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_AlphaProVault *AlphaProVaultSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AlphaProVault.Contract.UniswapV3MintCallback(&_AlphaProVault.TransactOpts, amount0, amount1, data)
}

// UniswapV3MintCallback is a paid mutator transaction binding the contract method 0xd3487997.
//
// Solidity: function uniswapV3MintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) UniswapV3MintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _AlphaProVault.Contract.UniswapV3MintCallback(&_AlphaProVault.TransactOpts, amount0, amount1, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_AlphaProVault *AlphaProVaultTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_AlphaProVault *AlphaProVaultSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _AlphaProVault.Contract.UniswapV3SwapCallback(&_AlphaProVault.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_AlphaProVault *AlphaProVaultTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _AlphaProVault.Contract.UniswapV3SwapCallback(&_AlphaProVault.TransactOpts, amount0Delta, amount1Delta, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultTransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.contract.Transact(opts, "withdraw", shares, amount0Min, amount1Min, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultSession) Withdraw(shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Withdraw(&_AlphaProVault.TransactOpts, shares, amount0Min, amount1Min, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd331bef7.
//
// Solidity: function withdraw(uint256 shares, uint256 amount0Min, uint256 amount1Min, address to) returns(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultTransactorSession) Withdraw(shares *big.Int, amount0Min *big.Int, amount1Min *big.Int, to common.Address) (*types.Transaction, error) {
	return _AlphaProVault.Contract.Withdraw(&_AlphaProVault.TransactOpts, shares, amount0Min, amount1Min, to)
}

// AlphaProVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AlphaProVault contract.
type AlphaProVaultApprovalIterator struct {
	Event *AlphaProVaultApproval // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultApproval)
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
		it.Event = new(AlphaProVaultApproval)
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
func (it *AlphaProVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultApproval represents a Approval event raised by the AlphaProVault contract.
type AlphaProVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AlphaProVault *AlphaProVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AlphaProVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultApprovalIterator{contract: _AlphaProVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AlphaProVault *AlphaProVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AlphaProVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultApproval)
				if err := _AlphaProVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_AlphaProVault *AlphaProVaultFilterer) ParseApproval(log types.Log) (*AlphaProVaultApproval, error) {
	event := new(AlphaProVaultApproval)
	if err := _AlphaProVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultCollectFeesIterator is returned from FilterCollectFees and is used to iterate over the raw logs and unpacked data for CollectFees events raised by the AlphaProVault contract.
type AlphaProVaultCollectFeesIterator struct {
	Event *AlphaProVaultCollectFees // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultCollectFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultCollectFees)
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
		it.Event = new(AlphaProVaultCollectFees)
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
func (it *AlphaProVaultCollectFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultCollectFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultCollectFees represents a CollectFees event raised by the AlphaProVault contract.
type AlphaProVaultCollectFees struct {
	FeesToVault0    *big.Int
	FeesToVault1    *big.Int
	FeesToProtocol0 *big.Int
	FeesToProtocol1 *big.Int
	FeesToManager0  *big.Int
	FeesToManager1  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCollectFees is a free log retrieval operation binding the contract event 0x3d6c49c42d0f03a60a09b6137b84abbd7a0bf738324602cb48343f830e3db470.
//
// Solidity: event CollectFees(uint256 feesToVault0, uint256 feesToVault1, uint256 feesToProtocol0, uint256 feesToProtocol1, uint256 feesToManager0, uint256 feesToManager1)
func (_AlphaProVault *AlphaProVaultFilterer) FilterCollectFees(opts *bind.FilterOpts) (*AlphaProVaultCollectFeesIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "CollectFees")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultCollectFeesIterator{contract: _AlphaProVault.contract, event: "CollectFees", logs: logs, sub: sub}, nil
}

// WatchCollectFees is a free log subscription operation binding the contract event 0x3d6c49c42d0f03a60a09b6137b84abbd7a0bf738324602cb48343f830e3db470.
//
// Solidity: event CollectFees(uint256 feesToVault0, uint256 feesToVault1, uint256 feesToProtocol0, uint256 feesToProtocol1, uint256 feesToManager0, uint256 feesToManager1)
func (_AlphaProVault *AlphaProVaultFilterer) WatchCollectFees(opts *bind.WatchOpts, sink chan<- *AlphaProVaultCollectFees) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "CollectFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultCollectFees)
				if err := _AlphaProVault.contract.UnpackLog(event, "CollectFees", log); err != nil {
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

// ParseCollectFees is a log parse operation binding the contract event 0x3d6c49c42d0f03a60a09b6137b84abbd7a0bf738324602cb48343f830e3db470.
//
// Solidity: event CollectFees(uint256 feesToVault0, uint256 feesToVault1, uint256 feesToProtocol0, uint256 feesToProtocol1, uint256 feesToManager0, uint256 feesToManager1)
func (_AlphaProVault *AlphaProVaultFilterer) ParseCollectFees(log types.Log) (*AlphaProVaultCollectFees, error) {
	event := new(AlphaProVaultCollectFees)
	if err := _AlphaProVault.contract.UnpackLog(event, "CollectFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultCollectManagerIterator is returned from FilterCollectManager and is used to iterate over the raw logs and unpacked data for CollectManager events raised by the AlphaProVault contract.
type AlphaProVaultCollectManagerIterator struct {
	Event *AlphaProVaultCollectManager // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultCollectManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultCollectManager)
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
		it.Event = new(AlphaProVaultCollectManager)
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
func (it *AlphaProVaultCollectManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultCollectManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultCollectManager represents a CollectManager event raised by the AlphaProVault contract.
type AlphaProVaultCollectManager struct {
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCollectManager is a free log retrieval operation binding the contract event 0x17f6807ba08fcbc6bc82aff717379b30802e04cff0756017bf22b2a490d0ab65.
//
// Solidity: event CollectManager(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultFilterer) FilterCollectManager(opts *bind.FilterOpts) (*AlphaProVaultCollectManagerIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "CollectManager")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultCollectManagerIterator{contract: _AlphaProVault.contract, event: "CollectManager", logs: logs, sub: sub}, nil
}

// WatchCollectManager is a free log subscription operation binding the contract event 0x17f6807ba08fcbc6bc82aff717379b30802e04cff0756017bf22b2a490d0ab65.
//
// Solidity: event CollectManager(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultFilterer) WatchCollectManager(opts *bind.WatchOpts, sink chan<- *AlphaProVaultCollectManager) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "CollectManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultCollectManager)
				if err := _AlphaProVault.contract.UnpackLog(event, "CollectManager", log); err != nil {
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

// ParseCollectManager is a log parse operation binding the contract event 0x17f6807ba08fcbc6bc82aff717379b30802e04cff0756017bf22b2a490d0ab65.
//
// Solidity: event CollectManager(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultFilterer) ParseCollectManager(log types.Log) (*AlphaProVaultCollectManager, error) {
	event := new(AlphaProVaultCollectManager)
	if err := _AlphaProVault.contract.UnpackLog(event, "CollectManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultCollectProtocolIterator is returned from FilterCollectProtocol and is used to iterate over the raw logs and unpacked data for CollectProtocol events raised by the AlphaProVault contract.
type AlphaProVaultCollectProtocolIterator struct {
	Event *AlphaProVaultCollectProtocol // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultCollectProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultCollectProtocol)
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
		it.Event = new(AlphaProVaultCollectProtocol)
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
func (it *AlphaProVaultCollectProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultCollectProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultCollectProtocol represents a CollectProtocol event raised by the AlphaProVault contract.
type AlphaProVaultCollectProtocol struct {
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCollectProtocol is a free log retrieval operation binding the contract event 0xd63986ca13f11502796aee70ba80ac7317d99f08e5871fd9fd602a2764c7ef30.
//
// Solidity: event CollectProtocol(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultFilterer) FilterCollectProtocol(opts *bind.FilterOpts) (*AlphaProVaultCollectProtocolIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "CollectProtocol")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultCollectProtocolIterator{contract: _AlphaProVault.contract, event: "CollectProtocol", logs: logs, sub: sub}, nil
}

// WatchCollectProtocol is a free log subscription operation binding the contract event 0xd63986ca13f11502796aee70ba80ac7317d99f08e5871fd9fd602a2764c7ef30.
//
// Solidity: event CollectProtocol(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultFilterer) WatchCollectProtocol(opts *bind.WatchOpts, sink chan<- *AlphaProVaultCollectProtocol) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "CollectProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultCollectProtocol)
				if err := _AlphaProVault.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
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

// ParseCollectProtocol is a log parse operation binding the contract event 0xd63986ca13f11502796aee70ba80ac7317d99f08e5871fd9fd602a2764c7ef30.
//
// Solidity: event CollectProtocol(uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultFilterer) ParseCollectProtocol(log types.Log) (*AlphaProVaultCollectProtocol, error) {
	event := new(AlphaProVaultCollectProtocol)
	if err := _AlphaProVault.contract.UnpackLog(event, "CollectProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the AlphaProVault contract.
type AlphaProVaultDepositIterator struct {
	Event *AlphaProVaultDeposit // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultDeposit)
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
		it.Event = new(AlphaProVaultDeposit)
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
func (it *AlphaProVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultDeposit represents a Deposit event raised by the AlphaProVault contract.
type AlphaProVaultDeposit struct {
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
func (_AlphaProVault *AlphaProVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*AlphaProVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultDepositIterator{contract: _AlphaProVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AlphaProVaultDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultDeposit)
				if err := _AlphaProVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_AlphaProVault *AlphaProVaultFilterer) ParseDeposit(log types.Log) (*AlphaProVaultDeposit, error) {
	event := new(AlphaProVaultDeposit)
	if err := _AlphaProVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AlphaProVault contract.
type AlphaProVaultInitializedIterator struct {
	Event *AlphaProVaultInitialized // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultInitialized)
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
		it.Event = new(AlphaProVaultInitialized)
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
func (it *AlphaProVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultInitialized represents a Initialized event raised by the AlphaProVault contract.
type AlphaProVaultInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AlphaProVault *AlphaProVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*AlphaProVaultInitializedIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultInitializedIterator{contract: _AlphaProVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AlphaProVault *AlphaProVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AlphaProVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultInitialized)
				if err := _AlphaProVault.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_AlphaProVault *AlphaProVaultFilterer) ParseInitialized(log types.Log) (*AlphaProVaultInitialized, error) {
	event := new(AlphaProVaultInitialized)
	if err := _AlphaProVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultSnapshotIterator is returned from FilterSnapshot and is used to iterate over the raw logs and unpacked data for Snapshot events raised by the AlphaProVault contract.
type AlphaProVaultSnapshotIterator struct {
	Event *AlphaProVaultSnapshot // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultSnapshot)
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
		it.Event = new(AlphaProVaultSnapshot)
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
func (it *AlphaProVaultSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultSnapshot represents a Snapshot event raised by the AlphaProVault contract.
type AlphaProVaultSnapshot struct {
	Tick         *big.Int
	TotalAmount0 *big.Int
	TotalAmount1 *big.Int
	TotalSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSnapshot is a free log retrieval operation binding the contract event 0x210f60adf1db7a02e9db9a49ec7c2eb2060c516cbcfd01a0c05288144738ee5d.
//
// Solidity: event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_AlphaProVault *AlphaProVaultFilterer) FilterSnapshot(opts *bind.FilterOpts) (*AlphaProVaultSnapshotIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultSnapshotIterator{contract: _AlphaProVault.contract, event: "Snapshot", logs: logs, sub: sub}, nil
}

// WatchSnapshot is a free log subscription operation binding the contract event 0x210f60adf1db7a02e9db9a49ec7c2eb2060c516cbcfd01a0c05288144738ee5d.
//
// Solidity: event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_AlphaProVault *AlphaProVaultFilterer) WatchSnapshot(opts *bind.WatchOpts, sink chan<- *AlphaProVaultSnapshot) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "Snapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultSnapshot)
				if err := _AlphaProVault.contract.UnpackLog(event, "Snapshot", log); err != nil {
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

// ParseSnapshot is a log parse operation binding the contract event 0x210f60adf1db7a02e9db9a49ec7c2eb2060c516cbcfd01a0c05288144738ee5d.
//
// Solidity: event Snapshot(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 totalSupply)
func (_AlphaProVault *AlphaProVaultFilterer) ParseSnapshot(log types.Log) (*AlphaProVaultSnapshot, error) {
	event := new(AlphaProVaultSnapshot)
	if err := _AlphaProVault.contract.UnpackLog(event, "Snapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AlphaProVault contract.
type AlphaProVaultTransferIterator struct {
	Event *AlphaProVaultTransfer // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultTransfer)
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
		it.Event = new(AlphaProVaultTransfer)
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
func (it *AlphaProVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultTransfer represents a Transfer event raised by the AlphaProVault contract.
type AlphaProVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AlphaProVault *AlphaProVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AlphaProVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultTransferIterator{contract: _AlphaProVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AlphaProVault *AlphaProVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AlphaProVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultTransfer)
				if err := _AlphaProVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_AlphaProVault *AlphaProVaultFilterer) ParseTransfer(log types.Log) (*AlphaProVaultTransfer, error) {
	event := new(AlphaProVaultTransfer)
	if err := _AlphaProVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateBaseThresholdIterator is returned from FilterUpdateBaseThreshold and is used to iterate over the raw logs and unpacked data for UpdateBaseThreshold events raised by the AlphaProVault contract.
type AlphaProVaultUpdateBaseThresholdIterator struct {
	Event *AlphaProVaultUpdateBaseThreshold // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateBaseThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateBaseThreshold)
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
		it.Event = new(AlphaProVaultUpdateBaseThreshold)
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
func (it *AlphaProVaultUpdateBaseThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateBaseThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateBaseThreshold represents a UpdateBaseThreshold event raised by the AlphaProVault contract.
type AlphaProVaultUpdateBaseThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateBaseThreshold is a free log retrieval operation binding the contract event 0xf1759909677b9c42577caba6b12efd5bcf3a398d02a2c1c97d23bbd312b561a7.
//
// Solidity: event UpdateBaseThreshold(int24 threshold)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateBaseThreshold(opts *bind.FilterOpts) (*AlphaProVaultUpdateBaseThresholdIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateBaseThreshold")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateBaseThresholdIterator{contract: _AlphaProVault.contract, event: "UpdateBaseThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdateBaseThreshold is a free log subscription operation binding the contract event 0xf1759909677b9c42577caba6b12efd5bcf3a398d02a2c1c97d23bbd312b561a7.
//
// Solidity: event UpdateBaseThreshold(int24 threshold)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateBaseThreshold(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateBaseThreshold) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateBaseThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateBaseThreshold)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateBaseThreshold", log); err != nil {
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

// ParseUpdateBaseThreshold is a log parse operation binding the contract event 0xf1759909677b9c42577caba6b12efd5bcf3a398d02a2c1c97d23bbd312b561a7.
//
// Solidity: event UpdateBaseThreshold(int24 threshold)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateBaseThreshold(log types.Log) (*AlphaProVaultUpdateBaseThreshold, error) {
	event := new(AlphaProVaultUpdateBaseThreshold)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateBaseThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateDepositDelegateIterator is returned from FilterUpdateDepositDelegate and is used to iterate over the raw logs and unpacked data for UpdateDepositDelegate events raised by the AlphaProVault contract.
type AlphaProVaultUpdateDepositDelegateIterator struct {
	Event *AlphaProVaultUpdateDepositDelegate // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateDepositDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateDepositDelegate)
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
		it.Event = new(AlphaProVaultUpdateDepositDelegate)
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
func (it *AlphaProVaultUpdateDepositDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateDepositDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateDepositDelegate represents a UpdateDepositDelegate event raised by the AlphaProVault contract.
type AlphaProVaultUpdateDepositDelegate struct {
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositDelegate is a free log retrieval operation binding the contract event 0x3b3e178c01607aa6eab8fd24b33e3de1d5f41d603b21010e65903bdce27dec13.
//
// Solidity: event UpdateDepositDelegate(address delegate)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateDepositDelegate(opts *bind.FilterOpts) (*AlphaProVaultUpdateDepositDelegateIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateDepositDelegate")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateDepositDelegateIterator{contract: _AlphaProVault.contract, event: "UpdateDepositDelegate", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositDelegate is a free log subscription operation binding the contract event 0x3b3e178c01607aa6eab8fd24b33e3de1d5f41d603b21010e65903bdce27dec13.
//
// Solidity: event UpdateDepositDelegate(address delegate)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateDepositDelegate(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateDepositDelegate) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateDepositDelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateDepositDelegate)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateDepositDelegate", log); err != nil {
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

// ParseUpdateDepositDelegate is a log parse operation binding the contract event 0x3b3e178c01607aa6eab8fd24b33e3de1d5f41d603b21010e65903bdce27dec13.
//
// Solidity: event UpdateDepositDelegate(address delegate)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateDepositDelegate(log types.Log) (*AlphaProVaultUpdateDepositDelegate, error) {
	event := new(AlphaProVaultUpdateDepositDelegate)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateDepositDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateLimitThresholdIterator is returned from FilterUpdateLimitThreshold and is used to iterate over the raw logs and unpacked data for UpdateLimitThreshold events raised by the AlphaProVault contract.
type AlphaProVaultUpdateLimitThresholdIterator struct {
	Event *AlphaProVaultUpdateLimitThreshold // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateLimitThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateLimitThreshold)
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
		it.Event = new(AlphaProVaultUpdateLimitThreshold)
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
func (it *AlphaProVaultUpdateLimitThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateLimitThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateLimitThreshold represents a UpdateLimitThreshold event raised by the AlphaProVault contract.
type AlphaProVaultUpdateLimitThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateLimitThreshold is a free log retrieval operation binding the contract event 0xc797ac22933a04255304effee6a7d3c3e4bbfaa0b2897c52c61a8b1fb027b9f2.
//
// Solidity: event UpdateLimitThreshold(int24 threshold)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateLimitThreshold(opts *bind.FilterOpts) (*AlphaProVaultUpdateLimitThresholdIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateLimitThreshold")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateLimitThresholdIterator{contract: _AlphaProVault.contract, event: "UpdateLimitThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdateLimitThreshold is a free log subscription operation binding the contract event 0xc797ac22933a04255304effee6a7d3c3e4bbfaa0b2897c52c61a8b1fb027b9f2.
//
// Solidity: event UpdateLimitThreshold(int24 threshold)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateLimitThreshold(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateLimitThreshold) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateLimitThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateLimitThreshold)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateLimitThreshold", log); err != nil {
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

// ParseUpdateLimitThreshold is a log parse operation binding the contract event 0xc797ac22933a04255304effee6a7d3c3e4bbfaa0b2897c52c61a8b1fb027b9f2.
//
// Solidity: event UpdateLimitThreshold(int24 threshold)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateLimitThreshold(log types.Log) (*AlphaProVaultUpdateLimitThreshold, error) {
	event := new(AlphaProVaultUpdateLimitThreshold)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateLimitThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateManagerIterator is returned from FilterUpdateManager and is used to iterate over the raw logs and unpacked data for UpdateManager events raised by the AlphaProVault contract.
type AlphaProVaultUpdateManagerIterator struct {
	Event *AlphaProVaultUpdateManager // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateManager)
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
		it.Event = new(AlphaProVaultUpdateManager)
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
func (it *AlphaProVaultUpdateManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateManager represents a UpdateManager event raised by the AlphaProVault contract.
type AlphaProVaultUpdateManager struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateManager is a free log retrieval operation binding the contract event 0x5c18ab5c697b63d102fc7e14c77bfaef0f1013206eca139920fd389277814e09.
//
// Solidity: event UpdateManager(address manager)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateManager(opts *bind.FilterOpts) (*AlphaProVaultUpdateManagerIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateManager")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateManagerIterator{contract: _AlphaProVault.contract, event: "UpdateManager", logs: logs, sub: sub}, nil
}

// WatchUpdateManager is a free log subscription operation binding the contract event 0x5c18ab5c697b63d102fc7e14c77bfaef0f1013206eca139920fd389277814e09.
//
// Solidity: event UpdateManager(address manager)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateManager(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateManager) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateManager)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateManager", log); err != nil {
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

// ParseUpdateManager is a log parse operation binding the contract event 0x5c18ab5c697b63d102fc7e14c77bfaef0f1013206eca139920fd389277814e09.
//
// Solidity: event UpdateManager(address manager)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateManager(log types.Log) (*AlphaProVaultUpdateManager, error) {
	event := new(AlphaProVaultUpdateManager)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateManagerFeeIterator is returned from FilterUpdateManagerFee and is used to iterate over the raw logs and unpacked data for UpdateManagerFee events raised by the AlphaProVault contract.
type AlphaProVaultUpdateManagerFeeIterator struct {
	Event *AlphaProVaultUpdateManagerFee // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateManagerFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateManagerFee)
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
		it.Event = new(AlphaProVaultUpdateManagerFee)
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
func (it *AlphaProVaultUpdateManagerFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateManagerFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateManagerFee represents a UpdateManagerFee event raised by the AlphaProVault contract.
type AlphaProVaultUpdateManagerFee struct {
	ManagerFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateManagerFee is a free log retrieval operation binding the contract event 0x218e2de5d78d8a72c25f717811481b5563d21963b758fddb0643d18fa38c7a67.
//
// Solidity: event UpdateManagerFee(uint24 managerFee)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateManagerFee(opts *bind.FilterOpts) (*AlphaProVaultUpdateManagerFeeIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateManagerFee")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateManagerFeeIterator{contract: _AlphaProVault.contract, event: "UpdateManagerFee", logs: logs, sub: sub}, nil
}

// WatchUpdateManagerFee is a free log subscription operation binding the contract event 0x218e2de5d78d8a72c25f717811481b5563d21963b758fddb0643d18fa38c7a67.
//
// Solidity: event UpdateManagerFee(uint24 managerFee)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateManagerFee(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateManagerFee) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateManagerFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateManagerFee)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateManagerFee", log); err != nil {
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

// ParseUpdateManagerFee is a log parse operation binding the contract event 0x218e2de5d78d8a72c25f717811481b5563d21963b758fddb0643d18fa38c7a67.
//
// Solidity: event UpdateManagerFee(uint24 managerFee)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateManagerFee(log types.Log) (*AlphaProVaultUpdateManagerFee, error) {
	event := new(AlphaProVaultUpdateManagerFee)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateManagerFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateMaxTotalSupplyIterator is returned from FilterUpdateMaxTotalSupply and is used to iterate over the raw logs and unpacked data for UpdateMaxTotalSupply events raised by the AlphaProVault contract.
type AlphaProVaultUpdateMaxTotalSupplyIterator struct {
	Event *AlphaProVaultUpdateMaxTotalSupply // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateMaxTotalSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateMaxTotalSupply)
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
		it.Event = new(AlphaProVaultUpdateMaxTotalSupply)
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
func (it *AlphaProVaultUpdateMaxTotalSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateMaxTotalSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateMaxTotalSupply represents a UpdateMaxTotalSupply event raised by the AlphaProVault contract.
type AlphaProVaultUpdateMaxTotalSupply struct {
	MaxTotalSupply *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxTotalSupply is a free log retrieval operation binding the contract event 0x49e8da6bc2b1ffc75cb81c88d1a8e64d5b1224c626dc9be8787d6ff982b46a39.
//
// Solidity: event UpdateMaxTotalSupply(uint256 maxTotalSupply)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateMaxTotalSupply(opts *bind.FilterOpts) (*AlphaProVaultUpdateMaxTotalSupplyIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateMaxTotalSupply")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateMaxTotalSupplyIterator{contract: _AlphaProVault.contract, event: "UpdateMaxTotalSupply", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxTotalSupply is a free log subscription operation binding the contract event 0x49e8da6bc2b1ffc75cb81c88d1a8e64d5b1224c626dc9be8787d6ff982b46a39.
//
// Solidity: event UpdateMaxTotalSupply(uint256 maxTotalSupply)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateMaxTotalSupply(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateMaxTotalSupply) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateMaxTotalSupply")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateMaxTotalSupply)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateMaxTotalSupply", log); err != nil {
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

// ParseUpdateMaxTotalSupply is a log parse operation binding the contract event 0x49e8da6bc2b1ffc75cb81c88d1a8e64d5b1224c626dc9be8787d6ff982b46a39.
//
// Solidity: event UpdateMaxTotalSupply(uint256 maxTotalSupply)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateMaxTotalSupply(log types.Log) (*AlphaProVaultUpdateMaxTotalSupply, error) {
	event := new(AlphaProVaultUpdateMaxTotalSupply)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateMaxTotalSupply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateMaxTwapDeviationIterator is returned from FilterUpdateMaxTwapDeviation and is used to iterate over the raw logs and unpacked data for UpdateMaxTwapDeviation events raised by the AlphaProVault contract.
type AlphaProVaultUpdateMaxTwapDeviationIterator struct {
	Event *AlphaProVaultUpdateMaxTwapDeviation // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateMaxTwapDeviationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateMaxTwapDeviation)
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
		it.Event = new(AlphaProVaultUpdateMaxTwapDeviation)
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
func (it *AlphaProVaultUpdateMaxTwapDeviationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateMaxTwapDeviationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateMaxTwapDeviation represents a UpdateMaxTwapDeviation event raised by the AlphaProVault contract.
type AlphaProVaultUpdateMaxTwapDeviation struct {
	MaxTwapDeviation *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxTwapDeviation is a free log retrieval operation binding the contract event 0x957dddabc1ac2b52bf67ebbe53150c1b46f57ec0e1dc487632d5bcc3a36a2d82.
//
// Solidity: event UpdateMaxTwapDeviation(int24 maxTwapDeviation)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateMaxTwapDeviation(opts *bind.FilterOpts) (*AlphaProVaultUpdateMaxTwapDeviationIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateMaxTwapDeviation")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateMaxTwapDeviationIterator{contract: _AlphaProVault.contract, event: "UpdateMaxTwapDeviation", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxTwapDeviation is a free log subscription operation binding the contract event 0x957dddabc1ac2b52bf67ebbe53150c1b46f57ec0e1dc487632d5bcc3a36a2d82.
//
// Solidity: event UpdateMaxTwapDeviation(int24 maxTwapDeviation)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateMaxTwapDeviation(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateMaxTwapDeviation) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateMaxTwapDeviation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateMaxTwapDeviation)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateMaxTwapDeviation", log); err != nil {
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

// ParseUpdateMaxTwapDeviation is a log parse operation binding the contract event 0x957dddabc1ac2b52bf67ebbe53150c1b46f57ec0e1dc487632d5bcc3a36a2d82.
//
// Solidity: event UpdateMaxTwapDeviation(int24 maxTwapDeviation)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateMaxTwapDeviation(log types.Log) (*AlphaProVaultUpdateMaxTwapDeviation, error) {
	event := new(AlphaProVaultUpdateMaxTwapDeviation)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateMaxTwapDeviation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateMinTickMoveIterator is returned from FilterUpdateMinTickMove and is used to iterate over the raw logs and unpacked data for UpdateMinTickMove events raised by the AlphaProVault contract.
type AlphaProVaultUpdateMinTickMoveIterator struct {
	Event *AlphaProVaultUpdateMinTickMove // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateMinTickMoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateMinTickMove)
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
		it.Event = new(AlphaProVaultUpdateMinTickMove)
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
func (it *AlphaProVaultUpdateMinTickMoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateMinTickMoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateMinTickMove represents a UpdateMinTickMove event raised by the AlphaProVault contract.
type AlphaProVaultUpdateMinTickMove struct {
	MinTickMove *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinTickMove is a free log retrieval operation binding the contract event 0x52a3cbe96b59de1055b9b043e8906557387b821f26af47beea677164a7f26b62.
//
// Solidity: event UpdateMinTickMove(int24 minTickMove)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateMinTickMove(opts *bind.FilterOpts) (*AlphaProVaultUpdateMinTickMoveIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateMinTickMove")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateMinTickMoveIterator{contract: _AlphaProVault.contract, event: "UpdateMinTickMove", logs: logs, sub: sub}, nil
}

// WatchUpdateMinTickMove is a free log subscription operation binding the contract event 0x52a3cbe96b59de1055b9b043e8906557387b821f26af47beea677164a7f26b62.
//
// Solidity: event UpdateMinTickMove(int24 minTickMove)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateMinTickMove(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateMinTickMove) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateMinTickMove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateMinTickMove)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateMinTickMove", log); err != nil {
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

// ParseUpdateMinTickMove is a log parse operation binding the contract event 0x52a3cbe96b59de1055b9b043e8906557387b821f26af47beea677164a7f26b62.
//
// Solidity: event UpdateMinTickMove(int24 minTickMove)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateMinTickMove(log types.Log) (*AlphaProVaultUpdateMinTickMove, error) {
	event := new(AlphaProVaultUpdateMinTickMove)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateMinTickMove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdatePendingManagerIterator is returned from FilterUpdatePendingManager and is used to iterate over the raw logs and unpacked data for UpdatePendingManager events raised by the AlphaProVault contract.
type AlphaProVaultUpdatePendingManagerIterator struct {
	Event *AlphaProVaultUpdatePendingManager // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdatePendingManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdatePendingManager)
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
		it.Event = new(AlphaProVaultUpdatePendingManager)
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
func (it *AlphaProVaultUpdatePendingManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdatePendingManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdatePendingManager represents a UpdatePendingManager event raised by the AlphaProVault contract.
type AlphaProVaultUpdatePendingManager struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdatePendingManager is a free log retrieval operation binding the contract event 0x4d3334a0a69f5f1c54636cf743914f0b34fb2e7849b55ee7c1faddd0e06b4dfd.
//
// Solidity: event UpdatePendingManager(address manager)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdatePendingManager(opts *bind.FilterOpts) (*AlphaProVaultUpdatePendingManagerIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdatePendingManager")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdatePendingManagerIterator{contract: _AlphaProVault.contract, event: "UpdatePendingManager", logs: logs, sub: sub}, nil
}

// WatchUpdatePendingManager is a free log subscription operation binding the contract event 0x4d3334a0a69f5f1c54636cf743914f0b34fb2e7849b55ee7c1faddd0e06b4dfd.
//
// Solidity: event UpdatePendingManager(address manager)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdatePendingManager(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdatePendingManager) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdatePendingManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdatePendingManager)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdatePendingManager", log); err != nil {
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

// ParseUpdatePendingManager is a log parse operation binding the contract event 0x4d3334a0a69f5f1c54636cf743914f0b34fb2e7849b55ee7c1faddd0e06b4dfd.
//
// Solidity: event UpdatePendingManager(address manager)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdatePendingManager(log types.Log) (*AlphaProVaultUpdatePendingManager, error) {
	event := new(AlphaProVaultUpdatePendingManager)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdatePendingManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdatePeriodIterator is returned from FilterUpdatePeriod and is used to iterate over the raw logs and unpacked data for UpdatePeriod events raised by the AlphaProVault contract.
type AlphaProVaultUpdatePeriodIterator struct {
	Event *AlphaProVaultUpdatePeriod // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdatePeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdatePeriod)
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
		it.Event = new(AlphaProVaultUpdatePeriod)
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
func (it *AlphaProVaultUpdatePeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdatePeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdatePeriod represents a UpdatePeriod event raised by the AlphaProVault contract.
type AlphaProVaultUpdatePeriod struct {
	Period uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdatePeriod is a free log retrieval operation binding the contract event 0xa55d4ed589bc280e365d3ff1dc7fe4f59dff2d75e22716e29d6a9158fe059846.
//
// Solidity: event UpdatePeriod(uint32 period)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdatePeriod(opts *bind.FilterOpts) (*AlphaProVaultUpdatePeriodIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdatePeriod")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdatePeriodIterator{contract: _AlphaProVault.contract, event: "UpdatePeriod", logs: logs, sub: sub}, nil
}

// WatchUpdatePeriod is a free log subscription operation binding the contract event 0xa55d4ed589bc280e365d3ff1dc7fe4f59dff2d75e22716e29d6a9158fe059846.
//
// Solidity: event UpdatePeriod(uint32 period)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdatePeriod(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdatePeriod) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdatePeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdatePeriod)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdatePeriod", log); err != nil {
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

// ParseUpdatePeriod is a log parse operation binding the contract event 0xa55d4ed589bc280e365d3ff1dc7fe4f59dff2d75e22716e29d6a9158fe059846.
//
// Solidity: event UpdatePeriod(uint32 period)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdatePeriod(log types.Log) (*AlphaProVaultUpdatePeriod, error) {
	event := new(AlphaProVaultUpdatePeriod)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdatePeriod", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateProtocolFeeIterator is returned from FilterUpdateProtocolFee and is used to iterate over the raw logs and unpacked data for UpdateProtocolFee events raised by the AlphaProVault contract.
type AlphaProVaultUpdateProtocolFeeIterator struct {
	Event *AlphaProVaultUpdateProtocolFee // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateProtocolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateProtocolFee)
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
		it.Event = new(AlphaProVaultUpdateProtocolFee)
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
func (it *AlphaProVaultUpdateProtocolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateProtocolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateProtocolFee represents a UpdateProtocolFee event raised by the AlphaProVault contract.
type AlphaProVaultUpdateProtocolFee struct {
	ProtocolFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateProtocolFee is a free log retrieval operation binding the contract event 0x24a0d123bf9f15cb6bd5c3bc9b5cfe044075db4114c7a2f2a83a2409de73ac92.
//
// Solidity: event UpdateProtocolFee(uint24 protocolFee)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateProtocolFee(opts *bind.FilterOpts) (*AlphaProVaultUpdateProtocolFeeIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateProtocolFee")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateProtocolFeeIterator{contract: _AlphaProVault.contract, event: "UpdateProtocolFee", logs: logs, sub: sub}, nil
}

// WatchUpdateProtocolFee is a free log subscription operation binding the contract event 0x24a0d123bf9f15cb6bd5c3bc9b5cfe044075db4114c7a2f2a83a2409de73ac92.
//
// Solidity: event UpdateProtocolFee(uint24 protocolFee)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateProtocolFee(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateProtocolFee) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateProtocolFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateProtocolFee)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateProtocolFee", log); err != nil {
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

// ParseUpdateProtocolFee is a log parse operation binding the contract event 0x24a0d123bf9f15cb6bd5c3bc9b5cfe044075db4114c7a2f2a83a2409de73ac92.
//
// Solidity: event UpdateProtocolFee(uint24 protocolFee)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateProtocolFee(log types.Log) (*AlphaProVaultUpdateProtocolFee, error) {
	event := new(AlphaProVaultUpdateProtocolFee)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateProtocolFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateRebalanceDelegateIterator is returned from FilterUpdateRebalanceDelegate and is used to iterate over the raw logs and unpacked data for UpdateRebalanceDelegate events raised by the AlphaProVault contract.
type AlphaProVaultUpdateRebalanceDelegateIterator struct {
	Event *AlphaProVaultUpdateRebalanceDelegate // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateRebalanceDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateRebalanceDelegate)
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
		it.Event = new(AlphaProVaultUpdateRebalanceDelegate)
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
func (it *AlphaProVaultUpdateRebalanceDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateRebalanceDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateRebalanceDelegate represents a UpdateRebalanceDelegate event raised by the AlphaProVault contract.
type AlphaProVaultUpdateRebalanceDelegate struct {
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateRebalanceDelegate is a free log retrieval operation binding the contract event 0x02d71f3b54d658df30617ce7b33fa5f9835dd21a1da2a6dce6368dc9e5a40a97.
//
// Solidity: event UpdateRebalanceDelegate(address delegate)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateRebalanceDelegate(opts *bind.FilterOpts) (*AlphaProVaultUpdateRebalanceDelegateIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateRebalanceDelegate")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateRebalanceDelegateIterator{contract: _AlphaProVault.contract, event: "UpdateRebalanceDelegate", logs: logs, sub: sub}, nil
}

// WatchUpdateRebalanceDelegate is a free log subscription operation binding the contract event 0x02d71f3b54d658df30617ce7b33fa5f9835dd21a1da2a6dce6368dc9e5a40a97.
//
// Solidity: event UpdateRebalanceDelegate(address delegate)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateRebalanceDelegate(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateRebalanceDelegate) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateRebalanceDelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateRebalanceDelegate)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateRebalanceDelegate", log); err != nil {
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

// ParseUpdateRebalanceDelegate is a log parse operation binding the contract event 0x02d71f3b54d658df30617ce7b33fa5f9835dd21a1da2a6dce6368dc9e5a40a97.
//
// Solidity: event UpdateRebalanceDelegate(address delegate)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateRebalanceDelegate(log types.Log) (*AlphaProVaultUpdateRebalanceDelegate, error) {
	event := new(AlphaProVaultUpdateRebalanceDelegate)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateRebalanceDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateTwapDurationIterator is returned from FilterUpdateTwapDuration and is used to iterate over the raw logs and unpacked data for UpdateTwapDuration events raised by the AlphaProVault contract.
type AlphaProVaultUpdateTwapDurationIterator struct {
	Event *AlphaProVaultUpdateTwapDuration // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateTwapDurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateTwapDuration)
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
		it.Event = new(AlphaProVaultUpdateTwapDuration)
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
func (it *AlphaProVaultUpdateTwapDurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateTwapDurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateTwapDuration represents a UpdateTwapDuration event raised by the AlphaProVault contract.
type AlphaProVaultUpdateTwapDuration struct {
	TwapDuration uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateTwapDuration is a free log retrieval operation binding the contract event 0x2402faf0100aca3dc010189fd1fb0e310ab05da61ce9a65457e0fe7018e920d9.
//
// Solidity: event UpdateTwapDuration(uint32 twapDuration)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateTwapDuration(opts *bind.FilterOpts) (*AlphaProVaultUpdateTwapDurationIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateTwapDuration")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateTwapDurationIterator{contract: _AlphaProVault.contract, event: "UpdateTwapDuration", logs: logs, sub: sub}, nil
}

// WatchUpdateTwapDuration is a free log subscription operation binding the contract event 0x2402faf0100aca3dc010189fd1fb0e310ab05da61ce9a65457e0fe7018e920d9.
//
// Solidity: event UpdateTwapDuration(uint32 twapDuration)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateTwapDuration(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateTwapDuration) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateTwapDuration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateTwapDuration)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateTwapDuration", log); err != nil {
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

// ParseUpdateTwapDuration is a log parse operation binding the contract event 0x2402faf0100aca3dc010189fd1fb0e310ab05da61ce9a65457e0fe7018e920d9.
//
// Solidity: event UpdateTwapDuration(uint32 twapDuration)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateTwapDuration(log types.Log) (*AlphaProVaultUpdateTwapDuration, error) {
	event := new(AlphaProVaultUpdateTwapDuration)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateTwapDuration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateWideRangeWeightIterator is returned from FilterUpdateWideRangeWeight and is used to iterate over the raw logs and unpacked data for UpdateWideRangeWeight events raised by the AlphaProVault contract.
type AlphaProVaultUpdateWideRangeWeightIterator struct {
	Event *AlphaProVaultUpdateWideRangeWeight // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateWideRangeWeightIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateWideRangeWeight)
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
		it.Event = new(AlphaProVaultUpdateWideRangeWeight)
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
func (it *AlphaProVaultUpdateWideRangeWeightIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateWideRangeWeightIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateWideRangeWeight represents a UpdateWideRangeWeight event raised by the AlphaProVault contract.
type AlphaProVaultUpdateWideRangeWeight struct {
	Weight *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateWideRangeWeight is a free log retrieval operation binding the contract event 0x0047af854b92d7c4ca4913ffd329caf1c2e2b5e262b29139ae6cb247f18b886c.
//
// Solidity: event UpdateWideRangeWeight(uint24 weight)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateWideRangeWeight(opts *bind.FilterOpts) (*AlphaProVaultUpdateWideRangeWeightIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateWideRangeWeight")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateWideRangeWeightIterator{contract: _AlphaProVault.contract, event: "UpdateWideRangeWeight", logs: logs, sub: sub}, nil
}

// WatchUpdateWideRangeWeight is a free log subscription operation binding the contract event 0x0047af854b92d7c4ca4913ffd329caf1c2e2b5e262b29139ae6cb247f18b886c.
//
// Solidity: event UpdateWideRangeWeight(uint24 weight)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateWideRangeWeight(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateWideRangeWeight) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateWideRangeWeight")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateWideRangeWeight)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateWideRangeWeight", log); err != nil {
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

// ParseUpdateWideRangeWeight is a log parse operation binding the contract event 0x0047af854b92d7c4ca4913ffd329caf1c2e2b5e262b29139ae6cb247f18b886c.
//
// Solidity: event UpdateWideRangeWeight(uint24 weight)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateWideRangeWeight(log types.Log) (*AlphaProVaultUpdateWideRangeWeight, error) {
	event := new(AlphaProVaultUpdateWideRangeWeight)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateWideRangeWeight", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultUpdateWideThresholdIterator is returned from FilterUpdateWideThreshold and is used to iterate over the raw logs and unpacked data for UpdateWideThreshold events raised by the AlphaProVault contract.
type AlphaProVaultUpdateWideThresholdIterator struct {
	Event *AlphaProVaultUpdateWideThreshold // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultUpdateWideThresholdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultUpdateWideThreshold)
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
		it.Event = new(AlphaProVaultUpdateWideThreshold)
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
func (it *AlphaProVaultUpdateWideThresholdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultUpdateWideThresholdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultUpdateWideThreshold represents a UpdateWideThreshold event raised by the AlphaProVault contract.
type AlphaProVaultUpdateWideThreshold struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateWideThreshold is a free log retrieval operation binding the contract event 0x6be14bdbb0283c52b91dff0090b8bda078b11ce7f87486a7c14ae7595fd0f938.
//
// Solidity: event UpdateWideThreshold(int24 threshold)
func (_AlphaProVault *AlphaProVaultFilterer) FilterUpdateWideThreshold(opts *bind.FilterOpts) (*AlphaProVaultUpdateWideThresholdIterator, error) {

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "UpdateWideThreshold")
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultUpdateWideThresholdIterator{contract: _AlphaProVault.contract, event: "UpdateWideThreshold", logs: logs, sub: sub}, nil
}

// WatchUpdateWideThreshold is a free log subscription operation binding the contract event 0x6be14bdbb0283c52b91dff0090b8bda078b11ce7f87486a7c14ae7595fd0f938.
//
// Solidity: event UpdateWideThreshold(int24 threshold)
func (_AlphaProVault *AlphaProVaultFilterer) WatchUpdateWideThreshold(opts *bind.WatchOpts, sink chan<- *AlphaProVaultUpdateWideThreshold) (event.Subscription, error) {

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "UpdateWideThreshold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultUpdateWideThreshold)
				if err := _AlphaProVault.contract.UnpackLog(event, "UpdateWideThreshold", log); err != nil {
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

// ParseUpdateWideThreshold is a log parse operation binding the contract event 0x6be14bdbb0283c52b91dff0090b8bda078b11ce7f87486a7c14ae7595fd0f938.
//
// Solidity: event UpdateWideThreshold(int24 threshold)
func (_AlphaProVault *AlphaProVaultFilterer) ParseUpdateWideThreshold(log types.Log) (*AlphaProVaultUpdateWideThreshold, error) {
	event := new(AlphaProVaultUpdateWideThreshold)
	if err := _AlphaProVault.contract.UnpackLog(event, "UpdateWideThreshold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlphaProVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AlphaProVault contract.
type AlphaProVaultWithdrawIterator struct {
	Event *AlphaProVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *AlphaProVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlphaProVaultWithdraw)
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
		it.Event = new(AlphaProVaultWithdraw)
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
func (it *AlphaProVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlphaProVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlphaProVaultWithdraw represents a Withdraw event raised by the AlphaProVault contract.
type AlphaProVaultWithdraw struct {
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
func (_AlphaProVault *AlphaProVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*AlphaProVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AlphaProVault.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AlphaProVaultWithdrawIterator{contract: _AlphaProVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_AlphaProVault *AlphaProVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AlphaProVaultWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AlphaProVault.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlphaProVaultWithdraw)
				if err := _AlphaProVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_AlphaProVault *AlphaProVaultFilterer) ParseWithdraw(log types.Log) (*AlphaProVaultWithdraw, error) {
	event := new(AlphaProVaultWithdraw)
	if err := _AlphaProVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
