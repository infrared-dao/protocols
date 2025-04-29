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

// Allocation is an auto generated low-level Go binding around an user-defined struct.
type Allocation struct {
	Index  *big.Int
	Amount *big.Int
}

// GraduatedFee is an auto generated low-level Go binding around an user-defined struct.
type GraduatedFee struct {
	LowerBound *big.Int
	UpperBound *big.Int
	Fee        uint64
}

// ReturnedRewards is an auto generated low-level Go binding around an user-defined struct.
type ReturnedRewards struct {
	RewardAddress common.Address
	RewardAmount  *big.Int
}

// Strategy is an auto generated low-level Go binding around an user-defined struct.
type Strategy struct {
	Strategy   common.Address
	Allocation Allocation
}

// VaultFees is an auto generated low-level Go binding around an user-defined struct.
type VaultFees struct {
	DepositFee     uint64
	WithdrawalFee  uint64
	ProtocolFee    uint64
	PerformanceFee []GraduatedFee
}

// ConcreteVaultMetaData contains all meta data concerning the ConcreteVault contract.
var ConcreteVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccumulatedFeeAccountedMustBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AdditionFail\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllotmentTotalTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApprovalFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BlueprintUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ClaimRouterUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20ApproveFail\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ImplementationAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ImplementationDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRequest\",\"type\":\"uint256\"}],\"name\":\"InsufficientQueueRequest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientUnderlyingBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InsufficientVaultFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAssetAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBeneficiary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidClaimRouterAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDefaultAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDepositLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFeeRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"argLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedLength\",\"type\":\"uint256\"}],\"name\":\"InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMultiSigAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParkingLot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRescuer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardTokenAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubstraction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTreasuryAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUserAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVaultFees\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVaultRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawlQueue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MultiSigUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MultipleProtectStrat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoProtectionStrategiesFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NotAvailableForWithdrawal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPassedYear\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"NotValidRewardToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"OnlyVault\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"PermitDeadlineExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ProtectUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoveFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardTokenAlreadyApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardTokenNotApproved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"StrategyHasLockedAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapperBaseRewardrate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapperBonusRewardrateCtToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapperBonusRewardrateSwapToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapperBonusRewardrateUser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapperMaxProgressionFactor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"TokenNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"TotalVaultsAllowedExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attacker\",\"type\":\"address\"}],\"name\":\"TreasuryAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"TreasuryChangeRequestCooldownNotElapsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"queue\",\"type\":\"address\"}],\"name\":\"UnfinalizedWithdrawl\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"UnregisteredTokensCannotBeRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultAssetMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"VaultByTokenLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultDeployInitFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"VaultDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultIsIdle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawalsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"DepositLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultName\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldMinQueueRequest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newMinQueueRequest\",\"type\":\"uint256\"}],\"name\":\"MinimunQueueRequestUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldParkingLot\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newParkingLot\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"successfulApproval\",\"type\":\"bool\"}],\"name\":\"ParkingLotUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protectStrategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RequestedFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"RewardsHarvested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newStrategy\",\"type\":\"address\"}],\"name\":\"StrategyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structAllocation[]\",\"name\":\"newAllocations\",\"type\":\"tuple[]\"}],\"name\":\"StrategyAllocationsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldStrategy\",\"type\":\"address\"}],\"name\":\"StrategyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pastValue\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newValue\",\"type\":\"bool\"}],\"name\":\"ToggleVaultIdle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pastValue\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newValue\",\"type\":\"bool\"}],\"name\":\"WithdrawalPausedToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldQueue\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newQueue\",\"type\":\"address\"}],\"name\":\"WithdrawalQueueUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"accruedPerformanceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accruedProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"replace_\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAllocation\",\"name\":\"allocation\",\"type\":\"tuple\"}],\"internalType\":\"structStrategy\",\"name\":\"newStrategy_\",\"type\":\"tuple\"}],\"name\":\"addStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRequests\",\"type\":\"uint256\"}],\"name\":\"batchClaimWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAllocation[]\",\"name\":\"allocations_\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"redistribute\",\"type\":\"bool\"}],\"name\":\"changeAllocations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimalOffset\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"forceEject_\",\"type\":\"bool\"}],\"name\":\"emergencyRemoveStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feesUpdatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAvailableAssetsForWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStrategies\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAllocation\",\"name\":\"allocation\",\"type\":\"tuple\"}],\"internalType\":\"structStrategy[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getTotalRewardsClaimed\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"internalType\":\"structReturnedRewards[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getUserRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"internalType\":\"structReturnedRewards[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultFees\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"depositFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"withdrawalFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"protocolFee\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"}],\"internalType\":\"structGraduatedFee[]\",\"name\":\"performanceFee\",\"type\":\"tuple[]\"}],\"internalType\":\"structVaultFees\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedData\",\"type\":\"bytes\"}],\"name\":\"harvestRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highWaterMark\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"baseAsset_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"shareName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"shareSymbol_\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"contractIStrategy\",\"name\":\"strategy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structAllocation\",\"name\":\"allocation\",\"type\":\"tuple\"}],\"internalType\":\"structStrategy[]\",\"name\":\"strategies_\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"feeRecipient_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"depositFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"withdrawalFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"protocolFee\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"}],\"internalType\":\"structGraduatedFee[]\",\"name\":\"performanceFee\",\"type\":\"tuple[]\"}],\"internalType\":\"structVaultFees\",\"name\":\"fees_\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"depositLimit_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minQueueRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parkingLot\",\"outputs\":[{\"internalType\":\"contractIParkingLot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protectStrategy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"pullFundsFromSingleStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pullFundsFromStrategies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"pushFundsIntoSingleStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"pushFundsIntoSingleStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pushFundsToStrategies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares_\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"removeStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLimit_\",\"type\":\"uint256\"}],\"name\":\"setDepositLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRecipient_\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minQueueRequest_\",\"type\":\"uint256\"}],\"name\":\"setMinimunQueueRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"parkingLot_\",\"type\":\"address\"}],\"name\":\"setParkingLot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"depositFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"withdrawalFee\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"protocolFee\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"lowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"upperBound\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"fee\",\"type\":\"uint64\"}],\"internalType\":\"structGraduatedFee[]\",\"name\":\"performanceFee\",\"type\":\"tuple[]\"}],\"internalType\":\"structVaultFees\",\"name\":\"newFees_\",\"type\":\"tuple\"}],\"name\":\"setVaultFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"withdrawalQueue_\",\"type\":\"address\"}],\"name\":\"setWithdrawalQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"takePortfolioAndProtocolFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleVaultIdle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"withdrawalsPaused_\",\"type\":\"bool\"}],\"name\":\"toggleWithdrawalsPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalRewardsClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultIdle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalQueue\",\"outputs\":[{\"internalType\":\"contractIWithdrawalQueue\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalsPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ConcreteVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ConcreteVaultMetaData.ABI instead.
var ConcreteVaultABI = ConcreteVaultMetaData.ABI

// ConcreteVault is an auto generated Go binding around an Ethereum contract.
type ConcreteVault struct {
	ConcreteVaultCaller     // Read-only binding to the contract
	ConcreteVaultTransactor // Write-only binding to the contract
	ConcreteVaultFilterer   // Log filterer for contract events
}

// ConcreteVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConcreteVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConcreteVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConcreteVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConcreteVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConcreteVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConcreteVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConcreteVaultSession struct {
	Contract     *ConcreteVault    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConcreteVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConcreteVaultCallerSession struct {
	Contract *ConcreteVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ConcreteVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConcreteVaultTransactorSession struct {
	Contract     *ConcreteVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ConcreteVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConcreteVaultRaw struct {
	Contract *ConcreteVault // Generic contract binding to access the raw methods on
}

// ConcreteVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConcreteVaultCallerRaw struct {
	Contract *ConcreteVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ConcreteVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConcreteVaultTransactorRaw struct {
	Contract *ConcreteVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConcreteVault creates a new instance of ConcreteVault, bound to a specific deployed contract.
func NewConcreteVault(address common.Address, backend bind.ContractBackend) (*ConcreteVault, error) {
	contract, err := bindConcreteVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConcreteVault{ConcreteVaultCaller: ConcreteVaultCaller{contract: contract}, ConcreteVaultTransactor: ConcreteVaultTransactor{contract: contract}, ConcreteVaultFilterer: ConcreteVaultFilterer{contract: contract}}, nil
}

// NewConcreteVaultCaller creates a new read-only instance of ConcreteVault, bound to a specific deployed contract.
func NewConcreteVaultCaller(address common.Address, caller bind.ContractCaller) (*ConcreteVaultCaller, error) {
	contract, err := bindConcreteVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultCaller{contract: contract}, nil
}

// NewConcreteVaultTransactor creates a new write-only instance of ConcreteVault, bound to a specific deployed contract.
func NewConcreteVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ConcreteVaultTransactor, error) {
	contract, err := bindConcreteVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultTransactor{contract: contract}, nil
}

// NewConcreteVaultFilterer creates a new log filterer instance of ConcreteVault, bound to a specific deployed contract.
func NewConcreteVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ConcreteVaultFilterer, error) {
	contract, err := bindConcreteVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultFilterer{contract: contract}, nil
}

// bindConcreteVault binds a generic wrapper to an already deployed contract.
func bindConcreteVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConcreteVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConcreteVault *ConcreteVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConcreteVault.Contract.ConcreteVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConcreteVault *ConcreteVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcreteVault.Contract.ConcreteVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConcreteVault *ConcreteVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConcreteVault.Contract.ConcreteVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConcreteVault *ConcreteVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConcreteVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConcreteVault *ConcreteVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcreteVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConcreteVault *ConcreteVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConcreteVault.Contract.contract.Transact(opts, method, params...)
}

// AccruedPerformanceFee is a free data retrieval call binding the contract method 0x0c14935e.
//
// Solidity: function accruedPerformanceFee() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) AccruedPerformanceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "accruedPerformanceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedPerformanceFee is a free data retrieval call binding the contract method 0x0c14935e.
//
// Solidity: function accruedPerformanceFee() view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) AccruedPerformanceFee() (*big.Int, error) {
	return _ConcreteVault.Contract.AccruedPerformanceFee(&_ConcreteVault.CallOpts)
}

// AccruedPerformanceFee is a free data retrieval call binding the contract method 0x0c14935e.
//
// Solidity: function accruedPerformanceFee() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) AccruedPerformanceFee() (*big.Int, error) {
	return _ConcreteVault.Contract.AccruedPerformanceFee(&_ConcreteVault.CallOpts)
}

// AccruedProtocolFee is a free data retrieval call binding the contract method 0xacfd7928.
//
// Solidity: function accruedProtocolFee() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) AccruedProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "accruedProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccruedProtocolFee is a free data retrieval call binding the contract method 0xacfd7928.
//
// Solidity: function accruedProtocolFee() view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) AccruedProtocolFee() (*big.Int, error) {
	return _ConcreteVault.Contract.AccruedProtocolFee(&_ConcreteVault.CallOpts)
}

