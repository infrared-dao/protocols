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

// IBaseCollateralVaultBaseInitParams is an auto generated low-level Go binding around an user-defined struct.
type IBaseCollateralVaultBaseInitParams struct {
	MinWithdrawFee     uint16
	MaxWithdrawFee     uint16
	WithdrawFee        uint16
	MetaBeraborrowCore common.Address
	Asset              common.Address
	SharesName         string
	SharesSymbol       string
}

// IInfraredCollateralVaultInfraredInitParams is an auto generated low-level Go binding around an user-defined struct.
type IInfraredCollateralVaultInfraredInitParams struct {
	BaseParams        IBaseCollateralVaultBaseInitParams
	MinPerformanceFee uint16
	MaxPerformanceFee uint16
	PerformanceFee    uint16
	IRedToken         common.Address
	InfraredVault     common.Address
	IbgtVault         common.Address
	InfraredWrapper   common.Address
}

// IInfraredCollateralVaultRebalanceParams is an auto generated low-level Go binding around an user-defined struct.
type IInfraredCollateralVaultRebalanceParams struct {
	SentCurrency common.Address
	SentAmount   *big.Int
	Swapper      common.Address
	Payload      []byte
}

// BeraBorrowCICVMetaData contains all meta data concerning the BeraBorrowCICV contract.
var BeraBorrowCICVMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AmountCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmissionRateExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawingLockedEmissions\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"EmissionsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"unlockRatePerSecond\",\"type\":\"uint64\"}],\"name\":\"NewUnlockRatePerSecond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PerformanceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sentCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sentValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedValue\",\"type\":\"uint256\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"_convertToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalanceOfWithFutureEmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFullProfitUnlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLockedEmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMetaBeraborrowCore\",\"outputs\":[{\"internalType\":\"contractIMetaBeraborrowCore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPerformanceFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceFeed\",\"outputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"iRedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibgt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibgtVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"infraredVault\",\"outputs\":[{\"internalType\":\"contractIInfraredVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"_minWithdrawFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_maxWithdrawFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_withdrawFee\",\"type\":\"uint16\"},{\"internalType\":\"contractIMetaBeraborrowCore\",\"name\":\"_metaBeraborrowCore\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_sharesName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sharesSymbol\",\"type\":\"string\"}],\"internalType\":\"structIBaseCollateralVault.BaseInitParams\",\"name\":\"_baseParams\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"_minPerformanceFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_maxPerformanceFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_performanceFee\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_iRedToken\",\"type\":\"address\"},{\"internalType\":\"contractIInfraredVault\",\"name\":\"_infraredVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ibgtVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_infraredWrapper\",\"type\":\"address\"}],\"internalType\":\"structIInfraredCollateralVault.InfraredInitParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"}],\"name\":\"internalizeDonations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sentCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sentAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structIInfraredCollateralVault.RebalanceParams\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"receiveDonations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iRedToken\",\"type\":\"address\"}],\"name\":\"setIRED\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"thresholdInBP\",\"type\":\"uint256\"}],\"name\":\"setPairThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_performanceFee\",\"type\":\"uint16\"}],\"name\":\"setPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_unlockRatePerSecond\",\"type\":\"uint64\"}],\"name\":\"setUnlockRatePerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_withdrawFee\",\"type\":\"uint16\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInAsset\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unlockRatePerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BeraBorrowCICVABI is the input ABI used to generate the binding from.
// Deprecated: Use BeraBorrowCICVMetaData.ABI instead.
var BeraBorrowCICVABI = BeraBorrowCICVMetaData.ABI

// BeraBorrowCICV is an auto generated Go binding around an Ethereum contract.
type BeraBorrowCICV struct {
	BeraBorrowCICVCaller     // Read-only binding to the contract
	BeraBorrowCICVTransactor // Write-only binding to the contract
	BeraBorrowCICVFilterer   // Log filterer for contract events
}

// BeraBorrowCICVCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeraBorrowCICVCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowCICVTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeraBorrowCICVTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowCICVFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeraBorrowCICVFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowCICVSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeraBorrowCICVSession struct {
	Contract     *BeraBorrowCICV   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeraBorrowCICVCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeraBorrowCICVCallerSession struct {
	Contract *BeraBorrowCICVCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BeraBorrowCICVTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeraBorrowCICVTransactorSession struct {
	Contract     *BeraBorrowCICVTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BeraBorrowCICVRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeraBorrowCICVRaw struct {
	Contract *BeraBorrowCICV // Generic contract binding to access the raw methods on
}

// BeraBorrowCICVCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeraBorrowCICVCallerRaw struct {
	Contract *BeraBorrowCICVCaller // Generic read-only contract binding to access the raw methods on
}

// BeraBorrowCICVTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeraBorrowCICVTransactorRaw struct {
	Contract *BeraBorrowCICVTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeraBorrowCICV creates a new instance of BeraBorrowCICV, bound to a specific deployed contract.
func NewBeraBorrowCICV(address common.Address, backend bind.ContractBackend) (*BeraBorrowCICV, error) {
	contract, err := bindBeraBorrowCICV(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICV{BeraBorrowCICVCaller: BeraBorrowCICVCaller{contract: contract}, BeraBorrowCICVTransactor: BeraBorrowCICVTransactor{contract: contract}, BeraBorrowCICVFilterer: BeraBorrowCICVFilterer{contract: contract}}, nil
}

// NewBeraBorrowCICVCaller creates a new read-only instance of BeraBorrowCICV, bound to a specific deployed contract.
func NewBeraBorrowCICVCaller(address common.Address, caller bind.ContractCaller) (*BeraBorrowCICVCaller, error) {
	contract, err := bindBeraBorrowCICV(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVCaller{contract: contract}, nil
}

// NewBeraBorrowCICVTransactor creates a new write-only instance of BeraBorrowCICV, bound to a specific deployed contract.
func NewBeraBorrowCICVTransactor(address common.Address, transactor bind.ContractTransactor) (*BeraBorrowCICVTransactor, error) {
	contract, err := bindBeraBorrowCICV(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVTransactor{contract: contract}, nil
}

// NewBeraBorrowCICVFilterer creates a new log filterer instance of BeraBorrowCICV, bound to a specific deployed contract.
func NewBeraBorrowCICVFilterer(address common.Address, filterer bind.ContractFilterer) (*BeraBorrowCICVFilterer, error) {
	contract, err := bindBeraBorrowCICV(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVFilterer{contract: contract}, nil
}

// bindBeraBorrowCICV binds a generic wrapper to an already deployed contract.
func bindBeraBorrowCICV(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeraBorrowCICVMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeraBorrowCICV *BeraBorrowCICVRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeraBorrowCICV.Contract.BeraBorrowCICVCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeraBorrowCICV *BeraBorrowCICVRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.BeraBorrowCICVTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeraBorrowCICV *BeraBorrowCICVRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.BeraBorrowCICVTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeraBorrowCICV *BeraBorrowCICVCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeraBorrowCICV.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeraBorrowCICV *BeraBorrowCICVTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeraBorrowCICV *BeraBorrowCICVTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BeraBorrowCICV *BeraBorrowCICVSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BeraBorrowCICV.Contract.UPGRADEINTERFACEVERSION(&_BeraBorrowCICV.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BeraBorrowCICV.Contract.UPGRADEINTERFACEVERSION(&_BeraBorrowCICV.CallOpts)
}

// ConvertToValue is a free data retrieval call binding the contract method 0xec05c341.
//
// Solidity: function _convertToValue(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) ConvertToValue(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "_convertToValue", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToValue is a free data retrieval call binding the contract method 0xec05c341.
//
// Solidity: function _convertToValue(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) ConvertToValue(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.ConvertToValue(&_BeraBorrowCICV.CallOpts, token)
}

// ConvertToValue is a free data retrieval call binding the contract method 0xec05c341.
//
// Solidity: function _convertToValue(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) ConvertToValue(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.ConvertToValue(&_BeraBorrowCICV.CallOpts, token)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.Allowance(&_BeraBorrowCICV.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.Allowance(&_BeraBorrowCICV.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Asset() (common.Address, error) {
	return _BeraBorrowCICV.Contract.Asset(&_BeraBorrowCICV.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) Asset() (common.Address, error) {
	return _BeraBorrowCICV.Contract.Asset(&_BeraBorrowCICV.CallOpts)
}

// AssetDecimals is a free data retrieval call binding the contract method 0xc2d41601.
//
// Solidity: function assetDecimals() view returns(uint8)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) AssetDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "assetDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AssetDecimals is a free data retrieval call binding the contract method 0xc2d41601.
//
// Solidity: function assetDecimals() view returns(uint8)
func (_BeraBorrowCICV *BeraBorrowCICVSession) AssetDecimals() (uint8, error) {
	return _BeraBorrowCICV.Contract.AssetDecimals(&_BeraBorrowCICV.CallOpts)
}

// AssetDecimals is a free data retrieval call binding the contract method 0xc2d41601.
//
// Solidity: function assetDecimals() view returns(uint8)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) AssetDecimals() (uint8, error) {
	return _BeraBorrowCICV.Contract.AssetDecimals(&_BeraBorrowCICV.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.BalanceOf(&_BeraBorrowCICV.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.BalanceOf(&_BeraBorrowCICV.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.ConvertToAssets(&_BeraBorrowCICV.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.ConvertToAssets(&_BeraBorrowCICV.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.ConvertToShares(&_BeraBorrowCICV.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.ConvertToShares(&_BeraBorrowCICV.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Decimals() (uint8, error) {
	return _BeraBorrowCICV.Contract.Decimals(&_BeraBorrowCICV.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) Decimals() (uint8, error) {
	return _BeraBorrowCICV.Contract.Decimals(&_BeraBorrowCICV.CallOpts)
}

// FetchPrice is a free data retrieval call binding the contract method 0x0fdb11cf.
//
// Solidity: function fetchPrice() view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) FetchPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "fetchPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchPrice is a free data retrieval call binding the contract method 0x0fdb11cf.
//
// Solidity: function fetchPrice() view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) FetchPrice() (*big.Int, error) {
	return _BeraBorrowCICV.Contract.FetchPrice(&_BeraBorrowCICV.CallOpts)
}

// FetchPrice is a free data retrieval call binding the contract method 0x0fdb11cf.
//
// Solidity: function fetchPrice() view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) FetchPrice() (*big.Int, error) {
	return _BeraBorrowCICV.Contract.FetchPrice(&_BeraBorrowCICV.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.GetBalance(&_BeraBorrowCICV.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.GetBalance(&_BeraBorrowCICV.CallOpts, token)
}

// GetBalanceOfWithFutureEmissions is a free data retrieval call binding the contract method 0x29d78c88.
//
// Solidity: function getBalanceOfWithFutureEmissions(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) GetBalanceOfWithFutureEmissions(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "getBalanceOfWithFutureEmissions", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOfWithFutureEmissions is a free data retrieval call binding the contract method 0x29d78c88.
//
// Solidity: function getBalanceOfWithFutureEmissions(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) GetBalanceOfWithFutureEmissions(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.GetBalanceOfWithFutureEmissions(&_BeraBorrowCICV.CallOpts, token)
}

// GetBalanceOfWithFutureEmissions is a free data retrieval call binding the contract method 0x29d78c88.
//
// Solidity: function getBalanceOfWithFutureEmissions(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) GetBalanceOfWithFutureEmissions(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.GetBalanceOfWithFutureEmissions(&_BeraBorrowCICV.CallOpts, token)
}

// GetFullProfitUnlockTimestamp is a free data retrieval call binding the contract method 0xd5c16d79.
//
// Solidity: function getFullProfitUnlockTimestamp(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) GetFullProfitUnlockTimestamp(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "getFullProfitUnlockTimestamp", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFullProfitUnlockTimestamp is a free data retrieval call binding the contract method 0xd5c16d79.
//
// Solidity: function getFullProfitUnlockTimestamp(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) GetFullProfitUnlockTimestamp(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.GetFullProfitUnlockTimestamp(&_BeraBorrowCICV.CallOpts, token)
}

// GetFullProfitUnlockTimestamp is a free data retrieval call binding the contract method 0xd5c16d79.
//
// Solidity: function getFullProfitUnlockTimestamp(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) GetFullProfitUnlockTimestamp(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.GetFullProfitUnlockTimestamp(&_BeraBorrowCICV.CallOpts, token)
}

// GetLockedEmissions is a free data retrieval call binding the contract method 0x403dd3bc.
//
// Solidity: function getLockedEmissions(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) GetLockedEmissions(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "getLockedEmissions", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedEmissions is a free data retrieval call binding the contract method 0x403dd3bc.
//
// Solidity: function getLockedEmissions(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) GetLockedEmissions(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.GetLockedEmissions(&_BeraBorrowCICV.CallOpts, token)
}

// GetLockedEmissions is a free data retrieval call binding the contract method 0x403dd3bc.
//
// Solidity: function getLockedEmissions(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) GetLockedEmissions(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.GetLockedEmissions(&_BeraBorrowCICV.CallOpts, token)
}

// GetMetaBeraborrowCore is a free data retrieval call binding the contract method 0x221085d7.
//
// Solidity: function getMetaBeraborrowCore() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) GetMetaBeraborrowCore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "getMetaBeraborrowCore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMetaBeraborrowCore is a free data retrieval call binding the contract method 0x221085d7.
//
// Solidity: function getMetaBeraborrowCore() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVSession) GetMetaBeraborrowCore() (common.Address, error) {
	return _BeraBorrowCICV.Contract.GetMetaBeraborrowCore(&_BeraBorrowCICV.CallOpts)
}

// GetMetaBeraborrowCore is a free data retrieval call binding the contract method 0x221085d7.
//
// Solidity: function getMetaBeraborrowCore() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) GetMetaBeraborrowCore() (common.Address, error) {
	return _BeraBorrowCICV.Contract.GetMetaBeraborrowCore(&_BeraBorrowCICV.CallOpts)
}

// GetPerformanceFee is a free data retrieval call binding the contract method 0x235c3603.
//
// Solidity: function getPerformanceFee() view returns(uint16)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) GetPerformanceFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "getPerformanceFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetPerformanceFee is a free data retrieval call binding the contract method 0x235c3603.
//
// Solidity: function getPerformanceFee() view returns(uint16)
func (_BeraBorrowCICV *BeraBorrowCICVSession) GetPerformanceFee() (uint16, error) {
	return _BeraBorrowCICV.Contract.GetPerformanceFee(&_BeraBorrowCICV.CallOpts)
}

// GetPerformanceFee is a free data retrieval call binding the contract method 0x235c3603.
//
// Solidity: function getPerformanceFee() view returns(uint16)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) GetPerformanceFee() (uint16, error) {
	return _BeraBorrowCICV.Contract.GetPerformanceFee(&_BeraBorrowCICV.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) GetPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "getPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) GetPrice(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.GetPrice(&_BeraBorrowCICV.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) GetPrice(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.GetPrice(&_BeraBorrowCICV.CallOpts, token)
}

// GetPriceFeed is a free data retrieval call binding the contract method 0x9e87a5cd.
//
// Solidity: function getPriceFeed() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) GetPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "getPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceFeed is a free data retrieval call binding the contract method 0x9e87a5cd.
//
// Solidity: function getPriceFeed() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVSession) GetPriceFeed() (common.Address, error) {
	return _BeraBorrowCICV.Contract.GetPriceFeed(&_BeraBorrowCICV.CallOpts)
}

// GetPriceFeed is a free data retrieval call binding the contract method 0x9e87a5cd.
//
// Solidity: function getPriceFeed() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) GetPriceFeed() (common.Address, error) {
	return _BeraBorrowCICV.Contract.GetPriceFeed(&_BeraBorrowCICV.CallOpts)
}

// GetWithdrawFee is a free data retrieval call binding the contract method 0x1540aa89.
//
// Solidity: function getWithdrawFee() view returns(uint16)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) GetWithdrawFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "getWithdrawFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetWithdrawFee is a free data retrieval call binding the contract method 0x1540aa89.
//
// Solidity: function getWithdrawFee() view returns(uint16)
func (_BeraBorrowCICV *BeraBorrowCICVSession) GetWithdrawFee() (uint16, error) {
	return _BeraBorrowCICV.Contract.GetWithdrawFee(&_BeraBorrowCICV.CallOpts)
}

// GetWithdrawFee is a free data retrieval call binding the contract method 0x1540aa89.
//
// Solidity: function getWithdrawFee() view returns(uint16)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) GetWithdrawFee() (uint16, error) {
	return _BeraBorrowCICV.Contract.GetWithdrawFee(&_BeraBorrowCICV.CallOpts)
}

// IRedToken is a free data retrieval call binding the contract method 0xe993e7ed.
//
// Solidity: function iRedToken() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) IRedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "iRedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IRedToken is a free data retrieval call binding the contract method 0xe993e7ed.
//
// Solidity: function iRedToken() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVSession) IRedToken() (common.Address, error) {
	return _BeraBorrowCICV.Contract.IRedToken(&_BeraBorrowCICV.CallOpts)
}

