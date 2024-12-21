// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gionnid/terraform-provider-notion/internal/provider/client (interfaces: NotionAPI)
//
// Generated by this command:
//
//	mockgen -destination=mocks/mock_notion_api.go -package=mocks github.com/gionnid/terraform-provider-notion/internal/provider/client NotionAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockNotionAPI is a mock of NotionAPI interface.
type MockNotionAPI struct {
	ctrl     *gomock.Controller
	recorder *MockNotionAPIMockRecorder
	isgomock struct{}
}

// MockNotionAPIMockRecorder is the mock recorder for MockNotionAPI.
type MockNotionAPIMockRecorder struct {
	mock *MockNotionAPI
}

// NewMockNotionAPI creates a new mock instance.
func NewMockNotionAPI(ctrl *gomock.Controller) *MockNotionAPI {
	mock := &MockNotionAPI{ctrl: ctrl}
	mock.recorder = &MockNotionAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotionAPI) EXPECT() *MockNotionAPIMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockNotionAPI) Get(url string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", url)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockNotionAPIMockRecorder) Get(url any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNotionAPI)(nil).Get), url)
}

// Init mocks base method.
func (m *MockNotionAPI) Init(token, version string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init", token, version)
}

// Init indicates an expected call of Init.
func (mr *MockNotionAPIMockRecorder) Init(token, version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockNotionAPI)(nil).Init), token, version)
}

// Patch mocks base method.
func (m *MockNotionAPI) Patch(url, body string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Patch", url, body)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockNotionAPIMockRecorder) Patch(url, body any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockNotionAPI)(nil).Patch), url, body)
}

// Post mocks base method.
func (m *MockNotionAPI) Post(url, body string) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", url, body)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Post indicates an expected call of Post.
func (mr *MockNotionAPIMockRecorder) Post(url, body any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockNotionAPI)(nil).Post), url, body)
}
