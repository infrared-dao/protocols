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

type BexPoolConfig struct {
	Base        string `json:"base"`
	Quote       string `json:"quote"`
	IDX         string `json:"idx"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// BexLPPriceProvider defines the provider for BEX LP price and Pool TVL.
type BexLPPriceProvider struct {
	queryAddress   common.Address
	lpTokenAddress common.Address
	logger         zerolog.Logger
	priceMap       map[string]Price
	block          *big.Int
	configBytes    []byte
	config         *BexPoolConfig
	queryContract  *sc.CrocQuery
	erc20Contract  *sc.ERC20
}

// NewBexLPPriceProvider creates a new instance of the BexLPPriceProvider.
func NewBexLPPriceProvider(crocqueryAddress common.Address, lpTokenAddress common.Address, prices map[string]Price, block *big.Int, logger zerolog.Logger, config []byte) *BexLPPriceProvider {
	b := &BexLPPriceProvider{
		queryAddress:   crocqueryAddress,
		lpTokenAddress: lpTokenAddress,
		logger:         logger,
		priceMap:       prices,
		block:          block,
		configBytes:    config,
	}
	return b
}

// Initialize checks the configuration/data and instantiates the CrocQuery and LP Token ERC20, CrocLPERC20 smart contracts.
func (b *BexLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	b.config = &BexPoolConfig{}
	err = json.Unmarshal(b.configBytes, b.config)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}
	_, ok := b.priceMap[b.config.Base]
	if !ok {
		err = fmt.Errorf("no price data found for base token (%s)", b.config.Base)
		b.logger.Error().Msg(err.Error())
		return err
	}
	_, ok = b.priceMap[b.config.Quote]
	if !ok {
		err = fmt.Errorf("no price data found for token1 (%s)", b.config.Quote)
		b.logger.Error().Msg(err.Error())
		return err
	}

	b.queryContract, err = sc.NewCrocQuery(b.queryAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate CrocQuery smart contract")
		return err
	}

	b.erc20Contract, err = sc.NewERC20(b.lpTokenAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate ERC20 smart contract on LP Token")
		return err
	}
	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD cents (1 USD = 100 cents).
func (b *BexLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Pending:     false,
		Context:     ctx,
		BlockNumber: b.block,
	}

	// Fetch total supply from ERC20 interface
	totalSupply, err := b.erc20Contract.ERC20Caller.TotalSupply(opts)
	if err != nil {
		return "", err
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		b.logger.Error().Err(err).Msg("Invalid totalSupply")
		return "", err
	}

	totalValue, err := b.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, b.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	b.logger.Info().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(8), nil
}

// TVL returns the Total Value Locked in the pool in USD cents (1 USD = 100 cents).
func (b *BexLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := b.totalValue(ctx)
	if err != nil {
		return "", err
	}

	b.logger.Info().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(8), nil
}

func (b *BexLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	erc20Contract, err := sc.NewERC20(common.HexToAddress(address), client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate ERC20 smart contract on LP Token")
		return nil, err
	}

	crocLPERC20Contract, err := sc.NewCrocLPERC20(common.HexToAddress(address), client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate CrocLPERC20 smart contract on LP Token")
		return nil, err
	}

	bpc := BexPoolConfig{}
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	// returns as *big.Int
	idx, err := crocLPERC20Contract.CrocLPERC20Caller.PoolType(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain pool type idx for bex pool %s, %v", address, err)
		return nil, err
	}
	bpc.IDX = idx.String()

	// base token address
	addr, err := crocLPERC20Contract.CrocLPERC20Caller.BaseToken(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain base token address for bex pool %s, %v", address, err)
		return nil, err
	}
	bpc.Base = strings.ToLower(addr.Hex())

	// quote token address
	addr, err = crocLPERC20Contract.CrocLPERC20Caller.QuoteToken(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain base token address for bex pool %s, %v", address, err)
		return nil, err
	}
	bpc.Quote = strings.ToLower(addr.Hex())

	// decimals is uint8
	decimals, err := erc20Contract.ERC20Caller.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	bpc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(bpc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (b *BexLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch underlying balances
	amountBase, amountQuote, err := b.getUnderlyingBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	priceBase, err := b.getPrice(b.config.Base)
	if err != nil {
		return decimal.Zero, err
	}
	amountBaseDecimal := NormalizeAmount(amountBase, priceBase.Decimals)

	priceQuote, err := b.getPrice(b.config.Quote)
	if err != nil {
		return decimal.Zero, err
	}
	amountQuoteDecimal := NormalizeAmount(amountQuote, priceQuote.Decimals)

	totalValue := amountBaseDecimal.Mul(priceBase.Price).Add(amountQuoteDecimal.Mul(priceQuote.Price))
	return totalValue, nil
}

func (b *BexLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := b.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		b.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getUnderlyingBalances fetches the underlying virtual token supply for each token.
func (b *BexLPPriceProvider) getUnderlyingBalances(ctx context.Context) (*big.Int, *big.Int, error) {
	opts := &bind.CallOpts{
		Pending:     false,
		Context:     ctx,
		BlockNumber: b.block,
	}

	// BEX Pools on CrocSwap do not actually have a fixed supply we can look up
	// They have multiple different liquidity types but we can compute virtual token supply at a point
	// 1. Get the price ratio at the current point on swap curve
	// 2. Get the virtual liquidity at that point on swap curve
	// 3. Compute the virtual token supplies from ratio and liquidity

	base := common.HexToAddress(b.config.Base)
	quote := common.HexToAddress(b.config.Quote)
	poolIdx, ok := new(big.Int).SetString(b.config.IDX, 10)
	if !ok {
		return nil, nil, errors.New("unable to parse poolIdx as *big.Int from string")
	}

	// 1. Get the price ratio at the current point on swap curve
	// pricePoint is a Q64.Q64 representation of the sqrt of price ratio in a *big.Int
	pricePoint, err := b.queryContract.CrocQueryCaller.QueryPrice(opts, base, quote, poolIdx)
	if err != nil {
		return nil, nil, err
	}

	numerator := decimal.NewFromBigInt(pricePoint, 0)
	divisor := decimal.NewFromInt(2).Pow(decimal.NewFromInt(64))
	sqrtRatio := numerator.Div(divisor)

	// 2. Get the virtual liquidity at that point on swap curve
	// Liquidity is square root of the product of base pool supply and quote pool supply
	liquidity, err := b.queryContract.CrocQueryCaller.QueryLiquidity(opts, base, quote, poolIdx)
	if err != nil {
		return nil, nil, err
	}
	b.logger.Info().Msgf("liquidity is %s", liquidity.String())
	liquidityDecimal := decimal.NewFromBigInt(liquidity, 0)

	baseSupplyDecimal := liquidityDecimal.Mul(sqrtRatio)
	quoteSupplyDecimal := liquidityDecimal.Div(sqrtRatio)

	baseSupply, ok := new(big.Int).SetString(baseSupplyDecimal.Round(0).String(), 10)
	if !ok {
		return nil, nil, errors.New("unable to parse baseSupply as *big.Int from string")
	}
	quoteSupply, ok := new(big.Int).SetString(quoteSupplyDecimal.Round(0).String(), 10)
	if !ok {
		return nil, nil, errors.New("unable to parse baseSupply as *big.Int from string")
	}

	b.logger.Info().Msgf("base and quote supply are %s %s", baseSupply.String(), quoteSupply.String())
	return baseSupply, quoteSupply, nil
}
