package main

import (
	"context"
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/infrared-dao/protocols"
	"github.com/infrared-dao/protocols/fetchers"
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

	defaultPricesArg := "0x1ce0a25d13ce4d52071ae7e02cf1f6606f4c79d3:1.0:18,"
	defaultPricesArg += "0x549943e04f40284185054145c6e4e9568c1d3241:1.0:6,"
	defaultPricesArg += "0xfcbd14dc51f0a4d49d5e53c2e0950e0bc26d0dce:1.0:18"

	// Command-line arguments
	lpTokenArg := flag.String("address", "0xd10e65a5f8ca6f835f2b1832e37cf150fb955f23", "LP Token address, ie. bex pool address")
	pricesArg := flag.String("prices", defaultPricesArg, "address:price:decimals, for each token. comma delimited list")
	rpcURLArg := flag.String("rpcurl", "https://  berachain-rpc-url", "Mainnet Berachain RPC URL")
	flag.Parse()

	// NECT-USDC-HONEY composable stable pool
	// burrbear -address=0xd10e65a5f8ca6f835f2b1832e37cf150fb955f23
	//       -prices=0x1ce0a25d13ce4d52071ae7e02cf1f6606f4c79d3:1.0:18,
	// 					0x549943e04f40284185054145c6e4e9568c1d3241:1.0:6,
	// 					0xfcbd14dc51f0a4d49d5e53c2e0950e0bc26d0dce:1.0:18
	//       -rpcurl=berachain-rpc-provider

	// Validate required arguments
	missingArgs := []string{}
	if *lpTokenArg == "" {
		missingArgs = append(missingArgs, "address")
	}
	if *pricesArg == "" {
		missingArgs = append(missingArgs, "prices")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -address <pool-address> -prices <token0:price0:decimals0,...> -rpcurl <rpc-url>").
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
	cp := protocols.BurrBearLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *lpTokenArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for BurrBearLPPriceProvider")
	}

	// Parse the smart contract addresses
	address := common.HexToAddress(*lpTokenArg)
	// Create a new BurrBearLPPriceProvider
	provider := protocols.NewBurrBearLPPriceProvider(address, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize BurrBearLPPriceProvider")
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

	// Test Offchain Kodiak API for fetching average pool APR
	stakingTokens := []string{
		"0xcbef1b65399065c2de2c495971e90466ff38f2d0",
		"0xb34fca9739249578d93a22ce02bf7fad277bad35",
		"0xc4c039178a79e614609e592d7689ecb3107758b1",
		"0x567f32e86be3e3963cdbc1887b5043b701f113d9",
		"0xac8a437d1e6905e419431662aead90fb04ac3008",
		"0xd10e65a5f8ca6f835f2b1832e37cf150fb955f23",
		"0xe416c064946112c1626d6700d1081a750b1b1dd7",
		"0xd170e25f6bcb5ace2108628c647be47d59900ade",
		"0x7ce7cb1893cfbd680cbfb9dd2a9ae6a62bde66a8",
	}
	burrbearAPRs, err := fetchers.FetchBurrBearAPRs(ctx, stakingTokens)
	if err != nil {
		logger.Fatal().Err(err).Msg("bad response from burrbear API")
	}
	logger.Info().
		Msgf("fetched burrbear fee base APRs from API %+v", burrbearAPRs)
}