// AccruedProtocolFee is a free data retrieval call binding the contract method 0xacfd7928.
//
// Solidity: function accruedProtocolFee() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) AccruedProtocolFee() (*big.Int, error) {
	return _ConcreteVault.Contract.AccruedProtocolFee(&_ConcreteVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.Allowance(&_ConcreteVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.Allowance(&_ConcreteVault.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ConcreteVault *ConcreteVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ConcreteVault *ConcreteVaultSession) Asset() (common.Address, error) {
	return _ConcreteVault.Contract.Asset(&_ConcreteVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ConcreteVault *ConcreteVaultCallerSession) Asset() (common.Address, error) {
	return _ConcreteVault.Contract.Asset(&_ConcreteVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.BalanceOf(&_ConcreteVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.BalanceOf(&_ConcreteVault.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.ConvertToAssets(&_ConcreteVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.ConvertToAssets(&_ConcreteVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.ConvertToShares(&_ConcreteVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.ConvertToShares(&_ConcreteVault.CallOpts, assets)
}

// DecimalOffset is a free data retrieval call binding the contract method 0xd1b39ae5.
//
// Solidity: function decimalOffset() view returns(uint8)
func (_ConcreteVault *ConcreteVaultCaller) DecimalOffset(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "decimalOffset")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DecimalOffset is a free data retrieval call binding the contract method 0xd1b39ae5.
//
// Solidity: function decimalOffset() view returns(uint8)
func (_ConcreteVault *ConcreteVaultSession) DecimalOffset() (uint8, error) {
	return _ConcreteVault.Contract.DecimalOffset(&_ConcreteVault.CallOpts)
}

// DecimalOffset is a free data retrieval call binding the contract method 0xd1b39ae5.
//
// Solidity: function decimalOffset() view returns(uint8)
func (_ConcreteVault *ConcreteVaultCallerSession) DecimalOffset() (uint8, error) {
	return _ConcreteVault.Contract.DecimalOffset(&_ConcreteVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ConcreteVault *ConcreteVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ConcreteVault *ConcreteVaultSession) Decimals() (uint8, error) {
	return _ConcreteVault.Contract.Decimals(&_ConcreteVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ConcreteVault *ConcreteVaultCallerSession) Decimals() (uint8, error) {
	return _ConcreteVault.Contract.Decimals(&_ConcreteVault.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "depositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) DepositLimit() (*big.Int, error) {
	return _ConcreteVault.Contract.DepositLimit(&_ConcreteVault.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) DepositLimit() (*big.Int, error) {
	return _ConcreteVault.Contract.DepositLimit(&_ConcreteVault.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_ConcreteVault *ConcreteVaultCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_ConcreteVault *ConcreteVaultSession) FeeRecipient() (common.Address, error) {
	return _ConcreteVault.Contract.FeeRecipient(&_ConcreteVault.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_ConcreteVault *ConcreteVaultCallerSession) FeeRecipient() (common.Address, error) {
	return _ConcreteVault.Contract.FeeRecipient(&_ConcreteVault.CallOpts)
}

// FeesUpdatedAt is a free data retrieval call binding the contract method 0x9f3d0a48.
//
// Solidity: function feesUpdatedAt() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) FeesUpdatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "feesUpdatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeesUpdatedAt is a free data retrieval call binding the contract method 0x9f3d0a48.
//
// Solidity: function feesUpdatedAt() view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) FeesUpdatedAt() (*big.Int, error) {
	return _ConcreteVault.Contract.FeesUpdatedAt(&_ConcreteVault.CallOpts)
}

// FeesUpdatedAt is a free data retrieval call binding the contract method 0x9f3d0a48.
//
// Solidity: function feesUpdatedAt() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) FeesUpdatedAt() (*big.Int, error) {
	return _ConcreteVault.Contract.FeesUpdatedAt(&_ConcreteVault.CallOpts)
}

// FirstDeposit is a free data retrieval call binding the contract method 0xa5f5be54.
//
// Solidity: function firstDeposit() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) FirstDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "firstDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstDeposit is a free data retrieval call binding the contract method 0xa5f5be54.
//
// Solidity: function firstDeposit() view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) FirstDeposit() (*big.Int, error) {
	return _ConcreteVault.Contract.FirstDeposit(&_ConcreteVault.CallOpts)
}

// FirstDeposit is a free data retrieval call binding the contract method 0xa5f5be54.
//
// Solidity: function firstDeposit() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) FirstDeposit() (*big.Int, error) {
	return _ConcreteVault.Contract.FirstDeposit(&_ConcreteVault.CallOpts)
}

// GetAvailableAssetsForWithdrawal is a free data retrieval call binding the contract method 0x37f1834e.
//
// Solidity: function getAvailableAssetsForWithdrawal() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) GetAvailableAssetsForWithdrawal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "getAvailableAssetsForWithdrawal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvailableAssetsForWithdrawal is a free data retrieval call binding the contract method 0x37f1834e.
//
// Solidity: function getAvailableAssetsForWithdrawal() view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) GetAvailableAssetsForWithdrawal() (*big.Int, error) {
	return _ConcreteVault.Contract.GetAvailableAssetsForWithdrawal(&_ConcreteVault.CallOpts)
}

// GetAvailableAssetsForWithdrawal is a free data retrieval call binding the contract method 0x37f1834e.
//
// Solidity: function getAvailableAssetsForWithdrawal() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) GetAvailableAssetsForWithdrawal() (*big.Int, error) {
	return _ConcreteVault.Contract.GetAvailableAssetsForWithdrawal(&_ConcreteVault.CallOpts)
}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_ConcreteVault *ConcreteVaultCaller) GetRewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "getRewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_ConcreteVault *ConcreteVaultSession) GetRewardTokens() ([]common.Address, error) {
	return _ConcreteVault.Contract.GetRewardTokens(&_ConcreteVault.CallOpts)
}

// GetRewardTokens is a free data retrieval call binding the contract method 0xc4f59f9b.
//
// Solidity: function getRewardTokens() view returns(address[])
func (_ConcreteVault *ConcreteVaultCallerSession) GetRewardTokens() ([]common.Address, error) {
	return _ConcreteVault.Contract.GetRewardTokens(&_ConcreteVault.CallOpts)
}

// GetStrategies is a free data retrieval call binding the contract method 0xb49a60bb.
//
// Solidity: function getStrategies() view returns((address,(uint256,uint256))[])
func (_ConcreteVault *ConcreteVaultCaller) GetStrategies(opts *bind.CallOpts) ([]Strategy, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "getStrategies")

	if err != nil {
		return *new([]Strategy), err
	}

	out0 := *abi.ConvertType(out[0], new([]Strategy)).(*[]Strategy)

	return out0, err

}

// GetStrategies is a free data retrieval call binding the contract method 0xb49a60bb.
//
// Solidity: function getStrategies() view returns((address,(uint256,uint256))[])
func (_ConcreteVault *ConcreteVaultSession) GetStrategies() ([]Strategy, error) {
	return _ConcreteVault.Contract.GetStrategies(&_ConcreteVault.CallOpts)
}

// GetStrategies is a free data retrieval call binding the contract method 0xb49a60bb.
//
// Solidity: function getStrategies() view returns((address,(uint256,uint256))[])
func (_ConcreteVault *ConcreteVaultCallerSession) GetStrategies() ([]Strategy, error) {
	return _ConcreteVault.Contract.GetStrategies(&_ConcreteVault.CallOpts)
}

// GetTotalRewardsClaimed is a free data retrieval call binding the contract method 0x72e71601.
//
// Solidity: function getTotalRewardsClaimed(address userAddress) view returns((address,uint256)[])
func (_ConcreteVault *ConcreteVaultCaller) GetTotalRewardsClaimed(opts *bind.CallOpts, userAddress common.Address) ([]ReturnedRewards, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "getTotalRewardsClaimed", userAddress)

	if err != nil {
		return *new([]ReturnedRewards), err
	}

	out0 := *abi.ConvertType(out[0], new([]ReturnedRewards)).(*[]ReturnedRewards)

	return out0, err

}

// GetTotalRewardsClaimed is a free data retrieval call binding the contract method 0x72e71601.
//
// Solidity: function getTotalRewardsClaimed(address userAddress) view returns((address,uint256)[])
func (_ConcreteVault *ConcreteVaultSession) GetTotalRewardsClaimed(userAddress common.Address) ([]ReturnedRewards, error) {
	return _ConcreteVault.Contract.GetTotalRewardsClaimed(&_ConcreteVault.CallOpts, userAddress)
}

// GetTotalRewardsClaimed is a free data retrieval call binding the contract method 0x72e71601.
//
// Solidity: function getTotalRewardsClaimed(address userAddress) view returns((address,uint256)[])
func (_ConcreteVault *ConcreteVaultCallerSession) GetTotalRewardsClaimed(userAddress common.Address) ([]ReturnedRewards, error) {
	return _ConcreteVault.Contract.GetTotalRewardsClaimed(&_ConcreteVault.CallOpts, userAddress)
}

// GetUserRewards is a free data retrieval call binding the contract method 0x078b0fb7.
//
// Solidity: function getUserRewards(address userAddress) view returns((address,uint256)[])
func (_ConcreteVault *ConcreteVaultCaller) GetUserRewards(opts *bind.CallOpts, userAddress common.Address) ([]ReturnedRewards, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "getUserRewards", userAddress)

	if err != nil {
		return *new([]ReturnedRewards), err
	}

	out0 := *abi.ConvertType(out[0], new([]ReturnedRewards)).(*[]ReturnedRewards)

	return out0, err

}

// GetUserRewards is a free data retrieval call binding the contract method 0x078b0fb7.
//
// Solidity: function getUserRewards(address userAddress) view returns((address,uint256)[])
func (_ConcreteVault *ConcreteVaultSession) GetUserRewards(userAddress common.Address) ([]ReturnedRewards, error) {
	return _ConcreteVault.Contract.GetUserRewards(&_ConcreteVault.CallOpts, userAddress)
}

// GetUserRewards is a free data retrieval call binding the contract method 0x078b0fb7.
//
// Solidity: function getUserRewards(address userAddress) view returns((address,uint256)[])
func (_ConcreteVault *ConcreteVaultCallerSession) GetUserRewards(userAddress common.Address) ([]ReturnedRewards, error) {
	return _ConcreteVault.Contract.GetUserRewards(&_ConcreteVault.CallOpts, userAddress)
}

// GetVaultFees is a free data retrieval call binding the contract method 0x13ee9df4.
//
// Solidity: function getVaultFees() view returns((uint64,uint64,uint64,(uint256,uint256,uint64)[]))
func (_ConcreteVault *ConcreteVaultCaller) GetVaultFees(opts *bind.CallOpts) (VaultFees, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "getVaultFees")

	if err != nil {
		return *new(VaultFees), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultFees)).(*VaultFees)

	return out0, err

}

// GetVaultFees is a free data retrieval call binding the contract method 0x13ee9df4.
//
// Solidity: function getVaultFees() view returns((uint64,uint64,uint64,(uint256,uint256,uint64)[]))
func (_ConcreteVault *ConcreteVaultSession) GetVaultFees() (VaultFees, error) {
	return _ConcreteVault.Contract.GetVaultFees(&_ConcreteVault.CallOpts)
}

// GetVaultFees is a free data retrieval call binding the contract method 0x13ee9df4.
//
// Solidity: function getVaultFees() view returns((uint64,uint64,uint64,(uint256,uint256,uint64)[]))
func (_ConcreteVault *ConcreteVaultCallerSession) GetVaultFees() (VaultFees, error) {
	return _ConcreteVault.Contract.GetVaultFees(&_ConcreteVault.CallOpts)
}

// HighWaterMark is a free data retrieval call binding the contract method 0x1e8410da.
//
// Solidity: function highWaterMark() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) HighWaterMark(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "highWaterMark")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighWaterMark is a free data retrieval call binding the contract method 0x1e8410da.
//
// Solidity: function highWaterMark() view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) HighWaterMark() (*big.Int, error) {
	return _ConcreteVault.Contract.HighWaterMark(&_ConcreteVault.CallOpts)
}

// HighWaterMark is a free data retrieval call binding the contract method 0x1e8410da.
//
// Solidity: function highWaterMark() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) HighWaterMark() (*big.Int, error) {
	return _ConcreteVault.Contract.HighWaterMark(&_ConcreteVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.MaxDeposit(&_ConcreteVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.MaxDeposit(&_ConcreteVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.MaxMint(&_ConcreteVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.MaxMint(&_ConcreteVault.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.MaxRedeem(&_ConcreteVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.MaxRedeem(&_ConcreteVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.MaxWithdraw(&_ConcreteVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.MaxWithdraw(&_ConcreteVault.CallOpts, owner)
}

// MinQueueRequest is a free data retrieval call binding the contract method 0x19e67246.
//
// Solidity: function minQueueRequest() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) MinQueueRequest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "minQueueRequest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQueueRequest is a free data retrieval call binding the contract method 0x19e67246.
//
// Solidity: function minQueueRequest() view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) MinQueueRequest() (*big.Int, error) {
	return _ConcreteVault.Contract.MinQueueRequest(&_ConcreteVault.CallOpts)
}

// MinQueueRequest is a free data retrieval call binding the contract method 0x19e67246.
//
// Solidity: function minQueueRequest() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) MinQueueRequest() (*big.Int, error) {
	return _ConcreteVault.Contract.MinQueueRequest(&_ConcreteVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ConcreteVault *ConcreteVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ConcreteVault *ConcreteVaultSession) Name() (string, error) {
	return _ConcreteVault.Contract.Name(&_ConcreteVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ConcreteVault *ConcreteVaultCallerSession) Name() (string, error) {
	return _ConcreteVault.Contract.Name(&_ConcreteVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConcreteVault *ConcreteVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConcreteVault *ConcreteVaultSession) Owner() (common.Address, error) {
	return _ConcreteVault.Contract.Owner(&_ConcreteVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConcreteVault *ConcreteVaultCallerSession) Owner() (common.Address, error) {
	return _ConcreteVault.Contract.Owner(&_ConcreteVault.CallOpts)
}

// ParkingLot is a free data retrieval call binding the contract method 0x49d0ea97.
//
// Solidity: function parkingLot() view returns(address)
func (_ConcreteVault *ConcreteVaultCaller) ParkingLot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "parkingLot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParkingLot is a free data retrieval call binding the contract method 0x49d0ea97.
//
// Solidity: function parkingLot() view returns(address)
func (_ConcreteVault *ConcreteVaultSession) ParkingLot() (common.Address, error) {
	return _ConcreteVault.Contract.ParkingLot(&_ConcreteVault.CallOpts)
}

// ParkingLot is a free data retrieval call binding the contract method 0x49d0ea97.
//
// Solidity: function parkingLot() view returns(address)
func (_ConcreteVault *ConcreteVaultCallerSession) ParkingLot() (common.Address, error) {
	return _ConcreteVault.Contract.ParkingLot(&_ConcreteVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ConcreteVault *ConcreteVaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ConcreteVault *ConcreteVaultSession) Paused() (bool, error) {
	return _ConcreteVault.Contract.Paused(&_ConcreteVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ConcreteVault *ConcreteVaultCallerSession) Paused() (bool, error) {
	return _ConcreteVault.Contract.Paused(&_ConcreteVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets_) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "previewDeposit", assets_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets_) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) PreviewDeposit(assets_ *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.PreviewDeposit(&_ConcreteVault.CallOpts, assets_)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets_) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) PreviewDeposit(assets_ *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.PreviewDeposit(&_ConcreteVault.CallOpts, assets_)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares_) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) PreviewMint(opts *bind.CallOpts, shares_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "previewMint", shares_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares_) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) PreviewMint(shares_ *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.PreviewMint(&_ConcreteVault.CallOpts, shares_)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares_) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) PreviewMint(shares_ *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.PreviewMint(&_ConcreteVault.CallOpts, shares_)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares_) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "previewRedeem", shares_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares_) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) PreviewRedeem(shares_ *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.PreviewRedeem(&_ConcreteVault.CallOpts, shares_)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares_) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) PreviewRedeem(shares_ *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.PreviewRedeem(&_ConcreteVault.CallOpts, shares_)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets_) view returns(uint256 shares)
func (_ConcreteVault *ConcreteVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "previewWithdraw", assets_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets_) view returns(uint256 shares)
func (_ConcreteVault *ConcreteVaultSession) PreviewWithdraw(assets_ *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.PreviewWithdraw(&_ConcreteVault.CallOpts, assets_)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets_) view returns(uint256 shares)
func (_ConcreteVault *ConcreteVaultCallerSession) PreviewWithdraw(assets_ *big.Int) (*big.Int, error) {
	return _ConcreteVault.Contract.PreviewWithdraw(&_ConcreteVault.CallOpts, assets_)
}

// ProtectStrategy is a free data retrieval call binding the contract method 0x9e2f0e22.
//
// Solidity: function protectStrategy() view returns(address)
func (_ConcreteVault *ConcreteVaultCaller) ProtectStrategy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "protectStrategy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtectStrategy is a free data retrieval call binding the contract method 0x9e2f0e22.
//
// Solidity: function protectStrategy() view returns(address)
func (_ConcreteVault *ConcreteVaultSession) ProtectStrategy() (common.Address, error) {
	return _ConcreteVault.Contract.ProtectStrategy(&_ConcreteVault.CallOpts)
}

// ProtectStrategy is a free data retrieval call binding the contract method 0x9e2f0e22.
//
// Solidity: function protectStrategy() view returns(address)
func (_ConcreteVault *ConcreteVaultCallerSession) ProtectStrategy() (common.Address, error) {
	return _ConcreteVault.Contract.ProtectStrategy(&_ConcreteVault.CallOpts)
}

// RewardIndex is a free data retrieval call binding the contract method 0xb3e32114.
//
// Solidity: function rewardIndex(address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) RewardIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "rewardIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardIndex is a free data retrieval call binding the contract method 0xb3e32114.
//
// Solidity: function rewardIndex(address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) RewardIndex(arg0 common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.RewardIndex(&_ConcreteVault.CallOpts, arg0)
}

// RewardIndex is a free data retrieval call binding the contract method 0xb3e32114.
//
// Solidity: function rewardIndex(address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) RewardIndex(arg0 common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.RewardIndex(&_ConcreteVault.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ConcreteVault *ConcreteVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ConcreteVault *ConcreteVaultSession) Symbol() (string, error) {
	return _ConcreteVault.Contract.Symbol(&_ConcreteVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ConcreteVault *ConcreteVaultCallerSession) Symbol() (string, error) {
	return _ConcreteVault.Contract.Symbol(&_ConcreteVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 total)
func (_ConcreteVault *ConcreteVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 total)
func (_ConcreteVault *ConcreteVaultSession) TotalAssets() (*big.Int, error) {
	return _ConcreteVault.Contract.TotalAssets(&_ConcreteVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 total)
func (_ConcreteVault *ConcreteVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _ConcreteVault.Contract.TotalAssets(&_ConcreteVault.CallOpts)
}

// TotalRewardsClaimed is a free data retrieval call binding the contract method 0x1d2af0b9.
//
// Solidity: function totalRewardsClaimed(address , address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) TotalRewardsClaimed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "totalRewardsClaimed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardsClaimed is a free data retrieval call binding the contract method 0x1d2af0b9.
//
// Solidity: function totalRewardsClaimed(address , address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) TotalRewardsClaimed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.TotalRewardsClaimed(&_ConcreteVault.CallOpts, arg0, arg1)
}

// TotalRewardsClaimed is a free data retrieval call binding the contract method 0x1d2af0b9.
//
// Solidity: function totalRewardsClaimed(address , address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) TotalRewardsClaimed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.TotalRewardsClaimed(&_ConcreteVault.CallOpts, arg0, arg1)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) TotalSupply() (*big.Int, error) {
	return _ConcreteVault.Contract.TotalSupply(&_ConcreteVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _ConcreteVault.Contract.TotalSupply(&_ConcreteVault.CallOpts)
}

// UserRewardIndex is a free data retrieval call binding the contract method 0xa039c129.
//
// Solidity: function userRewardIndex(address , address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCaller) UserRewardIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "userRewardIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardIndex is a free data retrieval call binding the contract method 0xa039c129.
//
// Solidity: function userRewardIndex(address , address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) UserRewardIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.UserRewardIndex(&_ConcreteVault.CallOpts, arg0, arg1)
}

// UserRewardIndex is a free data retrieval call binding the contract method 0xa039c129.
//
// Solidity: function userRewardIndex(address , address ) view returns(uint256)
func (_ConcreteVault *ConcreteVaultCallerSession) UserRewardIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ConcreteVault.Contract.UserRewardIndex(&_ConcreteVault.CallOpts, arg0, arg1)
}

// VaultIdle is a free data retrieval call binding the contract method 0x111c8453.
//
// Solidity: function vaultIdle() view returns(bool)
func (_ConcreteVault *ConcreteVaultCaller) VaultIdle(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "vaultIdle")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VaultIdle is a free data retrieval call binding the contract method 0x111c8453.
//
// Solidity: function vaultIdle() view returns(bool)
func (_ConcreteVault *ConcreteVaultSession) VaultIdle() (bool, error) {
	return _ConcreteVault.Contract.VaultIdle(&_ConcreteVault.CallOpts)
}

// VaultIdle is a free data retrieval call binding the contract method 0x111c8453.
//
// Solidity: function vaultIdle() view returns(bool)
func (_ConcreteVault *ConcreteVaultCallerSession) VaultIdle() (bool, error) {
	return _ConcreteVault.Contract.VaultIdle(&_ConcreteVault.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0x37d5fe99.
//
// Solidity: function withdrawalQueue() view returns(address)
func (_ConcreteVault *ConcreteVaultCaller) WithdrawalQueue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "withdrawalQueue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalQueue is a free data retrieval call binding the contract method 0x37d5fe99.
//
// Solidity: function withdrawalQueue() view returns(address)
func (_ConcreteVault *ConcreteVaultSession) WithdrawalQueue() (common.Address, error) {
	return _ConcreteVault.Contract.WithdrawalQueue(&_ConcreteVault.CallOpts)
}

// WithdrawalQueue is a free data retrieval call binding the contract method 0x37d5fe99.
//
// Solidity: function withdrawalQueue() view returns(address)
func (_ConcreteVault *ConcreteVaultCallerSession) WithdrawalQueue() (common.Address, error) {
	return _ConcreteVault.Contract.WithdrawalQueue(&_ConcreteVault.CallOpts)
}

// WithdrawalsPaused is a free data retrieval call binding the contract method 0xe9f2838e.
//
// Solidity: function withdrawalsPaused() view returns(bool)
func (_ConcreteVault *ConcreteVaultCaller) WithdrawalsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConcreteVault.contract.Call(opts, &out, "withdrawalsPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalsPaused is a free data retrieval call binding the contract method 0xe9f2838e.
//
// Solidity: function withdrawalsPaused() view returns(bool)
func (_ConcreteVault *ConcreteVaultSession) WithdrawalsPaused() (bool, error) {
	return _ConcreteVault.Contract.WithdrawalsPaused(&_ConcreteVault.CallOpts)
}

// WithdrawalsPaused is a free data retrieval call binding the contract method 0xe9f2838e.
//
// Solidity: function withdrawalsPaused() view returns(bool)
func (_ConcreteVault *ConcreteVaultCallerSession) WithdrawalsPaused() (bool, error) {
	return _ConcreteVault.Contract.WithdrawalsPaused(&_ConcreteVault.CallOpts)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xb45ba4a8.
//
// Solidity: function addStrategy(uint256 index_, bool replace_, (address,(uint256,uint256)) newStrategy_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) AddStrategy(opts *bind.TransactOpts, index_ *big.Int, replace_ bool, newStrategy_ Strategy) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "addStrategy", index_, replace_, newStrategy_)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xb45ba4a8.
//
// Solidity: function addStrategy(uint256 index_, bool replace_, (address,(uint256,uint256)) newStrategy_) returns()
func (_ConcreteVault *ConcreteVaultSession) AddStrategy(index_ *big.Int, replace_ bool, newStrategy_ Strategy) (*types.Transaction, error) {
	return _ConcreteVault.Contract.AddStrategy(&_ConcreteVault.TransactOpts, index_, replace_, newStrategy_)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xb45ba4a8.
//
// Solidity: function addStrategy(uint256 index_, bool replace_, (address,(uint256,uint256)) newStrategy_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) AddStrategy(index_ *big.Int, replace_ bool, newStrategy_ Strategy) (*types.Transaction, error) {
	return _ConcreteVault.Contract.AddStrategy(&_ConcreteVault.TransactOpts, index_, replace_, newStrategy_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ConcreteVault *ConcreteVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ConcreteVault *ConcreteVaultSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Approve(&_ConcreteVault.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ConcreteVault *ConcreteVaultTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Approve(&_ConcreteVault.TransactOpts, spender, value)
}

// BatchClaimWithdrawal is a paid mutator transaction binding the contract method 0xf1f741b4.
//
// Solidity: function batchClaimWithdrawal(uint256 maxRequests) returns()
func (_ConcreteVault *ConcreteVaultTransactor) BatchClaimWithdrawal(opts *bind.TransactOpts, maxRequests *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "batchClaimWithdrawal", maxRequests)
}

// BatchClaimWithdrawal is a paid mutator transaction binding the contract method 0xf1f741b4.
//
// Solidity: function batchClaimWithdrawal(uint256 maxRequests) returns()
func (_ConcreteVault *ConcreteVaultSession) BatchClaimWithdrawal(maxRequests *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.BatchClaimWithdrawal(&_ConcreteVault.TransactOpts, maxRequests)
}

// BatchClaimWithdrawal is a paid mutator transaction binding the contract method 0xf1f741b4.
//
// Solidity: function batchClaimWithdrawal(uint256 maxRequests) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) BatchClaimWithdrawal(maxRequests *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.BatchClaimWithdrawal(&_ConcreteVault.TransactOpts, maxRequests)
}

// ChangeAllocations is a paid mutator transaction binding the contract method 0xf4aa3efa.
//
// Solidity: function changeAllocations((uint256,uint256)[] allocations_, bool redistribute) returns()
func (_ConcreteVault *ConcreteVaultTransactor) ChangeAllocations(opts *bind.TransactOpts, allocations_ []Allocation, redistribute bool) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "changeAllocations", allocations_, redistribute)
}

// ChangeAllocations is a paid mutator transaction binding the contract method 0xf4aa3efa.
//
// Solidity: function changeAllocations((uint256,uint256)[] allocations_, bool redistribute) returns()
func (_ConcreteVault *ConcreteVaultSession) ChangeAllocations(allocations_ []Allocation, redistribute bool) (*types.Transaction, error) {
	return _ConcreteVault.Contract.ChangeAllocations(&_ConcreteVault.TransactOpts, allocations_, redistribute)
}

// ChangeAllocations is a paid mutator transaction binding the contract method 0xf4aa3efa.
//
// Solidity: function changeAllocations((uint256,uint256)[] allocations_, bool redistribute) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) ChangeAllocations(allocations_ []Allocation, redistribute bool) (*types.Transaction, error) {
	return _ConcreteVault.Contract.ChangeAllocations(&_ConcreteVault.TransactOpts, allocations_, redistribute)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_ConcreteVault *ConcreteVaultTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_ConcreteVault *ConcreteVaultSession) ClaimRewards() (*types.Transaction, error) {
	return _ConcreteVault.Contract.ClaimRewards(&_ConcreteVault.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _ConcreteVault.Contract.ClaimRewards(&_ConcreteVault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets_, address receiver_) returns(uint256 shares)
func (_ConcreteVault *ConcreteVaultTransactor) Deposit(opts *bind.TransactOpts, assets_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "deposit", assets_, receiver_)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets_, address receiver_) returns(uint256 shares)
func (_ConcreteVault *ConcreteVaultSession) Deposit(assets_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Deposit(&_ConcreteVault.TransactOpts, assets_, receiver_)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets_, address receiver_) returns(uint256 shares)
func (_ConcreteVault *ConcreteVaultTransactorSession) Deposit(assets_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Deposit(&_ConcreteVault.TransactOpts, assets_, receiver_)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets_) returns(uint256)
func (_ConcreteVault *ConcreteVaultTransactor) Deposit0(opts *bind.TransactOpts, assets_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "deposit0", assets_)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets_) returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) Deposit0(assets_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Deposit0(&_ConcreteVault.TransactOpts, assets_)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 assets_) returns(uint256)
func (_ConcreteVault *ConcreteVaultTransactorSession) Deposit0(assets_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Deposit0(&_ConcreteVault.TransactOpts, assets_)
}

// EmergencyRemoveStrategy is a paid mutator transaction binding the contract method 0x0070e595.
//
// Solidity: function emergencyRemoveStrategy(uint256 index_, bool forceEject_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) EmergencyRemoveStrategy(opts *bind.TransactOpts, index_ *big.Int, forceEject_ bool) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "emergencyRemoveStrategy", index_, forceEject_)
}

// EmergencyRemoveStrategy is a paid mutator transaction binding the contract method 0x0070e595.
//
// Solidity: function emergencyRemoveStrategy(uint256 index_, bool forceEject_) returns()
func (_ConcreteVault *ConcreteVaultSession) EmergencyRemoveStrategy(index_ *big.Int, forceEject_ bool) (*types.Transaction, error) {
	return _ConcreteVault.Contract.EmergencyRemoveStrategy(&_ConcreteVault.TransactOpts, index_, forceEject_)
}

// EmergencyRemoveStrategy is a paid mutator transaction binding the contract method 0x0070e595.
//
// Solidity: function emergencyRemoveStrategy(uint256 index_, bool forceEject_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) EmergencyRemoveStrategy(index_ *big.Int, forceEject_ bool) (*types.Transaction, error) {
	return _ConcreteVault.Contract.EmergencyRemoveStrategy(&_ConcreteVault.TransactOpts, index_, forceEject_)
}

