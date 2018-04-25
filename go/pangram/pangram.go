package pangram

// IsPangram checks a string s to see if it contains at least one of each char in [a-z]
func IsPangram(s string) bool {
	if len(s) < 26 {
		return false
	}

	var flags uint32 // instead of a map, use the lower 26 bits as an alphabet flag for each char
	var c byte
	for i := 0; i < len(s); i++ {
		c = s[i] | 32
		if 'a' <= c && c <= 'z' {
			flags |= 1 << (c - 'a')
		}

	}

	return flags == 1<<26-1
}
