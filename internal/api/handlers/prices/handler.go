package pricesHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesErrors "github.com/grandminingpool/pool-api/internal/api/handlers/prices/errors"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type Handler struct {
	pricesService             *pricesServices.PricesService
	blockchainPriceSerializer serializers.BaseSerializer[*pricesServices.BlockchainPriceDB, *apiModels.BlockchainPrice]
}

func (h *Handler) Get(ctx context.Context) apiModels.GetPricesRes {
	prices, err := h.pricesService.GetPrices(ctx)
	if err != nil {
		return pricesErrors.CreateGetPricesError(err)
	}

	pricesResponse := make([]apiModels.BlockchainPrice, 0, len(prices))
	for _, p := range prices {
		pricesResponse = append(pricesResponse, *h.blockchainPriceSerializer.Serialize(ctx, &p))
	}

	return &apiModels.PricesList{
		Prices: pricesResponse,
	}
}

func NewHandler(
	pricesService *pricesServices.PricesService,
	blockchainPriceSerializer serializers.BaseSerializer[*pricesServices.BlockchainPriceDB, *apiModels.BlockchainPrice],
) *Handler {
	return &Handler{
		pricesService:             pricesService,
		blockchainPriceSerializer: blockchainPriceSerializer,
	}
}
