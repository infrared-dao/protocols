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

// Golden: AquaBera batch path == direct math for the same inputs.
// Same shape as Kodiak (2-token weighted price) so the test is brief.
func TestAquaBera_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()

	p0, _ := decimal.NewFromString("5.00")
	p1, _ := decimal.NewFromString("2.00")
	a := &AquaBeraLPPriceProvider{
		address: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		logger:  zerolog.Nop(),
		config:  &AquaBeraConfig{Token0: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Token1: "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", LPTDecimals: 18},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: p0},
			"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {Decimals: 18, Price: p1},
		},
	}

	totalSupply := new(big.Int).Mul(big.NewInt(40), bigPow10(18))
	bal0 := new(big.Int).Mul(big.NewInt(10), bigPow10(18))
	bal1 := new(big.Int).Mul(big.NewInt(50), bigPow10(18))
	// $5*10 + $2*50 = $150 / 40 LP = $3.75/LP

	wantDirect, err := a.computeLPPriceFromReads(totalSupply, Balances{Amount0: bal0, Amount1: bal1})
	if err != nil {
		t.Fatalf("direct: %v", err)
	}

	abi, _ := sc.AquaBeraMetaData.GetAbi()
	encTS, _ := abi.Methods["totalSupply"].Outputs.Pack(totalSupply)
	encBal, _ := abi.Methods["getTotalAmounts"].Outputs.Pack(bal0, bal1)

	gotBatch, err := a.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTS},
		{Success: true, ReturnData: encBal},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !gotBatch.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s", gotBatch, wantDirect)
	}
	want, _ := decimal.NewFromString("3.75")
	if !gotBatch.Round(2).Equal(want) {
		t.Errorf("computed %s, want %s", gotBatch.Round(2), want)
	}
}
