package min_stack_155

import "slices"

type MinStack struct {
	stack []int
	mins  []int
}

// Modern constructor using slices.Grow for better performance
func Constructor() MinStack {
	stack := make([]int, 0)
	mins := make([]int, 0)
	// Pre-allocate some capacity for better performance
	slices.Grow(stack, 16)
	slices.Grow(mins, 16)
	return MinStack{
		stack: stack,
		mins:  mins,
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.mins) == 0 || x <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	
	// Use more efficient slice operations
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	
	if len(this.mins) > 0 && val == this.mins[len(this.mins)-1] {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
