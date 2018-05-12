package anagram

import (
	"sort"
)

// ByByte implements the sort interface
type ByByte []byte

func (r ByByte) Len() int           { return len(r) }
func (r ByByte) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByByte) Less(i, j int) bool { return r[i] < r[j] }

// Detect takes a subject and a list of candidates and returns all anagrams
func Detect(subject string, candidates []string) (anagrams []string) {
	for _, c := range candidates {
		if len(c) != len(subject) {
			continue
		}

		if isAnagram(subject, c) {
			anagrams = append(anagrams, c)
		}

	}
	return
}

func isAnagram(s1, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}

	var b1, b2 = make(ByByte, n), make(ByByte, n)
	var ok bool

	for i := 0; i < n; i++ {
		b1[i] = s1[i] | 32
		b2[i] = s2[i] | 32
		ok = ok || (b1[i] != b2[i]) // word is not anagram of itself
	}

	if !ok {
		return false
	}

	sort.Sort(b1)
	sort.Sort(b2)

	// check lowercase byte values for case insensitivity
	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}
