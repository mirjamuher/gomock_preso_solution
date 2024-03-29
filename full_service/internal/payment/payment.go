package payment

import (
	"net/http"
)

//go:generate mockery --with-expecter=true --name Payer
//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --with-expecter=true --name Payer

type Payer interface {
	ProcessPayment(p *Payment) (State, error)
	RefundPayment(p *Payment) error
	UpdateReason(r *Reason) error
	UpdateAndReturnReason(r *Reason) *Reason
}

type PaymentService struct {
	client *http.Client
}

func (ps *PaymentService) ProcessPayment(p *Payment) (State, error) {
	// process payment logic here
	return Succeeded, nil
}

func (ps *PaymentService) RefundPayment(p *Payment) error {
	// refund payment logic here
	return nil
}

func (ps *PaymentService) UpdateReason(r *Reason) error {
	// update passed in pointer
	r.Msg = "it's not you, it's me"
	return nil
}

func (ps *PaymentService) UpdateAndReturnReason(r *Reason) *Reason {
	// update passed in pointer
	r.Msg = "it's not you, it's me"
	return r
}
