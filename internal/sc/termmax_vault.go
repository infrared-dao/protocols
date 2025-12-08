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

// CurveCut is an auto generated low-level Go binding around an user-defined struct.
type CurveCut struct {
	XtReserve *big.Int
	LiqSquare *big.Int
	Offset    *big.Int
}

// CurveCuts is an auto generated low-level Go binding around an user-defined struct.
type CurveCuts struct {
	LendCurveCuts   []CurveCut
	BorrowCurveCuts []CurveCut
}

// OrderV2ConfigurationParams is an auto generated low-level Go binding around an user-defined struct.
type OrderV2ConfigurationParams struct {
	OriginalVirtualXtReserve *big.Int
	VirtualXtReserve         *big.Int
	MaxXtReserve             *big.Int
	CurveCuts                CurveCuts
}

// PendingAddress is an auto generated low-level Go binding around an user-defined struct.
type PendingAddress struct {
	Value   common.Address
	ValidAt uint64
}

// PendingUint192 is an auto generated low-level Go binding around an user-defined struct.
type PendingUint192 struct {
	Value   *big.Int
	ValidAt uint64
}

// VaultInitialParams is an auto generated low-level Go binding around an user-defined struct.
type VaultInitialParams struct {
	Admin              common.Address
	Curator            common.Address
	Timelock           *big.Int
	Asset              common.Address
	MaxCapacity        *big.Int
	Name               string
	Symbol             string
	PerformanceFeeRate uint64
}

// VaultInitialParamsV2 is an auto generated low-level Go binding around an user-defined struct.
type VaultInitialParamsV2 struct {
	Admin              common.Address
	Curator            common.Address
	Guardian           common.Address
	Timelock           *big.Int
	Asset              common.Address
	Pool               common.Address
	MaxCapacity        *big.Int
	Name               string
	Symbol             string
	PerformanceFeeRate uint64
	MinApy             uint64
}

// ERC4626VaultMetaData contains all meta data concerning the ERC4626Vault contract.
var ERC4626VaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ORDER_MANAGER_SINGLETON_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AboveMaxTimelock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BelowMinTimelock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanNotTransferUintMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralIsAsset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidActionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPendingValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCuratorRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGuardianRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PerformanceFeeRateExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oldActionId\",\"type\":\"uint256\"}],\"name\":\"ReentrantCallBetweenActions\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SupplyQueueNoLongerSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseApyInsteadOfApr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseVaultInitialParamsV2\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalQueueNoLongerSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"badDebt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collateralOut\",\"type\":\"uint256\"}],\"name\":\"DealBadDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"RevokePendingMarket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingMinApy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingPerformanceFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCapacity\",\"type\":\"uint256\"}],\"name\":\"SetCapacity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCurator\",\"type\":\"address\"}],\"name\":\"SetCurator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"SetGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"}],\"name\":\"SetMarketWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newMinApy\",\"type\":\"uint64\"}],\"name\":\"SetMinApy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPerformanceFeeRate\",\"type\":\"uint256\"}],\"name\":\"SetPerformanceFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"SetPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"SetTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"name\":\"SubmitGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"name\":\"SubmitMarketToWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newMinApy\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"name\":\"SubmitMinApy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"name\":\"SubmitPendingPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPerformanceFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"name\":\"SubmitPerformanceFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"name\":\"SubmitTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"order\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"WithdrawFts\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ORDER_MANAGER_SINGLETON\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"acceptMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptPendingMinApy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptPerformanceFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accretingPrincipal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ftReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"xtReserve\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"deltaFt\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"deltaXt\",\"type\":\"int256\"}],\"name\":\"afterSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"annualizedInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"badDebtMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITermMaxMarketV2\",\"name\":\"market\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"originalVirtualXtReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"virtualXtReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxXtReserve\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"xtReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liqSquare\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"offset\",\"type\":\"int256\"}],\"internalType\":\"structCurveCut[]\",\"name\":\"lendCurveCuts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"xtReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liqSquare\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"offset\",\"type\":\"int256\"}],\"internalType\":\"structCurveCut[]\",\"name\":\"borrowCurveCuts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCurveCuts\",\"name\":\"curveCuts\",\"type\":\"tuple\"}],\"internalType\":\"structOrderV2ConfigurationParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"createOrder\",\"outputs\":[{\"internalType\":\"contractITermMaxOrderV2\",\"name\":\"order\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"badDebtAmt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"dealBadDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"curator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxCapacity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"performanceFeeRate\",\"type\":\"uint64\"}],\"internalType\":\"structVaultInitialParams\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"curator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timelock\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"contractIERC4626\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxCapacity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"performanceFeeRate\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minApy\",\"type\":\"uint64\"}],\"internalType\":\"structVaultInitialParamsV2\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"marketWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minApy\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"order\",\"type\":\"address\"}],\"name\":\"orderMaturity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGuardian\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"internalType\":\"structPendingAddress\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"pendingMarkets\",\"outputs\":[{\"components\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"internalType\":\"structPendingUint192\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingMinApy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"internalType\":\"structPendingUint192\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingPerformanceFeeRate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"internalType\":\"structPendingUint192\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingPool\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"internalType\":\"structPendingAddress\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingPools\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"internalType\":\"structPendingAddress\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingTimelock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"internalType\":\"structPendingUint192\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"performanceFeeRate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITermMaxOrderV2\",\"name\":\"order\",\"type\":\"address\"}],\"name\":\"redeemOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"badDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deliveryCollateral\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"orders\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"removedLiquidities\",\"type\":\"uint256[]\"}],\"name\":\"removeLiquidityFromOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"}],\"name\":\"revokePendingMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingMinApy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingPerformanceFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCapacity\",\"type\":\"uint256\"}],\"name\":\"setCapacity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCurator\",\"type\":\"address\"}],\"name\":\"setCurator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"submitGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"market\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWhitelisted\",\"type\":\"bool\"}],\"name\":\"submitMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newMinApy\",\"type\":\"uint64\"}],\"name\":\"submitPendingMinApy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool_\",\"type\":\"address\"}],\"name\":\"submitPendingPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint184\",\"name\":\"newPerformanceFeeRate\",\"type\":\"uint184\"}],\"name\":\"submitPerformanceFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"submitTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supplyQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalFt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"orders\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"originalVirtualXtReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"virtualXtReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxXtReserve\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"xtReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liqSquare\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"offset\",\"type\":\"int256\"}],\"internalType\":\"structCurveCut[]\",\"name\":\"lendCurveCuts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"xtReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liqSquare\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"offset\",\"type\":\"int256\"}],\"internalType\":\"structCurveCut[]\",\"name\":\"borrowCurveCuts\",\"type\":\"tuple[]\"}],\"internalType\":\"structCurveCuts\",\"name\":\"curveCuts\",\"type\":\"tuple\"}],\"internalType\":\"structOrderV2ConfigurationParams[]\",\"name\":\"orderConfigs\",\"type\":\"tuple[]\"}],\"name\":\"updateOrdersConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"order\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawFts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ERC4626VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC4626VaultMetaData.ABI instead.
var ERC4626VaultABI = ERC4626VaultMetaData.ABI

// ERC4626Vault is an auto generated Go binding around an Ethereum contract.
type ERC4626Vault struct {
	ERC4626VaultCaller     // Read-only binding to the contract
	ERC4626VaultTransactor // Write-only binding to the contract
	ERC4626VaultFilterer   // Log filterer for contract events
}

// ERC4626VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC4626VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC4626VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC4626VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC4626VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC4626VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC4626VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC4626VaultSession struct {
	Contract     *ERC4626Vault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC4626VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC4626VaultCallerSession struct {
	Contract *ERC4626VaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ERC4626VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC4626VaultTransactorSession struct {
	Contract     *ERC4626VaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ERC4626VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC4626VaultRaw struct {
	Contract *ERC4626Vault // Generic contract binding to access the raw methods on
}

// ERC4626VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC4626VaultCallerRaw struct {
	Contract *ERC4626VaultCaller // Generic read-only contract binding to access the raw methods on
}

// ERC4626VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC4626VaultTransactorRaw struct {
	Contract *ERC4626VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC4626Vault creates a new instance of ERC4626Vault, bound to a specific deployed contract.
func NewERC4626Vault(address common.Address, backend bind.ContractBackend) (*ERC4626Vault, error) {
	contract, err := bindERC4626Vault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC4626Vault{ERC4626VaultCaller: ERC4626VaultCaller{contract: contract}, ERC4626VaultTransactor: ERC4626VaultTransactor{contract: contract}, ERC4626VaultFilterer: ERC4626VaultFilterer{contract: contract}}, nil
}

// NewERC4626VaultCaller creates a new read-only instance of ERC4626Vault, bound to a specific deployed contract.
func NewERC4626VaultCaller(address common.Address, caller bind.ContractCaller) (*ERC4626VaultCaller, error) {
	contract, err := bindERC4626Vault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultCaller{contract: contract}, nil
}

// NewERC4626VaultTransactor creates a new write-only instance of ERC4626Vault, bound to a specific deployed contract.
func NewERC4626VaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC4626VaultTransactor, error) {
	contract, err := bindERC4626Vault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultTransactor{contract: contract}, nil
}

// NewERC4626VaultFilterer creates a new log filterer instance of ERC4626Vault, bound to a specific deployed contract.
func NewERC4626VaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC4626VaultFilterer, error) {
	contract, err := bindERC4626Vault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultFilterer{contract: contract}, nil
}

// bindERC4626Vault binds a generic wrapper to an already deployed contract.
func bindERC4626Vault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC4626VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC4626Vault *ERC4626VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC4626Vault.Contract.ERC4626VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC4626Vault *ERC4626VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.ERC4626VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC4626Vault *ERC4626VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.ERC4626VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC4626Vault *ERC4626VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC4626Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC4626Vault *ERC4626VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC4626Vault *ERC4626VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.contract.Transact(opts, method, params...)
}

// ORDERMANAGERSINGLETON is a free data retrieval call binding the contract method 0x0f4f0c32.
//
// Solidity: function ORDER_MANAGER_SINGLETON() view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) ORDERMANAGERSINGLETON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "ORDER_MANAGER_SINGLETON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORDERMANAGERSINGLETON is a free data retrieval call binding the contract method 0x0f4f0c32.
//
// Solidity: function ORDER_MANAGER_SINGLETON() view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) ORDERMANAGERSINGLETON() (common.Address, error) {
	return _ERC4626Vault.Contract.ORDERMANAGERSINGLETON(&_ERC4626Vault.CallOpts)
}

// ORDERMANAGERSINGLETON is a free data retrieval call binding the contract method 0x0f4f0c32.
//
// Solidity: function ORDER_MANAGER_SINGLETON() view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) ORDERMANAGERSINGLETON() (common.Address, error) {
	return _ERC4626Vault.Contract.ORDERMANAGERSINGLETON(&_ERC4626Vault.CallOpts)
}

// AccretingPrincipal is a free data retrieval call binding the contract method 0x594d16f7.
//
// Solidity: function accretingPrincipal() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) AccretingPrincipal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "accretingPrincipal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccretingPrincipal is a free data retrieval call binding the contract method 0x594d16f7.
//
// Solidity: function accretingPrincipal() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) AccretingPrincipal() (*big.Int, error) {
	return _ERC4626Vault.Contract.AccretingPrincipal(&_ERC4626Vault.CallOpts)
}

// AccretingPrincipal is a free data retrieval call binding the contract method 0x594d16f7.
//
// Solidity: function accretingPrincipal() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) AccretingPrincipal() (*big.Int, error) {
	return _ERC4626Vault.Contract.AccretingPrincipal(&_ERC4626Vault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.Allowance(&_ERC4626Vault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.Allowance(&_ERC4626Vault.CallOpts, owner, spender)
}

// AnnualizedInterest is a free data retrieval call binding the contract method 0xe5c6c88d.
//
// Solidity: function annualizedInterest() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) AnnualizedInterest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "annualizedInterest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AnnualizedInterest is a free data retrieval call binding the contract method 0xe5c6c88d.
//
// Solidity: function annualizedInterest() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) AnnualizedInterest() (*big.Int, error) {
	return _ERC4626Vault.Contract.AnnualizedInterest(&_ERC4626Vault.CallOpts)
}

// AnnualizedInterest is a free data retrieval call binding the contract method 0xe5c6c88d.
//
// Solidity: function annualizedInterest() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) AnnualizedInterest() (*big.Int, error) {
	return _ERC4626Vault.Contract.AnnualizedInterest(&_ERC4626Vault.CallOpts)
}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) Apr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "apr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Apr() (*big.Int, error) {
	return _ERC4626Vault.Contract.Apr(&_ERC4626Vault.CallOpts)
}

// Apr is a free data retrieval call binding the contract method 0x57ded9c9.
//
// Solidity: function apr() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) Apr() (*big.Int, error) {
	return _ERC4626Vault.Contract.Apr(&_ERC4626Vault.CallOpts)
}

// Apy is a free data retrieval call binding the contract method 0x3bcfc4b8.
//
// Solidity: function apy() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) Apy(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "apy")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Apy is a free data retrieval call binding the contract method 0x3bcfc4b8.
//
// Solidity: function apy() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Apy() (*big.Int, error) {
	return _ERC4626Vault.Contract.Apy(&_ERC4626Vault.CallOpts)
}

// Apy is a free data retrieval call binding the contract method 0x3bcfc4b8.
//
// Solidity: function apy() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) Apy() (*big.Int, error) {
	return _ERC4626Vault.Contract.Apy(&_ERC4626Vault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) Asset() (common.Address, error) {
	return _ERC4626Vault.Contract.Asset(&_ERC4626Vault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) Asset() (common.Address, error) {
	return _ERC4626Vault.Contract.Asset(&_ERC4626Vault.CallOpts)
}

// BadDebtMapping is a free data retrieval call binding the contract method 0x618f9694.
//
// Solidity: function badDebtMapping(address collateral) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) BadDebtMapping(opts *bind.CallOpts, collateral common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "badDebtMapping", collateral)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BadDebtMapping is a free data retrieval call binding the contract method 0x618f9694.
//
// Solidity: function badDebtMapping(address collateral) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) BadDebtMapping(collateral common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.BadDebtMapping(&_ERC4626Vault.CallOpts, collateral)
}

