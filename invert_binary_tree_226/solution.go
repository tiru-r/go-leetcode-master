package invert_binary_tree_226

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree inverts binary tree using optimized recursive approach
// Optimized: O(n) time, O(h) space with Go 1.24 modern syntax
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	// Swap children using Go's multiple assignment
	root.Left, root.Right = root.Right, root.Left
	
	// Recursively invert subtrees
	invertTree(root.Left)
	invertTree(root.Right)
	
	return root
}
