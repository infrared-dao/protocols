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

// BaseGameContractProxyGameConfig is an auto generated low-level Go binding around an user-defined struct.
type BaseGameContractProxyGameConfig struct {
	Player           common.Address
	Token            common.Address
	Wager            *big.Int
	UserRandomNumber [32]byte
	Extra            uint8
	Count            uint8
}

// VaultManagerProxyVaultWithPrice is an auto generated low-level Go binding around an user-defined struct.
type VaultManagerProxyVaultWithPrice struct {
	VaultTokenAddress common.Address
	LiqTokenAddress   common.Address
	TotalBalance      *big.Int
	TotalVaultTokens  *big.Int
	VaultTokenPrice   *big.Int
}

// JunkyVaultManagerMetaData contains all meta data concerning the JunkyVaultManager contract.
var JunkyVaultManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payout\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"randomNumber\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"wonCount\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"totalCount\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"GameResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wager\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"GameStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liqTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liqTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liqTokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vaultTokenAddress\",\"type\":\"address\"}],\"name\":\"VaultCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"randomNumber\",\"type\":\"bytes32\"}],\"name\":\"_entropyCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liqTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"liqTokenAddress\",\"type\":\"address\"}],\"name\":\"createVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entropy\",\"outputs\":[{\"internalType\":\"contractIEntropy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entropyProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllVaults\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vaultTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liqTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVaultTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultTokenPrice\",\"type\":\"uint256\"}],\"internalType\":\"structVaultManagerProxy.VaultWithPrice[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"sequenceNumber\",\"type\":\"uint64\"}],\"name\":\"getGame\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wager\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userRandomNumber\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"extra\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"count\",\"type\":\"uint8\"}],\"internalType\":\"structBaseGameContractProxy.GameConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liqTokenAddress\",\"type\":\"address\"}],\"name\":\"getVault\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"vaultTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liqTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVaultTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultTokenPrice\",\"type\":\"uint256\"}],\"internalType\":\"structVaultManagerProxy.VaultWithPrice\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liqTokenAddress\",\"type\":\"address\"}],\"name\":\"getVaultTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"houseEdge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasuryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"entropyAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"entropyProviderAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liqTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vaultTokens\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newEdge\",\"type\":\"uint256\"}],\"name\":\"setHouseEdge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAmount\",\"type\":\"uint256\"}],\"name\":\"setMinAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasuryContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"contractTreasuryContractProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"vaultTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// JunkyVaultManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use JunkyVaultManagerMetaData.ABI instead.
var JunkyVaultManagerABI = JunkyVaultManagerMetaData.ABI

// JunkyVaultManager is an auto generated Go binding around an Ethereum contract.
type JunkyVaultManager struct {
	JunkyVaultManagerCaller     // Read-only binding to the contract
	JunkyVaultManagerTransactor // Write-only binding to the contract
	JunkyVaultManagerFilterer   // Log filterer for contract events
}

// JunkyVaultManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type JunkyVaultManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JunkyVaultManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JunkyVaultManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JunkyVaultManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JunkyVaultManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JunkyVaultManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JunkyVaultManagerSession struct {
	Contract     *JunkyVaultManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// JunkyVaultManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JunkyVaultManagerCallerSession struct {
	Contract *JunkyVaultManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// JunkyVaultManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JunkyVaultManagerTransactorSession struct {
	Contract     *JunkyVaultManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// JunkyVaultManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type JunkyVaultManagerRaw struct {
	Contract *JunkyVaultManager // Generic contract binding to access the raw methods on
}

// JunkyVaultManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JunkyVaultManagerCallerRaw struct {
	Contract *JunkyVaultManagerCaller // Generic read-only contract binding to access the raw methods on
}

// JunkyVaultManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JunkyVaultManagerTransactorRaw struct {
	Contract *JunkyVaultManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJunkyVaultManager creates a new instance of JunkyVaultManager, bound to a specific deployed contract.
func NewJunkyVaultManager(address common.Address, backend bind.ContractBackend) (*JunkyVaultManager, error) {
	contract, err := bindJunkyVaultManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManager{JunkyVaultManagerCaller: JunkyVaultManagerCaller{contract: contract}, JunkyVaultManagerTransactor: JunkyVaultManagerTransactor{contract: contract}, JunkyVaultManagerFilterer: JunkyVaultManagerFilterer{contract: contract}}, nil
}

// NewJunkyVaultManagerCaller creates a new read-only instance of JunkyVaultManager, bound to a specific deployed contract.
func NewJunkyVaultManagerCaller(address common.Address, caller bind.ContractCaller) (*JunkyVaultManagerCaller, error) {
	contract, err := bindJunkyVaultManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManagerCaller{contract: contract}, nil
}

// NewJunkyVaultManagerTransactor creates a new write-only instance of JunkyVaultManager, bound to a specific deployed contract.
func NewJunkyVaultManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*JunkyVaultManagerTransactor, error) {
	contract, err := bindJunkyVaultManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManagerTransactor{contract: contract}, nil
}

// NewJunkyVaultManagerFilterer creates a new log filterer instance of JunkyVaultManager, bound to a specific deployed contract.
func NewJunkyVaultManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*JunkyVaultManagerFilterer, error) {
	contract, err := bindJunkyVaultManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManagerFilterer{contract: contract}, nil
}

