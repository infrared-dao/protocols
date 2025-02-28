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

// BeraBorrowCDPMetaData contains all meta data concerning the BeraBorrowCDP contract.
var BeraBorrowCDPMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AmountCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmissionRateExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawingLockedEmissions\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"name\":\"EmissionsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"unlockRatePerSecond\",\"type\":\"uint64\"}],\"name\":\"NewUnlockRatePerSecond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PerformanceFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sentCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sentValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedValue\",\"type\":\"uint256\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"_convertToValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalanceOfWithFutureEmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getFullProfitUnlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getLockedEmissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMetaBeraborrowCore\",\"outputs\":[{\"internalType\":\"contractIMetaBeraborrowCore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPerformanceFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceFeed\",\"outputs\":[{\"internalType\":\"contractIPriceFeed\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"iRedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibgt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ibgtVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"infraredVault\",\"outputs\":[{\"internalType\":\"contractIInfraredVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"_minWithdrawFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_maxWithdrawFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_withdrawFee\",\"type\":\"uint16\"},{\"internalType\":\"contractIMetaBeraborrowCore\",\"name\":\"_metaBeraborrowCore\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_sharesName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_sharesSymbol\",\"type\":\"string\"}],\"internalType\":\"structIBaseCollateralVault.BaseInitParams\",\"name\":\"_baseParams\",\"type\":\"tuple\"},{\"internalType\":\"uint16\",\"name\":\"_minPerformanceFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_maxPerformanceFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_performanceFee\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_iRedToken\",\"type\":\"address\"},{\"internalType\":\"contractIInfraredVault\",\"name\":\"_infraredVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ibgtVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_infraredWrapper\",\"type\":\"address\"}],\"internalType\":\"structIInfraredCollateralVault.InfraredInitParams\",\"name\":\"baseParams\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint128[]\",\"name\":\"amounts\",\"type\":\"uint128[]\"}],\"name\":\"internalizeDonations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sentCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sentAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structIInfraredCollateralVault.RebalanceParams\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"receiveDonations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iRedToken\",\"type\":\"address\"}],\"name\":\"setIRED\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"thresholdInBP\",\"type\":\"uint256\"}],\"name\":\"setPairThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_performanceFee\",\"type\":\"uint16\"}],\"name\":\"setPerformanceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_unlockRatePerSecond\",\"type\":\"uint64\"}],\"name\":\"setUnlockRatePerSecond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_withdrawFee\",\"type\":\"uint16\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInAsset\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unlockRatePerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BeraBorrowCDPABI is the input ABI used to generate the binding from.
// Deprecated: Use BeraBorrowCDPMetaData.ABI instead.
var BeraBorrowCDPABI = BeraBorrowCDPMetaData.ABI

// BeraBorrowCDP is an auto generated Go binding around an Ethereum contract.
type BeraBorrowCDP struct {
	BeraBorrowCDPCaller     // Read-only binding to the contract
	BeraBorrowCDPTransactor // Write-only binding to the contract
	BeraBorrowCDPFilterer   // Log filterer for contract events
}

// BeraBorrowCDPCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeraBorrowCDPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowCDPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeraBorrowCDPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowCDPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeraBorrowCDPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowCDPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeraBorrowCDPSession struct {
	Contract     *BeraBorrowCDP    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeraBorrowCDPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeraBorrowCDPCallerSession struct {
	Contract *BeraBorrowCDPCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BeraBorrowCDPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeraBorrowCDPTransactorSession struct {
	Contract     *BeraBorrowCDPTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BeraBorrowCDPRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeraBorrowCDPRaw struct {
	Contract *BeraBorrowCDP // Generic contract binding to access the raw methods on
}

// BeraBorrowCDPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeraBorrowCDPCallerRaw struct {
	Contract *BeraBorrowCDPCaller // Generic read-only contract binding to access the raw methods on
}

// BeraBorrowCDPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeraBorrowCDPTransactorRaw struct {
	Contract *BeraBorrowCDPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeraBorrowCDP creates a new instance of BeraBorrowCDP, bound to a specific deployed contract.
func NewBeraBorrowCDP(address common.Address, backend bind.ContractBackend) (*BeraBorrowCDP, error) {
	contract, err := bindBeraBorrowCDP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDP{BeraBorrowCDPCaller: BeraBorrowCDPCaller{contract: contract}, BeraBorrowCDPTransactor: BeraBorrowCDPTransactor{contract: contract}, BeraBorrowCDPFilterer: BeraBorrowCDPFilterer{contract: contract}}, nil
}

// NewBeraBorrowCDPCaller creates a new read-only instance of BeraBorrowCDP, bound to a specific deployed contract.
func NewBeraBorrowCDPCaller(address common.Address, caller bind.ContractCaller) (*BeraBorrowCDPCaller, error) {
	contract, err := bindBeraBorrowCDP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPCaller{contract: contract}, nil
}

// NewBeraBorrowCDPTransactor creates a new write-only instance of BeraBorrowCDP, bound to a specific deployed contract.
func NewBeraBorrowCDPTransactor(address common.Address, transactor bind.ContractTransactor) (*BeraBorrowCDPTransactor, error) {
	contract, err := bindBeraBorrowCDP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPTransactor{contract: contract}, nil
}

// NewBeraBorrowCDPFilterer creates a new log filterer instance of BeraBorrowCDP, bound to a specific deployed contract.
func NewBeraBorrowCDPFilterer(address common.Address, filterer bind.ContractFilterer) (*BeraBorrowCDPFilterer, error) {
	contract, err := bindBeraBorrowCDP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPFilterer{contract: contract}, nil
}

// bindBeraBorrowCDP binds a generic wrapper to an already deployed contract.
func bindBeraBorrowCDP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeraBorrowCDPMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeraBorrowCDP *BeraBorrowCDPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeraBorrowCDP.Contract.BeraBorrowCDPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeraBorrowCDP *BeraBorrowCDPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.BeraBorrowCDPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeraBorrowCDP *BeraBorrowCDPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.BeraBorrowCDPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeraBorrowCDP *BeraBorrowCDPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeraBorrowCDP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeraBorrowCDP *BeraBorrowCDPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeraBorrowCDP *BeraBorrowCDPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BeraBorrowCDP *BeraBorrowCDPSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BeraBorrowCDP.Contract.UPGRADEINTERFACEVERSION(&_BeraBorrowCDP.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _BeraBorrowCDP.Contract.UPGRADEINTERFACEVERSION(&_BeraBorrowCDP.CallOpts)
}

// ConvertToValue is a free data retrieval call binding the contract method 0xec05c341.
//
// Solidity: function _convertToValue(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) ConvertToValue(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "_convertToValue", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToValue is a free data retrieval call binding the contract method 0xec05c341.
//
// Solidity: function _convertToValue(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) ConvertToValue(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.ConvertToValue(&_BeraBorrowCDP.CallOpts, token)
}

