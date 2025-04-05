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

// ILiquidStabilityPoolInitParams is an auto generated low-level Go binding around an user-defined struct.
type ILiquidStabilityPoolInitParams struct {
	Asset              common.Address
	SharesName         string
	SharesSymbol       string
	MetaBeraborrowCore common.Address
	LiquidationManager common.Address
	Factory            common.Address
	FeeReceiver        common.Address
}

// ILiquidStabilityPoolRebalanceParams is an auto generated low-level Go binding around an user-defined struct.
type ILiquidStabilityPoolRebalanceParams struct {
	SentCurrency     common.Address
	SentAmount       *big.Int
	ReceivedCurrency common.Address
	Swapper          common.Address
	Payload          []byte
}

// BeraBorrowSNECTMetaData contains all meta data concerning the BeraBorrowSNECT contract.
var BeraBorrowSNECTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceRemaining\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BelowThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BootstrapPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotFactory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotLM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralIsSunsetting\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollateralMustBeSunset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmissionRateExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExistingCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FactoryAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FactoryNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LMAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LMNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastTokenMustBeNect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPriceFeed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenCannotBeExtraAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenCannotBeNect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIsVesting\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenMustBeExtraAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawingLockedEmissions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroTotalSupply\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"AssetsWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCollateral\",\"type\":\"address\"}],\"name\":\"CollateralOverwritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"EmissionsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"EmissionsSub\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ExtraAssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ExtraAssetRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"unlockRatePerSecond\",\"type\":\"uint64\"}],\"name\":\"NewUnlockRatePerSecond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToOffset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collToAdd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collSurplusAmount\",\"type\":\"uint256\"}],\"name\":\"Offset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factoryRemoved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"LMremoved\",\"type\":\"address\"}],\"name\":\"ProtocolBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidationManager\",\"type\":\"address\"}],\"name\":\"ProtocolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sentCurrency\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receivedCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sentValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedValue\",\"type\":\"uint256\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SUNSET_DURATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_unlockRatePerSecond\",\"type\":\"uint64\"}],\"name\":\"addNewExtraAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_unlockRatePerSecond\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"forceThroughBalanceCheck\",\"type\":\"bool\"}],\"name\":\"enableCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"slots\",\"type\":\"bytes32[]\"}],\"name\":\"extSloads\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"res\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateralTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLockedEmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"scaledPriceInUsdWad\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDebtTokenDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_sharesName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sharesSymbol\",\"type\":\"string\"},{\"internalType\":\"contractIMetaBeraborrowCore\",\"name\":\"_metaBeraborrowCore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidationManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeReceiver\",\"type\":\"address\"}],\"internalType\":\"structILiquidStabilityPool.InitParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"amount\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"linearVestingExtraAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_debtToOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_collToAdd\",\"type\":\"uint256\"}],\"name\":\"offset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"netShares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sentCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sentAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receivedCurrency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structILiquidStabilityPool.RebalanceParams\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"preferredUnderlyingTokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeExtraAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_boycoVaults\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"enable\",\"type\":\"bool[]\"}],\"name\":\"setBoycoVaults\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"thresholdInBP\",\"type\":\"uint256\"}],\"name\":\"setPairThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_unlockRatePerSecond\",\"type\":\"uint64\"}],\"name\":\"setUnlockRatePerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"startCollateralSunset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInNect\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidationManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_register\",\"type\":\"bool\"}],\"name\":\"updateProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"preferredUnderlyingTokens\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BeraBorrowSNECTABI is the input ABI used to generate the binding from.
// Deprecated: Use BeraBorrowSNECTMetaData.ABI instead.
var BeraBorrowSNECTABI = BeraBorrowSNECTMetaData.ABI

// BeraBorrowSNECT is an auto generated Go binding around an Ethereum contract.
type BeraBorrowSNECT struct {
	BeraBorrowSNECTCaller     // Read-only binding to the contract
	BeraBorrowSNECTTransactor // Write-only binding to the contract
	BeraBorrowSNECTFilterer   // Log filterer for contract events
}

// BeraBorrowSNECTCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeraBorrowSNECTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowSNECTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeraBorrowSNECTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowSNECTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeraBorrowSNECTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowSNECTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeraBorrowSNECTSession struct {
	Contract     *BeraBorrowSNECT  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeraBorrowSNECTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeraBorrowSNECTCallerSession struct {
	Contract *BeraBorrowSNECTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BeraBorrowSNECTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeraBorrowSNECTTransactorSession struct {
	Contract     *BeraBorrowSNECTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BeraBorrowSNECTRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeraBorrowSNECTRaw struct {
	Contract *BeraBorrowSNECT // Generic contract binding to access the raw methods on
}

// BeraBorrowSNECTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeraBorrowSNECTCallerRaw struct {
	Contract *BeraBorrowSNECTCaller // Generic read-only contract binding to access the raw methods on
}

// BeraBorrowSNECTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeraBorrowSNECTTransactorRaw struct {
	Contract *BeraBorrowSNECTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeraBorrowSNECT creates a new instance of BeraBorrowSNECT, bound to a specific deployed contract.
func NewBeraBorrowSNECT(address common.Address, backend bind.ContractBackend) (*BeraBorrowSNECT, error) {
	contract, err := bindBeraBorrowSNECT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECT{BeraBorrowSNECTCaller: BeraBorrowSNECTCaller{contract: contract}, BeraBorrowSNECTTransactor: BeraBorrowSNECTTransactor{contract: contract}, BeraBorrowSNECTFilterer: BeraBorrowSNECTFilterer{contract: contract}}, nil
}

// NewBeraBorrowSNECTCaller creates a new read-only instance of BeraBorrowSNECT, bound to a specific deployed contract.
func NewBeraBorrowSNECTCaller(address common.Address, caller bind.ContractCaller) (*BeraBorrowSNECTCaller, error) {
	contract, err := bindBeraBorrowSNECT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTCaller{contract: contract}, nil
}

// NewBeraBorrowSNECTTransactor creates a new write-only instance of BeraBorrowSNECT, bound to a specific deployed contract.
func NewBeraBorrowSNECTTransactor(address common.Address, transactor bind.ContractTransactor) (*BeraBorrowSNECTTransactor, error) {
	contract, err := bindBeraBorrowSNECT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTTransactor{contract: contract}, nil
}

// NewBeraBorrowSNECTFilterer creates a new log filterer instance of BeraBorrowSNECT, bound to a specific deployed contract.
func NewBeraBorrowSNECTFilterer(address common.Address, filterer bind.ContractFilterer) (*BeraBorrowSNECTFilterer, error) {
	contract, err := bindBeraBorrowSNECT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTFilterer{contract: contract}, nil
}

// bindBeraBorrowSNECT binds a generic wrapper to an already deployed contract.
func bindBeraBorrowSNECT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeraBorrowSNECTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeraBorrowSNECT *BeraBorrowSNECTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeraBorrowSNECT.Contract.BeraBorrowSNECTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeraBorrowSNECT *BeraBorrowSNECTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.BeraBorrowSNECTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeraBorrowSNECT *BeraBorrowSNECTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.BeraBorrowSNECTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeraBorrowSNECT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.contract.Transact(opts, method, params...)
}

// SUNSETDURATION is a free data retrieval call binding the contract method 0xa7528a03.
//
// Solidity: function SUNSET_DURATION() view returns(uint128)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) SUNSETDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "SUNSET_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUNSETDURATION is a free data retrieval call binding the contract method 0xa7528a03.
//
// Solidity: function SUNSET_DURATION() view returns(uint128)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) SUNSETDURATION() (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.SUNSETDURATION(&_BeraBorrowSNECT.CallOpts)
}

