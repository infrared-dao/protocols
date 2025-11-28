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

// ERC4626VaultMetaData contains all meta data concerning the ERC4626Vault contract.
var ERC4626VaultMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOLOMITE_MARGIN\",\"outputs\":[{\"internalType\":\"contractIDolomiteMargin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOLOMITE_MARGIN_OWNER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dolomiteRegistry\",\"outputs\":[{\"internalType\":\"contractIDolomiteRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_marketId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dolomiteRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"isValidReceiver\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// DOLOMITEMARGIN is a free data retrieval call binding the contract method 0x15c14a4a.
//
// Solidity: function DOLOMITE_MARGIN() view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) DOLOMITEMARGIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "DOLOMITE_MARGIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DOLOMITEMARGIN is a free data retrieval call binding the contract method 0x15c14a4a.
//
// Solidity: function DOLOMITE_MARGIN() view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) DOLOMITEMARGIN() (common.Address, error) {
	return _ERC4626Vault.Contract.DOLOMITEMARGIN(&_ERC4626Vault.CallOpts)
}

// DOLOMITEMARGIN is a free data retrieval call binding the contract method 0x15c14a4a.
//
// Solidity: function DOLOMITE_MARGIN() view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) DOLOMITEMARGIN() (common.Address, error) {
	return _ERC4626Vault.Contract.DOLOMITEMARGIN(&_ERC4626Vault.CallOpts)
}

// DOLOMITEMARGINOWNER is a free data retrieval call binding the contract method 0xcbffd921.
//
// Solidity: function DOLOMITE_MARGIN_OWNER() view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) DOLOMITEMARGINOWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "DOLOMITE_MARGIN_OWNER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DOLOMITEMARGINOWNER is a free data retrieval call binding the contract method 0xcbffd921.
//
// Solidity: function DOLOMITE_MARGIN_OWNER() view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) DOLOMITEMARGINOWNER() (common.Address, error) {
	return _ERC4626Vault.Contract.DOLOMITEMARGINOWNER(&_ERC4626Vault.CallOpts)
}

