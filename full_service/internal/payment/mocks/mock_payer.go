// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment (interfaces: Payer)

// Package mock_payment is a generated GoMock package.
package mock_payment

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	payment "github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
)

// MockPayer is a mock of Payer interface.
type MockPayer struct {
	ctrl     *gomock.Controller
	recorder *MockPayerMockRecorder
}

// MockPayerMockRecorder is the mock recorder for MockPayer.
type MockPayerMockRecorder struct {
	mock *MockPayer
}

// NewMockPayer creates a new mock instance.
func NewMockPayer(ctrl *gomock.Controller) *MockPayer {
	mock := &MockPayer{ctrl: ctrl}
	mock.recorder = &MockPayerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPayer) EXPECT() *MockPayerMockRecorder {
	return m.recorder
}

// ProcessPayment mocks base method.
func (m *MockPayer) ProcessPayment(arg0 *payment.Payment) (payment.PaymentState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessPayment", arg0)
	ret0, _ := ret[0].(payment.PaymentState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessPayment indicates an expected call of ProcessPayment.
func (mr *MockPayerMockRecorder) ProcessPayment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessPayment", reflect.TypeOf((*MockPayer)(nil).ProcessPayment), arg0)
}
