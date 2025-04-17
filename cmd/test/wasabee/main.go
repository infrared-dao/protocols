package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
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
	addressArg := flag.String("address", "0xEC06041013b3a97c58b9ab61eAE9079Bc594EdA3", "Smart contract address of Wasabee pool")
	rpcURLArg := flag.String("rpcurl", DefaultRpcURL, "RPC URL for the blockchain")
	price0Arg := flag.String("price0", "0x2F6F07CDcf3588944Bf4C42aC74ff24bF56e7590:1474.66", "Price of token0 in USD")
	price1Arg := flag.String("price1", "0x6969696969696969696969696969696969696969:3.43", "Price of token1 in USD")
	flag.Parse()

	// Validate required arguments
	if *addressArg == "" {
		logger.Fatal().Msg("Missing required argument: -address <contract-address>")
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
		strings.ToLower(p0data[0]): protocols.Price{Decimals: 18, Price: price0},
		strings.ToLower(p1data[0]): protocols.Price{Decimals: 18, Price: price1},
	}

	// Connect to Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to RPC node")
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Pool address
	poolAddress := common.HexToAddress(*addressArg)

	// get the LP price provider config
	cp := protocols.WasabeeLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *addressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for WasabeeLPPriceProvider")
	}

	configData := &protocols.WasabeeConfig{}
	err = json.Unmarshal(configBytes, configData)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to deserialize config")
	}
	logger.Info().
		Str("Wasabee Config Data", fmt.Sprintf("%+v", configData)).
		Msg("successfully created Wasabee Config")

	// Initialize the provider
	provider := protocols.NewWasabeeLPPriceProvider(poolAddress, nil, pmap, logger, configBytes)
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize Wasabee adapter")
	}

	// Get LP token price
	lpPrice, err := provider.LPTokenPrice(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get LP token price")
	}
	fmt.Printf("LP Token Price: $%s\n", lpPrice)

	// Get TVL
	tvl, err := provider.TVL(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get TVL")
	}
	fmt.Printf("TVL: $%s\n", tvl)
}
