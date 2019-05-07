// Code generated by counterfeiter. DO NOT EDIT.
package constructfakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/stembuild/construct"
)

type FakeConstructMessenger struct {
	CreateProvisionDirStartedStub        func()
	createProvisionDirStartedMutex       sync.RWMutex
	createProvisionDirStartedArgsForCall []struct {
	}
	CreateProvisionDirSucceededStub        func()
	createProvisionDirSucceededMutex       sync.RWMutex
	createProvisionDirSucceededArgsForCall []struct {
	}
	EnableWinRMStartedStub        func()
	enableWinRMStartedMutex       sync.RWMutex
	enableWinRMStartedArgsForCall []struct {
	}
	EnableWinRMSucceededStub        func()
	enableWinRMSucceededMutex       sync.RWMutex
	enableWinRMSucceededArgsForCall []struct {
	}
	ExecuteScriptStartedStub        func()
	executeScriptStartedMutex       sync.RWMutex
	executeScriptStartedArgsForCall []struct {
	}
	ExecuteScriptSucceededStub        func()
	executeScriptSucceededMutex       sync.RWMutex
	executeScriptSucceededArgsForCall []struct {
	}
	ExtractArtifactsStartedStub        func()
	extractArtifactsStartedMutex       sync.RWMutex
	extractArtifactsStartedArgsForCall []struct {
	}
	ExtractArtifactsSucceededStub        func()
	extractArtifactsSucceededMutex       sync.RWMutex
	extractArtifactsSucceededArgsForCall []struct {
	}
	UploadArtifactsStartedStub        func()
	uploadArtifactsStartedMutex       sync.RWMutex
	uploadArtifactsStartedArgsForCall []struct {
	}
	UploadArtifactsSucceededStub        func()
	uploadArtifactsSucceededMutex       sync.RWMutex
	uploadArtifactsSucceededArgsForCall []struct {
	}
	UploadFileStartedStub        func(string)
	uploadFileStartedMutex       sync.RWMutex
	uploadFileStartedArgsForCall []struct {
		arg1 string
	}
	UploadFileSucceededStub        func()
	uploadFileSucceededMutex       sync.RWMutex
	uploadFileSucceededArgsForCall []struct {
	}
	ValidateVMConnectionStartedStub        func()
	validateVMConnectionStartedMutex       sync.RWMutex
	validateVMConnectionStartedArgsForCall []struct {
	}
	ValidateVMConnectionSucceededStub        func()
	validateVMConnectionSucceededMutex       sync.RWMutex
	validateVMConnectionSucceededArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConstructMessenger) CreateProvisionDirStarted() {
	fake.createProvisionDirStartedMutex.Lock()
	fake.createProvisionDirStartedArgsForCall = append(fake.createProvisionDirStartedArgsForCall, struct {
	}{})
	fake.recordInvocation("CreateProvisionDirStarted", []interface{}{})
	fake.createProvisionDirStartedMutex.Unlock()
	if fake.CreateProvisionDirStartedStub != nil {
		fake.CreateProvisionDirStartedStub()
	}
}

func (fake *FakeConstructMessenger) CreateProvisionDirStartedCallCount() int {
	fake.createProvisionDirStartedMutex.RLock()
	defer fake.createProvisionDirStartedMutex.RUnlock()
	return len(fake.createProvisionDirStartedArgsForCall)
}

func (fake *FakeConstructMessenger) CreateProvisionDirStartedCalls(stub func()) {
	fake.createProvisionDirStartedMutex.Lock()
	defer fake.createProvisionDirStartedMutex.Unlock()
	fake.CreateProvisionDirStartedStub = stub
}

func (fake *FakeConstructMessenger) CreateProvisionDirSucceeded() {
	fake.createProvisionDirSucceededMutex.Lock()
	fake.createProvisionDirSucceededArgsForCall = append(fake.createProvisionDirSucceededArgsForCall, struct {
	}{})
	fake.recordInvocation("CreateProvisionDirSucceeded", []interface{}{})
	fake.createProvisionDirSucceededMutex.Unlock()
	if fake.CreateProvisionDirSucceededStub != nil {
		fake.CreateProvisionDirSucceededStub()
	}
}

