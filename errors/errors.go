package errors

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"-" yaml:"-" xml:"-" toml:"-" bson:"-" mapstructure:"-" jsonb:"-"`
}

// NewError creates a new error instance
func NewError(code int, message string, detail string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

// relize the error interface, log print the detail
// but the detail is not show to the user in the web api response base on gin
// only show code and message to the user
func (e *Error) Error() string {
	return e.Detail
}

// WithDetail set the detail message, use error.Error() to print the detail
// the detail is not show to the user in the web api response base on gin
// only log print the detail, didn't use gin context.String to print this error
func (e *Error) WithDetail(tmp string, args ...interface{}) *Error {

	switch len(args) {
	case 0:
		e.Detail = tmp
	default:
		e.Detail = fmt.Sprintf(tmp, args...)
	}

	return e
}
