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

// CurveMathCurveState is an auto generated low-level Go binding around an user-defined struct.
type CurveMathCurveState struct {
	PriceRoot    *big.Int
	AmbientSeeds *big.Int
	ConcLiq      *big.Int
	SeedDeflator uint64
	ConcGrowth   uint64
}

// PoolSpecsPool is an auto generated low-level Go binding around an user-defined struct.
type PoolSpecsPool struct {
	Schema       uint8
	FeeRate      uint16
	ProtocolTake uint8
	TickSize     uint16
	JitThresh    uint8
	KnockoutBits uint8
	OracleFlags  uint8
}

// CrocQueryMetaData contains all meta data concerning the CrocQuery contract.
var CrocQueryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"dex\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dex_\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryAmbientPosition\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seeds\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryAmbientTokens\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"liq\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"baseQty\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"quoteQty\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryConcRewards\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lowerTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"upperTick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"liqRewards\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"baseRewards\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"quoteRewards\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryCurve\",\"inputs\":[{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"curve\",\"type\":\"tuple\",\"internalType\":\"structCurveMath.CurveState\",\"components\":[{\"name\":\"priceRoot_\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"ambientSeeds_\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"concLiq_\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"seedDeflator_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"concGrowth_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryCurveTick\",\"inputs\":[{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryKnockoutMerkle\",\"inputs\":[{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isBid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"root\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"pivot\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"fee\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryKnockoutPivot\",\"inputs\":[{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isBid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"lots\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"pivot\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"range\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryKnockoutPos\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pivot\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isBid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lowerTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"upperTick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"lots\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"mileage\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryKnockoutTokens\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pivot\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isBid\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"lowerTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"upperTick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"liq\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"baseQty\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"quoteQty\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"knockedOut\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryLevel\",\"inputs\":[{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"bidLots\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"askLots\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"odometer\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryLiquidity\",\"inputs\":[{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryPoolParams\",\"inputs\":[{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"pool\",\"type\":\"tuple\",\"internalType\":\"structPoolSpecs.Pool\",\"components\":[{\"name\":\"schema_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"feeRate_\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"protocolTake_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tickSize_\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"jitThresh_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"knockoutBits_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"oracleFlags_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryPoolTemplate\",\"inputs\":[{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"pool\",\"type\":\"tuple\",\"internalType\":\"structPoolSpecs.Pool\",\"components\":[{\"name\":\"schema_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"feeRate_\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"protocolTake_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"tickSize_\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"jitThresh_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"knockoutBits_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"oracleFlags_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryPrice\",\"inputs\":[{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryProtocolAccum\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryRangePosition\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lowerTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"upperTick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"liq\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"fee\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"atomic\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryRangeTokens\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"base\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"quote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"poolIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lowerTick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"upperTick\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"outputs\":[{\"name\":\"liq\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"baseQty\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"quoteQty\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"querySurplus\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"surplus\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queryVirtual\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"surplus\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"}]",
}

// CrocQueryABI is the input ABI used to generate the binding from.
// Deprecated: Use CrocQueryMetaData.ABI instead.
var CrocQueryABI = CrocQueryMetaData.ABI

// CrocQuery is an auto generated Go binding around an Ethereum contract.
type CrocQuery struct {
	CrocQueryCaller     // Read-only binding to the contract
	CrocQueryTransactor // Write-only binding to the contract
	CrocQueryFilterer   // Log filterer for contract events
}

// CrocQueryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrocQueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrocQueryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrocQueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrocQueryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrocQueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrocQuerySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrocQuerySession struct {
	Contract     *CrocQuery        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrocQueryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrocQueryCallerSession struct {
	Contract *CrocQueryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrocQueryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrocQueryTransactorSession struct {
	Contract     *CrocQueryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrocQueryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrocQueryRaw struct {
	Contract *CrocQuery // Generic contract binding to access the raw methods on
}

// CrocQueryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrocQueryCallerRaw struct {
	Contract *CrocQueryCaller // Generic read-only contract binding to access the raw methods on
}

// CrocQueryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrocQueryTransactorRaw struct {
	Contract *CrocQueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrocQuery creates a new instance of CrocQuery, bound to a specific deployed contract.
func NewCrocQuery(address common.Address, backend bind.ContractBackend) (*CrocQuery, error) {
	contract, err := bindCrocQuery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrocQuery{CrocQueryCaller: CrocQueryCaller{contract: contract}, CrocQueryTransactor: CrocQueryTransactor{contract: contract}, CrocQueryFilterer: CrocQueryFilterer{contract: contract}}, nil
}

// NewCrocQueryCaller creates a new read-only instance of CrocQuery, bound to a specific deployed contract.
func NewCrocQueryCaller(address common.Address, caller bind.ContractCaller) (*CrocQueryCaller, error) {
	contract, err := bindCrocQuery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrocQueryCaller{contract: contract}, nil
}

// NewCrocQueryTransactor creates a new write-only instance of CrocQuery, bound to a specific deployed contract.
func NewCrocQueryTransactor(address common.Address, transactor bind.ContractTransactor) (*CrocQueryTransactor, error) {
	contract, err := bindCrocQuery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrocQueryTransactor{contract: contract}, nil
}

// NewCrocQueryFilterer creates a new log filterer instance of CrocQuery, bound to a specific deployed contract.
func NewCrocQueryFilterer(address common.Address, filterer bind.ContractFilterer) (*CrocQueryFilterer, error) {
	contract, err := bindCrocQuery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrocQueryFilterer{contract: contract}, nil
}

// bindCrocQuery binds a generic wrapper to an already deployed contract.
func bindCrocQuery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CrocQueryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrocQuery *CrocQueryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrocQuery.Contract.CrocQueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrocQuery *CrocQueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrocQuery.Contract.CrocQueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrocQuery *CrocQueryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrocQuery.Contract.CrocQueryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrocQuery *CrocQueryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrocQuery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrocQuery *CrocQueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrocQuery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrocQuery *CrocQueryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrocQuery.Contract.contract.Transact(opts, method, params...)
}

