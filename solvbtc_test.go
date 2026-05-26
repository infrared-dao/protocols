package protocols

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func TestSolv_ZeroReadPassThrough(t *testing.T) {
	t.Parallel()
	want, _ := decimal.NewFromString("65100.25")
	s := &SolvLPPriceProvider{
		address: common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		logger:  zerolog.Nop(),
		config:  &SolvConfig{Asset: "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", LPTDecimals: 18},
		priceMap: map[string]Price{
			"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {Decimals: 18, Price: want},
		},
	}
	reads, err := s.PriceReads()
	if err != nil {
		t.Fatalf("PriceReads: %v", err)
	}
	if len(reads) != 0 {
		t.Fatalf("want 0 reads, got %d", len(reads))
	}
	got, err := s.ComputePrice(nil)
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !got.Equal(want) {
		t.Errorf("got %s, want %s", got, want)
	}
}
