package cryptosquare

import (
	"math"
	"strings"
)

// Encode returns an encoded string
func Encode(s string) string {
	s = normalize(s)

	row, col := RectDims(len(s))

	var b strings.Builder
	for i := 0; i < col; i++ {
		if i != 0 {
			b.WriteByte(' ')
		}
		for j := 0; i+j < row*col; j += col {
			if i+j < len(s) {
				b.WriteByte(s[i+j])

			} else {
				b.WriteByte(' ')
			}
		}
	}
	return b.String()
}

func normalize(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	var c byte

	for i := 0; i < len(s); i++ {
		c = s[i] | 32
		if 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
			b.WriteByte(c)
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
