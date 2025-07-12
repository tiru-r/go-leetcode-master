package intersection_of_two_linked_lists_160

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB

	// After both pointers switch lists once, they are equidistant from the tail.
	// Either they meet at the intersection node or hit nil together.
	for a != b {
		// switch lists when reaching the end
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a // nil if no intersection
}
