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
			alphaIdx, _ := strconv.Atoi(s[i : i+2])
			builder.WriteByte(a[alphaIdx-1])
			i += 3
		} else {
			alphaIdx, _ := strconv.Atoi(string(s[i]))
			builder.WriteByte(a[alphaIdx-1])
			i += 1
		}
	}

	return builder.String()
}
