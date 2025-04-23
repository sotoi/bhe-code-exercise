package main

// write integration test for the api
import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "ssse-exercise-sieve/api/router"
	"ssse-exercise-sieve/pkg/sieve"

	"github.com/stretchr/testify/assert"
)

func TestGetNthPrimeIntegration(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedPrime  int64
		expectError    *string
	}{
		{
			name:           "Valid input - 19th prime",
			queryParam:     "19",
			expectedStatus: http.StatusOK,
			expectedPrime:  71,
			expectError:    nil,
		},
		{
			name:           "Valid input - 1st prime",
			queryParam:     "1",
			expectedStatus: http.StatusOK,
			expectedPrime:  3,
			expectError:    nil,
		},
		{
			name:           "Invalid input - negative number",
			queryParam:     "-1",
			expectedStatus: http.StatusInternalServerError,
			expectedPrime:  0,
			expectError:    &handler.InternalServerError,
		},
		{
			name:           "Invalid input - non-numeric",
			queryParam:     "test",
			expectedStatus: http.StatusBadRequest,
			expectedPrime:  0,
			expectError:    &handler.InvalidParamError,
		},
		{
			name:           "Invalid input - empty",
			queryParam:     "",
			expectedStatus: http.StatusBadRequest,
			expectedPrime:  0,
			expectError:    &handler.MissingParamError,
		},
	}

	sieveService := sieve.NewSieve()
	handler := handler.NewPrimeHandler(sieveService)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create request with test query parameter
			req, err := http.NewRequest("GET", "/primes?n="+tt.queryParam, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler.GetNthPrime(rr, req)

			// Check status code
			assert.Equal(t, tt.expectedStatus, rr.Code)

			// Parse response

			if tt.expectError != nil {
				// Verify error response
				assert.Contains(t, rr.Body.String(), *tt.expectError)
			} else {
				var response map[string]interface{}
				err = json.Unmarshal(rr.Body.Bytes(), &response)
				if err != nil {
					t.Fatal(err)
				}
				// Verify success response
				prime, ok := response["nth_prime"].(float64)
				assert.True(t, ok)
				assert.Equal(t, float64(tt.expectedPrime), prime)
			}
		})
	}
}
