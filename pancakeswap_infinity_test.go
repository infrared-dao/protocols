package protocols

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestPancakeSwapInfinity_BatchablePriceProviderReturnsZero(t *testing.T) {
	t.Parallel()
	p := &PancakeSwapInfinityLPPriceProvider{}
	reads, err := p.PriceReads()
	if err != nil {
		t.Fatalf("PriceReads: %v", err)
	}
	if len(reads) != 0 {
		t.Fatalf("want 0 reads, got %d", len(reads))
	}
	got, err := p.ComputePrice(nil)
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !got.Equal(decimal.Zero) {
		t.Errorf("CLMM pool ComputePrice = %s, want 0", got)
	}
}
