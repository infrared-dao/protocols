package protocols

import (
	"context"
	"encoding/json"
	"errors"
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

var _ Protocol = &BendLPPriceProvider{}
var _ BatchablePriceProvider = &BendLPPriceProvider{}

type BendConfig struct {
	Asset       string `json:"token0"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// BendLPPriceProvider defines the provider for Bend Token price and TVL.
type BendLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *BendConfig
	contract    *sc.BendVault
}

// NewBendLPPriceProvider creates a new instance of the BendLPPriceProvider.
func NewBendLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *BendLPPriceProvider {
	w := &BendLPPriceProvider{
		address:     address,
		block:       block,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return w
}

// Initialize checks the configuration/data provided and instantiates the WeBera smart contract.
func (w *BendLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
	var err error

	w.config = &BendConfig{}
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

	w.contract, err = sc.NewBendVault(w.address, client)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to instantiate WeBera smart contract")
		return err
	}

	return nil
}

// LPTokenPrice — legacy path. Reads totalSupply + totalAssets sequentially.
// Batchable equivalent (PriceReads + ComputePrice) dispatches both in one
// Multicall3.aggregate3. Shared math via computeLPPriceFromReads.
func (w *BendLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	ts, err := w.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: w.block}
	totalAssets, err := w.contract.TotalAssets(opts)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to fetch totalAssets")
		return "", err
	}
	price, err := w.computeLPPriceFromReads(ts, totalAssets)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

// PriceReads returns the two calls (totalSupply + totalAssets) against the
// Bend vault contract.
func (w *BendLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	abi, err := sc.BendVaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("bend: get ABI: %w", err)
	}
	tsData, err := abi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("bend: pack totalSupply: %w", err)
	}
	taData, err := abi.Pack("totalAssets")
	if err != nil {
		return nil, fmt.Errorf("bend: pack totalAssets: %w", err)
	}
	return []multicall3.Call3{
		{Target: w.address, AllowFailure: false, CallData: tsData},
		{Target: w.address, AllowFailure: false, CallData: taData},
	}, nil
}

// ComputePrice decodes both responses and computes the LP price.
func (w *BendLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("bend: expected 2 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("bend: sub-call %d reverted in multicall", i)
		}
	}
	abi, err := sc.BendVaultMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("bend: get ABI: %w", err)
	}
	tsOut, err := abi.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("bend: unpack totalSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("bend: totalSupply type %T", tsOut[0])
	}
	taOut, err := abi.Methods["totalAssets"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("bend: unpack totalAssets: %w", err)
	}
	totalAssets, ok := taOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("bend: totalAssets type %T", taOut[0])
	}
	return w.computeLPPriceFromReads(totalSupply, totalAssets)
}

// computeLPPriceFromReads — shared math: tvl = totalAssets × assetPrice;
// price = tvl / totalSupply.
func (w *BendLPPriceProvider) computeLPPriceFromReads(totalSupply, totalAssets *big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		err := errors.New("total supply is zero")
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

	w.logger.Debug().
		Str("totalValue", tvl.String()).
		Str("totalSupply", totalSupply.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")
	return pricePerToken, nil
}

func (w *BendLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := w.tvl(ctx)
	if err != nil {
		return "", err
	}

	w.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (w *BendLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewBendVault(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate WeBera smart contract, %v", err)
		return nil, err
	}

	wc := &BendConfig{}
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

func (w *BendLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	w.block = block
	if prices != nil {
		w.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for Bend protocol
func (w *BendLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

///// Helpers

// tvl fetches the TVL from the Bend smart contract.
func (w *BendLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
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
func (w *BendLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := w.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		w.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (w *BendLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
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
