package protocols

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

const (
	V2PoolByteCodeSHA256    = "bffda5c9e111fa411890cc93d58cb5c58a14f97bcbfa642efe300c766c4397f6"
	V3IslandByteCodeSHA256  = "f1b68b8044f372c1831075be05a63c1757ef9c64cabd5dfbf9a749a93f25b1da"
	CharmPoolByteCodeSHA256 = "b8de812c5ad434d1d101b9dfb46de284637741a9b36b25e856fe1eca8a1388e0"
)

// enforce interface adherence
var _ Protocol = &KodiakLPPriceProvider{}
var _ KodiakContract = &KodiakV3Island{}
var _ KodiakContract = &KodiakV2Pool{}
var _ KodiakContract = &KodiakCharmPool{}

type Balances struct {
	Amount0 *big.Int
	Amount1 *big.Int
}

// Wrapper types which implement a common KodiakContract interface
type KodiakV3Island struct {
	*sc.KodiakIsland
}
type KodiakV2Pool struct {
	*sc.UniswapV2
}
type KodiakCharmPool struct {
	*sc.AlphaProVault
}

type KodiakContract interface {
	Decimals(opts *bind.CallOpts) (uint8, error)
	Token0(opts *bind.CallOpts) (common.Address, error)
	Token1(opts *bind.CallOpts) (common.Address, error)
	TotalSupply(opts *bind.CallOpts) (*big.Int, error)
	GetBalances(opts *bind.CallOpts) (Balances, error)
}

// Unify different implementations for getting internal balance amounts

func (v3 *KodiakV3Island) GetBalances(opts *bind.CallOpts) (Balances, error) {
	balances, err := v3.GetUnderlyingBalances(opts)
	if err != nil {
		return Balances{}, err
	}
	return Balances{
		Amount0: balances.Amount0Current,
		Amount1: balances.Amount1Current,
	}, nil
}

func (v2 *KodiakV2Pool) GetBalances(opts *bind.CallOpts) (Balances, error) {
	balances, err := v2.GetReserves(opts)
	if err != nil {
		return Balances{}, err
	}
	return Balances{
		Amount0: balances.Reserve0,
		Amount1: balances.Reserve1,
	}, nil
}

func (kc *KodiakCharmPool) GetBalances(opts *bind.CallOpts) (Balances, error) {
	balances, err := kc.GetTotalAmounts0(opts)
	if err != nil {
		return Balances{}, err
	}
	return Balances{
		Amount0: balances.Total0,
		Amount1: balances.Total1,
	}, nil
}

// Define core types for the Kodiak Adapter

type KodiakConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// KodiakLPPriceProvider defines the provider for Kodiak LP price and TVL.
type KodiakLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *KodiakConfig
	contract    KodiakContract
}

// NewKodiakLPPriceProvider creates a new instance of the KodiakLPPriceProvider.
func NewKodiakLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *KodiakLPPriceProvider {
	k := &KodiakLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return k
}

// Initialize checks the configuration/data provided and instantiates the KodiakV1 smart contract.
func (k *KodiakLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	k.config = &KodiakConfig{}
	err = json.Unmarshal(k.configBytes, k.config)
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}
	_, ok := k.priceMap[k.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", k.config.Token0)
		k.logger.Error().Msg(err.Error())
		return err
	}
	_, ok = k.priceMap[k.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", k.config.Token1)
		k.logger.Error().Msg(err.Error())
		return err
	}

	// initialize Kodiak contract of correct type
	k.contract, err = newKodiakContract(ctx, client, k.address)
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to instantiate Kodiak smart contract")
		return err
	}

	return nil
}

// LPTokenPrice returns the current price of LP token in USD
func (k *KodiakLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// Fetch total supply
	totalSupply, err := k.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		k.logger.Error().Err(err).Msg("Invalid totalSupply")
		return "", err
	}

	totalValue, err := k.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, k.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	k.logger.Debug().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the pool/island as USD
