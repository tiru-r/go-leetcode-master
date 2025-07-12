package linked_list_cycle_141

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {
	t.Run("cycle at tail", func(t *testing.T) {
		four := &ListNode{Val: -4}
		two := &ListNode{
			Val:  2,
			Next: &ListNode{Val: 0, Next: four},
		}
		four.Next = two
		head := &ListNode{Val: 3, Next: two}
		assert.True(t, hasCycle(head))
	})

	t.Run("two-node cycle", func(t *testing.T) {
		one := &ListNode{Val: 1}
		two := &ListNode{Val: 2, Next: one}
		one.Next = two
		assert.True(t, hasCycle(one))
	})

	t.Run("single node, no cycle", func(t *testing.T) {
		assert.False(t, hasCycle(&ListNode{Val: 1}))
	})

	t.Run("empty list", func(t *testing.T) {
		assert.False(t, hasCycle(nil))
	})
}
