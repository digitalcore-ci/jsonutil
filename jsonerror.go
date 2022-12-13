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
