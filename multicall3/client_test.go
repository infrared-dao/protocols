package multicall3

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// fakeBackend records the last request it received and returns canned bytes.
// Only CallContract is needed for testing the Multicall3 client.
type fakeBackend struct {
	gotCallMsg     ethereum.CallMsg
	gotBlockNumber *big.Int
	resp           []byte
	err            error
	calls          int
}

func (f *fakeBackend) CallContract(_ context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	f.gotCallMsg = call
	f.gotBlockNumber = blockNumber
	f.calls++
	return f.resp, f.err
}

const erc20BalanceOfABI = `[{"inputs":[{"name":"account","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`

func mustParseABI(t *testing.T, raw string) *abi.ABI {
	t.Helper()
	parsed, err := abi.JSON(strings.NewReader(raw))
	if err != nil {
		t.Fatalf("parse abi: %v", err)
	}
	return &parsed
}

// packAggregate3Response builds the raw bytes that the Multicall3 contract
// would return for an aggregate3 call producing the given Result3 list.
// Used by the fake backend to stand in for a real on-chain response.
func packAggregate3Response(t *testing.T, results []Result3) []byte {
	t.Helper()
	a := mustParseABI(t, aggregate3ABI)
	method := a.Methods["aggregate3"]
	anon := make([]struct {
		Success    bool
		ReturnData []byte
	}, len(results))
	for i, r := range results {
		anon[i].Success = r.Success
		anon[i].ReturnData = r.ReturnData
	}
	packed, err := method.Outputs.Pack(anon)
	if err != nil {
		t.Fatalf("pack aggregate3 response: %v", err)
	}
	return packed
}

func TestNewClient_RejectsNilBackend(t *testing.T) {
	t.Parallel()
	_, err := NewClient(nil)
	if err == nil {
		t.Fatal("expected error for nil backend, got nil")
	}
}

func TestAggregate3_EmptyInputIsNoop(t *testing.T) {
	t.Parallel()
	fake := &fakeBackend{}
	c, err := NewClient(fake)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	results, err := c.Aggregate3(context.Background(), nil, nil)
	if err != nil {
		t.Fatalf("Aggregate3 empty: %v", err)
	}
	if results != nil {
		t.Fatalf("expected nil results, got %v", results)
	}
	if fake.calls != 0 {
		t.Fatalf("backend should not have been called; got %d calls", fake.calls)
	}
}

func TestAggregate3_HappyPath(t *testing.T) {
	t.Parallel()
	target := common.HexToAddress("0x1111111111111111111111111111111111111111")
	wantBalance := big.NewInt(123456789)

	erc20 := mustParseABI(t, erc20BalanceOfABI)
	encodedBalance, err := erc20.Methods["balanceOf"].Outputs.Pack(wantBalance)
	if err != nil {
		t.Fatalf("pack balanceOf return: %v", err)
	}
	respBytes := packAggregate3Response(t, []Result3{
		{Success: true, ReturnData: encodedBalance},
	})

	fake := &fakeBackend{resp: respBytes}
	c, err := NewClient(fake)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	block := big.NewInt(12345)
	holder := common.HexToAddress("0x2222222222222222222222222222222222222222")
	callData, err := erc20.Pack("balanceOf", holder)
	if err != nil {
		t.Fatalf("pack balanceOf: %v", err)
	}
	results, err := c.Aggregate3(context.Background(), []Call3{
		{Target: target, AllowFailure: true, CallData: callData},
	}, block)
	if err != nil {
		t.Fatalf("Aggregate3: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("got %d results, want 1", len(results))
	}
	if !results[0].Success {
		t.Errorf("Success = false, want true")
	}
	out, err := erc20.Methods["balanceOf"].Outputs.Unpack(results[0].ReturnData)
	if err != nil {
		t.Fatalf("unpack balance: %v", err)
	}
	got := out[0].(*big.Int)
	if got.Cmp(wantBalance) != 0 {
		t.Errorf("balance = %v, want %v", got, wantBalance)
	}

	if fake.gotCallMsg.To == nil || *fake.gotCallMsg.To != common.HexToAddress(CanonicalAddress) {
		t.Errorf("backend To = %v, want %v", fake.gotCallMsg.To, CanonicalAddress)
	}
	if fake.gotBlockNumber == nil || fake.gotBlockNumber.Cmp(block) != 0 {
		t.Errorf("backend blockNumber = %v, want %v", fake.gotBlockNumber, block)
	}
}

func TestAggregate3_PerCallFailureNotTopLevel(t *testing.T) {
	t.Parallel()
	respBytes := packAggregate3Response(t, []Result3{
		{Success: true, ReturnData: []byte{0x01}},
		{Success: false, ReturnData: []byte{0x02, 0x03}},
		{Success: true, ReturnData: []byte{0x04}},
	})
	fake := &fakeBackend{resp: respBytes}
	c, err := NewClient(fake)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	results, err := c.Aggregate3(context.Background(), []Call3{
		{Target: common.Address{1}, AllowFailure: true, CallData: []byte{0xaa}},
		{Target: common.Address{2}, AllowFailure: true, CallData: []byte{0xbb}},
		{Target: common.Address{3}, AllowFailure: true, CallData: []byte{0xcc}},
	}, nil)
	if err != nil {
		t.Fatalf("Aggregate3 unexpected err: %v", err)
	}
	if len(results) != 3 {
		t.Fatalf("got %d results, want 3", len(results))
	}
	if !results[0].Success || results[1].Success || !results[2].Success {
		t.Errorf("Success = %v %v %v, want true false true",
			results[0].Success, results[1].Success, results[2].Success)
	}
}

func TestAggregate3_BackendErrorPropagates(t *testing.T) {
	t.Parallel()
	fake := &fakeBackend{err: errors.New("network down")}
	c, err := NewClient(fake)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	_, err = c.Aggregate3(context.Background(), []Call3{
		{Target: common.Address{1}, AllowFailure: true, CallData: []byte{0xaa}},
	}, nil)
	if err == nil {
		t.Fatal("expected error from backend, got nil")
	}
	if !strings.Contains(err.Error(), "network down") {
		t.Errorf("err = %q, want it to contain 'network down'", err)
	}
}

func TestNewClientAt_UsesCustomAddress(t *testing.T) {
	t.Parallel()
	fake := &fakeBackend{resp: packAggregate3Response(t, []Result3{
		{Success: true, ReturnData: []byte{0x42}},
	})}
	custom := common.HexToAddress("0xDeaDBEEfDEaDbeEfDeAdbEEFdeadbeEFDEADbEeF")
	c, err := NewClientAt(fake, custom)
	if err != nil {
		t.Fatalf("NewClientAt: %v", err)
	}
	if c.Address() != custom {
		t.Errorf("Address() = %v, want %v", c.Address(), custom)
	}
	_, _ = c.Aggregate3(context.Background(), []Call3{
		{Target: common.Address{1}, AllowFailure: true, CallData: []byte{0x01}},
	}, nil)
	if fake.gotCallMsg.To == nil || *fake.gotCallMsg.To != custom {
		t.Errorf("backend was called with To = %v, want custom address %v", fake.gotCallMsg.To, custom)
	}
}
