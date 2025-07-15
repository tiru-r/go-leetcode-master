package palindrome_partitioning_ii_132

import (
	"testing"
)

func TestMinCut(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "single character",
			input:    "a",
			expected: 0,
		},
		{
			name:     "two same characters",
			input:    "aa",
			expected: 0,
		},
		{
			name:     "two different characters",
			input:    "ab",
			expected: 1,
		},
		{
			name:     "palindrome",
			input:    "racecar",
			expected: 0,
		},
		{
			name:     "repeated characters",
			input:    "aaa",
			expected: 0,
		},
		{
			name:     "non-palindrome needing cuts",
			input:    "aab",
			expected: 1,
		},
		{
			name:     "complex string",
			input:    "aabbc",
			expected: 2,
		},
		{
			name:     "longer string",
			input:    "abacdc",
			expected: 1,
		},
		{
			name:     "all unique characters",
			input:    "abcd",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minCut(tt.input)
			if result != tt.expected {
				t.Errorf("minCut(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
