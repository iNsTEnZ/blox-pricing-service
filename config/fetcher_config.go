package config

import (
	"blox/util"
	"time"
)

type FetcherConfig struct {
	Address             string
	Secret              string
	MaxParallelRequests int
	RequestTimeout      time.Duration
}

func GetFetcherConfig() FetcherConfig {
	return FetcherConfig{
		Address:             util.GetEnv("FETCHER_ADDRESS", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest"),
		Secret:              util.GetEnv("FETCHER_SECRET", "5c16e994-9da9-41bd-8638-704e76275043"),
		MaxParallelRequests: util.IntEnv("MAX_PARALLEL_FETCH_REQUESTS", 1),
		RequestTimeout:      time.Second * time.Duration(util.IntEnv("FETCHER_TIMEOUT_SECONDS", 5)),
	}
}
