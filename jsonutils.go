package jsonutils

type ErrorCode string

const (
	ErrorInvalidInput   ErrorCode = "INVALID_INPUT"
	ErrorInternalError  ErrorCode = "INTERNAL_ERROR"
	ErrorInvalidRequest ErrorCode = "INVALID_REQUEST"
)

type JSONResponse map[string]interface{}

type JSONError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}
