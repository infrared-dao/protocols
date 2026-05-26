package protocols

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"

	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/infrared-dao/protocols/multicall3"
)

func TestIVX_BatchPathMatchesLegacy(t *testing.T) {
	t.Parallel()
	p := &IVXLPPriceProvider{
		lpMonitorAddress: common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		lpTokenAddress:   common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"),
		logger:           zerolog.Nop(),
		config:           &IVXLPConfig{LPTDecimals: 18},
	}
	sharePrice := new(big.Int).Mul(big.NewInt(125), bigPow10(16)) // 1.25e18 == $1.25
	wantDirect := p.computeLPPriceFromReads(sharePrice)
	abi, _ := sc.IVXLPMonitorMetaData.GetAbi()
	encSP, _ := abi.Methods["getSharePrice"].Outputs.Pack(sharePrice)
	got, err := p.ComputePrice([]multicall3.Result3{
		{Success: true, ReturnData: encSP},
	})
	if err != nil {
		t.Fatalf("ComputePrice: %v", err)
	}
	if !got.Equal(wantDirect) {
		t.Errorf("batch %s != direct %s", got, wantDirect)
	}
}
