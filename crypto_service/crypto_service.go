package crypto_service

import (
	"strings"

	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
)

type result struct {
	Name             string
	Symbol           string
	PercentChange1h  float64
	PercentChange24h float64
	PercentChange1w  float64
	Price            float64
	Rank             int64
}

func SearchByName(name string) ([]*result, error) {

	resultMap := []*result{}

	paprikaClient := coinpaprika.NewClient(nil)

	tickers, err := paprikaClient.Tickers.List(nil)
	if err != nil {
		panic(err)
	}
	for idx, t := range tickers {
		if t.Name == nil || t.Symbol == nil || t.Rank == nil {
			continue
		}
		if strings.ContainsAny(strings.ToLower(*t.Name), strings.ToLower(name)) {
			if quoteUSD, ok := t.Quotes["USD"]; ok {
				newResult := new(result)
				newResult.Name = *t.Name
				newResult.Rank = *t.Rank
				newResult.Symbol = *t.Symbol
				newResult.Price = *quoteUSD.Price
				newResult.PercentChange1h = *quoteUSD.PercentChange1h
				newResult.PercentChange1w = *quoteUSD.PercentChange7d
				newResult.PercentChange24h = *quoteUSD.PercentChange24h
				resultMap = append(resultMap, newResult)
			}
		}
		if idx >= 10 {
			break
		}
	}
	return resultMap, nil
}
