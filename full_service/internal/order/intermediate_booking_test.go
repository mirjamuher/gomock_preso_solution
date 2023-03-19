package order

//func TestProductService_CreateOrders(t *testing.T) {
//	twoOrders := Booking{
//		Quantity:      2,
//	}
//
//	type fields struct {
//		paymentService func() payment.Payer
//	}
//	type args struct {
//		order *Booking
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		{
//			name: "success: paid & stored multiple orders successfully",
//			fields: fields{
//				paymentService: func() payment.Payer {
//					ctrl := gomock.NewController(t)
//					p := mock_payment.NewMockPayer(ctrl)
//					gomock.InOrder(
//						p.EXPECT().ProcessPayment(gomock.Any()).Return(payment.Succeeded, nil).Times(2),
//					)
//					return p
//				},
//			},
//			args: args{
//				order: &twoOrders,
//			},
//			wantErr: false,
//		},
//		{
//			name: "success: one order paid successfully, one failed -> storing both, no error returned",
//			fields: fields{
//				paymentService: func() payment.Payer {
//					ctrl := gomock.NewController(t)
//					p := mock_payment.NewMockPayer(ctrl)
//					gomock.InOrder(
//						p.EXPECT().ProcessPayment(gomock.Any()).Return(payment.Succeeded, nil),
//						p.EXPECT().ProcessPayment(gomock.Any()).Return(payment.Failed, errors.New("")),
//					)
//					return p
//				},
//			},
//			args: args{
//				order: &twoOrders,
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			ps := &BookingService{
//				db:             &sql.DB{}, // in a full example, this would be mocked too
//				paymentService: tt.fields.paymentService(),
//			}
//			if err := ps.CreateOrders(tt.args.order); (err != nil) != tt.wantErr {
//				t.Errorf("CreateOrders() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
