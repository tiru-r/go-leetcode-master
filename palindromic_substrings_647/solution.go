package palindromic_substrings_647

func countSubstrings(s string) int {
	count := 0
	n := len(s)

	for i := range n {
		count += expandAroundCenter(s, i, i, n)
		count += expandAroundCenter(s, i, i+1, n)
	}

	return count
}

func expandAroundCenter(s string, left, right, n int) int {
	count := 0
	for left >= 0 && right < n && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}
