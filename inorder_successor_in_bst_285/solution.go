package inorder_successor_in_bst_285

// TreeNode is the standard BST node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderSuccessor returns the node with the smallest key strictly greater than p.Val.
// Works in O(h) time and O(1) extra space.
func inorderSuccessor(root, p *TreeNode) *TreeNode {
	if root == nil || p == nil {
		return nil
	}

	var succ *TreeNode
	for cur := root; cur != nil; {
		if cur.Val > p.Val {
			// cur is a candidate; move left to find smaller candidates
			succ = cur
			cur = cur.Left
		} else {
			// cur <= p.Val, successor must be in right subtree
			cur = cur.Right
		}
	}
	return succ
}
