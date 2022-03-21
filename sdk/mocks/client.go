// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"context"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"github.com/ONSdigital/dp-search-reindex-api/sdk"
	"sync"
)

var (
	lockClientMockPostJob sync.RWMutex
)

// Ensure, that ClientMock does implement sdk.Client.
// If this is not the case, regenerate this file with moq.
var _ sdk.Client = &ClientMock{}

// ClientMock is a mock implementation of sdk.Client.
//
//     func TestSomethingThatUsesClient(t *testing.T) {
//
//         // make and configure a mocked sdk.Client
//         mockedClient := &ClientMock{
//             PostJobFunc: func(ctx context.Context, headers sdk.Headers) (models.Job, error) {
// 	               panic("mock out the PostJob method")
//             },
//         }
//
//         // use mockedClient in code that requires sdk.Client
//         // and then make assertions.
//
//     }
type ClientMock struct {
	// PostJobFunc mocks the PostJob method.
	PostJobFunc func(ctx context.Context, headers sdk.Headers) (models.Job, error)

	// calls tracks calls to the methods.
	calls struct {
		// PostJob holds details about calls to the PostJob method.
		PostJob []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Headers is the headers argument value.
			Headers sdk.Headers
		}
	}
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