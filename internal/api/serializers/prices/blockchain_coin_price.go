package pricesSerializer

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	numericUtils "github.com/grandminingpool/pool-api/internal/common/utils/numeric"
)

type BlockchainCoinPriceSerializer struct {
	marketPriceSerializer *MarkerPriceSerializer
}

func (s *BlockchainCoinPriceSerializer) Serialize(ctx context.Context, blockchainCoinPrice *pricesServices.BlockchainCoinPrice) *apiModels.BlockchainCoinPrice {
	marketsResponse := make([]apiModels.MarketPrice, 0, len(blockchainCoinPrice.Markets))
	for _, m := range blockchainCoinPrice.Markets {
		marketsResponse = append(marketsResponse, *s.marketPriceSerializer.Serialize(ctx, &m))
	}

	return &apiModels.BlockchainCoinPrice{
		Price:                    blockchainCoinPrice.Price,
		PriceChange24hPercentage: numericUtils.ChangeFloatValueInPercentage(blockchainCoinPrice.Price24hAgo, blockchainCoinPrice.Price),
		Markets:                  marketsResponse,
	}
}

func NewBlockchainCoinPriceSerializer(marketPriceSerializer *MarkerPriceSerializer) *BlockchainCoinPriceSerializer {
	return &BlockchainCoinPriceSerializer{
		marketPriceSerializer: marketPriceSerializer,
	}
}
