package intersection_of_two_linked_lists_160

// ListNode represents a singly-linked list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode finds intersection using optimized two-pointer technique
// Optimized: O(m+n) time, O(1) space with Go 1.24 modern syntax
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pointerA, pointerB := headA, headB

	// Two-pointer "switch tracks" approach: each pointer traverses both lists
	// After switching, both pointers travel same distance (m+n)
	// They meet at intersection or both reach nil simultaneously
	for pointerA != pointerB {
		// Switch to other list when reaching end, creating equal path lengths
		if pointerA == nil {
			pointerA = headB
		} else {
			pointerA = pointerA.Next
		}

		if pointerB == nil {
			pointerB = headA
		} else {
			pointerB = pointerB.Next
		}
	}

	return pointerA // intersection node or nil if no intersection
}
