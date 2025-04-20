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

var _ Protocol = &D2LPPriceProvider{}

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
func (d2 *D2LPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
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

	if ts.Cmp(big.NewInt(0)) == 0 {
		err = fmt.Errorf("total supply is zero")
		d2.logger.Error().Err(err).Msg("failed to fetch total supply")
		return "", err
	}

	tvl, err := d2.tvl(ctx)
	if err != nil {
		return "", err
	}

	tsd := NormalizeAmount(ts, d2.config.LPTDecimals)
	price := tvl.Div(tsd)

	d2.logger.Debug().
		Str("totalValue", tvl.String()).
		Str("totalSupply", ts.String()).
		Str("pricePerToken", price.String()).
		Msg("LP token price calculated successfully")

	return price.StringFixed(roundingDecimals), nil
}

func (d2 *D2LPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := d2.tvl(ctx)
	if err != nil {
		return "", err
	}

	d2.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (d2 *D2LPPriceProvider) GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewD2Vault(common.HexToAddress(address), ethClient)
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
