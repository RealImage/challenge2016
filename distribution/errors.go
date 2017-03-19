package distribution

import (
	"fmt"
)

type ApplicationError interface {
	Error() string
	ErrorCode() string
	Message() string
}
type BaseApplicationError struct {
	message   string
	errorCode string
}

func (e *BaseApplicationError) ErrorCode() string {
	return e.errorCode
}
func (e *BaseApplicationError) Message() string {
	return e.message
}
func (e *BaseApplicationError) Error() string {
	return fmt.Sprintf("%v (%v)", e.message, e.errorCode)
}

func OSError(message string) ApplicationError {
	return &BaseApplicationError{
		errorCode: "application.os.error",
		message:   message,
	}
}

func InputError(message string) ApplicationError {
	return &BaseApplicationError{
		errorCode: "application.input.error",
		message:   message,
	}
}

func DistributionScopeError(location string) ApplicationError {
	return &BaseApplicationError{
		errorCode: "distribution.scope.error",
		message:   "Access denied to location: " + location,
	}
}
