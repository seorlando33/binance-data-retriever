package services

import (
	"context"
	"fmt"

	binance "github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/seorlando33/binance-data-retriever/cmd/futures/repository"
)

type SymbolService interface {
	InsertSymbol(...string) error
}

type symbolService struct {
	s repository.SymbolRepository
}

func NewSymbolService(sr repository.SymbolRepository) SymbolService {
	return &symbolService{s: sr}
}

func (s *symbolService) InsertSymbol(symbolNames ...string) error {

	exchangeInfo, err := binance.NewFuturesClient("", "").NewExchangeInfoService().Do(context.Background())
	if err != nil {
		return err
	}

	symbols, missing := filter(exchangeInfo.Symbols, symbolNames...)
	if missing != nil {
		return fmt.Errorf("the symbols: %v, doesn't exist", missing)
	}

	return s.s.InsertSymbol(symbols)
}

func filter(symbols []futures.Symbol, symbolNames ...string) ([]futures.Symbol, []string) {

	if symbolNames == nil {
		return symbols, nil
	}

	symbolsMap := make(map[string]futures.Symbol, len(symbols))
	result := make([]futures.Symbol, 0, len(symbolNames))
	var missing []string

	for _, sym := range symbols {
		symbolsMap[sym.Symbol] = sym
	}

	for _, name := range symbolNames {
		if sym, ok := symbolsMap[name]; ok {
			result = append(result, sym)
		} else {
			missing = append(missing, name)
		}
	}

	return result, missing
}
