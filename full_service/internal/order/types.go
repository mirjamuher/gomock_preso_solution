package product

import (
	"database/sql"
	"github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
)

type ProductService struct {
	db             *sql.DB
	paymentService payment.Payer
}

type Order struct {
	Product Product
	Quantity int
	PaymentMethod string
	PaymentState payment.PaymentState
}

func (o *Order) Validate() error {
	return nil
}

type Product struct{
	Price float64
}