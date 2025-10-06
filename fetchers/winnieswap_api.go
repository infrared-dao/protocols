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

const (
	winnieswapAPIBase = "https://sub.winnieswap.com/"
)

type winnieswapGraphQLQuery struct {
	Query string `json:"query"`
}

type winnieswapResponse struct {
	Data struct {
		Vaults struct {
			Items []struct {
				ID      string `json:"id"`
				Pool    string `json:"pool"`
				Name    string `json:"name"`
				PoolRef struct {
					PoolDayData struct {
						Items []struct {
							APR string `json:"apr"`
						} `json:"items"`
					} `json:"poolDayData"`
				} `json:"poolRef"`
			} `json:"items"`
		} `json:"stickyVaults"`
	} `json:"data"`
}

func FetchWinnieSwapAPRs(ctx context.Context, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	// Normalize token addresses to lowercase
	normalizedTokens := make([]string, len(stakingTokens))
	for idx, tokenAddress := range stakingTokens {
		normalizedTokens[idx] = strings.ToLower(tokenAddress)
	}

	// Build the GraphQL query
	query := `query VaultQuery {
		stickyVaults {
			items {
				id,
				pool,
				name,  
				poolRef {
					poolDayData(orderBy: "date", orderDirection: "desc", limit: 1) {
						items {
							apr
						}
					}
				}
			}  
		}}`

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

	var results winnieswapResponse
	err = json.Unmarshal(responseJSON, &results)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal WinnieSwap response: %w", err)
	}

	scalePercent := decimal.NewFromFloat(100.0)
	winnieswapAPRs := make(map[string]decimal.Decimal)

	// Process vaults and match against requested tokens
	for _, vault := range results.Data.Vaults.Items {
		vaultID := strings.ToLower(vault.ID)

		if !slices.Contains(normalizedTokens, vaultID) {
			continue
		}

		// Get the latest APR from poolDayData
		if len(vault.PoolRef.PoolDayData.Items) > 0 {
			aprString := vault.PoolRef.PoolDayData.Items[0].APR
			if aprString != "" {
				apr, err := decimal.NewFromString(aprString)
				if err != nil {
					log.Warn().
						Str("vaultID", vaultID).
						Str("aprString", aprString).
						Err(err).
						Msg("failed to parse APR value")
					continue
				}

				// returns as percentage (15.5% should be 0.155), so divide by 100
				baseAPR := apr.Div(scalePercent)

				winnieswapAPRs[vaultID] = baseAPR
			}
		}
	}

	return winnieswapAPRs, nil
}
