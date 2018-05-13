package wordcount

import (
	"strings"
)

// Frequency counts the number of occurrences of each word in a phrase.
type Frequency map[string]int

func isAlphaNumeric(b byte) bool {
	return 'a' <= b && b <= 'z' || '0' <= b && b <= '9'
}

// WordCount counts the frequency of each word
func WordCount(phrase string) (f Frequency) {
	f = make(Frequency)
	var b strings.Builder
	var ch byte

	for i := 0; i < len(phrase); i++ {
		ch = phrase[i] | 32

		if isAlphaNumeric(ch) {
			b.WriteByte(ch)
		} else if ch == '\'' && b.Len() > 0 && i+1 < len(phrase) && isAlphaNumeric(phrase[i+1]) {
			b.WriteByte(ch)
		} else if b.Len() > 0 {
			f[b.String()]++
			b.Reset()
		}
	}

	if b.Len() > 0 {
		f[b.String()]++
	}

	return
}