// bindJunkyVaultManager binds a generic wrapper to an already deployed contract.
func bindJunkyVaultManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JunkyVaultManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JunkyVaultManager *JunkyVaultManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JunkyVaultManager.Contract.JunkyVaultManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JunkyVaultManager *JunkyVaultManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.JunkyVaultManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JunkyVaultManager *JunkyVaultManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.JunkyVaultManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JunkyVaultManager *JunkyVaultManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JunkyVaultManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JunkyVaultManager *JunkyVaultManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JunkyVaultManager *JunkyVaultManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.contract.Transact(opts, method, params...)
}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerCaller) Entropy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "entropy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerSession) Entropy() (common.Address, error) {
	return _JunkyVaultManager.Contract.Entropy(&_JunkyVaultManager.CallOpts)
}

// Entropy is a free data retrieval call binding the contract method 0x47ce07cc.
//
// Solidity: function entropy() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) Entropy() (common.Address, error) {
	return _JunkyVaultManager.Contract.Entropy(&_JunkyVaultManager.CallOpts)
}

// EntropyProvider is a free data retrieval call binding the contract method 0x5bf414ac.
//
// Solidity: function entropyProvider() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerCaller) EntropyProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "entropyProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntropyProvider is a free data retrieval call binding the contract method 0x5bf414ac.
//
// Solidity: function entropyProvider() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerSession) EntropyProvider() (common.Address, error) {
	return _JunkyVaultManager.Contract.EntropyProvider(&_JunkyVaultManager.CallOpts)
}

// EntropyProvider is a free data retrieval call binding the contract method 0x5bf414ac.
//
// Solidity: function entropyProvider() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) EntropyProvider() (common.Address, error) {
	return _JunkyVaultManager.Contract.EntropyProvider(&_JunkyVaultManager.CallOpts)
}

// Games is a free data retrieval call binding the contract method 0xf54815b4.
//
// Solidity: function games(uint64 sequenceNumber) view returns(bytes)
func (_JunkyVaultManager *JunkyVaultManagerCaller) Games(opts *bind.CallOpts, sequenceNumber uint64) ([]byte, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "games", sequenceNumber)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Games is a free data retrieval call binding the contract method 0xf54815b4.
//
// Solidity: function games(uint64 sequenceNumber) view returns(bytes)
func (_JunkyVaultManager *JunkyVaultManagerSession) Games(sequenceNumber uint64) ([]byte, error) {
	return _JunkyVaultManager.Contract.Games(&_JunkyVaultManager.CallOpts, sequenceNumber)
}

// Games is a free data retrieval call binding the contract method 0xf54815b4.
//
// Solidity: function games(uint64 sequenceNumber) view returns(bytes)
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) Games(sequenceNumber uint64) ([]byte, error) {
	return _JunkyVaultManager.Contract.Games(&_JunkyVaultManager.CallOpts, sequenceNumber)
}

// GetAllVaults is a free data retrieval call binding the contract method 0x97331bf9.
//
// Solidity: function getAllVaults() view returns((address,address,uint256,uint256,uint256)[])
func (_JunkyVaultManager *JunkyVaultManagerCaller) GetAllVaults(opts *bind.CallOpts) ([]VaultManagerProxyVaultWithPrice, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "getAllVaults")

	if err != nil {
		return *new([]VaultManagerProxyVaultWithPrice), err
	}

	out0 := *abi.ConvertType(out[0], new([]VaultManagerProxyVaultWithPrice)).(*[]VaultManagerProxyVaultWithPrice)

	return out0, err

}

// GetAllVaults is a free data retrieval call binding the contract method 0x97331bf9.
//
// Solidity: function getAllVaults() view returns((address,address,uint256,uint256,uint256)[])
func (_JunkyVaultManager *JunkyVaultManagerSession) GetAllVaults() ([]VaultManagerProxyVaultWithPrice, error) {
	return _JunkyVaultManager.Contract.GetAllVaults(&_JunkyVaultManager.CallOpts)
}

// GetAllVaults is a free data retrieval call binding the contract method 0x97331bf9.
//
// Solidity: function getAllVaults() view returns((address,address,uint256,uint256,uint256)[])
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) GetAllVaults() ([]VaultManagerProxyVaultWithPrice, error) {
	return _JunkyVaultManager.Contract.GetAllVaults(&_JunkyVaultManager.CallOpts)
}

// GetGame is a free data retrieval call binding the contract method 0x49041903.
//
// Solidity: function getGame(uint64 sequenceNumber) view returns((address,address,uint256,bytes32,uint8,uint8))
func (_JunkyVaultManager *JunkyVaultManagerCaller) GetGame(opts *bind.CallOpts, sequenceNumber uint64) (BaseGameContractProxyGameConfig, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "getGame", sequenceNumber)

	if err != nil {
		return *new(BaseGameContractProxyGameConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BaseGameContractProxyGameConfig)).(*BaseGameContractProxyGameConfig)

	return out0, err

}

