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

// VaultStrategyParams is an auto generated low-level Go binding around an user-defined struct.
type VaultStrategyParams struct {
	IsActivated bool
	LastHarvest *big.Int
	CurrentDebt *big.Int
	MaxDebt     *big.Int
}

// WeberaVaultMetaData contains all meta data concerning the WeberaVault contract.
var WeberaVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDebt\",\"type\":\"uint256\"}],\"name\":\"DebtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRefunds\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"oldStrategyQueue\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newStrategyQueue\",\"type\":\"address[]\"}],\"name\":\"WithdrawQueueUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BPS_EXTENDED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STRATEGIES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountant\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDebt\",\"type\":\"uint256\"}],\"name\":\"addStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategy_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss_\",\"type\":\"uint256\"}],\"name\":\"emergencyStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxLoss_\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fullProfitUnlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHarvester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStrategies\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawQueue\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategy_\",\"type\":\"address\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"vaultTokenName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultTokenSymbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"vaultFactory_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profitMaxUnlockTime_\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpened\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastProfitHarvest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minTotalIdleAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitMaxUnlockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitUnlockingRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"revokeStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"harvester_\",\"type\":\"address\"}],\"name\":\"setHarvester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profitMaxUnlockTime_\",\"type\":\"uint256\"}],\"name\":\"setProfitMaxUnlockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vaultRouter_\",\"type\":\"address\"}],\"name\":\"setVaultRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"queue\",\"type\":\"address[]\"}],\"name\":\"setWithdrawQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"strategies\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategiesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategy_\",\"type\":\"address\"}],\"name\":\"strategyParams\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isActivated\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"lastHarvest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDebt\",\"type\":\"uint256\"}],\"internalType\":\"structVault.StrategyParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalIdleAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOutstandingDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockedShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategy_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetDebt_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss_\",\"type\":\"uint256\"}],\"name\":\"updateDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategy_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDebt_\",\"type\":\"uint256\"}],\"name\":\"updateStrategyMaxDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WeberaVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use WeberaVaultMetaData.ABI instead.
var WeberaVaultABI = WeberaVaultMetaData.ABI

// WeberaVault is an auto generated Go binding around an Ethereum contract.
type WeberaVault struct {
	WeberaVaultCaller     // Read-only binding to the contract
	WeberaVaultTransactor // Write-only binding to the contract
	WeberaVaultFilterer   // Log filterer for contract events
}

// WeberaVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type WeberaVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WeberaVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WeberaVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WeberaVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WeberaVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WeberaVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WeberaVaultSession struct {
	Contract     *WeberaVault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WeberaVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WeberaVaultCallerSession struct {
	Contract *WeberaVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WeberaVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WeberaVaultTransactorSession struct {
	Contract     *WeberaVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WeberaVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type WeberaVaultRaw struct {
	Contract *WeberaVault // Generic contract binding to access the raw methods on
}

// WeberaVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WeberaVaultCallerRaw struct {
	Contract *WeberaVaultCaller // Generic read-only contract binding to access the raw methods on
}

// WeberaVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WeberaVaultTransactorRaw struct {
	Contract *WeberaVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWeberaVault creates a new instance of WeberaVault, bound to a specific deployed contract.
func NewWeberaVault(address common.Address, backend bind.ContractBackend) (*WeberaVault, error) {
	contract, err := bindWeberaVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WeberaVault{WeberaVaultCaller: WeberaVaultCaller{contract: contract}, WeberaVaultTransactor: WeberaVaultTransactor{contract: contract}, WeberaVaultFilterer: WeberaVaultFilterer{contract: contract}}, nil
}

// NewWeberaVaultCaller creates a new read-only instance of WeberaVault, bound to a specific deployed contract.
func NewWeberaVaultCaller(address common.Address, caller bind.ContractCaller) (*WeberaVaultCaller, error) {
	contract, err := bindWeberaVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultCaller{contract: contract}, nil
}

// NewWeberaVaultTransactor creates a new write-only instance of WeberaVault, bound to a specific deployed contract.
func NewWeberaVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*WeberaVaultTransactor, error) {
	contract, err := bindWeberaVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultTransactor{contract: contract}, nil
}

// NewWeberaVaultFilterer creates a new log filterer instance of WeberaVault, bound to a specific deployed contract.
func NewWeberaVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*WeberaVaultFilterer, error) {
	contract, err := bindWeberaVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultFilterer{contract: contract}, nil
}

// bindWeberaVault binds a generic wrapper to an already deployed contract.
func bindWeberaVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WeberaVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WeberaVault *WeberaVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WeberaVault.Contract.WeberaVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WeberaVault *WeberaVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WeberaVault.Contract.WeberaVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WeberaVault *WeberaVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WeberaVault.Contract.WeberaVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WeberaVault *WeberaVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WeberaVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WeberaVault *WeberaVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WeberaVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WeberaVault *WeberaVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WeberaVault.Contract.contract.Transact(opts, method, params...)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) MAXBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "MAX_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) MAXBPS() (*big.Int, error) {
	return _WeberaVault.Contract.MAXBPS(&_WeberaVault.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) MAXBPS() (*big.Int, error) {
	return _WeberaVault.Contract.MAXBPS(&_WeberaVault.CallOpts)
}

// MAXBPSEXTENDED is a free data retrieval call binding the contract method 0x36fba084.
//
// Solidity: function MAX_BPS_EXTENDED() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) MAXBPSEXTENDED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "MAX_BPS_EXTENDED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPSEXTENDED is a free data retrieval call binding the contract method 0x36fba084.
//
// Solidity: function MAX_BPS_EXTENDED() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) MAXBPSEXTENDED() (*big.Int, error) {
	return _WeberaVault.Contract.MAXBPSEXTENDED(&_WeberaVault.CallOpts)
}

// MAXBPSEXTENDED is a free data retrieval call binding the contract method 0x36fba084.
//
// Solidity: function MAX_BPS_EXTENDED() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) MAXBPSEXTENDED() (*big.Int, error) {
	return _WeberaVault.Contract.MAXBPSEXTENDED(&_WeberaVault.CallOpts)
}

// MAXSTRATEGIES is a free data retrieval call binding the contract method 0x767f06ae.
//
// Solidity: function MAX_STRATEGIES() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) MAXSTRATEGIES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "MAX_STRATEGIES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTRATEGIES is a free data retrieval call binding the contract method 0x767f06ae.
//
// Solidity: function MAX_STRATEGIES() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) MAXSTRATEGIES() (*big.Int, error) {
	return _WeberaVault.Contract.MAXSTRATEGIES(&_WeberaVault.CallOpts)
}

