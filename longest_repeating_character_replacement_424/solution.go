package longest_repeating_character_replacement_424

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	count := [26]int{}
	left, maxCount, maxLength := 0, 0, 0

	for right := 0; right < len(s); right++ {
		// Add current character to window
		charIdx := s[right] - 'A'
		count[charIdx]++
		
		// Update max frequency in current window
		maxCount = max(maxCount, count[charIdx])
		
		// Check if window is valid (replacements needed <= k)
		windowSize := right - left + 1
		if windowSize-maxCount > k {
			// Shrink window from left
			count[s[left]-'A']--
			left++
		}
		
		// Update maximum valid window size
		maxLength = max(maxLength, right-left+1)
	}
	
	return maxLength
}
