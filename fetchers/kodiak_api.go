package fetchers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

const (
	kodiakAPI    = "https://api.goldsky.com/api/public/project_clpx84oel0al201r78jsl0r3i/subgraphs/kodiak-v3-berachain-mainnet/latest/gn"
	kodiakQuery  = `{ kodiakAprs(where: {id_in: [\"%s\"]}) { id, averageApr, timestamp } }`
	ageThreshold = 14400 // 4 hours, averageAPR observations older than this will be ignored
)

type kodiakAPR struct {
	ID        string `json:"id"`
	APR       string `json:"averageApr"`
	Timestamp string `json:"timestamp"`
}

type kodiakResponse struct {
	Data struct {
		APRs []kodiakAPR `json:"kodiakAprs"`
	} `json:"data"`
}

func FetchKodiakAPRs(ctx context.Context, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	for idx, tokenAddress := range stakingTokens {
		stakingTokens[idx] = strings.ToLower(tokenAddress)
	}
	tokensFilter := strings.Join(stakingTokens, `\", \"`)
	query := fmt.Sprintf(kodiakQuery, tokensFilter)
	jsonQuery := []byte(`{"query": "` + query + `"}`)

	params := HTTPParams{
		URL: kodiakAPI,
		Headers: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
			"Accept":       "application/json",
		},
		RequestBody: bytes.NewBuffer(jsonQuery).Bytes(),
		MaxWait:     DefaultRequestTimeout,
	}

	responseJSON, err := HTTPPost(ctx, params)
	if err != nil {
		err = fmt.Errorf("failed to fetch kodiak APR data, %w", err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	var results kodiakResponse
	err = json.Unmarshal(responseJSON, &results)
	if err != nil {
		return nil, err
	}

	scalePercent := decimal.NewFromFloat(100.0)

	kodiakAPRs := make(map[string]decimal.Decimal)
	for _, aprData := range results.Data.APRs {
		avgAPR, err := decimal.NewFromString(aprData.APR)
		if err != nil {
			continue // don't error if there are other results we could process
		}

		// Kodiak API returns literal percents but we divide by 100 for consistency
		avgAPR = avgAPR.Div(scalePercent)

		unixTimeInt, err := strconv.ParseInt(aprData.Timestamp, 10, 64)
		if err != nil {
			continue // don't error if there are other results we could process
		}
		dataTime := time.Unix(unixTimeInt, 0)
		now := time.Now()
		age := now.Sub(dataTime).Seconds()
		if age < ageThreshold {
			kodiakAPRs[aprData.ID] = avgAPR
		}
	}

	return kodiakAPRs, nil
}
