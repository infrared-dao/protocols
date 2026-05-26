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

// Golden: BrownFi 2-token pool, batch path == direct math.
func TestBrownFi_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	p0, _ := decimal.NewFromString("1.00")
	p1, _ := decimal.NewFromString("2.00")
	w := &BrownFiLPPriceProvider{
		address: common.HexToAddress("0x1234"),
		logger:  zerolog.Nop(),
		config:  &BrownFiConfig{Token0: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Token1: "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", LPTDecimals: 18},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: p0},
			"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {Decimals: 18, Price: p1},
		},
	}
	totalSupply := new(big.Int).Mul(big.NewInt(100), bigPow10(18))
	r0 := new(big.Int).Mul(big.NewInt(50), bigPow10(18))
	r1 := new(big.Int).Mul(big.NewInt(50), bigPow10(18))
	wantDirect, _ := w.computeLPPriceFromReads(totalSupply, Balances{Amount0: r0, Amount1: r1})

	abi, _ := sc.BrownFiPoolMetaData.GetAbi()
	encTS, _ := abi.Methods["totalSupply"].Outputs.Pack(totalSupply)
	// getReserves returns (uint112, uint112, uint32) — must pack all 3.
	encBal, _ := abi.Methods["getReserves"].Outputs.Pack(r0, r1, uint32(1700000000))

	gotBatch, err := w.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTS},
		{Success: true, ReturnData: encBal},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !gotBatch.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s", gotBatch, wantDirect)
	}
}
