package order

import (
	"context"
	"errors"
	"fmt"
	p "github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
	"time"
)

// The below is nonsensical, to be able to show more complex testing
func (ps *BookingService) ProcessBooking(booking *Booking) error {
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
	// if it went through, we end the flow
	if state == p.Succeeded {
		return nil
	}

	// If the state indicates it's pending, we do another payment for good measure
	if state == p.Initiated {
		_, err := ps.PaymentService.ProcessPayment(payment)
		if err != nil {
			return err
		}
	}

	// If the state indicates failure, we refund the payment
	if state == p.Failed {
		ch := make(chan error)
		ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		// call refund payment in a go function
		go func(ctx context.Context, ch chan error) {
			err = ps.PaymentService.RefundPayment(payment)
			if err != nil {
				ch <- err
				return
			}
			ch <- nil
		}(ctxTimeout, ch)

		// RefundPayment has 5 seconds to succeed, otherwise we end it
		select {
		case <-ctxTimeout.Done():
			// if we couldn't refund, we give up
			return errors.New("timeout")
		case err2 := <-ch:
			if err2 != nil {
				return nil
			}
			// Refund successful, so we try again
			_, err = ps.PaymentService.ProcessPayment(payment)
			if err != nil {
				return err
			}
		}
		return nil
	}

	// if the state is unknown, we pass a reason to be filled out
	if state == p.Unknown {
		reason := &p.Reason{}
		err := ps.PaymentService.UnmarshalReason(reason)
		if err != nil {
			reason2 := ps.PaymentService.UnmarshalAndReturnReason(reason)
			fmt.Println(reason2.Msg)
			return nil
		}
		fmt.Println(reason.Msg)
	}

	return nil
}
