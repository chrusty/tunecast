// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"context"
	"net/http"
	"sync"

	"github.com/chrusty/tunecast/api"
)

type FakeClientInterface struct {
	GetLibraryStub        func(context.Context) (*http.Response, error)
	getLibraryMutex       sync.RWMutex
	getLibraryArgsForCall []struct {
		arg1 context.Context
	}
	getLibraryReturns struct {
		result1 *http.Response
		result2 error
	}
	getLibraryReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClientInterface) GetLibrary(arg1 context.Context) (*http.Response, error) {
	fake.getLibraryMutex.Lock()
	ret, specificReturn := fake.getLibraryReturnsOnCall[len(fake.getLibraryArgsForCall)]
	fake.getLibraryArgsForCall = append(fake.getLibraryArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetLibraryStub
	fakeReturns := fake.getLibraryReturns
	fake.recordInvocation("GetLibrary", []interface{}{arg1})
	fake.getLibraryMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClientInterface) GetLibraryCallCount() int {
	fake.getLibraryMutex.RLock()
	defer fake.getLibraryMutex.RUnlock()
	return len(fake.getLibraryArgsForCall)
}

func (fake *FakeClientInterface) GetLibraryCalls(stub func(context.Context) (*http.Response, error)) {
	fake.getLibraryMutex.Lock()
	defer fake.getLibraryMutex.Unlock()
	fake.GetLibraryStub = stub
}

func (fake *FakeClientInterface) GetLibraryArgsForCall(i int) context.Context {
	fake.getLibraryMutex.RLock()
	defer fake.getLibraryMutex.RUnlock()
	argsForCall := fake.getLibraryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClientInterface) GetLibraryReturns(result1 *http.Response, result2 error) {
	fake.getLibraryMutex.Lock()
	defer fake.getLibraryMutex.Unlock()
	fake.GetLibraryStub = nil
	fake.getLibraryReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) GetLibraryReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.getLibraryMutex.Lock()
	defer fake.getLibraryMutex.Unlock()
	fake.GetLibraryStub = nil
	if fake.getLibraryReturnsOnCall == nil {
		fake.getLibraryReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.getLibraryReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeClientInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getLibraryMutex.RLock()
	defer fake.getLibraryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClientInterface) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ api.ClientInterface = new(FakeClientInterface)
