package valid_parentheses_20

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]byte, 0, len(s)/2)
	matches := map[byte]byte{')': '(', '}': '{', ']': '['}

	for i := range len(s) {
		c := s[i]

		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != matches[c] {
				return false
			}
		}
	}

	return len(stack) == 0
}
