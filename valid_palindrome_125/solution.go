package valid_palindrome_125

// Ultra-efficient O(1) space palindrome check with on-the-fly conversion
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Skip non-alphanumeric from left
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}

		// Skip non-alphanumeric from right
		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		// Compare characters with on-the-fly lowercase conversion
		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

// Optimized alphanumeric check using direct byte comparison
func isAlphaNumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

// Ultra-fast lowercase conversion without string allocation
func toLowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
