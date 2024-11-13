package serverHandelrs

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	serverErrors "github.com/grandminingpool/pool-api/internal/common/server/errors"
	"github.com/ogen-go/ogen/ogenerrors"
)

func setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func MethodNotAllowedHandler() func(w http.ResponseWriter, r *http.Request, allowed string) {
	return func(w http.ResponseWriter, r *http.Request, allowed string) {
		methodNotAllowedError := serverErrors.CreateNotAllowedError(serverErrors.MethodNotAllowed, errors.New(fmt.Sprintf("method not allowed, use one of the following methods: %s", allowed)))

		setContentType(w)
		methodNotAllowedError.WriteResponse(w)
	}
}

func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notFoundError := serverErrors.CreateNotFoundError(serverErrors.NotFoundErrorCode, errors.New("resource not found"))

		setContentType(w)
		notFoundError.WriteResponse(w)
	}
}

func ErrorHandler() ogenerrors.ErrorHandler {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
		var se serverErrors.ServerError

		switch t := err.(type) {
		case serverErrors.ServerError:
			se = t
		case *ogenerrors.DecodeParamsError:
		case *ogenerrors.DecodeParamError:
			se = serverErrors.CreateBadRequestError(serverErrors.DecodeParamsError, err)
		case *ogenerrors.DecodeRequestError:
			se = serverErrors.CreateBadRequestError(serverErrors.DecodeRequestErrorCode, err)
		default:
			se = serverErrors.CreateInternalServerError(serverErrors.UnknownErrorCode, err)
		}

		setContentType(w)
		se.WriteResponse(w)
	}
}
