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

func TestSteer_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	p0, _ := decimal.NewFromString("2.00")
	p1, _ := decimal.NewFromString("1.50")
	s := &SteerLPPriceProvider{
		address: common.HexToAddress("0xcccccccccccccccccccccccccccccccccccccccc"),
		logger:  zerolog.Nop(),
		config: &SteerConfig{
			Token0:      "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Token1:      "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			LPTDecimals: 18,
		},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: p0},
			"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {Decimals: 18, Price: p1},
		},
	}
	totalSupply := new(big.Int).Mul(big.NewInt(80), bigPow10(18))
	a0 := new(big.Int).Mul(big.NewInt(40), bigPow10(18))
	a1 := new(big.Int).Mul(big.NewInt(60), bigPow10(18))
	wantDirect, _ := s.computeLPPriceFromReads(totalSupply, a0, a1)
	abi, _ := sc.SteerPoolMetaData.GetAbi()
	encTS, _ := abi.Methods["totalSupply"].Outputs.Pack(totalSupply)
	encGTA, _ := abi.Methods["getTotalAmounts"].Outputs.Pack(a0, a1)
	got, err := s.ComputePrice([]multicall3.Result3{
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
