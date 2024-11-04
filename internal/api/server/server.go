package apiServer

import (
	apiModels "github.com/grandminingpool/pool-api/api/generated"
	serverHandelrs "github.com/grandminingpool/pool-api/internal/common/server/handlers"
)

func CreateServer(h apiModels.Handler) (*apiModels.Server, error) {
	return apiModels.NewServer(
		h,
		apiModels.WithErrorHandler(serverHandelrs.ErrorHandler()),
		apiModels.WithMethodNotAllowed(serverHandelrs.MethodNotAllowedHandler()),
		apiModels.WithNotFound(serverHandelrs.NotFoundHandler()),
	)
}
