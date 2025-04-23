package sieve

import "github.com/stretchr/testify/mock"

// MockSieve is a mock implementation of the sieve.Sieve interface
type MockSieve struct {
	mock.Mock
}

func NewMockSieve() *MockSieve {
	return &MockSieve{}
}

func (m *MockSieve) NthPrime(n int64) int64 {
	args := m.Called(n)
	return args.Get(0).(int64)
}