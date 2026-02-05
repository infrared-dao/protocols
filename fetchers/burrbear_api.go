package fetchers

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

// Pool APR data is kept on this public but secret endpoint not the graphql in their public docs
const (
	burrbearAPI   = "https://o27ufpzifl.execute-api.us-east-1.amazonaws.com/new-prd/api/graphql"
	burrbearQuery = `{ rewards { lps { apr { min, max, components { type, value, rewardName } }, token { address } } } }`
)

type burrbearResponse struct {
	Rewards []struct {
		Lps []struct {
			Apr struct {
				Max        string `json:"max"`
				Components []struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"components"`
			} `json:"apr"`
			Token struct {
				Address string `json:"address"`
			} `json:"token"`
		} `json:"lps"`
	} `json:"rewards"`
}

func FetchBurrBearAPRs(ctx context.Context, client HttpClient, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	for idx, tokenAddress := range stakingTokens {
		stakingTokens[idx] = strings.ToLower(tokenAddress)
	}

	var results burrbearResponse
	err := client.DoGraphQL(ctx, burrbearAPI, burrbearQuery, nil, &results,
		WithHeader("Content-Type", "application/json; charset=UTF-8"),
		WithHeader("Accept", "application/json"))
	if err != nil {
		err = fmt.Errorf("failed to fetch burrbear APR data, %w", err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	scalePercent := decimal.NewFromFloat(100.0)
	burrbearAPRs := make(map[string]decimal.Decimal)

	for _, rewards := range results.Rewards {
		for _, aprData := range rewards.Lps {
			token := strings.ToLower(aprData.Token.Address)
			for _, apr := range aprData.Apr.Components {
				if apr.Type != "SWAP_FEES" {
					// Ignore their "base APR", POL APR, points APR etc.
					// only swap fees are real baseAPR we want to record
					continue
				}
				swapAPR, err := strconv.ParseFloat(apr.Value, 64)
				if err != nil {
					// log the error but just continue in case others could be handled
					log.Error().Msgf("burrbear fetcher unable to parse %s as float for pool %s",
						apr.Value, token)
				}
				baseAPR := decimal.NewFromFloat(swapAPR)
				if slices.Contains(stakingTokens, token) {
					// this API returns 17.1% as 17.1 so divide by 100 before output
					baseAPR = baseAPR.Div(scalePercent)
					burrbearAPRs[token] = baseAPR
				}
			}
		}
	}

	return burrbearAPRs, nil
}
