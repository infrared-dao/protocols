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

func TestWasabee_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	p0, _ := decimal.NewFromString("1.00")
	p1, _ := decimal.NewFromString("2.00")
	w := &WasabeeLPPriceProvider{
		address: common.HexToAddress("0xdddddddddddddddddddddddddddddddddddddddd"),
		logger:  zerolog.Nop(),
		config: &WasabeeConfig{
			Token0:      "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Token1:      "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			LPTDecimals: 18,
		},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: p0},
			"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {Decimals: 18, Price: p1},
		},
	}
	totalSupply := new(big.Int).Mul(big.NewInt(100), bigPow10(18))
	a0 := new(big.Int).Mul(big.NewInt(50), bigPow10(18))
	a1 := new(big.Int).Mul(big.NewInt(50), bigPow10(18))
	wantDirect, _ := w.computeLPPriceFromReads(totalSupply, a0, a1)
	abi, _ := sc.WasabeeVaultMetaData.GetAbi()
	encTS, _ := abi.Methods["totalSupply"].Outputs.Pack(totalSupply)
	encGTA, _ := abi.Methods["getTotalAmounts"].Outputs.Pack(a0, a1)
	got, err := w.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTS},
		{Success: true, ReturnData: encGTA},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !got.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s", got, wantDirect)
	}
}
