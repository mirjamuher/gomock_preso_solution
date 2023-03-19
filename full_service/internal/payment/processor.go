package payment

import (
	"net/http"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name Payer

type Payer interface {
	ProcessPayment(p *Payment) (PaymentState, error)
}

type PaymentService struct {
	client *http.Client
}

func (ps *PaymentService) ProcessPayment(p *Payment) (PaymentState, error) {
	// process payment
	return Succeeded, nil
}

