package poolsErrors

import (
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

const GetPoolsError serverErrors.ServerErrorCode = "get_pools_error"

func CreateGetPoolsError(err error) *apiModels.GetPoolsInternalServerError {
	return &apiModels.GetPoolsInternalServerError{
		Code:    string(GetPoolsError),
		Message: err.Error(),
	}
}
