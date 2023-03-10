package payment

import "net/http"

func (ps *PaymentService) PersistPayment(p *Payment) error {
	// persisting logic
	return nil
}

func (ps *PaymentService) CreateRequest(p *Payment) (*http.Request, error) {
	// create http request
	return &http.Request{}, nil
}

func (ps *PaymentService) SendPaymentRequest(req *http.Request) (PaymentState, error) {
	// send payment request
	return Succeeded, nil
}
