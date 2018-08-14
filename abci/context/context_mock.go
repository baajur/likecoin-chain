// Code generated by MockGen. DO NOT EDIT.
// Source: context.go

// Package context is a generated GoMock package.
package context

import (
	gomock "github.com/golang/mock/gomock"
	iavl "github.com/tendermint/iavl"
	reflect "reflect"
)

// MockContext is a mock of Context interface
type MockContext struct {
	ctrl     *gomock.Controller
	recorder *MockContextMockRecorder
}

// MockContextMockRecorder is the mock recorder for MockContext
type MockContextMockRecorder struct {
	mock *MockContext
}

// NewMockContext creates a new mock instance
func NewMockContext(ctrl *gomock.Controller) *MockContext {
	mock := &MockContext{ctrl: ctrl}
	mock.recorder = &MockContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContext) EXPECT() *MockContextMockRecorder {
	return m.recorder
}

// StateTree mocks base method
func (m *MockContext) StateTree() *iavl.VersionedTree {
	ret := m.ctrl.Call(m, "StateTree")
	ret0, _ := ret[0].(*iavl.VersionedTree)
	return ret0
}

// StateTree indicates an expected call of StateTree
func (mr *MockContextMockRecorder) StateTree() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateTree", reflect.TypeOf((*MockContext)(nil).StateTree))
}

// WithdrawTree mocks base method
func (m *MockContext) WithdrawTree() *iavl.VersionedTree {
	ret := m.ctrl.Call(m, "WithdrawTree")
	ret0, _ := ret[0].(*iavl.VersionedTree)
	return ret0
}

// WithdrawTree indicates an expected call of WithdrawTree
func (mr *MockContextMockRecorder) WithdrawTree() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawTree", reflect.TypeOf((*MockContext)(nil).WithdrawTree))
}