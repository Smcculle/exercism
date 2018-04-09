package sublist

import (
	"reflect"
	"sort"
)

type Relation string

const (
	sublist   Relation = "sublist"
	superlist Relation = "superlist"
	equal     Relation = "equal"
	unequal   Relation = "unequal"
)

// containsList assumes len(listOne) < len(listTwo) and they are sorted
func containsList(listOne, listTwo []int) bool {
	n1, n2 := len(listOne), len(listTwo)
	if n1 > n2 {
		panic("listTwo should be bigger than listOne")
	}

	j := sort.Search(n2, func(i int) bool { return listTwo[i] == listOne[0] })

	if j == n2 {
		return false
	}

	for i, j := 0, j+1; i < n1; i, j = i+1, j+1 {
		if listOne[i] != listTwo[j] {
			return false
		}
	}
	return true
}

func checkEdgeCases(listOne, listTwo []int, n1, n2 int) Relation {

	if n1 == n2 && reflect.DeepEqual(listOne, listTwo) {
		return equal
	} else if n1 == 0 {
		return sublist
	} else if n2 == 0 {
		return superlist
	}
	return ""
}

func Sublist(listOne, listTwo []int) Relation {
	sort.Ints(listOne)
	sort.Ints(listTwo)
	n1, n2 := len(listOne), len(listTwo)

	if edgeCase := checkEdgeCases(listOne, listTwo, n1, n2); edgeCase != "" {
		return edgeCase
	}

	if n1 > n2 {
		if containsList(listTwo, listOne) {
			return superlist
		}
		return unequal
	}

	if containsList(listOne, listTwo) {
		return sublist
	}
	return unequal
}
