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
		slices.Reverse(r[i : j+1])
	}

	return string(r)
}
