package fetchers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

const (
	bexAPI   = "https://api.berachain.com/"
	bexQuery = `{ poolGetPools(first: 1000) { address, dynamicData { aprItems { apr } } } }`
	// There are currently less than 300 BEX pools, can get 1000 per page.... future todo: add paging logic
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
	Data struct {
		Pools []bexPool `json:"poolGetPools"`
	} `json:"data"`
}

func FetchBexAPRs(stakingTokens []string) (map[string]decimal.Decimal, error) {
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
	jsonQuery := []byte(`{"query": "` + bexQuery + `"}`)

	request, err := http.NewRequest("POST", bexAPI, bytes.NewBuffer(jsonQuery))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseJSON, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed to read bex response.Body: %v", response.Body)
		log.Error().Msg(err.Error())
		return nil, err
	}

	var results bexResponse
	err = json.Unmarshal(responseJSON, &results)
	if err != nil {
		return nil, err
	}

	bexAPRs := make(map[string]decimal.Decimal)
	for _, poolData := range results.Data.Pools {
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
