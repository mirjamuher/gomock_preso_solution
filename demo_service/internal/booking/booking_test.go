package booking

import (
	"errors"
	"github.com/mirjamuher/gomock_preso_solution/demo_service/internal/payment"
	"github.com/mirjamuher/gomock_preso_solution/demo_service/internal/payment/mocks"
	"testing"
)

func TestBookingService_CreateBooking(t *testing.T) {
	type fields struct {
		paymentService func() payment.Payer
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
			name: "success: payment was successful",
			fields: fields{
				paymentService: func() payment.Payer {
					ps := mocks.NewPayer(t)
					//ps.On("ProcessPayment", getValidPayload()).Return(payment.Succeeded, nil)
					ps.EXPECT().ProcessPayment(getValidPayload()).Return(payment.Succeeded, nil)
					return ps
				},
			},
			args: args{
				booking: getValidOrder(),
			},
			wantErr: false,
		},
		{
			name: "success: payment was successful",
			fields: fields{
				paymentService: func() payment.Payer {
					ps := mocks.NewPayer(t)
					//ps.On("ProcessPayment", getValidPayload()).Return(payment.Succeeded, nil)
					ps.EXPECT().ProcessPayment(getValidPayload()).Return(payment.Failed, errors.New("error"))
					return ps
				},
			},
			args: args{
				booking: getValidOrder(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &BookingService{
				paymentService: tt.fields.paymentService(),
			}
			if err := ps.CreateBooking(tt.args.booking); (err != nil) != tt.wantErr {
				t.Errorf("CreateBooking() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func getValidOrder() *Booking {
	return &Booking{
		product:       Product{
			price: 100,
		},
		quantity:      1,
		paymentMethod: "CREDIT",
		paymentState:  0,
	}
}

func getValidPayload() *payment.Payload {
	return &payment.Payload{
		TotalPrice: 100,
		Method:     "CREDIT",
	}
}