// IRedToken is a free data retrieval call binding the contract method 0xe993e7ed.
//
// Solidity: function iRedToken() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) IRedToken() (common.Address, error) {
	return _BeraBorrowCICV.Contract.IRedToken(&_BeraBorrowCICV.CallOpts)
}

// Ibgt is a free data retrieval call binding the contract method 0x3dafa4f3.
//
// Solidity: function ibgt() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) Ibgt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "ibgt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ibgt is a free data retrieval call binding the contract method 0x3dafa4f3.
//
// Solidity: function ibgt() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Ibgt() (common.Address, error) {
	return _BeraBorrowCICV.Contract.Ibgt(&_BeraBorrowCICV.CallOpts)
}

// Ibgt is a free data retrieval call binding the contract method 0x3dafa4f3.
//
// Solidity: function ibgt() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) Ibgt() (common.Address, error) {
	return _BeraBorrowCICV.Contract.Ibgt(&_BeraBorrowCICV.CallOpts)
}

// IbgtVault is a free data retrieval call binding the contract method 0xfd64c377.
//
// Solidity: function ibgtVault() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) IbgtVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "ibgtVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IbgtVault is a free data retrieval call binding the contract method 0xfd64c377.
//
// Solidity: function ibgtVault() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVSession) IbgtVault() (common.Address, error) {
	return _BeraBorrowCICV.Contract.IbgtVault(&_BeraBorrowCICV.CallOpts)
}

