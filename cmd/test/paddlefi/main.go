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
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	lpTokenArg := flag.String("address", "0xe16761787cF9bB0D3fC2E5C726dAe906ce81B102", "Pool contract address")
	pricesArg := flag.String("prices", "0x6969696969696969696969696969696969696969:2.05:18", "address:price:decimals")
	rpcURLArg := flag.String("rpcurl", "https://your-rpc-url", "RPC URL")
	flag.Parse()

	priceMap := make(map[string]protocols.Price)
	for _, p := range strings.Split(*pricesArg, ",") {
		parts := strings.Split(p, ":")
		price, _ := decimal.NewFromString(parts[1])
		decimals, _ := strconv.Atoi(parts[2])
		priceMap[strings.ToLower(parts[0])] = protocols.Price{Decimals: uint(decimals), Price: price}
	}

	client, err := ethclient.Dial(*rpcURLArg)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to RPC")
	}

	ctx := context.Background()

	adapter := &protocols.PaddleFiProvider{}
	configBytes, err := adapter.GetConfig(ctx, *lpTokenArg, client)
	if err != nil {
		logger.Fatal().Err(err).Msg("GetConfig failed")
	}

	addr := common.HexToAddress(*lpTokenArg)
	adapter = protocols.NewPaddleFiProvider(addr, nil, priceMap, logger, configBytes)
	if err := adapter.Initialize(ctx, client); err != nil {
		logger.Fatal().Err(err).Msg("Adapter init failed")
	}

	tvl, err := adapter.TVL(ctx)
	if err != nil {
		logger.Fatal().Err(err).Msg("TVL fetch failed")
	}
	logger.Info().Str("TVL (USD)", tvl).Msg("TVL fetched successfully")
}
