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

func TestBurrBear_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	tokenAAddr := common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	tokenBAddr := common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	pA, _ := decimal.NewFromString("2.50")
	pB, _ := decimal.NewFromString("1.00")
	bb := &BurrBearLPPriceProvider{
		poolAddress: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		logger:      zerolog.Nop(),
		config: &BurrBearPoolConfig{
			VaultContract: "0x2222222222222222222222222222222222222222",
			PoolID:        [32]byte{0x01, 0x02, 0x03},
			LPTDecimals:   18,
		},
		priceMap: map[string]Price{
			"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {Decimals: 18, Price: pA},
			"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {Decimals: 18, Price: pB},
		},
	}
	totalSupply := new(big.Int).Mul(big.NewInt(50), bigPow10(18))
	balA := new(big.Int).Mul(big.NewInt(20), bigPow10(18))
	balB := new(big.Int).Mul(big.NewInt(75), bigPow10(18))
	wantDirect, err := bb.computeLPPriceFromReads(totalSupply, map[string]*big.Int{
		"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": balA,
		"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": balB,
	})
	if err != nil {
		t.Fatalf("computeLPPriceFromReads: %v", err)
	}
	poolABI, _ := sc.BalancerBasePoolMetaData.GetAbi()
	encTotalSupply, _ := poolABI.Methods["getActualSupply"].Outputs.Pack(totalSupply)
	vaultABI, _ := sc.BalancerVaultMetaData.GetAbi()
	encPoolTokens, err := vaultABI.Methods["getPoolTokens"].Outputs.Pack(
		[]common.Address{tokenAAddr, tokenBAddr},
		[]*big.Int{balA, balB},
		big.NewInt(0),
	)
	if err != nil {
		t.Fatalf("pack getPoolTokens: %v", err)
	}
	got, err := bb.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encTotalSupply},
		{Success: true, ReturnData: encPoolTokens},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !got.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s", got, wantDirect)
	}
}
