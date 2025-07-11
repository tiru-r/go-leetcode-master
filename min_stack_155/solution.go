package min_stack_155

import "slices"

type MinStack struct {
	stack []int
	mins  []int
}

// Modern constructor using slices.Grow for better performance
func Constructor() MinStack {
	// Pre-allocate capacity for better performance
	stack := slices.Grow([]int(nil), 16)
	mins := slices.Grow([]int(nil), 16)
	return MinStack{
		stack: stack,
		mins:  mins,
	}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)
	if len(ms.mins) == 0 || x <= ms.mins[len(ms.mins)-1] {
		ms.mins = append(ms.mins, x)
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) == 0 {
		return
	}

	// Use more efficient slice operations
	val := ms.stack[len(ms.stack)-1]
	ms.stack = ms.stack[:len(ms.stack)-1]

	if len(ms.mins) > 0 && val == ms.mins[len(ms.mins)-1] {
		ms.mins = ms.mins[:len(ms.mins)-1]
	}
}

func (ms *MinStack) Top() int {
	if len(ms.stack) == 0 {
		return 0 // or panic("empty stack")
	}
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	if len(ms.mins) == 0 {
		return 0 // or panic("empty stack")
	}
	return ms.mins[len(ms.mins)-1]
}