// MAXSTRATEGIES is a free data retrieval call binding the contract method 0x767f06ae.
//
// Solidity: function MAX_STRATEGIES() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) MAXSTRATEGIES() (*big.Int, error) {
	return _WeberaVault.Contract.MAXSTRATEGIES(&_WeberaVault.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_WeberaVault *WeberaVaultCaller) Accountant(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "accountant")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_WeberaVault *WeberaVaultSession) Accountant() (common.Address, error) {
	return _WeberaVault.Contract.Accountant(&_WeberaVault.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_WeberaVault *WeberaVaultCallerSession) Accountant() (common.Address, error) {
	return _WeberaVault.Contract.Accountant(&_WeberaVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.Allowance(&_WeberaVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.Allowance(&_WeberaVault.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_WeberaVault *WeberaVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_WeberaVault *WeberaVaultSession) Asset() (common.Address, error) {
	return _WeberaVault.Contract.Asset(&_WeberaVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_WeberaVault *WeberaVaultCallerSession) Asset() (common.Address, error) {
	return _WeberaVault.Contract.Asset(&_WeberaVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.BalanceOf(&_WeberaVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.BalanceOf(&_WeberaVault.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.ConvertToAssets(&_WeberaVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.ConvertToAssets(&_WeberaVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.ConvertToShares(&_WeberaVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.ConvertToShares(&_WeberaVault.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WeberaVault *WeberaVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WeberaVault *WeberaVaultSession) Decimals() (uint8, error) {
	return _WeberaVault.Contract.Decimals(&_WeberaVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WeberaVault *WeberaVaultCallerSession) Decimals() (uint8, error) {
	return _WeberaVault.Contract.Decimals(&_WeberaVault.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_WeberaVault *WeberaVaultCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_WeberaVault *WeberaVaultSession) Factory() (common.Address, error) {
	return _WeberaVault.Contract.Factory(&_WeberaVault.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_WeberaVault *WeberaVaultCallerSession) Factory() (common.Address, error) {
	return _WeberaVault.Contract.Factory(&_WeberaVault.CallOpts)
}

// FullProfitUnlockTime is a free data retrieval call binding the contract method 0x19ec039b.
//
// Solidity: function fullProfitUnlockTime() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) FullProfitUnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "fullProfitUnlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FullProfitUnlockTime is a free data retrieval call binding the contract method 0x19ec039b.
//
// Solidity: function fullProfitUnlockTime() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) FullProfitUnlockTime() (*big.Int, error) {
	return _WeberaVault.Contract.FullProfitUnlockTime(&_WeberaVault.CallOpts)
}

// FullProfitUnlockTime is a free data retrieval call binding the contract method 0x19ec039b.
//
// Solidity: function fullProfitUnlockTime() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) FullProfitUnlockTime() (*big.Int, error) {
	return _WeberaVault.Contract.FullProfitUnlockTime(&_WeberaVault.CallOpts)
}

// GetHarvester is a free data retrieval call binding the contract method 0x5426cb2b.
//
// Solidity: function getHarvester() view returns(address)
func (_WeberaVault *WeberaVaultCaller) GetHarvester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "getHarvester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHarvester is a free data retrieval call binding the contract method 0x5426cb2b.
//
// Solidity: function getHarvester() view returns(address)
func (_WeberaVault *WeberaVaultSession) GetHarvester() (common.Address, error) {
	return _WeberaVault.Contract.GetHarvester(&_WeberaVault.CallOpts)
}

// GetHarvester is a free data retrieval call binding the contract method 0x5426cb2b.
//
// Solidity: function getHarvester() view returns(address)
func (_WeberaVault *WeberaVaultCallerSession) GetHarvester() (common.Address, error) {
	return _WeberaVault.Contract.GetHarvester(&_WeberaVault.CallOpts)
}

// GetStrategies is a free data retrieval call binding the contract method 0xb49a60bb.
//
// Solidity: function getStrategies() view returns(address[])
func (_WeberaVault *WeberaVaultCaller) GetStrategies(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "getStrategies")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStrategies is a free data retrieval call binding the contract method 0xb49a60bb.
//
// Solidity: function getStrategies() view returns(address[])
func (_WeberaVault *WeberaVaultSession) GetStrategies() ([]common.Address, error) {
	return _WeberaVault.Contract.GetStrategies(&_WeberaVault.CallOpts)
}

// GetStrategies is a free data retrieval call binding the contract method 0xb49a60bb.
//
// Solidity: function getStrategies() view returns(address[])
func (_WeberaVault *WeberaVaultCallerSession) GetStrategies() ([]common.Address, error) {
	return _WeberaVault.Contract.GetStrategies(&_WeberaVault.CallOpts)
}

// GetWithdrawQueue is a free data retrieval call binding the contract method 0x11183052.
//
// Solidity: function getWithdrawQueue() view returns(address[])
func (_WeberaVault *WeberaVaultCaller) GetWithdrawQueue(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "getWithdrawQueue")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWithdrawQueue is a free data retrieval call binding the contract method 0x11183052.
//
// Solidity: function getWithdrawQueue() view returns(address[])
func (_WeberaVault *WeberaVaultSession) GetWithdrawQueue() ([]common.Address, error) {
	return _WeberaVault.Contract.GetWithdrawQueue(&_WeberaVault.CallOpts)
}

// GetWithdrawQueue is a free data retrieval call binding the contract method 0x11183052.
//
// Solidity: function getWithdrawQueue() view returns(address[])
func (_WeberaVault *WeberaVaultCallerSession) GetWithdrawQueue() ([]common.Address, error) {
	return _WeberaVault.Contract.GetWithdrawQueue(&_WeberaVault.CallOpts)
}

// Harvester is a free data retrieval call binding the contract method 0x4bdaeac1.
//
// Solidity: function harvester() view returns(address)
func (_WeberaVault *WeberaVaultCaller) Harvester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "harvester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Harvester is a free data retrieval call binding the contract method 0x4bdaeac1.
//
// Solidity: function harvester() view returns(address)
func (_WeberaVault *WeberaVaultSession) Harvester() (common.Address, error) {
	return _WeberaVault.Contract.Harvester(&_WeberaVault.CallOpts)
}

// Harvester is a free data retrieval call binding the contract method 0x4bdaeac1.
//
// Solidity: function harvester() view returns(address)
func (_WeberaVault *WeberaVaultCallerSession) Harvester() (common.Address, error) {
	return _WeberaVault.Contract.Harvester(&_WeberaVault.CallOpts)
}

// IsOpened is a free data retrieval call binding the contract method 0x692aa97e.
//
// Solidity: function isOpened() view returns(bool)
func (_WeberaVault *WeberaVaultCaller) IsOpened(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "isOpened")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpened is a free data retrieval call binding the contract method 0x692aa97e.
//
// Solidity: function isOpened() view returns(bool)
func (_WeberaVault *WeberaVaultSession) IsOpened() (bool, error) {
	return _WeberaVault.Contract.IsOpened(&_WeberaVault.CallOpts)
}

// IsOpened is a free data retrieval call binding the contract method 0x692aa97e.
//
// Solidity: function isOpened() view returns(bool)
func (_WeberaVault *WeberaVaultCallerSession) IsOpened() (bool, error) {
	return _WeberaVault.Contract.IsOpened(&_WeberaVault.CallOpts)
}

// LastProfitHarvest is a free data retrieval call binding the contract method 0x50a5cb87.
//
// Solidity: function lastProfitHarvest() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) LastProfitHarvest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "lastProfitHarvest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProfitHarvest is a free data retrieval call binding the contract method 0x50a5cb87.
//
// Solidity: function lastProfitHarvest() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) LastProfitHarvest() (*big.Int, error) {
	return _WeberaVault.Contract.LastProfitHarvest(&_WeberaVault.CallOpts)
}

// LastProfitHarvest is a free data retrieval call binding the contract method 0x50a5cb87.
//
// Solidity: function lastProfitHarvest() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) LastProfitHarvest() (*big.Int, error) {
	return _WeberaVault.Contract.LastProfitHarvest(&_WeberaVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.MaxDeposit(&_WeberaVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.MaxDeposit(&_WeberaVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.MaxMint(&_WeberaVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.MaxMint(&_WeberaVault.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.MaxRedeem(&_WeberaVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.MaxRedeem(&_WeberaVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.MaxWithdraw(&_WeberaVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _WeberaVault.Contract.MaxWithdraw(&_WeberaVault.CallOpts, owner)
}

// MinTotalIdleAssets is a free data retrieval call binding the contract method 0x78824c06.
//
// Solidity: function minTotalIdleAssets() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) MinTotalIdleAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "minTotalIdleAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTotalIdleAssets is a free data retrieval call binding the contract method 0x78824c06.
//
// Solidity: function minTotalIdleAssets() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) MinTotalIdleAssets() (*big.Int, error) {
	return _WeberaVault.Contract.MinTotalIdleAssets(&_WeberaVault.CallOpts)
}

// MinTotalIdleAssets is a free data retrieval call binding the contract method 0x78824c06.
//
// Solidity: function minTotalIdleAssets() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) MinTotalIdleAssets() (*big.Int, error) {
	return _WeberaVault.Contract.MinTotalIdleAssets(&_WeberaVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WeberaVault *WeberaVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WeberaVault *WeberaVaultSession) Name() (string, error) {
	return _WeberaVault.Contract.Name(&_WeberaVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WeberaVault *WeberaVaultCallerSession) Name() (string, error) {
	return _WeberaVault.Contract.Name(&_WeberaVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WeberaVault *WeberaVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WeberaVault *WeberaVaultSession) Owner() (common.Address, error) {
	return _WeberaVault.Contract.Owner(&_WeberaVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WeberaVault *WeberaVaultCallerSession) Owner() (common.Address, error) {
	return _WeberaVault.Contract.Owner(&_WeberaVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.PreviewDeposit(&_WeberaVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.PreviewDeposit(&_WeberaVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.PreviewMint(&_WeberaVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.PreviewMint(&_WeberaVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.PreviewRedeem(&_WeberaVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.PreviewRedeem(&_WeberaVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_WeberaVault *WeberaVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.PreviewWithdraw(&_WeberaVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _WeberaVault.Contract.PreviewWithdraw(&_WeberaVault.CallOpts, assets)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) ProfitMaxUnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "profitMaxUnlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) ProfitMaxUnlockTime() (*big.Int, error) {
	return _WeberaVault.Contract.ProfitMaxUnlockTime(&_WeberaVault.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) ProfitMaxUnlockTime() (*big.Int, error) {
	return _WeberaVault.Contract.ProfitMaxUnlockTime(&_WeberaVault.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) ProfitUnlockingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "profitUnlockingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) ProfitUnlockingRate() (*big.Int, error) {
	return _WeberaVault.Contract.ProfitUnlockingRate(&_WeberaVault.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) ProfitUnlockingRate() (*big.Int, error) {
	return _WeberaVault.Contract.ProfitUnlockingRate(&_WeberaVault.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) ProtocolFee() (*big.Int, error) {
	return _WeberaVault.Contract.ProtocolFee(&_WeberaVault.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) ProtocolFee() (*big.Int, error) {
	return _WeberaVault.Contract.ProtocolFee(&_WeberaVault.CallOpts)
}

// Strategies is a free data retrieval call binding the contract method 0xd574ea3d.
//
// Solidity: function strategies(uint256 ) view returns(address)
func (_WeberaVault *WeberaVaultCaller) Strategies(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "strategies", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Strategies is a free data retrieval call binding the contract method 0xd574ea3d.
//
// Solidity: function strategies(uint256 ) view returns(address)
func (_WeberaVault *WeberaVaultSession) Strategies(arg0 *big.Int) (common.Address, error) {
	return _WeberaVault.Contract.Strategies(&_WeberaVault.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0xd574ea3d.
//
// Solidity: function strategies(uint256 ) view returns(address)
func (_WeberaVault *WeberaVaultCallerSession) Strategies(arg0 *big.Int) (common.Address, error) {
	return _WeberaVault.Contract.Strategies(&_WeberaVault.CallOpts, arg0)
}

// StrategiesLength is a free data retrieval call binding the contract method 0x3f74bb89.
//
// Solidity: function strategiesLength() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) StrategiesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "strategiesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrategiesLength is a free data retrieval call binding the contract method 0x3f74bb89.
//
// Solidity: function strategiesLength() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) StrategiesLength() (*big.Int, error) {
	return _WeberaVault.Contract.StrategiesLength(&_WeberaVault.CallOpts)
}

// StrategiesLength is a free data retrieval call binding the contract method 0x3f74bb89.
//
// Solidity: function strategiesLength() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) StrategiesLength() (*big.Int, error) {
	return _WeberaVault.Contract.StrategiesLength(&_WeberaVault.CallOpts)
}

// StrategyParams is a free data retrieval call binding the contract method 0xc4c2d9b9.
//
// Solidity: function strategyParams(address strategy_) view returns((bool,uint256,uint256,uint256))
func (_WeberaVault *WeberaVaultCaller) StrategyParams(opts *bind.CallOpts, strategy_ common.Address) (VaultStrategyParams, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "strategyParams", strategy_)

	if err != nil {
		return *new(VaultStrategyParams), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultStrategyParams)).(*VaultStrategyParams)

	return out0, err

}

// StrategyParams is a free data retrieval call binding the contract method 0xc4c2d9b9.
//
// Solidity: function strategyParams(address strategy_) view returns((bool,uint256,uint256,uint256))
func (_WeberaVault *WeberaVaultSession) StrategyParams(strategy_ common.Address) (VaultStrategyParams, error) {
	return _WeberaVault.Contract.StrategyParams(&_WeberaVault.CallOpts, strategy_)
}

// StrategyParams is a free data retrieval call binding the contract method 0xc4c2d9b9.
//
// Solidity: function strategyParams(address strategy_) view returns((bool,uint256,uint256,uint256))
func (_WeberaVault *WeberaVaultCallerSession) StrategyParams(strategy_ common.Address) (VaultStrategyParams, error) {
	return _WeberaVault.Contract.StrategyParams(&_WeberaVault.CallOpts, strategy_)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WeberaVault *WeberaVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WeberaVault *WeberaVaultSession) Symbol() (string, error) {
	return _WeberaVault.Contract.Symbol(&_WeberaVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WeberaVault *WeberaVaultCallerSession) Symbol() (string, error) {
	return _WeberaVault.Contract.Symbol(&_WeberaVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) TotalAssets() (*big.Int, error) {
	return _WeberaVault.Contract.TotalAssets(&_WeberaVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _WeberaVault.Contract.TotalAssets(&_WeberaVault.CallOpts)
}

// TotalIdleAssets is a free data retrieval call binding the contract method 0x35a22a7d.
//
// Solidity: function totalIdleAssets() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) TotalIdleAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "totalIdleAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIdleAssets is a free data retrieval call binding the contract method 0x35a22a7d.
//
// Solidity: function totalIdleAssets() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) TotalIdleAssets() (*big.Int, error) {
	return _WeberaVault.Contract.TotalIdleAssets(&_WeberaVault.CallOpts)
}

// TotalIdleAssets is a free data retrieval call binding the contract method 0x35a22a7d.
//
// Solidity: function totalIdleAssets() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) TotalIdleAssets() (*big.Int, error) {
	return _WeberaVault.Contract.TotalIdleAssets(&_WeberaVault.CallOpts)
}

// TotalOutstandingDebt is a free data retrieval call binding the contract method 0x859203e3.
//
// Solidity: function totalOutstandingDebt() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) TotalOutstandingDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "totalOutstandingDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOutstandingDebt is a free data retrieval call binding the contract method 0x859203e3.
//
// Solidity: function totalOutstandingDebt() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) TotalOutstandingDebt() (*big.Int, error) {
	return _WeberaVault.Contract.TotalOutstandingDebt(&_WeberaVault.CallOpts)
}

// TotalOutstandingDebt is a free data retrieval call binding the contract method 0x859203e3.
//
// Solidity: function totalOutstandingDebt() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) TotalOutstandingDebt() (*big.Int, error) {
	return _WeberaVault.Contract.TotalOutstandingDebt(&_WeberaVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) TotalSupply() (*big.Int, error) {
	return _WeberaVault.Contract.TotalSupply(&_WeberaVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _WeberaVault.Contract.TotalSupply(&_WeberaVault.CallOpts)
}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) UnlockedShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "unlockedShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) UnlockedShares() (*big.Int, error) {
	return _WeberaVault.Contract.UnlockedShares(&_WeberaVault.CallOpts)
}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) UnlockedShares() (*big.Int, error) {
	return _WeberaVault.Contract.UnlockedShares(&_WeberaVault.CallOpts)
}

// VaultRouter is a free data retrieval call binding the contract method 0xd37c441c.
//
// Solidity: function vaultRouter() view returns(address)
func (_WeberaVault *WeberaVaultCaller) VaultRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "vaultRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultRouter is a free data retrieval call binding the contract method 0xd37c441c.
//
// Solidity: function vaultRouter() view returns(address)
func (_WeberaVault *WeberaVaultSession) VaultRouter() (common.Address, error) {
	return _WeberaVault.Contract.VaultRouter(&_WeberaVault.CallOpts)
}

// VaultRouter is a free data retrieval call binding the contract method 0xd37c441c.
//
// Solidity: function vaultRouter() view returns(address)
func (_WeberaVault *WeberaVaultCallerSession) VaultRouter() (common.Address, error) {
	return _WeberaVault.Contract.VaultRouter(&_WeberaVault.CallOpts)
}

// WithdrawLimit is a free data retrieval call binding the contract method 0xf848d541.
//
// Solidity: function withdrawLimit() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) WithdrawLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "withdrawLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawLimit is a free data retrieval call binding the contract method 0xf848d541.
//
// Solidity: function withdrawLimit() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) WithdrawLimit() (*big.Int, error) {
	return _WeberaVault.Contract.WithdrawLimit(&_WeberaVault.CallOpts)
}

// WithdrawLimit is a free data retrieval call binding the contract method 0xf848d541.
//
// Solidity: function withdrawLimit() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) WithdrawLimit() (*big.Int, error) {
	return _WeberaVault.Contract.WithdrawLimit(&_WeberaVault.CallOpts)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(address)
func (_WeberaVault *WeberaVaultCaller) WithdrawQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "withdrawQueue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(address)
func (_WeberaVault *WeberaVaultSession) WithdrawQueue(arg0 *big.Int) (common.Address, error) {
	return _WeberaVault.Contract.WithdrawQueue(&_WeberaVault.CallOpts, arg0)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(address)
func (_WeberaVault *WeberaVaultCallerSession) WithdrawQueue(arg0 *big.Int) (common.Address, error) {
	return _WeberaVault.Contract.WithdrawQueue(&_WeberaVault.CallOpts, arg0)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_WeberaVault *WeberaVaultCaller) WithdrawQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WeberaVault.contract.Call(opts, &out, "withdrawQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_WeberaVault *WeberaVaultSession) WithdrawQueueLength() (*big.Int, error) {
	return _WeberaVault.Contract.WithdrawQueueLength(&_WeberaVault.CallOpts)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_WeberaVault *WeberaVaultCallerSession) WithdrawQueueLength() (*big.Int, error) {
	return _WeberaVault.Contract.WithdrawQueueLength(&_WeberaVault.CallOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xc9411e22.
//
// Solidity: function addStrategy(address strategy, uint256 maxDebt) returns()
func (_WeberaVault *WeberaVaultTransactor) AddStrategy(opts *bind.TransactOpts, strategy common.Address, maxDebt *big.Int) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "addStrategy", strategy, maxDebt)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xc9411e22.
//
// Solidity: function addStrategy(address strategy, uint256 maxDebt) returns()
func (_WeberaVault *WeberaVaultSession) AddStrategy(strategy common.Address, maxDebt *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.AddStrategy(&_WeberaVault.TransactOpts, strategy, maxDebt)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xc9411e22.
//
// Solidity: function addStrategy(address strategy, uint256 maxDebt) returns()
func (_WeberaVault *WeberaVaultTransactorSession) AddStrategy(strategy common.Address, maxDebt *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.AddStrategy(&_WeberaVault.TransactOpts, strategy, maxDebt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WeberaVault *WeberaVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WeberaVault *WeberaVaultSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.Approve(&_WeberaVault.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_WeberaVault *WeberaVaultTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.Approve(&_WeberaVault.TransactOpts, spender, value)
}

// CloseVault is a paid mutator transaction binding the contract method 0x18976fa2.
//
// Solidity: function closeVault() returns()
func (_WeberaVault *WeberaVaultTransactor) CloseVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "closeVault")
}

// CloseVault is a paid mutator transaction binding the contract method 0x18976fa2.
//
// Solidity: function closeVault() returns()
func (_WeberaVault *WeberaVaultSession) CloseVault() (*types.Transaction, error) {
	return _WeberaVault.Contract.CloseVault(&_WeberaVault.TransactOpts)
}

// CloseVault is a paid mutator transaction binding the contract method 0x18976fa2.
//
// Solidity: function closeVault() returns()
func (_WeberaVault *WeberaVaultTransactorSession) CloseVault() (*types.Transaction, error) {
	return _WeberaVault.Contract.CloseVault(&_WeberaVault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_WeberaVault *WeberaVaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_WeberaVault *WeberaVaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.Deposit(&_WeberaVault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_WeberaVault *WeberaVaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.Deposit(&_WeberaVault.TransactOpts, assets, receiver)
}

// EmergencyStrategy is a paid mutator transaction binding the contract method 0xb20fb167.
//
// Solidity: function emergencyStrategy(address strategy_, uint256 amount_, uint256 maxLoss_) returns()
func (_WeberaVault *WeberaVaultTransactor) EmergencyStrategy(opts *bind.TransactOpts, strategy_ common.Address, amount_ *big.Int, maxLoss_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "emergencyStrategy", strategy_, amount_, maxLoss_)
}

// EmergencyStrategy is a paid mutator transaction binding the contract method 0xb20fb167.
//
// Solidity: function emergencyStrategy(address strategy_, uint256 amount_, uint256 maxLoss_) returns()
func (_WeberaVault *WeberaVaultSession) EmergencyStrategy(strategy_ common.Address, amount_ *big.Int, maxLoss_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.EmergencyStrategy(&_WeberaVault.TransactOpts, strategy_, amount_, maxLoss_)
}

// EmergencyStrategy is a paid mutator transaction binding the contract method 0xb20fb167.
//
// Solidity: function emergencyStrategy(address strategy_, uint256 amount_, uint256 maxLoss_) returns()
func (_WeberaVault *WeberaVaultTransactorSession) EmergencyStrategy(strategy_ common.Address, amount_ *big.Int, maxLoss_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.EmergencyStrategy(&_WeberaVault.TransactOpts, strategy_, amount_, maxLoss_)
}

// EmergencyVault is a paid mutator transaction binding the contract method 0xfe6f68ba.
//
// Solidity: function emergencyVault() returns()
func (_WeberaVault *WeberaVaultTransactor) EmergencyVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "emergencyVault")
}

// EmergencyVault is a paid mutator transaction binding the contract method 0xfe6f68ba.
//
// Solidity: function emergencyVault() returns()
func (_WeberaVault *WeberaVaultSession) EmergencyVault() (*types.Transaction, error) {
	return _WeberaVault.Contract.EmergencyVault(&_WeberaVault.TransactOpts)
}

// EmergencyVault is a paid mutator transaction binding the contract method 0xfe6f68ba.
//
// Solidity: function emergencyVault() returns()
func (_WeberaVault *WeberaVaultTransactorSession) EmergencyVault() (*types.Transaction, error) {
	return _WeberaVault.Contract.EmergencyVault(&_WeberaVault.TransactOpts)
}

// EmergencyWithdrawAll is a paid mutator transaction binding the contract method 0x1e93facb.
//
// Solidity: function emergencyWithdrawAll(uint256 maxLoss_) returns()
func (_WeberaVault *WeberaVaultTransactor) EmergencyWithdrawAll(opts *bind.TransactOpts, maxLoss_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "emergencyWithdrawAll", maxLoss_)
}

// EmergencyWithdrawAll is a paid mutator transaction binding the contract method 0x1e93facb.
//
// Solidity: function emergencyWithdrawAll(uint256 maxLoss_) returns()
func (_WeberaVault *WeberaVaultSession) EmergencyWithdrawAll(maxLoss_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.EmergencyWithdrawAll(&_WeberaVault.TransactOpts, maxLoss_)
}

// EmergencyWithdrawAll is a paid mutator transaction binding the contract method 0x1e93facb.
//
// Solidity: function emergencyWithdrawAll(uint256 maxLoss_) returns()
func (_WeberaVault *WeberaVaultTransactorSession) EmergencyWithdrawAll(maxLoss_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.EmergencyWithdrawAll(&_WeberaVault.TransactOpts, maxLoss_)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address strategy_) returns()
func (_WeberaVault *WeberaVaultTransactor) Harvest(opts *bind.TransactOpts, strategy_ common.Address) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "harvest", strategy_)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address strategy_) returns()
func (_WeberaVault *WeberaVaultSession) Harvest(strategy_ common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.Harvest(&_WeberaVault.TransactOpts, strategy_)
}

// Harvest is a paid mutator transaction binding the contract method 0x0e5c011e.
//
// Solidity: function harvest(address strategy_) returns()
func (_WeberaVault *WeberaVaultTransactorSession) Harvest(strategy_ common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.Harvest(&_WeberaVault.TransactOpts, strategy_)
}

// HarvestAll is a paid mutator transaction binding the contract method 0x8ed955b9.
//
// Solidity: function harvestAll() returns()
func (_WeberaVault *WeberaVaultTransactor) HarvestAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "harvestAll")
}

// HarvestAll is a paid mutator transaction binding the contract method 0x8ed955b9.
//
// Solidity: function harvestAll() returns()
func (_WeberaVault *WeberaVaultSession) HarvestAll() (*types.Transaction, error) {
	return _WeberaVault.Contract.HarvestAll(&_WeberaVault.TransactOpts)
}

// HarvestAll is a paid mutator transaction binding the contract method 0x8ed955b9.
//
// Solidity: function harvestAll() returns()
func (_WeberaVault *WeberaVaultTransactorSession) HarvestAll() (*types.Transaction, error) {
	return _WeberaVault.Contract.HarvestAll(&_WeberaVault.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa01e3f74.
//
// Solidity: function initialize(address owner_, address asset_, string vaultTokenName_, string vaultTokenSymbol_, address vaultFactory_, uint256 profitMaxUnlockTime_) returns()
func (_WeberaVault *WeberaVaultTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, asset_ common.Address, vaultTokenName_ string, vaultTokenSymbol_ string, vaultFactory_ common.Address, profitMaxUnlockTime_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "initialize", owner_, asset_, vaultTokenName_, vaultTokenSymbol_, vaultFactory_, profitMaxUnlockTime_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa01e3f74.
//
// Solidity: function initialize(address owner_, address asset_, string vaultTokenName_, string vaultTokenSymbol_, address vaultFactory_, uint256 profitMaxUnlockTime_) returns()
func (_WeberaVault *WeberaVaultSession) Initialize(owner_ common.Address, asset_ common.Address, vaultTokenName_ string, vaultTokenSymbol_ string, vaultFactory_ common.Address, profitMaxUnlockTime_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.Initialize(&_WeberaVault.TransactOpts, owner_, asset_, vaultTokenName_, vaultTokenSymbol_, vaultFactory_, profitMaxUnlockTime_)
}

// Initialize is a paid mutator transaction binding the contract method 0xa01e3f74.
//
// Solidity: function initialize(address owner_, address asset_, string vaultTokenName_, string vaultTokenSymbol_, address vaultFactory_, uint256 profitMaxUnlockTime_) returns()
func (_WeberaVault *WeberaVaultTransactorSession) Initialize(owner_ common.Address, asset_ common.Address, vaultTokenName_ string, vaultTokenSymbol_ string, vaultFactory_ common.Address, profitMaxUnlockTime_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.Initialize(&_WeberaVault.TransactOpts, owner_, asset_, vaultTokenName_, vaultTokenSymbol_, vaultFactory_, profitMaxUnlockTime_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_WeberaVault *WeberaVaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_WeberaVault *WeberaVaultSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.Mint(&_WeberaVault.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_WeberaVault *WeberaVaultTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.Mint(&_WeberaVault.TransactOpts, shares, receiver)
}

// OpenVault is a paid mutator transaction binding the contract method 0x86f54712.
//
// Solidity: function openVault() returns()
func (_WeberaVault *WeberaVaultTransactor) OpenVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "openVault")
}

// OpenVault is a paid mutator transaction binding the contract method 0x86f54712.
//
// Solidity: function openVault() returns()
func (_WeberaVault *WeberaVaultSession) OpenVault() (*types.Transaction, error) {
	return _WeberaVault.Contract.OpenVault(&_WeberaVault.TransactOpts)
}

// OpenVault is a paid mutator transaction binding the contract method 0x86f54712.
//
// Solidity: function openVault() returns()
func (_WeberaVault *WeberaVaultTransactorSession) OpenVault() (*types.Transaction, error) {
	return _WeberaVault.Contract.OpenVault(&_WeberaVault.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_WeberaVault *WeberaVaultTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_WeberaVault *WeberaVaultSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.Redeem(&_WeberaVault.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_WeberaVault *WeberaVaultTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.Redeem(&_WeberaVault.TransactOpts, shares, receiver, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WeberaVault *WeberaVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WeberaVault *WeberaVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _WeberaVault.Contract.RenounceOwnership(&_WeberaVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WeberaVault *WeberaVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WeberaVault.Contract.RenounceOwnership(&_WeberaVault.TransactOpts)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_WeberaVault *WeberaVaultTransactor) RevokeStrategy(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "revokeStrategy", strategy)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_WeberaVault *WeberaVaultSession) RevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.RevokeStrategy(&_WeberaVault.TransactOpts, strategy)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0xbb994d48.
//
// Solidity: function revokeStrategy(address strategy) returns()
func (_WeberaVault *WeberaVaultTransactorSession) RevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.RevokeStrategy(&_WeberaVault.TransactOpts, strategy)
}

// SetHarvester is a paid mutator transaction binding the contract method 0x15de1daa.
//
// Solidity: function setHarvester(address harvester_) returns()
func (_WeberaVault *WeberaVaultTransactor) SetHarvester(opts *bind.TransactOpts, harvester_ common.Address) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "setHarvester", harvester_)
}

// SetHarvester is a paid mutator transaction binding the contract method 0x15de1daa.
//
// Solidity: function setHarvester(address harvester_) returns()
func (_WeberaVault *WeberaVaultSession) SetHarvester(harvester_ common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.SetHarvester(&_WeberaVault.TransactOpts, harvester_)
}

// SetHarvester is a paid mutator transaction binding the contract method 0x15de1daa.
//
// Solidity: function setHarvester(address harvester_) returns()
func (_WeberaVault *WeberaVaultTransactorSession) SetHarvester(harvester_ common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.SetHarvester(&_WeberaVault.TransactOpts, harvester_)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 profitMaxUnlockTime_) returns()
func (_WeberaVault *WeberaVaultTransactor) SetProfitMaxUnlockTime(opts *bind.TransactOpts, profitMaxUnlockTime_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "setProfitMaxUnlockTime", profitMaxUnlockTime_)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 profitMaxUnlockTime_) returns()
func (_WeberaVault *WeberaVaultSession) SetProfitMaxUnlockTime(profitMaxUnlockTime_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.SetProfitMaxUnlockTime(&_WeberaVault.TransactOpts, profitMaxUnlockTime_)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 profitMaxUnlockTime_) returns()
func (_WeberaVault *WeberaVaultTransactorSession) SetProfitMaxUnlockTime(profitMaxUnlockTime_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.SetProfitMaxUnlockTime(&_WeberaVault.TransactOpts, profitMaxUnlockTime_)
}

// SetVaultRouter is a paid mutator transaction binding the contract method 0x780e72ff.
//
// Solidity: function setVaultRouter(address vaultRouter_) returns()
func (_WeberaVault *WeberaVaultTransactor) SetVaultRouter(opts *bind.TransactOpts, vaultRouter_ common.Address) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "setVaultRouter", vaultRouter_)
}

// SetVaultRouter is a paid mutator transaction binding the contract method 0x780e72ff.
//
// Solidity: function setVaultRouter(address vaultRouter_) returns()
func (_WeberaVault *WeberaVaultSession) SetVaultRouter(vaultRouter_ common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.SetVaultRouter(&_WeberaVault.TransactOpts, vaultRouter_)
}

// SetVaultRouter is a paid mutator transaction binding the contract method 0x780e72ff.
//
// Solidity: function setVaultRouter(address vaultRouter_) returns()
func (_WeberaVault *WeberaVaultTransactorSession) SetVaultRouter(vaultRouter_ common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.SetVaultRouter(&_WeberaVault.TransactOpts, vaultRouter_)
}

// SetWithdrawQueue is a paid mutator transaction binding the contract method 0x80a30e30.
//
// Solidity: function setWithdrawQueue(address[] queue) returns()
func (_WeberaVault *WeberaVaultTransactor) SetWithdrawQueue(opts *bind.TransactOpts, queue []common.Address) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "setWithdrawQueue", queue)
}

// SetWithdrawQueue is a paid mutator transaction binding the contract method 0x80a30e30.
//
// Solidity: function setWithdrawQueue(address[] queue) returns()
func (_WeberaVault *WeberaVaultSession) SetWithdrawQueue(queue []common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.SetWithdrawQueue(&_WeberaVault.TransactOpts, queue)
}

// SetWithdrawQueue is a paid mutator transaction binding the contract method 0x80a30e30.
//
// Solidity: function setWithdrawQueue(address[] queue) returns()
func (_WeberaVault *WeberaVaultTransactorSession) SetWithdrawQueue(queue []common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.SetWithdrawQueue(&_WeberaVault.TransactOpts, queue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WeberaVault *WeberaVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WeberaVault *WeberaVaultSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.Transfer(&_WeberaVault.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_WeberaVault *WeberaVaultTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.Transfer(&_WeberaVault.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WeberaVault *WeberaVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WeberaVault *WeberaVaultSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.TransferFrom(&_WeberaVault.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_WeberaVault *WeberaVaultTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.TransferFrom(&_WeberaVault.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WeberaVault *WeberaVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WeberaVault *WeberaVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.TransferOwnership(&_WeberaVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WeberaVault *WeberaVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.TransferOwnership(&_WeberaVault.TransactOpts, newOwner)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x4492e200.
//
// Solidity: function updateDebt(address strategy_, uint256 targetDebt_, uint256 maxLoss_) returns()
func (_WeberaVault *WeberaVaultTransactor) UpdateDebt(opts *bind.TransactOpts, strategy_ common.Address, targetDebt_ *big.Int, maxLoss_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "updateDebt", strategy_, targetDebt_, maxLoss_)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x4492e200.
//
// Solidity: function updateDebt(address strategy_, uint256 targetDebt_, uint256 maxLoss_) returns()
func (_WeberaVault *WeberaVaultSession) UpdateDebt(strategy_ common.Address, targetDebt_ *big.Int, maxLoss_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.UpdateDebt(&_WeberaVault.TransactOpts, strategy_, targetDebt_, maxLoss_)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x4492e200.
//
// Solidity: function updateDebt(address strategy_, uint256 targetDebt_, uint256 maxLoss_) returns()
func (_WeberaVault *WeberaVaultTransactorSession) UpdateDebt(strategy_ common.Address, targetDebt_ *big.Int, maxLoss_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.UpdateDebt(&_WeberaVault.TransactOpts, strategy_, targetDebt_, maxLoss_)
}

// UpdateStrategyMaxDebt is a paid mutator transaction binding the contract method 0x1c9a9e38.
//
// Solidity: function updateStrategyMaxDebt(address strategy_, uint256 maxDebt_) returns()
func (_WeberaVault *WeberaVaultTransactor) UpdateStrategyMaxDebt(opts *bind.TransactOpts, strategy_ common.Address, maxDebt_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "updateStrategyMaxDebt", strategy_, maxDebt_)
}

// UpdateStrategyMaxDebt is a paid mutator transaction binding the contract method 0x1c9a9e38.
//
// Solidity: function updateStrategyMaxDebt(address strategy_, uint256 maxDebt_) returns()
func (_WeberaVault *WeberaVaultSession) UpdateStrategyMaxDebt(strategy_ common.Address, maxDebt_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.UpdateStrategyMaxDebt(&_WeberaVault.TransactOpts, strategy_, maxDebt_)
}

// UpdateStrategyMaxDebt is a paid mutator transaction binding the contract method 0x1c9a9e38.
//
// Solidity: function updateStrategyMaxDebt(address strategy_, uint256 maxDebt_) returns()
func (_WeberaVault *WeberaVaultTransactorSession) UpdateStrategyMaxDebt(strategy_ common.Address, maxDebt_ *big.Int) (*types.Transaction, error) {
	return _WeberaVault.Contract.UpdateStrategyMaxDebt(&_WeberaVault.TransactOpts, strategy_, maxDebt_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_WeberaVault *WeberaVaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WeberaVault.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_WeberaVault *WeberaVaultSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.Withdraw(&_WeberaVault.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_WeberaVault *WeberaVaultTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _WeberaVault.Contract.Withdraw(&_WeberaVault.TransactOpts, assets, receiver, owner)
}

// WeberaVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WeberaVault contract.
type WeberaVaultApprovalIterator struct {
	Event *WeberaVaultApproval // Event containing the contract specifics and raw log

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
func (it *WeberaVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultApproval)
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
		it.Event = new(WeberaVaultApproval)
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
func (it *WeberaVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultApproval represents a Approval event raised by the WeberaVault contract.
type WeberaVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WeberaVault *WeberaVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WeberaVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultApprovalIterator{contract: _WeberaVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WeberaVault *WeberaVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WeberaVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultApproval)
				if err := _WeberaVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_WeberaVault *WeberaVaultFilterer) ParseApproval(log types.Log) (*WeberaVaultApproval, error) {
	event := new(WeberaVaultApproval)
	if err := _WeberaVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeberaVaultDebtUpdatedIterator is returned from FilterDebtUpdated and is used to iterate over the raw logs and unpacked data for DebtUpdated events raised by the WeberaVault contract.
type WeberaVaultDebtUpdatedIterator struct {
	Event *WeberaVaultDebtUpdated // Event containing the contract specifics and raw log

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
func (it *WeberaVaultDebtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultDebtUpdated)
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
		it.Event = new(WeberaVaultDebtUpdated)
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
func (it *WeberaVaultDebtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultDebtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultDebtUpdated represents a DebtUpdated event raised by the WeberaVault contract.
type WeberaVaultDebtUpdated struct {
	Strategy common.Address
	OldDebt  *big.Int
	NewDebt  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDebtUpdated is a free log retrieval operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 oldDebt, uint256 newDebt)
func (_WeberaVault *WeberaVaultFilterer) FilterDebtUpdated(opts *bind.FilterOpts, strategy []common.Address) (*WeberaVaultDebtUpdatedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "DebtUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultDebtUpdatedIterator{contract: _WeberaVault.contract, event: "DebtUpdated", logs: logs, sub: sub}, nil
}

// WatchDebtUpdated is a free log subscription operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 oldDebt, uint256 newDebt)
func (_WeberaVault *WeberaVaultFilterer) WatchDebtUpdated(opts *bind.WatchOpts, sink chan<- *WeberaVaultDebtUpdated, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "DebtUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultDebtUpdated)
				if err := _WeberaVault.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
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

// ParseDebtUpdated is a log parse operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 oldDebt, uint256 newDebt)
func (_WeberaVault *WeberaVaultFilterer) ParseDebtUpdated(log types.Log) (*WeberaVaultDebtUpdated, error) {
	event := new(WeberaVaultDebtUpdated)
	if err := _WeberaVault.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeberaVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WeberaVault contract.
type WeberaVaultDepositIterator struct {
	Event *WeberaVaultDeposit // Event containing the contract specifics and raw log

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
func (it *WeberaVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultDeposit)
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
		it.Event = new(WeberaVaultDeposit)
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
func (it *WeberaVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultDeposit represents a Deposit event raised by the WeberaVault contract.
type WeberaVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_WeberaVault *WeberaVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*WeberaVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultDepositIterator{contract: _WeberaVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_WeberaVault *WeberaVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WeberaVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultDeposit)
				if err := _WeberaVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_WeberaVault *WeberaVaultFilterer) ParseDeposit(log types.Log) (*WeberaVaultDeposit, error) {
	event := new(WeberaVaultDeposit)
	if err := _WeberaVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeberaVaultHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the WeberaVault contract.
type WeberaVaultHarvestIterator struct {
	Event *WeberaVaultHarvest // Event containing the contract specifics and raw log

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
func (it *WeberaVaultHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultHarvest)
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
		it.Event = new(WeberaVaultHarvest)
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
func (it *WeberaVaultHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultHarvest represents a Harvest event raised by the WeberaVault contract.
type WeberaVaultHarvest struct {
	Strategy     common.Address
	Profit       *big.Int
	Loss         *big.Int
	CurrentDebt  *big.Int
	TotalFees    *big.Int
	TotalRefunds *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0xeaedd1267621f4a8ee97011faf0e67800ad3063714bd179f078c980bb9a3fa3d.
//
// Solidity: event Harvest(address indexed strategy, uint256 profit, uint256 loss, uint256 currentDebt, uint256 totalFees, uint256 totalRefunds)
func (_WeberaVault *WeberaVaultFilterer) FilterHarvest(opts *bind.FilterOpts, strategy []common.Address) (*WeberaVaultHarvestIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "Harvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultHarvestIterator{contract: _WeberaVault.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0xeaedd1267621f4a8ee97011faf0e67800ad3063714bd179f078c980bb9a3fa3d.
//
// Solidity: event Harvest(address indexed strategy, uint256 profit, uint256 loss, uint256 currentDebt, uint256 totalFees, uint256 totalRefunds)
func (_WeberaVault *WeberaVaultFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *WeberaVaultHarvest, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "Harvest", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultHarvest)
				if err := _WeberaVault.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0xeaedd1267621f4a8ee97011faf0e67800ad3063714bd179f078c980bb9a3fa3d.
//
// Solidity: event Harvest(address indexed strategy, uint256 profit, uint256 loss, uint256 currentDebt, uint256 totalFees, uint256 totalRefunds)
func (_WeberaVault *WeberaVaultFilterer) ParseHarvest(log types.Log) (*WeberaVaultHarvest, error) {
	event := new(WeberaVaultHarvest)
	if err := _WeberaVault.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeberaVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WeberaVault contract.
type WeberaVaultInitializedIterator struct {
	Event *WeberaVaultInitialized // Event containing the contract specifics and raw log

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
func (it *WeberaVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultInitialized)
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
		it.Event = new(WeberaVaultInitialized)
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
func (it *WeberaVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultInitialized represents a Initialized event raised by the WeberaVault contract.
type WeberaVaultInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_WeberaVault *WeberaVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*WeberaVaultInitializedIterator, error) {

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WeberaVaultInitializedIterator{contract: _WeberaVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_WeberaVault *WeberaVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WeberaVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultInitialized)
				if err := _WeberaVault.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WeberaVault *WeberaVaultFilterer) ParseInitialized(log types.Log) (*WeberaVaultInitialized, error) {
	event := new(WeberaVaultInitialized)
	if err := _WeberaVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeberaVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WeberaVault contract.
type WeberaVaultOwnershipTransferredIterator struct {
	Event *WeberaVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WeberaVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultOwnershipTransferred)
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
		it.Event = new(WeberaVaultOwnershipTransferred)
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
func (it *WeberaVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultOwnershipTransferred represents a OwnershipTransferred event raised by the WeberaVault contract.
type WeberaVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WeberaVault *WeberaVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WeberaVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultOwnershipTransferredIterator{contract: _WeberaVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WeberaVault *WeberaVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WeberaVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultOwnershipTransferred)
				if err := _WeberaVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WeberaVault *WeberaVaultFilterer) ParseOwnershipTransferred(log types.Log) (*WeberaVaultOwnershipTransferred, error) {
	event := new(WeberaVaultOwnershipTransferred)
	if err := _WeberaVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeberaVaultStrategyAddedIterator is returned from FilterStrategyAdded and is used to iterate over the raw logs and unpacked data for StrategyAdded events raised by the WeberaVault contract.
type WeberaVaultStrategyAddedIterator struct {
	Event *WeberaVaultStrategyAdded // Event containing the contract specifics and raw log

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
func (it *WeberaVaultStrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultStrategyAdded)
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
		it.Event = new(WeberaVaultStrategyAdded)
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
func (it *WeberaVaultStrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultStrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultStrategyAdded represents a StrategyAdded event raised by the WeberaVault contract.
type WeberaVaultStrategyAdded struct {
	Strategy common.Address
	Length   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyAdded is a free log retrieval operation binding the contract event 0x2f564a83158ad1831793ad3e69257b52f39ece5d49cb0d8746708ecb9ef964da.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 length)
func (_WeberaVault *WeberaVaultFilterer) FilterStrategyAdded(opts *bind.FilterOpts, strategy []common.Address) (*WeberaVaultStrategyAddedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultStrategyAddedIterator{contract: _WeberaVault.contract, event: "StrategyAdded", logs: logs, sub: sub}, nil
}

// WatchStrategyAdded is a free log subscription operation binding the contract event 0x2f564a83158ad1831793ad3e69257b52f39ece5d49cb0d8746708ecb9ef964da.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 length)
func (_WeberaVault *WeberaVaultFilterer) WatchStrategyAdded(opts *bind.WatchOpts, sink chan<- *WeberaVaultStrategyAdded, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "StrategyAdded", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultStrategyAdded)
				if err := _WeberaVault.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
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

// ParseStrategyAdded is a log parse operation binding the contract event 0x2f564a83158ad1831793ad3e69257b52f39ece5d49cb0d8746708ecb9ef964da.
//
// Solidity: event StrategyAdded(address indexed strategy, uint256 length)
func (_WeberaVault *WeberaVaultFilterer) ParseStrategyAdded(log types.Log) (*WeberaVaultStrategyAdded, error) {
	event := new(WeberaVaultStrategyAdded)
	if err := _WeberaVault.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeberaVaultStrategyRevokedIterator is returned from FilterStrategyRevoked and is used to iterate over the raw logs and unpacked data for StrategyRevoked events raised by the WeberaVault contract.
type WeberaVaultStrategyRevokedIterator struct {
	Event *WeberaVaultStrategyRevoked // Event containing the contract specifics and raw log

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
func (it *WeberaVaultStrategyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultStrategyRevoked)
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
		it.Event = new(WeberaVaultStrategyRevoked)
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
func (it *WeberaVaultStrategyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultStrategyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultStrategyRevoked represents a StrategyRevoked event raised by the WeberaVault contract.
type WeberaVaultStrategyRevoked struct {
	Strategy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStrategyRevoked is a free log retrieval operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_WeberaVault *WeberaVaultFilterer) FilterStrategyRevoked(opts *bind.FilterOpts, strategy []common.Address) (*WeberaVaultStrategyRevokedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "StrategyRevoked", strategyRule)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultStrategyRevokedIterator{contract: _WeberaVault.contract, event: "StrategyRevoked", logs: logs, sub: sub}, nil
}

// WatchStrategyRevoked is a free log subscription operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_WeberaVault *WeberaVaultFilterer) WatchStrategyRevoked(opts *bind.WatchOpts, sink chan<- *WeberaVaultStrategyRevoked, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "StrategyRevoked", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultStrategyRevoked)
				if err := _WeberaVault.contract.UnpackLog(event, "StrategyRevoked", log); err != nil {
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

// ParseStrategyRevoked is a log parse operation binding the contract event 0x4201c688d84c01154d321afa0c72f1bffe9eef53005c9de9d035074e71e9b32a.
//
// Solidity: event StrategyRevoked(address indexed strategy)
func (_WeberaVault *WeberaVaultFilterer) ParseStrategyRevoked(log types.Log) (*WeberaVaultStrategyRevoked, error) {
	event := new(WeberaVaultStrategyRevoked)
	if err := _WeberaVault.contract.UnpackLog(event, "StrategyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeberaVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WeberaVault contract.
type WeberaVaultTransferIterator struct {
	Event *WeberaVaultTransfer // Event containing the contract specifics and raw log

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
func (it *WeberaVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultTransfer)
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
		it.Event = new(WeberaVaultTransfer)
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
func (it *WeberaVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultTransfer represents a Transfer event raised by the WeberaVault contract.
type WeberaVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WeberaVault *WeberaVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WeberaVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultTransferIterator{contract: _WeberaVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WeberaVault *WeberaVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WeberaVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultTransfer)
				if err := _WeberaVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_WeberaVault *WeberaVaultFilterer) ParseTransfer(log types.Log) (*WeberaVaultTransfer, error) {
	event := new(WeberaVaultTransfer)
	if err := _WeberaVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeberaVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the WeberaVault contract.
type WeberaVaultWithdrawIterator struct {
	Event *WeberaVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *WeberaVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultWithdraw)
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
		it.Event = new(WeberaVaultWithdraw)
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
func (it *WeberaVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultWithdraw represents a Withdraw event raised by the WeberaVault contract.
type WeberaVaultWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_WeberaVault *WeberaVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*WeberaVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WeberaVaultWithdrawIterator{contract: _WeberaVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_WeberaVault *WeberaVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *WeberaVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultWithdraw)
				if err := _WeberaVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_WeberaVault *WeberaVaultFilterer) ParseWithdraw(log types.Log) (*WeberaVaultWithdraw, error) {
	event := new(WeberaVaultWithdraw)
	if err := _WeberaVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WeberaVaultWithdrawQueueUpdatedIterator is returned from FilterWithdrawQueueUpdated and is used to iterate over the raw logs and unpacked data for WithdrawQueueUpdated events raised by the WeberaVault contract.
type WeberaVaultWithdrawQueueUpdatedIterator struct {
	Event *WeberaVaultWithdrawQueueUpdated // Event containing the contract specifics and raw log

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
func (it *WeberaVaultWithdrawQueueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WeberaVaultWithdrawQueueUpdated)
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
		it.Event = new(WeberaVaultWithdrawQueueUpdated)
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
func (it *WeberaVaultWithdrawQueueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WeberaVaultWithdrawQueueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WeberaVaultWithdrawQueueUpdated represents a WithdrawQueueUpdated event raised by the WeberaVault contract.
type WeberaVaultWithdrawQueueUpdated struct {
	OldStrategyQueue []common.Address
	NewStrategyQueue []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawQueueUpdated is a free log retrieval operation binding the contract event 0x0ef29da578ca6365f361a762138684560226292af8e477c349267a8fa5c80488.
//
// Solidity: event WithdrawQueueUpdated(address[] oldStrategyQueue, address[] newStrategyQueue)
func (_WeberaVault *WeberaVaultFilterer) FilterWithdrawQueueUpdated(opts *bind.FilterOpts) (*WeberaVaultWithdrawQueueUpdatedIterator, error) {

	logs, sub, err := _WeberaVault.contract.FilterLogs(opts, "WithdrawQueueUpdated")
	if err != nil {
		return nil, err
	}
	return &WeberaVaultWithdrawQueueUpdatedIterator{contract: _WeberaVault.contract, event: "WithdrawQueueUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawQueueUpdated is a free log subscription operation binding the contract event 0x0ef29da578ca6365f361a762138684560226292af8e477c349267a8fa5c80488.
//
// Solidity: event WithdrawQueueUpdated(address[] oldStrategyQueue, address[] newStrategyQueue)
func (_WeberaVault *WeberaVaultFilterer) WatchWithdrawQueueUpdated(opts *bind.WatchOpts, sink chan<- *WeberaVaultWithdrawQueueUpdated) (event.Subscription, error) {

	logs, sub, err := _WeberaVault.contract.WatchLogs(opts, "WithdrawQueueUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WeberaVaultWithdrawQueueUpdated)
				if err := _WeberaVault.contract.UnpackLog(event, "WithdrawQueueUpdated", log); err != nil {
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

// ParseWithdrawQueueUpdated is a log parse operation binding the contract event 0x0ef29da578ca6365f361a762138684560226292af8e477c349267a8fa5c80488.
//
// Solidity: event WithdrawQueueUpdated(address[] oldStrategyQueue, address[] newStrategyQueue)
func (_WeberaVault *WeberaVaultFilterer) ParseWithdrawQueueUpdated(log types.Log) (*WeberaVaultWithdrawQueueUpdated, error) {
	event := new(WeberaVaultWithdrawQueueUpdated)
	if err := _WeberaVault.contract.UnpackLog(event, "WithdrawQueueUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
