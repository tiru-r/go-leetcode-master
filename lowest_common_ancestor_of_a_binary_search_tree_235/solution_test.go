package lowest_common_ancestor_of_a_binary_search_tree_235

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* ---------- helpers only for tests ---------- */

// build creates a BST from the given insertion order.
func build(vals ...int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	for _, v := range vals[1:] {
		insert(root, v)
	}
	return root
}

// insert puts v into the BST rooted at root.
func insert(root *TreeNode, v int) {
	for {
		if v < root.Val {
			if root.Left == nil {
				root.Left = &TreeNode{Val: v}
				return
			}
			root = root.Left
		} else {
			if root.Right == nil {
				root.Right = &TreeNode{Val: v}
				return
			}
			root = root.Right
		}
	}
}

// find returns the node whose Val == v (assumes v exists).
func find(root *TreeNode, v int) *TreeNode {
	for root != nil && root.Val != v {
		if v < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

/* ---------- test table ---------- */

func Test_lowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name string
		vals []int // BST insertion order
		p, q int
		want int
	}{
		{"example 1", []int{6, 2, 8, 0, 4, 7, 9, 3, 5}, 2, 8, 6},
		{"example 2", []int{6, 2, 8, 0, 4, 7, 9, 3, 5}, 2, 4, 2},
		{"parent is LCA", []int{6, 8, 10, 9, 12}, 8, 9, 8},
		{"root is LCA", []int{2, 1, 3}, 1, 3, 2},
		{"single node", []int{1}, 1, 1, 1},
		{"left chain", []int{5, 3, 1}, 3, 1, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := build(tt.vals...)
			p := find(root, tt.p)
			q := find(root, tt.q)
			assert.Equal(t, tt.want, lowestCommonAncestor(root, p, q).Val)
		})
	}
}

func BenchmarkLowestCommonAncestor(b *testing.B) {
	root := build(6, 2, 8, 0, 4, 7, 9, 3, 5)
	p := find(root, 2)
	q := find(root, 8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lowestCommonAncestor(root, p, q)
	}
}

func BenchmarkLowestCommonAncestorDeep(b *testing.B) {
	// Create a deep BST
	vals := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		vals[i] = i
	}
	root := build(vals...)
	p := find(root, 100)
	q := find(root, 900)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lowestCommonAncestor(root, p, q)
	}
}