// GetGame is a free data retrieval call binding the contract method 0x49041903.
//
// Solidity: function getGame(uint64 sequenceNumber) view returns((address,address,uint256,bytes32,uint8,uint8))
func (_JunkyVaultManager *JunkyVaultManagerSession) GetGame(sequenceNumber uint64) (BaseGameContractProxyGameConfig, error) {
	return _JunkyVaultManager.Contract.GetGame(&_JunkyVaultManager.CallOpts, sequenceNumber)
}

// GetGame is a free data retrieval call binding the contract method 0x49041903.
//
// Solidity: function getGame(uint64 sequenceNumber) view returns((address,address,uint256,bytes32,uint8,uint8))
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) GetGame(sequenceNumber uint64) (BaseGameContractProxyGameConfig, error) {
	return _JunkyVaultManager.Contract.GetGame(&_JunkyVaultManager.CallOpts, sequenceNumber)
}

// GetVault is a free data retrieval call binding the contract method 0x0eb9af38.
//
// Solidity: function getVault(address liqTokenAddress) view returns((address,address,uint256,uint256,uint256))
func (_JunkyVaultManager *JunkyVaultManagerCaller) GetVault(opts *bind.CallOpts, liqTokenAddress common.Address) (VaultManagerProxyVaultWithPrice, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "getVault", liqTokenAddress)

	if err != nil {
		return *new(VaultManagerProxyVaultWithPrice), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultManagerProxyVaultWithPrice)).(*VaultManagerProxyVaultWithPrice)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x0eb9af38.
//
// Solidity: function getVault(address liqTokenAddress) view returns((address,address,uint256,uint256,uint256))
func (_JunkyVaultManager *JunkyVaultManagerSession) GetVault(liqTokenAddress common.Address) (VaultManagerProxyVaultWithPrice, error) {
	return _JunkyVaultManager.Contract.GetVault(&_JunkyVaultManager.CallOpts, liqTokenAddress)
}

// GetVault is a free data retrieval call binding the contract method 0x0eb9af38.
//
// Solidity: function getVault(address liqTokenAddress) view returns((address,address,uint256,uint256,uint256))
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) GetVault(liqTokenAddress common.Address) (VaultManagerProxyVaultWithPrice, error) {
	return _JunkyVaultManager.Contract.GetVault(&_JunkyVaultManager.CallOpts, liqTokenAddress)
}

// GetVaultTokenPrice is a free data retrieval call binding the contract method 0x112f7b40.
//
// Solidity: function getVaultTokenPrice(address liqTokenAddress) view returns(uint256)
func (_JunkyVaultManager *JunkyVaultManagerCaller) GetVaultTokenPrice(opts *bind.CallOpts, liqTokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "getVaultTokenPrice", liqTokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultTokenPrice is a free data retrieval call binding the contract method 0x112f7b40.
//
// Solidity: function getVaultTokenPrice(address liqTokenAddress) view returns(uint256)
func (_JunkyVaultManager *JunkyVaultManagerSession) GetVaultTokenPrice(liqTokenAddress common.Address) (*big.Int, error) {
	return _JunkyVaultManager.Contract.GetVaultTokenPrice(&_JunkyVaultManager.CallOpts, liqTokenAddress)
}

// GetVaultTokenPrice is a free data retrieval call binding the contract method 0x112f7b40.
//
// Solidity: function getVaultTokenPrice(address liqTokenAddress) view returns(uint256)
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) GetVaultTokenPrice(liqTokenAddress common.Address) (*big.Int, error) {
	return _JunkyVaultManager.Contract.GetVaultTokenPrice(&_JunkyVaultManager.CallOpts, liqTokenAddress)
}

// HouseEdge is a free data retrieval call binding the contract method 0xd667dcd7.
//
// Solidity: function houseEdge() view returns(uint256)
func (_JunkyVaultManager *JunkyVaultManagerCaller) HouseEdge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "houseEdge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HouseEdge is a free data retrieval call binding the contract method 0xd667dcd7.
//
// Solidity: function houseEdge() view returns(uint256)
func (_JunkyVaultManager *JunkyVaultManagerSession) HouseEdge() (*big.Int, error) {
	return _JunkyVaultManager.Contract.HouseEdge(&_JunkyVaultManager.CallOpts)
}

// HouseEdge is a free data retrieval call binding the contract method 0xd667dcd7.
//
// Solidity: function houseEdge() view returns(uint256)
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) HouseEdge() (*big.Int, error) {
	return _JunkyVaultManager.Contract.HouseEdge(&_JunkyVaultManager.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_JunkyVaultManager *JunkyVaultManagerCaller) MinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "minAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_JunkyVaultManager *JunkyVaultManagerSession) MinAmount() (*big.Int, error) {
	return _JunkyVaultManager.Contract.MinAmount(&_JunkyVaultManager.CallOpts)
}

