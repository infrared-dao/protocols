package fetchers

import (
	"context"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

const (
	brownfiAPIBase = "https://bf-v2-api.brownfi.io/indexer"
)

type brownfiResponse struct {
	Pair struct {
		APR string `json:"apr"`
		FEE string `json:"protocolFee"`
	} `json:"pair"`
}

func FetchBrownfiAPRs(ctx context.Context, client HttpClient, stakingTokens []string) (map[string]decimal.Decimal, error) {
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
		pair(chainId: 80094, id: "%s") {
			apr,
			protocolFee
		}}`

	scalePercent := decimal.NewFromFloat(100.0)
	brownfiAPRs := make(map[string]decimal.Decimal)

	for _, tokenAddress := range normalizedTokens {
		queryString := fmt.Sprintf(query, tokenAddress)

		var results brownfiResponse
		err := client.DoGraphQL(ctx, brownfiAPIBase, queryString, nil, &results,
			WithHeader("Content-Type", "application/json"),
			WithHeader("Accept", "application/json"))
		if err != nil {
			err = fmt.Errorf("failed to fetch brownfi pool data, %w", err)
			log.Error().Msg(err.Error())
			return nil, err
		}

		// base APR = response_apr * (1 - response_fee)
		apr, err := decimal.NewFromString(results.Pair.APR)
		if err != nil {
			err = fmt.Errorf("failed to parse APR value, %w", err)
			log.Error().Msg(err.Error())
			return nil, err
		}
		fee, err := decimal.NewFromString(results.Pair.FEE)
		if err != nil {
			err = fmt.Errorf("failed to parse protocolFee value, %w", err)
			log.Error().Msg(err.Error())
			return nil, err
		}

		// returns as percentage (15.5% should be 0.155), so divide by 100
		baseAPR := apr.Div(scalePercent)
		baseAPR = baseAPR.Mul(decimal.NewFromFloat(1.0).Sub(fee))

		brownfiAPRs[tokenAddress] = baseAPR
	}

	return brownfiAPRs, nil
}
