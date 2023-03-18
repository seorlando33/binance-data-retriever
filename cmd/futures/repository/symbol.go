package repository

import (
	"database/sql"

	binance "github.com/adshao/go-binance/v2/futures"
)

type SymbolRepository interface {
	InsertSymbol([]binance.Symbol) error
}

type postgresSymbol struct {
	db *sql.DB
}

func NewPostgresSymbolRepository(db *sql.DB) SymbolRepository {

	return &postgresSymbol{db: db}
}

func (p *postgresSymbol) InsertSymbol(symbols []binance.Symbol) (err error) {

	tx, err := p.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	for _, s := range symbols {

		_, err = tx.Exec(
			insertSymbolQuery,
			s.Symbol,
			s.Pair,
			s.ContractType,
			s.DeliveryDate,
			s.OnboardDate,
			s.Status,
			s.BaseAsset,
			s.QuoteAsset,
			s.MarginAsset,
			s.PricePrecision,
			s.QuantityPrecision,
			s.BaseAssetPrecision,
			s.QuotePrecision,
			s.TriggerProtect,
			s.LiquidationFee,
		)
		if err != nil {
			return
		}

	}

	return
}

const insertSymbolQuery string = `
INSERT INTO 

	futures.symbol (
		sym_symbol,
		sym_pair,
		sym_contract_type,
		sym_delivery_date,
		sym_onboard_date,
		sym_status,
		sym_base_asset,
		sym_quote_asset,
		sym_margin_asset,
		sym_price_precision,
		sym_quantity_precision,
		sym_base_asset_precision,
		sym_quote_precision,
		sym_trigger_protect,
		sym_liquidation_fee
	) 
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10,
		$11,
		$12,
		$13,
		$14,
		$15
	)
`
