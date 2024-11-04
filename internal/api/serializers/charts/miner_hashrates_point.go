package chartsSerializers

import (
	"context"
	"math/big"
	"time"

	chartsProto "github.com/grandminingpool/pool-api-proto/generated/charts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type MinerHashratesPointSerializer struct{}

func (s *MinerHashratesPointSerializer) Serialize(ctx context.Context, point *chartsProto.MinerHashratesPoint) *apiModels.MinerHashratesPoint {
	return &apiModels.MinerHashratesPoint{
		Hashrate: new(big.Int).SetBytes(point.Hashrate).String(),
		Date:     point.Date.AsTime().Format(time.RFC3339),
	}
}
