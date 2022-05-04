// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/ONSdigital/dp-search-reindex-api/api"
	"github.com/ONSdigital/dp-search-reindex-api/models"
	"sync"
)

var (
	lockReindexRequestedProducerMockProduceReindexRequested sync.RWMutex
)

// Ensure, that ReindexRequestedProducerMock does implement ReindexRequestedProducer.
// If this is not the case, regenerate this file with moq.
var _ api.ReindexRequestedProducer = &ReindexRequestedProducerMock{}

// ReindexRequestedProducerMock is a mock implementation of api.ReindexRequestedProducer.
//
//     func TestSomethingThatUsesReindexRequestedProducer(t *testing.T) {
//
//         // make and configure a mocked api.ReindexRequestedProducer
//         mockedReindexRequestedProducer := &ReindexRequestedProducerMock{
//             ProduceReindexRequestedFunc: func(ctx context.Context, event models.ReindexRequested) error {
// 	               panic("mock out the ProduceReindexRequested method")
//             },
//         }
//
//         // use mockedReindexRequestedProducer in code that requires api.ReindexRequestedProducer
//         // and then make assertions.
//
//     }
type ReindexRequestedProducerMock struct {
	// ProduceReindexRequestedFunc mocks the ProduceReindexRequested method.
	ProduceReindexRequestedFunc func(ctx context.Context, event models.ReindexRequested) error

	// calls tracks calls to the methods.
	calls struct {
		// ProduceReindexRequested holds details about calls to the ProduceReindexRequested method.
		ProduceReindexRequested []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Event is the event argument value.
			Event models.ReindexRequested
		}
	}
}

// ProduceReindexRequested calls ProduceReindexRequestedFunc.
func (mock *ReindexRequestedProducerMock) ProduceReindexRequested(ctx context.Context, event models.ReindexRequested) error {
	if mock.ProduceReindexRequestedFunc == nil {
		panic("ReindexRequestedProducerMock.ProduceReindexRequestedFunc: method is nil but ReindexRequestedProducer.ProduceReindexRequested was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Event models.ReindexRequested
	}{
		Ctx:   ctx,
		Event: event,
	}
	lockReindexRequestedProducerMockProduceReindexRequested.Lock()
	mock.calls.ProduceReindexRequested = append(mock.calls.ProduceReindexRequested, callInfo)
	lockReindexRequestedProducerMockProduceReindexRequested.Unlock()
	return mock.ProduceReindexRequestedFunc(ctx, event)
}

// ProduceReindexRequestedCalls gets all the calls that were made to ProduceReindexRequested.
// Check the length with:
//     len(mockedReindexRequestedProducer.ProduceReindexRequestedCalls())
func (mock *ReindexRequestedProducerMock) ProduceReindexRequestedCalls() []struct {
	Ctx   context.Context
	Event models.ReindexRequested
} {
	var calls []struct {
		Ctx   context.Context
		Event models.ReindexRequested
	}
	lockReindexRequestedProducerMockProduceReindexRequested.RLock()
	calls = mock.calls.ProduceReindexRequested
	lockReindexRequestedProducerMockProduceReindexRequested.RUnlock()
	return calls
}
