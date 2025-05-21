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

// BaseIntegrations is an auto generated low-level Go binding around an user-defined struct.
type BaseIntegrations struct {
	Evc              common.Address
	ProtocolConfig   common.Address
	SequenceRegistry common.Address
	BalanceTracker   common.Address
	Permit2          common.Address
}

// DispatchDeployedModules is an auto generated low-level Go binding around an user-defined struct.
type DispatchDeployedModules struct {
	Initialize       common.Address
	Token            common.Address
	Vault            common.Address
	Borrowing        common.Address
	Liquidation      common.Address
	RiskManager      common.Address
	BalanceForwarder common.Address
	Governance       common.Address
}

// EulerVaultMetaData contains all meta data concerning the EulerVault contract.
var EulerVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"protocolConfig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequenceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"}],\"internalType\":\"structBase.Integrations\",\"name\":\"integrations\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"initialize\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowing\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"riskManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceForwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"internalType\":\"structDispatch.DeployedModules\",\"name\":\"modules\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"E_AccountLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_AmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadAssetReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadBorrowCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadMaxLiquidationDiscount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSharesOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSharesReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSupplyCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BorrowCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_CheckUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_CollateralDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ConfigAmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ControllerDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_DebtAmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_EmptyError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ExcessiveRepayAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_FlashLoanNotRepaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientCash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InvalidLTVAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LTVBorrow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LTVLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LiquidationCoolOff\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_MinYield\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NoLiability\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NoPriceOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotController\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotHookTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_OperationDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_OutstandingDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ProxyMetadata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_RepayTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SelfLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_TransientState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ViolatorLiquidityDeferred\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ZeroAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ZeroShares\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"BalanceForwarderStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolReceiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governorReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"governorShares\",\"type\":\"uint256\"}],\"name\":\"ConvertFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"DebtSocialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dToken\",\"type\":\"address\"}],\"name\":\"EVaultCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newSupplyCap\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newBorrowCap\",\"type\":\"uint16\"}],\"name\":\"GovSetCaps\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newConfigFlags\",\"type\":\"uint32\"}],\"name\":\"GovSetConfigFlags\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeReceiver\",\"type\":\"address\"}],\"name\":\"GovSetFeeReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernorAdmin\",\"type\":\"address\"}],\"name\":\"GovSetGovernorAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newHookTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newHookedOps\",\"type\":\"uint32\"}],\"name\":\"GovSetHookConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"}],\"name\":\"GovSetInterestFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"GovSetInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"initialLiquidationLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"targetTimestamp\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"name\":\"GovSetLTV\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newCoolOffTime\",\"type\":\"uint16\"}],\"name\":\"GovSetLiquidationCoolOffTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newDiscount\",\"type\":\"uint16\"}],\"name\":\"GovSetMaxLiquidationDiscount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"InterestAccrued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yieldBalance\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"PullDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accumulatedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"VaultStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EVC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVBorrow\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVFull\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"initialLiquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"targetTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVLiquidation\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LTVList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_BALANCE_FORWARDER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_BORROWING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_GOVERNANCE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_INITIALIZE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_LIQUIDATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_RISKMANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_VAULT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"liquidation\",\"type\":\"bool\"}],\"name\":\"accountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"liquidation\",\"type\":\"bool\"}],\"name\":\"accountLiquidityFull\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"collaterals\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"collateralValues\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedFeesAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceForwarderEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceTrackerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"caps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"supplyCap\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"borrowCap\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"collaterals\",\"type\":\"address[]\"}],\"name\":\"checkAccountStatus\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"checkLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRepay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxYield\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkVaultStatus\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configFlags\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convertFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtOfExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governorAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hookConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyCreator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestAccumulator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAssets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minYieldBalance\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationCoolOffTime\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidationDiscount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolConfigAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"pullDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"repayWithShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"supplyCap\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"borrowCap\",\"type\":\"uint16\"}],\"name\":\"setCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newConfigFlags\",\"type\":\"uint32\"}],\"name\":\"setConfigFlags\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernorAdmin\",\"type\":\"address\"}],\"name\":\"setGovernorAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newHookTarget\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newHookedOps\",\"type\":\"uint32\"}],\"name\":\"setHookConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"}],\"name\":\"setInterestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newModel\",\"type\":\"address\"}],\"name\":\"setInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"name\":\"setLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newCoolOffTime\",\"type\":\"uint16\"}],\"name\":\"setLiquidationCoolOffTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newDiscount\",\"type\":\"uint16\"}],\"name\":\"setMaxLiquidationDiscount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"touch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferFromMax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unitOfAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewDelegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EulerVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerVaultMetaData.ABI instead.
var EulerVaultABI = EulerVaultMetaData.ABI

// EulerVault is an auto generated Go binding around an Ethereum contract.
type EulerVault struct {
	EulerVaultCaller     // Read-only binding to the contract
	EulerVaultTransactor // Write-only binding to the contract
	EulerVaultFilterer   // Log filterer for contract events
}

// EulerVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerVaultSession struct {
	Contract     *EulerVault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EulerVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerVaultCallerSession struct {
	Contract *EulerVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EulerVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerVaultTransactorSession struct {
	Contract     *EulerVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EulerVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerVaultRaw struct {
	Contract *EulerVault // Generic contract binding to access the raw methods on
}

// EulerVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerVaultCallerRaw struct {
	Contract *EulerVaultCaller // Generic read-only contract binding to access the raw methods on
}

// EulerVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerVaultTransactorRaw struct {
	Contract *EulerVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerVault creates a new instance of EulerVault, bound to a specific deployed contract.
func NewEulerVault(address common.Address, backend bind.ContractBackend) (*EulerVault, error) {
	contract, err := bindEulerVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerVault{EulerVaultCaller: EulerVaultCaller{contract: contract}, EulerVaultTransactor: EulerVaultTransactor{contract: contract}, EulerVaultFilterer: EulerVaultFilterer{contract: contract}}, nil
}

// NewEulerVaultCaller creates a new read-only instance of EulerVault, bound to a specific deployed contract.
func NewEulerVaultCaller(address common.Address, caller bind.ContractCaller) (*EulerVaultCaller, error) {
	contract, err := bindEulerVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerVaultCaller{contract: contract}, nil
}

// NewEulerVaultTransactor creates a new write-only instance of EulerVault, bound to a specific deployed contract.
func NewEulerVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerVaultTransactor, error) {
	contract, err := bindEulerVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerVaultTransactor{contract: contract}, nil
}

// NewEulerVaultFilterer creates a new log filterer instance of EulerVault, bound to a specific deployed contract.
func NewEulerVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerVaultFilterer, error) {
	contract, err := bindEulerVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerVaultFilterer{contract: contract}, nil
}

