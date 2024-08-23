package poolsServices

import (
	"context"
	"fmt"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Pool struct {
	Info   *poolProto.PoolInfo
	Stats  *poolProto.PoolStats
	Slaves []*poolProto.PoolSlave
}

type BlockchainService struct{}

func (s *BlockchainService) getInfo(
	ctx context.Context,
	client poolProto.PoolServiceClient,
	poolInfoCh chan<- *poolProto.PoolInfo,
	errCh chan<- error,
) {
	select {
	case <-ctx.Done():
		return
	default:
		poolInfo, err := client.GetPoolInfo(ctx, &emptypb.Empty{})
		if err != nil {
			errCh <- fmt.Errorf("failed to get pool info: %w", err)

			return
		}

		poolInfoCh <- poolInfo
	}
}

func (s *BlockchainService) getStats(
	ctx context.Context,
	client poolProto.PoolServiceClient,
	poolStatsCh chan<- *poolProto.PoolStats,
	errCh chan<- error,
) {
	select {
	case <-ctx.Done():
		return
	default:
		poolStats, err := client.GetPoolStats(ctx, &emptypb.Empty{})
		if err != nil {
			errCh <- fmt.Errorf("failed to get pool stats: %w", err)

			return
		}

		poolStatsCh <- poolStats
	}
}

func (s *BlockchainService) getSlaves(
	ctx context.Context,
	client poolProto.PoolServiceClient,
	poolSlavesCh chan<- []*poolProto.PoolSlave,
	errCh chan<- error,
) {
	select {
	case <-ctx.Done():
		return
	default:
		poolSlaves, err := client.GetPoolSlaves(ctx, &emptypb.Empty{})
		if err != nil {
			errCh <- fmt.Errorf("failed to get pool slaves: %w", err)

			return
		}

		poolSlavesCh <- poolSlaves.Slaves
	}
}

func (s *BlockchainService) GetAll(ctx context.Context, blockchain *blockchains.Blockchain) (*Pool, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	pool := &Pool{
		Info:   nil,
		Stats:  nil,
		Slaves: nil,
	}

	poolInfoCh := make(chan *poolProto.PoolInfo, 1)
	poolStatsCh := make(chan *poolProto.PoolStats, 1)
	poolSlaves := make(chan []*poolProto.PoolSlave, 1)
	errCh := make(chan error, 3)
	defer close(poolInfoCh)
	defer close(poolStatsCh)
	defer close(poolSlaves)
	defer close(errCh)

	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go s.getInfo(newCtx, client, poolInfoCh, errCh)
	go s.getStats(newCtx, client, poolStatsCh, errCh)
	go s.getSlaves(newCtx, client, poolSlaves, errCh)

	for i := 0; i < 3; i++ {
		select {
		case err := <-errCh:
			return nil, fmt.Errorf("failed to get pool (blockchain: %s) data: %w", blockchain.GetInfo().Coin, err)
		case poolInfo := <-poolInfoCh:
			pool.Info = poolInfo
		case poolStats := <-poolStatsCh:
			pool.Stats = poolStats
		case poolSlaves := <-poolSlaves:
			pool.Slaves = poolSlaves
		}
	}

	return pool, nil
}

func (s *BlockchainService) GetInfo(ctx context.Context, blockchain *blockchains.Blockchain) (*poolProto.PoolInfo, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	poolInfo, err := client.GetPoolInfo(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool (blockchain: %s) info: %w", blockchain.GetInfo().Coin, err)
	}

	return poolInfo, nil
}

func (s *BlockchainService) GetStats(ctx context.Context, blockchain *blockchains.Blockchain) (*poolProto.PoolStats, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	poolStats, err := client.GetPoolStats(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool (blockchain: %s) stats: %w", blockchain.GetInfo().Coin, err)
	}

	return poolStats, nil
}

func (s *BlockchainService) GetSlaves(ctx context.Context, blockchain *blockchains.Blockchain) ([]*poolProto.PoolSlave, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	poolSlaves, err := client.GetPoolSlaves(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool (blockchain: %s) slaves: %w", blockchain.GetInfo().Coin, err)
	}

	return poolSlaves.Slaves, nil
}
