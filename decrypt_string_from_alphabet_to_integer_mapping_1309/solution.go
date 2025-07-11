package decrypt_string_from_alphabet_to_integer_mapping_1309

import (
	"strconv"
	"strings"
)

// Modern solution using strings.Builder for efficient string building
func freqAlphabets(s string) string {
	a := "abcdefghijklmnopqrstuvwxyz"
	var builder strings.Builder
	builder.Grow(len(s)) // Pre-allocate capacity

	i := 0
	for i < len(s) {
		if i < len(s)-2 && s[i+2] == '#' {
			alphaIdx, err := strconv.Atoi(s[i : i+2])
			if err != nil || alphaIdx < 1 || alphaIdx > 26 {
				break // Invalid input
			}
			builder.WriteByte(a[alphaIdx-1])
			i += 3
		} else {
			alphaIdx, err := strconv.Atoi(string(s[i]))
			if err != nil || alphaIdx < 1 || alphaIdx > 9 {
				break // Invalid input
			}
			builder.WriteByte(a[alphaIdx-1])
			i += 1
		}
	}

	return builder.String()
}
