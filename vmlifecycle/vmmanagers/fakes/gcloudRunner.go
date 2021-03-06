// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bytes"
	"sync"
)

type GCloudRunner struct {
	ExecuteStub        func([]interface{}) (*bytes.Buffer, *bytes.Buffer, error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 []interface{}
	}
	executeReturns struct {
		result1 *bytes.Buffer
		result2 *bytes.Buffer
		result3 error
	}
	executeReturnsOnCall map[int]struct {
		result1 *bytes.Buffer
		result2 *bytes.Buffer
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *GCloudRunner) Execute(arg1 []interface{}) (*bytes.Buffer, *bytes.Buffer, error) {
	var arg1Copy []interface{}
	if arg1 != nil {
		arg1Copy = make([]interface{}, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.executeMutex.Lock()
	ret, specificReturn := fake.executeReturnsOnCall[len(fake.executeArgsForCall)]
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 []interface{}
	}{arg1Copy})
	fake.recordInvocation("Execute", []interface{}{arg1Copy})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.executeReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *GCloudRunner) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *GCloudRunner) ExecuteCalls(stub func([]interface{}) (*bytes.Buffer, *bytes.Buffer, error)) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = stub
}

func (fake *GCloudRunner) ExecuteArgsForCall(i int) []interface{} {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	argsForCall := fake.executeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *GCloudRunner) ExecuteReturns(result1 *bytes.Buffer, result2 *bytes.Buffer, result3 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 *bytes.Buffer
		result2 *bytes.Buffer
		result3 error
	}{result1, result2, result3}
}

func (fake *GCloudRunner) ExecuteReturnsOnCall(i int, result1 *bytes.Buffer, result2 *bytes.Buffer, result3 error) {
	fake.executeMutex.Lock()
	defer fake.executeMutex.Unlock()
	fake.ExecuteStub = nil
	if fake.executeReturnsOnCall == nil {
		fake.executeReturnsOnCall = make(map[int]struct {
			result1 *bytes.Buffer
			result2 *bytes.Buffer
			result3 error
		})
	}
	fake.executeReturnsOnCall[i] = struct {
		result1 *bytes.Buffer
		result2 *bytes.Buffer
		result3 error
	}{result1, result2, result3}
}

func (fake *GCloudRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *GCloudRunner) recordInvocation(key string, args []interface{}) {
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
