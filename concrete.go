package protocols

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/infrared-dao/protocols/fetchers"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/infrared-dao/protocols/multicall3"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

var _ BatchablePriceProvider = &ConcreteLPPriceProvider{}

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
func (c *ConcreteLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
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

// LPTokenPrice — legacy path; batchable via PriceReads + ComputePrice.
func (c *ConcreteLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	ts, err := c.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: c.block}
	ta, err := c.contract.TotalAssets(opts)
	if err != nil {
		c.logger.Error().Err(err).Msg("failed to fetch totalAssets")
		return "", err
	}
	price, err := c.computeLPPriceFromReads(ts, ta)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

func (c *ConcreteLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	abi, err := sc.ConcreteVaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("concrete: ABI: %w", err)
	}
	tsData, err := abi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("concrete: pack totalSupply: %w", err)
	}
	taData, err := abi.Pack("totalAssets")
	if err != nil {
		return nil, fmt.Errorf("concrete: pack totalAssets: %w", err)
	}
	return []multicall3.Call3{
		{Target: c.address, AllowFailure: false, CallData: tsData},
		{Target: c.address, AllowFailure: false, CallData: taData},
	}, nil
}

func (c *ConcreteLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("concrete: expected 2 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("concrete: sub-call %d reverted", i)
		}
	}
	abi, err := sc.ConcreteVaultMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("concrete: ABI: %w", err)
	}
	tsOut, err := abi.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("concrete: unpack totalSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("concrete: totalSupply type %T", tsOut[0])
	}
	taOut, err := abi.Methods["totalAssets"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("concrete: unpack totalAssets: %w", err)
	}
	totalAssets, ok := taOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("concrete: totalAssets type %T", taOut[0])
	}
	return c.computeLPPriceFromReads(totalSupply, totalAssets)
}

// computeLPPriceFromReads — shared math: tvl = totalAssets * assetPrice;
// price = tvl / totalSupply. Returns ErrPriceNotReadyYet on zero supply
// (concrete-specific: temporary state expected on freshly-deployed vaults).
func (c *ConcreteLPPriceProvider) computeLPPriceFromReads(totalSupply, totalAssets *big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		return decimal.Zero, ErrPriceNotReadyYet
	}
	assetPrice, err := c.getPrice(c.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(totalAssets, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	tsd := NormalizeAmount(totalSupply, c.config.LPTDecimals)
	pricePerToken := tvl.Div(tsd)
	c.logger.Debug().Str("pricePerToken", pricePerToken.String()).Msg("LP token price calculated successfully")
	return pricePerToken, nil
}

func (c *ConcreteLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := c.tvl(ctx)
	if err != nil {
		return "", err
	}

	c.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (c *ConcreteLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewConcreteVault(common.HexToAddress(address), client)
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

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for Concrete protocol
func (c *ConcreteLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
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
