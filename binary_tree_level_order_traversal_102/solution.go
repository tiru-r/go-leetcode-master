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
func LevelOrder(root *TreeNode) [][]int {
	// Handle the edge case of an empty tree.
	if root == nil {
		return [][]int{}
	}

	// Initialize the queue for BFS.
	// We use a slice as a queue, and pre-allocate with a reasonable capacity
	// to reduce initial reallocations.
	queue := make([]*TreeNode, 0, 16)
	queue = append(queue, root) // Start with the root node in the queue

	// Slice to store the final result, where each inner slice represents a level.
	result := [][]int{}

	// 'head' acts as a pointer to the front of our logical queue within the slice.
	// Elements from index '0' to 'head-1' are considered "dequeued".
	// Elements from 'head' to 'len(queue)-1' are "enqueued".
	head := 0

	// Continue as long as there are nodes in the logical queue to process.
	for head < len(queue) {
		// 'levelSize' captures the number of nodes currently in the queue
		// that belong to the *current* level being processed.
		// This is crucial because new children for the *next* level will be
		// added to 'queue' during this loop iteration, but we only want to
		// process nodes that were present at the start of this level.
		levelSize := len(queue) - head

		// Create a slice to store the values for the current level.
		// Pre-allocate its capacity for efficiency.
		currentLevel := make([]int, 0, levelSize)

		// Process all nodes that belong to the current level.
		for i := 0; i < levelSize; i++ {
			// Get the node at the front of the logical queue.
			node := queue[head]
			head++ // Move the 'head' pointer forward, effectively "dequeuing" the node.

			// Add the node's value to the current level's result.
			currentLevel = append(currentLevel, node.Val)

			// Enqueue the left child if it exists.
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// Enqueue the right child if it exists.
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// After processing all nodes for the current level, add this level's
		// values to the final result.
		result = append(result, currentLevel)

		// Memory reuse optimization:
		// If 'head' has caught up to 'len(queue)', it means all nodes that were
		// initially in the queue at the start of this level's processing have been
		// fully processed. At this point, 'queue' only contains newly added
		// children for subsequent levels.
		//
		// By resetting 'queue = queue[:0]', we reset the slice's length to 0
		// while retaining its underlying array and capacity.
		// Then, by setting 'head = 0', we effectively "compact" the queue:
		// all nodes that were added for the next level(s) are now logically
		// at the beginning of the slice, ready to be processed in the next iteration.
		// This avoids continually growing the underlying array unnecessarily and
		// reuses allocated memory.
		if head == len(queue) {
			queue = queue[:0]
			head = 0
		}
	}

	return result
}
