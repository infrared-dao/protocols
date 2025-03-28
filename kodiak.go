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

type KodiakConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// KodiakLPPriceProvider defines the provider for Kodiak LP price and TVL.
type KodiakLPPriceProvider struct {
	address     common.Address
	logger      zerolog.Logger
	priceMap    map[string]Price
	configBytes []byte
	config      *KodiakConfig
	contract    *sc.KodiakV1
	contractV2  *sc.UniswapV2
}

// NewKodiakLPPriceProvider creates a new instance of the KodiakLPPriceProvider.
func NewKodiakLPPriceProvider(address common.Address, prices map[string]Price, logger zerolog.Logger, config []byte) *KodiakLPPriceProvider {
	k := &KodiakLPPriceProvider{
		address:     address,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return k
}

// Initialize checks the configuration/data provided and instantiates the KodiakV1 smart contract.
func (k *KodiakLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error
	opts := &bind.CallOpts{
		Context: ctx,
	}

	k.config = &KodiakConfig{}
	err = json.Unmarshal(k.configBytes, k.config)
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}
	_, ok := k.priceMap[k.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", k.config.Token0)
		k.logger.Error().Msg(err.Error())
		return err
	}
	_, ok = k.priceMap[k.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", k.config.Token1)
		k.logger.Error().Msg(err.Error())
		return err
	}

	k.contract, err = sc.NewKodiakV1(k.address, client)
	if err != nil {
		k.logger.Error().Err(err).Msg("failed to instantiate V3 Kodiak Island smart contract")
		return err
	}

	// Try to call the version() function... if error then not an Island V3 contract
	_, err = k.contract.Version(opts)
	if err != nil {
		// If not a V3 Island, initialize it as a Kodiak V2 Pool
		k.contractV2, err = sc.NewUniswapV2(k.address, client)
		if err != nil {
			k.logger.Error().Err(err).Msg("failed to instantiate V2 Kodiak Pool smart contract")
			return err
		}
	}

	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD cents (1 USD = 100 cents).
func (k *KodiakLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// Fetch total supply
	totalSupply, err := k.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		k.logger.Error().Err(err).Msg("Invalid totalSupply")
		return "", err
	}

	totalValue, err := k.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, k.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	k.logger.Info().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the protocol in USD cents (1 USD = 100 cents).
func (k *KodiakLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := k.totalValue(ctx)
	if err != nil {
		return "", err
	}

	k.logger.Info().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

func (k *KodiakLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}
	contract, err := sc.NewKodiakV1(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate KodiakV1 smart contract, %v", err)
		return nil, err
	}

	kc := KodiakConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// token0
	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for kodiak vault %s, %v", address, err)
		return nil, err
	}
	kc.Token0 = strings.ToLower(addr.Hex())

	// token1
	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for kodiak vault %s, %v", address, err)
		return nil, err
	}
	kc.Token1 = strings.ToLower(addr.Hex())

	// decimals
	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	kc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(kc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (k *KodiakLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch underlying balances
	amount0, amount1, err := k.getUnderlyingBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := k.getPrice(k.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	amount0Decimal := NormalizeAmount(amount0, price0.Decimals)

	price1, err := k.getPrice(k.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}
	amount1Decimal := NormalizeAmount(amount1, price1.Decimals)
	totalValue := amount0Decimal.Mul(price0.Price).Add(amount1Decimal.Mul(price1.Price))
	return totalValue, nil
}

func (k *KodiakLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := k.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		k.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (k *KodiakLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
	}
	ts, err := k.contract.TotalSupply(opts)
	if err != nil {
		k.logger.Error().Msgf("failed to obtain total supply for kodiak vault %s, %v", k.address.String(), err)
		return nil, fmt.Errorf("failed to get kodiak total supply, err: %w", err)
	}

	return ts, err
}

// getUnderlyingBalances fetches the underlying token balances.
func (k *KodiakLPPriceProvider) getUnderlyingBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
	}

	if k.contractV2 == nil { // V3 Island or Pool Logic
		ubs, err := k.contract.GetUnderlyingBalances(opts)
		if err != nil {
			k.logger.Error().Msgf("failed to obtain underlying balances for kodiak island %s, %v", k.address.String(), err)
			return nil, nil, fmt.Errorf("failed to get kodiak island underlying balances, err: %w", err)
		}
		return ubs.Amount0Current, ubs.Amount1Current, err
	} else { // V2 Pool Logic
		reserves, err := k.contractV2.GetReserves(opts)
		if err != nil {
			k.logger.Error().Msgf("failed to obtain reserve balances for kodiak V2 pool %s, %v", k.address.String(), err)
			return nil, nil, fmt.Errorf("failed to get kodiak V2 pool reserves, err: %w", err)
		}
		return reserves.Reserve0, reserves.Reserve1, err
	}
}
