package sieve

import (
	"math"
)

// Sieve returns a list of primes using Sieve of Eratosthenes up to limit
func Sieve(limit int) []int {
	flimit := float64(limit)
	approxPrimes := uint8(flimit / math.Log(flimit) - 1)
	primes := make([]int, 0, approxPrimes) // #primes up to n ~= n/(log(n)-1)
	list := make([]uint8, limit-1)
	for i := 0; i < len(list); i++ {
		if list[i] & 1 == 1 {
			continue
		}

		nextPrime := i + 2
		primes = append(primes, nextPrime)
		for j := (nextPrime - 2) + nextPrime; j < len(list); j += nextPrime {
			list[j] |= 1
		}
	}

	return primes
}