func (fake *FakeConstructMessenger) CreateProvisionDirSucceededCallCount() int {
	fake.createProvisionDirSucceededMutex.RLock()
	defer fake.createProvisionDirSucceededMutex.RUnlock()
	return len(fake.createProvisionDirSucceededArgsForCall)
}

func (fake *FakeConstructMessenger) CreateProvisionDirSucceededCalls(stub func()) {
	fake.createProvisionDirSucceededMutex.Lock()
	defer fake.createProvisionDirSucceededMutex.Unlock()
	fake.CreateProvisionDirSucceededStub = stub
}

func (fake *FakeConstructMessenger) EnableWinRMStarted() {
	fake.enableWinRMStartedMutex.Lock()
	fake.enableWinRMStartedArgsForCall = append(fake.enableWinRMStartedArgsForCall, struct {
	}{})
	fake.recordInvocation("EnableWinRMStarted", []interface{}{})
	fake.enableWinRMStartedMutex.Unlock()
	if fake.EnableWinRMStartedStub != nil {
		fake.EnableWinRMStartedStub()
	}
}

func (fake *FakeConstructMessenger) EnableWinRMStartedCallCount() int {
	fake.enableWinRMStartedMutex.RLock()
	defer fake.enableWinRMStartedMutex.RUnlock()
	return len(fake.enableWinRMStartedArgsForCall)
}

func (fake *FakeConstructMessenger) EnableWinRMStartedCalls(stub func()) {
	fake.enableWinRMStartedMutex.Lock()
	defer fake.enableWinRMStartedMutex.Unlock()
	fake.EnableWinRMStartedStub = stub
}

func (fake *FakeConstructMessenger) EnableWinRMSucceeded() {
	fake.enableWinRMSucceededMutex.Lock()
	fake.enableWinRMSucceededArgsForCall = append(fake.enableWinRMSucceededArgsForCall, struct {
	}{})
	fake.recordInvocation("EnableWinRMSucceeded", []interface{}{})
	fake.enableWinRMSucceededMutex.Unlock()
	if fake.EnableWinRMSucceededStub != nil {
		fake.EnableWinRMSucceededStub()
	}
}

func (fake *FakeConstructMessenger) EnableWinRMSucceededCallCount() int {
	fake.enableWinRMSucceededMutex.RLock()
	defer fake.enableWinRMSucceededMutex.RUnlock()
	return len(fake.enableWinRMSucceededArgsForCall)
}

func (fake *FakeConstructMessenger) EnableWinRMSucceededCalls(stub func()) {
	fake.enableWinRMSucceededMutex.Lock()
	defer fake.enableWinRMSucceededMutex.Unlock()
	fake.EnableWinRMSucceededStub = stub
}

func (fake *FakeConstructMessenger) ExecuteScriptStarted() {
	fake.executeScriptStartedMutex.Lock()
	fake.executeScriptStartedArgsForCall = append(fake.executeScriptStartedArgsForCall, struct {
	}{})
	fake.recordInvocation("ExecuteScriptStarted", []interface{}{})
	fake.executeScriptStartedMutex.Unlock()
	if fake.ExecuteScriptStartedStub != nil {
		fake.ExecuteScriptStartedStub()
	}
}

func (fake *FakeConstructMessenger) ExecuteScriptStartedCallCount() int {
	fake.executeScriptStartedMutex.RLock()
	defer fake.executeScriptStartedMutex.RUnlock()
	return len(fake.executeScriptStartedArgsForCall)
}

func (fake *FakeConstructMessenger) ExecuteScriptStartedCalls(stub func()) {
	fake.executeScriptStartedMutex.Lock()
	defer fake.executeScriptStartedMutex.Unlock()
	fake.ExecuteScriptStartedStub = stub
}

