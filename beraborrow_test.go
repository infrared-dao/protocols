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

// Golden: batchable ComputePrice and the direct computeLPPriceFromReads
// must produce identical prices for the CICV variant (most common).
// The ComputePrice path is reached via packed-then-unpacked ABI responses
// to ensure decode logic doesn't introduce drift from the shared math.
func TestBeraBorrow_BatchPathMatchesLegacy_CICV(t *testing.T) {
	t.Parallel()

	b := &BeraBorrowLPPriceProvider{
		LPTAddress:  common.HexToAddress("0x1111111111111111111111111111111111111111"),
		block:       nil,
		logger:      zerolog.Nop(),
		config: &BeraBorrowCDPConfig{
			ColVaultAddress: "0x2222222222222222222222222222222222222222", // CICV variant
			LPTDecimals:     18,
			CDPDecimals:     18,
		},
	}

	// Inputs: LP totalSupply = 100e18 LP tokens
	// CICV totalSupply = 200e18 underlying tokens
	// FetchPrice = 3.50e18 (3.50 USD per underlying with 18 decimal precision)
	// TVL = 200 * 3.50 = $700
	// pricePerLP = 700 / 100 = $7.00
	lpTotalSupply := new(big.Int).Mul(big.NewInt(100), bigPow10(18))
	cdpTotalSupply := new(big.Int).Mul(big.NewInt(200), bigPow10(18))
	pricePerCDP := new(big.Int).Mul(big.NewInt(350), bigPow10(16)) // 3.50 * 1e18

	// Direct: build tvl then call shared helper.
	pricePerToken := NormalizeAmount(pricePerCDP, USDPriceDecimals)
	numTokens := NormalizeAmount(cdpTotalSupply, b.config.CDPDecimals)
	tvl := numTokens.Mul(pricePerToken)
	wantDirect, err := b.computeLPPriceFromReads(lpTotalSupply, tvl)
	if err != nil {
		t.Fatalf("computeLPPriceFromReads: %v", err)
	}

	// Batch path: pack ABI responses for IW.totalSupply + CICV.fetchPrice +
	// CICV.totalSupply, hand them to ComputePrice.
	iwABI, _ := sc.BeraBorrowIWMetaData.GetAbi()
	cicvABI, _ := sc.BeraBorrowCICVMetaData.GetAbi()

	encLpTS, err := iwABI.Methods["totalSupply"].Outputs.Pack(lpTotalSupply)
	if err != nil {
		t.Fatalf("pack IW totalSupply: %v", err)
	}
	encFetchPrice, err := cicvABI.Methods["fetchPrice"].Outputs.Pack(pricePerCDP)
	if err != nil {
		t.Fatalf("pack CICV fetchPrice: %v", err)
	}
	encCdpTS, err := cicvABI.Methods["totalSupply"].Outputs.Pack(cdpTotalSupply)
	if err != nil {
		t.Fatalf("pack CICV totalSupply: %v", err)
	}

	gotBatch, err := b.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encLpTS},
		{Success: true, ReturnData: encFetchPrice},
		{Success: true, ReturnData: encCdpTS},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}

	if !gotBatch.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s — paths diverged!", gotBatch, wantDirect)
	}
	want, _ := decimal.NewFromString("7.00")
	if !gotBatch.Round(2).Equal(want) {
		t.Errorf("computed %s, want %s", gotBatch.Round(2), want)
	}
}

