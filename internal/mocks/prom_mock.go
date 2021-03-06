// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-link-api/internal/metrics (interfaces: Metrics)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// CreateSuccessResponseCounter mocks base method.
func (m *MockMetrics) CreateSuccessResponseCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateSuccessResponseCounter")
}

// CreateSuccessResponseCounter indicates an expected call of CreateSuccessResponseCounter.
func (mr *MockMetricsMockRecorder) CreateSuccessResponseCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSuccessResponseCounter", reflect.TypeOf((*MockMetrics)(nil).CreateSuccessResponseCounter))
}

// DescribeSuccessResponseCounter mocks base method.
func (m *MockMetrics) DescribeSuccessResponseCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DescribeSuccessResponseCounter")
}

// DescribeSuccessResponseCounter indicates an expected call of DescribeSuccessResponseCounter.
func (mr *MockMetricsMockRecorder) DescribeSuccessResponseCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSuccessResponseCounter", reflect.TypeOf((*MockMetrics)(nil).DescribeSuccessResponseCounter))
}

// ListSuccessResponseCounter mocks base method.
func (m *MockMetrics) ListSuccessResponseCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ListSuccessResponseCounter")
}

// ListSuccessResponseCounter indicates an expected call of ListSuccessResponseCounter.
func (mr *MockMetricsMockRecorder) ListSuccessResponseCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSuccessResponseCounter", reflect.TypeOf((*MockMetrics)(nil).ListSuccessResponseCounter))
}

// MultiCreateSuccessResponseCounter mocks base method.
func (m *MockMetrics) MultiCreateSuccessResponseCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MultiCreateSuccessResponseCounter")
}

// MultiCreateSuccessResponseCounter indicates an expected call of MultiCreateSuccessResponseCounter.
func (mr *MockMetricsMockRecorder) MultiCreateSuccessResponseCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiCreateSuccessResponseCounter", reflect.TypeOf((*MockMetrics)(nil).MultiCreateSuccessResponseCounter))
}

// RemoveSuccessResponseCounter mocks base method.
func (m *MockMetrics) RemoveSuccessResponseCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveSuccessResponseCounter")
}

// RemoveSuccessResponseCounter indicates an expected call of RemoveSuccessResponseCounter.
func (mr *MockMetricsMockRecorder) RemoveSuccessResponseCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSuccessResponseCounter", reflect.TypeOf((*MockMetrics)(nil).RemoveSuccessResponseCounter))
}

// UpdateSuccessResponseCounter mocks base method.
func (m *MockMetrics) UpdateSuccessResponseCounter() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateSuccessResponseCounter")
}

// UpdateSuccessResponseCounter indicates an expected call of UpdateSuccessResponseCounter.
func (mr *MockMetricsMockRecorder) UpdateSuccessResponseCounter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSuccessResponseCounter", reflect.TypeOf((*MockMetrics)(nil).UpdateSuccessResponseCounter))
}
