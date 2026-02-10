package fetchers

import (
	"context"
	"fmt"
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
func FetchDolomiteAPRs(ctx context.Context, client HttpClient, underlyingTokens []string) (map[string]decimal.Decimal, error) {
	if len(underlyingTokens) == 0 {
		return nil, nil
	}

	for idx, tokenAddress := range underlyingTokens {
		underlyingTokens[idx] = strings.ToLower(tokenAddress)
	}

	var results dolomiteResponse
	err := client.DoJSON(ctx, http.MethodGet, dolomiteAPI, nil, &results,
		WithHeader("Accept", "application/json"))
	if err != nil {
		err = fmt.Errorf("failed to fetch dolomite interest rate data, %w", err)
		log.Error().Msg(err.Error())
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
