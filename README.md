# Price Service
How to use
```
go build
```
```
go run blox
```

Environment variables
```
FETCHER_ADDRESS (default: https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest) - the address of the pricing API.
FETCHER_SECRET (default: "") - the secret for the pricing API. You must set this variable.
MAX_PARALLEL_FETCH_REQUESTS (default: 1) - the maximum allowed concurrent requests to the pricing API.
FETCHER_TIMEOUT_SECONDS (default: 5) - defines, in seconds, the timeout for a single request to the pricing API.
SERVICE_PORT (default: 8080) - defines the port on which the server would listen to requests.
FETCH_INTERVAL_SECONDS (default: 60) - defines, in seconds, the fetch interval from the pricing interval for specified symbols.
CACHE_SIZE (default: 20) - defines the maximum allowed symbols to be stored in memory.
FETCH_INTERVAL_CURRENCY (default: USD) - defines the default currency to fetch the prices with. 
FETCH_INTERVAL_SYMBOLS (default: BTC,ETH,BNB) - defines the symbol to fetch data for in a specified interval and save in memory.
```