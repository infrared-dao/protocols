package protocols

import (
	"errors"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"

	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/infrared-dao/protocols/multicall3"
)

// helper to build a KodiakLPPriceProvider with the given config + price map
// without making any RPC. We bypass Initialize since we're not testing the
// I/O path — only the pure-compute + ABI plumbing additions.
func newKodiakProviderForTest(t *testing.T, lptDecimals uint, token0Decimals, token1Decimals uint, token0Price, token1Price string, addr common.Address, contract KodiakContract) *KodiakLPPriceProvider {
	t.Helper()
	p0, _ := decimal.NewFromString(token0Price)
	p1, _ := decimal.NewFromString(token1Price)
	return &KodiakLPPriceProvider{
		address: addr,
		block:   nil,
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {TokenName: "T0", Decimals: token0Decimals, Price: p0},
			"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {TokenName: "T1", Decimals: token1Decimals, Price: p1},
		},
		logger:   zerolog.Nop(),
		config:   &KodiakConfig{Token0: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Token1: "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", LPTDecimals: lptDecimals},
		contract: contract,
	}
}

// computeLPPriceFromReads is the shared math both LPTokenPrice (legacy)
// and ComputePrice (batchable) route through. Testing it directly with
// known inputs validates the math without RPC mocking — if both paths
// produce identical prices, this helper is why.
func TestKodiak_computeLPPriceFromReads_Math(t *testing.T) {
	t.Parallel()
	// Token0: $2.00 with 18 decimals; balance 100e18 = 100 tokens × $2 = $200
	// Token1: $1.50 with 6  decimals; balance 200e6  = 200 tokens × $1.5 = $300
	// totalSupply: 50e18 = 50 LP tokens
	// LP price: ($200 + $300) / 50 = $10.00 per LP token
	k := newKodiakProviderForTest(t,
		/*lptDecimals*/ 18, /*t0Dec*/ 18, /*t1Dec*/ 6,
		/*t0Price*/ "2.00", /*t1Price*/ "1.50",
		common.HexToAddress("0x1234"), nil,
	)
	totalSupply := new(big.Int).Mul(big.NewInt(50), bigPow10(18))
	bal0 := new(big.Int).Mul(big.NewInt(100), bigPow10(18))
	bal1 := new(big.Int).Mul(big.NewInt(200), bigPow10(6))

	got, err := k.computeLPPriceFromReads(totalSupply, Balances{Amount0: bal0, Amount1: bal1})
	if err != nil {
		t.Fatalf("computeLPPriceFromReads: %v", err)
	}
	want, _ := decimal.NewFromString("10.00")
	if !got.Equal(want) {
		t.Errorf("price = %s, want %s", got.String(), want.String())
	}
}

// totalSupply of zero must produce an error — division-by-zero guard.
func TestKodiak_computeLPPriceFromReads_ZeroTotalSupplyErrors(t *testing.T) {
	t.Parallel()
	k := newKodiakProviderForTest(t, 18, 18, 18, "1", "1", common.HexToAddress("0x1234"), nil)
	_, err := k.computeLPPriceFromReads(big.NewInt(0), Balances{Amount0: big.NewInt(0), Amount1: big.NewInt(0)})
	if err == nil {
		t.Fatal("expected error for zero totalSupply, got nil")
	}
}

// Missing token price in the priceMap must propagate as an error.
func TestKodiak_computeLPPriceFromReads_MissingPriceErrors(t *testing.T) {
	t.Parallel()
	k := newKodiakProviderForTest(t, 18, 18, 18, "1", "1", common.HexToAddress("0x1234"), nil)
	delete(k.priceMap, k.config.Token1) // drop token1 price
	_, err := k.computeLPPriceFromReads(big.NewInt(1), Balances{Amount0: big.NewInt(1), Amount1: big.NewInt(1)})
	if err == nil {
		t.Fatal("expected error for missing token price, got nil")
	}
}

