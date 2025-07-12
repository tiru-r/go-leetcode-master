package intersection_of_two_linked_lists_160

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getIntersectionNode(t *testing.T) {
	// 1. Example from LeetCode – intersection at node 8.
	intersect := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersect}}
	headB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: intersect}}}
	assert.Same(t, intersect, getIntersectionNode(headA, headB))

	// 2. No intersection.
	a2 := &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	b2 := &ListNode{Val: 1, Next: &ListNode{Val: 5}}
	assert.Nil(t, getIntersectionNode(a2, b2))

	// 3. Intersection at the very first node (both lists identical).
	single := &ListNode{Val: 42}
	assert.Same(t, single, getIntersectionNode(single, single))

	// 4. One list is nil.
	assert.Nil(t, getIntersectionNode(nil, headA))
	assert.Nil(t, getIntersectionNode(headB, nil))

	// 5. Different lengths – intersection at tail.
	tail := &ListNode{Val: 7}
	short := &ListNode{Val: 1, Next: tail}
	long := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: tail}}}
	assert.Same(t, tail, getIntersectionNode(short, long))
}
