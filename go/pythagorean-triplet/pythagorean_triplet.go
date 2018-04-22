package pythagorean

import "fmt"

type Factors map[int]int
type Triplet [3] int

func Range(min, max int) []Triplet {
	var trips []Triplet
	fmt.Println("\t min max ", min, max )
	r := 2
	for {
		newTrips, valid := genTriplets(r)
		newTrips = filter(newTrips, min, max)
		if len(newTrips) == 0 && valid {
			break
		}
		trips = append(trips, newTrips...)
		r++
	}

	return trips
}

func Sum(p int ) []Triplet {
	return nil
}

func filter(trips []Triplet, min, max int)(valid []Triplet) {
	for _, t := range trips {
		isvalid := isValid(t, min, max)
		fmt.Printf("testing %v, isvalid? %v\n", t, isvalid)
		if isvalid {
			valid = append(valid, t)
		}
	}
	fmt.Println("returning set ", valid, " from set ", trips)
	return
}

func isValid(t Triplet, min, max int) bool {
	return t[0] >= min && t[2] <= max
}

func genTriplets(r int) ([]Triplet, bool){

	var trips []Triplet

	if r*r%2 != 0 {
		return nil, false
	}

	factors := getFactors(r * r / 2)
	for s, t := range factors {
		trips = append(trips, convertRST(r, s, t))
	}

	return trips, true
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

