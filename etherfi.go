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

var _ Protocol = &EtherfiLPPriceProvider{}
var _ BatchablePriceProvider = &EtherfiLPPriceProvider{}

var accountantContracts = []string{
	"0x88ea516DCb9f79CAFA9D0d19909A4dbd7B6890c8",
	"0x55ee6E1ADF848a2Fc831B07564223396ef6258d4",
}

type EtherfiConfig struct {
	Asset       string `json:"asset"`
	Accountant  string `json:"accountant"`
	LPTDecimals uint   `json:"lpt_decimals"`
}

// EtherfiLPPriceProvider defines the provider for Etherfi Token price and TVL.
type EtherfiLPPriceProvider struct {
	address     common.Address
	block       *big.Int
	priceMap    map[string]Price
	logger      zerolog.Logger
	configBytes []byte
	config      *EtherfiConfig
	contract    *sc.EtherfiVault
	accountant  *sc.EtherfiAccountant
}

// NewEtherfiLPPriceProvider creates a new instance of the EtherfiLPPriceProvider.
func NewEtherfiLPPriceProvider(
	address common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *EtherfiLPPriceProvider {
	e := &EtherfiLPPriceProvider{
		address:     address,
		block:       block,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
	return e
}

// Initialize checks the configuration/data provided and instantiates the Etherfi smart contract.
func (e *EtherfiLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
	var err error

	e.config = &EtherfiConfig{}
	err = json.Unmarshal(e.configBytes, e.config)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	_, ok := e.priceMap[e.config.Asset]
	if !ok {
		err = fmt.Errorf("no price data found for asset (%s)", e.config.Asset)
		e.logger.Error().Msg(err.Error())
		return err
	}

	e.contract, err = sc.NewEtherfiVault(e.address, client)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to instantiate Etherfi vault contract")
		return err
	}

	e.accountant, err = sc.NewEtherfiAccountant(common.HexToAddress(e.config.Accountant), client)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to instantiate Etherfi accountant contract")
		return err
	}

	return nil
}

func (e *EtherfiLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{Context: ctx, BlockNumber: e.block}
	ts, err := e.contract.TotalSupply(opts)
	if err != nil {
		return "", fmt.Errorf("etherfi: totalSupply: %w", err)
	}
	rate, err := e.accountant.GetRate(opts)
	if err != nil {
		return "", fmt.Errorf("etherfi: accountant.getRate: %w", err)
	}
	price, err := e.computeLPPriceFromReads(ts, rate)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

func (e *EtherfiLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	vaultABI, err := sc.EtherfiVaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("etherfi: vault ABI: %w", err)
	}
	tsData, err := vaultABI.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("etherfi: pack totalSupply: %w", err)
	}
	accABI, err := sc.EtherfiAccountantMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("etherfi: accountant ABI: %w", err)
	}
	rateData, err := accABI.Pack("getRate")
	if err != nil {
		return nil, fmt.Errorf("etherfi: pack getRate: %w", err)
	}
	return []multicall3.Call3{
		{Target: e.address, AllowFailure: true, CallData: tsData},
		{Target: common.HexToAddress(e.config.Accountant), AllowFailure: true, CallData: rateData},
	}, nil
}

func (e *EtherfiLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 2 {
		return decimal.Zero, fmt.Errorf("etherfi: expected 2 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("etherfi: sub-call %d reverted", i)
		}
	}
	vaultABI, err := sc.EtherfiVaultMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("etherfi: vault ABI: %w", err)
	}
	tsOut, err := vaultABI.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("etherfi: unpack totalSupply: %w", err)
	}
	totalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("etherfi: totalSupply type %T", tsOut[0])
	}
	accABI, err := sc.EtherfiAccountantMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("etherfi: accountant ABI: %w", err)
	}
	rateOut, err := accABI.Methods["getRate"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("etherfi: unpack getRate: %w", err)
	}
	rate, ok := rateOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("etherfi: rate type %T", rateOut[0])
	}
	return e.computeLPPriceFromReads(totalSupply, rate)
}

