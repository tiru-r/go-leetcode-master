package kth_smallest_element_in_a_bst_230

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// kthSmallest returns the k-th (1-based) smallest element in the BST.
// Runs in O(h + k) time and O(h) space (stack height h).
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) > 0 {
		// push left spine
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return curr.Val
		}

		curr = curr.Right
	}
	return -1 // unreachable if 1 <= k <= #nodes
}
