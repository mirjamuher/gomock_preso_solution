package order

import (
	"github.com/mirjamuher/gomock_preso_solution/demo_service/internal/payment"
)

type BookingService struct {
	PaymentService payment.PaymentService
}

type Booking struct {
	Product Product
	Quantity int
	PaymentMethod string
	PaymentState payment.State
}

func (o *Booking) Validate() error {
	return nil
}

type Product struct{
	Price float64
}