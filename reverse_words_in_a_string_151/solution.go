package reverse_words_in_a_string_151

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}
