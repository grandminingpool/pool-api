package chartsSerializers

import (
	"context"
	"time"

	chartsProto "github.com/grandminingpool/pool-api-proto/generated/charts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type MinerProfitabilitiesPointSerializer struct{}

func (s *MinerProfitabilitiesPointSerializer) Serialize(ctx context.Context, point *chartsProto.MinerProfitabilityPoint) *apiModels.MinerProfitabilitiesPoint {
	return &apiModels.MinerProfitabilitiesPoint{
		Balance: point.Balance,
		Date:    point.Date.AsTime().Format(time.RFC3339),
	}
}
