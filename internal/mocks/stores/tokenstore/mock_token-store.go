// Code generated by MockGen. DO NOT EDIT.
// Source: echo-starter/internal/contracts/stores/tokenstore (interfaces: ITokenStore,IInternalTokenStore)

// Package tokenstore is a generated GoMock package.
package tokenstore

import (
	context "context"
	models "echo-starter/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITokenStore is a mock of ITokenStore interface.
type MockITokenStore struct {
	ctrl     *gomock.Controller
	recorder *MockITokenStoreMockRecorder
}

// MockITokenStoreMockRecorder is the mock recorder for MockITokenStore.
type MockITokenStoreMockRecorder struct {
	mock *MockITokenStore
}

// NewMockITokenStore creates a new mock instance.
func NewMockITokenStore(ctrl *gomock.Controller) *MockITokenStore {
	mock := &MockITokenStore{ctrl: ctrl}
	mock.recorder = &MockITokenStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITokenStore) EXPECT() *MockITokenStoreMockRecorder {
	return m.recorder
}

// GetToken mocks base method.
func (m *MockITokenStore) GetToken(arg0 context.Context, arg1 string) (*models.TokenInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", arg0, arg1)
	ret0, _ := ret[0].(*models.TokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockITokenStoreMockRecorder) GetToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockITokenStore)(nil).GetToken), arg0, arg1)
}

// RemoveToken mocks base method.
func (m *MockITokenStore) RemoveToken(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveToken indicates an expected call of RemoveToken.
func (mr *MockITokenStoreMockRecorder) RemoveToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveToken", reflect.TypeOf((*MockITokenStore)(nil).RemoveToken), arg0, arg1)
}

// RemoveTokenByClientID mocks base method.
func (m *MockITokenStore) RemoveTokenByClientID(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTokenByClientID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTokenByClientID indicates an expected call of RemoveTokenByClientID.
func (mr *MockITokenStoreMockRecorder) RemoveTokenByClientID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTokenByClientID", reflect.TypeOf((*MockITokenStore)(nil).RemoveTokenByClientID), arg0, arg1)
}

// RemoveTokenByClientIdAndSubject mocks base method.
func (m *MockITokenStore) RemoveTokenByClientIdAndSubject(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTokenByClientIdAndSubject", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTokenByClientIdAndSubject indicates an expected call of RemoveTokenByClientIdAndSubject.
func (mr *MockITokenStoreMockRecorder) RemoveTokenByClientIdAndSubject(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTokenByClientIdAndSubject", reflect.TypeOf((*MockITokenStore)(nil).RemoveTokenByClientIdAndSubject), arg0, arg1, arg2)
}

// RemoveTokenBySubject mocks base method.
func (m *MockITokenStore) RemoveTokenBySubject(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTokenBySubject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTokenBySubject indicates an expected call of RemoveTokenBySubject.
func (mr *MockITokenStoreMockRecorder) RemoveTokenBySubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTokenBySubject", reflect.TypeOf((*MockITokenStore)(nil).RemoveTokenBySubject), arg0, arg1)
}

// StoreToken mocks base method.
func (m *MockITokenStore) StoreToken(arg0 context.Context, arg1 *models.TokenInfo) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreToken indicates an expected call of StoreToken.
func (mr *MockITokenStoreMockRecorder) StoreToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreToken", reflect.TypeOf((*MockITokenStore)(nil).StoreToken), arg0, arg1)
}

// UpdateToken mocks base method.
func (m *MockITokenStore) UpdateToken(arg0 context.Context, arg1 string, arg2 *models.TokenInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToken indicates an expected call of UpdateToken.
func (mr *MockITokenStoreMockRecorder) UpdateToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockITokenStore)(nil).UpdateToken), arg0, arg1, arg2)
}

// MockIInternalTokenStore is a mock of IInternalTokenStore interface.
type MockIInternalTokenStore struct {
	ctrl     *gomock.Controller
	recorder *MockIInternalTokenStoreMockRecorder
}

