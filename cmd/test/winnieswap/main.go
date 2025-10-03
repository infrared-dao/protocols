package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/infrared-dao/protocols"
	"github.com/infrared-dao/protocols/fetchers"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func main() {
	// Create a zerolog logger
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// Command-line arguments
	addressArg := flag.String("address", "0x46fbf6ff1fd62ec89af48c3bb0b63115052dab31", "Smart contract address")
	price0Arg := flag.String("price0", "0x7507c1dc16935B82698e4C63f2746A2fCf994dF8:8.0", "address:price of token 0 (wBERA)")
	price1Arg := flag.String("price1", "0x0E4aaF1351de4c0264C5c7056Ef3777b41BD8e03:1.0", "address:price of token 1 (HONEY)")
	rpcURLArg := flag.String("rpcurl", "https://bartio.rpc.berachain.com", "Berachain RPC URL")
	flag.Parse()

	// The WinnieSwap adapter handles StickyVault contracts

	// wBERA-HONEY WinnieSwap StickyVault (0.05% fee tier)
	// winnieswap -address=0x46fbf6ff1fd62ec89af48c3bb0b63115052dab31
	// 			-price0=0x7507c1dc16935B82698e4C63f2746A2fCf994dF8:8.0
	// 			-price1=0x0E4aaF1351de4c0264C5c7056Ef3777b41BD8e03:1.0
	//			-rpcurl=https://bartio.rpc.berachain.com

	// Validate required arguments
	missingArgs := []string{}
	if *addressArg == "" {
		missingArgs = append(missingArgs, "address")
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
			Str("usage", "go run main.go -address <contract-address> -price0 <price0> -price1 <price1> -rpcurl <rpc-url>").
			Msg("Missing required arguments")
	}

	// Parse prices
	p0data := strings.Split(*price0Arg, ":")
	if len(p0data) != 2 {
		logger.Fatal().Msgf("Invalid price0, '%s'", *price0Arg)
	}
	price0, err := decimal.NewFromString(p0data[1])
	if err != nil {
		logger.Fatal().Err(err).Str("price0", *price0Arg).Msg("Invalid price0")
	}
	p1data := strings.Split(*price1Arg, ":")
	if len(p1data) != 2 {
		logger.Fatal().Msgf("Invalid price1, '%s'", *price1Arg)
	}
	price1, err := decimal.NewFromString(p1data[1])
	if err != nil {
		logger.Fatal().Err(err).Str("price1", *price1Arg).Msg("Invalid price1")
	}

	pmap := map[string]protocols.Price{
		strings.ToLower(p0data[0]): protocols.Price{Decimals: 18, Price: price0},
		strings.ToLower(p1data[0]): protocols.Price{Decimals: 18, Price: price1},
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	// get the LP price provider config
	cp := protocols.WinnieSwapLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *addressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for WinnieSwapLPPriceProvider")
	}

	configData := &protocols.WinnieSwapConfig{}
	err = json.Unmarshal(configBytes, configData)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to deserialize config")
	}
	logger.Info().
		Str("WinnieSwap Config Data", fmt.Sprintf("%+v", configData)).
		Msg("successfully created WinnieSwap Config")

	// Parse the smart contract address
	address := common.HexToAddress(*addressArg)
	// Create a new WinnieSwapLPPriceProvider
	provider := protocols.NewWinnieSwapLPPriceProvider(address, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize WinnieSwapLPPriceProvider")
	}

	// Fetch LP token price
	lpPrice, err := provider.LPTokenPrice(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch LP token price")
	} else {
		logger.Info().
			Str("LPTokenPrice (USD)", lpPrice).
			Msg("successfully fetched LP token price")
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

	// Test Offchain WinnieSwap GraphQL API for fetching pool APRs
	stakingTokens := []string{
		"0x46fbf6ff1fd62ec89af48c3bb0b63115052dab31", // wBERA-HONEY 0.05%
	}
	winnieswapAPRs, err := fetchers.FetchWinnieSwapAPRs(ctx, stakingTokens)
	if err != nil {
		logger.Warn().Err(err).Msg("failed to fetch APRs from WinnieSwap API (this is optional)")
	} else {
		logger.Info().
			Msgf("fetched WinnieSwap APRs from API %+v", winnieswapAPRs)
	}
}
