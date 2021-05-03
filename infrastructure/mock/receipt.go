// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/receipt.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/softree-group/kitchen-plan-backend/domain/entity"
)

// MockReceiptReceiver is a mock of ReceiptReceiver interface.
type MockReceiptReceiver struct {
	ctrl     *gomock.Controller
	recorder *MockReceiptReceiverMockRecorder
}

// MockReceiptReceiverMockRecorder is the mock recorder for MockReceiptReceiver.
type MockReceiptReceiverMockRecorder struct {
	mock *MockReceiptReceiver
}

// NewMockReceiptReceiver creates a new mock instance.
func NewMockReceiptReceiver(ctrl *gomock.Controller) *MockReceiptReceiver {
	mock := &MockReceiptReceiver{ctrl: ctrl}
	mock.recorder = &MockReceiptReceiverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceiptReceiver) EXPECT() *MockReceiptReceiverMockRecorder {
	return m.recorder
}

// GetReceipt mocks base method.
func (m *MockReceiptReceiver) GetReceipt(id int) (*entity.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipt", id)
	ret0, _ := ret[0].(*entity.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipt indicates an expected call of GetReceipt.
func (mr *MockReceiptReceiverMockRecorder) GetReceipt(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipt", reflect.TypeOf((*MockReceiptReceiver)(nil).GetReceipt), id)
}

// GetReceipts mocks base method.
func (m *MockReceiptReceiver) GetReceipts(selection entity.Selection) ([]entity.Receipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReceipts", selection)
	ret0, _ := ret[0].([]entity.Receipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReceipts indicates an expected call of GetReceipts.
func (mr *MockReceiptReceiverMockRecorder) GetReceipts(selection interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReceipts", reflect.TypeOf((*MockReceiptReceiver)(nil).GetReceipts), selection)
}
