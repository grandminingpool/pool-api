package poolsHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolsErrors "github.com/grandminingpool/pool-api/internal/api/handlers/pools/errors"
	poolsSerializers "github.com/grandminingpool/pool-api/internal/api/serializers/pools"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
)

type Handler struct {
	poolsService                  *poolsServices.PoolsService
	blockchainPoolStatsSerializer *poolsSerializers.BlockchainPoolStatsSerializer
}

func (h *Handler) GetStats(ctx context.Context) apiModels.GetPoolsStatsRes {
	poolsStats, err := h.poolsService.GetStats(ctx)
	if err != nil {
		return poolsErrors.CreateGetPoolsStatsError(err)
	}

	statsResponse := make([]apiModels.BlockchainPoolStats, 0, len(poolsStats))
	for _, ps := range poolsStats {
		statsResponse = append(statsResponse, *h.blockchainPoolStatsSerializer.Serialize(ctx, ps))
	}

	return &apiModels.PoolsStatsList{
		Stats: statsResponse,
	}
}

func NewHandler(
	poolsService *poolsServices.PoolsService,
	blockchainPoolStatsSerializer *poolsSerializers.BlockchainPoolStatsSerializer,
) *Handler {
	return &Handler{
		poolsService:                  poolsService,
		blockchainPoolStatsSerializer: blockchainPoolStatsSerializer,
	}
}
