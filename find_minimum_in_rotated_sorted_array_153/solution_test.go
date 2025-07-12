package find_minimum_in_rotated_sorted_array_153

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"rotated mid", []int{3, 4, 5, 1, 2}, 1},
		{"rotated tail", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"no rotation", []int{1, 2, 3}, 1},
		{"two elements", []int{1, 0}, 0},
		{"single element", []int{42}, 42},
		{"all identical", []int{5, 5, 5, 5}, 5},
		{"rotated duplicates", []int{2, 2, 2, 0, 1, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findMin(tt.nums))
		})
	}
}
