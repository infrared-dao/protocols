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

func TestConcrete_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	assetPrice, _ := decimal.NewFromString("2.50")
	c := &ConcreteLPPriceProvider{
		address: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		logger:  zerolog.Nop(),
		config:  &ConcreteConfig{Asset: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", LPTDecimals: 18},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: assetPrice},
		},
	}
	totalSupply := new(big.Int).Mul(big.NewInt(40), bigPow10(18))
	totalAssets := new(big.Int).Mul(big.NewInt(100), bigPow10(18))
	wantDirect, _ := c.computeLPPriceFromReads(totalSupply, totalAssets)
	abi, _ := sc.ConcreteVaultMetaData.GetAbi()
	encTS, _ := abi.Methods["totalSupply"].Outputs.Pack(totalSupply)
	encTA, _ := abi.Methods["totalAssets"].Outputs.Pack(totalAssets)
	got, err := c.ComputePrice([]multicall3.Result3{
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

func TestConcrete_ZeroSupplyReturnsNotReady(t *testing.T) {
	t.Parallel()
	assetPrice, _ := decimal.NewFromString("1.00")
	c := &ConcreteLPPriceProvider{
		address: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		logger:  zerolog.Nop(),
		config:  &ConcreteConfig{Asset: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", LPTDecimals: 18},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: assetPrice},
		},
	}
	abi, _ := sc.ConcreteVaultMetaData.GetAbi()
	encTS, _ := abi.Methods["totalSupply"].Outputs.Pack(big.NewInt(0))
	encTA, _ := abi.Methods["totalAssets"].Outputs.Pack(big.NewInt(0))
	_, err := c.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTS},
		{Success: true, ReturnData: encTA},
	})
	if err != ErrPriceNotReadyYet {
		t.Fatalf("want ErrPriceNotReadyYet, got %v", err)
	}
}
