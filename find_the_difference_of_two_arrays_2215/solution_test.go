package find_the_difference_of_two_arrays_2215

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// sort2D sorts every inner slice so we can compare order-independently.
func sort2D(a [][]int) {
	for _, s := range a {
		sort.Ints(s)
	}
}

func TestFindDifference(t *testing.T) {
	tests := []struct {
		name         string
		nums1, nums2 []int
		want         [][]int
	}{
		{"empty inputs", []int{}, []int{}, [][]int{{}, {}}},
		{"identical arrays", []int{1, 2, 3}, []int{1, 2, 3}, [][]int{{}, {}}},
		{"no overlap", []int{1, 2}, []int{3, 4}, [][]int{{1, 2}, {3, 4}}},
		{"duplicates removed", []int{1, 1, 2}, []int{2, 2, 3}, [][]int{{1}, {3}}},
		{"mixed order", []int{3, 1, 2}, []int{2, 4, 3}, [][]int{{1}, {4}}},
		{"negative numbers", []int{-1, 0, 1}, []int{1, 2, 3}, [][]int{{-1, 0}, {2, 3}}},
		{"large n", []int{1, 2, 3, 4, 5}, []int{4, 5, 6, 7, 8}, [][]int{{1, 2, 3}, {6, 7, 8}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifference(tt.nums1, tt.nums2)
			sort2D(got)
			sort2D(tt.want)
			assert.Equal(t, tt.want, got)
		})
	}
}
