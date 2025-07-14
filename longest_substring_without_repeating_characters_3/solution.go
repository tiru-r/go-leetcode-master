package longest_substring_without_repeating_characters_3

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// Use byte array for ASCII optimization (common case)
	if isASCII(s) {
		return lengthOfLongestSubstringASCII(s)
	}

	// Fallback to Unicode-safe implementation
	return lengthOfLongestSubstringUnicode(s)
}

func isASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] >= 128 {
			return false
		}
	}
	return true
}

func lengthOfLongestSubstringASCII(s string) int {
	lastPos := [128]int{}
	for i := range lastPos {
		lastPos[i] = -1
	}
	
	left, maxLen := 0, 0
	
	for right := 0; right < len(s); right++ {
		char := s[right]
		if lastPos[char] >= left {
			left = lastPos[char] + 1
		}
		lastPos[char] = right
		maxLen = max(maxLen, right-left+1)
	}
	
	return maxLen
}

func lengthOfLongestSubstringUnicode(s string) int {
	lastPos := make(map[rune]int)
	left, maxLen := 0, 0
	
	runes := []rune(s)
	for right, char := range runes {
		if pos, exists := lastPos[char]; exists && pos >= left {
			left = pos + 1
		}
		lastPos[char] = right
		maxLen = max(maxLen, right-left+1)
	}
	
	return maxLen
}
