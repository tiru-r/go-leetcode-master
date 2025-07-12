package longest_substring_without_repeating_characters_3

func lengthOfLongestSubstring(s string) int {
	const size = 256
	last := [size]int{} // index of last occurrence; 0 means “not seen” yet
	for i := range last {
		last[i] = -1
	}

	left, best := 0, 0
	for right, ch := range []byte(s) {
		if idx := last[ch]; idx >= left {
			left = idx + 1
		}
		last[ch] = right
		if newLen := right - left + 1; newLen > best {
			best = newLen
		}
	}
	return best
}
