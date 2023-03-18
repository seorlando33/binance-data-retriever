package services

import (
	"reflect"
	"testing"

	"github.com/adshao/go-binance/v2/futures"
)

func TestFilter(t *testing.T) {
	symbols := []futures.Symbol{
		{Symbol: "BTCUSDT"},
		{Symbol: "ETHUSDT"},
		{Symbol: "BNBUSDT"},
	}

	cases := []struct {
		name        string
		symbolNames []string
		expectedRes []futures.Symbol
		expectedMis []string
	}{
		{
			name:        "no missing",
			symbolNames: []string{"BTCUSDT", "BNBUSDT"},
			expectedRes: []futures.Symbol{{Symbol: "BTCUSDT"}, {Symbol: "BNBUSDT"}},
			expectedMis: nil,
		},
		{
			name:        "some missing",
			symbolNames: []string{"ETHUSDT", "ADAUSDT", "BNBUSDT"},
			expectedRes: []futures.Symbol{{Symbol: "ETHUSDT"}, {Symbol: "BNBUSDT"}},
			expectedMis: []string{"ADAUSDT"},
		},
		{
			name:        "all missing",
			symbolNames: []string{"ADAUSDT", "ETCUSDT", "ALGOUSDT"},
			expectedRes: []futures.Symbol{},
			expectedMis: []string{"ADAUSDT", "ETCUSDT", "ALGOUSDT"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, mis := filter(symbols, tc.symbolNames...)

			if !reflect.DeepEqual(res, tc.expectedRes) {
				t.Errorf("expected %v but got %v", tc.expectedRes, res)
			}

			if !reflect.DeepEqual(mis, tc.expectedMis) {
				t.Errorf("expected %v but got %v", tc.expectedMis, mis)
			}
		})
	}
}
