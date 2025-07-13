package generate_parentheses_22


// generateParenthesis generates all valid parentheses combinations for n pairs
// Optimized: Uses Catalan number for capacity, pre-allocated byte slice for efficiency
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	
	// Pre-calculate Catalan number for exact capacity allocation
	capacity := catalan(n)
	result := make([]string, 0, capacity)
	
	// Optimized backtracking with pre-allocated byte slice
	path := make([]byte, 0, 2*n)
	
	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		// Base case: complete valid parentheses string
		if open == n && close == n {
			result = append(result, string(path))
			return
		}
		
		// Add opening parenthesis if we haven't used all n
		if open < n {
			path = append(path, '(')
			backtrack(open+1, close)
			path = path[:len(path)-1] // Backtrack
		}
		
		// Add closing parenthesis if it doesn't violate validity
		if close < open {
			path = append(path, ')')
			backtrack(open, close+1)
			path = path[:len(path)-1] // Backtrack
		}
	}
	
	backtrack(0, 0)
	return result
}

// catalan calculates the nth Catalan number efficiently
func catalan(n int) int {
	if n <= 1 {
		return 1
	}
	
	// Use optimized formula: C(n) = (2n)! / ((n+1)! * n!)
	// Simplified to: C(n) = (2n * (2n-1) * ... * (n+2)) / (n * (n-1) * ... * 1)
	result := 1
	for i := range n {
		result = result * (2*n - i) / (i + 1)
	}
	return result / (n + 1)
}
