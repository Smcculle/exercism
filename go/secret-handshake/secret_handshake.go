package secret

import (
	"fmt"
	"strconv"
)

var codes = []string{"jump", "close your eyes", "double blink", "wink"}

// Handshake loops through a binary string representation of code and builds a handshake from the given codes
func Handshake(code uint) (handshake []string) {

	mask := toBinString(code)

	if shouldReverse(mask) {
		handshake = loop(mask, len(mask)-1, 0, -1)
	} else {
		handshake = loop(mask, 1, len(mask), 1)
	}

	return
}


func loop(flags string, start, stop, step int) []string {
	var h []string
	for i := start; i != stop; i += step {
		if flags[i] == 49 {
			h = append(h, codes[i-1])
		}
	}

	return h
}

func toBinString(i uint) string {
	return fmt.Sprintf("%05s", strconv.FormatInt(int64(i), 2))
}

func shouldReverse(s string) bool {
	if s[0] == 48 {
		return true
	}
	return false
}
