package fetcher

import (
	"blox/model"
	"context"
)

type IFetcher interface {
	Fetch(ctx context.Context, currency string, symbols []string) (*model.FetcherResponse, error)
}
