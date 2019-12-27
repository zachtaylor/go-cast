package cast

import "errors"

// Error is a basic wrap-aware error
type Error struct {
	Text string
	Err  error
}

// NewError creates an error wrapper, using StringN
func NewError(err error, args ...interface{}) error {
	return &Error{StringN(args...), err}
}

// Error implements error
func (err *Error) Error() string {
	if err.Err == nil {
		return err.Text
	}

	return err.Text + ": " + err.Err.Error()
}

// Unwrap implements error unwrap
func (err *Error) Unwrap() error {
	return err.Err
}

// IsError = errors.Is
func IsError(err, target error) bool {
	return errors.Is(err, target)
}

// ErrorAs = errors.As
func ErrorAs(err error, target interface{}) bool {
	return errors.As(err, target)
}

// UnwrapError = errors.Unwrap
func UnwrapError(err error) error {
	return errors.Unwrap(err)
}
