package order

import (
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
	mock_payment "github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment/mocks"
	"testing"
)

func TestProductService_CreateOrder(t *testing.T) {
	validOrder := Order{}

	type fields struct {
		paymentService func() payment.Payer
	}
	type args struct {
		order *Order
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
					ctrl := gomock.NewController(t)
					p := mock_payment.NewMockPayer(ctrl)
					p.EXPECT().ProcessPayment(gomock.Any()).Return(payment.Succeeded, nil)
					return p
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
					ctrl := gomock.NewController(t)
					p := mock_payment.NewMockPayer(ctrl)
					p.EXPECT().ProcessPayment(gomock.Any()).Return(payment.Failed, errors.New(""))
					return p
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
					ctrl := gomock.NewController(t)
					p := mock_payment.NewMockPayer(ctrl)
					p.EXPECT().ProcessPayment(gomock.Any()).Return(payment.Unknown, nil)
					return p
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
			ps := &ProductService{
				db:             &sql.DB{},
				paymentService: tt.fields.paymentService(),
			}
			if err := ps.CreatePurchase(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("CreatePurchase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
