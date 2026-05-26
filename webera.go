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

var _ Protocol = &WeberaLPPriceProvider{}
var _ BatchablePriceProvider = &WeberaLPPriceProvider{}

type WeberaConfig struct {
	Asset       string `json:"token0"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// WeberaLPPriceProvider defines the provider for Webera Token price and TVL.
type WeberaLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *WeberaConfig
	contract    *sc.WeberaVault
}

// NewWeberaLPPriceProvider creates a new instance of the WeberaLPPriceProvider.
func NewWeberaLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *WeberaLPPriceProvider {
	w := &WeberaLPPriceProvider{
		address:     address,
		block:       block,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return w
}

// Initialize checks the configuration/data provided and instantiates the WeBera smart contract.
func (w *WeberaLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, httpClient fetchers.HttpClient) error {
	var err error

	w.config = &WeberaConfig{}
	err = json.Unmarshal(w.configBytes, w.config)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	_, ok := w.priceMap[w.config.Asset]
	if !ok {
		err = fmt.Errorf("no price data found for asset (%s)", w.config.Asset)
		w.logger.Error().Msg(err.Error())
		return err
	}

	w.contract, err = sc.NewWeberaVault(w.address, client)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to instantiate WeBera smart contract")
		return err
	}

	return nil
}

func (w *WeberaLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	ts, err := w.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: w.block}
	ta, err := w.contract.TotalAssets(opts)
	if err != nil {
		return "", fmt.Errorf("webera: totalAssets: %w", err)
	}
	price, err := w.computeLPPriceFromReads(ts, ta)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

func (w *WeberaLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	abi, err := sc.WeberaVaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("webera: ABI: %w", err)
	}
	tsData, err := abi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("webera: pack totalSupply: %w", err)
	}
	taData, err := abi.Pack("totalAssets")
	if err != nil {
		return nil, fmt.Errorf("webera: pack totalAssets: %w", err)
	}
	return []multicall3.Call3{
		{Target: w.address, AllowFailure: true, CallData: tsData},
		{Target: w.address, AllowFailure: true, CallData: taData},
	}, nil
}

func (w *WeberaLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("webera: expected 2 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("webera: sub-call %d reverted", i)
		}
	}
	abi, err := sc.WeberaVaultMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("webera: ABI: %w", err)
	}
	tsOut, err := abi.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("webera: unpack totalSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("webera: totalSupply type %T", tsOut[0])
	}
	taOut, err := abi.Methods["totalAssets"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("webera: unpack totalAssets: %w", err)
	}
	totalAssets, ok := taOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("webera: totalAssets type %T", taOut[0])
	}
	return w.computeLPPriceFromReads(totalSupply, totalAssets)
}

func (w *WeberaLPPriceProvider) computeLPPriceFromReads(totalSupply, totalAssets *big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		err := fmt.Errorf("total supply is zero")
		w.logger.Error().Err(err).Msg("invalid totalSupply")
		return decimal.Zero, err
	}
	assetPrice, err := w.getPrice(w.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(totalAssets, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	tsd := NormalizeAmount(totalSupply, w.config.LPTDecimals)
	pricePerToken := tvl.Div(tsd)
	w.logger.Debug().Str("pricePerToken", pricePerToken.String()).Msg("LP token price calculated successfully")
	return pricePerToken, nil
}

func (w *WeberaLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := w.tvl(ctx)
	if err != nil {
		return "", err
	}

	w.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (w *WeberaLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewWeberaVault(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate WeBera smart contract, %v", err)
		return nil, err
	}

	wc := &WeberaConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	assetAddress, err := contract.Asset(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch asset address, %v", err)
		return nil, err
	}
	wc.Asset = strings.ToLower(assetAddress.Hex())

	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch decimals, %v", err)
		return nil, err
	}
	wc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(wc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (w *WeberaLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	w.block = block
	if prices != nil {
		w.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for Webera protocol
func (w *WeberaLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

///// Helpers

// tvl fetches the TVL from the Webera smart contract.
func (w *WeberaLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: w.block,
	}

	assetAmount, err := w.contract.TotalAssets(opts)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to fetch wToken amount")
		return decimal.Zero, err
	}
	assetPrice, err := w.getPrice(w.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(assetAmount, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	return tvl, nil
}

// getPrice fetches the price of the token from the price map.
func (w *WeberaLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := w.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		w.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (w *WeberaLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: w.block,
	}
	totalSupply, err := w.contract.TotalSupply(opts)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to fetch total supply")
		return nil, err
	}
	return totalSupply, nil
}
