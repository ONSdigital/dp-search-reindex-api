// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/ONSdigital/dp-search-reindex-api/api"
	"github.com/ONSdigital/dp-search-reindex-api/config"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/dp-search-reindex-api/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

// Ensure, that DataStorerMock does implement api.DataStorer.
// If this is not the case, regenerate this file with moq.
var _ api.DataStorer = &DataStorerMock{}

// DataStorerMock is a mock implementation of api.DataStorer.
//
// 	func TestSomethingThatUsesDataStorer(t *testing.T) {
//
// 		// make and configure a mocked api.DataStorer
// 		mockedDataStorer := &DataStorerMock{
// 			AcquireJobLockFunc: func(ctx context.Context, id string) (string, error) {
// 				panic("mock out the AcquireJobLock method")
// 			},
// 			CheckInProgressJobFunc: func(ctx context.Context, cfg *config.Config) error {
// 				panic("mock out the CheckInProgressJob method")
// 			},
// 			CreateJobFunc: func(ctx context.Context, job models.Job) error {
// 				panic("mock out the CreateJob method")
// 			},
// 			GetJobFunc: func(ctx context.Context, id string) (*models.Job, error) {
// 				panic("mock out the GetJob method")
// 			},
// 			GetJobsFunc: func(ctx context.Context, options mongo.Options) (*models.Jobs, error) {
// 				panic("mock out the GetJobs method")
// 			},
// 			GetTaskFunc: func(ctx context.Context, jobID string, taskName string) (*models.Task, error) {
// 				panic("mock out the GetTask method")
// 			},
// 			GetTasksFunc: func(ctx context.Context, jobID string, options mongo.Options) (*models.Tasks, error) {
// 				panic("mock out the GetTasks method")
// 			},
// 			UnlockJobFunc: func(ctx context.Context, lockID string)  {
// 				panic("mock out the UnlockJob method")
// 			},
// 			UpdateJobFunc: func(ctx context.Context, id string, updates primitive.M) error {
// 				panic("mock out the UpdateJob method")
// 			},
// 			UpsertTaskFunc: func(ctx context.Context, jobID string, taskName string, task models.Task) error {
// 				panic("mock out the UpsertTask method")
// 			},
// 		}
//
// 		// use mockedDataStorer in code that requires api.DataStorer
// 		// and then make assertions.
//
// 	}
type DataStorerMock struct {
	// AcquireJobLockFunc mocks the AcquireJobLock method.
	AcquireJobLockFunc func(ctx context.Context, id string) (string, error)

	// CheckInProgressJobFunc mocks the CheckInProgressJob method.
	CheckInProgressJobFunc func(ctx context.Context, cfg *config.Config) error

	// CreateJobFunc mocks the CreateJob method.
	CreateJobFunc func(ctx context.Context, job models.Job) error

	// GetJobFunc mocks the GetJob method.
	GetJobFunc func(ctx context.Context, id string) (*models.Job, error)

	// GetJobsFunc mocks the GetJobs method.
	GetJobsFunc func(ctx context.Context, options mongo.Options) (*models.Jobs, error)

	// GetTaskFunc mocks the GetTask method.
	GetTaskFunc func(ctx context.Context, jobID string, taskName string) (*models.Task, error)

	// GetTasksFunc mocks the GetTasks method.
	GetTasksFunc func(ctx context.Context, jobID string, options mongo.Options) (*models.Tasks, error)

	// UnlockJobFunc mocks the UnlockJob method.
	UnlockJobFunc func(ctx context.Context, lockID string)

	// UpdateJobFunc mocks the UpdateJob method.
	UpdateJobFunc func(ctx context.Context, id string, updates primitive.M) error

	// UpsertTaskFunc mocks the UpsertTask method.
	UpsertTaskFunc func(ctx context.Context, jobID string, taskName string, task models.Task) error

	// calls tracks calls to the methods.
	calls struct {
		// AcquireJobLock holds details about calls to the AcquireJobLock method.
		AcquireJobLock []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// CheckInProgressJob holds details about calls to the CheckInProgressJob method.
		CheckInProgressJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Cfg is the cfg argument value.
			Cfg *config.Config
		}
		// CreateJob holds details about calls to the CreateJob method.
		CreateJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Job is the job argument value.
			Job models.Job
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
			// Options is the options argument value.
			Options mongo.Options
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
			// JobID is the jobID argument value.
			JobID string
			// Options is the options argument value.
			Options mongo.Options
		}
		// UnlockJob holds details about calls to the UnlockJob method.
		UnlockJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// LockID is the lockID argument value.
			LockID string
		}
		// UpdateJob holds details about calls to the UpdateJob method.
		UpdateJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
			// Updates is the updates argument value.
			Updates primitive.M
		}
		// UpsertTask holds details about calls to the UpsertTask method.
		UpsertTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// JobID is the jobID argument value.
			JobID string
			// TaskName is the taskName argument value.
			TaskName string
			// Task is the task argument value.
			Task models.Task
		}
	}
	lockAcquireJobLock     sync.RWMutex
	lockCheckInProgressJob sync.RWMutex
	lockCreateJob          sync.RWMutex
	lockGetJob             sync.RWMutex
	lockGetJobs            sync.RWMutex
	lockGetTask            sync.RWMutex
	lockGetTasks           sync.RWMutex
	lockUnlockJob          sync.RWMutex
	lockUpdateJob          sync.RWMutex
	lockUpsertTask         sync.RWMutex
}

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
//     len(mockedDataStorer.AcquireJobLockCalls())
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

