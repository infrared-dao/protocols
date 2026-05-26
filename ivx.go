package protocols

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/infrared-dao/protocols/fetchers"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/infrared-dao/protocols/multicall3"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

var _ Protocol = &IVXLPPriceProvider{}
var _ BatchablePriceProvider = &IVXLPPriceProvider{}

type IVXLPConfig struct {
	LPTDecimals uint `json:"lpt_decimals"`
}

// IVXLPPriceProvider defines the provider for IVXLP price and TVL.
type IVXLPPriceProvider struct {
	lpMonitorAddress  common.Address
	lpTokenAddress    common.Address
	block             *big.Int
	logger            zerolog.Logger
	configBytes       []byte
	config            *IVXLPConfig
	lpMonitorContract *sc.IVXLPMonitor
}

// NewIVXLPPriceProvider creates a new instance of the IVXLPPriceProvider.
func NewIVXLPPriceProvider(
	lpMonitorAddress common.Address,
	lpTokenAddress common.Address,
	block *big.Int,
	logger zerolog.Logger,
	config []byte,
) *IVXLPPriceProvider {
	return &IVXLPPriceProvider{
		lpMonitorAddress: lpMonitorAddress,
		lpTokenAddress:   lpTokenAddress,
		block:            block,
		logger:           logger,
		configBytes:      config,
	}
}

// Initialize checks the configuration/data and instantiates the LP Token ERC20 smart contract.
func (p *IVXLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
	var err error

	p.config = &IVXLPConfig{}
	err = json.Unmarshal(p.configBytes, p.config)
	if err != nil {
		p.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	// Initialize IVXLPMonitor contract
	p.lpMonitorContract, err = sc.NewIVXLPMonitor(p.lpMonitorAddress, client)
	if err != nil {
		p.logger.Error().Err(err).Msg("failed to instantiate IVXLPMonitor smart contract")
		return err
	}

	return nil
}

// LPTokenPrice returns the current price of the IVXLP token in USD.
func (p *IVXLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: p.block,
	}
	sharePrice, err := p.lpMonitorContract.GetSharePrice(opts)
	if err != nil {
		return "", fmt.Errorf("failed to get IVXLP share price, err: %w", err)
	}
	price := p.computeLPPriceFromReads(sharePrice)
	return price.StringFixed(roundingDecimals), nil
}

func (p *IVXLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	abi, err := sc.IVXLPMonitorMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("ivx: ABI: %w", err)
	}
	spData, err := abi.Pack("getSharePrice")
	if err != nil {
		return nil, fmt.Errorf("ivx: pack getSharePrice: %w", err)
	}
	return []multicall3.Call3{
		{Target: p.lpMonitorAddress, AllowFailure: true, CallData: spData},
	}, nil
}

func (p *IVXLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 1 {
		return decimal.Zero, fmt.Errorf("ivx: expected 1 response, got %d", len(responses))
	}
	if !responses[0].Success {
		return decimal.Zero, fmt.Errorf("ivx: getSharePrice reverted")
	}
	abi, err := sc.IVXLPMonitorMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("ivx: ABI: %w", err)
	}
	spOut, err := abi.Methods["getSharePrice"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("ivx: unpack getSharePrice: %w", err)
	}
	sharePrice, ok := spOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("ivx: sharePrice type %T", spOut[0])
	}
	return p.computeLPPriceFromReads(sharePrice), nil
}

func (p *IVXLPPriceProvider) computeLPPriceFromReads(sharePrice *big.Int) decimal.Decimal {
	pricePerToken := NormalizeAmount(sharePrice, p.config.LPTDecimals)
	p.logger.Debug().Str("Token Price", pricePerToken.String()).Msg("LP token price fetched successfully")
	return pricePerToken
}

// TVL returns the Total Value Locked in the pool in USD.
func (p *IVXLPPriceProvider) TVL(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: p.block,
	}

	// Fetch TVL from IVXLPMonitor contract
	tvl, err := p.lpMonitorContract.GetTVL(opts)
	if err != nil {
		return "", fmt.Errorf("failed to get TVL from IVXLPMonitor, err: %w", err)
	}

	tvlInUSD := NormalizeAmount(tvl, p.config.LPTDecimals)

	p.logger.Debug().
		Str("TVL In USD", tvlInUSD.String()).
		Msg("TVL fetched successfully")

	return tvlInUSD.StringFixed(roundingDecimals), nil
}

// GetConfig returns the configuration for the IVXLP pool.
func (p *IVXLPPriceProvider) GetConfig(ctx context.Context, address string, client bind.ContractBackend) ([]byte, error) {
	if !common.IsHexAddress(address) {
		return nil, fmt.Errorf("invalid smart contract address, '%s'", address)
	}

	erc20Contract, err := sc.NewERC20Caller(common.HexToAddress(address), client)
	if err != nil {
		p.logger.Error().Err(err).Msg("failed to instantiate ERC20 contract on IVX LP Token")
		return nil, err
	}

	// Construct the configuration object
	config := IVXLPConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	decimals, err := erc20Contract.Decimals(opts)
	if err != nil {
		p.logger.Error().Err(err).Msg("failed to get ERC20 decimals from IVX LP Token")
		return nil, err
	}
	config.LPTDecimals = uint(decimals)

	// Marshal the configuration into JSON
	body, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal configuration to JSON: %w", err)
	}

	return body, nil
}

func (b *IVXLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	b.block = block
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for IVX protocol
func (b *IVXLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}
