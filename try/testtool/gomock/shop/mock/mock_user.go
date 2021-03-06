// Code generated by MockGen. DO NOT EDIT.
// Source: usermock.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGuest is a mock of Guest interface
type MockGuest struct {
	ctrl     *gomock.Controller
	recorder *MockGuestMockRecorder
}

// MockGuestMockRecorder is the mock recorder for MockGuest
type MockGuestMockRecorder struct {
	mock *MockGuest
}

// NewMockGuest creates a new mock instance
func NewMockGuest(ctrl *gomock.Controller) *MockGuest {
	mock := &MockGuest{ctrl: ctrl}
	mock.recorder = &MockGuestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGuest) EXPECT() *MockGuestMockRecorder {
	return m.recorder
}

// Shopping mocks base method
func (m *MockGuest) Shopping(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shopping", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// Shopping indicates an expected call of Shopping
func (mr *MockGuestMockRecorder) Shopping(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shopping", reflect.TypeOf((*MockGuest)(nil).Shopping), name)
}
