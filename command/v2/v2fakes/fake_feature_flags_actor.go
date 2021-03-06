// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeFeatureFlagsActor struct {
	GetFeatureFlagsStub        func() ([]v2action.FeatureFlag, v2action.Warnings, error)
	getFeatureFlagsMutex       sync.RWMutex
	getFeatureFlagsArgsForCall []struct{}
	getFeatureFlagsReturns     struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}
	getFeatureFlagsReturnsOnCall map[int]struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFeatureFlagsActor) GetFeatureFlags() ([]v2action.FeatureFlag, v2action.Warnings, error) {
	fake.getFeatureFlagsMutex.Lock()
	ret, specificReturn := fake.getFeatureFlagsReturnsOnCall[len(fake.getFeatureFlagsArgsForCall)]
	fake.getFeatureFlagsArgsForCall = append(fake.getFeatureFlagsArgsForCall, struct{}{})
	fake.recordInvocation("GetFeatureFlags", []interface{}{})
	fake.getFeatureFlagsMutex.Unlock()
	if fake.GetFeatureFlagsStub != nil {
		return fake.GetFeatureFlagsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getFeatureFlagsReturns.result1, fake.getFeatureFlagsReturns.result2, fake.getFeatureFlagsReturns.result3
}

func (fake *FakeFeatureFlagsActor) GetFeatureFlagsCallCount() int {
	fake.getFeatureFlagsMutex.RLock()
	defer fake.getFeatureFlagsMutex.RUnlock()
	return len(fake.getFeatureFlagsArgsForCall)
}

func (fake *FakeFeatureFlagsActor) GetFeatureFlagsReturns(result1 []v2action.FeatureFlag, result2 v2action.Warnings, result3 error) {
	fake.GetFeatureFlagsStub = nil
	fake.getFeatureFlagsReturns = struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFeatureFlagsActor) GetFeatureFlagsReturnsOnCall(i int, result1 []v2action.FeatureFlag, result2 v2action.Warnings, result3 error) {
	fake.GetFeatureFlagsStub = nil
	if fake.getFeatureFlagsReturnsOnCall == nil {
		fake.getFeatureFlagsReturnsOnCall = make(map[int]struct {
			result1 []v2action.FeatureFlag
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getFeatureFlagsReturnsOnCall[i] = struct {
		result1 []v2action.FeatureFlag
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFeatureFlagsActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getFeatureFlagsMutex.RLock()
	defer fake.getFeatureFlagsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFeatureFlagsActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.FeatureFlagsActor = new(FakeFeatureFlagsActor)