// BadDebtMapping is a free data retrieval call binding the contract method 0x618f9694.
//
// Solidity: function badDebtMapping(address collateral) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) BadDebtMapping(collateral common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.BadDebtMapping(&_ERC4626Vault.CallOpts, collateral)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.BalanceOf(&_ERC4626Vault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.BalanceOf(&_ERC4626Vault.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.ConvertToAssets(&_ERC4626Vault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.ConvertToAssets(&_ERC4626Vault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.ConvertToShares(&_ERC4626Vault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.ConvertToShares(&_ERC4626Vault.CallOpts, assets)
}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) Curator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "curator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) Curator() (common.Address, error) {
	return _ERC4626Vault.Contract.Curator(&_ERC4626Vault.CallOpts)
}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) Curator() (common.Address, error) {
	return _ERC4626Vault.Contract.Curator(&_ERC4626Vault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC4626Vault *ERC4626VaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC4626Vault *ERC4626VaultSession) Decimals() (uint8, error) {
	return _ERC4626Vault.Contract.Decimals(&_ERC4626Vault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC4626Vault *ERC4626VaultCallerSession) Decimals() (uint8, error) {
	return _ERC4626Vault.Contract.Decimals(&_ERC4626Vault.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_ERC4626Vault *ERC4626VaultCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_ERC4626Vault *ERC4626VaultSession) GetVersion() (string, error) {
	return _ERC4626Vault.Contract.GetVersion(&_ERC4626Vault.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_ERC4626Vault *ERC4626VaultCallerSession) GetVersion() (string, error) {
	return _ERC4626Vault.Contract.GetVersion(&_ERC4626Vault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) Guardian() (common.Address, error) {
	return _ERC4626Vault.Contract.Guardian(&_ERC4626Vault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) Guardian() (common.Address, error) {
	return _ERC4626Vault.Contract.Guardian(&_ERC4626Vault.CallOpts)
}

// MarketWhitelist is a free data retrieval call binding the contract method 0x1908a2ee.
//
// Solidity: function marketWhitelist(address market) view returns(bool)
func (_ERC4626Vault *ERC4626VaultCaller) MarketWhitelist(opts *bind.CallOpts, market common.Address) (bool, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "marketWhitelist", market)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MarketWhitelist is a free data retrieval call binding the contract method 0x1908a2ee.
//
// Solidity: function marketWhitelist(address market) view returns(bool)
func (_ERC4626Vault *ERC4626VaultSession) MarketWhitelist(market common.Address) (bool, error) {
	return _ERC4626Vault.Contract.MarketWhitelist(&_ERC4626Vault.CallOpts, market)
}

// MarketWhitelist is a free data retrieval call binding the contract method 0x1908a2ee.
//
// Solidity: function marketWhitelist(address market) view returns(bool)
func (_ERC4626Vault *ERC4626VaultCallerSession) MarketWhitelist(market common.Address) (bool, error) {
	return _ERC4626Vault.Contract.MarketWhitelist(&_ERC4626Vault.CallOpts, market)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxDeposit(&_ERC4626Vault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxDeposit(&_ERC4626Vault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxMint(&_ERC4626Vault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxMint(&_ERC4626Vault.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxRedeem(&_ERC4626Vault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxRedeem(&_ERC4626Vault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxWithdraw(&_ERC4626Vault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxWithdraw(&_ERC4626Vault.CallOpts, owner)
}

// MinApy is a free data retrieval call binding the contract method 0x24524a96.
//
// Solidity: function minApy() view returns(uint64)
func (_ERC4626Vault *ERC4626VaultCaller) MinApy(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "minApy")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinApy is a free data retrieval call binding the contract method 0x24524a96.
//
// Solidity: function minApy() view returns(uint64)
func (_ERC4626Vault *ERC4626VaultSession) MinApy() (uint64, error) {
	return _ERC4626Vault.Contract.MinApy(&_ERC4626Vault.CallOpts)
}

// MinApy is a free data retrieval call binding the contract method 0x24524a96.
//
// Solidity: function minApy() view returns(uint64)
func (_ERC4626Vault *ERC4626VaultCallerSession) MinApy() (uint64, error) {
	return _ERC4626Vault.Contract.MinApy(&_ERC4626Vault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC4626Vault *ERC4626VaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC4626Vault *ERC4626VaultSession) Name() (string, error) {
	return _ERC4626Vault.Contract.Name(&_ERC4626Vault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC4626Vault *ERC4626VaultCallerSession) Name() (string, error) {
	return _ERC4626Vault.Contract.Name(&_ERC4626Vault.CallOpts)
}

// OrderMaturity is a free data retrieval call binding the contract method 0xac33207f.
//
// Solidity: function orderMaturity(address order) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) OrderMaturity(opts *bind.CallOpts, order common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "orderMaturity", order)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderMaturity is a free data retrieval call binding the contract method 0xac33207f.
//
// Solidity: function orderMaturity(address order) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) OrderMaturity(order common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.OrderMaturity(&_ERC4626Vault.CallOpts, order)
}

// OrderMaturity is a free data retrieval call binding the contract method 0xac33207f.
//
// Solidity: function orderMaturity(address order) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) OrderMaturity(order common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.OrderMaturity(&_ERC4626Vault.CallOpts, order)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) Owner() (common.Address, error) {
	return _ERC4626Vault.Contract.Owner(&_ERC4626Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) Owner() (common.Address, error) {
	return _ERC4626Vault.Contract.Owner(&_ERC4626Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC4626Vault *ERC4626VaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC4626Vault *ERC4626VaultSession) Paused() (bool, error) {
	return _ERC4626Vault.Contract.Paused(&_ERC4626Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ERC4626Vault *ERC4626VaultCallerSession) Paused() (bool, error) {
	return _ERC4626Vault.Contract.Paused(&_ERC4626Vault.CallOpts)
}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns((address,uint64))
func (_ERC4626Vault *ERC4626VaultCaller) PendingGuardian(opts *bind.CallOpts) (PendingAddress, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "pendingGuardian")

	if err != nil {
		return *new(PendingAddress), err
	}

	out0 := *abi.ConvertType(out[0], new(PendingAddress)).(*PendingAddress)

	return out0, err

}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns((address,uint64))
func (_ERC4626Vault *ERC4626VaultSession) PendingGuardian() (PendingAddress, error) {
	return _ERC4626Vault.Contract.PendingGuardian(&_ERC4626Vault.CallOpts)
}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns((address,uint64))
func (_ERC4626Vault *ERC4626VaultCallerSession) PendingGuardian() (PendingAddress, error) {
	return _ERC4626Vault.Contract.PendingGuardian(&_ERC4626Vault.CallOpts)
}

// PendingMarkets is a free data retrieval call binding the contract method 0x460bf7e1.
//
// Solidity: function pendingMarkets(address market) view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultCaller) PendingMarkets(opts *bind.CallOpts, market common.Address) (PendingUint192, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "pendingMarkets", market)

	if err != nil {
		return *new(PendingUint192), err
	}

	out0 := *abi.ConvertType(out[0], new(PendingUint192)).(*PendingUint192)

	return out0, err

}

// PendingMarkets is a free data retrieval call binding the contract method 0x460bf7e1.
//
// Solidity: function pendingMarkets(address market) view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultSession) PendingMarkets(market common.Address) (PendingUint192, error) {
	return _ERC4626Vault.Contract.PendingMarkets(&_ERC4626Vault.CallOpts, market)
}

// PendingMarkets is a free data retrieval call binding the contract method 0x460bf7e1.
//
// Solidity: function pendingMarkets(address market) view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultCallerSession) PendingMarkets(market common.Address) (PendingUint192, error) {
	return _ERC4626Vault.Contract.PendingMarkets(&_ERC4626Vault.CallOpts, market)
}

// PendingMinApy is a free data retrieval call binding the contract method 0xd209643a.
//
// Solidity: function pendingMinApy() view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultCaller) PendingMinApy(opts *bind.CallOpts) (PendingUint192, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "pendingMinApy")

	if err != nil {
		return *new(PendingUint192), err
	}

	out0 := *abi.ConvertType(out[0], new(PendingUint192)).(*PendingUint192)

	return out0, err

}

// PendingMinApy is a free data retrieval call binding the contract method 0xd209643a.
//
// Solidity: function pendingMinApy() view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultSession) PendingMinApy() (PendingUint192, error) {
	return _ERC4626Vault.Contract.PendingMinApy(&_ERC4626Vault.CallOpts)
}

// PendingMinApy is a free data retrieval call binding the contract method 0xd209643a.
//
// Solidity: function pendingMinApy() view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultCallerSession) PendingMinApy() (PendingUint192, error) {
	return _ERC4626Vault.Contract.PendingMinApy(&_ERC4626Vault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) PendingOwner() (common.Address, error) {
	return _ERC4626Vault.Contract.PendingOwner(&_ERC4626Vault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) PendingOwner() (common.Address, error) {
	return _ERC4626Vault.Contract.PendingOwner(&_ERC4626Vault.CallOpts)
}

// PendingPerformanceFeeRate is a free data retrieval call binding the contract method 0x3f38db41.
//
// Solidity: function pendingPerformanceFeeRate() view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultCaller) PendingPerformanceFeeRate(opts *bind.CallOpts) (PendingUint192, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "pendingPerformanceFeeRate")

	if err != nil {
		return *new(PendingUint192), err
	}

	out0 := *abi.ConvertType(out[0], new(PendingUint192)).(*PendingUint192)

	return out0, err

}

// PendingPerformanceFeeRate is a free data retrieval call binding the contract method 0x3f38db41.
//
// Solidity: function pendingPerformanceFeeRate() view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultSession) PendingPerformanceFeeRate() (PendingUint192, error) {
	return _ERC4626Vault.Contract.PendingPerformanceFeeRate(&_ERC4626Vault.CallOpts)
}

// PendingPerformanceFeeRate is a free data retrieval call binding the contract method 0x3f38db41.
//
// Solidity: function pendingPerformanceFeeRate() view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultCallerSession) PendingPerformanceFeeRate() (PendingUint192, error) {
	return _ERC4626Vault.Contract.PendingPerformanceFeeRate(&_ERC4626Vault.CallOpts)
}

// PendingPool is a free data retrieval call binding the contract method 0xaf9840e5.
//
// Solidity: function pendingPool() view returns((address,uint64))
func (_ERC4626Vault *ERC4626VaultCaller) PendingPool(opts *bind.CallOpts) (PendingAddress, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "pendingPool")

	if err != nil {
		return *new(PendingAddress), err
	}

	out0 := *abi.ConvertType(out[0], new(PendingAddress)).(*PendingAddress)

	return out0, err

}

// PendingPool is a free data retrieval call binding the contract method 0xaf9840e5.
//
// Solidity: function pendingPool() view returns((address,uint64))
func (_ERC4626Vault *ERC4626VaultSession) PendingPool() (PendingAddress, error) {
	return _ERC4626Vault.Contract.PendingPool(&_ERC4626Vault.CallOpts)
}

// PendingPool is a free data retrieval call binding the contract method 0xaf9840e5.
//
// Solidity: function pendingPool() view returns((address,uint64))
func (_ERC4626Vault *ERC4626VaultCallerSession) PendingPool() (PendingAddress, error) {
	return _ERC4626Vault.Contract.PendingPool(&_ERC4626Vault.CallOpts)
}

// PendingPools is a free data retrieval call binding the contract method 0xcbfd8011.
//
// Solidity: function pendingPools() view returns((address,uint64))
func (_ERC4626Vault *ERC4626VaultCaller) PendingPools(opts *bind.CallOpts) (PendingAddress, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "pendingPools")

	if err != nil {
		return *new(PendingAddress), err
	}

	out0 := *abi.ConvertType(out[0], new(PendingAddress)).(*PendingAddress)

	return out0, err

}

// PendingPools is a free data retrieval call binding the contract method 0xcbfd8011.
//
// Solidity: function pendingPools() view returns((address,uint64))
func (_ERC4626Vault *ERC4626VaultSession) PendingPools() (PendingAddress, error) {
	return _ERC4626Vault.Contract.PendingPools(&_ERC4626Vault.CallOpts)
}

// PendingPools is a free data retrieval call binding the contract method 0xcbfd8011.
//
// Solidity: function pendingPools() view returns((address,uint64))
func (_ERC4626Vault *ERC4626VaultCallerSession) PendingPools() (PendingAddress, error) {
	return _ERC4626Vault.Contract.PendingPools(&_ERC4626Vault.CallOpts)
}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultCaller) PendingTimelock(opts *bind.CallOpts) (PendingUint192, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "pendingTimelock")

	if err != nil {
		return *new(PendingUint192), err
	}

	out0 := *abi.ConvertType(out[0], new(PendingUint192)).(*PendingUint192)

	return out0, err

}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultSession) PendingTimelock() (PendingUint192, error) {
	return _ERC4626Vault.Contract.PendingTimelock(&_ERC4626Vault.CallOpts)
}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns((uint192,uint64))
func (_ERC4626Vault *ERC4626VaultCallerSession) PendingTimelock() (PendingUint192, error) {
	return _ERC4626Vault.Contract.PendingTimelock(&_ERC4626Vault.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) PerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "performanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) PerformanceFee() (*big.Int, error) {
	return _ERC4626Vault.Contract.PerformanceFee(&_ERC4626Vault.CallOpts)
}

// PerformanceFee is a free data retrieval call binding the contract method 0x87788782.
//
// Solidity: function performanceFee() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) PerformanceFee() (*big.Int, error) {
	return _ERC4626Vault.Contract.PerformanceFee(&_ERC4626Vault.CallOpts)
}

// PerformanceFeeRate is a free data retrieval call binding the contract method 0x0ffbfda4.
//
// Solidity: function performanceFeeRate() view returns(uint64)
func (_ERC4626Vault *ERC4626VaultCaller) PerformanceFeeRate(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "performanceFeeRate")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PerformanceFeeRate is a free data retrieval call binding the contract method 0x0ffbfda4.
//
// Solidity: function performanceFeeRate() view returns(uint64)
func (_ERC4626Vault *ERC4626VaultSession) PerformanceFeeRate() (uint64, error) {
	return _ERC4626Vault.Contract.PerformanceFeeRate(&_ERC4626Vault.CallOpts)
}

// PerformanceFeeRate is a free data retrieval call binding the contract method 0x0ffbfda4.
//
// Solidity: function performanceFeeRate() view returns(uint64)
func (_ERC4626Vault *ERC4626VaultCallerSession) PerformanceFeeRate() (uint64, error) {
	return _ERC4626Vault.Contract.PerformanceFeeRate(&_ERC4626Vault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) Pool() (common.Address, error) {
	return _ERC4626Vault.Contract.Pool(&_ERC4626Vault.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) Pool() (common.Address, error) {
	return _ERC4626Vault.Contract.Pool(&_ERC4626Vault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewDeposit(&_ERC4626Vault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewDeposit(&_ERC4626Vault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewMint(&_ERC4626Vault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewMint(&_ERC4626Vault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewRedeem(&_ERC4626Vault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewRedeem(&_ERC4626Vault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewWithdraw(&_ERC4626Vault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewWithdraw(&_ERC4626Vault.CallOpts, assets)
}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) SupplyQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "supplyQueue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) SupplyQueue(arg0 *big.Int) (common.Address, error) {
	return _ERC4626Vault.Contract.SupplyQueue(&_ERC4626Vault.CallOpts, arg0)
}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) SupplyQueue(arg0 *big.Int) (common.Address, error) {
	return _ERC4626Vault.Contract.SupplyQueue(&_ERC4626Vault.CallOpts, arg0)
}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) SupplyQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "supplyQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) SupplyQueueLength() (*big.Int, error) {
	return _ERC4626Vault.Contract.SupplyQueueLength(&_ERC4626Vault.CallOpts)
}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) SupplyQueueLength() (*big.Int, error) {
	return _ERC4626Vault.Contract.SupplyQueueLength(&_ERC4626Vault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC4626Vault *ERC4626VaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC4626Vault *ERC4626VaultSession) Symbol() (string, error) {
	return _ERC4626Vault.Contract.Symbol(&_ERC4626Vault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC4626Vault *ERC4626VaultCallerSession) Symbol() (string, error) {
	return _ERC4626Vault.Contract.Symbol(&_ERC4626Vault.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) Timelock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Timelock() (*big.Int, error) {
	return _ERC4626Vault.Contract.Timelock(&_ERC4626Vault.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) Timelock() (*big.Int, error) {
	return _ERC4626Vault.Contract.Timelock(&_ERC4626Vault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) TotalAssets() (*big.Int, error) {
	return _ERC4626Vault.Contract.TotalAssets(&_ERC4626Vault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) TotalAssets() (*big.Int, error) {
	return _ERC4626Vault.Contract.TotalAssets(&_ERC4626Vault.CallOpts)
}

// TotalFt is a free data retrieval call binding the contract method 0x69c42125.
//
// Solidity: function totalFt() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) TotalFt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "totalFt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFt is a free data retrieval call binding the contract method 0x69c42125.
//
// Solidity: function totalFt() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) TotalFt() (*big.Int, error) {
	return _ERC4626Vault.Contract.TotalFt(&_ERC4626Vault.CallOpts)
}

// TotalFt is a free data retrieval call binding the contract method 0x69c42125.
//
// Solidity: function totalFt() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) TotalFt() (*big.Int, error) {
	return _ERC4626Vault.Contract.TotalFt(&_ERC4626Vault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) TotalSupply() (*big.Int, error) {
	return _ERC4626Vault.Contract.TotalSupply(&_ERC4626Vault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC4626Vault.Contract.TotalSupply(&_ERC4626Vault.CallOpts)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) WithdrawQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "withdrawQueue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) WithdrawQueue(arg0 *big.Int) (common.Address, error) {
	return _ERC4626Vault.Contract.WithdrawQueue(&_ERC4626Vault.CallOpts, arg0)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) WithdrawQueue(arg0 *big.Int) (common.Address, error) {
	return _ERC4626Vault.Contract.WithdrawQueue(&_ERC4626Vault.CallOpts, arg0)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) WithdrawQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "withdrawQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) WithdrawQueueLength() (*big.Int, error) {
	return _ERC4626Vault.Contract.WithdrawQueueLength(&_ERC4626Vault.CallOpts)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) WithdrawQueueLength() (*big.Int, error) {
	return _ERC4626Vault.Contract.WithdrawQueueLength(&_ERC4626Vault.CallOpts)
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) AcceptGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "acceptGuardian")
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_ERC4626Vault *ERC4626VaultSession) AcceptGuardian() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptGuardian(&_ERC4626Vault.TransactOpts)
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) AcceptGuardian() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptGuardian(&_ERC4626Vault.TransactOpts)
}

// AcceptMarket is a paid mutator transaction binding the contract method 0xf6b6a7a4.
//
// Solidity: function acceptMarket(address market) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) AcceptMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "acceptMarket", market)
}

// AcceptMarket is a paid mutator transaction binding the contract method 0xf6b6a7a4.
//
// Solidity: function acceptMarket(address market) returns()
func (_ERC4626Vault *ERC4626VaultSession) AcceptMarket(market common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptMarket(&_ERC4626Vault.TransactOpts, market)
}

// AcceptMarket is a paid mutator transaction binding the contract method 0xf6b6a7a4.
//
// Solidity: function acceptMarket(address market) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) AcceptMarket(market common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptMarket(&_ERC4626Vault.TransactOpts, market)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ERC4626Vault *ERC4626VaultSession) AcceptOwnership() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptOwnership(&_ERC4626Vault.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptOwnership(&_ERC4626Vault.TransactOpts)
}

// AcceptPendingMinApy is a paid mutator transaction binding the contract method 0x163c8c77.
//
// Solidity: function acceptPendingMinApy() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) AcceptPendingMinApy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "acceptPendingMinApy")
}

// AcceptPendingMinApy is a paid mutator transaction binding the contract method 0x163c8c77.
//
// Solidity: function acceptPendingMinApy() returns()
func (_ERC4626Vault *ERC4626VaultSession) AcceptPendingMinApy() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptPendingMinApy(&_ERC4626Vault.TransactOpts)
}

// AcceptPendingMinApy is a paid mutator transaction binding the contract method 0x163c8c77.
//
// Solidity: function acceptPendingMinApy() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) AcceptPendingMinApy() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptPendingMinApy(&_ERC4626Vault.TransactOpts)
}

// AcceptPerformanceFeeRate is a paid mutator transaction binding the contract method 0xcb3cab67.
//
// Solidity: function acceptPerformanceFeeRate() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) AcceptPerformanceFeeRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "acceptPerformanceFeeRate")
}

