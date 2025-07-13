package first_unique_character_in_a_string_387

// firstUniqChar finds the first unique character's index in string
// Optimized: O(n) time, O(1) space using fixed-size array for lowercase letters
func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}
	
	// Fixed array for 26 lowercase letters - faster than map
	var freq [26]int
	
	// Count character frequencies
	for i := range len(s) {
		freq[s[i]-'a']++
	}
	
	// Find first character with frequency 1
	for i := range len(s) {
		if freq[s[i]-'a'] == 1 {
			return i
		}
	}
	
	return -1
}
