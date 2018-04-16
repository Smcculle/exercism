package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = -1 // not a triangle
	Sca Kind = 0  // scalene
	Iso Kind = 1  // isosceles
	Equ Kind = 3  // equilateral
)

// ShareWith should have a comment documenting it.
func KindFromSides(a, b, c float64) (k Kind) {

	if !isValid(a, b, c) {
		return NaT
	}

	if a == b {
		k++
	}
	if a == c {
		k++
	}
	if b == c {
		k++
	}
	return
}

func isValid(a, b, c float64) bool {

	sum := a + b + c

	if math.IsNaN(sum) || math.IsInf(sum, 1) || math.IsInf(sum, -1) {
		return false
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if a > b+c || b > a+c || c > a+b {
		return false
	}

	return true
}
