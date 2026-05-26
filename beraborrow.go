package protocols

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/infrared-dao/protocols/fetchers"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/infrared-dao/protocols/multicall3"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

var _ Protocol = &BeraBorrowLPPriceProvider{}
var _ BatchablePriceProvider = &BeraBorrowLPPriceProvider{}

// This adapter is for a very specific type of BeraBorrow token setup
// InfraredWrapper type tokens are the LP Token in this case
// InfraredWrappers are 1-1 with the CompoundingInfraredCollateralVault token type

// Extending it to also work with the singleton sNECT token important to beraborrow
// In the case of sNECT there is no ColVaultAddress so used to signal this is sNECT
const sNECTName string = "Staked Nectar"

// The fetchPrice() onchain call always returns a USD price with 18 decimals
const USDPriceDecimals uint = 18

// LPTDecimals are the decimals of the InfraredWrapper token
// CDPDecimals are the decimals of the CompoundingInfraredCollateralVault token
type BeraBorrowCDPConfig struct {
	ColVaultAddress string `json:"col_vault_address"`
	LPTDecimals     uint   `json:"lpt_decimals"`
	CDPDecimals     uint   `json:"cdp_decimals"`
}

// BeraBorrowLPPriceProvider defines the provider for BeraBorrow CDP price and CDP TVL.
type BeraBorrowLPPriceProvider struct {
	LPTAddress    common.Address
	block         *big.Int
	logger        zerolog.Logger
	configBytes   []byte
	config        *BeraBorrowCDPConfig
	lptContract   *sc.BeraBorrowIWCaller
	cdpContract   *sc.BeraBorrowCICVCaller
	snectContract *sc.BeraBorrowSNECTCaller

	// snectAsset is cached after Initialize for the sNECT variant — needed
	// as an argument to GetPrice. Caching here lets PriceReads describe its
	// reads without a chained dependency (asset doesn't change for a given
	// sNECT contract).
	snectAsset common.Address
}

// NewBeraBorrowLPPriceProvider creates a new instance of the BeraBorrowLPPriceProvider.
func NewBeraBorrowLPPriceProvider(
	LPTAddress common.Address,
	block *big.Int,
	logger zerolog.Logger,
	config []byte,
) *BeraBorrowLPPriceProvider {
	b := &BeraBorrowLPPriceProvider{
		LPTAddress:  LPTAddress,
		block:       block,
		logger:      logger,
		configBytes: config,
	}
	return b
}

// Initialize checks the configuration/data and instantiates the CDP contract.
func (b *BeraBorrowLPPriceProvider) Initialize(ctx context.Context, client bind.ContractBackend, _ fetchers.HttpClient) error {
	var err error

	b.config = &BeraBorrowCDPConfig{}
	err = json.Unmarshal(b.configBytes, b.config)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to deserialize config")
		return err
	}

	// Initialize the IW contract from the address in the inputs for the LP token address
	b.lptContract, err = sc.NewBeraBorrowIWCaller(b.LPTAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("adapter init failed to instantiate Infrared Wrapper contract")
		return err
	}

	// Initialize either the sNECT contract or the CICV contract but not both
	if b.config.ColVaultAddress == sNECTName {
		// Initialize the sNECT contract but no need for the CICV contract
		b.snectContract, err = sc.NewBeraBorrowSNECTCaller(b.LPTAddress, client)
		if err != nil {
			b.logger.Error().Err(err).Msg("adapter init failed to instantiate BeraBorrow sNECT contract")
			return err
		}

		// Cache the asset address (immutable for a given sNECT contract) so
		// PriceReads can describe GetPrice(asset) without a chained read.
		asset, err := b.snectContract.Asset(&bind.CallOpts{Context: ctx})
		if err != nil {
			b.logger.Error().Err(err).Msg("adapter init failed to fetch sNECT asset")
			return err
		}
		b.snectAsset = asset
	} else {
		// Initialize the CICV contract from the address in the config -- not the IW contract from the LP token address
		b.cdpContract, err = sc.NewBeraBorrowCICVCaller(common.HexToAddress(b.config.ColVaultAddress), client)
		if err != nil {
			b.logger.Error().Err(err).Msg("adapter init failed to instantiate Compounding Infrared Collateral Vault contract")
			return err
		}
	}

	return nil
}

// LPTokenPrice returns the current price of the protocol's LP token in USD.
//
// Legacy path — issues sequential eth_calls (3 per variant). The batchable
// path (PriceReads + ComputePrice) dispatches them as one
// Multicall3.aggregate3. Both paths share computeLPPriceFromReads so
// prices match for the same inputs.
func (b *BeraBorrowLPPriceProvider) LPTokenPrice(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: b.block,
	}

	// returns a *big.Int for total supply of the LP token
	lpTotalSupply, err := b.lptContract.TotalSupply(opts)
	if err != nil {
		return "", fmt.Errorf("failed to get beraborrow total supply, err: %w", err)
	}

	tvl, err := b.getTotalValue(ctx)
	if err != nil {
		return "", err
	}

	price, err := b.computeLPPriceFromReads(lpTotalSupply, tvl)
	if err != nil {
		return "", err
	}
	return price.StringFixed(roundingDecimals), nil
}

