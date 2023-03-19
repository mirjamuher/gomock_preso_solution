package order

import (
	"errors"
	"fmt"
	p "github.com/mirjamuher/gomock_preso_solution/demo_service/internal/payment"
)

func (ps *ProductService) CreateOrder(order *Order) error {
	// Validate the order
	if err := order.Validate(); err != nil {
		return err
	}

	// Process payment for the order
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
	if state != p.Success {
		return errors.New(fmt.Sprintf("PaymentState is %v", state))
	}

	// TODO: Take the DB stuff out
	// Create the order in the database
	if err := ps.InsertOrder(order, state); err != nil {
		return err
	}

	return nil
}