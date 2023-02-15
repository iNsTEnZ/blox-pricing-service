package service

import (
	"blox/cache"
	"blox/config"
	"blox/fetcher"
	"blox/model"
	pb "blox/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type Service struct {
	pb.UnimplementedCryptoPricingServer
	config   config.ServiceConfig
	cache    cache.ICache
	ticker   *time.Ticker
	fetcher  fetcher.IFetcher
	quitChan chan struct{}
}

func New(fetcher fetcher.IFetcher) *Service {
	cfg := config.GetServiceConfig()

	return &Service{
		config:   cfg,
		cache:    cache.NewLRUCache(cfg.CacheSizeLimit),
		ticker:   time.NewTicker(cfg.FetchInterval),
		fetcher:  fetcher,
		quitChan: make(chan struct{}),
	}
}

func (s *Service) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))

	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", s.config.Port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCryptoPricingServer(grpcServer, s)
	log.Printf("server listening at %v", lis.Addr())

	go func() {
		s.timedFetch()
		log.Printf("starting ticker for price fetching")

		for {
			select {
			case <-s.ticker.C:
				s.timedFetch()
			case <-s.quitChan:
				s.ticker.Stop()
				grpcServer.GracefulStop()
				return
			}
		}
	}()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Service) Stop() {
	close(s.quitChan)
}

func (s *Service) GetPrice(ctx context.Context, request *pb.PriceRequest) (*pb.PriceResponse, error) {
	response := &pb.PriceResponse{}
	currency := request.Currency
	symbols := request.Symbols
	var unavailableSym []string

	for _, symbol := range symbols {
		if cached, exist := s.cache.Get(fmt.Sprintf("%s_%s", currency, symbol)); exist {
			data := cached.(model.Data)
			response.Symbols = append(response.Symbols, &pb.Symbol{
				Name:     symbol,
				Price:    data.Quote[currency].Price,
				Currency: currency,
			})
		} else {
			unavailableSym = append(unavailableSym, symbol)
		}
	}

	if len(unavailableSym) == 0 {
		return response, nil
	}

	symbols = unavailableSym

	fetchResp, err := s.fetcher.Fetch(ctx, currency, symbols)

	if err != nil {
		return nil, err
	}

	for symbol, data := range fetchResp.Data {
		response.Symbols = append(response.Symbols, &pb.Symbol{
			Name:     symbol,
			Price:    data.Quote[currency].Price,
			Currency: currency,
		})
	}

	return response, nil
}

func (s *Service) timedFetch() {
	fetchResp, err := s.fetcher.Fetch(context.Background(), s.config.FetchIntervalCurrency, s.config.FetchIntervalSymbols)

	if err != nil {
		log.Printf("failed to fetch data: %v", err)
		return
	}

	for _, data := range fetchResp.Data {
		if key, err := data.CacheKey(s.config.FetchIntervalCurrency); err == nil {
			log.Printf("saving to cache: %s %f %s",
				data.Symbol,
				data.Quote[s.config.FetchIntervalCurrency].Price,
				s.config.FetchIntervalCurrency)
			s.cache.Set(key, data)
		}
	}
}
