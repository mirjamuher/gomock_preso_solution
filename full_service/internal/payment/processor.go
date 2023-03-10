package payment

import (
	"net/http"
)

// Important - explain in presentation the bug that makes --build_flags=--mod=mod required
//go:generate mockgen -destination=mocks/mock_payer.go --build_flags=--mod=mod github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment Payer

type Payer interface {
	ProcessPayment(p *Payment) (PaymentState, error)
}

type PaymentService struct {
	client *http.Client
}

func (ps *PaymentService) ProcessPayment(p *Payment) (PaymentState, error) {
	// Insert payment into DB
	if err := ps.PersistPayment(p); err != nil {
		return Failed, err
	}

	// Call external payment API
	req, err := ps.CreateRequest(p)
	if err != nil {
		return Failed, err
	}

	state, err := ps.SendPaymentRequest(req);
	if err != nil {
		return Failed, err
	}

	return state, nil
}

