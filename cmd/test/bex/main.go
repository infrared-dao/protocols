package main

import (
	"context"
	"flag"
	"os"
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
	contractArg := flag.String("contract", "", "Query Smart contract address")
	lpTokenArg := flag.String("address", "", "LP Token address")
	price0Arg := flag.String("price0", "", "address:price of token 0, colon delimited")
	price1Arg := flag.String("price1", "", "address:price of token 1, colon delimited")
	rpcURLArg := flag.String("rpcurl", "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID", "Ethereum RPC URL")
	flag.Parse()

	// Example call for HONEY-WBERA pool
	// bex -contract=0x8685CE9Db06D40CBa73e3d09e6868FE476B5dC89 -address=0xd28d852cbcc68dcec922f6d5c7a8185dbaa104b7
	//     -price0=0x0e4aaf1351de4c0264c5c7056ef3777b41bd8e03:1.0 -price1=0x7507c1dc16935b82698e4c63f2746a2fcf994df8:20.88
	//     -rpcurl=...

	// Validate required arguments
	missingArgs := []string{}
	if *contractArg == "" {
		missingArgs = append(missingArgs, "contract")
	}
	if *lpTokenArg == "" {
		missingArgs = append(missingArgs, "address")
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
			Str("usage", "go run main.go -address <contract-address> -abipath <path-to-abi> -price0 <price0> -price1 <price1> -rpcurl <rpc-url>").
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
		strings.ToLower(p0data[0]): protocols.Price{Decimals: 18, Price: price0},
		strings.ToLower(p1data[0]): protocols.Price{Decimals: 18, Price: price1},
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	// get the LP price provider config
	cp := protocols.BexLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *lpTokenArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for BexLPPriceProvider")
	}

	// Parse the smart contract addresses
	address := common.HexToAddress(*lpTokenArg)
	contract := common.HexToAddress(*contractArg)
	// Create a new KodiakLPPriceProvider
	provider := protocols.NewBexLPPriceProvider(contract, address, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize BexLPPriceProvider")
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

	// Fetch TVL Breakdown
	breakdown, err := provider.TVLBreakdown(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch TVL breakdown")
	} else {
		logger.Info().Msg("successfully fetched TVL breakdown")
		for tokenAddr, tokenTVL := range breakdown {
			logger.Info().
				Str("token", tokenAddr).
				Str("symbol", tokenTVL.TokenSymbol).
				Str("amount", tokenTVL.Amount.String()).
				Str("usdValue", tokenTVL.USDValue.String()).
				Str("ratio", tokenTVL.Ratio.String()).
				Msg("token breakdown")
		}
	}
}
