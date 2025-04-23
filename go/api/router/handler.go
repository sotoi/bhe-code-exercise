package handler

import (
	"encoding/json"
	"net/http"
	"ssse-exercise-sieve/pkg/sieve"
	"strconv"
)

var (
	MissingParamError   = "Missing parameter: n"
	InvalidParamError   = "Invalid parameter: n"
	InternalServerError = "Error calculating prime"
)

type PrimeHandler struct {
	SieveService sieve.Sieve
}

type PrimeResponse struct {
	NthPrime int64 `json:"nth_prime"`
}

func NewPrimeHandler(sieveService sieve.Sieve) *PrimeHandler {
	return &PrimeHandler{
		SieveService: sieveService,
	}
}

func (h *PrimeHandler) GetNthPrime(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		http.Error(w, MissingParamError, http.StatusBadRequest)
		return
	}
	n, err := strconv.ParseInt(nStr, 10, 64)
	if err != nil {
		http.Error(w, InvalidParamError, http.StatusBadRequest)
		return
	}
	prime := h.SieveService.NthPrime(n)
	if prime == -1 {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PrimeResponse{NthPrime: prime})
}
