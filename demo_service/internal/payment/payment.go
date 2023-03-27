package payment

import (
	"net/http"
)

type PaymentService struct {
	client *http.Client
}

func (ps *PaymentService) ProcessPayment(p *Payload) (State, error) {
	// process payment logic here
	return Succeeded, nil
}
