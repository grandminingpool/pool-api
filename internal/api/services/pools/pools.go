package poolsServices

import (
	"context"
	"fmt"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PoolsService struct {
	blockchainsService *blockchains.Service
}

type BlockchainPool struct {
	Info       *poolProto.PoolInfo
	Stats      *poolProto.PoolStats
	Blockchain string
}

func (s *PoolsService) getBlockchainPool(
	ctx context.Context,
	blockchain string,
	blockchainConn *grpc.ClientConn,
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
				return fmt.Errorf("failed to get pool info (blockchain: %s), error: %w", blockchain, err)
			}

			pool.Info = poolInfo

			return nil
		})
		g.Go(func() error {
			poolStats, err := client.GetPoolStats(gCtx, &emptypb.Empty{})
			if err != nil {
				return fmt.Errorf("failed to get pool stats (blockchain: %s), error: %w", blockchain, err)
			}

			pool.Stats = poolStats

			return nil
		})

		if err := g.Wait(); err != nil {
			errCh <- err

			return
		}

		poolsCh <- pool
	}
}

func (s *PoolsService) GetPools(ctx context.Context) ([]*BlockchainPool, error) {
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
