// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/v2/health"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/dp-search-reindex-api/sdk"
	"sync"
)

var (
	lockClientMockChecker        sync.RWMutex
	lockClientMockHealth         sync.RWMutex
	lockClientMockPostJob        sync.RWMutex
	lockClientMockPostTasksCount sync.RWMutex
	lockClientMockURL            sync.RWMutex
)

// Ensure, that ClientMock does implement Client.
// If this is not the case, regenerate this file with moq.
var _ sdk.Client = &ClientMock{}

// ClientMock is a mock implementation of sdk.Client.
//
//     func TestSomethingThatUsesClient(t *testing.T) {
//
//         // make and configure a mocked sdk.Client
//         mockedClient := &ClientMock{
//             CheckerFunc: func(ctx context.Context, check *healthcheck.CheckState) error {
// 	               panic("mock out the Checker method")
//             },
//             HealthFunc: func() *health.Client {
// 	               panic("mock out the Health method")
//             },
//             PostJobFunc: func(ctx context.Context, headers sdk.Headers) (models.Job, error) {
// 	               panic("mock out the PostJob method")
//             },
//             PostTasksCountFunc: func(ctx context.Context, headers sdk.Headers, jobID string) (models.Task, error) {
// 	               panic("mock out the PostTasksCount method")
//             },
//             URLFunc: func() string {
// 	               panic("mock out the URL method")
//             },
//         }
//
//         // use mockedClient in code that requires sdk.Client
//         // and then make assertions.
//
//     }
type ClientMock struct {
	// CheckerFunc mocks the Checker method.
	CheckerFunc func(ctx context.Context, check *healthcheck.CheckState) error

	// HealthFunc mocks the Health method.
	HealthFunc func() *health.Client

	// PostJobFunc mocks the PostJob method.
	PostJobFunc func(ctx context.Context, headers sdk.Headers) (models.Job, error)

	// PostTasksCountFunc mocks the PostTasksCount method.
	PostTasksCountFunc func(ctx context.Context, headers sdk.Headers, jobID string) (models.Task, error)

	// URLFunc mocks the URL method.
	URLFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Checker holds details about calls to the Checker method.
		Checker []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Check is the check argument value.
			Check *healthcheck.CheckState
		}
		// Health holds details about calls to the Health method.
		Health []struct {
		}
		// PostJob holds details about calls to the PostJob method.
		PostJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Headers is the headers argument value.
			Headers sdk.Headers
		}
		// PostTasksCount holds details about calls to the PostTasksCount method.
		PostTasksCount []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Headers is the headers argument value.
			Headers sdk.Headers
			// JobID is the jobID argument value.
			JobID string
		}
		// URL holds details about calls to the URL method.
		URL []struct {
		}
	}
}

