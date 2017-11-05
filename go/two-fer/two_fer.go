// Package twofer gives one to you and one to me.

package twofer

import (
	"fmt"
	"strings"
)

// ShareWith lets you know who to share with
func ShareWith(s string) string {

	if strings.TrimSpace(s) == "" {
		s = "you"
	}

	return fmt.Sprintf("One for %v, one for me.", s)
}
