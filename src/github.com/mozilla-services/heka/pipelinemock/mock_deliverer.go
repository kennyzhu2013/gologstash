// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/pipeline (interfaces: Deliverer)

package pipelinemock

import (
	pipeline "github.com/mozilla-services/heka/pipeline"
	gomock "github.com/rafrombrc/gomock/gomock"
)

// Mock of Deliverer interface
type MockDeliverer struct {
	ctrl     *gomock.Controller
	recorder *_MockDelivererRecorder
}

// Recorder for MockDeliverer (not exported)
type _MockDelivererRecorder struct {
	mock *MockDeliverer
}

func NewMockDeliverer(ctrl *gomock.Controller) *MockDeliverer {
	mock := &MockDeliverer{ctrl: ctrl}
	mock.recorder = &_MockDelivererRecorder{mock}
	return mock
}

func (_m *MockDeliverer) EXPECT() *_MockDelivererRecorder {
	return _m.recorder
}

func (_m *MockDeliverer) Deliver(_param0 *pipeline.PipelinePack) {
	_m.ctrl.Call(_m, "Deliver", _param0)
}

func (_mr *_MockDelivererRecorder) Deliver(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Deliver", arg0)
}

func (_m *MockDeliverer) DeliverFunc() pipeline.DeliverFunc {
	ret := _m.ctrl.Call(_m, "DeliverFunc")
	ret0, _ := ret[0].(pipeline.DeliverFunc)
	return ret0
}

func (_mr *_MockDelivererRecorder) DeliverFunc() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeliverFunc")
}

func (_m *MockDeliverer) Done() {
	_m.ctrl.Call(_m, "Done")
}

func (_mr *_MockDelivererRecorder) Done() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Done")
}