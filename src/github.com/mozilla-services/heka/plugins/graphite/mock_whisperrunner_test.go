// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/plugins/graphite (interfaces: WhisperRunner)

package graphite

import (
	whisper "github.com/rafrombrc/whisper-go/whisper"
	gomock "github.com/rafrombrc/gomock/gomock"
)

// Mock of WhisperRunner interface
type MockWhisperRunner struct {
	ctrl     *gomock.Controller
	recorder *_MockWhisperRunnerRecorder
}

// Recorder for MockWhisperRunner (not exported)
type _MockWhisperRunnerRecorder struct {
	mock *MockWhisperRunner
}

func NewMockWhisperRunner(ctrl *gomock.Controller) *MockWhisperRunner {
	mock := &MockWhisperRunner{ctrl: ctrl}
	mock.recorder = &_MockWhisperRunnerRecorder{mock}
	return mock
}

func (_m *MockWhisperRunner) EXPECT() *_MockWhisperRunnerRecorder {
	return _m.recorder
}

func (_m *MockWhisperRunner) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockWhisperRunnerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockWhisperRunner) InChan() chan *whisper.Point {
	ret := _m.ctrl.Call(_m, "InChan")
	ret0, _ := ret[0].(chan *whisper.Point)
	return ret0
}

func (_mr *_MockWhisperRunnerRecorder) InChan() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InChan")
}
