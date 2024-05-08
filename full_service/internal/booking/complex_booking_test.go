package booking

import (
	"errors"
	"github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
	"github.com/mirjamuher/gomock_preso_solution/full_service/mocks"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestBookingService_ProcessBooking(t *testing.T) {
	validOrder := Booking{
		Product: Product{
			Price: 100,
		},
		Quantity:      1,
		PaymentMethod: "CREDIT",
		PaymentState:  0,
	}
	validPayment := payment.Payment{
		TotalPrice: 100,
		Method:     "CREDIT",
	}

	type fields struct {
		PaymentService func(t *testing.T) payment.Payer
	}
	type args struct {
		booking *Booking
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success: payment initiated, so we do another one",
			fields: fields{
				PaymentService: func(t *testing.T) payment.Payer {
					ps := mocks_payment.NewPayer(t)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Initiated, nil).Times(2)
					return ps
				},
			},
			args: args{
				booking: &validOrder,
			},
			wantErr: false,
		},
		{
			// comment: what's interesting here is that the order doesn't matter if you don't specify
			name: "success: payment failed, so we do a call to Refund & then another payment",
			fields: fields{
				PaymentService: func(t *testing.T) payment.Payer {
					ps := mocks_payment.NewPayer(t)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Failed, nil).Twice()
					ps.EXPECT().RefundPayment(mock.Anything).Return(nil).Once()
					return ps
				},
			},
			args: args{
				booking: &validOrder,
			},
			wantErr: false,
		},
		{
			// comment: same test, this time we assert that it has to be in a certain order
			name: "success: payment failed, so we do a call to Refund & then another payment",
			fields: fields{
				PaymentService: func(t *testing.T) payment.Payer {
					ps := mocks_payment.NewPayer(t)
					call1 := ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Failed, nil).Once()
					call2 := ps.EXPECT().RefundPayment(mock.Anything).Return(nil).Once().NotBefore(call1)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Failed, nil).Once().NotBefore(call2)
					return ps
				},
			},
			args: args{
				booking: &validOrder,
			},
			wantErr: false,
		},
		{
			name: "error: (1) payment failed so we (2) try to refund but it (3) times out so we return an error",
			fields: fields{
				PaymentService: func(t *testing.T) payment.Payer {
					ps := mocks_payment.NewPayer(t)
					call1 := ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Failed, nil).Once()
					ps.EXPECT().RefundPayment(mock.Anything).Return(nil).Once().NotBefore(call1).After(time.Second * 6)
					return ps
				},
			},
			args: args{
				booking: &validOrder,
			},
			wantErr: true,
		},
		{
			name: "success: payment state unknown, we let payer fill out reason why",
			fields: fields{
				PaymentService: func(t *testing.T) payment.Payer {
					ps := mocks_payment.NewPayer(t)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Unknown, nil).Once()
					ps.EXPECT().UpdateReason(&payment.Reason{}).Return(nil).Run(func(r *payment.Reason) {
						r.Msg = "it's not you, it's us"
					})
					return ps
				},
			},
			args: args{
				booking: &validOrder,
			},
			wantErr: false,
		},
		{
			name: "success: payment state unknown, we let payer return the reason why",
			fields: fields{
				PaymentService: func(t *testing.T) payment.Payer {
					ps := mocks_payment.NewPayer(t)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Unknown, nil).Once()
					ps.EXPECT().UpdateReason(&payment.Reason{}).Return(errors.New("error"))
					ps.EXPECT().UpdateAndReturnReason(&payment.Reason{}).RunAndReturn(func(r *payment.Reason) *payment.Reason {
						r.Msg = "get the message!"
						return r
					})
					return ps
				},
			},
			args: args{
				booking: &validOrder,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &BookingService{
				PaymentService: tt.fields.PaymentService(t),
			}
			if err := ps.ProcessBooking(tt.args.booking); (err != nil) != tt.wantErr {
				t.Errorf("ProcessBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
