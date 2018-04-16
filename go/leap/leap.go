package leap

func IsLeapYear(n int) bool {
	if n % 4 != 0 {
		return false
	}

	if n % 100 == 0 && n % 400 != 0 {
		return false
	}

	return true
}