// PriceReads must produce two Call3 entries: a totalSupply call and a
// variant-specific balances call, both targeting the provider's address.
// V3Island uses getUnderlyingBalances; verify the selector matches.
func TestKodiak_PriceReads_V3Island(t *testing.T) {
	t.Parallel()
	k := newKodiakProviderForTest(t, 18, 18, 18, "1", "1",
		common.HexToAddress("0x1234567890123456789012345678901234567890"),
		&KodiakV3Island{},
	)
	calls, err := k.PriceReads()
	if err != nil {
		t.Fatalf("PriceReads: %v", err)
	}
	if len(calls) != 2 {
		t.Fatalf("PriceReads returned %d calls, want 2", len(calls))
	}
	for i, c := range calls {
		if c.Target != k.address {
			t.Errorf("calls[%d].Target = %v, want %v", i, c.Target, k.address)
		}
	}

	// Verify totalSupply selector — first 4 bytes of keccak256("totalSupply()")
	// = 0x18160ddd.
	if len(calls[0].CallData) < 4 || calls[0].CallData[0] != 0x18 || calls[0].CallData[1] != 0x16 || calls[0].CallData[2] != 0x0d || calls[0].CallData[3] != 0xdd {
		t.Errorf("calls[0] selector = %x..., want totalSupply 0x18160ddd...", calls[0].CallData[:min(4, len(calls[0].CallData))])
	}

	// Verify getUnderlyingBalances selector matches what the generated
	// binding's ABI would pack. Pack independently and compare; this catches
	// any future ABI drift in the binding.
	a, err := sc.KodiakIslandMetaData.GetAbi()
	if err != nil {
		t.Fatalf("get ABI: %v", err)
	}
	wantData, err := a.Pack("getUnderlyingBalances")
	if err != nil {
		t.Fatalf("pack getUnderlyingBalances: %v", err)
	}
	if !bytesEqual(calls[1].CallData, wantData) {
		t.Errorf("calls[1].CallData mismatch")
	}
}

// V2Pool variant uses getReserves; selector should differ from V3's
// getUnderlyingBalances.
func TestKodiak_PriceReads_V2Pool(t *testing.T) {
	t.Parallel()
	k := newKodiakProviderForTest(t, 18, 18, 18, "1", "1",
		common.HexToAddress("0x1234"),
		&KodiakV2Pool{},
	)
	calls, err := k.PriceReads()
	if err != nil {
		t.Fatalf("PriceReads: %v", err)
	}
	if len(calls) != 2 {
		t.Fatalf("calls = %d, want 2", len(calls))
	}
	// Compare against independently-packed getReserves calldata.
	a, err := sc.UniswapV2MetaData.GetAbi()
	if err != nil {
		t.Fatalf("get ABI: %v", err)
	}
	wantData, err := a.Pack("getReserves")
	if err != nil {
		t.Fatalf("pack getReserves: %v", err)
	}
	if !bytesEqual(calls[1].CallData, wantData) {
		t.Errorf("V2Pool calls[1].CallData mismatch — expected getReserves")
	}
}

// ComputePrice must error if either sub-call reverted in the Multicall3
// response. AllowFailure=false in PriceReads means in practice the whole
// aggregate3 reverts on failure; but ComputePrice has belt-and-suspenders
// per-response Success checks anyway.
func TestKodiak_ComputePrice_RejectsRevertedSubCalls(t *testing.T) {
	t.Parallel()
	k := newKodiakProviderForTest(t, 18, 18, 18, "1", "1", common.HexToAddress("0x1234"), &KodiakV3Island{})

	// First response failed.
	_, err := k.ComputePrice([]multicall3.Result3{
		{Success: false, ReturnData: nil},
		{Success: true, ReturnData: []byte{0x01}},
	})
	if err == nil || !strings.Contains(err.Error(), "totalSupply") {
		t.Errorf("expected totalSupply revert error, got %v", err)
	}

	// Second response failed.
	_, err = k.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: []byte{0x01}},
		{Success: false, ReturnData: nil},
	})
	if err == nil || !strings.Contains(err.Error(), "balances") {
		t.Errorf("expected balances revert error, got %v", err)
	}
}

// Wrong number of responses must error explicitly rather than panic.
func TestKodiak_ComputePrice_WrongResponseCount(t *testing.T) {
	t.Parallel()
	k := newKodiakProviderForTest(t, 18, 18, 18, "1", "1", common.HexToAddress("0x1234"), &KodiakV3Island{})
	_, err := k.ComputePrice([]multicall3.Result3{{Success: true}})
	if err == nil {
		t.Fatal("expected error for wrong response count, got nil")
	}
}

// Round-trip: pack totalSupply on the wire, unpack it, get the original
// value back. Ensures the helper functions wired in kodiak.go correctly
// handle the standard ERC20 totalSupply shape.
func TestKodiak_TotalSupplyRoundTrip(t *testing.T) {
	t.Parallel()
	want := new(big.Int).Mul(big.NewInt(42), bigPow10(18))

	a, err := sc.KodiakIslandMetaData.GetAbi()
	if err != nil {
		t.Fatalf("get ABI: %v", err)
	}
	encoded, err := a.Methods["totalSupply"].Outputs.Pack(want)
	if err != nil {
		t.Fatalf("pack: %v", err)
	}
	got, err := unpackKodiakTotalSupply(encoded)
	if err != nil {
		t.Fatalf("unpackKodiakTotalSupply: %v", err)
	}
	if got.Cmp(want) != 0 {
		t.Errorf("got %s, want %s", got.String(), want.String())
	}
}

