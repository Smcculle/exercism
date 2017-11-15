package luhn

import "strings"

func Valid(s string) bool {
	if len(strings.TrimSpace(s)) == 1 {
		return false
	}
	f := false
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		d := int(rune(s[i]) - '0')
		if d == -16 {
			continue
		}

		if f {
			sum += getVal(d * 2)
		} else {
			sum += d
		}
		f = !f
	}

	return sum%10 == 0

}

func getVal(i int) int {
	if i > 9 {
		return i - 9
	}
	return i
}
