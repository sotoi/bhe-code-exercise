package main

import (
	"net/http"
	handler "ssse-exercise-sieve/api/router"
	"ssse-exercise-sieve/pkg/sieve"
)

func main() {
	sieveService := sieve.NewSieve()
	router :=initializeRouter(sieveService)
	http.ListenAndServe(":8080", router)
}

func initializeRouter(sieve sieve.Sieve) *http.ServeMux {
	mux := http.NewServeMux()
	// Initialize the PrimeHandler with the SieveService
	primeHandler := handler.NewPrimeHandler(sieve)
	// Define the route for getting the nth prime number
	mux.HandleFunc("/primes", primeHandler.GetNthPrime)
	return mux
}
