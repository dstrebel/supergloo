// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/api/external/istio/authorization/v1alpha1/policy_client.sk.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	clients "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	v1alpha1 "github.com/solo-io/supergloo/pkg/api/external/istio/authorization/v1alpha1"
)

// MockPolicyWatcher is a mock of PolicyWatcher interface
type MockPolicyWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyWatcherMockRecorder
}

// MockPolicyWatcherMockRecorder is the mock recorder for MockPolicyWatcher
type MockPolicyWatcherMockRecorder struct {
	mock *MockPolicyWatcher
}

// NewMockPolicyWatcher creates a new mock instance
func NewMockPolicyWatcher(ctrl *gomock.Controller) *MockPolicyWatcher {
	mock := &MockPolicyWatcher{ctrl: ctrl}
	mock.recorder = &MockPolicyWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPolicyWatcher) EXPECT() *MockPolicyWatcherMockRecorder {
	return m.recorder
}

// Watch mocks base method
func (m *MockPolicyWatcher) Watch(namespace string, opts clients.WatchOpts) (<-chan v1alpha1.PolicyList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1alpha1.PolicyList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockPolicyWatcherMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockPolicyWatcher)(nil).Watch), namespace, opts)
}

// MockPolicyClient is a mock of PolicyClient interface
type MockPolicyClient struct {
	ctrl     *gomock.Controller
	recorder *MockPolicyClientMockRecorder
}

// MockPolicyClientMockRecorder is the mock recorder for MockPolicyClient
type MockPolicyClientMockRecorder struct {
	mock *MockPolicyClient
}

// NewMockPolicyClient creates a new mock instance
func NewMockPolicyClient(ctrl *gomock.Controller) *MockPolicyClient {
	mock := &MockPolicyClient{ctrl: ctrl}
	mock.recorder = &MockPolicyClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPolicyClient) EXPECT() *MockPolicyClientMockRecorder {
	return m.recorder
}

// BaseClient mocks base method
func (m *MockPolicyClient) BaseClient() clients.ResourceClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseClient")
	ret0, _ := ret[0].(clients.ResourceClient)
	return ret0
}

// BaseClient indicates an expected call of BaseClient
func (mr *MockPolicyClientMockRecorder) BaseClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseClient", reflect.TypeOf((*MockPolicyClient)(nil).BaseClient))
}

// Register mocks base method
func (m *MockPolicyClient) Register() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register")
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register
func (mr *MockPolicyClientMockRecorder) Register() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPolicyClient)(nil).Register))
}

// Read mocks base method
func (m *MockPolicyClient) Read(namespace, name string, opts clients.ReadOpts) (*v1alpha1.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", namespace, name, opts)
	ret0, _ := ret[0].(*v1alpha1.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockPolicyClientMockRecorder) Read(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockPolicyClient)(nil).Read), namespace, name, opts)
}

// Write mocks base method
func (m *MockPolicyClient) Write(resource *v1alpha1.Policy, opts clients.WriteOpts) (*v1alpha1.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", resource, opts)
	ret0, _ := ret[0].(*v1alpha1.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockPolicyClientMockRecorder) Write(resource, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockPolicyClient)(nil).Write), resource, opts)
}

// Delete mocks base method
func (m *MockPolicyClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", namespace, name, opts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPolicyClientMockRecorder) Delete(namespace, name, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPolicyClient)(nil).Delete), namespace, name, opts)
}

// List mocks base method
func (m *MockPolicyClient) List(namespace string, opts clients.ListOpts) (v1alpha1.PolicyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", namespace, opts)
	ret0, _ := ret[0].(v1alpha1.PolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockPolicyClientMockRecorder) List(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPolicyClient)(nil).List), namespace, opts)
}

// Watch mocks base method
func (m *MockPolicyClient) Watch(namespace string, opts clients.WatchOpts) (<-chan v1alpha1.PolicyList, <-chan error, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", namespace, opts)
	ret0, _ := ret[0].(<-chan v1alpha1.PolicyList)
	ret1, _ := ret[1].(<-chan error)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Watch indicates an expected call of Watch
func (mr *MockPolicyClientMockRecorder) Watch(namespace, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockPolicyClient)(nil).Watch), namespace, opts)
}
