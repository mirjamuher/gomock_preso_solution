package order

import (
	"github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
)

type BookingService struct {
	PaymentService payment.Payer
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