// MockIInternalTokenStoreMockRecorder is the mock recorder for MockIInternalTokenStore.
type MockIInternalTokenStoreMockRecorder struct {
	mock *MockIInternalTokenStore
}

// NewMockIInternalTokenStore creates a new mock instance.
func NewMockIInternalTokenStore(ctrl *gomock.Controller) *MockIInternalTokenStore {
	mock := &MockIInternalTokenStore{ctrl: ctrl}
	mock.recorder = &MockIInternalTokenStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIInternalTokenStore) EXPECT() *MockIInternalTokenStoreMockRecorder {
	return m.recorder
}

// GetToken mocks base method.
func (m *MockIInternalTokenStore) GetToken(arg0 context.Context, arg1 string) (*models.TokenInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", arg0, arg1)
	ret0, _ := ret[0].(*models.TokenInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockIInternalTokenStoreMockRecorder) GetToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockIInternalTokenStore)(nil).GetToken), arg0, arg1)
}

// RemoveExpired mocks base method.
func (m *MockIInternalTokenStore) RemoveExpired(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveExpired", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveExpired indicates an expected call of RemoveExpired.
func (mr *MockIInternalTokenStoreMockRecorder) RemoveExpired(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveExpired", reflect.TypeOf((*MockIInternalTokenStore)(nil).RemoveExpired), arg0)
}

// RemoveToken mocks base method.
func (m *MockIInternalTokenStore) RemoveToken(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveToken indicates an expected call of RemoveToken.
func (mr *MockIInternalTokenStoreMockRecorder) RemoveToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveToken", reflect.TypeOf((*MockIInternalTokenStore)(nil).RemoveToken), arg0, arg1)
}

// RemoveTokenByClientID mocks base method.
func (m *MockIInternalTokenStore) RemoveTokenByClientID(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTokenByClientID", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTokenByClientID indicates an expected call of RemoveTokenByClientID.
func (mr *MockIInternalTokenStoreMockRecorder) RemoveTokenByClientID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTokenByClientID", reflect.TypeOf((*MockIInternalTokenStore)(nil).RemoveTokenByClientID), arg0, arg1)
}

// RemoveTokenByClientIdAndSubject mocks base method.
func (m *MockIInternalTokenStore) RemoveTokenByClientIdAndSubject(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTokenByClientIdAndSubject", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTokenByClientIdAndSubject indicates an expected call of RemoveTokenByClientIdAndSubject.
func (mr *MockIInternalTokenStoreMockRecorder) RemoveTokenByClientIdAndSubject(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTokenByClientIdAndSubject", reflect.TypeOf((*MockIInternalTokenStore)(nil).RemoveTokenByClientIdAndSubject), arg0, arg1, arg2)
}

// RemoveTokenBySubject mocks base method.
func (m *MockIInternalTokenStore) RemoveTokenBySubject(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTokenBySubject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTokenBySubject indicates an expected call of RemoveTokenBySubject.
func (mr *MockIInternalTokenStoreMockRecorder) RemoveTokenBySubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTokenBySubject", reflect.TypeOf((*MockIInternalTokenStore)(nil).RemoveTokenBySubject), arg0, arg1)
}

// StoreToken mocks base method.
func (m *MockIInternalTokenStore) StoreToken(arg0 context.Context, arg1 *models.TokenInfo) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreToken indicates an expected call of StoreToken.
func (mr *MockIInternalTokenStoreMockRecorder) StoreToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreToken", reflect.TypeOf((*MockIInternalTokenStore)(nil).StoreToken), arg0, arg1)
}

// UpdateToken mocks base method.
func (m *MockIInternalTokenStore) UpdateToken(arg0 context.Context, arg1 string, arg2 *models.TokenInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateToken indicates an expected call of UpdateToken.
func (mr *MockIInternalTokenStoreMockRecorder) UpdateToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockIInternalTokenStore)(nil).UpdateToken), arg0, arg1, arg2)
}
