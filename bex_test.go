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

// newBexProviderForTest builds a BexLPPriceProvider that can compute prices
// from given inputs without making any RPC. Bypasses Initialize since these
// tests target the new BatchablePriceProvider methods + the shared math
// helper, not the I/O wiring.
func newBexProviderForTest(t *testing.T, lptDecimals uint, priceMap map[string]Price) *BexLPPriceProvider {
	t.Helper()
	return &BexLPPriceProvider{
		vaultAddress: common.HexToAddress("0x4Be03f781C497A489E3cB0287833452cA9B9E80B"),
		poolAddress:  common.HexToAddress("0x1111111111111111111111111111111111111111"),
		block:        nil,
		priceMap:     priceMap,
		logger:       zerolog.Nop(),
		config: &BexPoolConfig{
			PoolID:      [32]byte{0xaa, 0xbb, 0xcc},
			LPTDecimals: lptDecimals,
		},
	}
}

// Shared math: two-token pool, well-known prices/balances, verify the
// numeric LP price matches the expected calculation.
func TestBex_computeLPPriceFromReads_Math(t *testing.T) {
	t.Parallel()
	// Token A: $4.00 with 18 decimals, balance 50e18 → $200
	// Token B: $0.50 with  6 decimals, balance 1000e6 → $500
	// totalSupply: 100e18 (100 LP)
	// LP price: $700 / 100 = $7.00
	tokenA, _ := decimal.NewFromString("4.00")
	tokenB, _ := decimal.NewFromString("0.50")
	priceMap := map[string]Price{
		"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {TokenName: "A", Decimals: 18, Price: tokenA},
		"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {TokenName: "B", Decimals: 6, Price: tokenB},
	}
	b := newBexProviderForTest(t, 18, priceMap)

	totalSupply := new(big.Int).Mul(big.NewInt(100), bigPow10(18))
	balances := map[string]*big.Int{
		"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": new(big.Int).Mul(big.NewInt(50), bigPow10(18)),
		"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": new(big.Int).Mul(big.NewInt(1000), bigPow10(6)),
	}

	got, err := b.computeLPPriceFromReads(totalSupply, balances)
	if err != nil {
		t.Fatalf("computeLPPriceFromReads: %v", err)
	}
	want, _ := decimal.NewFromString("7.00")
	if !got.Equal(want) {
		t.Errorf("price = %s, want %s", got.String(), want.String())
	}
}

// PriceReads must produce two Call3 entries targeting different contracts:
// one for getActualSupply against the pool, one for getPoolTokens against
// the vault. Verify both selectors match what their respective ABIs pack
// independently.
func TestBex_PriceReads(t *testing.T) {
	t.Parallel()
	b := newBexProviderForTest(t, 18, map[string]Price{})
	calls, err := b.PriceReads()
	if err != nil {
		t.Fatalf("PriceReads: %v", err)
	}
	if len(calls) != 2 {
		t.Fatalf("calls = %d, want 2", len(calls))
	}
	if calls[0].Target != b.poolAddress {
		t.Errorf("calls[0].Target = %v, want pool %v", calls[0].Target, b.poolAddress)
	}
	if calls[1].Target != b.vaultAddress {
		t.Errorf("calls[1].Target = %v, want vault %v", calls[1].Target, b.vaultAddress)
	}

	// Compare against independently-packed calldata.
	poolABI, err := sc.BalancerBasePoolMetaData.GetAbi()
	if err != nil {
		t.Fatalf("get pool ABI: %v", err)
	}
	wantTotalSupply, err := poolABI.Pack("getActualSupply")
	if err != nil {
		t.Fatalf("pack getActualSupply: %v", err)
	}
	if !bytesEqual(calls[0].CallData, wantTotalSupply) {
		t.Errorf("calls[0].CallData mismatch with getActualSupply")
	}

	vaultABI, err := sc.BalancerVaultMetaData.GetAbi()
	if err != nil {
		t.Fatalf("get vault ABI: %v", err)
	}
	wantPoolTokens, err := vaultABI.Pack("getPoolTokens", b.config.PoolID)
	if err != nil {
		t.Fatalf("pack getPoolTokens: %v", err)
	}
	if !bytesEqual(calls[1].CallData, wantPoolTokens) {
		t.Errorf("calls[1].CallData mismatch with getPoolTokens(poolID)")
	}
}

// Reverted sub-calls must produce explicit errors per their position.
func TestBex_ComputePrice_RejectsRevertedSubCalls(t *testing.T) {
	t.Parallel()
	b := newBexProviderForTest(t, 18, map[string]Price{})
	if _, err := b.ComputePrice([]multicall3.Result3{
		{Success: false, ReturnData: nil},
		{Success: true, ReturnData: []byte{0x01}},
	}); err == nil {
		t.Errorf("expected error for reverted getActualSupply, got nil")
	}
	if _, err := b.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: []byte{0x01}},
		{Success: false, ReturnData: nil},
	}); err == nil {
		t.Errorf("expected error for reverted getPoolTokens, got nil")
	}
}

