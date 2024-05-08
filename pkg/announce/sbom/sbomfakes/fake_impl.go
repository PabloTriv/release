/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package sbomfakes

import (
	"sync"

	"sigs.k8s.io/bom/pkg/spdx"
)

type FakeImpl struct {
	docBuilderStub        func() *spdx.DocBuilder
	docBuilderMutex       sync.RWMutex
	docBuilderArgsForCall []struct {
	}
	docBuilderReturns struct {
		result1 *spdx.DocBuilder
	}
	docBuilderReturnsOnCall map[int]struct {
		result1 *spdx.DocBuilder
	}
	spdxClientStub        func() *spdx.SPDX
	spdxClientMutex       sync.RWMutex
	spdxClientArgsForCall []struct {
	}
	spdxClientReturns struct {
		result1 *spdx.SPDX
	}
	spdxClientReturnsOnCall map[int]struct {
		result1 *spdx.SPDX
	}
	tmpFileStub        func() (string, error)
	tmpFileMutex       sync.RWMutex
	tmpFileArgsForCall []struct {
	}
	tmpFileReturns struct {
		result1 string
		result2 error
	}
	tmpFileReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	writeFileStub        func(string, []byte) error
	writeFileMutex       sync.RWMutex
	writeFileArgsForCall []struct {
		arg1 string
		arg2 []byte
	}
	writeFileReturns struct {
		result1 error
	}
	writeFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) docBuilder() *spdx.DocBuilder {
	fake.docBuilderMutex.Lock()
	ret, specificReturn := fake.docBuilderReturnsOnCall[len(fake.docBuilderArgsForCall)]
	fake.docBuilderArgsForCall = append(fake.docBuilderArgsForCall, struct {
	}{})
	stub := fake.docBuilderStub
	fakeReturns := fake.docBuilderReturns
	fake.recordInvocation("docBuilder", []interface{}{})
	fake.docBuilderMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) DocBuilderCallCount() int {
	fake.docBuilderMutex.RLock()
	defer fake.docBuilderMutex.RUnlock()
	return len(fake.docBuilderArgsForCall)
}

func (fake *FakeImpl) DocBuilderCalls(stub func() *spdx.DocBuilder) {
	fake.docBuilderMutex.Lock()
	defer fake.docBuilderMutex.Unlock()
	fake.docBuilderStub = stub
}

func (fake *FakeImpl) DocBuilderReturns(result1 *spdx.DocBuilder) {
	fake.docBuilderMutex.Lock()
	defer fake.docBuilderMutex.Unlock()
	fake.docBuilderStub = nil
	fake.docBuilderReturns = struct {
		result1 *spdx.DocBuilder
	}{result1}
}

func (fake *FakeImpl) DocBuilderReturnsOnCall(i int, result1 *spdx.DocBuilder) {
	fake.docBuilderMutex.Lock()
	defer fake.docBuilderMutex.Unlock()
	fake.docBuilderStub = nil
	if fake.docBuilderReturnsOnCall == nil {
		fake.docBuilderReturnsOnCall = make(map[int]struct {
			result1 *spdx.DocBuilder
		})
	}
	fake.docBuilderReturnsOnCall[i] = struct {
		result1 *spdx.DocBuilder
	}{result1}
}

func (fake *FakeImpl) spdxClient() *spdx.SPDX {
	fake.spdxClientMutex.Lock()
	ret, specificReturn := fake.spdxClientReturnsOnCall[len(fake.spdxClientArgsForCall)]
	fake.spdxClientArgsForCall = append(fake.spdxClientArgsForCall, struct {
	}{})
	stub := fake.spdxClientStub
	fakeReturns := fake.spdxClientReturns
	fake.recordInvocation("spdxClient", []interface{}{})
	fake.spdxClientMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) SpdxClientCallCount() int {
	fake.spdxClientMutex.RLock()
	defer fake.spdxClientMutex.RUnlock()
	return len(fake.spdxClientArgsForCall)
}

func (fake *FakeImpl) SpdxClientCalls(stub func() *spdx.SPDX) {
	fake.spdxClientMutex.Lock()
	defer fake.spdxClientMutex.Unlock()
	fake.spdxClientStub = stub
}

