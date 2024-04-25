package application

import "github.com/Nikolay200669/go-ks/internal/app/domain"

type CalculateService struct{}

func (s *CalculateService) CalculateFactorials(request domain.CalculationRequest) domain.CalculationResponse {
	factorialA := factorial(request.A)
	factorialB := factorial(request.B)

	return domain.CalculationResponse{
		FactorialA: factorialA,
		FactorialB: factorialB,
	}
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
