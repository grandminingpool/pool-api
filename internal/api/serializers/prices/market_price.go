package pricesSerializer

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	numericUtils "github.com/grandminingpool/pool-api/internal/common/utils/numeric"
)

type MarkerPriceSerializer struct{}

func (s *MarkerPriceSerializer) Serialize(ctx context.Context, marketPrice *pricesServices.MarketPriceDB) *apiModels.MarketPrice {
	return &apiModels.MarketPrice{
		Price:                    marketPrice.Price,
		PriceChange24hPercentage: numericUtils.ChangeFloatValueInPercentage(marketPrice.Price24hAgo, marketPrice.Price),
		Ticker:                   marketPrice.MarketTicker,
	}
}
