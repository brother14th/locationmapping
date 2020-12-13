// Code generated by MockGen. DO NOT EDIT.
// Source: locationrepository.go

// Package db is a generated GoMock package.
package db

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLocationRepository is a mock of LocationRepository interface.
type MockLocationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLocationRepositoryMockRecorder
}

// MockLocationRepositoryMockRecorder is the mock recorder for MockLocationRepository.
type MockLocationRepositoryMockRecorder struct {
	mock *MockLocationRepository
}

// NewMockLocationRepository creates a new mock instance.
func NewMockLocationRepository(ctrl *gomock.Controller) *MockLocationRepository {
	mock := &MockLocationRepository{ctrl: ctrl}
	mock.recorder = &MockLocationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocationRepository) EXPECT() *MockLocationRepositoryMockRecorder {
	return m.recorder
}

// GetLocationReport mocks base method.
func (m *MockLocationRepository) GetLocationReport(location string) (LocationReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLocationReport", location)
	ret0, _ := ret[0].(LocationReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLocationReport indicates an expected call of GetLocationReport.
func (mr *MockLocationRepositoryMockRecorder) GetLocationReport(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLocationReport", reflect.TypeOf((*MockLocationRepository)(nil).GetLocationReport), location)
}