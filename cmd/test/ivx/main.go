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
	"github.com/rs/zerolog"
)

func main() {
	// Create a zerolog logger
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// Command-line arguments
	lpTokenArg := flag.String("address", "0x3b8B155E3C44f07f6EAd507570f4047C8B450A7F", "LP Token address")
	lpMonitorArg := flag.String("lpmonitor", "0xf30EC2B4363c7957dab4b83B3211a278e280802D", "LP Monitor address")
	rpcURLArg := flag.String("rpcurl", constant.DefaultBerachainRPCURL, "Berachain Mainnet RPC URL")
	flag.Parse()

	// Validate required arguments
	if *lpTokenArg == "" {
		logger.Fatal().
			Str("usage", "go run main.go -address <lp-token-address> -rpcurl <rpc-url>").
			Msg("Missing required arguments")
	}

	if *lpMonitorArg == "" {
		logger.Fatal().
			Str("usage", "go run main.go -lpmonitor <lp-monitor-address> -rpcurl <rpc-url>").
			Msg("Missing required arguments")
	}

	ctx := context.Background()

	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	// Parse the smart contract address
	lpTokenAddress := common.HexToAddress(*lpTokenArg)
	lpMonitorAddress := common.HexToAddress(*lpMonitorArg)

	// get the LP price provider config
	cp := protocols.IVXLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *lpTokenArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for IVXLPPriceProvider")
	}

	// Create a new IVXLPPriceProvider
	provider := protocols.NewIVXLPPriceProvider(lpMonitorAddress, lpTokenAddress, nil, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client, http.NewTestHttpClient())
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize IVXLPPriceProvider")
	}

	// Fetch LP token price
	lpPrice, err := provider.LPTokenPrice(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch LP token price")
	} else {
		logger.Info().
			Str("LPTokenPrice (USD)", lpPrice).
			Msg("Successfully fetched LP token price")
	}

	// Fetch TVL
	tvl, err := provider.TVL(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch TVL")
	} else {
		logger.Info().
			Str("TVL (USD)", tvl).
			Msg("Successfully fetched TVL")
	}
}
