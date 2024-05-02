package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Nikolay200669/go-ks/internal/app/application"
	iface "github.com/Nikolay200669/go-ks/internal/app/interfaces/http"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalcHandler(t *testing.T) {

	const url = "/calculate"

	t.Run("Run test OK", func(t *testing.T) {

		// Create an instance of the service
		calculateService := &application.CalculateService{}

		// Create a handler and apply middleware to it
		handler := iface.CalculateHandler(calculateService)
		mw := iface.ValidateInput(handler)

		// Create a test request
		reqBody := `{"a": 5, "b": 3}`
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(reqBody))
		require.NoError(t, err)

		// Create a recording response
		rr := httptest.NewRecorder()

		// Run the request
		mw(rr, req, httprouter.Params{})

		// Check the status code
		require.Equal(t, http.StatusOK, rr.Code)

		// Check the response body
		expectedBody := fmt.Sprint("{\"factorialA\":120,\"factorialB\":6}\n")
		assert.Equal(t, expectedBody, rr.Body.String())
	})

	t.Run("Run test middleware", func(t *testing.T) {

		calculateService := &application.CalculateService{}

		handler := iface.CalculateHandler(calculateService)
		mw := iface.ValidateInput(handler)

		reqBody := `{"a": "5", "b": 3}`
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(reqBody))
		require.NoError(t, err)

		rr := httptest.NewRecorder()

		mw(rr, req, httprouter.Params{})

		require.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "json: cannot unmarshal string into Go struct field CalculationRequest.a of type int\n", rr.Body.String())
	})
}
