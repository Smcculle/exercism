package lsproduct

import (
	"errors"
)

// LargestSeriesProduct calculates the largest product of `span` digits in digits
func LargestSeriesProduct(digits string, span int) (max int64, err error) {

	if span > len(digits) || span < 0 {
		return -1, errors.New("invalid span")
	}

	for i := 0; i+span <= len(digits); i++ {
		max, err = maxProduct(max, digits[i:i+span])
		if err != nil {
			return -1, err
		}
	}

	return
}

// maxProduct returns max of current and the product of each digit in digits
func maxProduct(current int64, digits string) (p int64, err error) {
	p = 1
	for i := 0; i < len(digits); i++ {
		digit := digits[i] - '0'
		if 0 <= digit && digit <= 9 {
			p *= int64(digit)
		} else {
			return 0, errors.New("non digit in string")
		}
	}

	if current >= p {
		return current, nil
	}

	return p, nil
}