// MinAmount is a free data retrieval call binding the contract method 0x9b2cb5d8.
//
// Solidity: function minAmount() view returns(uint256)
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) MinAmount() (*big.Int, error) {
	return _JunkyVaultManager.Contract.MinAmount(&_JunkyVaultManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerSession) Owner() (common.Address, error) {
	return _JunkyVaultManager.Contract.Owner(&_JunkyVaultManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) Owner() (common.Address, error) {
	return _JunkyVaultManager.Contract.Owner(&_JunkyVaultManager.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerSession) Treasury() (common.Address, error) {
	return _JunkyVaultManager.Contract.Treasury(&_JunkyVaultManager.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) Treasury() (common.Address, error) {
	return _JunkyVaultManager.Contract.Treasury(&_JunkyVaultManager.CallOpts)
}

// UserDeposits is a free data retrieval call binding the contract method 0x436d8039.
//
// Solidity: function userDeposits(address , address ) view returns(uint256 initialAmount, uint256 vaultTokens)
func (_JunkyVaultManager *JunkyVaultManagerCaller) UserDeposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	InitialAmount *big.Int
	VaultTokens   *big.Int
}, error) {
	var out []interface{}
	err := _JunkyVaultManager.contract.Call(opts, &out, "userDeposits", arg0, arg1)

	outstruct := new(struct {
		InitialAmount *big.Int
		VaultTokens   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VaultTokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserDeposits is a free data retrieval call binding the contract method 0x436d8039.
//
// Solidity: function userDeposits(address , address ) view returns(uint256 initialAmount, uint256 vaultTokens)
func (_JunkyVaultManager *JunkyVaultManagerSession) UserDeposits(arg0 common.Address, arg1 common.Address) (struct {
	InitialAmount *big.Int
	VaultTokens   *big.Int
}, error) {
	return _JunkyVaultManager.Contract.UserDeposits(&_JunkyVaultManager.CallOpts, arg0, arg1)
}

// UserDeposits is a free data retrieval call binding the contract method 0x436d8039.
//
// Solidity: function userDeposits(address , address ) view returns(uint256 initialAmount, uint256 vaultTokens)
func (_JunkyVaultManager *JunkyVaultManagerCallerSession) UserDeposits(arg0 common.Address, arg1 common.Address) (struct {
	InitialAmount *big.Int
	VaultTokens   *big.Int
}, error) {
	return _JunkyVaultManager.Contract.UserDeposits(&_JunkyVaultManager.CallOpts, arg0, arg1)
}

// EntropyCallback is a paid mutator transaction binding the contract method 0x52a5f1f8.
//
// Solidity: function _entropyCallback(uint64 sequence, address provider, bytes32 randomNumber) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactor) EntropyCallback(opts *bind.TransactOpts, sequence uint64, provider common.Address, randomNumber [32]byte) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.Transact(opts, "_entropyCallback", sequence, provider, randomNumber)
}

// EntropyCallback is a paid mutator transaction binding the contract method 0x52a5f1f8.
//
// Solidity: function _entropyCallback(uint64 sequence, address provider, bytes32 randomNumber) returns()
func (_JunkyVaultManager *JunkyVaultManagerSession) EntropyCallback(sequence uint64, provider common.Address, randomNumber [32]byte) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.EntropyCallback(&_JunkyVaultManager.TransactOpts, sequence, provider, randomNumber)
}

// EntropyCallback is a paid mutator transaction binding the contract method 0x52a5f1f8.
//
// Solidity: function _entropyCallback(uint64 sequence, address provider, bytes32 randomNumber) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) EntropyCallback(sequence uint64, provider common.Address, randomNumber [32]byte) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.EntropyCallback(&_JunkyVaultManager.TransactOpts, sequence, provider, randomNumber)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address liqTokenAddress, uint256 amount) payable returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactor) AddLiquidity(opts *bind.TransactOpts, liqTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.Transact(opts, "addLiquidity", liqTokenAddress, amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address liqTokenAddress, uint256 amount) payable returns()
func (_JunkyVaultManager *JunkyVaultManagerSession) AddLiquidity(liqTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.AddLiquidity(&_JunkyVaultManager.TransactOpts, liqTokenAddress, amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address liqTokenAddress, uint256 amount) payable returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) AddLiquidity(liqTokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.AddLiquidity(&_JunkyVaultManager.TransactOpts, liqTokenAddress, amount)
}

// CreateVault is a paid mutator transaction binding the contract method 0xfbd4834c.
//
// Solidity: function createVault(string name, string symbol, address liqTokenAddress) returns(address)
func (_JunkyVaultManager *JunkyVaultManagerTransactor) CreateVault(opts *bind.TransactOpts, name string, symbol string, liqTokenAddress common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.Transact(opts, "createVault", name, symbol, liqTokenAddress)
}

// CreateVault is a paid mutator transaction binding the contract method 0xfbd4834c.
//
// Solidity: function createVault(string name, string symbol, address liqTokenAddress) returns(address)
func (_JunkyVaultManager *JunkyVaultManagerSession) CreateVault(name string, symbol string, liqTokenAddress common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.CreateVault(&_JunkyVaultManager.TransactOpts, name, symbol, liqTokenAddress)
}

// CreateVault is a paid mutator transaction binding the contract method 0xfbd4834c.
//
// Solidity: function createVault(string name, string symbol, address liqTokenAddress) returns(address)
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) CreateVault(name string, symbol string, liqTokenAddress common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.CreateVault(&_JunkyVaultManager.TransactOpts, name, symbol, liqTokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address treasuryAddress, address entropyAddress, address entropyProviderAddress) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactor) Initialize(opts *bind.TransactOpts, treasuryAddress common.Address, entropyAddress common.Address, entropyProviderAddress common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.Transact(opts, "initialize", treasuryAddress, entropyAddress, entropyProviderAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address treasuryAddress, address entropyAddress, address entropyProviderAddress) returns()
func (_JunkyVaultManager *JunkyVaultManagerSession) Initialize(treasuryAddress common.Address, entropyAddress common.Address, entropyProviderAddress common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.Initialize(&_JunkyVaultManager.TransactOpts, treasuryAddress, entropyAddress, entropyProviderAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address treasuryAddress, address entropyAddress, address entropyProviderAddress) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) Initialize(treasuryAddress common.Address, entropyAddress common.Address, entropyProviderAddress common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.Initialize(&_JunkyVaultManager.TransactOpts, treasuryAddress, entropyAddress, entropyProviderAddress)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address liqTokenAddress, uint256 vaultTokens) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactor) RemoveLiquidity(opts *bind.TransactOpts, liqTokenAddress common.Address, vaultTokens *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.Transact(opts, "removeLiquidity", liqTokenAddress, vaultTokens)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address liqTokenAddress, uint256 vaultTokens) returns()
func (_JunkyVaultManager *JunkyVaultManagerSession) RemoveLiquidity(liqTokenAddress common.Address, vaultTokens *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.RemoveLiquidity(&_JunkyVaultManager.TransactOpts, liqTokenAddress, vaultTokens)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xa201ccf6.
//
// Solidity: function removeLiquidity(address liqTokenAddress, uint256 vaultTokens) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) RemoveLiquidity(liqTokenAddress common.Address, vaultTokens *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.RemoveLiquidity(&_JunkyVaultManager.TransactOpts, liqTokenAddress, vaultTokens)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JunkyVaultManager *JunkyVaultManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.RenounceOwnership(&_JunkyVaultManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.RenounceOwnership(&_JunkyVaultManager.TransactOpts)
}

// SetHouseEdge is a paid mutator transaction binding the contract method 0x6cd0f102.
//
// Solidity: function setHouseEdge(uint256 newEdge) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactor) SetHouseEdge(opts *bind.TransactOpts, newEdge *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.Transact(opts, "setHouseEdge", newEdge)
}

// SetHouseEdge is a paid mutator transaction binding the contract method 0x6cd0f102.
//
// Solidity: function setHouseEdge(uint256 newEdge) returns()
func (_JunkyVaultManager *JunkyVaultManagerSession) SetHouseEdge(newEdge *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.SetHouseEdge(&_JunkyVaultManager.TransactOpts, newEdge)
}

// SetHouseEdge is a paid mutator transaction binding the contract method 0x6cd0f102.
//
// Solidity: function setHouseEdge(uint256 newEdge) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) SetHouseEdge(newEdge *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.SetHouseEdge(&_JunkyVaultManager.TransactOpts, newEdge)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactor) SetMinAmount(opts *bind.TransactOpts, _minAmount *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.Transact(opts, "setMinAmount", _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_JunkyVaultManager *JunkyVaultManagerSession) SetMinAmount(_minAmount *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.SetMinAmount(&_JunkyVaultManager.TransactOpts, _minAmount)
}

