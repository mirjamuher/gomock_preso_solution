package payment

import (
	"net/http"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --with-expecter=true --name Payer

type PaymentService struct {
	client *http.Client
}

func (ps *PaymentService) ProcessPayment(p *Payment) (State, error) {
	// process payment logic here
	return Succeeded, nil
}
