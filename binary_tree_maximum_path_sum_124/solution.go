package binary_tree_maximum_path_sum_124

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxResult int

func maxPathSum(root *TreeNode) int {
	maxResult = root.Val // Initialize with root value instead of MinInt32
	maxPathSumHelper(root)
	return maxResult
}

func maxPathSumHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Including the left/right child in the current sum is not optimal
	// if either is negative, so take the max of the left/right sum and 0
	leftSum := max(maxPathSumHelper(root.Left), 0)
	rightSum := max(maxPathSumHelper(root.Right), 0)
	pathSum := rightSum + leftSum + root.Val
	maxResult = max(maxResult, pathSum)

	return root.Val + max(rightSum, leftSum)
}
