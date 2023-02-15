package main

import (
	"blox/fetcher"
	"blox/service"
)

func main() {
	s := service.New(fetcher.New())
	s.Start()
}
