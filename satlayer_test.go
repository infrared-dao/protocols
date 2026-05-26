package protocols

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"

	"github.com/infrared-dao/protocols/multicall3"
)

func TestSatLayer_ZeroReadPassThrough(t *testing.T) {
	t.Parallel()
	want, _ := decimal.NewFromString("65000.50")
	s := &SatLayerLPPriceProvider{
		address: common.HexToAddress("0x9999999999999999999999999999999999999999"),
		logger:  zerolog.Nop(),
		config:  &SatLayerConfig{Asset: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", LPTDecimals: 18},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: want},
		},
	}
	reads, err := s.PriceReads()
	if err != nil {
		t.Fatalf("PriceReads: %v", err)
	}
	if len(reads) != 0 {
		t.Fatalf("want 0 reads, got %d", len(reads))
	}
	got, err := s.ComputePrice([]multicall3.Result3{})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !got.Equal(want) {
		t.Errorf("got %s, want %s", got, want)
	}
}
