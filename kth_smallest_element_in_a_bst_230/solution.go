package kth_smallest_element_in_a_bst_230

// TreeNode represents a binary search tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthSmallest finds k-th smallest element using optimized inorder traversal
// Optimized: O(h+k) time, O(h) space with Go 1.24 modern syntax and early termination
func kthSmallest(root *TreeNode, k int) int {
	count := 0
	var result int
	
	// Recursive inorder with early termination
	var inorder func(*TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		
		// Traverse left subtree
		if inorder(node.Left) {
			return true
		}
		
		// Process current node
		count++
		if count == k {
			result = node.Val
			return true // early termination
		}
		
		// Traverse right subtree
		return inorder(node.Right)
	}
	
	inorder(root)
	return result
}
