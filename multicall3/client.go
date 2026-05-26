// Package multicall3 wraps the on-chain Multicall3 contract
// (https://github.com/mds1/multicall) deployed at the deterministic CREATE2
// address `0xcA11bde05977b3631167028862bE2a173976CA11` on every supported
// EVM chain.
//
// The contract accepts an array of (target, allowFailure, callData) tuples
// and atomically returns each call's (success, returnData) pair — all
// observed at the same block. From a provider's perspective the whole
// invocation is one `eth_call`, not N, which is the key win over per-call
// dispatch and over JSON-RPC batching (where the provider counts each
// sub-call individually).
//
// This package is intentionally minimal: it provides the on-chain
// invocation primitive (Aggregate3) and the typed Call3/Result3 structs.
// The library's higher-level BatchLPTokenPrice (see ../batch.go) is the
// idiomatic entry point for callers.
package multicall3

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// CanonicalAddress is the CREATE2-deterministic deployment address of the
// Multicall3 contract on every supported EVM chain (including Berachain
// mainnet). Reference: https://github.com/mds1/multicall.
const CanonicalAddress = "0xcA11bde05977b3631167028862bE2a173976CA11"

// Call3 is one entry in an aggregate3 invocation. Fields are ABI-tagged to
// match the on-chain struct so go-ethereum's reflect-based packer accepts
// the value directly.
type Call3 struct {
	Target       common.Address `abi:"target"`
	AllowFailure bool           `abi:"allowFailure"`
	CallData     []byte         `abi:"callData"`
}

// Result3 is the per-call output of aggregate3.
type Result3 struct {
	Success    bool   `abi:"success"`
	ReturnData []byte `abi:"returnData"`
}

// Backend is the minimum eth-client surface area required to call
// Multicall3.aggregate3. Standard go-ethereum eth client implementations
// (ethclient.Client, simulated.Backend, etc.) satisfy this interface; the
// narrow definition keeps testing simple.
type Backend interface {
	CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
}

// Client invokes aggregate3 against a deployed Multicall3 contract.
//
// Construct via NewClient (uses CanonicalAddress) or NewClientAt (for chains
// where Multicall3 has been deployed at a non-canonical address).
type Client struct {
	address common.Address
	abi     *abi.ABI
	backend Backend
}

// NewClient creates a Client targeting CanonicalAddress.
func NewClient(backend Backend) (*Client, error) {
	return NewClientAt(backend, common.HexToAddress(CanonicalAddress))
}

// NewClientAt creates a Client targeting a specific Multicall3 deployment
// address. Useful for chains where Multicall3 is deployed at a non-canonical
// address.
func NewClientAt(backend Backend, address common.Address) (*Client, error) {
	if backend == nil {
		return nil, errors.New("multicall3: backend must not be nil")
	}
	parsed, err := abi.JSON(strings.NewReader(aggregate3ABI))
	if err != nil {
		return nil, fmt.Errorf("multicall3: parse ABI: %w", err)
	}
	return &Client{
		address: address,
		abi:     &parsed,
		backend: backend,
	}, nil
}

// Address returns the deployed Multicall3 contract address this client
// targets.
func (c *Client) Address() common.Address { return c.address }

// Aggregate3 dispatches the given calls as a single eth_call to Multicall3's
// aggregate3 method and returns the per-call results.
//
// blockNumber may be nil (latest). All calls share the same block — that's
// the entire point of using Multicall3: atomic snapshots.
//
// Per-call failures are NOT errors at this level — they appear as
// Result3{Success: false, ReturnData: <revert-data>}. The caller is
// responsible for inspecting Success per result and acting on the revert
// data accordingly. Set Call3.AllowFailure = false to make the entire
// aggregate3 revert on any sub-call failure (rarely the right choice for
// batched reads).
//
// Errors returned from this function are network/encoding errors that
// affect the entire batch.
func (c *Client) Aggregate3(ctx context.Context, calls []Call3, blockNumber *big.Int) ([]Result3, error) {
	if len(calls) == 0 {
		return nil, nil
	}

	input, err := c.abi.Pack("aggregate3", calls)
	if err != nil {
		return nil, fmt.Errorf("multicall3: pack aggregate3 input: %w", err)
	}

	out, err := c.backend.CallContract(ctx, ethereum.CallMsg{
		To:   &c.address,
		Data: input,
	}, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("multicall3: aggregate3 call: %w", err)
	}

	unpacked, err := c.abi.Unpack("aggregate3", out)
	if err != nil {
		return nil, fmt.Errorf("multicall3: unpack aggregate3 output: %w", err)
	}
	if len(unpacked) != 1 {
		return nil, fmt.Errorf("multicall3: expected 1 output value, got %d", len(unpacked))
	}

	return convertResults(unpacked[0], len(calls))
}

// convertResults adapts the anonymous-struct slice go-ethereum's ABI decoder
// produces for aggregate3 into the typed Result3 slice we export. Defensive
// length + shape validation — a mismatch here indicates either Multicall3
// contract version drift or an ABI-configuration bug, both surfaced as
// errors rather than silently mangled.
func convertResults(raw any, expectedLen int) ([]Result3, error) {
	anon, ok := raw.([]struct {
		Success    bool    `json:"success"`
		ReturnData []uint8 `json:"returnData"`
	})
	if !ok {
		return nil, fmt.Errorf("multicall3: unexpected aggregate3 output type %T", raw)
	}
	if len(anon) != expectedLen {
		return nil, fmt.Errorf(
			"multicall3: aggregate3 returned %d results for %d calls", len(anon), expectedLen,
		)
	}
	out := make([]Result3, len(anon))
	for i, r := range anon {
		out[i] = Result3{Success: r.Success, ReturnData: r.ReturnData}
	}
	return out, nil
}

// aggregate3ABI is the minimum ABI needed to encode aggregate3 invocations
// and decode aggregate3 responses. Hand-written rather than generated so
// this sub-package is self-contained and doesn't pull in a full abigen
// output file.
//
// Source: https://github.com/mds1/multicall/blob/main/src/Multicall3.sol —
// the `aggregate3` function signature.
const aggregate3ABI = `[
  {
    "inputs": [
      {
        "components": [
          {"internalType": "address", "name": "target",       "type": "address"},
          {"internalType": "bool",    "name": "allowFailure", "type": "bool"},
          {"internalType": "bytes",   "name": "callData",     "type": "bytes"}
        ],
        "internalType": "struct Multicall3.Call3[]",
        "name": "calls",
        "type": "tuple[]"
      }
    ],
    "name": "aggregate3",
    "outputs": [
      {
        "components": [
          {"internalType": "bool",  "name": "success",    "type": "bool"},
          {"internalType": "bytes", "name": "returnData", "type": "bytes"}
        ],
        "internalType": "struct Multicall3.Result[]",
        "name": "returnData",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "payable",
    "type": "function"
  }
]`
