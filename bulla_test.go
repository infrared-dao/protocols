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

func TestBulla_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	p0, _ := decimal.NewFromString("3.00")
	p1, _ := decimal.NewFromString("1.00")
	b := &BullaLPPriceProvider{
		address: common.HexToAddress("0x1234"),
		logger:  zerolog.Nop(),
		config:  &BullaConfig{Token0: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Token1: "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", LPTDecimals: 18},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: p0},
			"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {Decimals: 18, Price: p1},
		},
	}
	totalSupply := new(big.Int).Mul(big.NewInt(60), bigPow10(18))
	t0 := new(big.Int).Mul(big.NewInt(40), bigPow10(18))
	t1 := new(big.Int).Mul(big.NewInt(120), bigPow10(18))
	wantDirect, _ := b.computeLPPriceFromReads(totalSupply, Balances{Amount0: t0, Amount1: t1})
	abi, _ := sc.BullaMetaData.GetAbi()
	encTS, _ := abi.Methods["totalSupply"].Outputs.Pack(totalSupply)
	encBal, _ := abi.Methods["getTotalAmounts"].Outputs.Pack(t0, t1)
	got, err := b.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTS},
		{Success: true, ReturnData: encBal},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !got.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s", got, wantDirect)
	}
}
