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

// EtherfiVaultMetaData contains all meta data concerning the EtherfiVault contract.
var EtherfiVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractAuthority\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"AuthorityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Enter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"internalType\":\"contractAuthority\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"}],\"name\":\"enter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractERC20\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shareAmount\",\"type\":\"uint256\"}],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hook\",\"outputs\":[{\"internalType\":\"contractBeforeTransferHook\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"manage\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"manage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractAuthority\",\"name\":\"newAuthority\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hook\",\"type\":\"address\"}],\"name\":\"setBeforeTransferHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// EtherfiVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherfiVaultMetaData.ABI instead.
var EtherfiVaultABI = EtherfiVaultMetaData.ABI

// EtherfiVault is an auto generated Go binding around an Ethereum contract.
type EtherfiVault struct {
	EtherfiVaultCaller     // Read-only binding to the contract
	EtherfiVaultTransactor // Write-only binding to the contract
	EtherfiVaultFilterer   // Log filterer for contract events
}

// EtherfiVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherfiVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherfiVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherfiVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherfiVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherfiVaultSession struct {
	Contract     *EtherfiVault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherfiVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherfiVaultCallerSession struct {
	Contract *EtherfiVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EtherfiVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherfiVaultTransactorSession struct {
	Contract     *EtherfiVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EtherfiVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherfiVaultRaw struct {
	Contract *EtherfiVault // Generic contract binding to access the raw methods on
}

// EtherfiVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherfiVaultCallerRaw struct {
	Contract *EtherfiVaultCaller // Generic read-only contract binding to access the raw methods on
}

// EtherfiVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherfiVaultTransactorRaw struct {
	Contract *EtherfiVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherfiVault creates a new instance of EtherfiVault, bound to a specific deployed contract.
func NewEtherfiVault(address common.Address, backend bind.ContractBackend) (*EtherfiVault, error) {
	contract, err := bindEtherfiVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherfiVault{EtherfiVaultCaller: EtherfiVaultCaller{contract: contract}, EtherfiVaultTransactor: EtherfiVaultTransactor{contract: contract}, EtherfiVaultFilterer: EtherfiVaultFilterer{contract: contract}}, nil
}

// NewEtherfiVaultCaller creates a new read-only instance of EtherfiVault, bound to a specific deployed contract.
func NewEtherfiVaultCaller(address common.Address, caller bind.ContractCaller) (*EtherfiVaultCaller, error) {
	contract, err := bindEtherfiVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiVaultCaller{contract: contract}, nil
}

// NewEtherfiVaultTransactor creates a new write-only instance of EtherfiVault, bound to a specific deployed contract.
func NewEtherfiVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherfiVaultTransactor, error) {
	contract, err := bindEtherfiVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherfiVaultTransactor{contract: contract}, nil
}

// NewEtherfiVaultFilterer creates a new log filterer instance of EtherfiVault, bound to a specific deployed contract.
func NewEtherfiVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherfiVaultFilterer, error) {
	contract, err := bindEtherfiVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherfiVaultFilterer{contract: contract}, nil
}

