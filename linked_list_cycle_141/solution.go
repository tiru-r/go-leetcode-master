package linked_list_cycle_141

type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle detects a cycle using Floyd’s Tortoise & Hare.
// Time: O(n)   Space: O(1)
func hasCycle(head *ListNode) bool {
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
