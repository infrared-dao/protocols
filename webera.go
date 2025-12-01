package protocols

import (
	"context"
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

var _ Protocol = &WeberaLPPriceProvider{}

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
func (w *WeberaLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
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

	if ts.Cmp(big.NewInt(0)) == 0 {
		err = fmt.Errorf("total supply is zero")
		w.logger.Error().Err(err).Msg("failed to fetch total supply")
		return "", err
	}

	tvl, err := w.tvl(ctx)
	if err != nil {
		return "", err
	}

	tsd := NormalizeAmount(ts, w.config.LPTDecimals)
	price := tvl.Div(tsd)

	w.logger.Debug().
		Str("totalValue", tvl.String()).
		Str("totalSupply", ts.String()).
		Str("pricePerToken", price.String()).
		Msg("LP token price calculated successfully")

	return price.StringFixed(roundingDecimals), nil
}

func (w *WeberaLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := w.tvl(ctx)
	if err != nil {
		return "", err
	}

	w.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (w *WeberaLPPriceProvider) GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewWeberaVault(common.HexToAddress(address), ethClient)
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
	return nil, errors.New("TVLBreakdown not yet implemented for Webera")
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
