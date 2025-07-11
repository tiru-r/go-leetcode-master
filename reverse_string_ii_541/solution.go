package reverse_string_ii_541

import "slices"

func reverseStr(s string, k int) string {
	r := []rune(s)
	for i := 0; i < len(s); i += 2 * k {
		// get the index to the end of k chars starting at i
		// pick the minimum for the case where the next k
		// characters falls off the end of s.
		j := min(i+k-1, len(s)-1)

		// Use modern slices.Reverse for efficient in-place reversal
		slices.Reverse(r[i:j+1])
	}

	return string(r)
}

// Note: original solution from mock interview.
// I learned how to reverse string in place by using []rune or []byte.
// Having to reverse by string concatenation make this problem tough.
// Also, the problem description was vague in cases with characters
// remaining to sort.
func reverseStrOrig(s string, k int) string {
	if len(s) < k {
		return s
	}

	// Convert to []rune for efficient manipulation
	r := []rune(s)
	
	for i := 0; i < len(r); i += 2 * k {
		end := min(i+k, len(r))
		// Use slices.Reverse for O(n) performance instead of string concatenation
		slices.Reverse(r[i:end])
	}
	
	s = string(r)

	return s
}
