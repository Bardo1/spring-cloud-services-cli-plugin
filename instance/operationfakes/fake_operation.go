// Code generated by counterfeiter. DO NOT EDIT.
package operationfakes

import (
	"sync"

	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/httpclient"
	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/instance"
)

type FakeOperation struct {
	IsLifecycleOperationStub        func() bool
	isLifecycleOperationMutex       sync.RWMutex
	isLifecycleOperationArgsForCall []struct {
	}
	isLifecycleOperationReturns struct {
		result1 bool
	}
	isLifecycleOperationReturnsOnCall map[int]struct {
		result1 bool
	}
	RunStub        func(httpclient.AuthenticatedClient, string, string) (string, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 httpclient.AuthenticatedClient
		arg2 string
		arg3 string
	}
	runReturns struct {
		result1 string
		result2 error
	}
	runReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOperation) IsLifecycleOperation() bool {
	fake.isLifecycleOperationMutex.Lock()
	ret, specificReturn := fake.isLifecycleOperationReturnsOnCall[len(fake.isLifecycleOperationArgsForCall)]
	fake.isLifecycleOperationArgsForCall = append(fake.isLifecycleOperationArgsForCall, struct {
	}{})
	fake.recordInvocation("IsLifecycleOperation", []interface{}{})
	fake.isLifecycleOperationMutex.Unlock()
	if fake.IsLifecycleOperationStub != nil {
		return fake.IsLifecycleOperationStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isLifecycleOperationReturns
	return fakeReturns.result1
}

func (fake *FakeOperation) IsLifecycleOperationCallCount() int {
	fake.isLifecycleOperationMutex.RLock()
	defer fake.isLifecycleOperationMutex.RUnlock()
	return len(fake.isLifecycleOperationArgsForCall)
}

func (fake *FakeOperation) IsLifecycleOperationCalls(stub func() bool) {
	fake.isLifecycleOperationMutex.Lock()
	defer fake.isLifecycleOperationMutex.Unlock()
	fake.IsLifecycleOperationStub = stub
}

func (fake *FakeOperation) IsLifecycleOperationReturns(result1 bool) {
	fake.isLifecycleOperationMutex.Lock()
	defer fake.isLifecycleOperationMutex.Unlock()
	fake.IsLifecycleOperationStub = nil
	fake.isLifecycleOperationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeOperation) IsLifecycleOperationReturnsOnCall(i int, result1 bool) {
	fake.isLifecycleOperationMutex.Lock()
	defer fake.isLifecycleOperationMutex.Unlock()
	fake.IsLifecycleOperationStub = nil
	if fake.isLifecycleOperationReturnsOnCall == nil {
		fake.isLifecycleOperationReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isLifecycleOperationReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeOperation) Run(arg1 httpclient.AuthenticatedClient, arg2 string, arg3 string) (string, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 httpclient.AuthenticatedClient
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Run", []interface{}{arg1, arg2, arg3})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.runReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeOperation) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeOperation) RunCalls(stub func(httpclient.AuthenticatedClient, string, string) (string, error)) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeOperation) RunArgsForCall(i int) (httpclient.AuthenticatedClient, string, string) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeOperation) RunReturns(result1 string, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOperation) RunReturnsOnCall(i int, result1 string, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeOperation) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isLifecycleOperationMutex.RLock()
	defer fake.isLifecycleOperationMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOperation) recordInvocation(key string, args []interface{}) {
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

var _ instance.Operation = new(FakeOperation)
