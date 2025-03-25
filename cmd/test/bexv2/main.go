package main

import (
	"context"
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/infrared-dao/protocols"
	"github.com/shopspring/decimal"

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
	contractArg := flag.String("contract", "", "Balancer Vault contract address")
	lpTokenArg := flag.String("address", "", "LP Token address, ie. bex pool address")
	pricesArg := flag.String("prices", "", "address:price:decimals, for each token. comma delimited list")
	rpcURLArg := flag.String("rpcurl", "https://  berchain-rpc-url", "Mainnet Berachain RPC URL")
	flag.Parse()

	// BEXv2 has several different types of pools like weighted pools, stable pools, liq bootstrapping pools, managed pools, etc.

	// WBERA-HONEY weighted pool
	// bexv2 -contract=0x9C8a5c82e797e074Fe3f121B326b140CEC4bcb33 -address=0x3aD1699779eF2c5a4600e649484402DFBd3c503C
	//       -prices=0x6969696969696969696969696969696969696969:6.9:18,0xd137593CDB341CcC78426c54Fb98435C60Da193c:1.0:18
	//       -rpcurl=berchain-rpc-provider

	// WBERA-WBTC weighted pool
	// bexv2 -contract=0x9C8a5c82e797e074Fe3f121B326b140CEC4bcb33 -address=0x4A782a6bA2e47367A4b2A1551815c27dc15F4795
	//       -prices=0x6969696969696969696969696969696969696969:6.9:18,0xfa5bf670a92aff186e5176aa55690e0277010040:82172.7:8
	//       -rpcurl=berchain-rpc-provider

	// USDC-HONEY composable stable pool
	// bexv2 -contract=0x9C8a5c82e797e074Fe3f121B326b140CEC4bcb33 -address=0xf7f214a9543c1153ef5df2edcd839074615f248c
	//       -prices=0x015fd589f4f1a33ce4487e12714e1b15129c9329:1.044:6,0xd137593cdb341ccc78426c54fb98435c60da193c:1.0:18
	//       -rpcurl=berchain-rpc-provider

	// Validate required arguments
	missingArgs := []string{}
	if *contractArg == "" {
		missingArgs = append(missingArgs, "contract")
	}
	if *lpTokenArg == "" {
		missingArgs = append(missingArgs, "address")
	}
	if *pricesArg == "" {
		missingArgs = append(missingArgs, "prices")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -contract <vault-contract-address> -address <pool-address> -prices <token0:price0:decimals0,...> -rpcurl <rpc-url>").
			Msg("Missing required arguments")
	}

	// Parse prices
	pdata := strings.Split(*pricesArg, ",")
	if len(pdata) < 2 {
		logger.Fatal().Msgf("Invalid or not enough prices, '%s'", *pricesArg)
	}
	var pmap = make(map[string]protocols.Price)
	for _, data := range pdata {
		parts := strings.Split(data, ":")
		token := strings.ToLower(parts[0])
		price, err := decimal.NewFromString(parts[1])
		if err != nil {
			logger.Fatal().Err(err).Str("price", data).Msg("Invalid price")
		}
		decimals, err := strconv.Atoi(parts[2])
		if err != nil {
			logger.Fatal().Err(err).Str("decimals", data).Msg("Invalid decimals")
		}
		pmap[token] = protocols.Price{Decimals: uint(decimals), Price: price}
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to RPC client")
	}

	// get the LP price provider config
	cp := protocols.BexV2LPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *lpTokenArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for BexV2LPPriceProvider")
	}

	// Parse the smart contract addresses
	address := common.HexToAddress(*lpTokenArg)
	contract := common.HexToAddress(*contractArg)
	// Create a new BexV2LPPriceProvider
	provider := protocols.NewBexV2LPPriceProvider(contract, address, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize BexV2LPPriceProvider")
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
