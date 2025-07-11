package valid_anagram_242

import "maps"

// Modern solution using maps.Equal for efficient comparison
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sFreq := make(map[rune]int)
	for _, r := range s {
		sFreq[r]++
	}

	tFreq := make(map[rune]int)
	for _, r := range t {
		tFreq[r]++
	}

	// Use maps.Equal for efficient comparison
	return maps.Equal(sFreq, tFreq)
}

func isAnagram0(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// build a character to frequency map
	freq := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		_, ok := freq[s[i]]
		if !ok {
			freq[s[i]] = 1
		} else {
			freq[s[i]] += 1
		}
	}

	// look over anagram in question and decrement frequencies
	for i := 0; i < len(s); i++ {
		f, ok := freq[t[i]]
		if !ok {
			return false
		}

		// if frequency drops below 0, return false early
		if f-1 < 0 {
			return false
		}

		freq[t[i]] -= 1
	}

	// verify that all frequencies are 0 for a valid anagram
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}

	return true
}
