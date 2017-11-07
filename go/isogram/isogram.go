package isogram

import (
	"unicode"
)

func IsIsogram(s string) bool {

	lset := make(map[rune]bool)

	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		lc := unicode.ToLower(c)
		if lset[lc] {
			return false
		} else {
			lset[lc] = true
		}

	}
	return true
}
