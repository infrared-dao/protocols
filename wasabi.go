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

// WasabiLPPriceProvider defines the provider for Wasabi Token price and TVL.
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
	w := &WasabiLPPriceProvider{
		address:     address,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return w
}

// Initialize checks the configuration/data provided and instantiates the Wasabi smart contract.
func (w *WasabiLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	w.config = &WasabiConfig{}
	err = json.Unmarshal(w.configBytes, w.config)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	_, ok := w.priceMap[w.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", w.config.Token0)
		w.logger.Error().Msg(err.Error())
		return err
	}

	w.contract, err = sc.NewERC4626(w.address, client)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to instantiate Wasabi smart contract")
		return err
	}

	return nil
}

func (w *WasabiLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
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

	w.logger.Info().
		Str("totalValue", tvl.String()).
		Str("totalSupply", ts.String()).
		Str("pricePerToken", price.String()).
		Msg("LP token price calculated successfully")

	return price.StringFixed(roundingDecimals), nil
}

func (w *WasabiLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := w.tvl(ctx)
	if err != nil {
		return "", err
	}

	w.logger.Info().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (w *WasabiLPPriceProvider) GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error) {
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

	wc := &WasabiConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	addr, err := contract.Asset(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch asset address, %v", err)
		return nil, err
	}
	wc.Token0 = strings.ToLower(addr.Hex())

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

///// Helpers

// tvl fetches the TVL from the Wasabi smart contract.
func (w *WasabiLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
	wTokenAmount, err := w.getUnderlyingBalances(ctx)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to fetch wToken amount")
		return decimal.Zero, err
	}
	wTokenPrice, err := w.getPrice(w.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	wTokenAmountDecimal := NormalizeAmount(wTokenAmount, wTokenPrice.Decimals)
	tvl := wTokenAmountDecimal.Mul(wTokenPrice.Price)
	return tvl, nil
}

// getPrice fetches the price of the token from the price map.
func (w *WasabiLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := w.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		w.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (w *WasabiLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
	}
	totalSupply, err := w.contract.TotalSupply(opts)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to fetch total supply")
		return nil, err
	}
	return totalSupply, nil
}

func (w *WasabiLPPriceProvider) getUnderlyingBalances(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
	}

	amount0, err := w.contract.TotalAssets(opts)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to fetch total assets")
		return nil, err
	}

	return amount0, nil
}
