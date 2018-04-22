package pythagorean

type Factors map[int]int
type Triplet [3]int

// Range returns all triplets where each side lies in [min..max]
func Range(min, max int) []Triplet {
	var trips []Triplet

	for r := 2; r < max/2; r++ {
		newTrips := genTriplets(r)
		trips = append(trips, newTrips...)
	}

	return filterRange(trips, min, max)
}

// Sum returns all triplets where x+y+z=p
func Sum(p int) []Triplet {
	return filterSum(Range(0, p/2), p)
}

func filterSum(trips []Triplet, p int) (valid []Triplet) {
	for _, t := range trips {
		if t[0]+t[1]+t[2] == p {
			valid = append(valid, t)
		}
	}
	return
}

func filterRange(trips []Triplet, min, max int) (valid []Triplet) {
	for _, t := range trips {
		if t[0] >= min && t[2] <= max {
			valid = append(valid, t)
		}
	}
	return
}

// genTriplets returns a set of pythagorean triplets using Dickson's method
func genTriplets(r int) []Triplet {

	var trips []Triplet

	if r*r%2 != 0 {
		return nil
	}

	for s, t := range getFactors(r * r / 2) {
		trips = append(trips, convertRST(r, s, t))
	}

	return trips
}

func convertRST(r, s, t int) Triplet {
	x, y, z := r+s, r+t, r+s+t
	return Triplet{x, y, z}
}

func getFactors(n int) Factors {
	factors := Factors{1: n}
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			f1, f2 := i, n/i
			if factors[f2] == 0 {
				factors[f1] = f2
			}
		}
	}
	return factors
}
