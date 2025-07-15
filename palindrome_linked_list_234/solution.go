package palindrome_linked_list_234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	mid := findMiddle(head)
	secondHalf := reverse(mid)
	defer func() { reverse(secondHalf) }()

	return compareLists(head, secondHalf)
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func compareLists(first, second *ListNode) bool {
	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	return true
}
