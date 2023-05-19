// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/filestorage/file_storage.go

// Package mock_filestorage is a generated GoMock package.
package mock_filestorage

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/vasiliyantufev/go-advanced-devops/internal/model"
	memstorage "github.com/vasiliyantufev/go-advanced-devops/internal/storage/memstorage"
)

// MockFileStorages is a mock of FileStorages interface.
type MockFileStorages struct {
	ctrl     *gomock.Controller
	recorder *MockFileStoragesMockRecorder
}

// MockFileStoragesMockRecorder is the mock recorder for MockFileStorages.
type MockFileStoragesMockRecorder struct {
	mock *MockFileStorages
}

// NewMockFileStorages creates a new mock instance.
func NewMockFileStorages(ctrl *gomock.Controller) *MockFileStorages {
	mock := &MockFileStorages{ctrl: ctrl}
	mock.recorder = &MockFileStoragesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStorages) EXPECT() *MockFileStoragesMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockFileStorages) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockFileStoragesMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFileStorages)(nil).Close))
}

// FileRestore mocks base method.
func (m *MockFileStorages) FileRestore(mem *memstorage.MemStorage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FileRestore", mem)
}

// FileRestore indicates an expected call of FileRestore.
func (mr *MockFileStoragesMockRecorder) FileRestore(mem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileRestore", reflect.TypeOf((*MockFileStorages)(nil).FileRestore), mem)
}

// FileStore mocks base method.
func (m *MockFileStorages) FileStore(mem *memstorage.MemStorage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FileStore", mem)
}

// FileStore indicates an expected call of FileStore.
func (mr *MockFileStoragesMockRecorder) FileStore(mem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileStore", reflect.TypeOf((*MockFileStorages)(nil).FileStore), mem)
}

// ReadMetric mocks base method.
func (m *MockFileStorages) ReadMetric() (*models.Metric, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMetric")
	ret0, _ := ret[0].(*models.Metric)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMetric indicates an expected call of ReadMetric.
func (mr *MockFileStoragesMockRecorder) ReadMetric() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMetric", reflect.TypeOf((*MockFileStorages)(nil).ReadMetric))
}

// RestoreMetricsFromFile mocks base method.
func (m *MockFileStorages) RestoreMetricsFromFile(mem *memstorage.MemStorage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RestoreMetricsFromFile", mem)
}

// RestoreMetricsFromFile indicates an expected call of RestoreMetricsFromFile.
func (mr *MockFileStoragesMockRecorder) RestoreMetricsFromFile(mem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreMetricsFromFile", reflect.TypeOf((*MockFileStorages)(nil).RestoreMetricsFromFile), mem)
}

// StoreMetricsToFile mocks base method.
func (m *MockFileStorages) StoreMetricsToFile(mem *memstorage.MemStorage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StoreMetricsToFile", mem)
}

// StoreMetricsToFile indicates an expected call of StoreMetricsToFile.
func (mr *MockFileStoragesMockRecorder) StoreMetricsToFile(mem interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreMetricsToFile", reflect.TypeOf((*MockFileStorages)(nil).StoreMetricsToFile), mem)
}

// WriteMetric mocks base method.
func (m *MockFileStorages) WriteMetric(event *models.Metric) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteMetric", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteMetric indicates an expected call of WriteMetric.
func (mr *MockFileStoragesMockRecorder) WriteMetric(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteMetric", reflect.TypeOf((*MockFileStorages)(nil).WriteMetric), event)
}
