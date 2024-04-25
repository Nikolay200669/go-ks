package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// CalcHandler handles the /calculate endpoint
func CalcHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Retrieve the parsed request body from the context
	reqBody := r.Context().Value("RequestBody").(struct {
		A uint64 `json:"a"`
		B uint64 `json:"b"`
	})

	// Calculate the factorial of A and B
	factorialA := factorial(reqBody.A)
	factorialB := factorial(reqBody.B)

	// Create a response body
	resBody := struct {
		FactorialA uint64 `json:"factorialA"`
		FactorialB uint64 `json:"factorialB"`
	}{
		FactorialA: factorialA,
		FactorialB: factorialB,
	}

	// Encode the response body
	err := json.NewEncoder(w).Encode(resBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the content type
	w.Header().Set("Content-Type", "application/json")

	// Set the status code
	w.WriteHeader(http.StatusOK)
}

func main() {
	router := httprouter.New()
	router.POST("/calculate", ValidateInput(CalcHandler))

	log.Fatal(http.ListenAndServe(":8989", router))
}

// Calculate the factorial of a number
func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func ValidateInput(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Parse the request body
		var reqBody struct {
			A uint64 `json:"a"`
			B uint64 `json:"b"`
		}

		// Decode the request body
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Check if A and B are non-negative
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

		// Add the parsed body to the request context
		r = r.WithContext(context.WithValue(r.Context(), "RequestBody", reqBody))

		// Call the next handler
		next(w, r, ps)
	}
}
