package poolsServices

import (
	"context"
	"fmt"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PoolsService struct {
	blockchainsService *blockchains.Service
}

type BlockchainPoolStats struct {
	PoolStats  *poolProto.PoolStats
	Blockchain string
}

func (s *PoolsService) getBlockchainPoolStats(
	ctx context.Context,
	blockchain string,
	blockchainConn *grpc.ClientConn,
	poolsCh chan<- *BlockchainPoolStats,
	errCh chan<- error,
) {
	select {
	case <-ctx.Done():
		return
	default:
		client := poolProto.NewPoolServiceClient(blockchainConn)
		poolStats, err := client.GetPoolStats(ctx, &emptypb.Empty{})
		if err != nil {
			errCh <- fmt.Errorf("failed to get pool stats (blockchain: %s), error: %w", blockchain, err)

			return
		}

		poolsCh <- &BlockchainPoolStats{
			PoolStats:  poolStats,
			Blockchain: blockchain,
		}
	}
}

func (s *PoolsService) GetStats(ctx context.Context) ([]*BlockchainPoolStats, error) {
	blockchains := s.blockchainsService.GetBlockchains()

	callsNum := len(blockchains)
	poolsCh := make(chan *BlockchainPoolStats, callsNum)
	errCh := make(chan error, callsNum)
	defer close(poolsCh)
	defer close(errCh)

	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, blockchain := range blockchains {
		go s.getBlockchainPoolStats(
			newCtx,
			blockchain.GetInfo().Blockchain,
			blockchain.GetConnection(),
			poolsCh,
			errCh,
		)
	}

	blockchainsPoolStatsMap := make(map[string]*BlockchainPoolStats)
	defer clear(blockchainsPoolStatsMap)

	for i := 0; i < callsNum; i++ {
		select {
		case err := <-errCh:
			return nil, err
		case blockchainPoolStats := <-poolsCh:
			blockchainsPoolStatsMap[blockchainPoolStats.Blockchain] = blockchainPoolStats
		}
	}

	blockchainsPoolStats := make([]*BlockchainPoolStats, 0, callsNum)
	for _, blockchain := range blockchains {
		blockchainPoolStats, ok := blockchainsPoolStatsMap[blockchain.GetInfo().Blockchain]
		if ok {
			blockchainsPoolStats = append(blockchainsPoolStats, blockchainPoolStats)
		}
	}

	return blockchainsPoolStats, nil
}

func NewPoolsService(blockchainsService *blockchains.Service) *PoolsService {
	return &PoolsService{
		blockchainsService: blockchainsService,
	}
}
