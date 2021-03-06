// Code generated by MockGen. DO NOT EDIT.
// Source: echo-starter/internal/contracts/tokenhandlers (interfaces: ITokenHandler,IClientCredentialsTokenHandler,IRefreshTokenHandler,ITokenExchangeTokenHandler,ITokenHandlerAccessor,IInternalTokenHandlerAccessor)

// Package tokenhandlers is a generated GoMock package.
package tokenhandlers

import (
	context "context"
	tokenhandlers "echo-starter/internal/contracts/tokenhandlers"
	models "echo-starter/internal/models"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITokenHandler is a mock of ITokenHandler interface.
type MockITokenHandler struct {
	ctrl     *gomock.Controller
	recorder *MockITokenHandlerMockRecorder
}

// MockITokenHandlerMockRecorder is the mock recorder for MockITokenHandler.
type MockITokenHandlerMockRecorder struct {
	mock *MockITokenHandler
}

// NewMockITokenHandler creates a new mock instance.
func NewMockITokenHandler(ctrl *gomock.Controller) *MockITokenHandler {
	mock := &MockITokenHandler{ctrl: ctrl}
	mock.recorder = &MockITokenHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITokenHandler) EXPECT() *MockITokenHandlerMockRecorder {
	return m.recorder
}

// ProcessTokenRequest mocks base method.
func (m *MockITokenHandler) ProcessTokenRequest(arg0 context.Context, arg1 *tokenhandlers.ValidatedTokenRequestResult) (models.IClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTokenRequest", arg0, arg1)
	ret0, _ := ret[0].(models.IClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessTokenRequest indicates an expected call of ProcessTokenRequest.
func (mr *MockITokenHandlerMockRecorder) ProcessTokenRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTokenRequest", reflect.TypeOf((*MockITokenHandler)(nil).ProcessTokenRequest), arg0, arg1)
}

// ValidationTokenRequest mocks base method.
func (m *MockITokenHandler) ValidationTokenRequest(arg0 *http.Request) (*tokenhandlers.ValidatedTokenRequestResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidationTokenRequest", arg0)
	ret0, _ := ret[0].(*tokenhandlers.ValidatedTokenRequestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidationTokenRequest indicates an expected call of ValidationTokenRequest.
func (mr *MockITokenHandlerMockRecorder) ValidationTokenRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidationTokenRequest", reflect.TypeOf((*MockITokenHandler)(nil).ValidationTokenRequest), arg0)
}

// MockIClientCredentialsTokenHandler is a mock of IClientCredentialsTokenHandler interface.
type MockIClientCredentialsTokenHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIClientCredentialsTokenHandlerMockRecorder
}

// MockIClientCredentialsTokenHandlerMockRecorder is the mock recorder for MockIClientCredentialsTokenHandler.
type MockIClientCredentialsTokenHandlerMockRecorder struct {
	mock *MockIClientCredentialsTokenHandler
}

// NewMockIClientCredentialsTokenHandler creates a new mock instance.
func NewMockIClientCredentialsTokenHandler(ctrl *gomock.Controller) *MockIClientCredentialsTokenHandler {
	mock := &MockIClientCredentialsTokenHandler{ctrl: ctrl}
	mock.recorder = &MockIClientCredentialsTokenHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClientCredentialsTokenHandler) EXPECT() *MockIClientCredentialsTokenHandlerMockRecorder {
	return m.recorder
}

// ProcessTokenRequest mocks base method.
func (m *MockIClientCredentialsTokenHandler) ProcessTokenRequest(arg0 context.Context, arg1 *tokenhandlers.ValidatedTokenRequestResult) (models.IClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTokenRequest", arg0, arg1)
	ret0, _ := ret[0].(models.IClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessTokenRequest indicates an expected call of ProcessTokenRequest.
func (mr *MockIClientCredentialsTokenHandlerMockRecorder) ProcessTokenRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTokenRequest", reflect.TypeOf((*MockIClientCredentialsTokenHandler)(nil).ProcessTokenRequest), arg0, arg1)
}

// ValidationTokenRequest mocks base method.
func (m *MockIClientCredentialsTokenHandler) ValidationTokenRequest(arg0 *http.Request) (*tokenhandlers.ValidatedTokenRequestResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidationTokenRequest", arg0)
	ret0, _ := ret[0].(*tokenhandlers.ValidatedTokenRequestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidationTokenRequest indicates an expected call of ValidationTokenRequest.
func (mr *MockIClientCredentialsTokenHandlerMockRecorder) ValidationTokenRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidationTokenRequest", reflect.TypeOf((*MockIClientCredentialsTokenHandler)(nil).ValidationTokenRequest), arg0)
}

// MockIRefreshTokenHandler is a mock of IRefreshTokenHandler interface.
type MockIRefreshTokenHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIRefreshTokenHandlerMockRecorder
}

// MockIRefreshTokenHandlerMockRecorder is the mock recorder for MockIRefreshTokenHandler.
type MockIRefreshTokenHandlerMockRecorder struct {
	mock *MockIRefreshTokenHandler
}

// NewMockIRefreshTokenHandler creates a new mock instance.
func NewMockIRefreshTokenHandler(ctrl *gomock.Controller) *MockIRefreshTokenHandler {
	mock := &MockIRefreshTokenHandler{ctrl: ctrl}
	mock.recorder = &MockIRefreshTokenHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRefreshTokenHandler) EXPECT() *MockIRefreshTokenHandlerMockRecorder {
	return m.recorder
}

// ProcessTokenRequest mocks base method.
func (m *MockIRefreshTokenHandler) ProcessTokenRequest(arg0 context.Context, arg1 *tokenhandlers.ValidatedTokenRequestResult) (models.IClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTokenRequest", arg0, arg1)
	ret0, _ := ret[0].(models.IClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessTokenRequest indicates an expected call of ProcessTokenRequest.
func (mr *MockIRefreshTokenHandlerMockRecorder) ProcessTokenRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTokenRequest", reflect.TypeOf((*MockIRefreshTokenHandler)(nil).ProcessTokenRequest), arg0, arg1)
}

// ValidationTokenRequest mocks base method.
func (m *MockIRefreshTokenHandler) ValidationTokenRequest(arg0 *http.Request) (*tokenhandlers.ValidatedTokenRequestResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidationTokenRequest", arg0)
	ret0, _ := ret[0].(*tokenhandlers.ValidatedTokenRequestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidationTokenRequest indicates an expected call of ValidationTokenRequest.
func (mr *MockIRefreshTokenHandlerMockRecorder) ValidationTokenRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidationTokenRequest", reflect.TypeOf((*MockIRefreshTokenHandler)(nil).ValidationTokenRequest), arg0)
}

// MockITokenExchangeTokenHandler is a mock of ITokenExchangeTokenHandler interface.
type MockITokenExchangeTokenHandler struct {
	ctrl     *gomock.Controller
	recorder *MockITokenExchangeTokenHandlerMockRecorder
}

// MockITokenExchangeTokenHandlerMockRecorder is the mock recorder for MockITokenExchangeTokenHandler.
type MockITokenExchangeTokenHandlerMockRecorder struct {
	mock *MockITokenExchangeTokenHandler
}

// NewMockITokenExchangeTokenHandler creates a new mock instance.
func NewMockITokenExchangeTokenHandler(ctrl *gomock.Controller) *MockITokenExchangeTokenHandler {
	mock := &MockITokenExchangeTokenHandler{ctrl: ctrl}
	mock.recorder = &MockITokenExchangeTokenHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITokenExchangeTokenHandler) EXPECT() *MockITokenExchangeTokenHandlerMockRecorder {
	return m.recorder
}

// ProcessTokenRequest mocks base method.
func (m *MockITokenExchangeTokenHandler) ProcessTokenRequest(arg0 context.Context, arg1 *tokenhandlers.ValidatedTokenRequestResult) (models.IClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessTokenRequest", arg0, arg1)
	ret0, _ := ret[0].(models.IClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessTokenRequest indicates an expected call of ProcessTokenRequest.
func (mr *MockITokenExchangeTokenHandlerMockRecorder) ProcessTokenRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessTokenRequest", reflect.TypeOf((*MockITokenExchangeTokenHandler)(nil).ProcessTokenRequest), arg0, arg1)
}

// ValidationTokenRequest mocks base method.
func (m *MockITokenExchangeTokenHandler) ValidationTokenRequest(arg0 *http.Request) (*tokenhandlers.ValidatedTokenRequestResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidationTokenRequest", arg0)
	ret0, _ := ret[0].(*tokenhandlers.ValidatedTokenRequestResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidationTokenRequest indicates an expected call of ValidationTokenRequest.
func (mr *MockITokenExchangeTokenHandlerMockRecorder) ValidationTokenRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidationTokenRequest", reflect.TypeOf((*MockITokenExchangeTokenHandler)(nil).ValidationTokenRequest), arg0)
}

// MockITokenHandlerAccessor is a mock of ITokenHandlerAccessor interface.
type MockITokenHandlerAccessor struct {
	ctrl     *gomock.Controller
	recorder *MockITokenHandlerAccessorMockRecorder
}

// MockITokenHandlerAccessorMockRecorder is the mock recorder for MockITokenHandlerAccessor.
type MockITokenHandlerAccessorMockRecorder struct {
	mock *MockITokenHandlerAccessor
}

// NewMockITokenHandlerAccessor creates a new mock instance.
func NewMockITokenHandlerAccessor(ctrl *gomock.Controller) *MockITokenHandlerAccessor {
	mock := &MockITokenHandlerAccessor{ctrl: ctrl}
	mock.recorder = &MockITokenHandlerAccessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITokenHandlerAccessor) EXPECT() *MockITokenHandlerAccessorMockRecorder {
	return m.recorder
}

// GetGrantType mocks base method.
func (m *MockITokenHandlerAccessor) GetGrantType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGrantType")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetGrantType indicates an expected call of GetGrantType.
func (mr *MockITokenHandlerAccessorMockRecorder) GetGrantType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrantType", reflect.TypeOf((*MockITokenHandlerAccessor)(nil).GetGrantType))
}

// GetTokenHandler mocks base method.
func (m *MockITokenHandlerAccessor) GetTokenHandler() tokenhandlers.ITokenHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokenHandler")
	ret0, _ := ret[0].(tokenhandlers.ITokenHandler)
	return ret0
}

// GetTokenHandler indicates an expected call of GetTokenHandler.
func (mr *MockITokenHandlerAccessorMockRecorder) GetTokenHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokenHandler", reflect.TypeOf((*MockITokenHandlerAccessor)(nil).GetTokenHandler))
}

// MockIInternalTokenHandlerAccessor is a mock of IInternalTokenHandlerAccessor interface.
type MockIInternalTokenHandlerAccessor struct {
	ctrl     *gomock.Controller
	recorder *MockIInternalTokenHandlerAccessorMockRecorder
}

// MockIInternalTokenHandlerAccessorMockRecorder is the mock recorder for MockIInternalTokenHandlerAccessor.
type MockIInternalTokenHandlerAccessorMockRecorder struct {
	mock *MockIInternalTokenHandlerAccessor
}

// NewMockIInternalTokenHandlerAccessor creates a new mock instance.
func NewMockIInternalTokenHandlerAccessor(ctrl *gomock.Controller) *MockIInternalTokenHandlerAccessor {
	mock := &MockIInternalTokenHandlerAccessor{ctrl: ctrl}
	mock.recorder = &MockIInternalTokenHandlerAccessorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIInternalTokenHandlerAccessor) EXPECT() *MockIInternalTokenHandlerAccessorMockRecorder {
	return m.recorder
}

// SetGrantType mocks base method.
func (m *MockIInternalTokenHandlerAccessor) SetGrantType(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGrantType", arg0)
}

// SetGrantType indicates an expected call of SetGrantType.
func (mr *MockIInternalTokenHandlerAccessorMockRecorder) SetGrantType(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGrantType", reflect.TypeOf((*MockIInternalTokenHandlerAccessor)(nil).SetGrantType), arg0)
}

// SetTokenHandler mocks base method.
func (m *MockIInternalTokenHandlerAccessor) SetTokenHandler(arg0 tokenhandlers.ITokenHandler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTokenHandler", arg0)
}

// SetTokenHandler indicates an expected call of SetTokenHandler.
func (mr *MockIInternalTokenHandlerAccessorMockRecorder) SetTokenHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTokenHandler", reflect.TypeOf((*MockIInternalTokenHandlerAccessor)(nil).SetTokenHandler), arg0)
}
