package palindrome_number_9

// Ultra-modern O(log n) solution - reverse only half the number for optimal performance
func isPalindrome(x int) bool {
	// Negative numbers and numbers ending in 0 (except 0 itself) cannot be palindromes
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversedHalf := 0
	for x > reversedHalf {
		reversedHalf = reversedHalf*10 + x%10
		x /= 10
	}

	// For even length: x == reversedHalf
	// For odd length: x == reversedHalf/10 (remove middle digit)
	return x == reversedHalf || x == reversedHalf/10
}

