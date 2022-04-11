// Code generated by MockGen. DO NOT EDIT.
// Source: echo-starter/internal/contracts/clients (interfaces: IClientStore,IClientRequest,IClientRequestInternal)

// Package clients is a generated GoMock package.
package clients

import (
	context "context"
	models "echo-starter/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIClientStore is a mock of IClientStore interface.
type MockIClientStore struct {
	ctrl     *gomock.Controller
	recorder *MockIClientStoreMockRecorder
}

// MockIClientStoreMockRecorder is the mock recorder for MockIClientStore.
type MockIClientStoreMockRecorder struct {
	mock *MockIClientStore
}

// NewMockIClientStore creates a new mock instance.
func NewMockIClientStore(ctrl *gomock.Controller) *MockIClientStore {
	mock := &MockIClientStore{ctrl: ctrl}
	mock.recorder = &MockIClientStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClientStore) EXPECT() *MockIClientStoreMockRecorder {
	return m.recorder
}

// GetClient mocks base method.
func (m *MockIClientStore) GetClient(arg0 context.Context, arg1 string) (*models.Client, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient", arg0, arg1)
	ret0, _ := ret[0].(*models.Client)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetClient indicates an expected call of GetClient.
func (mr *MockIClientStoreMockRecorder) GetClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockIClientStore)(nil).GetClient), arg0, arg1)
}

// MockIClientRequest is a mock of IClientRequest interface.
type MockIClientRequest struct {
	ctrl     *gomock.Controller
	recorder *MockIClientRequestMockRecorder
}

// MockIClientRequestMockRecorder is the mock recorder for MockIClientRequest.
type MockIClientRequestMockRecorder struct {
	mock *MockIClientRequest
}

// NewMockIClientRequest creates a new mock instance.
func NewMockIClientRequest(ctrl *gomock.Controller) *MockIClientRequest {
	mock := &MockIClientRequest{ctrl: ctrl}
	mock.recorder = &MockIClientRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClientRequest) EXPECT() *MockIClientRequestMockRecorder {
	return m.recorder
}

// GetClient mocks base method.
func (m *MockIClientRequest) GetClient() *models.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(*models.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockIClientRequestMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockIClientRequest)(nil).GetClient))
}

// MockIClientRequestInternal is a mock of IClientRequestInternal interface.
type MockIClientRequestInternal struct {
	ctrl     *gomock.Controller
	recorder *MockIClientRequestInternalMockRecorder
}

// MockIClientRequestInternalMockRecorder is the mock recorder for MockIClientRequestInternal.
type MockIClientRequestInternalMockRecorder struct {
	mock *MockIClientRequestInternal
}

// NewMockIClientRequestInternal creates a new mock instance.
func NewMockIClientRequestInternal(ctrl *gomock.Controller) *MockIClientRequestInternal {
	mock := &MockIClientRequestInternal{ctrl: ctrl}
	mock.recorder = &MockIClientRequestInternalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClientRequestInternal) EXPECT() *MockIClientRequestInternalMockRecorder {
	return m.recorder
}

// GetClient mocks base method.
func (m *MockIClientRequestInternal) GetClient() *models.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(*models.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockIClientRequestInternalMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockIClientRequestInternal)(nil).GetClient))
}

// SetClient mocks base method.
func (m *MockIClientRequestInternal) SetClient(arg0 *models.Client) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClient", arg0)
}

// SetClient indicates an expected call of SetClient.
func (mr *MockIClientRequestInternalMockRecorder) SetClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClient", reflect.TypeOf((*MockIClientRequestInternal)(nil).SetClient), arg0)
}
