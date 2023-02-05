// Code generated by MockGen. DO NOT EDIT.
// Source: crypto_stream.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/zzzz-26/quic-go/internal/protocol"
	wire "github.com/zzzz-26/quic-go/internal/wire"
)

// MockCryptoStream is a mock of CryptoStream interface.
type MockCryptoStream struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoStreamMockRecorder
}

// MockCryptoStreamMockRecorder is the mock recorder for MockCryptoStream.
type MockCryptoStreamMockRecorder struct {
	mock *MockCryptoStream
}

// NewMockCryptoStream creates a new mock instance.
func NewMockCryptoStream(ctrl *gomock.Controller) *MockCryptoStream {
	mock := &MockCryptoStream{ctrl: ctrl}
	mock.recorder = &MockCryptoStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoStream) EXPECT() *MockCryptoStreamMockRecorder {
	return m.recorder
}

// Finish mocks base method.
func (m *MockCryptoStream) Finish() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finish")
	ret0, _ := ret[0].(error)
	return ret0
}

// Finish indicates an expected call of Finish.
func (mr *MockCryptoStreamMockRecorder) Finish() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finish", reflect.TypeOf((*MockCryptoStream)(nil).Finish))
}

// GetCryptoData mocks base method.
func (m *MockCryptoStream) GetCryptoData() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCryptoData")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetCryptoData indicates an expected call of GetCryptoData.
func (mr *MockCryptoStreamMockRecorder) GetCryptoData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCryptoData", reflect.TypeOf((*MockCryptoStream)(nil).GetCryptoData))
}

// HandleCryptoFrame mocks base method.
func (m *MockCryptoStream) HandleCryptoFrame(arg0 *wire.CryptoFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleCryptoFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleCryptoFrame indicates an expected call of HandleCryptoFrame.
func (mr *MockCryptoStreamMockRecorder) HandleCryptoFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleCryptoFrame", reflect.TypeOf((*MockCryptoStream)(nil).HandleCryptoFrame), arg0)
}

// HasData mocks base method.
func (m *MockCryptoStream) HasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasData indicates an expected call of HasData.
func (mr *MockCryptoStreamMockRecorder) HasData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasData", reflect.TypeOf((*MockCryptoStream)(nil).HasData))
}

// PopCryptoFrame mocks base method.
func (m *MockCryptoStream) PopCryptoFrame(arg0 protocol.ByteCount) *wire.CryptoFrame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopCryptoFrame", arg0)
	ret0, _ := ret[0].(*wire.CryptoFrame)
	return ret0
}

// PopCryptoFrame indicates an expected call of PopCryptoFrame.
func (mr *MockCryptoStreamMockRecorder) PopCryptoFrame(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopCryptoFrame", reflect.TypeOf((*MockCryptoStream)(nil).PopCryptoFrame), arg0)
}

// Write mocks base method.
func (m *MockCryptoStream) Write(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockCryptoStreamMockRecorder) Write(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCryptoStream)(nil).Write), p)
}
