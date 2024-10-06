package poolsHandlers

import (
	"context"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolsErrors "github.com/grandminingpool/pool-api/internal/api/handlers/pools/errors"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

type BlockchainHandler struct {
	blockchainService   *poolsServices.BlockchainService
	poolSerializer      serializers.BaseSerializer[*poolsServices.Pool, *apiModels.Pool]
	poolInfoSerializer  serializers.BaseSerializer[*poolProto.PoolInfo, *apiModels.PoolInfo]
	poolStatsSerializer serializers.BaseSerializer[*poolProto.PoolStats, *apiModels.PoolStats]
	poolSlaveSerializer serializers.BaseSerializer[*poolProto.PoolSlave, *apiModels.PoolSlave]
}

func (h *BlockchainHandler) GetPool(ctx context.Context, blockchain *blockchains.Blockchain) (*apiModels.Pool, error) {
	pool, err := h.blockchainService.GetPool(ctx, blockchain)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(poolsErrors.GetPoolDataError, err)
	}

	return h.poolSerializer.Serialize(ctx, pool), nil
}

func (h *BlockchainHandler) GetPoolInfo(ctx context.Context, blockchain *blockchains.Blockchain) (*apiModels.PoolInfo, error) {
	poolInfo, err := h.blockchainService.GetPoolInfo(ctx, blockchain)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(poolsErrors.GetPoolInfoError, err)
	}

	return h.poolInfoSerializer.Serialize(ctx, poolInfo), nil
}

func (h *BlockchainHandler) GetPoolStats(ctx context.Context, blockchain *blockchains.Blockchain) (*apiModels.PoolStats, error) {
	poolStats, err := h.blockchainService.GetPoolStats(ctx, blockchain)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(poolsErrors.GetPoolStatsError, err)
	}

	return h.poolStatsSerializer.Serialize(ctx, poolStats), nil
}

func (h *BlockchainHandler) GetPoolSlaves(ctx context.Context, blockchain *blockchains.Blockchain) ([]apiModels.PoolSlave, error) {
	poolSlaves, err := h.blockchainService.GetPoolSlaves(ctx, blockchain)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(poolsErrors.GetPoolStatsError, err)
	}

	response := make([]apiModels.PoolSlave, 0, len(poolSlaves))
	for _, poolSlave := range poolSlaves {
		response = append(response, *h.poolSlaveSerializer.Serialize(ctx, poolSlave))
	}

	return response, nil
}