// PriceReads describes the eth_calls needed to compute the LP token price.
// Returns 3 calls in both variants:
//   - sNECT: lptContract.TotalSupply, snectContract.GetPrice(asset), snectContract.TotalAssets
//   - CICV:  lptContract.TotalSupply, cdpContract.FetchPrice,       cdpContract.TotalSupply
func (b *BeraBorrowLPPriceProvider) PriceReads() ([]multicall3.Call3, error) {
	iwABI, err := sc.BeraBorrowIWMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("beraborrow: get IW ABI: %w", err)
	}
	lpTotalSupplyData, err := iwABI.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("beraborrow: pack IW totalSupply: %w", err)
	}

	if b.config.ColVaultAddress == sNECTName {
		snectABI, err := sc.BeraBorrowSNECTMetaData.GetAbi()
		if err != nil {
			return nil, fmt.Errorf("beraborrow: get sNECT ABI: %w", err)
		}
		getPriceData, err := snectABI.Pack("getPrice", b.snectAsset)
		if err != nil {
			return nil, fmt.Errorf("beraborrow: pack sNECT getPrice: %w", err)
		}
		totalAssetsData, err := snectABI.Pack("totalAssets")
		if err != nil {
			return nil, fmt.Errorf("beraborrow: pack sNECT totalAssets: %w", err)
		}
		return []multicall3.Call3{
			{Target: b.LPTAddress, AllowFailure: true, CallData: lpTotalSupplyData},
			{Target: b.LPTAddress, AllowFailure: true, CallData: getPriceData},    // sNECT is at LPTAddress
			{Target: b.LPTAddress, AllowFailure: true, CallData: totalAssetsData}, // sNECT is at LPTAddress
		}, nil
	}

	// CICV variant
	cicvABI, err := sc.BeraBorrowCICVMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("beraborrow: get CICV ABI: %w", err)
	}
	fetchPriceData, err := cicvABI.Pack("fetchPrice")
	if err != nil {
		return nil, fmt.Errorf("beraborrow: pack CICV fetchPrice: %w", err)
	}
	cdpTotalSupplyData, err := cicvABI.Pack("totalSupply")
	if err != nil {
		return nil, fmt.Errorf("beraborrow: pack CICV totalSupply: %w", err)
	}
	cicvAddress := common.HexToAddress(b.config.ColVaultAddress)
	return []multicall3.Call3{
		{Target: b.LPTAddress, AllowFailure: true, CallData: lpTotalSupplyData},
		{Target: cicvAddress, AllowFailure: true, CallData: fetchPriceData},
		{Target: cicvAddress, AllowFailure: true, CallData: cdpTotalSupplyData},
	}, nil
}

// ComputePrice decodes the responses from PriceReads and computes the LP
// token price. Variant-aware: differs between sNECT and CICV per the
// variant taken in Initialize.
func (b *BeraBorrowLPPriceProvider) ComputePrice(responses []multicall3.Result3) (decimal.Decimal, error) {
	if len(responses) != 3 {
		return decimal.Zero, fmt.Errorf("beraborrow: expected 3 responses, got %d", len(responses))
	}
	for i, r := range responses {
		if !r.Success {
			return decimal.Zero, fmt.Errorf("beraborrow: sub-call %d reverted in multicall", i)
		}
	}

	// First response is the LP totalSupply (same for both variants).
	iwABI, err := sc.BeraBorrowIWMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("beraborrow: get IW ABI: %w", err)
	}
	tsOut, err := iwABI.Methods["totalSupply"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("beraborrow: unpack IW totalSupply: %w", err)
	}
	lpTotalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("beraborrow: IW totalSupply value type %T, want *big.Int", tsOut[0])
	}

	// Compute TVL from the variant-specific responses.
	var tvl decimal.Decimal
	if b.config.ColVaultAddress == sNECTName {
		tvl, err = b.computeSNECTTotalValue(responses[1:])
	} else {
		tvl, err = b.computeCICVTotalValue(responses[1:])
	}
	if err != nil {
		return decimal.Zero, err
	}

	return b.computeLPPriceFromReads(lpTotalSupply, tvl)
}

