package delete_node_in_a_linked_list_237

type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteNode removes the given node (not tail) in O(1) time.
func deleteNode(node *ListNode) {
	// node is guaranteed non-tail and non-nil by LeetCode.
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
