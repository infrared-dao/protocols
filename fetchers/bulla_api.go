package fetchers

import (
	"context"
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

// Bulla exchange in general has very narrow user defined positions not well suited to a single LP token
// All the bulla pools which can be used in POL make use of a partner "gamma" to move positions for users
// The gamma API is what we are really using to get the base APRs for these bulla pools
// Gamma API returns all pools it knows about which is only 14 currently, unlikely to ever grow past a single call

const (
	gammaAPI = "https://api.gamma.xyz/frontend/hypervisors/allDataSummary?chain=berachain&protocol=bulla"
)

type gammaAPR struct {
	Address   string  `json:"address"`
	APR       float64 `json:"feeApr"`
	Timestamp string  `json:"lastUpdated"`
}

func FetchBullaAPRs(ctx context.Context, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	for idx, tokenAddress := range stakingTokens {
		stakingTokens[idx] = strings.ToLower(tokenAddress)
	}

	params := HTTPParams{
		URL: gammaAPI,
		Headers: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
			"Accept":       "application/json",
		},
		MaxWait: DefaultRequestTimeout,
	}

	responseJSON, err := HTTPGet(ctx, params)
	if err != nil {
		err = fmt.Errorf("failed to fetch bulla APR data, %w", err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	var results []gammaAPR
	err = json.Unmarshal(responseJSON, &results)
	if err != nil {
		return nil, err
	}

	bullaAPRs := make(map[string]decimal.Decimal)
	for _, aprData := range results {
		avgAPR := decimal.NewFromFloat(aprData.APR)
		foundAddress := strings.ToLower(aprData.Address)
		if slices.Contains(stakingTokens, foundAddress) {
			bullaAPRs[foundAddress] = avgAPR
		}
	}

	return bullaAPRs, nil
}
