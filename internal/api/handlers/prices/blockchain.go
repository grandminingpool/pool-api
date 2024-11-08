package pricesHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesErrors "github.com/grandminingpool/pool-api/internal/api/handlers/prices/errors"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type BlockchainHandler struct {
	pricesService               *pricesServices.PricesService
	blockchainMarketsSerializer serializers.BaseSerializer[*pricesServices.BlockchainMarkets, *apiModels.BlockchainMarkets]
}

func (h *BlockchainHandler) GetMarkets(ctx context.Context, blockchainCoin string) apiModels.GetBlockchainMarketsRes {
	blockchainMarkets, err := h.pricesService.GetBlockchainMarkets(ctx, blockchainCoin)
	if err != nil {
		return pricesErrors.CreateGetBlockchainMarketsError(err)
	} else if blockchainMarkets == nil {
		return pricesErrors.CreateBlockchainMarketsNotFoundError(blockchainCoin)
	}

	return h.blockchainMarketsSerializer.Serialize(ctx, blockchainMarkets)
}

func NewBlockchainHandler(
	pricesService *pricesServices.PricesService,
	blockchainMarketsSerializer serializers.BaseSerializer[*pricesServices.BlockchainMarkets, *apiModels.BlockchainMarkets],
) *BlockchainHandler {
	return &BlockchainHandler{
		pricesService:               pricesService,
		blockchainMarketsSerializer: blockchainMarketsSerializer,
	}
}
