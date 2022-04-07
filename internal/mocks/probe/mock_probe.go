// Code generated by MockGen. DO NOT EDIT.
// Source: echo-starter/internal/contracts/probe (interfaces: IProbe)

// Package probe is a generated GoMock package.
package probe

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIProbe is a mock of IProbe interface.
type MockIProbe struct {
	ctrl     *gomock.Controller
	recorder *MockIProbeMockRecorder
}

// MockIProbeMockRecorder is the mock recorder for MockIProbe.
type MockIProbeMockRecorder struct {
	mock *MockIProbe
}

// NewMockIProbe creates a new mock instance.
func NewMockIProbe(ctrl *gomock.Controller) *MockIProbe {
	mock := &MockIProbe{ctrl: ctrl}
	mock.recorder = &MockIProbeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProbe) EXPECT() *MockIProbeMockRecorder {
	return m.recorder
}

// GetName mocks base method.
func (m *MockIProbe) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockIProbeMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockIProbe)(nil).GetName))
}

// Probe mocks base method.
func (m *MockIProbe) Probe() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Probe")
	ret0, _ := ret[0].(error)
	return ret0
}

// Probe indicates an expected call of Probe.
func (mr *MockIProbeMockRecorder) Probe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Probe", reflect.TypeOf((*MockIProbe)(nil).Probe))
}