// bindEulerVault binds a generic wrapper to an already deployed contract.
func bindEulerVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerVault *EulerVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerVault.Contract.EulerVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerVault *EulerVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVault.Contract.EulerVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerVault *EulerVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerVault.Contract.EulerVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerVault *EulerVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerVault *EulerVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerVault *EulerVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerVault.Contract.contract.Transact(opts, method, params...)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerVault *EulerVaultCaller) EVC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "EVC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerVault *EulerVaultSession) EVC() (common.Address, error) {
	return _EulerVault.Contract.EVC(&_EulerVault.CallOpts)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerVault *EulerVaultCallerSession) EVC() (common.Address, error) {
	return _EulerVault.Contract.EVC(&_EulerVault.CallOpts)
}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerVault *EulerVaultCaller) LTVBorrow(opts *bind.CallOpts, collateral common.Address) (uint16, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "LTVBorrow", collateral)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerVault *EulerVaultSession) LTVBorrow(collateral common.Address) (uint16, error) {
	return _EulerVault.Contract.LTVBorrow(&_EulerVault.CallOpts, collateral)
}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerVault *EulerVaultCallerSession) LTVBorrow(collateral common.Address) (uint16, error) {
	return _EulerVault.Contract.LTVBorrow(&_EulerVault.CallOpts, collateral)
}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerVault *EulerVaultCaller) LTVFull(opts *bind.CallOpts, collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "LTVFull", collateral)

	outstruct := new(struct {
		BorrowLTV             uint16
		LiquidationLTV        uint16
		InitialLiquidationLTV uint16
		TargetTimestamp       *big.Int
		RampDuration          uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BorrowLTV = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.LiquidationLTV = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.InitialLiquidationLTV = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.TargetTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RampDuration = *abi.ConvertType(out[4], new(uint32)).(*uint32)

	return *outstruct, err

}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerVault *EulerVaultSession) LTVFull(collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	return _EulerVault.Contract.LTVFull(&_EulerVault.CallOpts, collateral)
}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerVault *EulerVaultCallerSession) LTVFull(collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	return _EulerVault.Contract.LTVFull(&_EulerVault.CallOpts, collateral)
}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerVault *EulerVaultCaller) LTVLiquidation(opts *bind.CallOpts, collateral common.Address) (uint16, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "LTVLiquidation", collateral)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerVault *EulerVaultSession) LTVLiquidation(collateral common.Address) (uint16, error) {
	return _EulerVault.Contract.LTVLiquidation(&_EulerVault.CallOpts, collateral)
}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerVault *EulerVaultCallerSession) LTVLiquidation(collateral common.Address) (uint16, error) {
	return _EulerVault.Contract.LTVLiquidation(&_EulerVault.CallOpts, collateral)
}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerVault *EulerVaultCaller) LTVList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "LTVList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerVault *EulerVaultSession) LTVList() ([]common.Address, error) {
	return _EulerVault.Contract.LTVList(&_EulerVault.CallOpts)
}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerVault *EulerVaultCallerSession) LTVList() ([]common.Address, error) {
	return _EulerVault.Contract.LTVList(&_EulerVault.CallOpts)
}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerVault *EulerVaultCaller) MODULEBALANCEFORWARDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "MODULE_BALANCE_FORWARDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerVault *EulerVaultSession) MODULEBALANCEFORWARDER() (common.Address, error) {
	return _EulerVault.Contract.MODULEBALANCEFORWARDER(&_EulerVault.CallOpts)
}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerVault *EulerVaultCallerSession) MODULEBALANCEFORWARDER() (common.Address, error) {
	return _EulerVault.Contract.MODULEBALANCEFORWARDER(&_EulerVault.CallOpts)
}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerVault *EulerVaultCaller) MODULEBORROWING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "MODULE_BORROWING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerVault *EulerVaultSession) MODULEBORROWING() (common.Address, error) {
	return _EulerVault.Contract.MODULEBORROWING(&_EulerVault.CallOpts)
}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerVault *EulerVaultCallerSession) MODULEBORROWING() (common.Address, error) {
	return _EulerVault.Contract.MODULEBORROWING(&_EulerVault.CallOpts)
}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerVault *EulerVaultCaller) MODULEGOVERNANCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "MODULE_GOVERNANCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerVault *EulerVaultSession) MODULEGOVERNANCE() (common.Address, error) {
	return _EulerVault.Contract.MODULEGOVERNANCE(&_EulerVault.CallOpts)
}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerVault *EulerVaultCallerSession) MODULEGOVERNANCE() (common.Address, error) {
	return _EulerVault.Contract.MODULEGOVERNANCE(&_EulerVault.CallOpts)
}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerVault *EulerVaultCaller) MODULEINITIALIZE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "MODULE_INITIALIZE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerVault *EulerVaultSession) MODULEINITIALIZE() (common.Address, error) {
	return _EulerVault.Contract.MODULEINITIALIZE(&_EulerVault.CallOpts)
}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerVault *EulerVaultCallerSession) MODULEINITIALIZE() (common.Address, error) {
	return _EulerVault.Contract.MODULEINITIALIZE(&_EulerVault.CallOpts)
}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerVault *EulerVaultCaller) MODULELIQUIDATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "MODULE_LIQUIDATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerVault *EulerVaultSession) MODULELIQUIDATION() (common.Address, error) {
	return _EulerVault.Contract.MODULELIQUIDATION(&_EulerVault.CallOpts)
}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerVault *EulerVaultCallerSession) MODULELIQUIDATION() (common.Address, error) {
	return _EulerVault.Contract.MODULELIQUIDATION(&_EulerVault.CallOpts)
}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerVault *EulerVaultCaller) MODULERISKMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "MODULE_RISKMANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerVault *EulerVaultSession) MODULERISKMANAGER() (common.Address, error) {
	return _EulerVault.Contract.MODULERISKMANAGER(&_EulerVault.CallOpts)
}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerVault *EulerVaultCallerSession) MODULERISKMANAGER() (common.Address, error) {
	return _EulerVault.Contract.MODULERISKMANAGER(&_EulerVault.CallOpts)
}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerVault *EulerVaultCaller) MODULETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "MODULE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerVault *EulerVaultSession) MODULETOKEN() (common.Address, error) {
	return _EulerVault.Contract.MODULETOKEN(&_EulerVault.CallOpts)
}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerVault *EulerVaultCallerSession) MODULETOKEN() (common.Address, error) {
	return _EulerVault.Contract.MODULETOKEN(&_EulerVault.CallOpts)
}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerVault *EulerVaultCaller) MODULEVAULT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "MODULE_VAULT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerVault *EulerVaultSession) MODULEVAULT() (common.Address, error) {
	return _EulerVault.Contract.MODULEVAULT(&_EulerVault.CallOpts)
}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerVault *EulerVaultCallerSession) MODULEVAULT() (common.Address, error) {
	return _EulerVault.Contract.MODULEVAULT(&_EulerVault.CallOpts)
}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EulerVault *EulerVaultCaller) AccountLiquidity(opts *bind.CallOpts, account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "accountLiquidity", account, liquidation)

	outstruct := new(struct {
		CollateralValue *big.Int
		LiabilityValue  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollateralValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiabilityValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EulerVault *EulerVaultSession) AccountLiquidity(account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	return _EulerVault.Contract.AccountLiquidity(&_EulerVault.CallOpts, account, liquidation)
}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EulerVault *EulerVaultCallerSession) AccountLiquidity(account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	return _EulerVault.Contract.AccountLiquidity(&_EulerVault.CallOpts, account, liquidation)
}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EulerVault *EulerVaultCaller) AccountLiquidityFull(opts *bind.CallOpts, account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "accountLiquidityFull", account, liquidation)

	outstruct := new(struct {
		Collaterals      []common.Address
		CollateralValues []*big.Int
		LiabilityValue   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Collaterals = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CollateralValues = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.LiabilityValue = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EulerVault *EulerVaultSession) AccountLiquidityFull(account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	return _EulerVault.Contract.AccountLiquidityFull(&_EulerVault.CallOpts, account, liquidation)
}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EulerVault *EulerVaultCallerSession) AccountLiquidityFull(account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	return _EulerVault.Contract.AccountLiquidityFull(&_EulerVault.CallOpts, account, liquidation)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerVault *EulerVaultCaller) AccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "accumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerVault *EulerVaultSession) AccumulatedFees() (*big.Int, error) {
	return _EulerVault.Contract.AccumulatedFees(&_EulerVault.CallOpts)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) AccumulatedFees() (*big.Int, error) {
	return _EulerVault.Contract.AccumulatedFees(&_EulerVault.CallOpts)
}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerVault *EulerVaultCaller) AccumulatedFeesAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "accumulatedFeesAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerVault *EulerVaultSession) AccumulatedFeesAssets() (*big.Int, error) {
	return _EulerVault.Contract.AccumulatedFeesAssets(&_EulerVault.CallOpts)
}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) AccumulatedFeesAssets() (*big.Int, error) {
	return _EulerVault.Contract.AccumulatedFeesAssets(&_EulerVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerVault *EulerVaultCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerVault *EulerVaultSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _EulerVault.Contract.Allowance(&_EulerVault.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _EulerVault.Contract.Allowance(&_EulerVault.CallOpts, holder, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerVault *EulerVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerVault *EulerVaultSession) Asset() (common.Address, error) {
	return _EulerVault.Contract.Asset(&_EulerVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerVault *EulerVaultCallerSession) Asset() (common.Address, error) {
	return _EulerVault.Contract.Asset(&_EulerVault.CallOpts)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerVault *EulerVaultCaller) BalanceForwarderEnabled(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "balanceForwarderEnabled", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerVault *EulerVaultSession) BalanceForwarderEnabled(account common.Address) (bool, error) {
	return _EulerVault.Contract.BalanceForwarderEnabled(&_EulerVault.CallOpts, account)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerVault *EulerVaultCallerSession) BalanceForwarderEnabled(account common.Address) (bool, error) {
	return _EulerVault.Contract.BalanceForwarderEnabled(&_EulerVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerVault *EulerVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerVault *EulerVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EulerVault.Contract.BalanceOf(&_EulerVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EulerVault.Contract.BalanceOf(&_EulerVault.CallOpts, account)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerVault *EulerVaultCaller) BalanceTrackerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "balanceTrackerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerVault *EulerVaultSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerVault.Contract.BalanceTrackerAddress(&_EulerVault.CallOpts)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerVault *EulerVaultCallerSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerVault.Contract.BalanceTrackerAddress(&_EulerVault.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EulerVault *EulerVaultCaller) Caps(opts *bind.CallOpts) (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "caps")

	outstruct := new(struct {
		SupplyCap uint16
		BorrowCap uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SupplyCap = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.BorrowCap = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EulerVault *EulerVaultSession) Caps() (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	return _EulerVault.Contract.Caps(&_EulerVault.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EulerVault *EulerVaultCallerSession) Caps() (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	return _EulerVault.Contract.Caps(&_EulerVault.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerVault *EulerVaultCaller) Cash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerVault *EulerVaultSession) Cash() (*big.Int, error) {
	return _EulerVault.Contract.Cash(&_EulerVault.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) Cash() (*big.Int, error) {
	return _EulerVault.Contract.Cash(&_EulerVault.CallOpts)
}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerVault *EulerVaultCaller) CheckAccountStatus(opts *bind.CallOpts, account common.Address, collaterals []common.Address) ([4]byte, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "checkAccountStatus", account, collaterals)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerVault *EulerVaultSession) CheckAccountStatus(account common.Address, collaterals []common.Address) ([4]byte, error) {
	return _EulerVault.Contract.CheckAccountStatus(&_EulerVault.CallOpts, account, collaterals)
}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerVault *EulerVaultCallerSession) CheckAccountStatus(account common.Address, collaterals []common.Address) ([4]byte, error) {
	return _EulerVault.Contract.CheckAccountStatus(&_EulerVault.CallOpts, account, collaterals)
}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EulerVault *EulerVaultCaller) CheckLiquidation(opts *bind.CallOpts, liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "checkLiquidation", liquidator, violator, collateral)

	outstruct := new(struct {
		MaxRepay *big.Int
		MaxYield *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxRepay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxYield = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EulerVault *EulerVaultSession) CheckLiquidation(liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	return _EulerVault.Contract.CheckLiquidation(&_EulerVault.CallOpts, liquidator, violator, collateral)
}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EulerVault *EulerVaultCallerSession) CheckLiquidation(liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	return _EulerVault.Contract.CheckLiquidation(&_EulerVault.CallOpts, liquidator, violator, collateral)
}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerVault *EulerVaultCaller) ConfigFlags(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "configFlags")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerVault *EulerVaultSession) ConfigFlags() (uint32, error) {
	return _EulerVault.Contract.ConfigFlags(&_EulerVault.CallOpts)
}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerVault *EulerVaultCallerSession) ConfigFlags() (uint32, error) {
	return _EulerVault.Contract.ConfigFlags(&_EulerVault.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerVault *EulerVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerVault *EulerVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.ConvertToAssets(&_EulerVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.ConvertToAssets(&_EulerVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerVault *EulerVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerVault *EulerVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.ConvertToShares(&_EulerVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.ConvertToShares(&_EulerVault.CallOpts, assets)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerVault *EulerVaultCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerVault *EulerVaultSession) Creator() (common.Address, error) {
	return _EulerVault.Contract.Creator(&_EulerVault.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerVault *EulerVaultCallerSession) Creator() (common.Address, error) {
	return _EulerVault.Contract.Creator(&_EulerVault.CallOpts)
}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerVault *EulerVaultCaller) DToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "dToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerVault *EulerVaultSession) DToken() (common.Address, error) {
	return _EulerVault.Contract.DToken(&_EulerVault.CallOpts)
}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerVault *EulerVaultCallerSession) DToken() (common.Address, error) {
	return _EulerVault.Contract.DToken(&_EulerVault.CallOpts)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerVault *EulerVaultCaller) DebtOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "debtOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerVault *EulerVaultSession) DebtOf(account common.Address) (*big.Int, error) {
	return _EulerVault.Contract.DebtOf(&_EulerVault.CallOpts, account)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) DebtOf(account common.Address) (*big.Int, error) {
	return _EulerVault.Contract.DebtOf(&_EulerVault.CallOpts, account)
}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerVault *EulerVaultCaller) DebtOfExact(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "debtOfExact", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerVault *EulerVaultSession) DebtOfExact(account common.Address) (*big.Int, error) {
	return _EulerVault.Contract.DebtOfExact(&_EulerVault.CallOpts, account)
}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) DebtOfExact(account common.Address) (*big.Int, error) {
	return _EulerVault.Contract.DebtOfExact(&_EulerVault.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerVault *EulerVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerVault *EulerVaultSession) Decimals() (uint8, error) {
	return _EulerVault.Contract.Decimals(&_EulerVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerVault *EulerVaultCallerSession) Decimals() (uint8, error) {
	return _EulerVault.Contract.Decimals(&_EulerVault.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerVault *EulerVaultCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerVault *EulerVaultSession) FeeReceiver() (common.Address, error) {
	return _EulerVault.Contract.FeeReceiver(&_EulerVault.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerVault *EulerVaultCallerSession) FeeReceiver() (common.Address, error) {
	return _EulerVault.Contract.FeeReceiver(&_EulerVault.CallOpts)
}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerVault *EulerVaultCaller) GovernorAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "governorAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerVault *EulerVaultSession) GovernorAdmin() (common.Address, error) {
	return _EulerVault.Contract.GovernorAdmin(&_EulerVault.CallOpts)
}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerVault *EulerVaultCallerSession) GovernorAdmin() (common.Address, error) {
	return _EulerVault.Contract.GovernorAdmin(&_EulerVault.CallOpts)
}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EulerVault *EulerVaultCaller) HookConfig(opts *bind.CallOpts) (common.Address, uint32, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "hookConfig")

	if err != nil {
		return *new(common.Address), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EulerVault *EulerVaultSession) HookConfig() (common.Address, uint32, error) {
	return _EulerVault.Contract.HookConfig(&_EulerVault.CallOpts)
}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EulerVault *EulerVaultCallerSession) HookConfig() (common.Address, uint32, error) {
	return _EulerVault.Contract.HookConfig(&_EulerVault.CallOpts)
}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerVault *EulerVaultCaller) InterestAccumulator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "interestAccumulator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerVault *EulerVaultSession) InterestAccumulator() (*big.Int, error) {
	return _EulerVault.Contract.InterestAccumulator(&_EulerVault.CallOpts)
}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) InterestAccumulator() (*big.Int, error) {
	return _EulerVault.Contract.InterestAccumulator(&_EulerVault.CallOpts)
}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerVault *EulerVaultCaller) InterestFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "interestFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerVault *EulerVaultSession) InterestFee() (uint16, error) {
	return _EulerVault.Contract.InterestFee(&_EulerVault.CallOpts)
}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerVault *EulerVaultCallerSession) InterestFee() (uint16, error) {
	return _EulerVault.Contract.InterestFee(&_EulerVault.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerVault *EulerVaultCaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "interestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerVault *EulerVaultSession) InterestRate() (*big.Int, error) {
	return _EulerVault.Contract.InterestRate(&_EulerVault.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) InterestRate() (*big.Int, error) {
	return _EulerVault.Contract.InterestRate(&_EulerVault.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerVault *EulerVaultCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerVault *EulerVaultSession) InterestRateModel() (common.Address, error) {
	return _EulerVault.Contract.InterestRateModel(&_EulerVault.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerVault *EulerVaultCallerSession) InterestRateModel() (common.Address, error) {
	return _EulerVault.Contract.InterestRateModel(&_EulerVault.CallOpts)
}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerVault *EulerVaultCaller) LiquidationCoolOffTime(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "liquidationCoolOffTime")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerVault *EulerVaultSession) LiquidationCoolOffTime() (uint16, error) {
	return _EulerVault.Contract.LiquidationCoolOffTime(&_EulerVault.CallOpts)
}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerVault *EulerVaultCallerSession) LiquidationCoolOffTime() (uint16, error) {
	return _EulerVault.Contract.LiquidationCoolOffTime(&_EulerVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerVault *EulerVaultCaller) MaxDeposit(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "maxDeposit", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerVault *EulerVaultSession) MaxDeposit(account common.Address) (*big.Int, error) {
	return _EulerVault.Contract.MaxDeposit(&_EulerVault.CallOpts, account)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) MaxDeposit(account common.Address) (*big.Int, error) {
	return _EulerVault.Contract.MaxDeposit(&_EulerVault.CallOpts, account)
}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerVault *EulerVaultCaller) MaxLiquidationDiscount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "maxLiquidationDiscount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerVault *EulerVaultSession) MaxLiquidationDiscount() (uint16, error) {
	return _EulerVault.Contract.MaxLiquidationDiscount(&_EulerVault.CallOpts)
}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerVault *EulerVaultCallerSession) MaxLiquidationDiscount() (uint16, error) {
	return _EulerVault.Contract.MaxLiquidationDiscount(&_EulerVault.CallOpts)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerVault *EulerVaultCaller) MaxMint(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "maxMint", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerVault *EulerVaultSession) MaxMint(account common.Address) (*big.Int, error) {
	return _EulerVault.Contract.MaxMint(&_EulerVault.CallOpts, account)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) MaxMint(account common.Address) (*big.Int, error) {
	return _EulerVault.Contract.MaxMint(&_EulerVault.CallOpts, account)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerVault *EulerVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerVault *EulerVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _EulerVault.Contract.MaxRedeem(&_EulerVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _EulerVault.Contract.MaxRedeem(&_EulerVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerVault *EulerVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerVault *EulerVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _EulerVault.Contract.MaxWithdraw(&_EulerVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _EulerVault.Contract.MaxWithdraw(&_EulerVault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerVault *EulerVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerVault *EulerVaultSession) Name() (string, error) {
	return _EulerVault.Contract.Name(&_EulerVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerVault *EulerVaultCallerSession) Name() (string, error) {
	return _EulerVault.Contract.Name(&_EulerVault.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerVault *EulerVaultCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerVault *EulerVaultSession) Oracle() (common.Address, error) {
	return _EulerVault.Contract.Oracle(&_EulerVault.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerVault *EulerVaultCallerSession) Oracle() (common.Address, error) {
	return _EulerVault.Contract.Oracle(&_EulerVault.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerVault *EulerVaultCaller) Permit2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "permit2Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerVault *EulerVaultSession) Permit2Address() (common.Address, error) {
	return _EulerVault.Contract.Permit2Address(&_EulerVault.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerVault *EulerVaultCallerSession) Permit2Address() (common.Address, error) {
	return _EulerVault.Contract.Permit2Address(&_EulerVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerVault *EulerVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerVault *EulerVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.PreviewDeposit(&_EulerVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.PreviewDeposit(&_EulerVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerVault *EulerVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerVault *EulerVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.PreviewMint(&_EulerVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.PreviewMint(&_EulerVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerVault *EulerVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerVault *EulerVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.PreviewRedeem(&_EulerVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.PreviewRedeem(&_EulerVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerVault *EulerVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerVault *EulerVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.PreviewWithdraw(&_EulerVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _EulerVault.Contract.PreviewWithdraw(&_EulerVault.CallOpts, assets)
}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerVault *EulerVaultCaller) ProtocolConfigAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "protocolConfigAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerVault *EulerVaultSession) ProtocolConfigAddress() (common.Address, error) {
	return _EulerVault.Contract.ProtocolConfigAddress(&_EulerVault.CallOpts)
}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerVault *EulerVaultCallerSession) ProtocolConfigAddress() (common.Address, error) {
	return _EulerVault.Contract.ProtocolConfigAddress(&_EulerVault.CallOpts)
}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerVault *EulerVaultCaller) ProtocolFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "protocolFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerVault *EulerVaultSession) ProtocolFeeReceiver() (common.Address, error) {
	return _EulerVault.Contract.ProtocolFeeReceiver(&_EulerVault.CallOpts)
}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerVault *EulerVaultCallerSession) ProtocolFeeReceiver() (common.Address, error) {
	return _EulerVault.Contract.ProtocolFeeReceiver(&_EulerVault.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerVault *EulerVaultCaller) ProtocolFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "protocolFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerVault *EulerVaultSession) ProtocolFeeShare() (*big.Int, error) {
	return _EulerVault.Contract.ProtocolFeeShare(&_EulerVault.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) ProtocolFeeShare() (*big.Int, error) {
	return _EulerVault.Contract.ProtocolFeeShare(&_EulerVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerVault *EulerVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerVault *EulerVaultSession) Symbol() (string, error) {
	return _EulerVault.Contract.Symbol(&_EulerVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerVault *EulerVaultCallerSession) Symbol() (string, error) {
	return _EulerVault.Contract.Symbol(&_EulerVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerVault *EulerVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerVault *EulerVaultSession) TotalAssets() (*big.Int, error) {
	return _EulerVault.Contract.TotalAssets(&_EulerVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _EulerVault.Contract.TotalAssets(&_EulerVault.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerVault *EulerVaultCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerVault *EulerVaultSession) TotalBorrows() (*big.Int, error) {
	return _EulerVault.Contract.TotalBorrows(&_EulerVault.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) TotalBorrows() (*big.Int, error) {
	return _EulerVault.Contract.TotalBorrows(&_EulerVault.CallOpts)
}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerVault *EulerVaultCaller) TotalBorrowsExact(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "totalBorrowsExact")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerVault *EulerVaultSession) TotalBorrowsExact() (*big.Int, error) {
	return _EulerVault.Contract.TotalBorrowsExact(&_EulerVault.CallOpts)
}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) TotalBorrowsExact() (*big.Int, error) {
	return _EulerVault.Contract.TotalBorrowsExact(&_EulerVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerVault *EulerVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerVault *EulerVaultSession) TotalSupply() (*big.Int, error) {
	return _EulerVault.Contract.TotalSupply(&_EulerVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerVault *EulerVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _EulerVault.Contract.TotalSupply(&_EulerVault.CallOpts)
}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerVault *EulerVaultCaller) UnitOfAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVault.contract.Call(opts, &out, "unitOfAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerVault *EulerVaultSession) UnitOfAccount() (common.Address, error) {
	return _EulerVault.Contract.UnitOfAccount(&_EulerVault.CallOpts)
}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerVault *EulerVaultCallerSession) UnitOfAccount() (common.Address, error) {
	return _EulerVault.Contract.UnitOfAccount(&_EulerVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerVault *EulerVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerVault *EulerVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerVault.Contract.Approve(&_EulerVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerVault *EulerVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerVault.Contract.Approve(&_EulerVault.TransactOpts, spender, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "borrow", amount, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultSession) Borrow(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Borrow(&_EulerVault.TransactOpts, amount, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultTransactorSession) Borrow(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Borrow(&_EulerVault.TransactOpts, amount, receiver)
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerVault *EulerVaultTransactor) CheckVaultStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "checkVaultStatus")
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerVault *EulerVaultSession) CheckVaultStatus() (*types.Transaction, error) {
	return _EulerVault.Contract.CheckVaultStatus(&_EulerVault.TransactOpts)
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerVault *EulerVaultTransactorSession) CheckVaultStatus() (*types.Transaction, error) {
	return _EulerVault.Contract.CheckVaultStatus(&_EulerVault.TransactOpts)
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerVault *EulerVaultTransactor) ConvertFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "convertFees")
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerVault *EulerVaultSession) ConvertFees() (*types.Transaction, error) {
	return _EulerVault.Contract.ConvertFees(&_EulerVault.TransactOpts)
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerVault *EulerVaultTransactorSession) ConvertFees() (*types.Transaction, error) {
	return _EulerVault.Contract.ConvertFees(&_EulerVault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "deposit", amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Deposit(&_EulerVault.TransactOpts, amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultTransactorSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Deposit(&_EulerVault.TransactOpts, amount, receiver)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerVault *EulerVaultTransactor) DisableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "disableBalanceForwarder")
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerVault *EulerVaultSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerVault.Contract.DisableBalanceForwarder(&_EulerVault.TransactOpts)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerVault *EulerVaultTransactorSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerVault.Contract.DisableBalanceForwarder(&_EulerVault.TransactOpts)
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerVault *EulerVaultTransactor) DisableController(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "disableController")
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerVault *EulerVaultSession) DisableController() (*types.Transaction, error) {
	return _EulerVault.Contract.DisableController(&_EulerVault.TransactOpts)
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerVault *EulerVaultTransactorSession) DisableController() (*types.Transaction, error) {
	return _EulerVault.Contract.DisableController(&_EulerVault.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerVault *EulerVaultTransactor) EnableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "enableBalanceForwarder")
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerVault *EulerVaultSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerVault.Contract.EnableBalanceForwarder(&_EulerVault.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerVault *EulerVaultTransactorSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerVault.Contract.EnableBalanceForwarder(&_EulerVault.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerVault *EulerVaultTransactor) FlashLoan(opts *bind.TransactOpts, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "flashLoan", amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerVault *EulerVaultSession) FlashLoan(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerVault.Contract.FlashLoan(&_EulerVault.TransactOpts, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerVault *EulerVaultTransactorSession) FlashLoan(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerVault.Contract.FlashLoan(&_EulerVault.TransactOpts, amount, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerVault *EulerVaultTransactor) Initialize(opts *bind.TransactOpts, proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "initialize", proxyCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerVault *EulerVaultSession) Initialize(proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Initialize(&_EulerVault.TransactOpts, proxyCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerVault *EulerVaultTransactorSession) Initialize(proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Initialize(&_EulerVault.TransactOpts, proxyCreator)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerVault *EulerVaultTransactor) Liquidate(opts *bind.TransactOpts, violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "liquidate", violator, collateral, repayAssets, minYieldBalance)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerVault *EulerVaultSession) Liquidate(violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerVault.Contract.Liquidate(&_EulerVault.TransactOpts, violator, collateral, repayAssets, minYieldBalance)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerVault *EulerVaultTransactorSession) Liquidate(violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerVault.Contract.Liquidate(&_EulerVault.TransactOpts, violator, collateral, repayAssets, minYieldBalance)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultTransactor) Mint(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "mint", amount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultSession) Mint(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Mint(&_EulerVault.TransactOpts, amount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultTransactorSession) Mint(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Mint(&_EulerVault.TransactOpts, amount, receiver)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerVault *EulerVaultTransactor) PullDebt(opts *bind.TransactOpts, amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "pullDebt", amount, from)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerVault *EulerVaultSession) PullDebt(amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.PullDebt(&_EulerVault.TransactOpts, amount, from)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerVault *EulerVaultTransactorSession) PullDebt(amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.PullDebt(&_EulerVault.TransactOpts, amount, from)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerVault *EulerVaultTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "redeem", amount, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerVault *EulerVaultSession) Redeem(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Redeem(&_EulerVault.TransactOpts, amount, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerVault *EulerVaultTransactorSession) Redeem(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Redeem(&_EulerVault.TransactOpts, amount, receiver, owner)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultTransactor) Repay(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "repay", amount, receiver)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultSession) Repay(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Repay(&_EulerVault.TransactOpts, amount, receiver)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultTransactorSession) Repay(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Repay(&_EulerVault.TransactOpts, amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerVault *EulerVaultTransactor) RepayWithShares(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "repayWithShares", amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerVault *EulerVaultSession) RepayWithShares(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.RepayWithShares(&_EulerVault.TransactOpts, amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerVault *EulerVaultTransactorSession) RepayWithShares(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.RepayWithShares(&_EulerVault.TransactOpts, amount, receiver)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerVault *EulerVaultTransactor) SetCaps(opts *bind.TransactOpts, supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "setCaps", supplyCap, borrowCap)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerVault *EulerVaultSession) SetCaps(supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerVault.Contract.SetCaps(&_EulerVault.TransactOpts, supplyCap, borrowCap)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerVault *EulerVaultTransactorSession) SetCaps(supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerVault.Contract.SetCaps(&_EulerVault.TransactOpts, supplyCap, borrowCap)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerVault *EulerVaultTransactor) SetConfigFlags(opts *bind.TransactOpts, newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "setConfigFlags", newConfigFlags)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerVault *EulerVaultSession) SetConfigFlags(newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerVault.Contract.SetConfigFlags(&_EulerVault.TransactOpts, newConfigFlags)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerVault *EulerVaultTransactorSession) SetConfigFlags(newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerVault.Contract.SetConfigFlags(&_EulerVault.TransactOpts, newConfigFlags)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerVault *EulerVaultTransactor) SetFeeReceiver(opts *bind.TransactOpts, newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "setFeeReceiver", newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerVault *EulerVaultSession) SetFeeReceiver(newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.SetFeeReceiver(&_EulerVault.TransactOpts, newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerVault *EulerVaultTransactorSession) SetFeeReceiver(newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.SetFeeReceiver(&_EulerVault.TransactOpts, newFeeReceiver)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerVault *EulerVaultTransactor) SetGovernorAdmin(opts *bind.TransactOpts, newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "setGovernorAdmin", newGovernorAdmin)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerVault *EulerVaultSession) SetGovernorAdmin(newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.SetGovernorAdmin(&_EulerVault.TransactOpts, newGovernorAdmin)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerVault *EulerVaultTransactorSession) SetGovernorAdmin(newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.SetGovernorAdmin(&_EulerVault.TransactOpts, newGovernorAdmin)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerVault *EulerVaultTransactor) SetHookConfig(opts *bind.TransactOpts, newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "setHookConfig", newHookTarget, newHookedOps)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerVault *EulerVaultSession) SetHookConfig(newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerVault.Contract.SetHookConfig(&_EulerVault.TransactOpts, newHookTarget, newHookedOps)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerVault *EulerVaultTransactorSession) SetHookConfig(newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerVault.Contract.SetHookConfig(&_EulerVault.TransactOpts, newHookTarget, newHookedOps)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerVault *EulerVaultTransactor) SetInterestFee(opts *bind.TransactOpts, newFee uint16) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "setInterestFee", newFee)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerVault *EulerVaultSession) SetInterestFee(newFee uint16) (*types.Transaction, error) {
	return _EulerVault.Contract.SetInterestFee(&_EulerVault.TransactOpts, newFee)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerVault *EulerVaultTransactorSession) SetInterestFee(newFee uint16) (*types.Transaction, error) {
	return _EulerVault.Contract.SetInterestFee(&_EulerVault.TransactOpts, newFee)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerVault *EulerVaultTransactor) SetInterestRateModel(opts *bind.TransactOpts, newModel common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "setInterestRateModel", newModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerVault *EulerVaultSession) SetInterestRateModel(newModel common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.SetInterestRateModel(&_EulerVault.TransactOpts, newModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerVault *EulerVaultTransactorSession) SetInterestRateModel(newModel common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.SetInterestRateModel(&_EulerVault.TransactOpts, newModel)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerVault *EulerVaultTransactor) SetLTV(opts *bind.TransactOpts, collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "setLTV", collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerVault *EulerVaultSession) SetLTV(collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerVault.Contract.SetLTV(&_EulerVault.TransactOpts, collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerVault *EulerVaultTransactorSession) SetLTV(collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerVault.Contract.SetLTV(&_EulerVault.TransactOpts, collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerVault *EulerVaultTransactor) SetLiquidationCoolOffTime(opts *bind.TransactOpts, newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "setLiquidationCoolOffTime", newCoolOffTime)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerVault *EulerVaultSession) SetLiquidationCoolOffTime(newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerVault.Contract.SetLiquidationCoolOffTime(&_EulerVault.TransactOpts, newCoolOffTime)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerVault *EulerVaultTransactorSession) SetLiquidationCoolOffTime(newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerVault.Contract.SetLiquidationCoolOffTime(&_EulerVault.TransactOpts, newCoolOffTime)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerVault *EulerVaultTransactor) SetMaxLiquidationDiscount(opts *bind.TransactOpts, newDiscount uint16) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "setMaxLiquidationDiscount", newDiscount)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerVault *EulerVaultSession) SetMaxLiquidationDiscount(newDiscount uint16) (*types.Transaction, error) {
	return _EulerVault.Contract.SetMaxLiquidationDiscount(&_EulerVault.TransactOpts, newDiscount)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerVault *EulerVaultTransactorSession) SetMaxLiquidationDiscount(newDiscount uint16) (*types.Transaction, error) {
	return _EulerVault.Contract.SetMaxLiquidationDiscount(&_EulerVault.TransactOpts, newDiscount)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultTransactor) Skim(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "skim", amount, receiver)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultSession) Skim(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Skim(&_EulerVault.TransactOpts, amount, receiver)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerVault *EulerVaultTransactorSession) Skim(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Skim(&_EulerVault.TransactOpts, amount, receiver)
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerVault *EulerVaultTransactor) Touch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "touch")
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerVault *EulerVaultSession) Touch() (*types.Transaction, error) {
	return _EulerVault.Contract.Touch(&_EulerVault.TransactOpts)
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerVault *EulerVaultTransactorSession) Touch() (*types.Transaction, error) {
	return _EulerVault.Contract.Touch(&_EulerVault.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerVault *EulerVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerVault *EulerVaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerVault.Contract.Transfer(&_EulerVault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerVault *EulerVaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerVault.Contract.Transfer(&_EulerVault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerVault *EulerVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerVault *EulerVaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerVault.Contract.TransferFrom(&_EulerVault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerVault *EulerVaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerVault.Contract.TransferFrom(&_EulerVault.TransactOpts, from, to, amount)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerVault *EulerVaultTransactor) TransferFromMax(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "transferFromMax", from, to)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerVault *EulerVaultSession) TransferFromMax(from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.TransferFromMax(&_EulerVault.TransactOpts, from, to)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerVault *EulerVaultTransactorSession) TransferFromMax(from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.TransferFromMax(&_EulerVault.TransactOpts, from, to)
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerVault *EulerVaultTransactor) ViewDelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "viewDelegate")
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerVault *EulerVaultSession) ViewDelegate() (*types.Transaction, error) {
	return _EulerVault.Contract.ViewDelegate(&_EulerVault.TransactOpts)
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerVault *EulerVaultTransactorSession) ViewDelegate() (*types.Transaction, error) {
	return _EulerVault.Contract.ViewDelegate(&_EulerVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerVault *EulerVaultTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerVault.contract.Transact(opts, "withdraw", amount, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerVault *EulerVaultSession) Withdraw(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Withdraw(&_EulerVault.TransactOpts, amount, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerVault *EulerVaultTransactorSession) Withdraw(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerVault.Contract.Withdraw(&_EulerVault.TransactOpts, amount, receiver, owner)
}

// EulerVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EulerVault contract.
type EulerVaultApprovalIterator struct {
	Event *EulerVaultApproval // Event containing the contract specifics and raw log

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
func (it *EulerVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultApproval)
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
		it.Event = new(EulerVaultApproval)
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
func (it *EulerVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultApproval represents a Approval event raised by the EulerVault contract.
type EulerVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerVault *EulerVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EulerVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultApprovalIterator{contract: _EulerVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerVault *EulerVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EulerVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultApproval)
				if err := _EulerVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_EulerVault *EulerVaultFilterer) ParseApproval(log types.Log) (*EulerVaultApproval, error) {
	event := new(EulerVaultApproval)
	if err := _EulerVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultBalanceForwarderStatusIterator is returned from FilterBalanceForwarderStatus and is used to iterate over the raw logs and unpacked data for BalanceForwarderStatus events raised by the EulerVault contract.
type EulerVaultBalanceForwarderStatusIterator struct {
	Event *EulerVaultBalanceForwarderStatus // Event containing the contract specifics and raw log

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
func (it *EulerVaultBalanceForwarderStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultBalanceForwarderStatus)
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
		it.Event = new(EulerVaultBalanceForwarderStatus)
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
func (it *EulerVaultBalanceForwarderStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultBalanceForwarderStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultBalanceForwarderStatus represents a BalanceForwarderStatus event raised by the EulerVault contract.
type EulerVaultBalanceForwarderStatus struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalanceForwarderStatus is a free log retrieval operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EulerVault *EulerVaultFilterer) FilterBalanceForwarderStatus(opts *bind.FilterOpts, account []common.Address) (*EulerVaultBalanceForwarderStatusIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "BalanceForwarderStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultBalanceForwarderStatusIterator{contract: _EulerVault.contract, event: "BalanceForwarderStatus", logs: logs, sub: sub}, nil
}

// WatchBalanceForwarderStatus is a free log subscription operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EulerVault *EulerVaultFilterer) WatchBalanceForwarderStatus(opts *bind.WatchOpts, sink chan<- *EulerVaultBalanceForwarderStatus, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "BalanceForwarderStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultBalanceForwarderStatus)
				if err := _EulerVault.contract.UnpackLog(event, "BalanceForwarderStatus", log); err != nil {
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

// ParseBalanceForwarderStatus is a log parse operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EulerVault *EulerVaultFilterer) ParseBalanceForwarderStatus(log types.Log) (*EulerVaultBalanceForwarderStatus, error) {
	event := new(EulerVaultBalanceForwarderStatus)
	if err := _EulerVault.contract.UnpackLog(event, "BalanceForwarderStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the EulerVault contract.
type EulerVaultBorrowIterator struct {
	Event *EulerVaultBorrow // Event containing the contract specifics and raw log

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
func (it *EulerVaultBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultBorrow)
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
		it.Event = new(EulerVaultBorrow)
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
func (it *EulerVaultBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultBorrow represents a Borrow event raised by the EulerVault contract.
type EulerVaultBorrow struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) FilterBorrow(opts *bind.FilterOpts, account []common.Address) (*EulerVaultBorrowIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultBorrowIterator{contract: _EulerVault.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *EulerVaultBorrow, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultBorrow)
				if err := _EulerVault.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) ParseBorrow(log types.Log) (*EulerVaultBorrow, error) {
	event := new(EulerVaultBorrow)
	if err := _EulerVault.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultConvertFeesIterator is returned from FilterConvertFees and is used to iterate over the raw logs and unpacked data for ConvertFees events raised by the EulerVault contract.
type EulerVaultConvertFeesIterator struct {
	Event *EulerVaultConvertFees // Event containing the contract specifics and raw log

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
func (it *EulerVaultConvertFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultConvertFees)
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
		it.Event = new(EulerVaultConvertFees)
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
func (it *EulerVaultConvertFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultConvertFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultConvertFees represents a ConvertFees event raised by the EulerVault contract.
type EulerVaultConvertFees struct {
	Sender           common.Address
	ProtocolReceiver common.Address
	GovernorReceiver common.Address
	ProtocolShares   *big.Int
	GovernorShares   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConvertFees is a free log retrieval operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EulerVault *EulerVaultFilterer) FilterConvertFees(opts *bind.FilterOpts, sender []common.Address, protocolReceiver []common.Address, governorReceiver []common.Address) (*EulerVaultConvertFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var protocolReceiverRule []interface{}
	for _, protocolReceiverItem := range protocolReceiver {
		protocolReceiverRule = append(protocolReceiverRule, protocolReceiverItem)
	}
	var governorReceiverRule []interface{}
	for _, governorReceiverItem := range governorReceiver {
		governorReceiverRule = append(governorReceiverRule, governorReceiverItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "ConvertFees", senderRule, protocolReceiverRule, governorReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultConvertFeesIterator{contract: _EulerVault.contract, event: "ConvertFees", logs: logs, sub: sub}, nil
}

// WatchConvertFees is a free log subscription operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EulerVault *EulerVaultFilterer) WatchConvertFees(opts *bind.WatchOpts, sink chan<- *EulerVaultConvertFees, sender []common.Address, protocolReceiver []common.Address, governorReceiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var protocolReceiverRule []interface{}
	for _, protocolReceiverItem := range protocolReceiver {
		protocolReceiverRule = append(protocolReceiverRule, protocolReceiverItem)
	}
	var governorReceiverRule []interface{}
	for _, governorReceiverItem := range governorReceiver {
		governorReceiverRule = append(governorReceiverRule, governorReceiverItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "ConvertFees", senderRule, protocolReceiverRule, governorReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultConvertFees)
				if err := _EulerVault.contract.UnpackLog(event, "ConvertFees", log); err != nil {
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

// ParseConvertFees is a log parse operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EulerVault *EulerVaultFilterer) ParseConvertFees(log types.Log) (*EulerVaultConvertFees, error) {
	event := new(EulerVaultConvertFees)
	if err := _EulerVault.contract.UnpackLog(event, "ConvertFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultDebtSocializedIterator is returned from FilterDebtSocialized and is used to iterate over the raw logs and unpacked data for DebtSocialized events raised by the EulerVault contract.
type EulerVaultDebtSocializedIterator struct {
	Event *EulerVaultDebtSocialized // Event containing the contract specifics and raw log

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
func (it *EulerVaultDebtSocializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultDebtSocialized)
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
		it.Event = new(EulerVaultDebtSocialized)
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
func (it *EulerVaultDebtSocializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultDebtSocializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultDebtSocialized represents a DebtSocialized event raised by the EulerVault contract.
type EulerVaultDebtSocialized struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebtSocialized is a free log retrieval operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) FilterDebtSocialized(opts *bind.FilterOpts, account []common.Address) (*EulerVaultDebtSocializedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "DebtSocialized", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultDebtSocializedIterator{contract: _EulerVault.contract, event: "DebtSocialized", logs: logs, sub: sub}, nil
}

// WatchDebtSocialized is a free log subscription operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) WatchDebtSocialized(opts *bind.WatchOpts, sink chan<- *EulerVaultDebtSocialized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "DebtSocialized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultDebtSocialized)
				if err := _EulerVault.contract.UnpackLog(event, "DebtSocialized", log); err != nil {
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

// ParseDebtSocialized is a log parse operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) ParseDebtSocialized(log types.Log) (*EulerVaultDebtSocialized, error) {
	event := new(EulerVaultDebtSocialized)
	if err := _EulerVault.contract.UnpackLog(event, "DebtSocialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EulerVault contract.
type EulerVaultDepositIterator struct {
	Event *EulerVaultDeposit // Event containing the contract specifics and raw log

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
func (it *EulerVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultDeposit)
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
		it.Event = new(EulerVaultDeposit)
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
func (it *EulerVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultDeposit represents a Deposit event raised by the EulerVault contract.
type EulerVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerVault *EulerVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*EulerVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultDepositIterator{contract: _EulerVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerVault *EulerVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EulerVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultDeposit)
				if err := _EulerVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_EulerVault *EulerVaultFilterer) ParseDeposit(log types.Log) (*EulerVaultDeposit, error) {
	event := new(EulerVaultDeposit)
	if err := _EulerVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultEVaultCreatedIterator is returned from FilterEVaultCreated and is used to iterate over the raw logs and unpacked data for EVaultCreated events raised by the EulerVault contract.
type EulerVaultEVaultCreatedIterator struct {
	Event *EulerVaultEVaultCreated // Event containing the contract specifics and raw log

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
func (it *EulerVaultEVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultEVaultCreated)
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
		it.Event = new(EulerVaultEVaultCreated)
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
func (it *EulerVaultEVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultEVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultEVaultCreated represents a EVaultCreated event raised by the EulerVault contract.
type EulerVaultEVaultCreated struct {
	Creator common.Address
	Asset   common.Address
	DToken  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEVaultCreated is a free log retrieval operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EulerVault *EulerVaultFilterer) FilterEVaultCreated(opts *bind.FilterOpts, creator []common.Address, asset []common.Address) (*EulerVaultEVaultCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "EVaultCreated", creatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultEVaultCreatedIterator{contract: _EulerVault.contract, event: "EVaultCreated", logs: logs, sub: sub}, nil
}

// WatchEVaultCreated is a free log subscription operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EulerVault *EulerVaultFilterer) WatchEVaultCreated(opts *bind.WatchOpts, sink chan<- *EulerVaultEVaultCreated, creator []common.Address, asset []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "EVaultCreated", creatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultEVaultCreated)
				if err := _EulerVault.contract.UnpackLog(event, "EVaultCreated", log); err != nil {
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

// ParseEVaultCreated is a log parse operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EulerVault *EulerVaultFilterer) ParseEVaultCreated(log types.Log) (*EulerVaultEVaultCreated, error) {
	event := new(EulerVaultEVaultCreated)
	if err := _EulerVault.contract.UnpackLog(event, "EVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultGovSetCapsIterator is returned from FilterGovSetCaps and is used to iterate over the raw logs and unpacked data for GovSetCaps events raised by the EulerVault contract.
type EulerVaultGovSetCapsIterator struct {
	Event *EulerVaultGovSetCaps // Event containing the contract specifics and raw log

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
func (it *EulerVaultGovSetCapsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultGovSetCaps)
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
		it.Event = new(EulerVaultGovSetCaps)
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
func (it *EulerVaultGovSetCapsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultGovSetCapsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultGovSetCaps represents a GovSetCaps event raised by the EulerVault contract.
type EulerVaultGovSetCaps struct {
	NewSupplyCap uint16
	NewBorrowCap uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGovSetCaps is a free log retrieval operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EulerVault *EulerVaultFilterer) FilterGovSetCaps(opts *bind.FilterOpts) (*EulerVaultGovSetCapsIterator, error) {

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "GovSetCaps")
	if err != nil {
		return nil, err
	}
	return &EulerVaultGovSetCapsIterator{contract: _EulerVault.contract, event: "GovSetCaps", logs: logs, sub: sub}, nil
}

// WatchGovSetCaps is a free log subscription operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EulerVault *EulerVaultFilterer) WatchGovSetCaps(opts *bind.WatchOpts, sink chan<- *EulerVaultGovSetCaps) (event.Subscription, error) {

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "GovSetCaps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultGovSetCaps)
				if err := _EulerVault.contract.UnpackLog(event, "GovSetCaps", log); err != nil {
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

// ParseGovSetCaps is a log parse operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EulerVault *EulerVaultFilterer) ParseGovSetCaps(log types.Log) (*EulerVaultGovSetCaps, error) {
	event := new(EulerVaultGovSetCaps)
	if err := _EulerVault.contract.UnpackLog(event, "GovSetCaps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultGovSetConfigFlagsIterator is returned from FilterGovSetConfigFlags and is used to iterate over the raw logs and unpacked data for GovSetConfigFlags events raised by the EulerVault contract.
type EulerVaultGovSetConfigFlagsIterator struct {
	Event *EulerVaultGovSetConfigFlags // Event containing the contract specifics and raw log

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
func (it *EulerVaultGovSetConfigFlagsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultGovSetConfigFlags)
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
		it.Event = new(EulerVaultGovSetConfigFlags)
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
func (it *EulerVaultGovSetConfigFlagsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultGovSetConfigFlagsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultGovSetConfigFlags represents a GovSetConfigFlags event raised by the EulerVault contract.
type EulerVaultGovSetConfigFlags struct {
	NewConfigFlags uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetConfigFlags is a free log retrieval operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EulerVault *EulerVaultFilterer) FilterGovSetConfigFlags(opts *bind.FilterOpts) (*EulerVaultGovSetConfigFlagsIterator, error) {

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "GovSetConfigFlags")
	if err != nil {
		return nil, err
	}
	return &EulerVaultGovSetConfigFlagsIterator{contract: _EulerVault.contract, event: "GovSetConfigFlags", logs: logs, sub: sub}, nil
}

// WatchGovSetConfigFlags is a free log subscription operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EulerVault *EulerVaultFilterer) WatchGovSetConfigFlags(opts *bind.WatchOpts, sink chan<- *EulerVaultGovSetConfigFlags) (event.Subscription, error) {

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "GovSetConfigFlags")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultGovSetConfigFlags)
				if err := _EulerVault.contract.UnpackLog(event, "GovSetConfigFlags", log); err != nil {
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

// ParseGovSetConfigFlags is a log parse operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EulerVault *EulerVaultFilterer) ParseGovSetConfigFlags(log types.Log) (*EulerVaultGovSetConfigFlags, error) {
	event := new(EulerVaultGovSetConfigFlags)
	if err := _EulerVault.contract.UnpackLog(event, "GovSetConfigFlags", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultGovSetFeeReceiverIterator is returned from FilterGovSetFeeReceiver and is used to iterate over the raw logs and unpacked data for GovSetFeeReceiver events raised by the EulerVault contract.
type EulerVaultGovSetFeeReceiverIterator struct {
	Event *EulerVaultGovSetFeeReceiver // Event containing the contract specifics and raw log

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
func (it *EulerVaultGovSetFeeReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultGovSetFeeReceiver)
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
		it.Event = new(EulerVaultGovSetFeeReceiver)
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
func (it *EulerVaultGovSetFeeReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultGovSetFeeReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultGovSetFeeReceiver represents a GovSetFeeReceiver event raised by the EulerVault contract.
type EulerVaultGovSetFeeReceiver struct {
	NewFeeReceiver common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetFeeReceiver is a free log retrieval operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EulerVault *EulerVaultFilterer) FilterGovSetFeeReceiver(opts *bind.FilterOpts, newFeeReceiver []common.Address) (*EulerVaultGovSetFeeReceiverIterator, error) {

	var newFeeReceiverRule []interface{}
	for _, newFeeReceiverItem := range newFeeReceiver {
		newFeeReceiverRule = append(newFeeReceiverRule, newFeeReceiverItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "GovSetFeeReceiver", newFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultGovSetFeeReceiverIterator{contract: _EulerVault.contract, event: "GovSetFeeReceiver", logs: logs, sub: sub}, nil
}

// WatchGovSetFeeReceiver is a free log subscription operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EulerVault *EulerVaultFilterer) WatchGovSetFeeReceiver(opts *bind.WatchOpts, sink chan<- *EulerVaultGovSetFeeReceiver, newFeeReceiver []common.Address) (event.Subscription, error) {

	var newFeeReceiverRule []interface{}
	for _, newFeeReceiverItem := range newFeeReceiver {
		newFeeReceiverRule = append(newFeeReceiverRule, newFeeReceiverItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "GovSetFeeReceiver", newFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultGovSetFeeReceiver)
				if err := _EulerVault.contract.UnpackLog(event, "GovSetFeeReceiver", log); err != nil {
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

// ParseGovSetFeeReceiver is a log parse operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EulerVault *EulerVaultFilterer) ParseGovSetFeeReceiver(log types.Log) (*EulerVaultGovSetFeeReceiver, error) {
	event := new(EulerVaultGovSetFeeReceiver)
	if err := _EulerVault.contract.UnpackLog(event, "GovSetFeeReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultGovSetGovernorAdminIterator is returned from FilterGovSetGovernorAdmin and is used to iterate over the raw logs and unpacked data for GovSetGovernorAdmin events raised by the EulerVault contract.
type EulerVaultGovSetGovernorAdminIterator struct {
	Event *EulerVaultGovSetGovernorAdmin // Event containing the contract specifics and raw log

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
func (it *EulerVaultGovSetGovernorAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultGovSetGovernorAdmin)
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
		it.Event = new(EulerVaultGovSetGovernorAdmin)
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
func (it *EulerVaultGovSetGovernorAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultGovSetGovernorAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultGovSetGovernorAdmin represents a GovSetGovernorAdmin event raised by the EulerVault contract.
type EulerVaultGovSetGovernorAdmin struct {
	NewGovernorAdmin common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGovSetGovernorAdmin is a free log retrieval operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EulerVault *EulerVaultFilterer) FilterGovSetGovernorAdmin(opts *bind.FilterOpts, newGovernorAdmin []common.Address) (*EulerVaultGovSetGovernorAdminIterator, error) {

	var newGovernorAdminRule []interface{}
	for _, newGovernorAdminItem := range newGovernorAdmin {
		newGovernorAdminRule = append(newGovernorAdminRule, newGovernorAdminItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "GovSetGovernorAdmin", newGovernorAdminRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultGovSetGovernorAdminIterator{contract: _EulerVault.contract, event: "GovSetGovernorAdmin", logs: logs, sub: sub}, nil
}

// WatchGovSetGovernorAdmin is a free log subscription operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EulerVault *EulerVaultFilterer) WatchGovSetGovernorAdmin(opts *bind.WatchOpts, sink chan<- *EulerVaultGovSetGovernorAdmin, newGovernorAdmin []common.Address) (event.Subscription, error) {

	var newGovernorAdminRule []interface{}
	for _, newGovernorAdminItem := range newGovernorAdmin {
		newGovernorAdminRule = append(newGovernorAdminRule, newGovernorAdminItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "GovSetGovernorAdmin", newGovernorAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultGovSetGovernorAdmin)
				if err := _EulerVault.contract.UnpackLog(event, "GovSetGovernorAdmin", log); err != nil {
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

// ParseGovSetGovernorAdmin is a log parse operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EulerVault *EulerVaultFilterer) ParseGovSetGovernorAdmin(log types.Log) (*EulerVaultGovSetGovernorAdmin, error) {
	event := new(EulerVaultGovSetGovernorAdmin)
	if err := _EulerVault.contract.UnpackLog(event, "GovSetGovernorAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultGovSetHookConfigIterator is returned from FilterGovSetHookConfig and is used to iterate over the raw logs and unpacked data for GovSetHookConfig events raised by the EulerVault contract.
type EulerVaultGovSetHookConfigIterator struct {
	Event *EulerVaultGovSetHookConfig // Event containing the contract specifics and raw log

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
func (it *EulerVaultGovSetHookConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultGovSetHookConfig)
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
		it.Event = new(EulerVaultGovSetHookConfig)
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
func (it *EulerVaultGovSetHookConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultGovSetHookConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultGovSetHookConfig represents a GovSetHookConfig event raised by the EulerVault contract.
type EulerVaultGovSetHookConfig struct {
	NewHookTarget common.Address
	NewHookedOps  uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGovSetHookConfig is a free log retrieval operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EulerVault *EulerVaultFilterer) FilterGovSetHookConfig(opts *bind.FilterOpts, newHookTarget []common.Address) (*EulerVaultGovSetHookConfigIterator, error) {

	var newHookTargetRule []interface{}
	for _, newHookTargetItem := range newHookTarget {
		newHookTargetRule = append(newHookTargetRule, newHookTargetItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "GovSetHookConfig", newHookTargetRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultGovSetHookConfigIterator{contract: _EulerVault.contract, event: "GovSetHookConfig", logs: logs, sub: sub}, nil
}

// WatchGovSetHookConfig is a free log subscription operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EulerVault *EulerVaultFilterer) WatchGovSetHookConfig(opts *bind.WatchOpts, sink chan<- *EulerVaultGovSetHookConfig, newHookTarget []common.Address) (event.Subscription, error) {

	var newHookTargetRule []interface{}
	for _, newHookTargetItem := range newHookTarget {
		newHookTargetRule = append(newHookTargetRule, newHookTargetItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "GovSetHookConfig", newHookTargetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultGovSetHookConfig)
				if err := _EulerVault.contract.UnpackLog(event, "GovSetHookConfig", log); err != nil {
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

// ParseGovSetHookConfig is a log parse operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EulerVault *EulerVaultFilterer) ParseGovSetHookConfig(log types.Log) (*EulerVaultGovSetHookConfig, error) {
	event := new(EulerVaultGovSetHookConfig)
	if err := _EulerVault.contract.UnpackLog(event, "GovSetHookConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultGovSetInterestFeeIterator is returned from FilterGovSetInterestFee and is used to iterate over the raw logs and unpacked data for GovSetInterestFee events raised by the EulerVault contract.
type EulerVaultGovSetInterestFeeIterator struct {
	Event *EulerVaultGovSetInterestFee // Event containing the contract specifics and raw log

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
func (it *EulerVaultGovSetInterestFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultGovSetInterestFee)
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
		it.Event = new(EulerVaultGovSetInterestFee)
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
func (it *EulerVaultGovSetInterestFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultGovSetInterestFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultGovSetInterestFee represents a GovSetInterestFee event raised by the EulerVault contract.
type EulerVaultGovSetInterestFee struct {
	NewFee uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGovSetInterestFee is a free log retrieval operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EulerVault *EulerVaultFilterer) FilterGovSetInterestFee(opts *bind.FilterOpts) (*EulerVaultGovSetInterestFeeIterator, error) {

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "GovSetInterestFee")
	if err != nil {
		return nil, err
	}
	return &EulerVaultGovSetInterestFeeIterator{contract: _EulerVault.contract, event: "GovSetInterestFee", logs: logs, sub: sub}, nil
}

// WatchGovSetInterestFee is a free log subscription operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EulerVault *EulerVaultFilterer) WatchGovSetInterestFee(opts *bind.WatchOpts, sink chan<- *EulerVaultGovSetInterestFee) (event.Subscription, error) {

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "GovSetInterestFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultGovSetInterestFee)
				if err := _EulerVault.contract.UnpackLog(event, "GovSetInterestFee", log); err != nil {
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

// ParseGovSetInterestFee is a log parse operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EulerVault *EulerVaultFilterer) ParseGovSetInterestFee(log types.Log) (*EulerVaultGovSetInterestFee, error) {
	event := new(EulerVaultGovSetInterestFee)
	if err := _EulerVault.contract.UnpackLog(event, "GovSetInterestFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultGovSetInterestRateModelIterator is returned from FilterGovSetInterestRateModel and is used to iterate over the raw logs and unpacked data for GovSetInterestRateModel events raised by the EulerVault contract.
type EulerVaultGovSetInterestRateModelIterator struct {
	Event *EulerVaultGovSetInterestRateModel // Event containing the contract specifics and raw log

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
func (it *EulerVaultGovSetInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultGovSetInterestRateModel)
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
		it.Event = new(EulerVaultGovSetInterestRateModel)
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
func (it *EulerVaultGovSetInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultGovSetInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultGovSetInterestRateModel represents a GovSetInterestRateModel event raised by the EulerVault contract.
type EulerVaultGovSetInterestRateModel struct {
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGovSetInterestRateModel is a free log retrieval operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EulerVault *EulerVaultFilterer) FilterGovSetInterestRateModel(opts *bind.FilterOpts) (*EulerVaultGovSetInterestRateModelIterator, error) {

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "GovSetInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &EulerVaultGovSetInterestRateModelIterator{contract: _EulerVault.contract, event: "GovSetInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchGovSetInterestRateModel is a free log subscription operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EulerVault *EulerVaultFilterer) WatchGovSetInterestRateModel(opts *bind.WatchOpts, sink chan<- *EulerVaultGovSetInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "GovSetInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultGovSetInterestRateModel)
				if err := _EulerVault.contract.UnpackLog(event, "GovSetInterestRateModel", log); err != nil {
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

// ParseGovSetInterestRateModel is a log parse operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EulerVault *EulerVaultFilterer) ParseGovSetInterestRateModel(log types.Log) (*EulerVaultGovSetInterestRateModel, error) {
	event := new(EulerVaultGovSetInterestRateModel)
	if err := _EulerVault.contract.UnpackLog(event, "GovSetInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultGovSetLTVIterator is returned from FilterGovSetLTV and is used to iterate over the raw logs and unpacked data for GovSetLTV events raised by the EulerVault contract.
type EulerVaultGovSetLTVIterator struct {
	Event *EulerVaultGovSetLTV // Event containing the contract specifics and raw log

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
func (it *EulerVaultGovSetLTVIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultGovSetLTV)
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
		it.Event = new(EulerVaultGovSetLTV)
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
func (it *EulerVaultGovSetLTVIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultGovSetLTVIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultGovSetLTV represents a GovSetLTV event raised by the EulerVault contract.
type EulerVaultGovSetLTV struct {
	Collateral            common.Address
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovSetLTV is a free log retrieval operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerVault *EulerVaultFilterer) FilterGovSetLTV(opts *bind.FilterOpts, collateral []common.Address) (*EulerVaultGovSetLTVIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "GovSetLTV", collateralRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultGovSetLTVIterator{contract: _EulerVault.contract, event: "GovSetLTV", logs: logs, sub: sub}, nil
}

// WatchGovSetLTV is a free log subscription operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerVault *EulerVaultFilterer) WatchGovSetLTV(opts *bind.WatchOpts, sink chan<- *EulerVaultGovSetLTV, collateral []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "GovSetLTV", collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultGovSetLTV)
				if err := _EulerVault.contract.UnpackLog(event, "GovSetLTV", log); err != nil {
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

// ParseGovSetLTV is a log parse operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerVault *EulerVaultFilterer) ParseGovSetLTV(log types.Log) (*EulerVaultGovSetLTV, error) {
	event := new(EulerVaultGovSetLTV)
	if err := _EulerVault.contract.UnpackLog(event, "GovSetLTV", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultGovSetLiquidationCoolOffTimeIterator is returned from FilterGovSetLiquidationCoolOffTime and is used to iterate over the raw logs and unpacked data for GovSetLiquidationCoolOffTime events raised by the EulerVault contract.
type EulerVaultGovSetLiquidationCoolOffTimeIterator struct {
	Event *EulerVaultGovSetLiquidationCoolOffTime // Event containing the contract specifics and raw log

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
func (it *EulerVaultGovSetLiquidationCoolOffTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultGovSetLiquidationCoolOffTime)
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
		it.Event = new(EulerVaultGovSetLiquidationCoolOffTime)
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
func (it *EulerVaultGovSetLiquidationCoolOffTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultGovSetLiquidationCoolOffTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultGovSetLiquidationCoolOffTime represents a GovSetLiquidationCoolOffTime event raised by the EulerVault contract.
type EulerVaultGovSetLiquidationCoolOffTime struct {
	NewCoolOffTime uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetLiquidationCoolOffTime is a free log retrieval operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EulerVault *EulerVaultFilterer) FilterGovSetLiquidationCoolOffTime(opts *bind.FilterOpts) (*EulerVaultGovSetLiquidationCoolOffTimeIterator, error) {

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "GovSetLiquidationCoolOffTime")
	if err != nil {
		return nil, err
	}
	return &EulerVaultGovSetLiquidationCoolOffTimeIterator{contract: _EulerVault.contract, event: "GovSetLiquidationCoolOffTime", logs: logs, sub: sub}, nil
}

// WatchGovSetLiquidationCoolOffTime is a free log subscription operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EulerVault *EulerVaultFilterer) WatchGovSetLiquidationCoolOffTime(opts *bind.WatchOpts, sink chan<- *EulerVaultGovSetLiquidationCoolOffTime) (event.Subscription, error) {

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "GovSetLiquidationCoolOffTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultGovSetLiquidationCoolOffTime)
				if err := _EulerVault.contract.UnpackLog(event, "GovSetLiquidationCoolOffTime", log); err != nil {
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

// ParseGovSetLiquidationCoolOffTime is a log parse operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EulerVault *EulerVaultFilterer) ParseGovSetLiquidationCoolOffTime(log types.Log) (*EulerVaultGovSetLiquidationCoolOffTime, error) {
	event := new(EulerVaultGovSetLiquidationCoolOffTime)
	if err := _EulerVault.contract.UnpackLog(event, "GovSetLiquidationCoolOffTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultGovSetMaxLiquidationDiscountIterator is returned from FilterGovSetMaxLiquidationDiscount and is used to iterate over the raw logs and unpacked data for GovSetMaxLiquidationDiscount events raised by the EulerVault contract.
type EulerVaultGovSetMaxLiquidationDiscountIterator struct {
	Event *EulerVaultGovSetMaxLiquidationDiscount // Event containing the contract specifics and raw log

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
func (it *EulerVaultGovSetMaxLiquidationDiscountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultGovSetMaxLiquidationDiscount)
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
		it.Event = new(EulerVaultGovSetMaxLiquidationDiscount)
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
func (it *EulerVaultGovSetMaxLiquidationDiscountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultGovSetMaxLiquidationDiscountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultGovSetMaxLiquidationDiscount represents a GovSetMaxLiquidationDiscount event raised by the EulerVault contract.
type EulerVaultGovSetMaxLiquidationDiscount struct {
	NewDiscount uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGovSetMaxLiquidationDiscount is a free log retrieval operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EulerVault *EulerVaultFilterer) FilterGovSetMaxLiquidationDiscount(opts *bind.FilterOpts) (*EulerVaultGovSetMaxLiquidationDiscountIterator, error) {

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "GovSetMaxLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return &EulerVaultGovSetMaxLiquidationDiscountIterator{contract: _EulerVault.contract, event: "GovSetMaxLiquidationDiscount", logs: logs, sub: sub}, nil
}

// WatchGovSetMaxLiquidationDiscount is a free log subscription operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EulerVault *EulerVaultFilterer) WatchGovSetMaxLiquidationDiscount(opts *bind.WatchOpts, sink chan<- *EulerVaultGovSetMaxLiquidationDiscount) (event.Subscription, error) {

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "GovSetMaxLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultGovSetMaxLiquidationDiscount)
				if err := _EulerVault.contract.UnpackLog(event, "GovSetMaxLiquidationDiscount", log); err != nil {
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

// ParseGovSetMaxLiquidationDiscount is a log parse operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EulerVault *EulerVaultFilterer) ParseGovSetMaxLiquidationDiscount(log types.Log) (*EulerVaultGovSetMaxLiquidationDiscount, error) {
	event := new(EulerVaultGovSetMaxLiquidationDiscount)
	if err := _EulerVault.contract.UnpackLog(event, "GovSetMaxLiquidationDiscount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultInterestAccruedIterator is returned from FilterInterestAccrued and is used to iterate over the raw logs and unpacked data for InterestAccrued events raised by the EulerVault contract.
type EulerVaultInterestAccruedIterator struct {
	Event *EulerVaultInterestAccrued // Event containing the contract specifics and raw log

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
func (it *EulerVaultInterestAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultInterestAccrued)
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
		it.Event = new(EulerVaultInterestAccrued)
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
func (it *EulerVaultInterestAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultInterestAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultInterestAccrued represents a InterestAccrued event raised by the EulerVault contract.
type EulerVaultInterestAccrued struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInterestAccrued is a free log retrieval operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) FilterInterestAccrued(opts *bind.FilterOpts, account []common.Address) (*EulerVaultInterestAccruedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "InterestAccrued", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultInterestAccruedIterator{contract: _EulerVault.contract, event: "InterestAccrued", logs: logs, sub: sub}, nil
}

// WatchInterestAccrued is a free log subscription operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) WatchInterestAccrued(opts *bind.WatchOpts, sink chan<- *EulerVaultInterestAccrued, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "InterestAccrued", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultInterestAccrued)
				if err := _EulerVault.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
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

// ParseInterestAccrued is a log parse operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) ParseInterestAccrued(log types.Log) (*EulerVaultInterestAccrued, error) {
	event := new(EulerVaultInterestAccrued)
	if err := _EulerVault.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the EulerVault contract.
type EulerVaultLiquidateIterator struct {
	Event *EulerVaultLiquidate // Event containing the contract specifics and raw log

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
func (it *EulerVaultLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultLiquidate)
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
		it.Event = new(EulerVaultLiquidate)
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
func (it *EulerVaultLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultLiquidate represents a Liquidate event raised by the EulerVault contract.
type EulerVaultLiquidate struct {
	Liquidator   common.Address
	Violator     common.Address
	Collateral   common.Address
	RepayAssets  *big.Int
	YieldBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EulerVault *EulerVaultFilterer) FilterLiquidate(opts *bind.FilterOpts, liquidator []common.Address, violator []common.Address) (*EulerVaultLiquidateIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var violatorRule []interface{}
	for _, violatorItem := range violator {
		violatorRule = append(violatorRule, violatorItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "Liquidate", liquidatorRule, violatorRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultLiquidateIterator{contract: _EulerVault.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EulerVault *EulerVaultFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *EulerVaultLiquidate, liquidator []common.Address, violator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var violatorRule []interface{}
	for _, violatorItem := range violator {
		violatorRule = append(violatorRule, violatorItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "Liquidate", liquidatorRule, violatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultLiquidate)
				if err := _EulerVault.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EulerVault *EulerVaultFilterer) ParseLiquidate(log types.Log) (*EulerVaultLiquidate, error) {
	event := new(EulerVaultLiquidate)
	if err := _EulerVault.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultPullDebtIterator is returned from FilterPullDebt and is used to iterate over the raw logs and unpacked data for PullDebt events raised by the EulerVault contract.
type EulerVaultPullDebtIterator struct {
	Event *EulerVaultPullDebt // Event containing the contract specifics and raw log

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
func (it *EulerVaultPullDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultPullDebt)
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
		it.Event = new(EulerVaultPullDebt)
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
func (it *EulerVaultPullDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultPullDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultPullDebt represents a PullDebt event raised by the EulerVault contract.
type EulerVaultPullDebt struct {
	From   common.Address
	To     common.Address
	Assets *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPullDebt is a free log retrieval operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EulerVault *EulerVaultFilterer) FilterPullDebt(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EulerVaultPullDebtIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "PullDebt", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultPullDebtIterator{contract: _EulerVault.contract, event: "PullDebt", logs: logs, sub: sub}, nil
}

// WatchPullDebt is a free log subscription operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EulerVault *EulerVaultFilterer) WatchPullDebt(opts *bind.WatchOpts, sink chan<- *EulerVaultPullDebt, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "PullDebt", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultPullDebt)
				if err := _EulerVault.contract.UnpackLog(event, "PullDebt", log); err != nil {
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

// ParsePullDebt is a log parse operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EulerVault *EulerVaultFilterer) ParsePullDebt(log types.Log) (*EulerVaultPullDebt, error) {
	event := new(EulerVaultPullDebt)
	if err := _EulerVault.contract.UnpackLog(event, "PullDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the EulerVault contract.
type EulerVaultRepayIterator struct {
	Event *EulerVaultRepay // Event containing the contract specifics and raw log

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
func (it *EulerVaultRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultRepay)
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
		it.Event = new(EulerVaultRepay)
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
func (it *EulerVaultRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultRepay represents a Repay event raised by the EulerVault contract.
type EulerVaultRepay struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) FilterRepay(opts *bind.FilterOpts, account []common.Address) (*EulerVaultRepayIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultRepayIterator{contract: _EulerVault.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *EulerVaultRepay, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultRepay)
				if err := _EulerVault.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EulerVault *EulerVaultFilterer) ParseRepay(log types.Log) (*EulerVaultRepay, error) {
	event := new(EulerVaultRepay)
	if err := _EulerVault.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EulerVault contract.
type EulerVaultTransferIterator struct {
	Event *EulerVaultTransfer // Event containing the contract specifics and raw log

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
func (it *EulerVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultTransfer)
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
		it.Event = new(EulerVaultTransfer)
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
func (it *EulerVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultTransfer represents a Transfer event raised by the EulerVault contract.
type EulerVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerVault *EulerVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EulerVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultTransferIterator{contract: _EulerVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerVault *EulerVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EulerVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultTransfer)
				if err := _EulerVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_EulerVault *EulerVaultFilterer) ParseTransfer(log types.Log) (*EulerVaultTransfer, error) {
	event := new(EulerVaultTransfer)
	if err := _EulerVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultVaultStatusIterator is returned from FilterVaultStatus and is used to iterate over the raw logs and unpacked data for VaultStatus events raised by the EulerVault contract.
type EulerVaultVaultStatusIterator struct {
	Event *EulerVaultVaultStatus // Event containing the contract specifics and raw log

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
func (it *EulerVaultVaultStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultVaultStatus)
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
		it.Event = new(EulerVaultVaultStatus)
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
func (it *EulerVaultVaultStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultVaultStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultVaultStatus represents a VaultStatus event raised by the EulerVault contract.
type EulerVaultVaultStatus struct {
	TotalShares         *big.Int
	TotalBorrows        *big.Int
	AccumulatedFees     *big.Int
	Cash                *big.Int
	InterestAccumulator *big.Int
	InterestRate        *big.Int
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterVaultStatus is a free log retrieval operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EulerVault *EulerVaultFilterer) FilterVaultStatus(opts *bind.FilterOpts) (*EulerVaultVaultStatusIterator, error) {

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "VaultStatus")
	if err != nil {
		return nil, err
	}
	return &EulerVaultVaultStatusIterator{contract: _EulerVault.contract, event: "VaultStatus", logs: logs, sub: sub}, nil
}

// WatchVaultStatus is a free log subscription operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EulerVault *EulerVaultFilterer) WatchVaultStatus(opts *bind.WatchOpts, sink chan<- *EulerVaultVaultStatus) (event.Subscription, error) {

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "VaultStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultVaultStatus)
				if err := _EulerVault.contract.UnpackLog(event, "VaultStatus", log); err != nil {
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

// ParseVaultStatus is a log parse operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EulerVault *EulerVaultFilterer) ParseVaultStatus(log types.Log) (*EulerVaultVaultStatus, error) {
	event := new(EulerVaultVaultStatus)
	if err := _EulerVault.contract.UnpackLog(event, "VaultStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the EulerVault contract.
type EulerVaultWithdrawIterator struct {
	Event *EulerVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *EulerVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerVaultWithdraw)
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
		it.Event = new(EulerVaultWithdraw)
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
func (it *EulerVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerVaultWithdraw represents a Withdraw event raised by the EulerVault contract.
type EulerVaultWithdraw struct {
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
func (_EulerVault *EulerVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*EulerVaultWithdrawIterator, error) {

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

	logs, sub, err := _EulerVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerVaultWithdrawIterator{contract: _EulerVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EulerVault *EulerVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EulerVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EulerVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerVaultWithdraw)
				if err := _EulerVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_EulerVault *EulerVaultFilterer) ParseWithdraw(log types.Log) (*EulerVaultWithdraw, error) {
	event := new(EulerVaultWithdraw)
	if err := _EulerVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
