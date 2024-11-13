package serverErrors

type ServerErrorCode string

const (
	DecodeParamsError      ServerErrorCode = "decode_params_error"
	DecodeRequestErrorCode ServerErrorCode = "decode_request_error"
	UnknownErrorCode       ServerErrorCode = "unknown_error"
	NotFoundErrorCode      ServerErrorCode = "not_found"
	MethodNotAllowed       ServerErrorCode = "method_not_allowed"
)
