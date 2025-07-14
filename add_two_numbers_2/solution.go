package add_two_numbers_2

// ListNode defines a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers adds two numbers represented as linked lists
// Time: O(max(m,n)), Space: O(max(m,n))
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create a dummy node to simplify result list construction
	dummy := &ListNode{}
	// Keep a pointer to current position in result list
	current := dummy
	// Initialize carry to 0 for digit addition overflow
	carry := 0

	// Continue while there are digits in either list or carry exists
	for l1 != nil || l2 != nil || carry > 0 {
		// Start with carry from previous digit addition
		sum := carry

		// Add digit from first list if it exists
		if l1 != nil {
			sum += l1.Val        // Add current digit value to sum
			l1 = l1.Next         // Move to next digit in first list
		}

		// Add digit from second list if it exists
		if l2 != nil {
			sum += l2.Val        // Add current digit value to sum
			l2 = l2.Next         // Move to next digit in second list
		}

		// Calculate carry for next digit (1 if sum >= 10, 0 otherwise)
		carry = sum / 10
		// Create new node with current digit (sum modulo 10)
		current.Next = &ListNode{Val: sum % 10}
		// Move current pointer to the newly created node
		current = current.Next
	}

	// Return the actual result list (skip dummy node)
	return dummy.Next
}

func createList(arr []int) *ListNode {
	// Return nil for empty array
	if len(arr) == 0 {
		return nil
	}
	// Create dummy node to simplify list construction
	dummy := &ListNode{}
	// Keep pointer to current position in list
	current := dummy
	// Iterate through array values
	for _, val := range arr {
		// Create new node with current value
		current.Next = &ListNode{Val: val}
		// Move to newly created node
		current = current.Next
	}
	// Return actual list (skip dummy node)
	return dummy.Next
}

func listToArray(head *ListNode) []int {
	// Initialize empty slice to store list values
	var result []int
	// Traverse the linked list
	for head != nil {
		// Add current node's value to result array
		result = append(result, head.Val)
		// Move to next node
		head = head.Next
	}
	// Return the array containing all list values
	return result
}
