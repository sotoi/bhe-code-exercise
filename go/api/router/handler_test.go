package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"ssse-exercise-sieve/pkg/sieve"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNthPrime(t *testing.T) {
	mockSieve := sieve.NewMockSieve()
	handler := NewPrimeHandler(mockSieve)

	tests := []struct {
		name           string
		queryParam     string
		mockReturn     int64
		mockError      bool
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Missing parameter",
			queryParam:     "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   MissingParamError,
		},
		{
			name:           "Invalid parameter",
			queryParam:     "abc",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   InvalidParamError,
		},
		{
			name:           "Negative parameter",
			queryParam:     "-5",
			mockReturn:     -1,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   InternalServerError,
		},
		{
			name:           "Internal server error",
			queryParam:     "10",
			mockReturn:     -1,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   InternalServerError,
		},
		{
			name:           "Valid parameter",
			queryParam:     "5",
			mockReturn:     11,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"nth_prime":11}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockReturn != 0 || tt.mockError {
				n, _ := strconv.ParseInt(tt.queryParam, 10, 64)
				mockSieve.On("NthPrime", n).Return(tt.mockReturn)
			}

			req := httptest.NewRequest(http.MethodGet, "/prime?n="+tt.queryParam, nil)
			rec := httptest.NewRecorder()

			handler.GetNthPrime(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			assert.Equal(t, fmt.Sprintf("%v\n", tt.expectedBody), rec.Body.String())

		})
	}
}
