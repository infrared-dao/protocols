package protocols

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/rs/zerolog"
)

var _ Protocol = &IVXLPPriceProvider{}

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
func (p *IVXLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
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

	// Fetch share price from IVXLPMonitor contract
	sharePrice, err := p.lpMonitorContract.GetSharePrice(opts)
	if err != nil {
		return "", fmt.Errorf("failed to get IVXLP share price, err: %w", err)
	}

	pricePerToken := NormalizeAmount(sharePrice, p.config.LPTDecimals)

	p.logger.Debug().
		Str("Token Price", pricePerToken.String()).
		Msg("LP token price fetched successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
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
func (p *IVXLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
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
