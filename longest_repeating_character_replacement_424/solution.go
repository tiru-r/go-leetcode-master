package longest_repeating_character_replacement_424

func characterReplacement(s string, k int) int {
	// Use byte keys instead of string keys for better performance
	counts := make(map[byte]int)
	var start, maxLen, maxCnt int
	for end := 0; end < len(s); end++ {
		// Direct byte operations avoid string allocation
		counts[s[end]]++

		// max is the largest count of a single, unique character in the window
		maxCnt = max(maxCnt, counts[s[end]])

		// if the amount left to modify is greater than the k
		// modifications we can use, then shrink the window.
		windowSize := end - start + 1
		leftToModify := windowSize - maxCnt
		if leftToModify > k {
			counts[s[start]]--
			start++
		}

		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}

