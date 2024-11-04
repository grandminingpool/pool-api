package minersHandlers

import (
	"context"

	poolMinersProto "github.com/grandminingpool/pool-api-proto/generated/pool_miners"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	minersErrors "github.com/grandminingpool/pool-api/internal/api/handlers/miners/errors"
	minersServices "github.com/grandminingpool/pool-api/internal/api/services/miners"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type BlockchainHandler struct {
	blockchainService     *minersServices.BlockchainService
	minerSerializer       serializers.BaseSerializer[*poolMinersProto.Miner, *apiModels.Miner]
	minerWorkerSerializer serializers.BaseSerializer[*poolMinersProto.MinerWorker, *apiModels.MinerWorker]
}

func (h *BlockchainHandler) GetMiners(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	sorts *poolMinersProto.MinersSorts,
	filters *poolMinersProto.MinersFilters,
	limit, offset uint32,
) apiModels.GetBlockchainMinersRes {
	minersList, err := h.blockchainService.GetMiners(
		ctx,
		blockchain,
		sorts,
		filters,
		limit,
		offset,
	)
	if err != nil {
		return minersErrors.CreateGetMinersError(err)
	}

	minersResponse := make([]apiModels.Miner, 0, len(minersList.Miners))
	for _, m := range minersList.Miners {
		minersResponse = append(minersResponse, *h.minerSerializer.Serialize(ctx, m))
	}

	return &apiModels.MinersList{
		Miners: minersResponse,
		Limit:  minersList.Pagination.Limit,
		Offset: minersList.Pagination.Offset,
		Total:  minersList.Pagination.Total,
	}
}

func (h *BlockchainHandler) GetMiner(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	miner string,
) apiModels.GetBlockchainMinerRes {
	minerInfo, err := h.blockchainService.GetMiner(ctx, blockchain, miner)
	if err != nil {
		return minersErrors.CreateGetMinerError(err)
	} else if minerInfo == nil {
		return minersErrors.CreateMinerNotFoundError(miner)
	}

	return h.minerSerializer.Serialize(ctx, minerInfo)
}

func (h *BlockchainHandler) GetMinerWorkers(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	miner string,
) apiModels.GetBlockchainMinerWorkersRes {
	minerWorkers, err := h.blockchainService.GetMinerWorkers(ctx, blockchain, miner)
	if err != nil {
		return minersErrors.CreateGetMinerWorkersError(err)
	}

	minerWorkersResponse := make([]apiModels.MinerWorker, 0, len(minerWorkers))
	for _, mw := range minerWorkers {
		minerWorkersResponse = append(minerWorkersResponse, *h.minerWorkerSerializer.Serialize(ctx, mw))
	}

	return &apiModels.MinerWorkersList{
		Workers: minerWorkersResponse,
	}
}
