CREATE SCHEMA IF NOT EXISTS futures;

CREATE  TABLE futures.symbol ( 
	sym_id               integer DEFAULT nextval('futures.symbol_sym_id_seq'::regclass) NOT NULL  ,
	sym_symbol           varchar(20)  NOT NULL  ,
	sym_pair             varchar(20)  NOT NULL  ,
	sym_contract_type    varchar(20)  NOT NULL  ,
	sym_delivery_date    bigint  NOT NULL  ,
	sym_onboard_date     bigint  NOT NULL  ,
	sym_status           varchar(20)  NOT NULL  ,
	sym_base_asset       varchar(12)  NOT NULL  ,
	sym_quote_asset      varchar(12)  NOT NULL  ,
	sym_margin_asset     varchar(12)  NOT NULL  ,
	sym_price_precision  integer  NOT NULL  ,
	sym_quantity_precision integer  NOT NULL  ,
	sym_base_asset_precision integer  NOT NULL  ,
	sym_quote_precision  integer  NOT NULL  ,
	sym_trigger_protect  varchar(12)  NOT NULL  ,
	sym_liquidation_fee  varchar(12)  NOT NULL  ,
	CONSTRAINT pk_symbol PRIMARY KEY ( sym_id ),
	CONSTRAINT unq_symbol UNIQUE ( sym_symbol ) 
 );

CREATE  TABLE futures.taker_buy_sell_volume ( 
	bsv_id               integer DEFAULT nextval('futures.taker_buy_sell_volume_bsv_id_seq'::regclass) NOT NULL  ,
	sym_id               integer  NOT NULL  ,
	bsv_buy_sell_ratio   numeric(18,4)  NOT NULL  ,
	bsv_buy_volume       numeric(18,4)  NOT NULL  ,
	bsv_sell_volume      numeric(18,4)  NOT NULL  ,
	bsv_timestamp        bigint  NOT NULL  ,
	bsv_create_date      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_buy_sell_ratio PRIMARY KEY ( bsv_id )
 );

CREATE INDEX idx_taker_buy_sell_volume_0 ON futures.taker_buy_sell_volume  ( sym_id );

CREATE INDEX idx_taker_buy_sell_volume ON futures.taker_buy_sell_volume  ( bsv_timestamp );

CREATE  TABLE futures.top_trader_long_short_ratio_account ( 
	lsra_id              integer DEFAULT nextval('futures.top_trader_long_short_ratio_account_lsra_id_seq'::regclass) NOT NULL  ,
	sym_id               integer  NOT NULL  ,
	lsra_long_short_ratio numeric(6,4)  NOT NULL  ,
	lsra_long_account    numeric(6,4)  NOT NULL  ,
	lsra_short_account   numeric(6,4)  NOT NULL  ,
	lsra_timestamp       bigint  NOT NULL  ,
	lsra_create_date     timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_top_trader_long_short_ratio_account PRIMARY KEY ( lsra_id )
 );

CREATE INDEX idx_top_trader_long_short_ratio_account_0 ON futures.top_trader_long_short_ratio_account  ( sym_id );

CREATE INDEX idx_top_trader_long_short_ratio_account ON futures.top_trader_long_short_ratio_account  ( lsra_timestamp );

CREATE  TABLE futures.top_trader_long_short_ratio_position ( 
	lsrp_id              integer DEFAULT nextval('futures.top_trader_long_short_ratio_position_lsrp_id_seq'::regclass) NOT NULL  ,
	sym_id               integer  NOT NULL  ,
	lsrp_long_short_ratio numeric(6,4)  NOT NULL  ,
	lsrp_long_account    numeric(6,4)  NOT NULL  ,
	lsrp_short_account   numeric(6,4)  NOT NULL  ,
	lsrp_timestamp       bigint  NOT NULL  ,
	lsrp_create_date     timestamp DEFAULT CURRENT_TIMESTAMP   ,
	CONSTRAINT pk_top_trader_long_short_ratio_position PRIMARY KEY ( lsrp_id )
 );

CREATE INDEX idx_top_trader_long_short_ratio_position_0 ON futures.top_trader_long_short_ratio_position  ( sym_id );

