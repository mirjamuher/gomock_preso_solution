package payment

type Payment struct {
	TotalPrice float64
	Method string
}

type PaymentState int

const (
	Unknown PaymentState = iota
	Initiated
	Success
	Failed
)