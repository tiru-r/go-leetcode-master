package reverse_words_in_a_string_151

import (
	"slices"
	"strings"
)

// Modern solution using strings.Fields and slices.Reverse
func reverseWords(s string) string {
	// Use strings.Fields to split on any whitespace and remove empty strings
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}

// Alternative using strings.Builder for efficiency
func reverseWordsBuilder(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	
	var builder strings.Builder
	builder.Grow(len(s)) // Pre-allocate capacity
	
	for i, word := range words {
		if i > 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	
	return builder.String()
}

// Legacy implementation showing manual parsing (kept for comparison)
func reverseWordsArrayReversal(s string) string {
	s = strings.TrimSpace(s) // Use TrimSpace instead of Trim with " "

	tokens := make([]string, 0)
	tokens = slices.Grow(tokens, 16)
	start := 0
	
	for finish := 0; finish < len(s); finish++ {
		if s[finish] == ' ' { // Compare byte directly
			if start < finish { // Only add non-empty tokens
				tokens = append(tokens, s[start:finish])
			}
			
			// Skip multiple spaces
			for finish < len(s) && s[finish] == ' ' {
				finish++
			}
			start = finish
			finish-- // Compensate for loop increment
		}
	}
	
	// Add final token
	if start < len(s) {
		tokens = append(tokens, s[start:])
	}

	// Use slices.Reverse instead of manual reversal
	slices.Reverse(tokens)

	return strings.Join(tokens, " ")
}

// Manual reverse function (kept for educational purposes)
// Modern code should use slices.Reverse instead
func reverse(a []string) {
	slices.Reverse(a) // Use built-in function
}

// Original manual implementation (for comparison)
func reverseManual(a []string) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
