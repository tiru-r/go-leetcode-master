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

	result := [][]int{}               // Store the final result
	queue := make([]*TreeNode, 0, 10) // Initialize queue with pre-allocated capacity
	queue = append(queue, root)       // Start with the root node

	head := 0
	for head < len(queue) {
		levelSize := len(queue) - head
		currentLevel := make([]int, 0, levelSize) // Pre-allocate capacity for current level

		// Process all nodes at the current level
		for i := 0; i < levelSize; i++ {
			node := queue[head]
			head++

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
