package main

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/infrared-dao/protocols"
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
	// From Kodiak example: command-line arguments
	logger.Info().Msg("price conversion successful")
	poolIdArg := flag.String("poolId", "1", "poolId")
	oracleAddr := flag.String("oracle address", "0xA8655EF2354d679E2553C10b2d59a61C4345aF51", "address of pool token collateral USD-price")
	d8xMngrAddr := flag.String("mngr address", "0xb6329c7168b255Eca8e5c627b0CCe7A5289C8b7F", "address of d8x manager")
	rpcURLArg := flag.String("rpcurl", "https://rpc.berachain.com", "Mainnet Berachain RPC URL")
	flag.Parse()

	// get config
	tmp := protocols.D8xLPPriceProvider{}
	id, err := strconv.Atoi(*poolIdArg)
	if err != nil {
		logger.Fatal().Err(err).Str("poolIdArg", *poolIdArg).Msg("invalid pool id")
	}
	configBytes, err := tmp.GetConfig(uint8(id), *oracleAddr)
	if err != nil {
		logger.Fatal().Err(err).Msg("GetConfig failed")
	}
	mngr := common.HexToAddress(*d8xMngrAddr)
	provider := protocols.NewD8xLPPriceProvider(mngr, configBytes)

	ctx := context.Background()
	// Connect to the Ethereum client
	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Str("rpcurl", *rpcURLArg).Msg("Failed to connect to Ethereum client")
	}
	err = provider.Initialize(ctx, client)
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
