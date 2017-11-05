// Package acronym returns an acronym from a given string
package acronym

import "strings"

// Abbreviate returns an acronym from a given string
func Abbreviate(s string) (acronym string) {

	for _, v := range strings.Split(strings.Title(s), " ") {
		acronym += string(v[0])
	}
	return acronym
}
