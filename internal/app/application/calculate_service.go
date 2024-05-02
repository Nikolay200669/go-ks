package application

import (
	"math/big"

	"github.com/Nikolay200669/go-ks/internal/app/domain"
)

type CalculateService struct{}

func (s *CalculateService) CalculateFactorials(request domain.CalculationRequest) domain.CalculationResponse {
	return domain.CalculationResponse{
		FactorialA: factorial(request.A),
		FactorialB: factorial(request.B),
	}
}

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}
