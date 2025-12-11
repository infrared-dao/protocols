package protocols

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/infrared-dao/protocols/internal/sc"
	"github.com/rs/zerolog"
	// "github.com/shopspring/decimal"
)

type PaddleFiConfig struct {
	PaddlefiContract string `json:"paddlefi_contract"`
	WBERAAddress     string `json:"wbera_address"`
	Decimals         uint   `json:"wbera_decimals"`
	NFTAddress       string `json:"nft_address"`
}

type PaddleFiProvider struct {
	poolAddress      common.Address
	block            *big.Int
	priceMap         map[string]Price
	logger           zerolog.Logger
	configBytes      []byte
	config           *PaddleFiConfig
	wberaContract    *sc.ERC20
	nftContract      *sc.ERC20
	paddlefiContract *sc.Paddlefi
}

func NewPaddleFiProvider(
	poolAddress common.Address,
	block *big.Int,
	prices map[string]Price,
	logger zerolog.Logger,
	config []byte,
) *PaddleFiProvider {
	return &PaddleFiProvider{
		poolAddress: poolAddress,
		block:       block,
		priceMap:    prices,
		logger:      logger,
		configBytes: config,
	}
}

func (a *PaddleFiProvider) Initialize(ctx context.Context, client bind.ContractBackend) error {
	a.config = &PaddleFiConfig{}
	if err := json.Unmarshal(a.configBytes, a.config); err != nil {
		return fmt.Errorf("config unmarshal failed: %w", err)
	}

	var err error
	a.wberaContract, err = sc.NewERC20(common.HexToAddress(a.config.WBERAAddress), client)
	if err != nil {
		return fmt.Errorf("failed to bind WBERA contract: %w", err)
	}

	a.nftContract, err = sc.NewERC20(common.HexToAddress(a.config.NFTAddress), client)
	if err != nil {
		return fmt.Errorf("failed to instantiate ERC721/NFT contract: %w", err)
	}

	a.paddlefiContract, err = sc.NewPaddlefi(common.HexToAddress(a.config.PaddlefiContract), client)
	if err != nil {
		return fmt.Errorf("failed to instantiate PaddleFi contract: %w", err)
	}

	return nil
}

func (a *PaddleFiProvider) LPTokenPrice(ctx context.Context) (string, error) {
	return "", fmt.Errorf("LPTokenPrice not used in PaddleFiProvider")
}

func (a *PaddleFiProvider) TVL(ctx context.Context) (string, error) {
	opts := &bind.CallOpts{Context: ctx, BlockNumber: a.block}

	// WBERA balance in the pool
	wberaBalance, err := a.wberaContract.BalanceOf(opts, a.poolAddress)
	if err != nil {
		return "", fmt.Errorf("failed to get balance: %w", err)
	}

	// NFT balance in the pool
	nftBalance, err := a.nftContract.BalanceOf(opts, a.poolAddress)
	if err != nil {
		return "", fmt.Errorf("failed to fetch NFT balance: %w", err)
	}

	// nft price
	cTokens := []common.Address{a.poolAddress}
	result, err := a.paddlefiContract.CTokenUnderlyingPriceAll(opts, cTokens)
	if err != nil || len(result) == 0 {
		return "", fmt.Errorf("failed to call cTokenUnderlyingPriceAll: %w", err)
	}
	nftPriceRaw := result[0].UnderlyingNFTPrice
	nftPrice := NormalizeAmount(nftPriceRaw, a.config.Decimals)

	wberaAmount := NormalizeAmount(wberaBalance, a.config.Decimals)
	nftAmount := NormalizeAmount(nftBalance, 0)
	nftTotalValueInWbera := nftAmount.Mul(nftPrice)
	totalWbera := wberaAmount.Add(nftTotalValueInWbera)

	fmt.Printf("WBERA balance in pool: %s\n", wberaAmount)
	fmt.Printf("NFT amount in pool: %s\n", nftAmount)
	fmt.Printf("NFT value in pool: %s\n", nftTotalValueInWbera)
	fmt.Printf("NFT total value in WBERA: %s\n", totalWbera)

	priceInfo, ok := a.priceMap[strings.ToLower(a.config.WBERAAddress)]

	if !ok {
		return "", fmt.Errorf("no price found for WBERA")
	}

	tvl := totalWbera.Mul(priceInfo.Price)
	return tvl.StringFixed(6), nil
}

func (a *PaddleFiProvider) GetConfig(ctx context.Context, poolAddress string, client bind.ContractBackend) ([]byte, error) {
	config := PaddleFiConfig{
		PaddlefiContract: "0xAE9D7a721534907894eBBFf92b5eEB813c770Ba6",
		WBERAAddress:     "0x6969696969696969696969696969696969696969",
		NFTAddress:       "0x88888888A9361f15AAdBAca355A6B2938C6A674e",
		Decimals:         18,
	}
	return json.Marshal(config)
}

func (a *PaddleFiProvider) UpdateBlock(block *big.Int, prices map[string]Price) {
	a.block = block
	if prices != nil {
		a.priceMap = prices
	}
}

// TVLBreakdown returns the breakdown of TVL by underlying tokens.
// TODO: Implement TVL breakdown for PaddleFi protocol
func (a *PaddleFiProvider) TVLBreakdown(ctx context.Context) (map[string]TokenTVL, error) {
	return nil, ErrTVLBreakdownNotImplemented
}