func (k *KodiakLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := k.totalValue(ctx)
	if err != nil {
		return "", err
	}

	k.logger.Debug().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

func (k *KodiakLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	// initialize Kodiak contract of correct type
	contract, err := newKodiakContract(ctx, client, common.HexToAddress(address))
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to instantiate Kodiak smart contract")
		return nil, err
	}

	kc := KodiakConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// token0
	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for kodiak vault %s, %v", address, err)
		return nil, err
	}
	kc.Token0 = strings.ToLower(addr.Hex())

	// token1
	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for kodiak vault %s, %v", address, err)
		return nil, err
	}
	kc.Token1 = strings.ToLower(addr.Hex())

	// decimals
	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	kc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(kc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (k *KodiakLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	k.block = block
	if prices != nil {
		k.priceMap = prices
	}
}

// Internal Helper methods not able to be called except in this file

func (k *KodiakLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch internal balances
	amount0, amount1, err := k.getBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := k.getPrice(k.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	amount0Decimal := NormalizeAmount(amount0, price0.Decimals)

	price1, err := k.getPrice(k.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}
	amount1Decimal := NormalizeAmount(amount1, price1.Decimals)
	totalValue := amount0Decimal.Mul(price0.Price).Add(amount1Decimal.Mul(price1.Price))
	return totalValue, nil
}

func (k *KodiakLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := k.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		k.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (k *KodiakLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: k.block,
	}

	ts, err := k.contract.TotalSupply(opts)
	if err != nil {
		k.logger.Error().Msgf("failed to obtain total supply for kodiak contract %s, %v",
			k.address.String(), err)
		return nil, fmt.Errorf("failed to get kodiak contract total supply, err: %w", err)
	}

	return ts, err
}

// getUnderlyingBalances fetches the underlying token balances.
func (k *KodiakLPPriceProvider) getBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: k.block,
	}

	bs, err := k.contract.GetBalances(opts)
	if err != nil {
		k.logger.Error().Msgf("failed to obtain balances for kodiak contract %s, %v",
			k.address.String(), err)
		return nil, nil, fmt.Errorf("failed to get kodiak contract balances, err: %w", err)
	}
	return bs.Amount0, bs.Amount1, err
}

// Check type of contract and create new connection
func newKodiakContract(
	ctx context.Context,
	client *ethclient.Client,
	address common.Address,
) (KodiakContract, error) {
	contractHash := checkContractCodeHash(ctx, client, address)

	switch contractHash {
	case V2PoolByteCodeSHA256:
		// initialize contract as a Kodiak V2 Pool
		pool, err := sc.NewUniswapV2(address, client)
		if err != nil {
			return nil, err
		}
		return &KodiakV2Pool{pool}, nil
	case V3IslandByteCodeSHA256:
		// initialize contract as a Kodiak V3 Island
		island, err := sc.NewKodiakIsland(address, client)
		if err != nil {
			return nil, err
		}
		return &KodiakV3Island{island}, nil
	case CharmPoolByteCodeSHA256:
		// initialize contract as a Kodiak x Charm Pool (concentrated liquidity)
		pool, err := sc.NewAlphaProVault(address, client)
		if err != nil {
			return nil, err
		}
		return &KodiakCharmPool{pool}, nil
	default:
		// Error for yet unknown contract type being attempted
		return nil, errors.New("contract unrecognized: neither V2 Pool or V3 Island")
	}
}

// get contract byte code and SHA256 hash it for checks on the contract type against known values
func checkContractCodeHash(
	ctx context.Context,
	client *ethclient.Client,
	address common.Address,
) string {
	bytecode, err := client.CodeAt(ctx, address, nil)
	if err != nil {
		panic(err)
	}
	byteCodeString := hex.EncodeToString(bytecode)

	hash := sha256.New()
	hash.Write([]byte(byteCodeString))
	return hex.EncodeToString(hash.Sum(nil))
}
