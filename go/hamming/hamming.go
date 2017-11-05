// Package hamming calculates the Hamming distance between two given strings
package hamming

import "errors"

// Distance iterates over the args, returning the number of differences.
func Distance(a, b string) (int, error) {

	d := 0
	if len(a) != len(b) {
		return -1, errors.New("strings must have same length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d += 1
		}
	}
	return d, nil
}
