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

// High-performance memoized recursion - O(n²) time, O(n) space
func wordBreakMemoized(s string, wordDict []string) bool {
	wordSet := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	// Memoization cache to avoid recomputation
	memo := make(map[int]bool)
	return wordBreakMemo(s, wordSet, 0, memo)
}

func wordBreakMemo(s string, wordSet map[string]struct{}, start int, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}

	// Check memoization cache
	if result, exists := memo[start]; exists {
		return result
	}

	// Try all possible word segmentations
	for end := start + 1; end <= len(s); end++ {
		word := s[start:end]
		if _, exists := wordSet[word]; exists {
			if wordBreakMemo(s, wordSet, end, memo) {
				memo[start] = true
				return true
			}
		}
	}

	memo[start] = false
	return false
}

// Inefficient O(2^n) recursion - DO NOT USE (kept for comparison)
func wordBreakInefficient(s string, wordDict []string) bool {
	wordDictSet := make(map[string]struct{}, len(wordDict))
	for _, w := range wordDict {
		wordDictSet[w] = struct{}{}
	}
	return wordBreakHelperSlow(s, wordDictSet, 0)
}

func wordBreakHelperSlow(s string, wordDictSet map[string]struct{}, start int) bool {
	if start == len(s) {
		return true
	}
	for i := start; i <= len(s); i++ {
		if _, exists := wordDictSet[s[start:i]]; exists && wordBreakHelperSlow(s, wordDictSet, i) {
			return true
		}
	}
	return false
}
