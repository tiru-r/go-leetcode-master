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
		{"DuplicatesMidSlice", []int{1, 2, 3, 1}, true},
		{"DuplicatesNonAdjacent", []int{0, 4, 5, 0, 3, 6}, true},
		{"NoDuplicates", []int{1, 2, 3, 4, 5}, false},
		{"EmptySlice", []int{}, false},
		{"SingleElement", []int{42}, false},
		{"DuplicatesAtStart", []int{1, 1, 2, 3, 4}, true},
		{"DuplicatesAtEnd", []int{1, 2, 3, 4, 4}, true},
		{"NegativeNumbersWithDuplicates", []int{-1, 2, -1, 3, 4}, true},
		{"MultipleDuplicates", []int{1, 1, 1, 2, 2, 3}, true},
		{"LargeSliceNoDuplicates", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, containsDuplicate(tt.input))
		})
	}
}