// SetMinAmount is a paid mutator transaction binding the contract method 0x897b0637.
//
// Solidity: function setMinAmount(uint256 _minAmount) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) SetMinAmount(_minAmount *big.Int) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.SetMinAmount(&_JunkyVaultManager.TransactOpts, _minAmount)
}

// SetTreasuryContract is a paid mutator transaction binding the contract method 0x1ed65110.
//
// Solidity: function setTreasuryContract(address newTreasury) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactor) SetTreasuryContract(opts *bind.TransactOpts, newTreasury common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.Transact(opts, "setTreasuryContract", newTreasury)
}

// SetTreasuryContract is a paid mutator transaction binding the contract method 0x1ed65110.
//
// Solidity: function setTreasuryContract(address newTreasury) returns()
func (_JunkyVaultManager *JunkyVaultManagerSession) SetTreasuryContract(newTreasury common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.SetTreasuryContract(&_JunkyVaultManager.TransactOpts, newTreasury)
}

// SetTreasuryContract is a paid mutator transaction binding the contract method 0x1ed65110.
//
// Solidity: function setTreasuryContract(address newTreasury) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) SetTreasuryContract(newTreasury common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.SetTreasuryContract(&_JunkyVaultManager.TransactOpts, newTreasury)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JunkyVaultManager *JunkyVaultManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.TransferOwnership(&_JunkyVaultManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.TransferOwnership(&_JunkyVaultManager.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JunkyVaultManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_JunkyVaultManager *JunkyVaultManagerSession) Receive() (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.Receive(&_JunkyVaultManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_JunkyVaultManager *JunkyVaultManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _JunkyVaultManager.Contract.Receive(&_JunkyVaultManager.TransactOpts)
}

