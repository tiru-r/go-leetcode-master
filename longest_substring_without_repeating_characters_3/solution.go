package longest_substring_without_repeating_characters_3

func lengthOfLongestSubstring(s string) int {
	// Fast-path for empty or single-char strings
	if len(s) == 0 {
		return 0
	}

	// rune gives correct handling of Unicode code points
	// (but len(s) could still be used for ASCII-only optimizations)
	lastPos := make(map[rune]int) // zero value is enough; no need to pre-size
	left, maxLen := 0, 0

	for right, ch := range s { // range over string yields runes already
		if pos, ok := lastPos[ch]; ok && pos >= left {
			left = pos + 1 // jump the window’s left edge past the duplicate
		}
		lastPos[ch] = right
		if length := right - left + 1; length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}
