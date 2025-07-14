package longest_palindromic_substring_5

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"single char", "b", "b"},
		{"odd palindrome", "babad", "bab"},
		{"even palindrome", "cbbd", "bb"},
		{"whole string palindrome", "baaab", "baaab"},
		{"duplicates", "bb", "bb"},
		{"no palindrome longer than 1", "abcdef", "a"},
		{"mixed case", "abacdfgdcaba", "aba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Any correct answer is accepted when multiple exist
			got := longestPalindrome(tt.s)
			assert.True(t,
				got == tt.want || // exact match first
					(tt.name == "odd palindrome" && (got == "bab" || got == "aba")),
				"expected %q or equivalent palindrome, got %q", tt.want, got)
		})
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	s := "babad"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome(s)
	}
}

func BenchmarkLongestPalindromeLarge(b *testing.B) {
	s := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 100) + "racecar" + strings.Repeat("zyxwvutsrqponmlkjihgfedcba", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestPalindrome(s)
	}
}
