package merge_two_sorted_lists_21

// ListNode definition using Go 1.24 - no external dependencies needed
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// Note: Cool technique to not have to worry about front of list being null
	//       current.Next is always where the next will go
	holdFront := &ListNode{}
	current := holdFront

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}

		current = current.Next
	}

	// Append remaining nodes (one of l1 or l2 is nil)
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return holdFront.Next
}
