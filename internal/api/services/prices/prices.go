package pricesServices

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PriceDB struct {
	Price       float64 `db:"price"`
	Price24hAgo float64 `db:"price_24h_ago"`
}

type BlockchainPriceDB struct {
	PriceDB
	Blockchain string `db:"blockchain"`
}

type MarketPriceDB struct {
	PriceDB
	MarketTicker string `db:"market_ticker"`
}

type BlockchainMarkets struct {
	PriceDB
	Markets []MarketPriceDB
}

type PricesService struct {
	pgConn *sqlx.DB
}

func (s *PricesService) GetPrices(ctx context.Context) ([]BlockchainPriceDB, error) {
	prices := []BlockchainPriceDB{}
	err := s.pgConn.SelectContext(ctx, &prices, `SELECT blockchain, price, price_24h_ago FROM prices WHERE usdt = true`)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get blockchains prices: %w", err)
	}

	return prices, nil
}

func (s *PricesService) GetBlockchainMarkets(ctx context.Context, blockchain string) (*BlockchainMarkets, error) {
	marketPrices := []MarketPriceDB{}
	err := s.pgConn.SelectContext(ctx, &marketPrices, `SELECT market_ticker, price, price_24h_ago
		FROM prices WHERE blockchain = $1 ORDER BY usdt DESC`, blockchain)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get markets (blockchain: %s), error: %w", blockchain, err)
	}

	markets := marketPrices[1:]

	return &BlockchainMarkets{
		PriceDB: PriceDB{
			Price:       marketPrices[0].Price,
			Price24hAgo: marketPrices[0].Price24hAgo,
		},
		Markets: markets,
	}, nil
}

func NewPricesService(pgConn *sqlx.DB) *PricesService {
	return &PricesService{
		pgConn: pgConn,
	}
}
