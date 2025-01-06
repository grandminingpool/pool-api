package pricesSerializer

import (
	"context"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	numericUtils "github.com/grandminingpool/pool-api/internal/common/utils/numeric"
)

type BlockchainPriceSerializer struct{}

func (s *BlockchainPriceSerializer) Serialize(ctx context.Context, blockchainPrice pricesServices.BlockchainPriceDB) apiModels.BlockchainPrice {
	return apiModels.BlockchainPrice{
		Price:                    blockchainPrice.Price,
		PriceChange24hPercentage: numericUtils.ChangeFloatValueInPercentage(blockchainPrice.Price24hAgo, blockchainPrice.Price),
		Blockchain:               blockchainPrice.Blockchain,
	}
}
