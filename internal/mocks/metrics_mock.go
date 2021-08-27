// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-quiz-api/internal/metrics (interfaces: MetricsReporter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMetricsReporter is a mock of MetricsReporter interface.
type MockMetricsReporter struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsReporterMockRecorder
}

// MockMetricsReporterMockRecorder is the mock recorder for MockMetricsReporter.
type MockMetricsReporterMockRecorder struct {
	mock *MockMetricsReporter
}

// NewMockMetricsReporter creates a new mock instance.
func NewMockMetricsReporter(ctrl *gomock.Controller) *MockMetricsReporter {
	mock := &MockMetricsReporter{ctrl: ctrl}
	mock.recorder = &MockMetricsReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricsReporter) EXPECT() *MockMetricsReporterMockRecorder {
	return m.recorder
}

// IncCreate mocks base method.
func (m *MockMetricsReporter) IncCreate(arg0 uint, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncCreate", arg0, arg1)
}

// IncCreate indicates an expected call of IncCreate.
func (mr *MockMetricsReporterMockRecorder) IncCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncCreate", reflect.TypeOf((*MockMetricsReporter)(nil).IncCreate), arg0, arg1)
}

// IncList mocks base method.
func (m *MockMetricsReporter) IncList(arg0 uint, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncList", arg0, arg1)
}

// IncList indicates an expected call of IncList.
func (mr *MockMetricsReporterMockRecorder) IncList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncList", reflect.TypeOf((*MockMetricsReporter)(nil).IncList), arg0, arg1)
}

// IncRead mocks base method.
func (m *MockMetricsReporter) IncRead(arg0 uint, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncRead", arg0, arg1)
}

// IncRead indicates an expected call of IncRead.
func (mr *MockMetricsReporterMockRecorder) IncRead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncRead", reflect.TypeOf((*MockMetricsReporter)(nil).IncRead), arg0, arg1)
}

// IncRemove mocks base method.
func (m *MockMetricsReporter) IncRemove(arg0 uint, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncRemove", arg0, arg1)
}

// IncRemove indicates an expected call of IncRemove.
func (mr *MockMetricsReporterMockRecorder) IncRemove(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncRemove", reflect.TypeOf((*MockMetricsReporter)(nil).IncRemove), arg0, arg1)
}

// IncUpdate mocks base method.
func (m *MockMetricsReporter) IncUpdate(arg0 uint, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncUpdate", arg0, arg1)
}

// IncUpdate indicates an expected call of IncUpdate.
func (mr *MockMetricsReporterMockRecorder) IncUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncUpdate", reflect.TypeOf((*MockMetricsReporter)(nil).IncUpdate), arg0, arg1)
}
