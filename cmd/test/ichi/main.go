package main

import (
	"context"
	"flag"
	"os"

	"github.com/infrared-dao/protocols"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
)

func main() {
	// Create a zerolog logger
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// Command-line arguments
	ichiVaultArg := flag.String("address", "", "ICHI Vault address")
	rpcURLArg := flag.String("rpcurl", "https://mainnet-rpc-url", "Berachain RPC URL")
	flag.Parse()

	// Validate required arguments
	missingArgs := []string{}
	if *ichiVaultArg == "" {
		missingArgs = append(missingArgs, "address")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -address <ichi-vault-address> -rpcurl <rpc-url>").
			Msg("Missing required arguments")
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to RPC client")
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get current block header from Eth client")
	}
	blockNumber := header.Number

	// get the LP price provider config
	cp := protocols.ICHIPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *ichiVaultArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for ICHIPriceProvider")
	}

	// Parse the smart contract addresses
	address := common.HexToAddress(*ichiVaultArg)
	// Create a new ICHIPriceProvider
	provider := protocols.NewICHIPriceProvider(address, blockNumber, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize ICHIPriceProvider")
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
