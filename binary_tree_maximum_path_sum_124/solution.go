package binary_tree_maximum_path_sum_124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		
		leftGain := max(0, dfs(node.Left))
		rightGain := max(0, dfs(node.Right))
		
		pathSum := node.Val + leftGain + rightGain
		maxSum = max(maxSum, pathSum)
		
		return node.Val + max(leftGain, rightGain)
	}
	
	dfs(root)
	return maxSum
}
