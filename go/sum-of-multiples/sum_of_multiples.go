package summultiples

func addMults(limit int, divisor int, divisors []int) (sum int) {

	for multiple := divisor; multiple < limit; multiple += divisor {
		sum += multiple
		checkRepeat(divisors, divisor, multiple, &sum)
	}

	return
}

//checkRepeat removes any multiple that is common to another divisor from sum
func checkRepeat(divisors []int, divisor int, multiple int, sum *int) {
	for _, common := range divisors {
		if common > divisor && multiple%common == 0 {
			*sum -= multiple
			break
		}
	}
}

func SumMultiples(limit int, divisors ...int) (sum int) {

	for _, divisor := range divisors {
		sum += addMults(limit, divisor, divisors)
	}

	return
}
