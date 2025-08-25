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
	addressArg := flag.String("address", "0x98bdeede9a45c28d229285d9d6e9139e9f505391", "Smart contract address")
	price0Arg := flag.String("price0", "0x18878Df23e2a36f81e820e4b47b4A40576D3159C:26.32", "address:price of token 0")
	price1Arg := flag.String("price1", "0xFCBD14DC51f0A4d49d5E53C2E0950e0bC26d0Dce:1.0", "address:price of token 1")
	rpcURLArg := flag.String("rpcurl", "https://berchain-rpc-url", "Mainnet Berachain RPC URL")
	flag.Parse()

	// The kodiak adapter can handle both V3 Island and V2 Pool contracts in address

	// OHM-HONEY Kodiak V3 Island
	// kodiak -address=0x98bdeede9a45c28d229285d9d6e9139e9f505391
	// 			-price0=0x18878Df23e2a36f81e820e4b47b4A40576D3159C:26.32
	// 			-price1=0xFCBD14DC51f0A4d49d5E53C2E0950e0bC26d0Dce:1.0
	//			-rpcurl=berchain-rpc-provider

	// HOLD-WBERA Kodiak V2 Pool
	// kodiak -address=0xdca120bd3a13250b67f6faa5c29c1f38ec6ebece
	// 			-price0=0xff0a636dfc44bb0129b631cdd38d21b613290c98:1.446
	// 			-price1=0x6969696969696969696969696969696969696969:8.66
	//			-rpcurl=berchain-rpc-provider

	// CHRM WBERA-USDC Kodiak Charm Pool (concentrated liquidity)
	// kodiak -address=0x38920562047280f2f95b7aba7a9eaa8d0ae04a5c
	// 			-price0=0x549943e04f40284185054145c6E4e9568C1D3241:1.002  // NOTE: Decimals = 6
	// 			-price1=0x6969696969696969696969696969696969696969:1.93
	//			-rpcurl=berchain-rpc-provider

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
	cp := protocols.KodiakLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *addressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for KodiakLPPriceProvider")
	}

	configData := &protocols.KodiakConfig{}
	err = json.Unmarshal(configBytes, configData)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to deserialize config")
	}
	logger.Info().
		Str("Kodiak Config Data", fmt.Sprintf("%+v", configData)).
		Msg("successfully created Kodiak Config")

	// Parse the smart contract address
	address := common.HexToAddress(*addressArg)
	// Create a new KodiakLPPriceProvider
	provider := protocols.NewKodiakLPPriceProvider(address, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize KodiakLPPriceProvider")
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

	// Test Offchain Kodiak API for fetching average pool APR
	stakingTokens := []string{
		"0x98bdeede9a45c28d229285d9d6e9139e9f505391",
		"0x8b161685135e9fbc5475169e1addc0f2c4b7c343",
		"0xec8ba456b4e009408d0776cde8b91f8717d13fa1",
		"0xdca120bd3a13250b67f6faa5c29c1f38ec6ebece",
		"0xc60ea9801b3f1b01da373c05381e6ce8ff94d76f",
		"0x38920562047280f2f95b7aba7a9eaa8d0ae04a5c",
	}
	kodiakAPRs, err := fetchers.FetchKodiakAPRs(ctx, stakingTokens)
	if err != nil {
		logger.Fatal().Err(err).Msg("bad response from kodiak API")
	}
	logger.Info().
		Msgf("fetched kodiak average APRs from API %+v", kodiakAPRs)
}
