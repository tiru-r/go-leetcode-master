package longest_substring_without_repeating_characters_3

func lengthOfLongestSubstring(s string) int {
	var (
		last [128]int // ASCII fast-path
		m    map[rune]int
		useM bool

		left, maxLen int
	)

	// Initialize sentinel positions for ASCII array
	for i := range last {
		last[i] = -1
	}

	for right, ch := range s {
		if ch < 128 && !useM {
			// ASCII fast path using array lookup
			if pos := last[ch]; pos >= left {
				left = pos + 1
			}
			last[ch] = right
		} else {
			// Switch to map for Unicode support
			if !useM {
				// Estimate capacity based on remaining string length
				capacity := min(len(s)/4, 64) // reasonable default
				m = make(map[rune]int, capacity)
				// Migrate ASCII data to map
				for i, pos := range last {
					if pos != -1 {
						m[rune(i)] = pos
					}
				}
				useM = true
			}
			if pos, ok := m[ch]; ok && pos >= left {
				left = pos + 1
			}
			m[ch] = right
		}

		// Update maximum length using modern Go function
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}
