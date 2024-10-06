package pricesServices

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PriceDB struct {
	Price       float64 `db:"price"`
	Price24hAgo float64 `db:"price_usd_24h_ago"`
}

type CoinPriceDB struct {
	PriceDB
	Coin string `db:"blockchain_coin"`
}

type MarketPriceDB struct {
	PriceDB
	MarketTicker string `db:"market_ticker"`
}

type BlockchainCoinPrice struct {
	PriceDB
	Markets []MarketPriceDB
}

type PricesService struct {
	pgConn *sqlx.DB
}

func (s *PricesService) GetCoinPrices(ctx context.Context) ([]CoinPriceDB, error) {
	coinPrices := []CoinPriceDB{}
	err := s.pgConn.SelectContext(ctx, &coinPrices, `SELECT blockchain_coin, price, price_usd_24h_ago
		FROM coin_prices WHERE usdt = true`)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get coin prices: %w", err)
	}

	return coinPrices, nil
}

func (s *PricesService) GetBlockchainCoinPrice(ctx context.Context, coin string) (*BlockchainCoinPrice, error) {
	marketPrices := []MarketPriceDB{}
	err := s.pgConn.SelectContext(ctx, &marketPrices, `SELECT market_ticker, price, price_usd_24h_ago
		FROM coin_prices WHERE blockchain_coin = $1 ORDER BY usdt DESC`)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get blockchain (coin: %s) price: %w", coin, err)
	}

	coinMarkets := marketPrices[1:]

	return &BlockchainCoinPrice{
		PriceDB: PriceDB{
			Price:       marketPrices[0].Price,
			Price24hAgo: marketPrices[0].Price24hAgo,
		},
		Markets: coinMarkets,
	}, nil
}
