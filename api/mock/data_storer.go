// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"errors"
	"sync"

	"github.com/ONSdigital/dp-search-reindex-api/api"
	"github.com/ONSdigital/dp-search-reindex-api/models"
)

// Ensure, that DataStorerMock does implement api.DataStorer.
// If this is not the case, regenerate this file with moq.
var _ api.DataStorer = &DataStorerMock{}

// DataStorerMock is a mock implementation of api.DataStorer.
//
// 	func TestSomethingThatUsesJobStorer(t *testing.T) {
//
// 		// make and configure a mocked api.DataStorer
// 		mockedJobStorer := &DataStorerMock{
// 			AcquireJobLockFunc: func(ctx context.Context, id string) (string, error) {
// 				panic("mock out the AcquireJobLock method")
// 			},
// 			CreateJobFunc: func(ctx context.Context, id string) (models.Job, error) {
// 				panic("mock out the CreateJob method")
// 			},
// 			CreateTaskFunc: func(ctx context.Context, jobID string, taskName string, numDocuments int) (models.Task, error) {
// 				panic("mock out the CreateTask method")
// 			},
// 			GetJobFunc: func(ctx context.Context, id string) (models.Job, error) {
// 				panic("mock out the GetJob method")
// 			},
// 			GetJobsFunc: func(ctx context.Context, offsetParam string, limitParam string) (models.Jobs, error) {
// 				panic("mock out the GetJobs method")
// 			},
// 			GetTaskFunc: func(ctx context.Context, jobID string, taskName string) (models.Task, error) {
// 				panic("mock out the GetTask method")
// 			},
// 			PutNumberOfTasksFunc: func(ctx context.Context, id string, count int) error {
// 				panic("mock out the PutNumberOfTasks method")
// 			},
// 			UnlockJobFunc: func(lockID string) error {
// 				panic("mock out the UnlockJob method")
// 			},
// 		}
//
// 		// use mockedJobStorer in code that requires api.DataStorer
// 		// and then make assertions.
//
// 	}
type DataStorerMock struct {
	// AcquireJobLockFunc mocks the AcquireJobLock method.
	AcquireJobLockFunc func(ctx context.Context, id string) (string, error)

	// CreateJobFunc mocks the CreateJob method.
	CreateJobFunc func(ctx context.Context, id string) (models.Job, error)

	// CreateTaskFunc mocks the CreateTask method.
	CreateTaskFunc func(ctx context.Context, jobID string, taskName string, numDocuments int) (models.Task, error)

	// GetJobFunc mocks the GetJob method.
	GetJobFunc func(ctx context.Context, id string) (models.Job, error)

	// GetJobsFunc mocks the GetJobs method.
	GetJobsFunc func(ctx context.Context, offsetParam string, limitParam string) (models.Jobs, error)

	// GetTaskFunc mocks the GetTask method.
	GetTaskFunc func(ctx context.Context, jobID string, taskName string) (models.Task, error)

	// PutNumberOfTasksFunc mocks the PutNumberOfTasks method.
	PutNumberOfTasksFunc func(ctx context.Context, id string, count int) error

	// UnlockJobFunc mocks the UnlockJob method.
	UnlockJobFunc func(lockID string) error

	// calls tracks calls to the methods.
	calls struct {
		// AcquireJobLock holds details about calls to the AcquireJobLock method.
		AcquireJobLock []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// CreateJob holds details about calls to the CreateJob method.
		CreateJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// CreateTask holds details about calls to the CreateTask method.
		CreateTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// JobID is the jobID argument value.
			JobID string
			// TaskName is the taskName argument value.
			TaskName string
			// NumDocuments is the numDocuments argument value.
			NumDocuments int
		}
		// GetJob holds details about calls to the GetJob method.
		GetJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// GetJobs holds details about calls to the GetJobs method.
		GetJobs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OffsetParam is the offsetParam argument value.
			OffsetParam string
			// LimitParam is the limitParam argument value.
			LimitParam string
		}
		// GetTask holds details about calls to the GetTask method.
		GetTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// JobID is the jobID argument value.
			JobID string
			// TaskName is the taskName argument value.
			TaskName string
		}
		// PutNumberOfTasks holds details about calls to the PutNumberOfTasks method.
		PutNumberOfTasks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
			// Count is the count argument value.
			Count int
		}
		// UnlockJob holds details about calls to the UnlockJob method.
		UnlockJob []struct {
			// LockID is the lockID argument value.
			LockID string
		}
	}
	lockAcquireJobLock   sync.RWMutex
	lockCreateJob        sync.RWMutex
	lockCreateTask       sync.RWMutex
	lockGetJob           sync.RWMutex
	lockGetJobs          sync.RWMutex
	lockGetTask          sync.RWMutex
	lockPutNumberOfTasks sync.RWMutex
	lockUnlockJob        sync.RWMutex
}

