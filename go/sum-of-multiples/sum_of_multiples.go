package summultiples

type set map[int]bool

func (s *set) add(i int) {
	if !(*s)[i] {
		(*s)[i] = true
	}
}
func (s set) sum() (sum int) {
	for k := range s {
		sum += k
	}
	return
}

func (s *set) addMults(limit int, divisor int) {

	for multiple := divisor; multiple < limit; multiple += divisor {
		s.add(multiple)
	}

	return
}

func SumMultiples(limit int, divisors ...int) int {
	var mults = make(set)

	for _, divisor := range divisors {
		mults.addMults(limit, divisor)
	}

	return mults.sum()
}
