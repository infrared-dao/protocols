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

var _ Protocol = &EtherfiLPPriceProvider{}

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
func (e *EtherfiLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
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
	ts, err := e.getTotalSupply(ctx)
	if err != nil {
		return "", err
	}

	if ts.Cmp(big.NewInt(0)) == 0 {
		err = fmt.Errorf("total supply is zero")
		e.logger.Error().Err(err).Msg("failed to fetch total supply")
		return "", err
	}

	tvl, err := e.tvl(ctx)
	if err != nil {
		return "", err
	}

	tsd := NormalizeAmount(ts, e.config.LPTDecimals)
	price := tvl.Div(tsd)

	e.logger.Debug().
		Str("totalValue", tvl.String()).
		Str("totalSupply", ts.String()).
		Str("pricePerToken", price.String()).
		Msg("LP token price calculated successfully")

	return price.StringFixed(roundingDecimals), nil
}

func (e *EtherfiLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := e.tvl(ctx)
	if err != nil {
		return "", err
	}

	e.logger.Debug().Str("tvl", totalValue.String()).Msg("successfully fetched TVL")
	return totalValue.StringFixed(roundingDecimals), nil
}

func (w *EtherfiLPPriceProvider) GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error) {
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

		contract, err := sc.NewEtherfiAccountant(common.HexToAddress(accountantAddress), ethClient)
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
