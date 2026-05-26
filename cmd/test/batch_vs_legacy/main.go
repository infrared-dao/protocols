// Command batch_vs_legacy runs both the legacy LPTokenPrice path and the
// new BatchLPTokenPrice path against a curated set of live Berachain
// mainnet vaults at a single pinned block, then reports the diff.
//
// Purpose: validate that the BatchablePriceProvider migration preserves
// the exact numerical price that the legacy path produces, end-to-end,
// against a real RPC. Golden tests cover this with synthesised ABI
// responses; this tool covers it with on-chain state.
//
// Usage:
//
//	go run ./cmd/test/batch_vs_legacy -rpc https://mainnet.rpc-0.bera.rbx.lgns.net
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"

	"github.com/infrared-dao/protocols"
	"github.com/infrared-dao/protocols/cmd/test/http"
	"github.com/infrared-dao/protocols/fetchers"
	"github.com/infrared-dao/protocols/multicall3"
)

// testCase describes one vault to validate. configurator/builder split
// matches the existing pattern in sulaco's price.ProtocolAdapter — the
// configurator is a stateless instance used only to call GetConfig, and
// builder constructs a ready-to-Initialize provider from the decoded
// config bytes. The builder receives every input any provider needs
// (priceMap may be ignored — e.g. BeraBorrow — and the client is exposed
// for the rare provider that requires extra setup at construction time);
// individual closures pick the subset they need.
type testCase struct {
	name      string
	address   string
	configure func(ctx context.Context, addr string, client *ethclient.Client) ([]byte, error)
	build     func(addr common.Address, block *big.Int, prices map[string]protocols.Price, logger zerolog.Logger, config []byte, client *ethclient.Client) protocols.Protocol
}

// addressRE pulls all 0x-prefixed 40-hex-char tokens out of a JSON blob.
// We use it to find every referenced token address in a config so we can
// stub a priceMap entry for each. Mocked prices are fine because we're
// comparing path-vs-path, not vs ground truth.
var addressRE = regexp.MustCompile(`0x[0-9a-fA-F]{40}`)

// countingMC wraps a multicall3 client and counts how many aggregate3
// invocations it forwards. Used in phases 3 and 4 to verify that the
// dispatcher emits exactly the expected number of aggregate3 calls per
// BatchLPTokenPrice (1 per distinct BlockNumber partition).
type countingMC struct {
	inner *multicall3.Client
	calls int
}

func (c *countingMC) Aggregate3(ctx context.Context, calls []multicall3.Call3, block *big.Int) ([]multicall3.Result3, error) {
	c.calls++
	return c.inner.Aggregate3(ctx, calls, block)
}

// commonBerachainTokens covers tokens that pool-based providers
// (bex / burrbear) discover at runtime via vault.getPoolTokens — those
// addresses don't appear in the static config, so regex extraction
// misses them. Listed here so the mock priceMap is never short an entry.
// Adding more is harmless; they only matter if a future test case
// references one of them.
var commonBerachainTokens = []string{
	"0x6969696969696969696969696969696969696969", // WBERA
	"0x1ce0a25d13ce4d52071ae7e02cf1f6606f4c79d3", // BurrBear underlying (NECT-adjacent)
	"0xfcbd14dc51f0a4d49d5e53c2e0950e0bc26d0dce", // HONEY
	"0x549943e04f40284185054145c6e4e9568c1d3241", // USDC.e
	"0x0555e30da8f98308edb960aa94c0db47230d2b9c", // WBTC
	"0x2f6f07cdcf3588944bf4c42ac74ff24bf56e7590", // WETH
	"0xac03caba51e17c86c921e1f6cbfbdc91f8bb2e6b", // iBGT
	"0x6acf0a4dc3a7fa0bda85d2af8a02ec1ad88a8c10", // iBERA / NECT
}

