package delete_node_in_a_linked_list_237

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	*node = *node.Next
}
