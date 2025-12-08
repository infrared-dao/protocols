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

var _ Protocol = &TermMaxVaultPriceProvider{}

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
	contract    *sc.ERC4626Vault
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
func (t *TermMaxVaultPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
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

	t.contract, err = sc.NewERC4626Vault(t.address, client)
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

	if ts.Cmp(big.NewInt(0)) == 0 {
		err = fmt.Errorf("total supply is zero")
		t.logger.Error().Err(err).Msg("total supply is zero")
		return "", err
	}

	tvl, err := t.tvl(ctx)
	if err != nil {
		return "", err
	}

	tsd := NormalizeAmount(ts, t.config.ShareDecimals)
	price := tvl.Div(tsd)

	t.logger.Debug().
		Str("totalValue", tvl.String()).
		Str("totalSupply", ts.String()).
		Str("pricePerToken", price.String()).
		Msg("TermMax vault share price calculated successfully")

	return price.StringFixed(roundingDecimals), nil
}

func (t *TermMaxVaultPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := t.tvl(ctx)
	if err != nil {
		return "", err
	}

	t.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (t *TermMaxVaultPriceProvider) GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewERC4626Vault(common.HexToAddress(address), ethClient)
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
