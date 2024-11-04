package pricesHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesErrors "github.com/grandminingpool/pool-api/internal/api/handlers/prices/errors"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type BlockchainHandler struct {
	pricesService                 *pricesServices.PricesService
	blockchainCoinPriceSerializer serializers.BaseSerializer[*pricesServices.BlockchainCoinPrice, *apiModels.BlockchainCoinPrice]
}

func (h *BlockchainHandler) GetPrice(ctx context.Context, blockchainCoin string) apiModels.GetBlockchainCoinPriceRes {
	price, err := h.pricesService.GetBlockchainCoinPrice(ctx, blockchainCoin)
	if err != nil {
		return pricesErrors.CreateGetBlockchainCoinPriceError(err)
	} else if price == nil {
		return pricesErrors.CreateBlockchainCoinPriceNotFoundError(blockchainCoin)
	}

	return h.blockchainCoinPriceSerializer.Serialize(ctx, price)
}

func NewBlockchainHandler(
	pricesService *pricesServices.PricesService,
	blockchainCoinPriceSerializer serializers.BaseSerializer[*pricesServices.BlockchainCoinPrice, *apiModels.BlockchainCoinPrice],
) *BlockchainHandler {
	return &BlockchainHandler{
		pricesService:                 pricesService,
		blockchainCoinPriceSerializer: blockchainCoinPriceSerializer,
	}
}
