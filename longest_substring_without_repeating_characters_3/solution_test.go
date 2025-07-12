package longest_substring_without_repeating_characters_3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty string", "", 0},
		{"single char", "a", 1},
		{"all unique", "abcdef", 6},
		{"all same", "bbbbb", 1},
		{"example 1", "abcabcbb", 3},
		{"example 2", "pwwkew", 3},
		{"mixed", "adcab", 4},
		{"with spaces", "a b c a b", 3},
		{"unicode safe", "üöäüö", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, lengthOfLongestSubstring(tt.s))
		})
	}
}