// computeSNECTTotalValue decodes the sNECT-variant responses (GetPrice +
// TotalAssets) and computes the TVL.
func (b *BeraBorrowLPPriceProvider) computeSNECTTotalValue(responses []multicall3.Result3) (decimal.Decimal, error) {
	snectABI, err := sc.BeraBorrowSNECTMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("beraborrow: get sNECT ABI: %w", err)
	}

	priceOut, err := snectABI.Methods["getPrice"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("beraborrow: unpack sNECT getPrice: %w", err)
	}
	pricePerToken18, ok := priceOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("beraborrow: sNECT price value type %T, want *big.Int", priceOut[0])
	}
	pricePerToken := NormalizeAmount(pricePerToken18, USDPriceDecimals)

	assetsOut, err := snectABI.Methods["totalAssets"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("beraborrow: unpack sNECT totalAssets: %w", err)
	}
	totalAssets, ok := assetsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("beraborrow: sNECT totalAssets value type %T, want *big.Int", assetsOut[0])
	}
	numTokens := NormalizeAmount(totalAssets, b.config.CDPDecimals)

	return numTokens.Mul(pricePerToken), nil
}

// computeCICVTotalValue decodes the CICV-variant responses (FetchPrice +
// TotalSupply) and computes the TVL.
func (b *BeraBorrowLPPriceProvider) computeCICVTotalValue(responses []multicall3.Result3) (decimal.Decimal, error) {
	cicvABI, err := sc.BeraBorrowCICVMetaData.GetAbi()
	if err != nil {
		return decimal.Zero, fmt.Errorf("beraborrow: get CICV ABI: %w", err)
	}

	priceOut, err := cicvABI.Methods["fetchPrice"].Outputs.Unpack(responses[0].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("beraborrow: unpack CICV fetchPrice: %w", err)
	}
	pricePerToken18, ok := priceOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("beraborrow: CICV price value type %T, want *big.Int", priceOut[0])
	}
	pricePerToken := NormalizeAmount(pricePerToken18, USDPriceDecimals)

	tsOut, err := cicvABI.Methods["totalSupply"].Outputs.Unpack(responses[1].ReturnData)
	if err != nil {
		return decimal.Zero, fmt.Errorf("beraborrow: unpack CICV totalSupply: %w", err)
	}
	cdpTotalSupply, ok := tsOut[0].(*big.Int)
	if !ok {
		return decimal.Zero, fmt.Errorf("beraborrow: CICV totalSupply value type %T, want *big.Int", tsOut[0])
	}
	numTokens := NormalizeAmount(cdpTotalSupply, b.config.CDPDecimals)

	return numTokens.Mul(pricePerToken), nil
}

// computeLPPriceFromReads is the shared math between legacy LPTokenPrice
// and batchable ComputePrice. Takes the LP totalSupply and the computed
// TVL (which differs by variant), returns price per LP token.
func (b *BeraBorrowLPPriceProvider) computeLPPriceFromReads(lpTotalSupply *big.Int, tvl decimal.Decimal) (decimal.Decimal, error) {
	if lpTotalSupply.Sign() == 0 {
		err := errors.New("LP totalSupply is zero, cannot calculate LP token price")
		b.logger.Error().Err(err).Msg("Invalid totalSupply")
		return decimal.Zero, err
	}
	numTokens := NormalizeAmount(lpTotalSupply, b.config.LPTDecimals)
	pricePerToken := tvl.Div(numTokens)

	b.logger.Debug().
		Str("pricePerToken", pricePerToken.String()).
		Msg("LP token price calculated successfully")

	return pricePerToken, nil
}

// TVL returns the Total Value Locked in the CDP in USD
// Actually gets the TVL of the CICV which is 1-1 with the TVL of the IW which is the LP token
func (b *BeraBorrowLPPriceProvider) TVL(ctx context.Context) (string, error) {
	tvl, err := b.getTotalValue(ctx)
	if err != nil {
		return "", err
	}

	return tvl.StringFixed(roundingDecimals), nil
}

