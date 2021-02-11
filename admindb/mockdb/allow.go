// Code generated by counterfeiter. DO NOT EDIT.
package mockdb

import (
	"context"
	"sync"

	"github.com/ssb-ngi-pointer/go-ssb-room/admindb"
	refs "go.mindeco.de/ssb-refs"
)

type FakeAllowListService struct {
	AddStub        func(context.Context, refs.FeedRef) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 context.Context
		arg2 refs.FeedRef
	}
	addReturns struct {
		result1 error
	}
	addReturnsOnCall map[int]struct {
		result1 error
	}
	HasStub        func(context.Context, refs.FeedRef) bool
	hasMutex       sync.RWMutex
	hasArgsForCall []struct {
		arg1 context.Context
		arg2 refs.FeedRef
	}
	hasReturns struct {
		result1 bool
	}
	hasReturnsOnCall map[int]struct {
		result1 bool
	}
	ListStub        func(context.Context) ([]refs.FeedRef, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
	}
	listReturns struct {
		result1 []refs.FeedRef
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []refs.FeedRef
		result2 error
	}
	RemoveStub        func(context.Context, refs.FeedRef) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 context.Context
		arg2 refs.FeedRef
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAllowListService) Add(arg1 context.Context, arg2 refs.FeedRef) error {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 context.Context
		arg2 refs.FeedRef
	}{arg1, arg2})
	stub := fake.AddStub
	fakeReturns := fake.addReturns
	fake.recordInvocation("Add", []interface{}{arg1, arg2})
	fake.addMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAllowListService) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeAllowListService) AddCalls(stub func(context.Context, refs.FeedRef) error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *FakeAllowListService) AddArgsForCall(i int) (context.Context, refs.FeedRef) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAllowListService) AddReturns(result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAllowListService) AddReturnsOnCall(i int, result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAllowListService) Has(arg1 context.Context, arg2 refs.FeedRef) bool {
	fake.hasMutex.Lock()
	ret, specificReturn := fake.hasReturnsOnCall[len(fake.hasArgsForCall)]
	fake.hasArgsForCall = append(fake.hasArgsForCall, struct {
		arg1 context.Context
		arg2 refs.FeedRef
	}{arg1, arg2})
	stub := fake.HasStub
	fakeReturns := fake.hasReturns
	fake.recordInvocation("Has", []interface{}{arg1, arg2})
	fake.hasMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAllowListService) HasCallCount() int {
	fake.hasMutex.RLock()
	defer fake.hasMutex.RUnlock()
	return len(fake.hasArgsForCall)
}

func (fake *FakeAllowListService) HasCalls(stub func(context.Context, refs.FeedRef) bool) {
	fake.hasMutex.Lock()
	defer fake.hasMutex.Unlock()
	fake.HasStub = stub
}

func (fake *FakeAllowListService) HasArgsForCall(i int) (context.Context, refs.FeedRef) {
	fake.hasMutex.RLock()
	defer fake.hasMutex.RUnlock()
	argsForCall := fake.hasArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAllowListService) HasReturns(result1 bool) {
	fake.hasMutex.Lock()
	defer fake.hasMutex.Unlock()
	fake.HasStub = nil
	fake.hasReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAllowListService) HasReturnsOnCall(i int, result1 bool) {
	fake.hasMutex.Lock()
	defer fake.hasMutex.Unlock()
	fake.HasStub = nil
	if fake.hasReturnsOnCall == nil {
		fake.hasReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.hasReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAllowListService) List(arg1 context.Context) ([]refs.FeedRef, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAllowListService) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeAllowListService) ListCalls(stub func(context.Context) ([]refs.FeedRef, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeAllowListService) ListArgsForCall(i int) context.Context {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAllowListService) ListReturns(result1 []refs.FeedRef, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []refs.FeedRef
		result2 error
	}{result1, result2}
}

func (fake *FakeAllowListService) ListReturnsOnCall(i int, result1 []refs.FeedRef, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []refs.FeedRef
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []refs.FeedRef
		result2 error
	}{result1, result2}
}

func (fake *FakeAllowListService) Remove(arg1 context.Context, arg2 refs.FeedRef) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 context.Context
		arg2 refs.FeedRef
	}{arg1, arg2})
	stub := fake.RemoveStub
	fakeReturns := fake.removeReturns
	fake.recordInvocation("Remove", []interface{}{arg1, arg2})
	fake.removeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAllowListService) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeAllowListService) RemoveCalls(stub func(context.Context, refs.FeedRef) error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = stub
}

func (fake *FakeAllowListService) RemoveArgsForCall(i int) (context.Context, refs.FeedRef) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	argsForCall := fake.removeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAllowListService) RemoveReturns(result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAllowListService) RemoveReturnsOnCall(i int, result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAllowListService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.hasMutex.RLock()
	defer fake.hasMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAllowListService) recordInvocation(key string, args []interface{}) {
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

var _ admindb.AllowListService = new(FakeAllowListService)
