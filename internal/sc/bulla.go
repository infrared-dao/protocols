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

// BullaMetaData contains all meta data concerning the Bulla contract.
var BullaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"tickOutOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"Rebalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newFee\",\"type\":\"uint8\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"fee\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fees1\",\"type\":\"uint256\"}],\"name\":\"ZeroBurn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"inMin\",\"type\":\"uint256[2]\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"algebraMintCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"inMin\",\"type\":\"uint256[4]\"}],\"name\":\"compound\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"baseToken0Owed\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"baseToken1Owed\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"limitToken0Owed\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"limitToken1Owed\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTick\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit1\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"inMin\",\"type\":\"uint256[4]\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit0Max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit1Max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"directDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasePosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimitPosition\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitLower\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitUpper\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIAlgebraPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"uint128\",\"name\":\"shares\",\"type\":\"uint128\"},{\"internalType\":\"uint256[2]\",\"name\":\"amountMin\",\"type\":\"uint256[2]\"}],\"name\":\"pullLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"_baseLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_baseUpper\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_limitLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"_limitUpper\",\"type\":\"int24\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"inMin\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[4]\",\"name\":\"outMin\",\"type\":\"uint256[4]\"}],\"name\":\"rebalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeWhitelisted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"newFee\",\"type\":\"uint8\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"newTickSpacing\",\"type\":\"int24\"}],\"name\":\"setTickSpacing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleDirectDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelistedAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[4]\",\"name\":\"minAmounts\",\"type\":\"uint256[4]\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zeroBurn\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"baseLiquidity\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"limitLiquidity\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BullaABI is the input ABI used to generate the binding from.
// Deprecated: Use BullaMetaData.ABI instead.
var BullaABI = BullaMetaData.ABI

// Bulla is an auto generated Go binding around an Ethereum contract.
type Bulla struct {
	BullaCaller     // Read-only binding to the contract
	BullaTransactor // Write-only binding to the contract
	BullaFilterer   // Log filterer for contract events
}

