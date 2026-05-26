package protocols

import (
	"testing"
)

func TestPaddleFi_BatchablePriceProviderReturnsError(t *testing.T) {
	t.Parallel()
	p := &PaddleFiProvider{}
	reads, err := p.PriceReads()
	if err != nil {
		t.Fatalf("PriceReads: %v", err)
	}
	if len(reads) != 0 {
		t.Fatalf("want 0 reads, got %d", len(reads))
	}
	_, err = p.ComputePrice(nil)
	if err == nil {
		t.Fatal("expected ComputePrice to mirror LPTokenPrice's error, got nil")
	}
}
