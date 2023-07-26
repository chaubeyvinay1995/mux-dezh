// Code generated by MockGen. DO NOT EDIT.
// Source: http-util.go

// Package mock_impl is a generated GoMock package.
package mock_impl

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockHTTPClient is a mock of HTTPClient interface
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Post mocks base method
func (m *MockHTTPClient) Post(ctx context.Context, url string, requestData, responseData interface{}, httpHeaders ...map[string]string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, url, requestData, responseData}
	for _, a := range httpHeaders {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Post", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post
func (mr *MockHTTPClientMockRecorder) Post(ctx, url, requestData, responseData interface{}, httpHeaders ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, url, requestData, responseData}, httpHeaders...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockHTTPClient)(nil).Post), varargs...)
}

// Get mocks base method
func (m *MockHTTPClient) Get(ctx context.Context, url string, responseData interface{}, httpHeaders ...map[string]string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, url, responseData}
	for _, a := range httpHeaders {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockHTTPClientMockRecorder) Get(ctx, url, responseData interface{}, httpHeaders ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, url, responseData}, httpHeaders...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHTTPClient)(nil).Get), varargs...)
}
