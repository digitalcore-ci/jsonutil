package jsonutils

import (
	"fmt"
)

type ErrorCode string

// JSONError represents the "standard error response".
type JSONError struct {
	ErrCode ErrorCode `json:"errcode"`
	Message string    `json:"message"`
}

func (e JSONError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrCode, e.Message)
}

// InternalServerError returns an Unknow JSONError whithout error details
func InternalServerError() *JSONError {
	return Unknown("Internal Server Error")
}

// Unknown is an unexpected error
func Unknown(msg string) *JSONError {
	return &JSONError{"UNKNOWN_ERROR", msg}
}

// Forbidden is an error when the client tries to access a resource
// they are not allowed to access.
func Forbidden(msg string) *JSONError {
	return &JSONError{"FORBIDDEN", msg}
}

// BadJSON is an error when the client supplies malformed JSON.
func BadJSON(msg string) *JSONError {
	return &JSONError{"BAD_JSON", msg}
}

// BadAlias is an error when the client supplies a bad alias.
func BadAlias(msg string) *JSONError {
	return &JSONError{"BAD_ALIAS", msg}
}

// NotJSON is an error when the client supplies something that is not JSON
// to a JSON endpoint.
func NotJSON(msg string) *JSONError {
	return &JSONError{"NOT_JSON", msg}
}

// NotFound is an error when the client tries to access an unknown resource.
func NotFound(msg string) *JSONError {
	return &JSONError{"NOT_FOUND", msg}
}

// MissingArgument is an error when the client tries to access a resource
// without providing an argument that is required.
func MissingArgument(msg string) *JSONError {
	return &JSONError{"MISSING_ARGUMENT", msg}
}

// InvalidArgumentValue is an error when the client tries to provide an
// invalid value for a valid argument
func InvalidArgumentValue(msg string) *JSONError {
	return &JSONError{"INVALID_ARGUMENT_VALUE", msg}
}

// MissingToken is an error when the client tries to access a resource which
// requires authentication without supplying credentials.
func MissingToken(msg string) *JSONError {
	return &JSONError{"MISSING_TOKEN", msg}
}

// UnknownToken is an error when the client tries to access a resource which
// requires authentication and supplies an unrecognised token
func UnknownToken(msg string) *JSONError {
	return &JSONError{"UNKNOWN_TOKEN", msg}
}

// WeakPassword is an error which is returned when the client tries to register
// using a weak password. http://matrix.org/docs/spec/client_server/r0.2.0.html#password-based
func WeakPassword(msg string) *JSONError {
	return &JSONError{"WEAK_PASSWORD", msg}
}

// InvalidUsername is an error returned when the client tries to register an
// invalid username
func InvalidUsername(msg string) *JSONError {
	return &JSONError{"INVALID_USERNAME", msg}
}

// UserInUse is an error returned when the client tries to register an
// username that already exists
func UserInUse(msg string) *JSONError {
	return &JSONError{"USER_IN_USE", msg}
}

// ASExclusive is an error returned when an application service tries to
// register an username that is outside of its registered namespace, or if a
// user attempts to register a username or room alias within an exclusive
// namespace.
func ASExclusive(msg string) *JSONError {
	return &JSONError{"M_EXCLUSIVE", msg}
}

// GuestAccessForbidden is an error which is returned when the client is
// forbidden from accessing a resource as a guest.
func GuestAccessForbidden(msg string) *JSONError {
	return &JSONError{"M_GUEST_ACCESS_FORBIDDEN", msg}
}

// InvalidSignature is an error which is returned when the client tries
// to upload invalid signatures.
func InvalidSignature(msg string) *JSONError {
	return &JSONError{"M_INVALID_SIGNATURE", msg}
}

// InvalidParam is an error that is returned when a parameter was invalid,
// traditionally with cross-signing.
func InvalidParam(msg string) *JSONError {
	return &JSONError{"M_INVALID_PARAM", msg}
}

// MissingParam is an error that is returned when a parameter was incorrect,
// traditionally with cross-signing.
func MissingParam(msg string) *JSONError {
	return &JSONError{"M_MISSING_PARAM", msg}
}
