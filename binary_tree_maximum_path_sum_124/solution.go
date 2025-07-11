// Package binary_tree_maximum_path_sum_124 provides a solution to find the maximum path sum in a binary tree.
// A path is a sequence of nodes connected by parent-child edges, containing at least one node,
// and does not need to pass through the root. The path sum is the sum of the node values in the path.
package binary_tree_maximum_path_sum_124

import "math"

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxPathSum computes the maximum path sum in a binary tree.
// It initializes a global maximum variable and then calls a recursive helper
// function to traverse the tree and update this maximum.
//
// The problem guarantees at least one node.
//
// Returns:
//
//	The maximum sum of node values along any valid path in the tree.
func maxPathSum(root *TreeNode) int {
	// Initialize maxResult to the smallest possible integer. This is crucial
	// because node values can be negative, and the maximum path might be a
	// single negative node if all other paths are even more negative.
	maxResult := math.MinInt
	// Start the recursive traversal. maxResult is passed by reference so that
	// the helper function can update the global maximum directly.
	maxPathSumHelper(root, &maxResult)
	return maxResult
}

// maxPathSumHelper recursively computes the maximum path sum that can extend upward
// from the current node. It also updates the global 'maxResult' with the maximum
// path sum found that may 'turn' at the current node (i.e., include both left and
// right subtrees through the current node).
//
// Args:
//
//	root: The current node in the binary tree being processed.
//	maxResult: A pointer to the global maximum path sum found so far. This is
//	           updated whenever a new maximum path sum (potentially passing
//	           through the current node and its children) is discovered.
//
// Returns:
//
//	The maximum path sum that starts at the current 'root' node and extends
//	downward through at most one of its children. This value is used by the
//	parent of 'root' to potentially extend a path further up the tree.
func maxPathSumHelper(root *TreeNode, maxResult *int) int {
	// Base case: If the current node is nil (e.g., past a leaf), it contributes
	// nothing to a path, so we return 0. This also effectively prunes non-existent
	// branches.
	if root == nil {
		return 0
	}

	// Recursively calculate the maximum path sum from the left child.
	// We use max(0, ...) because if the path sum from a child is negative,
	// we choose to ignore it (treat it as 0) as including it would only
	// decrease our current path sum. This ensures we only consider positive
	// contributions from subtrees when extending a path.
	leftGain := max(0, maxPathSumHelper(root.Left, maxResult))
	// Recursively calculate the maximum path sum from the right child, with
	// the same logic for ignoring negative contributions.
	rightGain := max(0, maxPathSumHelper(root.Right, maxResult))

	// Calculate a potential maximum path sum that 'turns' at the current node.
	// This path includes the current node's value and potentially the maximum
	// positive gains from both its left and right subtrees. This represents
	// a complete path that doesn't need to extend further up.
	currentPathSum := root.Val + leftGain + rightGain

	// Update the global maximum path sum if the 'currentPathSum' (which represents
	// a path that 'turns' at this node) is greater than the current maxResult.
	// This ensures maxResult always holds the highest sum found anywhere in the tree.
	*maxResult = max(*maxResult, currentPathSum)

	// Return the maximum path sum that can be extended upwards from the current node.
	// This path consists of the current node's value plus the maximum positive gain
	// from *either* its left child or its right child (but not both, as a path
	// can only go upwards through one branch). This value is what the parent node
	// will use to calculate its own potential path sums.
	return root.Val + max(leftGain, rightGain)
}
