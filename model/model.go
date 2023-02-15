package model

import (
	"fmt"
	"time"
)

type Status struct {
	Timestamp    time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
}

type Data struct {
	Id          int              `json:"id"`
	Name        string           `json:"name"`
	Symbol      string           `json:"symbol"`
	LastUpdated time.Time        `json:"last_updated"`
	Quote       map[string]Quote `json:"quote"`
}

func (d *Data) CacheKey(currency string) (string, error) {
	if _, exist := d.Quote[currency]; exist {
		return fmt.Sprintf("%s_%s", currency, d.Symbol), nil
	}

	return "", fmt.Errorf("currency %s is not available", currency)
}

type Quote struct {
	Price                 float64   `json:"price"`
	Volume24H             float64   `json:"volume_24h"`
	VolumeChange24H       float64   `json:"volume_change_24h"`
	PercentChange1H       float64   `json:"percent_change_1h"`
	PercentChange24H      float64   `json:"percent_change_24h"`
	PercentChange7D       float64   `json:"percent_change_7d"`
	PercentChange30D      float64   `json:"percent_change_30d"`
	PercentChange60D      float64   `json:"percent_change_60d"`
	PercentChange90D      float64   `json:"percent_change_90d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time `json:"last_updated"`
}

type FetcherResponse struct {
	Status Status          `json:"status"`
	Data   map[string]Data `json:"data"`
}
