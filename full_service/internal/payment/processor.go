package payment

import (
	"net/http"
)

// Important - explain in presentation the bug that makes --build_flags=--mod=mod required
//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -destination=mocks/mock_payer.go --build_flags=--mod=mod github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment Payer

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

