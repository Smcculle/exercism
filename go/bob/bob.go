// Package bob package implements the complex language of Bob in 4 simple phrases.
// Bob responds to questions, yelling, or silence, or everything else.
package bob

import (
	"strings"
)

const testVersion = 3
const qAns = "Sure."
const yAns = "Whoa, chill out!"
const sAns = "Fine. Be that way!"
const elseAns = "Whatever."

// Hey tests a remark to see if it is all capitals (yelling), ends in a question mark (question), or has no text (blank).
// Returns the constant string associated with each case or a default answer if none match.
func Hey(remark string) (resp string) {

	switch trimmed := strings.TrimSpace(remark); {

	case isYelling(trimmed):
		resp = yAns
	case isQuestion(trimmed):
		resp = qAns
	case isSilence(trimmed):
		resp = sAns
	default:
		resp = elseAns
	}
	return
}
func isSilence(remark string) bool {
	return len(remark) == 0
}
func isQuestion(remark string) bool {
	return remark != "" && string(remark[len(remark)-1]) == "?"
}

func isYelling(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
}
