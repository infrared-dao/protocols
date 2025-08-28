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

var _ Protocol = &LrBGTLPPriceProvider{}

type LrBGTConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	Token2      string `json:"token2"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// LrBGTLPPriceProvider defines the provider for LrBGT LP price and TVL.
type LrBGTLPPriceProvider struct {
	address            common.Address
	address2           common.Address
	address3           common.Address
	block              *big.Int
	priceMap           map[string]Price
	logger             zerolog.Logger
	configBytes        []byte
	config             *LrBGTConfig
	lrBGT              *sc.LrBGT
	lrBGTManager       *sc.LrBGTManager
	lrBGTManagerHelper *sc.LrBGTManagerHelper
}

// NewLrBGTLPPriceProvider creates a new instance of the NewLrBGTLPPriceProvider.
func NewLrBGTLPPriceProvider(
	address common.Address,
	address2 common.Address,
	address3 common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *LrBGTLPPriceProvider {
	a := &LrBGTLPPriceProvider{
		address:     address,
		address2:    address2,
		address3:    address3,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return a
}

// Initialize checks the configuration/data provided and instantiates the LrBGT smart contract.
func (a *LrBGTLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	a.config = &LrBGTConfig{}
	err = json.Unmarshal(a.configBytes, a.config)
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	// Validate that we have price data for the tokens
	_, ok := a.priceMap[a.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", a.config.Token0)
		a.logger.Error().Msg(err.Error())
		return err
	}

	_, ok = a.priceMap[a.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", a.config.Token1)
		a.logger.Error().Msg(err.Error())
		return err
	}

	_, ok = a.priceMap[a.config.Token2]
	if !ok {
		err = fmt.Errorf("no price data found for token2 (%s)", a.config.Token2)
		a.logger.Error().Msg(err.Error())
		return err
	}

	a.lrBGT, err = sc.NewLrBGT(a.address, client)
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to instantiate LrBGT smart contract")
		return err
	}

	a.lrBGTManager, err = sc.NewLrBGTManager(a.address2, client)
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to instantiate LrBGTManager smart contract")
		return err
	}

	a.lrBGTManagerHelper, err = sc.NewLrBGTManagerHelper(a.address3, client)
	if err != nil {
		a.logger.Error().Err(err).Msg("failed to instantiate LrBGTManagerHelper smart contract")
		return err
	}

	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD.
func (a *LrBGTLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// Fetch total supply
	totalSupply, err := a.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	totalValue, err := a.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, a.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	a.logger.Debug().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the protocol in USD.
func (a *LrBGTLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := a.totalValue(ctx)
	if err != nil {
		return "", err
	}

	a.logger.Debug().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

// GetConfig fetches and returns the configuration for the LrBGT protocol.
func (a *LrBGTLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	lrBGTManager, err := sc.NewLrBGTManager(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate lrBGTManager smart contract, %v", err)
		return nil, err
	}

	ac := LrBGTConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	lrBGTAddress, err := lrBGTManager.LairBgtTokenAddress(opts)

	lrBGT, err := sc.NewLrBGT(common.HexToAddress(lrBGTAddress.String()), client)

	bgtTokenAddress, err := lrBGTManager.BgtTokenAddress(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain bgtTokenAddress %v", err)
		return nil, err
	}

	vaultInfo, err := lrBGTManager.VaultInfo(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain vaultInfo %v", err)
		return nil, err
	}

	ac.Token0 = strings.ToLower(bgtTokenAddress.Hex())
	ac.Token1 = strings.ToLower(vaultInfo.Token0Address.Hex())
	ac.Token2 = strings.ToLower(vaultInfo.Token1Address.Hex())

	// decimals
	decimals, err := lrBGT.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	ac.LPTDecimals = uint(decimals)

	body, err := json.Marshal(ac)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (a *LrBGTLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	a.block = block
	if prices != nil {
		a.priceMap = prices
	}
}

// Internal Helper methods not able to be called except in this file

// Helper method to calculate the total value of underlying assets
func (a *LrBGTLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	totalSupply, err := a.getTotalSupply(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	bgtAmount, lpAmount, err := a.PredictedUnStakeAmount(ctx, totalSupply)
	if err != nil {
		return decimal.Zero, err
	}

	token0amount, token1Amount, err := a.lpPairTokenAmount(ctx, lpAmount)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := a.getPrice(a.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	amount0Decimal := NormalizeAmount(bgtAmount, price0.Decimals)

	price1, err := a.getPrice(a.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}
	amount1Decimal := NormalizeAmount(token0amount, price1.Decimals)

	price2, err := a.getPrice(a.config.Token2)
	if err != nil {
		return decimal.Zero, err
	}
	amount2Decimal := NormalizeAmount(token1Amount, price1.Decimals)

	totalValue := amount0Decimal.Mul(price0.Price).Add(amount1Decimal.Mul(price1.Price)).Add(amount2Decimal.Mul(price2.Price))
	return totalValue, nil
}

func (a *LrBGTLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := a.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		a.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (a *LrBGTLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: a.block,
	}
	ts, err := a.lrBGT.TotalSupply(opts)
	if err != nil {
		a.logger.Error().Msgf("failed to obtain total supply for LairBGT %s, %v", a.address.String(), err)
		return nil, fmt.Errorf("failed to get LairBGT total supply, err: %w", err)
	}

	return ts, err
}

// PredictedUnStakeAmount fetches the token balances.
func (a *LrBGTLPPriceProvider) PredictedUnStakeAmount(ctx context.Context, amount *big.Int) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: a.block,
	}

	lairBGTTokenAddress, err := a.lrBGTManager.LairBgtTokenAddress(opts)
	if err != nil {
		a.logger.Error().Msgf("failed to obtain lairBGTTokenAddress %v", err)
		return nil, nil, fmt.Errorf("failed to obtain lairBGTTokenAddress, err: %w", err)
	}

	bgtVaultAddress, err := a.lrBGTManager.BgtVaultAddress(opts)
	if err != nil {
		a.logger.Error().Msgf("failed to obtain bgtVaultAddress %v", err)
		return nil, nil, fmt.Errorf("failed to obtain bgtVaultAddress, err: %w", err)
	}

	lpVaultAddress, err := a.lrBGTManager.LpVaultAddress(opts)
	if err != nil {
		a.logger.Error().Msgf("failed to obtain lpVaultAddress %v", err)
		return nil, nil, fmt.Errorf("failed to obtain lpVaultAddress, err: %w", err)
	}

	bgtTokenAddress, err := a.lrBGTManager.BgtTokenAddress(opts)
	if err != nil {
		a.logger.Error().Msgf("failed to obtain bgtTokenAddress %v", err)
		return nil, nil, fmt.Errorf("failed to obtain bgtTokenAddress, err: %w", err)
	}

	vaultInfo, err := a.lrBGTManager.VaultInfo(opts)
	if err != nil {
		a.logger.Error().Msgf("failed to obtain vaultInfo %v", err)
		return nil, nil, fmt.Errorf("failed to obtain vaultInfo, err: %w", err)
	}

	ubs, err := a.lrBGTManagerHelper.PredictedUnStakeAmount(opts,
		a.address2,
		lairBGTTokenAddress,
		bgtVaultAddress,
		lpVaultAddress,
		bgtTokenAddress,
		vaultInfo.LpTokenAddress,
		vaultInfo.Token0Address,
		vaultInfo.Token1Address,
		amount)

	if err != nil {
		a.logger.Error().Msgf("failed to obtain unstake balances for lrbgt manager %s, %v", a.address2.String(), err)
		return nil, nil, fmt.Errorf("failed to get lrbgt manager balances, err: %w", err)
	}
	return ubs.UnStakeBgtTokenAmount, ubs.UnStakeBgtTokenAmount, err
}

// PredictedUnStakeAmount fetches the token balances.
func (a *LrBGTLPPriceProvider) lpPairTokenAmount(ctx context.Context, amount *big.Int) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: a.block,
	}

	vaultInfo, err := a.lrBGTManager.VaultInfo(opts)
	if err != nil {
		a.logger.Error().Msgf("failed to obtain vaultInfo %v", err)
		return nil, nil, fmt.Errorf("failed to obtain vaultInfo, err: %w", err)
	}

	ubs, err := a.lrBGTManagerHelper.LpPairTokenAmount(opts, vaultInfo.LpTokenAddress, amount)
	if err != nil {
		a.logger.Error().Msgf("failed to obtain lpPairTokenAmount %s, %v", a.address2.String(), err)
		return nil, nil, fmt.Errorf("failed to lpPairTokenAmount, err: %w", err)
	}
	return ubs.Amount0, ubs.Amount1, err
}
