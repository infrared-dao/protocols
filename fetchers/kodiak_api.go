package fetchers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

const (
	kodiakAPIBase        = "https://backend.kodiak.finance"
	kodiakVaultsEndpoint = "/vaults"
	kodiakChainID        = "80094" // Berachain mainnet
	kodiakAgeThreshold   = 14400   // 4 hours, APR observations older than this will be ignored
)

type kodiakVault struct {
	ID          string  `json:"id"`
	TotalApr    float64 `json:"totalApr"`
	Apr         float64 `json:"apr"`
	TVL         float64 `json:"tvl"`
	LastUpdated string  `json:"lastUpdated"`
	LastIndexed string  `json:"lastIndexed"`
}

type kodiakVaultsResponse struct {
	Data  []kodiakVault `json:"data"`
	Count int           `json:"count"`
}

func FetchKodiakAPRs(ctx context.Context, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	// Normalize token addresses to lowercase
	normalizedTokens := make([]string, len(stakingTokens))
	for idx, tokenAddress := range stakingTokens {
		normalizedTokens[idx] = strings.ToLower(tokenAddress)
	}

	// Build the API URL
	apiURL := kodiakAPIBase + kodiakVaultsEndpoint
	params := url.Values{}
	params.Add("chainId", kodiakChainID)
	params.Add("limit", "200") // Get more vaults to ensure we find matches

	fullURL := fmt.Sprintf("%s?%s", apiURL, params.Encode())

	httpParams := HTTPParams{
		URL: fullURL,
		Headers: map[string]string{
			"Accept": "application/json",
		},
	}

	responseJSON, err := HTTPGet(ctx, httpParams)
	if err != nil {
		err = fmt.Errorf("failed to fetch kodiak vault data, %w", err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	var results kodiakVaultsResponse
	err = json.Unmarshal(responseJSON, &results)
	if err != nil {
		return nil, err
	}

	scalePercent := decimal.NewFromFloat(100.0)
	kodiakAPRs := make(map[string]decimal.Decimal)

	// Process vaults and match against requested tokens
	for _, vault := range results.Data {
		vaultID := strings.ToLower(vault.ID)

		// Check if this vault ID matches any of our requested tokens
		for _, requestedToken := range normalizedTokens {
			if vaultID == requestedToken {
				// Check if the data is recent enough
				if vault.LastUpdated != "" {
					updatedTime, err := time.Parse(time.RFC3339, vault.LastUpdated)
					if err != nil {
						continue
					}

					now := time.Now()
					age := now.Sub(updatedTime).Seconds()
					if age >= kodiakAgeThreshold {
						continue
					}
				}

				// Use apr which is the base APR from swap fees (excluding staking rewards)
				baseAPR := decimal.NewFromFloat(vault.Apr)

				// Kodiak REST API returns APR as percentages (e.g., 15.5 for 15.5%), so divide by 100
				baseAPR = baseAPR.Div(scalePercent)

				kodiakAPRs[vaultID] = baseAPR
				break
			}
		}
	}

	return kodiakAPRs, nil
}
