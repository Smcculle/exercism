package pangram

// IsPangram checks a string s to see if it contains at least one of each char in [a-z]
func IsPangram(s string) bool {
	if len(s) < 26 {
		return false
	}

	var found struct{}
	alphabet := make(map[int32]struct{})

	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			r = r | 32
		}
		if 'a' <= r && r <= 'z' {
			alphabet[r] = found
		}
	}

	return len(alphabet) == 26
}
