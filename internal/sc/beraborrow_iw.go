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

// BeraBorrowIWMetaData contains all meta data concerning the BeraBorrowIW contract.
var BeraBorrowIWMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_underlying\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_shareName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_shareSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_metaBeraborrowCore\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_infraredCollVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"infraredCollVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metaBeraborrowCore\",\"outputs\":[{\"internalType\":\"contractIMetaBeraborrowCore\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BeraBorrowIWABI is the input ABI used to generate the binding from.
// Deprecated: Use BeraBorrowIWMetaData.ABI instead.
var BeraBorrowIWABI = BeraBorrowIWMetaData.ABI

// BeraBorrowIW is an auto generated Go binding around an Ethereum contract.
type BeraBorrowIW struct {
	BeraBorrowIWCaller     // Read-only binding to the contract
	BeraBorrowIWTransactor // Write-only binding to the contract
	BeraBorrowIWFilterer   // Log filterer for contract events
}

// BeraBorrowIWCaller is an auto generated read-only Go binding around an Ethereum contract.
type BeraBorrowIWCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowIWTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BeraBorrowIWTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowIWFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BeraBorrowIWFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BeraBorrowIWSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BeraBorrowIWSession struct {
	Contract     *BeraBorrowIW     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BeraBorrowIWCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BeraBorrowIWCallerSession struct {
	Contract *BeraBorrowIWCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BeraBorrowIWTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BeraBorrowIWTransactorSession struct {
	Contract     *BeraBorrowIWTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BeraBorrowIWRaw is an auto generated low-level Go binding around an Ethereum contract.
type BeraBorrowIWRaw struct {
	Contract *BeraBorrowIW // Generic contract binding to access the raw methods on
}

// BeraBorrowIWCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BeraBorrowIWCallerRaw struct {
	Contract *BeraBorrowIWCaller // Generic read-only contract binding to access the raw methods on
}

// BeraBorrowIWTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BeraBorrowIWTransactorRaw struct {
	Contract *BeraBorrowIWTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBeraBorrowIW creates a new instance of BeraBorrowIW, bound to a specific deployed contract.
func NewBeraBorrowIW(address common.Address, backend bind.ContractBackend) (*BeraBorrowIW, error) {
	contract, err := bindBeraBorrowIW(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowIW{BeraBorrowIWCaller: BeraBorrowIWCaller{contract: contract}, BeraBorrowIWTransactor: BeraBorrowIWTransactor{contract: contract}, BeraBorrowIWFilterer: BeraBorrowIWFilterer{contract: contract}}, nil
}

// NewBeraBorrowIWCaller creates a new read-only instance of BeraBorrowIW, bound to a specific deployed contract.
func NewBeraBorrowIWCaller(address common.Address, caller bind.ContractCaller) (*BeraBorrowIWCaller, error) {
	contract, err := bindBeraBorrowIW(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowIWCaller{contract: contract}, nil
}

// NewBeraBorrowIWTransactor creates a new write-only instance of BeraBorrowIW, bound to a specific deployed contract.
func NewBeraBorrowIWTransactor(address common.Address, transactor bind.ContractTransactor) (*BeraBorrowIWTransactor, error) {
	contract, err := bindBeraBorrowIW(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowIWTransactor{contract: contract}, nil
}

// NewBeraBorrowIWFilterer creates a new log filterer instance of BeraBorrowIW, bound to a specific deployed contract.
func NewBeraBorrowIWFilterer(address common.Address, filterer bind.ContractFilterer) (*BeraBorrowIWFilterer, error) {
	contract, err := bindBeraBorrowIW(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowIWFilterer{contract: contract}, nil
}

// bindBeraBorrowIW binds a generic wrapper to an already deployed contract.
func bindBeraBorrowIW(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BeraBorrowIWMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeraBorrowIW *BeraBorrowIWRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeraBorrowIW.Contract.BeraBorrowIWCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeraBorrowIW *BeraBorrowIWRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.BeraBorrowIWTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeraBorrowIW *BeraBorrowIWRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.BeraBorrowIWTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BeraBorrowIW *BeraBorrowIWCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BeraBorrowIW.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BeraBorrowIW *BeraBorrowIWTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BeraBorrowIW *BeraBorrowIWTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowIW.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BeraBorrowIW.Contract.Allowance(&_BeraBorrowIW.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BeraBorrowIW.Contract.Allowance(&_BeraBorrowIW.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowIW.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BeraBorrowIW.Contract.BalanceOf(&_BeraBorrowIW.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BeraBorrowIW.Contract.BalanceOf(&_BeraBorrowIW.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowIW *BeraBorrowIWCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BeraBorrowIW.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowIW *BeraBorrowIWSession) Decimals() (uint8, error) {
	return _BeraBorrowIW.Contract.Decimals(&_BeraBorrowIW.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BeraBorrowIW *BeraBorrowIWCallerSession) Decimals() (uint8, error) {
	return _BeraBorrowIW.Contract.Decimals(&_BeraBorrowIW.CallOpts)
}

// InfraredCollVault is a free data retrieval call binding the contract method 0xac22c86d.
//
// Solidity: function infraredCollVault() view returns(address)
func (_BeraBorrowIW *BeraBorrowIWCaller) InfraredCollVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowIW.contract.Call(opts, &out, "infraredCollVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InfraredCollVault is a free data retrieval call binding the contract method 0xac22c86d.
//
// Solidity: function infraredCollVault() view returns(address)
func (_BeraBorrowIW *BeraBorrowIWSession) InfraredCollVault() (common.Address, error) {
	return _BeraBorrowIW.Contract.InfraredCollVault(&_BeraBorrowIW.CallOpts)
}

// InfraredCollVault is a free data retrieval call binding the contract method 0xac22c86d.
//
// Solidity: function infraredCollVault() view returns(address)
func (_BeraBorrowIW *BeraBorrowIWCallerSession) InfraredCollVault() (common.Address, error) {
	return _BeraBorrowIW.Contract.InfraredCollVault(&_BeraBorrowIW.CallOpts)
}

// MetaBeraborrowCore is a free data retrieval call binding the contract method 0xaf068f4f.
//
// Solidity: function metaBeraborrowCore() view returns(address)
func (_BeraBorrowIW *BeraBorrowIWCaller) MetaBeraborrowCore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowIW.contract.Call(opts, &out, "metaBeraborrowCore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetaBeraborrowCore is a free data retrieval call binding the contract method 0xaf068f4f.
//
// Solidity: function metaBeraborrowCore() view returns(address)
func (_BeraBorrowIW *BeraBorrowIWSession) MetaBeraborrowCore() (common.Address, error) {
	return _BeraBorrowIW.Contract.MetaBeraborrowCore(&_BeraBorrowIW.CallOpts)
}

// MetaBeraborrowCore is a free data retrieval call binding the contract method 0xaf068f4f.
//
// Solidity: function metaBeraborrowCore() view returns(address)
func (_BeraBorrowIW *BeraBorrowIWCallerSession) MetaBeraborrowCore() (common.Address, error) {
	return _BeraBorrowIW.Contract.MetaBeraborrowCore(&_BeraBorrowIW.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowIW *BeraBorrowIWCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowIW.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowIW *BeraBorrowIWSession) Name() (string, error) {
	return _BeraBorrowIW.Contract.Name(&_BeraBorrowIW.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BeraBorrowIW *BeraBorrowIWCallerSession) Name() (string, error) {
	return _BeraBorrowIW.Contract.Name(&_BeraBorrowIW.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowIW *BeraBorrowIWCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BeraBorrowIW.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowIW *BeraBorrowIWSession) Symbol() (string, error) {
	return _BeraBorrowIW.Contract.Symbol(&_BeraBorrowIW.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BeraBorrowIW *BeraBorrowIWCallerSession) Symbol() (string, error) {
	return _BeraBorrowIW.Contract.Symbol(&_BeraBorrowIW.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BeraBorrowIW.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWSession) TotalSupply() (*big.Int, error) {
	return _BeraBorrowIW.Contract.TotalSupply(&_BeraBorrowIW.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWCallerSession) TotalSupply() (*big.Int, error) {
	return _BeraBorrowIW.Contract.TotalSupply(&_BeraBorrowIW.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_BeraBorrowIW *BeraBorrowIWCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BeraBorrowIW.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_BeraBorrowIW *BeraBorrowIWSession) Underlying() (common.Address, error) {
	return _BeraBorrowIW.Contract.Underlying(&_BeraBorrowIW.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_BeraBorrowIW *BeraBorrowIWCallerSession) Underlying() (common.Address, error) {
	return _BeraBorrowIW.Contract.Underlying(&_BeraBorrowIW.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.Approve(&_BeraBorrowIW.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.Approve(&_BeraBorrowIW.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.DecreaseAllowance(&_BeraBorrowIW.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.DecreaseAllowance(&_BeraBorrowIW.TransactOpts, spender, subtractedValue)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address account, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactor) DepositFor(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.contract.Transact(opts, "depositFor", account, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address account, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWSession) DepositFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.DepositFor(&_BeraBorrowIW.TransactOpts, account, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address account, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactorSession) DepositFor(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.DepositFor(&_BeraBorrowIW.TransactOpts, account, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.IncreaseAllowance(&_BeraBorrowIW.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.IncreaseAllowance(&_BeraBorrowIW.TransactOpts, spender, addedValue)
}

// Recover is a paid mutator transaction binding the contract method 0x0cd865ec.
//
// Solidity: function recover(address account) returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWTransactor) Recover(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BeraBorrowIW.contract.Transact(opts, "recover", account)
}

// Recover is a paid mutator transaction binding the contract method 0x0cd865ec.
//
// Solidity: function recover(address account) returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWSession) Recover(account common.Address) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.Recover(&_BeraBorrowIW.TransactOpts, account)
}

// Recover is a paid mutator transaction binding the contract method 0x0cd865ec.
//
// Solidity: function recover(address account) returns(uint256)
func (_BeraBorrowIW *BeraBorrowIWTransactorSession) Recover(account common.Address) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.Recover(&_BeraBorrowIW.TransactOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.Transfer(&_BeraBorrowIW.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.Transfer(&_BeraBorrowIW.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.TransferFrom(&_BeraBorrowIW.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.TransferFrom(&_BeraBorrowIW.TransactOpts, from, to, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address account, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactor) WithdrawTo(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.contract.Transact(opts, "withdrawTo", account, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address account, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWSession) WithdrawTo(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.WithdrawTo(&_BeraBorrowIW.TransactOpts, account, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address account, uint256 amount) returns(bool)
func (_BeraBorrowIW *BeraBorrowIWTransactorSession) WithdrawTo(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BeraBorrowIW.Contract.WithdrawTo(&_BeraBorrowIW.TransactOpts, account, amount)
}

// BeraBorrowIWApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BeraBorrowIW contract.
type BeraBorrowIWApprovalIterator struct {
	Event *BeraBorrowIWApproval // Event containing the contract specifics and raw log

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
func (it *BeraBorrowIWApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowIWApproval)
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
		it.Event = new(BeraBorrowIWApproval)
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
func (it *BeraBorrowIWApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowIWApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowIWApproval represents a Approval event raised by the BeraBorrowIW contract.
type BeraBorrowIWApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeraBorrowIW *BeraBorrowIWFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BeraBorrowIWApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BeraBorrowIW.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowIWApprovalIterator{contract: _BeraBorrowIW.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BeraBorrowIW *BeraBorrowIWFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BeraBorrowIWApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BeraBorrowIW.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowIWApproval)
				if err := _BeraBorrowIW.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BeraBorrowIW *BeraBorrowIWFilterer) ParseApproval(log types.Log) (*BeraBorrowIWApproval, error) {
	event := new(BeraBorrowIWApproval)
	if err := _BeraBorrowIW.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BeraBorrowIWTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BeraBorrowIW contract.
type BeraBorrowIWTransferIterator struct {
	Event *BeraBorrowIWTransfer // Event containing the contract specifics and raw log

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
func (it *BeraBorrowIWTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BeraBorrowIWTransfer)
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
		it.Event = new(BeraBorrowIWTransfer)
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
func (it *BeraBorrowIWTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BeraBorrowIWTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BeraBorrowIWTransfer represents a Transfer event raised by the BeraBorrowIW contract.
type BeraBorrowIWTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BeraBorrowIW *BeraBorrowIWFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BeraBorrowIWTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeraBorrowIW.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BeraBorrowIWTransferIterator{contract: _BeraBorrowIW.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BeraBorrowIW *BeraBorrowIWFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BeraBorrowIWTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BeraBorrowIW.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BeraBorrowIWTransfer)
				if err := _BeraBorrowIW.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BeraBorrowIW *BeraBorrowIWFilterer) ParseTransfer(log types.Log) (*BeraBorrowIWTransfer, error) {
	event := new(BeraBorrowIWTransfer)
	if err := _BeraBorrowIW.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
