// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/abhinav/git-fu/git (interfaces: RebaseHandle)

package gittest

import (
	git "github.com/abhinav/git-fu/git"
	gomock "github.com/golang/mock/gomock"
)

// Mock of RebaseHandle interface
type MockRebaseHandle struct {
	ctrl     *gomock.Controller
	recorder *_MockRebaseHandleRecorder
}

// Recorder for MockRebaseHandle (not exported)
type _MockRebaseHandleRecorder struct {
	mock *MockRebaseHandle
}

func NewMockRebaseHandle(ctrl *gomock.Controller) *MockRebaseHandle {
	mock := &MockRebaseHandle{ctrl: ctrl}
	mock.recorder = &_MockRebaseHandleRecorder{mock}
	return mock
}

func (_m *MockRebaseHandle) EXPECT() *_MockRebaseHandleRecorder {
	return _m.recorder
}

func (_m *MockRebaseHandle) Base() string {
	ret := _m.ctrl.Call(_m, "Base")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockRebaseHandleRecorder) Base() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Base")
}

func (_m *MockRebaseHandle) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRebaseHandleRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Err")
}

func (_m *MockRebaseHandle) Rebase(_param0 string, _param1 string) git.RebaseHandle {
	ret := _m.ctrl.Call(_m, "Rebase", _param0, _param1)
	ret0, _ := ret[0].(git.RebaseHandle)
	return ret0
}

func (_mr *_MockRebaseHandleRecorder) Rebase(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rebase", arg0, arg1)
}
