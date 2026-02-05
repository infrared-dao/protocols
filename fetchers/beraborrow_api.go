package fetchers

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

// According to Beraborrow team -- only sNECT has a base APR from beraborrow themselves

// All other beraborrow vaults, the asset could have an internal APR (like wgBERA) but
// this is global for the underlying asset not from the protocol of beraborrow

// Therefore any staking token address other than sNECT we should explicitly return 0%

const (
	beraborrowAPI    = "https://api.goldsky.com/api/public/project_cm0v01jq86ry701rr6jta9tqm/subgraphs/bera-borrow-prod/1.0.20/gn"
	snectID          = "SharePool"
	snectAddress     = "0x597877ccf65be938bd214c4c46907669e3e62128"
	nectAddress      = "0x1ce0a25d13ce4d52071ae7e02cf1f6606f4c79d3"
	roundingDecimals = 8
)

type beraborrowResponse struct {
	Pools []struct {
		ID        string `json:"id"`
		Pool      string `json:"pool_address"`
		DebtToken string `json:"debt_token"`
		Latest    []struct {
			LastTime  string `json:"lastTime"`
			LastPrice string `json:"lastPrice"`
		} `json:"latest"`
		First []struct {
			StartTime  string `json:"startTime"`
			StartPrice string `json:"startPrice"`
		} `json:"first"`
	} `json:"sharePools"`
}

func FetchBeraborrowAPRs(ctx context.Context, client HttpClient, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	// Normalize token addresses to lowercase
	normalizedTokens := make([]string, len(stakingTokens))
	for idx, tokenAddress := range stakingTokens {
		normalizedTokens[idx] = strings.ToLower(tokenAddress)
	}

	// Build the GraphQL query with starting unix timestamp to be filled in
	query := `query SharePoolQuery {
		sharePools {
			id,
			pool_address,
			debt_token,
			latest: buckets(first: 1, orderBy: lastTime, orderDirection: desc) {
				lastTime,
				lastPrice
			},
			first: buckets(first: 1, orderBy: lastTime, orderDirection: desc, where: { startTime_lte: %d bucketType:DAILY}) {
				startTime
				startPrice
			}

		}}`

	// Beraborrow team suggests looking back at least 2 months because liquiditations happen
	// seldomly and at irregular intervals -- a long period is needed for accurate average
	startingPoint := time.Now().Add(time.Duration(-75*24) * time.Hour) // 75 days ago
	query = fmt.Sprintf(query, startingPoint.Unix())

	var results beraborrowResponse
	err := client.DoGraphQL(ctx, beraborrowAPI, query, nil, &results,
		WithHeader("Content-Type", "application/json"),
		WithHeader("Accept", "application/json"))
	if err != nil {
		err = fmt.Errorf("failed to fetch Beraborrow pool data, %w", err)
		log.Error().Msg(err.Error())
		return nil, err
	}

	beraborrowAPRs := make(map[string]decimal.Decimal)

	// Process staking tokens, assigning 0% base APR to all except sNECT
	for _, token := range normalizedTokens {
		if token != snectAddress {
			beraborrowAPRs[token] = decimal.Zero
			continue
		}

		// Validate the response for share pool matches expectations:
		// There should only be 1 share pool
		// It must match the addresses and IDs associated with sNECT
		// Both first and last bucket should only have one time and price
		if len(results.Pools) != 1 ||
			results.Pools[0].ID != snectID ||
			strings.ToLower(results.Pools[0].Pool) != snectAddress ||
			strings.ToLower(results.Pools[0].DebtToken) != nectAddress ||
			len(results.Pools[0].First) != 1 ||
			len(results.Pools[0].Latest) != 1 {
			log.Warn().
				Msg("response did not match expected format for sNECT")
			continue
		}

		firstBucket := results.Pools[0].First[0]
		lastBucket := results.Pools[0].Latest[0]

		// Get start and last time and price for computation
		startStamp, err := strconv.ParseInt(firstBucket.StartTime, 10, 64)
		if err != nil {
			log.Warn().
				Str("startTime", firstBucket.StartTime).
				Err(err).
				Msg("failed to parse start Unix timestamp")
			continue
		}
		startTime := time.Unix(startStamp, 0)

		lastStamp, err := strconv.ParseInt(lastBucket.LastTime, 10, 64)
		if err != nil {
			log.Warn().
				Str("lastTime", lastBucket.LastTime).
				Err(err).
				Msg("failed to parse last Unix timestamp")
			continue
		}
		lastTime := time.Unix(lastStamp, 0)

		startPrice, err := decimal.NewFromString(firstBucket.StartPrice)
		if err != nil {
			log.Warn().
				Str("startPrice", firstBucket.StartPrice).
				Err(err).
				Msg("failed to parse start price value")
			continue
		}

		lastPrice, err := decimal.NewFromString(lastBucket.LastPrice)
		if err != nil {
			log.Warn().
				Str("lastPrice", lastBucket.LastPrice).
				Err(err).
				Msg("failed to parse last price value")
			continue
		}

		// Computation of APY (which is close enough to APR for small values)
		// APY = (lastPrice / firstPrice) ^ (year / period length) - 1
		observedSeconds := lastTime.Sub(startTime).Seconds()
		secondsPerYear := 60.0 * 60.0 * 24.0 * 365.0
		annualizedInverse := decimal.NewFromFloat(secondsPerYear / observedSeconds)

		priceRatio := lastPrice.Div(startPrice)
		apyPlusOne := priceRatio.Pow(annualizedInverse)
		apy := apyPlusOne.Sub(decimal.NewFromInt(1))

		beraborrowAPRs[token] = apy.Round(roundingDecimals)
	}

	return beraborrowAPRs, nil
}
