package poolsSerializers

import (
	"context"
	"math/big"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type PoolStatsSerializer struct{}

func (s *PoolStatsSerializer) Serialize(ctx context.Context, poolStats *poolProto.PoolStats) apiModels.PoolStats {
	return apiModels.PoolStats{
		MinersCount:  poolStats.MinersCount,
		WorkersCount: poolStats.WorkersCount,
		Hashrate:     new(big.Int).SetBytes(poolStats.Hashrate).String(),
		AvgHashrate:  new(big.Int).SetBytes(poolStats.AvgHashrate).String(),
	}
}