// JunkyVaultManagerGameResultIterator is returned from FilterGameResult and is used to iterate over the raw logs and unpacked data for GameResult events raised by the JunkyVaultManager contract.
type JunkyVaultManagerGameResultIterator struct {
	Event *JunkyVaultManagerGameResult // Event containing the contract specifics and raw log

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
func (it *JunkyVaultManagerGameResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunkyVaultManagerGameResult)
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
		it.Event = new(JunkyVaultManagerGameResult)
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
func (it *JunkyVaultManagerGameResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunkyVaultManagerGameResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunkyVaultManagerGameResult represents a GameResult event raised by the JunkyVaultManager contract.
type JunkyVaultManagerGameResult struct {
	Player       common.Address
	Payout       *big.Int
	RandomNumber [32]byte
	WonCount     uint8
	TotalCount   uint8
	Token        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGameResult is a free log retrieval operation binding the contract event 0x2411364aba5a01e505ec0db89f7797ec95b07ed5aa9a61702a376428ef738f0a.
//
// Solidity: event GameResult(address indexed player, uint256 payout, bytes32 indexed randomNumber, uint8 wonCount, uint8 totalCount, address indexed token)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) FilterGameResult(opts *bind.FilterOpts, player []common.Address, randomNumber [][32]byte, token []common.Address) (*JunkyVaultManagerGameResultIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	var randomNumberRule []interface{}
	for _, randomNumberItem := range randomNumber {
		randomNumberRule = append(randomNumberRule, randomNumberItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.FilterLogs(opts, "GameResult", playerRule, randomNumberRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManagerGameResultIterator{contract: _JunkyVaultManager.contract, event: "GameResult", logs: logs, sub: sub}, nil
}

// WatchGameResult is a free log subscription operation binding the contract event 0x2411364aba5a01e505ec0db89f7797ec95b07ed5aa9a61702a376428ef738f0a.
//
// Solidity: event GameResult(address indexed player, uint256 payout, bytes32 indexed randomNumber, uint8 wonCount, uint8 totalCount, address indexed token)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) WatchGameResult(opts *bind.WatchOpts, sink chan<- *JunkyVaultManagerGameResult, player []common.Address, randomNumber [][32]byte, token []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	var randomNumberRule []interface{}
	for _, randomNumberItem := range randomNumber {
		randomNumberRule = append(randomNumberRule, randomNumberItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.WatchLogs(opts, "GameResult", playerRule, randomNumberRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunkyVaultManagerGameResult)
				if err := _JunkyVaultManager.contract.UnpackLog(event, "GameResult", log); err != nil {
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

// ParseGameResult is a log parse operation binding the contract event 0x2411364aba5a01e505ec0db89f7797ec95b07ed5aa9a61702a376428ef738f0a.
//
// Solidity: event GameResult(address indexed player, uint256 payout, bytes32 indexed randomNumber, uint8 wonCount, uint8 totalCount, address indexed token)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) ParseGameResult(log types.Log) (*JunkyVaultManagerGameResult, error) {
	event := new(JunkyVaultManagerGameResult)
	if err := _JunkyVaultManager.contract.UnpackLog(event, "GameResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunkyVaultManagerGameStartedIterator is returned from FilterGameStarted and is used to iterate over the raw logs and unpacked data for GameStarted events raised by the JunkyVaultManager contract.
type JunkyVaultManagerGameStartedIterator struct {
	Event *JunkyVaultManagerGameStarted // Event containing the contract specifics and raw log

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
func (it *JunkyVaultManagerGameStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunkyVaultManagerGameStarted)
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
		it.Event = new(JunkyVaultManagerGameStarted)
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
func (it *JunkyVaultManagerGameStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunkyVaultManagerGameStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunkyVaultManagerGameStarted represents a GameStarted event raised by the JunkyVaultManager contract.
type JunkyVaultManagerGameStarted struct {
	SequenceNumber uint64
	Player         common.Address
	Wager          *big.Int
	Count          uint8
	Token          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGameStarted is a free log retrieval operation binding the contract event 0x55523b392929e95430d6b23b94fd06bcd7e81430ca6e97ba40e82dedbb3b0614.
//
// Solidity: event GameStarted(uint64 sequenceNumber, address indexed player, uint256 wager, uint8 count, address indexed token)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) FilterGameStarted(opts *bind.FilterOpts, player []common.Address, token []common.Address) (*JunkyVaultManagerGameStartedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.FilterLogs(opts, "GameStarted", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManagerGameStartedIterator{contract: _JunkyVaultManager.contract, event: "GameStarted", logs: logs, sub: sub}, nil
}

// WatchGameStarted is a free log subscription operation binding the contract event 0x55523b392929e95430d6b23b94fd06bcd7e81430ca6e97ba40e82dedbb3b0614.
//
// Solidity: event GameStarted(uint64 sequenceNumber, address indexed player, uint256 wager, uint8 count, address indexed token)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) WatchGameStarted(opts *bind.WatchOpts, sink chan<- *JunkyVaultManagerGameStarted, player []common.Address, token []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.WatchLogs(opts, "GameStarted", playerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunkyVaultManagerGameStarted)
				if err := _JunkyVaultManager.contract.UnpackLog(event, "GameStarted", log); err != nil {
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

// ParseGameStarted is a log parse operation binding the contract event 0x55523b392929e95430d6b23b94fd06bcd7e81430ca6e97ba40e82dedbb3b0614.
//
// Solidity: event GameStarted(uint64 sequenceNumber, address indexed player, uint256 wager, uint8 count, address indexed token)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) ParseGameStarted(log types.Log) (*JunkyVaultManagerGameStarted, error) {
	event := new(JunkyVaultManagerGameStarted)
	if err := _JunkyVaultManager.contract.UnpackLog(event, "GameStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunkyVaultManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the JunkyVaultManager contract.
type JunkyVaultManagerInitializedIterator struct {
	Event *JunkyVaultManagerInitialized // Event containing the contract specifics and raw log

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
func (it *JunkyVaultManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunkyVaultManagerInitialized)
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
		it.Event = new(JunkyVaultManagerInitialized)
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
func (it *JunkyVaultManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunkyVaultManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunkyVaultManagerInitialized represents a Initialized event raised by the JunkyVaultManager contract.
type JunkyVaultManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*JunkyVaultManagerInitializedIterator, error) {

	logs, sub, err := _JunkyVaultManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManagerInitializedIterator{contract: _JunkyVaultManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *JunkyVaultManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _JunkyVaultManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunkyVaultManagerInitialized)
				if err := _JunkyVaultManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_JunkyVaultManager *JunkyVaultManagerFilterer) ParseInitialized(log types.Log) (*JunkyVaultManagerInitialized, error) {
	event := new(JunkyVaultManagerInitialized)
	if err := _JunkyVaultManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunkyVaultManagerLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the JunkyVaultManager contract.
type JunkyVaultManagerLiquidityAddedIterator struct {
	Event *JunkyVaultManagerLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *JunkyVaultManagerLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunkyVaultManagerLiquidityAdded)
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
		it.Event = new(JunkyVaultManagerLiquidityAdded)
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
func (it *JunkyVaultManagerLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunkyVaultManagerLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunkyVaultManagerLiquidityAdded represents a LiquidityAdded event raised by the JunkyVaultManager contract.
type JunkyVaultManagerLiquidityAdded struct {
	User            common.Address
	LiqTokenAddress common.Address
	Amount          *big.Int
	VaultTokens     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0x36f3b2e1a21c19137dd82ec243b0708a1d26b3d1fa1dc49c44c4c366a5878138.
//
// Solidity: event LiquidityAdded(address indexed user, address indexed liqTokenAddress, uint256 amount, uint256 vaultTokens)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, user []common.Address, liqTokenAddress []common.Address) (*JunkyVaultManagerLiquidityAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var liqTokenAddressRule []interface{}
	for _, liqTokenAddressItem := range liqTokenAddress {
		liqTokenAddressRule = append(liqTokenAddressRule, liqTokenAddressItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.FilterLogs(opts, "LiquidityAdded", userRule, liqTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManagerLiquidityAddedIterator{contract: _JunkyVaultManager.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0x36f3b2e1a21c19137dd82ec243b0708a1d26b3d1fa1dc49c44c4c366a5878138.
//
// Solidity: event LiquidityAdded(address indexed user, address indexed liqTokenAddress, uint256 amount, uint256 vaultTokens)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *JunkyVaultManagerLiquidityAdded, user []common.Address, liqTokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var liqTokenAddressRule []interface{}
	for _, liqTokenAddressItem := range liqTokenAddress {
		liqTokenAddressRule = append(liqTokenAddressRule, liqTokenAddressItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.WatchLogs(opts, "LiquidityAdded", userRule, liqTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunkyVaultManagerLiquidityAdded)
				if err := _JunkyVaultManager.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0x36f3b2e1a21c19137dd82ec243b0708a1d26b3d1fa1dc49c44c4c366a5878138.
//
// Solidity: event LiquidityAdded(address indexed user, address indexed liqTokenAddress, uint256 amount, uint256 vaultTokens)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) ParseLiquidityAdded(log types.Log) (*JunkyVaultManagerLiquidityAdded, error) {
	event := new(JunkyVaultManagerLiquidityAdded)
	if err := _JunkyVaultManager.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunkyVaultManagerLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the JunkyVaultManager contract.
type JunkyVaultManagerLiquidityRemovedIterator struct {
	Event *JunkyVaultManagerLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *JunkyVaultManagerLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunkyVaultManagerLiquidityRemoved)
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
		it.Event = new(JunkyVaultManagerLiquidityRemoved)
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
func (it *JunkyVaultManagerLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunkyVaultManagerLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunkyVaultManagerLiquidityRemoved represents a LiquidityRemoved event raised by the JunkyVaultManager contract.
type JunkyVaultManagerLiquidityRemoved struct {
	User            common.Address
	LiqTokenAddress common.Address
	Amount          *big.Int
	VaultTokens     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemoved is a free log retrieval operation binding the contract event 0x3b5c196aff80bb96c03b41c96906b66827014de931d1b36e0ede6ee8caeb4bf9.
//
// Solidity: event LiquidityRemoved(address indexed user, address indexed liqTokenAddress, uint256 amount, uint256 vaultTokens)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, user []common.Address, liqTokenAddress []common.Address) (*JunkyVaultManagerLiquidityRemovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var liqTokenAddressRule []interface{}
	for _, liqTokenAddressItem := range liqTokenAddress {
		liqTokenAddressRule = append(liqTokenAddressRule, liqTokenAddressItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.FilterLogs(opts, "LiquidityRemoved", userRule, liqTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManagerLiquidityRemovedIterator{contract: _JunkyVaultManager.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0x3b5c196aff80bb96c03b41c96906b66827014de931d1b36e0ede6ee8caeb4bf9.
//
// Solidity: event LiquidityRemoved(address indexed user, address indexed liqTokenAddress, uint256 amount, uint256 vaultTokens)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *JunkyVaultManagerLiquidityRemoved, user []common.Address, liqTokenAddress []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var liqTokenAddressRule []interface{}
	for _, liqTokenAddressItem := range liqTokenAddress {
		liqTokenAddressRule = append(liqTokenAddressRule, liqTokenAddressItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.WatchLogs(opts, "LiquidityRemoved", userRule, liqTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunkyVaultManagerLiquidityRemoved)
				if err := _JunkyVaultManager.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

// ParseLiquidityRemoved is a log parse operation binding the contract event 0x3b5c196aff80bb96c03b41c96906b66827014de931d1b36e0ede6ee8caeb4bf9.
//
// Solidity: event LiquidityRemoved(address indexed user, address indexed liqTokenAddress, uint256 amount, uint256 vaultTokens)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) ParseLiquidityRemoved(log types.Log) (*JunkyVaultManagerLiquidityRemoved, error) {
	event := new(JunkyVaultManagerLiquidityRemoved)
	if err := _JunkyVaultManager.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunkyVaultManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the JunkyVaultManager contract.
type JunkyVaultManagerOwnershipTransferredIterator struct {
	Event *JunkyVaultManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *JunkyVaultManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunkyVaultManagerOwnershipTransferred)
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
		it.Event = new(JunkyVaultManagerOwnershipTransferred)
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
func (it *JunkyVaultManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunkyVaultManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunkyVaultManagerOwnershipTransferred represents a OwnershipTransferred event raised by the JunkyVaultManager contract.
type JunkyVaultManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*JunkyVaultManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManagerOwnershipTransferredIterator{contract: _JunkyVaultManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *JunkyVaultManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunkyVaultManagerOwnershipTransferred)
				if err := _JunkyVaultManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_JunkyVaultManager *JunkyVaultManagerFilterer) ParseOwnershipTransferred(log types.Log) (*JunkyVaultManagerOwnershipTransferred, error) {
	event := new(JunkyVaultManagerOwnershipTransferred)
	if err := _JunkyVaultManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JunkyVaultManagerVaultCreatedIterator is returned from FilterVaultCreated and is used to iterate over the raw logs and unpacked data for VaultCreated events raised by the JunkyVaultManager contract.
type JunkyVaultManagerVaultCreatedIterator struct {
	Event *JunkyVaultManagerVaultCreated // Event containing the contract specifics and raw log

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
func (it *JunkyVaultManagerVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JunkyVaultManagerVaultCreated)
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
		it.Event = new(JunkyVaultManagerVaultCreated)
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
func (it *JunkyVaultManagerVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JunkyVaultManagerVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JunkyVaultManagerVaultCreated represents a VaultCreated event raised by the JunkyVaultManager contract.
type JunkyVaultManagerVaultCreated struct {
	LiqTokenAddress   common.Address
	VaultTokenAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVaultCreated is a free log retrieval operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address indexed liqTokenAddress, address indexed vaultTokenAddress)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) FilterVaultCreated(opts *bind.FilterOpts, liqTokenAddress []common.Address, vaultTokenAddress []common.Address) (*JunkyVaultManagerVaultCreatedIterator, error) {

	var liqTokenAddressRule []interface{}
	for _, liqTokenAddressItem := range liqTokenAddress {
		liqTokenAddressRule = append(liqTokenAddressRule, liqTokenAddressItem)
	}
	var vaultTokenAddressRule []interface{}
	for _, vaultTokenAddressItem := range vaultTokenAddress {
		vaultTokenAddressRule = append(vaultTokenAddressRule, vaultTokenAddressItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.FilterLogs(opts, "VaultCreated", liqTokenAddressRule, vaultTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &JunkyVaultManagerVaultCreatedIterator{contract: _JunkyVaultManager.contract, event: "VaultCreated", logs: logs, sub: sub}, nil
}

// WatchVaultCreated is a free log subscription operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address indexed liqTokenAddress, address indexed vaultTokenAddress)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) WatchVaultCreated(opts *bind.WatchOpts, sink chan<- *JunkyVaultManagerVaultCreated, liqTokenAddress []common.Address, vaultTokenAddress []common.Address) (event.Subscription, error) {

	var liqTokenAddressRule []interface{}
	for _, liqTokenAddressItem := range liqTokenAddress {
		liqTokenAddressRule = append(liqTokenAddressRule, liqTokenAddressItem)
	}
	var vaultTokenAddressRule []interface{}
	for _, vaultTokenAddressItem := range vaultTokenAddress {
		vaultTokenAddressRule = append(vaultTokenAddressRule, vaultTokenAddressItem)
	}

	logs, sub, err := _JunkyVaultManager.contract.WatchLogs(opts, "VaultCreated", liqTokenAddressRule, vaultTokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JunkyVaultManagerVaultCreated)
				if err := _JunkyVaultManager.contract.UnpackLog(event, "VaultCreated", log); err != nil {
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

// ParseVaultCreated is a log parse operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address indexed liqTokenAddress, address indexed vaultTokenAddress)
func (_JunkyVaultManager *JunkyVaultManagerFilterer) ParseVaultCreated(log types.Log) (*JunkyVaultManagerVaultCreated, error) {
	event := new(JunkyVaultManagerVaultCreated)
	if err := _JunkyVaultManager.contract.UnpackLog(event, "VaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