// Dex is a free data retrieval call binding the contract method 0x87834a0e.
//
// Solidity: function dex_() view returns(address)
func (_CrocQuery *CrocQueryCaller) Dex(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "dex_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dex is a free data retrieval call binding the contract method 0x87834a0e.
//
// Solidity: function dex_() view returns(address)
func (_CrocQuery *CrocQuerySession) Dex() (common.Address, error) {
	return _CrocQuery.Contract.Dex(&_CrocQuery.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x87834a0e.
//
// Solidity: function dex_() view returns(address)
func (_CrocQuery *CrocQueryCallerSession) Dex() (common.Address, error) {
	return _CrocQuery.Contract.Dex(&_CrocQuery.CallOpts)
}

// QueryAmbientPosition is a free data retrieval call binding the contract method 0x7f44601a.
//
// Solidity: function queryAmbientPosition(address owner, address base, address quote, uint256 poolIdx) view returns(uint128 seeds, uint32 timestamp)
func (_CrocQuery *CrocQueryCaller) QueryAmbientPosition(opts *bind.CallOpts, owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int) (struct {
	Seeds     *big.Int
	Timestamp uint32
}, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryAmbientPosition", owner, base, quote, poolIdx)

	outstruct := new(struct {
		Seeds     *big.Int
		Timestamp uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seeds = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// QueryAmbientPosition is a free data retrieval call binding the contract method 0x7f44601a.
//
// Solidity: function queryAmbientPosition(address owner, address base, address quote, uint256 poolIdx) view returns(uint128 seeds, uint32 timestamp)
func (_CrocQuery *CrocQuerySession) QueryAmbientPosition(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int) (struct {
	Seeds     *big.Int
	Timestamp uint32
}, error) {
	return _CrocQuery.Contract.QueryAmbientPosition(&_CrocQuery.CallOpts, owner, base, quote, poolIdx)
}

// QueryAmbientPosition is a free data retrieval call binding the contract method 0x7f44601a.
//
// Solidity: function queryAmbientPosition(address owner, address base, address quote, uint256 poolIdx) view returns(uint128 seeds, uint32 timestamp)
func (_CrocQuery *CrocQueryCallerSession) QueryAmbientPosition(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int) (struct {
	Seeds     *big.Int
	Timestamp uint32
}, error) {
	return _CrocQuery.Contract.QueryAmbientPosition(&_CrocQuery.CallOpts, owner, base, quote, poolIdx)
}

// QueryAmbientTokens is a free data retrieval call binding the contract method 0xebca95de.
//
// Solidity: function queryAmbientTokens(address owner, address base, address quote, uint256 poolIdx) view returns(uint128 liq, uint128 baseQty, uint128 quoteQty)
func (_CrocQuery *CrocQueryCaller) QueryAmbientTokens(opts *bind.CallOpts, owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int) (struct {
	Liq      *big.Int
	BaseQty  *big.Int
	QuoteQty *big.Int
}, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryAmbientTokens", owner, base, quote, poolIdx)

	outstruct := new(struct {
		Liq      *big.Int
		BaseQty  *big.Int
		QuoteQty *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liq = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BaseQty = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.QuoteQty = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QueryAmbientTokens is a free data retrieval call binding the contract method 0xebca95de.
//
// Solidity: function queryAmbientTokens(address owner, address base, address quote, uint256 poolIdx) view returns(uint128 liq, uint128 baseQty, uint128 quoteQty)
func (_CrocQuery *CrocQuerySession) QueryAmbientTokens(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int) (struct {
	Liq      *big.Int
	BaseQty  *big.Int
	QuoteQty *big.Int
}, error) {
	return _CrocQuery.Contract.QueryAmbientTokens(&_CrocQuery.CallOpts, owner, base, quote, poolIdx)
}

// QueryAmbientTokens is a free data retrieval call binding the contract method 0xebca95de.
//
// Solidity: function queryAmbientTokens(address owner, address base, address quote, uint256 poolIdx) view returns(uint128 liq, uint128 baseQty, uint128 quoteQty)
func (_CrocQuery *CrocQueryCallerSession) QueryAmbientTokens(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int) (struct {
	Liq      *big.Int
	BaseQty  *big.Int
	QuoteQty *big.Int
}, error) {
	return _CrocQuery.Contract.QueryAmbientTokens(&_CrocQuery.CallOpts, owner, base, quote, poolIdx)
}

// QueryConcRewards is a free data retrieval call binding the contract method 0x3dccd7d7.
//
// Solidity: function queryConcRewards(address owner, address base, address quote, uint256 poolIdx, int24 lowerTick, int24 upperTick) view returns(uint128 liqRewards, uint128 baseRewards, uint128 quoteRewards)
func (_CrocQuery *CrocQueryCaller) QueryConcRewards(opts *bind.CallOpts, owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, lowerTick *big.Int, upperTick *big.Int) (struct {
	LiqRewards   *big.Int
	BaseRewards  *big.Int
	QuoteRewards *big.Int
}, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryConcRewards", owner, base, quote, poolIdx, lowerTick, upperTick)

	outstruct := new(struct {
		LiqRewards   *big.Int
		BaseRewards  *big.Int
		QuoteRewards *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiqRewards = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BaseRewards = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.QuoteRewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QueryConcRewards is a free data retrieval call binding the contract method 0x3dccd7d7.
//
// Solidity: function queryConcRewards(address owner, address base, address quote, uint256 poolIdx, int24 lowerTick, int24 upperTick) view returns(uint128 liqRewards, uint128 baseRewards, uint128 quoteRewards)
func (_CrocQuery *CrocQuerySession) QueryConcRewards(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, lowerTick *big.Int, upperTick *big.Int) (struct {
	LiqRewards   *big.Int
	BaseRewards  *big.Int
	QuoteRewards *big.Int
}, error) {
	return _CrocQuery.Contract.QueryConcRewards(&_CrocQuery.CallOpts, owner, base, quote, poolIdx, lowerTick, upperTick)
}

// QueryConcRewards is a free data retrieval call binding the contract method 0x3dccd7d7.
//
// Solidity: function queryConcRewards(address owner, address base, address quote, uint256 poolIdx, int24 lowerTick, int24 upperTick) view returns(uint128 liqRewards, uint128 baseRewards, uint128 quoteRewards)
func (_CrocQuery *CrocQueryCallerSession) QueryConcRewards(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, lowerTick *big.Int, upperTick *big.Int) (struct {
	LiqRewards   *big.Int
	BaseRewards  *big.Int
	QuoteRewards *big.Int
}, error) {
	return _CrocQuery.Contract.QueryConcRewards(&_CrocQuery.CallOpts, owner, base, quote, poolIdx, lowerTick, upperTick)
}

// QueryCurve is a free data retrieval call binding the contract method 0x8e56c1c1.
//
// Solidity: function queryCurve(address base, address quote, uint256 poolIdx) view returns((uint128,uint128,uint128,uint64,uint64) curve)
func (_CrocQuery *CrocQueryCaller) QueryCurve(opts *bind.CallOpts, base common.Address, quote common.Address, poolIdx *big.Int) (CurveMathCurveState, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryCurve", base, quote, poolIdx)

	if err != nil {
		return *new(CurveMathCurveState), err
	}

	out0 := *abi.ConvertType(out[0], new(CurveMathCurveState)).(*CurveMathCurveState)

	return out0, err

}

// QueryCurve is a free data retrieval call binding the contract method 0x8e56c1c1.
//
// Solidity: function queryCurve(address base, address quote, uint256 poolIdx) view returns((uint128,uint128,uint128,uint64,uint64) curve)
func (_CrocQuery *CrocQuerySession) QueryCurve(base common.Address, quote common.Address, poolIdx *big.Int) (CurveMathCurveState, error) {
	return _CrocQuery.Contract.QueryCurve(&_CrocQuery.CallOpts, base, quote, poolIdx)
}

// QueryCurve is a free data retrieval call binding the contract method 0x8e56c1c1.
//
// Solidity: function queryCurve(address base, address quote, uint256 poolIdx) view returns((uint128,uint128,uint128,uint64,uint64) curve)
func (_CrocQuery *CrocQueryCallerSession) QueryCurve(base common.Address, quote common.Address, poolIdx *big.Int) (CurveMathCurveState, error) {
	return _CrocQuery.Contract.QueryCurve(&_CrocQuery.CallOpts, base, quote, poolIdx)
}

// QueryCurveTick is a free data retrieval call binding the contract method 0xdc91a6ad.
//
// Solidity: function queryCurveTick(address base, address quote, uint256 poolIdx) view returns(int24)
func (_CrocQuery *CrocQueryCaller) QueryCurveTick(opts *bind.CallOpts, base common.Address, quote common.Address, poolIdx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryCurveTick", base, quote, poolIdx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryCurveTick is a free data retrieval call binding the contract method 0xdc91a6ad.
//
// Solidity: function queryCurveTick(address base, address quote, uint256 poolIdx) view returns(int24)
func (_CrocQuery *CrocQuerySession) QueryCurveTick(base common.Address, quote common.Address, poolIdx *big.Int) (*big.Int, error) {
	return _CrocQuery.Contract.QueryCurveTick(&_CrocQuery.CallOpts, base, quote, poolIdx)
}

// QueryCurveTick is a free data retrieval call binding the contract method 0xdc91a6ad.
//
// Solidity: function queryCurveTick(address base, address quote, uint256 poolIdx) view returns(int24)
func (_CrocQuery *CrocQueryCallerSession) QueryCurveTick(base common.Address, quote common.Address, poolIdx *big.Int) (*big.Int, error) {
	return _CrocQuery.Contract.QueryCurveTick(&_CrocQuery.CallOpts, base, quote, poolIdx)
}

// QueryKnockoutMerkle is a free data retrieval call binding the contract method 0xab0a9898.
//
// Solidity: function queryKnockoutMerkle(address base, address quote, uint256 poolIdx, bool isBid, int24 tick) view returns(uint160 root, uint32 pivot, uint64 fee)
func (_CrocQuery *CrocQueryCaller) QueryKnockoutMerkle(opts *bind.CallOpts, base common.Address, quote common.Address, poolIdx *big.Int, isBid bool, tick *big.Int) (struct {
	Root  *big.Int
	Pivot uint32
	Fee   uint64
}, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryKnockoutMerkle", base, quote, poolIdx, isBid, tick)

	outstruct := new(struct {
		Root  *big.Int
		Pivot uint32
		Fee   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pivot = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Fee = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// QueryKnockoutMerkle is a free data retrieval call binding the contract method 0xab0a9898.
//
// Solidity: function queryKnockoutMerkle(address base, address quote, uint256 poolIdx, bool isBid, int24 tick) view returns(uint160 root, uint32 pivot, uint64 fee)
func (_CrocQuery *CrocQuerySession) QueryKnockoutMerkle(base common.Address, quote common.Address, poolIdx *big.Int, isBid bool, tick *big.Int) (struct {
	Root  *big.Int
	Pivot uint32
	Fee   uint64
}, error) {
	return _CrocQuery.Contract.QueryKnockoutMerkle(&_CrocQuery.CallOpts, base, quote, poolIdx, isBid, tick)
}

// QueryKnockoutMerkle is a free data retrieval call binding the contract method 0xab0a9898.
//
// Solidity: function queryKnockoutMerkle(address base, address quote, uint256 poolIdx, bool isBid, int24 tick) view returns(uint160 root, uint32 pivot, uint64 fee)
func (_CrocQuery *CrocQueryCallerSession) QueryKnockoutMerkle(base common.Address, quote common.Address, poolIdx *big.Int, isBid bool, tick *big.Int) (struct {
	Root  *big.Int
	Pivot uint32
	Fee   uint64
}, error) {
	return _CrocQuery.Contract.QueryKnockoutMerkle(&_CrocQuery.CallOpts, base, quote, poolIdx, isBid, tick)
}

// QueryKnockoutPivot is a free data retrieval call binding the contract method 0x10fc74f4.
//
// Solidity: function queryKnockoutPivot(address base, address quote, uint256 poolIdx, bool isBid, int24 tick) view returns(uint96 lots, uint32 pivot, uint16 range)
func (_CrocQuery *CrocQueryCaller) QueryKnockoutPivot(opts *bind.CallOpts, base common.Address, quote common.Address, poolIdx *big.Int, isBid bool, tick *big.Int) (struct {
	Lots  *big.Int
	Pivot uint32
	Range uint16
}, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryKnockoutPivot", base, quote, poolIdx, isBid, tick)

	outstruct := new(struct {
		Lots  *big.Int
		Pivot uint32
		Range uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lots = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Pivot = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Range = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// QueryKnockoutPivot is a free data retrieval call binding the contract method 0x10fc74f4.
//
// Solidity: function queryKnockoutPivot(address base, address quote, uint256 poolIdx, bool isBid, int24 tick) view returns(uint96 lots, uint32 pivot, uint16 range)
func (_CrocQuery *CrocQuerySession) QueryKnockoutPivot(base common.Address, quote common.Address, poolIdx *big.Int, isBid bool, tick *big.Int) (struct {
	Lots  *big.Int
	Pivot uint32
	Range uint16
}, error) {
	return _CrocQuery.Contract.QueryKnockoutPivot(&_CrocQuery.CallOpts, base, quote, poolIdx, isBid, tick)
}

// QueryKnockoutPivot is a free data retrieval call binding the contract method 0x10fc74f4.
//
// Solidity: function queryKnockoutPivot(address base, address quote, uint256 poolIdx, bool isBid, int24 tick) view returns(uint96 lots, uint32 pivot, uint16 range)
func (_CrocQuery *CrocQueryCallerSession) QueryKnockoutPivot(base common.Address, quote common.Address, poolIdx *big.Int, isBid bool, tick *big.Int) (struct {
	Lots  *big.Int
	Pivot uint32
	Range uint16
}, error) {
	return _CrocQuery.Contract.QueryKnockoutPivot(&_CrocQuery.CallOpts, base, quote, poolIdx, isBid, tick)
}

// QueryKnockoutPos is a free data retrieval call binding the contract method 0x391d582f.
//
// Solidity: function queryKnockoutPos(address owner, address base, address quote, uint256 poolIdx, uint32 pivot, bool isBid, int24 lowerTick, int24 upperTick) view returns(uint96 lots, uint64 mileage, uint32 timestamp)
func (_CrocQuery *CrocQueryCaller) QueryKnockoutPos(opts *bind.CallOpts, owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, pivot uint32, isBid bool, lowerTick *big.Int, upperTick *big.Int) (struct {
	Lots      *big.Int
	Mileage   uint64
	Timestamp uint32
}, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryKnockoutPos", owner, base, quote, poolIdx, pivot, isBid, lowerTick, upperTick)

	outstruct := new(struct {
		Lots      *big.Int
		Mileage   uint64
		Timestamp uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lots = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Mileage = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// QueryKnockoutPos is a free data retrieval call binding the contract method 0x391d582f.
//
// Solidity: function queryKnockoutPos(address owner, address base, address quote, uint256 poolIdx, uint32 pivot, bool isBid, int24 lowerTick, int24 upperTick) view returns(uint96 lots, uint64 mileage, uint32 timestamp)
func (_CrocQuery *CrocQuerySession) QueryKnockoutPos(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, pivot uint32, isBid bool, lowerTick *big.Int, upperTick *big.Int) (struct {
	Lots      *big.Int
	Mileage   uint64
	Timestamp uint32
}, error) {
	return _CrocQuery.Contract.QueryKnockoutPos(&_CrocQuery.CallOpts, owner, base, quote, poolIdx, pivot, isBid, lowerTick, upperTick)
}

// QueryKnockoutPos is a free data retrieval call binding the contract method 0x391d582f.
//
// Solidity: function queryKnockoutPos(address owner, address base, address quote, uint256 poolIdx, uint32 pivot, bool isBid, int24 lowerTick, int24 upperTick) view returns(uint96 lots, uint64 mileage, uint32 timestamp)
func (_CrocQuery *CrocQueryCallerSession) QueryKnockoutPos(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, pivot uint32, isBid bool, lowerTick *big.Int, upperTick *big.Int) (struct {
	Lots      *big.Int
	Mileage   uint64
	Timestamp uint32
}, error) {
	return _CrocQuery.Contract.QueryKnockoutPos(&_CrocQuery.CallOpts, owner, base, quote, poolIdx, pivot, isBid, lowerTick, upperTick)
}

// QueryKnockoutTokens is a free data retrieval call binding the contract method 0x93c33a71.
//
// Solidity: function queryKnockoutTokens(address owner, address base, address quote, uint256 poolIdx, uint32 pivot, bool isBid, int24 lowerTick, int24 upperTick) view returns(uint128 liq, uint128 baseQty, uint128 quoteQty, bool knockedOut)
func (_CrocQuery *CrocQueryCaller) QueryKnockoutTokens(opts *bind.CallOpts, owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, pivot uint32, isBid bool, lowerTick *big.Int, upperTick *big.Int) (struct {
	Liq        *big.Int
	BaseQty    *big.Int
	QuoteQty   *big.Int
	KnockedOut bool
}, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryKnockoutTokens", owner, base, quote, poolIdx, pivot, isBid, lowerTick, upperTick)

	outstruct := new(struct {
		Liq        *big.Int
		BaseQty    *big.Int
		QuoteQty   *big.Int
		KnockedOut bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liq = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BaseQty = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.QuoteQty = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.KnockedOut = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// QueryKnockoutTokens is a free data retrieval call binding the contract method 0x93c33a71.
//
// Solidity: function queryKnockoutTokens(address owner, address base, address quote, uint256 poolIdx, uint32 pivot, bool isBid, int24 lowerTick, int24 upperTick) view returns(uint128 liq, uint128 baseQty, uint128 quoteQty, bool knockedOut)
func (_CrocQuery *CrocQuerySession) QueryKnockoutTokens(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, pivot uint32, isBid bool, lowerTick *big.Int, upperTick *big.Int) (struct {
	Liq        *big.Int
	BaseQty    *big.Int
	QuoteQty   *big.Int
	KnockedOut bool
}, error) {
	return _CrocQuery.Contract.QueryKnockoutTokens(&_CrocQuery.CallOpts, owner, base, quote, poolIdx, pivot, isBid, lowerTick, upperTick)
}

// QueryKnockoutTokens is a free data retrieval call binding the contract method 0x93c33a71.
//
// Solidity: function queryKnockoutTokens(address owner, address base, address quote, uint256 poolIdx, uint32 pivot, bool isBid, int24 lowerTick, int24 upperTick) view returns(uint128 liq, uint128 baseQty, uint128 quoteQty, bool knockedOut)
func (_CrocQuery *CrocQueryCallerSession) QueryKnockoutTokens(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, pivot uint32, isBid bool, lowerTick *big.Int, upperTick *big.Int) (struct {
	Liq        *big.Int
	BaseQty    *big.Int
	QuoteQty   *big.Int
	KnockedOut bool
}, error) {
	return _CrocQuery.Contract.QueryKnockoutTokens(&_CrocQuery.CallOpts, owner, base, quote, poolIdx, pivot, isBid, lowerTick, upperTick)
}

// QueryLevel is a free data retrieval call binding the contract method 0x340bfa12.
//
// Solidity: function queryLevel(address base, address quote, uint256 poolIdx, int24 tick) view returns(uint96 bidLots, uint96 askLots, uint64 odometer)
func (_CrocQuery *CrocQueryCaller) QueryLevel(opts *bind.CallOpts, base common.Address, quote common.Address, poolIdx *big.Int, tick *big.Int) (struct {
	BidLots  *big.Int
	AskLots  *big.Int
	Odometer uint64
}, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryLevel", base, quote, poolIdx, tick)

	outstruct := new(struct {
		BidLots  *big.Int
		AskLots  *big.Int
		Odometer uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BidLots = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AskLots = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Odometer = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// QueryLevel is a free data retrieval call binding the contract method 0x340bfa12.
//
// Solidity: function queryLevel(address base, address quote, uint256 poolIdx, int24 tick) view returns(uint96 bidLots, uint96 askLots, uint64 odometer)
func (_CrocQuery *CrocQuerySession) QueryLevel(base common.Address, quote common.Address, poolIdx *big.Int, tick *big.Int) (struct {
	BidLots  *big.Int
	AskLots  *big.Int
	Odometer uint64
}, error) {
	return _CrocQuery.Contract.QueryLevel(&_CrocQuery.CallOpts, base, quote, poolIdx, tick)
}

// QueryLevel is a free data retrieval call binding the contract method 0x340bfa12.
//
// Solidity: function queryLevel(address base, address quote, uint256 poolIdx, int24 tick) view returns(uint96 bidLots, uint96 askLots, uint64 odometer)
func (_CrocQuery *CrocQueryCallerSession) QueryLevel(base common.Address, quote common.Address, poolIdx *big.Int, tick *big.Int) (struct {
	BidLots  *big.Int
	AskLots  *big.Int
	Odometer uint64
}, error) {
	return _CrocQuery.Contract.QueryLevel(&_CrocQuery.CallOpts, base, quote, poolIdx, tick)
}

// QueryLiquidity is a free data retrieval call binding the contract method 0x338adc67.
//
// Solidity: function queryLiquidity(address base, address quote, uint256 poolIdx) view returns(uint128)
func (_CrocQuery *CrocQueryCaller) QueryLiquidity(opts *bind.CallOpts, base common.Address, quote common.Address, poolIdx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryLiquidity", base, quote, poolIdx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryLiquidity is a free data retrieval call binding the contract method 0x338adc67.
//
// Solidity: function queryLiquidity(address base, address quote, uint256 poolIdx) view returns(uint128)
func (_CrocQuery *CrocQuerySession) QueryLiquidity(base common.Address, quote common.Address, poolIdx *big.Int) (*big.Int, error) {
	return _CrocQuery.Contract.QueryLiquidity(&_CrocQuery.CallOpts, base, quote, poolIdx)
}

// QueryLiquidity is a free data retrieval call binding the contract method 0x338adc67.
//
// Solidity: function queryLiquidity(address base, address quote, uint256 poolIdx) view returns(uint128)
func (_CrocQuery *CrocQueryCallerSession) QueryLiquidity(base common.Address, quote common.Address, poolIdx *big.Int) (*big.Int, error) {
	return _CrocQuery.Contract.QueryLiquidity(&_CrocQuery.CallOpts, base, quote, poolIdx)
}

// QueryPoolParams is a free data retrieval call binding the contract method 0xff163e14.
//
// Solidity: function queryPoolParams(address base, address quote, uint256 poolIdx) view returns((uint8,uint16,uint8,uint16,uint8,uint8,uint8) pool)
func (_CrocQuery *CrocQueryCaller) QueryPoolParams(opts *bind.CallOpts, base common.Address, quote common.Address, poolIdx *big.Int) (PoolSpecsPool, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryPoolParams", base, quote, poolIdx)

	if err != nil {
		return *new(PoolSpecsPool), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolSpecsPool)).(*PoolSpecsPool)

	return out0, err

}

// QueryPoolParams is a free data retrieval call binding the contract method 0xff163e14.
//
// Solidity: function queryPoolParams(address base, address quote, uint256 poolIdx) view returns((uint8,uint16,uint8,uint16,uint8,uint8,uint8) pool)
func (_CrocQuery *CrocQuerySession) QueryPoolParams(base common.Address, quote common.Address, poolIdx *big.Int) (PoolSpecsPool, error) {
	return _CrocQuery.Contract.QueryPoolParams(&_CrocQuery.CallOpts, base, quote, poolIdx)
}

// QueryPoolParams is a free data retrieval call binding the contract method 0xff163e14.
//
// Solidity: function queryPoolParams(address base, address quote, uint256 poolIdx) view returns((uint8,uint16,uint8,uint16,uint8,uint8,uint8) pool)
func (_CrocQuery *CrocQueryCallerSession) QueryPoolParams(base common.Address, quote common.Address, poolIdx *big.Int) (PoolSpecsPool, error) {
	return _CrocQuery.Contract.QueryPoolParams(&_CrocQuery.CallOpts, base, quote, poolIdx)
}

// QueryPoolTemplate is a free data retrieval call binding the contract method 0x2855bcf3.
//
// Solidity: function queryPoolTemplate(uint256 poolIdx) view returns((uint8,uint16,uint8,uint16,uint8,uint8,uint8) pool)
func (_CrocQuery *CrocQueryCaller) QueryPoolTemplate(opts *bind.CallOpts, poolIdx *big.Int) (PoolSpecsPool, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryPoolTemplate", poolIdx)

	if err != nil {
		return *new(PoolSpecsPool), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolSpecsPool)).(*PoolSpecsPool)

	return out0, err

}

// QueryPoolTemplate is a free data retrieval call binding the contract method 0x2855bcf3.
//
// Solidity: function queryPoolTemplate(uint256 poolIdx) view returns((uint8,uint16,uint8,uint16,uint8,uint8,uint8) pool)
func (_CrocQuery *CrocQuerySession) QueryPoolTemplate(poolIdx *big.Int) (PoolSpecsPool, error) {
	return _CrocQuery.Contract.QueryPoolTemplate(&_CrocQuery.CallOpts, poolIdx)
}

// QueryPoolTemplate is a free data retrieval call binding the contract method 0x2855bcf3.
//
// Solidity: function queryPoolTemplate(uint256 poolIdx) view returns((uint8,uint16,uint8,uint16,uint8,uint8,uint8) pool)
func (_CrocQuery *CrocQueryCallerSession) QueryPoolTemplate(poolIdx *big.Int) (PoolSpecsPool, error) {
	return _CrocQuery.Contract.QueryPoolTemplate(&_CrocQuery.CallOpts, poolIdx)
}

// QueryPrice is a free data retrieval call binding the contract method 0xf8c7efa7.
//
// Solidity: function queryPrice(address base, address quote, uint256 poolIdx) view returns(uint128)
func (_CrocQuery *CrocQueryCaller) QueryPrice(opts *bind.CallOpts, base common.Address, quote common.Address, poolIdx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryPrice", base, quote, poolIdx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryPrice is a free data retrieval call binding the contract method 0xf8c7efa7.
//
// Solidity: function queryPrice(address base, address quote, uint256 poolIdx) view returns(uint128)
func (_CrocQuery *CrocQuerySession) QueryPrice(base common.Address, quote common.Address, poolIdx *big.Int) (*big.Int, error) {
	return _CrocQuery.Contract.QueryPrice(&_CrocQuery.CallOpts, base, quote, poolIdx)
}

// QueryPrice is a free data retrieval call binding the contract method 0xf8c7efa7.
//
// Solidity: function queryPrice(address base, address quote, uint256 poolIdx) view returns(uint128)
func (_CrocQuery *CrocQueryCallerSession) QueryPrice(base common.Address, quote common.Address, poolIdx *big.Int) (*big.Int, error) {
	return _CrocQuery.Contract.QueryPrice(&_CrocQuery.CallOpts, base, quote, poolIdx)
}

// QueryProtocolAccum is a free data retrieval call binding the contract method 0xed8f58f4.
//
// Solidity: function queryProtocolAccum(address token) view returns(uint128)
func (_CrocQuery *CrocQueryCaller) QueryProtocolAccum(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryProtocolAccum", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryProtocolAccum is a free data retrieval call binding the contract method 0xed8f58f4.
//
// Solidity: function queryProtocolAccum(address token) view returns(uint128)
func (_CrocQuery *CrocQuerySession) QueryProtocolAccum(token common.Address) (*big.Int, error) {
	return _CrocQuery.Contract.QueryProtocolAccum(&_CrocQuery.CallOpts, token)
}

// QueryProtocolAccum is a free data retrieval call binding the contract method 0xed8f58f4.
//
// Solidity: function queryProtocolAccum(address token) view returns(uint128)
func (_CrocQuery *CrocQueryCallerSession) QueryProtocolAccum(token common.Address) (*big.Int, error) {
	return _CrocQuery.Contract.QueryProtocolAccum(&_CrocQuery.CallOpts, token)
}

// QueryRangePosition is a free data retrieval call binding the contract method 0x992236c5.
//
// Solidity: function queryRangePosition(address owner, address base, address quote, uint256 poolIdx, int24 lowerTick, int24 upperTick) view returns(uint128 liq, uint64 fee, uint32 timestamp, bool atomic)
func (_CrocQuery *CrocQueryCaller) QueryRangePosition(opts *bind.CallOpts, owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, lowerTick *big.Int, upperTick *big.Int) (struct {
	Liq       *big.Int
	Fee       uint64
	Timestamp uint32
	Atomic    bool
}, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryRangePosition", owner, base, quote, poolIdx, lowerTick, upperTick)

	outstruct := new(struct {
		Liq       *big.Int
		Fee       uint64
		Timestamp uint32
		Atomic    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liq = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Atomic = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// QueryRangePosition is a free data retrieval call binding the contract method 0x992236c5.
//
// Solidity: function queryRangePosition(address owner, address base, address quote, uint256 poolIdx, int24 lowerTick, int24 upperTick) view returns(uint128 liq, uint64 fee, uint32 timestamp, bool atomic)
func (_CrocQuery *CrocQuerySession) QueryRangePosition(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, lowerTick *big.Int, upperTick *big.Int) (struct {
	Liq       *big.Int
	Fee       uint64
	Timestamp uint32
	Atomic    bool
}, error) {
	return _CrocQuery.Contract.QueryRangePosition(&_CrocQuery.CallOpts, owner, base, quote, poolIdx, lowerTick, upperTick)
}

// QueryRangePosition is a free data retrieval call binding the contract method 0x992236c5.
//
// Solidity: function queryRangePosition(address owner, address base, address quote, uint256 poolIdx, int24 lowerTick, int24 upperTick) view returns(uint128 liq, uint64 fee, uint32 timestamp, bool atomic)
func (_CrocQuery *CrocQueryCallerSession) QueryRangePosition(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, lowerTick *big.Int, upperTick *big.Int) (struct {
	Liq       *big.Int
	Fee       uint64
	Timestamp uint32
	Atomic    bool
}, error) {
	return _CrocQuery.Contract.QueryRangePosition(&_CrocQuery.CallOpts, owner, base, quote, poolIdx, lowerTick, upperTick)
}

// QueryRangeTokens is a free data retrieval call binding the contract method 0xd7fd8d0f.
//
// Solidity: function queryRangeTokens(address owner, address base, address quote, uint256 poolIdx, int24 lowerTick, int24 upperTick) view returns(uint128 liq, uint128 baseQty, uint128 quoteQty)
func (_CrocQuery *CrocQueryCaller) QueryRangeTokens(opts *bind.CallOpts, owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, lowerTick *big.Int, upperTick *big.Int) (struct {
	Liq      *big.Int
	BaseQty  *big.Int
	QuoteQty *big.Int
}, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryRangeTokens", owner, base, quote, poolIdx, lowerTick, upperTick)

	outstruct := new(struct {
		Liq      *big.Int
		BaseQty  *big.Int
		QuoteQty *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Liq = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BaseQty = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.QuoteQty = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QueryRangeTokens is a free data retrieval call binding the contract method 0xd7fd8d0f.
//
// Solidity: function queryRangeTokens(address owner, address base, address quote, uint256 poolIdx, int24 lowerTick, int24 upperTick) view returns(uint128 liq, uint128 baseQty, uint128 quoteQty)
func (_CrocQuery *CrocQuerySession) QueryRangeTokens(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, lowerTick *big.Int, upperTick *big.Int) (struct {
	Liq      *big.Int
	BaseQty  *big.Int
	QuoteQty *big.Int
}, error) {
	return _CrocQuery.Contract.QueryRangeTokens(&_CrocQuery.CallOpts, owner, base, quote, poolIdx, lowerTick, upperTick)
}

// QueryRangeTokens is a free data retrieval call binding the contract method 0xd7fd8d0f.
//
// Solidity: function queryRangeTokens(address owner, address base, address quote, uint256 poolIdx, int24 lowerTick, int24 upperTick) view returns(uint128 liq, uint128 baseQty, uint128 quoteQty)
func (_CrocQuery *CrocQueryCallerSession) QueryRangeTokens(owner common.Address, base common.Address, quote common.Address, poolIdx *big.Int, lowerTick *big.Int, upperTick *big.Int) (struct {
	Liq      *big.Int
	BaseQty  *big.Int
	QuoteQty *big.Int
}, error) {
	return _CrocQuery.Contract.QueryRangeTokens(&_CrocQuery.CallOpts, owner, base, quote, poolIdx, lowerTick, upperTick)
}

// QuerySurplus is a free data retrieval call binding the contract method 0x56bf0f5b.
//
// Solidity: function querySurplus(address owner, address token) view returns(uint128 surplus)
func (_CrocQuery *CrocQueryCaller) QuerySurplus(opts *bind.CallOpts, owner common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "querySurplus", owner, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySurplus is a free data retrieval call binding the contract method 0x56bf0f5b.
//
// Solidity: function querySurplus(address owner, address token) view returns(uint128 surplus)
func (_CrocQuery *CrocQuerySession) QuerySurplus(owner common.Address, token common.Address) (*big.Int, error) {
	return _CrocQuery.Contract.QuerySurplus(&_CrocQuery.CallOpts, owner, token)
}

// QuerySurplus is a free data retrieval call binding the contract method 0x56bf0f5b.
//
// Solidity: function querySurplus(address owner, address token) view returns(uint128 surplus)
func (_CrocQuery *CrocQueryCallerSession) QuerySurplus(owner common.Address, token common.Address) (*big.Int, error) {
	return _CrocQuery.Contract.QuerySurplus(&_CrocQuery.CallOpts, owner, token)
}

// QueryVirtual is a free data retrieval call binding the contract method 0x6756e9b8.
//
// Solidity: function queryVirtual(address owner, address tracker, uint256 salt) view returns(uint128 surplus)
func (_CrocQuery *CrocQueryCaller) QueryVirtual(opts *bind.CallOpts, owner common.Address, tracker common.Address, salt *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CrocQuery.contract.Call(opts, &out, "queryVirtual", owner, tracker, salt)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryVirtual is a free data retrieval call binding the contract method 0x6756e9b8.
//
// Solidity: function queryVirtual(address owner, address tracker, uint256 salt) view returns(uint128 surplus)
func (_CrocQuery *CrocQuerySession) QueryVirtual(owner common.Address, tracker common.Address, salt *big.Int) (*big.Int, error) {
	return _CrocQuery.Contract.QueryVirtual(&_CrocQuery.CallOpts, owner, tracker, salt)
}

// QueryVirtual is a free data retrieval call binding the contract method 0x6756e9b8.
//
// Solidity: function queryVirtual(address owner, address tracker, uint256 salt) view returns(uint128 surplus)
func (_CrocQuery *CrocQueryCallerSession) QueryVirtual(owner common.Address, tracker common.Address, salt *big.Int) (*big.Int, error) {
	return _CrocQuery.Contract.QueryVirtual(&_CrocQuery.CallOpts, owner, tracker, salt)
}
