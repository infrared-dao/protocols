package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols"
	"github.com/infrared-dao/protocols/cmd/test/http"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func main() {
	// Create a zerolog logger
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// Command-line arguments
	poolIdArg := flag.String("poolid", "", "PancakeSwap Infinity Pool ID (bytes32 hex)")
	price0Arg := flag.String("price0", "", "address:price of token 0")
	price1Arg := flag.String("price1", "", "address:price of token 1")
	rpcURLArg := flag.String("rpcurl", "https://bsc-dataseed.binance.org/", "BNB Chain RPC URL")
	flag.Parse()

	// Example usage for $IR token pair on BNB Chain:
	// pancakeswapinfinity -poolid=0xd9b7a1d544432b0bacac321ccf543f3a23ae55de9709b2cf6664ab83f4cbe0a4
	//                     -price0=<token0-address>:<token0-price>
	//                     -price1=<token1-address>:<token1-price>
	//                     -rpcurl=https://bsc-dataseed.binance.org/

	// Validate required arguments
	missingArgs := []string{}
	if *poolIdArg == "" {
		missingArgs = append(missingArgs, "poolid")
	}
	if *price0Arg == "" {
		missingArgs = append(missingArgs, "price0")
	}
	if *price1Arg == "" {
		missingArgs = append(missingArgs, "price1")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -poolid <pool-id> -price0 <addr:price> -price1 <addr:price> -rpcurl <rpc-url>").
			Msg("Missing required arguments")
	}

	// Parse prices - format: address:price:name:decimals (name and decimals optional)
	p0data := strings.Split(*price0Arg, ":")
	if len(p0data) < 2 {
		logger.Fatal().Msgf("Invalid price0 format, expected 'address:price[:name:decimals]', got '%s'", *price0Arg)
	}
	price0, err := decimal.NewFromString(p0data[1])
	if err != nil {
		logger.Fatal().Err(err).Str("price0", *price0Arg).Msg("Invalid price0")
	}
	name0 := "Token0"
	decimals0 := uint(18)
	if len(p0data) >= 3 {
		name0 = p0data[2]
	}
	if len(p0data) >= 4 {
		var d int
		_, _ = fmt.Sscanf(p0data[3], "%d", &d)
		decimals0 = uint(d)
	}

	p1data := strings.Split(*price1Arg, ":")
	if len(p1data) < 2 {
		logger.Fatal().Msgf("Invalid price1 format, expected 'address:price[:name:decimals]', got '%s'", *price1Arg)
	}
	price1, err := decimal.NewFromString(p1data[1])
	if err != nil {
		logger.Fatal().Err(err).Str("price1", *price1Arg).Msg("Invalid price1")
	}
	name1 := "Token1"
	decimals1 := uint(18)
	if len(p1data) >= 3 {
		name1 = p1data[2]
	}
	if len(p1data) >= 4 {
		var d int
		_, _ = fmt.Sscanf(p1data[3], "%d", &d)
		decimals1 = uint(d)
	}

	pmap := map[string]protocols.Price{
		strings.ToLower(p0data[0]): {Decimals: decimals0, Price: price0, TokenName: name0},
		strings.ToLower(p1data[0]): {Decimals: decimals1, Price: price1, TokenName: name1},
	}

	ctx := context.Background()
	// Connect to the BNB Chain client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to BNB Chain client")
	}

	// Get the LP price provider config
	cp := protocols.PancakeSwapInfinityLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *poolIdArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for PancakeSwapInfinityLPPriceProvider")
	}

	configData := &protocols.PancakeSwapInfinityConfig{}
	err = json.Unmarshal(configBytes, configData)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to deserialize config")
	}
	logger.Info().
		Str("PancakeSwap Infinity Config", fmt.Sprintf("%+v", configData)).
		Msg("successfully created PancakeSwap Infinity Config")

	// Create a new PancakeSwapInfinityLPPriceProvider
	// Note: address parameter is not used for Infinity pools, poolId comes from config
	provider := protocols.NewPancakeSwapInfinityLPPriceProvider(common.Address{}, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client, http.NewTestHttpClient())
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize PancakeSwapInfinityLPPriceProvider")
	}

	// Fetch LP token price (note: CLMM pools don't have traditional LP tokens)
	lpPrice, err := provider.LPTokenPrice(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch LP token price")
	} else {
		logger.Info().
			Str("LPTokenPrice (USD)", lpPrice).
			Msg("LP token price (note: CLMM pools use NFT positions)")
	}

	// Fetch TVL
	tvl, err := provider.TVL(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch TVL")
	} else {
		logger.Info().
			Str("TVL (USD)", tvl).
			Msg("successfully fetched TVL")
	}

	// Fetch TVL Breakdown
	breakdown, err := provider.TVLBreakdown(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch TVL breakdown")
	} else {
		logger.Info().Msg("successfully fetched TVL breakdown")
		for tokenAddr, tokenTVL := range breakdown {
			logger.Info().
				Str("token", tokenAddr).
				Str("symbol", tokenTVL.TokenSymbol).
				Str("amount", tokenTVL.Amount.String()).
				Str("usdValue", tokenTVL.USDValue.String()).
				Str("ratio", tokenTVL.Ratio.String()).
				Msg("token breakdown")
		}
	}
}
