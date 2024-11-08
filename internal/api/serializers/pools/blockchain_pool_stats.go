package poolsSerializers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
)

type BlockchainPoolStatsSerializer struct {
	poolStatsSerializer *PoolStatsSerializer
}

func (s *BlockchainPoolStatsSerializer) Serialize(ctx context.Context, blockchainPoolStats *poolsServices.BlockchainPoolStats) *apiModels.BlockchainPoolStats {
	return &apiModels.BlockchainPoolStats{
		Blockchain: blockchainPoolStats.Blockchain,
		Stats:      *s.poolStatsSerializer.Serialize(ctx, blockchainPoolStats.PoolStats),
	}
}

func NewBlockchainPoolStatsSerializer(poolStatsSerializer *PoolStatsSerializer) *BlockchainPoolStatsSerializer {
	return &BlockchainPoolStatsSerializer{
		poolStatsSerializer: poolStatsSerializer,
	}
}
