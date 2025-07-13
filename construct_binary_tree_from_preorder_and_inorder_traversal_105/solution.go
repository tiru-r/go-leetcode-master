package construct_binary_tree_from_preorder_and_inorder_traversal_105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	
	indexMap := make(map[int]int, len(inorder))
	for i, val := range inorder {
		indexMap[val] = i
	}
	
	preIdx := 0
	
	var build func(start, end int) *TreeNode
	build = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}
		
		rootVal := preorder[preIdx]
		preIdx++
		
		root := &TreeNode{Val: rootVal}
		rootIdx := indexMap[rootVal]
		
		root.Left = build(start, rootIdx-1)
		root.Right = build(rootIdx+1, end)
		
		return root
	}
	
	return build(0, len(inorder)-1)
}
