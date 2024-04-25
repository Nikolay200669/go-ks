package http

import (
	"encoding/json"
	"net/http"

	"github.com/Nikolay200669/go-ks/internal/app/application"
	"github.com/Nikolay200669/go-ks/internal/app/domain"

	"github.com/julienschmidt/httprouter"
)

func CalculateHandler(service *application.CalculateService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Retrieve the request body from context
		reqBody := r.Context().Value("RequestBody").(domain.CalculationRequest)

		// Calculate factorials
		response := service.CalculateFactorials(reqBody)

		// Encode response
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set status code
		w.WriteHeader(http.StatusOK)
	}
}
