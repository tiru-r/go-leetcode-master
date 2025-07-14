package linked_list_cycle_141

// ListNode represents a singly linked list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle detects cycle using Floyd's algorithm (two pointers)
// Optimized: O(n) time, O(1) space with Go 1.24 modern syntax
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	
	// Floyd's cycle detection: tortoise and hare
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
