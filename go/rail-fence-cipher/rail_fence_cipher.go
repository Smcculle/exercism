package railfence

import (
	"sort"
	"strings"
)

// zipped represents a character and the rail it's on
type zipped struct {
	b    byte
	rail int
}

// zipSlice represents our rail fence, with each character belonging to a certain rail
type zipSlice []zipped

// SortRail sorts a slice of zipped bytes according its rail position
func (zs *zipSlice) SortRail() {

	sort.SliceStable(*zs, func(i, j int) bool {
		return (*zs)[i].rail < (*zs)[j].rail
	})
}

// Takes a string and a fence generator and combines them to form a rail fence.
// next is a cyclic function from {1..n..1} for n rails.
func (zs *zipSlice) Zip(s string, next func() int) {
	for i := 0; i < len(*zs); i++ {
		(*zs)[i] = zipped{s[i], next()}
	}
}

// Zips the position of a string instead of the character itself
func (zs *zipSlice) ZipIndex(next func() int) {
	for i := 0; i < len(*zs); i++ {
		(*zs)[i] = zipped{byte(i), next()}
	}
}

// String returns a string from the underlying byte array.
func (zs zipSlice) String() string {
	var b strings.Builder
	for _, z := range zs {
		b.WriteByte(z.b)
	}
	return b.String()
}

// GenEncoding returns a cyclic function that endlessly generates values
// from 1 to n and back down to 1, which will encode our string when sorted.
func GenEncoding(limit int) func() int {

	var current int
	var reverse bool

	return func() int {

		if !reverse {
			current++
			if current == limit {
				reverse = true
			}
		} else {
			current--
			if current == 1 {
				reverse = false
			}
		}
		return current
	}
}

// GenDecoding uses rail encode on positions {1..n} for an encoded string of length n.
// Returns a generator that outputs the encoded positions one at a time which
// will return the original string when sorted.
func GenDecoding(length, numRails int) func() int {
	z := make(zipSlice, length)
	z.ZipIndex(GenEncoding(numRails))
	z.SortRail()
	i := -1
	return func() int {
		i++
		return int(z[i].b)
	}
}

// Encode implements the rail cipher with the given number of rails
// BenchmarkEncode-4   5077 ns/op	    1416 B/op	      29 allocs/op
func Encode(s string, numRails int) string {
	z := make(zipSlice, len(s))
	z.Zip(s, GenEncoding(numRails))
	z.SortRail()
	return z.String()
}

// Decode reverses the rail cipher
// BenchmarkDecode-4  27272 ns/op	    4496 B/op	      52 allocs/op
func Decode(s string, numRails int) string {
	z := make(zipSlice, len(s))
	z.Zip(s, GenDecoding(len(s), numRails))
	z.SortRail()
	return z.String()
}