// HarvestRewards is a paid mutator transaction binding the contract method 0x2a81196c.
//
// Solidity: function harvestRewards(bytes encodedData) returns()
func (_ConcreteVault *ConcreteVaultTransactor) HarvestRewards(opts *bind.TransactOpts, encodedData []byte) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "harvestRewards", encodedData)
}

// HarvestRewards is a paid mutator transaction binding the contract method 0x2a81196c.
//
// Solidity: function harvestRewards(bytes encodedData) returns()
func (_ConcreteVault *ConcreteVaultSession) HarvestRewards(encodedData []byte) (*types.Transaction, error) {
	return _ConcreteVault.Contract.HarvestRewards(&_ConcreteVault.TransactOpts, encodedData)
}

// HarvestRewards is a paid mutator transaction binding the contract method 0x2a81196c.
//
// Solidity: function harvestRewards(bytes encodedData) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) HarvestRewards(encodedData []byte) (*types.Transaction, error) {
	return _ConcreteVault.Contract.HarvestRewards(&_ConcreteVault.TransactOpts, encodedData)
}

// Initialize is a paid mutator transaction binding the contract method 0xc6060762.
//
// Solidity: function initialize(address baseAsset_, string shareName_, string shareSymbol_, (address,(uint256,uint256))[] strategies_, address feeRecipient_, (uint64,uint64,uint64,(uint256,uint256,uint64)[]) fees_, uint256 depositLimit_, address owner_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) Initialize(opts *bind.TransactOpts, baseAsset_ common.Address, shareName_ string, shareSymbol_ string, strategies_ []Strategy, feeRecipient_ common.Address, fees_ VaultFees, depositLimit_ *big.Int, owner_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "initialize", baseAsset_, shareName_, shareSymbol_, strategies_, feeRecipient_, fees_, depositLimit_, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc6060762.
//
// Solidity: function initialize(address baseAsset_, string shareName_, string shareSymbol_, (address,(uint256,uint256))[] strategies_, address feeRecipient_, (uint64,uint64,uint64,(uint256,uint256,uint64)[]) fees_, uint256 depositLimit_, address owner_) returns()
func (_ConcreteVault *ConcreteVaultSession) Initialize(baseAsset_ common.Address, shareName_ string, shareSymbol_ string, strategies_ []Strategy, feeRecipient_ common.Address, fees_ VaultFees, depositLimit_ *big.Int, owner_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Initialize(&_ConcreteVault.TransactOpts, baseAsset_, shareName_, shareSymbol_, strategies_, feeRecipient_, fees_, depositLimit_, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc6060762.
//
// Solidity: function initialize(address baseAsset_, string shareName_, string shareSymbol_, (address,(uint256,uint256))[] strategies_, address feeRecipient_, (uint64,uint64,uint64,(uint256,uint256,uint64)[]) fees_, uint256 depositLimit_, address owner_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) Initialize(baseAsset_ common.Address, shareName_ string, shareSymbol_ string, strategies_ []Strategy, feeRecipient_ common.Address, fees_ VaultFees, depositLimit_ *big.Int, owner_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Initialize(&_ConcreteVault.TransactOpts, baseAsset_, shareName_, shareSymbol_, strategies_, feeRecipient_, fees_, depositLimit_, owner_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares_, address receiver_) returns(uint256 assets)
func (_ConcreteVault *ConcreteVaultTransactor) Mint(opts *bind.TransactOpts, shares_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "mint", shares_, receiver_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares_, address receiver_) returns(uint256 assets)
func (_ConcreteVault *ConcreteVaultSession) Mint(shares_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Mint(&_ConcreteVault.TransactOpts, shares_, receiver_)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares_, address receiver_) returns(uint256 assets)
func (_ConcreteVault *ConcreteVaultTransactorSession) Mint(shares_ *big.Int, receiver_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Mint(&_ConcreteVault.TransactOpts, shares_, receiver_)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 shares_) returns(uint256)
func (_ConcreteVault *ConcreteVaultTransactor) Mint0(opts *bind.TransactOpts, shares_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "mint0", shares_)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 shares_) returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) Mint0(shares_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Mint0(&_ConcreteVault.TransactOpts, shares_)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 shares_) returns(uint256)
func (_ConcreteVault *ConcreteVaultTransactorSession) Mint0(shares_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Mint0(&_ConcreteVault.TransactOpts, shares_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ConcreteVault *ConcreteVaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ConcreteVault *ConcreteVaultSession) Pause() (*types.Transaction, error) {
	return _ConcreteVault.Contract.Pause(&_ConcreteVault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) Pause() (*types.Transaction, error) {
	return _ConcreteVault.Contract.Pause(&_ConcreteVault.TransactOpts)
}

// PullFundsFromSingleStrategy is a paid mutator transaction binding the contract method 0x6d3300e5.
//
// Solidity: function pullFundsFromSingleStrategy(uint256 index_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) PullFundsFromSingleStrategy(opts *bind.TransactOpts, index_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "pullFundsFromSingleStrategy", index_)
}

// PullFundsFromSingleStrategy is a paid mutator transaction binding the contract method 0x6d3300e5.
//
// Solidity: function pullFundsFromSingleStrategy(uint256 index_) returns()
func (_ConcreteVault *ConcreteVaultSession) PullFundsFromSingleStrategy(index_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.PullFundsFromSingleStrategy(&_ConcreteVault.TransactOpts, index_)
}

// PullFundsFromSingleStrategy is a paid mutator transaction binding the contract method 0x6d3300e5.
//
// Solidity: function pullFundsFromSingleStrategy(uint256 index_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) PullFundsFromSingleStrategy(index_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.PullFundsFromSingleStrategy(&_ConcreteVault.TransactOpts, index_)
}

