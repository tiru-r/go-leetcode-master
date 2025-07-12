package longest_palindromic_substring_5

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	// helper expands around left, right and returns length
	expand := func(l, r int) int {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return r - l - 1 // length after final mismatch
	}

	for i := 0; i < len(s); i++ {
		if len1 := expand(i, i); len1 > maxLen {
			maxLen = len1
			start = i - (len1-1)/2
		}
		if len2 := expand(i, i+1); len2 > maxLen {
			maxLen = len2
			start = i - (len2-1)/2
		}
	}
	return s[start : start+maxLen]
}
