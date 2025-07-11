package longest_substring_without_repeating_characters_3

func lengthOfLongestSubstring(s string) int {
	front, rear := 0, 0
	longest := 0
	set := make(map[byte]bool)
	for front < len(s) {
		if set[s[front]] {
			for s[rear] != s[front] {
				delete(set, s[rear])
				rear++
			}
			rear++
		} else {
			set[s[front]] = true
			longest = max(longest, front-rear+1)
		}

		front++
	}

	return longest
}
