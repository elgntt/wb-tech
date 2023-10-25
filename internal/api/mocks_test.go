// Code generated by MockGen. DO NOT EDIT.
// Source: internal/api/deps.go

// Package api is a generated GoMock package.
package api

import (
	reflect "reflect"
	model "wb-tech/internal/model"

	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetFromCache mocks base method.
func (m *MockService) GetFromCache(uid string) (model.OrderData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFromCache", uid)
	ret0, _ := ret[0].(model.OrderData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromCache indicates an expected call of GetFromCache.
func (mr *MockServiceMockRecorder) GetFromCache(uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromCache", reflect.TypeOf((*MockService)(nil).GetFromCache), uid)
}
