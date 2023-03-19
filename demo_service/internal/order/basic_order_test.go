package order

import (
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/mirjamuher/gomock_preso_solution/demo_service/internal/payment"
	mock_payment "github.com/mirjamuher/gomock_preso_solution/demo_service/internal/payment/mocks"
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
			name: "success",
			fields: fields{
				paymentService: func() payment.Payer {
					ctrl := gomock.NewController(t)
					ps := mock_payment.NewMockPayer(ctrl)
					// todo: use actual mocking object
					ps.EXPECT().ProcessPayment(gomock.Any()).Return(payment.Success, nil)
					return ps
				},
			},
			args: args{
				order: &validOrder,
			},
			wantErr: false,
		},
		{
			name: "success",
			fields: fields{
				paymentService: func() payment.Payer {
					ctrl := gomock.NewController(t)
					ps := mock_payment.NewMockPayer(ctrl)
					ps.EXPECT().ProcessPayment(gomock.Any()).Return(payment.Failed, errors.New(""))
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
			ps := &ProductService{
				db:             &sql.DB{},
				paymentService: tt.fields.paymentService(),
			}
			if err := ps.CreateOrder(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
