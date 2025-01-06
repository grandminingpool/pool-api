package poolsHandlers

import (
	"context"

	poolProto "github.com/grandminingpool/pool-api-proto/generated/pool"
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolsErrors "github.com/grandminingpool/pool-api/internal/api/handlers/pools/errors"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type BlockchainHandler struct {
	blockchainService         *poolsServices.BlockchainService
	poolSerializer            serializers.BaseSerializer[poolsServices.Pool, apiModels.Pool]
	poolInfoSerializer        serializers.BaseSerializer[*poolProto.PoolInfo, apiModels.PoolInfo]
	poolStatsSerializer       serializers.BaseSerializer[*poolProto.PoolStats, apiModels.PoolStats]
	poolSlaveSerializer       serializers.BaseSerializer[*poolProto.PoolSlave, apiModels.PoolSlave]
	poolNetworkInfoSerializer serializers.BaseSerializer[*poolProto.NetworkInfo, apiModels.PoolNetworkInfo]
}

func (h *BlockchainHandler) GetPool(ctx context.Context, blockchain *blockchains.Blockchain, solo bool) apiModels.GetBlockchainPoolRes {
	pool, err := h.blockchainService.GetPool(ctx, blockchain, solo)
	if err != nil {
		return poolsErrors.CreateGetPoolDataError(err)
	}

	response := h.poolSerializer.Serialize(ctx, pool)
	return &response
}

func (h *BlockchainHandler) GetPoolInfo(ctx context.Context, blockchain *blockchains.Blockchain) apiModels.GetBlockchainPoolInfoRes {
	poolInfo, err := h.blockchainService.GetPoolInfo(ctx, blockchain)
	if err != nil {
		return poolsErrors.CreateGetPoolInfoError(err)
	}

	response := h.poolInfoSerializer.Serialize(ctx, poolInfo)
	return &response
}

func (h *BlockchainHandler) GetPoolStats(ctx context.Context, blockchain *blockchains.Blockchain, solo bool) apiModels.GetBlockchainPoolStatsRes {
	poolStats, err := h.blockchainService.GetPoolStats(ctx, blockchain, solo)
	if err != nil {
		return poolsErrors.CreateGetPoolStatsError(err)
	}

	response := h.poolStatsSerializer.Serialize(ctx, poolStats)
	return &response
}

func (h *BlockchainHandler) GetPoolNetworkInfo(ctx context.Context, blockchain *blockchains.Blockchain) apiModels.GetBlockchainPoolNetworkInfoRes {
	networkInfo, err := h.blockchainService.GetPoolNetworkInfo(ctx, blockchain)
	if err != nil {
		return poolsErrors.CreateGetPoolNetworkInfoError(err)
	}

	response := h.poolNetworkInfoSerializer.Serialize(ctx, networkInfo)
	return &response
}

func (h *BlockchainHandler) GetPoolSlaves(ctx context.Context, blockchain *blockchains.Blockchain, solo bool) apiModels.GetBlockchainPoolSlavesRes {
	poolSlaves, err := h.blockchainService.GetPoolSlaves(ctx, blockchain, solo)
	if err != nil {
		return poolsErrors.CreateGetPoolSlavesError(err)
	}

	poolSlavesResponse := make([]apiModels.PoolSlave, 0, len(poolSlaves))
	for _, poolSlave := range poolSlaves {
		poolSlavesResponse = append(poolSlavesResponse, h.poolSlaveSerializer.Serialize(ctx, poolSlave))
	}

	return &apiModels.PoolSlavesList{
		Slaves: poolSlavesResponse,
	}
}

func NewBlockchainHandler(
	blockchainService *poolsServices.BlockchainService,
	poolSerializer serializers.BaseSerializer[poolsServices.Pool, apiModels.Pool],
	poolInfoSerializer serializers.BaseSerializer[*poolProto.PoolInfo, apiModels.PoolInfo],
	poolStatsSerializer serializers.BaseSerializer[*poolProto.PoolStats, apiModels.PoolStats],
	poolSlaveSerializer serializers.BaseSerializer[*poolProto.PoolSlave, apiModels.PoolSlave],
	poolNetworkInfoSerializer serializers.BaseSerializer[*poolProto.NetworkInfo, apiModels.PoolNetworkInfo],
) *BlockchainHandler {
	return &BlockchainHandler{
		blockchainService:         blockchainService,
		poolSerializer:            poolSerializer,
		poolInfoSerializer:        poolInfoSerializer,
		poolStatsSerializer:       poolStatsSerializer,
		poolSlaveSerializer:       poolSlaveSerializer,
		poolNetworkInfoSerializer: poolNetworkInfoSerializer,
	}
}
