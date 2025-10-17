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

// MarketAllocation is an auto generated low-level Go binding around an user-defined struct.
type MarketAllocation struct {
	MarketParams MarketParams
	Assets       *big.Int
}

// MarketParams is an auto generated low-level Go binding around an user-defined struct.
type MarketParams struct {
	LoanToken       common.Address
	CollateralToken common.Address
	Oracle          common.Address
	Irm             common.Address
	Lltv            *big.Int
}

// BendVaultMetaData contains all meta data concerning the BendVault contract.
var BendVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"morpho\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feePartitioner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialTimelock\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"__name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"__symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AboveMaxTimelock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllCapsReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BelowMinTimelock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DuplicateMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ERC2612ExpiredSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC2612InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxRedeem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ERC4626ExceededMaxWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InconsistentAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InconsistentReallocation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalNonZeroCap\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalNonZeroSupply\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"InvalidMarketRemovalTimelockNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketNotCreated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"MarketNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxQueueLengthExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPendingValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllocatorRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCuratorNorGuardianRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCuratorRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGuardianRole\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"PendingCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRemoval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotElapsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"UnauthorizedMarket\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFeeRecipient\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeShares\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"suppliedAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"suppliedShares\",\"type\":\"uint256\"}],\"name\":\"ReallocateSupply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnShares\",\"type\":\"uint256\"}],\"name\":\"ReallocateWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RevokePendingCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RevokePendingMarketRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"RevokePendingTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"SetCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newCurator\",\"type\":\"address\"}],\"name\":\"SetCurator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"SetFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"SetGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"allocator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllocator\",\"type\":\"bool\"}],\"name\":\"SetIsAllocator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"SetName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSkimRecipient\",\"type\":\"address\"}],\"name\":\"SetSkimRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"Id[]\",\"name\":\"newSupplyQueue\",\"type\":\"bytes32[]\"}],\"name\":\"SetSupplyQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"SetSymbol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"SetTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"Id[]\",\"name\":\"newWithdrawQueue\",\"type\":\"bytes32[]\"}],\"name\":\"SetWithdrawQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Skim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"SubmitCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"SubmitGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"SubmitMarketRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"SubmitTimelock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedTotalAssets\",\"type\":\"uint256\"}],\"name\":\"UpdateLastTotalAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLostAssets\",\"type\":\"uint256\"}],\"name\":\"UpdateLostAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DECIMALS_OFFSET\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MORPHO\",\"outputs\":[{\"internalType\":\"contractIMorpho\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"acceptCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"config\",\"outputs\":[{\"internalType\":\"uint184\",\"name\":\"cap\",\"type\":\"uint184\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"removableAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"curator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAllocator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTotalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lostAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"pendingCap\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingTimelock\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"value\",\"type\":\"uint192\"},{\"internalType\":\"uint64\",\"name\":\"validAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"internalType\":\"structMarketAllocation[]\",\"name\":\"allocations\",\"type\":\"tuple[]\"}],\"name\":\"reallocate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"revokePendingCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"revokePendingMarketRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCurator\",\"type\":\"address\"}],\"name\":\"setCurator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAllocator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"newIsAllocator\",\"type\":\"bool\"}],\"name\":\"setIsAllocator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSkimRecipient\",\"type\":\"address\"}],\"name\":\"setSkimRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Id[]\",\"name\":\"newSupplyQueue\",\"type\":\"bytes32[]\"}],\"name\":\"setSupplyQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newSymbol\",\"type\":\"string\"}],\"name\":\"setSymbol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skimRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"newSupplyCap\",\"type\":\"uint256\"}],\"name\":\"submitCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"submitGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"irm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lltv\",\"type\":\"uint256\"}],\"internalType\":\"structMarketParams\",\"name\":\"marketParams\",\"type\":\"tuple\"}],\"name\":\"submitMarketRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimelock\",\"type\":\"uint256\"}],\"name\":\"submitTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supplyQueue\",\"outputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"supplyQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"indexes\",\"type\":\"uint256[]\"}],\"name\":\"updateWithdrawQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawQueue\",\"outputs\":[{\"internalType\":\"Id\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BendVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use BendVaultMetaData.ABI instead.
var BendVaultABI = BendVaultMetaData.ABI

// BendVault is an auto generated Go binding around an Ethereum contract.
type BendVault struct {
	BendVaultCaller     // Read-only binding to the contract
	BendVaultTransactor // Write-only binding to the contract
	BendVaultFilterer   // Log filterer for contract events
}

// BendVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type BendVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BendVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BendVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BendVaultSession struct {
	Contract     *BendVault        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BendVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BendVaultCallerSession struct {
	Contract *BendVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BendVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BendVaultTransactorSession struct {
	Contract     *BendVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BendVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type BendVaultRaw struct {
	Contract *BendVault // Generic contract binding to access the raw methods on
}

// BendVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BendVaultCallerRaw struct {
	Contract *BendVaultCaller // Generic read-only contract binding to access the raw methods on
}

// BendVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BendVaultTransactorRaw struct {
	Contract *BendVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBendVault creates a new instance of BendVault, bound to a specific deployed contract.
func NewBendVault(address common.Address, backend bind.ContractBackend) (*BendVault, error) {
	contract, err := bindBendVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BendVault{BendVaultCaller: BendVaultCaller{contract: contract}, BendVaultTransactor: BendVaultTransactor{contract: contract}, BendVaultFilterer: BendVaultFilterer{contract: contract}}, nil
}

// NewBendVaultCaller creates a new read-only instance of BendVault, bound to a specific deployed contract.
func NewBendVaultCaller(address common.Address, caller bind.ContractCaller) (*BendVaultCaller, error) {
	contract, err := bindBendVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BendVaultCaller{contract: contract}, nil
}

// NewBendVaultTransactor creates a new write-only instance of BendVault, bound to a specific deployed contract.
func NewBendVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*BendVaultTransactor, error) {
	contract, err := bindBendVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BendVaultTransactor{contract: contract}, nil
}

// NewBendVaultFilterer creates a new log filterer instance of BendVault, bound to a specific deployed contract.
func NewBendVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*BendVaultFilterer, error) {
	contract, err := bindBendVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BendVaultFilterer{contract: contract}, nil
}

// bindBendVault binds a generic wrapper to an already deployed contract.
func bindBendVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BendVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BendVault *BendVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BendVault.Contract.BendVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BendVault *BendVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendVault.Contract.BendVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BendVault *BendVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BendVault.Contract.BendVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BendVault *BendVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BendVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BendVault *BendVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BendVault *BendVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BendVault.Contract.contract.Transact(opts, method, params...)
}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_BendVault *BendVaultCaller) DECIMALSOFFSET(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "DECIMALS_OFFSET")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_BendVault *BendVaultSession) DECIMALSOFFSET() (uint8, error) {
	return _BendVault.Contract.DECIMALSOFFSET(&_BendVault.CallOpts)
}

