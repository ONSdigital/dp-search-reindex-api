package mongo

import (
	"errors"
)

// A list of error messages for Mongo
var (
	ErrJobNotFound           = errors.New("the job id could not be found in the jobs collection")
	ErrEmptyIDProvided       = errors.New("id must not be an empty string")
	ErrEmptyTaskNameProvided = errors.New("task name must not be an empty string")
	ErrDuplicateIDProvided   = errors.New("id must be unique")
	ErrExistingJobInProgress = errors.New("there is an existing job currently in progress")
)
