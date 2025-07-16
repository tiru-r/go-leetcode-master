package reverse_linked_list_206

// ListNode is a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseList reverses the list in-place and returns the new head.
// Time: O(n), Space: O(1).
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for curr := head; curr != nil; curr, prev, curr.Next = curr.Next, curr, prev {
	}
	return prev
}
