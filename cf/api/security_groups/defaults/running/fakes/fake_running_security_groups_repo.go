// This file was generated by counterfeiter
package fakes

import (
	"sync"

	. "github.com/cloudfoundry/cli/cf/api/security_groups/defaults"
	. "github.com/cloudfoundry/cli/cf/api/security_groups/defaults/running"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeRunningSecurityGroupsRepo struct {
	AddToRunningSetStub        func(string) error
	addToRunningSetMutex       sync.RWMutex
	addToRunningSetArgsForCall []struct {
		arg1 string
	}
	addToRunningSetReturns struct {
		result1 error
	}
	ListStub        func() ([]models.SecurityGroupFields, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct{}
	listReturns     struct {
		result1 []models.SecurityGroupFields
		result2 error
	}
	RemoveFromRunningSetStub        func(string) error
	removeFromRunningSetMutex       sync.RWMutex
	removeFromRunningSetArgsForCall []struct {
		arg1 string
	}
	removeFromRunningSetReturns struct {
		result1 error
	}
}

func (fake *FakeRunningSecurityGroupsRepo) AddToRunningSet(arg1 string) error {
	fake.addToRunningSetMutex.Lock()
	defer fake.addToRunningSetMutex.Unlock()
	fake.addToRunningSetArgsForCall = append(fake.addToRunningSetArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.AddToRunningSetStub != nil {
		return fake.AddToRunningSetStub(arg1)
	} else {
		return fake.addToRunningSetReturns.result1
	}
}

func (fake *FakeRunningSecurityGroupsRepo) AddToRunningSetCallCount() int {
	fake.addToRunningSetMutex.RLock()
	defer fake.addToRunningSetMutex.RUnlock()
	return len(fake.addToRunningSetArgsForCall)
}

func (fake *FakeRunningSecurityGroupsRepo) AddToRunningSetArgsForCall(i int) string {
	fake.addToRunningSetMutex.RLock()
	defer fake.addToRunningSetMutex.RUnlock()
	return fake.addToRunningSetArgsForCall[i].arg1
}

func (fake *FakeRunningSecurityGroupsRepo) AddToRunningSetReturns(result1 error) {
	fake.addToRunningSetReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRunningSecurityGroupsRepo) List() ([]models.SecurityGroupFields, error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct{}{})
	if fake.ListStub != nil {
		return fake.ListStub()
	} else {
		return fake.listReturns.result1, fake.listReturns.result2
	}
}

func (fake *FakeRunningSecurityGroupsRepo) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeRunningSecurityGroupsRepo) ListReturns(result1 []models.SecurityGroupFields, result2 error) {
	fake.listReturns = struct {
		result1 []models.SecurityGroupFields
		result2 error
	}{result1, result2}
}

func (fake *FakeRunningSecurityGroupsRepo) RemoveFromRunningSet(arg1 string) error {
	fake.removeFromRunningSetMutex.Lock()
	defer fake.removeFromRunningSetMutex.Unlock()
	fake.removeFromRunningSetArgsForCall = append(fake.removeFromRunningSetArgsForCall, struct {
		arg1 string
	}{arg1})
	if fake.RemoveFromRunningSetStub != nil {
		return fake.RemoveFromRunningSetStub(arg1)
	} else {
		return fake.removeFromRunningSetReturns.result1
	}
}

func (fake *FakeRunningSecurityGroupsRepo) RemoveFromRunningSetCallCount() int {
	fake.removeFromRunningSetMutex.RLock()
	defer fake.removeFromRunningSetMutex.RUnlock()
	return len(fake.removeFromRunningSetArgsForCall)
}

func (fake *FakeRunningSecurityGroupsRepo) RemoveFromRunningSetArgsForCall(i int) string {
	fake.removeFromRunningSetMutex.RLock()
	defer fake.removeFromRunningSetMutex.RUnlock()
	return fake.removeFromRunningSetArgsForCall[i].arg1
}

func (fake *FakeRunningSecurityGroupsRepo) RemoveFromRunningSetReturns(result1 error) {
	fake.removeFromRunningSetReturns = struct {
		result1 error
	}{result1}
}

var _ RunningSecurityGroupsRepo = new(FakeRunningSecurityGroupsRepo)