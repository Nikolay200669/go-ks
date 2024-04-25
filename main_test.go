package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalcHandler(t *testing.T) {

	const url = "/calculate"

	t.Run("Run test OK", func(t *testing.T) {
		// Prepare a sample request body
		reqBody := map[string]uint64{"a": 5, "b": 3}
		reqBodyJSON, err := json.Marshal(reqBody)
		require.NoError(t, err, "failed to marshal JSON")

		// Create a request with the sample body
		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBodyJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Create a router and handle the request
		router := httprouter.New()
		router.POST(url, ValidateInput(CalcHandler))
		router.ServeHTTP(rr, req)

		require.Equal(t, rr.Code, http.StatusOK, "unexpected status code")

		// Decode the response body
		var resBody struct {
			FactorialA uint64 `json:"factorialA"`
			FactorialB uint64 `json:"factorialB"`
		}
		err = json.NewDecoder(rr.Body).Decode(&resBody)
		require.NoError(t, err)

		// Check the factorial values
		expectedFactorialA := uint64(120) // 5!
		expectedFactorialB := uint64(6)   // 3!

		assert.Equal(t, resBody.FactorialA, expectedFactorialA)
		assert.Equal(t, resBody.FactorialB, expectedFactorialB)
	})

	t.Run("Run test middleware", func(t *testing.T) {
		// Prepare a sample request body
		reqBody := map[string]interface{}{"a": "5", "b": 3}
		reqBodyJSON, err := json.Marshal(reqBody)
		require.NoError(t, err)

		// Create a request with the sample body
		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBodyJSON))
		req.Header.Set("Content-Type", "application/json")

		// Create a ResponseRecorder to record the response
		rr := httptest.NewRecorder()

		// Create a router and handle the request
		router := httprouter.New()
		router.POST(url, ValidateInput(CalcHandler))
		router.ServeHTTP(rr, req)

		require.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "json: cannot unmarshal string into Go struct field .a of type uint64\n", rr.Body.String())
	})
}
