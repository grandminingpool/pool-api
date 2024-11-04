package pricesErrors

import (
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
)

const GetPricesError serverErrors.ServerErrorCode = "get_prices_error"

func CreateGetPricesError(err error) *apiModels.GetPricesInternalServerError {
	return &apiModels.GetPricesInternalServerError{
		Code:    string(GetPricesError),
		Message: err.Error(),
	}
}