// DECIMALSOFFSET is a free data retrieval call binding the contract method 0xaea70acc.
//
// Solidity: function DECIMALS_OFFSET() view returns(uint8)
func (_BendVault *BendVaultCallerSession) DECIMALSOFFSET() (uint8, error) {
	return _BendVault.Contract.DECIMALSOFFSET(&_BendVault.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BendVault *BendVaultCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BendVault *BendVaultSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BendVault.Contract.DOMAINSEPARATOR(&_BendVault.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BendVault *BendVaultCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BendVault.Contract.DOMAINSEPARATOR(&_BendVault.CallOpts)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_BendVault *BendVaultCaller) MORPHO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "MORPHO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_BendVault *BendVaultSession) MORPHO() (common.Address, error) {
	return _BendVault.Contract.MORPHO(&_BendVault.CallOpts)
}

// MORPHO is a free data retrieval call binding the contract method 0x3acb5624.
//
// Solidity: function MORPHO() view returns(address)
func (_BendVault *BendVaultCallerSession) MORPHO() (common.Address, error) {
	return _BendVault.Contract.MORPHO(&_BendVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BendVault *BendVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BendVault *BendVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BendVault.Contract.Allowance(&_BendVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BendVault *BendVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BendVault.Contract.Allowance(&_BendVault.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BendVault *BendVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BendVault *BendVaultSession) Asset() (common.Address, error) {
	return _BendVault.Contract.Asset(&_BendVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_BendVault *BendVaultCallerSession) Asset() (common.Address, error) {
	return _BendVault.Contract.Asset(&_BendVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BendVault *BendVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BendVault *BendVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BendVault.Contract.BalanceOf(&_BendVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BendVault *BendVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BendVault.Contract.BalanceOf(&_BendVault.CallOpts, account)
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_BendVault *BendVaultCaller) Config(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "config", arg0)

	outstruct := new(struct {
		Cap         *big.Int
		Enabled     bool
		RemovableAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RemovableAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_BendVault *BendVaultSession) Config(arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	return _BendVault.Contract.Config(&_BendVault.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0xcc718f76.
//
// Solidity: function config(bytes32 ) view returns(uint184 cap, bool enabled, uint64 removableAt)
func (_BendVault *BendVaultCallerSession) Config(arg0 [32]byte) (struct {
	Cap         *big.Int
	Enabled     bool
	RemovableAt uint64
}, error) {
	return _BendVault.Contract.Config(&_BendVault.CallOpts, arg0)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BendVault *BendVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BendVault *BendVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _BendVault.Contract.ConvertToAssets(&_BendVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_BendVault *BendVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _BendVault.Contract.ConvertToAssets(&_BendVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BendVault *BendVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BendVault *BendVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _BendVault.Contract.ConvertToShares(&_BendVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_BendVault *BendVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _BendVault.Contract.ConvertToShares(&_BendVault.CallOpts, assets)
}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_BendVault *BendVaultCaller) Curator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "curator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_BendVault *BendVaultSession) Curator() (common.Address, error) {
	return _BendVault.Contract.Curator(&_BendVault.CallOpts)
}

// Curator is a free data retrieval call binding the contract method 0xe66f53b7.
//
// Solidity: function curator() view returns(address)
func (_BendVault *BendVaultCallerSession) Curator() (common.Address, error) {
	return _BendVault.Contract.Curator(&_BendVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BendVault *BendVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BendVault *BendVaultSession) Decimals() (uint8, error) {
	return _BendVault.Contract.Decimals(&_BendVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BendVault *BendVaultCallerSession) Decimals() (uint8, error) {
	return _BendVault.Contract.Decimals(&_BendVault.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BendVault *BendVaultCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BendVault *BendVaultSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BendVault.Contract.Eip712Domain(&_BendVault.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_BendVault *BendVaultCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _BendVault.Contract.Eip712Domain(&_BendVault.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_BendVault *BendVaultCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_BendVault *BendVaultSession) Fee() (*big.Int, error) {
	return _BendVault.Contract.Fee(&_BendVault.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint96)
func (_BendVault *BendVaultCallerSession) Fee() (*big.Int, error) {
	return _BendVault.Contract.Fee(&_BendVault.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BendVault *BendVaultCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BendVault *BendVaultSession) FeeRecipient() (common.Address, error) {
	return _BendVault.Contract.FeeRecipient(&_BendVault.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BendVault *BendVaultCallerSession) FeeRecipient() (common.Address, error) {
	return _BendVault.Contract.FeeRecipient(&_BendVault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_BendVault *BendVaultCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_BendVault *BendVaultSession) Guardian() (common.Address, error) {
	return _BendVault.Contract.Guardian(&_BendVault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_BendVault *BendVaultCallerSession) Guardian() (common.Address, error) {
	return _BendVault.Contract.Guardian(&_BendVault.CallOpts)
}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_BendVault *BendVaultCaller) IsAllocator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "isAllocator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_BendVault *BendVaultSession) IsAllocator(arg0 common.Address) (bool, error) {
	return _BendVault.Contract.IsAllocator(&_BendVault.CallOpts, arg0)
}

// IsAllocator is a free data retrieval call binding the contract method 0x4dedf20e.
//
// Solidity: function isAllocator(address ) view returns(bool)
func (_BendVault *BendVaultCallerSession) IsAllocator(arg0 common.Address) (bool, error) {
	return _BendVault.Contract.IsAllocator(&_BendVault.CallOpts, arg0)
}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_BendVault *BendVaultCaller) LastTotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "lastTotalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_BendVault *BendVaultSession) LastTotalAssets() (*big.Int, error) {
	return _BendVault.Contract.LastTotalAssets(&_BendVault.CallOpts)
}

// LastTotalAssets is a free data retrieval call binding the contract method 0x568efc07.
//
// Solidity: function lastTotalAssets() view returns(uint256)
func (_BendVault *BendVaultCallerSession) LastTotalAssets() (*big.Int, error) {
	return _BendVault.Contract.LastTotalAssets(&_BendVault.CallOpts)
}

// LostAssets is a free data retrieval call binding the contract method 0x21cb4b14.
//
// Solidity: function lostAssets() view returns(uint256)
func (_BendVault *BendVaultCaller) LostAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "lostAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LostAssets is a free data retrieval call binding the contract method 0x21cb4b14.
//
// Solidity: function lostAssets() view returns(uint256)
func (_BendVault *BendVaultSession) LostAssets() (*big.Int, error) {
	return _BendVault.Contract.LostAssets(&_BendVault.CallOpts)
}

// LostAssets is a free data retrieval call binding the contract method 0x21cb4b14.
//
// Solidity: function lostAssets() view returns(uint256)
func (_BendVault *BendVaultCallerSession) LostAssets() (*big.Int, error) {
	return _BendVault.Contract.LostAssets(&_BendVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BendVault *BendVaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BendVault *BendVaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _BendVault.Contract.MaxDeposit(&_BendVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_BendVault *BendVaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _BendVault.Contract.MaxDeposit(&_BendVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BendVault *BendVaultCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BendVault *BendVaultSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _BendVault.Contract.MaxMint(&_BendVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_BendVault *BendVaultCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _BendVault.Contract.MaxMint(&_BendVault.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BendVault *BendVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BendVault *BendVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _BendVault.Contract.MaxRedeem(&_BendVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_BendVault *BendVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _BendVault.Contract.MaxRedeem(&_BendVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_BendVault *BendVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_BendVault *BendVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _BendVault.Contract.MaxWithdraw(&_BendVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256 assets)
func (_BendVault *BendVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _BendVault.Contract.MaxWithdraw(&_BendVault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BendVault *BendVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BendVault *BendVaultSession) Name() (string, error) {
	return _BendVault.Contract.Name(&_BendVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BendVault *BendVaultCallerSession) Name() (string, error) {
	return _BendVault.Contract.Name(&_BendVault.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BendVault *BendVaultCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BendVault *BendVaultSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BendVault.Contract.Nonces(&_BendVault.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_BendVault *BendVaultCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _BendVault.Contract.Nonces(&_BendVault.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BendVault *BendVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BendVault *BendVaultSession) Owner() (common.Address, error) {
	return _BendVault.Contract.Owner(&_BendVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BendVault *BendVaultCallerSession) Owner() (common.Address, error) {
	return _BendVault.Contract.Owner(&_BendVault.CallOpts)
}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_BendVault *BendVaultCaller) PendingCap(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "pendingCap", arg0)

	outstruct := new(struct {
		Value   *big.Int
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_BendVault *BendVaultSession) PendingCap(arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _BendVault.Contract.PendingCap(&_BendVault.CallOpts, arg0)
}

// PendingCap is a free data retrieval call binding the contract method 0xa31be5d6.
//
// Solidity: function pendingCap(bytes32 ) view returns(uint192 value, uint64 validAt)
func (_BendVault *BendVaultCallerSession) PendingCap(arg0 [32]byte) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _BendVault.Contract.PendingCap(&_BendVault.CallOpts, arg0)
}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_BendVault *BendVaultCaller) PendingGuardian(opts *bind.CallOpts) (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "pendingGuardian")

	outstruct := new(struct {
		Value   common.Address
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_BendVault *BendVaultSession) PendingGuardian() (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	return _BendVault.Contract.PendingGuardian(&_BendVault.CallOpts)
}

// PendingGuardian is a free data retrieval call binding the contract method 0x762c31ba.
//
// Solidity: function pendingGuardian() view returns(address value, uint64 validAt)
func (_BendVault *BendVaultCallerSession) PendingGuardian() (struct {
	Value   common.Address
	ValidAt uint64
}, error) {
	return _BendVault.Contract.PendingGuardian(&_BendVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BendVault *BendVaultCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BendVault *BendVaultSession) PendingOwner() (common.Address, error) {
	return _BendVault.Contract.PendingOwner(&_BendVault.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_BendVault *BendVaultCallerSession) PendingOwner() (common.Address, error) {
	return _BendVault.Contract.PendingOwner(&_BendVault.CallOpts)
}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_BendVault *BendVaultCaller) PendingTimelock(opts *bind.CallOpts) (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "pendingTimelock")

	outstruct := new(struct {
		Value   *big.Int
		ValidAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidAt = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_BendVault *BendVaultSession) PendingTimelock() (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _BendVault.Contract.PendingTimelock(&_BendVault.CallOpts)
}

// PendingTimelock is a free data retrieval call binding the contract method 0x7cc4d9a1.
//
// Solidity: function pendingTimelock() view returns(uint192 value, uint64 validAt)
func (_BendVault *BendVaultCallerSession) PendingTimelock() (struct {
	Value   *big.Int
	ValidAt uint64
}, error) {
	return _BendVault.Contract.PendingTimelock(&_BendVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BendVault *BendVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BendVault *BendVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _BendVault.Contract.PreviewDeposit(&_BendVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_BendVault *BendVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _BendVault.Contract.PreviewDeposit(&_BendVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_BendVault *BendVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_BendVault *BendVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _BendVault.Contract.PreviewMint(&_BendVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_BendVault *BendVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _BendVault.Contract.PreviewMint(&_BendVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BendVault *BendVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BendVault *BendVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _BendVault.Contract.PreviewRedeem(&_BendVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_BendVault *BendVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _BendVault.Contract.PreviewRedeem(&_BendVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BendVault *BendVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BendVault *BendVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _BendVault.Contract.PreviewWithdraw(&_BendVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_BendVault *BendVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _BendVault.Contract.PreviewWithdraw(&_BendVault.CallOpts, assets)
}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_BendVault *BendVaultCaller) SkimRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "skimRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_BendVault *BendVaultSession) SkimRecipient() (common.Address, error) {
	return _BendVault.Contract.SkimRecipient(&_BendVault.CallOpts)
}

// SkimRecipient is a free data retrieval call binding the contract method 0x388af5b5.
//
// Solidity: function skimRecipient() view returns(address)
func (_BendVault *BendVaultCallerSession) SkimRecipient() (common.Address, error) {
	return _BendVault.Contract.SkimRecipient(&_BendVault.CallOpts)
}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_BendVault *BendVaultCaller) SupplyQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "supplyQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_BendVault *BendVaultSession) SupplyQueue(arg0 *big.Int) ([32]byte, error) {
	return _BendVault.Contract.SupplyQueue(&_BendVault.CallOpts, arg0)
}

// SupplyQueue is a free data retrieval call binding the contract method 0xf7d18521.
//
// Solidity: function supplyQueue(uint256 ) view returns(bytes32)
func (_BendVault *BendVaultCallerSession) SupplyQueue(arg0 *big.Int) ([32]byte, error) {
	return _BendVault.Contract.SupplyQueue(&_BendVault.CallOpts, arg0)
}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_BendVault *BendVaultCaller) SupplyQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "supplyQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_BendVault *BendVaultSession) SupplyQueueLength() (*big.Int, error) {
	return _BendVault.Contract.SupplyQueueLength(&_BendVault.CallOpts)
}

// SupplyQueueLength is a free data retrieval call binding the contract method 0xa17b3130.
//
// Solidity: function supplyQueueLength() view returns(uint256)
func (_BendVault *BendVaultCallerSession) SupplyQueueLength() (*big.Int, error) {
	return _BendVault.Contract.SupplyQueueLength(&_BendVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BendVault *BendVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BendVault *BendVaultSession) Symbol() (string, error) {
	return _BendVault.Contract.Symbol(&_BendVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BendVault *BendVaultCallerSession) Symbol() (string, error) {
	return _BendVault.Contract.Symbol(&_BendVault.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_BendVault *BendVaultCaller) Timelock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_BendVault *BendVaultSession) Timelock() (*big.Int, error) {
	return _BendVault.Contract.Timelock(&_BendVault.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(uint256)
func (_BendVault *BendVaultCallerSession) Timelock() (*big.Int, error) {
	return _BendVault.Contract.Timelock(&_BendVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_BendVault *BendVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_BendVault *BendVaultSession) TotalAssets() (*big.Int, error) {
	return _BendVault.Contract.TotalAssets(&_BendVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_BendVault *BendVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _BendVault.Contract.TotalAssets(&_BendVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BendVault *BendVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BendVault *BendVaultSession) TotalSupply() (*big.Int, error) {
	return _BendVault.Contract.TotalSupply(&_BendVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BendVault *BendVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _BendVault.Contract.TotalSupply(&_BendVault.CallOpts)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_BendVault *BendVaultCaller) WithdrawQueue(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "withdrawQueue", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_BendVault *BendVaultSession) WithdrawQueue(arg0 *big.Int) ([32]byte, error) {
	return _BendVault.Contract.WithdrawQueue(&_BendVault.CallOpts, arg0)
}

// WithdrawQueue is a free data retrieval call binding the contract method 0x62518ddf.
//
// Solidity: function withdrawQueue(uint256 ) view returns(bytes32)
func (_BendVault *BendVaultCallerSession) WithdrawQueue(arg0 *big.Int) ([32]byte, error) {
	return _BendVault.Contract.WithdrawQueue(&_BendVault.CallOpts, arg0)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_BendVault *BendVaultCaller) WithdrawQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendVault.contract.Call(opts, &out, "withdrawQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_BendVault *BendVaultSession) WithdrawQueueLength() (*big.Int, error) {
	return _BendVault.Contract.WithdrawQueueLength(&_BendVault.CallOpts)
}

// WithdrawQueueLength is a free data retrieval call binding the contract method 0x33f91ebb.
//
// Solidity: function withdrawQueueLength() view returns(uint256)
func (_BendVault *BendVaultCallerSession) WithdrawQueueLength() (*big.Int, error) {
	return _BendVault.Contract.WithdrawQueueLength(&_BendVault.CallOpts)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_BendVault *BendVaultTransactor) AcceptCap(opts *bind.TransactOpts, marketParams MarketParams) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "acceptCap", marketParams)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_BendVault *BendVaultSession) AcceptCap(marketParams MarketParams) (*types.Transaction, error) {
	return _BendVault.Contract.AcceptCap(&_BendVault.TransactOpts, marketParams)
}

// AcceptCap is a paid mutator transaction binding the contract method 0x6fda3868.
//
// Solidity: function acceptCap((address,address,address,address,uint256) marketParams) returns()
func (_BendVault *BendVaultTransactorSession) AcceptCap(marketParams MarketParams) (*types.Transaction, error) {
	return _BendVault.Contract.AcceptCap(&_BendVault.TransactOpts, marketParams)
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_BendVault *BendVaultTransactor) AcceptGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "acceptGuardian")
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_BendVault *BendVaultSession) AcceptGuardian() (*types.Transaction, error) {
	return _BendVault.Contract.AcceptGuardian(&_BendVault.TransactOpts)
}

// AcceptGuardian is a paid mutator transaction binding the contract method 0xa5f31d61.
//
// Solidity: function acceptGuardian() returns()
func (_BendVault *BendVaultTransactorSession) AcceptGuardian() (*types.Transaction, error) {
	return _BendVault.Contract.AcceptGuardian(&_BendVault.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BendVault *BendVaultTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BendVault *BendVaultSession) AcceptOwnership() (*types.Transaction, error) {
	return _BendVault.Contract.AcceptOwnership(&_BendVault.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BendVault *BendVaultTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BendVault.Contract.AcceptOwnership(&_BendVault.TransactOpts)
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_BendVault *BendVaultTransactor) AcceptTimelock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "acceptTimelock")
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_BendVault *BendVaultSession) AcceptTimelock() (*types.Transaction, error) {
	return _BendVault.Contract.AcceptTimelock(&_BendVault.TransactOpts)
}

// AcceptTimelock is a paid mutator transaction binding the contract method 0x8a2c7b39.
//
// Solidity: function acceptTimelock() returns()
func (_BendVault *BendVaultTransactorSession) AcceptTimelock() (*types.Transaction, error) {
	return _BendVault.Contract.AcceptTimelock(&_BendVault.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BendVault *BendVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BendVault *BendVaultSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.Approve(&_BendVault.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BendVault *BendVaultTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.Approve(&_BendVault.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BendVault *BendVaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BendVault *BendVaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.Deposit(&_BendVault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_BendVault *BendVaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.Deposit(&_BendVault.TransactOpts, assets, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BendVault *BendVaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BendVault *BendVaultSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.Mint(&_BendVault.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_BendVault *BendVaultTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.Mint(&_BendVault.TransactOpts, shares, receiver)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BendVault *BendVaultTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BendVault *BendVaultSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _BendVault.Contract.Multicall(&_BendVault.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_BendVault *BendVaultTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _BendVault.Contract.Multicall(&_BendVault.TransactOpts, data)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BendVault *BendVaultTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BendVault *BendVaultSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BendVault.Contract.Permit(&_BendVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_BendVault *BendVaultTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _BendVault.Contract.Permit(&_BendVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_BendVault *BendVaultTransactor) Reallocate(opts *bind.TransactOpts, allocations []MarketAllocation) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "reallocate", allocations)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_BendVault *BendVaultSession) Reallocate(allocations []MarketAllocation) (*types.Transaction, error) {
	return _BendVault.Contract.Reallocate(&_BendVault.TransactOpts, allocations)
}

// Reallocate is a paid mutator transaction binding the contract method 0x7299aa31.
//
// Solidity: function reallocate(((address,address,address,address,uint256),uint256)[] allocations) returns()
func (_BendVault *BendVaultTransactorSession) Reallocate(allocations []MarketAllocation) (*types.Transaction, error) {
	return _BendVault.Contract.Reallocate(&_BendVault.TransactOpts, allocations)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_BendVault *BendVaultTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_BendVault *BendVaultSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.Redeem(&_BendVault.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_BendVault *BendVaultTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.Redeem(&_BendVault.TransactOpts, shares, receiver, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BendVault *BendVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BendVault *BendVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _BendVault.Contract.RenounceOwnership(&_BendVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BendVault *BendVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BendVault.Contract.RenounceOwnership(&_BendVault.TransactOpts)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_BendVault *BendVaultTransactor) RevokePendingCap(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "revokePendingCap", id)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_BendVault *BendVaultSession) RevokePendingCap(id [32]byte) (*types.Transaction, error) {
	return _BendVault.Contract.RevokePendingCap(&_BendVault.TransactOpts, id)
}

// RevokePendingCap is a paid mutator transaction binding the contract method 0x102f7b6c.
//
// Solidity: function revokePendingCap(bytes32 id) returns()
func (_BendVault *BendVaultTransactorSession) RevokePendingCap(id [32]byte) (*types.Transaction, error) {
	return _BendVault.Contract.RevokePendingCap(&_BendVault.TransactOpts, id)
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_BendVault *BendVaultTransactor) RevokePendingGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "revokePendingGuardian")
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_BendVault *BendVaultSession) RevokePendingGuardian() (*types.Transaction, error) {
	return _BendVault.Contract.RevokePendingGuardian(&_BendVault.TransactOpts)
}

// RevokePendingGuardian is a paid mutator transaction binding the contract method 0x1ecca77c.
//
// Solidity: function revokePendingGuardian() returns()
func (_BendVault *BendVaultTransactorSession) RevokePendingGuardian() (*types.Transaction, error) {
	return _BendVault.Contract.RevokePendingGuardian(&_BendVault.TransactOpts)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_BendVault *BendVaultTransactor) RevokePendingMarketRemoval(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "revokePendingMarketRemoval", id)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_BendVault *BendVaultSession) RevokePendingMarketRemoval(id [32]byte) (*types.Transaction, error) {
	return _BendVault.Contract.RevokePendingMarketRemoval(&_BendVault.TransactOpts, id)
}

// RevokePendingMarketRemoval is a paid mutator transaction binding the contract method 0x4b998de5.
//
// Solidity: function revokePendingMarketRemoval(bytes32 id) returns()
func (_BendVault *BendVaultTransactorSession) RevokePendingMarketRemoval(id [32]byte) (*types.Transaction, error) {
	return _BendVault.Contract.RevokePendingMarketRemoval(&_BendVault.TransactOpts, id)
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_BendVault *BendVaultTransactor) RevokePendingTimelock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "revokePendingTimelock")
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_BendVault *BendVaultSession) RevokePendingTimelock() (*types.Transaction, error) {
	return _BendVault.Contract.RevokePendingTimelock(&_BendVault.TransactOpts)
}

// RevokePendingTimelock is a paid mutator transaction binding the contract method 0xc9649aa9.
//
// Solidity: function revokePendingTimelock() returns()
func (_BendVault *BendVaultTransactorSession) RevokePendingTimelock() (*types.Transaction, error) {
	return _BendVault.Contract.RevokePendingTimelock(&_BendVault.TransactOpts)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_BendVault *BendVaultTransactor) SetCurator(opts *bind.TransactOpts, newCurator common.Address) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "setCurator", newCurator)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_BendVault *BendVaultSession) SetCurator(newCurator common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.SetCurator(&_BendVault.TransactOpts, newCurator)
}

// SetCurator is a paid mutator transaction binding the contract method 0xe90956cf.
//
// Solidity: function setCurator(address newCurator) returns()
func (_BendVault *BendVaultTransactorSession) SetCurator(newCurator common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.SetCurator(&_BendVault.TransactOpts, newCurator)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_BendVault *BendVaultTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_BendVault *BendVaultSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.SetFee(&_BendVault.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_BendVault *BendVaultTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.SetFee(&_BendVault.TransactOpts, newFee)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_BendVault *BendVaultTransactor) SetFeeRecipient(opts *bind.TransactOpts, newFeeRecipient common.Address) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "setFeeRecipient", newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_BendVault *BendVaultSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.SetFeeRecipient(&_BendVault.TransactOpts, newFeeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address newFeeRecipient) returns()
func (_BendVault *BendVaultTransactorSession) SetFeeRecipient(newFeeRecipient common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.SetFeeRecipient(&_BendVault.TransactOpts, newFeeRecipient)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_BendVault *BendVaultTransactor) SetIsAllocator(opts *bind.TransactOpts, newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "setIsAllocator", newAllocator, newIsAllocator)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_BendVault *BendVaultSession) SetIsAllocator(newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _BendVault.Contract.SetIsAllocator(&_BendVault.TransactOpts, newAllocator, newIsAllocator)
}

// SetIsAllocator is a paid mutator transaction binding the contract method 0xb192a84a.
//
// Solidity: function setIsAllocator(address newAllocator, bool newIsAllocator) returns()
func (_BendVault *BendVaultTransactorSession) SetIsAllocator(newAllocator common.Address, newIsAllocator bool) (*types.Transaction, error) {
	return _BendVault.Contract.SetIsAllocator(&_BendVault.TransactOpts, newAllocator, newIsAllocator)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_BendVault *BendVaultTransactor) SetName(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "setName", newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_BendVault *BendVaultSession) SetName(newName string) (*types.Transaction, error) {
	return _BendVault.Contract.SetName(&_BendVault.TransactOpts, newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_BendVault *BendVaultTransactorSession) SetName(newName string) (*types.Transaction, error) {
	return _BendVault.Contract.SetName(&_BendVault.TransactOpts, newName)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_BendVault *BendVaultTransactor) SetSkimRecipient(opts *bind.TransactOpts, newSkimRecipient common.Address) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "setSkimRecipient", newSkimRecipient)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_BendVault *BendVaultSession) SetSkimRecipient(newSkimRecipient common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.SetSkimRecipient(&_BendVault.TransactOpts, newSkimRecipient)
}

// SetSkimRecipient is a paid mutator transaction binding the contract method 0x2b30997b.
//
// Solidity: function setSkimRecipient(address newSkimRecipient) returns()
func (_BendVault *BendVaultTransactorSession) SetSkimRecipient(newSkimRecipient common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.SetSkimRecipient(&_BendVault.TransactOpts, newSkimRecipient)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_BendVault *BendVaultTransactor) SetSupplyQueue(opts *bind.TransactOpts, newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "setSupplyQueue", newSupplyQueue)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_BendVault *BendVaultSession) SetSupplyQueue(newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _BendVault.Contract.SetSupplyQueue(&_BendVault.TransactOpts, newSupplyQueue)
}

// SetSupplyQueue is a paid mutator transaction binding the contract method 0x2acc56f9.
//
// Solidity: function setSupplyQueue(bytes32[] newSupplyQueue) returns()
func (_BendVault *BendVaultTransactorSession) SetSupplyQueue(newSupplyQueue [][32]byte) (*types.Transaction, error) {
	return _BendVault.Contract.SetSupplyQueue(&_BendVault.TransactOpts, newSupplyQueue)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_BendVault *BendVaultTransactor) SetSymbol(opts *bind.TransactOpts, newSymbol string) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "setSymbol", newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_BendVault *BendVaultSession) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _BendVault.Contract.SetSymbol(&_BendVault.TransactOpts, newSymbol)
}

// SetSymbol is a paid mutator transaction binding the contract method 0xb84c8246.
//
// Solidity: function setSymbol(string newSymbol) returns()
func (_BendVault *BendVaultTransactorSession) SetSymbol(newSymbol string) (*types.Transaction, error) {
	return _BendVault.Contract.SetSymbol(&_BendVault.TransactOpts, newSymbol)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_BendVault *BendVaultTransactor) Skim(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "skim", token)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_BendVault *BendVaultSession) Skim(token common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.Skim(&_BendVault.TransactOpts, token)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address token) returns()
func (_BendVault *BendVaultTransactorSession) Skim(token common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.Skim(&_BendVault.TransactOpts, token)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_BendVault *BendVaultTransactor) SubmitCap(opts *bind.TransactOpts, marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "submitCap", marketParams, newSupplyCap)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_BendVault *BendVaultSession) SubmitCap(marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.SubmitCap(&_BendVault.TransactOpts, marketParams, newSupplyCap)
}

// SubmitCap is a paid mutator transaction binding the contract method 0x3b24c2bf.
//
// Solidity: function submitCap((address,address,address,address,uint256) marketParams, uint256 newSupplyCap) returns()
func (_BendVault *BendVaultTransactorSession) SubmitCap(marketParams MarketParams, newSupplyCap *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.SubmitCap(&_BendVault.TransactOpts, marketParams, newSupplyCap)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_BendVault *BendVaultTransactor) SubmitGuardian(opts *bind.TransactOpts, newGuardian common.Address) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "submitGuardian", newGuardian)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_BendVault *BendVaultSession) SubmitGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.SubmitGuardian(&_BendVault.TransactOpts, newGuardian)
}

// SubmitGuardian is a paid mutator transaction binding the contract method 0x9d6b4a45.
//
// Solidity: function submitGuardian(address newGuardian) returns()
func (_BendVault *BendVaultTransactorSession) SubmitGuardian(newGuardian common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.SubmitGuardian(&_BendVault.TransactOpts, newGuardian)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_BendVault *BendVaultTransactor) SubmitMarketRemoval(opts *bind.TransactOpts, marketParams MarketParams) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "submitMarketRemoval", marketParams)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_BendVault *BendVaultSession) SubmitMarketRemoval(marketParams MarketParams) (*types.Transaction, error) {
	return _BendVault.Contract.SubmitMarketRemoval(&_BendVault.TransactOpts, marketParams)
}

// SubmitMarketRemoval is a paid mutator transaction binding the contract method 0x84755b5f.
//
// Solidity: function submitMarketRemoval((address,address,address,address,uint256) marketParams) returns()
func (_BendVault *BendVaultTransactorSession) SubmitMarketRemoval(marketParams MarketParams) (*types.Transaction, error) {
	return _BendVault.Contract.SubmitMarketRemoval(&_BendVault.TransactOpts, marketParams)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_BendVault *BendVaultTransactor) SubmitTimelock(opts *bind.TransactOpts, newTimelock *big.Int) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "submitTimelock", newTimelock)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_BendVault *BendVaultSession) SubmitTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.SubmitTimelock(&_BendVault.TransactOpts, newTimelock)
}

// SubmitTimelock is a paid mutator transaction binding the contract method 0x7224a512.
//
// Solidity: function submitTimelock(uint256 newTimelock) returns()
func (_BendVault *BendVaultTransactorSession) SubmitTimelock(newTimelock *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.SubmitTimelock(&_BendVault.TransactOpts, newTimelock)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BendVault *BendVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BendVault *BendVaultSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.Transfer(&_BendVault.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BendVault *BendVaultTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.Transfer(&_BendVault.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BendVault *BendVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BendVault *BendVaultSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.TransferFrom(&_BendVault.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BendVault *BendVaultTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.TransferFrom(&_BendVault.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BendVault *BendVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BendVault *BendVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.TransferOwnership(&_BendVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BendVault *BendVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.TransferOwnership(&_BendVault.TransactOpts, newOwner)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_BendVault *BendVaultTransactor) UpdateWithdrawQueue(opts *bind.TransactOpts, indexes []*big.Int) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "updateWithdrawQueue", indexes)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_BendVault *BendVaultSession) UpdateWithdrawQueue(indexes []*big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.UpdateWithdrawQueue(&_BendVault.TransactOpts, indexes)
}

// UpdateWithdrawQueue is a paid mutator transaction binding the contract method 0x41b67833.
//
// Solidity: function updateWithdrawQueue(uint256[] indexes) returns()
func (_BendVault *BendVaultTransactorSession) UpdateWithdrawQueue(indexes []*big.Int) (*types.Transaction, error) {
	return _BendVault.Contract.UpdateWithdrawQueue(&_BendVault.TransactOpts, indexes)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_BendVault *BendVaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _BendVault.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_BendVault *BendVaultSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.Withdraw(&_BendVault.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_BendVault *BendVaultTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _BendVault.Contract.Withdraw(&_BendVault.TransactOpts, assets, receiver, owner)
}

// BendVaultAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the BendVault contract.
type BendVaultAccrueInterestIterator struct {
	Event *BendVaultAccrueInterest // Event containing the contract specifics and raw log

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
func (it *BendVaultAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultAccrueInterest)
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
		it.Event = new(BendVaultAccrueInterest)
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
func (it *BendVaultAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultAccrueInterest represents a AccrueInterest event raised by the BendVault contract.
type BendVaultAccrueInterest struct {
	NewTotalAssets *big.Int
	FeeShares      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_BendVault *BendVaultFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*BendVaultAccrueInterestIterator, error) {

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &BendVaultAccrueInterestIterator{contract: _BendVault.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_BendVault *BendVaultFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *BendVaultAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultAccrueInterest)
				if err := _BendVault.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0xf66f28b40975dbb933913542c7e6a0f50a1d0f20aa74ea6e0efe65ab616323ec.
//
// Solidity: event AccrueInterest(uint256 newTotalAssets, uint256 feeShares)
func (_BendVault *BendVaultFilterer) ParseAccrueInterest(log types.Log) (*BendVaultAccrueInterest, error) {
	event := new(BendVaultAccrueInterest)
	if err := _BendVault.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BendVault contract.
type BendVaultApprovalIterator struct {
	Event *BendVaultApproval // Event containing the contract specifics and raw log

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
func (it *BendVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultApproval)
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
		it.Event = new(BendVaultApproval)
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
func (it *BendVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultApproval represents a Approval event raised by the BendVault contract.
type BendVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BendVault *BendVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BendVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultApprovalIterator{contract: _BendVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BendVault *BendVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BendVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultApproval)
				if err := _BendVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BendVault *BendVaultFilterer) ParseApproval(log types.Log) (*BendVaultApproval, error) {
	event := new(BendVaultApproval)
	if err := _BendVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BendVault contract.
type BendVaultDepositIterator struct {
	Event *BendVaultDeposit // Event containing the contract specifics and raw log

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
func (it *BendVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultDeposit)
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
		it.Event = new(BendVaultDeposit)
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
func (it *BendVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultDeposit represents a Deposit event raised by the BendVault contract.
type BendVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_BendVault *BendVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*BendVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultDepositIterator{contract: _BendVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_BendVault *BendVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BendVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultDeposit)
				if err := _BendVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_BendVault *BendVaultFilterer) ParseDeposit(log types.Log) (*BendVaultDeposit, error) {
	event := new(BendVaultDeposit)
	if err := _BendVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the BendVault contract.
type BendVaultEIP712DomainChangedIterator struct {
	Event *BendVaultEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BendVaultEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultEIP712DomainChanged)
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
		it.Event = new(BendVaultEIP712DomainChanged)
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
func (it *BendVaultEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultEIP712DomainChanged represents a EIP712DomainChanged event raised by the BendVault contract.
type BendVaultEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BendVault *BendVaultFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BendVaultEIP712DomainChangedIterator, error) {

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BendVaultEIP712DomainChangedIterator{contract: _BendVault.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BendVault *BendVaultFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BendVaultEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultEIP712DomainChanged)
				if err := _BendVault.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_BendVault *BendVaultFilterer) ParseEIP712DomainChanged(log types.Log) (*BendVaultEIP712DomainChanged, error) {
	event := new(BendVaultEIP712DomainChanged)
	if err := _BendVault.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the BendVault contract.
type BendVaultOwnershipTransferStartedIterator struct {
	Event *BendVaultOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *BendVaultOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultOwnershipTransferStarted)
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
		it.Event = new(BendVaultOwnershipTransferStarted)
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
func (it *BendVaultOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the BendVault contract.
type BendVaultOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BendVault *BendVaultFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BendVaultOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultOwnershipTransferStartedIterator{contract: _BendVault.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_BendVault *BendVaultFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *BendVaultOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultOwnershipTransferStarted)
				if err := _BendVault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_BendVault *BendVaultFilterer) ParseOwnershipTransferStarted(log types.Log) (*BendVaultOwnershipTransferStarted, error) {
	event := new(BendVaultOwnershipTransferStarted)
	if err := _BendVault.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BendVault contract.
type BendVaultOwnershipTransferredIterator struct {
	Event *BendVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BendVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultOwnershipTransferred)
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
		it.Event = new(BendVaultOwnershipTransferred)
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
func (it *BendVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultOwnershipTransferred represents a OwnershipTransferred event raised by the BendVault contract.
type BendVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BendVault *BendVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BendVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultOwnershipTransferredIterator{contract: _BendVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BendVault *BendVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BendVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultOwnershipTransferred)
				if err := _BendVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BendVault *BendVaultFilterer) ParseOwnershipTransferred(log types.Log) (*BendVaultOwnershipTransferred, error) {
	event := new(BendVaultOwnershipTransferred)
	if err := _BendVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultReallocateSupplyIterator is returned from FilterReallocateSupply and is used to iterate over the raw logs and unpacked data for ReallocateSupply events raised by the BendVault contract.
type BendVaultReallocateSupplyIterator struct {
	Event *BendVaultReallocateSupply // Event containing the contract specifics and raw log

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
func (it *BendVaultReallocateSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultReallocateSupply)
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
		it.Event = new(BendVaultReallocateSupply)
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
func (it *BendVaultReallocateSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultReallocateSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultReallocateSupply represents a ReallocateSupply event raised by the BendVault contract.
type BendVaultReallocateSupply struct {
	Caller         common.Address
	Id             [32]byte
	SuppliedAssets *big.Int
	SuppliedShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReallocateSupply is a free log retrieval operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_BendVault *BendVaultFilterer) FilterReallocateSupply(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*BendVaultReallocateSupplyIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "ReallocateSupply", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultReallocateSupplyIterator{contract: _BendVault.contract, event: "ReallocateSupply", logs: logs, sub: sub}, nil
}

// WatchReallocateSupply is a free log subscription operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_BendVault *BendVaultFilterer) WatchReallocateSupply(opts *bind.WatchOpts, sink chan<- *BendVaultReallocateSupply, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "ReallocateSupply", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultReallocateSupply)
				if err := _BendVault.contract.UnpackLog(event, "ReallocateSupply", log); err != nil {
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

// ParseReallocateSupply is a log parse operation binding the contract event 0x89bf199df65bf65155e3e0a8abc4ad4a1be606220c8295840dba2ab5656c1f6d.
//
// Solidity: event ReallocateSupply(address indexed caller, bytes32 indexed id, uint256 suppliedAssets, uint256 suppliedShares)
func (_BendVault *BendVaultFilterer) ParseReallocateSupply(log types.Log) (*BendVaultReallocateSupply, error) {
	event := new(BendVaultReallocateSupply)
	if err := _BendVault.contract.UnpackLog(event, "ReallocateSupply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultReallocateWithdrawIterator is returned from FilterReallocateWithdraw and is used to iterate over the raw logs and unpacked data for ReallocateWithdraw events raised by the BendVault contract.
type BendVaultReallocateWithdrawIterator struct {
	Event *BendVaultReallocateWithdraw // Event containing the contract specifics and raw log

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
func (it *BendVaultReallocateWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultReallocateWithdraw)
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
		it.Event = new(BendVaultReallocateWithdraw)
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
func (it *BendVaultReallocateWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultReallocateWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultReallocateWithdraw represents a ReallocateWithdraw event raised by the BendVault contract.
type BendVaultReallocateWithdraw struct {
	Caller          common.Address
	Id              [32]byte
	WithdrawnAssets *big.Int
	WithdrawnShares *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReallocateWithdraw is a free log retrieval operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_BendVault *BendVaultFilterer) FilterReallocateWithdraw(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*BendVaultReallocateWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "ReallocateWithdraw", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultReallocateWithdrawIterator{contract: _BendVault.contract, event: "ReallocateWithdraw", logs: logs, sub: sub}, nil
}

// WatchReallocateWithdraw is a free log subscription operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_BendVault *BendVaultFilterer) WatchReallocateWithdraw(opts *bind.WatchOpts, sink chan<- *BendVaultReallocateWithdraw, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "ReallocateWithdraw", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultReallocateWithdraw)
				if err := _BendVault.contract.UnpackLog(event, "ReallocateWithdraw", log); err != nil {
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

// ParseReallocateWithdraw is a log parse operation binding the contract event 0xdd8bf5226dff861316e0fa7863fdb7dc7b87c614eb29a135f524eb79d5a1189a.
//
// Solidity: event ReallocateWithdraw(address indexed caller, bytes32 indexed id, uint256 withdrawnAssets, uint256 withdrawnShares)
func (_BendVault *BendVaultFilterer) ParseReallocateWithdraw(log types.Log) (*BendVaultReallocateWithdraw, error) {
	event := new(BendVaultReallocateWithdraw)
	if err := _BendVault.contract.UnpackLog(event, "ReallocateWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultRevokePendingCapIterator is returned from FilterRevokePendingCap and is used to iterate over the raw logs and unpacked data for RevokePendingCap events raised by the BendVault contract.
type BendVaultRevokePendingCapIterator struct {
	Event *BendVaultRevokePendingCap // Event containing the contract specifics and raw log

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
func (it *BendVaultRevokePendingCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultRevokePendingCap)
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
		it.Event = new(BendVaultRevokePendingCap)
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
func (it *BendVaultRevokePendingCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultRevokePendingCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultRevokePendingCap represents a RevokePendingCap event raised by the BendVault contract.
type BendVaultRevokePendingCap struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingCap is a free log retrieval operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_BendVault *BendVaultFilterer) FilterRevokePendingCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*BendVaultRevokePendingCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "RevokePendingCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultRevokePendingCapIterator{contract: _BendVault.contract, event: "RevokePendingCap", logs: logs, sub: sub}, nil
}

// WatchRevokePendingCap is a free log subscription operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_BendVault *BendVaultFilterer) WatchRevokePendingCap(opts *bind.WatchOpts, sink chan<- *BendVaultRevokePendingCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "RevokePendingCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultRevokePendingCap)
				if err := _BendVault.contract.UnpackLog(event, "RevokePendingCap", log); err != nil {
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

// ParseRevokePendingCap is a log parse operation binding the contract event 0x1026ceca5ed3747eb5edec555732d4a6f901ce1a875ecf981064628cadde1120.
//
// Solidity: event RevokePendingCap(address indexed caller, bytes32 indexed id)
func (_BendVault *BendVaultFilterer) ParseRevokePendingCap(log types.Log) (*BendVaultRevokePendingCap, error) {
	event := new(BendVaultRevokePendingCap)
	if err := _BendVault.contract.UnpackLog(event, "RevokePendingCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultRevokePendingGuardianIterator is returned from FilterRevokePendingGuardian and is used to iterate over the raw logs and unpacked data for RevokePendingGuardian events raised by the BendVault contract.
type BendVaultRevokePendingGuardianIterator struct {
	Event *BendVaultRevokePendingGuardian // Event containing the contract specifics and raw log

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
func (it *BendVaultRevokePendingGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultRevokePendingGuardian)
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
		it.Event = new(BendVaultRevokePendingGuardian)
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
func (it *BendVaultRevokePendingGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultRevokePendingGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultRevokePendingGuardian represents a RevokePendingGuardian event raised by the BendVault contract.
type BendVaultRevokePendingGuardian struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingGuardian is a free log retrieval operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_BendVault *BendVaultFilterer) FilterRevokePendingGuardian(opts *bind.FilterOpts, caller []common.Address) (*BendVaultRevokePendingGuardianIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "RevokePendingGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultRevokePendingGuardianIterator{contract: _BendVault.contract, event: "RevokePendingGuardian", logs: logs, sub: sub}, nil
}

// WatchRevokePendingGuardian is a free log subscription operation binding the contract event 0xc40a085ccfa20f5fd518ade5c3a77a7ecbdfbb4c75efcdca6146a8e3c841d663.
//
// Solidity: event RevokePendingGuardian(address indexed caller)
func (_BendVault *BendVaultFilterer) WatchRevokePendingGuardian(opts *bind.WatchOpts, sink chan<- *BendVaultRevokePendingGuardian, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "RevokePendingGuardian", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultRevokePendingGuardian)
				if err := _BendVault.contract.UnpackLog(event, "RevokePendingGuardian", log); err != nil {
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
func (_BendVault *BendVaultFilterer) ParseRevokePendingGuardian(log types.Log) (*BendVaultRevokePendingGuardian, error) {
	event := new(BendVaultRevokePendingGuardian)
	if err := _BendVault.contract.UnpackLog(event, "RevokePendingGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultRevokePendingMarketRemovalIterator is returned from FilterRevokePendingMarketRemoval and is used to iterate over the raw logs and unpacked data for RevokePendingMarketRemoval events raised by the BendVault contract.
type BendVaultRevokePendingMarketRemovalIterator struct {
	Event *BendVaultRevokePendingMarketRemoval // Event containing the contract specifics and raw log

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
func (it *BendVaultRevokePendingMarketRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultRevokePendingMarketRemoval)
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
		it.Event = new(BendVaultRevokePendingMarketRemoval)
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
func (it *BendVaultRevokePendingMarketRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultRevokePendingMarketRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultRevokePendingMarketRemoval represents a RevokePendingMarketRemoval event raised by the BendVault contract.
type BendVaultRevokePendingMarketRemoval struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingMarketRemoval is a free log retrieval operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_BendVault *BendVaultFilterer) FilterRevokePendingMarketRemoval(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*BendVaultRevokePendingMarketRemovalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "RevokePendingMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultRevokePendingMarketRemovalIterator{contract: _BendVault.contract, event: "RevokePendingMarketRemoval", logs: logs, sub: sub}, nil
}

// WatchRevokePendingMarketRemoval is a free log subscription operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_BendVault *BendVaultFilterer) WatchRevokePendingMarketRemoval(opts *bind.WatchOpts, sink chan<- *BendVaultRevokePendingMarketRemoval, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "RevokePendingMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultRevokePendingMarketRemoval)
				if err := _BendVault.contract.UnpackLog(event, "RevokePendingMarketRemoval", log); err != nil {
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

// ParseRevokePendingMarketRemoval is a log parse operation binding the contract event 0xcbeb8ecdaa5a3c133e62219b63bfc35bce3fda13065d2bed32e3b7dde60a59f4.
//
// Solidity: event RevokePendingMarketRemoval(address indexed caller, bytes32 indexed id)
func (_BendVault *BendVaultFilterer) ParseRevokePendingMarketRemoval(log types.Log) (*BendVaultRevokePendingMarketRemoval, error) {
	event := new(BendVaultRevokePendingMarketRemoval)
	if err := _BendVault.contract.UnpackLog(event, "RevokePendingMarketRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultRevokePendingTimelockIterator is returned from FilterRevokePendingTimelock and is used to iterate over the raw logs and unpacked data for RevokePendingTimelock events raised by the BendVault contract.
type BendVaultRevokePendingTimelockIterator struct {
	Event *BendVaultRevokePendingTimelock // Event containing the contract specifics and raw log

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
func (it *BendVaultRevokePendingTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultRevokePendingTimelock)
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
		it.Event = new(BendVaultRevokePendingTimelock)
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
func (it *BendVaultRevokePendingTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultRevokePendingTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultRevokePendingTimelock represents a RevokePendingTimelock event raised by the BendVault contract.
type BendVaultRevokePendingTimelock struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRevokePendingTimelock is a free log retrieval operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_BendVault *BendVaultFilterer) FilterRevokePendingTimelock(opts *bind.FilterOpts, caller []common.Address) (*BendVaultRevokePendingTimelockIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "RevokePendingTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultRevokePendingTimelockIterator{contract: _BendVault.contract, event: "RevokePendingTimelock", logs: logs, sub: sub}, nil
}

// WatchRevokePendingTimelock is a free log subscription operation binding the contract event 0x921828337692c347c634c5d2aacbc7b756014674bd236f3cc2058d8e284a951b.
//
// Solidity: event RevokePendingTimelock(address indexed caller)
func (_BendVault *BendVaultFilterer) WatchRevokePendingTimelock(opts *bind.WatchOpts, sink chan<- *BendVaultRevokePendingTimelock, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "RevokePendingTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultRevokePendingTimelock)
				if err := _BendVault.contract.UnpackLog(event, "RevokePendingTimelock", log); err != nil {
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
func (_BendVault *BendVaultFilterer) ParseRevokePendingTimelock(log types.Log) (*BendVaultRevokePendingTimelock, error) {
	event := new(BendVaultRevokePendingTimelock)
	if err := _BendVault.contract.UnpackLog(event, "RevokePendingTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetCapIterator is returned from FilterSetCap and is used to iterate over the raw logs and unpacked data for SetCap events raised by the BendVault contract.
type BendVaultSetCapIterator struct {
	Event *BendVaultSetCap // Event containing the contract specifics and raw log

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
func (it *BendVaultSetCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetCap)
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
		it.Event = new(BendVaultSetCap)
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
func (it *BendVaultSetCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetCap represents a SetCap event raised by the BendVault contract.
type BendVaultSetCap struct {
	Caller common.Address
	Id     [32]byte
	Cap    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetCap is a free log retrieval operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_BendVault *BendVaultFilterer) FilterSetCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*BendVaultSetCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSetCapIterator{contract: _BendVault.contract, event: "SetCap", logs: logs, sub: sub}, nil
}

// WatchSetCap is a free log subscription operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_BendVault *BendVaultFilterer) WatchSetCap(opts *bind.WatchOpts, sink chan<- *BendVaultSetCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetCap)
				if err := _BendVault.contract.UnpackLog(event, "SetCap", log); err != nil {
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

// ParseSetCap is a log parse operation binding the contract event 0xe86b6d3313d3098f4c5f689c935de8fde876a597c185def2cedab85efedac686.
//
// Solidity: event SetCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_BendVault *BendVaultFilterer) ParseSetCap(log types.Log) (*BendVaultSetCap, error) {
	event := new(BendVaultSetCap)
	if err := _BendVault.contract.UnpackLog(event, "SetCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetCuratorIterator is returned from FilterSetCurator and is used to iterate over the raw logs and unpacked data for SetCurator events raised by the BendVault contract.
type BendVaultSetCuratorIterator struct {
	Event *BendVaultSetCurator // Event containing the contract specifics and raw log

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
func (it *BendVaultSetCuratorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetCurator)
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
		it.Event = new(BendVaultSetCurator)
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
func (it *BendVaultSetCuratorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetCuratorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetCurator represents a SetCurator event raised by the BendVault contract.
type BendVaultSetCurator struct {
	NewCurator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetCurator is a free log retrieval operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_BendVault *BendVaultFilterer) FilterSetCurator(opts *bind.FilterOpts, newCurator []common.Address) (*BendVaultSetCuratorIterator, error) {

	var newCuratorRule []interface{}
	for _, newCuratorItem := range newCurator {
		newCuratorRule = append(newCuratorRule, newCuratorItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetCurator", newCuratorRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSetCuratorIterator{contract: _BendVault.contract, event: "SetCurator", logs: logs, sub: sub}, nil
}

// WatchSetCurator is a free log subscription operation binding the contract event 0xbd0a63c12948fbc9194a5839019f99c9d71db924e5c70018265bc778b8f1a506.
//
// Solidity: event SetCurator(address indexed newCurator)
func (_BendVault *BendVaultFilterer) WatchSetCurator(opts *bind.WatchOpts, sink chan<- *BendVaultSetCurator, newCurator []common.Address) (event.Subscription, error) {

	var newCuratorRule []interface{}
	for _, newCuratorItem := range newCurator {
		newCuratorRule = append(newCuratorRule, newCuratorItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetCurator", newCuratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetCurator)
				if err := _BendVault.contract.UnpackLog(event, "SetCurator", log); err != nil {
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
// Solidity: event SetCurator(address indexed newCurator)
func (_BendVault *BendVaultFilterer) ParseSetCurator(log types.Log) (*BendVaultSetCurator, error) {
	event := new(BendVaultSetCurator)
	if err := _BendVault.contract.UnpackLog(event, "SetCurator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the BendVault contract.
type BendVaultSetFeeIterator struct {
	Event *BendVaultSetFee // Event containing the contract specifics and raw log

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
func (it *BendVaultSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetFee)
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
		it.Event = new(BendVaultSetFee)
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
func (it *BendVaultSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetFee represents a SetFee event raised by the BendVault contract.
type BendVaultSetFee struct {
	Caller common.Address
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_BendVault *BendVaultFilterer) FilterSetFee(opts *bind.FilterOpts, caller []common.Address) (*BendVaultSetFeeIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetFee", callerRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSetFeeIterator{contract: _BendVault.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_BendVault *BendVaultFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *BendVaultSetFee, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetFee", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetFee)
				if err := _BendVault.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0x01fe2943baee27f47add82886c2200f910c749c461c9b63c5fe83901a53bdb49.
//
// Solidity: event SetFee(address indexed caller, uint256 newFee)
func (_BendVault *BendVaultFilterer) ParseSetFee(log types.Log) (*BendVaultSetFee, error) {
	event := new(BendVaultSetFee)
	if err := _BendVault.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetFeeRecipientIterator is returned from FilterSetFeeRecipient and is used to iterate over the raw logs and unpacked data for SetFeeRecipient events raised by the BendVault contract.
type BendVaultSetFeeRecipientIterator struct {
	Event *BendVaultSetFeeRecipient // Event containing the contract specifics and raw log

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
func (it *BendVaultSetFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetFeeRecipient)
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
		it.Event = new(BendVaultSetFeeRecipient)
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
func (it *BendVaultSetFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetFeeRecipient represents a SetFeeRecipient event raised by the BendVault contract.
type BendVaultSetFeeRecipient struct {
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeRecipient is a free log retrieval operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_BendVault *BendVaultFilterer) FilterSetFeeRecipient(opts *bind.FilterOpts, newFeeRecipient []common.Address) (*BendVaultSetFeeRecipientIterator, error) {

	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetFeeRecipient", newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSetFeeRecipientIterator{contract: _BendVault.contract, event: "SetFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchSetFeeRecipient is a free log subscription operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_BendVault *BendVaultFilterer) WatchSetFeeRecipient(opts *bind.WatchOpts, sink chan<- *BendVaultSetFeeRecipient, newFeeRecipient []common.Address) (event.Subscription, error) {

	var newFeeRecipientRule []interface{}
	for _, newFeeRecipientItem := range newFeeRecipient {
		newFeeRecipientRule = append(newFeeRecipientRule, newFeeRecipientItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetFeeRecipient", newFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetFeeRecipient)
				if err := _BendVault.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
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

// ParseSetFeeRecipient is a log parse operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address indexed newFeeRecipient)
func (_BendVault *BendVaultFilterer) ParseSetFeeRecipient(log types.Log) (*BendVaultSetFeeRecipient, error) {
	event := new(BendVaultSetFeeRecipient)
	if err := _BendVault.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetGuardianIterator is returned from FilterSetGuardian and is used to iterate over the raw logs and unpacked data for SetGuardian events raised by the BendVault contract.
type BendVaultSetGuardianIterator struct {
	Event *BendVaultSetGuardian // Event containing the contract specifics and raw log

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
func (it *BendVaultSetGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetGuardian)
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
		it.Event = new(BendVaultSetGuardian)
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
func (it *BendVaultSetGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetGuardian represents a SetGuardian event raised by the BendVault contract.
type BendVaultSetGuardian struct {
	Caller   common.Address
	Guardian common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetGuardian is a free log retrieval operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_BendVault *BendVaultFilterer) FilterSetGuardian(opts *bind.FilterOpts, caller []common.Address, guardian []common.Address) (*BendVaultSetGuardianIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetGuardian", callerRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSetGuardianIterator{contract: _BendVault.contract, event: "SetGuardian", logs: logs, sub: sub}, nil
}

// WatchSetGuardian is a free log subscription operation binding the contract event 0xcb11cc8aade2f5a556749d1b2380d108a16fac3431e6a5d5ce12ef9de0bd76e3.
//
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_BendVault *BendVaultFilterer) WatchSetGuardian(opts *bind.WatchOpts, sink chan<- *BendVaultSetGuardian, caller []common.Address, guardian []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetGuardian", callerRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetGuardian)
				if err := _BendVault.contract.UnpackLog(event, "SetGuardian", log); err != nil {
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
// Solidity: event SetGuardian(address indexed caller, address indexed guardian)
func (_BendVault *BendVaultFilterer) ParseSetGuardian(log types.Log) (*BendVaultSetGuardian, error) {
	event := new(BendVaultSetGuardian)
	if err := _BendVault.contract.UnpackLog(event, "SetGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetIsAllocatorIterator is returned from FilterSetIsAllocator and is used to iterate over the raw logs and unpacked data for SetIsAllocator events raised by the BendVault contract.
type BendVaultSetIsAllocatorIterator struct {
	Event *BendVaultSetIsAllocator // Event containing the contract specifics and raw log

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
func (it *BendVaultSetIsAllocatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetIsAllocator)
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
		it.Event = new(BendVaultSetIsAllocator)
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
func (it *BendVaultSetIsAllocatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetIsAllocatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetIsAllocator represents a SetIsAllocator event raised by the BendVault contract.
type BendVaultSetIsAllocator struct {
	Allocator   common.Address
	IsAllocator bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetIsAllocator is a free log retrieval operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_BendVault *BendVaultFilterer) FilterSetIsAllocator(opts *bind.FilterOpts, allocator []common.Address) (*BendVaultSetIsAllocatorIterator, error) {

	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetIsAllocator", allocatorRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSetIsAllocatorIterator{contract: _BendVault.contract, event: "SetIsAllocator", logs: logs, sub: sub}, nil
}

// WatchSetIsAllocator is a free log subscription operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_BendVault *BendVaultFilterer) WatchSetIsAllocator(opts *bind.WatchOpts, sink chan<- *BendVaultSetIsAllocator, allocator []common.Address) (event.Subscription, error) {

	var allocatorRule []interface{}
	for _, allocatorItem := range allocator {
		allocatorRule = append(allocatorRule, allocatorItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetIsAllocator", allocatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetIsAllocator)
				if err := _BendVault.contract.UnpackLog(event, "SetIsAllocator", log); err != nil {
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

// ParseSetIsAllocator is a log parse operation binding the contract event 0x74dc60cbc81a9472d04ad1d20e151d369c41104d655ed3f2f3091166a502cd8d.
//
// Solidity: event SetIsAllocator(address indexed allocator, bool isAllocator)
func (_BendVault *BendVaultFilterer) ParseSetIsAllocator(log types.Log) (*BendVaultSetIsAllocator, error) {
	event := new(BendVaultSetIsAllocator)
	if err := _BendVault.contract.UnpackLog(event, "SetIsAllocator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetNameIterator is returned from FilterSetName and is used to iterate over the raw logs and unpacked data for SetName events raised by the BendVault contract.
type BendVaultSetNameIterator struct {
	Event *BendVaultSetName // Event containing the contract specifics and raw log

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
func (it *BendVaultSetNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetName)
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
		it.Event = new(BendVaultSetName)
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
func (it *BendVaultSetNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetName represents a SetName event raised by the BendVault contract.
type BendVaultSetName struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetName is a free log retrieval operation binding the contract event 0x4df9dcd34ae35f40f2c756fd8ac83210ed0b76d065543ee73d868aec7c7fcf02.
//
// Solidity: event SetName(string name)
func (_BendVault *BendVaultFilterer) FilterSetName(opts *bind.FilterOpts) (*BendVaultSetNameIterator, error) {

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetName")
	if err != nil {
		return nil, err
	}
	return &BendVaultSetNameIterator{contract: _BendVault.contract, event: "SetName", logs: logs, sub: sub}, nil
}

// WatchSetName is a free log subscription operation binding the contract event 0x4df9dcd34ae35f40f2c756fd8ac83210ed0b76d065543ee73d868aec7c7fcf02.
//
// Solidity: event SetName(string name)
func (_BendVault *BendVaultFilterer) WatchSetName(opts *bind.WatchOpts, sink chan<- *BendVaultSetName) (event.Subscription, error) {

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetName)
				if err := _BendVault.contract.UnpackLog(event, "SetName", log); err != nil {
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

// ParseSetName is a log parse operation binding the contract event 0x4df9dcd34ae35f40f2c756fd8ac83210ed0b76d065543ee73d868aec7c7fcf02.
//
// Solidity: event SetName(string name)
func (_BendVault *BendVaultFilterer) ParseSetName(log types.Log) (*BendVaultSetName, error) {
	event := new(BendVaultSetName)
	if err := _BendVault.contract.UnpackLog(event, "SetName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetSkimRecipientIterator is returned from FilterSetSkimRecipient and is used to iterate over the raw logs and unpacked data for SetSkimRecipient events raised by the BendVault contract.
type BendVaultSetSkimRecipientIterator struct {
	Event *BendVaultSetSkimRecipient // Event containing the contract specifics and raw log

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
func (it *BendVaultSetSkimRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetSkimRecipient)
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
		it.Event = new(BendVaultSetSkimRecipient)
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
func (it *BendVaultSetSkimRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetSkimRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetSkimRecipient represents a SetSkimRecipient event raised by the BendVault contract.
type BendVaultSetSkimRecipient struct {
	NewSkimRecipient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetSkimRecipient is a free log retrieval operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_BendVault *BendVaultFilterer) FilterSetSkimRecipient(opts *bind.FilterOpts, newSkimRecipient []common.Address) (*BendVaultSetSkimRecipientIterator, error) {

	var newSkimRecipientRule []interface{}
	for _, newSkimRecipientItem := range newSkimRecipient {
		newSkimRecipientRule = append(newSkimRecipientRule, newSkimRecipientItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetSkimRecipient", newSkimRecipientRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSetSkimRecipientIterator{contract: _BendVault.contract, event: "SetSkimRecipient", logs: logs, sub: sub}, nil
}

// WatchSetSkimRecipient is a free log subscription operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_BendVault *BendVaultFilterer) WatchSetSkimRecipient(opts *bind.WatchOpts, sink chan<- *BendVaultSetSkimRecipient, newSkimRecipient []common.Address) (event.Subscription, error) {

	var newSkimRecipientRule []interface{}
	for _, newSkimRecipientItem := range newSkimRecipient {
		newSkimRecipientRule = append(newSkimRecipientRule, newSkimRecipientItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetSkimRecipient", newSkimRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetSkimRecipient)
				if err := _BendVault.contract.UnpackLog(event, "SetSkimRecipient", log); err != nil {
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

// ParseSetSkimRecipient is a log parse operation binding the contract event 0x2e7908865670e21b9779422cadf5f1cba271a62bb95c71eaaf615c0a1c48ebee.
//
// Solidity: event SetSkimRecipient(address indexed newSkimRecipient)
func (_BendVault *BendVaultFilterer) ParseSetSkimRecipient(log types.Log) (*BendVaultSetSkimRecipient, error) {
	event := new(BendVaultSetSkimRecipient)
	if err := _BendVault.contract.UnpackLog(event, "SetSkimRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetSupplyQueueIterator is returned from FilterSetSupplyQueue and is used to iterate over the raw logs and unpacked data for SetSupplyQueue events raised by the BendVault contract.
type BendVaultSetSupplyQueueIterator struct {
	Event *BendVaultSetSupplyQueue // Event containing the contract specifics and raw log

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
func (it *BendVaultSetSupplyQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetSupplyQueue)
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
		it.Event = new(BendVaultSetSupplyQueue)
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
func (it *BendVaultSetSupplyQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetSupplyQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetSupplyQueue represents a SetSupplyQueue event raised by the BendVault contract.
type BendVaultSetSupplyQueue struct {
	Caller         common.Address
	NewSupplyQueue [][32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetSupplyQueue is a free log retrieval operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_BendVault *BendVaultFilterer) FilterSetSupplyQueue(opts *bind.FilterOpts, caller []common.Address) (*BendVaultSetSupplyQueueIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetSupplyQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSetSupplyQueueIterator{contract: _BendVault.contract, event: "SetSupplyQueue", logs: logs, sub: sub}, nil
}

// WatchSetSupplyQueue is a free log subscription operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_BendVault *BendVaultFilterer) WatchSetSupplyQueue(opts *bind.WatchOpts, sink chan<- *BendVaultSetSupplyQueue, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetSupplyQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetSupplyQueue)
				if err := _BendVault.contract.UnpackLog(event, "SetSupplyQueue", log); err != nil {
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

// ParseSetSupplyQueue is a log parse operation binding the contract event 0x6ce31538fc7fba95714ddc8a275a09252b4b1fb8f33d2550aa58a5f62ad934de.
//
// Solidity: event SetSupplyQueue(address indexed caller, bytes32[] newSupplyQueue)
func (_BendVault *BendVaultFilterer) ParseSetSupplyQueue(log types.Log) (*BendVaultSetSupplyQueue, error) {
	event := new(BendVaultSetSupplyQueue)
	if err := _BendVault.contract.UnpackLog(event, "SetSupplyQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetSymbolIterator is returned from FilterSetSymbol and is used to iterate over the raw logs and unpacked data for SetSymbol events raised by the BendVault contract.
type BendVaultSetSymbolIterator struct {
	Event *BendVaultSetSymbol // Event containing the contract specifics and raw log

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
func (it *BendVaultSetSymbolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetSymbol)
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
		it.Event = new(BendVaultSetSymbol)
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
func (it *BendVaultSetSymbolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetSymbolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetSymbol represents a SetSymbol event raised by the BendVault contract.
type BendVaultSetSymbol struct {
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSymbol is a free log retrieval operation binding the contract event 0xadf3ae8bd543b3007d464f15cb8ea1db3f44e84d41d203164f40b95e27558ac6.
//
// Solidity: event SetSymbol(string symbol)
func (_BendVault *BendVaultFilterer) FilterSetSymbol(opts *bind.FilterOpts) (*BendVaultSetSymbolIterator, error) {

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetSymbol")
	if err != nil {
		return nil, err
	}
	return &BendVaultSetSymbolIterator{contract: _BendVault.contract, event: "SetSymbol", logs: logs, sub: sub}, nil
}

// WatchSetSymbol is a free log subscription operation binding the contract event 0xadf3ae8bd543b3007d464f15cb8ea1db3f44e84d41d203164f40b95e27558ac6.
//
// Solidity: event SetSymbol(string symbol)
func (_BendVault *BendVaultFilterer) WatchSetSymbol(opts *bind.WatchOpts, sink chan<- *BendVaultSetSymbol) (event.Subscription, error) {

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetSymbol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetSymbol)
				if err := _BendVault.contract.UnpackLog(event, "SetSymbol", log); err != nil {
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

// ParseSetSymbol is a log parse operation binding the contract event 0xadf3ae8bd543b3007d464f15cb8ea1db3f44e84d41d203164f40b95e27558ac6.
//
// Solidity: event SetSymbol(string symbol)
func (_BendVault *BendVaultFilterer) ParseSetSymbol(log types.Log) (*BendVaultSetSymbol, error) {
	event := new(BendVaultSetSymbol)
	if err := _BendVault.contract.UnpackLog(event, "SetSymbol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetTimelockIterator is returned from FilterSetTimelock and is used to iterate over the raw logs and unpacked data for SetTimelock events raised by the BendVault contract.
type BendVaultSetTimelockIterator struct {
	Event *BendVaultSetTimelock // Event containing the contract specifics and raw log

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
func (it *BendVaultSetTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetTimelock)
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
		it.Event = new(BendVaultSetTimelock)
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
func (it *BendVaultSetTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetTimelock represents a SetTimelock event raised by the BendVault contract.
type BendVaultSetTimelock struct {
	Caller      common.Address
	NewTimelock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetTimelock is a free log retrieval operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_BendVault *BendVaultFilterer) FilterSetTimelock(opts *bind.FilterOpts, caller []common.Address) (*BendVaultSetTimelockIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSetTimelockIterator{contract: _BendVault.contract, event: "SetTimelock", logs: logs, sub: sub}, nil
}

// WatchSetTimelock is a free log subscription operation binding the contract event 0xd28e9b90ee9b37c5936ff84392d71f29ff18117d7e76bcee60615262a90a3f75.
//
// Solidity: event SetTimelock(address indexed caller, uint256 newTimelock)
func (_BendVault *BendVaultFilterer) WatchSetTimelock(opts *bind.WatchOpts, sink chan<- *BendVaultSetTimelock, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetTimelock", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetTimelock)
				if err := _BendVault.contract.UnpackLog(event, "SetTimelock", log); err != nil {
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
func (_BendVault *BendVaultFilterer) ParseSetTimelock(log types.Log) (*BendVaultSetTimelock, error) {
	event := new(BendVaultSetTimelock)
	if err := _BendVault.contract.UnpackLog(event, "SetTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSetWithdrawQueueIterator is returned from FilterSetWithdrawQueue and is used to iterate over the raw logs and unpacked data for SetWithdrawQueue events raised by the BendVault contract.
type BendVaultSetWithdrawQueueIterator struct {
	Event *BendVaultSetWithdrawQueue // Event containing the contract specifics and raw log

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
func (it *BendVaultSetWithdrawQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSetWithdrawQueue)
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
		it.Event = new(BendVaultSetWithdrawQueue)
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
func (it *BendVaultSetWithdrawQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSetWithdrawQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSetWithdrawQueue represents a SetWithdrawQueue event raised by the BendVault contract.
type BendVaultSetWithdrawQueue struct {
	Caller           common.Address
	NewWithdrawQueue [][32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawQueue is a free log retrieval operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_BendVault *BendVaultFilterer) FilterSetWithdrawQueue(opts *bind.FilterOpts, caller []common.Address) (*BendVaultSetWithdrawQueueIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SetWithdrawQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSetWithdrawQueueIterator{contract: _BendVault.contract, event: "SetWithdrawQueue", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawQueue is a free log subscription operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_BendVault *BendVaultFilterer) WatchSetWithdrawQueue(opts *bind.WatchOpts, sink chan<- *BendVaultSetWithdrawQueue, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SetWithdrawQueue", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSetWithdrawQueue)
				if err := _BendVault.contract.UnpackLog(event, "SetWithdrawQueue", log); err != nil {
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

// ParseSetWithdrawQueue is a log parse operation binding the contract event 0xe0c2db6b54586be6d7d49943139fccf0dd315ba63e55364a76c73cd8fdba724d.
//
// Solidity: event SetWithdrawQueue(address indexed caller, bytes32[] newWithdrawQueue)
func (_BendVault *BendVaultFilterer) ParseSetWithdrawQueue(log types.Log) (*BendVaultSetWithdrawQueue, error) {
	event := new(BendVaultSetWithdrawQueue)
	if err := _BendVault.contract.UnpackLog(event, "SetWithdrawQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSkimIterator is returned from FilterSkim and is used to iterate over the raw logs and unpacked data for Skim events raised by the BendVault contract.
type BendVaultSkimIterator struct {
	Event *BendVaultSkim // Event containing the contract specifics and raw log

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
func (it *BendVaultSkimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSkim)
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
		it.Event = new(BendVaultSkim)
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
func (it *BendVaultSkimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSkimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSkim represents a Skim event raised by the BendVault contract.
type BendVaultSkim struct {
	Caller common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSkim is a free log retrieval operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_BendVault *BendVaultFilterer) FilterSkim(opts *bind.FilterOpts, caller []common.Address, token []common.Address) (*BendVaultSkimIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "Skim", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSkimIterator{contract: _BendVault.contract, event: "Skim", logs: logs, sub: sub}, nil
}

// WatchSkim is a free log subscription operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_BendVault *BendVaultFilterer) WatchSkim(opts *bind.WatchOpts, sink chan<- *BendVaultSkim, caller []common.Address, token []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "Skim", callerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSkim)
				if err := _BendVault.contract.UnpackLog(event, "Skim", log); err != nil {
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

// ParseSkim is a log parse operation binding the contract event 0x2ae72b44f59d038340fca5739135a1d51fc5ab720bb02d983e4c5ff4119ca7b8.
//
// Solidity: event Skim(address indexed caller, address indexed token, uint256 amount)
func (_BendVault *BendVaultFilterer) ParseSkim(log types.Log) (*BendVaultSkim, error) {
	event := new(BendVaultSkim)
	if err := _BendVault.contract.UnpackLog(event, "Skim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSubmitCapIterator is returned from FilterSubmitCap and is used to iterate over the raw logs and unpacked data for SubmitCap events raised by the BendVault contract.
type BendVaultSubmitCapIterator struct {
	Event *BendVaultSubmitCap // Event containing the contract specifics and raw log

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
func (it *BendVaultSubmitCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSubmitCap)
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
		it.Event = new(BendVaultSubmitCap)
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
func (it *BendVaultSubmitCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSubmitCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSubmitCap represents a SubmitCap event raised by the BendVault contract.
type BendVaultSubmitCap struct {
	Caller common.Address
	Id     [32]byte
	Cap    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitCap is a free log retrieval operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_BendVault *BendVaultFilterer) FilterSubmitCap(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*BendVaultSubmitCapIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SubmitCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSubmitCapIterator{contract: _BendVault.contract, event: "SubmitCap", logs: logs, sub: sub}, nil
}

// WatchSubmitCap is a free log subscription operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_BendVault *BendVaultFilterer) WatchSubmitCap(opts *bind.WatchOpts, sink chan<- *BendVaultSubmitCap, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SubmitCap", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSubmitCap)
				if err := _BendVault.contract.UnpackLog(event, "SubmitCap", log); err != nil {
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

// ParseSubmitCap is a log parse operation binding the contract event 0xe851bb5856808a50efd748be463b8f35bcfb5ec74c5bfde776fe0a4d2a26db27.
//
// Solidity: event SubmitCap(address indexed caller, bytes32 indexed id, uint256 cap)
func (_BendVault *BendVaultFilterer) ParseSubmitCap(log types.Log) (*BendVaultSubmitCap, error) {
	event := new(BendVaultSubmitCap)
	if err := _BendVault.contract.UnpackLog(event, "SubmitCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSubmitGuardianIterator is returned from FilterSubmitGuardian and is used to iterate over the raw logs and unpacked data for SubmitGuardian events raised by the BendVault contract.
type BendVaultSubmitGuardianIterator struct {
	Event *BendVaultSubmitGuardian // Event containing the contract specifics and raw log

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
func (it *BendVaultSubmitGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSubmitGuardian)
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
		it.Event = new(BendVaultSubmitGuardian)
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
func (it *BendVaultSubmitGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSubmitGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSubmitGuardian represents a SubmitGuardian event raised by the BendVault contract.
type BendVaultSubmitGuardian struct {
	NewGuardian common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitGuardian is a free log retrieval operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_BendVault *BendVaultFilterer) FilterSubmitGuardian(opts *bind.FilterOpts, newGuardian []common.Address) (*BendVaultSubmitGuardianIterator, error) {

	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SubmitGuardian", newGuardianRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSubmitGuardianIterator{contract: _BendVault.contract, event: "SubmitGuardian", logs: logs, sub: sub}, nil
}

// WatchSubmitGuardian is a free log subscription operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_BendVault *BendVaultFilterer) WatchSubmitGuardian(opts *bind.WatchOpts, sink chan<- *BendVaultSubmitGuardian, newGuardian []common.Address) (event.Subscription, error) {

	var newGuardianRule []interface{}
	for _, newGuardianItem := range newGuardian {
		newGuardianRule = append(newGuardianRule, newGuardianItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SubmitGuardian", newGuardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSubmitGuardian)
				if err := _BendVault.contract.UnpackLog(event, "SubmitGuardian", log); err != nil {
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

// ParseSubmitGuardian is a log parse operation binding the contract event 0x7633313af54753bce8a149927263b1a55eba857ba4ef1d13c6aee25d384d3c4b.
//
// Solidity: event SubmitGuardian(address indexed newGuardian)
func (_BendVault *BendVaultFilterer) ParseSubmitGuardian(log types.Log) (*BendVaultSubmitGuardian, error) {
	event := new(BendVaultSubmitGuardian)
	if err := _BendVault.contract.UnpackLog(event, "SubmitGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSubmitMarketRemovalIterator is returned from FilterSubmitMarketRemoval and is used to iterate over the raw logs and unpacked data for SubmitMarketRemoval events raised by the BendVault contract.
type BendVaultSubmitMarketRemovalIterator struct {
	Event *BendVaultSubmitMarketRemoval // Event containing the contract specifics and raw log

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
func (it *BendVaultSubmitMarketRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSubmitMarketRemoval)
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
		it.Event = new(BendVaultSubmitMarketRemoval)
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
func (it *BendVaultSubmitMarketRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSubmitMarketRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSubmitMarketRemoval represents a SubmitMarketRemoval event raised by the BendVault contract.
type BendVaultSubmitMarketRemoval struct {
	Caller common.Address
	Id     [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubmitMarketRemoval is a free log retrieval operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_BendVault *BendVaultFilterer) FilterSubmitMarketRemoval(opts *bind.FilterOpts, caller []common.Address, id [][32]byte) (*BendVaultSubmitMarketRemovalIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SubmitMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultSubmitMarketRemovalIterator{contract: _BendVault.contract, event: "SubmitMarketRemoval", logs: logs, sub: sub}, nil
}

// WatchSubmitMarketRemoval is a free log subscription operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_BendVault *BendVaultFilterer) WatchSubmitMarketRemoval(opts *bind.WatchOpts, sink chan<- *BendVaultSubmitMarketRemoval, caller []common.Address, id [][32]byte) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SubmitMarketRemoval", callerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSubmitMarketRemoval)
				if err := _BendVault.contract.UnpackLog(event, "SubmitMarketRemoval", log); err != nil {
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

// ParseSubmitMarketRemoval is a log parse operation binding the contract event 0x3240fc70754c5a2b4dab10bf7081a00024bfc8491581ee3d355360ec0dd91f16.
//
// Solidity: event SubmitMarketRemoval(address indexed caller, bytes32 indexed id)
func (_BendVault *BendVaultFilterer) ParseSubmitMarketRemoval(log types.Log) (*BendVaultSubmitMarketRemoval, error) {
	event := new(BendVaultSubmitMarketRemoval)
	if err := _BendVault.contract.UnpackLog(event, "SubmitMarketRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultSubmitTimelockIterator is returned from FilterSubmitTimelock and is used to iterate over the raw logs and unpacked data for SubmitTimelock events raised by the BendVault contract.
type BendVaultSubmitTimelockIterator struct {
	Event *BendVaultSubmitTimelock // Event containing the contract specifics and raw log

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
func (it *BendVaultSubmitTimelockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultSubmitTimelock)
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
		it.Event = new(BendVaultSubmitTimelock)
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
func (it *BendVaultSubmitTimelockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultSubmitTimelockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultSubmitTimelock represents a SubmitTimelock event raised by the BendVault contract.
type BendVaultSubmitTimelock struct {
	NewTimelock *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmitTimelock is a free log retrieval operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_BendVault *BendVaultFilterer) FilterSubmitTimelock(opts *bind.FilterOpts) (*BendVaultSubmitTimelockIterator, error) {

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "SubmitTimelock")
	if err != nil {
		return nil, err
	}
	return &BendVaultSubmitTimelockIterator{contract: _BendVault.contract, event: "SubmitTimelock", logs: logs, sub: sub}, nil
}

// WatchSubmitTimelock is a free log subscription operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_BendVault *BendVaultFilterer) WatchSubmitTimelock(opts *bind.WatchOpts, sink chan<- *BendVaultSubmitTimelock) (event.Subscription, error) {

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "SubmitTimelock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultSubmitTimelock)
				if err := _BendVault.contract.UnpackLog(event, "SubmitTimelock", log); err != nil {
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

// ParseSubmitTimelock is a log parse operation binding the contract event 0xb3aa0ade2442acf51d06713c2d1a5a3ec0373cce969d42b53f4689f97bccf380.
//
// Solidity: event SubmitTimelock(uint256 newTimelock)
func (_BendVault *BendVaultFilterer) ParseSubmitTimelock(log types.Log) (*BendVaultSubmitTimelock, error) {
	event := new(BendVaultSubmitTimelock)
	if err := _BendVault.contract.UnpackLog(event, "SubmitTimelock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BendVault contract.
type BendVaultTransferIterator struct {
	Event *BendVaultTransfer // Event containing the contract specifics and raw log

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
func (it *BendVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultTransfer)
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
		it.Event = new(BendVaultTransfer)
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
func (it *BendVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultTransfer represents a Transfer event raised by the BendVault contract.
type BendVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BendVault *BendVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BendVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultTransferIterator{contract: _BendVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BendVault *BendVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BendVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultTransfer)
				if err := _BendVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BendVault *BendVaultFilterer) ParseTransfer(log types.Log) (*BendVaultTransfer, error) {
	event := new(BendVaultTransfer)
	if err := _BendVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultUpdateLastTotalAssetsIterator is returned from FilterUpdateLastTotalAssets and is used to iterate over the raw logs and unpacked data for UpdateLastTotalAssets events raised by the BendVault contract.
type BendVaultUpdateLastTotalAssetsIterator struct {
	Event *BendVaultUpdateLastTotalAssets // Event containing the contract specifics and raw log

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
func (it *BendVaultUpdateLastTotalAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultUpdateLastTotalAssets)
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
		it.Event = new(BendVaultUpdateLastTotalAssets)
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
func (it *BendVaultUpdateLastTotalAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultUpdateLastTotalAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultUpdateLastTotalAssets represents a UpdateLastTotalAssets event raised by the BendVault contract.
type BendVaultUpdateLastTotalAssets struct {
	UpdatedTotalAssets *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateLastTotalAssets is a free log retrieval operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_BendVault *BendVaultFilterer) FilterUpdateLastTotalAssets(opts *bind.FilterOpts) (*BendVaultUpdateLastTotalAssetsIterator, error) {

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "UpdateLastTotalAssets")
	if err != nil {
		return nil, err
	}
	return &BendVaultUpdateLastTotalAssetsIterator{contract: _BendVault.contract, event: "UpdateLastTotalAssets", logs: logs, sub: sub}, nil
}

// WatchUpdateLastTotalAssets is a free log subscription operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_BendVault *BendVaultFilterer) WatchUpdateLastTotalAssets(opts *bind.WatchOpts, sink chan<- *BendVaultUpdateLastTotalAssets) (event.Subscription, error) {

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "UpdateLastTotalAssets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultUpdateLastTotalAssets)
				if err := _BendVault.contract.UnpackLog(event, "UpdateLastTotalAssets", log); err != nil {
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

// ParseUpdateLastTotalAssets is a log parse operation binding the contract event 0x15c027cc4fd826d986cad358803439f7326d3aa4ed969ff90dbee4bc150f68e9.
//
// Solidity: event UpdateLastTotalAssets(uint256 updatedTotalAssets)
func (_BendVault *BendVaultFilterer) ParseUpdateLastTotalAssets(log types.Log) (*BendVaultUpdateLastTotalAssets, error) {
	event := new(BendVaultUpdateLastTotalAssets)
	if err := _BendVault.contract.UnpackLog(event, "UpdateLastTotalAssets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultUpdateLostAssetsIterator is returned from FilterUpdateLostAssets and is used to iterate over the raw logs and unpacked data for UpdateLostAssets events raised by the BendVault contract.
type BendVaultUpdateLostAssetsIterator struct {
	Event *BendVaultUpdateLostAssets // Event containing the contract specifics and raw log

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
func (it *BendVaultUpdateLostAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultUpdateLostAssets)
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
		it.Event = new(BendVaultUpdateLostAssets)
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
func (it *BendVaultUpdateLostAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultUpdateLostAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultUpdateLostAssets represents a UpdateLostAssets event raised by the BendVault contract.
type BendVaultUpdateLostAssets struct {
	NewLostAssets *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateLostAssets is a free log retrieval operation binding the contract event 0x548669ea9bcc24888e6d74a69c9865fa98d795686853b8aa3eb87814261bbb71.
//
// Solidity: event UpdateLostAssets(uint256 newLostAssets)
func (_BendVault *BendVaultFilterer) FilterUpdateLostAssets(opts *bind.FilterOpts) (*BendVaultUpdateLostAssetsIterator, error) {

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "UpdateLostAssets")
	if err != nil {
		return nil, err
	}
	return &BendVaultUpdateLostAssetsIterator{contract: _BendVault.contract, event: "UpdateLostAssets", logs: logs, sub: sub}, nil
}

// WatchUpdateLostAssets is a free log subscription operation binding the contract event 0x548669ea9bcc24888e6d74a69c9865fa98d795686853b8aa3eb87814261bbb71.
//
// Solidity: event UpdateLostAssets(uint256 newLostAssets)
func (_BendVault *BendVaultFilterer) WatchUpdateLostAssets(opts *bind.WatchOpts, sink chan<- *BendVaultUpdateLostAssets) (event.Subscription, error) {

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "UpdateLostAssets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultUpdateLostAssets)
				if err := _BendVault.contract.UnpackLog(event, "UpdateLostAssets", log); err != nil {
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

// ParseUpdateLostAssets is a log parse operation binding the contract event 0x548669ea9bcc24888e6d74a69c9865fa98d795686853b8aa3eb87814261bbb71.
//
// Solidity: event UpdateLostAssets(uint256 newLostAssets)
func (_BendVault *BendVaultFilterer) ParseUpdateLostAssets(log types.Log) (*BendVaultUpdateLostAssets, error) {
	event := new(BendVaultUpdateLostAssets)
	if err := _BendVault.contract.UnpackLog(event, "UpdateLostAssets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BendVault contract.
type BendVaultWithdrawIterator struct {
	Event *BendVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *BendVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendVaultWithdraw)
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
		it.Event = new(BendVaultWithdraw)
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
func (it *BendVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendVaultWithdraw represents a Withdraw event raised by the BendVault contract.
type BendVaultWithdraw struct {
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
func (_BendVault *BendVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*BendVaultWithdrawIterator, error) {

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

	logs, sub, err := _BendVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &BendVaultWithdrawIterator{contract: _BendVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_BendVault *BendVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BendVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BendVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendVaultWithdraw)
				if err := _BendVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_BendVault *BendVaultFilterer) ParseWithdraw(log types.Log) (*BendVaultWithdraw, error) {
	event := new(BendVaultWithdraw)
	if err := _BendVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
