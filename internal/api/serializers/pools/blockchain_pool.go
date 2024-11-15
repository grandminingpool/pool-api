package poolsSerializers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
)

type BlockchainPoolSerializer struct {
	poolInfoSerializer  *PoolInfoSerializer
	poolStatsSerializer *PoolStatsSerializer
}

func (s *BlockchainPoolSerializer) Serialize(ctx context.Context, blockchainPool *poolsServices.BlockchainPool) *apiModels.BlockchainPool {
	return &apiModels.BlockchainPool{
		Info:  *s.poolInfoSerializer.Serialize(ctx, blockchainPool.Info),
		Stats: *s.poolStatsSerializer.Serialize(ctx, blockchainPool.Stats),
	}
}

func NewBlockchainPoolSerializer(
	poolInfoSerializer *PoolInfoSerializer,
	poolStatsSerializer *PoolStatsSerializer,
) *BlockchainPoolSerializer {
	return &BlockchainPoolSerializer{
		poolInfoSerializer:  poolInfoSerializer,
		poolStatsSerializer: poolStatsSerializer,
	}
}
