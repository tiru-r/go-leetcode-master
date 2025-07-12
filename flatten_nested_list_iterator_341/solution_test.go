package flatten_nested_list_iterator_341

import (
	"reflect"
	"testing"
)

/* ---------- test table ---------- */

func TestNestedIterator(t *testing.T) {
	tests := []struct {
		name string
		list []*NestedInteger
		want []int
	}{
		// 1. Empty / single element
		{"empty", []*NestedInteger{}, []int{}},
		{"single_int", []*NestedInteger{Int(8)}, []int{8}},

		// 2. Flat list
		{"flat_3", []*NestedInteger{Int(1), Int(2), Int(3)}, []int{1, 2, 3}},

		// 3. Single-level nesting
		{"nested_once", []*NestedInteger{List(Int(1), Int(2)), Int(3)}, []int{1, 2, 3}},
		{"nested_once_empty", []*NestedInteger{List(), Int(4)}, []int{4}},

		// 4. Deeply nested
		{"deep_3_levels", []*NestedInteger{List(List(List(Int(42))))}, []int{42}},
		{"deep_empty", []*NestedInteger{List(List(List()), List())}, []int{}},

		// 5. Mixed nesting
		{"mixed", []*NestedInteger{
			List(Int(1), List(Int(2), List(Int(3), Int(4))), Int(5)),
			Int(6),
			List(),
			List(Int(7), Int(8)),
		}, []int{1, 2, 3, 4, 5, 6, 7, 8}},

		// 6. Large flat list (LeetCode performance)
		{"large_flat", func() []*NestedInteger {
			out := make([]*NestedInteger, 2000)
			for i := 0; i < 2000; i++ {
				out[i] = Int(i + 1)
			}
			return out
		}(), makeRange(1, 2000)},

		// 7. Alternating nest / int pattern
		{"alternating", []*NestedInteger{
			List(Int(1)),
			List(List(Int(2))),
			List(List(List(Int(3)))),
			Int(4),
		}, []int{1, 2, 3, 4}},

		// 8. All empty lists
		{"all_empty", []*NestedInteger{List(), List(List()), List(List(List()))}, []int{}},

		// 9. Negative numbers
		{"negatives", []*NestedInteger{List(Int(-1), List(Int(-2))), Int(-3)}, []int{-1, -2, -3}},

		// 10. Zero values
		{"zeros", []*NestedInteger{Int(0), List(Int(0), List(Int(0)))}, []int{0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			it := Constructor(tt.list)
			got := drain(it)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

/* ---------- helper: drain iterator into slice ---------- */

func drain(it *NestedIterator) []int {
	var out []int
	for it.HasNext() {
		out = append(out, it.Next())
	}
	return out
}

/* ---------- helper: generate range for large tests ---------- */

func makeRange(lo, hi int) []int {
	r := make([]int, hi-lo+1)
	for i := range r {
		r[i] = lo + i
	}
	return r
}
