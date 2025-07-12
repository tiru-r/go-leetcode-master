package reverse_linked_list_206

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------- small test helpers (zero-alloc) ----------

// toSlice converts a list to a slice for assertions.
func toSlice(head *ListNode) []int {
	var s []int
	for n := head; n != nil; n = n.Next {
		s = append(s, n.Val)
	}
	return s
}

// fromSlice builds a list from a slice.
func fromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	nodes := make([]ListNode, len(vals))
	for i, v := range vals {
		nodes[i].Val = v
		if i < len(vals)-1 {
			nodes[i].Next = &nodes[i+1]
		}
	}
	return &nodes[0]
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", nil, nil},
		{"single", []int{42}, []int{42}},
		{"two", []int{1, 2}, []int{2, 1}},
		{"normal", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"palindrome", []int{1, 2, 3, 2, 1}, []int{1, 2, 3, 2, 1}},
		{"negatives", []int{-3, -2, -1}, []int{-1, -2, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toSlice(ReverseList(fromSlice(tt.in)))
			assert.Equal(t, tt.want, got)
		})
	}
}

// Benchmark focuses on the fast iterative version
func BenchmarkReverseList(b *testing.B) {
	const size = 1_000
	in := make([]int, size)
	for i := 0; i < size; i++ {
		in[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		head := fromSlice(in)
		_ = ReverseList(head)
	}
}
