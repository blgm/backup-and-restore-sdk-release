// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/s3-blobstore-backup-restore/incremental"
)

type FakeBackupDirectoryFinder struct {
	ListBlobsStub        func() ([]incremental.BackedUpBlob, error)
	listBlobsMutex       sync.RWMutex
	listBlobsArgsForCall []struct{}
	listBlobsReturns     struct {
		result1 []incremental.BackedUpBlob
		result2 error
	}
	listBlobsReturnsOnCall map[int]struct {
		result1 []incremental.BackedUpBlob
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBackupDirectoryFinder) ListBlobs() ([]incremental.BackedUpBlob, error) {
	fake.listBlobsMutex.Lock()
	ret, specificReturn := fake.listBlobsReturnsOnCall[len(fake.listBlobsArgsForCall)]
	fake.listBlobsArgsForCall = append(fake.listBlobsArgsForCall, struct{}{})
	fake.recordInvocation("ListBlobs", []interface{}{})
	fake.listBlobsMutex.Unlock()
	if fake.ListBlobsStub != nil {
		return fake.ListBlobsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listBlobsReturns.result1, fake.listBlobsReturns.result2
}

func (fake *FakeBackupDirectoryFinder) ListBlobsCallCount() int {
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	return len(fake.listBlobsArgsForCall)
}

func (fake *FakeBackupDirectoryFinder) ListBlobsReturns(result1 []incremental.BackedUpBlob, result2 error) {
	fake.ListBlobsStub = nil
	fake.listBlobsReturns = struct {
		result1 []incremental.BackedUpBlob
		result2 error
	}{result1, result2}
}

func (fake *FakeBackupDirectoryFinder) ListBlobsReturnsOnCall(i int, result1 []incremental.BackedUpBlob, result2 error) {
	fake.ListBlobsStub = nil
	if fake.listBlobsReturnsOnCall == nil {
		fake.listBlobsReturnsOnCall = make(map[int]struct {
			result1 []incremental.BackedUpBlob
			result2 error
		})
	}
	fake.listBlobsReturnsOnCall[i] = struct {
		result1 []incremental.BackedUpBlob
		result2 error
	}{result1, result2}
}

func (fake *FakeBackupDirectoryFinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBackupDirectoryFinder) recordInvocation(key string, args []interface{}) {
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

var _ incremental.BackupDirectoryFinder = new(FakeBackupDirectoryFinder)
