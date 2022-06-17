// Code generated by MockGen. DO NOT EDIT.
// Source: ./api.go

// Package mockapi_test is a generated GoMock package.
package mockapi_test

import (
	reflect "reflect"
	client "teststaticwongnai2/client"

	gomock "github.com/golang/mock/gomock"
)

// MockHandlerMock is a mock of HandlerMock interface.
type MockHandlerMock struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockMockRecorder
}

// MockHandlerMockMockRecorder is the mock recorder for MockHandlerMock.
type MockHandlerMockMockRecorder struct {
	mock *MockHandlerMock
}

// NewMockHandlerMock creates a new mock instance.
func NewMockHandlerMock(ctrl *gomock.Controller) *MockHandlerMock {
	mock := &MockHandlerMock{ctrl: ctrl}
	mock.recorder = &MockHandlerMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandlerMock) EXPECT() *MockHandlerMockMockRecorder {
	return m.recorder
}

// GetHttp mocks base method.
func (m *MockHandlerMock) GetHttp() (map[string]int, client.Age) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHttp")
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(client.Age)
	return ret0, ret1
}

// GetHttp indicates an expected call of GetHttp.
func (mr *MockHandlerMockMockRecorder) GetHttp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHttp", reflect.TypeOf((*MockHandlerMock)(nil).GetHttp))
}
