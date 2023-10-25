// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/deps.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"
	model "wb-tech/internal/model"

	gomock "go.uber.org/mock/gomock"
)

// Mockrepository is a mock of repository interface.
type Mockrepository struct {
	ctrl     *gomock.Controller
	recorder *MockrepositoryMockRecorder
}

// MockrepositoryMockRecorder is the mock recorder for Mockrepository.
type MockrepositoryMockRecorder struct {
	mock *Mockrepository
}

// NewMockrepository creates a new mock instance.
func NewMockrepository(ctrl *gomock.Controller) *Mockrepository {
	mock := &Mockrepository{ctrl: ctrl}
	mock.recorder = &MockrepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockrepository) EXPECT() *MockrepositoryMockRecorder {
	return m.recorder
}

// LoadCache mocks base method.
func (m *Mockrepository) LoadCache(ctx context.Context) (map[string]model.OrderData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadCache", ctx)
	ret0, _ := ret[0].(map[string]model.OrderData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadCache indicates an expected call of LoadCache.
func (mr *MockrepositoryMockRecorder) LoadCache(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadCache", reflect.TypeOf((*Mockrepository)(nil).LoadCache), ctx)
}

// Save mocks base method.
func (m *Mockrepository) Save(ctx context.Context, orderData model.OrderData) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, orderData)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockrepositoryMockRecorder) Save(ctx, orderData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*Mockrepository)(nil).Save), ctx, orderData)
}