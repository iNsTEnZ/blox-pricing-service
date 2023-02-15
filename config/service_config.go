package config

import (
	"blox/util"
	"time"
)

type ServiceConfig struct {
	Port                  int
	FetchInterval         time.Duration
	CacheSizeLimit        int
	FetchIntervalCurrency string
	FetchIntervalSymbols  []string
}

func GetServiceConfig() ServiceConfig {
	return ServiceConfig{
		Port:                  util.IntEnv("SERVICE_PORT", 8080),
		FetchInterval:         time.Second * time.Duration(util.IntEnv("FETCH_INTERVAL_SECONDS", 60)),
		CacheSizeLimit:        util.IntEnv("CACHE_SIZE", 20),
		FetchIntervalCurrency: util.GetEnv("FETCH_INTERVAL_CURRENCY", "USD"),
		FetchIntervalSymbols:  util.GetListEnv("FETCH_INTERVAL_SYMBOLS", "BTC,ETH,BNB"),
	}
}
