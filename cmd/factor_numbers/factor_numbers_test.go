package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_factorNumbersBothImpl(t *testing.T) {
	primes = sieveToPrimes(eulersSieve(20_000_000))

	tests := []struct {
		name    string
		number  int
		factors []int
	}{
		{
			name:    "312680865509917",
			number:  312680865509917,
			factors: []int{7791799, 40129483},
		},
		{
			name:    "12345678901234",
			number:  12345678901234,
			factors: []int{2, 7, 73, 12079920647},
		},
		{
			name:    "64374108854777",
			number:  64374108854777,
			factors: []int{64374108854777},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFactors(tt.number)
			assert.Equal(t, tt.factors, got)
			assert.Equal(t, multiplySlice(got), tt.number)

			got = findFactorsSieve(tt.number)
			assert.Equal(t, tt.factors, got)
			assert.Equal(t, multiplySlice(got), tt.number)
		})
	}
}
