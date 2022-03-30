// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
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
// 			PatchJobFunc: func(ctx context.Context, headers sdk.Headers, jobID string, body sdk.PatchOpsList) error {
// 				panic("mock out the PatchJob method")
// 			},
// 			PostJobFunc: func(ctx context.Context, headers sdk.Headers) (models.Job, error) {
// 				panic("mock out the PostJob method")
// 			},
// 		}
//
// 		// use mockedClient in code that requires sdk.Client
// 		// and then make assertions.
//
// 	}
type ClientMock struct {
	// PatchJobFunc mocks the PatchJob method.
	PatchJobFunc func(ctx context.Context, headers sdk.Headers, jobID string, body sdk.PatchOpsList) error

	// PostJobFunc mocks the PostJob method.
	PostJobFunc func(ctx context.Context, headers sdk.Headers) (models.Job, error)

	// calls tracks calls to the methods.
	calls struct {
		// PatchJob holds details about calls to the PatchJob method.
		PatchJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Headers is the headers argument value.
			Headers sdk.Headers
			// JobID is the jobID argument value.
			JobID string
			// Body is the body argument value.
			Body sdk.PatchOpsList
		}
		// PostJob holds details about calls to the PostJob method.
		PostJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Headers is the headers argument value.
			Headers sdk.Headers
		}
	}
	lockPatchJob sync.RWMutex
	lockPostJob  sync.RWMutex
}

// PatchJob calls PatchJobFunc.
func (mock *ClientMock) PatchJob(ctx context.Context, headers sdk.Headers, jobID string, body sdk.PatchOpsList) error {
	if mock.PatchJobFunc == nil {
		panic("ClientMock.PatchJobFunc: method is nil but Client.PatchJob was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Headers sdk.Headers
		JobID   string
		Body    sdk.PatchOpsList
	}{
		Ctx:     ctx,
		Headers: headers,
		JobID:   jobID,
		Body:    body,
	}
	mock.lockPatchJob.Lock()
	mock.calls.PatchJob = append(mock.calls.PatchJob, callInfo)
	mock.lockPatchJob.Unlock()
	return mock.PatchJobFunc(ctx, headers, jobID, body)
}

// PatchJobCalls gets all the calls that were made to PatchJob.
// Check the length with:
//     len(mockedClient.PatchJobCalls())
func (mock *ClientMock) PatchJobCalls() []struct {
	Ctx     context.Context
	Headers sdk.Headers
	JobID   string
	Body    sdk.PatchOpsList
} {
	var calls []struct {
		Ctx     context.Context
		Headers sdk.Headers
		JobID   string
		Body    sdk.PatchOpsList
	}
	mock.lockPatchJob.RLock()
	calls = mock.calls.PatchJob
	mock.lockPatchJob.RUnlock()
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
	mock.lockPostJob.Lock()
	mock.calls.PostJob = append(mock.calls.PostJob, callInfo)
	mock.lockPostJob.Unlock()
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
	mock.lockPostJob.RLock()
	calls = mock.calls.PostJob
	mock.lockPostJob.RUnlock()
	return calls
}
