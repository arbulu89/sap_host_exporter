// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/sapcontrol/webservice.go

// Package mock_sapcontrol is a generated GoMock package.
package mock_sapcontrol

import (
	sapcontrol "github.com/SUSE/sap_host_exporter/internal/sapcontrol"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWebService is a mock of WebService interface
type MockWebService struct {
	ctrl     *gomock.Controller
	recorder *MockWebServiceMockRecorder
}

// MockWebServiceMockRecorder is the mock recorder for MockWebService
type MockWebServiceMockRecorder struct {
	mock *MockWebService
}

// NewMockWebService creates a new mock instance
func NewMockWebService(ctrl *gomock.Controller) *MockWebService {
	mock := &MockWebService{ctrl: ctrl}
	mock.recorder = &MockWebServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWebService) EXPECT() *MockWebServiceMockRecorder {
	return m.recorder
}

// GetProcessList mocks base method
func (m *MockWebService) GetProcessList() (*sapcontrol.GetProcessListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcessList")
	ret0, _ := ret[0].(*sapcontrol.GetProcessListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcessList indicates an expected call of GetProcessList
func (mr *MockWebServiceMockRecorder) GetProcessList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessList", reflect.TypeOf((*MockWebService)(nil).GetProcessList))
}
