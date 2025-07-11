package binary_tree_level_order_traversal_102

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// LevelOrder performs a level-order (breadth-first) traversal of a binary tree.
// It returns a slice of slices, where each inner slice contains the values of nodes
// at a specific level, ordered from left to right.
// Example:
//
//	Input:  [3,9,20,null,null,15,7]
//	Output: [[3],[9,20],[15,7]]
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)        // Store the final result
	queue := make([]*TreeNode, 0, 10) // Initialize queue with pre-allocated capacity
	queue = append(queue, root)       // Start with the root node

	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevel := make([]int, 0, levelSize) // Pre-allocate capacity for current level

		// Process all nodes at the current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:] // Dequeue by re-slicing

			currentLevel = append(currentLevel, node.Val)

			// Enqueue non-nil children
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, currentLevel)
	}

	return result
}

// levelOrder0 performs a level-order traversal of a binary tree.
// Returns a slice of slices where each inner slice contains node values at a given level.
func levelOrder0(root *TreeNode) [][]int {
	// Base case: if the root is nil, return an empty slice of slices.
	if root == nil {
		return [][]int{}
	}

	// Initialize result to store all levels and currentLevel with the root node.
	var levels [][]int
	currentLevel := []*TreeNode{root}

	// Continue processing while there are nodes in the current level.
	for len(currentLevel) > 0 {
		// Initialize a slice for the next level's nodes and one for current level's values.
		var nextLevel []*TreeNode
		level := make([]int, 0, len(currentLevel)) // Pre-allocate capacity for efficiency.

		// Process each node in the current level.
		for _, node := range currentLevel {
			// Add the current node's value to the level's value slice.
			level = append(level, node.Val)

			// Add non-nil children to the next level.
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		// Add the current level's values to the result.
		levels = append(levels, level)
		// Move to the next level by updating currentLevel.
		currentLevel = nextLevel
	}

	// Return the collected levels.
	return levels
}
