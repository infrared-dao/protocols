package fetchers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

/*
Caution: Aquabera API absolutely requires the Mixed case ETH address of the token
Incoming addresses are mostly lowercase from DB to explicitly convert to mixed case
*/

const (
	aquaberaAPI = "https://a8ejzahwod.execute-api.us-east-1.amazonaws.com/prod/vaults/80094/%s"
	apiKey      = "tVmbB2ybYf1XnV9hNKHil3tnRcx9rGHS20kiP1pT"
)

type aquaberaResponse struct {
	Address   string `json:"address"`
	VaultName string `json:"displayName"`
	Apr       struct {
		WeekAvg  float64 `json:"7d"`
		MonthAvg float64 `json:"30d"`
	} `json:"apr"`
}

func FetchAquaberaAPRs(ctx context.Context, stakingTokens []string) (map[string]decimal.Decimal, error) {
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
		params := HTTPParams{
			URL: endpoint,
			Headers: map[string]string{
				"Content-Type": "application/json; charset=UTF-8",
				"Accept":       "application/json",
				"x-api-key":    apiKey,
			},
		}

		responseJSON, err := HTTPGet(ctx, params)
		if err != nil {
			err = fmt.Errorf("failed to fetch aquabera vault data, %w", err)
			log.Error().Msg(err.Error())
			continue
		}

		var results aquaberaResponse
		err = json.Unmarshal(responseJSON, &results)
		if err != nil {
			log.Error().Msg("fetcher unable to parse aquabera response: " + string(responseJSON))
			continue
		}

		baseAPR := decimal.NewFromFloat(results.Apr.WeekAvg)
		baseAPR = baseAPR.Div(scalePercent)
		aquaberaAPRs[tokenAddress] = baseAPR
	}

	return aquaberaAPRs, nil
}
