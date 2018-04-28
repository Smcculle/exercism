package sieve

// Sieve returns a list of primes using Sieve of Eratosthenes up to limit
func Sieve(limit int) (primes []int) {
	marked := make([]bool, limit+1)
	for next := 2; next < len(list); next++ {
		if marked[next] {
			continue
		}
		primes = append(primes, next)

		for j := next * next; j <= limit; j += i {
			marked[j] = true
		}
	}
	return
}
