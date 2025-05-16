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

// SolvBTCMetaData contains all meta data concerning the SolvBTC contract.
var SolvBTCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BlacklistableBlacklistedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"BlacklistableNotManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlacklistableZeroAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ERC3525NotReceivable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ERC721NotReceivable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PausablePauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"SolvBTCNotBlacklisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolvBTCZeroValueNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"BlacklistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newBlacklistManager\",\"type\":\"address\"}],\"name\":\"BlacklistManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"BlacklistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DestroyBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"SetAlias\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"SetOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLVBTC_MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOLVBTC_POOL_BURNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"addBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts_\",\"type\":\"address[]\"}],\"name\":\"addBlacklistBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blacklistManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"getSharesByValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"getValueByShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"isBlacklisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC3525Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"removeBlacklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts_\",\"type\":\"address[]\"}],\"name\":\"removeBlacklistBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"setAlias\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle_\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"setPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBlacklistManager_\",\"type\":\"address\"}],\"name\":\"updateBlacklistManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SolvBTCABI is the input ABI used to generate the binding from.
// Deprecated: Use SolvBTCMetaData.ABI instead.
var SolvBTCABI = SolvBTCMetaData.ABI

// SolvBTC is an auto generated Go binding around an Ethereum contract.
type SolvBTC struct {
	SolvBTCCaller     // Read-only binding to the contract
	SolvBTCTransactor // Write-only binding to the contract
	SolvBTCFilterer   // Log filterer for contract events
}

// SolvBTCCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolvBTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolvBTCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolvBTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolvBTCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolvBTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolvBTCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolvBTCSession struct {
	Contract     *SolvBTC          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolvBTCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolvBTCCallerSession struct {
	Contract *SolvBTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SolvBTCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolvBTCTransactorSession struct {
	Contract     *SolvBTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SolvBTCRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolvBTCRaw struct {
	Contract *SolvBTC // Generic contract binding to access the raw methods on
}

// SolvBTCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolvBTCCallerRaw struct {
	Contract *SolvBTCCaller // Generic read-only contract binding to access the raw methods on
}

// SolvBTCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolvBTCTransactorRaw struct {
	Contract *SolvBTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolvBTC creates a new instance of SolvBTC, bound to a specific deployed contract.
func NewSolvBTC(address common.Address, backend bind.ContractBackend) (*SolvBTC, error) {
	contract, err := bindSolvBTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolvBTC{SolvBTCCaller: SolvBTCCaller{contract: contract}, SolvBTCTransactor: SolvBTCTransactor{contract: contract}, SolvBTCFilterer: SolvBTCFilterer{contract: contract}}, nil
}

// NewSolvBTCCaller creates a new read-only instance of SolvBTC, bound to a specific deployed contract.
func NewSolvBTCCaller(address common.Address, caller bind.ContractCaller) (*SolvBTCCaller, error) {
	contract, err := bindSolvBTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolvBTCCaller{contract: contract}, nil
}

// NewSolvBTCTransactor creates a new write-only instance of SolvBTC, bound to a specific deployed contract.
func NewSolvBTCTransactor(address common.Address, transactor bind.ContractTransactor) (*SolvBTCTransactor, error) {
	contract, err := bindSolvBTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolvBTCTransactor{contract: contract}, nil
}

// NewSolvBTCFilterer creates a new log filterer instance of SolvBTC, bound to a specific deployed contract.
func NewSolvBTCFilterer(address common.Address, filterer bind.ContractFilterer) (*SolvBTCFilterer, error) {
	contract, err := bindSolvBTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolvBTCFilterer{contract: contract}, nil
}

// bindSolvBTC binds a generic wrapper to an already deployed contract.
func bindSolvBTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolvBTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolvBTC *SolvBTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolvBTC.Contract.SolvBTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolvBTC *SolvBTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolvBTC.Contract.SolvBTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolvBTC *SolvBTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolvBTC.Contract.SolvBTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolvBTC *SolvBTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolvBTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolvBTC *SolvBTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolvBTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolvBTC *SolvBTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolvBTC.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SolvBTC *SolvBTCCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SolvBTC *SolvBTCSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SolvBTC.Contract.DEFAULTADMINROLE(&_SolvBTC.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SolvBTC *SolvBTCCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SolvBTC.Contract.DEFAULTADMINROLE(&_SolvBTC.CallOpts)
}

// SOLVBTCMINTERROLE is a free data retrieval call binding the contract method 0x96c49597.
//
// Solidity: function SOLVBTC_MINTER_ROLE() view returns(bytes32)
func (_SolvBTC *SolvBTCCaller) SOLVBTCMINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "SOLVBTC_MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SOLVBTCMINTERROLE is a free data retrieval call binding the contract method 0x96c49597.
//
// Solidity: function SOLVBTC_MINTER_ROLE() view returns(bytes32)
func (_SolvBTC *SolvBTCSession) SOLVBTCMINTERROLE() ([32]byte, error) {
	return _SolvBTC.Contract.SOLVBTCMINTERROLE(&_SolvBTC.CallOpts)
}

// SOLVBTCMINTERROLE is a free data retrieval call binding the contract method 0x96c49597.
//
// Solidity: function SOLVBTC_MINTER_ROLE() view returns(bytes32)
func (_SolvBTC *SolvBTCCallerSession) SOLVBTCMINTERROLE() ([32]byte, error) {
	return _SolvBTC.Contract.SOLVBTCMINTERROLE(&_SolvBTC.CallOpts)
}

// SOLVBTCPOOLBURNERROLE is a free data retrieval call binding the contract method 0xa49630b2.
//
// Solidity: function SOLVBTC_POOL_BURNER_ROLE() view returns(bytes32)
func (_SolvBTC *SolvBTCCaller) SOLVBTCPOOLBURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "SOLVBTC_POOL_BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SOLVBTCPOOLBURNERROLE is a free data retrieval call binding the contract method 0xa49630b2.
//
// Solidity: function SOLVBTC_POOL_BURNER_ROLE() view returns(bytes32)
func (_SolvBTC *SolvBTCSession) SOLVBTCPOOLBURNERROLE() ([32]byte, error) {
	return _SolvBTC.Contract.SOLVBTCPOOLBURNERROLE(&_SolvBTC.CallOpts)
}

// SOLVBTCPOOLBURNERROLE is a free data retrieval call binding the contract method 0xa49630b2.
//
// Solidity: function SOLVBTC_POOL_BURNER_ROLE() view returns(bytes32)
func (_SolvBTC *SolvBTCCallerSession) SOLVBTCPOOLBURNERROLE() ([32]byte, error) {
	return _SolvBTC.Contract.SOLVBTCPOOLBURNERROLE(&_SolvBTC.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SolvBTC *SolvBTCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SolvBTC *SolvBTCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SolvBTC.Contract.Allowance(&_SolvBTC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SolvBTC *SolvBTCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SolvBTC.Contract.Allowance(&_SolvBTC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SolvBTC *SolvBTCCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SolvBTC *SolvBTCSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SolvBTC.Contract.BalanceOf(&_SolvBTC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SolvBTC *SolvBTCCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SolvBTC.Contract.BalanceOf(&_SolvBTC.CallOpts, account)
}

// BlacklistManager is a free data retrieval call binding the contract method 0xd9dbf657.
//
// Solidity: function blacklistManager() view returns(address)
func (_SolvBTC *SolvBTCCaller) BlacklistManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "blacklistManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlacklistManager is a free data retrieval call binding the contract method 0xd9dbf657.
//
// Solidity: function blacklistManager() view returns(address)
func (_SolvBTC *SolvBTCSession) BlacklistManager() (common.Address, error) {
	return _SolvBTC.Contract.BlacklistManager(&_SolvBTC.CallOpts)
}

// BlacklistManager is a free data retrieval call binding the contract method 0xd9dbf657.
//
// Solidity: function blacklistManager() view returns(address)
func (_SolvBTC *SolvBTCCallerSession) BlacklistManager() (common.Address, error) {
	return _SolvBTC.Contract.BlacklistManager(&_SolvBTC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolvBTC *SolvBTCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolvBTC *SolvBTCSession) Decimals() (uint8, error) {
	return _SolvBTC.Contract.Decimals(&_SolvBTC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SolvBTC *SolvBTCCallerSession) Decimals() (uint8, error) {
	return _SolvBTC.Contract.Decimals(&_SolvBTC.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_SolvBTC *SolvBTCCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_SolvBTC *SolvBTCSession) GetOracle() (common.Address, error) {
	return _SolvBTC.Contract.GetOracle(&_SolvBTC.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_SolvBTC *SolvBTCCallerSession) GetOracle() (common.Address, error) {
	return _SolvBTC.Contract.GetOracle(&_SolvBTC.CallOpts)
}

// GetOracleDecimals is a free data retrieval call binding the contract method 0x1a026b81.
//
// Solidity: function getOracleDecimals() view returns(uint8)
func (_SolvBTC *SolvBTCCaller) GetOracleDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "getOracleDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOracleDecimals is a free data retrieval call binding the contract method 0x1a026b81.
//
// Solidity: function getOracleDecimals() view returns(uint8)
func (_SolvBTC *SolvBTCSession) GetOracleDecimals() (uint8, error) {
	return _SolvBTC.Contract.GetOracleDecimals(&_SolvBTC.CallOpts)
}

// GetOracleDecimals is a free data retrieval call binding the contract method 0x1a026b81.
//
// Solidity: function getOracleDecimals() view returns(uint8)
func (_SolvBTC *SolvBTCCallerSession) GetOracleDecimals() (uint8, error) {
	return _SolvBTC.Contract.GetOracleDecimals(&_SolvBTC.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SolvBTC *SolvBTCCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SolvBTC *SolvBTCSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SolvBTC.Contract.GetRoleAdmin(&_SolvBTC.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SolvBTC *SolvBTCCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SolvBTC.Contract.GetRoleAdmin(&_SolvBTC.CallOpts, role)
}

// GetSharesByValue is a free data retrieval call binding the contract method 0xe31c3a90.
//
// Solidity: function getSharesByValue(uint256 value) view returns(uint256 shares)
func (_SolvBTC *SolvBTCCaller) GetSharesByValue(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "getSharesByValue", value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByValue is a free data retrieval call binding the contract method 0xe31c3a90.
//
// Solidity: function getSharesByValue(uint256 value) view returns(uint256 shares)
func (_SolvBTC *SolvBTCSession) GetSharesByValue(value *big.Int) (*big.Int, error) {
	return _SolvBTC.Contract.GetSharesByValue(&_SolvBTC.CallOpts, value)
}

// GetSharesByValue is a free data retrieval call binding the contract method 0xe31c3a90.
//
// Solidity: function getSharesByValue(uint256 value) view returns(uint256 shares)
func (_SolvBTC *SolvBTCCallerSession) GetSharesByValue(value *big.Int) (*big.Int, error) {
	return _SolvBTC.Contract.GetSharesByValue(&_SolvBTC.CallOpts, value)
}

// GetValueByShares is a free data retrieval call binding the contract method 0x71e26e00.
//
// Solidity: function getValueByShares(uint256 shares) view returns(uint256 value)
func (_SolvBTC *SolvBTCCaller) GetValueByShares(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "getValueByShares", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValueByShares is a free data retrieval call binding the contract method 0x71e26e00.
//
// Solidity: function getValueByShares(uint256 shares) view returns(uint256 value)
func (_SolvBTC *SolvBTCSession) GetValueByShares(shares *big.Int) (*big.Int, error) {
	return _SolvBTC.Contract.GetValueByShares(&_SolvBTC.CallOpts, shares)
}

// GetValueByShares is a free data retrieval call binding the contract method 0x71e26e00.
//
// Solidity: function getValueByShares(uint256 shares) view returns(uint256 value)
func (_SolvBTC *SolvBTCCallerSession) GetValueByShares(shares *big.Int) (*big.Int, error) {
	return _SolvBTC.Contract.GetValueByShares(&_SolvBTC.CallOpts, shares)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SolvBTC *SolvBTCCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SolvBTC *SolvBTCSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SolvBTC.Contract.HasRole(&_SolvBTC.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SolvBTC *SolvBTCCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SolvBTC.Contract.HasRole(&_SolvBTC.CallOpts, role, account)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account_) view returns(bool)
func (_SolvBTC *SolvBTCCaller) IsBlacklisted(opts *bind.CallOpts, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "isBlacklisted", account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account_) view returns(bool)
func (_SolvBTC *SolvBTCSession) IsBlacklisted(account_ common.Address) (bool, error) {
	return _SolvBTC.Contract.IsBlacklisted(&_SolvBTC.CallOpts, account_)
}

// IsBlacklisted is a free data retrieval call binding the contract method 0xfe575a87.
//
// Solidity: function isBlacklisted(address account_) view returns(bool)
func (_SolvBTC *SolvBTCCallerSession) IsBlacklisted(account_ common.Address) (bool, error) {
	return _SolvBTC.Contract.IsBlacklisted(&_SolvBTC.CallOpts, account_)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SolvBTC *SolvBTCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SolvBTC *SolvBTCSession) Name() (string, error) {
	return _SolvBTC.Contract.Name(&_SolvBTC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SolvBTC *SolvBTCCallerSession) Name() (string, error) {
	return _SolvBTC.Contract.Name(&_SolvBTC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SolvBTC *SolvBTCCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SolvBTC *SolvBTCSession) Owner() (common.Address, error) {
	return _SolvBTC.Contract.Owner(&_SolvBTC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SolvBTC *SolvBTCCallerSession) Owner() (common.Address, error) {
	return _SolvBTC.Contract.Owner(&_SolvBTC.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SolvBTC *SolvBTCCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SolvBTC *SolvBTCSession) Paused() (bool, error) {
	return _SolvBTC.Contract.Paused(&_SolvBTC.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SolvBTC *SolvBTCCallerSession) Paused() (bool, error) {
	return _SolvBTC.Contract.Paused(&_SolvBTC.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SolvBTC *SolvBTCCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SolvBTC *SolvBTCSession) PendingOwner() (common.Address, error) {
	return _SolvBTC.Contract.PendingOwner(&_SolvBTC.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SolvBTC *SolvBTCCallerSession) PendingOwner() (common.Address, error) {
	return _SolvBTC.Contract.PendingOwner(&_SolvBTC.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SolvBTC *SolvBTCCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SolvBTC *SolvBTCSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SolvBTC.Contract.SupportsInterface(&_SolvBTC.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SolvBTC *SolvBTCCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SolvBTC.Contract.SupportsInterface(&_SolvBTC.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SolvBTC *SolvBTCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SolvBTC *SolvBTCSession) Symbol() (string, error) {
	return _SolvBTC.Contract.Symbol(&_SolvBTC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SolvBTC *SolvBTCCallerSession) Symbol() (string, error) {
	return _SolvBTC.Contract.Symbol(&_SolvBTC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SolvBTC *SolvBTCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SolvBTC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SolvBTC *SolvBTCSession) TotalSupply() (*big.Int, error) {
	return _SolvBTC.Contract.TotalSupply(&_SolvBTC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SolvBTC *SolvBTCCallerSession) TotalSupply() (*big.Int, error) {
	return _SolvBTC.Contract.TotalSupply(&_SolvBTC.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SolvBTC *SolvBTCTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SolvBTC *SolvBTCSession) AcceptOwnership() (*types.Transaction, error) {
	return _SolvBTC.Contract.AcceptOwnership(&_SolvBTC.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SolvBTC *SolvBTCTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SolvBTC.Contract.AcceptOwnership(&_SolvBTC.TransactOpts)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x9cfe42da.
//
// Solidity: function addBlacklist(address account_) returns()
func (_SolvBTC *SolvBTCTransactor) AddBlacklist(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "addBlacklist", account_)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x9cfe42da.
//
// Solidity: function addBlacklist(address account_) returns()
func (_SolvBTC *SolvBTCSession) AddBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.AddBlacklist(&_SolvBTC.TransactOpts, account_)
}

// AddBlacklist is a paid mutator transaction binding the contract method 0x9cfe42da.
//
// Solidity: function addBlacklist(address account_) returns()
func (_SolvBTC *SolvBTCTransactorSession) AddBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.AddBlacklist(&_SolvBTC.TransactOpts, account_)
}

// AddBlacklistBatch is a paid mutator transaction binding the contract method 0x9999416f.
//
// Solidity: function addBlacklistBatch(address[] accounts_) returns()
func (_SolvBTC *SolvBTCTransactor) AddBlacklistBatch(opts *bind.TransactOpts, accounts_ []common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "addBlacklistBatch", accounts_)
}

// AddBlacklistBatch is a paid mutator transaction binding the contract method 0x9999416f.
//
// Solidity: function addBlacklistBatch(address[] accounts_) returns()
func (_SolvBTC *SolvBTCSession) AddBlacklistBatch(accounts_ []common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.AddBlacklistBatch(&_SolvBTC.TransactOpts, accounts_)
}

// AddBlacklistBatch is a paid mutator transaction binding the contract method 0x9999416f.
//
// Solidity: function addBlacklistBatch(address[] accounts_) returns()
func (_SolvBTC *SolvBTCTransactorSession) AddBlacklistBatch(accounts_ []common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.AddBlacklistBatch(&_SolvBTC.TransactOpts, accounts_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SolvBTC *SolvBTCTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SolvBTC *SolvBTCSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.Approve(&_SolvBTC.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_SolvBTC *SolvBTCTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.Approve(&_SolvBTC.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value_) returns()
func (_SolvBTC *SolvBTCTransactor) Burn(opts *bind.TransactOpts, value_ *big.Int) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "burn", value_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value_) returns()
func (_SolvBTC *SolvBTCSession) Burn(value_ *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.Burn(&_SolvBTC.TransactOpts, value_)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value_) returns()
func (_SolvBTC *SolvBTCTransactorSession) Burn(value_ *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.Burn(&_SolvBTC.TransactOpts, value_)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account_, uint256 value_) returns()
func (_SolvBTC *SolvBTCTransactor) Burn0(opts *bind.TransactOpts, account_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "burn0", account_, value_)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account_, uint256 value_) returns()
func (_SolvBTC *SolvBTCSession) Burn0(account_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.Burn0(&_SolvBTC.TransactOpts, account_, value_)
}

// Burn0 is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account_, uint256 value_) returns()
func (_SolvBTC *SolvBTCTransactorSession) Burn0(account_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.Burn0(&_SolvBTC.TransactOpts, account_, value_)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0x53d51e64.
//
// Solidity: function destroyBlackFunds(address account, uint256 amount) returns()
func (_SolvBTC *SolvBTCTransactor) DestroyBlackFunds(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "destroyBlackFunds", account, amount)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0x53d51e64.
//
// Solidity: function destroyBlackFunds(address account, uint256 amount) returns()
func (_SolvBTC *SolvBTCSession) DestroyBlackFunds(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.DestroyBlackFunds(&_SolvBTC.TransactOpts, account, amount)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0x53d51e64.
//
// Solidity: function destroyBlackFunds(address account, uint256 amount) returns()
func (_SolvBTC *SolvBTCTransactorSession) DestroyBlackFunds(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.DestroyBlackFunds(&_SolvBTC.TransactOpts, account, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SolvBTC *SolvBTCTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SolvBTC *SolvBTCSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.GrantRole(&_SolvBTC.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SolvBTC *SolvBTCTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.GrantRole(&_SolvBTC.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name_, string symbol_, address owner_) returns()
func (_SolvBTC *SolvBTCTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, owner_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "initialize", name_, symbol_, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name_, string symbol_, address owner_) returns()
func (_SolvBTC *SolvBTCSession) Initialize(name_ string, symbol_ string, owner_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.Initialize(&_SolvBTC.TransactOpts, name_, symbol_, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name_, string symbol_, address owner_) returns()
func (_SolvBTC *SolvBTCTransactorSession) Initialize(name_ string, symbol_ string, owner_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.Initialize(&_SolvBTC.TransactOpts, name_, symbol_, owner_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account_, uint256 value_) returns()
func (_SolvBTC *SolvBTCTransactor) Mint(opts *bind.TransactOpts, account_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "mint", account_, value_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account_, uint256 value_) returns()
func (_SolvBTC *SolvBTCSession) Mint(account_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.Mint(&_SolvBTC.TransactOpts, account_, value_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account_, uint256 value_) returns()
func (_SolvBTC *SolvBTCTransactorSession) Mint(account_ common.Address, value_ *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.Mint(&_SolvBTC.TransactOpts, account_, value_)
}

// OnERC3525Received is a paid mutator transaction binding the contract method 0x009ce20b.
//
// Solidity: function onERC3525Received(address , uint256 , uint256 , uint256 , bytes ) returns(bytes4)
func (_SolvBTC *SolvBTCTransactor) OnERC3525Received(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "onERC3525Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC3525Received is a paid mutator transaction binding the contract method 0x009ce20b.
//
// Solidity: function onERC3525Received(address , uint256 , uint256 , uint256 , bytes ) returns(bytes4)
func (_SolvBTC *SolvBTCSession) OnERC3525Received(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SolvBTC.Contract.OnERC3525Received(&_SolvBTC.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC3525Received is a paid mutator transaction binding the contract method 0x009ce20b.
//
// Solidity: function onERC3525Received(address , uint256 , uint256 , uint256 , bytes ) returns(bytes4)
func (_SolvBTC *SolvBTCTransactorSession) OnERC3525Received(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _SolvBTC.Contract.OnERC3525Received(&_SolvBTC.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_SolvBTC *SolvBTCTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_SolvBTC *SolvBTCSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _SolvBTC.Contract.OnERC721Received(&_SolvBTC.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_SolvBTC *SolvBTCTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _SolvBTC.Contract.OnERC721Received(&_SolvBTC.TransactOpts, arg0, arg1, arg2, arg3)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SolvBTC *SolvBTCTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SolvBTC *SolvBTCSession) Pause() (*types.Transaction, error) {
	return _SolvBTC.Contract.Pause(&_SolvBTC.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SolvBTC *SolvBTCTransactorSession) Pause() (*types.Transaction, error) {
	return _SolvBTC.Contract.Pause(&_SolvBTC.TransactOpts)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0xeb91e651.
//
// Solidity: function removeBlacklist(address account_) returns()
func (_SolvBTC *SolvBTCTransactor) RemoveBlacklist(opts *bind.TransactOpts, account_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "removeBlacklist", account_)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0xeb91e651.
//
// Solidity: function removeBlacklist(address account_) returns()
func (_SolvBTC *SolvBTCSession) RemoveBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.RemoveBlacklist(&_SolvBTC.TransactOpts, account_)
}

// RemoveBlacklist is a paid mutator transaction binding the contract method 0xeb91e651.
//
// Solidity: function removeBlacklist(address account_) returns()
func (_SolvBTC *SolvBTCTransactorSession) RemoveBlacklist(account_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.RemoveBlacklist(&_SolvBTC.TransactOpts, account_)
}

// RemoveBlacklistBatch is a paid mutator transaction binding the contract method 0xef2af922.
//
// Solidity: function removeBlacklistBatch(address[] accounts_) returns()
func (_SolvBTC *SolvBTCTransactor) RemoveBlacklistBatch(opts *bind.TransactOpts, accounts_ []common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "removeBlacklistBatch", accounts_)
}

// RemoveBlacklistBatch is a paid mutator transaction binding the contract method 0xef2af922.
//
// Solidity: function removeBlacklistBatch(address[] accounts_) returns()
func (_SolvBTC *SolvBTCSession) RemoveBlacklistBatch(accounts_ []common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.RemoveBlacklistBatch(&_SolvBTC.TransactOpts, accounts_)
}

// RemoveBlacklistBatch is a paid mutator transaction binding the contract method 0xef2af922.
//
// Solidity: function removeBlacklistBatch(address[] accounts_) returns()
func (_SolvBTC *SolvBTCTransactorSession) RemoveBlacklistBatch(accounts_ []common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.RemoveBlacklistBatch(&_SolvBTC.TransactOpts, accounts_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SolvBTC *SolvBTCTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SolvBTC *SolvBTCSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolvBTC.Contract.RenounceOwnership(&_SolvBTC.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SolvBTC *SolvBTCTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolvBTC.Contract.RenounceOwnership(&_SolvBTC.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SolvBTC *SolvBTCTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SolvBTC *SolvBTCSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.RenounceRole(&_SolvBTC.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SolvBTC *SolvBTCTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.RenounceRole(&_SolvBTC.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SolvBTC *SolvBTCTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SolvBTC *SolvBTCSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.RevokeRole(&_SolvBTC.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SolvBTC *SolvBTCTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.RevokeRole(&_SolvBTC.TransactOpts, role, account)
}

// SetAlias is a paid mutator transaction binding the contract method 0x677f8ac0.
//
// Solidity: function setAlias(string name_, string symbol_) returns()
func (_SolvBTC *SolvBTCTransactor) SetAlias(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "setAlias", name_, symbol_)
}

// SetAlias is a paid mutator transaction binding the contract method 0x677f8ac0.
//
// Solidity: function setAlias(string name_, string symbol_) returns()
func (_SolvBTC *SolvBTCSession) SetAlias(name_ string, symbol_ string) (*types.Transaction, error) {
	return _SolvBTC.Contract.SetAlias(&_SolvBTC.TransactOpts, name_, symbol_)
}

// SetAlias is a paid mutator transaction binding the contract method 0x677f8ac0.
//
// Solidity: function setAlias(string name_, string symbol_) returns()
func (_SolvBTC *SolvBTCTransactorSession) SetAlias(name_ string, symbol_ string) (*types.Transaction, error) {
	return _SolvBTC.Contract.SetAlias(&_SolvBTC.TransactOpts, name_, symbol_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_SolvBTC *SolvBTCTransactor) SetOracle(opts *bind.TransactOpts, oracle_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "setOracle", oracle_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_SolvBTC *SolvBTCSession) SetOracle(oracle_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.SetOracle(&_SolvBTC.TransactOpts, oracle_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_SolvBTC *SolvBTCTransactorSession) SetOracle(oracle_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.SetOracle(&_SolvBTC.TransactOpts, oracle_)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address pauser) returns()
func (_SolvBTC *SolvBTCTransactor) SetPauser(opts *bind.TransactOpts, pauser common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "setPauser", pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address pauser) returns()
func (_SolvBTC *SolvBTCSession) SetPauser(pauser common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.SetPauser(&_SolvBTC.TransactOpts, pauser)
}

// SetPauser is a paid mutator transaction binding the contract method 0x2d88af4a.
//
// Solidity: function setPauser(address pauser) returns()
func (_SolvBTC *SolvBTCTransactorSession) SetPauser(pauser common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.SetPauser(&_SolvBTC.TransactOpts, pauser)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SolvBTC *SolvBTCTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SolvBTC *SolvBTCSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.Transfer(&_SolvBTC.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_SolvBTC *SolvBTCTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.Transfer(&_SolvBTC.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SolvBTC *SolvBTCTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SolvBTC *SolvBTCSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.TransferFrom(&_SolvBTC.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_SolvBTC *SolvBTCTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _SolvBTC.Contract.TransferFrom(&_SolvBTC.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SolvBTC *SolvBTCTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SolvBTC *SolvBTCSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.TransferOwnership(&_SolvBTC.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SolvBTC *SolvBTCTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.TransferOwnership(&_SolvBTC.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SolvBTC *SolvBTCTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SolvBTC *SolvBTCSession) Unpause() (*types.Transaction, error) {
	return _SolvBTC.Contract.Unpause(&_SolvBTC.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SolvBTC *SolvBTCTransactorSession) Unpause() (*types.Transaction, error) {
	return _SolvBTC.Contract.Unpause(&_SolvBTC.TransactOpts)
}

// UpdateBlacklistManager is a paid mutator transaction binding the contract method 0x38b20518.
//
// Solidity: function updateBlacklistManager(address newBlacklistManager_) returns()
func (_SolvBTC *SolvBTCTransactor) UpdateBlacklistManager(opts *bind.TransactOpts, newBlacklistManager_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.contract.Transact(opts, "updateBlacklistManager", newBlacklistManager_)
}

// UpdateBlacklistManager is a paid mutator transaction binding the contract method 0x38b20518.
//
// Solidity: function updateBlacklistManager(address newBlacklistManager_) returns()
func (_SolvBTC *SolvBTCSession) UpdateBlacklistManager(newBlacklistManager_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.UpdateBlacklistManager(&_SolvBTC.TransactOpts, newBlacklistManager_)
}

// UpdateBlacklistManager is a paid mutator transaction binding the contract method 0x38b20518.
//
// Solidity: function updateBlacklistManager(address newBlacklistManager_) returns()
func (_SolvBTC *SolvBTCTransactorSession) UpdateBlacklistManager(newBlacklistManager_ common.Address) (*types.Transaction, error) {
	return _SolvBTC.Contract.UpdateBlacklistManager(&_SolvBTC.TransactOpts, newBlacklistManager_)
}

// SolvBTCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SolvBTC contract.
type SolvBTCApprovalIterator struct {
	Event *SolvBTCApproval // Event containing the contract specifics and raw log

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
func (it *SolvBTCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCApproval)
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
		it.Event = new(SolvBTCApproval)
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
func (it *SolvBTCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCApproval represents a Approval event raised by the SolvBTC contract.
type SolvBTCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SolvBTC *SolvBTCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SolvBTCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCApprovalIterator{contract: _SolvBTC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SolvBTC *SolvBTCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SolvBTCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCApproval)
				if err := _SolvBTC.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_SolvBTC *SolvBTCFilterer) ParseApproval(log types.Log) (*SolvBTCApproval, error) {
	event := new(SolvBTCApproval)
	if err := _SolvBTC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCBlacklistAddedIterator is returned from FilterBlacklistAdded and is used to iterate over the raw logs and unpacked data for BlacklistAdded events raised by the SolvBTC contract.
type SolvBTCBlacklistAddedIterator struct {
	Event *SolvBTCBlacklistAdded // Event containing the contract specifics and raw log

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
func (it *SolvBTCBlacklistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCBlacklistAdded)
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
		it.Event = new(SolvBTCBlacklistAdded)
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
func (it *SolvBTCBlacklistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCBlacklistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCBlacklistAdded represents a BlacklistAdded event raised by the SolvBTC contract.
type SolvBTCBlacklistAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklistAdded is a free log retrieval operation binding the contract event 0x44d5fe68b00f68950fb9c1ff0a61ef7f747b1a36359a7e3a7f3324db4b878967.
//
// Solidity: event BlacklistAdded(address indexed account_)
func (_SolvBTC *SolvBTCFilterer) FilterBlacklistAdded(opts *bind.FilterOpts, account_ []common.Address) (*SolvBTCBlacklistAddedIterator, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "BlacklistAdded", account_Rule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCBlacklistAddedIterator{contract: _SolvBTC.contract, event: "BlacklistAdded", logs: logs, sub: sub}, nil
}

// WatchBlacklistAdded is a free log subscription operation binding the contract event 0x44d5fe68b00f68950fb9c1ff0a61ef7f747b1a36359a7e3a7f3324db4b878967.
//
// Solidity: event BlacklistAdded(address indexed account_)
func (_SolvBTC *SolvBTCFilterer) WatchBlacklistAdded(opts *bind.WatchOpts, sink chan<- *SolvBTCBlacklistAdded, account_ []common.Address) (event.Subscription, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "BlacklistAdded", account_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCBlacklistAdded)
				if err := _SolvBTC.contract.UnpackLog(event, "BlacklistAdded", log); err != nil {
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

// ParseBlacklistAdded is a log parse operation binding the contract event 0x44d5fe68b00f68950fb9c1ff0a61ef7f747b1a36359a7e3a7f3324db4b878967.
//
// Solidity: event BlacklistAdded(address indexed account_)
func (_SolvBTC *SolvBTCFilterer) ParseBlacklistAdded(log types.Log) (*SolvBTCBlacklistAdded, error) {
	event := new(SolvBTCBlacklistAdded)
	if err := _SolvBTC.contract.UnpackLog(event, "BlacklistAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCBlacklistManagerChangedIterator is returned from FilterBlacklistManagerChanged and is used to iterate over the raw logs and unpacked data for BlacklistManagerChanged events raised by the SolvBTC contract.
type SolvBTCBlacklistManagerChangedIterator struct {
	Event *SolvBTCBlacklistManagerChanged // Event containing the contract specifics and raw log

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
func (it *SolvBTCBlacklistManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCBlacklistManagerChanged)
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
		it.Event = new(SolvBTCBlacklistManagerChanged)
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
func (it *SolvBTCBlacklistManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCBlacklistManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCBlacklistManagerChanged represents a BlacklistManagerChanged event raised by the SolvBTC contract.
type SolvBTCBlacklistManagerChanged struct {
	NewBlacklistManager common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBlacklistManagerChanged is a free log retrieval operation binding the contract event 0x5baec8c712a7efe1ef755579f2a5b46fe2ffb57d5d216b746e7130605ee9e976.
//
// Solidity: event BlacklistManagerChanged(address indexed newBlacklistManager)
func (_SolvBTC *SolvBTCFilterer) FilterBlacklistManagerChanged(opts *bind.FilterOpts, newBlacklistManager []common.Address) (*SolvBTCBlacklistManagerChangedIterator, error) {

	var newBlacklistManagerRule []interface{}
	for _, newBlacklistManagerItem := range newBlacklistManager {
		newBlacklistManagerRule = append(newBlacklistManagerRule, newBlacklistManagerItem)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "BlacklistManagerChanged", newBlacklistManagerRule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCBlacklistManagerChangedIterator{contract: _SolvBTC.contract, event: "BlacklistManagerChanged", logs: logs, sub: sub}, nil
}

// WatchBlacklistManagerChanged is a free log subscription operation binding the contract event 0x5baec8c712a7efe1ef755579f2a5b46fe2ffb57d5d216b746e7130605ee9e976.
//
// Solidity: event BlacklistManagerChanged(address indexed newBlacklistManager)
func (_SolvBTC *SolvBTCFilterer) WatchBlacklistManagerChanged(opts *bind.WatchOpts, sink chan<- *SolvBTCBlacklistManagerChanged, newBlacklistManager []common.Address) (event.Subscription, error) {

	var newBlacklistManagerRule []interface{}
	for _, newBlacklistManagerItem := range newBlacklistManager {
		newBlacklistManagerRule = append(newBlacklistManagerRule, newBlacklistManagerItem)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "BlacklistManagerChanged", newBlacklistManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCBlacklistManagerChanged)
				if err := _SolvBTC.contract.UnpackLog(event, "BlacklistManagerChanged", log); err != nil {
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

// ParseBlacklistManagerChanged is a log parse operation binding the contract event 0x5baec8c712a7efe1ef755579f2a5b46fe2ffb57d5d216b746e7130605ee9e976.
//
// Solidity: event BlacklistManagerChanged(address indexed newBlacklistManager)
func (_SolvBTC *SolvBTCFilterer) ParseBlacklistManagerChanged(log types.Log) (*SolvBTCBlacklistManagerChanged, error) {
	event := new(SolvBTCBlacklistManagerChanged)
	if err := _SolvBTC.contract.UnpackLog(event, "BlacklistManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCBlacklistRemovedIterator is returned from FilterBlacklistRemoved and is used to iterate over the raw logs and unpacked data for BlacklistRemoved events raised by the SolvBTC contract.
type SolvBTCBlacklistRemovedIterator struct {
	Event *SolvBTCBlacklistRemoved // Event containing the contract specifics and raw log

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
func (it *SolvBTCBlacklistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCBlacklistRemoved)
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
		it.Event = new(SolvBTCBlacklistRemoved)
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
func (it *SolvBTCBlacklistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCBlacklistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCBlacklistRemoved represents a BlacklistRemoved event raised by the SolvBTC contract.
type SolvBTCBlacklistRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBlacklistRemoved is a free log retrieval operation binding the contract event 0x1747ca720b1a174a464b6513ace29b1d3190b5f632b9f34147017c81425bfde8.
//
// Solidity: event BlacklistRemoved(address indexed account_)
func (_SolvBTC *SolvBTCFilterer) FilterBlacklistRemoved(opts *bind.FilterOpts, account_ []common.Address) (*SolvBTCBlacklistRemovedIterator, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "BlacklistRemoved", account_Rule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCBlacklistRemovedIterator{contract: _SolvBTC.contract, event: "BlacklistRemoved", logs: logs, sub: sub}, nil
}

// WatchBlacklistRemoved is a free log subscription operation binding the contract event 0x1747ca720b1a174a464b6513ace29b1d3190b5f632b9f34147017c81425bfde8.
//
// Solidity: event BlacklistRemoved(address indexed account_)
func (_SolvBTC *SolvBTCFilterer) WatchBlacklistRemoved(opts *bind.WatchOpts, sink chan<- *SolvBTCBlacklistRemoved, account_ []common.Address) (event.Subscription, error) {

	var account_Rule []interface{}
	for _, account_Item := range account_ {
		account_Rule = append(account_Rule, account_Item)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "BlacklistRemoved", account_Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCBlacklistRemoved)
				if err := _SolvBTC.contract.UnpackLog(event, "BlacklistRemoved", log); err != nil {
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

// ParseBlacklistRemoved is a log parse operation binding the contract event 0x1747ca720b1a174a464b6513ace29b1d3190b5f632b9f34147017c81425bfde8.
//
// Solidity: event BlacklistRemoved(address indexed account_)
func (_SolvBTC *SolvBTCFilterer) ParseBlacklistRemoved(log types.Log) (*SolvBTCBlacklistRemoved, error) {
	event := new(SolvBTCBlacklistRemoved)
	if err := _SolvBTC.contract.UnpackLog(event, "BlacklistRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCDestroyBlackFundsIterator is returned from FilterDestroyBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyBlackFunds events raised by the SolvBTC contract.
type SolvBTCDestroyBlackFundsIterator struct {
	Event *SolvBTCDestroyBlackFunds // Event containing the contract specifics and raw log

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
func (it *SolvBTCDestroyBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCDestroyBlackFunds)
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
		it.Event = new(SolvBTCDestroyBlackFunds)
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
func (it *SolvBTCDestroyBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCDestroyBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCDestroyBlackFunds represents a DestroyBlackFunds event raised by the SolvBTC contract.
type SolvBTCDestroyBlackFunds struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDestroyBlackFunds is a free log retrieval operation binding the contract event 0x11d33c4bdbad6892d3d8fe9b29fca9d1701c823ddea540b942b102366a4a47e2.
//
// Solidity: event DestroyBlackFunds(address indexed account, uint256 amount)
func (_SolvBTC *SolvBTCFilterer) FilterDestroyBlackFunds(opts *bind.FilterOpts, account []common.Address) (*SolvBTCDestroyBlackFundsIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "DestroyBlackFunds", accountRule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCDestroyBlackFundsIterator{contract: _SolvBTC.contract, event: "DestroyBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyBlackFunds is a free log subscription operation binding the contract event 0x11d33c4bdbad6892d3d8fe9b29fca9d1701c823ddea540b942b102366a4a47e2.
//
// Solidity: event DestroyBlackFunds(address indexed account, uint256 amount)
func (_SolvBTC *SolvBTCFilterer) WatchDestroyBlackFunds(opts *bind.WatchOpts, sink chan<- *SolvBTCDestroyBlackFunds, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "DestroyBlackFunds", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCDestroyBlackFunds)
				if err := _SolvBTC.contract.UnpackLog(event, "DestroyBlackFunds", log); err != nil {
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

// ParseDestroyBlackFunds is a log parse operation binding the contract event 0x11d33c4bdbad6892d3d8fe9b29fca9d1701c823ddea540b942b102366a4a47e2.
//
// Solidity: event DestroyBlackFunds(address indexed account, uint256 amount)
func (_SolvBTC *SolvBTCFilterer) ParseDestroyBlackFunds(log types.Log) (*SolvBTCDestroyBlackFunds, error) {
	event := new(SolvBTCDestroyBlackFunds)
	if err := _SolvBTC.contract.UnpackLog(event, "DestroyBlackFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SolvBTC contract.
type SolvBTCInitializedIterator struct {
	Event *SolvBTCInitialized // Event containing the contract specifics and raw log

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
func (it *SolvBTCInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCInitialized)
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
		it.Event = new(SolvBTCInitialized)
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
func (it *SolvBTCInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCInitialized represents a Initialized event raised by the SolvBTC contract.
type SolvBTCInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolvBTC *SolvBTCFilterer) FilterInitialized(opts *bind.FilterOpts) (*SolvBTCInitializedIterator, error) {

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SolvBTCInitializedIterator{contract: _SolvBTC.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolvBTC *SolvBTCFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SolvBTCInitialized) (event.Subscription, error) {

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCInitialized)
				if err := _SolvBTC.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SolvBTC *SolvBTCFilterer) ParseInitialized(log types.Log) (*SolvBTCInitialized, error) {
	event := new(SolvBTCInitialized)
	if err := _SolvBTC.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the SolvBTC contract.
type SolvBTCOwnershipTransferStartedIterator struct {
	Event *SolvBTCOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *SolvBTCOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCOwnershipTransferStarted)
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
		it.Event = new(SolvBTCOwnershipTransferStarted)
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
func (it *SolvBTCOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the SolvBTC contract.
type SolvBTCOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SolvBTC *SolvBTCFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SolvBTCOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCOwnershipTransferStartedIterator{contract: _SolvBTC.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SolvBTC *SolvBTCFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SolvBTCOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCOwnershipTransferStarted)
				if err := _SolvBTC.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_SolvBTC *SolvBTCFilterer) ParseOwnershipTransferStarted(log types.Log) (*SolvBTCOwnershipTransferStarted, error) {
	event := new(SolvBTCOwnershipTransferStarted)
	if err := _SolvBTC.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SolvBTC contract.
type SolvBTCOwnershipTransferredIterator struct {
	Event *SolvBTCOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SolvBTCOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCOwnershipTransferred)
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
		it.Event = new(SolvBTCOwnershipTransferred)
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
func (it *SolvBTCOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCOwnershipTransferred represents a OwnershipTransferred event raised by the SolvBTC contract.
type SolvBTCOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SolvBTC *SolvBTCFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SolvBTCOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCOwnershipTransferredIterator{contract: _SolvBTC.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SolvBTC *SolvBTCFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SolvBTCOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCOwnershipTransferred)
				if err := _SolvBTC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SolvBTC *SolvBTCFilterer) ParseOwnershipTransferred(log types.Log) (*SolvBTCOwnershipTransferred, error) {
	event := new(SolvBTCOwnershipTransferred)
	if err := _SolvBTC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SolvBTC contract.
type SolvBTCPausedIterator struct {
	Event *SolvBTCPaused // Event containing the contract specifics and raw log

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
func (it *SolvBTCPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCPaused)
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
		it.Event = new(SolvBTCPaused)
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
func (it *SolvBTCPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCPaused represents a Paused event raised by the SolvBTC contract.
type SolvBTCPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SolvBTC *SolvBTCFilterer) FilterPaused(opts *bind.FilterOpts) (*SolvBTCPausedIterator, error) {

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SolvBTCPausedIterator{contract: _SolvBTC.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SolvBTC *SolvBTCFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SolvBTCPaused) (event.Subscription, error) {

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCPaused)
				if err := _SolvBTC.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SolvBTC *SolvBTCFilterer) ParsePaused(log types.Log) (*SolvBTCPaused, error) {
	event := new(SolvBTCPaused)
	if err := _SolvBTC.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SolvBTC contract.
type SolvBTCRoleAdminChangedIterator struct {
	Event *SolvBTCRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SolvBTCRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCRoleAdminChanged)
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
		it.Event = new(SolvBTCRoleAdminChanged)
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
func (it *SolvBTCRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCRoleAdminChanged represents a RoleAdminChanged event raised by the SolvBTC contract.
type SolvBTCRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SolvBTC *SolvBTCFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SolvBTCRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCRoleAdminChangedIterator{contract: _SolvBTC.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SolvBTC *SolvBTCFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SolvBTCRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCRoleAdminChanged)
				if err := _SolvBTC.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SolvBTC *SolvBTCFilterer) ParseRoleAdminChanged(log types.Log) (*SolvBTCRoleAdminChanged, error) {
	event := new(SolvBTCRoleAdminChanged)
	if err := _SolvBTC.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SolvBTC contract.
type SolvBTCRoleGrantedIterator struct {
	Event *SolvBTCRoleGranted // Event containing the contract specifics and raw log

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
func (it *SolvBTCRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCRoleGranted)
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
		it.Event = new(SolvBTCRoleGranted)
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
func (it *SolvBTCRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCRoleGranted represents a RoleGranted event raised by the SolvBTC contract.
type SolvBTCRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolvBTC *SolvBTCFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SolvBTCRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCRoleGrantedIterator{contract: _SolvBTC.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolvBTC *SolvBTCFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SolvBTCRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCRoleGranted)
				if err := _SolvBTC.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolvBTC *SolvBTCFilterer) ParseRoleGranted(log types.Log) (*SolvBTCRoleGranted, error) {
	event := new(SolvBTCRoleGranted)
	if err := _SolvBTC.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SolvBTC contract.
type SolvBTCRoleRevokedIterator struct {
	Event *SolvBTCRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SolvBTCRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCRoleRevoked)
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
		it.Event = new(SolvBTCRoleRevoked)
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
func (it *SolvBTCRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCRoleRevoked represents a RoleRevoked event raised by the SolvBTC contract.
type SolvBTCRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolvBTC *SolvBTCFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SolvBTCRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCRoleRevokedIterator{contract: _SolvBTC.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolvBTC *SolvBTCFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SolvBTCRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCRoleRevoked)
				if err := _SolvBTC.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SolvBTC *SolvBTCFilterer) ParseRoleRevoked(log types.Log) (*SolvBTCRoleRevoked, error) {
	event := new(SolvBTCRoleRevoked)
	if err := _SolvBTC.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCSetAliasIterator is returned from FilterSetAlias and is used to iterate over the raw logs and unpacked data for SetAlias events raised by the SolvBTC contract.
type SolvBTCSetAliasIterator struct {
	Event *SolvBTCSetAlias // Event containing the contract specifics and raw log

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
func (it *SolvBTCSetAliasIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCSetAlias)
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
		it.Event = new(SolvBTCSetAlias)
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
func (it *SolvBTCSetAliasIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCSetAliasIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCSetAlias represents a SetAlias event raised by the SolvBTC contract.
type SolvBTCSetAlias struct {
	Name   string
	Symbol string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetAlias is a free log retrieval operation binding the contract event 0x841f96fea9cdc9cd8d9e403b87390322f34ff38914b7840853d1f1c9e7d54650.
//
// Solidity: event SetAlias(string name, string symbol)
func (_SolvBTC *SolvBTCFilterer) FilterSetAlias(opts *bind.FilterOpts) (*SolvBTCSetAliasIterator, error) {

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "SetAlias")
	if err != nil {
		return nil, err
	}
	return &SolvBTCSetAliasIterator{contract: _SolvBTC.contract, event: "SetAlias", logs: logs, sub: sub}, nil
}

// WatchSetAlias is a free log subscription operation binding the contract event 0x841f96fea9cdc9cd8d9e403b87390322f34ff38914b7840853d1f1c9e7d54650.
//
// Solidity: event SetAlias(string name, string symbol)
func (_SolvBTC *SolvBTCFilterer) WatchSetAlias(opts *bind.WatchOpts, sink chan<- *SolvBTCSetAlias) (event.Subscription, error) {

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "SetAlias")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCSetAlias)
				if err := _SolvBTC.contract.UnpackLog(event, "SetAlias", log); err != nil {
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

// ParseSetAlias is a log parse operation binding the contract event 0x841f96fea9cdc9cd8d9e403b87390322f34ff38914b7840853d1f1c9e7d54650.
//
// Solidity: event SetAlias(string name, string symbol)
func (_SolvBTC *SolvBTCFilterer) ParseSetAlias(log types.Log) (*SolvBTCSetAlias, error) {
	event := new(SolvBTCSetAlias)
	if err := _SolvBTC.contract.UnpackLog(event, "SetAlias", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCSetOracleIterator is returned from FilterSetOracle and is used to iterate over the raw logs and unpacked data for SetOracle events raised by the SolvBTC contract.
type SolvBTCSetOracleIterator struct {
	Event *SolvBTCSetOracle // Event containing the contract specifics and raw log

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
func (it *SolvBTCSetOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCSetOracle)
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
		it.Event = new(SolvBTCSetOracle)
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
func (it *SolvBTCSetOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCSetOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCSetOracle represents a SetOracle event raised by the SolvBTC contract.
type SolvBTCSetOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOracle is a free log retrieval operation binding the contract event 0xd3b5d1e0ffaeff528910f3663f0adace7694ab8241d58e17a91351ced2e08031.
//
// Solidity: event SetOracle(address indexed oracle)
func (_SolvBTC *SolvBTCFilterer) FilterSetOracle(opts *bind.FilterOpts, oracle []common.Address) (*SolvBTCSetOracleIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "SetOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCSetOracleIterator{contract: _SolvBTC.contract, event: "SetOracle", logs: logs, sub: sub}, nil
}

// WatchSetOracle is a free log subscription operation binding the contract event 0xd3b5d1e0ffaeff528910f3663f0adace7694ab8241d58e17a91351ced2e08031.
//
// Solidity: event SetOracle(address indexed oracle)
func (_SolvBTC *SolvBTCFilterer) WatchSetOracle(opts *bind.WatchOpts, sink chan<- *SolvBTCSetOracle, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "SetOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCSetOracle)
				if err := _SolvBTC.contract.UnpackLog(event, "SetOracle", log); err != nil {
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

// ParseSetOracle is a log parse operation binding the contract event 0xd3b5d1e0ffaeff528910f3663f0adace7694ab8241d58e17a91351ced2e08031.
//
// Solidity: event SetOracle(address indexed oracle)
func (_SolvBTC *SolvBTCFilterer) ParseSetOracle(log types.Log) (*SolvBTCSetOracle, error) {
	event := new(SolvBTCSetOracle)
	if err := _SolvBTC.contract.UnpackLog(event, "SetOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SolvBTC contract.
type SolvBTCTransferIterator struct {
	Event *SolvBTCTransfer // Event containing the contract specifics and raw log

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
func (it *SolvBTCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCTransfer)
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
		it.Event = new(SolvBTCTransfer)
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
func (it *SolvBTCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCTransfer represents a Transfer event raised by the SolvBTC contract.
type SolvBTCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SolvBTC *SolvBTCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SolvBTCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SolvBTCTransferIterator{contract: _SolvBTC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SolvBTC *SolvBTCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SolvBTCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCTransfer)
				if err := _SolvBTC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_SolvBTC *SolvBTCFilterer) ParseTransfer(log types.Log) (*SolvBTCTransfer, error) {
	event := new(SolvBTCTransfer)
	if err := _SolvBTC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolvBTCUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SolvBTC contract.
type SolvBTCUnpausedIterator struct {
	Event *SolvBTCUnpaused // Event containing the contract specifics and raw log

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
func (it *SolvBTCUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolvBTCUnpaused)
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
		it.Event = new(SolvBTCUnpaused)
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
func (it *SolvBTCUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolvBTCUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolvBTCUnpaused represents a Unpaused event raised by the SolvBTC contract.
type SolvBTCUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SolvBTC *SolvBTCFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SolvBTCUnpausedIterator, error) {

	logs, sub, err := _SolvBTC.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SolvBTCUnpausedIterator{contract: _SolvBTC.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SolvBTC *SolvBTCFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SolvBTCUnpaused) (event.Subscription, error) {

	logs, sub, err := _SolvBTC.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolvBTCUnpaused)
				if err := _SolvBTC.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SolvBTC *SolvBTCFilterer) ParseUnpaused(log types.Log) (*SolvBTCUnpaused, error) {
	event := new(SolvBTCUnpaused)
	if err := _SolvBTC.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
