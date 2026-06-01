package protocols

import (
	"context"
	"errors"
	"math/big"
	"sync/atomic"
	"testing"
	"time"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"

	"github.com/infrared-dao/protocols/fetchers"
	"github.com/infrared-dao/protocols/multicall3"
)

// fakeMulticall captures the calls handed to Aggregate3 and returns canned
// per-call results. Used to exercise BatchLPTokenPrice without a real
// network or contract.
type fakeMulticall struct {
	gotCalls       []multicall3.Call3
	gotBlockNumber *big.Int
	respFn         func(calls []multicall3.Call3) ([]multicall3.Result3, error)
	dispatches     int
}

func (f *fakeMulticall) Aggregate3(_ context.Context, calls []multicall3.Call3, block *big.Int) ([]multicall3.Result3, error) {
	f.gotCalls = append(f.gotCalls, calls...)
	f.gotBlockNumber = block
	f.dispatches++
	return f.respFn(calls)
}

// fakeBatchable is a minimal BatchablePriceProvider for tests. It always
// returns the configured reads/price and never touches a network. The
// legacy LPTokenPrice path is also implemented so tests can verify the
// batched path is actually preferred over the legacy fallback.
type fakeBatchable struct {
	reads         []multicall3.Call3
	priceFromResp func(responses []multicall3.Result3) (decimal.Decimal, error)
	legacyPrice   string
	legacyErr     error
}

// Protocol interface stubs (Initialize, GetConfig, etc.) — minimal
// implementations because BatchLPTokenPrice doesn't call them for already-
// initialized providers.
func (f *fakeBatchable) GetConfig(context.Context, string, bind.ContractBackend) ([]byte, error) {
	return nil, nil
}
func (f *fakeBatchable) Initialize(context.Context, bind.ContractBackend, fetchers.HttpClient) error {
	return nil
}
func (f *fakeBatchable) LPTokenPrice(context.Context) (string, error) {
	return f.legacyPrice, f.legacyErr
}
func (f *fakeBatchable) TVL(context.Context) (string, error)                       { return "", nil }
func (f *fakeBatchable) TVLBreakdown(context.Context) (map[string]TokenTVL, error) { return nil, nil }
func (f *fakeBatchable) UpdateBlock(*big.Int, map[string]Price)                    {}

// BatchablePriceProvider extension:
func (f *fakeBatchable) PriceReads() ([]multicall3.Call3, error) { return f.reads, nil }
func (f *fakeBatchable) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	return f.priceFromResp(responses)
}

// fakeLegacy is a minimal Protocol that does NOT implement BatchablePriceProvider.
// Exercises the legacy fallback path in BatchLPTokenPrice.
type fakeLegacy struct {
	price string
	err   error
}

func (f *fakeLegacy) GetConfig(context.Context, string, bind.ContractBackend) ([]byte, error) {
	return nil, nil
}
func (f *fakeLegacy) Initialize(context.Context, bind.ContractBackend, fetchers.HttpClient) error {
	return nil
}
func (f *fakeLegacy) LPTokenPrice(context.Context) (string, error)              { return f.price, f.err }
func (f *fakeLegacy) TVL(context.Context) (string, error)                       { return "", nil }
func (f *fakeLegacy) TVLBreakdown(context.Context) (map[string]TokenTVL, error) { return nil, nil }
func (f *fakeLegacy) UpdateBlock(*big.Int, map[string]Price)                    {}

func TestBatchLPTokenPrice_RejectsNilMulticall(t *testing.T) {
	t.Parallel()
	_, err := BatchLPTokenPrice(context.Background(), nil, nil, nil, nil)
	if err == nil {
		t.Fatal("expected error for nil multicall, got nil")
	}
}

func TestBatchLPTokenPrice_RejectsNilProvider(t *testing.T) {
	t.Parallel()
	mc := &fakeMulticall{}
	_, err := BatchLPTokenPrice(context.Background(), mc, nil, nil, []BatchPriceQuery{
		{Provider: nil},
	})
	if err == nil {
		t.Fatal("expected error for nil Provider, got nil")
	}
}

