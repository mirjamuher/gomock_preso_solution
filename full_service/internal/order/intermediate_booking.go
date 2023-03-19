package order

import (
	p "github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
)

func (ps *ProductService) CreateOrders(order *Order) error {
	// Validate the order
	if err := order.Validate(); err != nil {
		return err
	}
	if order.Quantity == 0 {
		return nil
	}

	// keep track of the orders we've processed payments for
	var paidOrders []*Order

	for i := 0; i < order.Quantity; i ++ {
		// Process payment for the order
		product := order.Product
		totalPrice := product.Price * float64(order.Quantity)

		payment := &p.Payment{
			TotalPrice: totalPrice,
			Method:     order.PaymentMethod,
		}
		state, err := ps.paymentService.ProcessPayment(payment)
		if err != nil ||  state != p.Succeeded  {
			order.PaymentState = p.Failed
		}

		// store successfully payed order
		order.PaymentState = state
		paidOrders = append(paidOrders, order)
	}

	// Create the order in the database
	if err := ps.InsertOrders(paidOrders); err != nil {
		return err
	}

	return nil
}