// SUNSETDURATION is a free data retrieval call binding the contract method 0xa7528a03.
//
// Solidity: function SUNSET_DURATION() view returns(uint128)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) SUNSETDURATION() (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.SUNSETDURATION(&_BeraBorrowSNECT.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BeraBorrowSNECT.Contract.UPGRADEINTERFACEVERSION(&_BeraBorrowSNECT.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BeraBorrowSNECT.Contract.UPGRADEINTERFACEVERSION(&_BeraBorrowSNECT.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.Allowance(&_BeraBorrowSNECT.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.Allowance(&_BeraBorrowSNECT.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Asset() (common.Address, error) {
	return _BeraBorrowSNECT.Contract.Asset(&_BeraBorrowSNECT.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) Asset() (common.Address, error) {
	return _BeraBorrowSNECT.Contract.Asset(&_BeraBorrowSNECT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.BalanceOf(&_BeraBorrowSNECT.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.BalanceOf(&_BeraBorrowSNECT.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.ConvertToAssets(&_BeraBorrowSNECT.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.ConvertToAssets(&_BeraBorrowSNECT.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.ConvertToShares(&_BeraBorrowSNECT.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.ConvertToShares(&_BeraBorrowSNECT.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Decimals() (uint8, error) {
	return _BeraBorrowSNECT.Contract.Decimals(&_BeraBorrowSNECT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) Decimals() (uint8, error) {
	return _BeraBorrowSNECT.Contract.Decimals(&_BeraBorrowSNECT.CallOpts)
}

// ExtSloads is a free data retrieval call binding the contract method 0x7784c685.
//
// Solidity: function extSloads(bytes32[] slots) view returns(bytes32[] res)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) ExtSloads(opts *bind.CallOpts, slots [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "extSloads", slots)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ExtSloads is a free data retrieval call binding the contract method 0x7784c685.
//
// Solidity: function extSloads(bytes32[] slots) view returns(bytes32[] res)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) ExtSloads(slots [][32]byte) ([][32]byte, error) {
	return _BeraBorrowSNECT.Contract.ExtSloads(&_BeraBorrowSNECT.CallOpts, slots)
}

// ExtSloads is a free data retrieval call binding the contract method 0x7784c685.
//
// Solidity: function extSloads(bytes32[] slots) view returns(bytes32[] res)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) ExtSloads(slots [][32]byte) ([][32]byte, error) {
	return _BeraBorrowSNECT.Contract.ExtSloads(&_BeraBorrowSNECT.CallOpts, slots)
}

// GetCollateralTokens is a free data retrieval call binding the contract method 0xb58eb63f.
//
// Solidity: function getCollateralTokens() view returns(address[])
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) GetCollateralTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "getCollateralTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCollateralTokens is a free data retrieval call binding the contract method 0xb58eb63f.
//
// Solidity: function getCollateralTokens() view returns(address[])
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) GetCollateralTokens() ([]common.Address, error) {
	return _BeraBorrowSNECT.Contract.GetCollateralTokens(&_BeraBorrowSNECT.CallOpts)
}

// GetCollateralTokens is a free data retrieval call binding the contract method 0xb58eb63f.
//
// Solidity: function getCollateralTokens() view returns(address[])
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) GetCollateralTokens() ([]common.Address, error) {
	return _BeraBorrowSNECT.Contract.GetCollateralTokens(&_BeraBorrowSNECT.CallOpts)
}

// GetLockedEmissions is a free data retrieval call binding the contract method 0x403dd3bc.
//
// Solidity: function getLockedEmissions(address token) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) GetLockedEmissions(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "getLockedEmissions", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedEmissions is a free data retrieval call binding the contract method 0x403dd3bc.
//
// Solidity: function getLockedEmissions(address token) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) GetLockedEmissions(token common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.GetLockedEmissions(&_BeraBorrowSNECT.CallOpts, token)
}

// GetLockedEmissions is a free data retrieval call binding the contract method 0x403dd3bc.
//
// Solidity: function getLockedEmissions(address token) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) GetLockedEmissions(token common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.GetLockedEmissions(&_BeraBorrowSNECT.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256 scaledPriceInUsdWad)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) GetPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "getPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256 scaledPriceInUsdWad)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) GetPrice(token common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.GetPrice(&_BeraBorrowSNECT.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256 scaledPriceInUsdWad)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) GetPrice(token common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.GetPrice(&_BeraBorrowSNECT.CallOpts, token)
}

// GetTotalDebtTokenDeposits is a free data retrieval call binding the contract method 0x0d9a6b35.
//
// Solidity: function getTotalDebtTokenDeposits() view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) GetTotalDebtTokenDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "getTotalDebtTokenDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDebtTokenDeposits is a free data retrieval call binding the contract method 0x0d9a6b35.
//
// Solidity: function getTotalDebtTokenDeposits() view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) GetTotalDebtTokenDeposits() (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.GetTotalDebtTokenDeposits(&_BeraBorrowSNECT.CallOpts)
}

// GetTotalDebtTokenDeposits is a free data retrieval call binding the contract method 0x0d9a6b35.
//
// Solidity: function getTotalDebtTokenDeposits() view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) GetTotalDebtTokenDeposits() (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.GetTotalDebtTokenDeposits(&_BeraBorrowSNECT.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.MaxDeposit(&_BeraBorrowSNECT.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.MaxDeposit(&_BeraBorrowSNECT.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.MaxMint(&_BeraBorrowSNECT.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.MaxMint(&_BeraBorrowSNECT.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.MaxRedeem(&_BeraBorrowSNECT.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.MaxRedeem(&_BeraBorrowSNECT.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.MaxWithdraw(&_BeraBorrowSNECT.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.MaxWithdraw(&_BeraBorrowSNECT.CallOpts, _owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Name() (string, error) {
	return _BeraBorrowSNECT.Contract.Name(&_BeraBorrowSNECT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) Name() (string, error) {
	return _BeraBorrowSNECT.Contract.Name(&_BeraBorrowSNECT.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.PreviewDeposit(&_BeraBorrowSNECT.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.PreviewDeposit(&_BeraBorrowSNECT.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 netShares) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) PreviewMint(opts *bind.CallOpts, netShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "previewMint", netShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 netShares) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) PreviewMint(netShares *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.PreviewMint(&_BeraBorrowSNECT.CallOpts, netShares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 netShares) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) PreviewMint(netShares *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.PreviewMint(&_BeraBorrowSNECT.CallOpts, netShares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.PreviewRedeem(&_BeraBorrowSNECT.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.PreviewRedeem(&_BeraBorrowSNECT.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.PreviewWithdraw(&_BeraBorrowSNECT.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.PreviewWithdraw(&_BeraBorrowSNECT.CallOpts, assets)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) ProxiableUUID() ([32]byte, error) {
	return _BeraBorrowSNECT.Contract.ProxiableUUID(&_BeraBorrowSNECT.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BeraBorrowSNECT.Contract.ProxiableUUID(&_BeraBorrowSNECT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Symbol() (string, error) {
	return _BeraBorrowSNECT.Contract.Symbol(&_BeraBorrowSNECT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) Symbol() (string, error) {
	return _BeraBorrowSNECT.Contract.Symbol(&_BeraBorrowSNECT.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 amountInNect)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 amountInNect)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) TotalAssets() (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.TotalAssets(&_BeraBorrowSNECT.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 amountInNect)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) TotalAssets() (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.TotalAssets(&_BeraBorrowSNECT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowSNECT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) TotalSupply() (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.TotalSupply(&_BeraBorrowSNECT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowSNECT *BeraBorrowSNECTCallerSession) TotalSupply() (*big.Int, error) {
	return _BeraBorrowSNECT.Contract.TotalSupply(&_BeraBorrowSNECT.CallOpts)
}

// AddNewExtraAsset is a paid mutator transaction binding the contract method 0xfcb99905.
//
// Solidity: function addNewExtraAsset(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) AddNewExtraAsset(opts *bind.TransactOpts, token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "addNewExtraAsset", token, _unlockRatePerSecond)
}

// AddNewExtraAsset is a paid mutator transaction binding the contract method 0xfcb99905.
//
// Solidity: function addNewExtraAsset(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) AddNewExtraAsset(token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.AddNewExtraAsset(&_BeraBorrowSNECT.TransactOpts, token, _unlockRatePerSecond)
}

// AddNewExtraAsset is a paid mutator transaction binding the contract method 0xfcb99905.
//
// Solidity: function addNewExtraAsset(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) AddNewExtraAsset(token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.AddNewExtraAsset(&_BeraBorrowSNECT.TransactOpts, token, _unlockRatePerSecond)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Approve(&_BeraBorrowSNECT.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Approve(&_BeraBorrowSNECT.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Deposit(&_BeraBorrowSNECT.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Deposit(&_BeraBorrowSNECT.TransactOpts, assets, receiver)
}

// EnableCollateral is a paid mutator transaction binding the contract method 0x602ecae5.
//
// Solidity: function enableCollateral(address _collateral, uint64 _unlockRatePerSecond, bool forceThroughBalanceCheck) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) EnableCollateral(opts *bind.TransactOpts, _collateral common.Address, _unlockRatePerSecond uint64, forceThroughBalanceCheck bool) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "enableCollateral", _collateral, _unlockRatePerSecond, forceThroughBalanceCheck)
}

// EnableCollateral is a paid mutator transaction binding the contract method 0x602ecae5.
//
// Solidity: function enableCollateral(address _collateral, uint64 _unlockRatePerSecond, bool forceThroughBalanceCheck) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) EnableCollateral(_collateral common.Address, _unlockRatePerSecond uint64, forceThroughBalanceCheck bool) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.EnableCollateral(&_BeraBorrowSNECT.TransactOpts, _collateral, _unlockRatePerSecond, forceThroughBalanceCheck)
}

// EnableCollateral is a paid mutator transaction binding the contract method 0x602ecae5.
//
// Solidity: function enableCollateral(address _collateral, uint64 _unlockRatePerSecond, bool forceThroughBalanceCheck) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) EnableCollateral(_collateral common.Address, _unlockRatePerSecond uint64, forceThroughBalanceCheck bool) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.EnableCollateral(&_BeraBorrowSNECT.TransactOpts, _collateral, _unlockRatePerSecond, forceThroughBalanceCheck)
}

// Initialize is a paid mutator transaction binding the contract method 0xfd694144.
//
// Solidity: function initialize((address,string,string,address,address,address,address) params) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Initialize(opts *bind.TransactOpts, params ILiquidStabilityPoolInitParams) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "initialize", params)
}

// Initialize is a paid mutator transaction binding the contract method 0xfd694144.
//
// Solidity: function initialize((address,string,string,address,address,address,address) params) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Initialize(params ILiquidStabilityPoolInitParams) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Initialize(&_BeraBorrowSNECT.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xfd694144.
//
// Solidity: function initialize((address,string,string,address,address,address,address) params) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Initialize(params ILiquidStabilityPoolInitParams) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Initialize(&_BeraBorrowSNECT.TransactOpts, params)
}

// LinearVestingExtraAssets is a paid mutator transaction binding the contract method 0xc5e6a767.
//
// Solidity: function linearVestingExtraAssets(address token, int256 amount, address recipient) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) LinearVestingExtraAssets(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "linearVestingExtraAssets", token, amount, recipient)
}

// LinearVestingExtraAssets is a paid mutator transaction binding the contract method 0xc5e6a767.
//
// Solidity: function linearVestingExtraAssets(address token, int256 amount, address recipient) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) LinearVestingExtraAssets(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.LinearVestingExtraAssets(&_BeraBorrowSNECT.TransactOpts, token, amount, recipient)
}

// LinearVestingExtraAssets is a paid mutator transaction binding the contract method 0xc5e6a767.
//
// Solidity: function linearVestingExtraAssets(address token, int256 amount, address recipient) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) LinearVestingExtraAssets(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.LinearVestingExtraAssets(&_BeraBorrowSNECT.TransactOpts, token, amount, recipient)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Mint(&_BeraBorrowSNECT.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Mint(&_BeraBorrowSNECT.TransactOpts, shares, receiver)
}

