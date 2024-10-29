package poolsSerializers

import (
	"context"
	"math/big"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type PoolStatsSerializer struct{}

func (s *PoolStatsSerializer) Serialize(ctx context.Context, poolStats *poolProto.PoolStats) *apiModels.PoolStats {
	soloMinersCount := apiModels.OptUint32{}
	if poolStats.SoloMinersCount != nil {
		soloMinersCount.SetTo(*poolStats.SoloMinersCount)
	}

	soloHashrate := apiModels.OptString{}
	if len(poolStats.SoloHashrate) > 0 {
		soloHashrate.SetTo(new(big.Int).SetBytes(poolStats.SoloHashrate).String())
	}

	soloAvgHashrate := apiModels.OptString{}
	if len(poolStats.SoloAvgHashrate) > 0 {
		soloAvgHashrate.SetTo(new(big.Int).SetBytes(poolStats.SoloAvgHashrate).String())
	}

	soloShareDifficulty := apiModels.OptUint64{}
	if poolStats.SoloShareDifficulty != nil {
		soloShareDifficulty.SetTo(*poolStats.SoloShareDifficulty)
	}

	return &apiModels.PoolStats{
		MinersCount:         poolStats.MinersCount,
		SoloMinersCount:     soloMinersCount,
		Hashrate:            new(big.Int).SetBytes(poolStats.Hashrate).String(),
		AvgHashrate:         new(big.Int).SetBytes(poolStats.AvgHashrate).String(),
		SoloHashrate:        soloHashrate,
		SoloAvgHashrate:     soloAvgHashrate,
		ShareDifficulty:     poolStats.ShareDifficulty,
		SoloShareDifficulty: soloShareDifficulty,
	}
}
