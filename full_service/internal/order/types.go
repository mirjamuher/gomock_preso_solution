package order

import (
	"database/sql"
	"github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
)

type BookingService struct {
	db             *sql.DB
	paymentService payment.Payer
}

type Booking struct {
	Product Product
	Quantity int
	PaymentMethod string
	PaymentState payment.PaymentState
}

func (o *Booking) Validate() error {
	return nil
}

type Product struct{
	Price float64
}