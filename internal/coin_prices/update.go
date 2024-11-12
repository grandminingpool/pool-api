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
	if err := s.pgConn.SelectContext(ctx, &marketTickers, `SELECT market_ticker FROM prices`); err != nil {
		return nil, fmt.Errorf("failed to get market tickers: %w", err)
	}

	return marketTickers, nil
}

func (s *UpdateService) parseCoinPriceFromTickerResponse(ticker string, response *GetTickerResponse) (*CoinPrice, error) {
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
		MarketTicker: ticker,
		Price:        price,
		Price24hAgo:  initialPrice,
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

		coinPrice, err := s.parseCoinPriceFromTickerResponse(ticker, &response)
		if err != nil {
			errCh <- fmt.Errorf("failed to parse coin price from response for ticker: %s", ticker)

			return
		}

		coinPricesCh <- coinPrice
	}
}

func (s *UpdateService) updateRowsInDB(ctx context.Context, tx *sqlx.Tx, newPrices []*CoinPrice) error {
	for _, coinPrice := range newPrices {
		if _, err := tx.ExecContext(ctx, `UPDATE prices SET price = $1, price_24h_ago = $2 WHERE market_ticker = $3`,
			coinPrice.Price,
			coinPrice.Price24hAgo,
			coinPrice.MarketTicker,
		); err != nil {
			return fmt.Errorf("failed to update price for ticker: %s, error: %w",
				coinPrice.MarketTicker,
				err,
			)
		}
	}

	return nil
}

func (s *UpdateService) Update(ctx context.Context) error {
	marketTickers, err := s.getMarketTickers(ctx)
	if err != nil {
		return err
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
			return err
		case price := <-coinPricesCh:
			newPrices = append(newPrices, price)
		}
	}

	tx, err := s.pgConn.BeginTxx(newCtx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return fmt.Errorf("failed to create transaction for update prices: %w", err)
	}

	if err := s.updateRowsInDB(newCtx, tx, newPrices); err != nil {
		tx.Rollback()

		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()

		return fmt.Errorf("failed to commit transaction with updated prices: %w", err)
	}

	return nil
}

func NewUpdateService(pgConn *sqlx.DB) *UpdateService {
	return &UpdateService{
		pgConn: pgConn,
	}
}