func (fake *FakeConstructMessenger) ExecuteScriptSucceeded() {
	fake.executeScriptSucceededMutex.Lock()
	fake.executeScriptSucceededArgsForCall = append(fake.executeScriptSucceededArgsForCall, struct {
	}{})
	fake.recordInvocation("ExecuteScriptSucceeded", []interface{}{})
	fake.executeScriptSucceededMutex.Unlock()
	if fake.ExecuteScriptSucceededStub != nil {
		fake.ExecuteScriptSucceededStub()
	}
}

func (fake *FakeConstructMessenger) ExecuteScriptSucceededCallCount() int {
	fake.executeScriptSucceededMutex.RLock()
	defer fake.executeScriptSucceededMutex.RUnlock()
	return len(fake.executeScriptSucceededArgsForCall)
}

func (fake *FakeConstructMessenger) ExecuteScriptSucceededCalls(stub func()) {
	fake.executeScriptSucceededMutex.Lock()
	defer fake.executeScriptSucceededMutex.Unlock()
	fake.ExecuteScriptSucceededStub = stub
}

func (fake *FakeConstructMessenger) ExtractArtifactsStarted() {
	fake.extractArtifactsStartedMutex.Lock()
	fake.extractArtifactsStartedArgsForCall = append(fake.extractArtifactsStartedArgsForCall, struct {
	}{})
	fake.recordInvocation("ExtractArtifactsStarted", []interface{}{})
	fake.extractArtifactsStartedMutex.Unlock()
	if fake.ExtractArtifactsStartedStub != nil {
		fake.ExtractArtifactsStartedStub()
	}
}

func (fake *FakeConstructMessenger) ExtractArtifactsStartedCallCount() int {
	fake.extractArtifactsStartedMutex.RLock()
	defer fake.extractArtifactsStartedMutex.RUnlock()
	return len(fake.extractArtifactsStartedArgsForCall)
}

func (fake *FakeConstructMessenger) ExtractArtifactsStartedCalls(stub func()) {
	fake.extractArtifactsStartedMutex.Lock()
	defer fake.extractArtifactsStartedMutex.Unlock()
	fake.ExtractArtifactsStartedStub = stub
}

func (fake *FakeConstructMessenger) ExtractArtifactsSucceeded() {
	fake.extractArtifactsSucceededMutex.Lock()
	fake.extractArtifactsSucceededArgsForCall = append(fake.extractArtifactsSucceededArgsForCall, struct {
	}{})
	fake.recordInvocation("ExtractArtifactsSucceeded", []interface{}{})
	fake.extractArtifactsSucceededMutex.Unlock()
	if fake.ExtractArtifactsSucceededStub != nil {
		fake.ExtractArtifactsSucceededStub()
	}
}

func (fake *FakeConstructMessenger) ExtractArtifactsSucceededCallCount() int {
	fake.extractArtifactsSucceededMutex.RLock()
	defer fake.extractArtifactsSucceededMutex.RUnlock()
	return len(fake.extractArtifactsSucceededArgsForCall)
}

func (fake *FakeConstructMessenger) ExtractArtifactsSucceededCalls(stub func()) {
	fake.extractArtifactsSucceededMutex.Lock()
	defer fake.extractArtifactsSucceededMutex.Unlock()
	fake.ExtractArtifactsSucceededStub = stub
}

func (fake *FakeConstructMessenger) UploadArtifactsStarted() {
	fake.uploadArtifactsStartedMutex.Lock()
	fake.uploadArtifactsStartedArgsForCall = append(fake.uploadArtifactsStartedArgsForCall, struct {
	}{})
	fake.recordInvocation("UploadArtifactsStarted", []interface{}{})
	fake.uploadArtifactsStartedMutex.Unlock()
	if fake.UploadArtifactsStartedStub != nil {
		fake.UploadArtifactsStartedStub()
	}
}

func (fake *FakeConstructMessenger) UploadArtifactsStartedCallCount() int {
	fake.uploadArtifactsStartedMutex.RLock()
	defer fake.uploadArtifactsStartedMutex.RUnlock()
	return len(fake.uploadArtifactsStartedArgsForCall)
}

