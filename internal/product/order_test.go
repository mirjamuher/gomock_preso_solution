package product

import (
	"SubscriptionService/internal/payment"
	"database/sql"
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
		// TODO: Add test cases.
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
