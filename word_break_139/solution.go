package word_break_139

// No additional imports needed

// Ultra-efficient O(n²) DP solution with modern Go patterns
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	// Optimized set creation with exact capacity
	wordSet := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	// DP array: dp[i] = true if s[0:i] can be segmented
	dp := make([]bool, len(s)+1)
	dp[0] = true

	// Modern range-over-int with early termination optimization
	for i := range len(s) {
		if !dp[i] {
			continue // Skip if current position is unreachable
		}

		// Check all possible word endings from current position
		for j := i + 1; j <= len(s); j++ {
			if _, exists := wordSet[s[i:j]]; exists {
				dp[j] = true
				if j == len(s) {
					return true // Early termination when we reach the end
				}
			}
		}
	}

	return dp[len(s)]
}