// Checker calls CheckerFunc.
func (mock *ClientMock) Checker(ctx context.Context, check *healthcheck.CheckState) error {
	if mock.CheckerFunc == nil {
		panic("ClientMock.CheckerFunc: method is nil but Client.Checker was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Check *healthcheck.CheckState
	}{
		Ctx:   ctx,
		Check: check,
	}
	lockClientMockChecker.Lock()
	mock.calls.Checker = append(mock.calls.Checker, callInfo)
	lockClientMockChecker.Unlock()
	return mock.CheckerFunc(ctx, check)
}

// CheckerCalls gets all the calls that were made to Checker.
// Check the length with:
//     len(mockedClient.CheckerCalls())
func (mock *ClientMock) CheckerCalls() []struct {
	Ctx   context.Context
	Check *healthcheck.CheckState
} {
	var calls []struct {
		Ctx   context.Context
		Check *healthcheck.CheckState
	}
	lockClientMockChecker.RLock()
	calls = mock.calls.Checker
	lockClientMockChecker.RUnlock()
	return calls
}

// Health calls HealthFunc.
func (mock *ClientMock) Health() *health.Client {
	if mock.HealthFunc == nil {
		panic("ClientMock.HealthFunc: method is nil but Client.Health was just called")
	}
	callInfo := struct {
	}{}
	lockClientMockHealth.Lock()
	mock.calls.Health = append(mock.calls.Health, callInfo)
	lockClientMockHealth.Unlock()
	return mock.HealthFunc()
}

// HealthCalls gets all the calls that were made to Health.
// Check the length with:
//     len(mockedClient.HealthCalls())
func (mock *ClientMock) HealthCalls() []struct {
} {
	var calls []struct {
	}
	lockClientMockHealth.RLock()
	calls = mock.calls.Health
	lockClientMockHealth.RUnlock()
	return calls
}

// PostJob calls PostJobFunc.
func (mock *ClientMock) PostJob(ctx context.Context, headers sdk.Headers) (models.Job, error) {
	if mock.PostJobFunc == nil {
		panic("ClientMock.PostJobFunc: method is nil but Client.PostJob was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Headers sdk.Headers
	}{
		Ctx:     ctx,
		Headers: headers,
	}
	lockClientMockPostJob.Lock()
	mock.calls.PostJob = append(mock.calls.PostJob, callInfo)
	lockClientMockPostJob.Unlock()
	return mock.PostJobFunc(ctx, headers)
}

// PostJobCalls gets all the calls that were made to PostJob.
// Check the length with:
//     len(mockedClient.PostJobCalls())
func (mock *ClientMock) PostJobCalls() []struct {
	Ctx     context.Context
	Headers sdk.Headers
} {
	var calls []struct {
		Ctx     context.Context
		Headers sdk.Headers
	}
	lockClientMockPostJob.RLock()
	calls = mock.calls.PostJob
	lockClientMockPostJob.RUnlock()
	return calls
}

// PostTasksCount calls PostTasksCountFunc.
func (mock *ClientMock) PostTasksCount(ctx context.Context, headers sdk.Headers, jobID string) (models.Task, error) {
	if mock.PostTasksCountFunc == nil {
		panic("ClientMock.PostTasksCountFunc: method is nil but Client.PostTasksCount was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Headers sdk.Headers
		JobID   string
	}{
		Ctx:     ctx,
		Headers: headers,
		JobID:   jobID,
	}
	lockClientMockPostTasksCount.Lock()
	mock.calls.PostTasksCount = append(mock.calls.PostTasksCount, callInfo)
	lockClientMockPostTasksCount.Unlock()
	return mock.PostTasksCountFunc(ctx, headers, jobID)
}

// PostTasksCountCalls gets all the calls that were made to PostTasksCount.
// Check the length with:
//     len(mockedClient.PostTasksCountCalls())
func (mock *ClientMock) PostTasksCountCalls() []struct {
	Ctx     context.Context
	Headers sdk.Headers
	JobID   string
} {
	var calls []struct {
		Ctx     context.Context
		Headers sdk.Headers
		JobID   string
	}
	lockClientMockPostTasksCount.RLock()
	calls = mock.calls.PostTasksCount
	lockClientMockPostTasksCount.RUnlock()
	return calls
}

// URL calls URLFunc.
func (mock *ClientMock) URL() string {
	if mock.URLFunc == nil {
		panic("ClientMock.URLFunc: method is nil but Client.URL was just called")
	}
	callInfo := struct {
	}{}
	lockClientMockURL.Lock()
	mock.calls.URL = append(mock.calls.URL, callInfo)
	lockClientMockURL.Unlock()
	return mock.URLFunc()
}

// URLCalls gets all the calls that were made to URL.
// Check the length with:
//     len(mockedClient.URLCalls())
func (mock *ClientMock) URLCalls() []struct {
} {
	var calls []struct {
	}
	lockClientMockURL.RLock()
	calls = mock.calls.URL
	lockClientMockURL.RUnlock()
	return calls
}
