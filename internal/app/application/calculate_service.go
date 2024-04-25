package application

import "github.com/Nikolay200669/go-ks/internal/app/domain"

type CalculateService struct{}

func (s *CalculateService) CalculateFactorials(request domain.CalculationRequest) domain.CalculationResponse {
	return domain.CalculationResponse{
		FactorialA: factorial(request.A),
		FactorialB: factorial(request.B),
	}
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
