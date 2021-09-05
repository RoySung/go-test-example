// Code generated by MockGen. DO NOT EDIT.
// Source: redis.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRedisRepository is a mock of RedisRepository interface.
type MockRedisRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRedisRepositoryMockRecorder
}

// MockRedisRepositoryMockRecorder is the mock recorder for MockRedisRepository.
type MockRedisRepositoryMockRecorder struct {
	mock *MockRedisRepository
}

// NewMockRedisRepository creates a new mock instance.
func NewMockRedisRepository(ctrl *gomock.Controller) *MockRedisRepository {
	mock := &MockRedisRepository{ctrl: ctrl}
	mock.recorder = &MockRedisRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisRepository) EXPECT() *MockRedisRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockRedisRepository) Get(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRedisRepositoryMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisRepository)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockRedisRepository) Set(key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisRepositoryMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisRepository)(nil).Set), key, value)
}