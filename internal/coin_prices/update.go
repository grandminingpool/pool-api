package coinPrices

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const apiURL = "https://tradeogre.com/api/v1"

type GetTickerResponse struct {
	Success      bool   `json:"success"`
	Price        string `json:"price"`
	InitialPrice string `json:"initialprice"`
}

type CoinPrice struct {
	MarketTicker string
	Price        float64
	Price24hAgo  float64
}

type UpdateService struct {
	pgConn *sqlx.DB
}

func (s *UpdateService) getMarketTickers(ctx context.Context) ([]string, error) {
	marketTickers := []string{}
	if err := s.pgConn.SelectContext(ctx, &marketTickers, `SELECT market_ticker FROM coin_prices`); err != nil {
		return nil, fmt.Errorf("failed to get market tickers: %w", err)
	}

	return marketTickers, nil
}

func (s *UpdateService) parseCoinPriceFromTickerResponse(response *GetTickerResponse) (*CoinPrice, error) {
	if !response.Success {
		return nil, fmt.Errorf("unsuccessful status in ticker response")
	}

	price, err := strconv.ParseFloat(response.Price, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse price from ticker response: %w", err)
	}

	initialPrice, err := strconv.ParseFloat(response.InitialPrice, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse initial price from ticker response: %w", err)
	}

	return &CoinPrice{
		Price:       price,
		Price24hAgo: initialPrice,
	}, nil
}

func (s *UpdateService) getTickerPrice(
	ctx context.Context,
	ticker string,
	coinPricesCh chan<- *CoinPrice,
	errCh chan<- error,
) {
	select {
	case <-ctx.Done():
		return
	default:
		resp, err := http.Get(fmt.Sprintf("%s/ticker/%s", apiURL, ticker))
		if err != nil {
			errCh <- fmt.Errorf("failed to make request to api to get price for ticker: %s, error: %w", ticker, err)

			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errCh <- fmt.Errorf("failed to read body from in response for ticker: %s, error: %w", ticker, err)

			return
		}

		var response GetTickerResponse
		if err := json.Unmarshal(body, &response); err != nil {
			errCh <- fmt.Errorf("failed to decode response for ticker: %s, error %w", ticker, err)

			return
		}

		coinPrice, err := s.parseCoinPriceFromTickerResponse(&response)
		if err != nil {
			errCh <- fmt.Errorf("failed to parse coin price from response for ticker: %s", ticker)

			return
		}

		coinPricesCh <- coinPrice
	}
}

func (s *UpdateService) updateRowsInDB(ctx context.Context, tx *sqlx.Tx, newPrices []*CoinPrice) error {
	for _, coinPrice := range newPrices {
		if _, err := tx.ExecContext(ctx, `UPDATE coin_prices SET price = ? WHERE market_ticker = ?`,
			coinPrice.Price,
			coinPrice.MarketTicker,
		); err != nil {
			return fmt.Errorf("failed to update price (value: %f) for ticker: %s, error: %w",
				coinPrice.Price,
				coinPrice.MarketTicker,
				err,
			)
		}
	}

	return nil
}

func (s *UpdateService) Update(ctx context.Context) {
	marketTickers, err := s.getMarketTickers(ctx)
	if err != nil {
		zap.L().Fatal("failed to get market tickers", zap.Error(err))

		return
	}

	errCh := make(chan error, len(marketTickers))
	coinPricesCh := make(chan *CoinPrice, len(marketTickers))
	defer close(errCh)
	defer close(coinPricesCh)

	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	for _, mt := range marketTickers {
		go s.getTickerPrice(newCtx, mt, coinPricesCh, errCh)
	}

	newPrices := make([]*CoinPrice, 0, len(marketTickers))
	for i := 0; i < len(marketTickers); i++ {
		select {
		case err := <-errCh:
			zap.L().Fatal("failed to get coin price from api", zap.Error(err))

			return
		case price := <-coinPricesCh:
			newPrices = append(newPrices, price)
		}
	}

	tx, err := s.pgConn.BeginTxx(newCtx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		zap.L().Fatal("failed to create transaction for update prices", zap.Error(err))

		return
	}

	if err := s.updateRowsInDB(newCtx, tx, newPrices); err != nil {
		zap.L().Fatal("failed to update database rows", zap.Error(err))

		return
	}

	if err := tx.Commit(); err != nil {
		zap.L().Fatal("failed to commit transaction with updated prices", zap.Error(err))
	}
}