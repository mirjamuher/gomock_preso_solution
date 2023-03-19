package order

import p "github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"

func (ps *BookingService) InsertOrder(order *Booking, state p.PaymentState) error {
	return nil
}

func (ps *BookingService) InsertOrders([]*Booking) error {
	return nil
}