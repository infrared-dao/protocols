package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols"
	"github.com/infrared-dao/protocols/cmd/test/constant"
	"github.com/infrared-dao/protocols/cmd/test/http"
	"github.com/rs/zerolog"
)

func main() {
	// From Kodiak example: create a zerolog logger
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	// test price conversion
	err := testPriceConversion()
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	logger.Info().Msg("price conversion successful")
	addressArg := flag.String("address", "0x26bbc26415c6316890565f5f73017f85ee70b60c", "Smart contract address")
	rpcURLArg := flag.String("rpcurl", constant.DefaultBerachainRPCURL, "Mainnet Berachain RPC URL")
	flag.Parse()

	// Validate required arguments
	missingArgs := []string{}
	if *addressArg == "" {
		missingArgs = append(missingArgs, "address")
	}
	if len(missingArgs) > 0 {
		logger.Fatal().
			Strs("missingArgs", missingArgs).
			Str("usage", "go run main.go -address <contract-address> -rpcurl <rpc-url>").
			Msg("Missing required arguments")
	}

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}

	cp := protocols.D8xLPPriceProvider{}
	configBytes, err := cp.GetConfig(ctx, *addressArg, client)
	if err != nil {
		logger.Fatal().Err(err).Str("address", *addressArg).Msg("Failed to get config")
	}

	// Parse the smart contract address
	address := common.HexToAddress(*addressArg)
	// Create a new DolomiteLPPriceProvider
	provider := protocols.NewD8xLPPriceProvider(address, nil, logger, configBytes)

	err = provider.Initialize(ctx, client, http.NewTestHttpClient())
	if err != nil {
		logger.Fatal().Err(err).Msg("Initialize failed")
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

func testPriceConversion() error {
	fromto := []string{
		"18446744073709551616", "1.0000",
		"18446744073709551", "0.0010",
		"-57950446507558556402", "-3.1415",
	}
	for k := 0; k < len(fromto); k += 2 {
		value, ok := new(big.Int).SetString(fromto[k], 10)
		if !ok {
			return fmt.Errorf("invalid number")
		}
		f, err := protocols.ABDKToDecimal(value)
		if err != nil {
			return err
		}
		fStr := f.StringFixed(4)
		if fStr != fromto[k+1] {
			return fmt.Errorf("%s != %s", fStr, fromto[k+1])
		}
	}
	return nil
}
