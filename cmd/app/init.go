package app

import (
	"github.com/seorlando33/binance-data-retriever/cmd/futures/services"
)

type Futures struct {
	Symbol services.SymbolService
}

type Services struct {
	Futures
}

var Service Services

func Run() {
	Service.Futures.Symbol = services.NewSymbolService()
}
