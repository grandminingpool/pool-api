package chartsErrors

import (
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

const (
	GetPoolStatsChartError            serverErrors.ServerErrorCode = "get_pool_stats_chart_error"
	GetPoolDifficultiesChartError     serverErrors.ServerErrorCode = "get_pool_difficulties_chart_error"
	GetRoundsChartError               serverErrors.ServerErrorCode = "get_rounds_chart_error"
	GetMinerHashratesChartError       serverErrors.ServerErrorCode = "get_miner_hashrates_chart_error"
	GetMinerWorkerHashratesChartError serverErrors.ServerErrorCode = "get_miner_worker_hashrates_chart_error"
	GetMinerSharesChartError          serverErrors.ServerErrorCode = "get_miner_shares_chart_error"
	GetMinerWorkerSharesChartError    serverErrors.ServerErrorCode = "get_miner_worker_shares_chart_error"
)

func CreateGetPoolStatsChartError(err error) *apiModels.GetBlockchainPoolStatsChartInternalServerError {
	return &apiModels.GetBlockchainPoolStatsChartInternalServerError{
		Code:    string(GetPoolStatsChartError),
		Message: err.Error(),
	}
}

func CreateGetPoolDifficultiesChartError(err error) *apiModels.GetBlockchainPoolDifficultiesChartInternalServerError {
	return &apiModels.GetBlockchainPoolDifficultiesChartInternalServerError{
		Code:    string(GetPoolDifficultiesChartError),
		Message: err.Error(),
	}
}

func CreateGetRoundsChartError(err error) *apiModels.GetBlockchainRoundsChartInternalServerError {
	return &apiModels.GetBlockchainRoundsChartInternalServerError{
		Code:    string(GetRoundsChartError),
		Message: err.Error(),
	}
}

func CreateGetMinerHashratesChartError(err error) *apiModels.GetBlockchainMinerHashratesChartInternalServerError {
	return &apiModels.GetBlockchainMinerHashratesChartInternalServerError{
		Code:    string(GetMinerHashratesChartError),
		Message: err.Error(),
	}
}

func CreateGetMinerWorkerHashratesChartError(err error) *apiModels.GetBlockchainMinerWorkerHashratesChartInternalServerError {
	return &apiModels.GetBlockchainMinerWorkerHashratesChartInternalServerError{
		Code:    string(GetMinerWorkerHashratesChartError),
		Message: err.Error(),
	}
}

func CreateGetMinerSharesChartError(err error) *apiModels.GetBlockchainMinerSharesChartInternalServerError {
	return &apiModels.GetBlockchainMinerSharesChartInternalServerError{
		Code:    string(GetMinerSharesChartError),
		Message: err.Error(),
	}
}

func CreateGetMinerWorkerSharesChartError(err error) *apiModels.GetBlockchainMinerWorkerSharesChartInternalServerError {
	return &apiModels.GetBlockchainMinerWorkerSharesChartInternalServerError{
		Code:    string(GetMinerWorkerSharesChartError),
		Message: err.Error(),
	}
}
