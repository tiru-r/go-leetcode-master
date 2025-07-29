package minimum_window_substring_76

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	// Count required characters (ASCII only – 128 slots)
	var need [128]int
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// Count unique characters needed
	required := 0
	for _, cnt := range need {
		if cnt > 0 {
			required++
		}
	}

	// Sliding-window state
	var have [128]int
	formed := 0 // Number of unique chars matching their required counts

	left, right := 0, 0
	minLen := len(s) + 1
	minStart := 0

	for right < len(s) {
		// Include s[right] in the window
		rc := s[right]
		have[rc]++
		if need[rc] > 0 && have[rc] == need[rc] {
			formed++
		}

		// Shrink from the left while the window is valid
		for left <= right && formed == required {
			winLen := right - left + 1
			if winLen < minLen {
				minLen = winLen
				minStart = left
			}

			lc := s[left]
			have[lc]--
			if need[lc] > 0 && have[lc] < need[lc] {
				formed--
			}
			left++
		}
		right++
	}

	if minLen == len(s)+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}
