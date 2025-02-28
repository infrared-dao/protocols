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

type BeraBorrowCDPConfig struct {
	LPTDecimals uint `json:"lpt_decimals"`
}

// BeraBorrowLPPriceProvider defines the provider for BeraBorrow CDP price and CDP TVL.
type BeraBorrowLPPriceProvider struct {
	cdpAddress  common.Address
	logger      zerolog.Logger
	block       *big.Int
	configBytes []byte
	config      *BeraBorrowCDPConfig
	cdpContract *sc.BeraBorrowCDP
}

// NewBeraBorrowLPPriceProvider creates a new instance of the BeraBorrowLPPriceProvider.
func NewBeraBorrowLPPriceProvider(cdpAddress common.Address, block *big.Int, logger zerolog.Logger, config []byte) *BeraBorrowLPPriceProvider {
	b := &BeraBorrowLPPriceProvider{
		cdpAddress:  cdpAddress,
		logger:      logger,
		block:       block,
		configBytes: config,
	}
	return b
}

// Initialize checks the configuration/data and instantiates the CDP contract.
func (b *BeraBorrowLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	b.config = &BeraBorrowCDPConfig{}
	err = json.Unmarshal(b.configBytes, b.config)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	b.cdpContract, err = sc.NewBeraBorrowCDP(b.cdpAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate Compounding Infrared Collateral Vault contract")
		return err
	}

	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD
func (b *BeraBorrowLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {

	opts := &bind.CallOpts{
		Pending:     false,
		Context:     ctx,
		BlockNumber: b.block,
	}

	// fetchPrice(opts) on Compounding Infrared Collateralized Vault returns USD price in decimals 18

	pricePerToken18, err := b.cdpContract.BeraBorrowCDPCaller.FetchPrice(opts)
	if err != nil {
		return "", err
	}
	pricePerToken := NormalizeAmount(pricePerToken18, uint(18))

	b.logger.Info().
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(8), nil
}

// TVL returns the Total Value Locked in the CDP in USD cents (1 USD = 100 cents).
func (b *BeraBorrowLPPriceProvider) TVL(ctx context.Context) (string, error) {

	opts := &bind.CallOpts{
		Pending:     false,
		Context:     ctx,
		BlockNumber: b.block,
	}

	pricePerToken18, err := b.cdpContract.BeraBorrowCDPCaller.FetchPrice(opts)
	if err != nil {
		return "", err
	}
	pricePerToken := NormalizeAmount(pricePerToken18, uint(18))

	totalSupply, err := b.cdpContract.BeraBorrowCDPCaller.TotalSupply(opts)
	if err != nil {
		return "", err
	}
	numTokens := NormalizeAmount(totalSupply, b.config.LPTDecimals)

	tvl := numTokens.Mul(pricePerToken)

	//b.logger.Info().
	//	Str("total value USD", tvl.String()).
	//	Msg("TVL calculated successfully")

	return tvl.StringFixed(8), nil
}

func (b *BeraBorrowLPPriceProvider) GetConfig(ctx context.Context, cdpAddress string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(cdpAddress) {
		err = fmt.Errorf("invalid smart contract address, '%s'", cdpAddress)
		return nil, err
	}

	cdpContract, err := sc.NewBeraBorrowCDP(common.HexToAddress(cdpAddress), client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate Compounding Infrared Collateral Vault contract on LP Token")
		return nil, err
	}

	bbcc := BeraBorrowCDPConfig{}
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	// decimals is uint8
	decimals, err := cdpContract.BeraBorrowCDPCaller.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", cdpAddress, err)
		return nil, err
	}
	bbcc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(bbcc)
	if err != nil {
		return nil, err
	}

	return body, nil
}
