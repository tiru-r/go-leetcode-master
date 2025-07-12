package invert_binary_tree_226

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------- helper: deep equality of two binary trees ----------
func treeEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return treeEqual(a.Left, b.Left) && treeEqual(a.Right, b.Right)
}

// ----------------------------------------------------------------
func Test_invertTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			name: "nil tree",
			root: nil,
			want: nil,
		},
		{
			name: "single node",
			root: &TreeNode{Val: 1},
			want: &TreeNode{Val: 1},
		},
		{
			name: "left skewed",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			want: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
		},
		{
			name: "right skewed",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			want: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
		},
		{
			name: "small complete",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			want: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
		},
		{
			name: "LeetCode example",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			want: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := invertTree(tt.root)
			assert.True(t, treeEqual(got, tt.want), "tree mismatch")
		})
	}
}