// Two batchable providers with the same block must be merged into a single
// aggregate3 dispatch, and their per-provider PriceReads concatenated in
// order with results distributed back by provider.
func TestBatchLPTokenPrice_BatchableProvidersShareOneDispatch(t *testing.T) {
	t.Parallel()

	p1 := &fakeBatchable{
		reads: []multicall3.Call3{
			{Target: common.Address{0xA1}, CallData: []byte{0xa1}},
			{Target: common.Address{0xA2}, CallData: []byte{0xa2}},
		},
		priceFromResp: func(responses []multicall3.Result3) (decimal.Decimal, error) {
			if len(responses) != 2 {
				return decimal.Zero, errors.New("p1 wrong response count")
			}
			return decimal.NewFromInt(100), nil
		},
	}
	p2 := &fakeBatchable{
		reads: []multicall3.Call3{
			{Target: common.Address{0xB1}, CallData: []byte{0xb1}},
		},
		priceFromResp: func(responses []multicall3.Result3) (decimal.Decimal, error) {
			if len(responses) != 1 {
				return decimal.Zero, errors.New("p2 wrong response count")
			}
			return decimal.NewFromInt(200), nil
		},
	}

	mc := &fakeMulticall{
		respFn: func(calls []multicall3.Call3) ([]multicall3.Result3, error) {
			out := make([]multicall3.Result3, len(calls))
			for i := range calls {
				out[i] = multicall3.Result3{Success: true, ReturnData: []byte{byte(i)}}
			}
			return out, nil
		},
	}

	results, err := BatchLPTokenPrice(context.Background(), mc, nil, nil, []BatchPriceQuery{
		{Provider: p1}, {Provider: p2},
	})
	if err != nil {
		t.Fatalf("BatchLPTokenPrice: %v", err)
	}
	if mc.dispatches != 1 {
		t.Errorf("dispatches = %d, want 1 (one aggregate3 for two same-block batchables)", mc.dispatches)
	}
	if len(mc.gotCalls) != 3 {
		t.Errorf("aggregate3 received %d calls, want 3 (2 from p1 + 1 from p2)", len(mc.gotCalls))
	}
	if !results[0].Price.Equal(decimal.NewFromInt(100)) {
		t.Errorf("results[0].Price = %v, want 100", results[0].Price)
	}
	if !results[1].Price.Equal(decimal.NewFromInt(200)) {
		t.Errorf("results[1].Price = %v, want 200", results[1].Price)
	}
}

// Providers with different blocks must be partitioned: one aggregate3 per
// block. This is fundamental — Multicall3 reads atomically at one block.
func TestBatchLPTokenPrice_DifferentBlocksPartitioned(t *testing.T) {
	t.Parallel()

	p1 := &fakeBatchable{
		reads:         []multicall3.Call3{{Target: common.Address{1}, CallData: []byte{1}}},
		priceFromResp: func([]multicall3.Result3) (decimal.Decimal, error) { return decimal.NewFromInt(1), nil },
	}
	p2 := &fakeBatchable{
		reads:         []multicall3.Call3{{Target: common.Address{2}, CallData: []byte{2}}},
		priceFromResp: func([]multicall3.Result3) (decimal.Decimal, error) { return decimal.NewFromInt(2), nil },
	}

	mc := &fakeMulticall{
		respFn: func(calls []multicall3.Call3) ([]multicall3.Result3, error) {
			return []multicall3.Result3{{Success: true, ReturnData: []byte{0}}}, nil
		},
	}

	_, err := BatchLPTokenPrice(context.Background(), mc, nil, nil, []BatchPriceQuery{
		{Provider: p1, BlockNumber: big.NewInt(100)},
		{Provider: p2, BlockNumber: big.NewInt(200)},
	})
	if err != nil {
		t.Fatalf("BatchLPTokenPrice: %v", err)
	}
	if mc.dispatches != 2 {
		t.Errorf("dispatches = %d, want 2 (different blocks must be separate aggregate3s)", mc.dispatches)
	}
}

