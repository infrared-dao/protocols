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

var _ Protocol = &D8xLPPriceProvider{}

const (
	One64x64                = "18446744073709551616" // 1 << 64
	Decimals18              = 1e18
	ShareTokenPriceDecimals = 18
)

type D8xConfig struct {
	PoolId         uint8          `json:"poolId"`
	PoolManager    common.Address `json:"pool_manager"`
	MarginToken    common.Address `json:"margin_token"`
	MarginDecimals uint8          `json:"margin_decimals"`
}

type D8xLPPriceProvider struct {
	address             common.Address
	block               *big.Int
	logger              zerolog.Logger
	configBytes         []byte
	config              D8xConfig
	marginTokenContract *sc.AggregatorV3
	poolManagerContract *sc.D8xPoolManager
}

// NewD8xLPPriceProvider returns a new instance of D8XLPPriceProvider with the assigned config
// and the D8X contract address
func NewD8xLPPriceProvider(
	address common.Address,
	block *big.Int,
	logger zerolog.Logger,
	config []byte,
) *D8xLPPriceProvider {
	d := &D8xLPPriceProvider{
		address:     address,
		block:       block,
		logger:      logger,
		configBytes: config,
	}
	return d
}

func (d8x *D8xLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	// extract config
	err := json.Unmarshal(d8x.configBytes, &d8x.config)
	if err != nil {
		return fmt.Errorf("unable to deserialize config: %v", err)
	}

	// create d8x pool manager contract instance
	d8x.poolManagerContract, err = sc.NewD8xPoolManager(d8x.config.PoolManager, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate d8x pool manager contract: %v", err)
	}

	// get oracle of pool token instance (ie. margin token contract)
	d8x.marginTokenContract, err = sc.NewAggregatorV3(d8x.config.MarginToken, client)
	if err != nil {
		return fmt.Errorf("unable to instantiate pool token oracle: %v", err)
	}

	return nil
}

// TVL returns the Total Value Locked in the pool in USD.
func (d8x *D8xLPPriceProvider) TVL(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: d8x.block,
	}

	// get current liquidity pool data
	lp, err := d8x.poolManagerContract.GetLiquidityPool(opts, d8x.config.PoolId)
	if err != nil {
		return "", fmt.Errorf("unable to extract liquidity pool data for TVL: %v", err)
	}

	tvlBase, err := ABDKToDecimal(lp.FPnLparticipantsCashCC)
	if err != nil {
		return "", fmt.Errorf("could not convert lp tokens to decimal: %v", err)
	}

	// get price from oracle
	oracleRes, err := d8x.marginTokenContract.LatestRoundData(opts)
	if err != nil {
		return "", fmt.Errorf("failed to call latestRoundData: %v", err)
	}

	// get conversion rate from pool token to USD
	px := NormalizeAmount(oracleRes.Answer, uint(d8x.config.MarginDecimals))

	// convert from base amount (e.g. some stablecoin) to USD
	tvl := tvlBase.Mul(px)
	return tvl.StringFixed(roundingDecimals), nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD
func (d8x *D8xLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// we call getShareTokenPriceD18(uint8 _poolId) which
	// returns the price of dbUSD in bUSD decimal 18 format
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: d8x.block,
	}

	// get the share token price in collateral currency
	px18, err := d8x.poolManagerContract.GetShareTokenPriceD18(opts, d8x.config.PoolId)
	if err != nil {
		return "", fmt.Errorf("unable to get share token price: %v", err)
	}
	pxCC := NormalizeAmount(px18, ShareTokenPriceDecimals)

	// we call the oracle to get the collateral price in USD
	oracleRes, err := d8x.marginTokenContract.LatestRoundData(opts)
	if err != nil {
		return "", fmt.Errorf("failed to call latestRoundData: %v", err)
	}
	px := NormalizeAmount(oracleRes.Answer, uint(d8x.config.MarginDecimals))
	pxUSD := pxCC.Mul(px)

	return pxUSD.StringFixed(roundingDecimals), nil
}

func (d8x *D8xLPPriceProvider) GetConfig(ctx context.Context, address string, ethClient *ethclient.Client) ([]byte, error) {
	var err error
	if !common.IsHexAddress(address) {
		err = fmt.Errorf("invalid smart contract address, '%s'", address)
		return nil, err
	}

	stContract, err := sc.NewD8xShareToken(common.HexToAddress(address), ethClient)
	if err != nil {
		err = fmt.Errorf("failed to instantiate D8x share token smart contract, %v", err)
		return nil, err
	}

	d8xc := &D8xConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	// 1. Get the poolID and pool manager address from the share token

	poolID, err := stContract.PoolId(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch poolID, %v", err)
		return nil, err
	}
	d8xc.PoolId = poolID

	poolManager, err := stContract.Owner(opts)
	if err != nil {
		err = fmt.Errorf("failed to fetch pool manager address (owner), %v", err)
		return nil, err
	}
	d8xc.PoolManager = poolManager

	// 2. Get the liquidity pool data from the pool manager and poolID
	//    Get margin token address and its decimals from liquidity pool

	pmContract, err := sc.NewD8xPoolManager(poolManager, ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate d8x pool manager contract: %v", err)
	}

	lp, err := pmContract.GetLiquidityPool(opts, poolID)
	if err != nil {
		return nil, fmt.Errorf("unable to extract liquidity pool data: %v", err)
	}
	d8xc.MarginToken = lp.MarginTokenAddress
	d8xc.MarginDecimals = lp.MarginTokenDecimals

	body, err := json.Marshal(d8xc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (d8x *D8xLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	d8x.block = block
	// prices param can be ignored since d8x uses onchain oracles
}

// ABDKToDecimal converts an ABDK fixed point 64.64 number to deicmal.Decimal
func ABDKToDecimal(xIn *big.Int) (decimal.Decimal, error) {
	if xIn.Cmp(big.NewInt(0)) == 0 {
		return decimal.Zero, nil
	}

	x := new(big.Int).Set(xIn)
	sgn := decimal.NewFromInt(int64(x.Sign()))
	x.Abs(x)

	var xInt, xRemainder, xDec big.Int
	one64x64 := new(big.Int)
	one64x64.SetString(One64x64, 10)

	xInt.Div(x, one64x64)
	xRemainder.Mod(x, one64x64)
	xDec.Mul(&xRemainder, decimal.NewFromInt(Decimals18).BigInt())
	xDec.Div(&xDec, one64x64)

	digits := 18 - len(xDec.String())
	pad := strings.Repeat("0", digits)

	numberStr := xInt.String() + "." + pad + xDec.String()

	res, err := decimal.NewFromString(numberStr)
	if err != nil {
		return decimal.Zero, err
	}

	return res.Mul(sgn), err
}
