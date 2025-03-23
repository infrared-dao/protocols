package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

const (
	// Use appropriate RPC URL for your network
	DefaultRpcURL = "https://rpc.berachain.com" 
)

func main() {
	// Setup logger
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Parse command-line arguments
	addressArg := flag.String("address", "", "Smart contract address of Bulla pool") // No default, must be provided
	rpcURLArg := flag.String("rpcurl", DefaultRpcURL, "RPC URL for the blockchain")
	price0Arg := flag.String("price0", "1.0", "Price of token0 in USD")
	price1Arg := flag.String("price1", "1.0", "Price of token1 in USD")
	flag.Parse()

	// Validate required arguments
	if *addressArg == "" {
		logger.Fatal().Msg("Missing required argument: -address <contract-address>")
	}

	// Parse price arguments
	price0, err := decimal.NewFromString(*price0Arg)
	if err != nil {
		logger.Fatal().Err(err).Str("price0", *price0Arg).Msg("Invalid price0 value")
	}
	
	price1, err := decimal.NewFromString(*price1Arg)
	if err != nil {
		logger.Fatal().Err(err).Str("price1", *price1Arg).Msg("Invalid price1 value")
	}

	// Connect to Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to Ethereum client")
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Pool address
	poolAddress := common.HexToAddress(*addressArg)

	// First call GetConfig to get the actual token addresses
	bullaProvider := protocols.NewBullaLPPriceProvider(poolAddress, nil, logger, nil)
	config, err := bullaProvider.GetConfig(ctx, *addressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config, please check your contract address")
		return
	}
	
	fmt.Printf("Config: %s\n", string(config))
	
	// Parse the returned config to extract token addresses
	var bullaConfig protocols.BullaConfig
	if err := json.Unmarshal(config, &bullaConfig); err != nil {
		logger.Fatal().Err(err).Msg("Failed to parse config")
		return
	}
	
	// Setup token prices based on the command-line arguments
	prices := map[string]protocols.Price{
		bullaConfig.Token0: {
			TokenName: "TOKEN0",
			Decimals:  18,
			Price:     price0,
		},
		bullaConfig.Token1: {
			TokenName: "TOKEN1",
			Decimals:  18,
			Price:     price1,
		},
	}

	logger.Info().
		Str("Token0", bullaConfig.Token0).
		Str("Price0", price0.String()).
		Str("Token1", bullaConfig.Token1).
		Str("Price1", price1.String()).
		Msg("Using token prices")

	// Create a new provider with the config and prices
	bullaProvider = protocols.NewBullaLPPriceProvider(poolAddress, prices, logger, config)

	// Initialize the provider
	err = bullaProvider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize Bulla adapter")
	}

	// Get LP token price
	lpPrice, err := bullaProvider.LPTokenPrice(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get LP token price")
	}
	fmt.Printf("LP Token Price: $%s\n", lpPrice)

	// Get TVL
	tvl, err := bullaProvider.TVL(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get TVL")
	}
	fmt.Printf("TVL: $%s\n", tvl)
} 