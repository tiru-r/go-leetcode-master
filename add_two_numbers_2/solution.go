package add_two_numbers_2

// ListNode defines a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers adds two numbers represented as linked lists.
// The digits are stored in reverse order, and each node contains a single digit.
// The function returns the sum as a new linked list.
//
// Constraints:
// - The number of nodes in each linked list is between 1 and 100.
// - 0 <= Node.val <= 9
// - It is guaranteed that the list represents a number that does not have leading zeros, except for the number 0 itself.
//
// Example:
// l1 = [2,4,3] (represents 342)
// l2 = [5,6,4] (represents 465)
// sum = [7,0,8] (represents 807)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Constraints guarantee l1 and l2 are non-nil (at least 1 node),
	// so no explicit nil checks are required. However, the loop condition
	// (l1 != nil || l2 != nil || carry > 0) naturally handles:
	// - Lists of different lengths (e.g., [1,2] + [5,6,7,8])
	// - Carry propagation after lists are exhausted (e.g., [9,9] + [1] = [0,0,1])
	// - Single-digit lists (e.g., [0] + [0] = [0])

	// Initialize dummy head for the result list to simplify construction.
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	// Loop until both lists are exhausted and no carry remains.
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		// Add l1's digit if available, advance l1.
		if l1 != nil {
			sum += l1.Val // Node.val is guaranteed to be 0-9 per constraints.
			l1 = l1.Next
		}

		// Add l2's digit if available, advance l2.
		if l2 != nil {
			sum += l2.Val // Node.val is guaranteed to be 0-9 per constraints.
			l2 = l2.Next
		}

		// Update carry and create new node with current digit.
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	// Return the result list, starting after dummy head.
	// No leading zeros are possible (except [0] case, which is handled correctly).
	return dummyHead.Next
}

// createList creates a linked list from an array of integers.
// This is a helper for testing purposes, not part of the core solution for LeetCode.
func createList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil // Not needed for LeetCode constraints, but included for robustness.
	}
	dummy := &ListNode{}
	current := dummy
	for _, val := range arr {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return dummy.Next
}

// listToArray converts a linked list to an array for easy comparison in tests.
// This is a helper for testing purposes, not part of the core solution for LeetCode.
func listToArray(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