func (fake *FakeConstructMessenger) UploadArtifactsStartedCalls(stub func()) {
	fake.uploadArtifactsStartedMutex.Lock()
	defer fake.uploadArtifactsStartedMutex.Unlock()
	fake.UploadArtifactsStartedStub = stub
}

func (fake *FakeConstructMessenger) UploadArtifactsSucceeded() {
	fake.uploadArtifactsSucceededMutex.Lock()
	fake.uploadArtifactsSucceededArgsForCall = append(fake.uploadArtifactsSucceededArgsForCall, struct {
	}{})
	fake.recordInvocation("UploadArtifactsSucceeded", []interface{}{})
	fake.uploadArtifactsSucceededMutex.Unlock()
	if fake.UploadArtifactsSucceededStub != nil {
		fake.UploadArtifactsSucceededStub()
	}
}

func (fake *FakeConstructMessenger) UploadArtifactsSucceededCallCount() int {
	fake.uploadArtifactsSucceededMutex.RLock()
	defer fake.uploadArtifactsSucceededMutex.RUnlock()
	return len(fake.uploadArtifactsSucceededArgsForCall)
}

func (fake *FakeConstructMessenger) UploadArtifactsSucceededCalls(stub func()) {
	fake.uploadArtifactsSucceededMutex.Lock()
	defer fake.uploadArtifactsSucceededMutex.Unlock()
	fake.UploadArtifactsSucceededStub = stub
}

