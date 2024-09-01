package blockchainsHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	"github.com/grandminingpool/pool-api/internal/blockchains"
)

type BlockchainsHandler struct {
	blockchainsService *blockchains.Service
}

func (h *BlockchainsHandler) Get(ctx context.Context) ([]apiModels.Blockchain, error) {
	h.blockchainsService.GetBlockchainsInfo()
}
