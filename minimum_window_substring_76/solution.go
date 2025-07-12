package minimum_window_substring_76

// minWindow returns the minimum substring of s that contains every
// character of t (including duplicates).  If no such window exists,
// the empty string is returned.
func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// 1. frequency table for t
	tFreq := make([]int, 128)
	for _, r := range t {
		tFreq[r]++
	}
	required := 0
	for _, v := range tFreq {
		if v > 0 {
			required++
		}
	}

	// 2. sliding-window state
	win := make([]int, 128)
	formed := 0
	left := 0
	minLen := len(s) + 1
	minStart := 0

	for right, r := range s {
		// expand window
		win[r]++
		if win[r] == tFreq[r] {
			formed++
		}

		// contract while valid
		for formed == required && left <= right {
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
			}
			leftR := s[left]
			win[leftR]--
			if win[leftR] == tFreq[leftR]-1 {
				formed--
			}
			left++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}
