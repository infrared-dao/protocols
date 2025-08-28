package main

import (
	"context"
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
	addressArg := flag.String("address", "0x66611Ba2aa5deB46e6138aD21a202b40ecE5b6AB", "LrBGT token contract address")
	addressArg2 := flag.String("address2", "0x381E9DC031D2667BB61d4Ab61e64D520Bbd7bFfC", "LrBGT Manager contract address")
	addressArg3 := flag.String("address3", "0xfc3Da0822fE3127334B95Ea3661060B957379c82", "LrBGT Manager Helper contract address")
	price0Arg := flag.String("price0", "0xac03CABA51e17c86c921E1f6CBFBdC91F8BB2E6b:2.699", "address:price of token 0, colon delimited")
	price1Arg := flag.String("price1", "0x6969696969696969696969696969696969696969:2.549", "address:price of token 1, colon delimited")
	price2Arg := flag.String("price2", "0xf3530788DEB3d21E8fA2c3CBBF93317FB38a0D3C:0.10", "address:price of token 2, colon delimited")
	rpcURLArg := flag.String("rpcurl", "https://rpc.berachain.com", "Berachain RPC URL")
	flag.Parse()

	// Validate required arguments
	missingArgs := []string{}
	if *addressArg == "" {
		missingArgs = append(missingArgs, "address")
	}
	if *addressArg2 == "" {
		missingArgs = append(missingArgs, "address2")
	}
	if *addressArg3 == "" {
		missingArgs = append(missingArgs, "address3")
	}
	if *price0Arg == "" {
		missingArgs = append(missingArgs, "price0")
	}
	if *price1Arg == "" {
		missingArgs = append(missingArgs, "price1")
	}
	if *price2Arg == "" {
		missingArgs = append(missingArgs, "price2")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -address <contract-address> -address2 <contract-address> -address3 <contract-address> -price0 <address:price> -price1 <address:price> -price2 <address:price> -rpcurl <rpc-url>").
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

	p2data := strings.Split(*price2Arg, ":")
	if len(p2data) != 2 {
		logger.Fatal().Msgf("Invalid price2, '%s'", *price2Arg)
	}
	price2, err := decimal.NewFromString(p2data[1])
	if err != nil {
		logger.Fatal().Err(err).Str("price2", *price1Arg).Msg("Invalid price2")
	}

	pmap := map[string]protocols.Price{
		strings.ToLower(p0data[0]): protocols.Price{Decimals: 18, Price: price0},
		strings.ToLower(p1data[0]): protocols.Price{Decimals: 18, Price: price1},
		strings.ToLower(p2data[0]): protocols.Price{Decimals: 18, Price: price2},
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	// Get the LP price provider config
	ap := protocols.LrBGTLPPriceProvider{}
	configBytes, err := ap.GetConfig(ctx, *addressArg2, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for LrBGTLPPriceProvider")
	}

	// Parse the smart contract address
	address := common.HexToAddress(*addressArg)
	// Parse the smart contract address
	address2 := common.HexToAddress(*addressArg2)
	// Parse the smart contract address
	address3 := common.HexToAddress(*addressArg3)

	// Create a new LrBGTLPPriceProvider
	provider := protocols.NewLrBGTLPPriceProvider(address, address2, address3, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize LrBGTLPPriceProvider")
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
