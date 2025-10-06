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
	winnieswapAPIBase = "https://sub.winnieswap.com/"
)

type winnieswapGraphQLQuery struct {
	Query string `json:"query"`
}

type winnieswapPoolDayData struct {
	APR string `json:"apr"`
}

type winnieswapPoolDayDataWrapper struct {
	Items []winnieswapPoolDayData `json:"items"`
}

type winnieswapPool struct {
	ID          string                           `json:"id"`
	PoolDayData winnieswapPoolDayDataWrapper     `json:"poolDayData"`
}

type winnieswapPoolsWrapper struct {
	Items []winnieswapPool `json:"items"`
}

type winnieswapGraphQLData struct {
	Pools winnieswapPoolsWrapper `json:"pools"`
}

type winnieswapGraphQLResponse struct {
	Data winnieswapGraphQLData `json:"data"`
}

func FetchWinnieSwapAPRs(ctx context.Context, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	// Normalize token addresses to lowercase
	normalizedTokens := make(map[string]bool)
	for _, tokenAddress := range stakingTokens {
		normalizedTokens[strings.ToLower(tokenAddress)] = true
	}

	// Build the GraphQL query
	query := `query MyQuery {
  pools {
    items {
      poolDayData(orderBy: "date", orderDirection: "desc", limit: 1) {
        items {
          apr
        }
      }
      id
    }
  }
}`

	requestBody := winnieswapGraphQLQuery{
		Query: query,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GraphQL query: %w", err)
	}

	httpParams := HTTPParams{
		URL: winnieswapAPIBase,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
		RequestBody: requestJSON,
	}

	responseJSON, err := HTTPPost(ctx, httpParams)
	if err != nil {
		err = fmt.Errorf("failed to fetch WinnieSwap pool data, %w", err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	var results winnieswapGraphQLResponse
	err = json.Unmarshal(responseJSON, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal WinnieSwap response: %w", err)
	}

	winnieswapAPRs := make(map[string]decimal.Decimal)

	// Process pools and match against requested tokens
	for _, pool := range results.Data.Pools.Items {
		poolID := strings.ToLower(pool.ID)

		// Check if this pool ID matches any of our requested tokens
		if _, found := normalizedTokens[poolID]; !found {
			continue
		}

		// Get the latest APR from poolDayData
		if len(pool.PoolDayData.Items) > 0 {
			aprString := pool.PoolDayData.Items[0].APR
			if aprString != "" {
				apr, err := decimal.NewFromString(aprString)
				if err != nil {
					log.Warn().
						Str("poolID", poolID).
						Str("aprString", aprString).
						Err(err).
						Msg("failed to parse APR value")
					continue
				}

				// Assuming the API returns APR as a decimal (e.g., 0.155 for 15.5%)
				// If it returns as percentage (15.5), divide by 100
				// Adjust this based on the actual API response format
				winnieswapAPRs[poolID] = apr
			}
		}
	}

	return winnieswapAPRs, nil
}
