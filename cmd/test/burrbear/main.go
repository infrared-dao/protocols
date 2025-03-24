package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
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
	vaultAddressArg := flag.String("vault", "", "Balancer vault contract address")
	poolIdArg := flag.String("poolid", "", "Balancer pool ID")
	price0Arg := flag.String("price0", "", "address:price of token 0, colon delimited")
	price1Arg := flag.String("price1", "", "address:price of token 1, colon delimited")
	price2Arg := flag.String("price2", "", "address:price of token 2, colon delimited (optional)")
	rpcURLArg := flag.String("rpcurl", "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID", "Ethereum RPC URL")
	flag.Parse()

	// Validate required arguments
	missingArgs := []string{}
	if *vaultAddressArg == "" {
		missingArgs = append(missingArgs, "vault")
	}
	if *poolIdArg == "" {
		missingArgs = append(missingArgs, "poolid")
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
			Str("usage", "go run main.go -vault <vault-address> -poolid <pool-id> -price0 <price0> -price1 <price1> [-price2 <price2>] -rpcurl <rpc-url>").
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
		strings.ToLower(p0data[0]): {Decimals: 18, Price: price0},
		strings.ToLower(p1data[0]): {Decimals: 18, Price: price1},
	}

	// Add optional third token price if provided
	if *price2Arg != "" {
		p2data := strings.Split(*price2Arg, ":")
		if len(p2data) != 2 {
			logger.Fatal().Msgf("Invalid price2, '%s'", *price2Arg)
		}
		price2, err := decimal.NewFromString(p2data[1])
		if err != nil {
			logger.Fatal().Err(err).Str("price2", *price2Arg).Msg("Invalid price2")
		}
		pmap[strings.ToLower(p2data[0])] = protocols.Price{Decimals: 18, Price: price2}
	}

	// Create BurrBear config
	config := protocols.BurrBearConfig{
		PoolId:      *poolIdArg,
		LPTDecimals: 18,
	}
	configBytes, err := json.Marshal(config)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to marshal config")
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	// Create and initialize the LP price provider
	provider := protocols.NewBurrBearLPPriceProvider(common.HexToAddress(*vaultAddressArg), pmap, logger, configBytes)
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
	// tvl, err := provider.TVL(ctx)
	// if err != nil {
	// 	logger.Fatal().Err(err).Msg("Failed to get TVL")
	// }
	// logger.Info().Str("tvl", tvl).Msg("Total Value Locked (USD cents)")
}
