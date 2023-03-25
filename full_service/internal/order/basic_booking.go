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

	// Process payment for the booking
	payment := &p.Payment{
		TotalPrice: booking.Product.Price * float64(booking.Quantity),
		Method:     booking.PaymentMethod,
	}
	state, err := ps.PaymentService.ProcessPayment(payment)
	if err != nil {
		return err
	}
	if state != p.Succeeded {
		return errors.New(fmt.Sprintf("State is %v", state))
	}

	return nil
}