// Constants for testing
const (
	notFoundID        = "NOT_FOUND_UUID"
	duplicateID       = "DUPLICATE_UUID"
	jobUpdatedFirstID = "JOB_UPDATED_FIRST_ID"
	jobUpdatedLastID  = "JOB_UPDATED_LAST_ID"
)

// AcquireJobLock calls AcquireJobLockFunc.
func (mock *DataStorerMock) AcquireJobLock(ctx context.Context, id string) (string, error) {
	if mock.AcquireJobLockFunc == nil {
		panic("DataStorerMock.AcquireJobLockFunc: method is nil but DataStorer.AcquireJobLock was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockAcquireJobLock.Lock()
	mock.calls.AcquireJobLock = append(mock.calls.AcquireJobLock, callInfo)
	mock.lockAcquireJobLock.Unlock()
	return mock.AcquireJobLockFunc(ctx, id)
}

// AcquireJobLockCalls gets all the calls that were made to AcquireJobLock.
// Check the length with:
//     len(mockedJobStorer.AcquireJobLockCalls())
func (mock *DataStorerMock) AcquireJobLockCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockAcquireJobLock.RLock()
	calls = mock.calls.AcquireJobLock
	mock.lockAcquireJobLock.RUnlock()
	return calls
}

// CreateJob calls CreateJobFunc.
func (mock *DataStorerMock) CreateJob(ctx context.Context, id string) (models.Job, error) {
	if mock.CreateJobFunc == nil {
		if id == "" {
			return models.Job{}, errors.New("id must not be an empty string")
		} else if id == duplicateID {
			return models.Job{}, errors.New("id must be unique")
		} else {
			return models.NewJob(id)
		}
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockCreateJob.Lock()
	mock.calls.CreateJob = append(mock.calls.CreateJob, callInfo)
	mock.lockCreateJob.Unlock()
	return mock.CreateJobFunc(ctx, id)
}

// CreateJobCalls gets all the calls that were made to CreateJob.
// Check the length with:
//     len(mockedJobStorer.CreateJobCalls())
func (mock *DataStorerMock) CreateJobCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockCreateJob.RLock()
	calls = mock.calls.CreateJob
	mock.lockCreateJob.RUnlock()
	return calls
}

// CreateTask calls CreateTaskFunc.
func (mock *DataStorerMock) CreateTask(ctx context.Context, jobID string, taskName string, numDocuments int) (models.Task, error) {
	if mock.CreateTaskFunc == nil {
		panic("DataStorerMock.CreateTaskFunc: method is nil but DataStorer.CreateTask was just called")
	}
	switch taskName {
	case "zebedee":
		jobID = "UUID1"
	case "dataset-api":
		jobID = "UUID3"
	}
	callInfo := struct {
		Ctx          context.Context
		JobID        string
		TaskName     string
		NumDocuments int
	}{
		Ctx:          ctx,
		JobID:        jobID,
		TaskName:     taskName,
		NumDocuments: numDocuments,
	}
	mock.lockCreateTask.Lock()
	mock.calls.CreateTask = append(mock.calls.CreateTask, callInfo)
	mock.lockCreateTask.Unlock()
	return mock.CreateTaskFunc(ctx, jobID, taskName, numDocuments)
}

// CreateTaskCalls gets all the calls that were made to CreateTask.
// Check the length with:
//     len(mockedJobStorer.CreateTaskCalls())
func (mock *DataStorerMock) CreateTaskCalls() []struct {
	Ctx          context.Context
	JobID        string
	TaskName     string
	NumDocuments int
} {
	var calls []struct {
		Ctx          context.Context
		JobID        string
		TaskName     string
		NumDocuments int
	}
	mock.lockCreateTask.RLock()
	calls = mock.calls.CreateTask
	mock.lockCreateTask.RUnlock()
	return calls
}

// GetJob calls GetJobFunc.
func (mock *DataStorerMock) GetJob(ctx context.Context, id string) (models.Job, error) {
	if mock.GetJobFunc == nil {
		if id == "" {
			return models.Job{}, errors.New("id must not be an empty string")
		} else if id == notFoundID {
			return models.Job{}, errors.New("the jobs collection does not contain the job id entered")
		} else {
			return models.NewJob(id)
		}
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetJob.Lock()
	mock.calls.GetJob = append(mock.calls.GetJob, callInfo)
	mock.lockGetJob.Unlock()
	return mock.GetJobFunc(ctx, id)
}

// GetJobCalls gets all the calls that were made to GetJob.
// Check the length with:
//     len(mockedJobStorer.GetJobCalls())
func (mock *DataStorerMock) GetJobCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGetJob.RLock()
	calls = mock.calls.GetJob
	mock.lockGetJob.RUnlock()
	return calls
}

// GetJobs calls GetJobsFunc.
func (mock *DataStorerMock) GetJobs(ctx context.Context, offsetParam string, limitParam string) (models.Jobs, error) {
	if mock.GetJobsFunc == nil {
		results := models.Jobs{}
		jobs := make([]models.Job, 2)
		jobUpdatedFirst, err := models.NewJob(jobUpdatedFirstID)
		jobs[0] = jobUpdatedFirst
		jobUpdatedLast, err := models.NewJob(jobUpdatedLastID)
		jobs[1] = jobUpdatedLast
		results.JobList = jobs
		return results, err
	}
	callInfo := struct {
		Ctx         context.Context
		OffsetParam string
		LimitParam  string
	}{
		Ctx:         ctx,
		OffsetParam: offsetParam,
		LimitParam:  limitParam,
	}
	mock.lockGetJobs.Lock()
	mock.calls.GetJobs = append(mock.calls.GetJobs, callInfo)
	mock.lockGetJobs.Unlock()
	return mock.GetJobsFunc(ctx, offsetParam, limitParam)
}

// GetJobsCalls gets all the calls that were made to GetJobs.
// Check the length with:
//     len(mockedJobStorer.GetJobsCalls())
func (mock *DataStorerMock) GetJobsCalls() []struct {
	Ctx         context.Context
	OffsetParam string
	LimitParam  string
} {
	var calls []struct {
		Ctx         context.Context
		OffsetParam string
		LimitParam  string
	}
	mock.lockGetJobs.RLock()
	calls = mock.calls.GetJobs
	mock.lockGetJobs.RUnlock()
	return calls
}

// GetTask calls GetTaskFunc.
func (mock *DataStorerMock) GetTask(ctx context.Context, jobID string, taskName string) (models.Task, error) {
	if mock.GetTaskFunc == nil {
		panic("DataStorerMock.GetTaskFunc: method is nil but DataStorer.GetTask was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		JobID    string
		TaskName string
	}{
		Ctx:      ctx,
		JobID:    jobID,
		TaskName: taskName,
	}
	mock.lockGetTask.Lock()
	mock.calls.GetTask = append(mock.calls.GetTask, callInfo)
	mock.lockGetTask.Unlock()
	return mock.GetTaskFunc(ctx, jobID, taskName)
}

// GetTaskCalls gets all the calls that were made to GetTask.
// Check the length with:
//     len(mockedJobStorer.GetTaskCalls())
func (mock *DataStorerMock) GetTaskCalls() []struct {
	Ctx      context.Context
	JobID    string
	TaskName string
} {
	var calls []struct {
		Ctx      context.Context
		JobID    string
		TaskName string
	}
	mock.lockGetTask.RLock()
	calls = mock.calls.GetTask
	mock.lockGetTask.RUnlock()
	return calls
}

// PutNumberOfTasks calls PutNumberOfTasksFunc.
func (mock *DataStorerMock) PutNumberOfTasks(ctx context.Context, id string, count int) error {
	if mock.PutNumberOfTasksFunc == nil {
		panic("DataStorerMock.PutNumberOfTasksFunc: method is nil but DataStorer.PutNumberOfTasks was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		ID    string
		Count int
	}{
		Ctx:   ctx,
		ID:    id,
		Count: count,
	}
	mock.lockPutNumberOfTasks.Lock()
	mock.calls.PutNumberOfTasks = append(mock.calls.PutNumberOfTasks, callInfo)
	mock.lockPutNumberOfTasks.Unlock()
	return mock.PutNumberOfTasksFunc(ctx, id, count)
}

// PutNumberOfTasksCalls gets all the calls that were made to PutNumberOfTasks.
// Check the length with:
//     len(mockedJobStorer.PutNumberOfTasksCalls())
func (mock *DataStorerMock) PutNumberOfTasksCalls() []struct {
	Ctx   context.Context
	ID    string
	Count int
} {
	var calls []struct {
		Ctx   context.Context
		ID    string
		Count int
	}
	mock.lockPutNumberOfTasks.RLock()
	calls = mock.calls.PutNumberOfTasks
	mock.lockPutNumberOfTasks.RUnlock()
	return calls
}

// UnlockJob calls UnlockJobFunc.
func (mock *DataStorerMock) UnlockJob(lockID string) error {
	if mock.UnlockJobFunc == nil {
		return nil
	}
	callInfo := struct {
		LockID string
	}{
		LockID: lockID,
	}
	mock.lockUnlockJob.Lock()
	mock.calls.UnlockJob = append(mock.calls.UnlockJob, callInfo)
	mock.lockUnlockJob.Unlock()
	return mock.UnlockJobFunc(lockID)
}

// UnlockJobCalls gets all the calls that were made to UnlockJob.
// Check the length with:
//     len(mockedJobStorer.UnlockJobCalls())
func (mock *DataStorerMock) UnlockJobCalls() []struct {
	LockID string
} {
	var calls []struct {
		LockID string
	}
	mock.lockUnlockJob.RLock()
	calls = mock.calls.UnlockJob
	mock.lockUnlockJob.RUnlock()
	return calls
}
