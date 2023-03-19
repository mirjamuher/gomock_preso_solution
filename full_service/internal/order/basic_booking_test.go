package order

import (
	"database/sql"
	"errors"
	"github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
	"github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductService_CreateOrder(t *testing.T) {
	validOrder := Booking{
		Product:       Product{
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
		paymentService func() payment.Payer
	}
	type args struct {
		order *Booking
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success: payment & persistence successful",
			fields: fields{
				paymentService: func() payment.Payer {
					ps := mocks.NewPayer(t)
					//ps.On("ProcessPayment", &validPayment).Return(payment.Succeeded, nil)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Succeeded, nil)
					return ps
				},
			},
			args: args{
				order: &validOrder,
			},
			wantErr: false,
		},
		{
			name: "error: process payment failed",
			fields: fields{
				paymentService: func() payment.Payer {
					ps := mocks.NewPayer(t)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Failed, errors.New("error"))
					return ps
				},
			},
			args: args{
				order: &validOrder,
			},
			wantErr: true,
		},
		{
			name: "error: process payment returned other state than success",
			fields: fields{
				paymentService: func() payment.Payer {
					ps := mocks.NewPayer(t)
					ps.EXPECT().ProcessPayment(&validPayment).Return(payment.Unknown, nil)
					return ps
				},
			},
			args: args{
				order: &validOrder,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &BookingService{
				db:             &sql.DB{},
				paymentService: tt.fields.paymentService(),
			}
			if err := ps.CreateBooking(tt.args.order); (err != nil) != tt.wantErr {
				if tt.wantErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}
