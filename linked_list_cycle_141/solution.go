package linked_list_cycle_141

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n)
// Space: O(1)
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// Standard Floyd's cycle detection: optimized rabbit and turtle
	turtle := head
	rabbit := head
	
	for rabbit != nil && rabbit.Next != nil {
		turtle = turtle.Next
		rabbit = rabbit.Next.Next
		
		if turtle == rabbit {
			return true
		}
	}

	return false
}
