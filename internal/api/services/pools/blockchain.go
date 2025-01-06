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
	Info        *poolProto.PoolInfo
	Stats       *poolProto.PoolStats
	NetworkInfo *poolProto.NetworkInfo
	Slaves      []*poolProto.PoolSlave
}

type BlockchainService struct{}

func (s *BlockchainService) GetPool(ctx context.Context, blockchain *blockchains.Blockchain, solo bool) (Pool, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	pool := Pool{
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
		poolStats, err := client.GetPoolStats(gCtx, &poolProto.GetPoolAssetRequest{
			Solo: solo,
		})
		if err == nil {
			pool.Stats = poolStats

			return nil
		}

		return fmt.Errorf("failed to get pool stats: %w", err)
	})
	g.Go(func() error {
		networkInfo, err := client.GetNetworkInfo(ctx, &emptypb.Empty{})
		if err == nil {
			pool.NetworkInfo = networkInfo

			return nil
		}

		return fmt.Errorf("failed to get pool network info: %w", err)
	})
	g.Go(func() error {
		poolSlaves, err := client.GetPoolSlaves(gCtx, &poolProto.GetPoolAssetRequest{
			Solo: solo,
		})
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

func (s *BlockchainService) GetPoolStats(ctx context.Context, blockchain *blockchains.Blockchain, solo bool) (*poolProto.PoolStats, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	poolStats, err := client.GetPoolStats(ctx, &poolProto.GetPoolAssetRequest{
		Solo: solo,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool (blockchain: %s, solo: %t) stats: %w", blockchain.GetInfo().Blockchain, solo, err)
	}

	return poolStats, nil
}

func (s *BlockchainService) GetPoolNetworkInfo(ctx context.Context, blockchain *blockchains.Blockchain) (*poolProto.NetworkInfo, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	networkInfo, err := client.GetNetworkInfo(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool (blockchain: %s) network info: %w", blockchain.GetInfo().Blockchain, err)
	}

	return networkInfo, nil
}

func (s *BlockchainService) GetPoolSlaves(ctx context.Context, blockchain *blockchains.Blockchain, solo bool) ([]*poolProto.PoolSlave, error) {
	client := poolProto.NewPoolServiceClient(blockchain.GetConnection())
	poolSlaves, err := client.GetPoolSlaves(ctx, &poolProto.GetPoolAssetRequest{
		Solo: solo,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get pool (blockchain: %s, solo: %t) slaves: %w", blockchain.GetInfo().Blockchain, solo, err)
	}

	return poolSlaves.Slaves, nil
}
