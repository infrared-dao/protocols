package main

import (
	"context"
	"flag"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func main() {
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	addressArg := flag.String("address", "", "TermMax vault contract address")
	price0Arg := flag.String("price0", "", "address:price of underlying asset (HONEY), colon delimited")
	rpcURLArg := flag.String("rpcurl", "https://rpc.berachain.com/", "Berachain Mainnet RPC URL")
	flag.Parse()

	// Example usage:
	// go run main.go -address=<TERMMAX_VAULT_ADDRESS> \
	//                -price0=0x0E4aaF1351de4c0264C5c7056Ef3777b41BD8e03:1.00 \
	//                -rpcurl=https://rpc.berachain.com/

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
			Str("usage", "go run main.go -address <vault-address> -price0 <honey-address:price> -rpcurl <rpc-url>").
			Msg("Missing required arguments")
	}

	p0data := strings.Split(*price0Arg, ":")
	if len(p0data) != 2 {
		logger.Fatal().Msgf("Invalid price0, '%s'", *price0Arg)
	}
	price0, err := decimal.NewFromString(p0data[1])
	if err != nil {
		logger.Fatal().Err(err).Str("price0", *price0Arg).Msg("Invalid price0")
	}
	pmap := map[string]protocols.Price{
		strings.ToLower(p0data[0]): {Decimals: 18, Price: price0},
	}

	ctx := context.Background()

	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Berachain client")
	}

	cp := protocols.TermMaxVaultPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *addressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Str("address", *addressArg).Msg("Failed to get config")
	}

	address := common.HexToAddress(*addressArg)
	provider := protocols.NewTermMaxVaultPriceProvider(address, nil, pmap, logger, configBytes)

	err = provider.Initialize(ctx, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to initialize TermMaxVaultPriceProvider")
	}

	lpPrice, err := provider.SharePrice(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch LP token price")
	} else {
		logger.Info().
			Str("SharePrice (USD)", lpPrice).
			Msg("successfully fetched LP token price")
	}

	tvl, err := provider.TVL(ctx)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to fetch TVL")
	} else {
		logger.Info().
			Str("TVL (USD)", tvl).
			Msg("successfully fetched TVL")
	}
}