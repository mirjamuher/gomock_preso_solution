package payment

type Payload struct {
	TotalPrice float64
	Method string
}

type State int

const (
	Unknown State = iota
	Initiated
	Succeeded
	Failed
)