package fetchers

import (
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
	dolomiteAPI = "https://api.dolomite.io/tokens/80094/interest-rates"
)

type dolomiteRates struct {
	Token struct {
		Address string `json:"tokenAddress"`
	} `json:"token"`
	SupplyAPR string `json:"totalSupplyInterestRate"`
}

type dolomiteResponse struct {
	Rates []dolomiteRates `json:"interestRates"`
}

// Dolomite API has no reference to the actual stake token, only the underlying token
func FetchDolomiteAPRs(underlyingTokens []string) (map[string]decimal.Decimal, error) {
	if len(underlyingTokens) == 0 {
		return nil, nil
	}

	for idx, tokenAddress := range underlyingTokens {
		underlyingTokens[idx] = strings.ToLower(tokenAddress)
	}

	request, err := http.NewRequest("GET", dolomiteAPI, nil)
	if err != nil {
		return nil, err
	}
	//request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseJSON, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response.Body: %v", response.Body)
		log.Error().Msg(err.Error())
		return nil, err
	}

	var results dolomiteResponse
	err = json.Unmarshal(responseJSON, &results)
	if err != nil {
		return nil, err
	}

	dolomiteAPRs := make(map[string]decimal.Decimal)
	for _, rateData := range results.Rates {
		address := strings.ToLower(rateData.Token.Address)
		if !slices.Contains(underlyingTokens, address) {
			continue
		}

		apr, err := decimal.NewFromString(rateData.SupplyAPR)
		if err != nil {
			continue // don't error if there are other results we could process
		}

		dolomiteAPRs[address] = apr
	}

	return dolomiteAPRs, nil
}
