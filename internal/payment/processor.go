package payment

import "net/http"

//go:generate mockgen -destination=mocks/mock_payer.go

type Payer interface {
	ProcessPayment(p *Payment) (*PaymentState, error)
}

type PaymentService struct {
	client *http.Client
}

func (ps *PaymentService) ProcessPayment(p *Payment) (*PaymentState, error) {
	// Insert payment into DB
	if err := ps.PersistPayment(p); err != nil {
		return nil, err
	}

	// Call external payment API
	req, err := ps.CreateRequest(p)
	if err != nil {
		return nil, err
	}

	state, err := ps.SendPaymentRequest(req);
	if err != nil {
		return nil, err
	}

	return &state, nil
}

