package domain

import "math/big"

type CalculationRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type CalculationResponse struct {
	FactorialA *big.Int `json:"factorialA"`
	FactorialB *big.Int `json:"factorialB"`
	_          struct{}
}