// CheckInProgressJob calls CheckInProgressJobFunc.
func (mock *DataStorerMock) CheckInProgressJob(ctx context.Context, cfg *config.Config) error {
	if mock.CheckInProgressJobFunc == nil {
		panic("DataStorerMock.CheckInProgressJobFunc: method is nil but DataStorer.CheckInProgressJob was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Cfg *config.Config
	}{
		Ctx: ctx,
		Cfg: cfg,
	}
	mock.lockCheckInProgressJob.Lock()
	mock.calls.CheckInProgressJob = append(mock.calls.CheckInProgressJob, callInfo)
	mock.lockCheckInProgressJob.Unlock()
	return mock.CheckInProgressJobFunc(ctx, cfg)
}

// CheckInProgressJobCalls gets all the calls that were made to CheckInProgressJob.
// Check the length with:
//     len(mockedDataStorer.CheckInProgressJobCalls())
func (mock *DataStorerMock) CheckInProgressJobCalls() []struct {
	Ctx context.Context
	Cfg *config.Config
} {
	var calls []struct {
		Ctx context.Context
		Cfg *config.Config
	}
	mock.lockCheckInProgressJob.RLock()
	calls = mock.calls.CheckInProgressJob
	mock.lockCheckInProgressJob.RUnlock()
	return calls
}

// CreateJob calls CreateJobFunc.
func (mock *DataStorerMock) CreateJob(ctx context.Context, job models.Job) error {
	if mock.CreateJobFunc == nil {
		panic("DataStorerMock.CreateJobFunc: method is nil but DataStorer.CreateJob was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Job models.Job
	}{
		Ctx: ctx,
		Job: job,
	}
	mock.lockCreateJob.Lock()
	mock.calls.CreateJob = append(mock.calls.CreateJob, callInfo)
	mock.lockCreateJob.Unlock()
	return mock.CreateJobFunc(ctx, job)
}

// CreateJobCalls gets all the calls that were made to CreateJob.
// Check the length with:
//     len(mockedDataStorer.CreateJobCalls())
func (mock *DataStorerMock) CreateJobCalls() []struct {
	Ctx context.Context
	Job models.Job
} {
	var calls []struct {
		Ctx context.Context
		Job models.Job
	}
	mock.lockCreateJob.RLock()
	calls = mock.calls.CreateJob
	mock.lockCreateJob.RUnlock()
	return calls
}

// GetJob calls GetJobFunc.
func (mock *DataStorerMock) GetJob(ctx context.Context, id string) (*models.Job, error) {
	if mock.GetJobFunc == nil {
		panic("DataStorerMock.GetJobFunc: method is nil but DataStorer.GetJob was just called")
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
//     len(mockedDataStorer.GetJobCalls())
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
func (mock *DataStorerMock) GetJobs(ctx context.Context, options mongo.Options) (*models.Jobs, error) {
	if mock.GetJobsFunc == nil {
		panic("DataStorerMock.GetJobsFunc: method is nil but DataStorer.GetJobs was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Options mongo.Options
	}{
		Ctx:     ctx,
		Options: options,
	}
	mock.lockGetJobs.Lock()
	mock.calls.GetJobs = append(mock.calls.GetJobs, callInfo)
	mock.lockGetJobs.Unlock()
	return mock.GetJobsFunc(ctx, options)
}

// GetJobsCalls gets all the calls that were made to GetJobs.
// Check the length with:
//     len(mockedDataStorer.GetJobsCalls())
func (mock *DataStorerMock) GetJobsCalls() []struct {
	Ctx     context.Context
	Options mongo.Options
} {
	var calls []struct {
		Ctx     context.Context
		Options mongo.Options
	}
	mock.lockGetJobs.RLock()
	calls = mock.calls.GetJobs
	mock.lockGetJobs.RUnlock()
	return calls
}

// GetTask calls GetTaskFunc.
func (mock *DataStorerMock) GetTask(ctx context.Context, jobID string, taskName string) (*models.Task, error) {
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
//     len(mockedDataStorer.GetTaskCalls())
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

// GetTasks calls GetTasksFunc.
func (mock *DataStorerMock) GetTasks(ctx context.Context, jobID string, options mongo.Options) (*models.Tasks, error) {
	if mock.GetTasksFunc == nil {
		panic("DataStorerMock.GetTasksFunc: method is nil but DataStorer.GetTasks was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		JobID   string
		Options mongo.Options
	}{
		Ctx:     ctx,
		JobID:   jobID,
		Options: options,
	}
	mock.lockGetTasks.Lock()
	mock.calls.GetTasks = append(mock.calls.GetTasks, callInfo)
	mock.lockGetTasks.Unlock()
	return mock.GetTasksFunc(ctx, jobID, options)
}

// GetTasksCalls gets all the calls that were made to GetTasks.
// Check the length with:
//     len(mockedDataStorer.GetTasksCalls())
func (mock *DataStorerMock) GetTasksCalls() []struct {
	Ctx     context.Context
	JobID   string
	Options mongo.Options
} {
	var calls []struct {
		Ctx     context.Context
		JobID   string
		Options mongo.Options
	}
	mock.lockGetTasks.RLock()
	calls = mock.calls.GetTasks
	mock.lockGetTasks.RUnlock()
	return calls
}

// UnlockJob calls UnlockJobFunc.
func (mock *DataStorerMock) UnlockJob(ctx context.Context, lockID string) {
	if mock.UnlockJobFunc == nil {
		panic("DataStorerMock.UnlockJobFunc: method is nil but DataStorer.UnlockJob was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		LockID string
	}{
		Ctx:    ctx,
		LockID: lockID,
	}
	mock.lockUnlockJob.Lock()
	mock.calls.UnlockJob = append(mock.calls.UnlockJob, callInfo)
	mock.lockUnlockJob.Unlock()
	mock.UnlockJobFunc(ctx, lockID)
}

// UnlockJobCalls gets all the calls that were made to UnlockJob.
// Check the length with:
//     len(mockedDataStorer.UnlockJobCalls())
func (mock *DataStorerMock) UnlockJobCalls() []struct {
	Ctx    context.Context
	LockID string
} {
	var calls []struct {
		Ctx    context.Context
		LockID string
	}
	mock.lockUnlockJob.RLock()
	calls = mock.calls.UnlockJob
	mock.lockUnlockJob.RUnlock()
	return calls
}

// UpdateJob calls UpdateJobFunc.
func (mock *DataStorerMock) UpdateJob(ctx context.Context, id string, updates primitive.M) error {
	if mock.UpdateJobFunc == nil {
		panic("DataStorerMock.UpdateJobFunc: method is nil but DataStorer.UpdateJob was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		ID      string
		Updates primitive.M
	}{
		Ctx:     ctx,
		ID:      id,
		Updates: updates,
	}
	mock.lockUpdateJob.Lock()
	mock.calls.UpdateJob = append(mock.calls.UpdateJob, callInfo)
	mock.lockUpdateJob.Unlock()
	return mock.UpdateJobFunc(ctx, id, updates)
}

// UpdateJobCalls gets all the calls that were made to UpdateJob.
// Check the length with:
//     len(mockedDataStorer.UpdateJobCalls())
func (mock *DataStorerMock) UpdateJobCalls() []struct {
	Ctx     context.Context
	ID      string
	Updates primitive.M
} {
	var calls []struct {
		Ctx     context.Context
		ID      string
		Updates primitive.M
	}
	mock.lockUpdateJob.RLock()
	calls = mock.calls.UpdateJob
	mock.lockUpdateJob.RUnlock()
	return calls
}

// UpsertTask calls UpsertTaskFunc.
func (mock *DataStorerMock) UpsertTask(ctx context.Context, jobID string, taskName string, task models.Task) error {
	if mock.UpsertTaskFunc == nil {
		panic("DataStorerMock.UpsertTaskFunc: method is nil but DataStorer.UpsertTask was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		JobID    string
		TaskName string
		Task     models.Task
	}{
		Ctx:      ctx,
		JobID:    jobID,
		TaskName: taskName,
		Task:     task,
	}
	mock.lockUpsertTask.Lock()
	mock.calls.UpsertTask = append(mock.calls.UpsertTask, callInfo)
	mock.lockUpsertTask.Unlock()
	return mock.UpsertTaskFunc(ctx, jobID, taskName, task)
}

// UpsertTaskCalls gets all the calls that were made to UpsertTask.
// Check the length with:
//     len(mockedDataStorer.UpsertTaskCalls())
func (mock *DataStorerMock) UpsertTaskCalls() []struct {
	Ctx      context.Context
	JobID    string
	TaskName string
	Task     models.Task
} {
	var calls []struct {
		Ctx      context.Context
		JobID    string
		TaskName string
		Task     models.Task
	}
	mock.lockUpsertTask.RLock()
	calls = mock.calls.UpsertTask
	mock.lockUpsertTask.RUnlock()
	return calls
}
