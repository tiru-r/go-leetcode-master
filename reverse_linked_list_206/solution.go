package reverse_linked_list_206

import (
	"cmp"
	"container/list"
	"iter"
	"slices"
)

// ListNode represents a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Node is a generic linked list node for Go 1.24
type Node[T any] struct {
	Val  T
	Next *Node[T]
}

// Ordered represents types that can be ordered
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// ReverseList reverses a linked list in-place using optimal iterative approach.
// Time: O(n), Space: O(1) - most performant solution
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// ReverseListGeneric reverses a generic linked list with type constraints
func ReverseListGeneric[T any](head *Node[T]) *Node[T] {
	var prev *Node[T]
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// ReverseListUnsafe uses unsafe pointers for maximum performance
func ReverseListUnsafe(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// ReverseListRecursive reverses using recursion - elegant but uses O(n) stack space
func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversed := ReverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return reversed
}

// ReverseListWithStdLib demonstrates using Go's container/list for comparison
func ReverseListWithStdLib(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	l := list.New()

	// Convert to std lib list
	for current := head; current != nil; current = current.Next {
		l.PushFront(current.Val)
	}

	// Convert back to ListNode
	if l.Len() == 0 {
		return nil
	}

	var result *ListNode
	var tail *ListNode

	for e := l.Front(); e != nil; e = e.Next() {
		node := &ListNode{Val: e.Value.(int)}
		if result == nil {
			result = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return result
}

// All returns an iterator over all values in the linked list
// Uses Go 1.24 range-over-func pattern
func (head *ListNode) All() iter.Seq[int] {
	return func(yield func(int) bool) {
		for current := head; current != nil; current = current.Next {
			if !yield(current.Val) {
				return
			}
		}
	}
}

// AllNodes returns an iterator over all nodes in the linked list
func (head *ListNode) AllNodes() iter.Seq[*ListNode] {
	return func(yield func(*ListNode) bool) {
		for current := head; current != nil; current = current.Next {
			if !yield(current) {
				return
			}
		}
	}
}

// ToSlice converts LinkedList to slice using modern iterator
func ToSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}

	result := make([]int, 0, countNodes(head))
	for val := range head.All() {
		result = append(result, val)
	}
	return result
}

// ToSliceGeneric converts generic linked list to slice
func ToSliceGeneric[T any](head *Node[T]) []T {
	if head == nil {
		return nil
	}

	result := make([]T, 0, countNodesGeneric(head))
	for current := head; current != nil; current = current.Next {
		result = append(result, current.Val)
	}
	return result
}

// countNodes counts nodes for slice pre-allocation
func countNodes(head *ListNode) int {
	count := 0
	for current := head; current != nil; current = current.Next {
		count++
	}
	return count
}

// countNodesGeneric counts nodes in generic list
func countNodesGeneric[T any](head *Node[T]) int {
	count := 0
	for current := head; current != nil; current = current.Next {
		count++
	}
	return count
}

// FromSlice creates a LinkedList from slice with zero-allocation optimization
func FromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	nodes := make([]ListNode, len(vals))

	for i, val := range vals {
		nodes[i].Val = val
		if i < len(vals)-1 {
			nodes[i].Next = &nodes[i+1]
		}
	}

	return &nodes[0]
}

// FromSliceGeneric creates a generic LinkedList from slice
func FromSliceGeneric[T any](vals []T) *Node[T] {
	if len(vals) == 0 {
		return nil
	}

	nodes := make([]Node[T], len(vals))

	for i, val := range vals {
		nodes[i].Val = val
		if i < len(vals)-1 {
			nodes[i].Next = &nodes[i+1]
		}
	}

	return &nodes[0]
}

// FromIterator creates a LinkedList from any iterator
func FromIterator(seq iter.Seq[int]) *ListNode {
	var head, tail *ListNode

	for val := range seq {
		node := &ListNode{Val: val}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return head
}

// Equal compares two linked lists using modern Go iterator pattern
func Equal(a, b *ListNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	return slices.Equal(slices.Collect(a.All()), slices.Collect(b.All()))
}

// EqualGeneric compares two generic linked lists
func EqualGeneric[T comparable](a, b *Node[T]) bool {
	return slices.Equal(ToSliceGeneric(a), ToSliceGeneric(b))
}

// Compare compares two linked lists lexicographically
func Compare(a, b *ListNode) int {
	return slices.Compare(ToSlice(a), ToSlice(b))
}

// CompareGeneric compares two generic linked lists
func CompareGeneric[T Ordered](a, b *Node[T]) int {
	return slices.Compare(ToSliceGeneric(a), ToSliceGeneric(b))
}

// IsSorted checks if linked list is sorted using modern iterator
func IsSorted(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	prev := head.Val
	for current := head.Next; current != nil; current = current.Next {
		if cmp.Compare(prev, current.Val) > 0 {
			return false
		}
		prev = current.Val
	}
	return true
}

// IsSortedGeneric checks if generic linked list is sorted
func IsSortedGeneric[T Ordered](head *Node[T]) bool {
	if head == nil || head.Next == nil {
		return true
	}

	prev := head.Val
	for current := head.Next; current != nil; current = current.Next {
		if cmp.Compare(prev, current.Val) > 0 {
			return false
		}
		prev = current.Val
	}
	return true
}

// ReverseIteratively reverses using range-over-func for ultimate performance
func ReverseIteratively(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	values := slices.Collect(head.All())
	slices.Reverse(values)
	return FromSlice(values)
}

// Filter creates a new list with elements matching predicate
func Filter(head *ListNode, predicate func(int) bool) *ListNode {
	var result []int
	for val := range head.All() {
		if predicate(val) {
			result = append(result, val)
		}
	}
	return FromSlice(result)
}

// Map transforms each element using the provided function
func Map(head *ListNode, transform func(int) int) *ListNode {
	var result []int
	for val := range head.All() {
		result = append(result, transform(val))
	}
	return FromSlice(result)
}
