package pricesHandlers

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesErrors "github.com/grandminingpool/pool-api/internal/api/handlers/prices/errors"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

type Handler struct {
	pricesService       *pricesServices.PricesService
	coinPriceSerializer serializers.BaseSerializer[*pricesServices.CoinPriceDB, *apiModels.CoinPrice]
}

func (h *Handler) Get(ctx context.Context) ([]apiModels.CoinPrice, error) {
	prices, err := h.pricesService.GetPrices(ctx)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(pricesErrors.GetPricesError, err)
	}

	response := make([]apiModels.CoinPrice, 0, len(prices))

	for _, p := range prices {
		response = append(response, *h.coinPriceSerializer.Serialize(ctx, &p))
	}

	return response, nil
}
