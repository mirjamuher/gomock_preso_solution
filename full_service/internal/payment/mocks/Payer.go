// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	payment "github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
	mock "github.com/stretchr/testify/mock"
)

// Payer is an autogenerated mock type for the Payer type
type Payer struct {
	mock.Mock
}

type Payer_Expecter struct {
	mock *mock.Mock
}

func (_m *Payer) EXPECT() *Payer_Expecter {
	return &Payer_Expecter{mock: &_m.Mock}
}

// ProcessPayment provides a mock function with given fields: p
func (_m *Payer) ProcessPayment(p *payment.Payment) (payment.PaymentState, error) {
	ret := _m.Called(p)

	var r0 payment.PaymentState
	var r1 error
	if rf, ok := ret.Get(0).(func(*payment.Payment) (payment.PaymentState, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func(*payment.Payment) payment.PaymentState); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(payment.PaymentState)
	}

	if rf, ok := ret.Get(1).(func(*payment.Payment) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Payer_ProcessPayment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessPayment'
type Payer_ProcessPayment_Call struct {
	*mock.Call
}

// ProcessPayment is a helper method to define mock.On call
//  - p *payment.Payment
func (_e *Payer_Expecter) ProcessPayment(p interface{}) *Payer_ProcessPayment_Call {
	return &Payer_ProcessPayment_Call{Call: _e.mock.On("ProcessPayment", p)}
}

func (_c *Payer_ProcessPayment_Call) Run(run func(p *payment.Payment)) *Payer_ProcessPayment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*payment.Payment))
	})
	return _c
}

func (_c *Payer_ProcessPayment_Call) Return(_a0 payment.PaymentState, _a1 error) *Payer_ProcessPayment_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Payer_ProcessPayment_Call) RunAndReturn(run func(*payment.Payment) (payment.PaymentState, error)) *Payer_ProcessPayment_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPayer interface {
	mock.TestingT
	Cleanup(func())
}

// NewPayer creates a new instance of Payer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPayer(t mockConstructorTestingTNewPayer) *Payer {
	mock := &Payer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
