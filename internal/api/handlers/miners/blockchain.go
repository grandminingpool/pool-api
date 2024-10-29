package minersHandlers

import (
	"context"
	"fmt"

	poolMinersProto "github.com/grandminingpool/pool-api-proto/generated/pool_miners"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	minersErrors "github.com/grandminingpool/pool-api/internal/api/handlers/miners/errors"
	minersServices "github.com/grandminingpool/pool-api/internal/api/services/miners"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
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
) (*apiModels.MinersList, error) {
	minersList, err := h.blockchainService.GetMiners(
		ctx,
		blockchain,
		sorts,
		filters,
		limit,
		offset,
	)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(minersErrors.GetMinersError, err)
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
	}, nil
}

func (h *BlockchainHandler) GetMiner(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	miner string,
) (*apiModels.Miner, error) {
	minerInfo, err := h.blockchainService.GetMiner(ctx, blockchain, miner)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(minersErrors.GetMinerError, err)
	} else if minerInfo == nil {
		return nil, serverErrors.CreateNotFoundError(minersErrors.MinerNotFoundError, fmt.Errorf("miner '%s' not found", miner))
	}

	return h.minerSerializer.Serialize(ctx, minerInfo), nil
}

func (h *BlockchainHandler) GetMinerWorkers(
	ctx context.Context,
	blockchain *blockchains.Blockchain,
	miner string,
) ([]apiModels.MinerWorker, error) {
	minerWorkers, err := h.blockchainService.GetMinerWorkers(ctx, blockchain, miner)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(minersErrors.GetMinerWorkersError, err)
	}

	response := make([]apiModels.MinerWorker, 0, len(minerWorkers))
	for _, mw := range minerWorkers {
		response = append(response, *h.minerWorkerSerializer.Serialize(ctx, mw))
	}

	return response, nil
}
