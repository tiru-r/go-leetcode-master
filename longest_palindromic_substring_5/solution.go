package longest_palindromic_substring_5

// Modern solution using built-in max function (Go 1.21+)

// expand at the center 2*n - 1 times for
// palindromes with even and odd lengths
func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		s1 := expandPalindrome(s, i, i)
		s2 := expandPalindrome(s, i, i+1)

		if len(s1) > len(longest) {
			longest = s1
		}
		if len(s2) > len(longest) {
			longest = s2
		}
	}

	return longest
}

func expandPalindrome(s string, i, j int) string {
	if i < 0 || j >= len(s) {
		return ""
	}

	// Don't expand if they're not already equal
	if s[i] != s[j] {
		return ""
	}

	// expand i and j until they're out-of-bounds
	for i > -1 && j < len(s) {
		// if expanding would make the string no longer
		// a palindrome or go out-of-bounds, break so
		// we can return the palindromic substring
		if i-1 < 0 || j+1 >= len(s) || s[i-1] != s[j+1] {
			break
		}

		i--
		j++
	}

	// return the palindromic substring
	return s[i : j+1]
}

