
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode returns an encoded string
func Encode(s string) string {
	s = fix(s)
	row, col := RectDims(len(s))
	for len(s) < row*col {
		s += " "
	}
	// s += strings.Repeat(" ", row*col-len(s)) // pad s with spaces
	var b strings.Builder
	for i := 0; i < col; i++ {
		if i != 0 {
			b.WriteByte(' ')
		}
		for j := 0; i+j < row*col; j += col {
			if i+j >= len(s) {
				b.WriteByte(' ')
			} else {
				b.WriteByte(s[i+j])
			}
		}
	}
	return b.String()
}

func fix(s string) string {
	var b strings.Builder

	for _, c := range s {
		c = c | 32
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			b.WriteRune(c)
		}
	}

	return b.String()
}

// RectDims returns dimensions for a properly-sized rectangle, 0<= row-col <= 1
func RectDims(textLen int) (row, col int) {
	row = sqrt(textLen)
	col = row
	if row*col < textLen {
		col++
	}

	if row*col < textLen {
		row++
	}

	return
}

func sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}
