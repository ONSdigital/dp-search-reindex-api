package apierrors

import (
	"errors"
)

// A list of error messages
var (
	ErrConflictWithJobETag   = errors.New("etag does not match with current state of job resource")
	ErrEmptyTaskNameProvided = errors.New("task name must not be an empty string")
	ErrInternalServer        = errors.New("internal server error")
	ErrTaskInvalidName       = errors.New("task name is not valid")
	ErrUnableToReadMessage   = errors.New("failed to read message body")
	ErrUnableToParseJSON     = errors.New("failed to parse json body")
)
