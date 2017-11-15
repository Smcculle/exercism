package diffsquares

import "math"

func SquareOfSums(n int) int {
	sum := 0
	for i := 1; i < n+1; i++ {
		sum += i
	}
	return pow(sum, 2)
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i < n+1; i++ {
		sum += pow(i, 2)
	}
	return sum

}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
