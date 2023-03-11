package product

import (
	"database/sql"
	"github.com/mirjamuher/gomock_preso_solution/full_service/internal/payment"
	"testing"
)

func TestProductService_CreateOrder(t *testing.T) {
	type fields struct {
		db             *sql.DB
		paymentService payment.PaymentService
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
			name: "",
			fields: fields{
				db:             &sql.DB{},
				paymentService: payment.PaymentService{},
			},
			args: args{
				order: &Order{
					Product: Product{
						Price: 0,
					},
					Quantity:      0,
					PaymentMethod: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &ProductService{
				db:             tt.fields.db,
				paymentService: tt.fields.paymentService,
			}
			if err := ps.CreateOrder(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
