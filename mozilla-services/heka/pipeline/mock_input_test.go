// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/pipeline (interfaces: Input)

package pipeline

import (
	gomock "github.com/rafrombrc/gomock/gomock"
)

// Mock of Input interface
type MockInput struct {
	ctrl     *gomock.Controller
	recorder *_MockInputRecorder
}

// Recorder for MockInput (not exported)
type _MockInputRecorder struct {
	mock *MockInput
}

func NewMockInput(ctrl *gomock.Controller) *MockInput {
	mock := &MockInput{ctrl: ctrl}
	mock.recorder = &_MockInputRecorder{mock}
	return mock
}

func (_m *MockInput) EXPECT() *_MockInputRecorder {
	return _m.recorder
}

func (_m *MockInput) Run(_param0 InputRunner, _param1 PluginHelper) error {
	ret := _m.ctrl.Call(_m, "Run", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockInputRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0, arg1)
}

func (_m *MockInput) Stop() {
	_m.ctrl.Call(_m, "Stop")
}

func (_mr *_MockInputRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}
