package main

import (
	"context"
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

	satSolvBTC := "0xFF21f46Bc9D78125705eEF6EfCA62f9420cfDB9b"
	solvBTC := "0x541FD749419CA806a8bc7da8ac23D346f2dF8B77"

	// Command-line arguments
	addressArg := flag.String("address", satSolvBTC, "Smart contract address (Default satSolvBTC.BERA)")
	price0Arg := flag.String("price0", fmt.Sprintf("%s:118805.23", solvBTC), "address:price of asset, colon delimited (SolvBTC)")
	rpcURLArg := flag.String("rpcurl", "https://rpc.berachain.com/", "Berachain Mainnet RPC URL")
	flag.Parse()

	// satSolvBTC.BERA is 1:1 with SolvBTC, so we proxy satSolvBTC.BERA price through SolvBTC price

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
	// SolvBTC is weird and does have 18 decimals although most BTC tokens have 8
	pmap := map[string]protocols.Price{
		strings.ToLower(p0data[0]): protocols.Price{Decimals: 18, Price: price0},
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	cp := protocols.SatLayerLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *addressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Str("address", *addressArg).Msg("Failed to get config")
	}

	// Parse the smart contract address
	address := common.HexToAddress(*addressArg)
	// Create a new NewSatLayerLPPriceProvider
	provider := protocols.NewSatLayerLPPriceProvider(address, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client, http.NewTestHttpClient())
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize NewSatLayerLPPriceProvider")
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

}