// AcceptPerformanceFeeRate is a paid mutator transaction binding the contract method 0xcb3cab67.
//
// Solidity: function acceptPerformanceFeeRate() returns()
func (_ERC4626Vault *ERC4626VaultSession) AcceptPerformanceFeeRate() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptPerformanceFeeRate(&_ERC4626Vault.TransactOpts)
}

// AcceptPerformanceFeeRate is a paid mutator transaction binding the contract method 0xcb3cab67.
//
// Solidity: function acceptPerformanceFeeRate() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) AcceptPerformanceFeeRate() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptPerformanceFeeRate(&_ERC4626Vault.TransactOpts)
}

// AcceptPool is a paid mutator transaction binding the contract method 0xb0a4709d.
//
// Solidity: function acceptPool() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) AcceptPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "acceptPool")
}

// AcceptPool is a paid mutator transaction binding the contract method 0xb0a4709d.
//
// Solidity: function acceptPool() returns()
func (_ERC4626Vault *ERC4626VaultSession) AcceptPool() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptPool(&_ERC4626Vault.TransactOpts)
}

// AcceptPool is a paid mutator transaction binding the contract method 0xb0a4709d.
//
// Solidity: function acceptPool() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) AcceptPool() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptPool(&_ERC4626Vault.TransactOpts)
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) AcceptTimelock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "acceptTimelock")
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_ERC4626Vault *ERC4626VaultSession) AcceptTimelock() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptTimelock(&_ERC4626Vault.TransactOpts)
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) AcceptTimelock() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AcceptTimelock(&_ERC4626Vault.TransactOpts)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xda3f224a.
//
// Solidity: function afterSwap(uint256 ftReserve, uint256 xtReserve, int256 deltaFt, int256 deltaXt) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) AfterSwap(opts *bind.TransactOpts, ftReserve *big.Int, xtReserve *big.Int, deltaFt *big.Int, deltaXt *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "afterSwap", ftReserve, xtReserve, deltaFt, deltaXt)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xda3f224a.
//
// Solidity: function afterSwap(uint256 ftReserve, uint256 xtReserve, int256 deltaFt, int256 deltaXt) returns()
func (_ERC4626Vault *ERC4626VaultSession) AfterSwap(ftReserve *big.Int, xtReserve *big.Int, deltaFt *big.Int, deltaXt *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AfterSwap(&_ERC4626Vault.TransactOpts, ftReserve, xtReserve, deltaFt, deltaXt)
}

// AfterSwap is a paid mutator transaction binding the contract method 0xda3f224a.
//
// Solidity: function afterSwap(uint256 ftReserve, uint256 xtReserve, int256 deltaFt, int256 deltaXt) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) AfterSwap(ftReserve *big.Int, xtReserve *big.Int, deltaFt *big.Int, deltaXt *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.AfterSwap(&_ERC4626Vault.TransactOpts, ftReserve, xtReserve, deltaFt, deltaXt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC4626Vault *ERC4626VaultSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Approve(&_ERC4626Vault.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Approve(&_ERC4626Vault.TransactOpts, spender, value)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xa7ebda97.
//
// Solidity: function createOrder(address market, (uint256,uint256,uint256,((uint256,uint256,int256)[],(uint256,uint256,int256)[])) params) returns(address order)
func (_ERC4626Vault *ERC4626VaultTransactor) CreateOrder(opts *bind.TransactOpts, market common.Address, params OrderV2ConfigurationParams) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "createOrder", market, params)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xa7ebda97.
//
// Solidity: function createOrder(address market, (uint256,uint256,uint256,((uint256,uint256,int256)[],(uint256,uint256,int256)[])) params) returns(address order)
func (_ERC4626Vault *ERC4626VaultSession) CreateOrder(market common.Address, params OrderV2ConfigurationParams) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.CreateOrder(&_ERC4626Vault.TransactOpts, market, params)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xa7ebda97.
//
// Solidity: function createOrder(address market, (uint256,uint256,uint256,((uint256,uint256,int256)[],(uint256,uint256,int256)[])) params) returns(address order)
func (_ERC4626Vault *ERC4626VaultTransactorSession) CreateOrder(market common.Address, params OrderV2ConfigurationParams) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.CreateOrder(&_ERC4626Vault.TransactOpts, market, params)
}

// DealBadDebt is a paid mutator transaction binding the contract method 0x7207fbb4.
//
// Solidity: function dealBadDebt(address collateral, uint256 badDebtAmt, address recipient, address owner) returns(uint256 shares, uint256 collateralOut)
func (_ERC4626Vault *ERC4626VaultTransactor) DealBadDebt(opts *bind.TransactOpts, collateral common.Address, badDebtAmt *big.Int, recipient common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "dealBadDebt", collateral, badDebtAmt, recipient, owner)
}

// DealBadDebt is a paid mutator transaction binding the contract method 0x7207fbb4.
//
// Solidity: function dealBadDebt(address collateral, uint256 badDebtAmt, address recipient, address owner) returns(uint256 shares, uint256 collateralOut)
func (_ERC4626Vault *ERC4626VaultSession) DealBadDebt(collateral common.Address, badDebtAmt *big.Int, recipient common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.DealBadDebt(&_ERC4626Vault.TransactOpts, collateral, badDebtAmt, recipient, owner)
}

