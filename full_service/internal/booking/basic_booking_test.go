package booking

import (
	"errors"
	"github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
	"github.com/mirjamuher/gomock_preso_solution/full_service/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductService_CreateOrder(t *testing.T) {
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
			name: "success: payment successful",
			fields: fields{
				PaymentService: func(t *testing.T) payment.Payer {
					ps := mocks_payment.NewPayer(t)
					//ps.On("ProcessPayment", &validPayment).Return(payment.Succeeded, nil)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Succeeded, nil)
					return ps
				},
			},
			args: args{
				booking: &validOrder,
			},
			wantErr: false,
		},
		{
			name: "error: process payment failed",
			fields: fields{
				PaymentService: func(t *testing.T) payment.Payer {
					ps := mocks_payment.NewPayer(t)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Failed, errors.New("error"))
					return ps
				},
			},
			args: args{
				booking: &validOrder,
			},
			wantErr: true,
		},
		{
			name: "error: process payment returned other state than success",
			fields: fields{
				PaymentService: func(t *testing.T) payment.Payer {
					ps := mocks_payment.NewPayer(t)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Unknown, nil)
					return ps
				},
			},
			args: args{
				booking: &validOrder,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &BookingService{
				PaymentService: tt.fields.PaymentService(t),
			}
			if err := ps.CreateBooking(tt.args.booking); (err != nil) != tt.wantErr {
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}
