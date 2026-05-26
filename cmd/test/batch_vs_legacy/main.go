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
	"flag"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"

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
	}

	thc := http.NewTestHttpClient()
	type row struct {
		name    string
		legacy  string
		batched string
		match   bool
		errMsg  string
	}
	rows := make([]row, 0, len(cases))

	for _, tc := range cases {
		r := row{name: tc.name}
		legacyPrice, batchedPrice, errMsg := runCase(ctx, tc, client, mc, block, logger, thc)
		r.legacy = legacyPrice
		r.batched = batchedPrice
		r.errMsg = errMsg
		r.match = errMsg == "" && legacyPrice == batchedPrice
		rows = append(rows, r)
	}

	// Report. Wider columns to fit long decimals.
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
	fmt.Printf("\n%d/%d passed", len(rows)-fails, len(rows))
	if fails > 0 {
		fmt.Printf(" — %d FAILED\n", fails)
		os.Exit(2)
	}
	fmt.Println()
}

// runCase fetches config from chain, builds the provider, then drives
// both paths at the same pinned block. Returns the two prices (legacy /
// batched) as strings plus any error encountered.
func runCase(
	ctx context.Context,
	tc testCase,
	client *ethclient.Client,
	mc *multicall3.Client,
	block *big.Int,
	logger zerolog.Logger,
	thc fetchers.HttpClient,
) (legacy, batched, errMsg string) {
	cfg, err := tc.configure(ctx, tc.address, client)
	if err != nil {
		return "", "", fmt.Sprintf("GetConfig: %v", err)
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
	provBatch := tc.build(addr, block, prices, logger, cfg, client)

	if err := provLegacy.Initialize(ctx, client, thc); err != nil {
		return "", "", fmt.Sprintf("Initialize (legacy): %v", err)
	}
	if err := provBatch.Initialize(ctx, client, thc); err != nil {
		return "", "", fmt.Sprintf("Initialize (batch): %v", err)
	}

	// Legacy path.
	legacy, lerr := provLegacy.LPTokenPrice(ctx)
	if lerr != nil {
		return "", "", fmt.Sprintf("legacy LPTokenPrice: %v", lerr)
	}

	// Batched path. One query, one provider — exercises dispatchBatch's
	// full PriceReads → aggregate3 → ComputePrice round-trip.
	results, err := protocols.BatchLPTokenPrice(ctx, mc, client, thc, []protocols.BatchPriceQuery{
		{Provider: provBatch, BlockNumber: block},
	})
	if err != nil {
		return legacy, "", fmt.Sprintf("BatchLPTokenPrice: %v", err)
	}
	if len(results) != 1 {
		return legacy, "", fmt.Sprintf("BatchLPTokenPrice: expected 1 result, got %d", len(results))
	}
	if results[0].Err != nil {
		return legacy, "", fmt.Sprintf("BatchLPTokenPrice query: %v", results[0].Err)
	}
	// The batched path returns a decimal.Decimal — convert to the same
	// fixed-decimal string the legacy path emits so the comparison is
	// apples-to-apples. The legacy path uses StringFixed(roundingDecimals)
	// internally; roundingDecimals is package-private (= 8 at the time of
	// writing) so we hard-code 8 here to match.
	batched = results[0].Price.StringFixed(8)
	return legacy, batched, ""
}

