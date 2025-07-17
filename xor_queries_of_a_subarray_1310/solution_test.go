package xor_queries_of_a_subarray_1310

import (
	"reflect"
	"testing"
)

func TestXorQueries(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		queries [][]int
		want    []int
	}{
		{"empty array", []int{}, [][]int{}, []int{}},
		{"single element", []int{5}, [][]int{{0, 0}}, []int{5}},
		{"duplicate numbers", []int{2, 2, 2}, [][]int{{0, 2}, {1, 1}}, []int{2, 2}},
		{"LC example 1",
			[]int{1, 3, 4, 8},
			[][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}},
			[]int{2, 7, 14, 8}},
		{"LC example 2",
			[]int{4, 8, 2, 10},
			[][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}},
			[]int{8, 0, 4, 4}},
		{"whole array", []int{6, 2, 7, 3, 8}, [][]int{{0, 4}}, []int{6 ^ 2 ^ 7 ^ 3 ^ 8}},
		{"many queries",
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}},
			[]int{1, 3, 0, 4, 1, 7, 0, 8}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := xorQueries(tc.arr, tc.queries)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// Benchmark with 3e4 elements, 3e4 queries (LeetCode upper limits)
func BenchmarkXorQueries(b *testing.B) {
	const n = 30_000
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i & 255 // pseudo-random bytes
	}

	queries := make([][]int, n)
	for i := 0; i < n; i++ {
		l := i
		r := n - 1 - i
		if l > r {
			l, r = r, l
		}
		queries[i] = []int{l, r}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xorQueries(arr, queries)
	}
}