// DealBadDebt is a paid mutator transaction binding the contract method 0x7207fbb4.
//
// Solidity: function dealBadDebt(address collateral, uint256 badDebtAmt, address recipient, address owner) returns(uint256 shares, uint256 collateralOut)
func (_ERC4626Vault *ERC4626VaultTransactorSession) DealBadDebt(collateral common.Address, badDebtAmt *big.Int, recipient common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.DealBadDebt(&_ERC4626Vault.TransactOpts, collateral, badDebtAmt, recipient, owner)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Deposit(&_ERC4626Vault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Deposit(&_ERC4626Vault.TransactOpts, assets, receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x77e1731b.
//
// Solidity: function initialize((address,address,uint256,address,uint256,string,string,uint64) ) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) Initialize(opts *bind.TransactOpts, arg0 VaultInitialParams) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "initialize", arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x77e1731b.
//
// Solidity: function initialize((address,address,uint256,address,uint256,string,string,uint64) ) returns()
func (_ERC4626Vault *ERC4626VaultSession) Initialize(arg0 VaultInitialParams) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Initialize(&_ERC4626Vault.TransactOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x77e1731b.
//
// Solidity: function initialize((address,address,uint256,address,uint256,string,string,uint64) ) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) Initialize(arg0 VaultInitialParams) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Initialize(&_ERC4626Vault.TransactOpts, arg0)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x9d8b21e9.
//
// Solidity: function initialize((address,address,address,uint256,address,address,uint256,string,string,uint64,uint64) params) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) Initialize0(opts *bind.TransactOpts, params VaultInitialParamsV2) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "initialize0", params)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x9d8b21e9.
//
// Solidity: function initialize((address,address,address,uint256,address,address,uint256,string,string,uint64,uint64) params) returns()
func (_ERC4626Vault *ERC4626VaultSession) Initialize0(params VaultInitialParamsV2) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Initialize0(&_ERC4626Vault.TransactOpts, params)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x9d8b21e9.
//
// Solidity: function initialize((address,address,address,uint256,address,address,uint256,string,string,uint64,uint64) params) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) Initialize0(params VaultInitialParamsV2) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Initialize0(&_ERC4626Vault.TransactOpts, params)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Mint(&_ERC4626Vault.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Mint(&_ERC4626Vault.TransactOpts, shares, receiver)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC4626Vault *ERC4626VaultSession) Pause() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Pause(&_ERC4626Vault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) Pause() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Pause(&_ERC4626Vault.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Redeem(&_ERC4626Vault.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Redeem(&_ERC4626Vault.TransactOpts, shares, receiver, owner)
}

// RedeemOrder is a paid mutator transaction binding the contract method 0x988a64c4.
//
// Solidity: function redeemOrder(address order) returns(uint256 badDebt, uint256 deliveryCollateral)
func (_ERC4626Vault *ERC4626VaultTransactor) RedeemOrder(opts *bind.TransactOpts, order common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "redeemOrder", order)
}

// RedeemOrder is a paid mutator transaction binding the contract method 0x988a64c4.
//
// Solidity: function redeemOrder(address order) returns(uint256 badDebt, uint256 deliveryCollateral)
func (_ERC4626Vault *ERC4626VaultSession) RedeemOrder(order common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RedeemOrder(&_ERC4626Vault.TransactOpts, order)
}

// RedeemOrder is a paid mutator transaction binding the contract method 0x988a64c4.
//
// Solidity: function redeemOrder(address order) returns(uint256 badDebt, uint256 deliveryCollateral)
func (_ERC4626Vault *ERC4626VaultTransactorSession) RedeemOrder(order common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RedeemOrder(&_ERC4626Vault.TransactOpts, order)
}

// RemoveLiquidityFromOrders is a paid mutator transaction binding the contract method 0x9a110973.
//
// Solidity: function removeLiquidityFromOrders(address[] orders, uint256[] removedLiquidities) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) RemoveLiquidityFromOrders(opts *bind.TransactOpts, orders []common.Address, removedLiquidities []*big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "removeLiquidityFromOrders", orders, removedLiquidities)
}

// RemoveLiquidityFromOrders is a paid mutator transaction binding the contract method 0x9a110973.
//
// Solidity: function removeLiquidityFromOrders(address[] orders, uint256[] removedLiquidities) returns()
func (_ERC4626Vault *ERC4626VaultSession) RemoveLiquidityFromOrders(orders []common.Address, removedLiquidities []*big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RemoveLiquidityFromOrders(&_ERC4626Vault.TransactOpts, orders, removedLiquidities)
}

// RemoveLiquidityFromOrders is a paid mutator transaction binding the contract method 0x9a110973.
//
// Solidity: function removeLiquidityFromOrders(address[] orders, uint256[] removedLiquidities) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) RemoveLiquidityFromOrders(orders []common.Address, removedLiquidities []*big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RemoveLiquidityFromOrders(&_ERC4626Vault.TransactOpts, orders, removedLiquidities)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC4626Vault *ERC4626VaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RenounceOwnership(&_ERC4626Vault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RenounceOwnership(&_ERC4626Vault.TransactOpts)
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) RevokePendingGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "revokePendingGuardian")
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_ERC4626Vault *ERC4626VaultSession) RevokePendingGuardian() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingGuardian(&_ERC4626Vault.TransactOpts)
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) RevokePendingGuardian() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingGuardian(&_ERC4626Vault.TransactOpts)
}

// RevokePendingMarket is a paid mutator transaction binding the contract method 0xe34c721b.
//
// Solidity: function revokePendingMarket(address market) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) RevokePendingMarket(opts *bind.TransactOpts, market common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "revokePendingMarket", market)
}

// RevokePendingMarket is a paid mutator transaction binding the contract method 0xe34c721b.
//
// Solidity: function revokePendingMarket(address market) returns()
func (_ERC4626Vault *ERC4626VaultSession) RevokePendingMarket(market common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingMarket(&_ERC4626Vault.TransactOpts, market)
}

// RevokePendingMarket is a paid mutator transaction binding the contract method 0xe34c721b.
//
// Solidity: function revokePendingMarket(address market) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) RevokePendingMarket(market common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingMarket(&_ERC4626Vault.TransactOpts, market)
}

// RevokePendingMinApy is a paid mutator transaction binding the contract method 0x86c058ed.
//
// Solidity: function revokePendingMinApy() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) RevokePendingMinApy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "revokePendingMinApy")
}

// RevokePendingMinApy is a paid mutator transaction binding the contract method 0x86c058ed.
//
// Solidity: function revokePendingMinApy() returns()
func (_ERC4626Vault *ERC4626VaultSession) RevokePendingMinApy() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingMinApy(&_ERC4626Vault.TransactOpts)
}

// RevokePendingMinApy is a paid mutator transaction binding the contract method 0x86c058ed.
//
// Solidity: function revokePendingMinApy() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) RevokePendingMinApy() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingMinApy(&_ERC4626Vault.TransactOpts)
}

// RevokePendingPerformanceFeeRate is a paid mutator transaction binding the contract method 0x80a58a5e.
//
// Solidity: function revokePendingPerformanceFeeRate() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) RevokePendingPerformanceFeeRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "revokePendingPerformanceFeeRate")
}

// RevokePendingPerformanceFeeRate is a paid mutator transaction binding the contract method 0x80a58a5e.
//
// Solidity: function revokePendingPerformanceFeeRate() returns()
func (_ERC4626Vault *ERC4626VaultSession) RevokePendingPerformanceFeeRate() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingPerformanceFeeRate(&_ERC4626Vault.TransactOpts)
}

// RevokePendingPerformanceFeeRate is a paid mutator transaction binding the contract method 0x80a58a5e.
//
// Solidity: function revokePendingPerformanceFeeRate() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) RevokePendingPerformanceFeeRate() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingPerformanceFeeRate(&_ERC4626Vault.TransactOpts)
}

// RevokePendingPool is a paid mutator transaction binding the contract method 0x63f04383.
//
// Solidity: function revokePendingPool() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) RevokePendingPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "revokePendingPool")
}

// RevokePendingPool is a paid mutator transaction binding the contract method 0x63f04383.
//
// Solidity: function revokePendingPool() returns()
func (_ERC4626Vault *ERC4626VaultSession) RevokePendingPool() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingPool(&_ERC4626Vault.TransactOpts)
}

// RevokePendingPool is a paid mutator transaction binding the contract method 0x63f04383.
//
// Solidity: function revokePendingPool() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) RevokePendingPool() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingPool(&_ERC4626Vault.TransactOpts)
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) RevokePendingTimelock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "revokePendingTimelock")
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_ERC4626Vault *ERC4626VaultSession) RevokePendingTimelock() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingTimelock(&_ERC4626Vault.TransactOpts)
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) RevokePendingTimelock() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.RevokePendingTimelock(&_ERC4626Vault.TransactOpts)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x91915ef8.
//
// Solidity: function setCapacity(uint256 newCapacity) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) SetCapacity(opts *bind.TransactOpts, newCapacity *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "setCapacity", newCapacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x91915ef8.
//
// Solidity: function setCapacity(uint256 newCapacity) returns()
func (_ERC4626Vault *ERC4626VaultSession) SetCapacity(newCapacity *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SetCapacity(&_ERC4626Vault.TransactOpts, newCapacity)
}

// SetCapacity is a paid mutator transaction binding the contract method 0x91915ef8.
//
// Solidity: function setCapacity(uint256 newCapacity) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) SetCapacity(newCapacity *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SetCapacity(&_ERC4626Vault.TransactOpts, newCapacity)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) SetCurator(opts *bind.TransactOpts, newCurator common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "setCurator", newCurator)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_ERC4626Vault *ERC4626VaultSession) SetCurator(newCurator common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SetCurator(&_ERC4626Vault.TransactOpts, newCurator)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) SetCurator(newCurator common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SetCurator(&_ERC4626Vault.TransactOpts, newCurator)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) SubmitGuardian(opts *bind.TransactOpts, newGuardian common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "submitGuardian", newGuardian)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_ERC4626Vault *ERC4626VaultSession) SubmitGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitGuardian(&_ERC4626Vault.TransactOpts, newGuardian)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) SubmitGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitGuardian(&_ERC4626Vault.TransactOpts, newGuardian)
}

// SubmitMarket is a paid mutator transaction binding the contract method 0x1124f92c.
//
// Solidity: function submitMarket(address market, bool isWhitelisted) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) SubmitMarket(opts *bind.TransactOpts, market common.Address, isWhitelisted bool) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "submitMarket", market, isWhitelisted)
}

// SubmitMarket is a paid mutator transaction binding the contract method 0x1124f92c.
//
// Solidity: function submitMarket(address market, bool isWhitelisted) returns()
func (_ERC4626Vault *ERC4626VaultSession) SubmitMarket(market common.Address, isWhitelisted bool) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitMarket(&_ERC4626Vault.TransactOpts, market, isWhitelisted)
}

// SubmitMarket is a paid mutator transaction binding the contract method 0x1124f92c.
//
// Solidity: function submitMarket(address market, bool isWhitelisted) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) SubmitMarket(market common.Address, isWhitelisted bool) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitMarket(&_ERC4626Vault.TransactOpts, market, isWhitelisted)
}

// SubmitPendingMinApy is a paid mutator transaction binding the contract method 0x89716a25.
//
// Solidity: function submitPendingMinApy(uint64 newMinApy) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) SubmitPendingMinApy(opts *bind.TransactOpts, newMinApy uint64) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "submitPendingMinApy", newMinApy)
}

// SubmitPendingMinApy is a paid mutator transaction binding the contract method 0x89716a25.
//
// Solidity: function submitPendingMinApy(uint64 newMinApy) returns()
func (_ERC4626Vault *ERC4626VaultSession) SubmitPendingMinApy(newMinApy uint64) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitPendingMinApy(&_ERC4626Vault.TransactOpts, newMinApy)
}

// SubmitPendingMinApy is a paid mutator transaction binding the contract method 0x89716a25.
//
// Solidity: function submitPendingMinApy(uint64 newMinApy) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) SubmitPendingMinApy(newMinApy uint64) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitPendingMinApy(&_ERC4626Vault.TransactOpts, newMinApy)
}

// SubmitPendingPool is a paid mutator transaction binding the contract method 0x45142231.
//
// Solidity: function submitPendingPool(address pool_) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) SubmitPendingPool(opts *bind.TransactOpts, pool_ common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "submitPendingPool", pool_)
}

// SubmitPendingPool is a paid mutator transaction binding the contract method 0x45142231.
//
// Solidity: function submitPendingPool(address pool_) returns()
func (_ERC4626Vault *ERC4626VaultSession) SubmitPendingPool(pool_ common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitPendingPool(&_ERC4626Vault.TransactOpts, pool_)
}

// SubmitPendingPool is a paid mutator transaction binding the contract method 0x45142231.
//
// Solidity: function submitPendingPool(address pool_) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) SubmitPendingPool(pool_ common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitPendingPool(&_ERC4626Vault.TransactOpts, pool_)
}

// SubmitPerformanceFeeRate is a paid mutator transaction binding the contract method 0xa9133f5e.
//
// Solidity: function submitPerformanceFeeRate(uint184 newPerformanceFeeRate) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) SubmitPerformanceFeeRate(opts *bind.TransactOpts, newPerformanceFeeRate *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "submitPerformanceFeeRate", newPerformanceFeeRate)
}

// SubmitPerformanceFeeRate is a paid mutator transaction binding the contract method 0xa9133f5e.
//
// Solidity: function submitPerformanceFeeRate(uint184 newPerformanceFeeRate) returns()
func (_ERC4626Vault *ERC4626VaultSession) SubmitPerformanceFeeRate(newPerformanceFeeRate *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitPerformanceFeeRate(&_ERC4626Vault.TransactOpts, newPerformanceFeeRate)
}

// SubmitPerformanceFeeRate is a paid mutator transaction binding the contract method 0xa9133f5e.
//
// Solidity: function submitPerformanceFeeRate(uint184 newPerformanceFeeRate) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) SubmitPerformanceFeeRate(newPerformanceFeeRate *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitPerformanceFeeRate(&_ERC4626Vault.TransactOpts, newPerformanceFeeRate)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) SubmitTimelock(opts *bind.TransactOpts, newTimelock *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "submitTimelock", newTimelock)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_ERC4626Vault *ERC4626VaultSession) SubmitTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitTimelock(&_ERC4626Vault.TransactOpts, newTimelock)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) SubmitTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.SubmitTimelock(&_ERC4626Vault.TransactOpts, newTimelock)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC4626Vault *ERC4626VaultSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Transfer(&_ERC4626Vault.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Transfer(&_ERC4626Vault.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC4626Vault *ERC4626VaultSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.TransferFrom(&_ERC4626Vault.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.TransferFrom(&_ERC4626Vault.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC4626Vault *ERC4626VaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.TransferOwnership(&_ERC4626Vault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.TransferOwnership(&_ERC4626Vault.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC4626Vault *ERC4626VaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC4626Vault *ERC4626VaultSession) Unpause() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Unpause(&_ERC4626Vault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Unpause(&_ERC4626Vault.TransactOpts)
}