// ConvertToValue is a free data retrieval call binding the contract method 0xec05c341.
//
// Solidity: function _convertToValue(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) ConvertToValue(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.ConvertToValue(&_BeraBorrowCDP.CallOpts, token)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.Allowance(&_BeraBorrowCDP.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.Allowance(&_BeraBorrowCDP.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Asset() (common.Address, error) {
	return _BeraBorrowCDP.Contract.Asset(&_BeraBorrowCDP.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) Asset() (common.Address, error) {
	return _BeraBorrowCDP.Contract.Asset(&_BeraBorrowCDP.CallOpts)
}

// AssetDecimals is a free data retrieval call binding the contract method 0xc2d41601.
//
// Solidity: function assetDecimals() view returns(uint8)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) AssetDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "assetDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AssetDecimals is a free data retrieval call binding the contract method 0xc2d41601.
//
// Solidity: function assetDecimals() view returns(uint8)
func (_BeraBorrowCDP *BeraBorrowCDPSession) AssetDecimals() (uint8, error) {
	return _BeraBorrowCDP.Contract.AssetDecimals(&_BeraBorrowCDP.CallOpts)
}

// AssetDecimals is a free data retrieval call binding the contract method 0xc2d41601.
//
// Solidity: function assetDecimals() view returns(uint8)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) AssetDecimals() (uint8, error) {
	return _BeraBorrowCDP.Contract.AssetDecimals(&_BeraBorrowCDP.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.BalanceOf(&_BeraBorrowCDP.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.BalanceOf(&_BeraBorrowCDP.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.ConvertToAssets(&_BeraBorrowCDP.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.ConvertToAssets(&_BeraBorrowCDP.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.ConvertToShares(&_BeraBorrowCDP.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.ConvertToShares(&_BeraBorrowCDP.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Decimals() (uint8, error) {
	return _BeraBorrowCDP.Contract.Decimals(&_BeraBorrowCDP.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) Decimals() (uint8, error) {
	return _BeraBorrowCDP.Contract.Decimals(&_BeraBorrowCDP.CallOpts)
}

// FetchPrice is a free data retrieval call binding the contract method 0x0fdb11cf.
//
// Solidity: function fetchPrice() view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) FetchPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "fetchPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FetchPrice is a free data retrieval call binding the contract method 0x0fdb11cf.
//
// Solidity: function fetchPrice() view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) FetchPrice() (*big.Int, error) {
	return _BeraBorrowCDP.Contract.FetchPrice(&_BeraBorrowCDP.CallOpts)
}

// FetchPrice is a free data retrieval call binding the contract method 0x0fdb11cf.
//
// Solidity: function fetchPrice() view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) FetchPrice() (*big.Int, error) {
	return _BeraBorrowCDP.Contract.FetchPrice(&_BeraBorrowCDP.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.GetBalance(&_BeraBorrowCDP.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.GetBalance(&_BeraBorrowCDP.CallOpts, token)
}

// GetBalanceOfWithFutureEmissions is a free data retrieval call binding the contract method 0x29d78c88.
//
// Solidity: function getBalanceOfWithFutureEmissions(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) GetBalanceOfWithFutureEmissions(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "getBalanceOfWithFutureEmissions", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOfWithFutureEmissions is a free data retrieval call binding the contract method 0x29d78c88.
//
// Solidity: function getBalanceOfWithFutureEmissions(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) GetBalanceOfWithFutureEmissions(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.GetBalanceOfWithFutureEmissions(&_BeraBorrowCDP.CallOpts, token)
}

// GetBalanceOfWithFutureEmissions is a free data retrieval call binding the contract method 0x29d78c88.
//
// Solidity: function getBalanceOfWithFutureEmissions(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) GetBalanceOfWithFutureEmissions(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.GetBalanceOfWithFutureEmissions(&_BeraBorrowCDP.CallOpts, token)
}

// GetFullProfitUnlockTimestamp is a free data retrieval call binding the contract method 0xd5c16d79.
//
// Solidity: function getFullProfitUnlockTimestamp(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) GetFullProfitUnlockTimestamp(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "getFullProfitUnlockTimestamp", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFullProfitUnlockTimestamp is a free data retrieval call binding the contract method 0xd5c16d79.
//
// Solidity: function getFullProfitUnlockTimestamp(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) GetFullProfitUnlockTimestamp(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.GetFullProfitUnlockTimestamp(&_BeraBorrowCDP.CallOpts, token)
}

// GetFullProfitUnlockTimestamp is a free data retrieval call binding the contract method 0xd5c16d79.
//
// Solidity: function getFullProfitUnlockTimestamp(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) GetFullProfitUnlockTimestamp(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.GetFullProfitUnlockTimestamp(&_BeraBorrowCDP.CallOpts, token)
}

// GetLockedEmissions is a free data retrieval call binding the contract method 0x403dd3bc.
//
// Solidity: function getLockedEmissions(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) GetLockedEmissions(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "getLockedEmissions", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedEmissions is a free data retrieval call binding the contract method 0x403dd3bc.
//
// Solidity: function getLockedEmissions(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) GetLockedEmissions(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.GetLockedEmissions(&_BeraBorrowCDP.CallOpts, token)
}

// GetLockedEmissions is a free data retrieval call binding the contract method 0x403dd3bc.
//
// Solidity: function getLockedEmissions(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) GetLockedEmissions(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.GetLockedEmissions(&_BeraBorrowCDP.CallOpts, token)
}

// GetMetaBeraborrowCore is a free data retrieval call binding the contract method 0x221085d7.
//
// Solidity: function getMetaBeraborrowCore() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) GetMetaBeraborrowCore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "getMetaBeraborrowCore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMetaBeraborrowCore is a free data retrieval call binding the contract method 0x221085d7.
//
// Solidity: function getMetaBeraborrowCore() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPSession) GetMetaBeraborrowCore() (common.Address, error) {
	return _BeraBorrowCDP.Contract.GetMetaBeraborrowCore(&_BeraBorrowCDP.CallOpts)
}

// GetMetaBeraborrowCore is a free data retrieval call binding the contract method 0x221085d7.
//
// Solidity: function getMetaBeraborrowCore() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) GetMetaBeraborrowCore() (common.Address, error) {
	return _BeraBorrowCDP.Contract.GetMetaBeraborrowCore(&_BeraBorrowCDP.CallOpts)
}

