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
	"github.com/infrared-dao/protocols/multicall3"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

var _ Protocol = &DolomiteLPPriceProvider{}
var _ BatchablePriceProvider = &DolomiteLPPriceProvider{}

type DolomiteConfig struct {
	Token0      string `json:"token0"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// DolomiteLPPriceProvider defines the provider for Dolomite dToken price and TVL.
type DolomiteLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *DolomiteConfig
	contract    *sc.ERC4626
}

// NewDolomiteLPPriceProvider creates a new instance of the DolomiteLPPriceProvider.
func NewDolomiteLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *DolomiteLPPriceProvider {
	d := &DolomiteLPPriceProvider{
		address:     address,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return d
}

// Initialize checks the configuration/data provided and instantiates the Dolomite smart contract.
func (d *DolomiteLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
	var err error

	d.config = &DolomiteConfig{}
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
		d.logger.Error().Err(err).Msg("failed to instantiate Dolomite smart contract")
		return err
	}

	return nil
}

func (d *DolomiteLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	ts, err := d.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}
	opts := &bind.CallOpts{Context: ctx, BlockNumber: d.block}
	ta, err := d.contract.TotalAssets(opts)
	if err != nil {
		return "", fmt.Errorf("dolomite: totalAssets: %w", err)
	}
	price, err := d.computeLPPriceFromReads(ts, ta)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

func (d *DolomiteLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	abi, err := sc.ERC4626MetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("dolomite: ABI: %w", err)
	}
	tsData, err := abi.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("dolomite: pack totalSupply: %w", err)
	}
	taData, err := abi.Pack("totalAssets")
	if err != nil {
		return nil, fmt.Errorf("dolomite: pack totalAssets: %w", err)
	}
	return []multicall3.Call3{
		{Target: d.address, AllowFailure: false, CallData: tsData},
		{Target: d.address, AllowFailure: false, CallData: taData},
	}, nil
}

func (d *DolomiteLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("dolomite: expected 2 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("dolomite: sub-call %d reverted", i)
		}
	}
	abi, err := sc.ERC4626MetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("dolomite: ABI: %w", err)
	}
	tsOut, err := abi.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("dolomite: unpack totalSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("dolomite: totalSupply type %T", tsOut[0])
	}
	taOut, err := abi.Methods["totalAssets"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("dolomite: unpack totalAssets: %w", err)
	}
	totalAssets, ok := taOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("dolomite: totalAssets type %T", taOut[0])
	}
	return d.computeLPPriceFromReads(totalSupply, totalAssets)
}

func (d *DolomiteLPPriceProvider) computeLPPriceFromReads(totalSupply, totalAssets *big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		err := fmt.Errorf("total supply is zero")
		d.logger.Error().Err(err).Msg("invalid totalSupply")
		return decimal.Zero, err
	}
	assetPrice, err := d.getPrice(d.config.Token0)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(totalAssets, assetPrice.Decimals)
	tvl := assetAmountDecimal.Mul(assetPrice.Price)
	tsd := NormalizeAmount(totalSupply, d.config.LPTDecimals)
	pricePerToken := tvl.Div(tsd)
	d.logger.Debug().Str("pricePerToken", pricePerToken.String()).Msg("LP token price calculated successfully")
	return pricePerToken, nil
}

func (d *DolomiteLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := d.tvl(ctx)
	if err != nil {
		return "", err
	}

	d.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (d *DolomiteLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	contract, err := sc.NewERC4626(common.HexToAddress(address), client)
	if err != nil {
		err = fmt.Errorf("failed to instantiate Dolomite smart contract, %v", err)
		return nil, err
	}

	dc := &DolomiteConfig{}
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

func (d *DolomiteLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	d.block = block
	if prices != nil {
		d.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for Dolomite protocol
func (d *DolomiteLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

// Internal Helper methods not able to be called except in this file

// tvl fetches the TVL from the Dolomite smart contract.
func (d *DolomiteLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
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
func (d *DolomiteLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := d.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		d.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (d *DolomiteLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: d.block,
	}
	totalSupply, err := d.contract.TotalSupply(opts)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to fetch total supply")
		return nil, err
	}
	return totalSupply, nil
}

func (d *DolomiteLPPriceProvider) getUnderlyingBalances(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: d.block,
	}

	amount0, err := d.contract.TotalAssets(opts)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to fetch total assets")
		return nil, err
	}

	return amount0, nil
}
