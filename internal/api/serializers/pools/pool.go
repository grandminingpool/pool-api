package poolsSerializers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
)

type PoolSerializer struct {
	poolInfoSerializer        *PoolInfoSerializer
	poolStatsSerializer       *PoolStatsSerializer
	poolNetworkInfoSerializer *PoolNetworkInfoSerialzier
	poolSlaveSerializer       *PoolSlaveSerialzier
}

func (s *PoolSerializer) Serialize(ctx context.Context, pool *poolsServices.Pool) *apiModels.Pool {
	slaves := make([]apiModels.PoolSlave, 0, len(pool.Slaves))
	for _, poolSlave := range pool.Slaves {
		slaves = append(slaves, *s.poolSlaveSerializer.Serialize(ctx, poolSlave))
	}

	return &apiModels.Pool{
		Info:        *s.poolInfoSerializer.Serialize(ctx, pool.Info),
		Stats:       *s.poolStatsSerializer.Serialize(ctx, pool.Stats),
		NetworkInfo: *s.poolNetworkInfoSerializer.Serialize(ctx, pool.NetworkInfo),
		Slaves:      slaves,
	}
}

func NewPoolSerializer(
	poolInfoSerializer *PoolInfoSerializer,
	poolStatsSerializer *PoolStatsSerializer,
	poolNetworkInfoSerializer *PoolNetworkInfoSerialzier,
	poolSlaveSerializer *PoolSlaveSerialzier,
) *PoolSerializer {
	return &PoolSerializer{
		poolInfoSerializer:        poolInfoSerializer,
		poolStatsSerializer:       poolStatsSerializer,
		poolNetworkInfoSerializer: poolNetworkInfoSerializer,
		poolSlaveSerializer:       poolSlaveSerializer,
	}
}