func (e *EtherfiLPPriceProvider) computeLPPriceFromReads(totalSupply, rate *big.Int) (decimal.Decimal, error) {
	if totalSupply.Sign() == 0 {
		err := fmt.Errorf("total supply is zero")
		e.logger.Error().Err(err).Msg("invalid totalSupply")
		return decimal.Zero, err
	}
	assetPrice, err := e.getPrice(e.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	// Mirrors the legacy tvl()+LPTokenPrice combo: TVL normalizes the
	// share supply by the underlying asset's decimals, while the divisor
	// normalizes the same supply by LPTDecimals. The two cancel out the
	// totalSupply factor and reduce to rate·assetPrice scaled by the
	// LPTDecimals−assetDecimals delta — kept literal here so the batched
	// path is bit-identical to the legacy path.
	assetAmountDecimal := NormalizeAmount(totalSupply, assetPrice.Decimals)
	rateDecimal := NormalizeAmount(rate, e.config.LPTDecimals)
	tvl := assetAmountDecimal.Mul(rateDecimal).Mul(assetPrice.Price)
	tsd := NormalizeAmount(totalSupply, e.config.LPTDecimals)
	pricePerToken := tvl.Div(tsd)
	e.logger.Debug().Str("pricePerToken", pricePerToken.String()).Msg("LP token price calculated successfully")
	return pricePerToken, nil
}

func (e *EtherfiLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := e.tvl(ctx)
	if err != nil {
		return "", err
	}

	e.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (w *EtherfiLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	var err error
	efc := &EtherfiConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	for _, accountantAddress := range accountantContracts {
		if !common.IsHexAddress(address) {
			err = fmt.Errorf("invalid smart contract address, '%s'", address)
			return nil, err
		}

		contract, err := sc.NewEtherfiAccountant(common.HexToAddress(accountantAddress), client)
		if err != nil {
			err = fmt.Errorf("failed to instantiate Etherfi smart contract, %v", err)
			return nil, err
		}

		vaultAddress, err := contract.Vault(opts)
		if err != nil {
			err = fmt.Errorf("failed to fetch vault address, %v", err)
			return nil, err
		}

		if !strings.EqualFold(address, vaultAddress.Hex()) {
			continue
		}

		// This accountant corresponds to the vault given by address
		efc.Accountant = strings.ToLower(accountantAddress)

		assetAddress, err := contract.Base(opts)
		if err != nil {
			err = fmt.Errorf("failed to fetch base asset address, %v", err)
			return nil, err
		}
		efc.Asset = strings.ToLower(assetAddress.Hex())

		decimals, err := contract.Decimals(opts)
		if err != nil {
			err = fmt.Errorf("failed to fetch decimals, %v", err)
			return nil, err
		}
		efc.LPTDecimals = uint(decimals)

		body, err := json.Marshal(efc)
		if err != nil {
			return nil, err
		}

		return body, nil
	}

	return nil, fmt.Errorf("failed to find accountant for vault %s", address)
}

func (w *EtherfiLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	w.block = block
	if prices != nil {
		w.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for Etherfi protocol
func (w *EtherfiLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

///// Helpers

// tvl fetches the TVL from the Etherfi smart contract.
func (e *EtherfiLPPriceProvider) tvl(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: e.block,
	}

	assetAmount, err := e.contract.TotalSupply(opts)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to fetch vault supply")
		return decimal.Zero, err
	}
	assetRate, err := e.accountant.GetRate(opts)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to fetch accountant rate")
		return decimal.Zero, err
	}
	assetPrice, err := e.getPrice(e.config.Asset)
	if err != nil {
		return decimal.Zero, err
	}
	assetAmountDecimal := NormalizeAmount(assetAmount, assetPrice.Decimals)
	assetRateDecimal := NormalizeAmount(assetRate, e.config.LPTDecimals)
	tvl := assetAmountDecimal.Mul(assetRateDecimal).Mul(assetPrice.Price)
	return tvl, nil
}

// getPrice fetches the price of the token from the price map.
func (e *EtherfiLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := e.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		e.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getTotalSupply fetches the total supply of the LP token.
func (e *EtherfiLPPriceProvider) getTotalSupply(ctx context.Context) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: e.block,
	}
	totalSupply, err := e.contract.TotalSupply(opts)
	if err != nil {
		e.logger.Error().Err(err).Msg("failed to fetch total supply")
		return nil, err
	}
	return totalSupply, nil
}
