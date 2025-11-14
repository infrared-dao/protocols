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
	brownfiAPIBase = "https://bf-v2-api.brownfi.io/indexer"
)

type brownfiGraphQLQuery struct {
	Query string `json:"query"`
}

type brownfiResponse struct {
	Data struct {
		Pair struct {
			APR float64 `json:"apr"`
			FEE float64 `json:"protocolFee"`
		} `json:"pair"`
	} `json:"data"`
}

func FetchBrownfiAPRs(ctx context.Context, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	// Normalize token addresses to lowercase
	normalizedTokens := make([]string, len(stakingTokens))
	for idx, tokenAddress := range stakingTokens {
		normalizedTokens[idx] = strings.ToLower(tokenAddress)
	}

	// Build the GraphQL query
	query := `{
		pair(chainId: 80094, address: "%s") {
			apr,
			protocolFee
		}}`

	scalePercent := decimal.NewFromFloat(100.0)
	brownfiAPRs := make(map[string]decimal.Decimal)

	for _, tokenAddress := range normalizedTokens {
		requestBody := brownfiGraphQLQuery{
			Query: fmt.Sprintf(query, tokenAddress),
		}

		requestJSON, err := json.Marshal(requestBody)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal GraphQL query: %w", err)
		}

		httpParams := HTTPParams{
			URL: brownfiAPIBase,
			Headers: map[string]string{
				"Content-Type": "application/json",
				"Accept":       "application/json",
			},
			RequestBody: requestJSON,
		}

		responseJSON, err := HTTPPost(ctx, httpParams)
		if err != nil {
			err = fmt.Errorf("failed to fetch brownfi pool data, %w", err)
			log.Error().Msg(err.Error())
			return nil, err
		}

		var results brownfiResponse
		err = json.Unmarshal(responseJSON, &results)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal brownfi response: %w", err)
		}

		// base APR = response_apr * (1 - response_fee)
		apr := decimal.NewFromFloat(results.Data.Pair.APR)
		fee := decimal.NewFromFloat(results.Data.Pair.FEE)

		// returns as percentage (15.5% should be 0.155), so divide by 100
		baseAPR := apr.Div(scalePercent)
		baseAPR = baseAPR.Mul(decimal.NewFromFloat(1.0).Sub(fee))

		brownfiAPRs[tokenAddress] = baseAPR
	}

	return brownfiAPRs, nil
}
