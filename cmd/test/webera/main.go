package main

import (
	"context"
	"flag"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols"
	"github.com/infrared-dao/protocols/cmd/test/http"
	"github.com/infrared-dao/protocols/fetchers"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func main() {
	// Create a zerolog logger
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// Command-line arguments
	addressArg := flag.String("address", "0x55a050f76541c2554e9dfa3a0b4e665914bf92ea", "Smart contract address (Default wBera Vault)")
	price0Arg := flag.String("price0", "0x6969696969696969696969696969696969696969:3.687", "address:price of token 0, colon delimited (Bera price)")
	rpcURLArg := flag.String("rpcurl", "https://rpc.berachain.com/", "Berachain Mainnet RPC URL")
	flag.Parse()

	// WeBera adapter can handle WeBera vaults

	// weBERA vault
	// webera -address=0x55a050f76541c2554e9dfa3a0b4e665914bf92ea
	// 		  -price0=0x6969696969696969696969696969696969696969:3.687
	//		  -rpcurl=berchain-rpc-provider

	// weiBERA vault
	// webera -address=0x396a3d0b799b1a0b1eaa17e75b4dea412400860b
	// 		  -price0=0x9b6761bf2397Bb5a6624a856cC84A3A14Dcd3fe5:3.724
	//		  -rpcurl=berchain-rpc-provider

	// Validate required arguments
	missingArgs := []string{}
	if *addressArg == "" {
		missingArgs = append(missingArgs, "address")
	}
	if *price0Arg == "" {
		missingArgs = append(missingArgs, "price0")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -address <contract-address> -price0 <price0> -rpcurl <rpc-url>").
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
	pmap := map[string]protocols.Price{
		strings.ToLower(p0data[0]): protocols.Price{Decimals: 18, Price: price0},
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	cp := protocols.WeberaLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *addressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Str("address", *addressArg).Msg("Failed to get config")
	}

	// Parse the smart contract address
	address := common.HexToAddress(*addressArg)
	// Create a new NewWeberaLPPriceProvider
	provider := protocols.NewWeberaLPPriceProvider(address, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client, http.NewTestHttpClient())
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

	// Test Offchain Webera API for fetching vault APR
	stakingTokens := []string{
		"0x396A3D0B799B1a0B1EaA17e75B4DEa412400860b",
		"0x55a050f76541C2554e9dfA3A0b4e665914bF92EA",
	}
	weberaAPRs, err := fetchers.FetchWeberaAPRs(ctx, http.NewTestHttpClient(), stakingTokens)
	if err != nil {
		logger.Fatal().Err(err).Msg("bad response from webera API")
	}
	logger.Info().
		Msgf("fetched webera vault APRs from API %+v", weberaAPRs)

}
