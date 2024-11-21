package poolsHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolsErrors "github.com/grandminingpool/pool-api/internal/api/handlers/pools/errors"
	poolsSerializers "github.com/grandminingpool/pool-api/internal/api/serializers/pools"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
)

type Handler struct {
	poolsService             *poolsServices.PoolsService
	blockchainPoolSerializer *poolsSerializers.BlockchainPoolSerializer
}

func (h *Handler) GetPools(ctx context.Context, includeSoloStats, includeNetworkInfo bool) apiModels.GetPoolsRes {
	pools, err := h.poolsService.GetPools(ctx, includeSoloStats, includeNetworkInfo)
	if err != nil {
		return poolsErrors.CreateGetPoolsError(err)
	}

	poolsResponse := make([]apiModels.BlockchainPool, 0, len(pools))
	for _, p := range pools {
		poolsResponse = append(poolsResponse, *h.blockchainPoolSerializer.Serialize(ctx, p))
	}

	return &apiModels.PoolsList{
		Pools: poolsResponse,
	}
}

func NewHandler(
	poolsService *poolsServices.PoolsService,
	blockchainPoolSerializer *poolsSerializers.BlockchainPoolSerializer,
) *Handler {
	return &Handler{
		poolsService:             poolsService,
		blockchainPoolSerializer: blockchainPoolSerializer,
	}
}
