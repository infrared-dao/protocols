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

func TestTermMax_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	assetPrice, _ := decimal.NewFromString("1.05")
	tm := &TermMaxVaultPriceProvider{
		address: common.HexToAddress("0x6666666666666666666666666666666666666666"),
		logger:  zerolog.Nop(),
		config:  &TermMaxVaultConfig{Asset: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ShareDecimals: 18},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: assetPrice},
		},
	}
	totalSupply := new(big.Int).Mul(big.NewInt(500), bigPow10(18))
	totalAssets := new(big.Int).Mul(big.NewInt(525), bigPow10(18))
	wantDirect, _ := tm.computeLPPriceFromReads(totalSupply, totalAssets)
	abi, _ := sc.ERC4626MetaData.GetAbi()
	encTS, _ := abi.Methods["totalSupply"].Outputs.Pack(totalSupply)
	encTA, _ := abi.Methods["totalAssets"].Outputs.Pack(totalAssets)
	got, err := tm.ComputePrice([]multicall3.Result3{
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
