package intersection_of_two_arrays_ii_350

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func normalize(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	sort.Ints(b)
	return b
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name    string
		nums1   []int
		nums2   []int
		wantAny []int // order-independent
	}{
		// 1. LeetCode official
		{"LC ex1", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{"LC ex2", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},

		// 2. Empty slices
		{"empty nums1", []int{}, []int{1, 2}, []int{}},
		{"empty nums2", []int{1, 2}, []int{}, []int{}},
		{"both empty", []int{}, []int{}, []int{}},

		// 3. No overlap
		{"no overlap", []int{1, 2}, []int{3, 4}, []int{}},

		// 4. Single element
		{"single match", []int{5}, []int{5}, []int{5}},
		{"single no match", []int{5}, []int{6}, []int{}},

		// 5. Duplicates
		{"duplicates many", []int{2, 2, 2}, []int{2, 2}, []int{2, 2}},

		// 6. Large values / large counts
		{"large nums", []int{1000, 1000, 999}, []int{1000, 999, 999}, []int{999, 1000}},

		// 7. Large input (performance smoke)
		{"large 1e4", make([]int, 10000), make([]int, 10000), make([]int, 10000)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := intersect(tc.nums1, tc.nums2)
			assert.ElementsMatch(t, normalize(tc.wantAny), normalize(got))
		})
	}
}
