package main

import (
	"context"
	"flag"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols"
	"github.com/infrared-dao/protocols/cmd/test/constant"
	"github.com/infrared-dao/protocols/cmd/test/http"
	"github.com/infrared-dao/protocols/fetchers"
	"github.com/rs/zerolog"
)

func main() {
	// Create a zerolog logger
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// Command-line arguments
	lpTokenArg := flag.String("address", "0xc2605ed80880bd6b1523d52aef8d624ed468a935", "Wrapped LP Token address, ie. pendle pool address")
	rpcURLArg := flag.String("rpcurl", constant.DefaultBerachainRPCURL, "Mainnet Berachain RPC URL")
	flag.Parse()

	// Pendle iBGT pool
	// pendle -address=0xc2605ed80880bd6b1523d52aef8d624ed468a935
	//        -rpcurl=berachain-rpc-provider

	// Validate required arguments
	missingArgs := []string{}
	if *lpTokenArg == "" {
		missingArgs = append(missingArgs, "address")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -address <pool-address> -rpcurl <rpc-url>").
			Msg("Missing required arguments")
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to RPC client")
	}

	// get the LP price provider config
	cp := protocols.PendleLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *lpTokenArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for PendleLPPriceProvider")
	}

	// Parse the smart contract addresses
	address := common.HexToAddress(*lpTokenArg)
	// Create a new PendleLPPriceProvider
	provider := protocols.NewPendleLPPriceProvider(address, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client, http.NewTestHttpClient())
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize PendleLPPriceProvider")
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

	// Test fetcher using Pendle API for base APRs
	pools := []string{
		"0x02293d88656860a840ed457d38f51e3e6bee8461",
		"0x5200c9900436f649b4659dfc79213837dcccaab1",
	}
	pendleAPRs, err := fetchers.FetchPendleAPRs(ctx, http.NewTestHttpClient(), pools)
	if err != nil {
		logger.Fatal().Err(err).Msg("bad response from pendle API")
	}
	logger.Info().
		Msgf("fetched pendle base APRs from API %+v", pendleAPRs)
}
