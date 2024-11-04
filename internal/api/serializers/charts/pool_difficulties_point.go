package chartsSerializers

import (
	"context"
	"time"

	chartsProto "github.com/grandminingpool/pool-api-proto/generated/charts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type PoolDifficultiesPointSerializer struct{}

func (s *PoolDifficultiesPointSerializer) Serialize(ctx context.Context, point *chartsProto.PoolDifficultiesPoint) *apiModels.PoolDifficultiesPoint {
	return &apiModels.PoolDifficultiesPoint{
		ShareDifficulty: point.ShareDifficulty,
		Date:            point.Date.AsTime().Format(time.RFC3339),
	}
}
