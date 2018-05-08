package brackets

type stack []byte

func (s *stack) Push(b byte) {
	*s = append(*s, b)
}

func (s *stack) Pop() (b byte) {
	lastIndex := len(*s) - 1
	b = (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return
}

func isOpenBracket(b byte) bool {
	return b == '(' || b == '[' || b == '{'
}

func isClosedBracket(b byte) bool {
	return b == ')' || b == ']' || b == '}'
}

func closedToOpen(b byte) byte {
	if b == ')' {
		return b - 1
	}
	return b - 2
}

// Bracket matches each type using a stack
func Bracket(s string) (bool, error) {
	var bstack stack
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isOpenBracket(c) {
			bstack.Push(c)
		} else if isClosedBracket(c) {
			if len(bstack) == 0 || bstack.Pop() != closedToOpen(c) {
				return false, nil
			}
		}
	}
	return len(bstack) == 0, nil
}
