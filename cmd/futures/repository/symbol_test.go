package repository

import (
	"context"
	"database/sql"
	"testing"

	binance "github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/seorlando33/binance-data-retriever/cmd/futures/db"
)

func TestInsertSymbols(t *testing.T) {

	client := binance.NewFuturesClient("", "")
	exchangeInfo, err := client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		t.Fatal("Error getting the exchange info")
	}

	conn := db.GetDBConnection()

	tests := []struct {
		name    string
		args    []futures.Symbol
		db      *sql.DB
		wantErr bool
	}{
		{
			name:    "success insert",
			args:    exchangeInfo.Symbols,
			db:      conn,
			wantErr: false,
		},
		{
			name: "fail insert incomplete payload",
			args: []futures.Symbol{
				{
					Symbol: "BTCUSDT",
				},
			},
			db:      conn,
			wantErr: true,
		},
		{
			name:    "fail insert symbols should be unique",
			args:    exchangeInfo.Symbols,
			db:      conn,
			wantErr: true,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			err := NewPostgresSymbolRepository(tt.db).InsertSymbol(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("test: %s, failed", tt.name)
			}
		})

	}

	_, err = conn.Exec(deleteSymbols)
	if err != nil {
		t.Error("Error deleting the symbols")
	}
}

const deleteSymbols = `DELETE FROM futures.symbol`
