package reorder_list_143

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	cur := head
	var end *ListNode
	for end != cur {
		end = head

		// move e until one before the end
		for end.Next != nil && end.Next.Next != nil {
			end = end.Next
		}

		if end != cur {
			end.Next.Next = cur.Next
			cur.Next = end.Next
			end.Next = nil
			cur = cur.Next.Next
		}
	}
}

// Modern slice-based solution with optimal performance
func reorderListModern(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Collect nodes in slice for efficient random access
	nodes := make([]*ListNode, 0, 64) // Pre-allocate reasonable capacity
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	// Reorder using slice indices for O(1) access
	left, right := 0, len(nodes)-1
	current := head
	
	for left < right {
		// Take from left
		if current != nodes[left] {
			current.Next = nodes[left]
			current = current.Next
		}
		left++
		
		// Take from right if still valid
		if left <= right {
			current.Next = nodes[right]
			current = current.Next
			right--
		}
	}
	
	// Handle middle element for odd length
	if left == right {
		current.Next = nodes[left]
		current = current.Next
	}
	
	// Terminate the list
	current.Next = nil
}
