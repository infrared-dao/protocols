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
	lpTokenArg := flag.String("address", "0xB318Cd79dC0743De041A26D3F0d467d49955E5bC", "LP Token address, ie. beraborrow infrared wrapper token address")
	rpcURLArg := flag.String("rpcurl", "https://  mainnet-rpc-url", "Berachain RPC URL")
	flag.Parse()

	// BeraBorrow CDP -- specifically a Compounding Infrared Collateralized Vault

	// 0xB318Cd79dC0743De041A26D3F0d467d49955E5bC is the address of the InfraredWrapper token used as LP token
	// 0xCE1e426E35eBc9f512944F59527304E3B771EA12 would be the address of the Compounding Infrared Collateral Vault
	// do not use the CICV as the LP token, only the IW type token should be used as an LP token for this adapter

	// beraborrow pumpBTC
	// beraborrow -address=0xB318Cd79dC0743De041A26D3F0d467d49955E5bC
	//       -rpcurl=berachain-rpc-provider

	// Validate required arguments
	missingArgs := []string{}
	if *lpTokenArg == "" {
		missingArgs = append(missingArgs, "address")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -address <cdp-address> -rpcurl <rpc-url>").
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
	cp := protocols.BeraBorrowLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *lpTokenArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for BeraBorrowLPPriceProvider")
	}

	// Parse the smart contract addresses
	address := common.HexToAddress(*lpTokenArg)
	// Create a new BeraBorrowLPPriceProvider
	provider := protocols.NewBeraBorrowLPPriceProvider(address, blockNumber, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize BeraBorrowLPPriceProvider")
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