// Same golden test for the sNECT variant — different code path, different
// responses, same architectural guarantee.
func TestBeraBorrow_BatchPathMatchesLegacy_SNECT(t *testing.T) {
	t.Parallel()

	asset := common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")

	b := &BeraBorrowLPPriceProvider{
		LPTAddress:  common.HexToAddress("0x3333333333333333333333333333333333333333"),
		block:       nil,
		logger:      zerolog.Nop(),
		config: &BeraBorrowCDPConfig{
			ColVaultAddress: sNECTName,
			LPTDecimals:     18,
			CDPDecimals:     18,
		},
		snectAsset: asset,
	}

	// Inputs: LP totalSupply = 50e18
	// TotalAssets = 100e18 NECT
	// GetPrice = 1.50e18 (NECT @ $1.50 with 18-dec precision)
	// TVL = 100 * 1.50 = $150
	// pricePerLP = 150 / 50 = $3.00
	lpTotalSupply := new(big.Int).Mul(big.NewInt(50), bigPow10(18))
	totalAssets := new(big.Int).Mul(big.NewInt(100), bigPow10(18))
	nectPrice := new(big.Int).Mul(big.NewInt(150), bigPow10(16))

	pricePerToken := NormalizeAmount(nectPrice, USDPriceDecimals)
	numTokens := NormalizeAmount(totalAssets, b.config.CDPDecimals)
	tvl := numTokens.Mul(pricePerToken)
	wantDirect, _ := b.computeLPPriceFromReads(lpTotalSupply, tvl)

	iwABI, _ := sc.BeraBorrowIWMetaData.GetAbi()
	snectABI, _ := sc.BeraBorrowSNECTMetaData.GetAbi()

	encLpTS, _ := iwABI.Methods["totalSupply"].Outputs.Pack(lpTotalSupply)
	encGetPrice, _ := snectABI.Methods["getPrice"].Outputs.Pack(nectPrice)
	encTotalAssets, _ := snectABI.Methods["totalAssets"].Outputs.Pack(totalAssets)

	gotBatch, err := b.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encLpTS},
		{Success: true, ReturnData: encGetPrice},
		{Success: true, ReturnData: encTotalAssets},
	})
	if err != nil {
		t.Fatalf("ComputePrice (sNECT): %v", err)
	}

	if !gotBatch.Equal(wantDirect) {
		t.Errorf("sNECT batch %s != direct %s", gotBatch, wantDirect)
	}
	want, _ := decimal.NewFromString("3.00")
	if !gotBatch.Round(2).Equal(want) {
		t.Errorf("computed %s, want %s", gotBatch.Round(2), want)
	}
}

// PriceReads must produce 3 calls in both variants, with correct targets:
// CICV target the CICV address; sNECT target the LP address (sNECT
// contract is at LPTAddress).
func TestBeraBorrow_PriceReads_TargetsByVariant(t *testing.T) {
	t.Parallel()

	cicv := &BeraBorrowLPPriceProvider{
		LPTAddress: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		logger:     zerolog.Nop(),
		config: &BeraBorrowCDPConfig{
			ColVaultAddress: "0x2222222222222222222222222222222222222222",
			LPTDecimals:     18, CDPDecimals: 18,
		},
	}
	calls, err := cicv.PriceReads()
	if err != nil {
		t.Fatalf("CICV PriceReads: %v", err)
	}
	if len(calls) != 3 {
		t.Fatalf("CICV calls = %d, want 3", len(calls))
	}
	if calls[0].Target != cicv.LPTAddress {
		t.Errorf("CICV call[0] target = %v, want LPT %v", calls[0].Target, cicv.LPTAddress)
	}
	cicvAddr := common.HexToAddress(cicv.config.ColVaultAddress)
	if calls[1].Target != cicvAddr || calls[2].Target != cicvAddr {
		t.Errorf("CICV calls[1,2] targets = %v %v, want CICV %v",
			calls[1].Target, calls[2].Target, cicvAddr)
	}

	snect := &BeraBorrowLPPriceProvider{
		LPTAddress: common.HexToAddress("0x3333333333333333333333333333333333333333"),
		logger:     zerolog.Nop(),
		config: &BeraBorrowCDPConfig{
			ColVaultAddress: sNECTName,
			LPTDecimals:     18, CDPDecimals: 18,
		},
		snectAsset: common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	}
	scalls, err := snect.PriceReads()
	if err != nil {
		t.Fatalf("sNECT PriceReads: %v", err)
	}
	if len(scalls) != 3 {
		t.Fatalf("sNECT calls = %d, want 3", len(scalls))
	}
	for i, c := range scalls {
		if c.Target != snect.LPTAddress {
			t.Errorf("sNECT calls[%d] target = %v, want LPT %v", i, c.Target, snect.LPTAddress)
		}
	}
}
