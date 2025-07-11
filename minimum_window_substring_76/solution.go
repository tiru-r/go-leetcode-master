package minimum_window_substring_76

// No additional imports needed

// Second solution. Intuition is to open the window
// until character frequencies in t are all less than zero.
// After, close the window until any single character frequency
// of t is larger than 0.
func minWindow(s string, t string) string {
	tFreq := make(map[rune]int)
	for _, r := range t {
		tFreq[r]++
	}

	required := len(tFreq)
	formed := 0
	windowCounts := make(map[rune]int)
	
	left, right := 0, 0
	minLen := len(s) + 1
	minStart := 0

	for right < len(s) {
		char := rune(s[right])
		windowCounts[char]++

		if freq, exists := tFreq[char]; exists && windowCounts[char] == freq {
			formed++
		}

		for formed == required && left <= right {
			if right-left+1 < minLen {
				minLen = right - left + 1
				minStart = left
			}

			leftChar := rune(s[left])
			windowCounts[leftChar]--
			if freq, exists := tFreq[leftChar]; exists && windowCounts[leftChar] < freq {
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


// Initial solution. Pretty close, but overly complicated
// and assumed that duplicates of t were not allowed, meaning
// the substring of s must have exactly (and no more) the same
// frequency of unique characters in t.
func minWindow0(s string, t string) string {
	tFreq := make(map[rune]int)
	for _, r := range t {
		tFreq[r]++
	}

	left := 0
	for left < len(s) {
		if _, exists := tFreq[rune(s[left])]; exists {
			break
		}
		left++
	}

	right := left
	minWindow := ""
	windowCounts := make(map[rune]int)

	for right < len(s) {
		char := rune(s[right])
		if _, exists := tFreq[char]; exists {
			windowCounts[char]++

			for isValidWindow(tFreq, windowCounts) {
				current := s[left : right+1]
				if minWindow == "" || len(current) < len(minWindow) {
					minWindow = current
				}

				leftChar := rune(s[left])
				if _, exists := tFreq[leftChar]; exists {
					windowCounts[leftChar]--
				}
				left++
			}
		}
		right++
	}

	return minWindow
}

func isValidWindow(tFreq, windowCounts map[rune]int) bool {
	// Check if all characters have sufficient count
	for char, needed := range tFreq {
		if windowCounts[char] < needed {
			return false
		}
	}
	return true
}
