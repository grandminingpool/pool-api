package chartsHandlers

import (
	"context"

	chartsProto "github.com/grandminingpool/pool-api-proto/generated/charts"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	chartsErrors "github.com/grandminingpool/pool-api/internal/api/handlers/charts/errors"
	chartsServices "github.com/grandminingpool/pool-api/internal/api/services/charts"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type BlockchainHandler struct {
	blockchainService               *chartsServices.BlockchainService
	poolStatsPointSerializer        serializers.BaseSerializer[*chartsProto.PoolStatsPoint, *apiModels.PoolStatsPoint]
	poolDifficultiesPointSerializer serializers.BaseSerializer[*chartsProto.PoolDifficultiesPoint, *apiModels.PoolDifficultiesPoint]
	roundsPointSerializer           serializers.BaseSerializer[*chartsProto.RoundsPoint, *apiModels.RoundsPoint]
	minerHashratesPointSerializer   serializers.BaseSerializer[*chartsProto.MinerHashratesPoint, *apiModels.MinerHashratesPoint]
	minerSharesPointSerializer      serializers.BaseSerializer[*chartsProto.MinerSharesPoint, *apiModels.MinerSharesPoint]
}

func (h *BlockchainHandler) minerHashratesResponse(ctx context.Context, hashratesPoints []*chartsProto.MinerHashratesPoint) *apiModels.MinerHashratesPoints {
	points := make([]apiModels.MinerHashratesPoint, 0, len(hashratesPoints))
	for _, p := range hashratesPoints {
		points = append(points, *h.minerHashratesPointSerializer.Serialize(ctx, p))
	}

	return &apiModels.MinerHashratesPoints{
		Points: points,
	}
}

func (h *BlockchainHandler) minerSharesResponse(ctx context.Context, sharesPoints []*chartsProto.MinerSharesPoint) *apiModels.MinerSharesPoints {
	points := make([]apiModels.MinerSharesPoint, 0, len(sharesPoints))
	for _, p := range sharesPoints {
		points = append(points, *h.minerSharesPointSerializer.Serialize(ctx, p))
	}

	return &apiModels.MinerSharesPoints{
		Points: points,
	}
}

func (h *BlockchainHandler) GetPoolStatsChart(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	solo *apiModels.OptBool,
) apiModels.GetBlockchainPoolStatsChartRes {
	poolStatsPoints, err := h.blockchainService.GetPoolStatsChartPoints(ctx, blockchain, period, solo)
	if err != nil {
		return chartsErrors.CreateGetPoolStatsChartError(err)
	}

	points := make([]apiModels.PoolStatsPoint, 0, len(poolStatsPoints))
	for _, p := range poolStatsPoints {
		points = append(points, *h.poolStatsPointSerializer.Serialize(ctx, p))
	}

	return &apiModels.PoolStatsPoints{
		Points: points,
	}
}

func (h *BlockchainHandler) GetPoolDifficultiesChart(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	solo *apiModels.OptBool,
) apiModels.GetBlockchainPoolDifficultiesChartRes {
	poolDifficultiesPoints, err := h.blockchainService.GetPoolDifficultiesChartPoints(ctx, blockchain, period, solo)
	if err != nil {
		return chartsErrors.CreateGetPoolDifficultiesChartError(err)
	}

	points := make([]apiModels.PoolDifficultiesPoint, 0, len(poolDifficultiesPoints))
	for _, p := range poolDifficultiesPoints {
		points = append(points, *h.poolDifficultiesPointSerializer.Serialize(ctx, p))
	}

	return &apiModels.PoolDifficultiesPoints{
		Points: points,
	}
}

func (h *BlockchainHandler) GetRoundsChart(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
) apiModels.GetBlockchainRoundsChartRes {
	roundsPoints, err := h.blockchainService.GetRoundsChartPoints(ctx, blockchain, period)
	if err != nil {
		return chartsErrors.CreateGetRoundsChartError(err)
	}

	points := make([]apiModels.RoundsPoint, 0, len(roundsPoints))
	for _, p := range roundsPoints {
		points = append(points, *h.roundsPointSerializer.Serialize(ctx, p))
	}

	return &apiModels.RoundsPoints{
		Points: points,
	}
}

func (h *BlockchainHandler) GetMinerHashratesChart(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	miner string,
	solo *apiModels.OptBool,
) apiModels.GetBlockchainMinerHashratesChartRes {
	minerHashratesPoints, err := h.blockchainService.GetMinerHashratesChartPoints(ctx, blockchain, period, miner, solo)
	if err != nil {
		return chartsErrors.CreateGetMinerHashratesChartError(err)
	}

	return h.minerHashratesResponse(ctx, minerHashratesPoints)
}

func (h *BlockchainHandler) GetMinerWorkerHashratesChart(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	miner, worker string,
) apiModels.GetBlockchainMinerWorkerHashratesChartRes {
	minerWorkerHashratesPoints, err := h.blockchainService.GetMinerWorkerHashratesChartPoints(ctx, blockchain, period, miner, worker)
	if err != nil {
		return chartsErrors.CreateGetMinerWorkerHashratesChartError(err)
	}

	return h.minerHashratesResponse(ctx, minerWorkerHashratesPoints)
}

func (h *BlockchainHandler) GetMinerSharesChart(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	miner string,
	solo *apiModels.OptBool,
) apiModels.GetBlockchainMinerSharesChartRes {
	minerSharesPoints, err := h.blockchainService.GetMinerSharesChartPoints(ctx, blockchain, period, miner, solo)
	if err != nil {
		return chartsErrors.CreateGetMinerSharesChartError(err)
	}

	return h.minerSharesResponse(ctx, minerSharesPoints)
}

func (h *BlockchainHandler) GetMinerWorkerSharesChart(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	period *apiModels.ChartPeriod,
	miner, worker string,
) apiModels.GetBlockchainMinerWorkerSharesChartRes {
	minerWorkerSharesPoints, err := h.blockchainService.GetMinerWorkerSharesChartPoints(ctx, blockchain, period, miner, worker)
	if err != nil {
		return chartsErrors.CreateGetMinerWorkerSharesChartError(err)
	}

	return h.minerSharesResponse(ctx, minerWorkerSharesPoints)
}