// PullFundsFromStrategies is a paid mutator transaction binding the contract method 0x66584f8c.
//
// Solidity: function pullFundsFromStrategies() returns()
func (_ConcreteVault *ConcreteVaultTransactor) PullFundsFromStrategies(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "pullFundsFromStrategies")
}

// PullFundsFromStrategies is a paid mutator transaction binding the contract method 0x66584f8c.
//
// Solidity: function pullFundsFromStrategies() returns()
func (_ConcreteVault *ConcreteVaultSession) PullFundsFromStrategies() (*types.Transaction, error) {
	return _ConcreteVault.Contract.PullFundsFromStrategies(&_ConcreteVault.TransactOpts)
}

// PullFundsFromStrategies is a paid mutator transaction binding the contract method 0x66584f8c.
//
// Solidity: function pullFundsFromStrategies() returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) PullFundsFromStrategies() (*types.Transaction, error) {
	return _ConcreteVault.Contract.PullFundsFromStrategies(&_ConcreteVault.TransactOpts)
}

// PushFundsIntoSingleStrategy is a paid mutator transaction binding the contract method 0x839e950f.
//
// Solidity: function pushFundsIntoSingleStrategy(uint256 index_, uint256 amount) returns()
func (_ConcreteVault *ConcreteVaultTransactor) PushFundsIntoSingleStrategy(opts *bind.TransactOpts, index_ *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "pushFundsIntoSingleStrategy", index_, amount)
}

// PushFundsIntoSingleStrategy is a paid mutator transaction binding the contract method 0x839e950f.
//
// Solidity: function pushFundsIntoSingleStrategy(uint256 index_, uint256 amount) returns()
func (_ConcreteVault *ConcreteVaultSession) PushFundsIntoSingleStrategy(index_ *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.PushFundsIntoSingleStrategy(&_ConcreteVault.TransactOpts, index_, amount)
}

// PushFundsIntoSingleStrategy is a paid mutator transaction binding the contract method 0x839e950f.
//
// Solidity: function pushFundsIntoSingleStrategy(uint256 index_, uint256 amount) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) PushFundsIntoSingleStrategy(index_ *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.PushFundsIntoSingleStrategy(&_ConcreteVault.TransactOpts, index_, amount)
}

// PushFundsIntoSingleStrategy0 is a paid mutator transaction binding the contract method 0xb39ed6d9.
//
// Solidity: function pushFundsIntoSingleStrategy(uint256 index_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) PushFundsIntoSingleStrategy0(opts *bind.TransactOpts, index_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "pushFundsIntoSingleStrategy0", index_)
}

// PushFundsIntoSingleStrategy0 is a paid mutator transaction binding the contract method 0xb39ed6d9.
//
// Solidity: function pushFundsIntoSingleStrategy(uint256 index_) returns()
func (_ConcreteVault *ConcreteVaultSession) PushFundsIntoSingleStrategy0(index_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.PushFundsIntoSingleStrategy0(&_ConcreteVault.TransactOpts, index_)
}

// PushFundsIntoSingleStrategy0 is a paid mutator transaction binding the contract method 0xb39ed6d9.
//
// Solidity: function pushFundsIntoSingleStrategy(uint256 index_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) PushFundsIntoSingleStrategy0(index_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.PushFundsIntoSingleStrategy0(&_ConcreteVault.TransactOpts, index_)
}

// PushFundsToStrategies is a paid mutator transaction binding the contract method 0x55b8b781.
//
// Solidity: function pushFundsToStrategies() returns()
func (_ConcreteVault *ConcreteVaultTransactor) PushFundsToStrategies(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "pushFundsToStrategies")
}

// PushFundsToStrategies is a paid mutator transaction binding the contract method 0x55b8b781.
//
// Solidity: function pushFundsToStrategies() returns()
func (_ConcreteVault *ConcreteVaultSession) PushFundsToStrategies() (*types.Transaction, error) {
	return _ConcreteVault.Contract.PushFundsToStrategies(&_ConcreteVault.TransactOpts)
}

// PushFundsToStrategies is a paid mutator transaction binding the contract method 0x55b8b781.
//
// Solidity: function pushFundsToStrategies() returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) PushFundsToStrategies() (*types.Transaction, error) {
	return _ConcreteVault.Contract.PushFundsToStrategies(&_ConcreteVault.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares_, address receiver_, address owner_) returns(uint256 assets)
func (_ConcreteVault *ConcreteVaultTransactor) Redeem(opts *bind.TransactOpts, shares_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "redeem", shares_, receiver_, owner_)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares_, address receiver_, address owner_) returns(uint256 assets)
func (_ConcreteVault *ConcreteVaultSession) Redeem(shares_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Redeem(&_ConcreteVault.TransactOpts, shares_, receiver_, owner_)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares_, address receiver_, address owner_) returns(uint256 assets)
func (_ConcreteVault *ConcreteVaultTransactorSession) Redeem(shares_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Redeem(&_ConcreteVault.TransactOpts, shares_, receiver_, owner_)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 shares_) returns(uint256)
func (_ConcreteVault *ConcreteVaultTransactor) Redeem0(opts *bind.TransactOpts, shares_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "redeem0", shares_)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 shares_) returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) Redeem0(shares_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Redeem0(&_ConcreteVault.TransactOpts, shares_)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 shares_) returns(uint256)
func (_ConcreteVault *ConcreteVaultTransactorSession) Redeem0(shares_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Redeem0(&_ConcreteVault.TransactOpts, shares_)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0xc0cbbca6.
//
// Solidity: function removeStrategy(uint256 index_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) RemoveStrategy(opts *bind.TransactOpts, index_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "removeStrategy", index_)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0xc0cbbca6.
//
// Solidity: function removeStrategy(uint256 index_) returns()
func (_ConcreteVault *ConcreteVaultSession) RemoveStrategy(index_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.RemoveStrategy(&_ConcreteVault.TransactOpts, index_)
}

// RemoveStrategy is a paid mutator transaction binding the contract method 0xc0cbbca6.
//
// Solidity: function removeStrategy(uint256 index_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) RemoveStrategy(index_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.RemoveStrategy(&_ConcreteVault.TransactOpts, index_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConcreteVault *ConcreteVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConcreteVault *ConcreteVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConcreteVault.Contract.RenounceOwnership(&_ConcreteVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ConcreteVault.Contract.RenounceOwnership(&_ConcreteVault.TransactOpts)
}

