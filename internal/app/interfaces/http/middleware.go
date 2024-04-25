package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Nikolay200669/go-ks/internal/app/domain"
	"github.com/julienschmidt/httprouter"
)

// ValidateInput validates the input of the request
func ValidateInput(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var reqBody domain.CalculationRequest

		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if reqBody.A < 0 || reqBody.B < 0 {
			errorMsg := struct {
				Error string `json:"error"`
			}{
				Error: "Incorrect input",
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorMsg)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "RequestBody", reqBody))
		next(w, r, ps)
	}
}
