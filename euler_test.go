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

func TestEuler_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	assetPrice, _ := decimal.NewFromString("0.75")
	e := &EulerLPPriceProvider{
		address: common.HexToAddress("0x5555555555555555555555555555555555555555"),
		logger:  zerolog.Nop(),
		config:  &EulerConfig{Asset: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", LPTDecimals: 18},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: assetPrice},
		},
	}
	totalSupply := new(big.Int).Mul(big.NewInt(1000), bigPow10(18))
	totalAssets := new(big.Int).Mul(big.NewInt(1200), bigPow10(18))
	wantDirect, _ := e.computeLPPriceFromReads(totalSupply, totalAssets)
	abi, _ := sc.EulerVaultMetaData.GetAbi()
	encTS, _ := abi.Methods["totalSupply"].Outputs.Pack(totalSupply)
	encTA, _ := abi.Methods["totalAssets"].Outputs.Pack(totalAssets)
	got, err := e.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTS},
		{Success: true, ReturnData: encTA},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !got.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s", got, wantDirect)
	}
}