func (b *BeraBorrowLPPriceProvider) GetConfig(ctx context.Context, lpAddress string, client bind.ContractBackend) ([]byte, error) {
	var err error
	if !common.IsHexAddress(lpAddress) {
		err = fmt.Errorf("invalid smart contract address, '%s'", lpAddress)
		return nil, err
	}

	bbcc := BeraBorrowCDPConfig{}
	opts := &bind.CallOpts{
		Context: ctx,
	}

	erc20Contract, err := sc.NewERC20Caller(common.HexToAddress(lpAddress), client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instantiate ERC20 contract on LP Token")
		return nil, err
	}

	name, err := erc20Contract.Name(opts)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to get ERC20 name from LP Token")
		return nil, err
	}

	// For the special sNECT token, different logic is needed for each function and config
	if name == sNECTName {
		snectDecimals, err := erc20Contract.Decimals(opts)
		if err != nil {
			b.logger.Error().Err(err).Msg("failed to get ERC20 decimals from sNECT LP Token")
			return nil, err
		}
		bbcc.LPTDecimals = uint(snectDecimals)

		snectContract, err := sc.NewBeraBorrowSNECTCaller(common.HexToAddress(lpAddress), client)
		if err != nil {
			b.logger.Error().Err(err).Msg("adapter init failed to instantiate BeraBorrow sNECT contract")
			return nil, err
		}

		asset, err := snectContract.Asset(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to get underlying asset in sNECT config, err: %w", err)
		}

		nectContract, err := sc.NewERC20Caller(asset, client)
		if err != nil {
			b.logger.Error().Err(err).Msg("failed to instantiate ERC20 contract on asset token NECT")
			return nil, err
		}

		nectDecimals, err := nectContract.Decimals(opts)
		if err != nil {
			b.logger.Error().Err(err).Msg("failed to get ERC20 decimals from sNECT LP Token")
			return nil, err
		}
		bbcc.CDPDecimals = uint(nectDecimals)

		// using ColVaultAddress to signal to other functions this is the special sNECT token
		bbcc.ColVaultAddress = sNECTName

		body, err := json.Marshal(bbcc)
		if err != nil {
			return nil, err
		}

		return body, nil
	}

	// Normal config logic for all other beraborrow tokens other than sNECT

	iwContract, err := sc.NewBeraBorrowIWCaller(common.HexToAddress(lpAddress), client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instantiate Infrared Wrapper contract on LP Token")
		return nil, err
	}

	// decimals is uint8
	lptDecimals, err := iwContract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for LP token %s, %v", lpAddress, err)
		return nil, err
	}
	bbcc.LPTDecimals = uint(lptDecimals)

	icvAddress, err := iwContract.InfraredCollVault(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain the InfraredCollVault addressfor LP token %s, %v", lpAddress, err)
		return nil, err
	}
	bbcc.ColVaultAddress = icvAddress.Hex()

	// Initialize the CICV contract from address gotten from the IW contract
	cdpContract, err := sc.NewBeraBorrowCICVCaller(icvAddress, client)
	if err != nil {
		b.logger.Error().Err(err).Msg("failed to instantiate Compounding Infrared Collateral Vault contract for config")
		return nil, err
	}

	// decimals is uint8
	cdpDecimals, err := cdpContract.Decimals(opts)
	if err != nil {
		err = fmt.Errorf("failed to obtain number of decimals for CDP (CICV) token %s, %v", icvAddress.Hex(), err)
		return nil, err
	}
	bbcc.CDPDecimals = uint(cdpDecimals)

	body, err := json.Marshal(bbcc)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (b *BeraBorrowLPPriceProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	b.block = block
	// prices param can be ignored because beraborrow uses pricing onchain
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for BeraBorrow protocol
func (b *BeraBorrowLPPriceProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}

// Internal Helper methods not able to be called except in this file

func (b *BeraBorrowLPPriceProvider) getTotalValue(ctx context.Context) (decimal.Decimal, error) {
	var value decimal.Decimal
	var err error

	if b.config.ColVaultAddress == sNECTName {
		value, err = b.sNECTTotalValue(ctx)
	} else {
		value, err = b.cicvTotalValue(ctx)
	}

	return value, err
}

func (b *BeraBorrowLPPriceProvider) sNECTTotalValue(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: b.block,
	}

	asset, err := b.snectContract.Asset(opts)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to get underlying asset for sNECT, err: %w", err)
	}

	pricePerToken18, err := b.snectContract.GetPrice(opts, asset)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to get underlying token price for sNECT, err: %w", err)
	}
	pricePerToken := NormalizeAmount(pricePerToken18, USDPriceDecimals)

	totalAssets, err := b.snectContract.TotalAssets(opts)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to get beraborrow total supply from CICV contract, err: %w", err)
	}
	numTokens := NormalizeAmount(totalAssets, b.config.CDPDecimals)

	totalValue := numTokens.Mul(pricePerToken)

	return totalValue, nil
}

func (b *BeraBorrowLPPriceProvider) cicvTotalValue(ctx context.Context) (decimal.Decimal, error) {
	opts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: b.block,
	}

	pricePerToken18, err := b.cdpContract.FetchPrice(opts)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to fetchPrice from beraborrow CICV contract, err: %w", err)
	}
	pricePerToken := NormalizeAmount(pricePerToken18, USDPriceDecimals)

	cdpTotalSupply, err := b.cdpContract.TotalSupply(opts)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to get beraborrow total supply from CICV contract, err: %w", err)
	}
	numTokens := NormalizeAmount(cdpTotalSupply, b.config.CDPDecimals)

	cdpTotalValue := numTokens.Mul(pricePerToken)

	return cdpTotalValue, nil
}
