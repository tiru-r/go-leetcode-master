package generate_parentheses_22

func generateParenthesis(n int) []string {
	result := make([]string, 0, 1<<n) // Catalan capacity
	var backtrack func(cur []byte, open, close int)
	backtrack = func(cur []byte, open, close int) {
		if len(cur) == 2*n {
			result = append(result, string(cur))
			return
		}
		if open < n {
			backtrack(append(cur, '('), open+1, close)
		}
		if close < open {
			backtrack(append(cur, ')'), open, close+1)
		}
	}
	backtrack(make([]byte, 0, 2*n), 0, 0)
	return result
}
