package merge_k_sorted_lists_23

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{"empty input", [][]int{}, nil},
		{"single list", [][]int{{1, 4, 5}}, []int{1, 4, 5}},
		{"two lists", [][]int{{1, 4, 5}, {1, 3, 4}}, []int{1, 1, 3, 4, 4, 5}},
		{"three lists", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"with nil", [][]int{{}, {0}}, []int{0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var heads []*ListNode
			for _, nums := range tt.lists {
				heads = append(heads, fromSlice(nums))
			}
			got := toSlice(mergeKLists(heads))
			assert.Equal(t, tt.want, got)
		})
	}
}

// ---------- lightweight helpers ----------

func fromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	nodes := make([]ListNode, len(nums))
	for i, v := range nums {
		nodes[i].Val = v
		if i < len(nums)-1 {
			nodes[i].Next = &nodes[i+1]
		}
	}
	return &nodes[0]
}

func toSlice(head *ListNode) []int {
	var res []int
	for cur := head; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}
	return res
}
