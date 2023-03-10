package product

type Order struct {
	Product Product
	Quantity int
	PaymentMethod string
}

func (o *Order) Validate() error {
	return nil
}

type Product struct{
	Price float64
}