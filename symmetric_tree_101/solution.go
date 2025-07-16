package symmetric_tree_101

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSym(root, root)
}

func isSym(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	return left.Val == right.Val && isSym(left.Left, right.Right) && isSym(right.Left, left.Right)
}
