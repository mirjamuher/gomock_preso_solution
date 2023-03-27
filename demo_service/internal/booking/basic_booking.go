package booking

import (
	"errors"
	"fmt"
	p "github.com/mirjamuher/gomock_preso_solution/demo_service/internal/payment"
)

func (ps *BookingService) CreateBooking(booking *Booking) error {
	// Validate the booking
	if err := booking.Validate(); err != nil {
		return err
	}

	// Process payment for the booking
	payment := &p.Payment{
		TotalPrice: booking.product.price * float64(booking.quantity),
		Method:     booking.paymentMethod,
	}
	state, err := ps.paymentService.ProcessPayment(payment)
	if err != nil {
		return err
	}
	if state != p.Succeeded {
		return errors.New(fmt.Sprintf("State is %v", state))
	}

	return nil
}
