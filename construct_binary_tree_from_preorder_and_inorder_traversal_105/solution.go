package construct_binary_tree_from_preorder_and_inorder_traversal_105

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTree constructs a binary tree from preorder and inorder traversal slices.
func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// Build a map of inorder values to their indices for O(1) lookups.
	inorderMap := make(map[int]int, len(inorder))
	for i, val := range inorder {
		inorderMap[val] = i
	}

	// Helper function to avoid passing map repeatedly.
	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd || inStart > inEnd {
			return nil
		}

		// Create root from preorder's first element.
		root := &TreeNode{Val: preorder[preStart]}

		// Find root's index in inorder.
		rootIdx, exists := inorderMap[root.Val]
		if !exists {
			return nil // Invalid input.
		}

		// Calculate size of left subtree.
		leftSize := rootIdx - inStart

		// Recursively build left and right subtrees.
		root.Left = build(preStart+1, preStart+leftSize, inStart, rootIdx-1)
		root.Right = build(preStart+leftSize+1, preEnd, rootIdx+1, inEnd)

		return root
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
