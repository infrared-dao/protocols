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

type BexV2PoolConfig struct {
	PoolID      [32]byte `json:"poolid"`
	LPTDecimals uint     `json:"lpt_decimals"`
}

// BexLPPriceProvider defines the provider for BEX LP price and Pool TVL.
type BexV2LPPriceProvider struct {
	vaultAddress  common.Address
	poolAddress   common.Address
	logger        zerolog.Logger
	priceMap      map[string]Price
	configBytes   []byte
	config        *BexV2PoolConfig
	vaultContract *sc.BalancerVault
	poolContract  *sc.BalancerBasePool
}

// NewBexLPPriceProvider creates a new instance of the BexLPPriceProvider.
func NewBexV2LPPriceProvider(vaultAddress common.Address, poolAddress common.Address, prices map[string]Price, logger zerolog.Logger, config []byte) *BexV2LPPriceProvider {
	b := &BexV2LPPriceProvider{
		vaultAddress: vaultAddress,
		poolAddress:  poolAddress,
		logger:       logger,
		priceMap:     prices,
		configBytes:  config,
	}
	return b
}

// Initialize checks the configuration/data and instantiates the Vault and Base Pool contracts.
func (b *BexV2LPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	b.config = &BexV2PoolConfig{}
	err = json.Unmarshal(b.configBytes, b.config)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	b.vaultContract, err = sc.NewBalancerVault(b.vaultAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate Balancer Vault contract")
		return err
	}

	b.poolContract, err = sc.NewBalancerBasePool(b.poolAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate Balancer Base Pool contract on LP Token")
		return err
	}
	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD
func (b *BexV2LPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {

	// Fetch total supply from Balancer Base Pool which implements ERC20 interface
	//totalSupply, err := b.poolContract.BalancerBasePoolCaller.TotalSupply(&bind.CallOpts{})

	// Using GetActualSupply because this is how much is circulating for pools which lock up some LP tokens
	totalSupply, err := b.poolContract.BalancerBasePoolCaller.GetActualSupply(&bind.CallOpts{})
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
func (b *BexV2LPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := b.totalValue(ctx)
	if err != nil {
		return "", err
	}

	//b.logger.Info().
	//	Str("totalValue", totalValue.String()).
	//	Msg("TVL calculated successfully")

	return totalValue.StringFixed(8), nil
}

func (b *BexV2LPPriceProvider) GetConfig(ctx context.Context, poolAddress string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(poolAddress) {
		err = fmt.Errorf("invalid smart contract address, '%s'", poolAddress)
		return nil, err
	}

	poolContract, err := sc.NewBalancerBasePool(common.HexToAddress(poolAddress), client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instatiate Balancer Base Pool contract on LP Token")
		return nil, err
	}

	bpc := BexV2PoolConfig{}
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	// returns as [32]byte
	poolID, err := poolContract.BalancerBasePoolCaller.GetPoolId(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain poolID for bex pool %s, %v", poolAddress, err)
		return nil, err
	}
	bpc.PoolID = poolID

	// decimals is uint8
	decimals, err := poolContract.BalancerBasePoolCaller.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", poolAddress, err)
		return nil, err
	}
	bpc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(bpc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (b *BexV2LPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch underlying balances as map[string]*big.Int
	balanceData, err := b.getUnderlyingBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	b.logger.Info().
		Msgf("Token Balances: %+v", balanceData)

	totalValue := decimal.Zero
	for token, balance := range balanceData {
		price, err := b.getPrice(token)
		if err != nil {
			return decimal.Zero, err
		}
		balanceDecimal := NormalizeAmount(balance, price.Decimals)
		totalValue = totalValue.Add(balanceDecimal.Mul(price.Price))
	}

	return totalValue, nil
}

func (b *BexV2LPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := b.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		b.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getUnderlyingBalances fetches the underlying virtual token supply for each token.
func (b *BexV2LPPriceProvider) getUnderlyingBalances(ctx context.Context) (map[string]*big.Int, error) {
	opts := &bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	/********************************************
		Returns data of type:
		struct {
			Tokens          []common.Address
			Balances        []*big.Int
			LastChangeBlock *big.Int
		}
	********************************************/
	poolTokens, err := b.vaultContract.BalancerVaultCaller.GetPoolTokens(opts, b.config.PoolID)
	if err != nil {
		return nil, err
	}

	var balanceData = make(map[string]*big.Int)
	for i, tokenAddress := range poolTokens.Tokens {
		token := strings.ToLower(tokenAddress.Hex())

		//verify this is always sound for all pool types
		if token == strings.ToLower(b.poolAddress.Hex()) {
			// ignore when some of the LP token is locked in pool itself
			// this is why should use actualSupply instead of totalSupply
			continue
		}

		balance := poolTokens.Balances[i]
		balanceData[token] = balance
	}

	return balanceData, nil
}
