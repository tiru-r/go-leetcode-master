package construct_binary_tree_from_preorder_and_inorder_traversal_105

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test_buildTree tests the buildTree function with various input scenarios.
func Test_buildTree(t *testing.T) {
	// Define a struct for test cases.
	type args struct {
		preorder []int
		inorder  []int
	}

	// Helper function to create a TreeNode for concise test case definitions.
	newNode := func(val int, left, right *TreeNode) *TreeNode {
		return &TreeNode{Val: val, Left: left, Right: right}
	}

	// Define test cases.
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "standard case with left and right subtrees",
			args: args{
				preorder: []int{3, 9, 20, 15, 7},
				inorder:  []int{9, 3, 15, 20, 7},
			},
			want: newNode(3,
				newNode(9, nil, nil),
				newNode(20,
					newNode(15, nil, nil),
					newNode(7, nil, nil),
				),
			),
		},
		{
			name: "empty tree",
			args: args{
				preorder: []int{},
				inorder:  []int{},
			},
			want: nil,
		},
		{
			name: "single node",
			args: args{
				preorder: []int{1},
				inorder:  []int{1},
			},
			want: newNode(1, nil, nil),
		},
		{
			name: "left-only subtree",
			args: args{
				preorder: []int{1, 2, 3},
				inorder:  []int{3, 2, 1},
			},
			want: newNode(1,
				newNode(2,
					newNode(3, nil, nil),
					nil,
				),
				nil,
			),
		},
		{
			name: "right-only subtree",
			args: args{
				preorder: []int{1, 2, 3},
				inorder:  []int{1, 2, 3},
			},
			want: newNode(1,
				nil,
				newNode(2,
					nil,
					newNode(3, nil, nil),
				),
			),
		},
		{
			name: "invalid input (mismatched lengths)",
			args: args{
				preorder: []int{1, 2},
				inorder:  []int{1},
			},
			want: nil,
		},
	}

	// Run test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTree(tt.args.preorder, tt.args.inorder)
			assert.Equal(t, tt.want, got, "Tree structures should match")
		})
	}
}
