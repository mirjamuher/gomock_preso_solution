package payment

import (
	"net/http"
)

// Important - explain in presentation the bug that makes --build_flags=--mod=mod required
//go:generate go run github.com/stretchr/testify@v1.7.0 -name Payer

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

