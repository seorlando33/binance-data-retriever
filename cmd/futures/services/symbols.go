package services

import (
	"context"
	"fmt"

	binance "github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)

type SymbolService interface {
	GetSymbols(...string) ([]futures.Symbol, error)
}

type symbolService struct{}

func NewSymbolService() SymbolService {
	return &symbolService{}
}

func (s *symbolService) GetSymbols(symbolNames ...string) ([]futures.Symbol, error) {

	exchangeInfo, err := binance.NewFuturesClient("", "").NewExchangeInfoService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	symbols, missing := filter(exchangeInfo.Symbols, symbolNames...)
	if missing != nil {
		return nil, fmt.Errorf("the symbols: %v, doesn't exist", missing)
	}

	return symbols, nil
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
