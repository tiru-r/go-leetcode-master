package subtree_of_another_tree_572

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return equalTrees(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func equalTrees(t1, t2 *TreeNode) bool {
	if t1 == nil || t2 == nil {
		return t1 == t2
	}
	return t1.Val == t2.Val && equalTrees(t1.Left, t2.Left) && equalTrees(t1.Right, t2.Right)
}
