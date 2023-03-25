package payment

type Payment struct {
	TotalPrice float64
	Method string
}

type Reason struct {
	Msg string
}

type PaymentState int

const (
	Unknown PaymentState = iota
	Initiated
	Succeeded
	Failed
)