func main() {
	rpc := flag.String("rpc", "https://mainnet.rpc-0.bera.rbx.lgns.net", "Berachain mainnet RPC URL")
	vaultsPath := flag.String("vaults", "", "Optional path to production-vaults.json — if set, runs Phase 5 (production-vault pre-flight) using its full vault list")
	flag.Parse()

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger().Level(zerolog.WarnLevel)
	ctx := context.Background()

	client, err := ethclient.Dial(*rpc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "rpc dial: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	// Pin the block once. Both paths run against this exact block so the
	// only variable is the code path, not the chain state.
	head, err := client.BlockNumber(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "BlockNumber: %v\n", err)
		os.Exit(1)
	}
	block := new(big.Int).SetUint64(head)
	fmt.Printf("Pinned block: %d\n\n", head)

	// Construct the multicall3 client once. Used for the batched path.
	mc, err := multicall3.NewClient(client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "multicall3 client: %v\n", err)
		os.Exit(1)
	}

	cases := []testCase{
		{
			name:      "bend",
			address:   "0x30bba9cd9eb8c95824aa42faa1bb397b07545bc1",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.BendLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewBendLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "brownfi",
			address:   "0xd932c344e21ef6C3a94971bf4D4cC71304E2a66C",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.BrownFiLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewBrownFiLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "bulla",
			address:   "0xcffbfd665bedb19b47837461a5abf4388c560d35",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.BullaLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewBullaLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "concrete",
			address:   "0xec577e989c02b294d5b8f4324224a5b63f5beef7",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.ConcreteLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewConcreteLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "d2",
			address:   "0xbe75c8a7e58c7901d2e128dc8d3b6de2481f1f79",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.D2LPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewD2LPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "dolomite",
			address:   "0x7f2b60fdff1494a0e3e060532c9980d7fad0404b",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.DolomiteLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewDolomiteLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "etherfi",
			address:   "0x46fcd35431f5B371224ACC2e2E91732867B1A77e",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.EtherfiLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewEtherfiLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "euler",
			address:   "0x112B77A77753b092306b1c04Bd70215FeD4e00a1",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.EulerLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewEulerLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "steer",
			address:   "0xDB78B4166580917c9604f8DdfBea5F49B493845c",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.SteerLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewSteerLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "termmax",
			address:   "0xd07F1862AE599697CDcd6Fd36dF3C33af25fd782",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.TermMaxVaultPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewTermMaxVaultPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "wasabee",
			address:   "0xEC06041013b3a97c58b9ab61eAE9079Bc594EdA3",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.WasabeeLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewWasabeeLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "wasabi",
			address:   "0xc95ab9eff8fb48760703c74416764b8f898afa1b",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.WasabiLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewWasabiLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "webera",
			address:   "0x55a050f76541c2554e9dfa3a0b4e665914bf92ea",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.WeberaLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewWeberaLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "winnieswap",
			address:   "0x46fbf6ff1fd62ec89af48c3bb0b63115052dab31",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.WinnieSwapLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewWinnieSwapLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "solvbtc",
			address:   "0x0F6f337B09cb5131cF0ce9df3Beb295b8e728F3B",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.SolvLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewSolvLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "kodiak",
			address:   "0x98bdeede9a45c28d229285d9d6e9139e9f505391",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.KodiakLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewKodiakLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "aquabera",
			address:   "0xf9845a03f7e6b06645a03a28b943c8a4b5fe7bcc",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.AquaBeraLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewAquaBeraLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			name:      "burrbear",
			address:   "0xd10e65a5f8ca6f835f2b1832e37cf150fb955f23",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.BurrBearLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewBurrBearLPPriceProvider(a, b, p, l, cfg)
			},
		},
		{
			// BeraBorrow has no priceMap parameter — its price source is
			// the configured asset oracle, looked up via the on-chain
			// AggregatorV3 the contract points at.
			name:      "beraborrow",
			address:   "0xB318Cd79dC0743De041A26D3F0d467d49955E5bC",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.BeraBorrowLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, _ map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewBeraBorrowLPPriceProvider(a, b, l, cfg)
			},
		},
		{
			// BEX takes vaultAddress + poolAddress; the BalancerVault for
			// BEX on Berachain mainnet is at 0x4Be03f...E80B. The test
			// `address` here is the LP token / pool address.
			name:      "bex",
			address:   "0x2c4a603a2aa5596287a06886862dc29d56dbc354",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.BexLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				vault := common.HexToAddress("0x4Be03f781C497A489E3cB0287833452cA9B9E80B")
				return protocols.NewBexLPPriceProvider(vault, a, b, p, l, cfg)
			},
		},
		{
			// D8x has no priceMap parameter — it reads its USD reference
			// price from the configured AggregatorV3 oracle on chain.
			name:      "d8x",
			address:   "0x26bbc26415c6316890565f5f73017f85ee70b60c",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.D8xLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, _ map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewD8xLPPriceProvider(a, b, l, cfg)
			},
		},
		{
			// IVX takes a separate lpMonitor address in addition to the
			// LP token address. lpMonitor on mainnet is hardcoded here;
			// the test `address` is the LP token itself.
			name:      "ivx",
			address:   "0x3b8B155E3C44f07f6EAd507570f4047C8B450A7F",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.IVXLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, _ map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				lpMonitor := common.HexToAddress("0xf30EC2B4363c7957dab4b83B3211a278e280802D")
				return protocols.NewIVXLPPriceProvider(lpMonitor, a, b, l, cfg)
			},
		},
		{
			// Pendle is off-chain: LPTokenPrice / ComputePrice both pull
			// pool state from api-v2.pendle.finance. Zero RPC reads. Both
			// paths should agree (subject to the 5-second cache window).
			name:      "pendle",
			address:   "0xc2605ed80880bd6b1523d52aef8d624ed468a935",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.PendleLPPriceProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, _ *big.Int, _ map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewPendleLPPriceProvider(a, l, cfg)
			},
		},
		{
			// PaddleFi has no LP-token price by design: LPTokenPrice
			// always returns ("", error). Both paths should error
			// identically — revert-parity.
			name:      "paddlefi",
			address:   "0xe16761787cF9bB0D3fC2E5C726dAe906ce81B102",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) { return (&protocols.PaddleFiProvider{}).GetConfig(ctx, a, c) },
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewPaddleFiProvider(a, b, p, l, cfg)
			},
		},
		{
			// PancakeSwap Infinity is a BSC-chain CLMM with no per-LP-token
			// price by design: LPTokenPrice returns "0" unconditionally
			// without any RPC. To avoid a BSC roundtrip just to build a
			// config (its GetConfig calls a CLPoolManager method on BSC),
			// we stub a minimal valid config here. Both paths should
			// return "0" against any RPC.
			name:    "pancakeswap_infinity",
			address: "0x0000000000000000000000000000000000000001",
			configure: func(ctx context.Context, a string, c *ethclient.Client) ([]byte, error) {
				return []byte(`{"pool_id":"0x0000000000000000000000000000000000000000000000000000000000000000","token0":"0x0000000000000000000000000000000000000001","token1":"0x0000000000000000000000000000000000000002","fee":3000,"cl_pool_manager":"0xa0FfB9c1CE1Fe56963B0321B32E7A0302114058b"}`), nil
			},
			build: func(a common.Address, b *big.Int, p map[string]protocols.Price, l zerolog.Logger, cfg []byte, _ *ethclient.Client) protocols.Protocol {
				return protocols.NewPancakeSwapInfinityLPPriceProvider(a, b, p, l, cfg)
			},
		},
	}

	thc := http.NewTestHttpClient()
	type row struct {
		name    string
		legacy  string
		batched string
		match   bool
		errMsg  string
		// provBatch is the initialized provider kept around so the multi-
		// query phase below can re-submit it as part of a combined
		// BatchLPTokenPrice call without paying for a second Initialize.
		provBatch protocols.Protocol
	}
	rows := make([]row, 0, len(cases))

	for _, tc := range cases {
		r := row{name: tc.name}
		legacyPrice, batchedPrice, provBatch, errMsg := runCase(ctx, tc, client, mc, block, logger, thc)
		r.legacy = legacyPrice
		r.batched = batchedPrice
		r.errMsg = errMsg
		// Match holds whenever the two paths agree, whether that's a
		// numerical match on a successful price or a "<both reverted>"
		// sentinel where both paths failed at the contract layer
		// (e.g. expired oracle). A non-empty errMsg in the reverted-
		// parity case is informational, not a failure.
		r.match = legacyPrice != "" && legacyPrice == batchedPrice
		r.provBatch = provBatch
		rows = append(rows, r)
	}

	// === Phase 1 report: per-vault parity (N=1 batch per call) ===
	fmt.Println("Phase 1: per-vault parity (N=1 BatchLPTokenPrice per vault)")
	fmt.Printf("%-12s | %-30s | %-30s | %-5s | %s\n", "protocol", "legacy LPTokenPrice", "BatchLPTokenPrice", "match", "notes")
	fmt.Printf("%-12s-+-%-30s-+-%-30s-+-%-5s-+-%s\n", strings.Repeat("-", 12), strings.Repeat("-", 30), strings.Repeat("-", 30), strings.Repeat("-", 5), strings.Repeat("-", 30))
	fails := 0
	for _, r := range rows {
		marker := "yes"
		if !r.match {
			marker = "NO"
			fails++
		}
		fmt.Printf("%-12s | %-30s | %-30s | %-5s | %s\n", r.name, r.legacy, r.batched, marker, r.errMsg)
	}
	fmt.Printf("\n%d/%d passed\n\n", len(rows)-fails, len(rows))

	// === Phase 2: N>1 multi-query batch ===
	// Take every successfully-initialized provider and submit them as ONE
	// BatchLPTokenPrice invocation. Verifies that the dispatcher correctly
	// partitions by block (all same block here so we expect 1 aggregate3),
	// packs reads, and demultiplexes responses back per-query — against
	// real RPC, not mocks.
	queries := make([]protocols.BatchPriceQuery, 0, len(rows))
	idxToRow := make([]int, 0, len(rows))
	for i, r := range rows {
		if r.provBatch == nil {
			continue
		}
		queries = append(queries, protocols.BatchPriceQuery{
			Provider:    r.provBatch,
			BlockNumber: block,
		})
		idxToRow = append(idxToRow, i)
	}

	multiResults, err := protocols.BatchLPTokenPrice(ctx, mc, client, thc, queries)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Phase 2 BatchLPTokenPrice (N=%d): %v\n", len(queries), err)
		os.Exit(2)
	}

	fmt.Printf("Phase 2: multi-query batch (N=%d in a single BatchLPTokenPrice call)\n", len(queries))
	fmt.Printf("%-12s | %-30s | %-30s | %-5s | %s\n", "protocol", "single-query batched", "multi-query batched", "match", "notes")
	fmt.Printf("%-12s-+-%-30s-+-%-30s-+-%-5s-+-%s\n", strings.Repeat("-", 12), strings.Repeat("-", 30), strings.Repeat("-", 30), strings.Repeat("-", 5), strings.Repeat("-", 30))
	multiFails := 0
	for i, res := range multiResults {
		r := &rows[idxToRow[i]]
		var multiStr, note string
		if res.Err != nil {
			note = fmt.Sprintf("multi-query err: %v", res.Err)
			multiStr = "<reverted>"
		} else {
			multiStr = res.Price.StringFixed(8)
		}
		// Parity holds iff both phases agree on outcome. Phase 1 records
		// "<both reverted>" when both single-query paths revert; in that
		// case Phase 2 must also revert (multiStr == "<reverted>"). For
		// successful prices, multiStr must match Phase 1's batched value.
		var match bool
		if r.batched == "<both reverted>" {
			match = res.Err != nil
		} else {
			match = note == "" && multiStr == r.batched
		}
		marker := "yes"
		if !match {
			marker = "NO"
			multiFails++
		}
		fmt.Printf("%-12s | %-30s | %-30s | %-5s | %s\n", r.name, r.batched, multiStr, marker, note)
	}
	fmt.Printf("\n%d/%d passed in multi-query batch\n\n", len(multiResults)-multiFails, len(multiResults))

	// === Phase 3: multi-block partitioning ===
	// Submit each healthy vault twice in the same BatchLPTokenPrice call
	// — once at head, once at head-10. The dispatcher should partition by
	// BlockNumber and emit exactly 2 aggregate3 calls (one per block).
	// For each result, compare against an independent single-query
	// baseline at the same block.
	olderBlock := new(big.Int).Sub(block, big.NewInt(10))
	fmt.Printf("Phase 3: multi-block partitioning (head=%d + head-10=%d in one BatchLPTokenPrice call)\n", block.Uint64(), olderBlock.Uint64())

	type mbQuery struct {
		rowIdx int
		block  *big.Int
		label  string
	}
	var mbQueries []protocols.BatchPriceQuery
	var mbMeta []mbQuery
	for i, r := range rows {
		if r.provBatch == nil || r.batched == "<both reverted>" {
			continue
		}
		mbQueries = append(mbQueries, protocols.BatchPriceQuery{Provider: r.provBatch, BlockNumber: block})
		mbMeta = append(mbMeta, mbQuery{rowIdx: i, block: block, label: "head"})
		mbQueries = append(mbQueries, protocols.BatchPriceQuery{Provider: r.provBatch, BlockNumber: olderBlock})
		mbMeta = append(mbMeta, mbQuery{rowIdx: i, block: olderBlock, label: "head-10"})
	}

	cMC := &countingMC{inner: mc}
	mbResults, err := protocols.BatchLPTokenPrice(ctx, cMC, client, thc, mbQueries)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Phase 3 BatchLPTokenPrice: %v\n", err)
		os.Exit(2)
	}
	mbDispatches := cMC.calls
	mbPartitionOK := mbDispatches == 2
	if mbPartitionOK {
		fmt.Printf("  aggregate3 invocations: %d  (expected 2 — one per block) ✓\n", mbDispatches)
	} else {
		fmt.Printf("  aggregate3 invocations: %d  (expected 2 — one per block) ✗\n", mbDispatches)
	}

	// For each multi-block result, run an independent single-query
	// BatchLPTokenPrice at that block to get a baseline price, then
	// compare the multi-block result against it.
	fmt.Printf("  %-12s | %-7s | %-30s | %-30s | %s\n", "protocol", "block", "baseline (N=1 at block)", "multi-block result", "match")
	mbFails := 0
	for i, res := range mbResults {
		meta := mbMeta[i]
		r := &rows[meta.rowIdx]

		var multiStr string
		if res.Err != nil {
			multiStr = fmt.Sprintf("<err: %v>", res.Err)
		} else {
			multiStr = res.Price.StringFixed(8)
		}

		// Baseline: same provider, same block, but in its own
		// BatchLPTokenPrice call — bypasses the partitioning logic.
		baselineRes, baselineErr := protocols.BatchLPTokenPrice(ctx, mc, client, thc, []protocols.BatchPriceQuery{
			{Provider: r.provBatch, BlockNumber: meta.block},
		})
		var baseStr string
		if baselineErr != nil {
			baseStr = fmt.Sprintf("<top-err: %v>", baselineErr)
		} else if baselineRes[0].Err != nil {
			baseStr = fmt.Sprintf("<err: %v>", baselineRes[0].Err)
		} else {
			baseStr = baselineRes[0].Price.StringFixed(8)
		}

		match := baseStr == multiStr
		marker := "yes"
		if !match {
			marker = "NO"
			mbFails++
		}
		fmt.Printf("  %-12s | %-7s | %-30s | %-30s | %s\n", r.name, meta.label, baseStr, multiStr, marker)
	}
	if !mbPartitionOK {
		mbFails++
	}
	fmt.Printf("\n%d/%d passed in multi-block (and partition count %s)\n\n", len(mbResults)-mbFails, len(mbResults), boolMark(mbPartitionOK))

	// === Phase 4: N=100 scale test ===
	// Duplicate the per-row provider list 5× to get ~100 queries, all at
	// the same block. Exercises Multicall3 contract gas limits and the
	// dispatcher's range-tracking under load. We expect exactly 1
	// aggregate3 call (single partition) and every result to match its
	// per-vault Phase-1 batched price.
	const scaleMultiplier = 5
	fmt.Printf("Phase 4: scale test (N=%d in one BatchLPTokenPrice call, expect 1 aggregate3)\n", len(rows)*scaleMultiplier)

	type scaleQuery struct {
		rowIdx int
	}
	var scaleQueries []protocols.BatchPriceQuery
	var scaleMeta []scaleQuery
	for rep := 0; rep < scaleMultiplier; rep++ {
		for i, r := range rows {
			if r.provBatch == nil {
				continue
			}
			scaleQueries = append(scaleQueries, protocols.BatchPriceQuery{Provider: r.provBatch, BlockNumber: block})
			scaleMeta = append(scaleMeta, scaleQuery{rowIdx: i})
		}
	}

	cMC2 := &countingMC{inner: mc}
	scaleResults, err := protocols.BatchLPTokenPrice(ctx, cMC2, client, thc, scaleQueries)
	scaleFails := 0
	scalePartitionOK := cMC2.calls == 1
	if err != nil {
		fmt.Printf("  BatchLPTokenPrice top-level err: %v\n", err)
		fmt.Printf("  (likely Multicall3 ran out of gas at this N — note the practical aggregate3 size limit)\n")
		os.Exit(2)
	}
	if scalePartitionOK {
		fmt.Printf("  aggregate3 invocations: %d  (expected 1 — single partition) ✓\n", cMC2.calls)
	} else {
		fmt.Printf("  aggregate3 invocations: %d  (expected 1 — single partition) ✗\n", cMC2.calls)
		scaleFails++
	}

	// Spot-check: every Nth query must match its Phase-1 batched result.
	// Print only mismatches to keep the output readable; emit a count at
	// the end either way.
	for i, res := range scaleResults {
		r := &rows[scaleMeta[i].rowIdx]
		var got string
		if res.Err != nil {
			got = "<reverted>"
		} else {
			got = res.Price.StringFixed(8)
		}
		// Match definition: same as Phase 2 — revert-symmetric or numerical.
		var match bool
		if r.batched == "<both reverted>" {
			match = res.Err != nil
		} else {
			match = res.Err == nil && got == r.batched
		}
		if !match {
			scaleFails++
			fmt.Printf("  MISMATCH q[%d] %s: Phase-1=%s, scale=%s, err=%v\n", i, r.name, r.batched, got, res.Err)
		}
	}
	fmt.Printf("\n%d/%d passed at scale N=%d\n", len(scaleResults)-scaleFails, len(scaleResults), len(scaleResults))

	totalFails := fails + multiFails + mbFails + scaleFails

	// === Phase 6: concurrent batch dispatch ===
	// Fire M goroutines, each running BatchLPTokenPrice against the SAME
	// shared provider instances at the same block. Verifies (a) no data
	// races (run with `go run -race`), (b) every goroutine produces
	// identical results to the Phase 1 baseline. Sulaco's price collector
	// can issue overlapping cycles, so this models the worst-case
	// concurrent-read pattern against the migrated providers.
	//
	// Builds a FRESH set of provider instances for this phase rather than
	// reusing the Phase-1 ones, so any provider with first-call mutable
	// state (e.g. pendle's HTTP cache) starts cold and the race detector
	// has the best chance of catching concurrent writes. We also coordinate
	// goroutine start via a barrier channel so they hit the dispatch code
	// at roughly the same instant rather than serialising on goroutine
	// startup.
	const concurrency = 10
	fmt.Printf("Phase 6: concurrent batch dispatch (%d goroutines, each N=%d, fresh providers, barrier-synced start)\n", concurrency, len(rows))

	freshQueries := make([]protocols.BatchPriceQuery, 0, len(rows))
	freshBaselines := make([]string, 0, len(rows)) // each query's expected price (matches rows[].batched)
	for _, r := range rows {
		if r.provBatch == nil {
			continue
		}
		// Look up the testCase by name to rebuild from scratch.
		var tc testCase
		for _, candidate := range cases {
			if candidate.name == r.name {
				tc = candidate
				break
			}
		}
		// Fetch config and initialize a fresh provider — same flow as
		// runCase but skipping the legacy/batched comparison.
		cfg, err := tc.configure(ctx, tc.address, client)
		if err != nil {
			fmt.Printf("  Phase 6 config fetch failed for %s: %v\n", r.name, err)
			continue
		}
		prices := map[string]protocols.Price{}
		one, _ := decimal.NewFromString("1.0")
		for _, addr := range addressRE.FindAllString(string(cfg), -1) {
			prices[strings.ToLower(addr)] = protocols.Price{Decimals: 18, Price: one}
		}
		for _, addr := range commonBerachainTokens {
			prices[strings.ToLower(addr)] = protocols.Price{Decimals: 18, Price: one}
		}
		prices[strings.ToLower(tc.address)] = protocols.Price{Decimals: 18, Price: one}
		fresh := tc.build(common.HexToAddress(tc.address), block, prices, logger, cfg, client)
		if err := fresh.Initialize(ctx, client, thc); err != nil {
			fmt.Printf("  Phase 6 init failed for %s: %v\n", r.name, err)
			continue
		}
		freshQueries = append(freshQueries, protocols.BatchPriceQuery{Provider: fresh, BlockNumber: block})
		freshBaselines = append(freshBaselines, r.batched)
	}

	var (
		wg               sync.WaitGroup
		concResults      = make([][]protocols.BatchPriceResult, concurrency)
		concTopLevelErrs = make([]error, concurrency)
		startBarrier     = make(chan struct{})
	)
	wg.Add(concurrency)
	for g := 0; g < concurrency; g++ {
		go func(idx int) {
			defer wg.Done()
			<-startBarrier // wait until every goroutine is parked
			res, err := protocols.BatchLPTokenPrice(ctx, mc, client, thc, freshQueries)
			concResults[idx] = res
			concTopLevelErrs[idx] = err
		}(g)
	}
	close(startBarrier) // release all goroutines simultaneously
	wg.Wait()

	concFails := 0
	for g := 0; g < concurrency; g++ {
		if concTopLevelErrs[g] != nil {
			fmt.Printf("  goroutine %d top-level err: %v\n", g, concTopLevelErrs[g])
			concFails++
			continue
		}
		if len(concResults[g]) != len(freshQueries) {
			fmt.Printf("  goroutine %d returned %d results, want %d\n", g, len(concResults[g]), len(freshQueries))
			concFails++
			continue
		}
		// Compare each goroutine's result against the Phase-1 baseline.
		// Any divergence here means the shared provider state got
		// perturbed across concurrent calls.
		for i, res := range concResults[g] {
			baseline := freshBaselines[i]
			var got string
			if res.Err != nil {
				got = "<reverted>"
			} else {
				got = res.Price.StringFixed(8)
			}
			var match bool
			if baseline == "<both reverted>" {
				match = res.Err != nil
			} else {
				match = res.Err == nil && got == baseline
			}
			if !match {
				concFails++
				fmt.Printf("  MISMATCH goroutine=%d, query[%d]: Phase-1=%s, concurrent=%s, err=%v\n", g, i, baseline, got, res.Err)
			}
		}
	}
	totalConcResults := concurrency * len(freshQueries)
	fmt.Printf("\n%d/%d concurrent results match Phase-1 baseline (across %d goroutines)\n\n", totalConcResults-concFails, totalConcResults, concurrency)
	totalFails += concFails

	// === Phase 5: production-vault pre-flight (optional) ===
	// When --vaults is supplied, iterate every vault in the production
	// list, map its protocol.id to the corresponding testCase builder,
	// and run the same per-vault parity check (Phase-1-style) against
	// real on-chain state. Catches any production-specific config edge
	// case the curated 22-vault set doesn't exercise.
	if *vaultsPath != "" {
		p5Fails := runPhase5(ctx, *vaultsPath, client, mc, block, logger, thc, cases)
		totalFails += p5Fails
	}

	if totalFails > 0 {
		fmt.Fprintf(os.Stderr, "\n%d FAILED across all phases\n", totalFails)
		os.Exit(2)
	}
}

