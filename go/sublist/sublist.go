package sublist

type Relation string

const (
	sublist   Relation = "sublist"
	superlist Relation = "superlist"
	equal     Relation = "equal"
	unequal   Relation = "unequal"
)

// containsList assumes len(listOne) < len(listTwo)
func containsList(listOne, listTwo []int) bool {
	n1, n2 := len(listOne), len(listTwo)
	if n1 > n2 {
		panic("listTwo should be bigger than listOne")
	}

	for i := 0; i < n2-n1+1; i++ {
		slice := listTwo[i : i+n1]
		if eq(listOne, slice) {
			return true
		}
	}

	return false
}

func Sublist(listOne, listTwo []int) Relation {
	n1, n2 := len(listOne), len(listTwo)

	if n1 > n2 {
		if containsList(listTwo, listOne) {
			return superlist
		}
	} else if containsList(listOne, listTwo) {
		if n1 == n2 {
			return equal
		}
		return sublist
	}

	return unequal
}

func eq(s1, s2 []int) bool {
	n1, n2 := len(s1), len(s2)
	if n1 != n2 {
		return false
	}

	for i := 0; i < n1; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
