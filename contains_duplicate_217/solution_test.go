package contains_duplicate_217

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "DuplicatesMidSlice",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "DuplicatesNonAdjacent",
			input:    []int{0, 4, 5, 0, 3, 6},
			expected: true,
		},
		{
			name:     "NoDuplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "EmptySlice",
			input:    []int{},
			expected: false,
		},
		{
			name:     "SingleElement",
			input:    []int{42},
			expected: false,
		},
		{
			name:     "DuplicatesAtStart",
			input:    []int{1, 1, 2, 3, 4},
			expected: true,
		},
		{
			name:     "DuplicatesAtEnd",
			input:    []int{1, 2, 3, 4, 4},
			expected: true,
		},
		{
			name:     "NegativeNumbersWithDuplicates",
			input:    []int{-1, 2, -1, 3, 4},
			expected: true,
		},
		{
			name:     "MultipleDuplicates",
			input:    []int{1, 1, 1, 2, 2, 3},
			expected: true,
		},
		{
			name:     "LargeSliceNoDuplicates",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, containsDuplicate(tt.input), "Failed for input: %v", tt.input)
		})
	}
}
