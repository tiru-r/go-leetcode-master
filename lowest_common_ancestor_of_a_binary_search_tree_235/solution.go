package lowest_common_ancestor_of_a_binary_search_tree_235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Ensure p has smaller value for optimization
	if p.Val > q.Val {
		p, q = q, p
	}
	
	// Iterative solution with single comparison optimization
	for root != nil {
		rootVal := root.Val
		
		// Both nodes are in left subtree
		if q.Val < rootVal {
			root = root.Left
		} else if p.Val > rootVal {
			// Both nodes are in right subtree
			root = root.Right
		} else {
			// Split point found: p.Val <= rootVal <= q.Val
			return root
		}
	}
	
	return nil
}
