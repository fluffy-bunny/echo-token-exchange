// Code generated by MockGen. DO NOT EDIT.
// Source: echo-starter/internal/contracts/background/tasks (interfaces: ISingletonTask,ITaskClient,ITaskEngineFactory)

// Package tasks is a generated GoMock package.
package tasks

import (
	context "context"
	reflect "reflect"

	hashset "github.com/fluffy-bunny/grpcdotnetgo/pkg/gods/sets/hashset"
	gomock "github.com/golang/mock/gomock"
	asynq "github.com/hibiken/asynq"
)

// MockISingletonTask is a mock of ISingletonTask interface.
type MockISingletonTask struct {
	ctrl     *gomock.Controller
	recorder *MockISingletonTaskMockRecorder
}

// MockISingletonTaskMockRecorder is the mock recorder for MockISingletonTask.
type MockISingletonTaskMockRecorder struct {
	mock *MockISingletonTask
}

// NewMockISingletonTask creates a new mock instance.
func NewMockISingletonTask(ctrl *gomock.Controller) *MockISingletonTask {
	mock := &MockISingletonTask{ctrl: ctrl}
	mock.recorder = &MockISingletonTaskMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISingletonTask) EXPECT() *MockISingletonTaskMockRecorder {
	return m.recorder
}

// EnqueTask mocks base method.
func (m *MockISingletonTask) EnqueTask(arg0 interface{}, arg1 ...asynq.Option) (*asynq.TaskInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnqueTask", varargs...)
	ret0, _ := ret[0].(*asynq.TaskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnqueTask indicates an expected call of EnqueTask.
func (mr *MockISingletonTaskMockRecorder) EnqueTask(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueTask", reflect.TypeOf((*MockISingletonTask)(nil).EnqueTask), varargs...)
}

// GetPatterns mocks base method.
func (m *MockISingletonTask) GetPatterns() *hashset.StringSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPatterns")
	ret0, _ := ret[0].(*hashset.StringSet)
	return ret0
}

// GetPatterns indicates an expected call of GetPatterns.
func (mr *MockISingletonTaskMockRecorder) GetPatterns() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPatterns", reflect.TypeOf((*MockISingletonTask)(nil).GetPatterns))
}

// ProcessTask mocks base method.
func (m *MockISingletonTask) ProcessTask(arg0 context.Context, arg1 *asynq.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessTask indicates an expected call of ProcessTask.
func (mr *MockISingletonTaskMockRecorder) ProcessTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTask", reflect.TypeOf((*MockISingletonTask)(nil).ProcessTask), arg0, arg1)
}

// MockITaskClient is a mock of ITaskClient interface.
type MockITaskClient struct {
	ctrl     *gomock.Controller
	recorder *MockITaskClientMockRecorder
}

// MockITaskClientMockRecorder is the mock recorder for MockITaskClient.
type MockITaskClientMockRecorder struct {
	mock *MockITaskClient
}

// NewMockITaskClient creates a new mock instance.
func NewMockITaskClient(ctrl *gomock.Controller) *MockITaskClient {
	mock := &MockITaskClient{ctrl: ctrl}
	mock.recorder = &MockITaskClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITaskClient) EXPECT() *MockITaskClientMockRecorder {
	return m.recorder
}

// EnqueTask mocks base method.
func (m *MockITaskClient) EnqueTask(arg0 *asynq.Task, arg1 ...asynq.Option) (*asynq.TaskInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnqueTask", varargs...)
	ret0, _ := ret[0].(*asynq.TaskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnqueTask indicates an expected call of EnqueTask.
func (mr *MockITaskClientMockRecorder) EnqueTask(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnqueTask", reflect.TypeOf((*MockITaskClient)(nil).EnqueTask), varargs...)
}

// MockITaskEngineFactory is a mock of ITaskEngineFactory interface.
type MockITaskEngineFactory struct {
	ctrl     *gomock.Controller
	recorder *MockITaskEngineFactoryMockRecorder
}

// MockITaskEngineFactoryMockRecorder is the mock recorder for MockITaskEngineFactory.
type MockITaskEngineFactoryMockRecorder struct {
	mock *MockITaskEngineFactory
}

// NewMockITaskEngineFactory creates a new mock instance.
func NewMockITaskEngineFactory(ctrl *gomock.Controller) *MockITaskEngineFactory {
	mock := &MockITaskEngineFactory{ctrl: ctrl}
	mock.recorder = &MockITaskEngineFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITaskEngineFactory) EXPECT() *MockITaskEngineFactoryMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockITaskEngineFactory) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockITaskEngineFactoryMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockITaskEngineFactory)(nil).Start))
}

// Stop mocks base method.
func (m *MockITaskEngineFactory) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockITaskEngineFactoryMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockITaskEngineFactory)(nil).Stop))
}
