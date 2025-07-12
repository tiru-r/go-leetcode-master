package two_sum_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------- twoSum (unsorted) ----------

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"classic", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"duplicate pair", []int{3, 3}, 6, []int{0, 1}},
		{"duplicate values", []int{2, 2, 11, 15}, 4, []int{0, 1}},
		{"negative numbers", []int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
		{"endpoints", []int{2, 3, 11, 15}, 17, []int{0, 3}},
		{"no solution", []int{1, 2, 3}, 7, []int{}},
		{"single element", []int{5}, 10, []int{}},
		{"empty", []int{}, 0, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, twoSum(tt.nums, tt.target))
		})
	}
}

// ---------- TwoSumSorted (pre-sorted) ----------

func Test_TwoSumSorted(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		i, j   int
	}{
		{"classic", []int{2, 7, 11, 15}, 9, 0, 1},
		{"duplicate pair", []int{3, 3}, 6, 0, 1},
		{"negative numbers", []int{-5, -3, -1, 2, 4, 11}, 10, 2, 5},
		{"no solution", []int{1, 2, 3, 4}, 10, -1, -1},
		{"single element", []int{5}, 10, -1, -1},
		{"empty", []int{}, 0, -1, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, gotJ := twoSumSorted(tt.nums, tt.target)
			assert.Equal(t, tt.i, gotI)
			assert.Equal(t, tt.j, gotJ)
		})
	}
}
