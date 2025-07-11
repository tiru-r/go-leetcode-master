package delete_node_in_a_linked_list_237

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	c := node
	for c.Next.Next != nil {
		c.Val = c.Next.Val
		c = c.Next
	}

	c.Val = c.Next.Val
	c.Next = nil
}
