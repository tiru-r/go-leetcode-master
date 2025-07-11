package word_break_139

// No additional imports needed

// Note: study again. DP.
func wordBreak(s string, wordDict []string) bool {
	// create a set of words for fast lookup
	wordDictSet := make(map[string]struct{}, len(wordDict))
	for _, w := range wordDict {
		wordDictSet[w] = struct{}{}
	}

	// create a dp array which stores whether the i-th
	dp := make([]bool, len(s)+1)
	dp[0] = true

	// i represents the length of the current substring bring considered
	for i := 1; i <= len(s); i++ {
		// j represents the location of the partition in s[0:i]
		for j := 0; j < i; j++ {

			// if the partition up to j is in the wordDict
			// and the s[j:i] is in the wordDict, then set
			// dp[i] = true.
			partition1 := dp[j]
			partition2 := s[j:i]
			if partition1 {
				if _, exists := wordDictSet[partition2]; exists {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[len(dp)-1]
}

// Recursive solution, which exceeds the time limit.
func wordBreak0(s string, wordDict []string) bool {
	// create a set of words in wordDict for fast lookup using modern map operations
	wordDictSet := make(map[string]struct{}, len(wordDict))
	for _, w := range wordDict {
		wordDictSet[w] = struct{}{}
	}

	return wordBreakHelper(s, wordDictSet, 0)
}

func wordBreakHelper(s string, wordDictSet map[string]struct{}, start int) bool {
	// if we've reached the start of s, the string has been broken up
	if start == len(s) {
		return true
	}

	// for each character in s
	for i := start; i <= len(s); i++ {
		if _, exists := wordDictSet[s[start:i]]; exists && wordBreakHelper(s, wordDictSet, i) {
			return true
		}
	}

	return false
}
