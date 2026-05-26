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

var _ Protocol = &D2LPPriceProvider{}
var _ BatchablePriceProvider = &D2LPPriceProvider{}

type D2Config struct {
	Asset       string `json:"asset"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// D2LPPriceProvider defines the provider for D2 strategy vault price and TVL.
type D2LPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *D2Config
	contract    *sc.D2Vault
}

// NewD2LPPriceProvider creates a new instance of the D2LPPriceProvider.
func NewD2LPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *D2LPPriceProvider {
	d2 := &D2LPPriceProvider{
		address:     address,
		block:       block,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return d2
}

// Initialize checks the configuration/data provided and instantiates the D2 vault smart contract.
func (d2 *D2LPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
	var err error

	d2.config = &D2Config{}
	err = json.Unmarshal(d2.configBytes, d2.config)
	if err != nil {
		d2.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	_, ok := d2.priceMap[d2.config.Asset]
	if !ok {
		err = fmt.Errorf("no price data found for asset (%s)", d2.config.Asset)
		d2.logger.Error().Msg(err.Error())
		return err
	}

	d2.contract, err = sc.NewD2Vault(d2.address, client)
	if err != nil {
		d2.logger.Error().Err(err).Msg("failed to instantiate D2 vault smart contract")
		return err
	}

	return nil
}

func (d2 *D2LPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	ts, err := d2.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: d2.block}
	ta, err := d2.contract.TotalAssets(opts)
	if err != nil {
		return "", fmt.Errorf("d2: totalAssets: %w", err)
	}
	price, err := d2.computeLPPriceFromReads(ts, ta)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

func (d2 *D2LPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	abi, err := sc.D2VaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("d2: ABI: %w", err)
	}
	tsData, err := abi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("d2: pack totalSupply: %w", err)
	}
	taData, err := abi.Pack("totalAssets")
	if err != nil {
		return nil, fmt.Errorf("d2: pack totalAssets: %w", err)
	}
	return []multicall3.Call3{
		{Target: d2.address, AllowFailure: true, CallData: tsData},
		{Target: d2.address, AllowFailure: true, CallData: taData},
	}, nil
}

func (d2 *D2LPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("d2: expected 2 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("d2: sub-call %d reverted", i)
		}
	}
	abi, err := sc.D2VaultMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("d2: ABI: %w", err)
	}
	tsOut, err := abi.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("d2: unpack totalSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("d2: totalSupply type %T", tsOut[0])
	}
	taOut, err := abi.Methods["totalAssets"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("d2: unpack totalAssets: %w", err)
	}
	totalAssets, ok := taOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("d2: totalAssets type %T", taOut[0])
	}
	return d2.computeLPPriceFromReads(totalSupply, totalAssets)
}

func (d2 *D2LPPriceProvider) computeLPPriceFromReads(totalSupply, totalAssets *big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		err := fmt.Errorf("total supply is zero")
		d2.logger.Error().Err(err).Msg("invalid totalSupply")
		return decimal.Zero, err
	}
	assetPrice, err := d2.getPrice(d2.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(totalAssets, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	tsd := NormalizeAmount(totalSupply, d2.config.LPTDecimals)
	pricePerToken := tvl.Div(tsd)
	d2.logger.Debug().Str("pricePerToken", pricePerToken.String()).Msg("LP token price calculated successfully")
	return pricePerToken, nil
}

func (d2 *D2LPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := d2.tvl(ctx)
	if err != nil {
		return "", err
	}

	d2.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (d2 *D2LPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewD2Vault(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate D2 vault smart contract, %v", err)
		return nil, err
	}

	d2c := &D2Config{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	assetAddress, err := contract.Asset(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch asset address, %v", err)
		return nil, err
	}
	d2c.Asset = strings.ToLower(assetAddress.Hex())

	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch decimals, %v", err)
		return nil, err
	}
	d2c.LPTDecimals = uint(decimals)

	body, err := json.Marshal(d2c)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (d2 *D2LPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	d2.block = block
	if prices != nil {
		d2.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for D2 protocol
func (d2 *D2LPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

///// Helpers

// tvl fetches the TVL from the D2 vault smart contract.
func (d2 *D2LPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: d2.block,
	}

	assetAmount, err := d2.contract.TotalAssets(opts)
	if err != nil {
		d2.logger.Error().Err(err).Msg("failed to fetch TotalAssets amount")
		return decimal.Zero, err
	}
	assetPrice, err := d2.getPrice(d2.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(assetAmount, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	return tvl, nil
}

// getPrice fetches the price of the token from the price map.
func (d2 *D2LPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := d2.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		d2.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (d2 *D2LPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: d2.block,
	}
	totalSupply, err := d2.contract.TotalSupply(opts)
	if err != nil {
		d2.logger.Error().Err(err).Msg("failed to fetch total supply")
		return nil, err
	}
	return totalSupply, nil
}
