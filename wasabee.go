package protocols

import (
	"context"
	"encoding/json"
	"errors"
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

var _ Protocol = &WasabeeLPPriceProvider{}
var _ BatchablePriceProvider = &WasabeeLPPriceProvider{}

// WasabeeConfig defines the configuration for Wasabee adapter
type WasabeeConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// WasabeeLPPriceProvider defines the provider for Wasabee LP price and TVL.
type WasabeeLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *WasabeeConfig
	contract    *sc.WasabeeVault
}

// NewWasabeeLPPriceProvider creates a new instance of the WasabeeLPPriceProvider.
func NewWasabeeLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *WasabeeLPPriceProvider {
	w := &WasabeeLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return w
}

// Initialize checks the configuration/data provided and instantiates the Wasabee smart contract.
func (w *WasabeeLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, httpClient fetchers.HttpClient) error {
	var err error

	w.config = &WasabeeConfig{}
	err = json.Unmarshal(w.configBytes, w.config)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	// Validate that we have price data for the tokens
	_, ok := w.priceMap[w.config.Token0]
	if !ok {
		err = fmt.Errorf("no price data found for token0 (%s)", w.config.Token0)
		w.logger.Error().Msg(err.Error())
		return err
	}

	_, ok = w.priceMap[w.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", w.config.Token1)
		w.logger.Error().Msg(err.Error())
		return err
	}

	// Initialize Bulla contract
	w.contract, err = sc.NewWasabeeVault(w.address, client)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to instantiate Wasabee smart contract")
		return err
	}

	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD.
func (w *WasabeeLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	totalSupply, err := w.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}
	a0, a1, err := w.getPoolBalances(ctx)
	if err != nil {
		return "", err
	}
	price, err := w.computeLPPriceFromReads(totalSupply, a0, a1)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

func (w *WasabeeLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	abi, err := sc.WasabeeVaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("wasabee: ABI: %w", err)
	}
	tsData, err := abi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("wasabee: pack totalSupply: %w", err)
	}
	gtaData, err := abi.Pack("getTotalAmounts")
	if err != nil {
		return nil, fmt.Errorf("wasabee: pack getTotalAmounts: %w", err)
	}
	return []multicall3.Call3{
		{Target: w.address, AllowFailure: true, CallData: tsData},
		{Target: w.address, AllowFailure: true, CallData: gtaData},
	}, nil
}

func (w *WasabeeLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("wasabee: expected 2 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("wasabee: sub-call %d reverted", i)
		}
	}
	abi, err := sc.WasabeeVaultMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("wasabee: ABI: %w", err)
	}
	tsOut, err := abi.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("wasabee: unpack totalSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("wasabee: totalSupply type %T", tsOut[0])
	}
	gtaOut, err := abi.Methods["getTotalAmounts"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("wasabee: unpack getTotalAmounts: %w", err)
	}
	if len(gtaOut) != 2 {
		return decimal.Zero, fmt.Errorf("wasabee: getTotalAmounts returned %d fields, want 2", len(gtaOut))
	}
	a0, ok := gtaOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("wasabee: amount0 type %T", gtaOut[0])
	}
	a1, ok := gtaOut[1].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("wasabee: amount1 type %T", gtaOut[1])
	}
	return w.computeLPPriceFromReads(totalSupply, a0, a1)
}

func (w *WasabeeLPPriceProvider) computeLPPriceFromReads(totalSupply, amount0, amount1 *big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		w.logger.Error().Err(err).Msg("Invalid totalSupply")
		return decimal.Zero, err
	}
	price0, err := w.getPrice(w.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	price1, err := w.getPrice(w.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}
	a0d := NormalizeAmount(amount0, price0.Decimals)
	a1d := NormalizeAmount(amount1, price1.Decimals)
	totalValue := a0d.Mul(price0.Price).Add(a1d.Mul(price1.Price))
	tsd := NormalizeAmount(totalSupply, w.config.LPTDecimals)
	pricePerToken := totalValue.Div(tsd)
	w.logger.Debug().Str("pricePerToken", pricePerToken.String()).Msg("LP token price calculated successfully")
	return pricePerToken, nil
}

// TVL returns the Total Value Locked in the protocol in USD.
func (w *WasabeeLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := w.totalValue(ctx)
	if err != nil {
		return "", err
	}

	w.logger.Debug().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

// GetConfig returns the configuration for the Wasabee adapter.
func (w *WasabeeLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewWasabeeVault(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Wasabee smart contract, %v", err)
		return nil, err
	}

	wc := WasabeeConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// Token0
	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for Wasabee vault %s, %v", address, err)
		return nil, err
	}
	wc.Token0 = strings.ToLower(addr.Hex())

	// Token1
	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for Wasabee vault %s, %v", address, err)
		return nil, err
	}
	wc.Token1 = strings.ToLower(addr.Hex())

	// Decimals
	decimals, err := contract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	wc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(wc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (w *WasabeeLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	w.block = block
	if prices != nil {
		w.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for Wasabee protocol
func (w *WasabeeLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

// Internal Helper methods not able to be called except in this file

// totalValue calculates the total value of assets in the pool
func (w *WasabeeLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	total0, total1, err := w.getPoolBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := w.getPrice(w.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}

	price1, err := w.getPrice(w.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}

	amount0Decimal := NormalizeAmount(total0, price0.Decimals)
	amount1Decimal := NormalizeAmount(total1, price1.Decimals)

	value0 := amount0Decimal.Mul(price0.Price)
	value1 := amount1Decimal.Mul(price1.Price)

	totalValue := value0.Add(value1)

	return totalValue, nil
}

// getPrice retrieves the price for a given token
func (w *WasabeeLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := w.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		w.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token
func (w *WasabeeLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: w.block,
	}

	totalSupply, err := w.contract.TotalSupply(opts)
	if err != nil {
		w.logger.Error().Msgf("failed to obtain total supply for Wasabee pool %s, %v", w.address.String(), err)
		return nil, fmt.Errorf("failed to get Wasabee total supply, err: %w", err)
	}

	return totalSupply, nil
}

// getPoolBalances fetches the base and quote token balances in the pool
func (w *WasabeeLPPriceProvider) getPoolBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: w.block,
	}

	totalAmounts, err := w.contract.GetTotalAmounts(opts)
	if err != nil {
		w.logger.Error().Msgf("failed to obtain total amounts for Wasabee pool %s, %v", w.address.String(), err)
		return nil, nil, fmt.Errorf("failed to get Wasabee total amounts, err: %w", err)
	}

	return totalAmounts.Total0, totalAmounts.Total1, nil
}
