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

		if isAnagramByPrime(subject, c) {
			anagrams = append(anagrams, c)
		}

	}
	return
}

// first 26 primes used to calculate a prime hash
var prime = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}


// primeAnagram uses fundamental theorem of arithmetic to calculate a hash unique among anagrams. 
// i.e., if len(s1)==len(s2), then h(s1)==h(s2) iff s1 and s2 are anagrams. 
//BenchmarkDetectAnagrams-4  	729 ns/op	     224 B/op	       9 allocs/op
func isAnagramByPrime(s1, s2 string) bool {
	n := len(s1)
	h1, h2 := 1, 1
	var ok bool 
	for i:=0; i < n; i++ {
		c1 := s1[i] | 32
		c2 := s2[i] | 32 
		ok = ok || (c1 != c2) // word is not anagram of itself
		
		h1 *= prime[c1 - 'a']
		h2 *= prime[c2 - 'a']
	}
	if ! ok {
		return false
	}

	return h1 == h2 
}

// isAnagram converts each string to a lowercase byte array, sorts it, and compares the values
//BenchmarkDetectAnagrams-4	      4381 ns/op	    1472 B/op	      71 allocs/op
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
