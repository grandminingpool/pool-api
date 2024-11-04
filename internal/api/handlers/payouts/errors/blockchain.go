package payoutsErrors

import (
	"fmt"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

const (
	GetPayoutsError           serverErrors.ServerErrorCode = "get_payouts_error"
	GetMinerBalanceError      serverErrors.ServerErrorCode = "get_miner_balance_error"
	MinerBalanceNotFoundError serverErrors.ServerErrorCode = "miner_balance_not_found"
)

func CreateGetPayoutsError(err error) *apiModels.GetBlockchainPayoutsInternalServerError {
	return &apiModels.GetBlockchainPayoutsInternalServerError{
		Code:    string(GetPayoutsError),
		Message: err.Error(),
	}
}

func CreateMinerBalanceNotFoundError(miner string) *apiModels.GetBlockchainMinerBalanceNotFound {
	return &apiModels.GetBlockchainMinerBalanceNotFound{
		Code:    string(GetMinerBalanceError),
		Message: fmt.Sprintf("miner '%s' balance not found", miner),
	}
}

func CreateGetMinerBalanceError(err error) *apiModels.GetBlockchainMinerBalanceInternalServerError {
	return &apiModels.GetBlockchainMinerBalanceInternalServerError{
		Code:    string(GetMinerBalanceError),
		Message: err.Error(),
	}
}
