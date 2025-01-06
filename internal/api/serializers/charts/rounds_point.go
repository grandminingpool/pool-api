package chartsSerializers

import (
	"context"
	"time"

	chartsProto "github.com/grandminingpool/pool-api-proto/generated/charts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type RoundsPointSerializer struct{}

func (s *RoundsPointSerializer) Serialize(ctx context.Context, point *chartsProto.RoundsPoint) apiModels.RoundsPoint {
	return apiModels.RoundsPoint{
		RoundsCount:      point.RoundsCount,
		MinersCount:      point.MinersCount,
		AvgRoundInterval: point.AvgRoundInterval,
		Date:             point.Date.AsTime().Format(time.RFC3339),
	}
}