func (fake *FakeImpl) SpdxClientReturns(result1 *spdx.SPDX) {
	fake.spdxClientMutex.Lock()
	defer fake.spdxClientMutex.Unlock()
	fake.spdxClientStub = nil
	fake.spdxClientReturns = struct {
		result1 *spdx.SPDX
	}{result1}
}

func (fake *FakeImpl) SpdxClientReturnsOnCall(i int, result1 *spdx.SPDX) {
	fake.spdxClientMutex.Lock()
	defer fake.spdxClientMutex.Unlock()
	fake.spdxClientStub = nil
	if fake.spdxClientReturnsOnCall == nil {
		fake.spdxClientReturnsOnCall = make(map[int]struct {
			result1 *spdx.SPDX
		})
	}
	fake.spdxClientReturnsOnCall[i] = struct {
		result1 *spdx.SPDX
	}{result1}
}

func (fake *FakeImpl) tmpFile() (string, error) {
	fake.tmpFileMutex.Lock()
	ret, specificReturn := fake.tmpFileReturnsOnCall[len(fake.tmpFileArgsForCall)]
	fake.tmpFileArgsForCall = append(fake.tmpFileArgsForCall, struct {
	}{})
	stub := fake.tmpFileStub
	fakeReturns := fake.tmpFileReturns
	fake.recordInvocation("tmpFile", []interface{}{})
	fake.tmpFileMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) TmpFileCallCount() int {
	fake.tmpFileMutex.RLock()
	defer fake.tmpFileMutex.RUnlock()
	return len(fake.tmpFileArgsForCall)
}

func (fake *FakeImpl) TmpFileCalls(stub func() (string, error)) {
	fake.tmpFileMutex.Lock()
	defer fake.tmpFileMutex.Unlock()
	fake.tmpFileStub = stub
}

func (fake *FakeImpl) TmpFileReturns(result1 string, result2 error) {
	fake.tmpFileMutex.Lock()
	defer fake.tmpFileMutex.Unlock()
	fake.tmpFileStub = nil
	fake.tmpFileReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) TmpFileReturnsOnCall(i int, result1 string, result2 error) {
	fake.tmpFileMutex.Lock()
	defer fake.tmpFileMutex.Unlock()
	fake.tmpFileStub = nil
	if fake.tmpFileReturnsOnCall == nil {
		fake.tmpFileReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.tmpFileReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) writeFile(arg1 string, arg2 []byte) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeFileMutex.Lock()
	ret, specificReturn := fake.writeFileReturnsOnCall[len(fake.writeFileArgsForCall)]
	fake.writeFileArgsForCall = append(fake.writeFileArgsForCall, struct {
		arg1 string
		arg2 []byte
	}{arg1, arg2Copy})
	stub := fake.writeFileStub
	fakeReturns := fake.writeFileReturns
	fake.recordInvocation("writeFile", []interface{}{arg1, arg2Copy})
	fake.writeFileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) WriteFileCallCount() int {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return len(fake.writeFileArgsForCall)
}

func (fake *FakeImpl) WriteFileCalls(stub func(string, []byte) error) {
	fake.writeFileMutex.Lock()
	defer fake.writeFileMutex.Unlock()
	fake.writeFileStub = stub
}

func (fake *FakeImpl) WriteFileArgsForCall(i int) (string, []byte) {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	argsForCall := fake.writeFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) WriteFileReturns(result1 error) {
	fake.writeFileMutex.Lock()
	defer fake.writeFileMutex.Unlock()
	fake.writeFileStub = nil
	fake.writeFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) WriteFileReturnsOnCall(i int, result1 error) {
	fake.writeFileMutex.Lock()
	defer fake.writeFileMutex.Unlock()
	fake.writeFileStub = nil
	if fake.writeFileReturnsOnCall == nil {
		fake.writeFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.docBuilderMutex.RLock()
	defer fake.docBuilderMutex.RUnlock()
	fake.spdxClientMutex.RLock()
	defer fake.spdxClientMutex.RUnlock()
	fake.tmpFileMutex.RLock()
	defer fake.tmpFileMutex.RUnlock()
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
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