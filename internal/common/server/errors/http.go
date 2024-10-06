package serverErrors

import (
	"encoding/json"
	"net/http"
)

type ServerError interface {
	error
	WriteResponse(w http.ResponseWriter)
}

type HttpErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type HttpError struct {
	StatusCode int
	Code       ServerErrorCode
	Err        error
}

func (e *HttpError) Error() string {
	return e.Err.Error()
}

func (e *HttpError) WriteResponse(w http.ResponseWriter) {
	w.WriteHeader(e.StatusCode)
	json.NewEncoder(w).Encode(HttpErrorResponse{
		Code:    string(e.Code),
		Message: e.Error(),
	})
}

func CreateHttpError(statusCode int, code ServerErrorCode, err error) *HttpError {
	return &HttpError{
		StatusCode: statusCode,
		Code:       code,
		Err:        err,
	}
}

func CreateBadRequestError(code ServerErrorCode, err error) *HttpError {
	return CreateHttpError(http.StatusBadRequest, code, err)
}

func CreateNotFoundError(code ServerErrorCode, err error) *HttpError {
	return CreateHttpError(http.StatusNotFound, code, err)
}

func CreateForbiddenError(code ServerErrorCode, err error) *HttpError {
	return CreateHttpError(http.StatusForbidden, code, err)
}

func CreateInternalServerError(code ServerErrorCode, err error) *HttpError {
	return CreateHttpError(http.StatusInternalServerError, code, err)
}

func CreateUnauthorizedError(code ServerErrorCode, err error) *HttpError {
	return CreateHttpError(http.StatusUnauthorized, code, err)
}

func CreateNotAllowedError(code ServerErrorCode, err error) *HttpError {
	return CreateHttpError(http.StatusMethodNotAllowed, code, err)
}