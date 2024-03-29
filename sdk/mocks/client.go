// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	healthcheck "github.com/ONSdigital/dp-api-clients-go/v2/health"
	health "github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/dp-search-reindex-api/sdk"
	"sync"
)

// Ensure, that ClientMock does implement sdk.Client.
// If this is not the case, regenerate this file with moq.
var _ sdk.Client = &ClientMock{}

// ClientMock is a mock implementation of sdk.Client.
//
// 	func TestSomethingThatUsesClient(t *testing.T) {
//
// 		// make and configure a mocked sdk.Client
// 		mockedClient := &ClientMock{
// 			CheckerFunc: func(ctx context.Context, check *health.CheckState) error {
// 				panic("mock out the Checker method")
// 			},
// 			GetJobFunc: func(ctx context.Context, reqheader sdk.Headers, jobID string) (*sdk.RespHeaders, *models.Job, error) {
// 				panic("mock out the GetJob method")
// 			},
// 			GetJobsFunc: func(ctx context.Context, reqheader sdk.Headers, options sdk.Options) (*sdk.RespHeaders, *models.Jobs, error) {
// 				panic("mock out the GetJobs method")
// 			},
// 			GetTaskFunc: func(ctx context.Context, reqHeaders sdk.Headers, jobID string, taskName string) (*sdk.RespHeaders, *models.Task, error) {
// 				panic("mock out the GetTask method")
// 			},
// 			GetTasksFunc: func(ctx context.Context, reqHeaders sdk.Headers, jobID string) (*sdk.RespHeaders, *models.Tasks, error) {
// 				panic("mock out the GetTasks method")
// 			},
// 			HealthFunc: func() *healthcheck.Client {
// 				panic("mock out the Health method")
// 			},
// 			PatchJobFunc: func(ctx context.Context, reqHeaders sdk.Headers, jobID string, patchList []sdk.PatchOperation) (*sdk.RespHeaders, error) {
// 				panic("mock out the PatchJob method")
// 			},
// 			PostJobFunc: func(ctx context.Context, reqHeaders sdk.Headers) (*sdk.RespHeaders, *models.Job, error) {
// 				panic("mock out the PostJob method")
// 			},
// 			PostTaskFunc: func(ctx context.Context, reqHeaders sdk.Headers, jobID string, taskToCreate models.TaskToCreate) (*sdk.RespHeaders, *models.Task, error) {
// 				panic("mock out the PostTask method")
// 			},
// 			PutJobNumberOfTasksFunc: func(ctx context.Context, reqHeaders sdk.Headers, jobID string, numTasks string) (*sdk.RespHeaders, error) {
// 				panic("mock out the PutJobNumberOfTasks method")
// 			},
// 			PutTaskNumberOfDocsFunc: func(ctx context.Context, reqHeaders sdk.Headers, jobID string, taskName string, docCount string) (*sdk.RespHeaders, error) {
// 				panic("mock out the PutTaskNumberOfDocs method")
// 			},
// 			URLFunc: func() string {
// 				panic("mock out the URL method")
// 			},
// 		}
//
// 		// use mockedClient in code that requires sdk.Client
// 		// and then make assertions.
//
// 	}
type ClientMock struct {
	// CheckerFunc mocks the Checker method.
	CheckerFunc func(ctx context.Context, check *health.CheckState) error

	// GetJobFunc mocks the GetJob method.
	GetJobFunc func(ctx context.Context, reqheader sdk.Headers, jobID string) (*sdk.RespHeaders, *models.Job, error)

	// GetJobsFunc mocks the GetJobs method.
	GetJobsFunc func(ctx context.Context, reqheader sdk.Headers, options sdk.Options) (*sdk.RespHeaders, *models.Jobs, error)

	// GetTaskFunc mocks the GetTask method.
	GetTaskFunc func(ctx context.Context, reqHeaders sdk.Headers, jobID string, taskName string) (*sdk.RespHeaders, *models.Task, error)

	// GetTasksFunc mocks the GetTasks method.
	GetTasksFunc func(ctx context.Context, reqHeaders sdk.Headers, jobID string) (*sdk.RespHeaders, *models.Tasks, error)

	// HealthFunc mocks the Health method.
	HealthFunc func() *healthcheck.Client

	// PatchJobFunc mocks the PatchJob method.
	PatchJobFunc func(ctx context.Context, reqHeaders sdk.Headers, jobID string, patchList []sdk.PatchOperation) (*sdk.RespHeaders, error)

	// PostJobFunc mocks the PostJob method.
	PostJobFunc func(ctx context.Context, reqHeaders sdk.Headers) (*sdk.RespHeaders, *models.Job, error)

	// PostTaskFunc mocks the PostTask method.
	PostTaskFunc func(ctx context.Context, reqHeaders sdk.Headers, jobID string, taskToCreate models.TaskToCreate) (*sdk.RespHeaders, *models.Task, error)

	// PutJobNumberOfTasksFunc mocks the PutJobNumberOfTasks method.
	PutJobNumberOfTasksFunc func(ctx context.Context, reqHeaders sdk.Headers, jobID string, numTasks string) (*sdk.RespHeaders, error)

	// PutTaskNumberOfDocsFunc mocks the PutTaskNumberOfDocs method.
	PutTaskNumberOfDocsFunc func(ctx context.Context, reqHeaders sdk.Headers, jobID string, taskName string, docCount string) (*sdk.RespHeaders, error)

	// URLFunc mocks the URL method.
	URLFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Checker holds details about calls to the Checker method.
		Checker []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Check is the check argument value.
			Check *health.CheckState
		}
		// GetJob holds details about calls to the GetJob method.
		GetJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Reqheader is the reqheader argument value.
			Reqheader sdk.Headers
			// JobID is the jobID argument value.
			JobID string
		}
		// GetJobs holds details about calls to the GetJobs method.
		GetJobs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Reqheader is the reqheader argument value.
			Reqheader sdk.Headers
			// Options is the options argument value.
			Options sdk.Options
		}
		// GetTask holds details about calls to the GetTask method.
		GetTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReqHeaders is the reqHeaders argument value.
			ReqHeaders sdk.Headers
			// JobID is the jobID argument value.
			JobID string
			// TaskName is the taskName argument value.
			TaskName string
		}
		// GetTasks holds details about calls to the GetTasks method.
		GetTasks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReqHeaders is the reqHeaders argument value.
			ReqHeaders sdk.Headers
			// JobID is the jobID argument value.
			JobID string
		}
		// Health holds details about calls to the Health method.
		Health []struct {
		}
		// PatchJob holds details about calls to the PatchJob method.
		PatchJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReqHeaders is the reqHeaders argument value.
			ReqHeaders sdk.Headers
			// JobID is the jobID argument value.
			JobID string
			// PatchList is the patchList argument value.
			PatchList []sdk.PatchOperation
		}
		// PostJob holds details about calls to the PostJob method.
		PostJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReqHeaders is the reqHeaders argument value.
			ReqHeaders sdk.Headers
		}
		// PostTask holds details about calls to the PostTask method.
		PostTask []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReqHeaders is the reqHeaders argument value.
			ReqHeaders sdk.Headers
			// JobID is the jobID argument value.
			JobID string
			// TaskToCreate is the taskToCreate argument value.
			TaskToCreate models.TaskToCreate
		}
		// PutJobNumberOfTasks holds details about calls to the PutJobNumberOfTasks method.
		PutJobNumberOfTasks []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReqHeaders is the reqHeaders argument value.
			ReqHeaders sdk.Headers
			// JobID is the jobID argument value.
			JobID string
			// NumTasks is the numTasks argument value.
			NumTasks string
		}
		// PutTaskNumberOfDocs holds details about calls to the PutTaskNumberOfDocs method.
		PutTaskNumberOfDocs []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ReqHeaders is the reqHeaders argument value.
			ReqHeaders sdk.Headers
			// JobID is the jobID argument value.
			JobID string
			// TaskName is the taskName argument value.
			TaskName string
			// DocCount is the docCount argument value.
			DocCount string
		}
		// URL holds details about calls to the URL method.
		URL []struct {
		}
	}
	lockChecker             sync.RWMutex
	lockGetJob              sync.RWMutex
	lockGetJobs             sync.RWMutex
	lockGetTask             sync.RWMutex
	lockGetTasks            sync.RWMutex
	lockHealth              sync.RWMutex
	lockPatchJob            sync.RWMutex
	lockPostJob             sync.RWMutex
	lockPostTask            sync.RWMutex
	lockPutJobNumberOfTasks sync.RWMutex
	lockPutTaskNumberOfDocs sync.RWMutex
	lockURL                 sync.RWMutex
}

