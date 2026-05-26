package protocols

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"

	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/infrared-dao/protocols/multicall3"
)

// Golden: Bend batch path == direct math for the same inputs.
// Single-asset vault wrapper (totalAssets * assetPrice / totalSupply).
func TestBend_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()

	assetPrice, _ := decimal.NewFromString("1.25")
	b := &BendLPPriceProvider{
		address: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		logger:  zerolog.Nop(),
		config:  &BendConfig{Asset: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", LPTDecimals: 18},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: assetPrice},
		},
	}

	totalSupply := new(big.Int).Mul(big.NewInt(80), bigPow10(18))
	totalAssets := new(big.Int).Mul(big.NewInt(100), bigPow10(18))
	// $1.25 * 100 / 80 = $1.5625

	wantDirect, _ := b.computeLPPriceFromReads(totalSupply, totalAssets)

	abi, _ := sc.BendVaultMetaData.GetAbi()
	encTS, _ := abi.Methods["totalSupply"].Outputs.Pack(totalSupply)
	encTA, _ := abi.Methods["totalAssets"].Outputs.Pack(totalAssets)

	gotBatch, err := b.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTS},
		{Success: true, ReturnData: encTA},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !gotBatch.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s", gotBatch, wantDirect)
	}
	want, _ := decimal.NewFromString("1.5625")
	if !gotBatch.Round(4).Equal(want) {
		t.Errorf("computed %s, want %s", gotBatch.Round(4), want)
	}
}
