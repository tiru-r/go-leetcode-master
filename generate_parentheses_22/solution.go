package generate_parentheses_22

func generateParenthesis(n int) []string {
	result := make([]string, 0, n*n)
	backtrack(&result, make([]byte, 0, n*2), 0, 0, n)
	return result
}

// Modern backtracking with proper pruning - O(4^n / √n) optimal time complexity
func backtrack(result *[]string, current []byte, open, close, n int) {
	// Base case: we've built a complete valid combination
	if len(current) == n*2 {
		*result = append(*result, string(current))
		return
	}

	// Add opening parenthesis if we haven't reached the limit
	if open < n {
		current = append(current, '(')
		backtrack(result, current, open+1, close, n)
		current = current[:len(current)-1] // backtrack
	}

	// Add closing parenthesis if it would keep the string valid
	if close < open {
		current = append(current, ')')
		backtrack(result, current, open, close+1, n)
		current = current[:len(current)-1] // backtrack
	}
}