// RequestFunds is a paid mutator transaction binding the contract method 0xbf3bd25d.
//
// Solidity: function requestFunds(uint256 amount) returns()
func (_ConcreteVault *ConcreteVaultTransactor) RequestFunds(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "requestFunds", amount)
}

// RequestFunds is a paid mutator transaction binding the contract method 0xbf3bd25d.
//
// Solidity: function requestFunds(uint256 amount) returns()
func (_ConcreteVault *ConcreteVaultSession) RequestFunds(amount *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.RequestFunds(&_ConcreteVault.TransactOpts, amount)
}

// RequestFunds is a paid mutator transaction binding the contract method 0xbf3bd25d.
//
// Solidity: function requestFunds(uint256 amount) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) RequestFunds(amount *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.RequestFunds(&_ConcreteVault.TransactOpts, amount)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 newLimit_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) SetDepositLimit(opts *bind.TransactOpts, newLimit_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "setDepositLimit", newLimit_)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 newLimit_) returns()
func (_ConcreteVault *ConcreteVaultSession) SetDepositLimit(newLimit_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetDepositLimit(&_ConcreteVault.TransactOpts, newLimit_)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 newLimit_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) SetDepositLimit(newLimit_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetDepositLimit(&_ConcreteVault.TransactOpts, newLimit_)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newRecipient_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) SetFeeRecipient(opts *bind.TransactOpts, newRecipient_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "setFeeRecipient", newRecipient_)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newRecipient_) returns()
func (_ConcreteVault *ConcreteVaultSession) SetFeeRecipient(newRecipient_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetFeeRecipient(&_ConcreteVault.TransactOpts, newRecipient_)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newRecipient_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) SetFeeRecipient(newRecipient_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetFeeRecipient(&_ConcreteVault.TransactOpts, newRecipient_)
}

// SetMinimunQueueRequest is a paid mutator transaction binding the contract method 0xf745d6ca.
//
// Solidity: function setMinimunQueueRequest(uint256 minQueueRequest_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) SetMinimunQueueRequest(opts *bind.TransactOpts, minQueueRequest_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "setMinimunQueueRequest", minQueueRequest_)
}

// SetMinimunQueueRequest is a paid mutator transaction binding the contract method 0xf745d6ca.
//
// Solidity: function setMinimunQueueRequest(uint256 minQueueRequest_) returns()
func (_ConcreteVault *ConcreteVaultSession) SetMinimunQueueRequest(minQueueRequest_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetMinimunQueueRequest(&_ConcreteVault.TransactOpts, minQueueRequest_)
}

// SetMinimunQueueRequest is a paid mutator transaction binding the contract method 0xf745d6ca.
//
// Solidity: function setMinimunQueueRequest(uint256 minQueueRequest_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) SetMinimunQueueRequest(minQueueRequest_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetMinimunQueueRequest(&_ConcreteVault.TransactOpts, minQueueRequest_)
}

// SetParkingLot is a paid mutator transaction binding the contract method 0x59f7c845.
//
// Solidity: function setParkingLot(address parkingLot_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) SetParkingLot(opts *bind.TransactOpts, parkingLot_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "setParkingLot", parkingLot_)
}

// SetParkingLot is a paid mutator transaction binding the contract method 0x59f7c845.
//
// Solidity: function setParkingLot(address parkingLot_) returns()
func (_ConcreteVault *ConcreteVaultSession) SetParkingLot(parkingLot_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetParkingLot(&_ConcreteVault.TransactOpts, parkingLot_)
}

// SetParkingLot is a paid mutator transaction binding the contract method 0x59f7c845.
//
// Solidity: function setParkingLot(address parkingLot_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) SetParkingLot(parkingLot_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetParkingLot(&_ConcreteVault.TransactOpts, parkingLot_)
}

// SetVaultFees is a paid mutator transaction binding the contract method 0x6bedc46c.
//
// Solidity: function setVaultFees((uint64,uint64,uint64,(uint256,uint256,uint64)[]) newFees_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) SetVaultFees(opts *bind.TransactOpts, newFees_ VaultFees) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "setVaultFees", newFees_)
}

// SetVaultFees is a paid mutator transaction binding the contract method 0x6bedc46c.
//
// Solidity: function setVaultFees((uint64,uint64,uint64,(uint256,uint256,uint64)[]) newFees_) returns()
func (_ConcreteVault *ConcreteVaultSession) SetVaultFees(newFees_ VaultFees) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetVaultFees(&_ConcreteVault.TransactOpts, newFees_)
}

// SetVaultFees is a paid mutator transaction binding the contract method 0x6bedc46c.
//
// Solidity: function setVaultFees((uint64,uint64,uint64,(uint256,uint256,uint64)[]) newFees_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) SetVaultFees(newFees_ VaultFees) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetVaultFees(&_ConcreteVault.TransactOpts, newFees_)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x5337e670.
//
// Solidity: function setWithdrawalQueue(address withdrawalQueue_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) SetWithdrawalQueue(opts *bind.TransactOpts, withdrawalQueue_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "setWithdrawalQueue", withdrawalQueue_)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x5337e670.
//
// Solidity: function setWithdrawalQueue(address withdrawalQueue_) returns()
func (_ConcreteVault *ConcreteVaultSession) SetWithdrawalQueue(withdrawalQueue_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetWithdrawalQueue(&_ConcreteVault.TransactOpts, withdrawalQueue_)
}

// SetWithdrawalQueue is a paid mutator transaction binding the contract method 0x5337e670.
//
// Solidity: function setWithdrawalQueue(address withdrawalQueue_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) SetWithdrawalQueue(withdrawalQueue_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.SetWithdrawalQueue(&_ConcreteVault.TransactOpts, withdrawalQueue_)
}

// TakePortfolioAndProtocolFees is a paid mutator transaction binding the contract method 0x46571899.
//
// Solidity: function takePortfolioAndProtocolFees() returns()
func (_ConcreteVault *ConcreteVaultTransactor) TakePortfolioAndProtocolFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "takePortfolioAndProtocolFees")
}

// TakePortfolioAndProtocolFees is a paid mutator transaction binding the contract method 0x46571899.
//
// Solidity: function takePortfolioAndProtocolFees() returns()
func (_ConcreteVault *ConcreteVaultSession) TakePortfolioAndProtocolFees() (*types.Transaction, error) {
	return _ConcreteVault.Contract.TakePortfolioAndProtocolFees(&_ConcreteVault.TransactOpts)
}

// TakePortfolioAndProtocolFees is a paid mutator transaction binding the contract method 0x46571899.
//
// Solidity: function takePortfolioAndProtocolFees() returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) TakePortfolioAndProtocolFees() (*types.Transaction, error) {
	return _ConcreteVault.Contract.TakePortfolioAndProtocolFees(&_ConcreteVault.TransactOpts)
}

// ToggleVaultIdle is a paid mutator transaction binding the contract method 0x6e2a6fe5.
//
// Solidity: function toggleVaultIdle() returns()
func (_ConcreteVault *ConcreteVaultTransactor) ToggleVaultIdle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "toggleVaultIdle")
}

// ToggleVaultIdle is a paid mutator transaction binding the contract method 0x6e2a6fe5.
//
// Solidity: function toggleVaultIdle() returns()
func (_ConcreteVault *ConcreteVaultSession) ToggleVaultIdle() (*types.Transaction, error) {
	return _ConcreteVault.Contract.ToggleVaultIdle(&_ConcreteVault.TransactOpts)
}

// ToggleVaultIdle is a paid mutator transaction binding the contract method 0x6e2a6fe5.
//
// Solidity: function toggleVaultIdle() returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) ToggleVaultIdle() (*types.Transaction, error) {
	return _ConcreteVault.Contract.ToggleVaultIdle(&_ConcreteVault.TransactOpts)
}

// ToggleWithdrawalsPaused is a paid mutator transaction binding the contract method 0x75a046a5.
//
// Solidity: function toggleWithdrawalsPaused(bool withdrawalsPaused_) returns()
func (_ConcreteVault *ConcreteVaultTransactor) ToggleWithdrawalsPaused(opts *bind.TransactOpts, withdrawalsPaused_ bool) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "toggleWithdrawalsPaused", withdrawalsPaused_)
}

// ToggleWithdrawalsPaused is a paid mutator transaction binding the contract method 0x75a046a5.
//
// Solidity: function toggleWithdrawalsPaused(bool withdrawalsPaused_) returns()
func (_ConcreteVault *ConcreteVaultSession) ToggleWithdrawalsPaused(withdrawalsPaused_ bool) (*types.Transaction, error) {
	return _ConcreteVault.Contract.ToggleWithdrawalsPaused(&_ConcreteVault.TransactOpts, withdrawalsPaused_)
}

