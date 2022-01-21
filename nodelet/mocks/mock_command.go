// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/platform9/nodelet/pkg/utils/command (interfaces: CLI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCLI is a mock of CLI interface
type MockCLI struct {
	ctrl     *gomock.Controller
	recorder *MockCLIMockRecorder
}

// MockCLIMockRecorder is the mock recorder for MockCLI
type MockCLIMockRecorder struct {
	mock *MockCLI
}

// NewMockCLI creates a new mock instance
func NewMockCLI(ctrl *gomock.Controller) *MockCLI {
	mock := &MockCLI{ctrl: ctrl}
	mock.recorder = &MockCLIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCLI) EXPECT() *MockCLIMockRecorder {
	return m.recorder
}

// RunCommand mocks base method
func (m *MockCLI) RunCommand(arg0 context.Context, arg1 map[string]string, arg2 int, arg3, arg4 string, arg5 ...string) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCommand", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunCommand indicates an expected call of RunCommand
func (mr *MockCLIMockRecorder) RunCommand(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommand", reflect.TypeOf((*MockCLI)(nil).RunCommand), varargs...)
}

// RunCommandWithStdErr mocks base method
func (m *MockCLI) RunCommandWithStdErr(arg0 context.Context, arg1 map[string]string, arg2 int, arg3, arg4 string, arg5 ...string) (int, []string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCommandWithStdErr", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RunCommandWithStdErr indicates an expected call of RunCommandWithStdErr
func (mr *MockCLIMockRecorder) RunCommandWithStdErr(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommandWithStdErr", reflect.TypeOf((*MockCLI)(nil).RunCommandWithStdErr), varargs...)
}

// RunCommandWithStdOut mocks base method
func (m *MockCLI) RunCommandWithStdOut(arg0 context.Context, arg1 map[string]string, arg2 int, arg3, arg4 string, arg5 ...string) (int, []string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCommandWithStdOut", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RunCommandWithStdOut indicates an expected call of RunCommandWithStdOut
func (mr *MockCLIMockRecorder) RunCommandWithStdOut(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommandWithStdOut", reflect.TypeOf((*MockCLI)(nil).RunCommandWithStdOut), varargs...)
}

// RunCommandWithStdOutStdErr mocks base method
func (m *MockCLI) RunCommandWithStdOutStdErr(arg0 context.Context, arg1 map[string]string, arg2 int, arg3, arg4 string, arg5 ...string) (int, []string, []string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RunCommandWithStdOutStdErr", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].([]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// RunCommandWithStdOutStdErr indicates an expected call of RunCommandWithStdOutStdErr
func (mr *MockCLIMockRecorder) RunCommandWithStdOutStdErr(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommandWithStdOutStdErr", reflect.TypeOf((*MockCLI)(nil).RunCommandWithStdOutStdErr), varargs...)
}