package booking

import (
	"errors"
	"fmt"
	"github.com/mirjamuher/gomock_preso_solution/demo_service/internal/payment"
)

type BookingService struct {
	paymentService payment.Payer
}

func (ps *BookingService) CreateBooking(booking *Booking) error {
	// Validate the booking
	if err := booking.Validate(); err != nil {
		return err
	}

	// Process payment for the booking
	p := &payment.Payload{
		TotalPrice: booking.product.price * float64(booking.quantity),
		Method:     booking.paymentMethod,
	}
	state, err := ps.paymentService.ProcessPayment(p)
	if err != nil {
		return err
	}
	if state != payment.Succeeded {
		return errors.New(fmt.Sprintf("State is %v", state))
	}

	return nil
}
