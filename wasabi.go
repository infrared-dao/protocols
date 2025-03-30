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

type WasabiConfig struct {
	Token0      string `json:"token0"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// WasabiLPPriceProvider defines the provider for Wasabi dToken price and TVL.
type WasabiLPPriceProvider struct {
	address     common.Address
	logger      zerolog.Logger
	priceMap    map[string]Price
	configBytes []byte
	config      *WasabiConfig
	contract    *sc.ERC4626
}

// NewWasabiLPPriceProvider creates a new instance of the WasabiLPPriceProvider.
func NewWasabiLPPriceProvider(address common.Address, prices map[string]Price, logger zerolog.Logger, config []byte) *WasabiLPPriceProvider {
	d := &WasabiLPPriceProvider{
		address:     address,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return d
}

// Initialize checks the configuration/data provided and instantiates the Wasabi smart contract.
func (d *WasabiLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	d.config = &WasabiConfig{}
	err = json.Unmarshal(d.configBytes, d.config)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	_, ok := d.priceMap[d.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", d.config.Token0)
		d.logger.Error().Msg(err.Error())
		return err
	}

	d.contract, err = sc.NewERC4626(d.address, client)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to instantiate Wasabi smart contract")
		return err
	}

	return nil
}

func (d *WasabiLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	ts, err := d.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	if ts.Cmp(big.NewInt(0)) == 0 {
		err = fmt.Errorf("total supply is zero")
		d.logger.Error().Err(err).Msg("failed to fetch total supply")
		return "", err
	}

	tvl, err := d.tvl(ctx)
	if err != nil {
		return "", err
	}

	tsd := NormalizeAmount(ts, d.config.LPTDecimals)
	price := tvl.Div(tsd)

	d.logger.Info().
		Str("totalValue", tvl.String()).
		Str("totalSupply", ts.String()).
		Str("pricePerToken", price.String()).
		Msg("LP token price calculated successfully")

	return price.StringFixed(roundingDecimals), nil
}

func (d *WasabiLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := d.tvl(ctx)
	if err != nil {
		return "", err
	}

	d.logger.Info().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (d *WasabiLPPriceProvider) GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewERC4626(common.HexToAddress(address), ethClient)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Wasabi smart contract, %v", err)
		return nil, err
	}

	dc := &WasabiConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	addr, err := contract.Asset(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch asset address, %v", err)
		return nil, err
	}
	dc.Token0 = strings.ToLower(addr.Hex())

	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch decimals, %v", err)
		return nil, err
	}
	dc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(dc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

///// Helpers

// tvl fetches the TVL from the Wasabi smart contract.
func (d *WasabiLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
	dTokenAmount, err := d.getUnderlyingBalances(ctx)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to fetch dToken amount")
		return decimal.Zero, err
	}
	dTokenPrice, err := d.getPrice(d.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	dTokenAmountDecimal := NormalizeAmount(dTokenAmount, dTokenPrice.Decimals)
	tvl := dTokenAmountDecimal.Mul(dTokenPrice.Price)
	return tvl, nil
}

// getPrice fetches the price of the token from the price map.
func (d *WasabiLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := d.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		d.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (d *WasabiLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
	}
	totalSupply, err := d.contract.TotalSupply(opts)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to fetch total supply")
		return nil, err
	}
	return totalSupply, nil
}

func (d *WasabiLPPriceProvider) getUnderlyingBalances(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
	}

	amount0, err := d.contract.TotalAssets(opts)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to fetch total assets")
		return nil, err
	}

	return amount0, nil
}
