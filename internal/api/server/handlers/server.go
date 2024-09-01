package apiServerHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	blockchainsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/blockchains"
	poolsHandlers "github.com/grandminingpool/pool-api/internal/api/handlers/pools"
	apiResolvers "github.com/grandminingpool/pool-api/internal/api/resolvers"
)

type ServerHandler struct {
	blockchainsHandler        *blockchainsHandlers.BlockchainsHandler
	poolsBlockchainHandler    *poolsHandlers.BlockchainHandler
	getBlockchainPoolResolver *apiResolvers.BlockchainResolver[apiModels.Pool]
}

func (h *ServerHandler) GetBlockchains(ctx context.Context) ([]apiModels.Blockchain, error) {
	return h.blockchainsHandler.Get(ctx)
}

func (h *ServerHandler) GetBlockchainPool(ctx context.Context, params apiModels.GetBlockchainPoolParams) (*apiModels.Pool, error) {
	return h.getBlockchainPoolResolver.Handler(ctx, params.Blockchain, h.poolsBlockchainHandler.Get)
}
