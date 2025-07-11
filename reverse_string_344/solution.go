package reverse_string_344

import "slices"

func reverseString(s []byte) {
	slices.Reverse(s)
}

// Alternative manual implementation
func reverseStringManual(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
