package poolsErrors

import (
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

const GetPoolsStatsError serverErrors.ServerErrorCode = "get_pools_stats_error"

func CreateGetPoolsStatsError(err error) *apiModels.GetPoolsStatsInternalServerError {
	return &apiModels.GetPoolsStatsInternalServerError{
		Code:    string(GetPoolsStatsError),
		Message: err.Error(),
	}
}