func (fake *FakeConstructMessenger) UploadFileStarted(arg1 string) {
	fake.uploadFileStartedMutex.Lock()
	fake.uploadFileStartedArgsForCall = append(fake.uploadFileStartedArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UploadFileStarted", []interface{}{arg1})
	fake.uploadFileStartedMutex.Unlock()
	if fake.UploadFileStartedStub != nil {
		fake.UploadFileStartedStub(arg1)
	}
}

func (fake *FakeConstructMessenger) UploadFileStartedCallCount() int {
	fake.uploadFileStartedMutex.RLock()
	defer fake.uploadFileStartedMutex.RUnlock()
	return len(fake.uploadFileStartedArgsForCall)
}

func (fake *FakeConstructMessenger) UploadFileStartedCalls(stub func(string)) {
	fake.uploadFileStartedMutex.Lock()
	defer fake.uploadFileStartedMutex.Unlock()
	fake.UploadFileStartedStub = stub
}

func (fake *FakeConstructMessenger) UploadFileStartedArgsForCall(i int) string {
	fake.uploadFileStartedMutex.RLock()
	defer fake.uploadFileStartedMutex.RUnlock()
	argsForCall := fake.uploadFileStartedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConstructMessenger) UploadFileSucceeded() {
	fake.uploadFileSucceededMutex.Lock()
	fake.uploadFileSucceededArgsForCall = append(fake.uploadFileSucceededArgsForCall, struct {
	}{})
	fake.recordInvocation("UploadFileSucceeded", []interface{}{})
	fake.uploadFileSucceededMutex.Unlock()
	if fake.UploadFileSucceededStub != nil {
		fake.UploadFileSucceededStub()
	}
}

func (fake *FakeConstructMessenger) UploadFileSucceededCallCount() int {
	fake.uploadFileSucceededMutex.RLock()
	defer fake.uploadFileSucceededMutex.RUnlock()
	return len(fake.uploadFileSucceededArgsForCall)
}

func (fake *FakeConstructMessenger) UploadFileSucceededCalls(stub func()) {
	fake.uploadFileSucceededMutex.Lock()
	defer fake.uploadFileSucceededMutex.Unlock()
	fake.UploadFileSucceededStub = stub
}

func (fake *FakeConstructMessenger) ValidateVMConnectionStarted() {
	fake.validateVMConnectionStartedMutex.Lock()
	fake.validateVMConnectionStartedArgsForCall = append(fake.validateVMConnectionStartedArgsForCall, struct {
	}{})
	fake.recordInvocation("ValidateVMConnectionStarted", []interface{}{})
	fake.validateVMConnectionStartedMutex.Unlock()
	if fake.ValidateVMConnectionStartedStub != nil {
		fake.ValidateVMConnectionStartedStub()
	}
}

func (fake *FakeConstructMessenger) ValidateVMConnectionStartedCallCount() int {
	fake.validateVMConnectionStartedMutex.RLock()
	defer fake.validateVMConnectionStartedMutex.RUnlock()
	return len(fake.validateVMConnectionStartedArgsForCall)
}

func (fake *FakeConstructMessenger) ValidateVMConnectionStartedCalls(stub func()) {
	fake.validateVMConnectionStartedMutex.Lock()
	defer fake.validateVMConnectionStartedMutex.Unlock()
	fake.ValidateVMConnectionStartedStub = stub
}

func (fake *FakeConstructMessenger) ValidateVMConnectionSucceeded() {
	fake.validateVMConnectionSucceededMutex.Lock()
	fake.validateVMConnectionSucceededArgsForCall = append(fake.validateVMConnectionSucceededArgsForCall, struct {
	}{})
	fake.recordInvocation("ValidateVMConnectionSucceeded", []interface{}{})
	fake.validateVMConnectionSucceededMutex.Unlock()
	if fake.ValidateVMConnectionSucceededStub != nil {
		fake.ValidateVMConnectionSucceededStub()
	}
}

func (fake *FakeConstructMessenger) ValidateVMConnectionSucceededCallCount() int {
	fake.validateVMConnectionSucceededMutex.RLock()
	defer fake.validateVMConnectionSucceededMutex.RUnlock()
	return len(fake.validateVMConnectionSucceededArgsForCall)
}

func (fake *FakeConstructMessenger) ValidateVMConnectionSucceededCalls(stub func()) {
	fake.validateVMConnectionSucceededMutex.Lock()
	defer fake.validateVMConnectionSucceededMutex.Unlock()
	fake.ValidateVMConnectionSucceededStub = stub
}

func (fake *FakeConstructMessenger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createProvisionDirStartedMutex.RLock()
	defer fake.createProvisionDirStartedMutex.RUnlock()
	fake.createProvisionDirSucceededMutex.RLock()
	defer fake.createProvisionDirSucceededMutex.RUnlock()
	fake.enableWinRMStartedMutex.RLock()
	defer fake.enableWinRMStartedMutex.RUnlock()
	fake.enableWinRMSucceededMutex.RLock()
	defer fake.enableWinRMSucceededMutex.RUnlock()
	fake.executeScriptStartedMutex.RLock()
	defer fake.executeScriptStartedMutex.RUnlock()
	fake.executeScriptSucceededMutex.RLock()
	defer fake.executeScriptSucceededMutex.RUnlock()
	fake.extractArtifactsStartedMutex.RLock()
	defer fake.extractArtifactsStartedMutex.RUnlock()
	fake.extractArtifactsSucceededMutex.RLock()
	defer fake.extractArtifactsSucceededMutex.RUnlock()
	fake.uploadArtifactsStartedMutex.RLock()
	defer fake.uploadArtifactsStartedMutex.RUnlock()
	fake.uploadArtifactsSucceededMutex.RLock()
	defer fake.uploadArtifactsSucceededMutex.RUnlock()
	fake.uploadFileStartedMutex.RLock()
	defer fake.uploadFileStartedMutex.RUnlock()
	fake.uploadFileSucceededMutex.RLock()
	defer fake.uploadFileSucceededMutex.RUnlock()
	fake.validateVMConnectionStartedMutex.RLock()
	defer fake.validateVMConnectionStartedMutex.RUnlock()
	fake.validateVMConnectionSucceededMutex.RLock()
	defer fake.validateVMConnectionSucceededMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConstructMessenger) recordInvocation(key string, args []interface{}) {
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

var _ construct.ConstructMessenger = new(FakeConstructMessenger)