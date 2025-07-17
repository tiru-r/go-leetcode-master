package word_break_139

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	wordSet := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := range len(s) {
		if !dp[i] {
			continue
		}

		for j := i + 1; j <= len(s); j++ {
			if _, exists := wordSet[s[i:j]]; exists {
				dp[j] = true
				if j == len(s) {
					return true
				}
			}
		}
	}

	return dp[len(s)]
}
