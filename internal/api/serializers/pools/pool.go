package poolsSerializers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
)

type PoolSerializer struct {
	poolInfoSerializer  *PoolInfoSerializer
	poolStatsSerializer *PoolStatsSerializer
	poolSlaveSerializer *PoolSlaveSerialzier
}

func (s *PoolSerializer) Serialize(ctx context.Context, pool *poolsServices.Pool) *apiModels.Pool {
	slaves := make([]apiModels.PoolSlave, 0, len(pool.Slaves))
	for _, poolSlave := range pool.Slaves {
		slaves = append(slaves, *s.poolSlaveSerializer.Serialize(ctx, poolSlave))
	}

	return &apiModels.Pool{
		Info:   *s.poolInfoSerializer.Serialize(ctx, pool.Info),
		Stats:  *s.poolStatsSerializer.Serialize(ctx, pool.Stats),
		Slaves: slaves,
	}
}

func NewPoolSerializer(
	poolInfoSerializer *PoolInfoSerializer,
	poolStatsSerializer *PoolStatsSerializer,
	poolSlaveSerializer *PoolSlaveSerialzier,
) *PoolSerializer {
	return &PoolSerializer{
		poolInfoSerializer:  poolInfoSerializer,
		poolStatsSerializer: poolStatsSerializer,
		poolSlaveSerializer: poolSlaveSerializer,
	}
}
