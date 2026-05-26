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

var _ Protocol = &EulerLPPriceProvider{}
var _ BatchablePriceProvider = &EulerLPPriceProvider{}

type EulerConfig struct {
	Asset       string `json:"asset"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// EulerLPPriceProvider defines the provider for Euler Token price and TVL.
type EulerLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *EulerConfig
	contract    *sc.EulerVault
}

// NewEulerLPPriceProvider creates a new instance of the EulerLPPriceProvider.
func NewEulerLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *EulerLPPriceProvider {
	e := &EulerLPPriceProvider{
		address:     address,
		block:       block,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return e
}

// Initialize checks the configuration/data provided and instantiates the Euler smart contract.
func (e *EulerLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
	var err error

	e.config = &EulerConfig{}
	err = json.Unmarshal(e.configBytes, e.config)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	_, ok := e.priceMap[e.config.Asset]
	if !ok {
		err = fmt.Errorf("no price data found for asset (%s)", e.config.Asset)
		e.logger.Error().Msg(err.Error())
		return err
	}

	e.contract, err = sc.NewEulerVault(e.address, client)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to instantiate Euler smart contract")
		return err
	}

	return nil
}

func (e *EulerLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	ts, err := e.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: e.block}
	ta, err := e.contract.TotalAssets(opts)
	if err != nil {
		return "", fmt.Errorf("euler: totalAssets: %w", err)
	}
	price, err := e.computeLPPriceFromReads(ts, ta)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

func (e *EulerLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	abi, err := sc.EulerVaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("euler: ABI: %w", err)
	}
	tsData, err := abi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("euler: pack totalSupply: %w", err)
	}
	taData, err := abi.Pack("totalAssets")
	if err != nil {
		return nil, fmt.Errorf("euler: pack totalAssets: %w", err)
	}
	return []multicall3.Call3{
		{Target: e.address, AllowFailure: true, CallData: tsData},
		{Target: e.address, AllowFailure: true, CallData: taData},
	}, nil
}

func (e *EulerLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("euler: expected 2 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("euler: sub-call %d reverted", i)
		}
	}
	abi, err := sc.EulerVaultMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("euler: ABI: %w", err)
	}
	tsOut, err := abi.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("euler: unpack totalSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("euler: totalSupply type %T", tsOut[0])
	}
	taOut, err := abi.Methods["totalAssets"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("euler: unpack totalAssets: %w", err)
	}
	totalAssets, ok := taOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("euler: totalAssets type %T", taOut[0])
	}
	return e.computeLPPriceFromReads(totalSupply, totalAssets)
}

func (e *EulerLPPriceProvider) computeLPPriceFromReads(totalSupply, totalAssets *big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		err := fmt.Errorf("total supply is zero")
		e.logger.Error().Err(err).Msg("invalid totalSupply")
		return decimal.Zero, err
	}
	assetPrice, err := e.getPrice(e.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(totalAssets, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	tsd := NormalizeAmount(totalSupply, e.config.LPTDecimals)
	pricePerToken := tvl.Div(tsd)
	e.logger.Debug().Str("pricePerToken", pricePerToken.String()).Msg("LP token price calculated successfully")
	return pricePerToken, nil
}

func (e *EulerLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := e.tvl(ctx)
	if err != nil {
		return "", err
	}

	e.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (e *EulerLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewEulerVault(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Euler smart contract, %v", err)
		return nil, err
	}

	ec := &EulerConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	assetAddress, err := contract.Asset(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch asset address, %v", err)
		return nil, err
	}
	ec.Asset = strings.ToLower(assetAddress.Hex())

	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch decimals, %v", err)
		return nil, err
	}
	ec.LPTDecimals = uint(decimals)

	body, err := json.Marshal(ec)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (e *EulerLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	e.block = block
	if prices != nil {
		e.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for Euler protocol
func (e *EulerLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

///// Helpers

// tvl fetches the TVL from the Euler smart contract.
func (e *EulerLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: e.block,
	}

	assetAmount, err := e.contract.TotalAssets(opts)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to fetch wToken amount")
		return decimal.Zero, err
	}
	assetPrice, err := e.getPrice(e.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(assetAmount, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	return tvl, nil
}

// getPrice fetches the price of the token from the price map.
func (e *EulerLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := e.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		e.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (e *EulerLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: e.block,
	}
	totalSupply, err := e.contract.TotalSupply(opts)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to fetch total supply")
		return nil, err
	}
	return totalSupply, nil
}