// UpdateOrdersConfiguration is a paid mutator transaction binding the contract method 0x0edc3c64.
//
// Solidity: function updateOrdersConfiguration(address[] orders, (uint256,uint256,uint256,((uint256,uint256,int256)[],(uint256,uint256,int256)[]))[] orderConfigs) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) UpdateOrdersConfiguration(opts *bind.TransactOpts, orders []common.Address, orderConfigs []OrderV2ConfigurationParams) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "updateOrdersConfiguration", orders, orderConfigs)
}

// UpdateOrdersConfiguration is a paid mutator transaction binding the contract method 0x0edc3c64.
//
// Solidity: function updateOrdersConfiguration(address[] orders, (uint256,uint256,uint256,((uint256,uint256,int256)[],(uint256,uint256,int256)[]))[] orderConfigs) returns()
func (_ERC4626Vault *ERC4626VaultSession) UpdateOrdersConfiguration(orders []common.Address, orderConfigs []OrderV2ConfigurationParams) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.UpdateOrdersConfiguration(&_ERC4626Vault.TransactOpts, orders, orderConfigs)
}

// UpdateOrdersConfiguration is a paid mutator transaction binding the contract method 0x0edc3c64.
//
// Solidity: function updateOrdersConfiguration(address[] orders, (uint256,uint256,uint256,((uint256,uint256,int256)[],(uint256,uint256,int256)[]))[] orderConfigs) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) UpdateOrdersConfiguration(orders []common.Address, orderConfigs []OrderV2ConfigurationParams) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.UpdateOrdersConfiguration(&_ERC4626Vault.TransactOpts, orders, orderConfigs)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Withdraw(&_ERC4626Vault.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Withdraw(&_ERC4626Vault.TransactOpts, assets, receiver, owner)
}

// WithdrawFts is a paid mutator transaction binding the contract method 0x94e854f2.
//
// Solidity: function withdrawFts(address order, uint256 amount, address recipient, address owner) returns(uint256 shares)
func (_ERC4626Vault *ERC4626VaultTransactor) WithdrawFts(opts *bind.TransactOpts, order common.Address, amount *big.Int, recipient common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "withdrawFts", order, amount, recipient, owner)
}

// WithdrawFts is a paid mutator transaction binding the contract method 0x94e854f2.
//
// Solidity: function withdrawFts(address order, uint256 amount, address recipient, address owner) returns(uint256 shares)
func (_ERC4626Vault *ERC4626VaultSession) WithdrawFts(order common.Address, amount *big.Int, recipient common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.WithdrawFts(&_ERC4626Vault.TransactOpts, order, amount, recipient, owner)
}

// WithdrawFts is a paid mutator transaction binding the contract method 0x94e854f2.
//
// Solidity: function withdrawFts(address order, uint256 amount, address recipient, address owner) returns(uint256 shares)
func (_ERC4626Vault *ERC4626VaultTransactorSession) WithdrawFts(order common.Address, amount *big.Int, recipient common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.WithdrawFts(&_ERC4626Vault.TransactOpts, order, amount, recipient, owner)
}

// WithdrawPerformanceFee is a paid mutator transaction binding the contract method 0x93bd007b.
//
// Solidity: function withdrawPerformanceFee(address recipient, uint256 amount) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) WithdrawPerformanceFee(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "withdrawPerformanceFee", recipient, amount)
}

// WithdrawPerformanceFee is a paid mutator transaction binding the contract method 0x93bd007b.
//
// Solidity: function withdrawPerformanceFee(address recipient, uint256 amount) returns()
func (_ERC4626Vault *ERC4626VaultSession) WithdrawPerformanceFee(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.WithdrawPerformanceFee(&_ERC4626Vault.TransactOpts, recipient, amount)
}

// WithdrawPerformanceFee is a paid mutator transaction binding the contract method 0x93bd007b.
//
// Solidity: function withdrawPerformanceFee(address recipient, uint256 amount) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) WithdrawPerformanceFee(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.WithdrawPerformanceFee(&_ERC4626Vault.TransactOpts, recipient, amount)
}

