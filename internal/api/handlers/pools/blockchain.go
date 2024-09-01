package poolsHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	poolsServices "github.com/grandminingpool/pool-api/internal/api/services/pools"
	"github.com/grandminingpool/pool-api/internal/blockchains"
)

type BlockchainHandler struct {
	blockchainService *poolsServices.BlockchainService
}

func (h *BlockchainHandler) Get(ctx context.Context, blockchain *blockchains.Blockchain) (*apiModels.Pool, error) {
	pool, err := h.blockchainService.GetAll(ctx, blockchain)
	if err != nil {

	}

	// todo: serialize pool
}
