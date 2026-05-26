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

func TestEtherfi_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	assetPrice, _ := decimal.NewFromString("2500.00")
	// Deliberately use LPTDecimals (8) != asset Decimals (18) to catch any
	// regression where the share supply is normalized by the wrong decimals
	// in computeLPPriceFromReads. The legacy tvl() formula uses
	// assetPrice.Decimals for that normalization; the batched path must match.
	e := &EtherfiLPPriceProvider{
		address: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		logger:  zerolog.Nop(),
		config: &EtherfiConfig{
			Asset:       "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			Accountant:  "0x2222222222222222222222222222222222222222",
			LPTDecimals: 8,
		},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: assetPrice},
		},
	}
	totalSupply := new(big.Int).Mul(big.NewInt(10), bigPow10(8))
	rate := new(big.Int).Mul(big.NewInt(1), bigPow10(8))
	wantDirect, _ := e.computeLPPriceFromReads(totalSupply, rate)
	vaultABI, _ := sc.EtherfiVaultMetaData.GetAbi()
	encTS, _ := vaultABI.Methods["totalSupply"].Outputs.Pack(totalSupply)
	accABI, _ := sc.EtherfiAccountantMetaData.GetAbi()
	encRate, _ := accABI.Methods["getRate"].Outputs.Pack(rate)
	got, err := e.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTS},
		{Success: true, ReturnData: encRate},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !got.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s", got, wantDirect)
	}
}
