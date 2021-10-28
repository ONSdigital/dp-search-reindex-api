// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/dp-search-reindex-api/service"
	"sync"
)

// Ensure, that MongoDataStorerMock does implement service.MongoDataStorer.
// If this is not the case, regenerate this file with moq.
var _ service.MongoDataStorer = &MongoDataStorerMock{}

// MongoDataStorerMock is a mock implementation of service.MongoDataStorer.
//
// 	func TestSomethingThatUsesMongoDataStorer(t *testing.T) {
//
// 		// make and configure a mocked service.MongoDataStorer
// 		mockedMongoDataStorer := &MongoDataStorerMock{
// 			AcquireJobLockFunc: func(ctx context.Context, id string) (string, error) {
// 				panic("mock out the AcquireJobLock method")
// 			},
// 			CheckerFunc: func(ctx context.Context, state *healthcheck.CheckState) error {
// 				panic("mock out the Checker method")
// 			},
// 			CloseFunc: func(ctx context.Context) error {
// 				panic("mock out the Close method")
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
// 			GetTasksFunc: func(ctx context.Context, offsetParam string, limitParam string) (models.Tasks, error) {
// 				panic("mock out the GetTasks method")
// 			},
// 			PutNumberOfTasksFunc: func(ctx context.Context, id string, count int) error {
// 				panic("mock out the PutNumberOfTasks method")
// 			},
// 			UnlockJobFunc: func(lockID string) error {
// 				panic("mock out the UnlockJob method")
// 			},
// 		}
//
// 		// use mockedMongoDataStorer in code that requires service.MongoDataStorer
// 		// and then make assertions.
//
// 	}
type MongoDataStorerMock struct {
	// AcquireJobLockFunc mocks the AcquireJobLock method.
	AcquireJobLockFunc func(ctx context.Context, id string) (string, error)

	// CheckerFunc mocks the Checker method.
	CheckerFunc func(ctx context.Context, state *healthcheck.CheckState) error

	// CloseFunc mocks the Close method.
	CloseFunc func(ctx context.Context) error

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

	// GetTasksFunc mocks the GetTasks method.
	GetTasksFunc func(ctx context.Context, offsetParam string, limitParam string, jobID string) (models.Tasks, error)

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
		// Checker holds details about calls to the Checker method.
		Checker []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// State is the state argument value.
			State *healthcheck.CheckState
		}
		// Close holds details about calls to the Close method.
		Close []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
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
		// GetTasks holds details about calls to the GetTasks method.
		GetTasks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OffsetParam is the offsetParam argument value.
			OffsetParam string
			// LimitParam is the limitParam argument value.
			LimitParam string
			// JobID is the jobID argument value.
			JobID string
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
	lockChecker          sync.RWMutex
	lockClose            sync.RWMutex
	lockCreateJob        sync.RWMutex
	lockCreateTask       sync.RWMutex
	lockGetJob           sync.RWMutex
	lockGetJobs          sync.RWMutex
	lockGetTask          sync.RWMutex
	lockGetTasks         sync.RWMutex
	lockPutNumberOfTasks sync.RWMutex
	lockUnlockJob        sync.RWMutex
}

