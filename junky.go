package protocols

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/rs/zerolog"
)

// The Junky adapter is simpler than the AMM style adapters for several reasons:

// The only specific contract we need is the singleton JunkyVaultManager contract
// Junky vaults only have a single token inside, this underlying token is called the 'liqToken'
// The only config information we need is the decimals for the staking token called the 'vaultToken'

// A VaultManager getter directly tells us a scaled vaultToken price relative to the liqToken

// For example: juJNKY is the staking token of an infrared vault with only JNKY as an underlying token
// In their terminology JNKY is the "liqToken" and juJNKY is the "vaultToken"
// We can pass the address of liqToken JNKY into getVaultTokenPrice() to get the scaled conversion rate

// JunkyVaultManager returns conversion rate scaled by 10^5 to represent as uint instead of a decimal
const scaleFactor = uint(5)

type JunkyVaultConfig struct {
	LPTDecimals uint `json:"lpt_decimals"`
}

// JunkyLPPriceProvider defines the provider for Junky LP price and TVL.
type JunkyLPPriceProvider struct {
	address     common.Address
	liqAddress  common.Address
	logger      zerolog.Logger
	priceMap    map[string]Price
	configBytes []byte
	config      *JunkyVaultConfig
	contract    *sc.JunkyVaultManager
}

// NewJunkyLPPriceProvider creates a new instance of the JunkyLPPriceProvider.
// Note that unlike other adapters which use the address of the vault to connect,
// this address being passed in is the singleton JunkyVaultManager contract address
func NewJunkyLPPriceProvider(address common.Address, prices map[string]Price, logger zerolog.Logger, config []byte) *JunkyLPPriceProvider {
	return &JunkyLPPriceProvider{
		address:     address,
		logger:      logger,
		priceMap:    prices,
		configBytes: config,
	}
}

// Initialize checks the configuration/data provided and instantiates the Junky smart contract.
func (j *JunkyLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	j.config = &JunkyVaultConfig{}
	err = json.Unmarshal(j.configBytes, j.config)
	if err != nil {
		j.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	// Junky vaults should only have single underlying token, don't care about its price, just address
	liqTokenAddress := ""
	for underlyingAddress := range j.priceMap {
		liqTokenAddress = strings.ToLower(underlyingAddress)
	}
	if !common.IsHexAddress(liqTokenAddress) {
		err = fmt.Errorf("invalid underlying 'liqToken' address, '%s'", liqTokenAddress)
		j.logger.Error().Msg(err.Error())
		return err
	} else {
		j.liqAddress = common.HexToAddress(liqTokenAddress)
	}

	j.contract, err = sc.NewJunkyVaultManager(j.address, client)
	if err != nil {
		j.logger.Error().Err(err).Msg("failed to instatiate JunkyVaultManager smart contract")
		return err
	}
	return nil
}

// LPTokenPrice returns the current price of the Junky vaultToken in USD cents (1 USD = 100 cents).
func (j *JunkyLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {

	totalValue, err := j.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, j.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	j.logger.Info().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(8), nil
}

// TVL returns the Total Value Locked in the Junky vault in USD cents (1 USD = 100 cents).
func (j *JunkyLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := j.totalValue(ctx)
	if err != nil {
		return "", err
	}

	j.logger.Info().
		Str("totalValue", totalValue.String()).
		Msg("TVL calculated successfully")

	return totalValue.StringFixed(8), nil
}

// Just need to get the decimals in the LP staking token from ERC20 call
func (j *JunkyLPPriceProvider) GetConfig(ctx context.Context, address string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid LP Token address, '%s'", address)
		return nil, err
	}

	erc20Contract, err := sc.NewERC20(common.HexToAddress(address), client)
	if err != nil {
		j.logger.Error().Err(err).Msg("failed to instatiate ERC20 smart contract on LP Token")
		return nil, err
	}

	jvc := JunkyVaultConfig{}
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	// decimals comes out as uint8, cast to uint
	decimals, err := erc20Contract.ERC20Caller.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", address, err)
		return nil, err
	}
	jvc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(jvc)
	if err != nil {
		return nil, err
	}

	return body, nil
}
