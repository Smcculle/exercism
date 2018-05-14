package railfence

// action specifies 0 for decode or 1 for encode
type action int

const (
	decode action = iota
	encode
)

// Encode implements the rail cipher with the given number of rails
// Benchmark-Sorting  5077 ns/op	    1416 B/op	      29 allocs/op
// Benchmark-Swap  	   188 ns/op	      80 B/op	       3 allocs/op
func Encode(s string, numRails int) string {
	return Swap(s, numRails, encode)
}

// Decode reverses the rail cipher
// BenchmarkDecode-Sorting  27272 ns/op	    4496 B/op	      52 allocs/op
// BenchmarkDecode-Swap       294 ns/op	     192 B/op	       4 allocs/op
func Decode(s string, numRails int) string {
	return Swap(s, numRails, decode)
}

// Swap creates a mutable copy of s and swaps each element according to the given action
// Action is either encode or decode.
func Swap(s string, numRails int, a action) string {
	var i, index, delta int

	// moving pointers which track and response position and string position
	var rpos, spos *int
	if a == encode {
		rpos = &index
		spos = &i
	} else { // decoding is just the opposite of encoding, swap the pointers
		rpos = &i
		spos = &index
	}

	cycle := (numRails - 1) * 2
	resp := []byte(s)
	for rail := 0; rail < numRails; rail++ {
		delta = rail * 2
		for i = rail; i < len(s); i += delta {
			resp[*rpos] = s[*spos]
			index++
			if delta == cycle {
				continue
			}
			delta = cycle - delta
		}
	}
	return string(resp)
}
