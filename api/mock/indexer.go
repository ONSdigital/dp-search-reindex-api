// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	dpHTTP "github.com/ONSdigital/dp-net/http"
	"github.com/ONSdigital/dp-search-reindex-api/api"
	"io"
	"net/http"
	"sync"
)

// Ensure, that IndexerMock does implement api.Indexer.
// If this is not the case, regenerate this file with moq.
var _ api.Indexer = &IndexerMock{}

// IndexerMock is a mock implementation of api.Indexer.
//
// 	func TestSomethingThatUsesIndexer(t *testing.T) {
//
// 		// make and configure a mocked api.Indexer
// 		mockedIndexer := &IndexerMock{
// 			CreateIndexFunc: func(ctx context.Context, serviceAuthToken string, searchAPISearchURL string, httpClient dpHTTP.Clienter) (*http.Response, error) {
// 				panic("mock out the CreateIndex method")
// 			},
// 			GetIndexNameFromResponseFunc: func(ctx context.Context, body io.ReadCloser) (string, error) {
// 				panic("mock out the GetIndexNameFromResponse method")
// 			},
// 		}
//
// 		// use mockedIndexer in code that requires api.Indexer
// 		// and then make assertions.
//
// 	}
type IndexerMock struct {
	// CreateIndexFunc mocks the CreateIndex method.
	CreateIndexFunc func(ctx context.Context, serviceAuthToken string, searchAPISearchURL string, httpClient dpHTTP.Clienter) (*http.Response, error)

	// GetIndexNameFromResponseFunc mocks the GetIndexNameFromResponse method.
	GetIndexNameFromResponseFunc func(ctx context.Context, body io.ReadCloser) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateIndex holds details about calls to the CreateIndex method.
		CreateIndex []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ServiceAuthToken is the serviceAuthToken argument value.
			ServiceAuthToken string
			// SearchAPISearchURL is the searchAPISearchURL argument value.
			SearchAPISearchURL string
			// HttpClient is the httpClient argument value.
			HttpClient dpHTTP.Clienter
		}
		// GetIndexNameFromResponse holds details about calls to the GetIndexNameFromResponse method.
		GetIndexNameFromResponse []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Body is the body argument value.
			Body io.ReadCloser
		}
	}
	lockCreateIndex              sync.RWMutex
	lockGetIndexNameFromResponse sync.RWMutex
}

// CreateIndex calls CreateIndexFunc.
func (mock *IndexerMock) CreateIndex(ctx context.Context, serviceAuthToken string, searchAPISearchURL string, httpClient dpHTTP.Clienter) (*http.Response, error) {
	if mock.CreateIndexFunc == nil {
		panic("IndexerMock.CreateIndexFunc: method is nil but Indexer.CreateIndex was just called")
	}
	callInfo := struct {
		Ctx                context.Context
		ServiceAuthToken   string
		SearchAPISearchURL string
		HttpClient         dpHTTP.Clienter
	}{
		Ctx:                ctx,
		ServiceAuthToken:   serviceAuthToken,
		SearchAPISearchURL: searchAPISearchURL,
		HttpClient:         httpClient,
	}
	mock.lockCreateIndex.Lock()
	mock.calls.CreateIndex = append(mock.calls.CreateIndex, callInfo)
	mock.lockCreateIndex.Unlock()
	return mock.CreateIndexFunc(ctx, serviceAuthToken, searchAPISearchURL, httpClient)
}

// CreateIndexCalls gets all the calls that were made to CreateIndex.
// Check the length with:
//     len(mockedIndexer.CreateIndexCalls())
func (mock *IndexerMock) CreateIndexCalls() []struct {
	Ctx                context.Context
	ServiceAuthToken   string
	SearchAPISearchURL string
	HttpClient         dpHTTP.Clienter
} {
	var calls []struct {
		Ctx                context.Context
		ServiceAuthToken   string
		SearchAPISearchURL string
		HttpClient         dpHTTP.Clienter
	}
	mock.lockCreateIndex.RLock()
	calls = mock.calls.CreateIndex
	mock.lockCreateIndex.RUnlock()
	return calls
}

// GetIndexNameFromResponse calls GetIndexNameFromResponseFunc.
func (mock *IndexerMock) GetIndexNameFromResponse(ctx context.Context, body io.ReadCloser) (string, error) {
	if mock.GetIndexNameFromResponseFunc == nil {
		panic("IndexerMock.GetIndexNameFromResponseFunc: method is nil but Indexer.GetIndexNameFromResponse was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Body io.ReadCloser
	}{
		Ctx:  ctx,
		Body: body,
	}
	mock.lockGetIndexNameFromResponse.Lock()
	mock.calls.GetIndexNameFromResponse = append(mock.calls.GetIndexNameFromResponse, callInfo)
	mock.lockGetIndexNameFromResponse.Unlock()
	return mock.GetIndexNameFromResponseFunc(ctx, body)
}

// GetIndexNameFromResponseCalls gets all the calls that were made to GetIndexNameFromResponse.
// Check the length with:
//     len(mockedIndexer.GetIndexNameFromResponseCalls())
func (mock *IndexerMock) GetIndexNameFromResponseCalls() []struct {
	Ctx  context.Context
	Body io.ReadCloser
} {
	var calls []struct {
		Ctx  context.Context
		Body io.ReadCloser
	}
	mock.lockGetIndexNameFromResponse.RLock()
	calls = mock.calls.GetIndexNameFromResponse
	mock.lockGetIndexNameFromResponse.RUnlock()
	return calls
}
