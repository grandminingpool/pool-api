package poolsSerializers

import (
	"context"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
)

type PoolInfoSerializer struct{}

func (s *PoolInfoSerializer) Serialize(ctx context.Context, poolInfo *poolProto.PoolInfo) *apiModels.PoolInfo {
	var payoutMode apiModels.PayoutMode
	switch poolInfo.PayoutMode {
	case poolProto.PayoutMode_PROP:
		payoutMode = apiModels.PayoutModeProp
		break
	default:
		payoutMode = apiModels.PayoutModePplns
	}

	poolFee := apiModels.PoolFee{
		Fee:     poolInfo.Fee.Fee,
		SoloFee: apiModels.OptFloat64{},
	}

	if poolInfo.Fee.SoloFee != nil {
		poolFee.SoloFee.SetTo(*poolInfo.Fee.SoloFee)
	}

	payoutsInfo := apiModels.PayoutsInfo{
		Interval:  int(poolInfo.PayoutsInfo.Interval),
		MinPayout: apiModels.OptInt{},
		MaxPayout: apiModels.OptInt{},
	}

	if poolInfo.PayoutsInfo.MinPayout != nil {
		payoutsInfo.MinPayout.SetTo(int(*poolInfo.PayoutsInfo.MinPayout))
	}

	if poolInfo.PayoutsInfo.MaxPayout != nil {
		payoutsInfo.MaxPayout.SetTo(int(*poolInfo.PayoutsInfo.MaxPayout))
	}

	return &apiModels.PoolInfo{
		Blockchain:  poolInfo.Blockchain,
		Host:        poolInfo.Host,
		Algos:       poolInfo.Algos,
		PayoutMode:  payoutMode,
		Solo:        poolInfo.Solo,
		Fee:         poolFee,
		PayoutsInfo: payoutsInfo,
		Agents:      poolInfo.Agents,
	}
}