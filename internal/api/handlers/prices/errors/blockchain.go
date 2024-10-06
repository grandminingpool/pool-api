package pricesErrors

import serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"

const (
	BlockchainCoinPriceNotFound serverErrors.ServerErrorCode = "blockchain_coin_price_not_found"
	GetBlockchainCoinPriceError serverErrors.ServerErrorCode = "get_blockchain_coin_price_error"
)