// bindEtherfiVault binds a generic wrapper to an already deployed contract.
func bindEtherfiVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherfiVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiVault *EtherfiVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiVault.Contract.EtherfiVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiVault *EtherfiVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiVault.Contract.EtherfiVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiVault *EtherfiVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiVault.Contract.EtherfiVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherfiVault *EtherfiVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherfiVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherfiVault *EtherfiVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherfiVault *EtherfiVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherfiVault.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EtherfiVault *EtherfiVaultCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EtherfiVault *EtherfiVaultSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _EtherfiVault.Contract.DOMAINSEPARATOR(&_EtherfiVault.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_EtherfiVault *EtherfiVaultCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _EtherfiVault.Contract.DOMAINSEPARATOR(&_EtherfiVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_EtherfiVault *EtherfiVaultCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_EtherfiVault *EtherfiVaultSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EtherfiVault.Contract.Allowance(&_EtherfiVault.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_EtherfiVault *EtherfiVaultCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EtherfiVault.Contract.Allowance(&_EtherfiVault.CallOpts, arg0, arg1)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_EtherfiVault *EtherfiVaultCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "authority")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_EtherfiVault *EtherfiVaultSession) Authority() (common.Address, error) {
	return _EtherfiVault.Contract.Authority(&_EtherfiVault.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() view returns(address)
func (_EtherfiVault *EtherfiVaultCallerSession) Authority() (common.Address, error) {
	return _EtherfiVault.Contract.Authority(&_EtherfiVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_EtherfiVault *EtherfiVaultCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_EtherfiVault *EtherfiVaultSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _EtherfiVault.Contract.BalanceOf(&_EtherfiVault.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_EtherfiVault *EtherfiVaultCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _EtherfiVault.Contract.BalanceOf(&_EtherfiVault.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EtherfiVault *EtherfiVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EtherfiVault *EtherfiVaultSession) Decimals() (uint8, error) {
	return _EtherfiVault.Contract.Decimals(&_EtherfiVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EtherfiVault *EtherfiVaultCallerSession) Decimals() (uint8, error) {
	return _EtherfiVault.Contract.Decimals(&_EtherfiVault.CallOpts)
}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_EtherfiVault *EtherfiVaultCaller) Hook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "hook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_EtherfiVault *EtherfiVaultSession) Hook() (common.Address, error) {
	return _EtherfiVault.Contract.Hook(&_EtherfiVault.CallOpts)
}

// Hook is a free data retrieval call binding the contract method 0x7f5a7c7b.
//
// Solidity: function hook() view returns(address)
func (_EtherfiVault *EtherfiVaultCallerSession) Hook() (common.Address, error) {
	return _EtherfiVault.Contract.Hook(&_EtherfiVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EtherfiVault *EtherfiVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EtherfiVault *EtherfiVaultSession) Name() (string, error) {
	return _EtherfiVault.Contract.Name(&_EtherfiVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EtherfiVault *EtherfiVaultCallerSession) Name() (string, error) {
	return _EtherfiVault.Contract.Name(&_EtherfiVault.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_EtherfiVault *EtherfiVaultCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_EtherfiVault *EtherfiVaultSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _EtherfiVault.Contract.Nonces(&_EtherfiVault.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_EtherfiVault *EtherfiVaultCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _EtherfiVault.Contract.Nonces(&_EtherfiVault.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiVault *EtherfiVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiVault *EtherfiVaultSession) Owner() (common.Address, error) {
	return _EtherfiVault.Contract.Owner(&_EtherfiVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherfiVault *EtherfiVaultCallerSession) Owner() (common.Address, error) {
	return _EtherfiVault.Contract.Owner(&_EtherfiVault.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EtherfiVault *EtherfiVaultCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EtherfiVault *EtherfiVaultSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EtherfiVault.Contract.SupportsInterface(&_EtherfiVault.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_EtherfiVault *EtherfiVaultCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _EtherfiVault.Contract.SupportsInterface(&_EtherfiVault.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EtherfiVault *EtherfiVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EtherfiVault *EtherfiVaultSession) Symbol() (string, error) {
	return _EtherfiVault.Contract.Symbol(&_EtherfiVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EtherfiVault *EtherfiVaultCallerSession) Symbol() (string, error) {
	return _EtherfiVault.Contract.Symbol(&_EtherfiVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EtherfiVault *EtherfiVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EtherfiVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EtherfiVault *EtherfiVaultSession) TotalSupply() (*big.Int, error) {
	return _EtherfiVault.Contract.TotalSupply(&_EtherfiVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EtherfiVault *EtherfiVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _EtherfiVault.Contract.TotalSupply(&_EtherfiVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EtherfiVault *EtherfiVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EtherfiVault *EtherfiVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Approve(&_EtherfiVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EtherfiVault *EtherfiVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Approve(&_EtherfiVault.TransactOpts, spender, amount)
}

// Enter is a paid mutator transaction binding the contract method 0x39d6ba32.
//
// Solidity: function enter(address from, address asset, uint256 assetAmount, address to, uint256 shareAmount) returns()
func (_EtherfiVault *EtherfiVaultTransactor) Enter(opts *bind.TransactOpts, from common.Address, asset common.Address, assetAmount *big.Int, to common.Address, shareAmount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "enter", from, asset, assetAmount, to, shareAmount)
}

// Enter is a paid mutator transaction binding the contract method 0x39d6ba32.
//
// Solidity: function enter(address from, address asset, uint256 assetAmount, address to, uint256 shareAmount) returns()
func (_EtherfiVault *EtherfiVaultSession) Enter(from common.Address, asset common.Address, assetAmount *big.Int, to common.Address, shareAmount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Enter(&_EtherfiVault.TransactOpts, from, asset, assetAmount, to, shareAmount)
}

// Enter is a paid mutator transaction binding the contract method 0x39d6ba32.
//
// Solidity: function enter(address from, address asset, uint256 assetAmount, address to, uint256 shareAmount) returns()
func (_EtherfiVault *EtherfiVaultTransactorSession) Enter(from common.Address, asset common.Address, assetAmount *big.Int, to common.Address, shareAmount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Enter(&_EtherfiVault.TransactOpts, from, asset, assetAmount, to, shareAmount)
}

// Exit is a paid mutator transaction binding the contract method 0x18457e61.
//
// Solidity: function exit(address to, address asset, uint256 assetAmount, address from, uint256 shareAmount) returns()
func (_EtherfiVault *EtherfiVaultTransactor) Exit(opts *bind.TransactOpts, to common.Address, asset common.Address, assetAmount *big.Int, from common.Address, shareAmount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "exit", to, asset, assetAmount, from, shareAmount)
}

// Exit is a paid mutator transaction binding the contract method 0x18457e61.
//
// Solidity: function exit(address to, address asset, uint256 assetAmount, address from, uint256 shareAmount) returns()
func (_EtherfiVault *EtherfiVaultSession) Exit(to common.Address, asset common.Address, assetAmount *big.Int, from common.Address, shareAmount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Exit(&_EtherfiVault.TransactOpts, to, asset, assetAmount, from, shareAmount)
}

// Exit is a paid mutator transaction binding the contract method 0x18457e61.
//
// Solidity: function exit(address to, address asset, uint256 assetAmount, address from, uint256 shareAmount) returns()
func (_EtherfiVault *EtherfiVaultTransactorSession) Exit(to common.Address, asset common.Address, assetAmount *big.Int, from common.Address, shareAmount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Exit(&_EtherfiVault.TransactOpts, to, asset, assetAmount, from, shareAmount)
}

// Manage is a paid mutator transaction binding the contract method 0x224d8703.
//
// Solidity: function manage(address[] targets, bytes[] data, uint256[] values) returns(bytes[] results)
func (_EtherfiVault *EtherfiVaultTransactor) Manage(opts *bind.TransactOpts, targets []common.Address, data [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "manage", targets, data, values)
}

// Manage is a paid mutator transaction binding the contract method 0x224d8703.
//
// Solidity: function manage(address[] targets, bytes[] data, uint256[] values) returns(bytes[] results)
func (_EtherfiVault *EtherfiVaultSession) Manage(targets []common.Address, data [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Manage(&_EtherfiVault.TransactOpts, targets, data, values)
}

// Manage is a paid mutator transaction binding the contract method 0x224d8703.
//
// Solidity: function manage(address[] targets, bytes[] data, uint256[] values) returns(bytes[] results)
func (_EtherfiVault *EtherfiVaultTransactorSession) Manage(targets []common.Address, data [][]byte, values []*big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Manage(&_EtherfiVault.TransactOpts, targets, data, values)
}

// Manage0 is a paid mutator transaction binding the contract method 0xf6e715d0.
//
// Solidity: function manage(address target, bytes data, uint256 value) returns(bytes result)
func (_EtherfiVault *EtherfiVaultTransactor) Manage0(opts *bind.TransactOpts, target common.Address, data []byte, value *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "manage0", target, data, value)
}

// Manage0 is a paid mutator transaction binding the contract method 0xf6e715d0.
//
// Solidity: function manage(address target, bytes data, uint256 value) returns(bytes result)
func (_EtherfiVault *EtherfiVaultSession) Manage0(target common.Address, data []byte, value *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Manage0(&_EtherfiVault.TransactOpts, target, data, value)
}

// Manage0 is a paid mutator transaction binding the contract method 0xf6e715d0.
//
// Solidity: function manage(address target, bytes data, uint256 value) returns(bytes result)
func (_EtherfiVault *EtherfiVaultTransactorSession) Manage0(target common.Address, data []byte, value *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Manage0(&_EtherfiVault.TransactOpts, target, data, value)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_EtherfiVault *EtherfiVaultTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_EtherfiVault *EtherfiVaultSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _EtherfiVault.Contract.OnERC1155BatchReceived(&_EtherfiVault.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_EtherfiVault *EtherfiVaultTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _EtherfiVault.Contract.OnERC1155BatchReceived(&_EtherfiVault.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_EtherfiVault *EtherfiVaultTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_EtherfiVault *EtherfiVaultSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _EtherfiVault.Contract.OnERC1155Received(&_EtherfiVault.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_EtherfiVault *EtherfiVaultTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _EtherfiVault.Contract.OnERC1155Received(&_EtherfiVault.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_EtherfiVault *EtherfiVaultTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_EtherfiVault *EtherfiVaultSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _EtherfiVault.Contract.OnERC721Received(&_EtherfiVault.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_EtherfiVault *EtherfiVaultTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _EtherfiVault.Contract.OnERC721Received(&_EtherfiVault.TransactOpts, arg0, arg1, arg2, arg3)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_EtherfiVault *EtherfiVaultTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_EtherfiVault *EtherfiVaultSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Permit(&_EtherfiVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_EtherfiVault *EtherfiVaultTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Permit(&_EtherfiVault.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_EtherfiVault *EtherfiVaultTransactor) SetAuthority(opts *bind.TransactOpts, newAuthority common.Address) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "setAuthority", newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_EtherfiVault *EtherfiVaultSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _EtherfiVault.Contract.SetAuthority(&_EtherfiVault.TransactOpts, newAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address newAuthority) returns()
func (_EtherfiVault *EtherfiVaultTransactorSession) SetAuthority(newAuthority common.Address) (*types.Transaction, error) {
	return _EtherfiVault.Contract.SetAuthority(&_EtherfiVault.TransactOpts, newAuthority)
}

// SetBeforeTransferHook is a paid mutator transaction binding the contract method 0x8929565f.
//
// Solidity: function setBeforeTransferHook(address _hook) returns()
func (_EtherfiVault *EtherfiVaultTransactor) SetBeforeTransferHook(opts *bind.TransactOpts, _hook common.Address) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "setBeforeTransferHook", _hook)
}

// SetBeforeTransferHook is a paid mutator transaction binding the contract method 0x8929565f.
//
// Solidity: function setBeforeTransferHook(address _hook) returns()
func (_EtherfiVault *EtherfiVaultSession) SetBeforeTransferHook(_hook common.Address) (*types.Transaction, error) {
	return _EtherfiVault.Contract.SetBeforeTransferHook(&_EtherfiVault.TransactOpts, _hook)
}

// SetBeforeTransferHook is a paid mutator transaction binding the contract method 0x8929565f.
//
// Solidity: function setBeforeTransferHook(address _hook) returns()
func (_EtherfiVault *EtherfiVaultTransactorSession) SetBeforeTransferHook(_hook common.Address) (*types.Transaction, error) {
	return _EtherfiVault.Contract.SetBeforeTransferHook(&_EtherfiVault.TransactOpts, _hook)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EtherfiVault *EtherfiVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EtherfiVault *EtherfiVaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Transfer(&_EtherfiVault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EtherfiVault *EtherfiVaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.Transfer(&_EtherfiVault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EtherfiVault *EtherfiVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EtherfiVault *EtherfiVaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.TransferFrom(&_EtherfiVault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EtherfiVault *EtherfiVaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherfiVault.Contract.TransferFrom(&_EtherfiVault.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiVault *EtherfiVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiVault *EtherfiVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiVault.Contract.TransferOwnership(&_EtherfiVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EtherfiVault *EtherfiVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherfiVault.Contract.TransferOwnership(&_EtherfiVault.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherfiVault *EtherfiVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherfiVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherfiVault *EtherfiVaultSession) Receive() (*types.Transaction, error) {
	return _EtherfiVault.Contract.Receive(&_EtherfiVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherfiVault *EtherfiVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _EtherfiVault.Contract.Receive(&_EtherfiVault.TransactOpts)
}

// EtherfiVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EtherfiVault contract.
type EtherfiVaultApprovalIterator struct {
	Event *EtherfiVaultApproval // Event containing the contract specifics and raw log

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
func (it *EtherfiVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiVaultApproval)
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
		it.Event = new(EtherfiVaultApproval)
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
func (it *EtherfiVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiVaultApproval represents a Approval event raised by the EtherfiVault contract.
type EtherfiVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EtherfiVault *EtherfiVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EtherfiVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EtherfiVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiVaultApprovalIterator{contract: _EtherfiVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EtherfiVault *EtherfiVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EtherfiVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EtherfiVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiVaultApproval)
				if err := _EtherfiVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EtherfiVault *EtherfiVaultFilterer) ParseApproval(log types.Log) (*EtherfiVaultApproval, error) {
	event := new(EtherfiVaultApproval)
	if err := _EtherfiVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiVaultAuthorityUpdatedIterator is returned from FilterAuthorityUpdated and is used to iterate over the raw logs and unpacked data for AuthorityUpdated events raised by the EtherfiVault contract.
type EtherfiVaultAuthorityUpdatedIterator struct {
	Event *EtherfiVaultAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *EtherfiVaultAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiVaultAuthorityUpdated)
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
		it.Event = new(EtherfiVaultAuthorityUpdated)
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
func (it *EtherfiVaultAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiVaultAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiVaultAuthorityUpdated represents a AuthorityUpdated event raised by the EtherfiVault contract.
type EtherfiVaultAuthorityUpdated struct {
	User         common.Address
	NewAuthority common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuthorityUpdated is a free log retrieval operation binding the contract event 0xa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b76389980198.
//
// Solidity: event AuthorityUpdated(address indexed user, address indexed newAuthority)
func (_EtherfiVault *EtherfiVaultFilterer) FilterAuthorityUpdated(opts *bind.FilterOpts, user []common.Address, newAuthority []common.Address) (*EtherfiVaultAuthorityUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAuthorityRule []interface{}
	for _, newAuthorityItem := range newAuthority {
		newAuthorityRule = append(newAuthorityRule, newAuthorityItem)
	}

	logs, sub, err := _EtherfiVault.contract.FilterLogs(opts, "AuthorityUpdated", userRule, newAuthorityRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiVaultAuthorityUpdatedIterator{contract: _EtherfiVault.contract, event: "AuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorityUpdated is a free log subscription operation binding the contract event 0xa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b76389980198.
//
// Solidity: event AuthorityUpdated(address indexed user, address indexed newAuthority)
func (_EtherfiVault *EtherfiVaultFilterer) WatchAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *EtherfiVaultAuthorityUpdated, user []common.Address, newAuthority []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newAuthorityRule []interface{}
	for _, newAuthorityItem := range newAuthority {
		newAuthorityRule = append(newAuthorityRule, newAuthorityItem)
	}

	logs, sub, err := _EtherfiVault.contract.WatchLogs(opts, "AuthorityUpdated", userRule, newAuthorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiVaultAuthorityUpdated)
				if err := _EtherfiVault.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
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

// ParseAuthorityUpdated is a log parse operation binding the contract event 0xa3396fd7f6e0a21b50e5089d2da70d5ac0a3bbbd1f617a93f134b76389980198.
//
// Solidity: event AuthorityUpdated(address indexed user, address indexed newAuthority)
func (_EtherfiVault *EtherfiVaultFilterer) ParseAuthorityUpdated(log types.Log) (*EtherfiVaultAuthorityUpdated, error) {
	event := new(EtherfiVaultAuthorityUpdated)
	if err := _EtherfiVault.contract.UnpackLog(event, "AuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiVaultEnterIterator is returned from FilterEnter and is used to iterate over the raw logs and unpacked data for Enter events raised by the EtherfiVault contract.
type EtherfiVaultEnterIterator struct {
	Event *EtherfiVaultEnter // Event containing the contract specifics and raw log

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
func (it *EtherfiVaultEnterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiVaultEnter)
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
		it.Event = new(EtherfiVaultEnter)
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
func (it *EtherfiVaultEnterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiVaultEnterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiVaultEnter represents a Enter event raised by the EtherfiVault contract.
type EtherfiVaultEnter struct {
	From   common.Address
	Asset  common.Address
	Amount *big.Int
	To     common.Address
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnter is a free log retrieval operation binding the contract event 0xea00f88768a86184a6e515238a549c171769fe7460a011d6fd0bcd48ca078ea4.
//
// Solidity: event Enter(address indexed from, address indexed asset, uint256 amount, address indexed to, uint256 shares)
func (_EtherfiVault *EtherfiVaultFilterer) FilterEnter(opts *bind.FilterOpts, from []common.Address, asset []common.Address, to []common.Address) (*EtherfiVaultEnterIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherfiVault.contract.FilterLogs(opts, "Enter", fromRule, assetRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiVaultEnterIterator{contract: _EtherfiVault.contract, event: "Enter", logs: logs, sub: sub}, nil
}

// WatchEnter is a free log subscription operation binding the contract event 0xea00f88768a86184a6e515238a549c171769fe7460a011d6fd0bcd48ca078ea4.
//
// Solidity: event Enter(address indexed from, address indexed asset, uint256 amount, address indexed to, uint256 shares)
func (_EtherfiVault *EtherfiVaultFilterer) WatchEnter(opts *bind.WatchOpts, sink chan<- *EtherfiVaultEnter, from []common.Address, asset []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherfiVault.contract.WatchLogs(opts, "Enter", fromRule, assetRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiVaultEnter)
				if err := _EtherfiVault.contract.UnpackLog(event, "Enter", log); err != nil {
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

// ParseEnter is a log parse operation binding the contract event 0xea00f88768a86184a6e515238a549c171769fe7460a011d6fd0bcd48ca078ea4.
//
// Solidity: event Enter(address indexed from, address indexed asset, uint256 amount, address indexed to, uint256 shares)
func (_EtherfiVault *EtherfiVaultFilterer) ParseEnter(log types.Log) (*EtherfiVaultEnter, error) {
	event := new(EtherfiVaultEnter)
	if err := _EtherfiVault.contract.UnpackLog(event, "Enter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiVaultExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the EtherfiVault contract.
type EtherfiVaultExitIterator struct {
	Event *EtherfiVaultExit // Event containing the contract specifics and raw log

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
func (it *EtherfiVaultExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiVaultExit)
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
		it.Event = new(EtherfiVaultExit)
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
func (it *EtherfiVaultExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiVaultExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiVaultExit represents a Exit event raised by the EtherfiVault contract.
type EtherfiVaultExit struct {
	To     common.Address
	Asset  common.Address
	Amount *big.Int
	From   common.Address
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0xe0c82280a1164680e0cf43be7db4c4c9f985423623ad7a544fb76c772bdc6043.
//
// Solidity: event Exit(address indexed to, address indexed asset, uint256 amount, address indexed from, uint256 shares)
func (_EtherfiVault *EtherfiVaultFilterer) FilterExit(opts *bind.FilterOpts, to []common.Address, asset []common.Address, from []common.Address) (*EtherfiVaultExitIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _EtherfiVault.contract.FilterLogs(opts, "Exit", toRule, assetRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiVaultExitIterator{contract: _EtherfiVault.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0xe0c82280a1164680e0cf43be7db4c4c9f985423623ad7a544fb76c772bdc6043.
//
// Solidity: event Exit(address indexed to, address indexed asset, uint256 amount, address indexed from, uint256 shares)
func (_EtherfiVault *EtherfiVaultFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *EtherfiVaultExit, to []common.Address, asset []common.Address, from []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _EtherfiVault.contract.WatchLogs(opts, "Exit", toRule, assetRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiVaultExit)
				if err := _EtherfiVault.contract.UnpackLog(event, "Exit", log); err != nil {
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

// ParseExit is a log parse operation binding the contract event 0xe0c82280a1164680e0cf43be7db4c4c9f985423623ad7a544fb76c772bdc6043.
//
// Solidity: event Exit(address indexed to, address indexed asset, uint256 amount, address indexed from, uint256 shares)
func (_EtherfiVault *EtherfiVaultFilterer) ParseExit(log types.Log) (*EtherfiVaultExit, error) {
	event := new(EtherfiVaultExit)
	if err := _EtherfiVault.contract.UnpackLog(event, "Exit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EtherfiVault contract.
type EtherfiVaultOwnershipTransferredIterator struct {
	Event *EtherfiVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EtherfiVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiVaultOwnershipTransferred)
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
		it.Event = new(EtherfiVaultOwnershipTransferred)
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
func (it *EtherfiVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiVaultOwnershipTransferred represents a OwnershipTransferred event raised by the EtherfiVault contract.
type EtherfiVaultOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_EtherfiVault *EtherfiVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*EtherfiVaultOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherfiVault.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiVaultOwnershipTransferredIterator{contract: _EtherfiVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_EtherfiVault *EtherfiVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EtherfiVaultOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EtherfiVault.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiVaultOwnershipTransferred)
				if err := _EtherfiVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_EtherfiVault *EtherfiVaultFilterer) ParseOwnershipTransferred(log types.Log) (*EtherfiVaultOwnershipTransferred, error) {
	event := new(EtherfiVaultOwnershipTransferred)
	if err := _EtherfiVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EtherfiVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EtherfiVault contract.
type EtherfiVaultTransferIterator struct {
	Event *EtherfiVaultTransfer // Event containing the contract specifics and raw log

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
func (it *EtherfiVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherfiVaultTransfer)
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
		it.Event = new(EtherfiVaultTransfer)
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
func (it *EtherfiVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherfiVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherfiVaultTransfer represents a Transfer event raised by the EtherfiVault contract.
type EtherfiVaultTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_EtherfiVault *EtherfiVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EtherfiVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherfiVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EtherfiVaultTransferIterator{contract: _EtherfiVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_EtherfiVault *EtherfiVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EtherfiVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EtherfiVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherfiVaultTransfer)
				if err := _EtherfiVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_EtherfiVault *EtherfiVaultFilterer) ParseTransfer(log types.Log) (*EtherfiVaultTransfer, error) {
	event := new(EtherfiVaultTransfer)
	if err := _EtherfiVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
