package fetchers

import (
	"context"
	"fmt"
	"net/http"
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

func FetchWeberaAPRs(ctx context.Context, client HttpClient, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	weberaAPRs := make(map[string]decimal.Decimal)
	for _, tokenAddress := range stakingTokens {
		tokenAddress = strings.ToLower(tokenAddress)

		var results weberaResponse
		err := client.DoJSON(ctx, http.MethodGet, weberaAPI+tokenAddress+"/apr", nil, &results,
			WithHeader("Content-Type", "application/json; charset=UTF-8"),
			WithHeader("Accept", "application/json"))
		if err != nil {
			err = fmt.Errorf("failed to fetch webera APR data, %w", err)
			log.Error().Msg(err.Error())
			return nil, err
		}

		avgAPR := decimal.NewFromFloat(results.Data.APR)
		weberaAPRs[tokenAddress] = avgAPR
	}

	return weberaAPRs, nil
}
