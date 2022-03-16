package apierrors

import (
	"errors"
)

// A list of error messages for JobStore API
var (
	ErrEmptyTaskNameProvided = errors.New("task name must not be an empty string")
	ErrInternalServer        = errors.New("internal server error")
	ErrTaskInvalidName       = errors.New("task name is not valid")
	ErrUnableToParseJSON     = errors.New("failed to parse json body")
	ErrUnableToReadMessage   = errors.New("failed to read message body")
)
