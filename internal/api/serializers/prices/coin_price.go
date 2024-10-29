package pricesSerializer

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	numericUtils "github.com/grandminingpool/pool-api/internal/common/utils/numeric"
)

type CoinPriceSerializer struct{}

func (s *CoinPriceSerializer) Serialize(ctx context.Context, coinPrice *pricesServices.CoinPriceDB) *apiModels.CoinPrice {
	return &apiModels.CoinPrice{
		Price:                    coinPrice.Price,
		PriceChange24hPercentage: numericUtils.ChangeFloatValueInPercentage(coinPrice.Price24hAgo, coinPrice.Price),
		Coin:                     coinPrice.Coin,
	}
}
