package fetchers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

/*
Pendle API gives many different APY values but the base APR which users get just for holding
LP tokens without anything additional (no vePendle holding "Boosts") is called "aggregatedApy"
*/

const (
	pendleAPI  = "https://api-v2.pendle.finance/core/v2/80094/markets/%s/data"
	ageTimeout = 14400 // 4 hours, averageAPR observations older than this will be ignored
)

type pendleResponse struct {
	BaseAPR   float64 `json:"aggregatedAPY"`
	Timestamp string  `json:"timestamp"`
}

// Important: pendle tokens are wrapped so the staking token is the wrapped token
// Addresses in pools[] must be unwrapped tokens, ie. underlying token to the staking token
func FetchPendleAPRs(ctx context.Context, client HttpClient, pools []string) (map[string]decimal.Decimal, error) {
	if len(pools) == 0 {
		return nil, nil
	}

	pendleAPRs := make(map[string]decimal.Decimal)
	for _, poolAddress := range pools {
		endpoint := fmt.Sprintf(pendleAPI, strings.ToLower(poolAddress))

		var results pendleResponse
		err := client.DoJSON(ctx, http.MethodGet, endpoint, nil, &results,
			WithHeader("Content-Type", "application/json; charset=UTF-8"),
			WithHeader("Accept", "application/json"))
		if err != nil {
			err = fmt.Errorf("failed to fetch pendle pool data, %w", err)
			log.Error().Msg(err.Error())
			return nil, err
		}

		baseAPR := decimal.NewFromFloat(results.BaseAPR)
		dataTime, err := time.Parse("2006-01-02T15:04:05.000Z", results.Timestamp)
		if err != nil {
			log.Error().Msg(err.Error())
			continue // don't error if there are other results we could process
		}
		now := time.Now()
		age := now.Sub(dataTime).Seconds()
		if age < ageTimeout {
			pendleAPRs[poolAddress] = baseAPR
		}
	}

	return pendleAPRs, nil
}