// AcquireJobLock calls AcquireJobLockFunc.
func (mock *MongoDataStorerMock) AcquireJobLock(ctx context.Context, id string) (string, error) {
	if mock.AcquireJobLockFunc == nil {
		panic("MongoDataStorerMock.AcquireJobLockFunc: method is nil but MongoDataStorer.AcquireJobLock was just called")
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
//     len(mockedMongoDataStorer.AcquireJobLockCalls())
func (mock *MongoDataStorerMock) AcquireJobLockCalls() []struct {
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

// Checker calls CheckerFunc.
func (mock *MongoDataStorerMock) Checker(ctx context.Context, state *healthcheck.CheckState) error {
	if mock.CheckerFunc == nil {
		panic("MongoDataStorerMock.CheckerFunc: method is nil but MongoDataStorer.Checker was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		State *healthcheck.CheckState
	}{
		Ctx:   ctx,
		State: state,
	}
	mock.lockChecker.Lock()
	mock.calls.Checker = append(mock.calls.Checker, callInfo)
	mock.lockChecker.Unlock()
	return mock.CheckerFunc(ctx, state)
}

// CheckerCalls gets all the calls that were made to Checker.
// Check the length with:
//     len(mockedMongoDataStorer.CheckerCalls())
func (mock *MongoDataStorerMock) CheckerCalls() []struct {
	Ctx   context.Context
	State *healthcheck.CheckState
} {
	var calls []struct {
		Ctx   context.Context
		State *healthcheck.CheckState
	}
	mock.lockChecker.RLock()
	calls = mock.calls.Checker
	mock.lockChecker.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *MongoDataStorerMock) Close(ctx context.Context) error {
	if mock.CloseFunc == nil {
		panic("MongoDataStorerMock.CloseFunc: method is nil but MongoDataStorer.Close was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	mock.lockClose.Unlock()
	return mock.CloseFunc(ctx)
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedMongoDataStorer.CloseCalls())
func (mock *MongoDataStorerMock) CloseCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockClose.RLock()
	calls = mock.calls.Close
	mock.lockClose.RUnlock()
	return calls
}

// CreateJob calls CreateJobFunc.
func (mock *MongoDataStorerMock) CreateJob(ctx context.Context, id string) (models.Job, error) {
	if mock.CreateJobFunc == nil {
		panic("MongoDataStorerMock.CreateJobFunc: method is nil but MongoDataStorer.CreateJob was just called")
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
//     len(mockedMongoDataStorer.CreateJobCalls())
func (mock *MongoDataStorerMock) CreateJobCalls() []struct {
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
func (mock *MongoDataStorerMock) CreateTask(ctx context.Context, jobID string, taskName string, numDocuments int) (models.Task, error) {
	if mock.CreateTaskFunc == nil {
		panic("MongoDataStorerMock.CreateTaskFunc: method is nil but MongoDataStorer.CreateTask was just called")
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
//     len(mockedMongoDataStorer.CreateTaskCalls())
func (mock *MongoDataStorerMock) CreateTaskCalls() []struct {
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
func (mock *MongoDataStorerMock) GetJob(ctx context.Context, id string) (models.Job, error) {
	if mock.GetJobFunc == nil {
		panic("MongoDataStorerMock.GetJobFunc: method is nil but MongoDataStorer.GetJob was just called")
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
//     len(mockedMongoDataStorer.GetJobCalls())
func (mock *MongoDataStorerMock) GetJobCalls() []struct {
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
func (mock *MongoDataStorerMock) GetJobs(ctx context.Context, offsetParam string, limitParam string) (models.Jobs, error) {
	if mock.GetJobsFunc == nil {
		panic("MongoDataStorerMock.GetJobsFunc: method is nil but MongoDataStorer.GetJobs was just called")
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
//     len(mockedMongoDataStorer.GetJobsCalls())
func (mock *MongoDataStorerMock) GetJobsCalls() []struct {
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
func (mock *MongoDataStorerMock) GetTask(ctx context.Context, jobID string, taskName string) (models.Task, error) {
	if mock.GetTaskFunc == nil {
		panic("MongoDataStorerMock.GetTaskFunc: method is nil but MongoDataStorer.GetTask was just called")
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
//     len(mockedMongoDataStorer.GetTaskCalls())
func (mock *MongoDataStorerMock) GetTaskCalls() []struct {
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

// GetTasks calls GetTasksFunc.
func (mock *MongoDataStorerMock) GetTasks(ctx context.Context, offsetParam string, limitParam string, jobID string) (models.Tasks, error) {
	if mock.GetTasksFunc == nil {
		panic("MongoDataStorerMock.GetTasksFunc: method is nil but MongoDataStorer.GetTasks was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		OffsetParam string
		LimitParam  string
		JobID       string
	}{
		Ctx:         ctx,
		OffsetParam: offsetParam,
		LimitParam:  limitParam,
		JobID:       jobID,
	}
	mock.lockGetTasks.Lock()
	mock.calls.GetTasks = append(mock.calls.GetTasks, callInfo)
	mock.lockGetTasks.Unlock()
	return mock.GetTasksFunc(ctx, offsetParam, limitParam, jobID)
}

// GetTasksCalls gets all the calls that were made to GetTasks.
// Check the length with:
//     len(mockedMongoDataStorer.GetTasksCalls())
func (mock *MongoDataStorerMock) GetTasksCalls() []struct {
	Ctx         context.Context
	OffsetParam string
	LimitParam  string
	JobID       string
} {
	var calls []struct {
		Ctx         context.Context
		OffsetParam string
		LimitParam  string
		JobID       string
	}
	mock.lockGetTasks.RLock()
	calls = mock.calls.GetTasks
	mock.lockGetTasks.RUnlock()
	return calls
}

// PutNumberOfTasks calls PutNumberOfTasksFunc.
func (mock *MongoDataStorerMock) PutNumberOfTasks(ctx context.Context, id string, count int) error {
	if mock.PutNumberOfTasksFunc == nil {
		panic("MongoDataStorerMock.PutNumberOfTasksFunc: method is nil but MongoDataStorer.PutNumberOfTasks was just called")
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
//     len(mockedMongoDataStorer.PutNumberOfTasksCalls())
func (mock *MongoDataStorerMock) PutNumberOfTasksCalls() []struct {
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
func (mock *MongoDataStorerMock) UnlockJob(lockID string) error {
	if mock.UnlockJobFunc == nil {
		panic("MongoDataStorerMock.UnlockJobFunc: method is nil but MongoDataStorer.UnlockJob was just called")
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
//     len(mockedMongoDataStorer.UnlockJobCalls())
func (mock *MongoDataStorerMock) UnlockJobCalls() []struct {
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
