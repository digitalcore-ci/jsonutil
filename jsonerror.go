package jsonutil

import (
	"fmt"
)

type ErrorCode string

// Error represents the "standard error response".
type Error struct {
	ErrCode ErrorCode `json:"errcode"`
	Message string    `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrCode, e.Message)
}

// InternalServerError returns an Unknow Error whithout error details
func InternalServerError() *Error {
	return UnknownError("Internal Server Error")
}

// InvalidContentTypeError is an error when the type of content being uploaded
// is not allowed by the server.
func InvalidContentTypeError(msg string) *Error {
	return &Error{"INVALID_CONTENT_TYPE", msg}
}

// FileNotFoundError is an error when client is requesting a file
// that is not found on the server
func FileNotFoundError(msg string) *Error {
	return &Error{"FILE_NOT_FOUND", msg}
}

// MaxFileSizeExcedeedError is an error when used when
// the API enforces a maximum file size limit and the user
// has attempted to upload a file that exceeds this limit.
// The error message could include the maximum allowed file size
func MaxFileSizeExcedeedError(msg string) *Error {
	return &Error{"MAX_FILE_SIZE_EXCEEDED", msg}
}

// Unknown is an unexpected error
func UnknownError(msg string) *Error {
	return &Error{"UNKNOWN_ERROR", msg}
}

// Forbidden is an error when the client tries to access a resource
// they are not allowed to access.
func ForbiddenError(msg string) *Error {
	return &Error{"FORBIDDEN", msg}
}

// BadJSON is an error when the client supplies malformed JSON.
func BadJSONError(msg string) *Error {
	return &Error{"BAD_JSON", msg}
}

// NotJSON is an error when the client supplies something that is not JSON
// to a JSON endpoint.
func NotJSONError(msg string) *Error {
	return &Error{"NOT_JSON", msg}
}

// NotFound is an error when the client tries to access an unknown resource.
func NotFoundError(msg string) *Error {
	return &Error{"NOT_FOUND", msg}
}
