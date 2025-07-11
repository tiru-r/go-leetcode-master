package longest_common_prefix_14

import (
	"cmp"
	"slices"
	"strings"
)

// Modern solution using slices operations and efficient string building
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// Find the shortest string length for bounds checking using cmp.Compare
	minLen := slices.MinFunc(strs, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	})
	if len(minLen) == 0 {
		return ""
	}

	// Use strings.Builder for efficient string construction
	var prefix strings.Builder
	prefix.Grow(len(minLen)) // Pre-allocate capacity

	// Check each position up to the shortest string length
	for i := 0; i < len(minLen); i++ {
		char := strs[0][i]
		// Check if all strings have the same character at position i
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != char {
				return prefix.String()
			}
		}
		prefix.WriteByte(char)
	}

	return prefix.String()
}

// Alternative implementation using strings.Cut for binary search approach
func longestCommonPrefixBinary(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := len(slices.MinFunc(strs, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}))

	low, high := 0, minLen
	for low < high {
		mid := (low + high + 1) / 2
		if isCommonPrefix(strs, mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return strs[0][:low]
}

func isCommonPrefix(strs []string, length int) bool {
	prefix := strs[0][:length]
	for i := 1; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], prefix) {
			return false
		}
	}
	return true
}
