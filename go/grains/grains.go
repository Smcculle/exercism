package grains

import "errors"

func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("invalid input")
	}

	return 1 << uint64(n-1), nil
}

// Total returns sum of the grains, i.e. Sum(2^n) for n=1->63, which is 2^n+1 -1 or 2^64 - 1
func Total() uint64 {
	return 1 << uint64(64) - 1
}