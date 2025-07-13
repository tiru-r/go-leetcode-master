package wildcard_matching_44

// IsMatch checks if string s matches pattern p with wildcards
// Time: O(m*n), Space: O(n) - optimized space using rolling array
func IsMatch(s string, p string) bool {
	m, n := len(s), len(p)
	
	// Optimize pattern by removing consecutive *
	pattern := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		if p[i] == '*' && len(pattern) > 0 && pattern[len(pattern)-1] == '*' {
			continue
		}
		pattern = append(pattern, p[i])
	}
	n = len(pattern)
	
	// Use two arrays for space optimization
	prev := make([]bool, n+1)
	curr := make([]bool, n+1)
	
	// Base case: empty string matches empty pattern
	prev[0] = true
	
	// Handle patterns that start with *
	for j := 1; j <= n; j++ {
		prev[j] = prev[j-1] && pattern[j-1] == '*'
	}
	
	// Fill DP table
	for i := 1; i <= m; i++ {
		curr[0] = false // Non-empty string cannot match empty pattern
		
		for j := 1; j <= n; j++ {
			if pattern[j-1] == '*' {
				// * can match empty sequence or any character
				curr[j] = curr[j-1] || prev[j]
			} else if pattern[j-1] == '?' || pattern[j-1] == s[i-1] {
				// ? matches any character, or exact character match
				curr[j] = prev[j-1]
			} else {
				// No match
				curr[j] = false
			}
		}
		
		// Swap arrays for next iteration
		prev, curr = curr, prev
	}
	
	return prev[n]
}