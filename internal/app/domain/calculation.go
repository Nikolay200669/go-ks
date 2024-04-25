package domain

type CalculationRequest struct {
	A uint64 `json:"a"`
	B uint64 `json:"b"`
}

type CalculationResponse struct {
	FactorialA uint64 `json:"factorialA"`
	FactorialB uint64 `json:"factorialB"`
	_          struct{}
}
