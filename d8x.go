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
	"github.com/shopspring/decimal"
)

const (
	ONE_64x64   = "18446744073709551616" // 1 << 64
	DECIMALS_18 = 1e18
)

type D8xConfig struct {
	PoolId      uint8          `json:"poolId"`
	OrclPoolTkn common.Address `json:"orclPoolTkn"`
}

type D8xLPPriceProvider struct {
	poolMngr    common.Address
	poolTknDec  uint8
	orclDec     uint8
	configBytes []byte
	config      D8xConfig
	orclCtrct   *sc.AggregatorV3
	d8xCtrct    *sc.D8x
}

// NewD8xLPPriceProvider returns a new instance of D8XLPPriceProvider with the assigned config
// and the D8X contract address
func NewD8xLPPriceProvider(d8xMngr common.Address, config []byte) *D8xLPPriceProvider {
	d := &D8xLPPriceProvider{
		poolMngr:    d8xMngr,
		configBytes: config,
	}
	return d
}

// GetConfig returns the configuration for the given parameters for D8X protocol.
// PoolId is the LP pool identification, orclPoolTknAddr is the chainlink-type
// oracle address for the pool token price
func (d8x *D8xLPPriceProvider) GetConfig(poolId uint8, orclPoolTknAddr string) ([]byte, error) {
	if !common.IsHexAddress(orclPoolTknAddr) {
		return nil, fmt.Errorf("invalid smart contract address, '%s'", orclPoolTknAddr)
	}
	c := D8xConfig{
		PoolId:      poolId,
		OrclPoolTkn: common.HexToAddress(orclPoolTknAddr),
	}
	return json.Marshal(c)
}

// Initialize requires d8x.configBytes to be set and initializes the price provider
// by querying onchain data
func (d8x *D8xLPPriceProvider) Initialize(ctx context.Context, client *ethclient.Client) error {
	// extract config
	err := json.Unmarshal(d8x.configBytes, &d8x.config)
	if err != nil {
		return fmt.Errorf("unable to deserialize config: %v", err)
	}
	// create d8x manager instance
	d8x.d8xCtrct, err = sc.NewD8x(d8x.poolMngr, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate d8x contract: %v", err)
	}
	opts := &bind.CallOpts{
		Context: ctx,
	}
	lp, err := d8x.d8xCtrct.GetLiquidityPool(opts, d8x.config.PoolId)
	if err != nil {
		return fmt.Errorf("unable to extract liquidity pool data: %v", err)
	}
	// get pool token decimals
	poolTkn, err := sc.NewERC20(lp.MarginTokenAddress, client)
	if err != nil {
		return fmt.Errorf("unable to instantiate ERC20 from margin token: %v", err)
	}
	d8x.poolTknDec, err = poolTkn.Decimals(opts)
	if err != nil {
		return fmt.Errorf("cannot get pool token decimals: %v", err)
	}
	// get oracle of pool token and decimals
	d8x.orclCtrct, err = sc.NewAggregatorV3(d8x.config.OrclPoolTkn, client)
	if err != nil {
		return fmt.Errorf("unable to instantiate pool token oracle: %v", err)
	}
	d8x.orclDec, err = d8x.orclCtrct.Decimals(opts)
	if err != nil {
		return fmt.Errorf("cannot get oracle decimals: %v", err)
	}
	return nil
}

// TVL returns the Total Value Locked in the pool in USD.
func (d8x *D8xLPPriceProvider) TVL(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Context: ctx,
	}
	// get lp data
	lp, err := d8x.d8xCtrct.GetLiquidityPool(opts, d8x.config.PoolId)
	if err != nil {
		return "", fmt.Errorf("unable to extract liquidity pool data for TVL: %v", err)
	}
	tvlBase, err := ABDKToDecimal(lp.FPnLparticipantsCashCC)
	if err != nil {
		return "", fmt.Errorf("could not convert lp tokens to decimal: %v", err)
	}
	// get price from oracle
	oracleRes, err := d8x.orclCtrct.LatestRoundData(opts)
	if err != nil {
		return "", fmt.Errorf("failed to call latestRoundData: %v", err)
	}
	// get conversion rate from pool token to USD
	px := NormalizeAmount(oracleRes.Answer, uint(d8x.orclDec))
	// convert from base amount (e.g. some stablecoin) to USD
	tvl := tvlBase.Mul(px)
	return tvl.StringFixed(roundingDecimals), nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD
func (d8x *D8xLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	// we call getShareTokenPriceD18(uint8 _poolId) which
	// returns the price of dbUSD in bUSD decimal 18 format
	opts := &bind.CallOpts{
		Context: ctx,
	}
	// get the share token price in collateral currency
	px18, err := d8x.d8xCtrct.GetShareTokenPriceD18(opts, d8x.config.PoolId)
	if err != nil {
		return "", fmt.Errorf("unable to get share token price: %v", err)
	}
	pxCC := NormalizeAmount(px18, 18)
	// we call the oracle to get the collateral price in USD
	oracleRes, err := d8x.orclCtrct.LatestRoundData(opts)
	if err != nil {
		return "", fmt.Errorf("failed to call latestRoundData: %v", err)
	}
	px := NormalizeAmount(oracleRes.Answer, uint(d8x.orclDec))
	pxUSD := pxCC.Mul(px)
	return pxUSD.StringFixed(roundingDecimals), nil
}

// ABDKToDecimal converts an ABDK fixed point 64.64 number to
// a decimal number
func ABDKToDecimal(xIn *big.Int) (decimal.Decimal, error) {
	if xIn.Cmp(big.NewInt(0)) == 0 {
		return decimal.Zero, nil
	}
	x := new(big.Int).Set(xIn)
	sgn := decimal.NewFromInt(int64(x.Sign()))
	x.Abs(x)

	var xInt, xRemainder, xDec big.Int
	one64x64 := new(big.Int)
	one64x64.SetString(ONE_64x64, 10)

	xInt.Div(x, one64x64)
	xRemainder.Mod(x, one64x64)
	xDec.Mul(&xRemainder, decimal.NewFromInt(DECIMALS_18).BigInt())
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
