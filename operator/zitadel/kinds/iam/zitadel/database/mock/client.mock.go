// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package databasemock is a generated GoMock package.
package databasemock

import (
	mntr "github.com/caos/orbos/mntr"
	kubernetes "github.com/caos/orbos/pkg/kubernetes"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetConnectionInfo mocks base method
func (m *MockClient) GetConnectionInfo(monitor mntr.Monitor, k8sClient kubernetes.ClientInt) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnectionInfo", monitor, k8sClient)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetConnectionInfo indicates an expected call of GetConnectionInfo
func (mr *MockClientMockRecorder) GetConnectionInfo(monitor, k8sClient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectionInfo", reflect.TypeOf((*MockClient)(nil).GetConnectionInfo), monitor, k8sClient)
}

// DeleteUser mocks base method
func (m *MockClient) DeleteUser(monitor mntr.Monitor, user string, k8sClient kubernetes.ClientInt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", monitor, user, k8sClient)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockClientMockRecorder) DeleteUser(monitor, user, k8sClient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockClient)(nil).DeleteUser), monitor, user, k8sClient)
}

// AddUser mocks base method
func (m *MockClient) AddUser(monitor mntr.Monitor, user string, k8sClient kubernetes.ClientInt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", monitor, user, k8sClient)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser
func (mr *MockClientMockRecorder) AddUser(monitor, user, k8sClient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockClient)(nil).AddUser), monitor, user, k8sClient)
}

// ListUsers mocks base method
func (m *MockClient) ListUsers(monitor mntr.Monitor, k8sClient kubernetes.ClientInt) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", monitor, k8sClient)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers
func (mr *MockClientMockRecorder) ListUsers(monitor, k8sClient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockClient)(nil).ListUsers), monitor, k8sClient)
}