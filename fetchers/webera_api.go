package fetchers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

const (
	weberaAPI = "https://be-prod.webera.finance/api/v1/optimizer/"
)

type weberaResponse struct {
	Data struct {
		APR float64 `json:"apr"`
	} `json:"data"`
}

func FetchWeberaAPRs(ctx context.Context, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	weberaAPRs := make(map[string]decimal.Decimal)
	for _, tokenAddress := range stakingTokens {
		tokenAddress = strings.ToLower(tokenAddress)

		params := HTTPParams{
			URL: weberaAPI + tokenAddress + "/apr",
			Headers: map[string]string{
				"Content-Type": "application/json; charset=UTF-8",
				"Accept":       "application/json",
			},
		}

		responseJSON, err := HTTPGet(ctx, params)
		if err != nil {
			err = fmt.Errorf("failed to fetch webera APR data, %w", err)
			log.Error().Msg(err.Error())
			return nil, err
		}

		var results weberaResponse
		err = json.Unmarshal(responseJSON, &results)
		if err != nil {
			return nil, err
		}

		avgAPR := decimal.NewFromFloat(results.Data.APR)
		weberaAPRs[tokenAddress] = avgAPR
	}

	return weberaAPRs, nil
}
