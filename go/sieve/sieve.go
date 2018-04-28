package sieve

// Sieve returns a list of primes using Sieve of Eratosthenes up to limit
func Sieve(limit int) (primes []int) {
	list := make([]uint8, limit+1)
	for i := 2; i < len(list); i++ {
		if list[i]&1 == 1 {
			continue
		}

		primes = append(primes, i)
		for j := i * 2; j < len(list); j += i {
			list[j] |= 1
		}
	}

	return primes
}