// DOLOMITEMARGINOWNER is a free data retrieval call binding the contract method 0xcbffd921.
//
// Solidity: function DOLOMITE_MARGIN_OWNER() view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) DOLOMITEMARGINOWNER() (common.Address, error) {
	return _ERC4626Vault.Contract.DOLOMITEMARGINOWNER(&_ERC4626Vault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.Allowance(&_ERC4626Vault.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.Allowance(&_ERC4626Vault.CallOpts, _owner, _spender)
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

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.BalanceOf(&_ERC4626Vault.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.BalanceOf(&_ERC4626Vault.CallOpts, _account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) ConvertToAssets(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "convertToAssets", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.ConvertToAssets(&_ERC4626Vault.CallOpts, _shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.ConvertToAssets(&_ERC4626Vault.CallOpts, _shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) ConvertToShares(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "convertToShares", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.ConvertToShares(&_ERC4626Vault.CallOpts, _assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.ConvertToShares(&_ERC4626Vault.CallOpts, _assets)
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

// DolomiteRegistry is a free data retrieval call binding the contract method 0xbd12584f.
//
// Solidity: function dolomiteRegistry() view returns(address)
func (_ERC4626Vault *ERC4626VaultCaller) DolomiteRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "dolomiteRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DolomiteRegistry is a free data retrieval call binding the contract method 0xbd12584f.
//
// Solidity: function dolomiteRegistry() view returns(address)
func (_ERC4626Vault *ERC4626VaultSession) DolomiteRegistry() (common.Address, error) {
	return _ERC4626Vault.Contract.DolomiteRegistry(&_ERC4626Vault.CallOpts)
}

// DolomiteRegistry is a free data retrieval call binding the contract method 0xbd12584f.
//
// Solidity: function dolomiteRegistry() view returns(address)
func (_ERC4626Vault *ERC4626VaultCallerSession) DolomiteRegistry() (common.Address, error) {
	return _ERC4626Vault.Contract.DolomiteRegistry(&_ERC4626Vault.CallOpts)
}

// IsValidReceiver is a free data retrieval call binding the contract method 0xeb9b449c.
//
// Solidity: function isValidReceiver(address _receiver) view returns(bool)
func (_ERC4626Vault *ERC4626VaultCaller) IsValidReceiver(opts *bind.CallOpts, _receiver common.Address) (bool, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "isValidReceiver", _receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceiver is a free data retrieval call binding the contract method 0xeb9b449c.
//
// Solidity: function isValidReceiver(address _receiver) view returns(bool)
func (_ERC4626Vault *ERC4626VaultSession) IsValidReceiver(_receiver common.Address) (bool, error) {
	return _ERC4626Vault.Contract.IsValidReceiver(&_ERC4626Vault.CallOpts, _receiver)
}

// IsValidReceiver is a free data retrieval call binding the contract method 0xeb9b449c.
//
// Solidity: function isValidReceiver(address _receiver) view returns(bool)
func (_ERC4626Vault *ERC4626VaultCallerSession) IsValidReceiver(_receiver common.Address) (bool, error) {
	return _ERC4626Vault.Contract.IsValidReceiver(&_ERC4626Vault.CallOpts, _receiver)
}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) MarketId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "marketId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) MarketId() (*big.Int, error) {
	return _ERC4626Vault.Contract.MarketId(&_ERC4626Vault.CallOpts)
}

// MarketId is a free data retrieval call binding the contract method 0x6ed71ede.
//
// Solidity: function marketId() view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) MarketId() (*big.Int, error) {
	return _ERC4626Vault.Contract.MarketId(&_ERC4626Vault.CallOpts)
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
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) MaxRedeem(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "maxRedeem", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxRedeem(&_ERC4626Vault.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxRedeem(&_ERC4626Vault.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxWithdraw(&_ERC4626Vault.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _ERC4626Vault.Contract.MaxWithdraw(&_ERC4626Vault.CallOpts, _owner)
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

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) PreviewDeposit(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "previewDeposit", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewDeposit(&_ERC4626Vault.CallOpts, _assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewDeposit(&_ERC4626Vault.CallOpts, _assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) PreviewMint(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "previewMint", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewMint(&_ERC4626Vault.CallOpts, _shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewMint(&_ERC4626Vault.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) PreviewRedeem(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "previewRedeem", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) PreviewRedeem(_shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewRedeem(&_ERC4626Vault.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _shares) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) PreviewRedeem(_shares *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewRedeem(&_ERC4626Vault.CallOpts, _shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCaller) PreviewWithdraw(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC4626Vault.contract.Call(opts, &out, "previewWithdraw", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewWithdraw(&_ERC4626Vault.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_ERC4626Vault *ERC4626VaultCallerSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _ERC4626Vault.Contract.PreviewWithdraw(&_ERC4626Vault.CallOpts, _assets)
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

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ERC4626Vault *ERC4626VaultSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Approve(&_ERC4626Vault.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Approve(&_ERC4626Vault.TransactOpts, _spender, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactor) Deposit(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "deposit", _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Deposit(&_ERC4626Vault.TransactOpts, _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Deposit(&_ERC4626Vault.TransactOpts, _assets, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xf3571819.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, uint256 _marketId, address _dolomiteRegistry) returns()
func (_ERC4626Vault *ERC4626VaultTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, _marketId *big.Int, _dolomiteRegistry common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "initialize", _name, _symbol, _decimals, _marketId, _dolomiteRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xf3571819.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, uint256 _marketId, address _dolomiteRegistry) returns()
func (_ERC4626Vault *ERC4626VaultSession) Initialize(_name string, _symbol string, _decimals uint8, _marketId *big.Int, _dolomiteRegistry common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Initialize(&_ERC4626Vault.TransactOpts, _name, _symbol, _decimals, _marketId, _dolomiteRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xf3571819.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, uint256 _marketId, address _dolomiteRegistry) returns()
func (_ERC4626Vault *ERC4626VaultTransactorSession) Initialize(_name string, _symbol string, _decimals uint8, _marketId *big.Int, _dolomiteRegistry common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Initialize(&_ERC4626Vault.TransactOpts, _name, _symbol, _decimals, _marketId, _dolomiteRegistry)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactor) Mint(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "mint", _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Mint(&_ERC4626Vault.TransactOpts, _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Mint(&_ERC4626Vault.TransactOpts, _shares, _receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactor) Redeem(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "redeem", _shares, _receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Redeem(_shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Redeem(&_ERC4626Vault.TransactOpts, _shares, _receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _shares, address _receiver, address _owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Redeem(_shares *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Redeem(&_ERC4626Vault.TransactOpts, _shares, _receiver, _owner)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ERC4626Vault *ERC4626VaultSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Transfer(&_ERC4626Vault.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Transfer(&_ERC4626Vault.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 _amount) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "transferFrom", from, to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 _amount) returns(bool)
func (_ERC4626Vault *ERC4626VaultSession) TransferFrom(from common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.TransferFrom(&_ERC4626Vault.TransactOpts, from, to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 _amount) returns(bool)
func (_ERC4626Vault *ERC4626VaultTransactorSession) TransferFrom(from common.Address, to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.TransferFrom(&_ERC4626Vault.TransactOpts, from, to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactor) Withdraw(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.contract.Transact(opts, "withdraw", _assets, _receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultSession) Withdraw(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Withdraw(&_ERC4626Vault.TransactOpts, _assets, _receiver, _owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_ERC4626Vault *ERC4626VaultTransactorSession) Withdraw(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _ERC4626Vault.Contract.Withdraw(&_ERC4626Vault.TransactOpts, _assets, _receiver, _owner)
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
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC4626Vault *ERC4626VaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC4626VaultInitializedIterator, error) {

	logs, sub, err := _ERC4626Vault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC4626VaultInitializedIterator{contract: _ERC4626Vault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC4626Vault *ERC4626VaultFilterer) ParseInitialized(log types.Log) (*ERC4626VaultInitialized, error) {
	event := new(ERC4626VaultInitialized)
	if err := _ERC4626Vault.contract.UnpackLog(event, "Initialized", log); err != nil {
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
