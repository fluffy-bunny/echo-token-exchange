// Code generated by MockGen. DO NOT EDIT.
// Source: echo-starter/internal/contracts/background/tasks/removetokens (interfaces: IRemoveTokensSingletonTask)

// Package removetokens is a generated GoMock package.
package removetokens

import (
	context "context"
	removetokens "echo-starter/internal/contracts/background/tasks/removetokens"
	reflect "reflect"

	hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	gomock "github.com/golang/mock/gomock"
	asynq "github.com/hibiken/asynq"
)

// MockIRemoveTokensSingletonTask is a mock of IRemoveTokensSingletonTask interface.
type MockIRemoveTokensSingletonTask struct {
	ctrl     *gomock.Controller
	recorder *MockIRemoveTokensSingletonTaskMockRecorder
}

// MockIRemoveTokensSingletonTaskMockRecorder is the mock recorder for MockIRemoveTokensSingletonTask.
type MockIRemoveTokensSingletonTaskMockRecorder struct {
	mock *MockIRemoveTokensSingletonTask
}

// NewMockIRemoveTokensSingletonTask creates a new mock instance.
func NewMockIRemoveTokensSingletonTask(ctrl *gomock.Controller) *MockIRemoveTokensSingletonTask {
	mock := &MockIRemoveTokensSingletonTask{ctrl: ctrl}
	mock.recorder = &MockIRemoveTokensSingletonTaskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRemoveTokensSingletonTask) EXPECT() *MockIRemoveTokensSingletonTaskMockRecorder {
	return m.recorder
}

// EnqueTask mocks base method.
func (m *MockIRemoveTokensSingletonTask) EnqueTask(arg0 context.Context, arg1 interface{}, arg2 ...asynq.Option) (*asynq.TaskInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnqueTask", varargs...)
	ret0, _ := ret[0].(*asynq.TaskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnqueTask indicates an expected call of EnqueTask.
func (mr *MockIRemoveTokensSingletonTaskMockRecorder) EnqueTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueTask", reflect.TypeOf((*MockIRemoveTokensSingletonTask)(nil).EnqueTask), varargs...)
}

// EnqueTaskTokenRemoveByClientID mocks base method.
func (m *MockIRemoveTokensSingletonTask) EnqueTaskTokenRemoveByClientID(arg0 context.Context, arg1 *removetokens.TokenRemoveByClientID, arg2 ...asynq.Option) (*asynq.TaskInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnqueTaskTokenRemoveByClientID", varargs...)
	ret0, _ := ret[0].(*asynq.TaskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnqueTaskTokenRemoveByClientID indicates an expected call of EnqueTaskTokenRemoveByClientID.
func (mr *MockIRemoveTokensSingletonTaskMockRecorder) EnqueTaskTokenRemoveByClientID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueTaskTokenRemoveByClientID", reflect.TypeOf((*MockIRemoveTokensSingletonTask)(nil).EnqueTaskTokenRemoveByClientID), varargs...)
}

// EnqueTaskTokenRemoveByClientIDAndSubject mocks base method.
func (m *MockIRemoveTokensSingletonTask) EnqueTaskTokenRemoveByClientIDAndSubject(arg0 context.Context, arg1 *removetokens.TokenRemoveByClientIDAndSubject, arg2 ...asynq.Option) (*asynq.TaskInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnqueTaskTokenRemoveByClientIDAndSubject", varargs...)
	ret0, _ := ret[0].(*asynq.TaskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnqueTaskTokenRemoveByClientIDAndSubject indicates an expected call of EnqueTaskTokenRemoveByClientIDAndSubject.
func (mr *MockIRemoveTokensSingletonTaskMockRecorder) EnqueTaskTokenRemoveByClientIDAndSubject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueTaskTokenRemoveByClientIDAndSubject", reflect.TypeOf((*MockIRemoveTokensSingletonTask)(nil).EnqueTaskTokenRemoveByClientIDAndSubject), varargs...)
}

// EnqueTaskTypeRemoveTokenBySubject mocks base method.
func (m *MockIRemoveTokensSingletonTask) EnqueTaskTypeRemoveTokenBySubject(arg0 context.Context, arg1 *removetokens.TokenRemoveBySubject, arg2 ...asynq.Option) (*asynq.TaskInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnqueTaskTypeRemoveTokenBySubject", varargs...)
	ret0, _ := ret[0].(*asynq.TaskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnqueTaskTypeRemoveTokenBySubject indicates an expected call of EnqueTaskTypeRemoveTokenBySubject.
func (mr *MockIRemoveTokensSingletonTaskMockRecorder) EnqueTaskTypeRemoveTokenBySubject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueTaskTypeRemoveTokenBySubject", reflect.TypeOf((*MockIRemoveTokensSingletonTask)(nil).EnqueTaskTypeRemoveTokenBySubject), varargs...)
}

// GetPatterns mocks base method.
func (m *MockIRemoveTokensSingletonTask) GetPatterns() *hashset.StringSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPatterns")
	ret0, _ := ret[0].(*hashset.StringSet)
	return ret0
}

// GetPatterns indicates an expected call of GetPatterns.
func (mr *MockIRemoveTokensSingletonTaskMockRecorder) GetPatterns() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPatterns", reflect.TypeOf((*MockIRemoveTokensSingletonTask)(nil).GetPatterns))
}

// ProcessTask mocks base method.
func (m *MockIRemoveTokensSingletonTask) ProcessTask(arg0 context.Context, arg1 *asynq.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessTask indicates an expected call of ProcessTask.
func (mr *MockIRemoveTokensSingletonTaskMockRecorder) ProcessTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTask", reflect.TypeOf((*MockIRemoveTokensSingletonTask)(nil).ProcessTask), arg0, arg1)
}
