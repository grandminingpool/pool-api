package pricesErrors

import (
	"fmt"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

const (
	BlockchainMarketsNotFoundError serverErrors.ServerErrorCode = "blockchain_markets_not_found"
	GetBlockchainMarketsError      serverErrors.ServerErrorCode = "get_blockchain_markets_error"
)

func CreateBlockchainMarketsNotFoundError(blockchain string) *apiModels.GetBlockchainMarketsNotFound {
	return &apiModels.GetBlockchainMarketsNotFound{
		Code:    string(BlockchainMarketsNotFoundError),
		Message: fmt.Sprintf("blockchain '%s' price not found", blockchain),
	}
}

func CreateGetBlockchainMarketsError(err error) *apiModels.GetBlockchainMarketsInternalServerError {
	return &apiModels.GetBlockchainMarketsInternalServerError{
		Code:    string(GetBlockchainMarketsError),
		Message: err.Error(),
	}
}
