package pricesHandlers

import (
	"context"
	"fmt"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	pricesErrors "github.com/grandminingpool/pool-api/internal/api/handlers/prices/errors"
	pricesServices "github.com/grandminingpool/pool-api/internal/api/services/prices"
	"github.com/grandminingpool/pool-api/internal/common/serializers"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

type BlockchainHandler struct {
	pricesService                 *pricesServices.PricesService
	blockchainCoinPriceSerializer serializers.BaseSerializer[*pricesServices.BlockchainCoinPrice, *apiModels.BlockchainCoinPrice]
}

func (h *BlockchainHandler) GetPrice(ctx context.Context, blockchainCoin string) (*apiModels.BlockchainCoinPrice, error) {
	price, err := h.pricesService.GetBlockchainCoinPrice(ctx, blockchainCoin)
	if err != nil {
		return nil, serverErrors.CreateInternalServerError(pricesErrors.GetBlockchainCoinPriceError, err)
	} else if price == nil {
		return nil, serverErrors.CreateNotFoundError(pricesErrors.BlockchainCoinPriceNotFound, fmt.Errorf("blockchain (coin: %s) price not found", blockchainCoin))
	}

	return h.blockchainCoinPriceSerializer.Serialize(ctx, price), nil
}
