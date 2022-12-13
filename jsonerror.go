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
	return Unknown("Internal Server Error")
}

// Unknown is an unexpected error
func Unknown(msg string) *Error {
	return &Error{"UNKNOWN_ERROR", msg}
}

// Forbidden is an error when the client tries to access a resource
// they are not allowed to access.
func Forbidden(msg string) *Error {
	return &Error{"FORBIDDEN", msg}
}

// BadJSON is an error when the client supplies malformed JSON.
func BadJSON(msg string) *Error {
	return &Error{"BAD_JSON", msg}
}

// BadAlias is an error when the client supplies a bad alias.
func BadAlias(msg string) *Error {
	return &Error{"BAD_ALIAS", msg}
}

// NotJSON is an error when the client supplies something that is not JSON
// to a JSON endpoint.
func NotJSON(msg string) *Error {
	return &Error{"NOT_JSON", msg}
}

// NotFound is an error when the client tries to access an unknown resource.
func NotFound(msg string) *Error {
	return &Error{"NOT_FOUND", msg}
}
