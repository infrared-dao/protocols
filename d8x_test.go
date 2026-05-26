package protocols

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"

	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/infrared-dao/protocols/multicall3"
)

func TestD8x_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	d8x := &D8xLPPriceProvider{
		address: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		logger:  zerolog.Nop(),
		config: D8xConfig{
			PoolId:         1,
			PoolManager:    common.HexToAddress("0x2222222222222222222222222222222222222222"),
			MarginToken:    common.HexToAddress("0x3333333333333333333333333333333333333333"),
			MarginDecimals: 8,
		},
	}
	// share price = 1.05e18, oracle answer = 1e8 (1.0 USD)
	px18 := new(big.Int).Mul(big.NewInt(105), bigPow10(16))
	answer := big.NewInt(100000000)
	wantDirect := d8x.computeLPPriceFromReads(px18, answer)
	pmABI, _ := sc.D8xPoolManagerMetaData.GetAbi()
	encSP, _ := pmABI.Methods["getShareTokenPriceD18"].Outputs.Pack(px18)
	oracleABI, _ := sc.AggregatorV3MetaData.GetAbi()
	encLRD, _ := oracleABI.Methods["latestRoundData"].Outputs.Pack(
		big.NewInt(1),
		answer,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(1),
	)
	got, err := d8x.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encSP},
		{Success: true, ReturnData: encLRD},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !got.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s", got, wantDirect)
	}
}
