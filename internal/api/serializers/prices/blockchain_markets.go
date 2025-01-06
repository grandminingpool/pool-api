package pricesSerializer

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	numericUtils "github.com/grandminingpool/pool-api/internal/common/utils/numeric"
)

type BlockchainMarketsSerializer struct {
	marketPriceSerializer *MarkerPriceSerializer
}

func (s *BlockchainMarketsSerializer) Serialize(ctx context.Context, blockchainMarkets *pricesServices.BlockchainMarkets) apiModels.BlockchainMarkets {
	marketsResponse := make([]apiModels.MarketPrice, 0, len(blockchainMarkets.Markets))
	for _, m := range blockchainMarkets.Markets {
		marketsResponse = append(marketsResponse, s.marketPriceSerializer.Serialize(ctx, m))
	}

	return apiModels.BlockchainMarkets{
		Price:                    blockchainMarkets.Price,
		PriceChange24hPercentage: numericUtils.ChangeFloatValueInPercentage(blockchainMarkets.Price24hAgo, blockchainMarkets.Price),
		Markets:                  marketsResponse,
	}
}

func NewBlockchainMarketsSerializer(marketPriceSerializer *MarkerPriceSerializer) *BlockchainMarketsSerializer {
	return &BlockchainMarketsSerializer{
		marketPriceSerializer: marketPriceSerializer,
	}
}
