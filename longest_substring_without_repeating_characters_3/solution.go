package longest_substring_without_repeating_characters_3

func lengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	left, maxLen := 0, 0

	for right := 0; right < len(s); right++ {
		if idx, found := charIndex[s[right]]; found && idx >= left {
			left = idx + 1
		}
		charIndex[s[right]] = right
		if currentLen := right - left + 1; currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}
