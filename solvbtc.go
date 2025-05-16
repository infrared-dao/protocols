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

// This adapter is used for SolvBTC.BERA -- the only known solv protocol asset so far
// According to them this is a yield bearing token which should always be 1:1 with the underlying SolvBTC
// It isn't rebasing, they pay the yield directly to user wallets, we get SolvBTC price from oracles already
// Therefore, this adapter is literally just a pass through for the LP price

const solvBTC = "0x541FD749419CA806a8bc7da8ac23D346f2dF8B77"

var _ Protocol = &SolvLPPriceProvider{}

type SolvConfig struct {
	Asset       string `json:"asset"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// SolvLPPriceProvider defines the provider for Solv Token price and TVL.
type SolvLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *SolvConfig
	contract    *sc.SolvBTC
}

// NewSolvLPPriceProvider creates a new instance of the SolvLPPriceProvider.
func NewSolvLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *SolvLPPriceProvider {
	s := &SolvLPPriceProvider{
		address:     address,
		block:       block,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return s
}

// Initialize checks the configuration/data provided and instantiates the Solv smart contract.
func (s *SolvLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	s.config = &SolvConfig{}
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
		s.logger.Error().Err(err).Msg("failed to instantiate Solv smart contract")
		return err
	}

	return nil
}

func (s *SolvLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	price, err := s.getPrice(s.config.Asset)
	if err != nil {
		return "", err
	}

	s.logger.Debug().
		Str("pricePerToken", price.Price.String()).
		Msg("LP token price calculated successfully")

	return price.Price.StringFixed(roundingDecimals), nil
}

func (s *SolvLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := s.tvl(ctx)
	if err != nil {
		return "", err
	}

	s.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (s *SolvLPPriceProvider) GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewSolvBTC(common.HexToAddress(address), ethClient)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Solv smart contract, %v", err)
		return nil, err
	}

	sc := &SolvConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	sc.Asset = strings.ToLower(solvBTC)

	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch decimals, %v", err)
		return nil, err
	}
	sc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(sc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *SolvLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	s.block = block
	if prices != nil {
		s.priceMap = prices
	}
}

///// Helpers

// tvl fetches the TVL from the Solv smart contract.
func (s *SolvLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
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
func (s *SolvLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := s.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		s.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}
