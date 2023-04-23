// Code generated by MockGen. DO NOT EDIT.
// Source: internal/db/database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	configserver "github.com/vasiliyantufev/go-advanced-devops/internal/config/configserver"
	memstorage "github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

// MockDBS is a mock of DBS interface.
type MockDBS struct {
	ctrl     *gomock.Controller
	recorder *MockDBSMockRecorder
}

// MockDBSMockRecorder is the mock recorder for MockDBS.
type MockDBSMockRecorder struct {
	mock *MockDBS
}

// NewMockDBS creates a new mock instance.
func NewMockDBS(ctrl *gomock.Controller) *MockDBS {
	mock := &MockDBS{ctrl: ctrl}
	mock.recorder = &MockDBSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBS) EXPECT() *MockDBSMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDBS) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDBSMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDBS)(nil).Close))
}

// CreateTablesMigration mocks base method.
func (m *MockDBS) CreateTablesMigration(cfg *configserver.ConfigServer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateTablesMigration", cfg)
}

// CreateTablesMigration indicates an expected call of CreateTablesMigration.
func (mr *MockDBSMockRecorder) CreateTablesMigration(cfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTablesMigration", reflect.TypeOf((*MockDBS)(nil).CreateTablesMigration), cfg)
}

// InsertOrUpdateMetrics mocks base method.
func (m *MockDBS) InsertOrUpdateMetrics(metrics *memstorage.MemStorage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrUpdateMetrics", metrics)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOrUpdateMetrics indicates an expected call of InsertOrUpdateMetrics.
func (mr *MockDBSMockRecorder) InsertOrUpdateMetrics(metrics interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrUpdateMetrics", reflect.TypeOf((*MockDBS)(nil).InsertOrUpdateMetrics), metrics)
}

// Ping mocks base method.
func (m *MockDBS) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockDBSMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDBS)(nil).Ping))
}
