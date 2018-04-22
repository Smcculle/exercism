package etl

import "strings"

func Transform(in map[int][]string) map[string]int {

	out := make(map[string]int)
	for point, vals := range in {
		for _, letter := range vals {
			out[strings.ToLower(letter)] = point
		}
	}

	return out
}
