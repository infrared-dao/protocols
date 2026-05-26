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

// enforce interface adherence
var _ Protocol = &BrownFiLPPriceProvider{}
var _ BatchablePriceProvider = &BrownFiLPPriceProvider{}

// Define core types for the BrownFi Adapter

type BrownFiConfig struct {
	Token0      string `json:"token0"`
	Token1      string `json:"token1"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// BrownFiLPPriceProvider defines the provider for BrownFi LP price and TVL.
type BrownFiLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *BrownFiConfig
	contract    *sc.BrownFiPool
}

// NewBrownFiLPPriceProvider creates a new instance of the BrownFiLPPriceProvider.
func NewBrownFiLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *BrownFiLPPriceProvider {
	w := &BrownFiLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return w
}

// Initialize checks the configuration/data provided and instantiates the BrownFiPool smart contract.
func (w *BrownFiLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
	var err error

	w.config = &BrownFiConfig{}
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
	_, ok = w.priceMap[w.config.Token1]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", w.config.Token1)
		w.logger.Error().Msg(err.Error())
		return err
	}

	// initialize BrownFiPool contract
	vault, err := sc.NewBrownFiPool(w.address, client)
	if err != nil {
		return err
	}
	w.contract = vault

	return nil
}

// LPTokenPrice — legacy path. Batchable: PriceReads + ComputePrice via Multicall3.
func (w *BrownFiLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	totalSupply, err := w.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}
	amount0, amount1, err := w.getBalances(ctx)
	if err != nil {
		return "", err
	}
	price, err := w.computeLPPriceFromReads(totalSupply, Balances{Amount0: amount0, Amount1: amount1})
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

func (w *BrownFiLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	abi, err := sc.BrownFiPoolMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("brownfi: ABI: %w", err)
	}
	tsData, err := abi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("brownfi: pack totalSupply: %w", err)
	}
	balData, err := abi.Pack("getReserves")
	if err != nil {
		return nil, fmt.Errorf("brownfi: pack getReserves: %w", err)
	}
	return []multicall3.Call3{
		{Target: w.address, AllowFailure: false, CallData: tsData},
		{Target: w.address, AllowFailure: false, CallData: balData},
	}, nil
}

func (w *BrownFiLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("brownfi: expected 2 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("brownfi: sub-call %d reverted", i)
		}
	}
	abi, err := sc.BrownFiPoolMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("brownfi: ABI: %w", err)
	}
	tsOut, err := abi.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("brownfi: unpack totalSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("brownfi: totalSupply type %T", tsOut[0])
	}
	balOut, err := abi.Methods["getReserves"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("brownfi: unpack getReserves: %w", err)
	}
	if len(balOut) < 2 {
		return decimal.Zero, fmt.Errorf("brownfi: getReserves returned %d values", len(balOut))
	}
	r0, ok := balOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("brownfi: reserve0 type %T", balOut[0])
	}
	r1, ok := balOut[1].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("brownfi: reserve1 type %T", balOut[1])
	}
	return w.computeLPPriceFromReads(totalSupply, Balances{Amount0: r0, Amount1: r1})
}

func (w *BrownFiLPPriceProvider) computeLPPriceFromReads(totalSupply *big.Int, balances Balances) (decimal.Decimal, error) {
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
	amount0Decimal := NormalizeAmount(balances.Amount0, price0.Decimals)
	amount1Decimal := NormalizeAmount(balances.Amount1, price1.Decimals)
	totalValue := amount0Decimal.Mul(price0.Price).Add(amount1Decimal.Mul(price1.Price))
	totalSupplyDecimal := NormalizeAmount(totalSupply, w.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)
	w.logger.Debug().Str("pricePerToken", pricePerToken.String()).Msg("LP token price calculated successfully")
	return pricePerToken, nil
}

// TVL returns the Total Value Locked in the vault as USD
func (w *BrownFiLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := w.totalValue(ctx)
	if err != nil {
		return "", err
	}

	w.logger.Debug().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(roundingDecimals), nil
}

func (w *BrownFiLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	// initialize BrownFiPool contract
	contract, err := sc.NewBrownFiPool(common.HexToAddress(address), client)
	if err != nil {
		w.logger.Error().Err(err).Msg("failed to instantiate BrownFiPool contract")
		return nil, err
	}

	wc := BrownFiConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// token0
	addr, err := contract.Token0(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token0 address for BrownFi vault %s, %v", address, err)
		return nil, err
	}
	wc.Token0 = strings.ToLower(addr.Hex())

	// token1
	addr, err = contract.Token1(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain token1 address for BrownFi vault %s, %v", address, err)
		return nil, err
	}
	wc.Token1 = strings.ToLower(addr.Hex())

	// decimals
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

func (w *BrownFiLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	w.block = block
	if prices != nil {
		w.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for BrownFi protocol
func (w *BrownFiLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

// Internal Helper methods not able to be called except in this file

func (w *BrownFiLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch internal balances
	amount0, amount1, err := w.getBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	price0, err := w.getPrice(w.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	amount0Decimal := NormalizeAmount(amount0, price0.Decimals)

	price1, err := w.getPrice(w.config.Token1)
	if err != nil {
		return decimal.Zero, err
	}
	amount1Decimal := NormalizeAmount(amount1, price1.Decimals)
	totalValue := amount0Decimal.Mul(price0.Price).Add(amount1Decimal.Mul(price1.Price))
	return totalValue, nil
}

func (w *BrownFiLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := w.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		w.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (w *BrownFiLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: w.block,
	}

	ts, err := w.contract.TotalSupply(opts)
	if err != nil {
		w.logger.Error().Msgf("failed to obtain total supply for BrownFi vault %s, %v",
			w.address.String(), err)
		return nil, fmt.Errorf("failed to get BrownFi vault total supply, err: %w", err)
	}

	return ts, err
}

// getBalances fetches the underlying token balances.
func (w *BrownFiLPPriceProvider) getBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: w.block,
	}

	balances, err := w.contract.GetReserves(opts)
	if err != nil {
		w.logger.Error().Msgf("failed to obtain balances for BrownFi vault %s, %v",
			w.address.String(), err)
		return nil, nil, fmt.Errorf("failed to get BrownFi vault balances, err: %w", err)
	}
	return balances.Reserve0, balances.Reserve1, err
}
