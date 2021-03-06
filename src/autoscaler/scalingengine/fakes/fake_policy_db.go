// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"autoscaler/db"
	"autoscaler/models"
	"sync"
)

type FakePolicyDB struct {
	GetAppIdsStub        func() (map[string]bool, error)
	getAppIdsMutex       sync.RWMutex
	getAppIdsArgsForCall []struct{}
	getAppIdsReturns     struct {
		result1 map[string]bool
		result2 error
	}
	getAppIdsReturnsOnCall map[int]struct {
		result1 map[string]bool
		result2 error
	}
	GetAppPolicyStub        func(appId string) (*models.ScalingPolicy, error)
	getAppPolicyMutex       sync.RWMutex
	getAppPolicyArgsForCall []struct {
		appId string
	}
	getAppPolicyReturns struct {
		result1 *models.ScalingPolicy
		result2 error
	}
	getAppPolicyReturnsOnCall map[int]struct {
		result1 *models.ScalingPolicy
		result2 error
	}
	RetrievePoliciesStub        func() ([]*models.PolicyJson, error)
	retrievePoliciesMutex       sync.RWMutex
	retrievePoliciesArgsForCall []struct{}
	retrievePoliciesReturns     struct {
		result1 []*models.PolicyJson
		result2 error
	}
	retrievePoliciesReturnsOnCall map[int]struct {
		result1 []*models.PolicyJson
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePolicyDB) GetAppIds() (map[string]bool, error) {
	fake.getAppIdsMutex.Lock()
	ret, specificReturn := fake.getAppIdsReturnsOnCall[len(fake.getAppIdsArgsForCall)]
	fake.getAppIdsArgsForCall = append(fake.getAppIdsArgsForCall, struct{}{})
	fake.recordInvocation("GetAppIds", []interface{}{})
	fake.getAppIdsMutex.Unlock()
	if fake.GetAppIdsStub != nil {
		return fake.GetAppIdsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAppIdsReturns.result1, fake.getAppIdsReturns.result2
}

func (fake *FakePolicyDB) GetAppIdsCallCount() int {
	fake.getAppIdsMutex.RLock()
	defer fake.getAppIdsMutex.RUnlock()
	return len(fake.getAppIdsArgsForCall)
}

func (fake *FakePolicyDB) GetAppIdsReturns(result1 map[string]bool, result2 error) {
	fake.GetAppIdsStub = nil
	fake.getAppIdsReturns = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) GetAppIdsReturnsOnCall(i int, result1 map[string]bool, result2 error) {
	fake.GetAppIdsStub = nil
	if fake.getAppIdsReturnsOnCall == nil {
		fake.getAppIdsReturnsOnCall = make(map[int]struct {
			result1 map[string]bool
			result2 error
		})
	}
	fake.getAppIdsReturnsOnCall[i] = struct {
		result1 map[string]bool
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) GetAppPolicy(appId string) (*models.ScalingPolicy, error) {
	fake.getAppPolicyMutex.Lock()
	ret, specificReturn := fake.getAppPolicyReturnsOnCall[len(fake.getAppPolicyArgsForCall)]
	fake.getAppPolicyArgsForCall = append(fake.getAppPolicyArgsForCall, struct {
		appId string
	}{appId})
	fake.recordInvocation("GetAppPolicy", []interface{}{appId})
	fake.getAppPolicyMutex.Unlock()
	if fake.GetAppPolicyStub != nil {
		return fake.GetAppPolicyStub(appId)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAppPolicyReturns.result1, fake.getAppPolicyReturns.result2
}

func (fake *FakePolicyDB) GetAppPolicyCallCount() int {
	fake.getAppPolicyMutex.RLock()
	defer fake.getAppPolicyMutex.RUnlock()
	return len(fake.getAppPolicyArgsForCall)
}

func (fake *FakePolicyDB) GetAppPolicyArgsForCall(i int) string {
	fake.getAppPolicyMutex.RLock()
	defer fake.getAppPolicyMutex.RUnlock()
	return fake.getAppPolicyArgsForCall[i].appId
}

func (fake *FakePolicyDB) GetAppPolicyReturns(result1 *models.ScalingPolicy, result2 error) {
	fake.GetAppPolicyStub = nil
	fake.getAppPolicyReturns = struct {
		result1 *models.ScalingPolicy
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) GetAppPolicyReturnsOnCall(i int, result1 *models.ScalingPolicy, result2 error) {
	fake.GetAppPolicyStub = nil
	if fake.getAppPolicyReturnsOnCall == nil {
		fake.getAppPolicyReturnsOnCall = make(map[int]struct {
			result1 *models.ScalingPolicy
			result2 error
		})
	}
	fake.getAppPolicyReturnsOnCall[i] = struct {
		result1 *models.ScalingPolicy
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) RetrievePolicies() ([]*models.PolicyJson, error) {
	fake.retrievePoliciesMutex.Lock()
	ret, specificReturn := fake.retrievePoliciesReturnsOnCall[len(fake.retrievePoliciesArgsForCall)]
	fake.retrievePoliciesArgsForCall = append(fake.retrievePoliciesArgsForCall, struct{}{})
	fake.recordInvocation("RetrievePolicies", []interface{}{})
	fake.retrievePoliciesMutex.Unlock()
	if fake.RetrievePoliciesStub != nil {
		return fake.RetrievePoliciesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.retrievePoliciesReturns.result1, fake.retrievePoliciesReturns.result2
}

func (fake *FakePolicyDB) RetrievePoliciesCallCount() int {
	fake.retrievePoliciesMutex.RLock()
	defer fake.retrievePoliciesMutex.RUnlock()
	return len(fake.retrievePoliciesArgsForCall)
}

func (fake *FakePolicyDB) RetrievePoliciesReturns(result1 []*models.PolicyJson, result2 error) {
	fake.RetrievePoliciesStub = nil
	fake.retrievePoliciesReturns = struct {
		result1 []*models.PolicyJson
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) RetrievePoliciesReturnsOnCall(i int, result1 []*models.PolicyJson, result2 error) {
	fake.RetrievePoliciesStub = nil
	if fake.retrievePoliciesReturnsOnCall == nil {
		fake.retrievePoliciesReturnsOnCall = make(map[int]struct {
			result1 []*models.PolicyJson
			result2 error
		})
	}
	fake.retrievePoliciesReturnsOnCall[i] = struct {
		result1 []*models.PolicyJson
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyDB) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakePolicyDB) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakePolicyDB) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePolicyDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAppIdsMutex.RLock()
	defer fake.getAppIdsMutex.RUnlock()
	fake.getAppPolicyMutex.RLock()
	defer fake.getAppPolicyMutex.RUnlock()
	fake.retrievePoliciesMutex.RLock()
	defer fake.retrievePoliciesMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePolicyDB) recordInvocation(key string, args []interface{}) {
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

var _ db.PolicyDB = new(FakePolicyDB)
