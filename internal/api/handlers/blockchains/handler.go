package blockchainsHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	"github.com/grandminingpool/pool-api/internal/blockchains"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type Handler struct {
	blockchainsService       *blockchains.Service
	blockchainInfoSerializer serializers.BaseSerializer[*blockchains.BlockchainInfo, *apiModels.BlockchainInfo]
}

func (h *Handler) Get(ctx context.Context) *apiModels.BlockchainsList {
	blockchainsInfos := h.blockchainsService.GetBlockchainsInfos()
	blockchainsResponse := make([]apiModels.BlockchainInfo, 0, len(blockchainsInfos))
	for _, blockchainInfo := range blockchainsInfos {
		blockchainsResponse = append(blockchainsResponse, *h.blockchainInfoSerializer.Serialize(ctx, &blockchainInfo))
	}

	return &apiModels.BlockchainsList{
		Blockchains: blockchainsResponse,
	}
}

func NewHandler(
	blockchainsService *blockchains.Service,
	blockchainInfoSerializer serializers.BaseSerializer[*blockchains.BlockchainInfo, *apiModels.BlockchainInfo],
) *Handler {
	return &Handler{
		blockchainsService:       blockchainsService,
		blockchainInfoSerializer: blockchainInfoSerializer,
	}
}