// Checker calls CheckerFunc.
func (mock *ClientMock) Checker(ctx context.Context, check *health.CheckState) error {
	if mock.CheckerFunc == nil {
		panic("ClientMock.CheckerFunc: method is nil but Client.Checker was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Check *health.CheckState
	}{
		Ctx:   ctx,
		Check: check,
	}
	mock.lockChecker.Lock()
	mock.calls.Checker = append(mock.calls.Checker, callInfo)
	mock.lockChecker.Unlock()
	return mock.CheckerFunc(ctx, check)
}

// CheckerCalls gets all the calls that were made to Checker.
// Check the length with:
//     len(mockedClient.CheckerCalls())
func (mock *ClientMock) CheckerCalls() []struct {
	Ctx   context.Context
	Check *health.CheckState
} {
	var calls []struct {
		Ctx   context.Context
		Check *health.CheckState
	}
	mock.lockChecker.RLock()
	calls = mock.calls.Checker
	mock.lockChecker.RUnlock()
	return calls
}

// GetJob calls GetJobFunc.
func (mock *ClientMock) GetJob(ctx context.Context, reqheader sdk.Headers, jobID string) (*sdk.RespHeaders, *models.Job, error) {
	if mock.GetJobFunc == nil {
		panic("ClientMock.GetJobFunc: method is nil but Client.GetJob was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Reqheader sdk.Headers
		JobID     string
	}{
		Ctx:       ctx,
		Reqheader: reqheader,
		JobID:     jobID,
	}
	mock.lockGetJob.Lock()
	mock.calls.GetJob = append(mock.calls.GetJob, callInfo)
	mock.lockGetJob.Unlock()
	return mock.GetJobFunc(ctx, reqheader, jobID)
}

// GetJobCalls gets all the calls that were made to GetJob.
// Check the length with:
//     len(mockedClient.GetJobCalls())
func (mock *ClientMock) GetJobCalls() []struct {
	Ctx       context.Context
	Reqheader sdk.Headers
	JobID     string
} {
	var calls []struct {
		Ctx       context.Context
		Reqheader sdk.Headers
		JobID     string
	}
	mock.lockGetJob.RLock()
	calls = mock.calls.GetJob
	mock.lockGetJob.RUnlock()
	return calls
}

// GetJobs calls GetJobsFunc.
func (mock *ClientMock) GetJobs(ctx context.Context, reqheader sdk.Headers, options sdk.Options) (*sdk.RespHeaders, *models.Jobs, error) {
	if mock.GetJobsFunc == nil {
		panic("ClientMock.GetJobsFunc: method is nil but Client.GetJobs was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Reqheader sdk.Headers
		Options   sdk.Options
	}{
		Ctx:       ctx,
		Reqheader: reqheader,
		Options:   options,
	}
	mock.lockGetJobs.Lock()
	mock.calls.GetJobs = append(mock.calls.GetJobs, callInfo)
	mock.lockGetJobs.Unlock()
	return mock.GetJobsFunc(ctx, reqheader, options)
}

// GetJobsCalls gets all the calls that were made to GetJobs.
// Check the length with:
//     len(mockedClient.GetJobsCalls())
func (mock *ClientMock) GetJobsCalls() []struct {
	Ctx       context.Context
	Reqheader sdk.Headers
	Options   sdk.Options
} {
	var calls []struct {
		Ctx       context.Context
		Reqheader sdk.Headers
		Options   sdk.Options
	}
	mock.lockGetJobs.RLock()
	calls = mock.calls.GetJobs
	mock.lockGetJobs.RUnlock()
	return calls
}

// GetTask calls GetTaskFunc.
func (mock *ClientMock) GetTask(ctx context.Context, reqHeaders sdk.Headers, jobID string, taskName string) (*sdk.RespHeaders, *models.Task, error) {
	if mock.GetTaskFunc == nil {
		panic("ClientMock.GetTaskFunc: method is nil but Client.GetTask was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
		JobID      string
		TaskName   string
	}{
		Ctx:        ctx,
		ReqHeaders: reqHeaders,
		JobID:      jobID,
		TaskName:   taskName,
	}
	mock.lockGetTask.Lock()
	mock.calls.GetTask = append(mock.calls.GetTask, callInfo)
	mock.lockGetTask.Unlock()
	return mock.GetTaskFunc(ctx, reqHeaders, jobID, taskName)
}

// GetTaskCalls gets all the calls that were made to GetTask.
// Check the length with:
//     len(mockedClient.GetTaskCalls())
func (mock *ClientMock) GetTaskCalls() []struct {
	Ctx        context.Context
	ReqHeaders sdk.Headers
	JobID      string
	TaskName   string
} {
	var calls []struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
		JobID      string
		TaskName   string
	}
	mock.lockGetTask.RLock()
	calls = mock.calls.GetTask
	mock.lockGetTask.RUnlock()
	return calls
}

// GetTasks calls GetTasksFunc.
func (mock *ClientMock) GetTasks(ctx context.Context, reqHeaders sdk.Headers, jobID string) (*sdk.RespHeaders, *models.Tasks, error) {
	if mock.GetTasksFunc == nil {
		panic("ClientMock.GetTasksFunc: method is nil but Client.GetTasks was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
		JobID      string
	}{
		Ctx:        ctx,
		ReqHeaders: reqHeaders,
		JobID:      jobID,
	}
	mock.lockGetTasks.Lock()
	mock.calls.GetTasks = append(mock.calls.GetTasks, callInfo)
	mock.lockGetTasks.Unlock()
	return mock.GetTasksFunc(ctx, reqHeaders, jobID)
}

// GetTasksCalls gets all the calls that were made to GetTasks.
// Check the length with:
//     len(mockedClient.GetTasksCalls())
func (mock *ClientMock) GetTasksCalls() []struct {
	Ctx        context.Context
	ReqHeaders sdk.Headers
	JobID      string
} {
	var calls []struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
		JobID      string
	}
	mock.lockGetTasks.RLock()
	calls = mock.calls.GetTasks
	mock.lockGetTasks.RUnlock()
	return calls
}

// Health calls HealthFunc.
func (mock *ClientMock) Health() *healthcheck.Client {
	if mock.HealthFunc == nil {
		panic("ClientMock.HealthFunc: method is nil but Client.Health was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHealth.Lock()
	mock.calls.Health = append(mock.calls.Health, callInfo)
	mock.lockHealth.Unlock()
	return mock.HealthFunc()
}

// HealthCalls gets all the calls that were made to Health.
// Check the length with:
//     len(mockedClient.HealthCalls())
func (mock *ClientMock) HealthCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHealth.RLock()
	calls = mock.calls.Health
	mock.lockHealth.RUnlock()
	return calls
}

// PatchJob calls PatchJobFunc.
func (mock *ClientMock) PatchJob(ctx context.Context, reqHeaders sdk.Headers, jobID string, patchList []sdk.PatchOperation) (*sdk.RespHeaders, error) {
	if mock.PatchJobFunc == nil {
		panic("ClientMock.PatchJobFunc: method is nil but Client.PatchJob was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
		JobID      string
		PatchList  []sdk.PatchOperation
	}{
		Ctx:        ctx,
		ReqHeaders: reqHeaders,
		JobID:      jobID,
		PatchList:  patchList,
	}
	mock.lockPatchJob.Lock()
	mock.calls.PatchJob = append(mock.calls.PatchJob, callInfo)
	mock.lockPatchJob.Unlock()
	return mock.PatchJobFunc(ctx, reqHeaders, jobID, patchList)
}

// PatchJobCalls gets all the calls that were made to PatchJob.
// Check the length with:
//     len(mockedClient.PatchJobCalls())
func (mock *ClientMock) PatchJobCalls() []struct {
	Ctx        context.Context
	ReqHeaders sdk.Headers
	JobID      string
	PatchList  []sdk.PatchOperation
} {
	var calls []struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
		JobID      string
		PatchList  []sdk.PatchOperation
	}
	mock.lockPatchJob.RLock()
	calls = mock.calls.PatchJob
	mock.lockPatchJob.RUnlock()
	return calls
}

// PostJob calls PostJobFunc.
func (mock *ClientMock) PostJob(ctx context.Context, reqHeaders sdk.Headers) (*sdk.RespHeaders, *models.Job, error) {
	if mock.PostJobFunc == nil {
		panic("ClientMock.PostJobFunc: method is nil but Client.PostJob was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
	}{
		Ctx:        ctx,
		ReqHeaders: reqHeaders,
	}
	mock.lockPostJob.Lock()
	mock.calls.PostJob = append(mock.calls.PostJob, callInfo)
	mock.lockPostJob.Unlock()
	return mock.PostJobFunc(ctx, reqHeaders)
}

// PostJobCalls gets all the calls that were made to PostJob.
// Check the length with:
//     len(mockedClient.PostJobCalls())
func (mock *ClientMock) PostJobCalls() []struct {
	Ctx        context.Context
	ReqHeaders sdk.Headers
} {
	var calls []struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
	}
	mock.lockPostJob.RLock()
	calls = mock.calls.PostJob
	mock.lockPostJob.RUnlock()
	return calls
}

// PostTask calls PostTaskFunc.
func (mock *ClientMock) PostTask(ctx context.Context, reqHeaders sdk.Headers, jobID string, taskToCreate models.TaskToCreate) (*sdk.RespHeaders, *models.Task, error) {
	if mock.PostTaskFunc == nil {
		panic("ClientMock.PostTaskFunc: method is nil but Client.PostTask was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		ReqHeaders   sdk.Headers
		JobID        string
		TaskToCreate models.TaskToCreate
	}{
		Ctx:          ctx,
		ReqHeaders:   reqHeaders,
		JobID:        jobID,
		TaskToCreate: taskToCreate,
	}
	mock.lockPostTask.Lock()
	mock.calls.PostTask = append(mock.calls.PostTask, callInfo)
	mock.lockPostTask.Unlock()
	return mock.PostTaskFunc(ctx, reqHeaders, jobID, taskToCreate)
}

// PostTaskCalls gets all the calls that were made to PostTask.
// Check the length with:
//     len(mockedClient.PostTaskCalls())
func (mock *ClientMock) PostTaskCalls() []struct {
	Ctx          context.Context
	ReqHeaders   sdk.Headers
	JobID        string
	TaskToCreate models.TaskToCreate
} {
	var calls []struct {
		Ctx          context.Context
		ReqHeaders   sdk.Headers
		JobID        string
		TaskToCreate models.TaskToCreate
	}
	mock.lockPostTask.RLock()
	calls = mock.calls.PostTask
	mock.lockPostTask.RUnlock()
	return calls
}

// PutJobNumberOfTasks calls PutJobNumberOfTasksFunc.
func (mock *ClientMock) PutJobNumberOfTasks(ctx context.Context, reqHeaders sdk.Headers, jobID string, numTasks string) (*sdk.RespHeaders, error) {
	if mock.PutJobNumberOfTasksFunc == nil {
		panic("ClientMock.PutJobNumberOfTasksFunc: method is nil but Client.PutJobNumberOfTasks was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
		JobID      string
		NumTasks   string
	}{
		Ctx:        ctx,
		ReqHeaders: reqHeaders,
		JobID:      jobID,
		NumTasks:   numTasks,
	}
	mock.lockPutJobNumberOfTasks.Lock()
	mock.calls.PutJobNumberOfTasks = append(mock.calls.PutJobNumberOfTasks, callInfo)
	mock.lockPutJobNumberOfTasks.Unlock()
	return mock.PutJobNumberOfTasksFunc(ctx, reqHeaders, jobID, numTasks)
}

// PutJobNumberOfTasksCalls gets all the calls that were made to PutJobNumberOfTasks.
// Check the length with:
//     len(mockedClient.PutJobNumberOfTasksCalls())
func (mock *ClientMock) PutJobNumberOfTasksCalls() []struct {
	Ctx        context.Context
	ReqHeaders sdk.Headers
	JobID      string
	NumTasks   string
} {
	var calls []struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
		JobID      string
		NumTasks   string
	}
	mock.lockPutJobNumberOfTasks.RLock()
	calls = mock.calls.PutJobNumberOfTasks
	mock.lockPutJobNumberOfTasks.RUnlock()
	return calls
}

// PutTaskNumberOfDocs calls PutTaskNumberOfDocsFunc.
func (mock *ClientMock) PutTaskNumberOfDocs(ctx context.Context, reqHeaders sdk.Headers, jobID string, taskName string, docCount string) (*sdk.RespHeaders, error) {
	if mock.PutTaskNumberOfDocsFunc == nil {
		panic("ClientMock.PutTaskNumberOfDocsFunc: method is nil but Client.PutTaskNumberOfDocs was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
		JobID      string
		TaskName   string
		DocCount   string
	}{
		Ctx:        ctx,
		ReqHeaders: reqHeaders,
		JobID:      jobID,
		TaskName:   taskName,
		DocCount:   docCount,
	}
	mock.lockPutTaskNumberOfDocs.Lock()
	mock.calls.PutTaskNumberOfDocs = append(mock.calls.PutTaskNumberOfDocs, callInfo)
	mock.lockPutTaskNumberOfDocs.Unlock()
	return mock.PutTaskNumberOfDocsFunc(ctx, reqHeaders, jobID, taskName, docCount)
}

// PutTaskNumberOfDocsCalls gets all the calls that were made to PutTaskNumberOfDocs.
// Check the length with:
//     len(mockedClient.PutTaskNumberOfDocsCalls())
func (mock *ClientMock) PutTaskNumberOfDocsCalls() []struct {
	Ctx        context.Context
	ReqHeaders sdk.Headers
	JobID      string
	TaskName   string
	DocCount   string
} {
	var calls []struct {
		Ctx        context.Context
		ReqHeaders sdk.Headers
		JobID      string
		TaskName   string
		DocCount   string
	}
	mock.lockPutTaskNumberOfDocs.RLock()
	calls = mock.calls.PutTaskNumberOfDocs
	mock.lockPutTaskNumberOfDocs.RUnlock()
	return calls
}

// URL calls URLFunc.
func (mock *ClientMock) URL() string {
	if mock.URLFunc == nil {
		panic("ClientMock.URLFunc: method is nil but Client.URL was just called")
	}
	callInfo := struct {
	}{}
	mock.lockURL.Lock()
	mock.calls.URL = append(mock.calls.URL, callInfo)
	mock.lockURL.Unlock()
	return mock.URLFunc()
}

// URLCalls gets all the calls that were made to URL.
// Check the length with:
//     len(mockedClient.URLCalls())
func (mock *ClientMock) URLCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockURL.RLock()
	calls = mock.calls.URL
	mock.lockURL.RUnlock()
	return calls
}
