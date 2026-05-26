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

var _ Protocol = &TermMaxVaultPriceProvider{}
var _ BatchablePriceProvider = &TermMaxVaultPriceProvider{}

type TermMaxVaultConfig struct {
	Asset         string `json:"asset"`
	ShareDecimals uint   `json:"share_decimals"`
}

// TermMaxVaultPriceProvider defines the provider for TermMax Vault price and TVL.
type TermMaxVaultPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *TermMaxVaultConfig
	contract    *sc.ERC4626
}

// NewTermMaxVaultPriceProvider creates a new instance of the TermMaxVaultPriceProvider.
func NewTermMaxVaultPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *TermMaxVaultPriceProvider {
	t := &TermMaxVaultPriceProvider{
		address:     address,
		block:       block,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return t
}

// Initialize checks the configuration/data provided and instantiates the ERC4626 smart contract.
func (t *TermMaxVaultPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, httpClient fetchers.HttpClient) error {
	var err error

	t.config = &TermMaxVaultConfig{}
	err = json.Unmarshal(t.configBytes, t.config)
	if err != nil {
		t.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	_, ok := t.priceMap[t.config.Asset]
	if !ok {
		err = fmt.Errorf("no price data found for asset (%s)", t.config.Asset)
		t.logger.Error().Msg(err.Error())
		return err
	}

	t.contract, err = sc.NewERC4626(t.address, client)
	if err != nil {
		t.logger.Error().Err(err).Msg("failed to instantiate TermMax vault contract")
		return err
	}

	return nil
}

func (t *TermMaxVaultPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	ts, err := t.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: t.block}
	ta, err := t.contract.TotalAssets(opts)
	if err != nil {
		return "", fmt.Errorf("termmax: totalAssets: %w", err)
	}
	price, err := t.computeLPPriceFromReads(ts, ta)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

func (t *TermMaxVaultPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	abi, err := sc.ERC4626MetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("termmax: ABI: %w", err)
	}
	tsData, err := abi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("termmax: pack totalSupply: %w", err)
	}
	taData, err := abi.Pack("totalAssets")
	if err != nil {
		return nil, fmt.Errorf("termmax: pack totalAssets: %w", err)
	}
	return []multicall3.Call3{
		{Target: t.address, AllowFailure: true, CallData: tsData},
		{Target: t.address, AllowFailure: true, CallData: taData},
	}, nil
}

func (t *TermMaxVaultPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("termmax: expected 2 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("termmax: sub-call %d reverted", i)
		}
	}
	abi, err := sc.ERC4626MetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("termmax: ABI: %w", err)
	}
	tsOut, err := abi.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("termmax: unpack totalSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("termmax: totalSupply type %T", tsOut[0])
	}
	taOut, err := abi.Methods["totalAssets"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("termmax: unpack totalAssets: %w", err)
	}
	totalAssets, ok := taOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("termmax: totalAssets type %T", taOut[0])
	}
	return t.computeLPPriceFromReads(totalSupply, totalAssets)
}

func (t *TermMaxVaultPriceProvider) computeLPPriceFromReads(totalSupply, totalAssets *big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		err := fmt.Errorf("total supply is zero")
		t.logger.Error().Err(err).Msg("total supply is zero")
		return decimal.Zero, err
	}
	assetPrice, err := t.getPrice(t.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(totalAssets, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	tsd := NormalizeAmount(totalSupply, t.config.ShareDecimals)
	pricePerToken := tvl.Div(tsd)
	t.logger.Debug().Str("pricePerToken", pricePerToken.String()).Msg("TermMax vault share price calculated successfully")
	return pricePerToken, nil
}

func (t *TermMaxVaultPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := t.tvl(ctx)
	if err != nil {
		return "", err
	}

	t.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (t *TermMaxVaultPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewERC4626(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate TermMax vault contract, %v", err)
		return nil, err
	}

	tc := &TermMaxVaultConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	assetAddress, err := contract.Asset(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch asset address, %v", err)
		return nil, err
	}
	tc.Asset = strings.ToLower(assetAddress.Hex())

	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch decimals, %v", err)
		return nil, err
	}
	tc.ShareDecimals = uint(decimals)

	body, err := json.Marshal(tc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (t *TermMaxVaultPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	t.block = block
	if prices != nil {
		t.priceMap = prices
	}
}

///// Helpers

// tvl fetches the TVL from the TermMax vault.
func (t *TermMaxVaultPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: t.block,
	}

	totalAssets, err := t.contract.TotalAssets(opts)
	if err != nil {
		t.logger.Error().Err(err).Msg("failed to fetch total assets")
		return decimal.Zero, err
	}

	assetPrice, err := t.getPrice(t.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}

	assetAmountDecimal := NormalizeAmount(totalAssets, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)

	return tvl, nil
}

// getPrice fetches the price of the token from the price map.
func (t *TermMaxVaultPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := t.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		t.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of vault shares.
func (t *TermMaxVaultPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: t.block,
	}
	totalSupply, err := t.contract.TotalSupply(opts)
	if err != nil {
		t.logger.Error().Err(err).Msg("failed to fetch total supply")
		return nil, err
	}
	return totalSupply, nil
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for TermMax protocol
func (w *TermMaxVaultPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}