CREATE INDEX idx_top_trader_long_short_ratio_position ON futures.top_trader_long_short_ratio_position  ( lsrp_timestamp );

CREATE  TABLE futures.kline ( 
	kld_id               integer DEFAULT nextval('futures.kline_kld_id_seq1'::regclass) NOT NULL  ,
	sym_id               integer  NOT NULL  ,
	kld_open_time        bigint  NOT NULL  ,
	kld_open             numeric(36,8)  NOT NULL  ,
	kld_high             numeric(36,8)  NOT NULL  ,
	kld_low              numeric(36,8)  NOT NULL  ,
	kld_close            numeric(36,8)  NOT NULL  ,
	kld_volume           numeric(36,8)  NOT NULL  ,
	kld_close_time       bigint  NOT NULL  ,
	kld_quote_asset_volume numeric(36,8)  NOT NULL  ,
	kld_number_of_trades integer  NOT NULL  ,
	kld_taker_buy_base_asset_volume numeric(36,8)  NOT NULL  ,
	kld_taker_buy_quote_asset_volume numeric(36,8)  NOT NULL  ,
	kld_create_date      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_kline PRIMARY KEY ( kld_id )
 );

CREATE INDEX idx_kline_0 ON futures.kline  ( sym_id );

CREATE INDEX idx_kline ON futures.kline  ( kld_open_time, kld_close_time );

CREATE  TABLE futures.long_short_ratio ( 
	lsr_id               integer DEFAULT nextval('futures.long_short_ratio_lsr_id_seq'::regclass) NOT NULL  ,
	sym_id               integer  NOT NULL  ,
	lsr_long_short_ratio numeric(6,4)  NOT NULL  ,
	lsr_long_account     numeric(6,4)  NOT NULL  ,
	lsr_short_account    numeric(6,4)  NOT NULL  ,
	lsr_timestamp        bigint  NOT NULL  ,
	lsr_create_date      timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_long_short_ratio PRIMARY KEY ( lsr_id )
 );

CREATE INDEX idx_long_short_ratio_0 ON futures.long_short_ratio  ( sym_id );

CREATE INDEX idx_long_short_ratio ON futures.long_short_ratio  ( lsr_timestamp );

CREATE  TABLE futures.open_interest ( 
	oi_id                integer DEFAULT nextval('futures.open_interest_oi_id_seq1'::regclass) NOT NULL  ,
	sym_id               integer  NOT NULL  ,
	oi_sum_open_interest numeric(36,8)  NOT NULL  ,
	oi_sum_open_interest_value numeric(36,8)  NOT NULL  ,
	oi_timestamp         bigint  NOT NULL  ,
	oi_create_date       timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL  ,
	CONSTRAINT pk_open_interest PRIMARY KEY ( oi_id )
 );

CREATE INDEX idx_open_interest_0 ON futures.open_interest  ( sym_id );

CREATE INDEX idx_open_interest ON futures.open_interest  ( oi_timestamp );

ALTER TABLE futures.kline ADD CONSTRAINT fk_kline_symbol FOREIGN KEY ( sym_id ) REFERENCES futures.symbol( sym_id );

ALTER TABLE futures.long_short_ratio ADD CONSTRAINT fk_long_short_ratio_symbol FOREIGN KEY ( sym_id ) REFERENCES futures.symbol( sym_id );

ALTER TABLE futures.open_interest ADD CONSTRAINT fk_open_interest_symbol FOREIGN KEY ( sym_id ) REFERENCES futures.symbol( sym_id );

ALTER TABLE futures.taker_buy_sell_volume ADD CONSTRAINT fk_taker_buy_sell_volume FOREIGN KEY ( sym_id ) REFERENCES futures.symbol( sym_id );

ALTER TABLE futures.top_trader_long_short_ratio_account ADD CONSTRAINT fk_top_trader_long_short_ratio_account FOREIGN KEY ( sym_id ) REFERENCES futures.symbol( sym_id );

ALTER TABLE futures.top_trader_long_short_ratio_position ADD CONSTRAINT fk_top_trader_long_short_ratio_position FOREIGN KEY ( sym_id ) REFERENCES futures.symbol( sym_id );

