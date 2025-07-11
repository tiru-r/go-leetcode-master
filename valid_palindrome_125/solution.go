package valid_palindrome_125

import "strings"

func isPalindrome(s string) bool {
	front := 0
	back := len(s) - 1
	s = strings.ToLower(s)

	for front <= back {
		// Use direct byte operations for better performance
		for front < len(s) && !isAlphaNumeric(s[front]) {
			front++
		}

		for back > -1 && !isAlphaNumeric(s[back]) {
			back--
		}

		// "a:a" case needs initial condition here
		if front <= back && s[front] != s[back] {
			return false
		}

		front++
		back--
	}

	return true
}

// Optimized function using direct byte comparison instead of regex
func isAlphaNumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}
