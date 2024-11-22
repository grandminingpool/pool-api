package poolsSerializers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
)

type BlockchainPoolSerializer struct {
	poolInfoSerializer        *PoolInfoSerializer
	poolStatsSerializer       *PoolStatsSerializer
	poolNetworkInfoSerializer *PoolNetworkInfoSerialzier
}

func (s *BlockchainPoolSerializer) Serialize(ctx context.Context, blockchainPool *poolsServices.BlockchainPool) *apiModels.BlockchainPool {
	response := &apiModels.BlockchainPool{
		Info:        *s.poolInfoSerializer.Serialize(ctx, blockchainPool.Info),
		Stats:       *s.poolStatsSerializer.Serialize(ctx, blockchainPool.Stats),
		SoloStats:   apiModels.OptPoolStats{},
		NetworkInfo: apiModels.OptPoolNetworkInfo{},
	}

	if blockchainPool.SoloStats != nil {
		response.SoloStats = apiModels.NewOptPoolStats(*s.poolStatsSerializer.Serialize(ctx, blockchainPool.SoloStats))
	}

	if blockchainPool.NetworkInfo != nil {
		response.NetworkInfo = apiModels.NewOptPoolNetworkInfo(*s.poolNetworkInfoSerializer.Serialize(ctx, blockchainPool.NetworkInfo))
	}

	return response
}

func NewBlockchainPoolSerializer(
	poolInfoSerializer *PoolInfoSerializer,
	poolStatsSerializer *PoolStatsSerializer,
	poolNetworkInfoSerializer *PoolNetworkInfoSerialzier,
) *BlockchainPoolSerializer {
	return &BlockchainPoolSerializer{
		poolInfoSerializer:        poolInfoSerializer,
		poolStatsSerializer:       poolStatsSerializer,
		poolNetworkInfoSerializer: poolNetworkInfoSerializer,
	}
}
