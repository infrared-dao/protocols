package protocols

import (
	"context"
	"encoding/json"
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

var _ Protocol = &GoldilocksLPPriceProvider{}

type GoldilocksConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// GoldilocksLPPriceProvider defines the provider for Goldilocks oriBGT/oriBGT-OT LP token price and TVL.
type GoldilocksLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *GoldilocksConfig
	contract    *sc.SteerVault
}

// NewGoldilocksLPPriceProvider creates a new instance of the GoldilocksLPPriceProvider.
func NewGoldilocksLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *GoldilocksLPPriceProvider {
	g := &GoldilocksLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return g
}

// Initialize checks the configuration/data provided and instantiates the Steer Vault smart contract.
func (g *GoldilocksLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	g.config = &GoldilocksConfig{}
	err = json.Unmarshal(g.configBytes, g.config)
	if err != nil {
		g.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	_, ok := g.priceMap[g.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", g.config.Token0)
		g.logger.Error().Msg(err.Error())
		return err
	}
	_, ok = g.priceMap[g.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", g.config.Token1)
		g.logger.Error().Msg(err.Error())
		return err
	}

	g.contract, err = sc.NewSteerVault(g.address, client)
	if err != nil {
		g.logger.Error().Err(err).Msg("failed to instantiate Steer Vault smart contract")
		return err
	}

	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD.
func (g *GoldilocksLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	ts, err := g.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	if ts.Cmp(big.NewInt(0)) == 0 {
		err = fmt.Errorf("total supply is zero")
		g.logger.Error().Err(err).Msg("failed to fetch token supply")
		return "", err
	}

	tvl, err := g.tvl(ctx)
	if err != nil {
		return "", err
	}

	tsd := NormalizeAmount(ts, g.config.LPTDecimals)
	price := tvl.Div(tsd)

	g.logger.Debug().
		Str("totalValue", tvl.String()).
		Str("totalSupply", ts.String()).
		Str("pricePerToken", price.String()).
		Msg("LP token price calculated successfully")

	return price.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the protocol in USD.
func (g *GoldilocksLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := g.tvl(ctx)
	if err != nil {
		return "", err
	}

	g.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

// GetConfig fetches and returns the configuration for the Goldilocks protocol.
func (g *GoldilocksLPPriceProvider) GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewSteerVault(common.HexToAddress(address), ethClient)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Steer Vault smart contract, %v", err)
		return nil, err
	}

	gc := &GoldilocksConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for steer vault %s, %v", address, err)
		return nil, err
	}
	gc.Token0 = strings.ToLower(addr.Hex())

	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for steer vault %s, %v", address, err)
		return nil, err
	}
	gc.Token1 = strings.ToLower(addr.Hex())

	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	gc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(gc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (g *GoldilocksLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	g.block = block
	if prices != nil {
		g.priceMap = prices
	}
}

///// Helpers

func (g *GoldilocksLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
	var err error

	amount0, amount1, err := g.getTotalAmounts(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := g.getPrice(g.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	amount0Decimal := NormalizeAmount(amount0, price0.Decimals)

	price1, err := g.getPrice(g.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}
	amount1Decimal := NormalizeAmount(amount1, price1.Decimals)
	totalValue := amount0Decimal.Mul(price0.Price).Add(amount1Decimal.Mul(price1.Price))
	return totalValue, nil
}

// getPrice fetches the price of the token from the price map.
func (g *GoldilocksLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := g.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		g.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (g *GoldilocksLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: g.block,
	}
	ts, err := g.contract.TotalSupply(opts)
	if err != nil {
		g.logger.Error().Msgf("failed to obtain total supply for steer vault %s, %v", g.address.String(), err)
		return nil, fmt.Errorf("failed to get steer total supply, err: %w", err)
	}

	return ts, err
}

// getTotalAmounts fetches the underlying token balances.
func (g *GoldilocksLPPriceProvider) getTotalAmounts(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: g.block,
	}
	ubs, err := g.contract.GetTotalAmounts(opts)
	if err != nil {
		g.logger.Error().Msgf("failed to obtain underlying balances for steer vault %s, %v", g.address.String(), err)
		return nil, nil, fmt.Errorf("failed to get steer underlying balances, err: %w", err)
	}
	return ubs.Total0, ubs.Total1, err
}