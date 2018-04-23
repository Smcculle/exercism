package pangram

import "unicode"

func IsPangram(s string) bool {
	if len(s) < 26 {
		return false
	}

	var found struct{}
	alphabet := make(map[int32]struct{})

	for _, r := range s {
		if unicode.IsLetter(r) {
			alphabet[r|32] = found
		}
	}

	return len(alphabet) == 26
}
