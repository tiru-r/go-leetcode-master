package minimum_window_substring_76

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	tFreq := make(map[byte]int)
	for i := range t {
		tFreq[t[i]]++
	}

	required := len(tFreq)
	formed := 0
	windowCounts := make(map[byte]int)
	
	left, right := 0, 0
	minLen := len(s) + 1
	minLeft := 0

	for right < len(s) {
		char := s[right]
		windowCounts[char]++
		
		if count, exists := tFreq[char]; exists && windowCounts[char] == count {
			formed++
		}

		for left <= right && formed == required {
			if right-left+1 < minLen {
				minLen = right - left + 1
				minLeft = left
			}

			leftChar := s[left]
			windowCounts[leftChar]--
			if count, exists := tFreq[leftChar]; exists && windowCounts[leftChar] < count {
				formed--
			}
			left++
		}
		right++
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[minLeft : minLeft+minLen]
}
