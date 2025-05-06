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

var _ Protocol = &ConcreteLPPriceProvider{}

type ConcreteConfig struct {
	Asset       string `json:"asset"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// ConcreteLPPriceProvider defines the provider for Concrete Token price and TVL.
type ConcreteLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *ConcreteConfig
	contract    *sc.ConcreteVault
}

// NewConcreteLPPriceProvider creates a new instance of the ConcreteLPPriceProvider.
func NewConcreteLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *ConcreteLPPriceProvider {
	c := &ConcreteLPPriceProvider{
		address:     address,
		block:       block,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return c
}

// Initialize checks the configuration/data provided and instantiates the Concrete Vault smart contract.
func (c *ConcreteLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	c.config = &ConcreteConfig{}
	err = json.Unmarshal(c.configBytes, c.config)
	if err != nil {
		c.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	_, ok := c.priceMap[c.config.Asset]
	if !ok {
		err = fmt.Errorf("no price data found for asset (%s)", c.config.Asset)
		c.logger.Error().Msg(err.Error())
		return err
	}

	c.contract, err = sc.NewConcreteVault(c.address, client)
	if err != nil {
		c.logger.Error().Err(err).Msg("failed to instantiate Concrete Vault smart contract")
		return err
	}

	return nil
}

func (c *ConcreteLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	ts, err := c.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	if ts.Cmp(big.NewInt(0)) == 0 {
		// it's expected to temporarily have zero total supply
		return "", ErrPriceNotReadyYet
	}

	tvl, err := c.tvl(ctx)
	if err != nil {
		return "", err
	}

	tsd := NormalizeAmount(ts, c.config.LPTDecimals)
	price := tvl.Div(tsd)

	c.logger.Debug().
		Str("totalValue", tvl.String()).
		Str("totalSupply", ts.String()).
		Str("pricePerToken", price.String()).
		Msg("LP token price calculated successfully")

	return price.StringFixed(roundingDecimals), nil
}

func (c *ConcreteLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := c.tvl(ctx)
	if err != nil {
		return "", err
	}

	c.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (c *ConcreteLPPriceProvider) GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewConcreteVault(common.HexToAddress(address), ethClient)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Concrete smart contract, %v", err)
		return nil, err
	}

	cc := &ConcreteConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	assetAddress, err := contract.Asset(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch asset address, %v", err)
		return nil, err
	}
	cc.Asset = strings.ToLower(assetAddress.Hex())

	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch decimals, %v", err)
		return nil, err
	}
	cc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(cc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *ConcreteLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	c.block = block
	if prices != nil {
		c.priceMap = prices
	}
}

///// Helpers

// tvl fetches the TVL from the Concrete smart contract.
func (c *ConcreteLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: c.block,
	}

	assetAmount, err := c.contract.TotalAssets(opts)
	if err != nil {
		c.logger.Error().Err(err).Msg("failed to fetch total assets in vault")
		return decimal.Zero, err
	}
	assetPrice, err := c.getPrice(c.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(assetAmount, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	return tvl, nil
}

// getPrice fetches the price of the token from the price map.
func (c *ConcreteLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := c.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		c.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (c *ConcreteLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: c.block,
	}
	totalSupply, err := c.contract.TotalSupply(opts)
	if err != nil {
		c.logger.Error().Err(err).Msg("failed to fetch total supply")
		return nil, err
	}
	return totalSupply, nil
}