// Offset is a paid mutator transaction binding the contract method 0xe6666733.
//
// Solidity: function offset(address collateral, uint256 _debtToOffset, uint256 _collToAdd) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Offset(opts *bind.TransactOpts, collateral common.Address, _debtToOffset *big.Int, _collToAdd *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "offset", collateral, _debtToOffset, _collToAdd)
}

// Offset is a paid mutator transaction binding the contract method 0xe6666733.
//
// Solidity: function offset(address collateral, uint256 _debtToOffset, uint256 _collToAdd) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Offset(collateral common.Address, _debtToOffset *big.Int, _collToAdd *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Offset(&_BeraBorrowSNECT.TransactOpts, collateral, _debtToOffset, _collToAdd)
}

// Offset is a paid mutator transaction binding the contract method 0xe6666733.
//
// Solidity: function offset(address collateral, uint256 _debtToOffset, uint256 _collToAdd) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Offset(collateral common.Address, _debtToOffset *big.Int, _collToAdd *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Offset(&_BeraBorrowSNECT.TransactOpts, collateral, _debtToOffset, _collToAdd)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7facd79b.
//
// Solidity: function rebalance((address,uint256,address,address,bytes) p) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Rebalance(opts *bind.TransactOpts, p ILiquidStabilityPoolRebalanceParams) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "rebalance", p)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7facd79b.
//
// Solidity: function rebalance((address,uint256,address,address,bytes) p) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Rebalance(p ILiquidStabilityPoolRebalanceParams) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Rebalance(&_BeraBorrowSNECT.TransactOpts, p)
}

// Rebalance is a paid mutator transaction binding the contract method 0x7facd79b.
//
// Solidity: function rebalance((address,uint256,address,address,bytes) p) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Rebalance(p ILiquidStabilityPoolRebalanceParams) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Rebalance(&_BeraBorrowSNECT.TransactOpts, p)
}

// Redeem is a paid mutator transaction binding the contract method 0x31e95162.
//
// Solidity: function redeem(uint256 shares, address[] preferredUnderlyingTokens, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, preferredUnderlyingTokens []common.Address, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "redeem", shares, preferredUnderlyingTokens, receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0x31e95162.
//
// Solidity: function redeem(uint256 shares, address[] preferredUnderlyingTokens, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Redeem(shares *big.Int, preferredUnderlyingTokens []common.Address, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Redeem(&_BeraBorrowSNECT.TransactOpts, shares, preferredUnderlyingTokens, receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0x31e95162.
//
// Solidity: function redeem(uint256 shares, address[] preferredUnderlyingTokens, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Redeem(shares *big.Int, preferredUnderlyingTokens []common.Address, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Redeem(&_BeraBorrowSNECT.TransactOpts, shares, preferredUnderlyingTokens, receiver, _owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Redeem0(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "redeem0", shares, receiver, _owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Redeem0(shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Redeem0(&_BeraBorrowSNECT.TransactOpts, shares, receiver, _owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Redeem0(shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Redeem0(&_BeraBorrowSNECT.TransactOpts, shares, receiver, _owner)
}

