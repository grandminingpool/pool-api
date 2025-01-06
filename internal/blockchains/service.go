package blockchains

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	appConfig "github.com/grandminingpool/pool-api/configs/app"
	poolAPIClient "github.com/grandminingpool/pool-api/internal/clients/pool_api"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type PoolAPIDB struct {
	Address string `db:"pool_api_address"`
	TLSCA   string `db:"pool_api_tls_ca"`
}

type BlockchainDB struct {
	Blockchain string `db:"blockchain"`
	Name       string `db:"name"`
	Ticker     string `db:"ticker"`
	AtomicUnit uint16 `db:"atomic_unit"`
	RowOrder   int    `db:"row_order"`
	PoolAPIDB
}

type BlockchainInfo struct {
	Blockchain string
	Name       string
	Ticker     string
	AtomicUnit uint16
}

type Blockchain struct {
	info BlockchainInfo
	conn *grpc.ClientConn
}

func (b *Blockchain) GetInfo() BlockchainInfo {
	return b.info
}

func (b *Blockchain) GetConnection() *grpc.ClientConn {
	return b.conn
}

type Service struct {
	pgConn         *sqlx.DB
	blockchainsMap map[string]Blockchain
	blockchains    []BlockchainInfo
	config         *appConfig.PoolAPIConfig
}

func (s *Service) getBlockchainsFromDB(ctx context.Context) ([]BlockchainDB, error) {
	blockchains := []BlockchainDB{}
	if err := s.pgConn.SelectContext(ctx, &blockchains, "SELECT * FROM blockchains ORDER BY row_order"); err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to query blockchains: %w", err)
	}

	return blockchains, nil
}

func (s *Service) GetBlockchains() []Blockchain {
	blockchains := make([]Blockchain, 0, len(s.blockchains))
	for _, b := range s.blockchains {
		blockchain, ok := s.blockchainsMap[b.Blockchain]
		if ok {
			blockchains = append(blockchains, blockchain)
		}
	}

	return blockchains
}

func (s *Service) GetBlockchainsInfos() []BlockchainInfo {
	return s.blockchains
}

func (s *Service) GetBlockchain(blockchain string) (*Blockchain, error) {
	b, ok := s.blockchainsMap[blockchain]
	if !ok {
		return nil, fmt.Errorf("failed to get blockchain: %s, error: not found", blockchain)
	}

	return &b, nil
}

func (s *Service) Start(ctx context.Context, certsPath string) error {
	blockchains, err := s.getBlockchainsFromDB(ctx)
	if err != nil {
		return err
	}

	for _, b := range blockchains {
		conn, err := poolAPIClient.NewClient(
			b.PoolAPIDB.Address,
			certsPath,
			b.PoolAPIDB.TLSCA,
			time.Duration(s.config.RequestTimeout)*time.Second,
		)
		if err != nil {
			s.Close()

			return fmt.Errorf("failed to create blockchain pool api client (blockchain: %s), error: %w", b.Blockchain, err)
		}

		blockchainInfo := BlockchainInfo{
			Blockchain: b.Blockchain,
			Name:       b.Name,
			Ticker:     b.Ticker,
			AtomicUnit: b.AtomicUnit,
		}
		s.blockchainsMap[b.Blockchain] = Blockchain{
			info: blockchainInfo,
			conn: conn,
		}
		s.blockchains = append(s.blockchains, blockchainInfo)
	}

	return nil
}

func (s *Service) Close() {
	for _, b := range s.blockchainsMap {
		b.conn.Close()
	}

	clear(s.blockchains)
	s.blockchains = nil
}

func NewService(pgConn *sqlx.DB, config *appConfig.PoolAPIConfig) *Service {
	return &Service{
		pgConn:         pgConn,
		blockchainsMap: make(map[string]Blockchain),
		blockchains:    []BlockchainInfo{},
		config:         config,
	}
}
