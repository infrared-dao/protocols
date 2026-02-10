package fetchers

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

const (
	wasabeeAPIBase = "https://api.dodoex.io/frontend-graphql?opname=FetchAllVaults"
)

type wasabeeResponse struct {
	Vaults []struct {
		ID     string `json:"id"`
		APR_7D string `json:"feeApr_7d"`
	} `json:"berachain_alm_allVaults"`
}

func FetchWasabeeAPRs(ctx context.Context, client HttpClient, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	// Normalize token addresses to lowercase
	normalizedTokens := make([]string, len(stakingTokens))
	for idx, tokenAddress := range stakingTokens {
		normalizedTokens[idx] = strings.ToLower(tokenAddress)
	}

	// Simplified GraphQL query
	query := `{ berachain_alm_allVaults { id, feeApr_7d } }`

	var results wasabeeResponse
	err := client.DoGraphQL(ctx, wasabeeAPIBase, query, nil, &results,
		WithHeader("Content-Type", "application/json"),
		WithHeader("Accept", "application/json"))
	if err != nil {
		err = fmt.Errorf("failed to fetch Wasabee pool data, %w", err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	scalePercent := decimal.NewFromFloat(100.0)
	wasabeeAPRs := make(map[string]decimal.Decimal)

	// Process vaults and match against requested tokens
	for _, vault := range results.Vaults {
		vaultID := strings.ToLower(vault.ID)

		if !slices.Contains(normalizedTokens, vaultID) {
			continue
		}

		// Get the 7d fee APR for the vault
		aprString := vault.APR_7D
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

			wasabeeAPRs[vaultID] = baseAPR
		}
	}

	return wasabeeAPRs, nil
}