// ERC4626VaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC4626Vault contract.
type ERC4626VaultApprovalIterator struct {
	Event *ERC4626VaultApproval // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultApproval)
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
		it.Event = new(ERC4626VaultApproval)
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
func (it *ERC4626VaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultApproval represents a Approval event raised by the ERC4626Vault contract.
type ERC4626VaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC4626VaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultApprovalIterator{contract: _ERC4626Vault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC4626VaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultApproval)
				if err := _ERC4626Vault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC4626Vault *ERC4626VaultFilterer) ParseApproval(log types.Log) (*ERC4626VaultApproval, error) {
	event := new(ERC4626VaultApproval)
	if err := _ERC4626Vault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultDealBadDebtIterator is returned from FilterDealBadDebt and is used to iterate over the raw logs and unpacked data for DealBadDebt events raised by the ERC4626Vault contract.
type ERC4626VaultDealBadDebtIterator struct {
	Event *ERC4626VaultDealBadDebt // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultDealBadDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultDealBadDebt)
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
		it.Event = new(ERC4626VaultDealBadDebt)
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
func (it *ERC4626VaultDealBadDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultDealBadDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultDealBadDebt represents a DealBadDebt event raised by the ERC4626Vault contract.
type ERC4626VaultDealBadDebt struct {
	Caller        common.Address
	Recipient     common.Address
	Collateral    common.Address
	BadDebt       *big.Int
	Shares        *big.Int
	CollateralOut *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDealBadDebt is a free log retrieval operation binding the contract event 0xaf2e30fae2dfd1a90059cf53415e90c4ee9d151c1b1861df8f8a5963069c47f4.
//
// Solidity: event DealBadDebt(address indexed caller, address indexed recipient, address indexed collateral, uint256 badDebt, uint256 shares, uint256 collateralOut)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterDealBadDebt(opts *bind.FilterOpts, caller []common.Address, recipient []common.Address, collateral []common.Address) (*ERC4626VaultDealBadDebtIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "DealBadDebt", callerRule, recipientRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultDealBadDebtIterator{contract: _ERC4626Vault.contract, event: "DealBadDebt", logs: logs, sub: sub}, nil
}

// WatchDealBadDebt is a free log subscription operation binding the contract event 0xaf2e30fae2dfd1a90059cf53415e90c4ee9d151c1b1861df8f8a5963069c47f4.
//
// Solidity: event DealBadDebt(address indexed caller, address indexed recipient, address indexed collateral, uint256 badDebt, uint256 shares, uint256 collateralOut)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchDealBadDebt(opts *bind.WatchOpts, sink chan<- *ERC4626VaultDealBadDebt, caller []common.Address, recipient []common.Address, collateral []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "DealBadDebt", callerRule, recipientRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultDealBadDebt)
				if err := _ERC4626Vault.contract.UnpackLog(event, "DealBadDebt", log); err != nil {
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

// ParseDealBadDebt is a log parse operation binding the contract event 0xaf2e30fae2dfd1a90059cf53415e90c4ee9d151c1b1861df8f8a5963069c47f4.
//
// Solidity: event DealBadDebt(address indexed caller, address indexed recipient, address indexed collateral, uint256 badDebt, uint256 shares, uint256 collateralOut)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseDealBadDebt(log types.Log) (*ERC4626VaultDealBadDebt, error) {
	event := new(ERC4626VaultDealBadDebt)
	if err := _ERC4626Vault.contract.UnpackLog(event, "DealBadDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ERC4626Vault contract.
type ERC4626VaultDepositIterator struct {
	Event *ERC4626VaultDeposit // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultDeposit)
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
		it.Event = new(ERC4626VaultDeposit)
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
func (it *ERC4626VaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultDeposit represents a Deposit event raised by the ERC4626Vault contract.
type ERC4626VaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*ERC4626VaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultDepositIterator{contract: _ERC4626Vault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ERC4626VaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultDeposit)
				if err := _ERC4626Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_ERC4626Vault *ERC4626VaultFilterer) ParseDeposit(log types.Log) (*ERC4626VaultDeposit, error) {
	event := new(ERC4626VaultDeposit)
	if err := _ERC4626Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC4626Vault contract.
type ERC4626VaultInitializedIterator struct {
	Event *ERC4626VaultInitialized // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultInitialized)
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
		it.Event = new(ERC4626VaultInitialized)
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
func (it *ERC4626VaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultInitialized represents a Initialized event raised by the ERC4626Vault contract.
type ERC4626VaultInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC4626VaultInitializedIterator, error) {

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultInitializedIterator{contract: _ERC4626Vault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC4626VaultInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultInitialized)
				if err := _ERC4626Vault.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ERC4626Vault *ERC4626VaultFilterer) ParseInitialized(log types.Log) (*ERC4626VaultInitialized, error) {
	event := new(ERC4626VaultInitialized)
	if err := _ERC4626Vault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the ERC4626Vault contract.
type ERC4626VaultOwnershipTransferStartedIterator struct {
	Event *ERC4626VaultOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultOwnershipTransferStarted)
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
		it.Event = new(ERC4626VaultOwnershipTransferStarted)
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
func (it *ERC4626VaultOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the ERC4626Vault contract.
type ERC4626VaultOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC4626VaultOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultOwnershipTransferStartedIterator{contract: _ERC4626Vault.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *ERC4626VaultOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultOwnershipTransferStarted)
				if err := _ERC4626Vault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseOwnershipTransferStarted(log types.Log) (*ERC4626VaultOwnershipTransferStarted, error) {
	event := new(ERC4626VaultOwnershipTransferStarted)
	if err := _ERC4626Vault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC4626Vault contract.
type ERC4626VaultOwnershipTransferredIterator struct {
	Event *ERC4626VaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultOwnershipTransferred)
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
		it.Event = new(ERC4626VaultOwnershipTransferred)
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
func (it *ERC4626VaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultOwnershipTransferred represents a OwnershipTransferred event raised by the ERC4626Vault contract.
type ERC4626VaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC4626VaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultOwnershipTransferredIterator{contract: _ERC4626Vault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC4626VaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultOwnershipTransferred)
				if err := _ERC4626Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC4626Vault *ERC4626VaultFilterer) ParseOwnershipTransferred(log types.Log) (*ERC4626VaultOwnershipTransferred, error) {
	event := new(ERC4626VaultOwnershipTransferred)
	if err := _ERC4626Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ERC4626Vault contract.
type ERC4626VaultPausedIterator struct {
	Event *ERC4626VaultPaused // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultPaused)
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
		it.Event = new(ERC4626VaultPaused)
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
func (it *ERC4626VaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultPaused represents a Paused event raised by the ERC4626Vault contract.
type ERC4626VaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterPaused(opts *bind.FilterOpts) (*ERC4626VaultPausedIterator, error) {

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultPausedIterator{contract: _ERC4626Vault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ERC4626VaultPaused) (event.Subscription, error) {

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultPaused)
				if err := _ERC4626Vault.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ERC4626Vault *ERC4626VaultFilterer) ParsePaused(log types.Log) (*ERC4626VaultPaused, error) {
	event := new(ERC4626VaultPaused)
	if err := _ERC4626Vault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultRevokePendingGuardianIterator is returned from FilterRevokePendingGuardian and is used to iterate over the raw logs and unpacked data for RevokePendingGuardian events raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingGuardianIterator struct {
	Event *ERC4626VaultRevokePendingGuardian // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultRevokePendingGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultRevokePendingGuardian)
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
		it.Event = new(ERC4626VaultRevokePendingGuardian)
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
func (it *ERC4626VaultRevokePendingGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultRevokePendingGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultRevokePendingGuardian represents a RevokePendingGuardian event raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingGuardian struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingGuardian is a free log retrieval operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterRevokePendingGuardian(opts *bind.FilterOpts, caller []common.Address) (*ERC4626VaultRevokePendingGuardianIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "RevokePendingGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultRevokePendingGuardianIterator{contract: _ERC4626Vault.contract, event: "RevokePendingGuardian", logs: logs, sub: sub}, nil
}

// WatchRevokePendingGuardian is a free log subscription operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchRevokePendingGuardian(opts *bind.WatchOpts, sink chan<- *ERC4626VaultRevokePendingGuardian, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "RevokePendingGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultRevokePendingGuardian)
				if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingGuardian", log); err != nil {
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

// ParseRevokePendingGuardian is a log parse operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseRevokePendingGuardian(log types.Log) (*ERC4626VaultRevokePendingGuardian, error) {
	event := new(ERC4626VaultRevokePendingGuardian)
	if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultRevokePendingMarketIterator is returned from FilterRevokePendingMarket and is used to iterate over the raw logs and unpacked data for RevokePendingMarket events raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingMarketIterator struct {
	Event *ERC4626VaultRevokePendingMarket // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultRevokePendingMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultRevokePendingMarket)
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
		it.Event = new(ERC4626VaultRevokePendingMarket)
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
func (it *ERC4626VaultRevokePendingMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultRevokePendingMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultRevokePendingMarket represents a RevokePendingMarket event raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingMarket struct {
	Caller common.Address
	Market common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingMarket is a free log retrieval operation binding the contract event 0x10d2cd24a2375b12d91635518e47506f9aebfe8af364c6109b93ac41e8b0b86f.
//
// Solidity: event RevokePendingMarket(address indexed caller, address indexed market)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterRevokePendingMarket(opts *bind.FilterOpts, caller []common.Address, market []common.Address) (*ERC4626VaultRevokePendingMarketIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "RevokePendingMarket", callerRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultRevokePendingMarketIterator{contract: _ERC4626Vault.contract, event: "RevokePendingMarket", logs: logs, sub: sub}, nil
}

// WatchRevokePendingMarket is a free log subscription operation binding the contract event 0x10d2cd24a2375b12d91635518e47506f9aebfe8af364c6109b93ac41e8b0b86f.
//
// Solidity: event RevokePendingMarket(address indexed caller, address indexed market)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchRevokePendingMarket(opts *bind.WatchOpts, sink chan<- *ERC4626VaultRevokePendingMarket, caller []common.Address, market []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "RevokePendingMarket", callerRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultRevokePendingMarket)
				if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingMarket", log); err != nil {
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

// ParseRevokePendingMarket is a log parse operation binding the contract event 0x10d2cd24a2375b12d91635518e47506f9aebfe8af364c6109b93ac41e8b0b86f.
//
// Solidity: event RevokePendingMarket(address indexed caller, address indexed market)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseRevokePendingMarket(log types.Log) (*ERC4626VaultRevokePendingMarket, error) {
	event := new(ERC4626VaultRevokePendingMarket)
	if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultRevokePendingMinApyIterator is returned from FilterRevokePendingMinApy and is used to iterate over the raw logs and unpacked data for RevokePendingMinApy events raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingMinApyIterator struct {
	Event *ERC4626VaultRevokePendingMinApy // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultRevokePendingMinApyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultRevokePendingMinApy)
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
		it.Event = new(ERC4626VaultRevokePendingMinApy)
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
func (it *ERC4626VaultRevokePendingMinApyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultRevokePendingMinApyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultRevokePendingMinApy represents a RevokePendingMinApy event raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingMinApy struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingMinApy is a free log retrieval operation binding the contract event 0x6e2daa1b3bca5239ed02028dbe091b3e5d41ae2702e4394b0a28cee04bd59fb2.
//
// Solidity: event RevokePendingMinApy(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterRevokePendingMinApy(opts *bind.FilterOpts, caller []common.Address) (*ERC4626VaultRevokePendingMinApyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "RevokePendingMinApy", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultRevokePendingMinApyIterator{contract: _ERC4626Vault.contract, event: "RevokePendingMinApy", logs: logs, sub: sub}, nil
}

// WatchRevokePendingMinApy is a free log subscription operation binding the contract event 0x6e2daa1b3bca5239ed02028dbe091b3e5d41ae2702e4394b0a28cee04bd59fb2.
//
// Solidity: event RevokePendingMinApy(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchRevokePendingMinApy(opts *bind.WatchOpts, sink chan<- *ERC4626VaultRevokePendingMinApy, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "RevokePendingMinApy", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultRevokePendingMinApy)
				if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingMinApy", log); err != nil {
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

// ParseRevokePendingMinApy is a log parse operation binding the contract event 0x6e2daa1b3bca5239ed02028dbe091b3e5d41ae2702e4394b0a28cee04bd59fb2.
//
// Solidity: event RevokePendingMinApy(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseRevokePendingMinApy(log types.Log) (*ERC4626VaultRevokePendingMinApy, error) {
	event := new(ERC4626VaultRevokePendingMinApy)
	if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingMinApy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultRevokePendingPerformanceFeeRateIterator is returned from FilterRevokePendingPerformanceFeeRate and is used to iterate over the raw logs and unpacked data for RevokePendingPerformanceFeeRate events raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingPerformanceFeeRateIterator struct {
	Event *ERC4626VaultRevokePendingPerformanceFeeRate // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultRevokePendingPerformanceFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultRevokePendingPerformanceFeeRate)
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
		it.Event = new(ERC4626VaultRevokePendingPerformanceFeeRate)
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
func (it *ERC4626VaultRevokePendingPerformanceFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultRevokePendingPerformanceFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultRevokePendingPerformanceFeeRate represents a RevokePendingPerformanceFeeRate event raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingPerformanceFeeRate struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingPerformanceFeeRate is a free log retrieval operation binding the contract event 0xb7cc052ba5dd4e2c200f771e69c6c4dddd930567bd8bea00e527944967c57db8.
//
// Solidity: event RevokePendingPerformanceFeeRate(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterRevokePendingPerformanceFeeRate(opts *bind.FilterOpts, caller []common.Address) (*ERC4626VaultRevokePendingPerformanceFeeRateIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "RevokePendingPerformanceFeeRate", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultRevokePendingPerformanceFeeRateIterator{contract: _ERC4626Vault.contract, event: "RevokePendingPerformanceFeeRate", logs: logs, sub: sub}, nil
}

// WatchRevokePendingPerformanceFeeRate is a free log subscription operation binding the contract event 0xb7cc052ba5dd4e2c200f771e69c6c4dddd930567bd8bea00e527944967c57db8.
//
// Solidity: event RevokePendingPerformanceFeeRate(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchRevokePendingPerformanceFeeRate(opts *bind.WatchOpts, sink chan<- *ERC4626VaultRevokePendingPerformanceFeeRate, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "RevokePendingPerformanceFeeRate", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultRevokePendingPerformanceFeeRate)
				if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingPerformanceFeeRate", log); err != nil {
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

// ParseRevokePendingPerformanceFeeRate is a log parse operation binding the contract event 0xb7cc052ba5dd4e2c200f771e69c6c4dddd930567bd8bea00e527944967c57db8.
//
// Solidity: event RevokePendingPerformanceFeeRate(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseRevokePendingPerformanceFeeRate(log types.Log) (*ERC4626VaultRevokePendingPerformanceFeeRate, error) {
	event := new(ERC4626VaultRevokePendingPerformanceFeeRate)
	if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingPerformanceFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultRevokePendingPoolIterator is returned from FilterRevokePendingPool and is used to iterate over the raw logs and unpacked data for RevokePendingPool events raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingPoolIterator struct {
	Event *ERC4626VaultRevokePendingPool // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultRevokePendingPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultRevokePendingPool)
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
		it.Event = new(ERC4626VaultRevokePendingPool)
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
func (it *ERC4626VaultRevokePendingPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultRevokePendingPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultRevokePendingPool represents a RevokePendingPool event raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingPool struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingPool is a free log retrieval operation binding the contract event 0xa2364ce3a039aaebd156a368d9ee05971530a28068e2304d9a5c0d0269ef002e.
//
// Solidity: event RevokePendingPool(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterRevokePendingPool(opts *bind.FilterOpts, caller []common.Address) (*ERC4626VaultRevokePendingPoolIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "RevokePendingPool", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultRevokePendingPoolIterator{contract: _ERC4626Vault.contract, event: "RevokePendingPool", logs: logs, sub: sub}, nil
}

// WatchRevokePendingPool is a free log subscription operation binding the contract event 0xa2364ce3a039aaebd156a368d9ee05971530a28068e2304d9a5c0d0269ef002e.
//
// Solidity: event RevokePendingPool(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchRevokePendingPool(opts *bind.WatchOpts, sink chan<- *ERC4626VaultRevokePendingPool, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "RevokePendingPool", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultRevokePendingPool)
				if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingPool", log); err != nil {
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

// ParseRevokePendingPool is a log parse operation binding the contract event 0xa2364ce3a039aaebd156a368d9ee05971530a28068e2304d9a5c0d0269ef002e.
//
// Solidity: event RevokePendingPool(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseRevokePendingPool(log types.Log) (*ERC4626VaultRevokePendingPool, error) {
	event := new(ERC4626VaultRevokePendingPool)
	if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultRevokePendingTimelockIterator is returned from FilterRevokePendingTimelock and is used to iterate over the raw logs and unpacked data for RevokePendingTimelock events raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingTimelockIterator struct {
	Event *ERC4626VaultRevokePendingTimelock // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultRevokePendingTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultRevokePendingTimelock)
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
		it.Event = new(ERC4626VaultRevokePendingTimelock)
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
func (it *ERC4626VaultRevokePendingTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultRevokePendingTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultRevokePendingTimelock represents a RevokePendingTimelock event raised by the ERC4626Vault contract.
type ERC4626VaultRevokePendingTimelock struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingTimelock is a free log retrieval operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterRevokePendingTimelock(opts *bind.FilterOpts, caller []common.Address) (*ERC4626VaultRevokePendingTimelockIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "RevokePendingTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultRevokePendingTimelockIterator{contract: _ERC4626Vault.contract, event: "RevokePendingTimelock", logs: logs, sub: sub}, nil
}

// WatchRevokePendingTimelock is a free log subscription operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchRevokePendingTimelock(opts *bind.WatchOpts, sink chan<- *ERC4626VaultRevokePendingTimelock, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "RevokePendingTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultRevokePendingTimelock)
				if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingTimelock", log); err != nil {
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

// ParseRevokePendingTimelock is a log parse operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseRevokePendingTimelock(log types.Log) (*ERC4626VaultRevokePendingTimelock, error) {
	event := new(ERC4626VaultRevokePendingTimelock)
	if err := _ERC4626Vault.contract.UnpackLog(event, "RevokePendingTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSetCapacityIterator is returned from FilterSetCapacity and is used to iterate over the raw logs and unpacked data for SetCapacity events raised by the ERC4626Vault contract.
type ERC4626VaultSetCapacityIterator struct {
	Event *ERC4626VaultSetCapacity // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSetCapacityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSetCapacity)
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
		it.Event = new(ERC4626VaultSetCapacity)
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
func (it *ERC4626VaultSetCapacityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSetCapacityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSetCapacity represents a SetCapacity event raised by the ERC4626Vault contract.
type ERC4626VaultSetCapacity struct {
	Caller      common.Address
	NewCapacity *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetCapacity is a free log retrieval operation binding the contract event 0x51fa58fd85e72d533eb3933ef4dd0bb83a9614f46e076ffa025b0dcbb8dff315.
//
// Solidity: event SetCapacity(address indexed caller, uint256 newCapacity)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSetCapacity(opts *bind.FilterOpts, caller []common.Address) (*ERC4626VaultSetCapacityIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SetCapacity", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSetCapacityIterator{contract: _ERC4626Vault.contract, event: "SetCapacity", logs: logs, sub: sub}, nil
}

// WatchSetCapacity is a free log subscription operation binding the contract event 0x51fa58fd85e72d533eb3933ef4dd0bb83a9614f46e076ffa025b0dcbb8dff315.
//
// Solidity: event SetCapacity(address indexed caller, uint256 newCapacity)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSetCapacity(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSetCapacity, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SetCapacity", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSetCapacity)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SetCapacity", log); err != nil {
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

// ParseSetCapacity is a log parse operation binding the contract event 0x51fa58fd85e72d533eb3933ef4dd0bb83a9614f46e076ffa025b0dcbb8dff315.
//
// Solidity: event SetCapacity(address indexed caller, uint256 newCapacity)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSetCapacity(log types.Log) (*ERC4626VaultSetCapacity, error) {
	event := new(ERC4626VaultSetCapacity)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SetCapacity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSetCuratorIterator is returned from FilterSetCurator and is used to iterate over the raw logs and unpacked data for SetCurator events raised by the ERC4626Vault contract.
type ERC4626VaultSetCuratorIterator struct {
	Event *ERC4626VaultSetCurator // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSetCuratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSetCurator)
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
		it.Event = new(ERC4626VaultSetCurator)
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
func (it *ERC4626VaultSetCuratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSetCuratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSetCurator represents a SetCurator event raised by the ERC4626Vault contract.
type ERC4626VaultSetCurator struct {
	NewCurator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetCurator is a free log retrieval operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address newCurator)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSetCurator(opts *bind.FilterOpts) (*ERC4626VaultSetCuratorIterator, error) {

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SetCurator")
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSetCuratorIterator{contract: _ERC4626Vault.contract, event: "SetCurator", logs: logs, sub: sub}, nil
}

// WatchSetCurator is a free log subscription operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address newCurator)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSetCurator(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSetCurator) (event.Subscription, error) {

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SetCurator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSetCurator)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SetCurator", log); err != nil {
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

// ParseSetCurator is a log parse operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address newCurator)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSetCurator(log types.Log) (*ERC4626VaultSetCurator, error) {
	event := new(ERC4626VaultSetCurator)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SetCurator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSetGuardianIterator is returned from FilterSetGuardian and is used to iterate over the raw logs and unpacked data for SetGuardian events raised by the ERC4626Vault contract.
type ERC4626VaultSetGuardianIterator struct {
	Event *ERC4626VaultSetGuardian // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSetGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSetGuardian)
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
		it.Event = new(ERC4626VaultSetGuardian)
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
func (it *ERC4626VaultSetGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSetGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSetGuardian represents a SetGuardian event raised by the ERC4626Vault contract.
type ERC4626VaultSetGuardian struct {
	Caller      common.Address
	NewGuardian common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetGuardian is a free log retrieval operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address newGuardian)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSetGuardian(opts *bind.FilterOpts, caller []common.Address) (*ERC4626VaultSetGuardianIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SetGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSetGuardianIterator{contract: _ERC4626Vault.contract, event: "SetGuardian", logs: logs, sub: sub}, nil
}

// WatchSetGuardian is a free log subscription operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address newGuardian)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSetGuardian(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSetGuardian, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SetGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSetGuardian)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SetGuardian", log); err != nil {
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

// ParseSetGuardian is a log parse operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address newGuardian)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSetGuardian(log types.Log) (*ERC4626VaultSetGuardian, error) {
	event := new(ERC4626VaultSetGuardian)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SetGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSetMarketWhitelistIterator is returned from FilterSetMarketWhitelist and is used to iterate over the raw logs and unpacked data for SetMarketWhitelist events raised by the ERC4626Vault contract.
type ERC4626VaultSetMarketWhitelistIterator struct {
	Event *ERC4626VaultSetMarketWhitelist // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSetMarketWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSetMarketWhitelist)
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
		it.Event = new(ERC4626VaultSetMarketWhitelist)
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
func (it *ERC4626VaultSetMarketWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSetMarketWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSetMarketWhitelist represents a SetMarketWhitelist event raised by the ERC4626Vault contract.
type ERC4626VaultSetMarketWhitelist struct {
	Caller        common.Address
	Market        common.Address
	IsWhitelisted bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetMarketWhitelist is a free log retrieval operation binding the contract event 0x7276f1d77fbb1794919362294fb870305193d894417c47c88716aa858f4272d3.
//
// Solidity: event SetMarketWhitelist(address indexed caller, address indexed market, bool isWhitelisted)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSetMarketWhitelist(opts *bind.FilterOpts, caller []common.Address, market []common.Address) (*ERC4626VaultSetMarketWhitelistIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SetMarketWhitelist", callerRule, marketRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSetMarketWhitelistIterator{contract: _ERC4626Vault.contract, event: "SetMarketWhitelist", logs: logs, sub: sub}, nil
}

// WatchSetMarketWhitelist is a free log subscription operation binding the contract event 0x7276f1d77fbb1794919362294fb870305193d894417c47c88716aa858f4272d3.
//
// Solidity: event SetMarketWhitelist(address indexed caller, address indexed market, bool isWhitelisted)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSetMarketWhitelist(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSetMarketWhitelist, caller []common.Address, market []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SetMarketWhitelist", callerRule, marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSetMarketWhitelist)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SetMarketWhitelist", log); err != nil {
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

// ParseSetMarketWhitelist is a log parse operation binding the contract event 0x7276f1d77fbb1794919362294fb870305193d894417c47c88716aa858f4272d3.
//
// Solidity: event SetMarketWhitelist(address indexed caller, address indexed market, bool isWhitelisted)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSetMarketWhitelist(log types.Log) (*ERC4626VaultSetMarketWhitelist, error) {
	event := new(ERC4626VaultSetMarketWhitelist)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SetMarketWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSetMinApyIterator is returned from FilterSetMinApy and is used to iterate over the raw logs and unpacked data for SetMinApy events raised by the ERC4626Vault contract.
type ERC4626VaultSetMinApyIterator struct {
	Event *ERC4626VaultSetMinApy // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSetMinApyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSetMinApy)
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
		it.Event = new(ERC4626VaultSetMinApy)
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
func (it *ERC4626VaultSetMinApyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSetMinApyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSetMinApy represents a SetMinApy event raised by the ERC4626Vault contract.
type ERC4626VaultSetMinApy struct {
	Caller    common.Address
	NewMinApy uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetMinApy is a free log retrieval operation binding the contract event 0x0cdc8c1df674ca6c4f86f6d47b02a15083c5d3bf098ea38f02c2762bad0afad2.
//
// Solidity: event SetMinApy(address indexed caller, uint64 newMinApy)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSetMinApy(opts *bind.FilterOpts, caller []common.Address) (*ERC4626VaultSetMinApyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SetMinApy", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSetMinApyIterator{contract: _ERC4626Vault.contract, event: "SetMinApy", logs: logs, sub: sub}, nil
}

// WatchSetMinApy is a free log subscription operation binding the contract event 0x0cdc8c1df674ca6c4f86f6d47b02a15083c5d3bf098ea38f02c2762bad0afad2.
//
// Solidity: event SetMinApy(address indexed caller, uint64 newMinApy)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSetMinApy(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSetMinApy, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SetMinApy", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSetMinApy)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SetMinApy", log); err != nil {
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

// ParseSetMinApy is a log parse operation binding the contract event 0x0cdc8c1df674ca6c4f86f6d47b02a15083c5d3bf098ea38f02c2762bad0afad2.
//
// Solidity: event SetMinApy(address indexed caller, uint64 newMinApy)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSetMinApy(log types.Log) (*ERC4626VaultSetMinApy, error) {
	event := new(ERC4626VaultSetMinApy)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SetMinApy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSetPerformanceFeeRateIterator is returned from FilterSetPerformanceFeeRate and is used to iterate over the raw logs and unpacked data for SetPerformanceFeeRate events raised by the ERC4626Vault contract.
type ERC4626VaultSetPerformanceFeeRateIterator struct {
	Event *ERC4626VaultSetPerformanceFeeRate // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSetPerformanceFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSetPerformanceFeeRate)
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
		it.Event = new(ERC4626VaultSetPerformanceFeeRate)
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
func (it *ERC4626VaultSetPerformanceFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSetPerformanceFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSetPerformanceFeeRate represents a SetPerformanceFeeRate event raised by the ERC4626Vault contract.
type ERC4626VaultSetPerformanceFeeRate struct {
	Caller                common.Address
	NewPerformanceFeeRate *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSetPerformanceFeeRate is a free log retrieval operation binding the contract event 0x20affe2401825617c69366f8c3a3493d9822d1021d0b3023c4e77ea5b3d0fbc5.
//
// Solidity: event SetPerformanceFeeRate(address indexed caller, uint256 newPerformanceFeeRate)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSetPerformanceFeeRate(opts *bind.FilterOpts, caller []common.Address) (*ERC4626VaultSetPerformanceFeeRateIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SetPerformanceFeeRate", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSetPerformanceFeeRateIterator{contract: _ERC4626Vault.contract, event: "SetPerformanceFeeRate", logs: logs, sub: sub}, nil
}

// WatchSetPerformanceFeeRate is a free log subscription operation binding the contract event 0x20affe2401825617c69366f8c3a3493d9822d1021d0b3023c4e77ea5b3d0fbc5.
//
// Solidity: event SetPerformanceFeeRate(address indexed caller, uint256 newPerformanceFeeRate)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSetPerformanceFeeRate(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSetPerformanceFeeRate, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SetPerformanceFeeRate", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSetPerformanceFeeRate)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SetPerformanceFeeRate", log); err != nil {
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

// ParseSetPerformanceFeeRate is a log parse operation binding the contract event 0x20affe2401825617c69366f8c3a3493d9822d1021d0b3023c4e77ea5b3d0fbc5.
//
// Solidity: event SetPerformanceFeeRate(address indexed caller, uint256 newPerformanceFeeRate)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSetPerformanceFeeRate(log types.Log) (*ERC4626VaultSetPerformanceFeeRate, error) {
	event := new(ERC4626VaultSetPerformanceFeeRate)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SetPerformanceFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSetPoolIterator is returned from FilterSetPool and is used to iterate over the raw logs and unpacked data for SetPool events raised by the ERC4626Vault contract.
type ERC4626VaultSetPoolIterator struct {
	Event *ERC4626VaultSetPool // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSetPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSetPool)
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
		it.Event = new(ERC4626VaultSetPool)
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
func (it *ERC4626VaultSetPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSetPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSetPool represents a SetPool event raised by the ERC4626Vault contract.
type ERC4626VaultSetPool struct {
	Caller common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPool is a free log retrieval operation binding the contract event 0x390ace337562623e4cf938891cfa7e80b7b2e6ff395963aba93e537ce67e842c.
//
// Solidity: event SetPool(address indexed caller, address indexed pool)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSetPool(opts *bind.FilterOpts, caller []common.Address, pool []common.Address) (*ERC4626VaultSetPoolIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SetPool", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSetPoolIterator{contract: _ERC4626Vault.contract, event: "SetPool", logs: logs, sub: sub}, nil
}

// WatchSetPool is a free log subscription operation binding the contract event 0x390ace337562623e4cf938891cfa7e80b7b2e6ff395963aba93e537ce67e842c.
//
// Solidity: event SetPool(address indexed caller, address indexed pool)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSetPool(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSetPool, caller []common.Address, pool []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SetPool", callerRule, poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSetPool)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SetPool", log); err != nil {
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

// ParseSetPool is a log parse operation binding the contract event 0x390ace337562623e4cf938891cfa7e80b7b2e6ff395963aba93e537ce67e842c.
//
// Solidity: event SetPool(address indexed caller, address indexed pool)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSetPool(log types.Log) (*ERC4626VaultSetPool, error) {
	event := new(ERC4626VaultSetPool)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SetPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSetTimelockIterator is returned from FilterSetTimelock and is used to iterate over the raw logs and unpacked data for SetTimelock events raised by the ERC4626Vault contract.
type ERC4626VaultSetTimelockIterator struct {
	Event *ERC4626VaultSetTimelock // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSetTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSetTimelock)
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
		it.Event = new(ERC4626VaultSetTimelock)
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
func (it *ERC4626VaultSetTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSetTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSetTimelock represents a SetTimelock event raised by the ERC4626Vault contract.
type ERC4626VaultSetTimelock struct {
	Caller      common.Address
	NewTimelock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetTimelock is a free log retrieval operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSetTimelock(opts *bind.FilterOpts, caller []common.Address) (*ERC4626VaultSetTimelockIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SetTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSetTimelockIterator{contract: _ERC4626Vault.contract, event: "SetTimelock", logs: logs, sub: sub}, nil
}

// WatchSetTimelock is a free log subscription operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSetTimelock(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSetTimelock, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SetTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSetTimelock)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SetTimelock", log); err != nil {
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

// ParseSetTimelock is a log parse operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSetTimelock(log types.Log) (*ERC4626VaultSetTimelock, error) {
	event := new(ERC4626VaultSetTimelock)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SetTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSubmitGuardianIterator is returned from FilterSubmitGuardian and is used to iterate over the raw logs and unpacked data for SubmitGuardian events raised by the ERC4626Vault contract.
type ERC4626VaultSubmitGuardianIterator struct {
	Event *ERC4626VaultSubmitGuardian // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSubmitGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSubmitGuardian)
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
		it.Event = new(ERC4626VaultSubmitGuardian)
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
func (it *ERC4626VaultSubmitGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSubmitGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSubmitGuardian represents a SubmitGuardian event raised by the ERC4626Vault contract.
type ERC4626VaultSubmitGuardian struct {
	NewGuardian common.Address
	ValidAt     uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitGuardian is a free log retrieval operation binding the contract event 0x14279aa98f18dee77127cf315bcced708f417d07da24929c2f6460b481d0c13e.
//
// Solidity: event SubmitGuardian(address newGuardian, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSubmitGuardian(opts *bind.FilterOpts) (*ERC4626VaultSubmitGuardianIterator, error) {

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SubmitGuardian")
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSubmitGuardianIterator{contract: _ERC4626Vault.contract, event: "SubmitGuardian", logs: logs, sub: sub}, nil
}

// WatchSubmitGuardian is a free log subscription operation binding the contract event 0x14279aa98f18dee77127cf315bcced708f417d07da24929c2f6460b481d0c13e.
//
// Solidity: event SubmitGuardian(address newGuardian, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSubmitGuardian(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSubmitGuardian) (event.Subscription, error) {

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SubmitGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSubmitGuardian)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitGuardian", log); err != nil {
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

// ParseSubmitGuardian is a log parse operation binding the contract event 0x14279aa98f18dee77127cf315bcced708f417d07da24929c2f6460b481d0c13e.
//
// Solidity: event SubmitGuardian(address newGuardian, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSubmitGuardian(log types.Log) (*ERC4626VaultSubmitGuardian, error) {
	event := new(ERC4626VaultSubmitGuardian)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSubmitMarketToWhitelistIterator is returned from FilterSubmitMarketToWhitelist and is used to iterate over the raw logs and unpacked data for SubmitMarketToWhitelist events raised by the ERC4626Vault contract.
type ERC4626VaultSubmitMarketToWhitelistIterator struct {
	Event *ERC4626VaultSubmitMarketToWhitelist // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSubmitMarketToWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSubmitMarketToWhitelist)
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
		it.Event = new(ERC4626VaultSubmitMarketToWhitelist)
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
func (it *ERC4626VaultSubmitMarketToWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSubmitMarketToWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSubmitMarketToWhitelist represents a SubmitMarketToWhitelist event raised by the ERC4626Vault contract.
type ERC4626VaultSubmitMarketToWhitelist struct {
	Market  common.Address
	ValidAt uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitMarketToWhitelist is a free log retrieval operation binding the contract event 0x24c5ea3dd9a5ba11b8a603674e63d62accaa480c1b05d18de5861f3894868d79.
//
// Solidity: event SubmitMarketToWhitelist(address indexed market, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSubmitMarketToWhitelist(opts *bind.FilterOpts, market []common.Address) (*ERC4626VaultSubmitMarketToWhitelistIterator, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SubmitMarketToWhitelist", marketRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSubmitMarketToWhitelistIterator{contract: _ERC4626Vault.contract, event: "SubmitMarketToWhitelist", logs: logs, sub: sub}, nil
}

// WatchSubmitMarketToWhitelist is a free log subscription operation binding the contract event 0x24c5ea3dd9a5ba11b8a603674e63d62accaa480c1b05d18de5861f3894868d79.
//
// Solidity: event SubmitMarketToWhitelist(address indexed market, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSubmitMarketToWhitelist(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSubmitMarketToWhitelist, market []common.Address) (event.Subscription, error) {

	var marketRule []interface{}
	for _, marketItem := range market {
		marketRule = append(marketRule, marketItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SubmitMarketToWhitelist", marketRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSubmitMarketToWhitelist)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitMarketToWhitelist", log); err != nil {
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

// ParseSubmitMarketToWhitelist is a log parse operation binding the contract event 0x24c5ea3dd9a5ba11b8a603674e63d62accaa480c1b05d18de5861f3894868d79.
//
// Solidity: event SubmitMarketToWhitelist(address indexed market, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSubmitMarketToWhitelist(log types.Log) (*ERC4626VaultSubmitMarketToWhitelist, error) {
	event := new(ERC4626VaultSubmitMarketToWhitelist)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitMarketToWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSubmitMinApyIterator is returned from FilterSubmitMinApy and is used to iterate over the raw logs and unpacked data for SubmitMinApy events raised by the ERC4626Vault contract.
type ERC4626VaultSubmitMinApyIterator struct {
	Event *ERC4626VaultSubmitMinApy // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSubmitMinApyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSubmitMinApy)
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
		it.Event = new(ERC4626VaultSubmitMinApy)
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
func (it *ERC4626VaultSubmitMinApyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSubmitMinApyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSubmitMinApy represents a SubmitMinApy event raised by the ERC4626Vault contract.
type ERC4626VaultSubmitMinApy struct {
	NewMinApy uint64
	ValidAt   uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubmitMinApy is a free log retrieval operation binding the contract event 0x4cedca5aa9a1cac41bbc05fef13c4ddccee26353a3de3d7615435bdf3f92ad63.
//
// Solidity: event SubmitMinApy(uint64 newMinApy, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSubmitMinApy(opts *bind.FilterOpts) (*ERC4626VaultSubmitMinApyIterator, error) {

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SubmitMinApy")
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSubmitMinApyIterator{contract: _ERC4626Vault.contract, event: "SubmitMinApy", logs: logs, sub: sub}, nil
}

// WatchSubmitMinApy is a free log subscription operation binding the contract event 0x4cedca5aa9a1cac41bbc05fef13c4ddccee26353a3de3d7615435bdf3f92ad63.
//
// Solidity: event SubmitMinApy(uint64 newMinApy, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSubmitMinApy(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSubmitMinApy) (event.Subscription, error) {

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SubmitMinApy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSubmitMinApy)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitMinApy", log); err != nil {
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

// ParseSubmitMinApy is a log parse operation binding the contract event 0x4cedca5aa9a1cac41bbc05fef13c4ddccee26353a3de3d7615435bdf3f92ad63.
//
// Solidity: event SubmitMinApy(uint64 newMinApy, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSubmitMinApy(log types.Log) (*ERC4626VaultSubmitMinApy, error) {
	event := new(ERC4626VaultSubmitMinApy)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitMinApy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSubmitPendingPoolIterator is returned from FilterSubmitPendingPool and is used to iterate over the raw logs and unpacked data for SubmitPendingPool events raised by the ERC4626Vault contract.
type ERC4626VaultSubmitPendingPoolIterator struct {
	Event *ERC4626VaultSubmitPendingPool // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSubmitPendingPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSubmitPendingPool)
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
		it.Event = new(ERC4626VaultSubmitPendingPool)
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
func (it *ERC4626VaultSubmitPendingPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSubmitPendingPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSubmitPendingPool represents a SubmitPendingPool event raised by the ERC4626Vault contract.
type ERC4626VaultSubmitPendingPool struct {
	Pool    common.Address
	ValidAt uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSubmitPendingPool is a free log retrieval operation binding the contract event 0x8bdad4b09e0168e2b64d8340297b3f8224ddb6cbdc855e6fff7e7624c6d8339d.
//
// Solidity: event SubmitPendingPool(address indexed pool, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSubmitPendingPool(opts *bind.FilterOpts, pool []common.Address) (*ERC4626VaultSubmitPendingPoolIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SubmitPendingPool", poolRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSubmitPendingPoolIterator{contract: _ERC4626Vault.contract, event: "SubmitPendingPool", logs: logs, sub: sub}, nil
}

// WatchSubmitPendingPool is a free log subscription operation binding the contract event 0x8bdad4b09e0168e2b64d8340297b3f8224ddb6cbdc855e6fff7e7624c6d8339d.
//
// Solidity: event SubmitPendingPool(address indexed pool, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSubmitPendingPool(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSubmitPendingPool, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SubmitPendingPool", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSubmitPendingPool)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitPendingPool", log); err != nil {
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

// ParseSubmitPendingPool is a log parse operation binding the contract event 0x8bdad4b09e0168e2b64d8340297b3f8224ddb6cbdc855e6fff7e7624c6d8339d.
//
// Solidity: event SubmitPendingPool(address indexed pool, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSubmitPendingPool(log types.Log) (*ERC4626VaultSubmitPendingPool, error) {
	event := new(ERC4626VaultSubmitPendingPool)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitPendingPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSubmitPerformanceFeeRateIterator is returned from FilterSubmitPerformanceFeeRate and is used to iterate over the raw logs and unpacked data for SubmitPerformanceFeeRate events raised by the ERC4626Vault contract.
type ERC4626VaultSubmitPerformanceFeeRateIterator struct {
	Event *ERC4626VaultSubmitPerformanceFeeRate // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSubmitPerformanceFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSubmitPerformanceFeeRate)
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
		it.Event = new(ERC4626VaultSubmitPerformanceFeeRate)
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
func (it *ERC4626VaultSubmitPerformanceFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSubmitPerformanceFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSubmitPerformanceFeeRate represents a SubmitPerformanceFeeRate event raised by the ERC4626Vault contract.
type ERC4626VaultSubmitPerformanceFeeRate struct {
	NewPerformanceFeeRate *big.Int
	ValidAt               uint64
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSubmitPerformanceFeeRate is a free log retrieval operation binding the contract event 0x7424f23ff556b5fe97d2f5a4b22548445b948e612029151105b441519f1bfb99.
//
// Solidity: event SubmitPerformanceFeeRate(uint256 newPerformanceFeeRate, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSubmitPerformanceFeeRate(opts *bind.FilterOpts) (*ERC4626VaultSubmitPerformanceFeeRateIterator, error) {

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SubmitPerformanceFeeRate")
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSubmitPerformanceFeeRateIterator{contract: _ERC4626Vault.contract, event: "SubmitPerformanceFeeRate", logs: logs, sub: sub}, nil
}

// WatchSubmitPerformanceFeeRate is a free log subscription operation binding the contract event 0x7424f23ff556b5fe97d2f5a4b22548445b948e612029151105b441519f1bfb99.
//
// Solidity: event SubmitPerformanceFeeRate(uint256 newPerformanceFeeRate, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSubmitPerformanceFeeRate(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSubmitPerformanceFeeRate) (event.Subscription, error) {

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SubmitPerformanceFeeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSubmitPerformanceFeeRate)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitPerformanceFeeRate", log); err != nil {
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

// ParseSubmitPerformanceFeeRate is a log parse operation binding the contract event 0x7424f23ff556b5fe97d2f5a4b22548445b948e612029151105b441519f1bfb99.
//
// Solidity: event SubmitPerformanceFeeRate(uint256 newPerformanceFeeRate, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSubmitPerformanceFeeRate(log types.Log) (*ERC4626VaultSubmitPerformanceFeeRate, error) {
	event := new(ERC4626VaultSubmitPerformanceFeeRate)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitPerformanceFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultSubmitTimelockIterator is returned from FilterSubmitTimelock and is used to iterate over the raw logs and unpacked data for SubmitTimelock events raised by the ERC4626Vault contract.
type ERC4626VaultSubmitTimelockIterator struct {
	Event *ERC4626VaultSubmitTimelock // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultSubmitTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultSubmitTimelock)
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
		it.Event = new(ERC4626VaultSubmitTimelock)
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
func (it *ERC4626VaultSubmitTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultSubmitTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultSubmitTimelock represents a SubmitTimelock event raised by the ERC4626Vault contract.
type ERC4626VaultSubmitTimelock struct {
	NewTimelock *big.Int
	ValidAt     uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitTimelock is a free log retrieval operation binding the contract event 0x6ed11f5df0bdefbbb4873e90566a9cfafbb8305c164922c173437f3c45f90a35.
//
// Solidity: event SubmitTimelock(uint256 newTimelock, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterSubmitTimelock(opts *bind.FilterOpts) (*ERC4626VaultSubmitTimelockIterator, error) {

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "SubmitTimelock")
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultSubmitTimelockIterator{contract: _ERC4626Vault.contract, event: "SubmitTimelock", logs: logs, sub: sub}, nil
}

// WatchSubmitTimelock is a free log subscription operation binding the contract event 0x6ed11f5df0bdefbbb4873e90566a9cfafbb8305c164922c173437f3c45f90a35.
//
// Solidity: event SubmitTimelock(uint256 newTimelock, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchSubmitTimelock(opts *bind.WatchOpts, sink chan<- *ERC4626VaultSubmitTimelock) (event.Subscription, error) {

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "SubmitTimelock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultSubmitTimelock)
				if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitTimelock", log); err != nil {
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

// ParseSubmitTimelock is a log parse operation binding the contract event 0x6ed11f5df0bdefbbb4873e90566a9cfafbb8305c164922c173437f3c45f90a35.
//
// Solidity: event SubmitTimelock(uint256 newTimelock, uint64 validAt)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseSubmitTimelock(log types.Log) (*ERC4626VaultSubmitTimelock, error) {
	event := new(ERC4626VaultSubmitTimelock)
	if err := _ERC4626Vault.contract.UnpackLog(event, "SubmitTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC4626Vault contract.
type ERC4626VaultTransferIterator struct {
	Event *ERC4626VaultTransfer // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultTransfer)
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
		it.Event = new(ERC4626VaultTransfer)
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
func (it *ERC4626VaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultTransfer represents a Transfer event raised by the ERC4626Vault contract.
type ERC4626VaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC4626VaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultTransferIterator{contract: _ERC4626Vault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC4626VaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultTransfer)
				if err := _ERC4626Vault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC4626Vault *ERC4626VaultFilterer) ParseTransfer(log types.Log) (*ERC4626VaultTransfer, error) {
	event := new(ERC4626VaultTransfer)
	if err := _ERC4626Vault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ERC4626Vault contract.
type ERC4626VaultUnpausedIterator struct {
	Event *ERC4626VaultUnpaused // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultUnpaused)
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
		it.Event = new(ERC4626VaultUnpaused)
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
func (it *ERC4626VaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultUnpaused represents a Unpaused event raised by the ERC4626Vault contract.
type ERC4626VaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ERC4626VaultUnpausedIterator, error) {

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultUnpausedIterator{contract: _ERC4626Vault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ERC4626VaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultUnpaused)
				if err := _ERC4626Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ERC4626Vault *ERC4626VaultFilterer) ParseUnpaused(log types.Log) (*ERC4626VaultUnpaused, error) {
	event := new(ERC4626VaultUnpaused)
	if err := _ERC4626Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ERC4626Vault contract.
type ERC4626VaultWithdrawIterator struct {
	Event *ERC4626VaultWithdraw // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultWithdraw)
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
		it.Event = new(ERC4626VaultWithdraw)
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
func (it *ERC4626VaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultWithdraw represents a Withdraw event raised by the ERC4626Vault contract.
type ERC4626VaultWithdraw struct {
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
func (_ERC4626Vault *ERC4626VaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*ERC4626VaultWithdrawIterator, error) {

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

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultWithdrawIterator{contract: _ERC4626Vault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ERC4626VaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultWithdraw)
				if err := _ERC4626Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_ERC4626Vault *ERC4626VaultFilterer) ParseWithdraw(log types.Log) (*ERC4626VaultWithdraw, error) {
	event := new(ERC4626VaultWithdraw)
	if err := _ERC4626Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC4626VaultWithdrawFtsIterator is returned from FilterWithdrawFts and is used to iterate over the raw logs and unpacked data for WithdrawFts events raised by the ERC4626Vault contract.
type ERC4626VaultWithdrawFtsIterator struct {
	Event *ERC4626VaultWithdrawFts // Event containing the contract specifics and raw log

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
func (it *ERC4626VaultWithdrawFtsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC4626VaultWithdrawFts)
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
		it.Event = new(ERC4626VaultWithdrawFts)
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
func (it *ERC4626VaultWithdrawFtsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC4626VaultWithdrawFtsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC4626VaultWithdrawFts represents a WithdrawFts event raised by the ERC4626Vault contract.
type ERC4626VaultWithdrawFts struct {
	Caller    common.Address
	Recipient common.Address
	Order     common.Address
	Amount    *big.Int
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFts is a free log retrieval operation binding the contract event 0x53239297447654f3a1c8342314051bc2fe9134b7bbe4a390eade008bb5eca1f2.
//
// Solidity: event WithdrawFts(address indexed caller, address indexed recipient, address indexed order, uint256 amount, uint256 shares)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterWithdrawFts(opts *bind.FilterOpts, caller []common.Address, recipient []common.Address, order []common.Address) (*ERC4626VaultWithdrawFtsIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var orderRule []interface{}
	for _, orderItem := range order {
		orderRule = append(orderRule, orderItem)
	}

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "WithdrawFts", callerRule, recipientRule, orderRule)
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultWithdrawFtsIterator{contract: _ERC4626Vault.contract, event: "WithdrawFts", logs: logs, sub: sub}, nil
}

// WatchWithdrawFts is a free log subscription operation binding the contract event 0x53239297447654f3a1c8342314051bc2fe9134b7bbe4a390eade008bb5eca1f2.
//
// Solidity: event WithdrawFts(address indexed caller, address indexed recipient, address indexed order, uint256 amount, uint256 shares)
func (_ERC4626Vault *ERC4626VaultFilterer) WatchWithdrawFts(opts *bind.WatchOpts, sink chan<- *ERC4626VaultWithdrawFts, caller []common.Address, recipient []common.Address, order []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var orderRule []interface{}
	for _, orderItem := range order {
		orderRule = append(orderRule, orderItem)
	}

	logs, sub, err := _ERC4626Vault.contract.WatchLogs(opts, "WithdrawFts", callerRule, recipientRule, orderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC4626VaultWithdrawFts)
				if err := _ERC4626Vault.contract.UnpackLog(event, "WithdrawFts", log); err != nil {
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

// ParseWithdrawFts is a log parse operation binding the contract event 0x53239297447654f3a1c8342314051bc2fe9134b7bbe4a390eade008bb5eca1f2.
//
// Solidity: event WithdrawFts(address indexed caller, address indexed recipient, address indexed order, uint256 amount, uint256 shares)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseWithdrawFts(log types.Log) (*ERC4626VaultWithdrawFts, error) {
	event := new(ERC4626VaultWithdrawFts)
	if err := _ERC4626Vault.contract.UnpackLog(event, "WithdrawFts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
