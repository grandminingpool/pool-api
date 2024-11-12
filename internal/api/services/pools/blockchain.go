package poolsServices

import (
	"context"
	"fmt"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Pool struct {
	Info   *poolProto.PoolInfo
	Stats  *poolProto.PoolStats
	Slaves []*poolProto.PoolSlave
}

type BlockchainService struct{}

func (s *BlockchainService) GetPool(ctx context.Context, blockchain *blockchains.Blockchain) (*Pool, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	pool := &Pool{
		Info:   nil,
		Stats:  nil,
		Slaves: nil,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		poolInfo, err := client.GetPoolInfo(gCtx, &emptypb.Empty{})
		if err == nil {
			pool.Info = poolInfo

			return nil
		}

		return fmt.Errorf("failed to get pool info: %w", err)
	})
	g.Go(func() error {
		poolStats, err := client.GetPoolStats(gCtx, &emptypb.Empty{})
		if err == nil {
			pool.Stats = poolStats

			return nil
		}

		return fmt.Errorf("failed to get pool stats: %w", err)
	})
	g.Go(func() error {
		poolSlaves, err := client.GetPoolSlaves(gCtx, &emptypb.Empty{})
		if err == nil {
			pool.Slaves = poolSlaves.Slaves

			return nil
		}

		return fmt.Errorf("failed to get pool slaves: %w", err)
	})

	return pool, g.Wait()
}

func (s *BlockchainService) GetPoolInfo(ctx context.Context, blockchain *blockchains.Blockchain) (*poolProto.PoolInfo, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	poolInfo, err := client.GetPoolInfo(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool (blockchain: %s) info: %w", blockchain.GetInfo().Blockchain, err)
	}

	return poolInfo, nil
}

func (s *BlockchainService) GetPoolStats(ctx context.Context, blockchain *blockchains.Blockchain) (*poolProto.PoolStats, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	poolStats, err := client.GetPoolStats(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool (blockchain: %s) stats: %w", blockchain.GetInfo().Blockchain, err)
	}

	return poolStats, nil
}

func (s *BlockchainService) GetPoolSlaves(ctx context.Context, blockchain *blockchains.Blockchain) ([]*poolProto.PoolSlave, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	poolSlaves, err := client.GetPoolSlaves(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool (blockchain: %s) slaves: %w", blockchain.GetInfo().Blockchain, err)
	}

	return poolSlaves.Slaves, nil
}
