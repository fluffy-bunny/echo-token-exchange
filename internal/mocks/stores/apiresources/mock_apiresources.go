// Code generated by MockGen. DO NOT EDIT.
// Source: echo-starter/internal/contracts/stores/apiresources (interfaces: IAPIResources)

// Package apiresources is a generated GoMock package.
package apiresources

import (
	models "echo-starter/internal/models"
	reflect "reflect"

	hashset "github.com/fluffy-bunny/fluffycore/gods/sets/hashset"
	gomock "github.com/golang/mock/gomock"
)

// MockIAPIResources is a mock of IAPIResources interface.
type MockIAPIResources struct {
	ctrl     *gomock.Controller
	recorder *MockIAPIResourcesMockRecorder
}

// MockIAPIResourcesMockRecorder is the mock recorder for MockIAPIResources.
type MockIAPIResourcesMockRecorder struct {
	mock *MockIAPIResources
}

// NewMockIAPIResources creates a new mock instance.
func NewMockIAPIResources(ctrl *gomock.Controller) *MockIAPIResources {
	mock := &MockIAPIResources{ctrl: ctrl}
	mock.recorder = &MockIAPIResourcesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAPIResources) EXPECT() *MockIAPIResourcesMockRecorder {
	return m.recorder
}

// GetAPIResource mocks base method.
func (m *MockIAPIResources) GetAPIResource(arg0 string) (*models.APIResource, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIResource", arg0)
	ret0, _ := ret[0].(*models.APIResource)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAPIResource indicates an expected call of GetAPIResource.
func (mr *MockIAPIResourcesMockRecorder) GetAPIResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIResource", reflect.TypeOf((*MockIAPIResources)(nil).GetAPIResource), arg0)
}

// GetAPIResources mocks base method.
func (m *MockIAPIResources) GetAPIResources() ([]models.APIResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIResources")
	ret0, _ := ret[0].([]models.APIResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIResources indicates an expected call of GetAPIResources.
func (mr *MockIAPIResourcesMockRecorder) GetAPIResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIResources", reflect.TypeOf((*MockIAPIResources)(nil).GetAPIResources))
}

// GetApiResourceByScope mocks base method.
func (m *MockIAPIResources) GetApiResourceByScope(arg0 string) (*models.APIResource, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApiResourceByScope", arg0)
	ret0, _ := ret[0].(*models.APIResource)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetApiResourceByScope indicates an expected call of GetApiResourceByScope.
func (mr *MockIAPIResourcesMockRecorder) GetApiResourceByScope(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiResourceByScope", reflect.TypeOf((*MockIAPIResources)(nil).GetApiResourceByScope), arg0)
}

// GetApiResourceScopes mocks base method.
func (m *MockIAPIResources) GetApiResourceScopes() (*hashset.StringSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApiResourceScopes")
	ret0, _ := ret[0].(*hashset.StringSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApiResourceScopes indicates an expected call of GetApiResourceScopes.
func (mr *MockIAPIResourcesMockRecorder) GetApiResourceScopes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiResourceScopes", reflect.TypeOf((*MockIAPIResources)(nil).GetApiResourceScopes))
}
