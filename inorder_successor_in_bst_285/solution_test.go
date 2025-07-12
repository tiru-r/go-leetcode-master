package inorder_successor_in_bst_285

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// helper: build a BST from slice (level-order for brevity)
func buildBST(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	q := []*TreeNode{root}
	i := 1
	for len(q) > 0 && i < len(vals) {
		curr := q[0]
		q = q[1:]
		if vals[i] != -1 { // -1 means nil
			curr.Left = &TreeNode{Val: vals[i]}
			q = append(q, curr.Left)
		}
		i++
		if i < len(vals) && vals[i] != -1 {
			curr.Right = &TreeNode{Val: vals[i]}
			q = append(q, curr.Right)
		}
		i++
	}
	return root
}

// helper: pointer to node with given value (first occurrence)
func find(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		}
		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}

func TestInorderSuccessor(t *testing.T) {
	tests := []struct {
		name    string
		tree    []int // level-order, -1 = nil
		pVal    int
		wantVal int
		wantNil bool
	}{
		{"LC 2-1-3", []int{2, 1, 3}, 1, 2, false},
		{"LC 5-3-6-2-4-nil-nil-1", []int{5, 3, 6, 2, 4, -1, -1, 1}, 6, 0, true}, // successor nil
		{"root only", []int{1}, 1, 0, true},
		{"right skew", []int{1, -1, 2, -1, -1, -1, 3}, 2, 3, false},
		{"left skew", []int{3, 2, -1, 1}, 1, 2, false},
		{"missing successor", []int{5, 3, 7}, 7, 0, true},
		{"duplicate values (first)", []int{2, 2, 3}, 2, 3, false},
		{"full tree", []int{4, 2, 6, 1, 3, 5, 7}, 3, 4, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := buildBST(tc.tree)
			p := find(root, tc.pVal)
			if p == nil {
				t.Fatalf("node %d not found in tree", tc.pVal)
			}
			got := inorderSuccessor(root, p)
			if tc.wantNil {
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
				assert.Equal(t, tc.wantVal, got.Val)
			}
		})
	}
}
