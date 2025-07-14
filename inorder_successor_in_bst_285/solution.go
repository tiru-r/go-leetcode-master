package inorder_successor_in_bst_285

// TreeNode represents a binary search tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderSuccessor finds the smallest node value > p.Val using optimized BST traversal
// Optimized: O(h) time, O(1) space with Go 1.24 modern syntax
func inorderSuccessor(root, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}

	var successor *TreeNode
	for curr := root; curr != nil; {
		if curr.Val > p.Val {
			// Found potential successor - search left for closer match
			successor = curr
			curr = curr.Left
		} else {
			// Current value ≤ target - successor must be in right subtree
			curr = curr.Right
		}
	}
	
	return successor
}
