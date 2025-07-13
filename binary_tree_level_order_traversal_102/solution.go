package binary_tree_level_order_traversal_102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := make([][]int, 0, 8)

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)
		
		for range levelSize {
			node := queue[0]
			queue = queue[1:]
			
			level = append(level, node.Val)
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		
		result = append(result, level)
	}

	return result
}
