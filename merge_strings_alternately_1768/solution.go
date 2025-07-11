package merge_strings_alternately_1768

import "strings"

// Modern solution using strings.Builder for efficient string building
func mergeAlternately(word1 string, word2 string) string {
	var builder strings.Builder
	builder.Grow(len(word1) + len(word2)) // Pre-allocate capacity
	
	i, j := 0, 0
	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			builder.WriteByte(word1[i])
			i++
		}
		if j < len(word2) {
			builder.WriteByte(word2[j])
			j++
		}
	}

	return builder.String()
}