// Non-batchable providers fall back to LPTokenPrice via errgroup. Mix of
// batchable + legacy in one BatchLPTokenPrice invocation should produce
// correct prices for both shapes, with the batchable ones aggregated and
// the legacy ones invoked individually.
func TestBatchLPTokenPrice_LegacyFallbackForNonBatchable(t *testing.T) {
	t.Parallel()

	batched := &fakeBatchable{
		reads:         []multicall3.Call3{{Target: common.Address{1}, CallData: []byte{1}}},
		priceFromResp: func([]multicall3.Result3) (decimal.Decimal, error) { return decimal.NewFromInt(42), nil },
	}
	legacy := &fakeLegacy{price: "99.5"}

	mc := &fakeMulticall{
		respFn: func(calls []multicall3.Call3) ([]multicall3.Result3, error) {
			return []multicall3.Result3{{Success: true, ReturnData: []byte{0}}}, nil
		},
	}

	results, err := BatchLPTokenPrice(context.Background(), mc, nil, nil, []BatchPriceQuery{
		{Provider: batched}, {Provider: legacy},
	})
	if err != nil {
		t.Fatalf("BatchLPTokenPrice: %v", err)
	}
	if !results[0].Price.Equal(decimal.NewFromInt(42)) {
		t.Errorf("batched result = %v, want 42", results[0].Price)
	}
	wantLegacy, _ := decimal.NewFromString("99.5")
	if !results[1].Price.Equal(wantLegacy) {
		t.Errorf("legacy result = %v, want 99.5", results[1].Price)
	}
}

// A per-query error from the legacy path should be reported in
// results[i].Err without affecting other queries' results or the top-level
// return.
func TestBatchLPTokenPrice_LegacyErrorIsolated(t *testing.T) {
	t.Parallel()

	good := &fakeLegacy{price: "5"}
	bad := &fakeLegacy{err: errors.New("legacy failed")}

	mc := &fakeMulticall{}
	results, err := BatchLPTokenPrice(context.Background(), mc, nil, nil, []BatchPriceQuery{
		{Provider: good}, {Provider: bad},
	})
	if err != nil {
		t.Fatalf("BatchLPTokenPrice unexpected err: %v", err)
	}
	if results[0].Err != nil {
		t.Errorf("good query had err: %v", results[0].Err)
	}
	if results[1].Err == nil {
		t.Errorf("bad query expected err, got nil")
	}
}

// Top-level aggregate3 failure should mark every batched query in that
// partition with the error, but leave queries in other partitions
// (different block) and legacy queries unaffected.
func TestBatchLPTokenPrice_BatchedDispatchErrorIsolated(t *testing.T) {
	t.Parallel()

	p1 := &fakeBatchable{
		reads:         []multicall3.Call3{{Target: common.Address{1}, CallData: []byte{1}}},
		priceFromResp: func([]multicall3.Result3) (decimal.Decimal, error) { return decimal.NewFromInt(1), nil },
	}
	p2 := &fakeBatchable{
		reads:         []multicall3.Call3{{Target: common.Address{2}, CallData: []byte{2}}},
		priceFromResp: func([]multicall3.Result3) (decimal.Decimal, error) { return decimal.NewFromInt(2), nil },
	}

	mc := &fakeMulticall{
		respFn: func(calls []multicall3.Call3) ([]multicall3.Result3, error) {
			return nil, errors.New("multicall down")
		},
	}

	results, err := BatchLPTokenPrice(context.Background(), mc, nil, nil, []BatchPriceQuery{
		{Provider: p1}, {Provider: p2},
	})
	if err != nil {
		t.Fatalf("top-level err: %v", err)
	}
	for i, r := range results {
		if r.Err == nil {
			t.Errorf("results[%d] expected err from failed dispatch, got Price=%v", i, r.Price)
		}
	}
}