// GetPerformanceFee is a free data retrieval call binding the contract method 0x235c3603.
//
// Solidity: function getPerformanceFee() view returns(uint16)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) GetPerformanceFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "getPerformanceFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetPerformanceFee is a free data retrieval call binding the contract method 0x235c3603.
//
// Solidity: function getPerformanceFee() view returns(uint16)
func (_BeraBorrowCDP *BeraBorrowCDPSession) GetPerformanceFee() (uint16, error) {
	return _BeraBorrowCDP.Contract.GetPerformanceFee(&_BeraBorrowCDP.CallOpts)
}

// GetPerformanceFee is a free data retrieval call binding the contract method 0x235c3603.
//
// Solidity: function getPerformanceFee() view returns(uint16)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) GetPerformanceFee() (uint16, error) {
	return _BeraBorrowCDP.Contract.GetPerformanceFee(&_BeraBorrowCDP.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) GetPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "getPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) GetPrice(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.GetPrice(&_BeraBorrowCDP.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) GetPrice(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.GetPrice(&_BeraBorrowCDP.CallOpts, token)
}

// GetPriceFeed is a free data retrieval call binding the contract method 0x9e87a5cd.
//
// Solidity: function getPriceFeed() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) GetPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "getPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceFeed is a free data retrieval call binding the contract method 0x9e87a5cd.
//
// Solidity: function getPriceFeed() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPSession) GetPriceFeed() (common.Address, error) {
	return _BeraBorrowCDP.Contract.GetPriceFeed(&_BeraBorrowCDP.CallOpts)
}

// GetPriceFeed is a free data retrieval call binding the contract method 0x9e87a5cd.
//
// Solidity: function getPriceFeed() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) GetPriceFeed() (common.Address, error) {
	return _BeraBorrowCDP.Contract.GetPriceFeed(&_BeraBorrowCDP.CallOpts)
}

// GetWithdrawFee is a free data retrieval call binding the contract method 0x1540aa89.
//
// Solidity: function getWithdrawFee() view returns(uint16)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) GetWithdrawFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "getWithdrawFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetWithdrawFee is a free data retrieval call binding the contract method 0x1540aa89.
//
// Solidity: function getWithdrawFee() view returns(uint16)
func (_BeraBorrowCDP *BeraBorrowCDPSession) GetWithdrawFee() (uint16, error) {
	return _BeraBorrowCDP.Contract.GetWithdrawFee(&_BeraBorrowCDP.CallOpts)
}

// GetWithdrawFee is a free data retrieval call binding the contract method 0x1540aa89.
//
// Solidity: function getWithdrawFee() view returns(uint16)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) GetWithdrawFee() (uint16, error) {
	return _BeraBorrowCDP.Contract.GetWithdrawFee(&_BeraBorrowCDP.CallOpts)
}

// IRedToken is a free data retrieval call binding the contract method 0xe993e7ed.
//
// Solidity: function iRedToken() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) IRedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "iRedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IRedToken is a free data retrieval call binding the contract method 0xe993e7ed.
//
// Solidity: function iRedToken() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPSession) IRedToken() (common.Address, error) {
	return _BeraBorrowCDP.Contract.IRedToken(&_BeraBorrowCDP.CallOpts)
}

// IRedToken is a free data retrieval call binding the contract method 0xe993e7ed.
//
// Solidity: function iRedToken() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) IRedToken() (common.Address, error) {
	return _BeraBorrowCDP.Contract.IRedToken(&_BeraBorrowCDP.CallOpts)
}

// Ibgt is a free data retrieval call binding the contract method 0x3dafa4f3.
//
// Solidity: function ibgt() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) Ibgt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "ibgt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ibgt is a free data retrieval call binding the contract method 0x3dafa4f3.
//
// Solidity: function ibgt() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Ibgt() (common.Address, error) {
	return _BeraBorrowCDP.Contract.Ibgt(&_BeraBorrowCDP.CallOpts)
}

// Ibgt is a free data retrieval call binding the contract method 0x3dafa4f3.
//
// Solidity: function ibgt() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) Ibgt() (common.Address, error) {
	return _BeraBorrowCDP.Contract.Ibgt(&_BeraBorrowCDP.CallOpts)
}

// IbgtVault is a free data retrieval call binding the contract method 0xfd64c377.
//
// Solidity: function ibgtVault() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) IbgtVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "ibgtVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IbgtVault is a free data retrieval call binding the contract method 0xfd64c377.
//
// Solidity: function ibgtVault() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPSession) IbgtVault() (common.Address, error) {
	return _BeraBorrowCDP.Contract.IbgtVault(&_BeraBorrowCDP.CallOpts)
}

