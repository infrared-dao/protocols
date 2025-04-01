package fetchers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

const (
	kodiakAPI    = "https://api.goldsky.com/api/public/project_clpx84oel0al201r78jsl0r3i/subgraphs/kodiak-v3-berachain-mainnet/latest/gn"
	queryAPRs    = `{ kodiakAprs(where: {id_in: [\"%s\"]}) { id, averageApr, timestamp } }`
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

func FetchKodiakAPRs(stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	for idx, tokenAddress := range stakingTokens {
		stakingTokens[idx] = strings.ToLower(tokenAddress)
	}
	tokensFilter := strings.Join(stakingTokens, `\", \"`)
	query := fmt.Sprintf(queryAPRs, tokensFilter)
	jsonQuery := []byte(`{"query": "` + query + `"}`)

	request, err := http.NewRequest("POST", kodiakAPI, bytes.NewBuffer(jsonQuery))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseJSON, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response.Body: %v", response.Body)
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
