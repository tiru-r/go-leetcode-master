package reverse_words_in_a_string_151

import (
	"slices"
	"strings"
)

// reverseWords returns the string with word order reversed.
// Runs in O(n) time and O(n) space (output slice + Builder).
func reverseWords(s string) string {
	// strings.Fields collapses any run of Unicode whitespace.
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}

// reverseWordsBuilder avoids the intermediate slice allocation
// by using strings.Builder.  Still O(n) but ~½ the allocs.
func reverseWordsBuilder(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)

	var b strings.Builder
	b.Grow(len(s)) // worst-case exact capacity
	for i, w := range words {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(w)
	}
	return b.String()
}
