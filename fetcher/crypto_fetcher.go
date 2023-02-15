package fetcher

import (
	"blox/config"
	"blox/model"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type CryptoFetcher struct {
	httpClient  *http.Client
	config      config.FetcherConfig
	rateLimiter chan struct{}
}

func New() *CryptoFetcher {
	cfg := config.GetFetcherConfig()

	return &CryptoFetcher{
		httpClient: &http.Client{
			Timeout: cfg.RequestTimeout,
		},
		config:      cfg,
		rateLimiter: make(chan struct{}, cfg.MaxParallelRequests),
	}
}

func (f *CryptoFetcher) Fetch(ctx context.Context, currency string, symbols []string) (*model.FetcherResponse, error) {
	f.rateLimiter <- struct{}{}
	defer func() { <-f.rateLimiter }()
	log.Printf("fetching data for currency %s and symbols %v", currency, symbols)
	req, err := http.NewRequestWithContext(ctx, "GET", f.config.Address, nil)

	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("symbol", strings.Join(symbols, ","))
	q.Add("convert", currency)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", f.config.Secret)
	req.URL.RawQuery = q.Encode()

	resp, err := f.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	result := &model.FetcherResponse{}

	if err := json.Unmarshal(respBody, result); err != nil {
		return nil, err
	}

	return result, nil
}
