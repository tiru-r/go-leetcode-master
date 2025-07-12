package valid_parentheses_20

// No imports needed - using Go 1.24 built-in slice operations

// Modern solution using Go 1.24 slice operations instead of custom stack
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	// Pre-allocate stack with capacity to avoid reallocations
	stack := make([]byte, 0, len(s)/2)
	matches := map[byte]byte{')': '(', '}': '{', ']': '['}

	for i := range len(s) {
		c := s[i]

		switch c {
		case '(', '{', '[':
			// Push: append to slice
			stack = append(stack, c)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			// Pop: get last element and shrink slice
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top != matches[c] {
				return false
			}
		}
	}

	return len(stack) == 0
}