// Wrong response count guards against caller bugs.
func TestBex_ComputePrice_WrongResponseCount(t *testing.T) {
	t.Parallel()
	b := newBexProviderForTest(t, 18, map[string]Price{})
	if _, err := b.ComputePrice([]multicall3.Result3{{Success: true}}); err == nil {
		t.Fatal("expected error for 1 response, got nil")
	}
}

// Round-trip: pack BEX ABI responses, run them through ComputePrice,
// verify the price matches what computeLPPriceFromReads would have
// produced for the same inputs. This is the BEX equivalent of Kodiak's
// golden-master test.
func TestBex_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()

	tokenAAddr := common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	tokenBAddr := common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")

	tokenA, _ := decimal.NewFromString("2.50")
	tokenB, _ := decimal.NewFromString("1.00")
	priceMap := map[string]Price{
		"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {TokenName: "A", Decimals: 18, Price: tokenA},
		"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {TokenName: "B", Decimals: 18, Price: tokenB},
	}
	b := newBexProviderForTest(t, 18, priceMap)

	totalSupply := new(big.Int).Mul(big.NewInt(50), bigPow10(18))
	balA := new(big.Int).Mul(big.NewInt(20), bigPow10(18))
	balB := new(big.Int).Mul(big.NewInt(75), bigPow10(18))

	// Direct math: $50 + $75 = $125 / 50 LP = $2.50/LP
	wantDirect, err := b.computeLPPriceFromReads(totalSupply, map[string]*big.Int{
		"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": balA,
		"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": balB,
	})
	if err != nil {
		t.Fatalf("computeLPPriceFromReads: %v", err)
	}

	// Simulate batch path: pack ABI responses as if returned by Multicall3.
	poolABI, _ := sc.BalancerBasePoolMetaData.GetAbi()
	encTotalSupply, err := poolABI.Methods["getActualSupply"].Outputs.Pack(totalSupply)
	if err != nil {
		t.Fatalf("pack totalSupply: %v", err)
	}
	vaultABI, _ := sc.BalancerVaultMetaData.GetAbi()
	encPoolTokens, err := vaultABI.Methods["getPoolTokens"].Outputs.Pack(
		[]common.Address{tokenAAddr, tokenBAddr},
		[]*big.Int{balA, balB},
		big.NewInt(0), // lastChangeBlock — unused by ComputePrice
	)
	if err != nil {
		t.Fatalf("pack getPoolTokens: %v", err)
	}

	gotBatch, err := b.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTotalSupply},
		{Success: true, ReturnData: encPoolTokens},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}

	if !gotBatch.Equal(wantDirect) {
		t.Errorf("batch %s != legacy %s — paths diverged!",
			gotBatch.String(), wantDirect.String())
	}

	want, _ := decimal.NewFromString("2.50")
	if !gotBatch.Round(2).Equal(want) {
		t.Errorf("computed price %s rounded to %s, want %s", gotBatch, gotBatch.Round(2), want)
	}
}

// The pool's own LP-token entry in getPoolTokens must be filtered out
// (BEX locks some LP tokens in the pool itself; getActualSupply already
// accounts for this, so the pool-token entry in balances must be ignored).
func TestBex_ComputePrice_FiltersPoolOwnToken(t *testing.T) {
	t.Parallel()

	tokenAAddr := common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	tokenA, _ := decimal.NewFromString("1.00")
	priceMap := map[string]Price{
		"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {TokenName: "A", Decimals: 18, Price: tokenA},
	}
	b := newBexProviderForTest(t, 18, priceMap)
	// b.poolAddress is 0x1111... — the pool token listed in getPoolTokens
	// should be filtered out by ComputePrice so the absence of a price for
	// it (we deliberately did NOT add it to priceMap) doesn't error out.

	totalSupply := new(big.Int).Mul(big.NewInt(10), bigPow10(18))
	balA := new(big.Int).Mul(big.NewInt(5), bigPow10(18))
	balPool := new(big.Int).Mul(big.NewInt(1000), bigPow10(18)) // ignored

	poolABI, _ := sc.BalancerBasePoolMetaData.GetAbi()
	encTotalSupply, _ := poolABI.Methods["getActualSupply"].Outputs.Pack(totalSupply)
	vaultABI, _ := sc.BalancerVaultMetaData.GetAbi()
	encPoolTokens, _ := vaultABI.Methods["getPoolTokens"].Outputs.Pack(
		[]common.Address{tokenAAddr, b.poolAddress},
		[]*big.Int{balA, balPool},
		big.NewInt(0),
	)

	got, err := b.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTotalSupply},
		{Success: true, ReturnData: encPoolTokens},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	// Only token A's balance counts: $1 × 5 = $5 / 10 LP = $0.50/LP
	want, _ := decimal.NewFromString("0.50")
	if !got.Equal(want) {
		t.Errorf("price = %s, want %s (pool-own token should have been filtered)", got, want)
	}
}
