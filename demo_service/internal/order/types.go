package order

import (
	"database/sql"
	"github.com/mirjamuher/gomock_preso_solution/demo_service/internal/payment"
)

type ProductService struct {
	db             *sql.DB
	paymentService payment.PaymentService
}

type Order struct {
	Product Product
	Quantity int
	PaymentMethod string
}

func (o *Order) Validate() error {
	return nil
}

type Product struct{
	Price float64
}