// IbgtVault is a free data retrieval call binding the contract method 0xfd64c377.
//
// Solidity: function ibgtVault() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) IbgtVault() (common.Address, error) {
	return _BeraBorrowCDP.Contract.IbgtVault(&_BeraBorrowCDP.CallOpts)
}

// InfraredVault is a free data retrieval call binding the contract method 0x6ab3b0e5.
//
// Solidity: function infraredVault() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) InfraredVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "infraredVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InfraredVault is a free data retrieval call binding the contract method 0x6ab3b0e5.
//
// Solidity: function infraredVault() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPSession) InfraredVault() (common.Address, error) {
	return _BeraBorrowCDP.Contract.InfraredVault(&_BeraBorrowCDP.CallOpts)
}

// InfraredVault is a free data retrieval call binding the contract method 0x6ab3b0e5.
//
// Solidity: function infraredVault() view returns(address)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) InfraredVault() (common.Address, error) {
	return _BeraBorrowCDP.Contract.InfraredVault(&_BeraBorrowCDP.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.MaxDeposit(&_BeraBorrowCDP.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.MaxDeposit(&_BeraBorrowCDP.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.MaxMint(&_BeraBorrowCDP.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.MaxMint(&_BeraBorrowCDP.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.MaxRedeem(&_BeraBorrowCDP.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.MaxRedeem(&_BeraBorrowCDP.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.MaxWithdraw(&_BeraBorrowCDP.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.MaxWithdraw(&_BeraBorrowCDP.CallOpts, _owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Name() (string, error) {
	return _BeraBorrowCDP.Contract.Name(&_BeraBorrowCDP.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) Name() (string, error) {
	return _BeraBorrowCDP.Contract.Name(&_BeraBorrowCDP.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.PreviewDeposit(&_BeraBorrowCDP.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.PreviewDeposit(&_BeraBorrowCDP.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.PreviewMint(&_BeraBorrowCDP.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.PreviewMint(&_BeraBorrowCDP.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.PreviewRedeem(&_BeraBorrowCDP.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.PreviewRedeem(&_BeraBorrowCDP.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.PreviewWithdraw(&_BeraBorrowCDP.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.PreviewWithdraw(&_BeraBorrowCDP.CallOpts, assets)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BeraBorrowCDP *BeraBorrowCDPSession) ProxiableUUID() ([32]byte, error) {
	return _BeraBorrowCDP.Contract.ProxiableUUID(&_BeraBorrowCDP.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BeraBorrowCDP.Contract.ProxiableUUID(&_BeraBorrowCDP.CallOpts)
}

// RewardedTokens is a free data retrieval call binding the contract method 0xaaffe076.
//
// Solidity: function rewardedTokens() view returns(address[])
func (_BeraBorrowCDP *BeraBorrowCDPCaller) RewardedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "rewardedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RewardedTokens is a free data retrieval call binding the contract method 0xaaffe076.
//
// Solidity: function rewardedTokens() view returns(address[])
func (_BeraBorrowCDP *BeraBorrowCDPSession) RewardedTokens() ([]common.Address, error) {
	return _BeraBorrowCDP.Contract.RewardedTokens(&_BeraBorrowCDP.CallOpts)
}

// RewardedTokens is a free data retrieval call binding the contract method 0xaaffe076.
//
// Solidity: function rewardedTokens() view returns(address[])
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) RewardedTokens() ([]common.Address, error) {
	return _BeraBorrowCDP.Contract.RewardedTokens(&_BeraBorrowCDP.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Symbol() (string, error) {
	return _BeraBorrowCDP.Contract.Symbol(&_BeraBorrowCDP.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) Symbol() (string, error) {
	return _BeraBorrowCDP.Contract.Symbol(&_BeraBorrowCDP.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 amountInAsset)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 amountInAsset)
func (_BeraBorrowCDP *BeraBorrowCDPSession) TotalAssets() (*big.Int, error) {
	return _BeraBorrowCDP.Contract.TotalAssets(&_BeraBorrowCDP.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256 amountInAsset)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) TotalAssets() (*big.Int, error) {
	return _BeraBorrowCDP.Contract.TotalAssets(&_BeraBorrowCDP.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) TotalSupply() (*big.Int, error) {
	return _BeraBorrowCDP.Contract.TotalSupply(&_BeraBorrowCDP.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) TotalSupply() (*big.Int, error) {
	return _BeraBorrowCDP.Contract.TotalSupply(&_BeraBorrowCDP.CallOpts)
}

// UnlockRatePerSecond is a free data retrieval call binding the contract method 0xd9c00197.
//
// Solidity: function unlockRatePerSecond(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCaller) UnlockRatePerSecond(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowCDP.contract.Call(opts, &out, "unlockRatePerSecond", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockRatePerSecond is a free data retrieval call binding the contract method 0xd9c00197.
//
// Solidity: function unlockRatePerSecond(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPSession) UnlockRatePerSecond(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.UnlockRatePerSecond(&_BeraBorrowCDP.CallOpts, token)
}

// UnlockRatePerSecond is a free data retrieval call binding the contract method 0xd9c00197.
//
// Solidity: function unlockRatePerSecond(address token) view returns(uint256)
func (_BeraBorrowCDP *BeraBorrowCDPCallerSession) UnlockRatePerSecond(token common.Address) (*big.Int, error) {
	return _BeraBorrowCDP.Contract.UnlockRatePerSecond(&_BeraBorrowCDP.CallOpts, token)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Approve(&_BeraBorrowCDP.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Approve(&_BeraBorrowCDP.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Deposit(&_BeraBorrowCDP.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Deposit(&_BeraBorrowCDP.TransactOpts, assets, receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xd67524a9.
//
// Solidity: function initialize(((uint16,uint16,uint16,address,address,string,string),uint16,uint16,uint16,address,address,address,address) baseParams) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) Initialize(opts *bind.TransactOpts, baseParams IInfraredCollateralVaultInfraredInitParams) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "initialize", baseParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xd67524a9.
//
// Solidity: function initialize(((uint16,uint16,uint16,address,address,string,string),uint16,uint16,uint16,address,address,address,address) baseParams) returns()
func (_BeraBorrowCDP *BeraBorrowCDPSession) Initialize(baseParams IInfraredCollateralVaultInfraredInitParams) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Initialize(&_BeraBorrowCDP.TransactOpts, baseParams)
}

// Initialize is a paid mutator transaction binding the contract method 0xd67524a9.
//
// Solidity: function initialize(((uint16,uint16,uint16,address,address,string,string),uint16,uint16,uint16,address,address,address,address) baseParams) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) Initialize(baseParams IInfraredCollateralVaultInfraredInitParams) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Initialize(&_BeraBorrowCDP.TransactOpts, baseParams)
}

// InternalizeDonations is a paid mutator transaction binding the contract method 0x3e92f95b.
//
// Solidity: function internalizeDonations(address[] tokens, uint128[] amounts) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) InternalizeDonations(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "internalizeDonations", tokens, amounts)
}

// InternalizeDonations is a paid mutator transaction binding the contract method 0x3e92f95b.
//
// Solidity: function internalizeDonations(address[] tokens, uint128[] amounts) returns()
func (_BeraBorrowCDP *BeraBorrowCDPSession) InternalizeDonations(tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.InternalizeDonations(&_BeraBorrowCDP.TransactOpts, tokens, amounts)
}

// InternalizeDonations is a paid mutator transaction binding the contract method 0x3e92f95b.
//
// Solidity: function internalizeDonations(address[] tokens, uint128[] amounts) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) InternalizeDonations(tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.InternalizeDonations(&_BeraBorrowCDP.TransactOpts, tokens, amounts)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Mint(&_BeraBorrowCDP.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Mint(&_BeraBorrowCDP.TransactOpts, shares, receiver)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb2c06ba0.
//
// Solidity: function rebalance((address,uint256,address,bytes) p) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) Rebalance(opts *bind.TransactOpts, p IInfraredCollateralVaultRebalanceParams) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "rebalance", p)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb2c06ba0.
//
// Solidity: function rebalance((address,uint256,address,bytes) p) returns()
func (_BeraBorrowCDP *BeraBorrowCDPSession) Rebalance(p IInfraredCollateralVaultRebalanceParams) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Rebalance(&_BeraBorrowCDP.TransactOpts, p)
}

// Rebalance is a paid mutator transaction binding the contract method 0xb2c06ba0.
//
// Solidity: function rebalance((address,uint256,address,bytes) p) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) Rebalance(p IInfraredCollateralVaultRebalanceParams) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Rebalance(&_BeraBorrowCDP.TransactOpts, p)
}

// ReceiveDonations is a paid mutator transaction binding the contract method 0x32e1fef2.
//
// Solidity: function receiveDonations(address[] tokens, uint256[] amounts, address receiver) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) ReceiveDonations(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "receiveDonations", tokens, amounts, receiver)
}

// ReceiveDonations is a paid mutator transaction binding the contract method 0x32e1fef2.
//
// Solidity: function receiveDonations(address[] tokens, uint256[] amounts, address receiver) returns()
func (_BeraBorrowCDP *BeraBorrowCDPSession) ReceiveDonations(tokens []common.Address, amounts []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.ReceiveDonations(&_BeraBorrowCDP.TransactOpts, tokens, amounts, receiver)
}

// ReceiveDonations is a paid mutator transaction binding the contract method 0x32e1fef2.
//
// Solidity: function receiveDonations(address[] tokens, uint256[] amounts, address receiver) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) ReceiveDonations(tokens []common.Address, amounts []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.ReceiveDonations(&_BeraBorrowCDP.TransactOpts, tokens, amounts, receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "redeem", shares, receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Redeem(shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Redeem(&_BeraBorrowCDP.TransactOpts, shares, receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address _owner) returns(uint256 assets)
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) Redeem(shares *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Redeem(&_BeraBorrowCDP.TransactOpts, shares, receiver, _owner)
}

// SetIRED is a paid mutator transaction binding the contract method 0x41519a65.
//
// Solidity: function setIRED(address _iRedToken) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) SetIRED(opts *bind.TransactOpts, _iRedToken common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "setIRED", _iRedToken)
}

// SetIRED is a paid mutator transaction binding the contract method 0x41519a65.
//
// Solidity: function setIRED(address _iRedToken) returns()
func (_BeraBorrowCDP *BeraBorrowCDPSession) SetIRED(_iRedToken common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.SetIRED(&_BeraBorrowCDP.TransactOpts, _iRedToken)
}

// SetIRED is a paid mutator transaction binding the contract method 0x41519a65.
//
// Solidity: function setIRED(address _iRedToken) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) SetIRED(_iRedToken common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.SetIRED(&_BeraBorrowCDP.TransactOpts, _iRedToken)
}

// SetPairThreshold is a paid mutator transaction binding the contract method 0x9a400b0f.
//
// Solidity: function setPairThreshold(address tokenIn, uint256 thresholdInBP) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) SetPairThreshold(opts *bind.TransactOpts, tokenIn common.Address, thresholdInBP *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "setPairThreshold", tokenIn, thresholdInBP)
}

// SetPairThreshold is a paid mutator transaction binding the contract method 0x9a400b0f.
//
// Solidity: function setPairThreshold(address tokenIn, uint256 thresholdInBP) returns()
func (_BeraBorrowCDP *BeraBorrowCDPSession) SetPairThreshold(tokenIn common.Address, thresholdInBP *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.SetPairThreshold(&_BeraBorrowCDP.TransactOpts, tokenIn, thresholdInBP)
}

// SetPairThreshold is a paid mutator transaction binding the contract method 0x9a400b0f.
//
// Solidity: function setPairThreshold(address tokenIn, uint256 thresholdInBP) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) SetPairThreshold(tokenIn common.Address, thresholdInBP *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.SetPairThreshold(&_BeraBorrowCDP.TransactOpts, tokenIn, thresholdInBP)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 _performanceFee) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) SetPerformanceFee(opts *bind.TransactOpts, _performanceFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "setPerformanceFee", _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 _performanceFee) returns()
func (_BeraBorrowCDP *BeraBorrowCDPSession) SetPerformanceFee(_performanceFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.SetPerformanceFee(&_BeraBorrowCDP.TransactOpts, _performanceFee)
}

// SetPerformanceFee is a paid mutator transaction binding the contract method 0xaa290e6d.
//
// Solidity: function setPerformanceFee(uint16 _performanceFee) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) SetPerformanceFee(_performanceFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.SetPerformanceFee(&_BeraBorrowCDP.TransactOpts, _performanceFee)
}

// SetUnlockRatePerSecond is a paid mutator transaction binding the contract method 0x0b983a74.
//
// Solidity: function setUnlockRatePerSecond(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) SetUnlockRatePerSecond(opts *bind.TransactOpts, token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "setUnlockRatePerSecond", token, _unlockRatePerSecond)
}

// SetUnlockRatePerSecond is a paid mutator transaction binding the contract method 0x0b983a74.
//
// Solidity: function setUnlockRatePerSecond(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowCDP *BeraBorrowCDPSession) SetUnlockRatePerSecond(token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.SetUnlockRatePerSecond(&_BeraBorrowCDP.TransactOpts, token, _unlockRatePerSecond)
}

// SetUnlockRatePerSecond is a paid mutator transaction binding the contract method 0x0b983a74.
//
// Solidity: function setUnlockRatePerSecond(address token, uint64 _unlockRatePerSecond) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) SetUnlockRatePerSecond(token common.Address, _unlockRatePerSecond uint64) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.SetUnlockRatePerSecond(&_BeraBorrowCDP.TransactOpts, token, _unlockRatePerSecond)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0x322ce8f0.
//
// Solidity: function setWithdrawFee(uint16 _withdrawFee) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) SetWithdrawFee(opts *bind.TransactOpts, _withdrawFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "setWithdrawFee", _withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0x322ce8f0.
//
// Solidity: function setWithdrawFee(uint16 _withdrawFee) returns()
func (_BeraBorrowCDP *BeraBorrowCDPSession) SetWithdrawFee(_withdrawFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.SetWithdrawFee(&_BeraBorrowCDP.TransactOpts, _withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0x322ce8f0.
//
// Solidity: function setWithdrawFee(uint16 _withdrawFee) returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) SetWithdrawFee(_withdrawFee uint16) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.SetWithdrawFee(&_BeraBorrowCDP.TransactOpts, _withdrawFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Transfer(&_BeraBorrowCDP.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Transfer(&_BeraBorrowCDP.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BeraBorrowCDP *BeraBorrowCDPSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.TransferFrom(&_BeraBorrowCDP.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.TransferFrom(&_BeraBorrowCDP.TransactOpts, from, to, value)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BeraBorrowCDP *BeraBorrowCDPSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.UpgradeToAndCall(&_BeraBorrowCDP.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.UpgradeToAndCall(&_BeraBorrowCDP.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowCDP *BeraBorrowCDPTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.contract.Transact(opts, "withdraw", assets, receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowCDP *BeraBorrowCDPSession) Withdraw(assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Withdraw(&_BeraBorrowCDP.TransactOpts, assets, receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address _owner) returns(uint256 shares)
func (_BeraBorrowCDP *BeraBorrowCDPTransactorSession) Withdraw(assets *big.Int, receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _BeraBorrowCDP.Contract.Withdraw(&_BeraBorrowCDP.TransactOpts, assets, receiver, _owner)
}

// BeraBorrowCDPApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BeraBorrowCDP contract.
type BeraBorrowCDPApprovalIterator struct {
	Event *BeraBorrowCDPApproval // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCDPApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCDPApproval)
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
		it.Event = new(BeraBorrowCDPApproval)
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
func (it *BeraBorrowCDPApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCDPApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCDPApproval represents a Approval event raised by the BeraBorrowCDP contract.
type BeraBorrowCDPApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BeraBorrowCDPApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPApprovalIterator{contract: _BeraBorrowCDP.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BeraBorrowCDPApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCDPApproval)
				if err := _BeraBorrowCDP.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) ParseApproval(log types.Log) (*BeraBorrowCDPApproval, error) {
	event := new(BeraBorrowCDPApproval)
	if err := _BeraBorrowCDP.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCDPDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BeraBorrowCDP contract.
type BeraBorrowCDPDepositIterator struct {
	Event *BeraBorrowCDPDeposit // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCDPDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCDPDeposit)
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
		it.Event = new(BeraBorrowCDPDeposit)
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
func (it *BeraBorrowCDPDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCDPDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCDPDeposit represents a Deposit event raised by the BeraBorrowCDP contract.
type BeraBorrowCDPDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*BeraBorrowCDPDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPDepositIterator{contract: _BeraBorrowCDP.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BeraBorrowCDPDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCDPDeposit)
				if err := _BeraBorrowCDP.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) ParseDeposit(log types.Log) (*BeraBorrowCDPDeposit, error) {
	event := new(BeraBorrowCDPDeposit)
	if err := _BeraBorrowCDP.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCDPEmissionsAddedIterator is returned from FilterEmissionsAdded and is used to iterate over the raw logs and unpacked data for EmissionsAdded events raised by the BeraBorrowCDP contract.
type BeraBorrowCDPEmissionsAddedIterator struct {
	Event *BeraBorrowCDPEmissionsAdded // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCDPEmissionsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCDPEmissionsAdded)
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
		it.Event = new(BeraBorrowCDPEmissionsAdded)
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
func (it *BeraBorrowCDPEmissionsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCDPEmissionsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCDPEmissionsAdded represents a EmissionsAdded event raised by the BeraBorrowCDP contract.
type BeraBorrowCDPEmissionsAdded struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmissionsAdded is a free log retrieval operation binding the contract event 0x693ffe037fd29f4846b006bd3ada57d4fd4c3622277227342f0e4fed9011dc20.
//
// Solidity: event EmissionsAdded(address indexed token, uint128 amount)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) FilterEmissionsAdded(opts *bind.FilterOpts, token []common.Address) (*BeraBorrowCDPEmissionsAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.FilterLogs(opts, "EmissionsAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPEmissionsAddedIterator{contract: _BeraBorrowCDP.contract, event: "EmissionsAdded", logs: logs, sub: sub}, nil
}

// WatchEmissionsAdded is a free log subscription operation binding the contract event 0x693ffe037fd29f4846b006bd3ada57d4fd4c3622277227342f0e4fed9011dc20.
//
// Solidity: event EmissionsAdded(address indexed token, uint128 amount)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) WatchEmissionsAdded(opts *bind.WatchOpts, sink chan<- *BeraBorrowCDPEmissionsAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.WatchLogs(opts, "EmissionsAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCDPEmissionsAdded)
				if err := _BeraBorrowCDP.contract.UnpackLog(event, "EmissionsAdded", log); err != nil {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) ParseEmissionsAdded(log types.Log) (*BeraBorrowCDPEmissionsAdded, error) {
	event := new(BeraBorrowCDPEmissionsAdded)
	if err := _BeraBorrowCDP.contract.UnpackLog(event, "EmissionsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCDPInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BeraBorrowCDP contract.
type BeraBorrowCDPInitializedIterator struct {
	Event *BeraBorrowCDPInitialized // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCDPInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCDPInitialized)
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
		it.Event = new(BeraBorrowCDPInitialized)
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
func (it *BeraBorrowCDPInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCDPInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCDPInitialized represents a Initialized event raised by the BeraBorrowCDP contract.
type BeraBorrowCDPInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) FilterInitialized(opts *bind.FilterOpts) (*BeraBorrowCDPInitializedIterator, error) {

	logs, sub, err := _BeraBorrowCDP.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPInitializedIterator{contract: _BeraBorrowCDP.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BeraBorrowCDPInitialized) (event.Subscription, error) {

	logs, sub, err := _BeraBorrowCDP.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCDPInitialized)
				if err := _BeraBorrowCDP.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) ParseInitialized(log types.Log) (*BeraBorrowCDPInitialized, error) {
	event := new(BeraBorrowCDPInitialized)
	if err := _BeraBorrowCDP.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCDPNewUnlockRatePerSecondIterator is returned from FilterNewUnlockRatePerSecond and is used to iterate over the raw logs and unpacked data for NewUnlockRatePerSecond events raised by the BeraBorrowCDP contract.
type BeraBorrowCDPNewUnlockRatePerSecondIterator struct {
	Event *BeraBorrowCDPNewUnlockRatePerSecond // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCDPNewUnlockRatePerSecondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCDPNewUnlockRatePerSecond)
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
		it.Event = new(BeraBorrowCDPNewUnlockRatePerSecond)
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
func (it *BeraBorrowCDPNewUnlockRatePerSecondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCDPNewUnlockRatePerSecondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCDPNewUnlockRatePerSecond represents a NewUnlockRatePerSecond event raised by the BeraBorrowCDP contract.
type BeraBorrowCDPNewUnlockRatePerSecond struct {
	Token               common.Address
	UnlockRatePerSecond uint64
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewUnlockRatePerSecond is a free log retrieval operation binding the contract event 0x5577d4c8f6e5397effa5c71df8fe221e1162e18aaa0aabe87026cfb0c6762150.
//
// Solidity: event NewUnlockRatePerSecond(address indexed token, uint64 unlockRatePerSecond)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) FilterNewUnlockRatePerSecond(opts *bind.FilterOpts, token []common.Address) (*BeraBorrowCDPNewUnlockRatePerSecondIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.FilterLogs(opts, "NewUnlockRatePerSecond", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPNewUnlockRatePerSecondIterator{contract: _BeraBorrowCDP.contract, event: "NewUnlockRatePerSecond", logs: logs, sub: sub}, nil
}

// WatchNewUnlockRatePerSecond is a free log subscription operation binding the contract event 0x5577d4c8f6e5397effa5c71df8fe221e1162e18aaa0aabe87026cfb0c6762150.
//
// Solidity: event NewUnlockRatePerSecond(address indexed token, uint64 unlockRatePerSecond)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) WatchNewUnlockRatePerSecond(opts *bind.WatchOpts, sink chan<- *BeraBorrowCDPNewUnlockRatePerSecond, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.WatchLogs(opts, "NewUnlockRatePerSecond", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCDPNewUnlockRatePerSecond)
				if err := _BeraBorrowCDP.contract.UnpackLog(event, "NewUnlockRatePerSecond", log); err != nil {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) ParseNewUnlockRatePerSecond(log types.Log) (*BeraBorrowCDPNewUnlockRatePerSecond, error) {
	event := new(BeraBorrowCDPNewUnlockRatePerSecond)
	if err := _BeraBorrowCDP.contract.UnpackLog(event, "NewUnlockRatePerSecond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCDPPerformanceFeeIterator is returned from FilterPerformanceFee and is used to iterate over the raw logs and unpacked data for PerformanceFee events raised by the BeraBorrowCDP contract.
type BeraBorrowCDPPerformanceFeeIterator struct {
	Event *BeraBorrowCDPPerformanceFee // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCDPPerformanceFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCDPPerformanceFee)
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
		it.Event = new(BeraBorrowCDPPerformanceFee)
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
func (it *BeraBorrowCDPPerformanceFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCDPPerformanceFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCDPPerformanceFee represents a PerformanceFee event raised by the BeraBorrowCDP contract.
type BeraBorrowCDPPerformanceFee struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPerformanceFee is a free log retrieval operation binding the contract event 0x620154c7367c3805aa5a57be3421bcd8ae0de705b07bb7880c7ebdd1eff84237.
//
// Solidity: event PerformanceFee(address indexed token, uint256 amount)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) FilterPerformanceFee(opts *bind.FilterOpts, token []common.Address) (*BeraBorrowCDPPerformanceFeeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.FilterLogs(opts, "PerformanceFee", tokenRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPPerformanceFeeIterator{contract: _BeraBorrowCDP.contract, event: "PerformanceFee", logs: logs, sub: sub}, nil
}

// WatchPerformanceFee is a free log subscription operation binding the contract event 0x620154c7367c3805aa5a57be3421bcd8ae0de705b07bb7880c7ebdd1eff84237.
//
// Solidity: event PerformanceFee(address indexed token, uint256 amount)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) WatchPerformanceFee(opts *bind.WatchOpts, sink chan<- *BeraBorrowCDPPerformanceFee, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.WatchLogs(opts, "PerformanceFee", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCDPPerformanceFee)
				if err := _BeraBorrowCDP.contract.UnpackLog(event, "PerformanceFee", log); err != nil {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) ParsePerformanceFee(log types.Log) (*BeraBorrowCDPPerformanceFee, error) {
	event := new(BeraBorrowCDPPerformanceFee)
	if err := _BeraBorrowCDP.contract.UnpackLog(event, "PerformanceFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCDPRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the BeraBorrowCDP contract.
type BeraBorrowCDPRebalanceIterator struct {
	Event *BeraBorrowCDPRebalance // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCDPRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCDPRebalance)
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
		it.Event = new(BeraBorrowCDPRebalance)
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
func (it *BeraBorrowCDPRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCDPRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCDPRebalance represents a Rebalance event raised by the BeraBorrowCDP contract.
type BeraBorrowCDPRebalance struct {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) FilterRebalance(opts *bind.FilterOpts, sentCurrency []common.Address) (*BeraBorrowCDPRebalanceIterator, error) {

	var sentCurrencyRule []interface{}
	for _, sentCurrencyItem := range sentCurrency {
		sentCurrencyRule = append(sentCurrencyRule, sentCurrencyItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.FilterLogs(opts, "Rebalance", sentCurrencyRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPRebalanceIterator{contract: _BeraBorrowCDP.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0x90fea00fe6e24ccfd7c40b40e07486e8b69ada2a3ad693dd908b6479cac36b6b.
//
// Solidity: event Rebalance(address indexed sentCurrency, uint256 sent, uint256 received, uint256 sentValue, uint256 receivedValue)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *BeraBorrowCDPRebalance, sentCurrency []common.Address) (event.Subscription, error) {

	var sentCurrencyRule []interface{}
	for _, sentCurrencyItem := range sentCurrency {
		sentCurrencyRule = append(sentCurrencyRule, sentCurrencyItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.WatchLogs(opts, "Rebalance", sentCurrencyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCDPRebalance)
				if err := _BeraBorrowCDP.contract.UnpackLog(event, "Rebalance", log); err != nil {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) ParseRebalance(log types.Log) (*BeraBorrowCDPRebalance, error) {
	event := new(BeraBorrowCDPRebalance)
	if err := _BeraBorrowCDP.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCDPTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BeraBorrowCDP contract.
type BeraBorrowCDPTransferIterator struct {
	Event *BeraBorrowCDPTransfer // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCDPTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCDPTransfer)
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
		it.Event = new(BeraBorrowCDPTransfer)
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
func (it *BeraBorrowCDPTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCDPTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCDPTransfer represents a Transfer event raised by the BeraBorrowCDP contract.
type BeraBorrowCDPTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BeraBorrowCDPTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPTransferIterator{contract: _BeraBorrowCDP.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BeraBorrowCDPTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCDPTransfer)
				if err := _BeraBorrowCDP.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) ParseTransfer(log types.Log) (*BeraBorrowCDPTransfer, error) {
	event := new(BeraBorrowCDPTransfer)
	if err := _BeraBorrowCDP.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCDPUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BeraBorrowCDP contract.
type BeraBorrowCDPUpgradedIterator struct {
	Event *BeraBorrowCDPUpgraded // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCDPUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCDPUpgraded)
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
		it.Event = new(BeraBorrowCDPUpgraded)
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
func (it *BeraBorrowCDPUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCDPUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCDPUpgraded represents a Upgraded event raised by the BeraBorrowCDP contract.
type BeraBorrowCDPUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BeraBorrowCDPUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPUpgradedIterator{contract: _BeraBorrowCDP.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BeraBorrowCDPUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BeraBorrowCDP.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCDPUpgraded)
				if err := _BeraBorrowCDP.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) ParseUpgraded(log types.Log) (*BeraBorrowCDPUpgraded, error) {
	event := new(BeraBorrowCDPUpgraded)
	if err := _BeraBorrowCDP.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowCDPWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BeraBorrowCDP contract.
type BeraBorrowCDPWithdrawIterator struct {
	Event *BeraBorrowCDPWithdraw // Event containing the contract specifics and raw log

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
func (it *BeraBorrowCDPWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowCDPWithdraw)
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
		it.Event = new(BeraBorrowCDPWithdraw)
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
func (it *BeraBorrowCDPWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowCDPWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowCDPWithdraw represents a Withdraw event raised by the BeraBorrowCDP contract.
type BeraBorrowCDPWithdraw struct {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*BeraBorrowCDPWithdrawIterator, error) {

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

	logs, sub, err := _BeraBorrowCDP.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowCDPWithdrawIterator{contract: _BeraBorrowCDP.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BeraBorrowCDPWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BeraBorrowCDP.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowCDPWithdraw)
				if err := _BeraBorrowCDP.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_BeraBorrowCDP *BeraBorrowCDPFilterer) ParseWithdraw(log types.Log) (*BeraBorrowCDPWithdraw, error) {
	event := new(BeraBorrowCDPWithdraw)
	if err := _BeraBorrowCDP.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
