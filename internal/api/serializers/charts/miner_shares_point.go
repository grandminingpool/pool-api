package chartsSerializers

import (
	"context"
	"time"

	chartsProto "github.com/grandminingpool/pool-api-proto/generated/charts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type MinerSharesPointSerializer struct{}

func (s *MinerSharesPointSerializer) Serialize(ctx context.Context, point *chartsProto.MinerSharesPoint) *apiModels.MinerSharesPoint {
	return &apiModels.MinerSharesPoint{
		AcceptedSharesCount:   point.AcceptedSharesCount,
		RejectedSharesCount:   point.RejectedSharesCount,
		StaleSharesCount:      point.StaleSharesCount,
		ValidBlockSharesCount: point.ValidBlockSharesCount,
		Date:                  point.Date.AsTime().Format(time.RFC3339),
	}
}
