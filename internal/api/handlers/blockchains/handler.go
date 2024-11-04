package blockchainsHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type Handler struct {
	blockchainsService   *blockchains.Service
	blockchainSerializer serializers.BaseSerializer[*blockchains.BlockchainInfo, *apiModels.Blockchain]
}

func (h *Handler) Get(ctx context.Context) *apiModels.BlockchainsList {
	blockchainsInfo := h.blockchainsService.GetBlockchainsInfo()
	blockchainsResponse := make([]apiModels.Blockchain, 0, len(blockchainsInfo))
	for _, blockchainInfo := range blockchainsInfo {
		blockchainsResponse = append(blockchainsResponse, *h.blockchainSerializer.Serialize(ctx, &blockchainInfo))
	}

	return &apiModels.BlockchainsList{
		Blockchains: blockchainsResponse,
	}
}