// ToggleWithdrawalsPaused is a paid mutator transaction binding the contract method 0x75a046a5.
//
// Solidity: function toggleWithdrawalsPaused(bool withdrawalsPaused_) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) ToggleWithdrawalsPaused(withdrawalsPaused_ bool) (*types.Transaction, error) {
	return _ConcreteVault.Contract.ToggleWithdrawalsPaused(&_ConcreteVault.TransactOpts, withdrawalsPaused_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ConcreteVault *ConcreteVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ConcreteVault *ConcreteVaultSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Transfer(&_ConcreteVault.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ConcreteVault *ConcreteVaultTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Transfer(&_ConcreteVault.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ConcreteVault *ConcreteVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ConcreteVault *ConcreteVaultSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.TransferFrom(&_ConcreteVault.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ConcreteVault *ConcreteVaultTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.TransferFrom(&_ConcreteVault.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConcreteVault *ConcreteVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConcreteVault *ConcreteVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.TransferOwnership(&_ConcreteVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.TransferOwnership(&_ConcreteVault.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ConcreteVault *ConcreteVaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ConcreteVault *ConcreteVaultSession) Unpause() (*types.Transaction, error) {
	return _ConcreteVault.Contract.Unpause(&_ConcreteVault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ConcreteVault *ConcreteVaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _ConcreteVault.Contract.Unpause(&_ConcreteVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 assets_) returns(uint256)
func (_ConcreteVault *ConcreteVaultTransactor) Withdraw(opts *bind.TransactOpts, assets_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "withdraw", assets_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 assets_) returns(uint256)
func (_ConcreteVault *ConcreteVaultSession) Withdraw(assets_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Withdraw(&_ConcreteVault.TransactOpts, assets_)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 assets_) returns(uint256)
func (_ConcreteVault *ConcreteVaultTransactorSession) Withdraw(assets_ *big.Int) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Withdraw(&_ConcreteVault.TransactOpts, assets_)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets_, address receiver_, address owner_) returns(uint256 shares)
func (_ConcreteVault *ConcreteVaultTransactor) Withdraw0(opts *bind.TransactOpts, assets_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.contract.Transact(opts, "withdraw0", assets_, receiver_, owner_)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets_, address receiver_, address owner_) returns(uint256 shares)
func (_ConcreteVault *ConcreteVaultSession) Withdraw0(assets_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Withdraw0(&_ConcreteVault.TransactOpts, assets_, receiver_, owner_)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets_, address receiver_, address owner_) returns(uint256 shares)
func (_ConcreteVault *ConcreteVaultTransactorSession) Withdraw0(assets_ *big.Int, receiver_ common.Address, owner_ common.Address) (*types.Transaction, error) {
	return _ConcreteVault.Contract.Withdraw0(&_ConcreteVault.TransactOpts, assets_, receiver_, owner_)
}

// ConcreteVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ConcreteVault contract.
type ConcreteVaultApprovalIterator struct {
	Event *ConcreteVaultApproval // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultApproval)
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
		it.Event = new(ConcreteVaultApproval)
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
func (it *ConcreteVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultApproval represents a Approval event raised by the ConcreteVault contract.
type ConcreteVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ConcreteVault *ConcreteVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ConcreteVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultApprovalIterator{contract: _ConcreteVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ConcreteVault *ConcreteVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ConcreteVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultApproval)
				if err := _ConcreteVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ConcreteVault *ConcreteVaultFilterer) ParseApproval(log types.Log) (*ConcreteVaultApproval, error) {
	event := new(ConcreteVaultApproval)
	if err := _ConcreteVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ConcreteVault contract.
type ConcreteVaultDepositIterator struct {
	Event *ConcreteVaultDeposit // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultDeposit)
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
		it.Event = new(ConcreteVaultDeposit)
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
func (it *ConcreteVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultDeposit represents a Deposit event raised by the ConcreteVault contract.
type ConcreteVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ConcreteVault *ConcreteVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*ConcreteVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultDepositIterator{contract: _ConcreteVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ConcreteVault *ConcreteVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ConcreteVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultDeposit)
				if err := _ConcreteVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_ConcreteVault *ConcreteVaultFilterer) ParseDeposit(log types.Log) (*ConcreteVaultDeposit, error) {
	event := new(ConcreteVaultDeposit)
	if err := _ConcreteVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultDepositLimitSetIterator is returned from FilterDepositLimitSet and is used to iterate over the raw logs and unpacked data for DepositLimitSet events raised by the ConcreteVault contract.
type ConcreteVaultDepositLimitSetIterator struct {
	Event *ConcreteVaultDepositLimitSet // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultDepositLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultDepositLimitSet)
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
		it.Event = new(ConcreteVaultDepositLimitSet)
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
func (it *ConcreteVaultDepositLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultDepositLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultDepositLimitSet represents a DepositLimitSet event raised by the ConcreteVault contract.
type ConcreteVaultDepositLimitSet struct {
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDepositLimitSet is a free log retrieval operation binding the contract event 0x5d2e73196f8ba1b44e887e2bcc9bcd1f3e657add341d4a0a632ffff00d6903f2.
//
// Solidity: event DepositLimitSet(uint256 limit)
func (_ConcreteVault *ConcreteVaultFilterer) FilterDepositLimitSet(opts *bind.FilterOpts) (*ConcreteVaultDepositLimitSetIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "DepositLimitSet")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultDepositLimitSetIterator{contract: _ConcreteVault.contract, event: "DepositLimitSet", logs: logs, sub: sub}, nil
}

// WatchDepositLimitSet is a free log subscription operation binding the contract event 0x5d2e73196f8ba1b44e887e2bcc9bcd1f3e657add341d4a0a632ffff00d6903f2.
//
// Solidity: event DepositLimitSet(uint256 limit)
func (_ConcreteVault *ConcreteVaultFilterer) WatchDepositLimitSet(opts *bind.WatchOpts, sink chan<- *ConcreteVaultDepositLimitSet) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "DepositLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultDepositLimitSet)
				if err := _ConcreteVault.contract.UnpackLog(event, "DepositLimitSet", log); err != nil {
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

// ParseDepositLimitSet is a log parse operation binding the contract event 0x5d2e73196f8ba1b44e887e2bcc9bcd1f3e657add341d4a0a632ffff00d6903f2.
//
// Solidity: event DepositLimitSet(uint256 limit)
func (_ConcreteVault *ConcreteVaultFilterer) ParseDepositLimitSet(log types.Log) (*ConcreteVaultDepositLimitSet, error) {
	event := new(ConcreteVaultDepositLimitSet)
	if err := _ConcreteVault.contract.UnpackLog(event, "DepositLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultFeeRecipientUpdatedIterator is returned from FilterFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientUpdated events raised by the ConcreteVault contract.
type ConcreteVaultFeeRecipientUpdatedIterator struct {
	Event *ConcreteVaultFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultFeeRecipientUpdated)
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
		it.Event = new(ConcreteVaultFeeRecipientUpdated)
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
func (it *ConcreteVaultFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the ConcreteVault contract.
type ConcreteVaultFeeRecipientUpdated struct {
	OldRecipient common.Address
	NewRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientUpdated is a free log retrieval operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed oldRecipient, address indexed newRecipient)
func (_ConcreteVault *ConcreteVaultFilterer) FilterFeeRecipientUpdated(opts *bind.FilterOpts, oldRecipient []common.Address, newRecipient []common.Address) (*ConcreteVaultFeeRecipientUpdatedIterator, error) {

	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "FeeRecipientUpdated", oldRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultFeeRecipientUpdatedIterator{contract: _ConcreteVault.contract, event: "FeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientUpdated is a free log subscription operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed oldRecipient, address indexed newRecipient)
func (_ConcreteVault *ConcreteVaultFilterer) WatchFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *ConcreteVaultFeeRecipientUpdated, oldRecipient []common.Address, newRecipient []common.Address) (event.Subscription, error) {

	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "FeeRecipientUpdated", oldRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultFeeRecipientUpdated)
				if err := _ConcreteVault.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
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

// ParseFeeRecipientUpdated is a log parse operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed oldRecipient, address indexed newRecipient)
func (_ConcreteVault *ConcreteVaultFilterer) ParseFeeRecipientUpdated(log types.Log) (*ConcreteVaultFeeRecipientUpdated, error) {
	event := new(ConcreteVaultFeeRecipientUpdated)
	if err := _ConcreteVault.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ConcreteVault contract.
type ConcreteVaultInitializedIterator struct {
	Event *ConcreteVaultInitialized // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultInitialized)
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
		it.Event = new(ConcreteVaultInitialized)
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
func (it *ConcreteVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultInitialized represents a Initialized event raised by the ConcreteVault contract.
type ConcreteVaultInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ConcreteVault *ConcreteVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*ConcreteVaultInitializedIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultInitializedIterator{contract: _ConcreteVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ConcreteVault *ConcreteVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ConcreteVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultInitialized)
				if err := _ConcreteVault.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ConcreteVault *ConcreteVaultFilterer) ParseInitialized(log types.Log) (*ConcreteVaultInitialized, error) {
	event := new(ConcreteVaultInitialized)
	if err := _ConcreteVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultInitialized0Iterator is returned from FilterInitialized0 and is used to iterate over the raw logs and unpacked data for Initialized0 events raised by the ConcreteVault contract.
type ConcreteVaultInitialized0Iterator struct {
	Event *ConcreteVaultInitialized0 // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultInitialized0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultInitialized0)
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
		it.Event = new(ConcreteVaultInitialized0)
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
func (it *ConcreteVaultInitialized0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultInitialized0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultInitialized0 represents a Initialized0 event raised by the ConcreteVault contract.
type ConcreteVaultInitialized0 struct {
	VaultName       common.Address
	UnderlyingAsset common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInitialized0 is a free log retrieval operation binding the contract event 0x3cd5ec01b1ae7cfec6ca1863e2cd6aa25d6d1702825803ff2b7cc95010fffdc2.
//
// Solidity: event Initialized(address indexed vaultName, address indexed underlyingAsset)
func (_ConcreteVault *ConcreteVaultFilterer) FilterInitialized0(opts *bind.FilterOpts, vaultName []common.Address, underlyingAsset []common.Address) (*ConcreteVaultInitialized0Iterator, error) {

	var vaultNameRule []interface{}
	for _, vaultNameItem := range vaultName {
		vaultNameRule = append(vaultNameRule, vaultNameItem)
	}
	var underlyingAssetRule []interface{}
	for _, underlyingAssetItem := range underlyingAsset {
		underlyingAssetRule = append(underlyingAssetRule, underlyingAssetItem)
	}

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "Initialized0", vaultNameRule, underlyingAssetRule)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultInitialized0Iterator{contract: _ConcreteVault.contract, event: "Initialized0", logs: logs, sub: sub}, nil
}

// WatchInitialized0 is a free log subscription operation binding the contract event 0x3cd5ec01b1ae7cfec6ca1863e2cd6aa25d6d1702825803ff2b7cc95010fffdc2.
//
// Solidity: event Initialized(address indexed vaultName, address indexed underlyingAsset)
func (_ConcreteVault *ConcreteVaultFilterer) WatchInitialized0(opts *bind.WatchOpts, sink chan<- *ConcreteVaultInitialized0, vaultName []common.Address, underlyingAsset []common.Address) (event.Subscription, error) {

	var vaultNameRule []interface{}
	for _, vaultNameItem := range vaultName {
		vaultNameRule = append(vaultNameRule, vaultNameItem)
	}
	var underlyingAssetRule []interface{}
	for _, underlyingAssetItem := range underlyingAsset {
		underlyingAssetRule = append(underlyingAssetRule, underlyingAssetItem)
	}

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "Initialized0", vaultNameRule, underlyingAssetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultInitialized0)
				if err := _ConcreteVault.contract.UnpackLog(event, "Initialized0", log); err != nil {
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

// ParseInitialized0 is a log parse operation binding the contract event 0x3cd5ec01b1ae7cfec6ca1863e2cd6aa25d6d1702825803ff2b7cc95010fffdc2.
//
// Solidity: event Initialized(address indexed vaultName, address indexed underlyingAsset)
func (_ConcreteVault *ConcreteVaultFilterer) ParseInitialized0(log types.Log) (*ConcreteVaultInitialized0, error) {
	event := new(ConcreteVaultInitialized0)
	if err := _ConcreteVault.contract.UnpackLog(event, "Initialized0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultMinimunQueueRequestUpdatedIterator is returned from FilterMinimunQueueRequestUpdated and is used to iterate over the raw logs and unpacked data for MinimunQueueRequestUpdated events raised by the ConcreteVault contract.
type ConcreteVaultMinimunQueueRequestUpdatedIterator struct {
	Event *ConcreteVaultMinimunQueueRequestUpdated // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultMinimunQueueRequestUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultMinimunQueueRequestUpdated)
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
		it.Event = new(ConcreteVaultMinimunQueueRequestUpdated)
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
func (it *ConcreteVaultMinimunQueueRequestUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultMinimunQueueRequestUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultMinimunQueueRequestUpdated represents a MinimunQueueRequestUpdated event raised by the ConcreteVault contract.
type ConcreteVaultMinimunQueueRequestUpdated struct {
	OldMinQueueRequest *big.Int
	NewMinQueueRequest *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMinimunQueueRequestUpdated is a free log retrieval operation binding the contract event 0x47ecb16e79d3795bad1b15b763dad4136c346b10a729a5e22a60ae3416b571ba.
//
// Solidity: event MinimunQueueRequestUpdated(uint256 _oldMinQueueRequest, uint256 _newMinQueueRequest)
func (_ConcreteVault *ConcreteVaultFilterer) FilterMinimunQueueRequestUpdated(opts *bind.FilterOpts) (*ConcreteVaultMinimunQueueRequestUpdatedIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "MinimunQueueRequestUpdated")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultMinimunQueueRequestUpdatedIterator{contract: _ConcreteVault.contract, event: "MinimunQueueRequestUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimunQueueRequestUpdated is a free log subscription operation binding the contract event 0x47ecb16e79d3795bad1b15b763dad4136c346b10a729a5e22a60ae3416b571ba.
//
// Solidity: event MinimunQueueRequestUpdated(uint256 _oldMinQueueRequest, uint256 _newMinQueueRequest)
func (_ConcreteVault *ConcreteVaultFilterer) WatchMinimunQueueRequestUpdated(opts *bind.WatchOpts, sink chan<- *ConcreteVaultMinimunQueueRequestUpdated) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "MinimunQueueRequestUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultMinimunQueueRequestUpdated)
				if err := _ConcreteVault.contract.UnpackLog(event, "MinimunQueueRequestUpdated", log); err != nil {
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

// ParseMinimunQueueRequestUpdated is a log parse operation binding the contract event 0x47ecb16e79d3795bad1b15b763dad4136c346b10a729a5e22a60ae3416b571ba.
//
// Solidity: event MinimunQueueRequestUpdated(uint256 _oldMinQueueRequest, uint256 _newMinQueueRequest)
func (_ConcreteVault *ConcreteVaultFilterer) ParseMinimunQueueRequestUpdated(log types.Log) (*ConcreteVaultMinimunQueueRequestUpdated, error) {
	event := new(ConcreteVaultMinimunQueueRequestUpdated)
	if err := _ConcreteVault.contract.UnpackLog(event, "MinimunQueueRequestUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ConcreteVault contract.
type ConcreteVaultOwnershipTransferredIterator struct {
	Event *ConcreteVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultOwnershipTransferred)
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
		it.Event = new(ConcreteVaultOwnershipTransferred)
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
func (it *ConcreteVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultOwnershipTransferred represents a OwnershipTransferred event raised by the ConcreteVault contract.
type ConcreteVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConcreteVault *ConcreteVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ConcreteVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultOwnershipTransferredIterator{contract: _ConcreteVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ConcreteVault *ConcreteVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ConcreteVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultOwnershipTransferred)
				if err := _ConcreteVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ConcreteVault *ConcreteVaultFilterer) ParseOwnershipTransferred(log types.Log) (*ConcreteVaultOwnershipTransferred, error) {
	event := new(ConcreteVaultOwnershipTransferred)
	if err := _ConcreteVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultParkingLotUpdatedIterator is returned from FilterParkingLotUpdated and is used to iterate over the raw logs and unpacked data for ParkingLotUpdated events raised by the ConcreteVault contract.
type ConcreteVaultParkingLotUpdatedIterator struct {
	Event *ConcreteVaultParkingLotUpdated // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultParkingLotUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultParkingLotUpdated)
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
		it.Event = new(ConcreteVaultParkingLotUpdated)
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
func (it *ConcreteVaultParkingLotUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultParkingLotUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultParkingLotUpdated represents a ParkingLotUpdated event raised by the ConcreteVault contract.
type ConcreteVaultParkingLotUpdated struct {
	OldParkingLot      common.Address
	NewParkingLot      common.Address
	SuccessfulApproval bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterParkingLotUpdated is a free log retrieval operation binding the contract event 0x8a8682c8e0b406a8cfcc7c7c4ac90372d088ddb8dfaf1baea5435b69e832a43a.
//
// Solidity: event ParkingLotUpdated(address indexed oldParkingLot, address indexed newParkingLot, bool successfulApproval)
func (_ConcreteVault *ConcreteVaultFilterer) FilterParkingLotUpdated(opts *bind.FilterOpts, oldParkingLot []common.Address, newParkingLot []common.Address) (*ConcreteVaultParkingLotUpdatedIterator, error) {

	var oldParkingLotRule []interface{}
	for _, oldParkingLotItem := range oldParkingLot {
		oldParkingLotRule = append(oldParkingLotRule, oldParkingLotItem)
	}
	var newParkingLotRule []interface{}
	for _, newParkingLotItem := range newParkingLot {
		newParkingLotRule = append(newParkingLotRule, newParkingLotItem)
	}

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "ParkingLotUpdated", oldParkingLotRule, newParkingLotRule)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultParkingLotUpdatedIterator{contract: _ConcreteVault.contract, event: "ParkingLotUpdated", logs: logs, sub: sub}, nil
}

// WatchParkingLotUpdated is a free log subscription operation binding the contract event 0x8a8682c8e0b406a8cfcc7c7c4ac90372d088ddb8dfaf1baea5435b69e832a43a.
//
// Solidity: event ParkingLotUpdated(address indexed oldParkingLot, address indexed newParkingLot, bool successfulApproval)
func (_ConcreteVault *ConcreteVaultFilterer) WatchParkingLotUpdated(opts *bind.WatchOpts, sink chan<- *ConcreteVaultParkingLotUpdated, oldParkingLot []common.Address, newParkingLot []common.Address) (event.Subscription, error) {

	var oldParkingLotRule []interface{}
	for _, oldParkingLotItem := range oldParkingLot {
		oldParkingLotRule = append(oldParkingLotRule, oldParkingLotItem)
	}
	var newParkingLotRule []interface{}
	for _, newParkingLotItem := range newParkingLot {
		newParkingLotRule = append(newParkingLotRule, newParkingLotItem)
	}

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "ParkingLotUpdated", oldParkingLotRule, newParkingLotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultParkingLotUpdated)
				if err := _ConcreteVault.contract.UnpackLog(event, "ParkingLotUpdated", log); err != nil {
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

// ParseParkingLotUpdated is a log parse operation binding the contract event 0x8a8682c8e0b406a8cfcc7c7c4ac90372d088ddb8dfaf1baea5435b69e832a43a.
//
// Solidity: event ParkingLotUpdated(address indexed oldParkingLot, address indexed newParkingLot, bool successfulApproval)
func (_ConcreteVault *ConcreteVaultFilterer) ParseParkingLotUpdated(log types.Log) (*ConcreteVaultParkingLotUpdated, error) {
	event := new(ConcreteVaultParkingLotUpdated)
	if err := _ConcreteVault.contract.UnpackLog(event, "ParkingLotUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ConcreteVault contract.
type ConcreteVaultPausedIterator struct {
	Event *ConcreteVaultPaused // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultPaused)
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
		it.Event = new(ConcreteVaultPaused)
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
func (it *ConcreteVaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultPaused represents a Paused event raised by the ConcreteVault contract.
type ConcreteVaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ConcreteVault *ConcreteVaultFilterer) FilterPaused(opts *bind.FilterOpts) (*ConcreteVaultPausedIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultPausedIterator{contract: _ConcreteVault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ConcreteVault *ConcreteVaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ConcreteVaultPaused) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultPaused)
				if err := _ConcreteVault.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_ConcreteVault *ConcreteVaultFilterer) ParsePaused(log types.Log) (*ConcreteVaultPaused, error) {
	event := new(ConcreteVaultPaused)
	if err := _ConcreteVault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultRequestedFundsIterator is returned from FilterRequestedFunds and is used to iterate over the raw logs and unpacked data for RequestedFunds events raised by the ConcreteVault contract.
type ConcreteVaultRequestedFundsIterator struct {
	Event *ConcreteVaultRequestedFunds // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultRequestedFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultRequestedFunds)
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
		it.Event = new(ConcreteVaultRequestedFunds)
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
func (it *ConcreteVaultRequestedFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultRequestedFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultRequestedFunds represents a RequestedFunds event raised by the ConcreteVault contract.
type ConcreteVaultRequestedFunds struct {
	ProtectStrategy common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRequestedFunds is a free log retrieval operation binding the contract event 0x986c3f7d8e5489362da2e0e75ed9d0aa29fbf30326fc391ae913868a4730bdb8.
//
// Solidity: event RequestedFunds(address indexed protectStrategy, uint256 amount)
func (_ConcreteVault *ConcreteVaultFilterer) FilterRequestedFunds(opts *bind.FilterOpts, protectStrategy []common.Address) (*ConcreteVaultRequestedFundsIterator, error) {

	var protectStrategyRule []interface{}
	for _, protectStrategyItem := range protectStrategy {
		protectStrategyRule = append(protectStrategyRule, protectStrategyItem)
	}

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "RequestedFunds", protectStrategyRule)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultRequestedFundsIterator{contract: _ConcreteVault.contract, event: "RequestedFunds", logs: logs, sub: sub}, nil
}

// WatchRequestedFunds is a free log subscription operation binding the contract event 0x986c3f7d8e5489362da2e0e75ed9d0aa29fbf30326fc391ae913868a4730bdb8.
//
// Solidity: event RequestedFunds(address indexed protectStrategy, uint256 amount)
func (_ConcreteVault *ConcreteVaultFilterer) WatchRequestedFunds(opts *bind.WatchOpts, sink chan<- *ConcreteVaultRequestedFunds, protectStrategy []common.Address) (event.Subscription, error) {

	var protectStrategyRule []interface{}
	for _, protectStrategyItem := range protectStrategy {
		protectStrategyRule = append(protectStrategyRule, protectStrategyItem)
	}

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "RequestedFunds", protectStrategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultRequestedFunds)
				if err := _ConcreteVault.contract.UnpackLog(event, "RequestedFunds", log); err != nil {
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

// ParseRequestedFunds is a log parse operation binding the contract event 0x986c3f7d8e5489362da2e0e75ed9d0aa29fbf30326fc391ae913868a4730bdb8.
//
// Solidity: event RequestedFunds(address indexed protectStrategy, uint256 amount)
func (_ConcreteVault *ConcreteVaultFilterer) ParseRequestedFunds(log types.Log) (*ConcreteVaultRequestedFunds, error) {
	event := new(ConcreteVaultRequestedFunds)
	if err := _ConcreteVault.contract.UnpackLog(event, "RequestedFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultRewardsHarvestedIterator is returned from FilterRewardsHarvested and is used to iterate over the raw logs and unpacked data for RewardsHarvested events raised by the ConcreteVault contract.
type ConcreteVaultRewardsHarvestedIterator struct {
	Event *ConcreteVaultRewardsHarvested // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultRewardsHarvestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultRewardsHarvested)
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
		it.Event = new(ConcreteVaultRewardsHarvested)
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
func (it *ConcreteVaultRewardsHarvestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultRewardsHarvestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultRewardsHarvested represents a RewardsHarvested event raised by the ConcreteVault contract.
type ConcreteVaultRewardsHarvested struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRewardsHarvested is a free log retrieval operation binding the contract event 0x429eabcd6b69e5ca232c6a9d2c5af1805e375e951df8e9bfe831b445e5f69b2b.
//
// Solidity: event RewardsHarvested()
func (_ConcreteVault *ConcreteVaultFilterer) FilterRewardsHarvested(opts *bind.FilterOpts) (*ConcreteVaultRewardsHarvestedIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "RewardsHarvested")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultRewardsHarvestedIterator{contract: _ConcreteVault.contract, event: "RewardsHarvested", logs: logs, sub: sub}, nil
}

// WatchRewardsHarvested is a free log subscription operation binding the contract event 0x429eabcd6b69e5ca232c6a9d2c5af1805e375e951df8e9bfe831b445e5f69b2b.
//
// Solidity: event RewardsHarvested()
func (_ConcreteVault *ConcreteVaultFilterer) WatchRewardsHarvested(opts *bind.WatchOpts, sink chan<- *ConcreteVaultRewardsHarvested) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "RewardsHarvested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultRewardsHarvested)
				if err := _ConcreteVault.contract.UnpackLog(event, "RewardsHarvested", log); err != nil {
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

// ParseRewardsHarvested is a log parse operation binding the contract event 0x429eabcd6b69e5ca232c6a9d2c5af1805e375e951df8e9bfe831b445e5f69b2b.
//
// Solidity: event RewardsHarvested()
func (_ConcreteVault *ConcreteVaultFilterer) ParseRewardsHarvested(log types.Log) (*ConcreteVaultRewardsHarvested, error) {
	event := new(ConcreteVaultRewardsHarvested)
	if err := _ConcreteVault.contract.UnpackLog(event, "RewardsHarvested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultStrategyAddedIterator is returned from FilterStrategyAdded and is used to iterate over the raw logs and unpacked data for StrategyAdded events raised by the ConcreteVault contract.
type ConcreteVaultStrategyAddedIterator struct {
	Event *ConcreteVaultStrategyAdded // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultStrategyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultStrategyAdded)
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
		it.Event = new(ConcreteVaultStrategyAdded)
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
func (it *ConcreteVaultStrategyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultStrategyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultStrategyAdded represents a StrategyAdded event raised by the ConcreteVault contract.
type ConcreteVaultStrategyAdded struct {
	NewStrategy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyAdded is a free log retrieval operation binding the contract event 0x3f008fd510eae7a9e7bee13513d7b83bef8003d488b5a3d0b0da4de71d6846f1.
//
// Solidity: event StrategyAdded(address newStrategy)
func (_ConcreteVault *ConcreteVaultFilterer) FilterStrategyAdded(opts *bind.FilterOpts) (*ConcreteVaultStrategyAddedIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "StrategyAdded")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultStrategyAddedIterator{contract: _ConcreteVault.contract, event: "StrategyAdded", logs: logs, sub: sub}, nil
}

// WatchStrategyAdded is a free log subscription operation binding the contract event 0x3f008fd510eae7a9e7bee13513d7b83bef8003d488b5a3d0b0da4de71d6846f1.
//
// Solidity: event StrategyAdded(address newStrategy)
func (_ConcreteVault *ConcreteVaultFilterer) WatchStrategyAdded(opts *bind.WatchOpts, sink chan<- *ConcreteVaultStrategyAdded) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "StrategyAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultStrategyAdded)
				if err := _ConcreteVault.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
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

// ParseStrategyAdded is a log parse operation binding the contract event 0x3f008fd510eae7a9e7bee13513d7b83bef8003d488b5a3d0b0da4de71d6846f1.
//
// Solidity: event StrategyAdded(address newStrategy)
func (_ConcreteVault *ConcreteVaultFilterer) ParseStrategyAdded(log types.Log) (*ConcreteVaultStrategyAdded, error) {
	event := new(ConcreteVaultStrategyAdded)
	if err := _ConcreteVault.contract.UnpackLog(event, "StrategyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultStrategyAllocationsChangedIterator is returned from FilterStrategyAllocationsChanged and is used to iterate over the raw logs and unpacked data for StrategyAllocationsChanged events raised by the ConcreteVault contract.
type ConcreteVaultStrategyAllocationsChangedIterator struct {
	Event *ConcreteVaultStrategyAllocationsChanged // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultStrategyAllocationsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultStrategyAllocationsChanged)
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
		it.Event = new(ConcreteVaultStrategyAllocationsChanged)
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
func (it *ConcreteVaultStrategyAllocationsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultStrategyAllocationsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultStrategyAllocationsChanged represents a StrategyAllocationsChanged event raised by the ConcreteVault contract.
type ConcreteVaultStrategyAllocationsChanged struct {
	NewAllocations []Allocation
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyAllocationsChanged is a free log retrieval operation binding the contract event 0xb02d56028d2fba1da0ff6a7619d01fd6a273956cc6f10d93112c0e764c519661.
//
// Solidity: event StrategyAllocationsChanged((uint256,uint256)[] newAllocations)
func (_ConcreteVault *ConcreteVaultFilterer) FilterStrategyAllocationsChanged(opts *bind.FilterOpts) (*ConcreteVaultStrategyAllocationsChangedIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "StrategyAllocationsChanged")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultStrategyAllocationsChangedIterator{contract: _ConcreteVault.contract, event: "StrategyAllocationsChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyAllocationsChanged is a free log subscription operation binding the contract event 0xb02d56028d2fba1da0ff6a7619d01fd6a273956cc6f10d93112c0e764c519661.
//
// Solidity: event StrategyAllocationsChanged((uint256,uint256)[] newAllocations)
func (_ConcreteVault *ConcreteVaultFilterer) WatchStrategyAllocationsChanged(opts *bind.WatchOpts, sink chan<- *ConcreteVaultStrategyAllocationsChanged) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "StrategyAllocationsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultStrategyAllocationsChanged)
				if err := _ConcreteVault.contract.UnpackLog(event, "StrategyAllocationsChanged", log); err != nil {
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

// ParseStrategyAllocationsChanged is a log parse operation binding the contract event 0xb02d56028d2fba1da0ff6a7619d01fd6a273956cc6f10d93112c0e764c519661.
//
// Solidity: event StrategyAllocationsChanged((uint256,uint256)[] newAllocations)
func (_ConcreteVault *ConcreteVaultFilterer) ParseStrategyAllocationsChanged(log types.Log) (*ConcreteVaultStrategyAllocationsChanged, error) {
	event := new(ConcreteVaultStrategyAllocationsChanged)
	if err := _ConcreteVault.contract.UnpackLog(event, "StrategyAllocationsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultStrategyRemovedIterator is returned from FilterStrategyRemoved and is used to iterate over the raw logs and unpacked data for StrategyRemoved events raised by the ConcreteVault contract.
type ConcreteVaultStrategyRemovedIterator struct {
	Event *ConcreteVaultStrategyRemoved // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultStrategyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultStrategyRemoved)
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
		it.Event = new(ConcreteVaultStrategyRemoved)
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
func (it *ConcreteVaultStrategyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultStrategyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultStrategyRemoved represents a StrategyRemoved event raised by the ConcreteVault contract.
type ConcreteVaultStrategyRemoved struct {
	OldStrategy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStrategyRemoved is a free log retrieval operation binding the contract event 0x09a1db4b80c32706328728508c941a6b954f31eb5affd32f236c1fd405f8fea4.
//
// Solidity: event StrategyRemoved(address oldStrategy)
func (_ConcreteVault *ConcreteVaultFilterer) FilterStrategyRemoved(opts *bind.FilterOpts) (*ConcreteVaultStrategyRemovedIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "StrategyRemoved")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultStrategyRemovedIterator{contract: _ConcreteVault.contract, event: "StrategyRemoved", logs: logs, sub: sub}, nil
}

// WatchStrategyRemoved is a free log subscription operation binding the contract event 0x09a1db4b80c32706328728508c941a6b954f31eb5affd32f236c1fd405f8fea4.
//
// Solidity: event StrategyRemoved(address oldStrategy)
func (_ConcreteVault *ConcreteVaultFilterer) WatchStrategyRemoved(opts *bind.WatchOpts, sink chan<- *ConcreteVaultStrategyRemoved) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "StrategyRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultStrategyRemoved)
				if err := _ConcreteVault.contract.UnpackLog(event, "StrategyRemoved", log); err != nil {
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

// ParseStrategyRemoved is a log parse operation binding the contract event 0x09a1db4b80c32706328728508c941a6b954f31eb5affd32f236c1fd405f8fea4.
//
// Solidity: event StrategyRemoved(address oldStrategy)
func (_ConcreteVault *ConcreteVaultFilterer) ParseStrategyRemoved(log types.Log) (*ConcreteVaultStrategyRemoved, error) {
	event := new(ConcreteVaultStrategyRemoved)
	if err := _ConcreteVault.contract.UnpackLog(event, "StrategyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultToggleVaultIdleIterator is returned from FilterToggleVaultIdle and is used to iterate over the raw logs and unpacked data for ToggleVaultIdle events raised by the ConcreteVault contract.
type ConcreteVaultToggleVaultIdleIterator struct {
	Event *ConcreteVaultToggleVaultIdle // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultToggleVaultIdleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultToggleVaultIdle)
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
		it.Event = new(ConcreteVaultToggleVaultIdle)
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
func (it *ConcreteVaultToggleVaultIdleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultToggleVaultIdleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultToggleVaultIdle represents a ToggleVaultIdle event raised by the ConcreteVault contract.
type ConcreteVaultToggleVaultIdle struct {
	PastValue bool
	NewValue  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterToggleVaultIdle is a free log retrieval operation binding the contract event 0x40f0175b5fde6802a96ba1ba9eeaaa9564b53a1c021b6fcde51422b574abcdbd.
//
// Solidity: event ToggleVaultIdle(bool pastValue, bool newValue)
func (_ConcreteVault *ConcreteVaultFilterer) FilterToggleVaultIdle(opts *bind.FilterOpts) (*ConcreteVaultToggleVaultIdleIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "ToggleVaultIdle")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultToggleVaultIdleIterator{contract: _ConcreteVault.contract, event: "ToggleVaultIdle", logs: logs, sub: sub}, nil
}

// WatchToggleVaultIdle is a free log subscription operation binding the contract event 0x40f0175b5fde6802a96ba1ba9eeaaa9564b53a1c021b6fcde51422b574abcdbd.
//
// Solidity: event ToggleVaultIdle(bool pastValue, bool newValue)
func (_ConcreteVault *ConcreteVaultFilterer) WatchToggleVaultIdle(opts *bind.WatchOpts, sink chan<- *ConcreteVaultToggleVaultIdle) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "ToggleVaultIdle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultToggleVaultIdle)
				if err := _ConcreteVault.contract.UnpackLog(event, "ToggleVaultIdle", log); err != nil {
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

// ParseToggleVaultIdle is a log parse operation binding the contract event 0x40f0175b5fde6802a96ba1ba9eeaaa9564b53a1c021b6fcde51422b574abcdbd.
//
// Solidity: event ToggleVaultIdle(bool pastValue, bool newValue)
func (_ConcreteVault *ConcreteVaultFilterer) ParseToggleVaultIdle(log types.Log) (*ConcreteVaultToggleVaultIdle, error) {
	event := new(ConcreteVaultToggleVaultIdle)
	if err := _ConcreteVault.contract.UnpackLog(event, "ToggleVaultIdle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ConcreteVault contract.
type ConcreteVaultTransferIterator struct {
	Event *ConcreteVaultTransfer // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultTransfer)
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
		it.Event = new(ConcreteVaultTransfer)
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
func (it *ConcreteVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultTransfer represents a Transfer event raised by the ConcreteVault contract.
type ConcreteVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ConcreteVault *ConcreteVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ConcreteVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultTransferIterator{contract: _ConcreteVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ConcreteVault *ConcreteVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ConcreteVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultTransfer)
				if err := _ConcreteVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ConcreteVault *ConcreteVaultFilterer) ParseTransfer(log types.Log) (*ConcreteVaultTransfer, error) {
	event := new(ConcreteVaultTransfer)
	if err := _ConcreteVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ConcreteVault contract.
type ConcreteVaultUnpausedIterator struct {
	Event *ConcreteVaultUnpaused // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultUnpaused)
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
		it.Event = new(ConcreteVaultUnpaused)
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
func (it *ConcreteVaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultUnpaused represents a Unpaused event raised by the ConcreteVault contract.
type ConcreteVaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ConcreteVault *ConcreteVaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ConcreteVaultUnpausedIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultUnpausedIterator{contract: _ConcreteVault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ConcreteVault *ConcreteVaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ConcreteVaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultUnpaused)
				if err := _ConcreteVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_ConcreteVault *ConcreteVaultFilterer) ParseUnpaused(log types.Log) (*ConcreteVaultUnpaused, error) {
	event := new(ConcreteVaultUnpaused)
	if err := _ConcreteVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ConcreteVault contract.
type ConcreteVaultWithdrawIterator struct {
	Event *ConcreteVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultWithdraw)
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
		it.Event = new(ConcreteVaultWithdraw)
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
func (it *ConcreteVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultWithdraw represents a Withdraw event raised by the ConcreteVault contract.
type ConcreteVaultWithdraw struct {
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
func (_ConcreteVault *ConcreteVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*ConcreteVaultWithdrawIterator, error) {

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

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultWithdrawIterator{contract: _ConcreteVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ConcreteVault *ConcreteVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ConcreteVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultWithdraw)
				if err := _ConcreteVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_ConcreteVault *ConcreteVaultFilterer) ParseWithdraw(log types.Log) (*ConcreteVaultWithdraw, error) {
	event := new(ConcreteVaultWithdraw)
	if err := _ConcreteVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultWithdrawalPausedToggledIterator is returned from FilterWithdrawalPausedToggled and is used to iterate over the raw logs and unpacked data for WithdrawalPausedToggled events raised by the ConcreteVault contract.
type ConcreteVaultWithdrawalPausedToggledIterator struct {
	Event *ConcreteVaultWithdrawalPausedToggled // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultWithdrawalPausedToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultWithdrawalPausedToggled)
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
		it.Event = new(ConcreteVaultWithdrawalPausedToggled)
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
func (it *ConcreteVaultWithdrawalPausedToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultWithdrawalPausedToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultWithdrawalPausedToggled represents a WithdrawalPausedToggled event raised by the ConcreteVault contract.
type ConcreteVaultWithdrawalPausedToggled struct {
	PastValue bool
	NewValue  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalPausedToggled is a free log retrieval operation binding the contract event 0xfdc806f366660aa3e77ed7f1da541577007de0e4b7e0c5d351ae99d2d01cc8f3.
//
// Solidity: event WithdrawalPausedToggled(bool pastValue, bool newValue)
func (_ConcreteVault *ConcreteVaultFilterer) FilterWithdrawalPausedToggled(opts *bind.FilterOpts) (*ConcreteVaultWithdrawalPausedToggledIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "WithdrawalPausedToggled")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultWithdrawalPausedToggledIterator{contract: _ConcreteVault.contract, event: "WithdrawalPausedToggled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalPausedToggled is a free log subscription operation binding the contract event 0xfdc806f366660aa3e77ed7f1da541577007de0e4b7e0c5d351ae99d2d01cc8f3.
//
// Solidity: event WithdrawalPausedToggled(bool pastValue, bool newValue)
func (_ConcreteVault *ConcreteVaultFilterer) WatchWithdrawalPausedToggled(opts *bind.WatchOpts, sink chan<- *ConcreteVaultWithdrawalPausedToggled) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "WithdrawalPausedToggled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultWithdrawalPausedToggled)
				if err := _ConcreteVault.contract.UnpackLog(event, "WithdrawalPausedToggled", log); err != nil {
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

// ParseWithdrawalPausedToggled is a log parse operation binding the contract event 0xfdc806f366660aa3e77ed7f1da541577007de0e4b7e0c5d351ae99d2d01cc8f3.
//
// Solidity: event WithdrawalPausedToggled(bool pastValue, bool newValue)
func (_ConcreteVault *ConcreteVaultFilterer) ParseWithdrawalPausedToggled(log types.Log) (*ConcreteVaultWithdrawalPausedToggled, error) {
	event := new(ConcreteVaultWithdrawalPausedToggled)
	if err := _ConcreteVault.contract.UnpackLog(event, "WithdrawalPausedToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConcreteVaultWithdrawalQueueUpdatedIterator is returned from FilterWithdrawalQueueUpdated and is used to iterate over the raw logs and unpacked data for WithdrawalQueueUpdated events raised by the ConcreteVault contract.
type ConcreteVaultWithdrawalQueueUpdatedIterator struct {
	Event *ConcreteVaultWithdrawalQueueUpdated // Event containing the contract specifics and raw log

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
func (it *ConcreteVaultWithdrawalQueueUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConcreteVaultWithdrawalQueueUpdated)
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
		it.Event = new(ConcreteVaultWithdrawalQueueUpdated)
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
func (it *ConcreteVaultWithdrawalQueueUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConcreteVaultWithdrawalQueueUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConcreteVaultWithdrawalQueueUpdated represents a WithdrawalQueueUpdated event raised by the ConcreteVault contract.
type ConcreteVaultWithdrawalQueueUpdated struct {
	OldQueue common.Address
	NewQueue common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalQueueUpdated is a free log retrieval operation binding the contract event 0x90ee8c46ab78acce150d0041456437c06fbd2a3642739b948bd373357e4ed699.
//
// Solidity: event WithdrawalQueueUpdated(address oldQueue, address newQueue)
func (_ConcreteVault *ConcreteVaultFilterer) FilterWithdrawalQueueUpdated(opts *bind.FilterOpts) (*ConcreteVaultWithdrawalQueueUpdatedIterator, error) {

	logs, sub, err := _ConcreteVault.contract.FilterLogs(opts, "WithdrawalQueueUpdated")
	if err != nil {
		return nil, err
	}
	return &ConcreteVaultWithdrawalQueueUpdatedIterator{contract: _ConcreteVault.contract, event: "WithdrawalQueueUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalQueueUpdated is a free log subscription operation binding the contract event 0x90ee8c46ab78acce150d0041456437c06fbd2a3642739b948bd373357e4ed699.
//
// Solidity: event WithdrawalQueueUpdated(address oldQueue, address newQueue)
func (_ConcreteVault *ConcreteVaultFilterer) WatchWithdrawalQueueUpdated(opts *bind.WatchOpts, sink chan<- *ConcreteVaultWithdrawalQueueUpdated) (event.Subscription, error) {

	logs, sub, err := _ConcreteVault.contract.WatchLogs(opts, "WithdrawalQueueUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConcreteVaultWithdrawalQueueUpdated)
				if err := _ConcreteVault.contract.UnpackLog(event, "WithdrawalQueueUpdated", log); err != nil {
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

// ParseWithdrawalQueueUpdated is a log parse operation binding the contract event 0x90ee8c46ab78acce150d0041456437c06fbd2a3642739b948bd373357e4ed699.
//
// Solidity: event WithdrawalQueueUpdated(address oldQueue, address newQueue)
func (_ConcreteVault *ConcreteVaultFilterer) ParseWithdrawalQueueUpdated(log types.Log) (*ConcreteVaultWithdrawalQueueUpdated, error) {
	event := new(ConcreteVaultWithdrawalQueueUpdated)
	if err := _ConcreteVault.contract.UnpackLog(event, "WithdrawalQueueUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
