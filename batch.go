package protocols

import (
	"context"
	"fmt"
	"math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"

	"github.com/infrared-dao/protocols/fetchers"
	"github.com/infrared-dao/protocols/multicall3"
)

// BatchablePriceProvider is an opt-in extension of [Protocol] that
// decomposes LP-price computation into a declarative reads-list builder
// (PriceReads) and a pure compute-from-responses function (ComputePrice).
//
// Protocols that implement BatchablePriceProvider can be batched with
// other batchable protocols into a single Multicall3.aggregate3 invocation
// by [BatchLPTokenPrice]. Protocols that do NOT implement
// BatchablePriceProvider are processed individually via the legacy
// Protocol.LPTokenPrice() path by [BatchLPTokenPrice] (fan-out via
// errgroup), preserving backwards compatibility.
//
// Implementation contract:
//   - PriceReads MUST return exactly the set of eth_calls needed to compute
//     the LP price. Its returned slice's order is significant; ComputePrice
//     receives the corresponding []multicall3.Result3 in the same order.
//   - PriceReads is called once per BatchLPTokenPrice invocation; it must
//     not perform I/O.
//   - ComputePrice is pure: no network I/O, no mutation of provider state
//     other than what UpdateBlock already does. Per-call response failures
//     (Result3.Success == false) are the provider's responsibility to
//     handle — typically by returning an error.
//   - The provider's existing Initialize/GetConfig/UpdateBlock methods are
//     still called by the caller; PriceReads and ComputePrice represent
//     only the live-state read+compute step, replacing LPTokenPrice's
//     internal eth_call dispatch.
//
// Implementers should typically factor their existing LPTokenPrice math
// into a private helper that is shared with ComputePrice, so both paths
// produce identical results for the same inputs.
type BatchablePriceProvider interface {
	Protocol

	// PriceReads describes the eth_calls required to compute this
	// provider's LP token price at the block set via UpdateBlock.
	PriceReads() ([]multicall3.Call3, error)

	// ComputePrice computes the LP token price from the responses to the
	// calls returned by PriceReads. Inputs are responses[i] corresponding
	// to PriceReads()[i] by position.
	ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error)
}

// BatchPriceQuery describes one staking-token price request handed to
// [BatchLPTokenPrice]. Provider is the per-token Protocol instance (already
// initialized — caller is responsible for calling Initialize + UpdateBlock
// before dispatch). BlockNumber is the snapshot block; nil means latest.
type BatchPriceQuery struct {
	Provider    Protocol
	BlockNumber *big.Int
}

// BatchPriceResult holds the outcome of one staking-token price computation.
// Err is non-nil if the per-query computation failed; otherwise Price is
// the LP token price in USD.
type BatchPriceResult struct {
	Price decimal.Decimal
	Err   error
}

// BatchLPTokenPrice computes LP token prices for N queries.
//
// Queries whose Provider implements [BatchablePriceProvider] are dispatched
// in a single Multicall3.aggregate3 invocation (per shared block number —
// queries are partitioned by block, and one aggregate3 is emitted per
// partition). Queries whose Provider does NOT implement
// BatchablePriceProvider fall back to the legacy LPTokenPrice path,
// invoked individually in parallel via errgroup.
//
// Returns one result per input query, in the same order. Per-query errors
// are reported in BatchPriceResult.Err. A non-nil top-level error indicates
// invalid input (e.g., nil Provider); per-query errors do not.
//
// Caller is responsible for having already called Provider.Initialize and
// Provider.UpdateBlock on each query's Provider. legacyClient and
// legacyHTTPClient are needed only for the fallback path; pass nil if all
// providers are guaranteed to be BatchablePriceProvider.
func BatchLPTokenPrice(
	ctx context.Context,
	multicall multicall3Caller,
	legacyClient bind.ContractBackend,
	legacyHTTPClient fetchers.HttpClient,
	queries []BatchPriceQuery,
) ([]BatchPriceResult, error) {
	if multicall == nil {
		return nil, fmt.Errorf("BatchLPTokenPrice: multicall must not be nil")
	}
	for i, q := range queries {
		if q.Provider == nil {
			return nil, fmt.Errorf("BatchLPTokenPrice: queries[%d].Provider is nil", i)
		}
	}

	results := make([]BatchPriceResult, len(queries))

	// Partition queries: (1) batchable indices grouped by block, (2)
	// non-batchable indices for parallel legacy fan-out.
	batchableByBlock := make(map[string][]int) // block-key → query indices
	nonBatchable := []int{}
	for i, q := range queries {
		if _, ok := q.Provider.(BatchablePriceProvider); ok {
			key := blockKey(q.BlockNumber)
			batchableByBlock[key] = append(batchableByBlock[key], i)
		} else {
			nonBatchable = append(nonBatchable, i)
		}
	}

	// Dispatch the batched path: one aggregate3 per block partition.
	for _, indices := range batchableByBlock {
		dispatchBatch(ctx, multicall, queries, indices, results)
	}

	// Fan out the legacy path for non-batchable queries. Each runs the
	// existing LPTokenPrice synchronously; errgroup just parallelises them
	// up to a sane cap so a long tail doesn't dominate wall-clock time.
	if len(nonBatchable) > 0 {
		const legacyParallelism = 8
		eg, egCtx := errgroup.WithContext(ctx)
		eg.SetLimit(legacyParallelism)
		for _, i := range nonBatchable {
			eg.Go(func() error {
				results[i] = runLegacyPriceQuery(egCtx, queries[i])
				return nil
			})
		}
		_ = eg.Wait() //nolint:errcheck // per-query errors already recorded in results[i].Err
	}

	return results, nil
}

