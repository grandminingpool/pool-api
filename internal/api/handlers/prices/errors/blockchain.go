package pricesErrors

import (
	"fmt"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

const (
	BlockchainCoinPriceNotFoundError serverErrors.ServerErrorCode = "blockchain_coin_price_not_found"
	GetBlockchainCoinPriceError      serverErrors.ServerErrorCode = "get_blockchain_coin_price_error"
)

func CreateBlockchainCoinPriceNotFoundError(blockchainCoin string) *apiModels.GetBlockchainCoinPriceNotFound {
	return &apiModels.GetBlockchainCoinPriceNotFound{
		Code:    string(BlockchainCoinPriceNotFoundError),
		Message: fmt.Sprintf("blockchain coin: '%s' price not found", blockchainCoin),
	}
}

func CreateGetBlockchainCoinPriceError(err error) *apiModels.GetBlockchainCoinPriceInternalServerError {
	return &apiModels.GetBlockchainCoinPriceInternalServerError{
		Code:    string(GetBlockchainCoinPriceError),
		Message: err.Error(),
	}
}
