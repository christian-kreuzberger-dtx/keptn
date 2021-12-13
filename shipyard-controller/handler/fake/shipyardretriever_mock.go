// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	keptnv2 "github.com/keptn/go-utils/pkg/lib/v0_2_0"
	"sync"
)

// IShipyardRetrieverMock is a mock implementation of handler.IShipyardRetriever.
//
// 	func TestSomethingThatUsesIShipyardRetriever(t *testing.T) {
//
// 		// make and configure a mocked handler.IShipyardRetriever
// 		mockedIShipyardRetriever := &IShipyardRetrieverMock{
// 			GetCachedShipyardFunc: func(projectName string) (*keptnv2.Shipyard, error) {
// 				panic("mock out the GetCachedShipyard method")
// 			},
// 			GetShipyardFunc: func(projectName string) (*keptnv2.Shipyard, error) {
// 				panic("mock out the GetShipyard method")
// 			},
// 		}
//
// 		// use mockedIShipyardRetriever in code that requires handler.IShipyardRetriever
// 		// and then make assertions.
//
// 	}
type IShipyardRetrieverMock struct {
	// GetCachedShipyardFunc mocks the GetCachedShipyard method.
	GetCachedShipyardFunc func(projectName string) (*keptnv2.Shipyard, error)

	// GetShipyardFunc mocks the GetShipyard method.
	GetShipyardFunc func(projectName string) (*keptnv2.Shipyard, string, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetCachedShipyard holds details about calls to the GetCachedShipyard method.
		GetCachedShipyard []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
		}
		// GetShipyard holds details about calls to the GetShipyard method.
		GetShipyard []struct {
			// ProjectName is the projectName argument value.
			ProjectName string
		}
	}
	lockGetCachedShipyard sync.RWMutex
	lockGetShipyard       sync.RWMutex
}

// GetCachedShipyard calls GetCachedShipyardFunc.
func (mock *IShipyardRetrieverMock) GetCachedShipyard(projectName string) (*keptnv2.Shipyard, error) {
	if mock.GetCachedShipyardFunc == nil {
		panic("IShipyardRetrieverMock.GetCachedShipyardFunc: method is nil but IShipyardRetriever.GetCachedShipyard was just called")
	}
	callInfo := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}
	mock.lockGetCachedShipyard.Lock()
	mock.calls.GetCachedShipyard = append(mock.calls.GetCachedShipyard, callInfo)
	mock.lockGetCachedShipyard.Unlock()
	return mock.GetCachedShipyardFunc(projectName)
}

// GetCachedShipyardCalls gets all the calls that were made to GetCachedShipyard.
// Check the length with:
//     len(mockedIShipyardRetriever.GetCachedShipyardCalls())
func (mock *IShipyardRetrieverMock) GetCachedShipyardCalls() []struct {
	ProjectName string
} {
	var calls []struct {
		ProjectName string
	}
	mock.lockGetCachedShipyard.RLock()
	calls = mock.calls.GetCachedShipyard
	mock.lockGetCachedShipyard.RUnlock()
	return calls
}

// GetShipyard calls GetShipyardFunc.
func (mock *IShipyardRetrieverMock) GetShipyard(projectName string) (*keptnv2.Shipyard, string, error) {
	if mock.GetShipyardFunc == nil {
		panic("IShipyardRetrieverMock.GetShipyardFunc: method is nil but IShipyardRetriever.GetShipyard was just called")
	}
	callInfo := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}
	mock.lockGetShipyard.Lock()
	mock.calls.GetShipyard = append(mock.calls.GetShipyard, callInfo)
	mock.lockGetShipyard.Unlock()
	return mock.GetShipyardFunc(projectName)
}

// GetShipyardCalls gets all the calls that were made to GetShipyard.
// Check the length with:
//     len(mockedIShipyardRetriever.GetShipyardCalls())
func (mock *IShipyardRetrieverMock) GetShipyardCalls() []struct {
	ProjectName string
} {
	var calls []struct {
		ProjectName string
	}
	mock.lockGetShipyard.RLock()
	calls = mock.calls.GetShipyard
	mock.lockGetShipyard.RUnlock()
	return calls
}
