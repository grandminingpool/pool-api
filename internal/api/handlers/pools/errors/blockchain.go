package poolsErrors

import (
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

const (
	GetPoolDataError        serverErrors.ServerErrorCode = "get_pool_data_error"
	GetPoolInfoError        serverErrors.ServerErrorCode = "get_pool_info_error"
	GetPoolStatsError       serverErrors.ServerErrorCode = "get_pool_stats_error"
	GetPoolNetworkInfoError serverErrors.ServerErrorCode = "get_pool_network_info_error"
	GetPoolSlavesError      serverErrors.ServerErrorCode = "get_pool_slaves_error"
)

func CreateGetPoolDataError(err error) *apiModels.GetBlockchainPoolInternalServerError {
	return &apiModels.GetBlockchainPoolInternalServerError{
		Code:    string(GetPoolDataError),
		Message: err.Error(),
	}
}

func CreateGetPoolInfoError(err error) *apiModels.GetBlockchainPoolInfoInternalServerError {
	return &apiModels.GetBlockchainPoolInfoInternalServerError{
		Code:    string(GetPoolInfoError),
		Message: err.Error(),
	}
}

func CreateGetPoolStatsError(err error) *apiModels.GetBlockchainPoolStatsInternalServerError {
	return &apiModels.GetBlockchainPoolStatsInternalServerError{
		Code:    string(GetPoolStatsError),
		Message: err.Error(),
	}
}

func CreateGetPoolNetworkInfoError(err error) *apiModels.GetBlockchainPoolNetworkInfoInternalServerError {
	return &apiModels.GetBlockchainPoolNetworkInfoInternalServerError{
		Code:    string(GetPoolNetworkInfoError),
		Message: err.Error(),
	}
}

func CreateGetPoolSlavesError(err error) *apiModels.GetBlockchainPoolSlavesInternalServerError {
	return &apiModels.GetBlockchainPoolSlavesInternalServerError{
		Code:    string(GetPoolSlavesError),
		Message: err.Error(),
	}
}