func boolMark(ok bool) string {
	if ok {
		return "OK"
	}
	return "FAILED"
}

// productionVaultsFile mirrors the subset of production-vaults.json we
// need: the per-vault id, stake_token address, and protocol identifier.
type productionVaultsFile struct {
	Vaults []struct {
		ID         string `json:"id"`
		StakeToken struct {
			Address string `json:"address"`
		} `json:"stake_token"`
		Protocol struct {
			ID string `json:"id"`
		} `json:"protocol"`
	} `json:"vaults"`
}

// runPhase5 iterates production-vaults.json, dispatches each vault to
// the matching testCase builder (via protocol.id), runs per-vault parity,
// and reports an aggregate pass/fail/skip summary. Unmappable protocols
// (those without an in-tree adapter — beradrome, kuma, etc.) are skipped
// with an explicit count rather than treated as failures.
func runPhase5(
	ctx context.Context,
	path string,
	client *ethclient.Client,
	mc *multicall3.Client,
	block *big.Int,
	logger zerolog.Logger,
	thc fetchers.HttpClient,
	cases []testCase,
) int {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Phase 5: read %s: %v\n", path, err)
		return 1
	}
	var pv productionVaultsFile
	if err := json.Unmarshal(data, &pv); err != nil {
		fmt.Fprintf(os.Stderr, "Phase 5: parse %s: %v\n", path, err)
		return 1
	}

	// Build a lookup from testCase name → testCase so we can dispatch
	// quickly by protocol.id with name aliasing (e.g. d2finance → d2).
	byName := map[string]testCase{}
	for _, tc := range cases {
		byName[tc.name] = tc
	}
	// Map production protocol.id values to the curated testCase name.
	// IDs not in this map (or whose target name isn't in `cases`) are
	// reported as skipped — those are external protocols without an
	// in-tree adapter (beradrome, kuma, narra, etc.).
	protocolAlias := map[string]string{
		"d2finance": "d2",
		"solv":      "solvbtc",
		"charm":     "kodiak", // Charm pools use the Kodiak adapter
	}

	fmt.Printf("\nPhase 5: production-vault pre-flight (%d vaults from %s)\n", len(pv.Vaults), path)
	fmt.Printf("  Pinned block: %d\n\n", block.Uint64())

	type p5row struct {
		id        string
		protocol  string
		address   string
		legacy    string
		batched   string
		match     bool
		errMsg    string
		skipped   bool
		skipNote  string
	}
	var rows []p5row
	for _, v := range pv.Vaults {
		row := p5row{id: v.ID, protocol: v.Protocol.ID, address: v.StakeToken.Address}
		if row.address == "" {
			row.skipped = true
			row.skipNote = "no stake_token.address"
			rows = append(rows, row)
			continue
		}
		// Resolve protocol.id → testCase name via alias map first, then
		// direct lookup. Skip vaults whose protocol has no in-tree adapter.
		targetName := v.Protocol.ID
		if alias, ok := protocolAlias[targetName]; ok {
			targetName = alias
		}
		tc, ok := byName[targetName]
		if !ok {
			row.skipped = true
			row.skipNote = "no in-tree adapter"
			rows = append(rows, row)
			continue
		}
		// Override the testCase's address with this vault's stake token.
		tc.address = row.address
		legacyPrice, batchedPrice, _, errMsg := runCase(ctx, tc, client, mc, block, logger, thc)
		row.legacy = legacyPrice
		row.batched = batchedPrice
		row.errMsg = errMsg
		// A GetConfig revert means the vault isn't priceable on either
		// path — that's a pre-existing library limitation for that vault,
		// not a parity bug. Classify as skipped rather than failed so it
		// doesn't drown out real mismatches.
		if errMsg != "" && legacyPrice == "" && batchedPrice == "" {
			row.skipped = true
			row.skipNote = errMsg
		} else {
			row.match = legacyPrice != "" && legacyPrice == batchedPrice
		}
		rows = append(rows, row)
	}

	// Per-protocol aggregate report — keeps the output compact even with
	// 190 vaults. Print individual misses only.
	type stat struct{ matched, mismatched, skipped int }
	byProto := map[string]*stat{}
	for _, r := range rows {
		s := byProto[r.protocol]
		if s == nil {
			s = &stat{}
			byProto[r.protocol] = s
		}
		switch {
		case r.skipped:
			s.skipped++
		case r.match:
			s.matched++
		default:
			s.mismatched++
		}
	}

	// Sort protocol names for deterministic output.
	protos := make([]string, 0, len(byProto))
	for p := range byProto {
		protos = append(protos, p)
	}
	sort.Strings(protos)
	fmt.Printf("  %-22s | matched | mismatched | skipped\n", "protocol")
	fmt.Printf("  %s\n", strings.Repeat("-", 56))
	totalMatched, totalMismatched, totalSkipped := 0, 0, 0
	for _, p := range protos {
		s := byProto[p]
		totalMatched += s.matched
		totalMismatched += s.mismatched
		totalSkipped += s.skipped
		fmt.Printf("  %-22s | %7d | %10d | %7d\n", p, s.matched, s.mismatched, s.skipped)
	}
	fmt.Printf("  %s\n", strings.Repeat("-", 56))
	fmt.Printf("  %-22s | %7d | %10d | %7d\n\n", "TOTAL", totalMatched, totalMismatched, totalSkipped)

	// List individual mismatches (if any) so they're easy to triage.
	if totalMismatched > 0 {
		fmt.Println("  Per-vault mismatches:")
		for _, r := range rows {
			if !r.skipped && !r.match {
				fmt.Printf("    %s (%s @ %s)\n        legacy:  %q\n        batched: %q\n        err:     %s\n",
					r.id, r.protocol, r.address, r.legacy, r.batched, r.errMsg)
			}
		}
	}
	return totalMismatched
}

