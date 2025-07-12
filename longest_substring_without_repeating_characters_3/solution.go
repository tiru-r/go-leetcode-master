package longest_substring_without_repeating_characters_3

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	
	var charIndex [256]int // ASCII character set
	for i := range charIndex {
		charIndex[i] = -1
	}
	
	left, maxLen := 0, 0
	
	for right := range len(s) {
		char := s[right]
		if charIndex[char] >= left {
			left = charIndex[char] + 1
		}
		charIndex[char] = right
		maxLen = max(maxLen, right-left+1)
	}
	
	return maxLen
}