// dispatchBatch issues one aggregate3 for the given batchable query
// indices (which share a block number), then distributes per-call results
// back to each provider's ComputePrice.
func dispatchBatch(
	ctx context.Context,
	multicall multicall3Caller,
	queries []BatchPriceQuery,
	indices []int,
	results []BatchPriceResult,
) {
	// Build the aggregate Call3 list. callRanges records, per query, the
	// (start, end) slice into the response array so we can give each
	// provider its own response window.
	var allCalls []multicall3.Call3
	type callRange struct{ queryIdx, start, end int }
	ranges := make([]callRange, 0, len(indices))

	for _, idx := range indices {
		bp := queries[idx].Provider.(BatchablePriceProvider)
		reads, err := bp.PriceReads()
		if err != nil {
			results[idx] = BatchPriceResult{Err: fmt.Errorf("PriceReads: %w", err)}
			continue
		}
		ranges = append(ranges, callRange{
			queryIdx: idx,
			start:    len(allCalls),
			end:      len(allCalls) + len(reads),
		})
		allCalls = append(allCalls, reads...)
	}

	if len(allCalls) == 0 {
		return // all queries in this partition errored at PriceReads
	}

	// Block number is the same for all queries in this partition by
	// construction.
	block := queries[indices[0]].BlockNumber
	resp, err := multicall.Aggregate3(ctx, allCalls, block)
	if err != nil {
		// Network/encoding failure: every batched query in this partition
		// gets the same error.
		for _, r := range ranges {
			results[r.queryIdx] = BatchPriceResult{Err: fmt.Errorf("aggregate3: %w", err)}
		}
		return
	}

	// Distribute responses to each provider's ComputePrice.
	for _, r := range ranges {
		bp := queries[r.queryIdx].Provider.(BatchablePriceProvider)
		slice := resp[r.start:r.end]
		price, err := bp.ComputePrice(slice)
		if err != nil {
			results[r.queryIdx] = BatchPriceResult{Err: fmt.Errorf("ComputePrice: %w", err)}
			continue
		}
		results[r.queryIdx] = BatchPriceResult{Price: price}
	}
}

// runLegacyPriceQuery invokes the existing LPTokenPrice path on a provider
// that doesn't (yet) implement BatchablePriceProvider. Used by
// BatchLPTokenPrice's fallback fan-out for unmigrated protocols.
func runLegacyPriceQuery(ctx context.Context, q BatchPriceQuery) BatchPriceResult {
	priceStr, err := q.Provider.LPTokenPrice(ctx)
	if err != nil {
		return BatchPriceResult{Err: fmt.Errorf("LPTokenPrice: %w", err)}
	}
	price, err := decimal.NewFromString(priceStr)
	if err != nil {
		return BatchPriceResult{Err: fmt.Errorf("parse price %q: %w", priceStr, err)}
	}
	return BatchPriceResult{Price: price}
}

// multicall3Caller is the minimum dependency BatchLPTokenPrice needs from
// a Multicall3 client. Defined here (rather than imported from the
// multicall3 sub-package as a *Client) so consumers can inject their own
// implementations (e.g., retry-wrapped, metrics-wrapped) without depending
// on the concrete type.
type multicall3Caller interface {
	Aggregate3(ctx context.Context, calls []multicall3.Call3, blockNumber *big.Int) ([]multicall3.Result3, error)
}

// blockKey returns a stable string for grouping queries by block number.
// Treats nil (= latest) as its own key so all nil-block queries are batched
// together.
func blockKey(b *big.Int) string {
	if b == nil {
		return "latest"
	}
	return b.String()
}
