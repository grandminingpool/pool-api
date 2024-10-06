package blockchainsHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type BlockchainsHandler struct {
	blockchainsService   *blockchains.Service
	blockchainSerializer serializers.BaseSerializer[*blockchains.BlockchainInfo, *apiModels.Blockchain]
}

func (h *BlockchainsHandler) Get(ctx context.Context) ([]apiModels.Blockchain, error) {
	blockchainsInfo := h.blockchainsService.GetBlockchainsInfo()
	response := make([]apiModels.Blockchain, 0, len(blockchainsInfo))

	for _, blockchainInfo := range blockchainsInfo {
		response = append(response, *h.blockchainSerializer.Serialize(ctx, &blockchainInfo))
	}

	return response, nil
}
