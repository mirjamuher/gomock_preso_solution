package order

import (
	"errors"
	"fmt"
	p "github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
)

func (ps *BookingService) CreateBooking(booking *Booking) error {
	// Validate the booking
	if err := booking.Validate(); err != nil {
		return err
	}

	// Process p for the booking
	product := booking.Product
	totalPrice := product.Price * float64(booking.Quantity)

	payment := &p.Payment{
		TotalPrice: totalPrice,
		Method:     booking.PaymentMethod,
	}
	state, err := ps.paymentService.ProcessPayment(payment)
	if err != nil {
		return err
	}
	if state != p.Succeeded {
		return errors.New(fmt.Sprintf("PaymentState is %v", state))
	}

	// Create the booking in the database
	if err := ps.InsertOrder(booking, state); err != nil {
		return err
	}

	return nil
}
