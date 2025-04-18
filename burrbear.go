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

var _ Protocol = &BurrBearLPPriceProvider{}

// BurrBear is based on BalancerV2 which is the same codebase which BEX uses on mainnet
// Decided to implement it as a parallel code instead of a wrapper so it can get in config
//  the VaultContract address off the contract itself so we don't need to pass this in as
//  an env var like we originally did for BEX. This leads to different New function signature
// It can use the same Balancer contracts so there are no new abi or smart contract bindings

type BurrBearPoolConfig struct {
	VaultContract string   `json:"vault_contract"`
	PoolID        [32]byte `json:"poolid"`
	LPTDecimals   uint     `json:"lpt_decimals"`
}

// BurrBearLPPriceProvider defines the provider for BEX LP price and Pool TVL.
type BurrBearLPPriceProvider struct {
	poolAddress   common.Address
	block         *big.Int
	priceMap      map[string]Price
	logger        zerolog.Logger
	configBytes   []byte
	config        *BurrBearPoolConfig
	vaultContract *sc.BalancerVault
	poolContract  *sc.BalancerBasePool
}

// NewBurrBearLPPriceProvider creates a new instance of the BurrBearLPPriceProvider.
func NewBurrBearLPPriceProvider(
	poolAddress common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *BurrBearLPPriceProvider {
	b := &BurrBearLPPriceProvider{
		poolAddress: poolAddress,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
	return b
}

// Initialize checks the configuration/data and instantiates the Vault and Base Pool contracts.
func (bb *BurrBearLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	var err error

	bb.config = &BurrBearPoolConfig{}
	err = json.Unmarshal(bb.configBytes, bb.config)
	if err != nil {
		bb.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	vaultAddress := common.HexToAddress(bb.config.VaultContract)

	bb.vaultContract, err = sc.NewBalancerVault(vaultAddress, client)
	if err != nil {
		bb.logger.Error().Err(err).Msg("failed to instantiate Balancer Vault contract")
		return err
	}

	bb.poolContract, err = sc.NewBalancerBasePool(bb.poolAddress, client)
	if err != nil {
		bb.logger.Error().Err(err).Msg("failed to instantiate Balancer Base Pool contract on LP Token")
		return err
	}
	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD
func (bb *BurrBearLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: bb.block,
	}

	// Using GetActualSupply because this is how much is circulating for pools which lock up some LP tokens
	totalSupply, err := bb.poolContract.BalancerBasePoolCaller.GetActualSupply(opts)
	if err != nil {
		return "", err
	}

	// Avoid division by zero
	if totalSupply.Sign() == 0 {
		err := errors.New("totalSupply is zero, cannot calculate LP token price")
		bb.logger.Error().Err(err).Msg("Invalid totalSupply")
		return "", err
	}

	totalValue, err := bb.totalValue(ctx)
	if err != nil {
		return "", err
	}

	totalSupplyDecimal := NormalizeAmount(totalSupply, bb.config.LPTDecimals)
	pricePerToken := totalValue.Div(totalSupplyDecimal)

	bb.logger.Debug().
		Str("totalValue", totalValue.String()).
		Str("totalSupply", totalSupplyDecimal.String()).
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken.StringFixed(roundingDecimals), nil
}

// TVL returns the Total Value Locked in the pool in USD cents (1 USD = 100 cents).
func (bb *BurrBearLPPriceProvider) TVL(ctx context.Context) (string, error) {
	totalValue, err := bb.totalValue(ctx)
	if err != nil {
		return "", err
	}

	return totalValue.StringFixed(roundingDecimals), nil
}

func (bb *BurrBearLPPriceProvider) GetConfig(ctx context.Context, poolAddress string, client *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(poolAddress) {
		err = fmt.Errorf("invalid smart contract address, '%s'", poolAddress)
		return nil, err
	}

	poolContract, err := sc.NewBalancerBasePool(common.HexToAddress(poolAddress), client)
	if err != nil {
		bb.logger.Error().Err(err).Msg("failed to instantiate Balancer Base Pool contract on LP Token")
		return nil, err
	}

	bbpc := BurrBearPoolConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// returns as common.Address, need to get a string for saving with Hex() call
	vaultAddress, err := poolContract.BalancerBasePoolCaller.GetVault(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain poolID for bex pool %s, %v", poolAddress, err)
		return nil, err
	}
	bbpc.VaultContract = vaultAddress.Hex()

	// returns as [32]byte
	poolID, err := poolContract.BalancerBasePoolCaller.GetPoolId(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain poolID for bex pool %s, %v", poolAddress, err)
		return nil, err
	}
	bbpc.PoolID = poolID

	// decimals is uint8
	decimals, err := poolContract.BalancerBasePoolCaller.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", poolAddress, err)
		return nil, err
	}
	bbpc.LPTDecimals = uint(decimals)

	body, err := json.Marshal(bbpc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (bb *BurrBearLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	bb.block = block
	if prices != nil {
		bb.priceMap = prices
	}
}

// Internal Helper methods not able to be called except in this file

func (bb *BurrBearLPPriceProvider) totalValue(ctx context.Context) (decimal.Decimal, error) {
	var err error

	// Fetch underlying balances as map[string]*big.Int
	balanceData, err := bb.getUnderlyingBalances(ctx)
	if err != nil {
		return decimal.Zero, err
	}

	bb.logger.Debug().
		Msgf("Token Balances: %+v", balanceData)

	totalValue := decimal.Zero
	for token, balance := range balanceData {
		price, err := bb.getPrice(token)
		if err != nil {
			return decimal.Zero, err
		}
		balanceDecimal := NormalizeAmount(balance, price.Decimals)
		totalValue = totalValue.Add(balanceDecimal.Mul(price.Price))
	}

	return totalValue, nil
}

func (bb *BurrBearLPPriceProvider) getPrice(tokenKey string) (*Price, error) {
	price, ok := bb.priceMap[tokenKey]
	if !ok {
		err := fmt.Errorf("no price data found for token (%s)", tokenKey)
		bb.logger.Error().Msg(err.Error())
		return nil, err
	}
	return &price, nil
}

// getUnderlyingBalances fetches the underlying virtual token supply for each token.
func (bb *BurrBearLPPriceProvider) getUnderlyingBalances(ctx context.Context) (map[string]*big.Int, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: bb.block,
	}

	/********************************************
		Returns data of type:
		struct {
			Tokens          []common.Address
			Balances        []*big.Int
			LastChangeBlock *big.Int
		}
	********************************************/
	poolTokens, err := bb.vaultContract.BalancerVaultCaller.GetPoolTokens(opts, bb.config.PoolID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pool tokens and balances from bex, err: %w", err)
	}

	var balanceData = make(map[string]*big.Int)
	for i, tokenAddress := range poolTokens.Tokens {
		token := strings.ToLower(tokenAddress.Hex())

		//verify this is always sound for all pool types
		if token == strings.ToLower(bb.poolAddress.Hex()) {
			// ignore when some of the LP token is locked in pool itself
			// this is why should use actualSupply instead of totalSupply
			continue
		}

		balance := poolTokens.Balances[i]
		balanceData[token] = balance
	}

	return balanceData, nil
}
