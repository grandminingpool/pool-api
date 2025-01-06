package chartsSerializers

import (
	"context"
	"math/big"
	"time"

	chartsProto "github.com/grandminingpool/pool-api-proto/generated/charts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type PoolStatsPointSerializer struct{}

func (s *PoolStatsPointSerializer) Serialize(ctx context.Context, point *chartsProto.PoolStatsPoint) apiModels.PoolStatsPoint {
	return apiModels.PoolStatsPoint{
		Hashrate:    new(big.Int).SetBytes(point.Hashrate).String(),
		AvgHashrate: new(big.Int).SetBytes(point.AvgHashrate).String(),
		MinersCount: point.MinersCount,
		Date:        point.Date.AsTime().Format(time.RFC3339),
	}
}