// BullaCaller is an auto generated read-only Go binding around an Ethereum contract.
type BullaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BullaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BullaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BullaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BullaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BullaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BullaSession struct {
	Contract     *Bulla            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BullaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BullaCallerSession struct {
	Contract *BullaCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BullaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BullaTransactorSession struct {
	Contract     *BullaTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BullaRaw is an auto generated low-level Go binding around an Ethereum contract.
type BullaRaw struct {
	Contract *Bulla // Generic contract binding to access the raw methods on
}

// BullaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BullaCallerRaw struct {
	Contract *BullaCaller // Generic read-only contract binding to access the raw methods on
}

// BullaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BullaTransactorRaw struct {
	Contract *BullaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBulla creates a new instance of Bulla, bound to a specific deployed contract.
func NewBulla(address common.Address, backend bind.ContractBackend) (*Bulla, error) {
	contract, err := bindBulla(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bulla{BullaCaller: BullaCaller{contract: contract}, BullaTransactor: BullaTransactor{contract: contract}, BullaFilterer: BullaFilterer{contract: contract}}, nil
}

// NewBullaCaller creates a new read-only instance of Bulla, bound to a specific deployed contract.
func NewBullaCaller(address common.Address, caller bind.ContractCaller) (*BullaCaller, error) {
	contract, err := bindBulla(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BullaCaller{contract: contract}, nil
}

// NewBullaTransactor creates a new write-only instance of Bulla, bound to a specific deployed contract.
func NewBullaTransactor(address common.Address, transactor bind.ContractTransactor) (*BullaTransactor, error) {
	contract, err := bindBulla(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BullaTransactor{contract: contract}, nil
}

// NewBullaFilterer creates a new log filterer instance of Bulla, bound to a specific deployed contract.
func NewBullaFilterer(address common.Address, filterer bind.ContractFilterer) (*BullaFilterer, error) {
	contract, err := bindBulla(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BullaFilterer{contract: contract}, nil
}

// bindBulla binds a generic wrapper to an already deployed contract.
func bindBulla(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BullaMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bulla *BullaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bulla.Contract.BullaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bulla *BullaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bulla.Contract.BullaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bulla *BullaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bulla.Contract.BullaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bulla *BullaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bulla.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bulla *BullaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bulla.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bulla *BullaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bulla.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bulla *BullaCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bulla *BullaSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bulla.Contract.DOMAINSEPARATOR(&_Bulla.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Bulla *BullaCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Bulla.Contract.DOMAINSEPARATOR(&_Bulla.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Bulla *BullaCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Bulla *BullaSession) PRECISION() (*big.Int, error) {
	return _Bulla.Contract.PRECISION(&_Bulla.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_Bulla *BullaCallerSession) PRECISION() (*big.Int, error) {
	return _Bulla.Contract.PRECISION(&_Bulla.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bulla *BullaCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bulla *BullaSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bulla.Contract.Allowance(&_Bulla.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Bulla *BullaCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Bulla.Contract.Allowance(&_Bulla.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bulla *BullaCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bulla *BullaSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bulla.Contract.BalanceOf(&_Bulla.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Bulla *BullaCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Bulla.Contract.BalanceOf(&_Bulla.CallOpts, account)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_Bulla *BullaCaller) BaseLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "baseLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_Bulla *BullaSession) BaseLower() (*big.Int, error) {
	return _Bulla.Contract.BaseLower(&_Bulla.CallOpts)
}

// BaseLower is a free data retrieval call binding the contract method 0xfa082743.
//
// Solidity: function baseLower() view returns(int24)
func (_Bulla *BullaCallerSession) BaseLower() (*big.Int, error) {
	return _Bulla.Contract.BaseLower(&_Bulla.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_Bulla *BullaCaller) BaseUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "baseUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_Bulla *BullaSession) BaseUpper() (*big.Int, error) {
	return _Bulla.Contract.BaseUpper(&_Bulla.CallOpts)
}

// BaseUpper is a free data retrieval call binding the contract method 0x888a9134.
//
// Solidity: function baseUpper() view returns(int24)
func (_Bulla *BullaCallerSession) BaseUpper() (*big.Int, error) {
	return _Bulla.Contract.BaseUpper(&_Bulla.CallOpts)
}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_Bulla *BullaCaller) CurrentTick(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "currentTick")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_Bulla *BullaSession) CurrentTick() (*big.Int, error) {
	return _Bulla.Contract.CurrentTick(&_Bulla.CallOpts)
}

// CurrentTick is a free data retrieval call binding the contract method 0x065e5360.
//
// Solidity: function currentTick() view returns(int24 tick)
func (_Bulla *BullaCallerSession) CurrentTick() (*big.Int, error) {
	return _Bulla.Contract.CurrentTick(&_Bulla.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bulla *BullaCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bulla *BullaSession) Decimals() (uint8, error) {
	return _Bulla.Contract.Decimals(&_Bulla.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Bulla *BullaCallerSession) Decimals() (uint8, error) {
	return _Bulla.Contract.Decimals(&_Bulla.CallOpts)
}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_Bulla *BullaCaller) Deposit0Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "deposit0Max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_Bulla *BullaSession) Deposit0Max() (*big.Int, error) {
	return _Bulla.Contract.Deposit0Max(&_Bulla.CallOpts)
}

// Deposit0Max is a free data retrieval call binding the contract method 0x648cab85.
//
// Solidity: function deposit0Max() view returns(uint256)
func (_Bulla *BullaCallerSession) Deposit0Max() (*big.Int, error) {
	return _Bulla.Contract.Deposit0Max(&_Bulla.CallOpts)
}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_Bulla *BullaCaller) Deposit1Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "deposit1Max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_Bulla *BullaSession) Deposit1Max() (*big.Int, error) {
	return _Bulla.Contract.Deposit1Max(&_Bulla.CallOpts)
}

// Deposit1Max is a free data retrieval call binding the contract method 0x4d461fbb.
//
// Solidity: function deposit1Max() view returns(uint256)
func (_Bulla *BullaCallerSession) Deposit1Max() (*big.Int, error) {
	return _Bulla.Contract.Deposit1Max(&_Bulla.CallOpts)
}

// DirectDeposit is a free data retrieval call binding the contract method 0x6d90a39c.
//
// Solidity: function directDeposit() view returns(bool)
func (_Bulla *BullaCaller) DirectDeposit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "directDeposit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DirectDeposit is a free data retrieval call binding the contract method 0x6d90a39c.
//
// Solidity: function directDeposit() view returns(bool)
func (_Bulla *BullaSession) DirectDeposit() (bool, error) {
	return _Bulla.Contract.DirectDeposit(&_Bulla.CallOpts)
}

// DirectDeposit is a free data retrieval call binding the contract method 0x6d90a39c.
//
// Solidity: function directDeposit() view returns(bool)
func (_Bulla *BullaCallerSession) DirectDeposit() (bool, error) {
	return _Bulla.Contract.DirectDeposit(&_Bulla.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Bulla *BullaCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "eip712Domain")

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
func (_Bulla *BullaSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Bulla.Contract.Eip712Domain(&_Bulla.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Bulla *BullaCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Bulla.Contract.Eip712Domain(&_Bulla.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint8)
func (_Bulla *BullaCaller) Fee(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint8)
func (_Bulla *BullaSession) Fee() (uint8, error) {
	return _Bulla.Contract.Fee(&_Bulla.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint8)
func (_Bulla *BullaCallerSession) Fee() (uint8, error) {
	return _Bulla.Contract.Fee(&_Bulla.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Bulla *BullaCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Bulla *BullaSession) FeeRecipient() (common.Address, error) {
	return _Bulla.Contract.FeeRecipient(&_Bulla.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Bulla *BullaCallerSession) FeeRecipient() (common.Address, error) {
	return _Bulla.Contract.FeeRecipient(&_Bulla.CallOpts)
}

// GetBasePosition is a free data retrieval call binding the contract method 0xd2eabcfc.
//
// Solidity: function getBasePosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Bulla *BullaCaller) GetBasePosition(opts *bind.CallOpts) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "getBasePosition")

	outstruct := new(struct {
		Liquidity *big.Int
		Amount0   *big.Int
		Amount1   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount0 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBasePosition is a free data retrieval call binding the contract method 0xd2eabcfc.
//
// Solidity: function getBasePosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Bulla *BullaSession) GetBasePosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _Bulla.Contract.GetBasePosition(&_Bulla.CallOpts)
}

// GetBasePosition is a free data retrieval call binding the contract method 0xd2eabcfc.
//
// Solidity: function getBasePosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Bulla *BullaCallerSession) GetBasePosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _Bulla.Contract.GetBasePosition(&_Bulla.CallOpts)
}

// GetLimitPosition is a free data retrieval call binding the contract method 0xa049de6b.
//
// Solidity: function getLimitPosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Bulla *BullaCaller) GetLimitPosition(opts *bind.CallOpts) (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "getLimitPosition")

	outstruct := new(struct {
		Liquidity *big.Int
		Amount0   *big.Int
		Amount1   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount0 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetLimitPosition is a free data retrieval call binding the contract method 0xa049de6b.
//
// Solidity: function getLimitPosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Bulla *BullaSession) GetLimitPosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _Bulla.Contract.GetLimitPosition(&_Bulla.CallOpts)
}

// GetLimitPosition is a free data retrieval call binding the contract method 0xa049de6b.
//
// Solidity: function getLimitPosition() view returns(uint128 liquidity, uint256 amount0, uint256 amount1)
func (_Bulla *BullaCallerSession) GetLimitPosition() (struct {
	Liquidity *big.Int
	Amount0   *big.Int
	Amount1   *big.Int
}, error) {
	return _Bulla.Contract.GetLimitPosition(&_Bulla.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_Bulla *BullaCaller) GetTotalAmounts(opts *bind.CallOpts) (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "getTotalAmounts")

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

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_Bulla *BullaSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _Bulla.Contract.GetTotalAmounts(&_Bulla.CallOpts)
}

// GetTotalAmounts is a free data retrieval call binding the contract method 0xc4a7761e.
//
// Solidity: function getTotalAmounts() view returns(uint256 total0, uint256 total1)
func (_Bulla *BullaCallerSession) GetTotalAmounts() (struct {
	Total0 *big.Int
	Total1 *big.Int
}, error) {
	return _Bulla.Contract.GetTotalAmounts(&_Bulla.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_Bulla *BullaCaller) LimitLower(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "limitLower")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_Bulla *BullaSession) LimitLower() (*big.Int, error) {
	return _Bulla.Contract.LimitLower(&_Bulla.CallOpts)
}

// LimitLower is a free data retrieval call binding the contract method 0x51e87af7.
//
// Solidity: function limitLower() view returns(int24)
func (_Bulla *BullaCallerSession) LimitLower() (*big.Int, error) {
	return _Bulla.Contract.LimitLower(&_Bulla.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_Bulla *BullaCaller) LimitUpper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "limitUpper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_Bulla *BullaSession) LimitUpper() (*big.Int, error) {
	return _Bulla.Contract.LimitUpper(&_Bulla.CallOpts)
}

// LimitUpper is a free data retrieval call binding the contract method 0x0f35bcac.
//
// Solidity: function limitUpper() view returns(int24)
func (_Bulla *BullaCallerSession) LimitUpper() (*big.Int, error) {
	return _Bulla.Contract.LimitUpper(&_Bulla.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Bulla *BullaCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Bulla *BullaSession) MaxTotalSupply() (*big.Int, error) {
	return _Bulla.Contract.MaxTotalSupply(&_Bulla.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_Bulla *BullaCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _Bulla.Contract.MaxTotalSupply(&_Bulla.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bulla *BullaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bulla *BullaSession) Name() (string, error) {
	return _Bulla.Contract.Name(&_Bulla.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bulla *BullaCallerSession) Name() (string, error) {
	return _Bulla.Contract.Name(&_Bulla.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bulla *BullaCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bulla *BullaSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Bulla.Contract.Nonces(&_Bulla.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Bulla *BullaCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Bulla.Contract.Nonces(&_Bulla.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bulla *BullaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bulla *BullaSession) Owner() (common.Address, error) {
	return _Bulla.Contract.Owner(&_Bulla.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bulla *BullaCallerSession) Owner() (common.Address, error) {
	return _Bulla.Contract.Owner(&_Bulla.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Bulla *BullaCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Bulla *BullaSession) Pool() (common.Address, error) {
	return _Bulla.Contract.Pool(&_Bulla.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Bulla *BullaCallerSession) Pool() (common.Address, error) {
	return _Bulla.Contract.Pool(&_Bulla.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bulla *BullaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bulla *BullaSession) Symbol() (string, error) {
	return _Bulla.Contract.Symbol(&_Bulla.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bulla *BullaCallerSession) Symbol() (string, error) {
	return _Bulla.Contract.Symbol(&_Bulla.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Bulla *BullaCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Bulla *BullaSession) TickSpacing() (*big.Int, error) {
	return _Bulla.Contract.TickSpacing(&_Bulla.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(int24)
func (_Bulla *BullaCallerSession) TickSpacing() (*big.Int, error) {
	return _Bulla.Contract.TickSpacing(&_Bulla.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Bulla *BullaCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Bulla *BullaSession) Token0() (common.Address, error) {
	return _Bulla.Contract.Token0(&_Bulla.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Bulla *BullaCallerSession) Token0() (common.Address, error) {
	return _Bulla.Contract.Token0(&_Bulla.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Bulla *BullaCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Bulla *BullaSession) Token1() (common.Address, error) {
	return _Bulla.Contract.Token1(&_Bulla.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Bulla *BullaCallerSession) Token1() (common.Address, error) {
	return _Bulla.Contract.Token1(&_Bulla.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bulla *BullaCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bulla *BullaSession) TotalSupply() (*big.Int, error) {
	return _Bulla.Contract.TotalSupply(&_Bulla.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bulla *BullaCallerSession) TotalSupply() (*big.Int, error) {
	return _Bulla.Contract.TotalSupply(&_Bulla.CallOpts)
}

// WhitelistedAddress is a free data retrieval call binding the contract method 0x86a29081.
//
// Solidity: function whitelistedAddress() view returns(address)
func (_Bulla *BullaCaller) WhitelistedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bulla.contract.Call(opts, &out, "whitelistedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedAddress is a free data retrieval call binding the contract method 0x86a29081.
//
// Solidity: function whitelistedAddress() view returns(address)
func (_Bulla *BullaSession) WhitelistedAddress() (common.Address, error) {
	return _Bulla.Contract.WhitelistedAddress(&_Bulla.CallOpts)
}

// WhitelistedAddress is a free data retrieval call binding the contract method 0x86a29081.
//
// Solidity: function whitelistedAddress() view returns(address)
func (_Bulla *BullaCallerSession) WhitelistedAddress() (common.Address, error) {
	return _Bulla.Contract.WhitelistedAddress(&_Bulla.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x63e96836.
//
// Solidity: function addLiquidity(int24 tickLower, int24 tickUpper, uint256 amount0, uint256 amount1, uint256[2] inMin) returns()
func (_Bulla *BullaTransactor) AddLiquidity(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, amount0 *big.Int, amount1 *big.Int, inMin [2]*big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "addLiquidity", tickLower, tickUpper, amount0, amount1, inMin)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x63e96836.
//
// Solidity: function addLiquidity(int24 tickLower, int24 tickUpper, uint256 amount0, uint256 amount1, uint256[2] inMin) returns()
func (_Bulla *BullaSession) AddLiquidity(tickLower *big.Int, tickUpper *big.Int, amount0 *big.Int, amount1 *big.Int, inMin [2]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.AddLiquidity(&_Bulla.TransactOpts, tickLower, tickUpper, amount0, amount1, inMin)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x63e96836.
//
// Solidity: function addLiquidity(int24 tickLower, int24 tickUpper, uint256 amount0, uint256 amount1, uint256[2] inMin) returns()
func (_Bulla *BullaTransactorSession) AddLiquidity(tickLower *big.Int, tickUpper *big.Int, amount0 *big.Int, amount1 *big.Int, inMin [2]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.AddLiquidity(&_Bulla.TransactOpts, tickLower, tickUpper, amount0, amount1, inMin)
}

// AlgebraMintCallback is a paid mutator transaction binding the contract method 0x3dd657c5.
//
// Solidity: function algebraMintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Bulla *BullaTransactor) AlgebraMintCallback(opts *bind.TransactOpts, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "algebraMintCallback", amount0, amount1, data)
}

// AlgebraMintCallback is a paid mutator transaction binding the contract method 0x3dd657c5.
//
// Solidity: function algebraMintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Bulla *BullaSession) AlgebraMintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bulla.Contract.AlgebraMintCallback(&_Bulla.TransactOpts, amount0, amount1, data)
}

// AlgebraMintCallback is a paid mutator transaction binding the contract method 0x3dd657c5.
//
// Solidity: function algebraMintCallback(uint256 amount0, uint256 amount1, bytes data) returns()
func (_Bulla *BullaTransactorSession) AlgebraMintCallback(amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Bulla.Contract.AlgebraMintCallback(&_Bulla.TransactOpts, amount0, amount1, data)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bulla *BullaTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bulla *BullaSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Approve(&_Bulla.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Bulla *BullaTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Approve(&_Bulla.TransactOpts, spender, amount)
}

// Compound is a paid mutator transaction binding the contract method 0x513ea884.
//
// Solidity: function compound(uint256[4] inMin) returns(uint128 baseToken0Owed, uint128 baseToken1Owed, uint128 limitToken0Owed, uint128 limitToken1Owed)
func (_Bulla *BullaTransactor) Compound(opts *bind.TransactOpts, inMin [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "compound", inMin)
}

// Compound is a paid mutator transaction binding the contract method 0x513ea884.
//
// Solidity: function compound(uint256[4] inMin) returns(uint128 baseToken0Owed, uint128 baseToken1Owed, uint128 limitToken0Owed, uint128 limitToken1Owed)
func (_Bulla *BullaSession) Compound(inMin [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Compound(&_Bulla.TransactOpts, inMin)
}

// Compound is a paid mutator transaction binding the contract method 0x513ea884.
//
// Solidity: function compound(uint256[4] inMin) returns(uint128 baseToken0Owed, uint128 baseToken1Owed, uint128 limitToken0Owed, uint128 limitToken1Owed)
func (_Bulla *BullaTransactorSession) Compound(inMin [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Compound(&_Bulla.TransactOpts, inMin)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bulla *BullaTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bulla *BullaSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.DecreaseAllowance(&_Bulla.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Bulla *BullaTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.DecreaseAllowance(&_Bulla.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x8e3c92e4.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to, address from, uint256[4] inMin) returns(uint256 shares)
func (_Bulla *BullaTransactor) Deposit(opts *bind.TransactOpts, deposit0 *big.Int, deposit1 *big.Int, to common.Address, from common.Address, inMin [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "deposit", deposit0, deposit1, to, from, inMin)
}

// Deposit is a paid mutator transaction binding the contract method 0x8e3c92e4.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to, address from, uint256[4] inMin) returns(uint256 shares)
func (_Bulla *BullaSession) Deposit(deposit0 *big.Int, deposit1 *big.Int, to common.Address, from common.Address, inMin [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Deposit(&_Bulla.TransactOpts, deposit0, deposit1, to, from, inMin)
}

// Deposit is a paid mutator transaction binding the contract method 0x8e3c92e4.
//
// Solidity: function deposit(uint256 deposit0, uint256 deposit1, address to, address from, uint256[4] inMin) returns(uint256 shares)
func (_Bulla *BullaTransactorSession) Deposit(deposit0 *big.Int, deposit1 *big.Int, to common.Address, from common.Address, inMin [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Deposit(&_Bulla.TransactOpts, deposit0, deposit1, to, from, inMin)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bulla *BullaTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bulla *BullaSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.IncreaseAllowance(&_Bulla.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Bulla *BullaTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.IncreaseAllowance(&_Bulla.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bulla *BullaTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bulla *BullaSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bulla.Contract.Permit(&_Bulla.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Bulla *BullaTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bulla.Contract.Permit(&_Bulla.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// PullLiquidity is a paid mutator transaction binding the contract method 0x95235656.
//
// Solidity: function pullLiquidity(int24 tickLower, int24 tickUpper, uint128 shares, uint256[2] amountMin) returns(uint256 amount0, uint256 amount1)
func (_Bulla *BullaTransactor) PullLiquidity(opts *bind.TransactOpts, tickLower *big.Int, tickUpper *big.Int, shares *big.Int, amountMin [2]*big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "pullLiquidity", tickLower, tickUpper, shares, amountMin)
}

// PullLiquidity is a paid mutator transaction binding the contract method 0x95235656.
//
// Solidity: function pullLiquidity(int24 tickLower, int24 tickUpper, uint128 shares, uint256[2] amountMin) returns(uint256 amount0, uint256 amount1)
func (_Bulla *BullaSession) PullLiquidity(tickLower *big.Int, tickUpper *big.Int, shares *big.Int, amountMin [2]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.PullLiquidity(&_Bulla.TransactOpts, tickLower, tickUpper, shares, amountMin)
}

// PullLiquidity is a paid mutator transaction binding the contract method 0x95235656.
//
// Solidity: function pullLiquidity(int24 tickLower, int24 tickUpper, uint128 shares, uint256[2] amountMin) returns(uint256 amount0, uint256 amount1)
func (_Bulla *BullaTransactorSession) PullLiquidity(tickLower *big.Int, tickUpper *big.Int, shares *big.Int, amountMin [2]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.PullLiquidity(&_Bulla.TransactOpts, tickLower, tickUpper, shares, amountMin)
}

// Rebalance is a paid mutator transaction binding the contract method 0x85919c5d.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, address _feeRecipient, uint256[4] inMin, uint256[4] outMin) returns()
func (_Bulla *BullaTransactor) Rebalance(opts *bind.TransactOpts, _baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, _feeRecipient common.Address, inMin [4]*big.Int, outMin [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "rebalance", _baseLower, _baseUpper, _limitLower, _limitUpper, _feeRecipient, inMin, outMin)
}

// Rebalance is a paid mutator transaction binding the contract method 0x85919c5d.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, address _feeRecipient, uint256[4] inMin, uint256[4] outMin) returns()
func (_Bulla *BullaSession) Rebalance(_baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, _feeRecipient common.Address, inMin [4]*big.Int, outMin [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Rebalance(&_Bulla.TransactOpts, _baseLower, _baseUpper, _limitLower, _limitUpper, _feeRecipient, inMin, outMin)
}

// Rebalance is a paid mutator transaction binding the contract method 0x85919c5d.
//
// Solidity: function rebalance(int24 _baseLower, int24 _baseUpper, int24 _limitLower, int24 _limitUpper, address _feeRecipient, uint256[4] inMin, uint256[4] outMin) returns()
func (_Bulla *BullaTransactorSession) Rebalance(_baseLower *big.Int, _baseUpper *big.Int, _limitLower *big.Int, _limitUpper *big.Int, _feeRecipient common.Address, inMin [4]*big.Int, outMin [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Rebalance(&_Bulla.TransactOpts, _baseLower, _baseUpper, _limitLower, _limitUpper, _feeRecipient, inMin, outMin)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0xc5241e29.
//
// Solidity: function removeWhitelisted() returns()
func (_Bulla *BullaTransactor) RemoveWhitelisted(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "removeWhitelisted")
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0xc5241e29.
//
// Solidity: function removeWhitelisted() returns()
func (_Bulla *BullaSession) RemoveWhitelisted() (*types.Transaction, error) {
	return _Bulla.Contract.RemoveWhitelisted(&_Bulla.TransactOpts)
}

// RemoveWhitelisted is a paid mutator transaction binding the contract method 0xc5241e29.
//
// Solidity: function removeWhitelisted() returns()
func (_Bulla *BullaTransactorSession) RemoveWhitelisted() (*types.Transaction, error) {
	return _Bulla.Contract.RemoveWhitelisted(&_Bulla.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0xcb122a09.
//
// Solidity: function setFee(uint8 newFee) returns()
func (_Bulla *BullaTransactor) SetFee(opts *bind.TransactOpts, newFee uint8) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0xcb122a09.
//
// Solidity: function setFee(uint8 newFee) returns()
func (_Bulla *BullaSession) SetFee(newFee uint8) (*types.Transaction, error) {
	return _Bulla.Contract.SetFee(&_Bulla.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0xcb122a09.
//
// Solidity: function setFee(uint8 newFee) returns()
func (_Bulla *BullaTransactorSession) SetFee(newFee uint8) (*types.Transaction, error) {
	return _Bulla.Contract.SetFee(&_Bulla.TransactOpts, newFee)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_Bulla *BullaTransactor) SetTickSpacing(opts *bind.TransactOpts, newTickSpacing *big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "setTickSpacing", newTickSpacing)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_Bulla *BullaSession) SetTickSpacing(newTickSpacing *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.SetTickSpacing(&_Bulla.TransactOpts, newTickSpacing)
}

// SetTickSpacing is a paid mutator transaction binding the contract method 0xf085a610.
//
// Solidity: function setTickSpacing(int24 newTickSpacing) returns()
func (_Bulla *BullaTransactorSession) SetTickSpacing(newTickSpacing *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.SetTickSpacing(&_Bulla.TransactOpts, newTickSpacing)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _address) returns()
func (_Bulla *BullaTransactor) SetWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "setWhitelist", _address)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _address) returns()
func (_Bulla *BullaSession) SetWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Bulla.Contract.SetWhitelist(&_Bulla.TransactOpts, _address)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(address _address) returns()
func (_Bulla *BullaTransactorSession) SetWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Bulla.Contract.SetWhitelist(&_Bulla.TransactOpts, _address)
}

// ToggleDirectDeposit is a paid mutator transaction binding the contract method 0xb1a3d533.
//
// Solidity: function toggleDirectDeposit() returns()
func (_Bulla *BullaTransactor) ToggleDirectDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "toggleDirectDeposit")
}

// ToggleDirectDeposit is a paid mutator transaction binding the contract method 0xb1a3d533.
//
// Solidity: function toggleDirectDeposit() returns()
func (_Bulla *BullaSession) ToggleDirectDeposit() (*types.Transaction, error) {
	return _Bulla.Contract.ToggleDirectDeposit(&_Bulla.TransactOpts)
}

// ToggleDirectDeposit is a paid mutator transaction binding the contract method 0xb1a3d533.
//
// Solidity: function toggleDirectDeposit() returns()
func (_Bulla *BullaTransactorSession) ToggleDirectDeposit() (*types.Transaction, error) {
	return _Bulla.Contract.ToggleDirectDeposit(&_Bulla.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bulla *BullaTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bulla *BullaSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Transfer(&_Bulla.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Bulla *BullaTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Transfer(&_Bulla.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bulla *BullaTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bulla *BullaSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.TransferFrom(&_Bulla.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Bulla *BullaTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.TransferFrom(&_Bulla.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bulla *BullaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bulla *BullaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bulla.Contract.TransferOwnership(&_Bulla.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bulla *BullaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bulla.Contract.TransferOwnership(&_Bulla.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa8559872.
//
// Solidity: function withdraw(uint256 shares, address to, address from, uint256[4] minAmounts) returns(uint256 amount0, uint256 amount1)
func (_Bulla *BullaTransactor) Withdraw(opts *bind.TransactOpts, shares *big.Int, to common.Address, from common.Address, minAmounts [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "withdraw", shares, to, from, minAmounts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa8559872.
//
// Solidity: function withdraw(uint256 shares, address to, address from, uint256[4] minAmounts) returns(uint256 amount0, uint256 amount1)
func (_Bulla *BullaSession) Withdraw(shares *big.Int, to common.Address, from common.Address, minAmounts [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Withdraw(&_Bulla.TransactOpts, shares, to, from, minAmounts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa8559872.
//
// Solidity: function withdraw(uint256 shares, address to, address from, uint256[4] minAmounts) returns(uint256 amount0, uint256 amount1)
func (_Bulla *BullaTransactorSession) Withdraw(shares *big.Int, to common.Address, from common.Address, minAmounts [4]*big.Int) (*types.Transaction, error) {
	return _Bulla.Contract.Withdraw(&_Bulla.TransactOpts, shares, to, from, minAmounts)
}

// ZeroBurn is a paid mutator transaction binding the contract method 0xcce6776e.
//
// Solidity: function zeroBurn() returns(uint128 baseLiquidity, uint128 limitLiquidity)
func (_Bulla *BullaTransactor) ZeroBurn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bulla.contract.Transact(opts, "zeroBurn")
}

// ZeroBurn is a paid mutator transaction binding the contract method 0xcce6776e.
//
// Solidity: function zeroBurn() returns(uint128 baseLiquidity, uint128 limitLiquidity)
func (_Bulla *BullaSession) ZeroBurn() (*types.Transaction, error) {
	return _Bulla.Contract.ZeroBurn(&_Bulla.TransactOpts)
}

// ZeroBurn is a paid mutator transaction binding the contract method 0xcce6776e.
//
// Solidity: function zeroBurn() returns(uint128 baseLiquidity, uint128 limitLiquidity)
func (_Bulla *BullaTransactorSession) ZeroBurn() (*types.Transaction, error) {
	return _Bulla.Contract.ZeroBurn(&_Bulla.TransactOpts)
}

// BullaApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bulla contract.
type BullaApprovalIterator struct {
	Event *BullaApproval // Event containing the contract specifics and raw log

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
func (it *BullaApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BullaApproval)
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
		it.Event = new(BullaApproval)
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
func (it *BullaApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BullaApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BullaApproval represents a Approval event raised by the Bulla contract.
type BullaApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bulla *BullaFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BullaApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bulla.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BullaApprovalIterator{contract: _Bulla.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Bulla *BullaFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BullaApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Bulla.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BullaApproval)
				if err := _Bulla.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Bulla *BullaFilterer) ParseApproval(log types.Log) (*BullaApproval, error) {
	event := new(BullaApproval)
	if err := _Bulla.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BullaDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Bulla contract.
type BullaDepositIterator struct {
	Event *BullaDeposit // Event containing the contract specifics and raw log

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
func (it *BullaDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BullaDeposit)
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
		it.Event = new(BullaDeposit)
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
func (it *BullaDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BullaDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BullaDeposit represents a Deposit event raised by the Bulla contract.
type BullaDeposit struct {
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
func (_Bulla *BullaFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*BullaDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bulla.contract.FilterLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BullaDepositIterator{contract: _Bulla.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x4e2ca0515ed1aef1395f66b5303bb5d6f1bf9d61a353fa53f73f8ac9973fa9f6.
//
// Solidity: event Deposit(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Bulla *BullaFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BullaDeposit, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bulla.contract.WatchLogs(opts, "Deposit", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BullaDeposit)
				if err := _Bulla.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Bulla *BullaFilterer) ParseDeposit(log types.Log) (*BullaDeposit, error) {
	event := new(BullaDeposit)
	if err := _Bulla.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BullaEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Bulla contract.
type BullaEIP712DomainChangedIterator struct {
	Event *BullaEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *BullaEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BullaEIP712DomainChanged)
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
		it.Event = new(BullaEIP712DomainChanged)
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
func (it *BullaEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BullaEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BullaEIP712DomainChanged represents a EIP712DomainChanged event raised by the Bulla contract.
type BullaEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Bulla *BullaFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BullaEIP712DomainChangedIterator, error) {

	logs, sub, err := _Bulla.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BullaEIP712DomainChangedIterator{contract: _Bulla.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Bulla *BullaFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BullaEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Bulla.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BullaEIP712DomainChanged)
				if err := _Bulla.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_Bulla *BullaFilterer) ParseEIP712DomainChanged(log types.Log) (*BullaEIP712DomainChanged, error) {
	event := new(BullaEIP712DomainChanged)
	if err := _Bulla.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BullaRebalanceIterator is returned from FilterRebalance and is used to iterate over the raw logs and unpacked data for Rebalance events raised by the Bulla contract.
type BullaRebalanceIterator struct {
	Event *BullaRebalance // Event containing the contract specifics and raw log

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
func (it *BullaRebalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BullaRebalance)
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
		it.Event = new(BullaRebalance)
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
func (it *BullaRebalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BullaRebalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BullaRebalance represents a Rebalance event raised by the Bulla contract.
type BullaRebalance struct {
	Tick         *big.Int
	TotalAmount0 *big.Int
	TotalAmount1 *big.Int
	FeeAmount0   *big.Int
	FeeAmount1   *big.Int
	TotalSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRebalance is a free log retrieval operation binding the contract event 0xbc4c20ad04f161d631d9ce94d27659391196415aa3c42f6a71c62e905ece782d.
//
// Solidity: event Rebalance(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 feeAmount0, uint256 feeAmount1, uint256 totalSupply)
func (_Bulla *BullaFilterer) FilterRebalance(opts *bind.FilterOpts) (*BullaRebalanceIterator, error) {

	logs, sub, err := _Bulla.contract.FilterLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return &BullaRebalanceIterator{contract: _Bulla.contract, event: "Rebalance", logs: logs, sub: sub}, nil
}

// WatchRebalance is a free log subscription operation binding the contract event 0xbc4c20ad04f161d631d9ce94d27659391196415aa3c42f6a71c62e905ece782d.
//
// Solidity: event Rebalance(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 feeAmount0, uint256 feeAmount1, uint256 totalSupply)
func (_Bulla *BullaFilterer) WatchRebalance(opts *bind.WatchOpts, sink chan<- *BullaRebalance) (event.Subscription, error) {

	logs, sub, err := _Bulla.contract.WatchLogs(opts, "Rebalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BullaRebalance)
				if err := _Bulla.contract.UnpackLog(event, "Rebalance", log); err != nil {
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

// ParseRebalance is a log parse operation binding the contract event 0xbc4c20ad04f161d631d9ce94d27659391196415aa3c42f6a71c62e905ece782d.
//
// Solidity: event Rebalance(int24 tick, uint256 totalAmount0, uint256 totalAmount1, uint256 feeAmount0, uint256 feeAmount1, uint256 totalSupply)
func (_Bulla *BullaFilterer) ParseRebalance(log types.Log) (*BullaRebalance, error) {
	event := new(BullaRebalance)
	if err := _Bulla.contract.UnpackLog(event, "Rebalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BullaSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the Bulla contract.
type BullaSetFeeIterator struct {
	Event *BullaSetFee // Event containing the contract specifics and raw log

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
func (it *BullaSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BullaSetFee)
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
		it.Event = new(BullaSetFee)
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
func (it *BullaSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BullaSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BullaSetFee represents a SetFee event raised by the Bulla contract.
type BullaSetFee struct {
	NewFee uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x91f2ade82ab0e77bb6823899e6daddc07e3da0e3ad998577e7c09c2f38943c43.
//
// Solidity: event SetFee(uint8 newFee)
func (_Bulla *BullaFilterer) FilterSetFee(opts *bind.FilterOpts) (*BullaSetFeeIterator, error) {

	logs, sub, err := _Bulla.contract.FilterLogs(opts, "SetFee")
	if err != nil {
		return nil, err
	}
	return &BullaSetFeeIterator{contract: _Bulla.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x91f2ade82ab0e77bb6823899e6daddc07e3da0e3ad998577e7c09c2f38943c43.
//
// Solidity: event SetFee(uint8 newFee)
func (_Bulla *BullaFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *BullaSetFee) (event.Subscription, error) {

	logs, sub, err := _Bulla.contract.WatchLogs(opts, "SetFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BullaSetFee)
				if err := _Bulla.contract.UnpackLog(event, "SetFee", log); err != nil {
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

// ParseSetFee is a log parse operation binding the contract event 0x91f2ade82ab0e77bb6823899e6daddc07e3da0e3ad998577e7c09c2f38943c43.
//
// Solidity: event SetFee(uint8 newFee)
func (_Bulla *BullaFilterer) ParseSetFee(log types.Log) (*BullaSetFee, error) {
	event := new(BullaSetFee)
	if err := _Bulla.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BullaTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bulla contract.
type BullaTransferIterator struct {
	Event *BullaTransfer // Event containing the contract specifics and raw log

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
func (it *BullaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BullaTransfer)
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
		it.Event = new(BullaTransfer)
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
func (it *BullaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BullaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BullaTransfer represents a Transfer event raised by the Bulla contract.
type BullaTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bulla *BullaFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BullaTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bulla.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BullaTransferIterator{contract: _Bulla.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Bulla *BullaFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BullaTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bulla.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BullaTransfer)
				if err := _Bulla.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Bulla *BullaFilterer) ParseTransfer(log types.Log) (*BullaTransfer, error) {
	event := new(BullaTransfer)
	if err := _Bulla.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BullaWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Bulla contract.
type BullaWithdrawIterator struct {
	Event *BullaWithdraw // Event containing the contract specifics and raw log

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
func (it *BullaWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BullaWithdraw)
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
		it.Event = new(BullaWithdraw)
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
func (it *BullaWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BullaWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BullaWithdraw represents a Withdraw event raised by the Bulla contract.
type BullaWithdraw struct {
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
func (_Bulla *BullaFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*BullaWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bulla.contract.FilterLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BullaWithdrawIterator{contract: _Bulla.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed sender, address indexed to, uint256 shares, uint256 amount0, uint256 amount1)
func (_Bulla *BullaFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BullaWithdraw, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bulla.contract.WatchLogs(opts, "Withdraw", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BullaWithdraw)
				if err := _Bulla.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Bulla *BullaFilterer) ParseWithdraw(log types.Log) (*BullaWithdraw, error) {
	event := new(BullaWithdraw)
	if err := _Bulla.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BullaZeroBurnIterator is returned from FilterZeroBurn and is used to iterate over the raw logs and unpacked data for ZeroBurn events raised by the Bulla contract.
type BullaZeroBurnIterator struct {
	Event *BullaZeroBurn // Event containing the contract specifics and raw log

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
func (it *BullaZeroBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BullaZeroBurn)
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
		it.Event = new(BullaZeroBurn)
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
func (it *BullaZeroBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BullaZeroBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BullaZeroBurn represents a ZeroBurn event raised by the Bulla contract.
type BullaZeroBurn struct {
	Fee   uint8
	Fees0 *big.Int
	Fees1 *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterZeroBurn is a free log retrieval operation binding the contract event 0x4606b8a47eb284e8e80929101ece6ab5fe8d4f8735acc56bd0c92ca872f2cfe7.
//
// Solidity: event ZeroBurn(uint8 fee, uint256 fees0, uint256 fees1)
func (_Bulla *BullaFilterer) FilterZeroBurn(opts *bind.FilterOpts) (*BullaZeroBurnIterator, error) {

	logs, sub, err := _Bulla.contract.FilterLogs(opts, "ZeroBurn")
	if err != nil {
		return nil, err
	}
	return &BullaZeroBurnIterator{contract: _Bulla.contract, event: "ZeroBurn", logs: logs, sub: sub}, nil
}

// WatchZeroBurn is a free log subscription operation binding the contract event 0x4606b8a47eb284e8e80929101ece6ab5fe8d4f8735acc56bd0c92ca872f2cfe7.
//
// Solidity: event ZeroBurn(uint8 fee, uint256 fees0, uint256 fees1)
func (_Bulla *BullaFilterer) WatchZeroBurn(opts *bind.WatchOpts, sink chan<- *BullaZeroBurn) (event.Subscription, error) {

	logs, sub, err := _Bulla.contract.WatchLogs(opts, "ZeroBurn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BullaZeroBurn)
				if err := _Bulla.contract.UnpackLog(event, "ZeroBurn", log); err != nil {
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

// ParseZeroBurn is a log parse operation binding the contract event 0x4606b8a47eb284e8e80929101ece6ab5fe8d4f8735acc56bd0c92ca872f2cfe7.
//
// Solidity: event ZeroBurn(uint8 fee, uint256 fees0, uint256 fees1)
func (_Bulla *BullaFilterer) ParseZeroBurn(log types.Log) (*BullaZeroBurn, error) {
	event := new(BullaZeroBurn)
	if err := _Bulla.contract.UnpackLog(event, "ZeroBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
