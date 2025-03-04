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
	"github.com/shopspring/decimal"
)

// This adapter is for a very specific type of BeraBorrow token setup
// InfraredWrapper type tokens are the LP Token in this case
// InfraredWrappers are 1-1 with the CompoundingInfraredCollateralVault token type

// The fetchPrice() onchain call always returns a USD price with 18 decimals
const USDPriceDecimals uint = 18

// LPTDecimals are the decimals of the InfraredWrapper token
// CDPDecimals are the decimals of the CompoundingInfraredCollateralVault token
type BeraBorrowCDPConfig struct {
	ColVaultAddress string `json:"col_vault_address"`
	LPTDecimals     uint   `json:"lpt_decimals"`
	CDPDecimals     uint   `json:"cdp_decimals"`
}

// BeraBorrowLPPriceProvider defines the provider for BeraBorrow CDP price and CDP TVL.
type BeraBorrowLPPriceProvider struct {
	LPTAddress  common.Address
	logger      zerolog.Logger
	block       *big.Int
	configBytes []byte
	config      *BeraBorrowCDPConfig
	lptContract *sc.BeraBorrowIW
	cdpContract *sc.BeraBorrowCICV
}

// NewBeraBorrowLPPriceProvider creates a new instance of the BeraBorrowLPPriceProvider.
func NewBeraBorrowLPPriceProvider(LPTAddress common.Address, block *big.Int, logger zerolog.Logger, config []byte) *BeraBorrowLPPriceProvider {
	b := &BeraBorrowLPPriceProvider{
		LPTAddress:  LPTAddress,
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

	// Initialize the IW contract from the address in the inputs for the LP token address
	b.lptContract, err = sc.NewBeraBorrowIW(b.LPTAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("adapter init failed to instatiate Infrared Wrapper contract")
		return err
	}

	// Initialize the CICV contract from the address in the config -- not the IW contract from the LP token address
	b.cdpContract, err = sc.NewBeraBorrowCICV(common.HexToAddress(b.config.ColVaultAddress), client)
	if err != nil {
		b.logger.Error().Err(err).Msg("adapter init failed to instatiate Compounding Infrared Collateral Vault contract")
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

	tvl, err := b.cicvTotalValue(ctx)
	if err != nil {
		return "", err
	}

	// returns a *big.Int for total supply of the LP token
	lpTotalSupply, err := b.lptContract.BeraBorrowIWCaller.TotalSupply(opts)
	if err != nil {
		return "", err
	}
	numTokens := NormalizeAmount(lpTotalSupply, b.config.LPTDecimals)
	pricePerToken := tvl.Div(numTokens)

	b.logger.Info().
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the CDP in USD
// Actually gets the TVL of the CICV which is 1-1 with the TVL of the IW which is the LP token
func (b *BeraBorrowLPPriceProvider) TVL(ctx context.Context) (string, error) {
	tvl, err := b.cicvTotalValue(ctx)
	if err != nil {
		return "", err
	}

	return tvl.StringFixed(roundingDecimals), nil
}

func (b *BeraBorrowLPPriceProvider) cicvTotalValue(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Pending:     false,
		Context:     ctx,
		BlockNumber: b.block,
	}

	pricePerToken18, err := b.cdpContract.BeraBorrowCICVCaller.FetchPrice(opts)
	if err != nil {
		return decimal.Zero, err
	}
	pricePerToken := NormalizeAmount(pricePerToken18, USDPriceDecimals)

	cdpTotalSupply, err := b.cdpContract.BeraBorrowCICVCaller.TotalSupply(opts)
	if err != nil {
		return decimal.Zero, err
	}
	numTokens := NormalizeAmount(cdpTotalSupply, b.config.CDPDecimals)

	cdpTotalValue := numTokens.Mul(pricePerToken)

	return cdpTotalValue, nil
}

func (b *BeraBorrowLPPriceProvider) GetConfig(ctx context.Context, lpAddress string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(lpAddress) {
		err = fmt.Errorf("invalid smart contract address, '%s'", lpAddress)
		return nil, err
	}

	bbcc := BeraBorrowCDPConfig{}

	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	iwContract, err := sc.NewBeraBorrowIW(common.HexToAddress(lpAddress), client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate Infrared Wrapper contract on LP Token")
		return nil, err
	}

	// decimals is uint8
	lptDecimals, err := iwContract.BeraBorrowIWCaller.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", lpAddress, err)
		return nil, err
	}
	bbcc.LPTDecimals = uint(lptDecimals)

	icvAddress, err := iwContract.BeraBorrowIWCaller.InfraredCollVault(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain the InfraredCollVault addressfor LP token %s, %v", lpAddress, err)
		return nil, err
	}
	bbcc.ColVaultAddress = icvAddress.Hex()

	// Initialize the CICV contract from address gotten from the IW contract
	cdpContract, err := sc.NewBeraBorrowCICV(icvAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate Compounding Infrared Collateral Vault contract for config")
		return nil, err
	}

	// decimals is uint8
	cdpDecimals, err := cdpContract.BeraBorrowCICVCaller.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for CDP (CICV) token %s, %v", icvAddress.Hex(), err)
		return nil, err
	}
	bbcc.CDPDecimals = uint(cdpDecimals)

	body, err := json.Marshal(bbcc)
	if err != nil {
		return nil, err
	}

	return body, nil
}
