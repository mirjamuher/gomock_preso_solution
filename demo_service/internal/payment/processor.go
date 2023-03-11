package payment

import "net/http"

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