package longest_palindromic_substring_5

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	// Optimized expand function
	expand := func(l, r int) (int, int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return l + 1, r - 1
	}

	for i := 0; i < len(s); i++ {
		// Skip if remaining characters can't form a longer palindrome
		if len(s)-i <= maxLen/2 {
			break
		}

		// Check odd-length palindromes (center at i)
		if left, right := expand(i, i); right-left+1 > maxLen {
			maxLen = right - left + 1
			start = left
		}

		// Check even-length palindromes (center between i and i+1)
		if i+1 < len(s) {
			if left, right := expand(i, i+1); right-left+1 > maxLen {
				maxLen = right - left + 1
				start = left
			}
		}
	}

	return s[start : start+maxLen]
}
