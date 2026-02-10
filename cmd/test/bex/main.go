package main

import (
	"context"
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols"
	"github.com/infrared-dao/protocols/cmd/test/constant"
	"github.com/infrared-dao/protocols/cmd/test/http"
	"github.com/infrared-dao/protocols/fetchers"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func main() {
	// Create a zerolog logger
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	examplePricesArg := "0x6969696969696969696969696969696969696969:6.9:18,0xfcbd14dc51f0a4d49d5e53c2e0950e0bc26d0dce:1.0:18"

	// Command-line arguments
	contractArg := flag.String("contract", "0x4Be03f781C497A489E3cB0287833452cA9B9E80B", "Balancer Vault contract address")
	lpTokenArg := flag.String("address", "0x2c4a603a2aa5596287a06886862dc29d56dbc354", "LP Token address, ie. bex pool address")
	pricesArg := flag.String("prices", examplePricesArg, "address:price:decimals, for each token. comma delimited list")
	rpcURLArg := flag.String("rpcurl", constant.DefaultBerachainRPCURL, "Mainnet Berachain RPC URL")
	flag.Parse()

	// BEX has several different types of pools like weighted pools, stable pools, liq bootstrapping pools, managed pools, etc.

	// WBERA-HONEY weighted pool
	// bex -contract=0x4Be03f781C497A489E3cB0287833452cA9B9E80B -address=0x2c4a603a2aa5596287a06886862dc29d56dbc354
	//       -prices=0x6969696969696969696969696969696969696969:6.9:18,0xfcbd14dc51f0a4d49d5e53c2e0950e0bc26d0dce:1.0:18

	// WBERA-WBTC weighted pool
	// bex -contract=0x4Be03f781C497A489E3cB0287833452cA9B9E80B -address=0x38fdd999fe8783037db1bbfe465759e312f2d809
	//       -prices=0x6969696969696969696969696969696969696969:6.9:18,0x0555e30da8f98308edb960aa94c0db47230d2b9c:82172.7:8

	// USDC.e-HONEY composable stable pool
	// bex -contract=0x4Be03f781C497A489E3cB0287833452cA9B9E80B -address=0xf961a8f6d8c69e7321e78d254ecafbcc3a637621
	//       -prices=0x549943e04f40284185054145c6e4e9568c1d3241:1.044:6,0xfcbd14dc51f0a4d49d5e53c2e0950e0bc26d0dce:1.0:18

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
	cp := protocols.BexLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *lpTokenArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to get config for BexLPPriceProvider")
	}

	// Parse the smart contract addresses
	address := common.HexToAddress(*lpTokenArg)
	contract := common.HexToAddress(*contractArg)
	// Create a new BexLPPriceProvider
	provider := protocols.NewBexLPPriceProvider(contract, address, nil, pmap, logger, configBytes)

	// Initialize the provider
	err = provider.Initialize(ctx, client, http.NewTestHttpClient())
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

	// Test Offchain BEX API for fetching current pool APR
	stakingTokens := []string{
		"0x1207c619086a52edef4a4b7af881b5ddd367a919",
		"0xdd70a5ef7d8cfe5c5134b5f9874b09fb5ce812b4",
		"0x2461e93d5963c2bb69de499676763e67a63c7ba5",
		"0x62c030b29a6fef1b32677499e4a1f1852a8808c0",
	}
	bexAPRs, err := fetchers.FetchBexAPRs(ctx, http.NewTestHttpClient(), stakingTokens)
	if err != nil {
		logger.Fatal().Err(err).Msg("bad response from bex API")
	}
	logger.Info().
		Msgf("fetched bex APRs from API %+v", bexAPRs)
}
