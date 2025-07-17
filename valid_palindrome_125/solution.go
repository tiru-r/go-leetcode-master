package valid_palindrome_125

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}

		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}

		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func toLowerCase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