// IbgtVault is a free data retrieval call binding the contract method 0xfd64c377.
//
// Solidity: function ibgtVault() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) IbgtVault() (common.Address, error) {
	return _BeraBorrowCICV.Contract.IbgtVault(&_BeraBorrowCICV.CallOpts)
}

// InfraredVault is a free data retrieval call binding the contract method 0x6ab3b0e5.
//
// Solidity: function infraredVault() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) InfraredVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "infraredVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InfraredVault is a free data retrieval call binding the contract method 0x6ab3b0e5.
//
// Solidity: function infraredVault() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVSession) InfraredVault() (common.Address, error) {
	return _BeraBorrowCICV.Contract.InfraredVault(&_BeraBorrowCICV.CallOpts)
}

// InfraredVault is a free data retrieval call binding the contract method 0x6ab3b0e5.
//
// Solidity: function infraredVault() view returns(address)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) InfraredVault() (common.Address, error) {
	return _BeraBorrowCICV.Contract.InfraredVault(&_BeraBorrowCICV.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.MaxDeposit(&_BeraBorrowCICV.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.MaxDeposit(&_BeraBorrowCICV.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.MaxMint(&_BeraBorrowCICV.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.MaxMint(&_BeraBorrowCICV.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.MaxRedeem(&_BeraBorrowCICV.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.MaxRedeem(&_BeraBorrowCICV.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.MaxWithdraw(&_BeraBorrowCICV.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.MaxWithdraw(&_BeraBorrowCICV.CallOpts, _owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Name() (string, error) {
	return _BeraBorrowCICV.Contract.Name(&_BeraBorrowCICV.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) Name() (string, error) {
	return _BeraBorrowCICV.Contract.Name(&_BeraBorrowCICV.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.PreviewDeposit(&_BeraBorrowCICV.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.PreviewDeposit(&_BeraBorrowCICV.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.PreviewMint(&_BeraBorrowCICV.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.PreviewMint(&_BeraBorrowCICV.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.PreviewRedeem(&_BeraBorrowCICV.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.PreviewRedeem(&_BeraBorrowCICV.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.PreviewWithdraw(&_BeraBorrowCICV.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.PreviewWithdraw(&_BeraBorrowCICV.CallOpts, assets)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BeraBorrowCICV *BeraBorrowCICVSession) ProxiableUUID() ([32]byte, error) {
	return _BeraBorrowCICV.Contract.ProxiableUUID(&_BeraBorrowCICV.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BeraBorrowCICV.Contract.ProxiableUUID(&_BeraBorrowCICV.CallOpts)
}

// RewardedTokens is a free data retrieval call binding the contract method 0xaaffe076.
//
// Solidity: function rewardedTokens() view returns(address[])
func (_BeraBorrowCICV *BeraBorrowCICVCaller) RewardedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "rewardedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RewardedTokens is a free data retrieval call binding the contract method 0xaaffe076.
//
// Solidity: function rewardedTokens() view returns(address[])
func (_BeraBorrowCICV *BeraBorrowCICVSession) RewardedTokens() ([]common.Address, error) {
	return _BeraBorrowCICV.Contract.RewardedTokens(&_BeraBorrowCICV.CallOpts)
}

// RewardedTokens is a free data retrieval call binding the contract method 0xaaffe076.
//
// Solidity: function rewardedTokens() view returns(address[])
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) RewardedTokens() ([]common.Address, error) {
	return _BeraBorrowCICV.Contract.RewardedTokens(&_BeraBorrowCICV.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Symbol() (string, error) {
	return _BeraBorrowCICV.Contract.Symbol(&_BeraBorrowCICV.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) Symbol() (string, error) {
	return _BeraBorrowCICV.Contract.Symbol(&_BeraBorrowCICV.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 amountInAsset)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 amountInAsset)
func (_BeraBorrowCICV *BeraBorrowCICVSession) TotalAssets() (*big.Int, error) {
	return _BeraBorrowCICV.Contract.TotalAssets(&_BeraBorrowCICV.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 amountInAsset)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) TotalAssets() (*big.Int, error) {
	return _BeraBorrowCICV.Contract.TotalAssets(&_BeraBorrowCICV.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) TotalSupply() (*big.Int, error) {
	return _BeraBorrowCICV.Contract.TotalSupply(&_BeraBorrowCICV.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) TotalSupply() (*big.Int, error) {
	return _BeraBorrowCICV.Contract.TotalSupply(&_BeraBorrowCICV.CallOpts)
}

// UnlockRatePerSecond is a free data retrieval call binding the contract method 0xd9c00197.
//
// Solidity: function unlockRatePerSecond(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCaller) UnlockRatePerSecond(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCICV.contract.Call(opts, &out, "unlockRatePerSecond", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockRatePerSecond is a free data retrieval call binding the contract method 0xd9c00197.
//
// Solidity: function unlockRatePerSecond(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVSession) UnlockRatePerSecond(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.UnlockRatePerSecond(&_BeraBorrowCICV.CallOpts, token)
}

// UnlockRatePerSecond is a free data retrieval call binding the contract method 0xd9c00197.
//
// Solidity: function unlockRatePerSecond(address token) view returns(uint256)
func (_BeraBorrowCICV *BeraBorrowCICVCallerSession) UnlockRatePerSecond(token common.Address) (*big.Int, error) {
	return _BeraBorrowCICV.Contract.UnlockRatePerSecond(&_BeraBorrowCICV.CallOpts, token)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Approve(&_BeraBorrowCICV.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Approve(&_BeraBorrowCICV.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Deposit(&_BeraBorrowCICV.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Deposit(&_BeraBorrowCICV.TransactOpts, assets, receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xd67524a9.
//
// Solidity: function initialize(((uint16,uint16,uint16,address,address,string,string),uint16,uint16,uint16,address,address,address,address) baseParams) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) Initialize(opts *bind.TransactOpts, baseParams IInfraredCollateralVaultInfraredInitParams) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "initialize", baseParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xd67524a9.
//
// Solidity: function initialize(((uint16,uint16,uint16,address,address,string,string),uint16,uint16,uint16,address,address,address,address) baseParams) returns()
func (_BeraBorrowCICV *BeraBorrowCICVSession) Initialize(baseParams IInfraredCollateralVaultInfraredInitParams) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Initialize(&_BeraBorrowCICV.TransactOpts, baseParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xd67524a9.
//
// Solidity: function initialize(((uint16,uint16,uint16,address,address,string,string),uint16,uint16,uint16,address,address,address,address) baseParams) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) Initialize(baseParams IInfraredCollateralVaultInfraredInitParams) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Initialize(&_BeraBorrowCICV.TransactOpts, baseParams)
}

// InternalizeDonations is a paid mutator transaction binding the contract method 0x3e92f95b.
//
// Solidity: function internalizeDonations(address[] tokens, uint128[] amounts) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) InternalizeDonations(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "internalizeDonations", tokens, amounts)
}

// InternalizeDonations is a paid mutator transaction binding the contract method 0x3e92f95b.
//
// Solidity: function internalizeDonations(address[] tokens, uint128[] amounts) returns()
func (_BeraBorrowCICV *BeraBorrowCICVSession) InternalizeDonations(tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.InternalizeDonations(&_BeraBorrowCICV.TransactOpts, tokens, amounts)
}

// InternalizeDonations is a paid mutator transaction binding the contract method 0x3e92f95b.
//
// Solidity: function internalizeDonations(address[] tokens, uint128[] amounts) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) InternalizeDonations(tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.InternalizeDonations(&_BeraBorrowCICV.TransactOpts, tokens, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Mint(&_BeraBorrowCICV.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Mint(&_BeraBorrowCICV.TransactOpts, shares, receiver)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb2c06ba0.
//
// Solidity: function rebalance((address,uint256,address,bytes) p) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) Rebalance(opts *bind.TransactOpts, p IInfraredCollateralVaultRebalanceParams) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "rebalance", p)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb2c06ba0.
//
// Solidity: function rebalance((address,uint256,address,bytes) p) returns()
func (_BeraBorrowCICV *BeraBorrowCICVSession) Rebalance(p IInfraredCollateralVaultRebalanceParams) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Rebalance(&_BeraBorrowCICV.TransactOpts, p)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb2c06ba0.
//
// Solidity: function rebalance((address,uint256,address,bytes) p) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) Rebalance(p IInfraredCollateralVaultRebalanceParams) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Rebalance(&_BeraBorrowCICV.TransactOpts, p)
}

// ReceiveDonations is a paid mutator transaction binding the contract method 0x32e1fef2.
//
// Solidity: function receiveDonations(address[] tokens, uint256[] amounts, address receiver) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) ReceiveDonations(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "receiveDonations", tokens, amounts, receiver)
}

// ReceiveDonations is a paid mutator transaction binding the contract method 0x32e1fef2.
//
// Solidity: function receiveDonations(address[] tokens, uint256[] amounts, address receiver) returns()
func (_BeraBorrowCICV *BeraBorrowCICVSession) ReceiveDonations(tokens []common.Address, amounts []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.ReceiveDonations(&_BeraBorrowCICV.TransactOpts, tokens, amounts, receiver)
}

// ReceiveDonations is a paid mutator transaction binding the contract method 0x32e1fef2.
//
// Solidity: function receiveDonations(address[] tokens, uint256[] amounts, address receiver) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) ReceiveDonations(tokens []common.Address, amounts []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.ReceiveDonations(&_BeraBorrowCICV.TransactOpts, tokens, amounts, receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "redeem", shares, receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Redeem(shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Redeem(&_BeraBorrowCICV.TransactOpts, shares, receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) Redeem(shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Redeem(&_BeraBorrowCICV.TransactOpts, shares, receiver, _owner)
}

// SetIRED is a paid mutator transaction binding the contract method 0x41519a65.
//
// Solidity: function setIRED(address _iRedToken) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) SetIRED(opts *bind.TransactOpts, _iRedToken common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "setIRED", _iRedToken)
}

// SetIRED is a paid mutator transaction binding the contract method 0x41519a65.
//
// Solidity: function setIRED(address _iRedToken) returns()
func (_BeraBorrowCICV *BeraBorrowCICVSession) SetIRED(_iRedToken common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.SetIRED(&_BeraBorrowCICV.TransactOpts, _iRedToken)
}

// SetIRED is a paid mutator transaction binding the contract method 0x41519a65.
//
// Solidity: function setIRED(address _iRedToken) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) SetIRED(_iRedToken common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.SetIRED(&_BeraBorrowCICV.TransactOpts, _iRedToken)
}

// SetPairThreshold is a paid mutator transaction binding the contract method 0x9a400b0f.
//
// Solidity: function setPairThreshold(address tokenIn, uint256 thresholdInBP) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) SetPairThreshold(opts *bind.TransactOpts, tokenIn common.Address, thresholdInBP *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "setPairThreshold", tokenIn, thresholdInBP)
}

// SetPairThreshold is a paid mutator transaction binding the contract method 0x9a400b0f.
//
// Solidity: function setPairThreshold(address tokenIn, uint256 thresholdInBP) returns()
func (_BeraBorrowCICV *BeraBorrowCICVSession) SetPairThreshold(tokenIn common.Address, thresholdInBP *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.SetPairThreshold(&_BeraBorrowCICV.TransactOpts, tokenIn, thresholdInBP)
}

// SetPairThreshold is a paid mutator transaction binding the contract method 0x9a400b0f.
//
// Solidity: function setPairThreshold(address tokenIn, uint256 thresholdInBP) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) SetPairThreshold(tokenIn common.Address, thresholdInBP *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.SetPairThreshold(&_BeraBorrowCICV.TransactOpts, tokenIn, thresholdInBP)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 _performanceFee) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) SetPerformanceFee(opts *bind.TransactOpts, _performanceFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "setPerformanceFee", _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 _performanceFee) returns()
func (_BeraBorrowCICV *BeraBorrowCICVSession) SetPerformanceFee(_performanceFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.SetPerformanceFee(&_BeraBorrowCICV.TransactOpts, _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 _performanceFee) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) SetPerformanceFee(_performanceFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.SetPerformanceFee(&_BeraBorrowCICV.TransactOpts, _performanceFee)
}

// SetUnlockRatePerSecond is a paid mutator transaction binding the contract method 0x0b983a74.
//
// Solidity: function setUnlockRatePerSecond(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) SetUnlockRatePerSecond(opts *bind.TransactOpts, token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "setUnlockRatePerSecond", token, _unlockRatePerSecond)
}

// SetUnlockRatePerSecond is a paid mutator transaction binding the contract method 0x0b983a74.
//
// Solidity: function setUnlockRatePerSecond(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowCICV *BeraBorrowCICVSession) SetUnlockRatePerSecond(token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.SetUnlockRatePerSecond(&_BeraBorrowCICV.TransactOpts, token, _unlockRatePerSecond)
}

// SetUnlockRatePerSecond is a paid mutator transaction binding the contract method 0x0b983a74.
//
// Solidity: function setUnlockRatePerSecond(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) SetUnlockRatePerSecond(token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.SetUnlockRatePerSecond(&_BeraBorrowCICV.TransactOpts, token, _unlockRatePerSecond)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0x322ce8f0.
//
// Solidity: function setWithdrawFee(uint16 _withdrawFee) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) SetWithdrawFee(opts *bind.TransactOpts, _withdrawFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "setWithdrawFee", _withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0x322ce8f0.
//
// Solidity: function setWithdrawFee(uint16 _withdrawFee) returns()
func (_BeraBorrowCICV *BeraBorrowCICVSession) SetWithdrawFee(_withdrawFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.SetWithdrawFee(&_BeraBorrowCICV.TransactOpts, _withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0x322ce8f0.
//
// Solidity: function setWithdrawFee(uint16 _withdrawFee) returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) SetWithdrawFee(_withdrawFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.SetWithdrawFee(&_BeraBorrowCICV.TransactOpts, _withdrawFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Transfer(&_BeraBorrowCICV.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Transfer(&_BeraBorrowCICV.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BeraBorrowCICV *BeraBorrowCICVSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.TransferFrom(&_BeraBorrowCICV.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.TransferFrom(&_BeraBorrowCICV.TransactOpts, from, to, value)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BeraBorrowCICV *BeraBorrowCICVSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.UpgradeToAndCall(&_BeraBorrowCICV.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.UpgradeToAndCall(&_BeraBorrowCICV.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowCICV *BeraBorrowCICVTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.contract.Transact(opts, "withdraw", assets, receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowCICV *BeraBorrowCICVSession) Withdraw(assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Withdraw(&_BeraBorrowCICV.TransactOpts, assets, receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowCICV *BeraBorrowCICVTransactorSession) Withdraw(assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCICV.Contract.Withdraw(&_BeraBorrowCICV.TransactOpts, assets, receiver, _owner)
}

// BeraBorrowCICVApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BeraBorrowCICV contract.
type BeraBorrowCICVApprovalIterator struct {
	Event *BeraBorrowCICVApproval // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCICVApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCICVApproval)
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
		it.Event = new(BeraBorrowCICVApproval)
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
func (it *BeraBorrowCICVApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCICVApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCICVApproval represents a Approval event raised by the BeraBorrowCICV contract.
type BeraBorrowCICVApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BeraBorrowCICVApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVApprovalIterator{contract: _BeraBorrowCICV.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BeraBorrowCICVApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCICVApproval)
				if err := _BeraBorrowCICV.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) ParseApproval(log types.Log) (*BeraBorrowCICVApproval, error) {
	event := new(BeraBorrowCICVApproval)
	if err := _BeraBorrowCICV.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCICVDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BeraBorrowCICV contract.
type BeraBorrowCICVDepositIterator struct {
	Event *BeraBorrowCICVDeposit // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCICVDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCICVDeposit)
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
		it.Event = new(BeraBorrowCICVDeposit)
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
func (it *BeraBorrowCICVDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCICVDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCICVDeposit represents a Deposit event raised by the BeraBorrowCICV contract.
type BeraBorrowCICVDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*BeraBorrowCICVDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVDepositIterator{contract: _BeraBorrowCICV.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BeraBorrowCICVDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCICVDeposit)
				if err := _BeraBorrowCICV.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) ParseDeposit(log types.Log) (*BeraBorrowCICVDeposit, error) {
	event := new(BeraBorrowCICVDeposit)
	if err := _BeraBorrowCICV.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCICVEmissionsAddedIterator is returned from FilterEmissionsAdded and is used to iterate over the raw logs and unpacked data for EmissionsAdded events raised by the BeraBorrowCICV contract.
type BeraBorrowCICVEmissionsAddedIterator struct {
	Event *BeraBorrowCICVEmissionsAdded // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCICVEmissionsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCICVEmissionsAdded)
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
		it.Event = new(BeraBorrowCICVEmissionsAdded)
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
func (it *BeraBorrowCICVEmissionsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCICVEmissionsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCICVEmissionsAdded represents a EmissionsAdded event raised by the BeraBorrowCICV contract.
type BeraBorrowCICVEmissionsAdded struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmissionsAdded is a free log retrieval operation binding the contract event 0x693ffe037fd29f4846b006bd3ada57d4fd4c3622277227342f0e4fed9011dc20.
//
// Solidity: event EmissionsAdded(address indexed token, uint128 amount)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) FilterEmissionsAdded(opts *bind.FilterOpts, token []common.Address) (*BeraBorrowCICVEmissionsAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.FilterLogs(opts, "EmissionsAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVEmissionsAddedIterator{contract: _BeraBorrowCICV.contract, event: "EmissionsAdded", logs: logs, sub: sub}, nil
}

// WatchEmissionsAdded is a free log subscription operation binding the contract event 0x693ffe037fd29f4846b006bd3ada57d4fd4c3622277227342f0e4fed9011dc20.
//
// Solidity: event EmissionsAdded(address indexed token, uint128 amount)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) WatchEmissionsAdded(opts *bind.WatchOpts, sink chan<- *BeraBorrowCICVEmissionsAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.WatchLogs(opts, "EmissionsAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCICVEmissionsAdded)
				if err := _BeraBorrowCICV.contract.UnpackLog(event, "EmissionsAdded", log); err != nil {
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
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) ParseEmissionsAdded(log types.Log) (*BeraBorrowCICVEmissionsAdded, error) {
	event := new(BeraBorrowCICVEmissionsAdded)
	if err := _BeraBorrowCICV.contract.UnpackLog(event, "EmissionsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCICVInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BeraBorrowCICV contract.
type BeraBorrowCICVInitializedIterator struct {
	Event *BeraBorrowCICVInitialized // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCICVInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCICVInitialized)
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
		it.Event = new(BeraBorrowCICVInitialized)
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
func (it *BeraBorrowCICVInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCICVInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCICVInitialized represents a Initialized event raised by the BeraBorrowCICV contract.
type BeraBorrowCICVInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) FilterInitialized(opts *bind.FilterOpts) (*BeraBorrowCICVInitializedIterator, error) {

	logs, sub, err := _BeraBorrowCICV.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVInitializedIterator{contract: _BeraBorrowCICV.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BeraBorrowCICVInitialized) (event.Subscription, error) {

	logs, sub, err := _BeraBorrowCICV.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCICVInitialized)
				if err := _BeraBorrowCICV.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) ParseInitialized(log types.Log) (*BeraBorrowCICVInitialized, error) {
	event := new(BeraBorrowCICVInitialized)
	if err := _BeraBorrowCICV.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCICVNewUnlockRatePerSecondIterator is returned from FilterNewUnlockRatePerSecond and is used to iterate over the raw logs and unpacked data for NewUnlockRatePerSecond events raised by the BeraBorrowCICV contract.
type BeraBorrowCICVNewUnlockRatePerSecondIterator struct {
	Event *BeraBorrowCICVNewUnlockRatePerSecond // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCICVNewUnlockRatePerSecondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCICVNewUnlockRatePerSecond)
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
		it.Event = new(BeraBorrowCICVNewUnlockRatePerSecond)
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
func (it *BeraBorrowCICVNewUnlockRatePerSecondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCICVNewUnlockRatePerSecondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCICVNewUnlockRatePerSecond represents a NewUnlockRatePerSecond event raised by the BeraBorrowCICV contract.
type BeraBorrowCICVNewUnlockRatePerSecond struct {
	Token               common.Address
	UnlockRatePerSecond uint64
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewUnlockRatePerSecond is a free log retrieval operation binding the contract event 0x5577d4c8f6e5397effa5c71df8fe221e1162e18aaa0aabe87026cfb0c6762150.
//
// Solidity: event NewUnlockRatePerSecond(address indexed token, uint64 unlockRatePerSecond)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) FilterNewUnlockRatePerSecond(opts *bind.FilterOpts, token []common.Address) (*BeraBorrowCICVNewUnlockRatePerSecondIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.FilterLogs(opts, "NewUnlockRatePerSecond", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVNewUnlockRatePerSecondIterator{contract: _BeraBorrowCICV.contract, event: "NewUnlockRatePerSecond", logs: logs, sub: sub}, nil
}

// WatchNewUnlockRatePerSecond is a free log subscription operation binding the contract event 0x5577d4c8f6e5397effa5c71df8fe221e1162e18aaa0aabe87026cfb0c6762150.
//
// Solidity: event NewUnlockRatePerSecond(address indexed token, uint64 unlockRatePerSecond)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) WatchNewUnlockRatePerSecond(opts *bind.WatchOpts, sink chan<- *BeraBorrowCICVNewUnlockRatePerSecond, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.WatchLogs(opts, "NewUnlockRatePerSecond", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCICVNewUnlockRatePerSecond)
				if err := _BeraBorrowCICV.contract.UnpackLog(event, "NewUnlockRatePerSecond", log); err != nil {
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
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) ParseNewUnlockRatePerSecond(log types.Log) (*BeraBorrowCICVNewUnlockRatePerSecond, error) {
	event := new(BeraBorrowCICVNewUnlockRatePerSecond)
	if err := _BeraBorrowCICV.contract.UnpackLog(event, "NewUnlockRatePerSecond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCICVPerformanceFeeIterator is returned from FilterPerformanceFee and is used to iterate over the raw logs and unpacked data for PerformanceFee events raised by the BeraBorrowCICV contract.
type BeraBorrowCICVPerformanceFeeIterator struct {
	Event *BeraBorrowCICVPerformanceFee // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCICVPerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCICVPerformanceFee)
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
		it.Event = new(BeraBorrowCICVPerformanceFee)
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
func (it *BeraBorrowCICVPerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCICVPerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCICVPerformanceFee represents a PerformanceFee event raised by the BeraBorrowCICV contract.
type BeraBorrowCICVPerformanceFee struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPerformanceFee is a free log retrieval operation binding the contract event 0x620154c7367c3805aa5a57be3421bcd8ae0de705b07bb7880c7ebdd1eff84237.
//
// Solidity: event PerformanceFee(address indexed token, uint256 amount)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) FilterPerformanceFee(opts *bind.FilterOpts, token []common.Address) (*BeraBorrowCICVPerformanceFeeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.FilterLogs(opts, "PerformanceFee", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVPerformanceFeeIterator{contract: _BeraBorrowCICV.contract, event: "PerformanceFee", logs: logs, sub: sub}, nil
}

// WatchPerformanceFee is a free log subscription operation binding the contract event 0x620154c7367c3805aa5a57be3421bcd8ae0de705b07bb7880c7ebdd1eff84237.
//
// Solidity: event PerformanceFee(address indexed token, uint256 amount)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) WatchPerformanceFee(opts *bind.WatchOpts, sink chan<- *BeraBorrowCICVPerformanceFee, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.WatchLogs(opts, "PerformanceFee", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCICVPerformanceFee)
				if err := _BeraBorrowCICV.contract.UnpackLog(event, "PerformanceFee", log); err != nil {
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

// ParsePerformanceFee is a log parse operation binding the contract event 0x620154c7367c3805aa5a57be3421bcd8ae0de705b07bb7880c7ebdd1eff84237.
//
// Solidity: event PerformanceFee(address indexed token, uint256 amount)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) ParsePerformanceFee(log types.Log) (*BeraBorrowCICVPerformanceFee, error) {
	event := new(BeraBorrowCICVPerformanceFee)
	if err := _BeraBorrowCICV.contract.UnpackLog(event, "PerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCICVRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the BeraBorrowCICV contract.
type BeraBorrowCICVRebalanceIterator struct {
	Event *BeraBorrowCICVRebalance // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCICVRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCICVRebalance)
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
		it.Event = new(BeraBorrowCICVRebalance)
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
func (it *BeraBorrowCICVRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCICVRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCICVRebalance represents a Rebalance event raised by the BeraBorrowCICV contract.
type BeraBorrowCICVRebalance struct {
	SentCurrency  common.Address
	Sent          *big.Int
	Received      *big.Int
	SentValue     *big.Int
	ReceivedValue *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRebalance is a free log retrieval operation binding the contract event 0x90fea00fe6e24ccfd7c40b40e07486e8b69ada2a3ad693dd908b6479cac36b6b.
//
// Solidity: event Rebalance(address indexed sentCurrency, uint256 sent, uint256 received, uint256 sentValue, uint256 receivedValue)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) FilterRebalance(opts *bind.FilterOpts, sentCurrency []common.Address) (*BeraBorrowCICVRebalanceIterator, error) {

	var sentCurrencyRule []interface{}
	for _, sentCurrencyItem := range sentCurrency {
		sentCurrencyRule = append(sentCurrencyRule, sentCurrencyItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.FilterLogs(opts, "Rebalance", sentCurrencyRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVRebalanceIterator{contract: _BeraBorrowCICV.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0x90fea00fe6e24ccfd7c40b40e07486e8b69ada2a3ad693dd908b6479cac36b6b.
//
// Solidity: event Rebalance(address indexed sentCurrency, uint256 sent, uint256 received, uint256 sentValue, uint256 receivedValue)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *BeraBorrowCICVRebalance, sentCurrency []common.Address) (event.Subscription, error) {

	var sentCurrencyRule []interface{}
	for _, sentCurrencyItem := range sentCurrency {
		sentCurrencyRule = append(sentCurrencyRule, sentCurrencyItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.WatchLogs(opts, "Rebalance", sentCurrencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCICVRebalance)
				if err := _BeraBorrowCICV.contract.UnpackLog(event, "Rebalance", log); err != nil {
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

// ParseRebalance is a log parse operation binding the contract event 0x90fea00fe6e24ccfd7c40b40e07486e8b69ada2a3ad693dd908b6479cac36b6b.
//
// Solidity: event Rebalance(address indexed sentCurrency, uint256 sent, uint256 received, uint256 sentValue, uint256 receivedValue)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) ParseRebalance(log types.Log) (*BeraBorrowCICVRebalance, error) {
	event := new(BeraBorrowCICVRebalance)
	if err := _BeraBorrowCICV.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCICVTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BeraBorrowCICV contract.
type BeraBorrowCICVTransferIterator struct {
	Event *BeraBorrowCICVTransfer // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCICVTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCICVTransfer)
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
		it.Event = new(BeraBorrowCICVTransfer)
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
func (it *BeraBorrowCICVTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCICVTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCICVTransfer represents a Transfer event raised by the BeraBorrowCICV contract.
type BeraBorrowCICVTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BeraBorrowCICVTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVTransferIterator{contract: _BeraBorrowCICV.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BeraBorrowCICVTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCICVTransfer)
				if err := _BeraBorrowCICV.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) ParseTransfer(log types.Log) (*BeraBorrowCICVTransfer, error) {
	event := new(BeraBorrowCICVTransfer)
	if err := _BeraBorrowCICV.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCICVUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BeraBorrowCICV contract.
type BeraBorrowCICVUpgradedIterator struct {
	Event *BeraBorrowCICVUpgraded // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCICVUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCICVUpgraded)
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
		it.Event = new(BeraBorrowCICVUpgraded)
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
func (it *BeraBorrowCICVUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCICVUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCICVUpgraded represents a Upgraded event raised by the BeraBorrowCICV contract.
type BeraBorrowCICVUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BeraBorrowCICVUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVUpgradedIterator{contract: _BeraBorrowCICV.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BeraBorrowCICVUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BeraBorrowCICV.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCICVUpgraded)
				if err := _BeraBorrowCICV.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) ParseUpgraded(log types.Log) (*BeraBorrowCICVUpgraded, error) {
	event := new(BeraBorrowCICVUpgraded)
	if err := _BeraBorrowCICV.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCICVWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BeraBorrowCICV contract.
type BeraBorrowCICVWithdrawIterator struct {
	Event *BeraBorrowCICVWithdraw // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCICVWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCICVWithdraw)
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
		it.Event = new(BeraBorrowCICVWithdraw)
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
func (it *BeraBorrowCICVWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCICVWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCICVWithdraw represents a Withdraw event raised by the BeraBorrowCICV contract.
type BeraBorrowCICVWithdraw struct {
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
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*BeraBorrowCICVWithdrawIterator, error) {

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

	logs, sub, err := _BeraBorrowCICV.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCICVWithdrawIterator{contract: _BeraBorrowCICV.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BeraBorrowCICVWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BeraBorrowCICV.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCICVWithdraw)
				if err := _BeraBorrowCICV.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_BeraBorrowCICV *BeraBorrowCICVFilterer) ParseWithdraw(log types.Log) (*BeraBorrowCICVWithdraw, error) {
	event := new(BeraBorrowCICVWithdraw)
	if err := _BeraBorrowCICV.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
