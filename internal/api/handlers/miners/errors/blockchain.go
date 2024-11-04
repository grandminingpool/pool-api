package minersErrors

import (
	"fmt"

	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

const (
	GetMinersError            serverErrors.ServerErrorCode = "get_miners_error"
	GetMinerError             serverErrors.ServerErrorCode = "get_miner_error"
	MinerNotFoundError        serverErrors.ServerErrorCode = "miner_not_found"
	GetMinerWorkersError      serverErrors.ServerErrorCode = "get_miner_workers_error"
	MinerWorkersNotFoundError serverErrors.ServerErrorCode = "miner_workers_not_found"
)

func CreateGetMinersError(err error) *apiModels.GetBlockchainMinersInternalServerError {
	return &apiModels.GetBlockchainMinersInternalServerError{
		Code:    string(GetMinersError),
		Message: err.Error(),
	}
}

func CreateGetMinerError(err error) *apiModels.GetBlockchainMinerInternalServerError {
	return &apiModels.GetBlockchainMinerInternalServerError{
		Code:    string(GetMinerError),
		Message: err.Error(),
	}
}

func CreateMinerNotFoundError(miner string) *apiModels.GetBlockchainMinerNotFound {
	return &apiModels.GetBlockchainMinerNotFound{
		Code:    string(MinerNotFoundError),
		Message: fmt.Sprintf("miner '%s' not found", miner),
	}
}

func CreateGetMinerWorkersError(err error) *apiModels.GetBlockchainMinerWorkersInternalServerError {
	return &apiModels.GetBlockchainMinerWorkersInternalServerError{
		Code:    string(GetMinerWorkersError),
		Message: err.Error(),
	}
}
