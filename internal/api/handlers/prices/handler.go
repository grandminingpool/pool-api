package pricesHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesErrors "github.com/grandminingpool/pool-api/internal/api/handlers/prices/errors"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
)

type Handler struct {
	pricesService       *pricesServices.PricesService
	coinPriceSerializer serializers.BaseSerializer[*pricesServices.CoinPriceDB, *apiModels.CoinPrice]
}

func (h *Handler) Get(ctx context.Context) apiModels.GetPricesRes {
	prices, err := h.pricesService.GetPrices(ctx)
	if err != nil {
		return pricesErrors.CreateGetPricesError(err)
	}

	pricesResponse := make([]apiModels.CoinPrice, 0, len(prices))
	for _, p := range prices {
		pricesResponse = append(pricesResponse, *h.coinPriceSerializer.Serialize(ctx, &p))
	}

	return &apiModels.CoinPricesList{
		Prices: pricesResponse,
	}
}
