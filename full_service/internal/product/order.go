package product

import (
	"database/sql"
	p "github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
)

type ProductService struct {
	db             *sql.DB
	paymentService p.PaymentService
}

func (ps *ProductService) CreateOrder(order *Order) error {
	// Validate the order
	if err := order.Validate(); err != nil {
		return err
	}

	// Process p for the order
	product := order.Product
	totalPrice := product.Price * float64(order.Quantity)

	payment := &p.Payment{
		TotalPrice: totalPrice,
		Method: order.PaymentMethod,
	}
	state, err := ps.paymentService.ProcessPayment(payment)
	if err != nil {
		return err
	}

	// Create the order in the database
	if err := ps.InsertOrder(order, *state); err != nil {
		return err
	}

	return nil
}