// RemoveExtraAsset is a paid mutator transaction binding the contract method 0x6cd611be.
//
// Solidity: function removeExtraAsset(address token) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) RemoveExtraAsset(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "removeExtraAsset", token)
}

// RemoveExtraAsset is a paid mutator transaction binding the contract method 0x6cd611be.
//
// Solidity: function removeExtraAsset(address token) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) RemoveExtraAsset(token common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.RemoveExtraAsset(&_BeraBorrowSNECT.TransactOpts, token)
}

// RemoveExtraAsset is a paid mutator transaction binding the contract method 0x6cd611be.
//
// Solidity: function removeExtraAsset(address token) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) RemoveExtraAsset(token common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.RemoveExtraAsset(&_BeraBorrowSNECT.TransactOpts, token)
}

// SetBoycoVaults is a paid mutator transaction binding the contract method 0x7d4601c0.
//
// Solidity: function setBoycoVaults(address[] _boycoVaults, bool[] enable) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) SetBoycoVaults(opts *bind.TransactOpts, _boycoVaults []common.Address, enable []bool) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "setBoycoVaults", _boycoVaults, enable)
}

// SetBoycoVaults is a paid mutator transaction binding the contract method 0x7d4601c0.
//
// Solidity: function setBoycoVaults(address[] _boycoVaults, bool[] enable) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) SetBoycoVaults(_boycoVaults []common.Address, enable []bool) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.SetBoycoVaults(&_BeraBorrowSNECT.TransactOpts, _boycoVaults, enable)
}

// SetBoycoVaults is a paid mutator transaction binding the contract method 0x7d4601c0.
//
// Solidity: function setBoycoVaults(address[] _boycoVaults, bool[] enable) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) SetBoycoVaults(_boycoVaults []common.Address, enable []bool) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.SetBoycoVaults(&_BeraBorrowSNECT.TransactOpts, _boycoVaults, enable)
}

// SetPairThreshold is a paid mutator transaction binding the contract method 0xbadd8b2d.
//
// Solidity: function setPairThreshold(address tokenIn, address tokenOut, uint256 thresholdInBP) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) SetPairThreshold(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, thresholdInBP *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "setPairThreshold", tokenIn, tokenOut, thresholdInBP)
}

// SetPairThreshold is a paid mutator transaction binding the contract method 0xbadd8b2d.
//
// Solidity: function setPairThreshold(address tokenIn, address tokenOut, uint256 thresholdInBP) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) SetPairThreshold(tokenIn common.Address, tokenOut common.Address, thresholdInBP *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.SetPairThreshold(&_BeraBorrowSNECT.TransactOpts, tokenIn, tokenOut, thresholdInBP)
}

// SetPairThreshold is a paid mutator transaction binding the contract method 0xbadd8b2d.
//
// Solidity: function setPairThreshold(address tokenIn, address tokenOut, uint256 thresholdInBP) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) SetPairThreshold(tokenIn common.Address, tokenOut common.Address, thresholdInBP *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.SetPairThreshold(&_BeraBorrowSNECT.TransactOpts, tokenIn, tokenOut, thresholdInBP)
}

// SetUnlockRatePerSecond is a paid mutator transaction binding the contract method 0x0b983a74.
//
// Solidity: function setUnlockRatePerSecond(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) SetUnlockRatePerSecond(opts *bind.TransactOpts, token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "setUnlockRatePerSecond", token, _unlockRatePerSecond)
}

// SetUnlockRatePerSecond is a paid mutator transaction binding the contract method 0x0b983a74.
//
// Solidity: function setUnlockRatePerSecond(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) SetUnlockRatePerSecond(token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.SetUnlockRatePerSecond(&_BeraBorrowSNECT.TransactOpts, token, _unlockRatePerSecond)
}

// SetUnlockRatePerSecond is a paid mutator transaction binding the contract method 0x0b983a74.
//
// Solidity: function setUnlockRatePerSecond(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) SetUnlockRatePerSecond(token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.SetUnlockRatePerSecond(&_BeraBorrowSNECT.TransactOpts, token, _unlockRatePerSecond)
}

// StartCollateralSunset is a paid mutator transaction binding the contract method 0x19f27b3b.
//
// Solidity: function startCollateralSunset(address collateral) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) StartCollateralSunset(opts *bind.TransactOpts, collateral common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "startCollateralSunset", collateral)
}

// StartCollateralSunset is a paid mutator transaction binding the contract method 0x19f27b3b.
//
// Solidity: function startCollateralSunset(address collateral) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) StartCollateralSunset(collateral common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.StartCollateralSunset(&_BeraBorrowSNECT.TransactOpts, collateral)
}

// StartCollateralSunset is a paid mutator transaction binding the contract method 0x19f27b3b.
//
// Solidity: function startCollateralSunset(address collateral) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) StartCollateralSunset(collateral common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.StartCollateralSunset(&_BeraBorrowSNECT.TransactOpts, collateral)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Transfer(&_BeraBorrowSNECT.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Transfer(&_BeraBorrowSNECT.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.TransferFrom(&_BeraBorrowSNECT.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.TransferFrom(&_BeraBorrowSNECT.TransactOpts, from, to, value)
}

// UpdateProtocol is a paid mutator transaction binding the contract method 0x7b8b8b74.
//
// Solidity: function updateProtocol(address _liquidationManager, address _factory, bool _register) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) UpdateProtocol(opts *bind.TransactOpts, _liquidationManager common.Address, _factory common.Address, _register bool) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "updateProtocol", _liquidationManager, _factory, _register)
}

// UpdateProtocol is a paid mutator transaction binding the contract method 0x7b8b8b74.
//
// Solidity: function updateProtocol(address _liquidationManager, address _factory, bool _register) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) UpdateProtocol(_liquidationManager common.Address, _factory common.Address, _register bool) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.UpdateProtocol(&_BeraBorrowSNECT.TransactOpts, _liquidationManager, _factory, _register)
}

// UpdateProtocol is a paid mutator transaction binding the contract method 0x7b8b8b74.
//
// Solidity: function updateProtocol(address _liquidationManager, address _factory, bool _register) returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) UpdateProtocol(_liquidationManager common.Address, _factory common.Address, _register bool) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.UpdateProtocol(&_BeraBorrowSNECT.TransactOpts, _liquidationManager, _factory, _register)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.UpgradeToAndCall(&_BeraBorrowSNECT.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.UpgradeToAndCall(&_BeraBorrowSNECT.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "withdraw", assets, receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Withdraw(assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Withdraw(&_BeraBorrowSNECT.TransactOpts, assets, receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Withdraw(assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Withdraw(&_BeraBorrowSNECT.TransactOpts, assets, receiver, _owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xddda679b.
//
// Solidity: function withdraw(uint256 assets, address[] preferredUnderlyingTokens, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactor) Withdraw0(opts *bind.TransactOpts, assets *big.Int, preferredUnderlyingTokens []common.Address, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.contract.Transact(opts, "withdraw0", assets, preferredUnderlyingTokens, receiver, _owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xddda679b.
//
// Solidity: function withdraw(uint256 assets, address[] preferredUnderlyingTokens, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTSession) Withdraw0(assets *big.Int, preferredUnderlyingTokens []common.Address, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Withdraw0(&_BeraBorrowSNECT.TransactOpts, assets, preferredUnderlyingTokens, receiver, _owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xddda679b.
//
// Solidity: function withdraw(uint256 assets, address[] preferredUnderlyingTokens, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTTransactorSession) Withdraw0(assets *big.Int, preferredUnderlyingTokens []common.Address, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowSNECT.Contract.Withdraw0(&_BeraBorrowSNECT.TransactOpts, assets, preferredUnderlyingTokens, receiver, _owner)
}

// BeraBorrowSNECTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTApprovalIterator struct {
	Event *BeraBorrowSNECTApproval // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTApproval)
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
		it.Event = new(BeraBorrowSNECTApproval)
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
func (it *BeraBorrowSNECTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTApproval represents a Approval event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BeraBorrowSNECTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTApprovalIterator{contract: _BeraBorrowSNECT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTApproval)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseApproval(log types.Log) (*BeraBorrowSNECTApproval, error) {
	event := new(BeraBorrowSNECTApproval)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTAssetsWithdrawIterator is returned from FilterAssetsWithdraw and is used to iterate over the raw logs and unpacked data for AssetsWithdraw events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTAssetsWithdrawIterator struct {
	Event *BeraBorrowSNECTAssetsWithdraw // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTAssetsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTAssetsWithdraw)
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
		it.Event = new(BeraBorrowSNECTAssetsWithdraw)
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
func (it *BeraBorrowSNECTAssetsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTAssetsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTAssetsWithdraw represents a AssetsWithdraw event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTAssetsWithdraw struct {
	Receiver common.Address
	Shares   *big.Int
	Tokens   []common.Address
	Amounts  []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAssetsWithdraw is a free log retrieval operation binding the contract event 0x86bcb277da75a9fbb738b8bb82beb731d82a0a89516b848730df92849f966bf0.
//
// Solidity: event AssetsWithdraw(address indexed receiver, uint256 shares, address[] tokens, uint256[] amounts)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterAssetsWithdraw(opts *bind.FilterOpts, receiver []common.Address) (*BeraBorrowSNECTAssetsWithdrawIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "AssetsWithdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTAssetsWithdrawIterator{contract: _BeraBorrowSNECT.contract, event: "AssetsWithdraw", logs: logs, sub: sub}, nil
}

// WatchAssetsWithdraw is a free log subscription operation binding the contract event 0x86bcb277da75a9fbb738b8bb82beb731d82a0a89516b848730df92849f966bf0.
//
// Solidity: event AssetsWithdraw(address indexed receiver, uint256 shares, address[] tokens, uint256[] amounts)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchAssetsWithdraw(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTAssetsWithdraw, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "AssetsWithdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTAssetsWithdraw)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "AssetsWithdraw", log); err != nil {
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

// ParseAssetsWithdraw is a log parse operation binding the contract event 0x86bcb277da75a9fbb738b8bb82beb731d82a0a89516b848730df92849f966bf0.
//
// Solidity: event AssetsWithdraw(address indexed receiver, uint256 shares, address[] tokens, uint256[] amounts)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseAssetsWithdraw(log types.Log) (*BeraBorrowSNECTAssetsWithdraw, error) {
	event := new(BeraBorrowSNECTAssetsWithdraw)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "AssetsWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTCollateralOverwrittenIterator is returned from FilterCollateralOverwritten and is used to iterate over the raw logs and unpacked data for CollateralOverwritten events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTCollateralOverwrittenIterator struct {
	Event *BeraBorrowSNECTCollateralOverwritten // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTCollateralOverwrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTCollateralOverwritten)
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
		it.Event = new(BeraBorrowSNECTCollateralOverwritten)
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
func (it *BeraBorrowSNECTCollateralOverwrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTCollateralOverwrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTCollateralOverwritten represents a CollateralOverwritten event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTCollateralOverwritten struct {
	OldCollateral common.Address
	NewCollateral common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollateralOverwritten is a free log retrieval operation binding the contract event 0x9e147d339c63698deb55c3d0d44ed3eba29bac2a068a88c4bc5bde17d6331e19.
//
// Solidity: event CollateralOverwritten(address oldCollateral, address newCollateral)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterCollateralOverwritten(opts *bind.FilterOpts) (*BeraBorrowSNECTCollateralOverwrittenIterator, error) {

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "CollateralOverwritten")
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTCollateralOverwrittenIterator{contract: _BeraBorrowSNECT.contract, event: "CollateralOverwritten", logs: logs, sub: sub}, nil
}

// WatchCollateralOverwritten is a free log subscription operation binding the contract event 0x9e147d339c63698deb55c3d0d44ed3eba29bac2a068a88c4bc5bde17d6331e19.
//
// Solidity: event CollateralOverwritten(address oldCollateral, address newCollateral)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchCollateralOverwritten(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTCollateralOverwritten) (event.Subscription, error) {

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "CollateralOverwritten")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTCollateralOverwritten)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "CollateralOverwritten", log); err != nil {
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

// ParseCollateralOverwritten is a log parse operation binding the contract event 0x9e147d339c63698deb55c3d0d44ed3eba29bac2a068a88c4bc5bde17d6331e19.
//
// Solidity: event CollateralOverwritten(address oldCollateral, address newCollateral)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseCollateralOverwritten(log types.Log) (*BeraBorrowSNECTCollateralOverwritten, error) {
	event := new(BeraBorrowSNECTCollateralOverwritten)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "CollateralOverwritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTDepositIterator struct {
	Event *BeraBorrowSNECTDeposit // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTDeposit)
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
		it.Event = new(BeraBorrowSNECTDeposit)
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
func (it *BeraBorrowSNECTDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTDeposit represents a Deposit event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*BeraBorrowSNECTDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTDepositIterator{contract: _BeraBorrowSNECT.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTDeposit)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseDeposit(log types.Log) (*BeraBorrowSNECTDeposit, error) {
	event := new(BeraBorrowSNECTDeposit)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTEmissionsAddedIterator is returned from FilterEmissionsAdded and is used to iterate over the raw logs and unpacked data for EmissionsAdded events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTEmissionsAddedIterator struct {
	Event *BeraBorrowSNECTEmissionsAdded // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTEmissionsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTEmissionsAdded)
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
		it.Event = new(BeraBorrowSNECTEmissionsAdded)
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
func (it *BeraBorrowSNECTEmissionsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTEmissionsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTEmissionsAdded represents a EmissionsAdded event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTEmissionsAdded struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmissionsAdded is a free log retrieval operation binding the contract event 0x693ffe037fd29f4846b006bd3ada57d4fd4c3622277227342f0e4fed9011dc20.
//
// Solidity: event EmissionsAdded(address indexed token, uint128 amount)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterEmissionsAdded(opts *bind.FilterOpts, token []common.Address) (*BeraBorrowSNECTEmissionsAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "EmissionsAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTEmissionsAddedIterator{contract: _BeraBorrowSNECT.contract, event: "EmissionsAdded", logs: logs, sub: sub}, nil
}

// WatchEmissionsAdded is a free log subscription operation binding the contract event 0x693ffe037fd29f4846b006bd3ada57d4fd4c3622277227342f0e4fed9011dc20.
//
// Solidity: event EmissionsAdded(address indexed token, uint128 amount)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchEmissionsAdded(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTEmissionsAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "EmissionsAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTEmissionsAdded)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "EmissionsAdded", log); err != nil {
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

// ParseEmissionsAdded is a log parse operation binding the contract event 0x693ffe037fd29f4846b006bd3ada57d4fd4c3622277227342f0e4fed9011dc20.
//
// Solidity: event EmissionsAdded(address indexed token, uint128 amount)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseEmissionsAdded(log types.Log) (*BeraBorrowSNECTEmissionsAdded, error) {
	event := new(BeraBorrowSNECTEmissionsAdded)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "EmissionsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTEmissionsSubIterator is returned from FilterEmissionsSub and is used to iterate over the raw logs and unpacked data for EmissionsSub events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTEmissionsSubIterator struct {
	Event *BeraBorrowSNECTEmissionsSub // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTEmissionsSubIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTEmissionsSub)
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
		it.Event = new(BeraBorrowSNECTEmissionsSub)
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
func (it *BeraBorrowSNECTEmissionsSubIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTEmissionsSubIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTEmissionsSub represents a EmissionsSub event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTEmissionsSub struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmissionsSub is a free log retrieval operation binding the contract event 0x5f4e7177e0f8e013ddb6d29e468fa7a45f8df4e00e7895b7f20c5979cab21c6c.
//
// Solidity: event EmissionsSub(address indexed token, uint128 amount)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterEmissionsSub(opts *bind.FilterOpts, token []common.Address) (*BeraBorrowSNECTEmissionsSubIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "EmissionsSub", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTEmissionsSubIterator{contract: _BeraBorrowSNECT.contract, event: "EmissionsSub", logs: logs, sub: sub}, nil
}

// WatchEmissionsSub is a free log subscription operation binding the contract event 0x5f4e7177e0f8e013ddb6d29e468fa7a45f8df4e00e7895b7f20c5979cab21c6c.
//
// Solidity: event EmissionsSub(address indexed token, uint128 amount)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchEmissionsSub(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTEmissionsSub, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "EmissionsSub", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTEmissionsSub)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "EmissionsSub", log); err != nil {
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

// ParseEmissionsSub is a log parse operation binding the contract event 0x5f4e7177e0f8e013ddb6d29e468fa7a45f8df4e00e7895b7f20c5979cab21c6c.
//
// Solidity: event EmissionsSub(address indexed token, uint128 amount)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseEmissionsSub(log types.Log) (*BeraBorrowSNECTEmissionsSub, error) {
	event := new(BeraBorrowSNECTEmissionsSub)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "EmissionsSub", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTExtraAssetAddedIterator is returned from FilterExtraAssetAdded and is used to iterate over the raw logs and unpacked data for ExtraAssetAdded events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTExtraAssetAddedIterator struct {
	Event *BeraBorrowSNECTExtraAssetAdded // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTExtraAssetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTExtraAssetAdded)
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
		it.Event = new(BeraBorrowSNECTExtraAssetAdded)
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
func (it *BeraBorrowSNECTExtraAssetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTExtraAssetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTExtraAssetAdded represents a ExtraAssetAdded event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTExtraAssetAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtraAssetAdded is a free log retrieval operation binding the contract event 0x252fb22f1e5dcdba04908f13259852204aead54fea1342d028eb2f49510bee97.
//
// Solidity: event ExtraAssetAdded(address token)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterExtraAssetAdded(opts *bind.FilterOpts) (*BeraBorrowSNECTExtraAssetAddedIterator, error) {

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "ExtraAssetAdded")
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTExtraAssetAddedIterator{contract: _BeraBorrowSNECT.contract, event: "ExtraAssetAdded", logs: logs, sub: sub}, nil
}

// WatchExtraAssetAdded is a free log subscription operation binding the contract event 0x252fb22f1e5dcdba04908f13259852204aead54fea1342d028eb2f49510bee97.
//
// Solidity: event ExtraAssetAdded(address token)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchExtraAssetAdded(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTExtraAssetAdded) (event.Subscription, error) {

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "ExtraAssetAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTExtraAssetAdded)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "ExtraAssetAdded", log); err != nil {
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

// ParseExtraAssetAdded is a log parse operation binding the contract event 0x252fb22f1e5dcdba04908f13259852204aead54fea1342d028eb2f49510bee97.
//
// Solidity: event ExtraAssetAdded(address token)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseExtraAssetAdded(log types.Log) (*BeraBorrowSNECTExtraAssetAdded, error) {
	event := new(BeraBorrowSNECTExtraAssetAdded)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "ExtraAssetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTExtraAssetRemovedIterator is returned from FilterExtraAssetRemoved and is used to iterate over the raw logs and unpacked data for ExtraAssetRemoved events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTExtraAssetRemovedIterator struct {
	Event *BeraBorrowSNECTExtraAssetRemoved // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTExtraAssetRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTExtraAssetRemoved)
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
		it.Event = new(BeraBorrowSNECTExtraAssetRemoved)
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
func (it *BeraBorrowSNECTExtraAssetRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTExtraAssetRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTExtraAssetRemoved represents a ExtraAssetRemoved event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTExtraAssetRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExtraAssetRemoved is a free log retrieval operation binding the contract event 0xfc9138846a97b86614d19b78419b88e555c50bbd80b03feffd6264cd43064380.
//
// Solidity: event ExtraAssetRemoved(address token)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterExtraAssetRemoved(opts *bind.FilterOpts) (*BeraBorrowSNECTExtraAssetRemovedIterator, error) {

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "ExtraAssetRemoved")
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTExtraAssetRemovedIterator{contract: _BeraBorrowSNECT.contract, event: "ExtraAssetRemoved", logs: logs, sub: sub}, nil
}

// WatchExtraAssetRemoved is a free log subscription operation binding the contract event 0xfc9138846a97b86614d19b78419b88e555c50bbd80b03feffd6264cd43064380.
//
// Solidity: event ExtraAssetRemoved(address token)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchExtraAssetRemoved(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTExtraAssetRemoved) (event.Subscription, error) {

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "ExtraAssetRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTExtraAssetRemoved)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "ExtraAssetRemoved", log); err != nil {
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

// ParseExtraAssetRemoved is a log parse operation binding the contract event 0xfc9138846a97b86614d19b78419b88e555c50bbd80b03feffd6264cd43064380.
//
// Solidity: event ExtraAssetRemoved(address token)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseExtraAssetRemoved(log types.Log) (*BeraBorrowSNECTExtraAssetRemoved, error) {
	event := new(BeraBorrowSNECTExtraAssetRemoved)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "ExtraAssetRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTInitializedIterator struct {
	Event *BeraBorrowSNECTInitialized // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTInitialized)
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
		it.Event = new(BeraBorrowSNECTInitialized)
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
func (it *BeraBorrowSNECTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTInitialized represents a Initialized event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterInitialized(opts *bind.FilterOpts) (*BeraBorrowSNECTInitializedIterator, error) {

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTInitializedIterator{contract: _BeraBorrowSNECT.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTInitialized) (event.Subscription, error) {

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTInitialized)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseInitialized(log types.Log) (*BeraBorrowSNECTInitialized, error) {
	event := new(BeraBorrowSNECTInitialized)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTNewUnlockRatePerSecondIterator is returned from FilterNewUnlockRatePerSecond and is used to iterate over the raw logs and unpacked data for NewUnlockRatePerSecond events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTNewUnlockRatePerSecondIterator struct {
	Event *BeraBorrowSNECTNewUnlockRatePerSecond // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTNewUnlockRatePerSecondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTNewUnlockRatePerSecond)
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
		it.Event = new(BeraBorrowSNECTNewUnlockRatePerSecond)
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
func (it *BeraBorrowSNECTNewUnlockRatePerSecondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTNewUnlockRatePerSecondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTNewUnlockRatePerSecond represents a NewUnlockRatePerSecond event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTNewUnlockRatePerSecond struct {
	Token               common.Address
	UnlockRatePerSecond uint64
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewUnlockRatePerSecond is a free log retrieval operation binding the contract event 0x5577d4c8f6e5397effa5c71df8fe221e1162e18aaa0aabe87026cfb0c6762150.
//
// Solidity: event NewUnlockRatePerSecond(address indexed token, uint64 unlockRatePerSecond)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterNewUnlockRatePerSecond(opts *bind.FilterOpts, token []common.Address) (*BeraBorrowSNECTNewUnlockRatePerSecondIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "NewUnlockRatePerSecond", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTNewUnlockRatePerSecondIterator{contract: _BeraBorrowSNECT.contract, event: "NewUnlockRatePerSecond", logs: logs, sub: sub}, nil
}

// WatchNewUnlockRatePerSecond is a free log subscription operation binding the contract event 0x5577d4c8f6e5397effa5c71df8fe221e1162e18aaa0aabe87026cfb0c6762150.
//
// Solidity: event NewUnlockRatePerSecond(address indexed token, uint64 unlockRatePerSecond)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchNewUnlockRatePerSecond(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTNewUnlockRatePerSecond, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "NewUnlockRatePerSecond", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTNewUnlockRatePerSecond)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "NewUnlockRatePerSecond", log); err != nil {
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

// ParseNewUnlockRatePerSecond is a log parse operation binding the contract event 0x5577d4c8f6e5397effa5c71df8fe221e1162e18aaa0aabe87026cfb0c6762150.
//
// Solidity: event NewUnlockRatePerSecond(address indexed token, uint64 unlockRatePerSecond)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseNewUnlockRatePerSecond(log types.Log) (*BeraBorrowSNECTNewUnlockRatePerSecond, error) {
	event := new(BeraBorrowSNECTNewUnlockRatePerSecond)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "NewUnlockRatePerSecond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTOffsetIterator is returned from FilterOffset and is used to iterate over the raw logs and unpacked data for Offset events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTOffsetIterator struct {
	Event *BeraBorrowSNECTOffset // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTOffsetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTOffset)
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
		it.Event = new(BeraBorrowSNECTOffset)
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
func (it *BeraBorrowSNECTOffsetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTOffsetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTOffset represents a Offset event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTOffset struct {
	Collateral        common.Address
	DebtToOffset      *big.Int
	CollToAdd         *big.Int
	CollSurplusAmount *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOffset is a free log retrieval operation binding the contract event 0xc659bef2facfde65b659c8c5160cf21ac8232b38f8331aac0cea195e1d929665.
//
// Solidity: event Offset(address collateral, uint256 debtToOffset, uint256 collToAdd, uint256 collSurplusAmount)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterOffset(opts *bind.FilterOpts) (*BeraBorrowSNECTOffsetIterator, error) {

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "Offset")
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTOffsetIterator{contract: _BeraBorrowSNECT.contract, event: "Offset", logs: logs, sub: sub}, nil
}

// WatchOffset is a free log subscription operation binding the contract event 0xc659bef2facfde65b659c8c5160cf21ac8232b38f8331aac0cea195e1d929665.
//
// Solidity: event Offset(address collateral, uint256 debtToOffset, uint256 collToAdd, uint256 collSurplusAmount)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchOffset(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTOffset) (event.Subscription, error) {

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "Offset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTOffset)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Offset", log); err != nil {
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

// ParseOffset is a log parse operation binding the contract event 0xc659bef2facfde65b659c8c5160cf21ac8232b38f8331aac0cea195e1d929665.
//
// Solidity: event Offset(address collateral, uint256 debtToOffset, uint256 collToAdd, uint256 collSurplusAmount)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseOffset(log types.Log) (*BeraBorrowSNECTOffset, error) {
	event := new(BeraBorrowSNECTOffset)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Offset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTProtocolBlacklistedIterator is returned from FilterProtocolBlacklisted and is used to iterate over the raw logs and unpacked data for ProtocolBlacklisted events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTProtocolBlacklistedIterator struct {
	Event *BeraBorrowSNECTProtocolBlacklisted // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTProtocolBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTProtocolBlacklisted)
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
		it.Event = new(BeraBorrowSNECTProtocolBlacklisted)
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
func (it *BeraBorrowSNECTProtocolBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTProtocolBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTProtocolBlacklisted represents a ProtocolBlacklisted event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTProtocolBlacklisted struct {
	FactoryRemoved common.Address
	LMremoved      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProtocolBlacklisted is a free log retrieval operation binding the contract event 0x88ecee496061ddc38d88503f7cf6a1f4f6ade60ee216c946c5a0e8de6049595c.
//
// Solidity: event ProtocolBlacklisted(address indexed factoryRemoved, address indexed LMremoved)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterProtocolBlacklisted(opts *bind.FilterOpts, factoryRemoved []common.Address, LMremoved []common.Address) (*BeraBorrowSNECTProtocolBlacklistedIterator, error) {

	var factoryRemovedRule []interface{}
	for _, factoryRemovedItem := range factoryRemoved {
		factoryRemovedRule = append(factoryRemovedRule, factoryRemovedItem)
	}
	var LMremovedRule []interface{}
	for _, LMremovedItem := range LMremoved {
		LMremovedRule = append(LMremovedRule, LMremovedItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "ProtocolBlacklisted", factoryRemovedRule, LMremovedRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTProtocolBlacklistedIterator{contract: _BeraBorrowSNECT.contract, event: "ProtocolBlacklisted", logs: logs, sub: sub}, nil
}

// WatchProtocolBlacklisted is a free log subscription operation binding the contract event 0x88ecee496061ddc38d88503f7cf6a1f4f6ade60ee216c946c5a0e8de6049595c.
//
// Solidity: event ProtocolBlacklisted(address indexed factoryRemoved, address indexed LMremoved)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchProtocolBlacklisted(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTProtocolBlacklisted, factoryRemoved []common.Address, LMremoved []common.Address) (event.Subscription, error) {

	var factoryRemovedRule []interface{}
	for _, factoryRemovedItem := range factoryRemoved {
		factoryRemovedRule = append(factoryRemovedRule, factoryRemovedItem)
	}
	var LMremovedRule []interface{}
	for _, LMremovedItem := range LMremoved {
		LMremovedRule = append(LMremovedRule, LMremovedItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "ProtocolBlacklisted", factoryRemovedRule, LMremovedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTProtocolBlacklisted)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "ProtocolBlacklisted", log); err != nil {
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

// ParseProtocolBlacklisted is a log parse operation binding the contract event 0x88ecee496061ddc38d88503f7cf6a1f4f6ade60ee216c946c5a0e8de6049595c.
//
// Solidity: event ProtocolBlacklisted(address indexed factoryRemoved, address indexed LMremoved)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseProtocolBlacklisted(log types.Log) (*BeraBorrowSNECTProtocolBlacklisted, error) {
	event := new(BeraBorrowSNECTProtocolBlacklisted)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "ProtocolBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTProtocolRegisteredIterator is returned from FilterProtocolRegistered and is used to iterate over the raw logs and unpacked data for ProtocolRegistered events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTProtocolRegisteredIterator struct {
	Event *BeraBorrowSNECTProtocolRegistered // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTProtocolRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTProtocolRegistered)
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
		it.Event = new(BeraBorrowSNECTProtocolRegistered)
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
func (it *BeraBorrowSNECTProtocolRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTProtocolRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTProtocolRegistered represents a ProtocolRegistered event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTProtocolRegistered struct {
	Factory            common.Address
	LiquidationManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterProtocolRegistered is a free log retrieval operation binding the contract event 0xd6c91941062a66dc4c4344f6b10af4b565b256816a2f9080ba7f83e1d6a2bdc6.
//
// Solidity: event ProtocolRegistered(address indexed factory, address indexed liquidationManager)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterProtocolRegistered(opts *bind.FilterOpts, factory []common.Address, liquidationManager []common.Address) (*BeraBorrowSNECTProtocolRegisteredIterator, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}
	var liquidationManagerRule []interface{}
	for _, liquidationManagerItem := range liquidationManager {
		liquidationManagerRule = append(liquidationManagerRule, liquidationManagerItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "ProtocolRegistered", factoryRule, liquidationManagerRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTProtocolRegisteredIterator{contract: _BeraBorrowSNECT.contract, event: "ProtocolRegistered", logs: logs, sub: sub}, nil
}

// WatchProtocolRegistered is a free log subscription operation binding the contract event 0xd6c91941062a66dc4c4344f6b10af4b565b256816a2f9080ba7f83e1d6a2bdc6.
//
// Solidity: event ProtocolRegistered(address indexed factory, address indexed liquidationManager)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchProtocolRegistered(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTProtocolRegistered, factory []common.Address, liquidationManager []common.Address) (event.Subscription, error) {

	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}
	var liquidationManagerRule []interface{}
	for _, liquidationManagerItem := range liquidationManager {
		liquidationManagerRule = append(liquidationManagerRule, liquidationManagerItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "ProtocolRegistered", factoryRule, liquidationManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTProtocolRegistered)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "ProtocolRegistered", log); err != nil {
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

// ParseProtocolRegistered is a log parse operation binding the contract event 0xd6c91941062a66dc4c4344f6b10af4b565b256816a2f9080ba7f83e1d6a2bdc6.
//
// Solidity: event ProtocolRegistered(address indexed factory, address indexed liquidationManager)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseProtocolRegistered(log types.Log) (*BeraBorrowSNECTProtocolRegistered, error) {
	event := new(BeraBorrowSNECTProtocolRegistered)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "ProtocolRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTRebalanceIterator struct {
	Event *BeraBorrowSNECTRebalance // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTRebalance)
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
		it.Event = new(BeraBorrowSNECTRebalance)
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
func (it *BeraBorrowSNECTRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTRebalance represents a Rebalance event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTRebalance struct {
	SentCurrency     common.Address
	ReceivedCurrency common.Address
	SentAmount       *big.Int
	ReceivedAmount   *big.Int
	SentValue        *big.Int
	ReceivedValue    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRebalance is a free log retrieval operation binding the contract event 0xb8c3fd52c06cd7e35d81a3fc31542187d197c9deef253587a27e0214677d0f6b.
//
// Solidity: event Rebalance(address indexed sentCurrency, address indexed receivedCurrency, uint256 sentAmount, uint256 receivedAmount, uint256 sentValue, uint256 receivedValue)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterRebalance(opts *bind.FilterOpts, sentCurrency []common.Address, receivedCurrency []common.Address) (*BeraBorrowSNECTRebalanceIterator, error) {

	var sentCurrencyRule []interface{}
	for _, sentCurrencyItem := range sentCurrency {
		sentCurrencyRule = append(sentCurrencyRule, sentCurrencyItem)
	}
	var receivedCurrencyRule []interface{}
	for _, receivedCurrencyItem := range receivedCurrency {
		receivedCurrencyRule = append(receivedCurrencyRule, receivedCurrencyItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "Rebalance", sentCurrencyRule, receivedCurrencyRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTRebalanceIterator{contract: _BeraBorrowSNECT.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0xb8c3fd52c06cd7e35d81a3fc31542187d197c9deef253587a27e0214677d0f6b.
//
// Solidity: event Rebalance(address indexed sentCurrency, address indexed receivedCurrency, uint256 sentAmount, uint256 receivedAmount, uint256 sentValue, uint256 receivedValue)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTRebalance, sentCurrency []common.Address, receivedCurrency []common.Address) (event.Subscription, error) {

	var sentCurrencyRule []interface{}
	for _, sentCurrencyItem := range sentCurrency {
		sentCurrencyRule = append(sentCurrencyRule, sentCurrencyItem)
	}
	var receivedCurrencyRule []interface{}
	for _, receivedCurrencyItem := range receivedCurrency {
		receivedCurrencyRule = append(receivedCurrencyRule, receivedCurrencyItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "Rebalance", sentCurrencyRule, receivedCurrencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTRebalance)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Rebalance", log); err != nil {
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

// ParseRebalance is a log parse operation binding the contract event 0xb8c3fd52c06cd7e35d81a3fc31542187d197c9deef253587a27e0214677d0f6b.
//
// Solidity: event Rebalance(address indexed sentCurrency, address indexed receivedCurrency, uint256 sentAmount, uint256 receivedAmount, uint256 sentValue, uint256 receivedValue)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseRebalance(log types.Log) (*BeraBorrowSNECTRebalance, error) {
	event := new(BeraBorrowSNECTRebalance)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTTransferIterator struct {
	Event *BeraBorrowSNECTTransfer // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTTransfer)
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
		it.Event = new(BeraBorrowSNECTTransfer)
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
func (it *BeraBorrowSNECTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTTransfer represents a Transfer event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BeraBorrowSNECTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTTransferIterator{contract: _BeraBorrowSNECT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTTransfer)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseTransfer(log types.Log) (*BeraBorrowSNECTTransfer, error) {
	event := new(BeraBorrowSNECTTransfer)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTUpgradedIterator struct {
	Event *BeraBorrowSNECTUpgraded // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTUpgraded)
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
		it.Event = new(BeraBorrowSNECTUpgraded)
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
func (it *BeraBorrowSNECTUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTUpgraded represents a Upgraded event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BeraBorrowSNECTUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTUpgradedIterator{contract: _BeraBorrowSNECT.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTUpgraded)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseUpgraded(log types.Log) (*BeraBorrowSNECTUpgraded, error) {
	event := new(BeraBorrowSNECTUpgraded)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowSNECTWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTWithdrawIterator struct {
	Event *BeraBorrowSNECTWithdraw // Event containing the contract specifics and raw log

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
func (it *BeraBorrowSNECTWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowSNECTWithdraw)
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
		it.Event = new(BeraBorrowSNECTWithdraw)
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
func (it *BeraBorrowSNECTWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowSNECTWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowSNECTWithdraw represents a Withdraw event raised by the BeraBorrowSNECT contract.
type BeraBorrowSNECTWithdraw struct {
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
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*BeraBorrowSNECTWithdrawIterator, error) {

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

	logs, sub, err := _BeraBorrowSNECT.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowSNECTWithdrawIterator{contract: _BeraBorrowSNECT.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BeraBorrowSNECTWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BeraBorrowSNECT.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowSNECTWithdraw)
				if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_BeraBorrowSNECT *BeraBorrowSNECTFilterer) ParseWithdraw(log types.Log) (*BeraBorrowSNECTWithdraw, error) {
	event := new(BeraBorrowSNECTWithdraw)
	if err := _BeraBorrowSNECT.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
