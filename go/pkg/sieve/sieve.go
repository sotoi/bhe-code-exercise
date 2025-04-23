package sieve

import "math"

type Sieve interface {
	NthPrime(n int64) int64
}
type SieveImpl struct {
	primes []int64
}

func NewSieve() Sieve {
	return &SieveImpl{
		primes: []int64{2}, // start with 2, the only even prime
	}
}

func (s *SieveImpl) NthPrime(n int64) int64 {
	if n < 0 {
		return -1
	}

	if int64(len(s.primes)) > n {
		return s.primes[n]
	}

	// Estimate upper bound using prime number theorem
	// p(n) ~ n * log(n) + n * log(log(n))
	est := float64(n)
	limit := int(est*math.Log(est) + est*math.Log(math.Log(est)))
	if limit < 100 {
		limit = 100
	}

	sieve := make([]bool, (limit+1)/2) // Only odd indices, index i -> number 2*i + 1

	// Perform Sieve of Eratosthenes for odd numbers
	for i := 3; i*i <= limit; i += 2 {
		if !sieve[i/2] {
			for j := i * i; j <= limit; j += 2 * i {
				sieve[j/2] = true
			}
		}
	}

	// Collect primes
	s.primes = []int64{2}
	for i := 1; 2*i+1 <= limit; i++ {
		if !sieve[i] {
			s.primes = append(s.primes, int64(2*i+1))
		}
		if int64(len(s.primes)) > n {
			break
		}
	}

	if int64(len(s.primes)) > n {
		return s.primes[n]
	}
	return -1
}
