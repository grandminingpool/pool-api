package serverErrors

type ServerErrorCode string

const (
	DecodeRequestErrorCode ServerErrorCode = "decode_request_error"
	InternalErrorCode      ServerErrorCode = "internal_error"
	UnknownErrorCode       ServerErrorCode = "unknown_error"
	NotFoundErrorCode      ServerErrorCode = "not_found"
	MethodNotAllowed       ServerErrorCode = "method_not_allowed"
)