// A partition containing only zero-PriceReads providers (pure pricemap
// pass-throughs like satlayer, solvbtc) must still receive a ComputePrice
// invocation each: skipping ComputePrice would leak zero-value results.
func TestBatchLPTokenPrice_ZeroReadPartitionStillComputes(t *testing.T) {
	t.Parallel()

	p1 := &fakeBatchable{
		reads:         nil, // no on-chain reads
		priceFromResp: func([]multicall3.Result3) (decimal.Decimal, error) { return decimal.NewFromInt(7), nil },
	}
	p2 := &fakeBatchable{
		reads:         nil,
		priceFromResp: func([]multicall3.Result3) (decimal.Decimal, error) { return decimal.NewFromInt(11), nil },
	}

	mc := &fakeMulticall{
		respFn: func(calls []multicall3.Call3) ([]multicall3.Result3, error) {
			t.Errorf("Aggregate3 must not be called when all reads are empty")
			return nil, nil
		},
	}

	results, err := BatchLPTokenPrice(context.Background(), mc, nil, nil, []BatchPriceQuery{
		{Provider: p1}, {Provider: p2},
	})
	if err != nil {
		t.Fatalf("BatchLPTokenPrice: %v", err)
	}
	if mc.dispatches != 0 {
		t.Errorf("dispatches = %d, want 0 (no calls to aggregate)", mc.dispatches)
	}
	if !results[0].Price.Equal(decimal.NewFromInt(7)) {
		t.Errorf("results[0].Price = %v, want 7", results[0].Price)
	}
	if !results[1].Price.Equal(decimal.NewFromInt(11)) {
		t.Errorf("results[1].Price = %v, want 11", results[1].Price)
	}
}

// BatchOptions.LegacyParallelism must cap concurrent legacy goroutines.
// With LegacyParallelism=1, peak observed in-flight count for N>1 legacy
// providers must be exactly 1.
func TestBatchLPTokenPriceWithOptions_LegacyParallelism(t *testing.T) {
	t.Parallel()

	var inFlight, peak int64
	gate := make(chan struct{})

	slowImpl := func(_ context.Context) (string, error) {
		cur := atomic.AddInt64(&inFlight, 1)
		for {
			p := atomic.LoadInt64(&peak)
			if cur <= p || atomic.CompareAndSwapInt64(&peak, p, cur) {
				break
			}
		}
		<-gate
		atomic.AddInt64(&inFlight, -1)
		return "1", nil
	}

	const n = 4
	queries := make([]BatchPriceQuery, n)
	for i := range queries {
		queries[i] = BatchPriceQuery{Provider: &gatedLegacy{run: slowImpl}}
	}

	go func() {
		// Brief wait so the dispatcher launches and parks the first
		// goroutine on <-gate before we release them all at once.
		time.Sleep(50 * time.Millisecond)
		close(gate)
	}()

	results, err := BatchLPTokenPriceWithOptions(
		context.Background(),
		&fakeMulticall{},
		nil, nil,
		queries,
		BatchOptions{LegacyParallelism: 1},
	)
	if err != nil {
		t.Fatalf("BatchLPTokenPriceWithOptions: %v", err)
	}
	if len(results) != n {
		t.Fatalf("len(results) = %d, want %d", len(results), n)
	}
	if got := atomic.LoadInt64(&peak); got != 1 {
		t.Errorf("peak in-flight = %d, want 1 (LegacyParallelism=1)", got)
	}
}

// gatedLegacy is a Protocol that delegates LPTokenPrice to a caller-supplied
// func. Used by the LegacyParallelism test to observe in-flight goroutine
// counts.
type gatedLegacy struct {
	run func(context.Context) (string, error)
}

func (g *gatedLegacy) GetConfig(context.Context, string, bind.ContractBackend) ([]byte, error) {
	return nil, nil
}
func (g *gatedLegacy) Initialize(context.Context, bind.ContractBackend, fetchers.HttpClient) error {
	return nil
}
func (g *gatedLegacy) LPTokenPrice(ctx context.Context) (string, error)          { return g.run(ctx) }
func (g *gatedLegacy) TVL(context.Context) (string, error)                       { return "", nil }
func (g *gatedLegacy) TVLBreakdown(context.Context) (map[string]TokenTVL, error) { return nil, nil }
func (g *gatedLegacy) UpdateBlock(*big.Int, map[string]Price)                    {}

// Suppress the unused-import warning in case zerolog is removed by a later
// refactor — keep it referenced so the import block stays stable across
// follow-up edits.
var _ = zerolog.Logger{}
