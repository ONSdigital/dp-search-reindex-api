package models

import dprequest "github.com/ONSdigital/dp-net/v2/request"

// JOB STATE - Possible values of a job's state
const (
	JobCreatedState   = "created" // this is the default value of state in a new job
	JobFailedState    = "failed"
	JobCompletedState = "completed"
)

var (
	// ValidJobStates is used for logging available job states
	ValidJobStates = []string{JobCreatedState, JobFailedState, JobCompletedState}

	// ValidJobStatesMap is used for searching available job states
	ValidJobStatesMap = map[string]int{
		JobCreatedState:   1,
		JobFailedState:    1,
		JobCompletedState: 1,
	}
)

// PATCH-OPERATIONS - Possible patch operations available in search-reindex-api
var (
	// ValidPatchOps is used for logging available patch operations
	ValidPatchOps = []string{
		dprequest.OpReplace.String(),
	}

	// ValidPatchOpsMap is used for searching available patch operations
	ValidPatchOpsMap = map[string]bool{
		dprequest.OpReplace.String(): true,
	}
)
