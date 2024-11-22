package poolsServices

import (
	"context"
	"fmt"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PoolsService struct {
	blockchainsService *blockchains.Service
}

type BlockchainPool struct {
	Info        *poolProto.PoolInfo
	Stats       *poolProto.PoolStats
	SoloStats   *poolProto.PoolStats
	NetworkInfo *poolProto.NetworkInfo
	Blockchain  string
}

func (s *PoolsService) getBlockchainPool(
	ctx context.Context,
	blockchain string,
	blockchainConn *grpc.ClientConn,
	includeSoloStats, includeNetworkInfo bool,
	poolsCh chan<- *BlockchainPool,
	errCh chan<- error,
) {
	select {
	case <-ctx.Done():
		return
	default:
		pool := &BlockchainPool{}
		client := poolProto.NewPoolServiceClient(blockchainConn)
		g, gCtx := errgroup.WithContext(ctx)
		g.Go(func() error {
			poolInfo, err := client.GetPoolInfo(gCtx, &emptypb.Empty{})
			if err != nil {
				return fmt.Errorf("failed to get pool (blockchain: %s) info, error: %w", blockchain, err)
			}

			pool.Info = poolInfo

			return nil
		})
		g.Go(func() error {
			poolStats, err := client.GetPoolStats(ctx, &poolProto.GetPoolAssetRequest{
				Solo: false,
			})
			if err != nil {
				return fmt.Errorf("failed to get pool (blockchain: %s) stats, error: %w", blockchain, err)
			}

			pool.Stats = poolStats

			return nil
		})

		if includeSoloStats {
			g.Go(func() error {
				poolSoloStats, err := client.GetPoolStats(ctx, &poolProto.GetPoolAssetRequest{
					Solo: true,
				})

				if err != nil {
					e, ok := status.FromError(err)
					if ok && e.Code() == codes.Unimplemented {
						return nil
					}

					return fmt.Errorf("failed to get pool (blockchain: %s) solo stats, error: %w", blockchain, err)
				}

				pool.SoloStats = poolSoloStats

				return nil
			})
		}

		if includeNetworkInfo {
			g.Go(func() error {
				networkInfo, err := client.GetNetworkInfo(ctx, &emptypb.Empty{})
				if err != nil {
					return fmt.Errorf("failed to get pool (blockchain: %s) network info, error: %w", blockchain, err)
				}

				pool.NetworkInfo = networkInfo

				return nil
			})
		}

		if err := g.Wait(); err != nil {
			errCh <- err

			return
		}

		poolsCh <- pool
	}
}

func (s *PoolsService) GetPools(ctx context.Context, includeSoloStats, includeNetworkInfo bool) ([]*BlockchainPool, error) {
	blockchains := s.blockchainsService.GetBlockchains()

	callsNum := len(blockchains)
	poolsCh := make(chan *BlockchainPool, callsNum)
	errCh := make(chan error, callsNum)
	defer close(poolsCh)
	defer close(errCh)

	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, blockchain := range blockchains {
		go s.getBlockchainPool(
			newCtx,
			blockchain.GetInfo().Blockchain,
			blockchain.GetConnection(),
			includeSoloStats,
			includeNetworkInfo,
			poolsCh,
			errCh,
		)
	}

	blockchainsPoolsMap := make(map[string]*BlockchainPool)
	defer clear(blockchainsPoolsMap)

	for i := 0; i < callsNum; i++ {
		select {
		case err := <-errCh:
			return nil, err
		case blockchainPool := <-poolsCh:
			blockchainsPoolsMap[blockchainPool.Info.Blockchain] = blockchainPool
		}
	}

	blockchainsPools := make([]*BlockchainPool, 0, callsNum)
	for _, blockchain := range blockchains {
		blockchainPool, ok := blockchainsPoolsMap[blockchain.GetInfo().Blockchain]
		if ok {
			blockchainsPools = append(blockchainsPools, blockchainPool)
		}
	}

	return blockchainsPools, nil
}

func NewPoolsService(blockchainsService *blockchains.Service) *PoolsService {
	return &PoolsService{
		blockchainsService: blockchainsService,
	}
}
