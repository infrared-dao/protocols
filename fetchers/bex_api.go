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
	bexAPI   = "https://api.berachain.com/"
	bexQuery = `{ poolGetPools(first: 10000) { address, dynamicData { aprItems { apr } } } }`
	// There are currently less than 300 BEX pools, can get 10000 per page....
	// TODO: add paging logic when number of BEX pools gets close
)

type bexPool struct {
	Address string `json:"address"`
	Dynamic struct {
		Items []struct {
			APR float64 `json:"apr"`
		} `json:"aprItems"`
	} `json:"dynamicData"`
}

type bexResponse struct {
	Pools []bexPool `json:"poolGetPools"`
}

func FetchBexAPRs(ctx context.Context, client HttpClient, stakingTokens []string) (map[string]decimal.Decimal, error) {
	if len(stakingTokens) == 0 {
		return nil, nil
	}

	for idx, tokenAddress := range stakingTokens {
		stakingTokens[idx] = strings.ToLower(tokenAddress)
	}

	// Unfortunately BEX API doesn't have very useful filtering
	// Would like to filter by token or address but neither is reliable
	// Best filter is by ID but would need to get ID from the bex config data
	// Since there are only about 300 Bex pools just load all and filter in code
	var results bexResponse
	err := client.DoGraphQL(ctx, bexAPI, bexQuery, nil, &results,
		WithHeader("Content-Type", "application/json; charset=UTF-8"),
		WithHeader("Accept", "application/json"))
	if err != nil {
		return nil, err
	}

	// log a warning as we get close to page limit but continue processing
	if len(results.Pools) > 9000 {
		err = fmt.Errorf("over 9000 bex pools found, approaching limit")
		log.Error().Msg(err.Error())
	}

	bexAPRs := make(map[string]decimal.Decimal)
	for _, poolData := range results.Pools {
		if len(poolData.Dynamic.Items) == 0 {
			continue
		}

		address := strings.ToLower(poolData.Address)
		if !slices.Contains(stakingTokens, address) {
			continue
		}

		// BEX already returns percents scaled by 100, no need to divide
		aprFloat := poolData.Dynamic.Items[0].APR
		apr := decimal.NewFromFloat(aprFloat)

		bexAPRs[address] = apr
	}

	return bexAPRs, nil
}
