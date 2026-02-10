package fetchers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

/*
Caution: Aquabera API absolutely requires the Mixed case ETH address of the token
Incoming addresses are mostly lowercase from DB to explicitly convert to mixed case
*/

const aquaberaAPI = "https://a8ejzahwod.execute-api.us-east-1.amazonaws.com/prod/vaults/80094/%s"

type aquaberaResponse struct {
	Address   string `json:"address"`
	VaultName string `json:"displayName"`
	Apr       struct {
		WeekAvg  float64 `json:"7d"`
		MonthAvg float64 `json:"30d"`
	} `json:"apr"`
}

func BuildAquaberaAPRsFetcher(apiKey string) func(context.Context, HttpClient, []string) (map[string]decimal.Decimal, error) {
	return func(ctx context.Context, client HttpClient, stakingTokens []string) (map[string]decimal.Decimal, error) {
		return FetchAquaberaAPRs(ctx, client, stakingTokens, apiKey)
	}
}

func FetchAquaberaAPRs(ctx context.Context, client HttpClient, stakingTokens []string, apiKey string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	scalePercent := decimal.NewFromFloat(100.0)
	aquaberaAPRs := make(map[string]decimal.Decimal)

	for _, tokenAddress := range stakingTokens {
		// We must convert the address to the MixedCase address for checksum
		address := common.HexToAddress(tokenAddress)
		mixedAddress := common.NewMixedcaseAddress(address)
		mixedAddressHex := mixedAddress.Address().Hex()

		endpoint := fmt.Sprintf(aquaberaAPI, mixedAddressHex)

		var results aquaberaResponse
		err := client.DoJSON(ctx, http.MethodGet, endpoint, nil, &results,
			WithHeader("Content-Type", "application/json; charset=UTF-8"),
			WithHeader("Accept", "application/json"),
			WithHeader("x-api-key", apiKey))
		if err != nil {
			err = fmt.Errorf("failed to fetch aquabera vault data, %w", err)
			log.Error().Msg(err.Error())
			continue
		}

		baseAPR := decimal.NewFromFloat(results.Apr.WeekAvg)
		baseAPR = baseAPR.Div(scalePercent)
		aquaberaAPRs[tokenAddress] = baseAPR
	}

	return aquaberaAPRs, nil
}
