// Code generated by MockGen. DO NOT EDIT.
// Source: internal/storage/memstorage/mem_storage.go

// Package mock_memstorage is a generated GoMock package.
package mock_memstorage

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/vasiliyantufev/go-advanced-devops/internal/model"
)

// MockMemStorages is a mock of MemStorages interface.
type MockMemStorages struct {
	ctrl     *gomock.Controller
	recorder *MockMemStoragesMockRecorder
}

// MockMemStoragesMockRecorder is the mock recorder for MockMemStorages.
type MockMemStoragesMockRecorder struct {
	mock *MockMemStorages
}

// NewMockMemStorages creates a new mock instance.
func NewMockMemStorages(ctrl *gomock.Controller) *MockMemStorages {
	mock := &MockMemStorages{ctrl: ctrl}
	mock.recorder = &MockMemStoragesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMemStorages) EXPECT() *MockMemStoragesMockRecorder {
	return m.recorder
}

// GetAllMetrics mocks base method.
func (m *MockMemStorages) GetAllMetrics() []models.Metric {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMetrics")
	ret0, _ := ret[0].([]models.Metric)
	return ret0
}

// GetAllMetrics indicates an expected call of GetAllMetrics.
func (mr *MockMemStoragesMockRecorder) GetAllMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMetrics", reflect.TypeOf((*MockMemStorages)(nil).GetAllMetrics))
}

// GetMetricsCount mocks base method.
func (m *MockMemStorages) GetMetricsCount(id string) (int64, string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetricsCount", id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// GetMetricsCount indicates an expected call of GetMetricsCount.
func (mr *MockMemStoragesMockRecorder) GetMetricsCount(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricsCount", reflect.TypeOf((*MockMemStorages)(nil).GetMetricsCount), id)
}

// GetMetricsGauge mocks base method.
func (m *MockMemStorages) GetMetricsGauge(id string) (float64, string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetricsGauge", id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// GetMetricsGauge indicates an expected call of GetMetricsGauge.
func (mr *MockMemStoragesMockRecorder) GetMetricsGauge(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricsGauge", reflect.TypeOf((*MockMemStorages)(nil).GetMetricsGauge), id)
}

// PutMetricsCount mocks base method.
func (m *MockMemStorages) PutMetricsCount(id string, o int64, h string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutMetricsCount", id, o, h)
}

// PutMetricsCount indicates an expected call of PutMetricsCount.
func (mr *MockMemStoragesMockRecorder) PutMetricsCount(id, o, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMetricsCount", reflect.TypeOf((*MockMemStorages)(nil).PutMetricsCount), id, o, h)
}

// PutMetricsGauge mocks base method.
func (m *MockMemStorages) PutMetricsGauge(id string, o float64, h string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutMetricsGauge", id, o, h)
}

// PutMetricsGauge indicates an expected call of PutMetricsGauge.
func (mr *MockMemStoragesMockRecorder) PutMetricsGauge(id, o, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMetricsGauge", reflect.TypeOf((*MockMemStorages)(nil).PutMetricsGauge), id, o, h)
}
