package main

import (
	"context"
	"flag"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols"
	"github.com/infrared-dao/protocols/cmd/test/http"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func main() {
	// Create a zerolog logger
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// Command-line arguments
	addressArg := flag.String("address", "0x46fcd35431f5B371224ACC2e2E91732867B1A77e", "Smart contract address (Default primeLiquidBeraBTC)")
	price0Arg := flag.String("price0", "0x0555E30da8f98308EdB960aa94C0Db47230d2B9c:104413.52", "address:price of asset, colon delimited")
	rpcURLArg := flag.String("rpcurl", "https://rpc.berachain.com/", "Berachain Mainnet RPC URL")
	flag.Parse()

	// etherfi adapter can handle etherfi vaults

	// etherfi vault: primeLiquidBeraBTC
	// etherfi -address=0x46fcd35431f5B371224ACC2e2E91732867B1A77e
	// 		  -price0=0x0555E30da8f98308EdB960aa94C0Db47230d2B9c:104413.52
	//		  -rpcurl=berchain-rpc-provider

	// etherfi vault: primeLiquidBeraETH
	// etherfi -address=0xB83742330443f7413DBD2aBdfc046dB0474a944e
	// 		  -price0=0x2F6F07CDcf3588944Bf4C42aC74ff24bF56e7590:2614.14
	//		  -rpcurl=berchain-rpc-provider

	// Validate required arguments
	missingArgs := []string{}
	if *addressArg == "" {
		missingArgs = append(missingArgs, "address")
	}
	if *price0Arg == "" {
		missingArgs = append(missingArgs, "price0")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -address <contract-address> -price0 <price0> -rpcurl <rpc-url>").
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
	decimals := 18
	if strings.EqualFold(p0data[0], "0x0555E30da8f98308EdB960aa94C0Db47230d2B9c") {
		decimals = 8 // WBTC has 8 decimals
	}
	pmap := map[string]protocols.Price{
		strings.ToLower(p0data[0]): protocols.Price{Decimals: uint(decimals), Price: price0},
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	cp := protocols.EtherfiLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *addressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Str("address", *addressArg).Msg("Failed to get config")
	}

	// Parse the smart contract address
	address := common.HexToAddress(*addressArg)
	// Create a new NewetherfiLPPriceProvider
	provider := protocols.NewEtherfiLPPriceProvider(address, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client, http.NewTestHttpClient())
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize KodiakLPPriceProvider")
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
