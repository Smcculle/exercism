// Package reverse handles reversing strings
package reverse

import "bytes"

// Reverse the given string
func String(s string) string {

	var buffer bytes.Buffer
	for i := len(s) - 1; i >= 0; i-- {
		buffer.WriteByte(s[i])
	}

	return buffer.String()
}
