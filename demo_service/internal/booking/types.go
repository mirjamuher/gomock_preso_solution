package booking

import (
	"github.com/mirjamuher/gomock_preso_solution/demo_service/internal/payment"
)

type Booking struct {
	product       Product
	quantity      int
	paymentMethod string
	paymentState  payment.State
}

func (o *Booking) Validate() error {
	return nil
}

type Product struct{
	price float64
}