package perfect

import (
	"errors"
	"math"
)

// Classification is an enum for number classification
type Classification int

// ErrOnlyPositive is an error that states input was not positive
var ErrOnlyPositive = errors.New("input must be positive")

// ClassificationPerfect
const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

func mod(n int64) (m int64) {
	return 0
}

// AliquotSum calculates the sum of each factor of n except n itself
func AliquotSum(n int64) (sum int64) {

	sum = -n
	for i := int64(1); i <= sqrt(n); i++ {
		if sum > n {
			return
		}
		
		if n%i == 0 {
			f1, f2 := i, n/i
			if f1 == f2 {
				sum += f1
			} else {
				sum = sum + f1 + f2
			}
		}
	}

	return
}

// Classify calculates the aliquot sum of num and returns a Classification
func Classify(num int64) (c Classification, err error) {
	if num <= 0 {
		return c, ErrOnlyPositive
	}

	asum := AliquotSum(num)

	if asum > num {
		c = ClassificationAbundant
	} else if asum < num {
		c = ClassificationDeficient
	} else {
		c = ClassificationPerfect
	}

	return
}

func sqrt(i int64) int64 {
	return int64(math.Sqrt(float64(i)))
}
