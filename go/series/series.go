package series

func All(n int, s string) []string {
	series := make([]string, len(s))
	var i int
	for ; i+n <= len(s); i++ {
		series[i] = s[i : i+n]
	}

	return series[:i]
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}
