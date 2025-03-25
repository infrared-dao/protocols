package main

import (
	"context"
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/infrared-dao/protocols"

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
	vaultAddressArg := flag.String("contract", "", "Balancer vault contract address")
	lpTokenArg := flag.String("address", "", "LP Token address")
	pricesArg := flag.String("prices", "", "address:price:decimals, for each token. comma delimited list")
	rpcURLArg := flag.String("rpcurl", "https://rpc.berachain.com/", "Ethereum RPC URL")
	flag.Parse()

	// Validate required arguments
	missingArgs := []string{}
	if *vaultAddressArg == "" {
		missingArgs = append(missingArgs, "vault")
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
			Str("usage", "go run main.go -contract <vault-address> -address <pool-address> -prices <token0:price0:decimals0,...> -rpcurl <rpc-url>").
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
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	// get the LP price provider config
	cp := protocols.BurrBearLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *lpTokenArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for BexV2LPPriceProvider")
	}

	// Create and initialize the LP price provider
	provider := protocols.NewBurrBearLPPriceProvider(common.HexToAddress(*vaultAddressArg), common.HexToAddress(*lpTokenArg), pmap, logger, configBytes)
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize price provider")
	}

	// Get LP token price
	price, err := provider.LPTokenPrice(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get LP token price")
	}
	logger.Info().Str("price", price).Msg("LP token price (USD cents)")

	// Get TVL
	tvl, err := provider.TVL(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get TVL")
	}
	logger.Info().Str("tvl", tvl).Msg("Total Value Locked (USD cents)")
}
