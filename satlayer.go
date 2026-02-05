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
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

// This adapter is used for satSolvBTC.BERA -- a 1:1 receipt token of SolvBTC.BERA
// Similar to Infrared's existing SolvBTC.BERA adapter
// satSolvBTC.BERA which should always be 1:1 with the underlying SolvBTC.BERA, and in turn, 1:1 with SolvBTC
// It isn't rebasing, they pay the yield directly to user wallets, we get SolvBTC price from oracles already
// Therefore, this adapter is literally just a pass through for the LP price

var _ Protocol = &SatLayerLPPriceProvider{}

type SatLayerConfig struct {
	Asset       string `json:"asset"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// SatLayerLPPriceProvider defines the provider for Token price and TVL.
type SatLayerLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *SatLayerConfig
	contract    *sc.SolvBTC
}

// NewSatLayerLPPriceProvider creates a new instance of the SatLayerLPPriceProvider.
func NewSatLayerLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *SatLayerLPPriceProvider {
	s := &SatLayerLPPriceProvider{
		address:     address,
		block:       block,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return s
}

// Initialize checks the configuration/data provided and instantiates the SatLayer smart contract.
func (s *SatLayerLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, httpClient fetchers.HttpClient) error {
	var err error

	s.config = &SatLayerConfig{}
	if len(s.configBytes) == 0 {
		err = fmt.Errorf("no configuration data provided for SatLayerLPPriceProvider")
		s.logger.Error().Msg(err.Error())
		return err
	}

	err = json.Unmarshal(s.configBytes, s.config)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	_, ok := s.priceMap[s.config.Asset]
	if !ok {
		err = fmt.Errorf("no price data found for asset (%s)", s.config.Asset)
		s.logger.Error().Msg(err.Error())
		return err
	}

	s.contract, err = sc.NewSolvBTC(s.address, client)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to instantiate SatLayer smart contract")
		return err
	}

	return nil
}

func (s *SatLayerLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	price, err := s.getPrice(s.config.Asset)
	if err != nil {
		return "", err
	}

	s.logger.Debug().
		Str("pricePerToken", price.Price.String()).
		Msg("LP token price calculated successfully")

	return price.Price.StringFixed(roundingDecimals), nil
}

func (s *SatLayerLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := s.tvl(ctx)
	if err != nil {
		return "", err
	}

	s.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (s *SatLayerLPPriceProvider) GetConfig(ctx context.Context, _ string, client bind.ContractBackend) ([]byte, error) {
	// Address is hardcoded to solvBTC because price feed of satSolvBTC.BERA is always 1:1 with SolvBTC
	var err error

	contract, err := sc.NewSolvBTC(common.HexToAddress(solvBTC), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Solv smart contract, %v", err)
		return nil, err
	}

	slc := &SatLayerConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	slc.Asset = strings.ToLower(solvBTC)

	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch decimals, %v", err)
		return nil, err
	}
	slc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(slc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *SatLayerLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	s.block = block
	if prices != nil {
		s.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for SatLayer protocol
func (s *SatLayerLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

///// Helpers

// tvl fetches the TVL from the SatLayer smart contract.
func (s *SatLayerLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: s.block,
	}

	assetAmount, err := s.contract.TotalSupply(opts)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to fetch total supply")
		return decimal.Zero, err
	}
	assetPrice, err := s.getPrice(s.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(assetAmount, s.config.LPTDecimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	return tvl, nil
}

// getPrice fetches the price of the token from the price map.
func (s *SatLayerLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := s.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		s.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}
