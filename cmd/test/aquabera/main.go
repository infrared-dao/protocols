package main

import (
	"context"
	"flag"
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
	addressArg := flag.String("address", "0xf9845a03f7e6b06645a03a28b943c8a4b5fe7bcc", "AquaBera contract address")
	price0Arg := flag.String("price0", "0x6969696969696969696969696969696969696969:2.58", "address:price of token 0, colon delimited")
	price1Arg := flag.String("price1", "0x1f7210257fa157227d09449229a9266b0d581337:0.000274", "address:price of token 1, colon delimited")
	apiKeyArg := flag.String("apikey", "", "A valid aquabera api key for the x-api-key request headers")
	rpcURLArg := flag.String("rpcurl", "https://rpc.berachain.com", "Berachain RPC URL")
	flag.Parse()

	// WBERA-BERAMO kodiak LP aquabera pool "AB-KODIAK-WBERA-BERAMO-10000"
	// aquabera -address=0xf9845a03f7e6b06645a03a28b943c8a4b5fe7bcc
	//       -price0=0x6969696969696969696969696969696969696969:2.58
	//       -price1=0x1f7210257fa157227d09449229a9266b0d581337:0.000274
	//       -apikey=xxxxxxxxxxxxxxxxxxxxxxxxxxx
	//       -rpcurl=berachain-rpc-provider

	// Validate required arguments
	missingArgs := []string{}
	if *addressArg == "" {
		missingArgs = append(missingArgs, "address")
	}
	if *apiKeyArg == "" {
		missingArgs = append(missingArgs, "apikey")
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
			Str("usage", "go run main.go -address <contract-address> -price0 <address:price> -price1 <address:price> -apikey <key> -rpcurl <rpc-url>").
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

	// Get the LP price provider config
	ap := protocols.AquaBeraLPPriceProvider{}
	configBytes, err := ap.GetConfig(ctx, *addressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for AquaBeraLPPriceProvider")
	}

	// Parse the smart contract address
	address := common.HexToAddress(*addressArg)
	// Create a new AquaBeraLPPriceProvider
	provider := protocols.NewAquaBeraLPPriceProvider(address, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize AquaBeraLPPriceProvider")
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

	// Test Offchain Aquabera API for 7d base APR
	stakingTokens := []string{
		"0x04fd6a7b02e2e48caedad7135420604de5f834f8",
		"0xf9845a03f7e6b06645a03a28b943c8a4b5fe7bcc",
	}
	fetchFunction := fetchers.BuildAquaberaAPRsFetcher(*apiKeyArg)
	aquaberaAPRs, err := fetchFunction(ctx, stakingTokens)
	if err != nil {
		logger.Fatal().Err(err).Msg("bad response from aquabera API")
	}
	logger.Info().
		Msgf("fetched aquabera 7d base APRs from API %+v", aquaberaAPRs)

}