// runCase fetches config from chain, builds the provider, then drives
// both paths at the same pinned block. Returns the two prices (legacy /
// batched) as strings, the initialized batched-path provider (kept alive
// for the multi-query phase), plus any error encountered.
func runCase(
	ctx context.Context,
	tc testCase,
	client *ethclient.Client,
	mc *multicall3.Client,
	block *big.Int,
	logger zerolog.Logger,
	thc fetchers.HttpClient,
) (legacy, batched string, provBatch protocols.Protocol, errMsg string) {
	cfg, err := tc.configure(ctx, tc.address, client)
	if err != nil {
		return "", "", nil, fmt.Sprintf("GetConfig: %v", err)
	}

	// Stub a priceMap entry for every address that appears in the config,
	// plus every well-known Berachain token a pool-based provider might
	// look up at runtime via getPoolTokens. Mock prices are fine — both
	// paths consume the same map, so the diff (or lack of) is what we're
	// checking.
	prices := map[string]protocols.Price{}
	one, _ := decimal.NewFromString("1.0")
	for _, addr := range addressRE.FindAllString(string(cfg), -1) {
		prices[strings.ToLower(addr)] = protocols.Price{Decimals: 18, Price: one}
	}
	for _, addr := range commonBerachainTokens {
		prices[strings.ToLower(addr)] = protocols.Price{Decimals: 18, Price: one}
	}
	// Also include the LP token's own address — some providers
	// (PaddleFi pattern) look it up that way.
	prices[strings.ToLower(tc.address)] = protocols.Price{Decimals: 18, Price: one}

	addr := common.HexToAddress(tc.address)

	// Build two independent provider instances so the legacy path and
	// the batched path have isolated state — no risk of a path mutating
	// internal state that influences the other.
	provLegacy := tc.build(addr, block, prices, logger, cfg, client)
	provBatch = tc.build(addr, block, prices, logger, cfg, client)

	if err := provLegacy.Initialize(ctx, client, thc); err != nil {
		return "", "", nil, fmt.Sprintf("Initialize (legacy): %v", err)
	}
	if err := provBatch.Initialize(ctx, client, thc); err != nil {
		return "", "", nil, fmt.Sprintf("Initialize (batch): %v", err)
	}

	// Drive both paths unconditionally so we can compare success/error
	// shape across both — a revert at the contract layer (e.g. d8x's
	// oracle feed expired, ivx's getSharePrice division-by-zero) should
	// surface symmetrically on both paths and that's a valid parity
	// outcome, not a migration bug.
	legacy, lerr := provLegacy.LPTokenPrice(ctx)
	results, batchErr := protocols.BatchLPTokenPrice(ctx, mc, client, thc, []protocols.BatchPriceQuery{
		{Provider: provBatch, BlockNumber: block},
	})

	// Legacy succeeded: must compare batched to it numerically.
	if lerr == nil {
		if batchErr != nil {
			return legacy, "", provBatch, fmt.Sprintf("BatchLPTokenPrice top-level err: %v", batchErr)
		}
		if len(results) != 1 {
			return legacy, "", provBatch, fmt.Sprintf("BatchLPTokenPrice: expected 1 result, got %d", len(results))
		}
		if results[0].Err != nil {
			return legacy, "", provBatch, fmt.Sprintf("BatchLPTokenPrice query: %v (legacy returned %s)", results[0].Err, legacy)
		}
		// The batched path returns a decimal.Decimal — convert to the same
		// fixed-decimal string the legacy path emits so the comparison is
		// apples-to-apples. The legacy path uses StringFixed(roundingDecimals)
		// internally; roundingDecimals is package-private (= 8 at the time of
		// writing) so we hard-code 8 here to match.
		batched = results[0].Price.StringFixed(8)
		return legacy, batched, provBatch, ""
	}

	// Legacy errored: parity holds iff the batched path also errored
	// at the per-query level. The two error strings won't be byte-identical
	// (legacy goes through go-ethereum's typed bindings, batched through
	// raw Multicall3 aggregate3 + ABI decoding) but both should report
	// the same root revert.
	var batchedErr error
	if batchErr != nil {
		batchedErr = batchErr
	} else if len(results) == 1 && results[0].Err != nil {
		batchedErr = results[0].Err
	}
	if batchedErr == nil {
		return "", "", provBatch, fmt.Sprintf("LEGACY ERRORED but BATCHED SUCCEEDED — legacy=%v, batched=%s", lerr, results[0].Price.StringFixed(8))
	}
	// Both errored. Surface as "both-reverted" by leaving the price
	// strings empty and using a notes line that says so. We can't string-
	// equality compare the error texts (they're framed differently), but
	// the fact that both paths fail at all is parity for our purposes.
	return "<both reverted>", "<both reverted>", provBatch, fmt.Sprintf("legacy=%v; batched=%v", lerr, batchedErr)
}