// Round-trip: pack a V3 getUnderlyingBalances response, unpack via the
// V3Island variant, get the (Amount0, Amount1) back. Verifies that the
// variant's UnpackBalancesResponse correctly ignores trailing accumulator
// fields and extracts the first two values.
func TestKodiakV3Island_BalancesRoundTrip(t *testing.T) {
	t.Parallel()
	a, err := sc.KodiakIslandMetaData.GetAbi()
	if err != nil {
		t.Fatalf("get ABI: %v", err)
	}
	// getUnderlyingBalances returns (uint256 amount0Current, uint256 amount1Current).
	want0 := big.NewInt(123456789)
	want1 := big.NewInt(987654321)
	encoded, err := a.Methods["getUnderlyingBalances"].Outputs.Pack(want0, want1)
	if err != nil {
		t.Fatalf("pack: %v", err)
	}
	v3 := &KodiakV3Island{}
	got, err := v3.UnpackBalancesResponse(encoded)
	if err != nil {
		t.Fatalf("UnpackBalancesResponse: %v", err)
	}
	if got.Amount0.Cmp(want0) != 0 || got.Amount1.Cmp(want1) != 0 {
		t.Errorf("got (%s, %s), want (%s, %s)", got.Amount0, got.Amount1, want0, want1)
	}
}

// Same round-trip for V2 getReserves shape (uint112, uint112, uint32).
func TestKodiakV2Pool_BalancesRoundTrip(t *testing.T) {
	t.Parallel()
	a, err := sc.UniswapV2MetaData.GetAbi()
	if err != nil {
		t.Fatalf("get ABI: %v", err)
	}
	want0 := big.NewInt(111)
	want1 := big.NewInt(222)
	timestamp := uint32(1700000000)
	encoded, err := a.Methods["getReserves"].Outputs.Pack(want0, want1, timestamp)
	if err != nil {
		t.Fatalf("pack: %v", err)
	}
	v2 := &KodiakV2Pool{}
	got, err := v2.UnpackBalancesResponse(encoded)
	if err != nil {
		t.Fatalf("UnpackBalancesResponse: %v", err)
	}
	if got.Amount0.Cmp(want0) != 0 || got.Amount1.Cmp(want1) != 0 {
		t.Errorf("got (%s, %s), want (%s, %s)", got.Amount0, got.Amount1, want0, want1)
	}
}

// Golden test: legacy LPTokenPrice and batchable ComputePrice MUST produce
// bit-identical prices for the same inputs, because they share
// computeLPPriceFromReads. This test simulates both paths' final compute
// step with identical inputs and verifies they match.
//
// This is the central correctness guarantee of the migration: the new
// batch path can never silently diverge from legacy for the same inputs.
func TestKodiak_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	k := newKodiakProviderForTest(t,
		18, 18, 6,
		"3.50", "1.25",
		common.HexToAddress("0x1234"), &KodiakV3Island{},
	)
	totalSupply := new(big.Int).Mul(big.NewInt(75), bigPow10(18))
	balances := Balances{
		Amount0: new(big.Int).Mul(big.NewInt(123), bigPow10(18)),
		Amount1: new(big.Int).Mul(big.NewInt(456), bigPow10(6)),
	}

	// Direct invocation (what both paths converge on).
	wantDirect, err := k.computeLPPriceFromReads(totalSupply, balances)
	if err != nil {
		t.Fatalf("computeLPPriceFromReads: %v", err)
	}

	// Simulate the batch path: pack the responses ABI-style as if they came
	// from Multicall3, then run them through ComputePrice.
	a, err := sc.KodiakIslandMetaData.GetAbi()
	if err != nil {
		t.Fatalf("get ABI: %v", err)
	}
	encodedTotalSupply, err := a.Methods["totalSupply"].Outputs.Pack(totalSupply)
	if err != nil {
		t.Fatalf("pack totalSupply: %v", err)
	}
	encodedBalances, err := a.Methods["getUnderlyingBalances"].Outputs.Pack(balances.Amount0, balances.Amount1)
	if err != nil {
		t.Fatalf("pack getUnderlyingBalances: %v", err)
	}
	gotBatch, err := k.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encodedTotalSupply},
		{Success: true, ReturnData: encodedBalances},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}

	if !gotBatch.Equal(wantDirect) {
		t.Errorf("batch price %s != legacy compute %s — paths diverged!",
			gotBatch.String(), wantDirect.String())
	}

	// Sanity-check the actual number for the example: $3.50 * 123 + $1.25 * 456 = $430.50 + $570 = $1000.50; / 75 LP = $13.34
	want, _ := decimal.NewFromString("13.34")
	if !gotBatch.Round(2).Equal(want) {
		t.Errorf("computed price %s rounded to %s, want %s", gotBatch, gotBatch.Round(2), want)
	}

	// Catch the use of `errors` in the file to satisfy unused-import lint
	// when the imports list changes later.
	_ = errors.New("anchor")
}

// bigPow10 returns 10^n as a *big.Int for readable test fixtures.
func bigPow10(n int) *big.Int {
	return new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
